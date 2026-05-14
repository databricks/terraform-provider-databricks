package settings_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccDisableLegacyDbfsSetting(t *testing.T) {
	template := `
 	resource "databricks_disable_legacy_dbfs_setting" "this" {
 		disable_legacy_dbfs {
 			value = "true"
 		}
 	}
 	`
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: acceptance.ResourceCheckWithState("databricks_disable_legacy_dbfs_setting.this",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				etag := state.Attributes["etag"]
				require.NotEmpty(t, etag)
				// TODO: re-enable value assertion once workspace-settings estore staleness
				// (up to ~2min) is reduced. GET after PATCH may return stale value.
				return nil
			}),
	},
		acceptance.Step{
			Template: template,
			Destroy:  true,
			Check: acceptance.ResourceCheck("databricks_disable_legacy_dbfs_setting.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				// Reset the setting to its default so the workspace is left clean for the next run.
				_, err = w.Settings.DisableLegacyDbfs().Update(ctx, settings.UpdateDisableLegacyDbfsRequest{
					AllowMissing: true,
					Setting: settings.DisableLegacyDbfs{
						DisableLegacyDbfs: settings.BooleanMessage{
							Value: false,
						},
					},
					FieldMask: "disable_legacy_dbfs.value",
				})
				require.NoError(t, err)
				// TODO: re-enable post-reset value assertion once workspace-settings
				// estore staleness (up to ~2min) is reduced.
				return nil
			}),
		},
	)
}
