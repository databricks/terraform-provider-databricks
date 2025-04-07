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

func TestAccDisableLegacyAccessSetting(t *testing.T) {
	template := `
 	resource "databricks_disable_legacy_access_setting" "this" {
 		disable_legacy_access {
 			value = "true"
 		}
 	}
 	`
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: acceptance.ResourceCheckWithState("databricks_disable_legacy_access_setting.this",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				etag := state.Attributes["etag"]
				require.NotEmpty(t, etag)
				res, err := w.Settings.DisableLegacyAccess().Get(ctx, settings.GetDisableLegacyAccessRequest{
					Etag: etag,
				})
				require.NoError(t, err)
				// Check that the resource has been created and that it has the correct value.
				assert.Equal(t, res.DisableLegacyAccess.Value, "true")
				return nil
			}),
	},
		acceptance.Step{
			Template: template,
			Destroy:  true,
			Check: acceptance.ResourceCheck("databricks_disable_legacy_access_setting.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				// Terraform Check returns the latest resource status before it is destroyed, which has an outdated eTag.
				// We are making an update call to get the correct eTag in the response error.
				_, err = w.Settings.DisableLegacyAccess().Update(ctx, settings.UpdateDisableLegacyAccessRequest{
					AllowMissing: true,
					Setting: settings.DisableLegacyAccess{
						DisableLegacyAccess: settings.BooleanMessage{
							Value: false,
						},
					},
					FieldMask: "disable_legacy_access.value",
				})
				assert.Error(t, err)
				var aerr *apierr.APIError
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}
				etag := aerr.Details[0].Metadata["etag"]
				res, err := w.Settings.DisableLegacyAccess().Get(ctx, settings.GetDisableLegacyAccessRequest{
					Etag: etag,
				})
				// we should not be getting any error
				assert.NoError(t, err)
				// setting should go back to default
				assert.Equal(t, res.DisableLegacyAccess.Value, false)
				return nil
			}),
		},
	)
}
