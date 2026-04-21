package sharing

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceShare_ImplementsResourceWithModifyPlan(t *testing.T) {
	r := ResourceShare()
	_, ok := r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "ShareResource must implement ResourceWithModifyPlan")
}

func TestResourceShare_SchemaPreserved(t *testing.T) {
	r := ResourceShare()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	assert.Equal(t, int64(0), s.Version, "schema version should be 0 (bidirectional migration with SDKv2)")

	// Verify name is required with RequiresReplace
	nameAttr, ok := s.Attributes["name"]
	require.True(t, ok, "name attribute must exist")
	strAttr, ok := nameAttr.(schema.StringAttribute)
	require.True(t, ok, "name must be a string attribute")
	assert.True(t, strAttr.Required, "name should be required")
	assert.Len(t, strAttr.PlanModifiers, 1, "name should have RequiresReplace plan modifier")

	// Verify computed fields with UseStateForUnknown
	createdAtAttr, ok := s.Attributes["created_at"]
	require.True(t, ok, "created_at must exist")
	int64Attr, ok := createdAtAttr.(schema.Int64Attribute)
	require.True(t, ok, "created_at must be int64")
	assert.Len(t, int64Attr.PlanModifiers, 1, "created_at should have UseStateForUnknown")

	createdByAttr, ok := s.Attributes["created_by"]
	require.True(t, ok, "created_by must exist")
	cbStr, ok := createdByAttr.(schema.StringAttribute)
	require.True(t, ok, "created_by must be string")
	assert.Len(t, cbStr.PlanModifiers, 1, "created_by should have UseStateForUnknown")

	// Verify id is computed
	idAttr, ok := s.Attributes["id"]
	require.True(t, ok, "id must exist")
	idStr, ok := idAttr.(schema.StringAttribute)
	require.True(t, ok, "id must be string")
	assert.True(t, idStr.Computed, "id should be computed")

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

func TestResourceShare_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &ShareResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}
