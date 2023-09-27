package acceptance

import (
	"context"
	"testing"

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
	})
}
