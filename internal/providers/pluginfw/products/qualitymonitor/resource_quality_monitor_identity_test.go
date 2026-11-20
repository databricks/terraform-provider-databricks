package qualitymonitor

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
)

type testIdentity struct {
	ID          types.String `tfsdk:"id"`
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

type importedMonitor struct {
	Name        string
	WorkspaceID string
}

func TestQualityMonitorImportState(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		importID string
		identity *testIdentity
		want     importedMonitor
	}{
		{
			name:     "plain ID",
			importID: "main.default.my_table",
			want:     importedMonitor{Name: "main.default.my_table"},
		},
		{
			name:     "identity with workspace_id",
			identity: &testIdentity{ID: types.StringValue("main.default.my_table"), WorkspaceID: types.StringValue("123")},
			want:     importedMonitor{Name: "main.default.my_table", WorkspaceID: "123"},
		},
		{
			name:     "identity without workspace_id",
			identity: &testIdentity{ID: types.StringValue("main.default.my_table"), WorkspaceID: types.StringNull()},
			want:     importedMonitor{Name: "main.default.my_table"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &QualityMonitorResource{}
			var schemaResp resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
			identitySchema := tfschema.WorkspaceResourceIdentitySchema()
			identity := &tfsdk.ResourceIdentity{Schema: identitySchema, Raw: tftypes.NewValue(identitySchema.Type().TerraformType(ctx), nil)}
			if tc.identity != nil {
				identity.Set(ctx, *tc.identity)
			}
			resp := resource.ImportStateResponse{State: tfsdk.State{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
			}}

			r.ImportState(ctx, resource.ImportStateRequest{ID: tc.importID, Identity: identity}, &resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("unexpected error: %v", resp.Diagnostics)
			}
			var name types.String
			var providerConfig types.List
			resp.State.GetAttribute(ctx, path.Root("table_name"), &name)
			resp.State.GetAttribute(ctx, path.Root("provider_config"), &providerConfig)
			workspaceID, _ := tfschema.GetWorkspaceID_SdkV2(ctx, providerConfig)
			if diff := cmp.Diff(tc.want, importedMonitor{Name: name.ValueString(), WorkspaceID: workspaceID}); diff != "" {
				t.Errorf("imported state mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
