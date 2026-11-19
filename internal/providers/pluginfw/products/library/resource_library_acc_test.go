package library

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/qa"
)

const testClusterID = "test-cluster"

func librarySchema(t *testing.T, ctx context.Context) schema.Schema {
	t.Helper()
	resp := &resource.SchemaResponse{}
	ResourceLibrary().Schema(ctx, resource.SchemaRequest{}, resp)
	require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)
	return resp.Schema
}

func libraryState(t *testing.T, ctx context.Context, library compute.Library) LibraryExtended {
	t.Helper()
	var model LibraryExtended
	diags := converters.GoSdkToTfSdkStruct(ctx, library, &model)
	require.False(t, diags.HasError(), diags)
	model.ClusterId = types.StringValue(testClusterID)
	model.ID = types.StringValue(library.String())
	model.ProviderConfig = types.ListNull(tfschema.ProviderConfig{}.Type(ctx))
	return model
}

func libraryPlan(t *testing.T, ctx context.Context, model LibraryExtended) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: librarySchema(t, ctx)}
	diags := plan.Set(ctx, &model)
	require.False(t, diags.HasError(), diags)
	return plan
}

func TestLibraryCreate(t *testing.T) {
	ctx := context.Background()
	pypi := compute.Library{Pypi: &compute.PythonPyPiLibrary{Package: "databricks-sdk"}}
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=" + testClusterID,
			Response: compute.ClusterDetails{ClusterId: testClusterID, State: compute.StateRunning},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.0/libraries/install",
			ExpectedRequest: compute.InstallLibraries{ClusterId: testClusterID, Libraries: []compute.Library{pypi}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + testClusterID,
			Response: compute.ClusterLibraryStatuses{ClusterId: testClusterID, LibraryStatuses: []compute.LibraryFullStatus{
				{Library: &pypi, Status: compute.LibraryInstallStatusInstalled},
			}},
		},
	}, func(_ context.Context, client *common.DatabricksClient) {
		s := librarySchema(t, ctx)
		resp := &resource.CreateResponse{State: tfsdk.State{Schema: s}}
		(&LibraryResource{Client: client}).Create(ctx, resource.CreateRequest{
			Plan: libraryPlan(t, ctx, libraryState(t, ctx, pypi)),
		}, resp)
		require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)

		var state LibraryExtended
		diags := resp.State.Get(ctx, &state)
		require.False(t, diags.HasError(), diags)
		assert.Equal(t, testClusterID, state.ClusterId.ValueString())
		assert.Equal(t, pypi.String(), state.ID.ValueString())
	})
}

func TestLibraryCreateStartsTerminatedCluster(t *testing.T) {
	ctx := context.Background()
	pypi := compute.Library{Pypi: &compute.PythonPyPiLibrary{Package: "networkx"}}
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=" + testClusterID,
			Response: compute.ClusterDetails{ClusterId: testClusterID, State: compute.StateTerminated},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.1/clusters/start",
			ExpectedRequest: compute.StartCluster{ClusterId: testClusterID},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=" + testClusterID,
			Response: compute.ClusterDetails{ClusterId: testClusterID, State: compute.StateRunning},
		},
		{
			Method:          "POST",
			Resource:        "/api/2.0/libraries/install",
			ExpectedRequest: compute.InstallLibraries{ClusterId: testClusterID, Libraries: []compute.Library{pypi}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + testClusterID,
			Response: compute.ClusterLibraryStatuses{ClusterId: testClusterID, LibraryStatuses: []compute.LibraryFullStatus{
				{Library: &pypi, Status: compute.LibraryInstallStatusInstalled},
			}},
		},
	}, func(_ context.Context, client *common.DatabricksClient) {
		resp := &resource.CreateResponse{State: tfsdk.State{Schema: librarySchema(t, ctx)}}
		(&LibraryResource{Client: client}).Create(ctx, resource.CreateRequest{
			Plan: libraryPlan(t, ctx, libraryState(t, ctx, pypi)),
		}, resp)
		require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)
	})
}

func TestLibraryDeleteSkipsMissingCluster(t *testing.T) {
	ctx := context.Background()
	pypi := compute.Library{Pypi: &compute.PythonPyPiLibrary{Package: "networkx"}}
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/get?cluster_id=" + testClusterID,
			Status:   404,
		},
	}, func(_ context.Context, client *common.DatabricksClient) {
		state := tfsdk.State{Schema: librarySchema(t, ctx)}
		diags := state.Set(ctx, libraryState(t, ctx, pypi))
		require.False(t, diags.HasError(), diags)
		resp := &resource.DeleteResponse{}
		(&LibraryResource{Client: client}).Delete(ctx, resource.DeleteRequest{State: state}, resp)
		assert.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)
		require.Len(t, resp.Diagnostics.Warnings(), 1)
		assert.Contains(t, resp.Diagnostics.Warnings()[0].Detail(), "skipping library uninstallation")
	})
}

func TestLibraryUpdateUsesReplacementAndPreservesSdkV2StateShape(t *testing.T) {
	ctx := context.Background()
	s := librarySchema(t, ctx)
	for _, name := range []string{"cran", "maven", "pypi"} {
		block, isBlock := s.Blocks[name]
		_, isAttribute := s.Attributes[name]
		assert.True(t, isBlock, "%q must remain an SDKv2-compatible block", name)
		assert.False(t, isAttribute, "%q must not become an attribute", name)

		listBlock, ok := block.(schema.ListNestedBlock)
		require.True(t, ok, "%q must remain a list block", name)
		assert.NotEmpty(t, listBlock.PlanModifiers, "%q changes must replace the library", name)
	}
}
