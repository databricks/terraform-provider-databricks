package acceptance

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/workspace"
	"github.com/stretchr/testify/assert"
)

func TestAccWorkspaceConfFullLifecycle(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_workspace_conf" "this" {
				custom_config = {
					"enableIpAccessLists": true
				}
			}`,
			Check: acceptance.ResourceCheck("databricks_workspace_conf.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					conf := map[string]interface{}{
						"enableIpAccessLists": nil,
					}
					err := workspace.NewWorkspaceConfAPI(ctx, client).Read(&conf)
					assert.NoError(t, err)
					assert.Len(t, conf, 1)
					assert.Equal(t, conf["enableIpAccessLists"], "true")
					return nil
				}),
		},
	})
}
