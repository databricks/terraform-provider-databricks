package qualitymonitor

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// v1.113.0-style state JSON: provider_config is a list (because the schema
// used Namespace_SdkV2 → ListNestedBlock). Issue #5669.
const v1113QMStateEmptyPC = `{
  "table_name": "main.default.users",
  "assets_dir": "/Users/me/qm",
  "output_schema_name": "main.default",
  "provider_config": []
}`

const v1113QMStateWithWorkspaceID = `{
  "table_name": "main.default.users",
  "assets_dir": "/Users/me/qm",
  "output_schema_name": "main.default",
  "provider_config": [{"workspace_id": "12345"}]
}`

func TestResourceQualityMonitor_PassthroughDecodeFailsForV1_113_0State(t *testing.T) {
	ctx := context.Background()

	r := ResourceQualityMonitor()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	schemaType := schemaResp.Schema.Type().TerraformType(ctx)

	for _, tc := range []struct {
		name string
		json string
	}{
		{"empty_provider_config_list", v1113QMStateEmptyPC},
		{"populated_provider_config_list", v1113QMStateWithWorkspaceID},
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
			assert.Contains(t, err.Error(), "provider_config")
			t.Logf("decode error users see: %v", err)
		})
	}
}

func TestResourceQualityMonitor_UpgradeStateConvertsListToObject(t *testing.T) {
	ctx := context.Background()

	r, ok := ResourceQualityMonitor().(resource.ResourceWithUpgradeState)
	require.True(t, ok, "QualityMonitorResource must implement ResourceWithUpgradeState")

	upgraders := r.UpgradeState(ctx)
	upgrader, ok := upgraders[0]
	require.True(t, ok, "quality_monitor must register an upgrader for prior schema version 0")

	schemaResp := &resource.SchemaResponse{}
	r.(resource.Resource).Schema(ctx, resource.SchemaRequest{}, schemaResp)

	for _, tc := range []struct {
		name              string
		json              string
		expectWorkspaceID string
	}{
		{"empty_list_becomes_null", v1113QMStateEmptyPC, ""},
		{"list_with_object_becomes_object", v1113QMStateWithWorkspaceID, "12345"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: []byte(tc.json)}}
			resp := resource.UpgradeStateResponse{
				State: tfsdk.State{Schema: schemaResp.Schema},
			}

			upgrader.StateUpgrader(ctx, req, &resp)

			require.False(t, resp.Diagnostics.HasError(),
				"upgrader returned diagnostics: %s", resp.Diagnostics.Errors())
			require.NotNil(t, resp.State.Raw.Type(), "upgrader must populate State.Raw")

			pcVal, err := resp.State.Raw.ApplyTerraform5AttributePathStep(
				tftypes.AttributeName("provider_config"))
			require.NoError(t, err)
			pcTfVal, ok := pcVal.(tftypes.Value)
			require.True(t, ok)

			if tc.expectWorkspaceID == "" {
				assert.True(t, pcTfVal.IsNull(),
					"empty list must upgrade to null object, got: %v", pcTfVal)
				return
			}

			require.False(t, pcTfVal.IsNull())
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
