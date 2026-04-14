package qualitymonitor

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

	assert.Equal(t, int64(0), s.Version, "schema version should be 0 (bidirectional migration with SDKv2)")

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

	// Verify provider_config is a SingleNestedAttribute (types.Object)
	pcAttr, ok := s.Attributes["provider_config"]
	require.True(t, ok, "provider_config attribute must exist")
	pcSNA, ok := pcAttr.(schema.SingleNestedAttribute)
	require.True(t, ok, "provider_config must be a SingleNestedAttribute")
	assert.True(t, pcSNA.Optional, "provider_config should be optional")
	assert.True(t, pcSNA.Computed, "provider_config should be computed")
	assert.Len(t, pcSNA.PlanModifiers, 1, "provider_config should have ProviderConfigPlanModifier")

	// Verify workspace_id inside provider_config
	wsAttr, ok := pcSNA.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be string")
	assert.True(t, wsStr.Optional, "workspace_id should be optional")
	assert.True(t, wsStr.Computed, "workspace_id should be computed")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 2, "workspace_id should have 2 validators")
}

func TestResourceQualityMonitor_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &QualityMonitorResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}
