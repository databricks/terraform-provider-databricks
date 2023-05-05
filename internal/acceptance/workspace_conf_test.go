package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestAccWorkspaceConfFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_workspace_conf" "this" {
			custom_config = {
				"enableIpAccessLists": true
			}
		}`,
		Check: resourceCheck("databricks_workspace_conf.this",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				if err != nil {
					return err
				}
				conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
					Keys: "enableIpAccessLists",
				})
				if err != nil {
					return err
				}
				assert.Len(t, *conf, 1)
				assert.Equal(t, (*conf)["enableIpAccessLists"], "true")
				return nil
			}),
	})
}
