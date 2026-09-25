package genie_space

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostUpdateState_TrustsPlanForUserControlledFieldsWhenGetIsStale(t *testing.T) {
	plan := GenieSpace{
		SpaceId:         types.StringValue("space-123"),
		Title:           types.StringValue("new-title"),
		Description:     types.StringValue("new-desc"),
		WarehouseId:     types.StringValue("new-warehouse"),
		ParentPath:      types.StringValue("/Users/me/folder"),
		SerializedSpace: types.StringValue(`{"version":1}`),
	}
	// updated == response from UpdateSpace; reflects the just-applied write.
	updated := &dashboards.GenieSpace{
		SpaceId: "space-123",
		Etag:    "etag-after-update",
	}
	// fetched == post-update GET; simulates a stale cache that still has the
	// OLD title/description/warehouse_id. SerializedSpace is the API-canonical
	// (possibly re-serialized) form.
	fetched := &dashboards.GenieSpace{
		SpaceId:         "space-123",
		Title:           "old-title",
		Description:     "old-desc",
		WarehouseId:     "old-warehouse",
		SerializedSpace: `{ "version": 1 }`, // semantically equal to plan
		Etag:            "etag-from-get",
	}

	got := postUpdateState(plan, updated, fetched)

	assert.Equal(t, "new-title", got.Title.ValueString(), "title must come from plan, not stale GET")
	assert.Equal(t, "new-desc", got.Description.ValueString(), "description must come from plan, not stale GET")
	assert.Equal(t, "new-warehouse", got.WarehouseId.ValueString(), "warehouse_id must come from plan, not stale GET")
	assert.Equal(t, "/Users/me/folder", got.ParentPath.ValueString(), "parent_path must come from plan (not returned by API)")
	assert.Equal(t, "space-123", got.SpaceId.ValueString())
	assert.Equal(t, "etag-from-get", got.Etag.ValueString(), "etag must come from GET (API-managed)")
	assert.Equal(t, `{ "version": 1 }`, got.SerializedSpace.ValueString(), "serialized_space comes from GET; normalization is applied later by Sync")
}

func TestResourceGenieSpace_ImplementsRequiredInterfaces(t *testing.T) {
	r := ResourceGenieSpace()
	_, ok := r.(resource.ResourceWithConfigure)
	assert.True(t, ok, "ResourceGenieSpace must implement ResourceWithConfigure")
	_, ok = r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "ResourceGenieSpace must implement ResourceWithModifyPlan")
	_, ok = r.(resource.ResourceWithImportState)
	assert.True(t, ok, "ResourceGenieSpace must implement ResourceWithImportState")
}

func TestResourceGenieSpace_Metadata(t *testing.T) {
	r := ResourceGenieSpace()
	resp := &resource.MetadataResponse{}
	r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "databricks"}, resp)
	assert.Equal(t, "databricks_genie_space", resp.TypeName)
}

func TestResourceGenieSpace_Schema(t *testing.T) {
	r := ResourceGenieSpace()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	requiredFields := []string{"title", "warehouse_id", "serialized_space", "parent_path"}
	for _, name := range requiredFields {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", name)
		assert.True(t, attr.Required, "%s should be required", name)
	}

	for _, name := range []string{"description"} {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", name)
		assert.True(t, attr.Optional, "%s should be optional", name)
	}

	for _, name := range []string{"space_id", "etag"} {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", name)
		assert.True(t, attr.Computed, "%s should be computed", name)
		assert.Len(t, attr.PlanModifiers, 1, "%s should have UseStateForUnknown plan modifier", name)
	}

	parentPath, ok := s.Attributes["parent_path"].(schema.StringAttribute)
	require.True(t, ok)
	assert.Len(t, parentPath.PlanModifiers, 1, "parent_path should have RequiresReplace plan modifier")

	serialized, ok := s.Attributes["serialized_space"].(schema.StringAttribute)
	require.True(t, ok)
	assert.Len(t, serialized.PlanModifiers, 1, "serialized_space should have JSON-semantic-equality plan modifier")

	pc, ok := s.Attributes["provider_config"].(schema.SingleNestedAttribute)
	require.True(t, ok, "provider_config must be a single nested attribute")
	assert.True(t, pc.Optional, "provider_config should be optional")
	assert.True(t, pc.Computed, "provider_config should be computed")
	wsID, ok := pc.Attributes["workspace_id"].(schema.StringAttribute)
	require.True(t, ok, "workspace_id must exist in provider_config")
	assert.True(t, wsID.Optional)
	assert.True(t, wsID.Computed)
}

func TestResourceGenieSpace_ImportState_Valid(t *testing.T) {
	ctx := context.Background()
	r := ResourceGenieSpace()
	importer, ok := r.(resource.ResourceWithImportState)
	require.True(t, ok)
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

	resp := &resource.ImportStateResponse{
		State: tfsdk.State{
			Schema: schemaResp.Schema,
			Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
		},
	}
	importer.ImportState(ctx, resource.ImportStateRequest{ID: "abc-123"}, resp)
	require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)

	var id types.String
	diags := resp.State.GetAttribute(ctx, path.Root("space_id"), &id)
	require.False(t, diags.HasError(), "diagnostics: %v", diags)
	assert.Equal(t, "abc-123", id.ValueString())
}

