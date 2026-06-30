package library

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceLibrary_ImplementsResourceWithModifyPlan(t *testing.T) {
	r := ResourceLibrary()
	_, ok := r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "LibraryResource must implement ResourceWithModifyPlan")
}

func TestResourceLibrary_SchemaPreserved(t *testing.T) {
	r := ResourceLibrary()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	// Verify cluster_id exists and is required
	clusterAttr, ok := s.Attributes["cluster_id"]
	require.True(t, ok, "cluster_id attribute must exist")
	strAttr, ok := clusterAttr.(schema.StringAttribute)
	require.True(t, ok, "cluster_id must be a string attribute")
	assert.True(t, strAttr.Required, "cluster_id should be required")
	assert.Len(t, strAttr.PlanModifiers, 1, "cluster_id should have RequiresReplace plan modifier")

	// Verify all string attributes have RequiresReplace
	for _, field := range []string{"jar", "whl", "requirements"} {
		attr, ok := s.Attributes[field]
		if !ok {
			continue // some may be nested differently
		}
		strA, ok := attr.(schema.StringAttribute)
		if ok {
			assert.Len(t, strA.PlanModifiers, 1, "%s should have RequiresReplace plan modifier", field)
		}
	}

	// Verify id is computed and optional
	idAttr, ok := s.Attributes["id"]
	require.True(t, ok, "id attribute must exist")
	idStr, ok := idAttr.(schema.StringAttribute)
	require.True(t, ok, "id must be a string attribute")
	assert.True(t, idStr.Computed, "id should be computed")
	assert.True(t, idStr.Optional, "id should be optional")

	// Verify SDKv2-block compat: ConfigureAsSdkV2Compatible() must keep these
	// nested types as blocks (not attributes) so that terraform.tfstate files
	// written by SDKv2-era releases of databricks_library still decode.
	for _, name := range []string{"cran", "maven", "pypi"} {
		_, isBlock := s.Blocks[name]
		_, isAttr := s.Attributes[name]
		assert.True(t, isBlock, "%q must be a block (SDKv2-state compat)", name)
		assert.False(t, isAttr, "%q must NOT be an attribute", name)
	}

	// Verify provider_config block exists (SdkV2 compatible = list nested block)
	pcBlock, ok := s.Blocks["provider_config"]
	require.True(t, ok, "provider_config block must exist")
	pcList, ok := pcBlock.(schema.ListNestedBlock)
	require.True(t, ok, "provider_config must be a list nested block (SdkV2 compatible)")
	assert.Len(t, pcList.Validators, 1, "provider_config should have SizeAtMost(1) validator")

	// Verify workspace_id inside provider_config
	wsAttr, ok := pcList.NestedObject.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be a string attribute")
	assert.True(t, wsStr.Required, "workspace_id should be required")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 1, "workspace_id should have LengthAtLeast(1) validator only")
}

func TestResourceLibrary_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &LibraryResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}

func TestReadLibrary_NotFoundDuringRefresh(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=test-cluster",
			ReuseRequest: true,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: "test-cluster",
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Status: "INSTALLED",
						Library: &compute.Library{
							Jar: "other.jar",
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)

		// During refresh (Read), a missing library should return nil with a warning, not an error.
		// This allows Terraform to detect drift and plan re-creation.
		result, diags := readLibrary(ctx, w, compute.Wait{
			ClusterID: "test-cluster",
			IsRefresh: true,
		}, "mvn:net.snowflake:spark-snowflake_2.12:3.1.0")

		assert.Nil(t, result)
		assert.False(t, diags.HasError(), "should not return error during refresh when library is not found")
		assert.True(t, len(diags.Warnings()) > 0, "should return a warning when library is not found")
		assert.Contains(t, diags.Warnings()[0].Detail(), "not found on cluster")
	})
}

func TestReadLibrary_NotFoundDuringCreate(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=test-cluster",
			ReuseRequest: true,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: "test-cluster",
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Status: "INSTALLED",
						Library: &compute.Library{
							Jar: "other.jar",
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)

		// During create (IsRunning=true), a missing library should return an error.
		result, diags := readLibrary(ctx, w, compute.Wait{
			ClusterID: "test-cluster",
			IsRunning: true,
		}, "mvn:net.snowflake:spark-snowflake_2.12:3.1.0")

		assert.Nil(t, result)
		assert.True(t, diags.HasError(), "should return error during create when library is not found")
		assert.Contains(t, diags.Errors()[0].Detail(), "failed to find")
	})
}

func TestReadLibrary_FoundDuringRefresh(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=test-cluster",
			ReuseRequest: true,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: "test-cluster",
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Status: "INSTALLED",
						Library: &compute.Library{
							Maven: &compute.MavenLibrary{
								Coordinates: "net.snowflake:spark-snowflake_2.12:3.1.0",
							},
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)

		result, diags := readLibrary(ctx, w, compute.Wait{
			ClusterID: "test-cluster",
			IsRefresh: true,
		}, "mvn:net.snowflake:spark-snowflake_2.12:3.1.0")

		assert.False(t, diags.HasError())
		assert.NotNil(t, result)
		assert.Equal(t, "test-cluster", result.ClusterId.ValueString())
	})
}
