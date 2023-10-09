package acceptance

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestAccDefaultNamespaceSetting(t *testing.T) {
	t.Skip("Automatic test disabled until feature is enabled")
	workspaceLevel(t, step{
		Template: `
		resource "databricks_default_namespace_settings" "this" {
			namespace {
				value = "namespace_value"
			}
		}
		`,
		Check: resourceCheck("databricks_default_namespace_settings.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)
			res, err := w.Settings.ReadDefaultWorkspaceNamespace(ctx, settings.ReadDefaultWorkspaceNamespaceRequest{
				Etag: id,
			})
			assert.NoError(t, err)
			assert.Equal(t, res.Namespace.Value, "namespace_value")
			return nil
		}),
	},
		step{
			Destroy: true,
			Check: resourceCheck("databricks_default_namespace_settings.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				assert.NoError(t, err)
				// Terraform Check returns the latest resource status before it is destroyed, which has an outdated eTag.
				// We are making an update call to get the correct eTag in the response error.
				_, err = w.Settings.UpdateDefaultWorkspaceNamespace(ctx, settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					Setting: &settings.DefaultNamespaceSetting{
						Namespace: settings.StringMessage{
							Value: "this_call_should_fail",
						},
					},
					FieldMask: "namespace.value",
				})
				assert.Error(t, err)
				var aerr *apierr.APIError
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}
				etag := aerr.Details[0].Metadata["etag"]
				_, err = w.Settings.ReadDefaultWorkspaceNamespace(ctx, settings.ReadDefaultWorkspaceNamespaceRequest{
					Etag: etag,
				})
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}

				assert.Equal(t, aerr.ErrorCode, "RESOURCE_DOES_NOT_EXIST")
				return nil
			}),
		},
	)
}
