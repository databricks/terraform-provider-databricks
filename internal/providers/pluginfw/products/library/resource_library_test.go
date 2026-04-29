package library

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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

	// Schema version was bumped to 1 to migrate provider_config from list
	// (v1.113.0 ListNestedBlock) to object (SingleNestedAttribute).
	// See https://github.com/databricks/terraform-provider-databricks/issues/5669.
	assert.Equal(t, int64(1), s.Version, "schema version should be 1 after the provider_config list→object migration")

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
	require.True(t, ok, "workspace_id must be a string attribute")
	assert.True(t, wsStr.Optional, "workspace_id should be optional")
	assert.True(t, wsStr.Computed, "workspace_id should be computed")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 2, "workspace_id should have 2 validators")
}

func TestResourceLibrary_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &LibraryResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}

// v1.113.0-style state JSON: provider_config is a list (because the schema
// used Namespace_SdkV2 → ListNestedBlock). issue #5669 reports that this
// state can no longer be decoded by v1.114.0+, which switched the schema to
// types.Object → SingleNestedAttribute.
const v1113PriorStateJSON = `{
  "cluster_id": "0214-145447-tnvgr2qz",
  "id": "pypi:lxml~=6.0.0",
  "pypi": [{"package": "lxml~=6.0.0", "repo": ""}],
  "provider_config": []
}`

const v1113PriorStateWithWorkspaceIDJSON = `{
  "cluster_id": "0214-145447-tnvgr2qz",
  "id": "pypi:lxml~=6.0.0",
  "pypi": [{"package": "lxml~=6.0.0", "repo": ""}],
  "provider_config": [{"workspace_id": "12345"}]
}`

// TestResourceLibrary_PassthroughDecodeFailsForV1_113_0State demonstrates
// the regression behind issue #5669: when no UpgradeState is registered, the
// framework's UpgradeResourceState falls back to a passthrough decode using
// the current schema. v1.113.0 wrote provider_config as a list, but the
// current schema expects an object — the cty decoder fails on the type tag.
func TestResourceLibrary_PassthroughDecodeFailsForV1_113_0State(t *testing.T) {
	ctx := context.Background()

	r := ResourceLibrary()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	schemaType := schemaResp.Schema.Type().TerraformType(ctx)

	for _, tc := range []struct {
		name string
		json string
	}{
		{"empty_provider_config_list", v1113PriorStateJSON},
		{"populated_provider_config_list", v1113PriorStateWithWorkspaceIDJSON},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rawState := &tfprotov6.RawState{JSON: []byte(tc.json)}

			_, err := rawState.UnmarshalWithOpts(schemaType, tfprotov6.UnmarshalOpts{
				ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
					IgnoreUndefinedAttributes: true,
				},
			})

			require.Error(t, err,
				"v1.113.0 state with list-shaped provider_config must fail to decode against the v1.114.0+ object schema")
			assert.Contains(t, err.Error(), "provider_config",
				"error must point at provider_config")
			t.Logf("decode error users see: %v", err)
		})
	}
}

// TestResourceLibrary_UpgradeStateConvertsListToObject exercises the v0→v1
// state upgrader: a v1.113.0 state with provider_config as a list must
// upgrade cleanly to the v1.114.x+ object shape. Empty list → null; a list
// with a single object → object.
func TestResourceLibrary_UpgradeStateConvertsListToObject(t *testing.T) {
	ctx := context.Background()

	r, ok := ResourceLibrary().(resource.ResourceWithUpgradeState)
	require.True(t, ok, "LibraryResource must implement ResourceWithUpgradeState")

	upgraders := r.UpgradeState(ctx)
	upgrader, ok := upgraders[0]
	require.True(t, ok, "library must register an upgrader for prior schema version 0")

	schemaResp := &resource.SchemaResponse{}
	r.(resource.Resource).Schema(ctx, resource.SchemaRequest{}, schemaResp)

	for _, tc := range []struct {
		name              string
		json              string
		expectWorkspaceID string // empty → expect provider_config = null
	}{
		{"empty_list_becomes_null", v1113PriorStateJSON, ""},
		{"list_with_object_becomes_object", v1113PriorStateWithWorkspaceIDJSON, "12345"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: []byte(tc.json)}}
			resp := resource.UpgradeStateResponse{
				State: tfsdk.State{Schema: schemaResp.Schema},
			}

			upgrader.StateUpgrader(ctx, req, &resp)

			require.False(t, resp.Diagnostics.HasError(),
				"upgrader returned diagnostics: %s", resp.Diagnostics.Errors())

			// The upgraded state must use the current schema's terraform type
			// (object-shaped provider_config). If we got here, decode worked.
			require.NotNil(t, resp.State.Raw.Type(),
				"upgrader must populate State.Raw")

			// Pull provider_config out of the upgraded raw value.
			pcVal, err := resp.State.Raw.ApplyTerraform5AttributePathStep(
				tftypes.AttributeName("provider_config"))
			require.NoError(t, err)
			pcTfVal, ok := pcVal.(tftypes.Value)
			require.True(t, ok, "provider_config must be a tftypes.Value")

			if tc.expectWorkspaceID == "" {
				assert.True(t, pcTfVal.IsNull(),
					"empty list must upgrade to null object, got: %v", pcTfVal)
				return
			}

			require.False(t, pcTfVal.IsNull(),
				"populated list must upgrade to a non-null object")
			wsVal, err := pcTfVal.ApplyTerraform5AttributePathStep(
				tftypes.AttributeName("workspace_id"))
			require.NoError(t, err)
			wsTfVal, ok := wsVal.(tftypes.Value)
			require.True(t, ok)
			var got string
			require.NoError(t, wsTfVal.As(&got))
			assert.Equal(t, tc.expectWorkspaceID, got)
		})
	}
}
