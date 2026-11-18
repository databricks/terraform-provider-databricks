package qualitymonitor

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceQualityMonitor_ImplementsResourceWithModifyPlan(t *testing.T) {
	r := ResourceQualityMonitor()
	_, ok := r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "QualityMonitorResource must implement ResourceWithModifyPlan")
}

func TestResourceQualityMonitor_SchemaPreserved(t *testing.T) {
	r := ResourceQualityMonitor()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	// Verify key required attributes
	assetsAttr, ok := s.Attributes["assets_dir"]
	require.True(t, ok, "assets_dir attribute must exist")
	strAttr, ok := assetsAttr.(schema.StringAttribute)
	require.True(t, ok, "assets_dir must be a string attribute")
	assert.True(t, strAttr.Required, "assets_dir should be required")

	// Verify read-only computed fields
	for _, field := range []string{"monitor_version", "drift_metrics_table_name", "profile_metrics_table_name", "status", "dashboard_id"} {
		attr, ok := s.Attributes[field]
		require.True(t, ok, "%s attribute must exist", field)
		strA, ok := attr.(schema.StringAttribute)
		if ok {
			assert.True(t, strA.Computed, "%s should be computed", field)
		}
	}

	// Verify optional fields
	whAttr, ok := s.Attributes["warehouse_id"]
	require.True(t, ok, "warehouse_id must exist")
	whStr, ok := whAttr.(schema.StringAttribute)
	require.True(t, ok, "warehouse_id must be string")
	assert.True(t, whStr.Optional, "warehouse_id should be optional")

	// Verify id
	idAttr, ok := s.Attributes["id"]
	require.True(t, ok, "id must exist")
	idStr, ok := idAttr.(schema.StringAttribute)
	require.True(t, ok, "id must be string")
	assert.True(t, idStr.Computed, "id should be computed")
	assert.True(t, idStr.Optional, "id should be optional")

	// Verify provider_config block exists (SdkV2 compatible)
	pcBlock, ok := s.Blocks["provider_config"]
	require.True(t, ok, "provider_config block must exist")
	pcList, ok := pcBlock.(schema.ListNestedBlock)
	require.True(t, ok, "provider_config must be a list nested block (SdkV2 compatible)")
	assert.Len(t, pcList.Validators, 1, "provider_config should have SizeAtMost(1) validator")

	// Verify workspace_id inside provider_config
	wsAttr, ok := pcList.NestedObject.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be string")
	assert.True(t, wsStr.Required, "workspace_id should be required")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 1, "workspace_id should have LengthAtLeast(1) validator only")
}

func TestResourceQualityMonitor_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &QualityMonitorResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}

func TestResourceQualityMonitor_ModifyPlan_UnknownWorkspaceID(t *testing.T) {
	ctx := context.Background()
	pcType := tfschema.ProviderConfig{}.Type(ctx)
	tests := []struct {
		name           string
		providerConfig types.List
		wantErr        string
	}{
		{
			name: "unknown workspace_id",
			providerConfig: types.ListValueMust(pcType, []attr.Value{
				tfschema.ProviderConfig{WorkspaceID: types.StringUnknown()}.ToObjectValue(ctx)}),
		},
		{name: "unknown provider_config", providerConfig: types.ListUnknown(pcType)},
		{name: "no provider_config", providerConfig: types.ListNull(pcType), wantErr: "none was found"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// An account-level provider without workspace_id fails a workspace client
			// lookup without a request.
			r := &QualityMonitorResource{Client: &common.DatabricksClient{
				DatabricksClient: &client.DatabricksClient{Config: &config.Config{
					Host:      "https://accounts.cloud.databricks.com",
					AccountID: "test-account-id",
					Token:     "test-token",
				}},
			}}
			schemaResp := &resource.SchemaResponse{}
			r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
			plan := tfsdk.Plan{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
			}
			require.False(t, plan.SetAttribute(ctx, path.Root("provider_config"), tc.providerConfig).HasError())
			resp := &resource.ModifyPlanResponse{Plan: plan}

			r.ModifyPlan(ctx, resource.ModifyPlanRequest{Plan: plan}, resp)

			if tc.wantErr == "" {
				assert.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				return
			}
			require.True(t, resp.Diagnostics.HasError())
			assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), tc.wantErr)
		})
	}
}
