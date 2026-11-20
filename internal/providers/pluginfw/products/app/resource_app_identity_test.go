package app

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

type importedApp struct {
	Name        string
	WorkspaceID string
}

func TestAppImportState(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		importID string
		identity *testIdentity
		want     importedApp
	}{
		{
			name:     "plain ID",
			importID: "my-app",
			want:     importedApp{Name: "my-app"},
		},
		{
			name:     "identity with workspace_id",
			identity: &testIdentity{ID: types.StringValue("my-app"), WorkspaceID: types.StringValue("123")},
			want:     importedApp{Name: "my-app", WorkspaceID: "123"},
		},
		{
			name:     "identity without workspace_id",
			identity: &testIdentity{ID: types.StringValue("my-app"), WorkspaceID: types.StringNull()},
			want:     importedApp{Name: "my-app"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &resourceApp{}
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
			var providerConfig types.Object
			resp.State.GetAttribute(ctx, path.Root("name"), &name)
			resp.State.GetAttribute(ctx, path.Root("provider_config"), &providerConfig)
			workspaceID, _ := tfschema.GetWorkspaceIDResource(ctx, providerConfig)
			if diff := cmp.Diff(tc.want, importedApp{Name: name.ValueString(), WorkspaceID: workspaceID}); diff != "" {
				t.Errorf("imported state mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
