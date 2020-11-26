package acceptance

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccWorkspaceConfFullLifecycle(t *testing.T) {
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `
				resource "databricks_workspace_conf" "this" {
					custom_config = {
						"enableIpAccessLists": true
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					acceptance.ResourceCheck("databricks_workspace_conf.this",
						func(client *common.DatabricksClient, id string) error {
							conf := map[string]interface{}{
								"enableIpAccessLists": nil,
							}
							err := workspace.NewWorkspaceConfAPI(context.Background(), client).Read(&conf)
							assert.NoError(t, err)
							assert.Len(t, conf, 1)
							assert.Equal(t, conf["enableIpAccessLists"], "true")
							return nil
						}),
				),
			},
		},
	})
}
