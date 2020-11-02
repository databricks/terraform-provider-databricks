package acceptance

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfFullLifecycle(t *testing.T) {
	acceptance.AccTest(t, resource.TestCase{

		Steps: []resource.TestStep{
			{
				Config: `
				resource "databricks_workspace_conf" "features" {
					enable_ip_access_lists = "true"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					acceptance.ResourceCheck("databricks_workspace_conf.features",
						func(client *common.DatabricksClient, id string) error {
							workspaceConf, err := workspace.NewWorkspaceConfAPI(client).Read("enableIpAccessLists")
							if err != nil {
								return err
							}
							assert.Len(t, workspaceConf, 1)
							assert.Equal(t, workspaceConf["enableIpAccessLists"], "true")
							return nil
						}),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
