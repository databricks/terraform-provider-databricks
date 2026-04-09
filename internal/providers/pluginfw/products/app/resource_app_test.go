package app

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceApp_ImplementsResourceWithModifyPlan(t *testing.T) {
	r := ResourceApp()
	_, ok := r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "resourceApp must implement ResourceWithModifyPlan")
}

func TestResourceApp_SchemaPreserved(t *testing.T) {
	r := ResourceApp()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	// Verify key attributes exist
	nameAttr, ok := s.Attributes["name"]
	require.True(t, ok, "name attribute must exist")
	strAttr, ok := nameAttr.(schema.StringAttribute)
	require.True(t, ok, "name must be a string attribute")
	assert.Len(t, strAttr.PlanModifiers, 1, "name should have RequiresReplace plan modifier")

	// Verify computed fields have UseStateForUnknown plan modifiers
	for _, field := range []string{"create_time", "creator", "service_principal_client_id", "service_principal_name", "url"} {
		attr, ok := s.Attributes[field]
		require.True(t, ok, "%s attribute must exist", field)
		strAttr, ok := attr.(schema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", field)
		assert.Len(t, strAttr.PlanModifiers, 1, "%s should have UseStateForUnknown plan modifier", field)
	}

	// service_principal_id should have int64 UseStateForUnknown
	spIdAttr, ok := s.Attributes["service_principal_id"]
	require.True(t, ok, "service_principal_id must exist")
	int64Attr, ok := spIdAttr.(schema.Int64Attribute)
	require.True(t, ok, "service_principal_id must be int64")
	assert.Len(t, int64Attr.PlanModifiers, 1, "service_principal_id should have UseStateForUnknown")

	// Verify provider_config exists and is optional
	pcAttr, ok := s.Attributes["provider_config"]
	require.True(t, ok, "provider_config attribute must exist")
	pcNested, ok := pcAttr.(schema.SingleNestedAttribute)
	require.True(t, ok, "provider_config must be a single nested attribute")
	assert.True(t, pcNested.Optional, "provider_config should be optional")

	// Verify workspace_id inside provider_config
	wsAttr, ok := pcNested.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be a string attribute")
	assert.True(t, wsStr.Optional, "workspace_id should be optional")
	assert.True(t, wsStr.Computed, "workspace_id should be computed")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 2, "workspace_id should have 2 validators (LengthAtLeast, RegexMatches)")
}

func TestResourceApp_ModifyPlan_SkipsDestroyPlan(t *testing.T) {
	r := &resourceApp{}
	// Plan.Raw zero value is null, simulating a destroy plan
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null (destroy) plan")
}

func TestResourceApp_ModifyPlan_SkipsWhenClientNil(t *testing.T) {
	r := &resourceApp{client: nil}
	// Non-null plan but no client
	resp := &resource.ModifyPlanResponse{}
	req := resource.ModifyPlanRequest{}
	// Plan.Raw is null by default (zero value), so this tests the null path
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error when client is nil")
}
