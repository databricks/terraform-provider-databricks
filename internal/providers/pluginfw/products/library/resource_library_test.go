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

	// Verify schema version is 1 (we introduced a v0→v1 upgrader).
	assert.Equal(t, int64(1), s.Version, "schema version should be 1")

	// Verify provider_config is a SingleNestedAttribute (types.Object). The v0
	// ListNestedBlock shape is gone in v1 — see UpgradeState for the migration.
	pcAttr, ok := s.Attributes["provider_config"]
	require.True(t, ok, "provider_config attribute must exist")
	pcObj, ok := pcAttr.(schema.SingleNestedAttribute)
	require.True(t, ok, "provider_config must be a SingleNestedAttribute")
	assert.True(t, pcObj.Optional, "provider_config should be Optional")
	assert.True(t, pcObj.Computed, "provider_config should be Computed")
	assert.Len(t, pcObj.PlanModifiers, 1, "provider_config should have ProviderConfigPlanModifier")

	// Verify workspace_id inside provider_config is Optional+Computed with a plan modifier and length validator.
	wsAttr, ok := pcObj.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be a string attribute")
	assert.True(t, wsStr.Optional, "workspace_id should be Optional")
	assert.True(t, wsStr.Computed, "workspace_id should be Computed")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 1, "workspace_id should have LengthAtLeast(1) validator only")
}

func TestResourceLibrary_UpgradeStateV0ToV1(t *testing.T) {
	r := ResourceLibrary().(*LibraryResource)
	upgraders := r.UpgradeState(context.Background())
	require.Contains(t, upgraders, int64(0), "must register a v0 upgrader")
	v0 := upgraders[0]
	require.NotNil(t, v0.PriorSchema, "v0 upgrader must declare PriorSchema")
	require.NotNil(t, v0.StateUpgrader, "v0 upgrader must declare a StateUpgrader func")
	// v0 PriorSchema must have provider_config as a ListNestedBlock so that
	// state files written when v0 was the active schema can decode against it.
	pcBlock, isBlock := v0.PriorSchema.Blocks["provider_config"]
	require.True(t, isBlock, "v0 PriorSchema must have provider_config as a block")
	_, isList := pcBlock.(schema.ListNestedBlock)
	require.True(t, isList, "v0 provider_config must be a ListNestedBlock")
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