func TestResourceGenieSpace_ImportState_Invalid(t *testing.T) {
	r := ResourceGenieSpace()
	importer := r.(resource.ResourceWithImportState)
	resp := &resource.ImportStateResponse{}
	importer.ImportState(context.Background(), resource.ImportStateRequest{ID: ""}, resp)
	assert.True(t, resp.Diagnostics.HasError(), "empty ID must produce an error")

	resp = &resource.ImportStateResponse{}
	importer.ImportState(context.Background(), resource.ImportStateRequest{ID: "a,b"}, resp)
	assert.True(t, resp.Diagnostics.HasError(), "multi-part ID must produce an error")
}

func TestResourceGenieSpace_ModifyPlan_NilClient(t *testing.T) {
	r := &GenieSpaceResource{Client: nil}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), resource.ModifyPlanRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError(), "ModifyPlan must be safe to call with nil client (e.g. during destroy)")
}

func TestNormalizeJSON(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"empty", "", "", false},
		{"object", `{"b":1,"a":2}`, `{"a":2,"b":1}`, false},
		{"nested", `{"b":1,"a":{"y":2,"x":1}}`, `{"a":{"x":1,"y":2},"b":1}`, false},
		{"whitespace", "{\n  \"a\": 1\n}", `{"a":1}`, false},
		{"array", `[3,1,2]`, `[3,1,2]`, false},
		{"invalid", `{not json`, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := normalizeJSON(tc.in)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestJSONSemanticallyEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want bool
	}{
		{"identical", `{"a":1}`, `{"a":1}`, true},
		{"whitespace_only", `{"a":1}`, "{\n  \"a\": 1\n}", true},
		{"key_order", `{"a":1,"b":2}`, `{"b":2,"a":1}`, true},
		{"nested_key_order", `{"o":{"y":2,"x":1}}`, `{"o":{"x":1,"y":2}}`, true},
		{"different_value", `{"a":1}`, `{"a":2}`, false},
		{"one_invalid_json", `{"a":1}`, `not json`, false},
		{"both_empty", "", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, jsonSemanticallyEqual(tc.a, tc.b))
		})
	}
}

func TestJSONSerializedSpacePlanModifier_PreservesStateWhenEquivalent(t *testing.T) {
	m := jsonSerializedSpacePlanModifier{}
	stateValue := `{"a":1,"b":2}`
	planValue := `{"b":2,"a":1}`

	req := planmodifier.StringRequest{
		StateValue: types.StringValue(stateValue),
		PlanValue:  types.StringValue(planValue),
	}
	resp := &planmodifier.StringResponse{PlanValue: req.PlanValue}
	m.PlanModifyString(context.Background(), req, resp)

	assert.Equal(t, stateValue, resp.PlanValue.ValueString(),
		"plan should snap back to state value when JSON is semantically equal")
}

func TestJSONSerializedSpacePlanModifier_KeepsPlanWhenDifferent(t *testing.T) {
	m := jsonSerializedSpacePlanModifier{}
	stateValue := `{"a":1}`
	planValue := `{"a":2}`

	req := planmodifier.StringRequest{
		StateValue: types.StringValue(stateValue),
		PlanValue:  types.StringValue(planValue),
	}
	resp := &planmodifier.StringResponse{PlanValue: req.PlanValue}
	m.PlanModifyString(context.Background(), req, resp)

	assert.Equal(t, planValue, resp.PlanValue.ValueString(),
		"plan value should be kept when JSON values are not semantically equal")
}

func TestJSONSerializedSpacePlanModifier_SkipsNullOrUnknown(t *testing.T) {
	m := jsonSerializedSpacePlanModifier{}
	planValue := `{"a":1}`

	req := planmodifier.StringRequest{
		StateValue: types.StringNull(),
		PlanValue:  types.StringValue(planValue),
	}
	resp := &planmodifier.StringResponse{PlanValue: req.PlanValue}
	m.PlanModifyString(context.Background(), req, resp)
	assert.Equal(t, planValue, resp.PlanValue.ValueString())

	req = planmodifier.StringRequest{
		StateValue: types.StringValue(planValue),
		PlanValue:  types.StringUnknown(),
	}
	resp = &planmodifier.StringResponse{PlanValue: req.PlanValue}
	m.PlanModifyString(context.Background(), req, resp)
	assert.True(t, resp.PlanValue.IsUnknown(), "unknown plan must stay unknown")
}

func TestNewStateFromGenieSpace_PreservesParentPath(t *testing.T) {
	apiResponse := &dashboards.GenieSpace{
		SpaceId:         "id-1",
		Title:           "My Space",
		WarehouseId:     "wh-1",
		SerializedSpace: `{"a":1}`,
		Description:     "desc",
		Etag:            "etag-1",
	}
	preserved := types.StringValue("/Users/me/spaces")
	state := newStateFromGenieSpace(apiResponse, preserved)

	assert.Equal(t, "id-1", state.SpaceId.ValueString())
	assert.Equal(t, "My Space", state.Title.ValueString())
	assert.Equal(t, "wh-1", state.WarehouseId.ValueString())
	assert.Equal(t, `{"a":1}`, state.SerializedSpace.ValueString())
	assert.Equal(t, "desc", state.Description.ValueString())
	assert.Equal(t, "/Users/me/spaces", state.ParentPath.ValueString())
	assert.Equal(t, "etag-1", state.Etag.ValueString())
}

func TestOptionalString(t *testing.T) {
	assert.True(t, optionalString("").IsNull(), "empty input must become null")
	assert.Equal(t, "value", optionalString("value").ValueString())
}
