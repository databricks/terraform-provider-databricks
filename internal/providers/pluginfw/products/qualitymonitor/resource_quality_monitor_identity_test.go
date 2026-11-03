package qualitymonitor

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/databricks/terraform-provider-databricks/common"
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

func testClient(host string) *common.DatabricksClient {
	return &common.DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: &config.Config{Host: host, AccountID: "abc"}}}
}

func TestQualityMonitorImportState(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		host     string
		importID string
		identity *testIdentity
		want     importedMonitor
		wantErr  string
	}{
		{
			name:     "plain ID on a workspace host",
			host:     "https://workspace.cloud.databricks.com",
			importID: "main.default.my_table",
			want:     importedMonitor{Name: "main.default.my_table"},
		},
		{
			name:     "identity with workspace_id on an account host",
			host:     "https://accounts.cloud.databricks.com",
			identity: &testIdentity{ID: types.StringValue("main.default.my_table"), WorkspaceID: types.StringValue("123")},
			want:     importedMonitor{Name: "main.default.my_table", WorkspaceID: "123"},
		},
		{
			name:     "plain ID on an account host",
			host:     "https://accounts.cloud.databricks.com",
			importID: "main.default.my_table",
			wantErr:  "Use an import block with identity",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &QualityMonitorResource{Client: testClient(tc.host)}
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

			if tc.wantErr != "" {
				if !resp.Diagnostics.HasError() || !strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tc.wantErr) {
					t.Fatalf("want error containing %q, got %v", tc.wantErr, resp.Diagnostics)
				}
				return
			}
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
