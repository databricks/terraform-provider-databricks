package acceptance

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRestrictWorkspaceAdminsSetting(t *testing.T) {
	template := `
	resource "databricks_restrict_workspace_admins_setting" "this" {
		restrict_workspace_admins {
			status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
		}
	}
	`
	workspaceLevel(t, step{
		Template: template,
		Check: resourceCheckWithState("databricks_restrict_workspace_admins_setting.this",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				etag := state.Attributes["etag"]
				require.NotEmpty(t, etag)
				res, err := w.Settings.GetRestrictWorkspaceAdminsSetting(ctx, settings.GetRestrictWorkspaceAdminsSettingRequest{
					Etag: etag,
				})
				assert.NoError(t, err)
				// Check that the resource has been created and that it has the correct value.
				assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", res.RestrictWorkspaceAdmins.Status.String())
				return nil
			}),
	},
		step{
			Template: template,
			Destroy:  true,
			Check: resourceCheck("databricks_restrict_workspace_admins_setting.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					ctx = context.WithValue(ctx, common.Api, common.API_2_1)
					w, err := client.WorkspaceClient()
					assert.NoError(t, err)
					// Terraform Check returns the latest resource status before it is destroyed, which has an outdated eTag.
					// We are making an update call to get the correct eTag in the response error.
					_, err = w.Settings.UpdateRestrictWorkspaceAdminsSetting(ctx, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
						AllowMissing: true,
						Setting: settings.RestrictWorkspaceAdminsSetting{
							RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
								Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
							},
						},
						FieldMask: "restrict_workspace_admins.status",
					})
					assert.Error(t, err)
					var aerr *apierr.APIError
					if !errors.As(err, &aerr) {
						assert.FailNow(t, "cannot parse error message %v", err)
					}
					etag := aerr.Details[0].Metadata["etag"]
					res, err := w.Settings.GetRestrictWorkspaceAdminsSetting(ctx, settings.GetRestrictWorkspaceAdminsSettingRequest{
						Etag: etag,
					})
					// we should not be getting any error
					assert.NoError(t, err)
					// workspace should go back to default
					assert.Equal(t, res.RestrictWorkspaceAdmins.Status.String(), "ALLOW_ALL")
					return nil
				}),
		},
	)
}
