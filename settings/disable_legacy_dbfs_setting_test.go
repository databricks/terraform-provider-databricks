package settings_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
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
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				etag := state.Attributes["etag"]
				require.NotEmpty(t, etag)
				res, err := w.Settings.DisableLegacyDbfs().Get(ctx, settings.GetDisableLegacyDbfsRequest{
					Etag: etag,
				})
				require.NoError(t, err)
				// Check that the resource has been created and that it has the correct value.
				assert.Equal(t, res.DisableLegacyDbfs.Value, true)
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
				// Terraform Check returns the latest resource status before it is destroyed, which has an outdated eTag.
				// We are making an update call to get the correct eTag in the response error.
				_, err = w.Settings.DisableLegacyDbfs().Update(ctx, settings.UpdateDisableLegacyDbfsRequest{
					AllowMissing: true,
					Setting: settings.DisableLegacyDbfs{
						DisableLegacyDbfs: settings.BooleanMessage{
							Value: false,
						},
					},
					FieldMask: "disable_legacy_dbfs.value",
				})
				assert.Error(t, err)
				var aerr *apierr.APIError
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}
				etag := aerr.Details[0].Metadata["etag"]
				res, err := w.Settings.DisableLegacyDbfs().Get(ctx, settings.GetDisableLegacyDbfsRequest{
					Etag: etag,
				})
				// we should not be getting any error
				assert.NoError(t, err)
				// setting should go back to default
				assert.Equal(t, res.DisableLegacyDbfs.Value, false)
				return nil
			}),
		},
	)
}
