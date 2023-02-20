package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
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
				conf := map[string]any{
					"enableIpAccessLists": nil,
				}
				err := workspace.NewWorkspaceConfAPI(ctx, client).Read(&conf)
				assert.NoError(t, err)
				assert.Len(t, conf, 1)
				assert.Equal(t, conf["enableIpAccessLists"], "true")
				return nil
			}),
	})
}
