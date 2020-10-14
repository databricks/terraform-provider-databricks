package acceptance

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestIPACLListsResourceFullLifecycle(t *testing.T) {
	randomName1 := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	randomName2 := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	acceptance.AccTest(t, resource.TestCase{

		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "databricks_workspace_conf" "features" {
				enable_ip_access_lists = "true"
				}
				  
				resource "databricks_ip_access_list" "%[1]s" {
					label = "%[1]s"
					list_type = "BLOCK"
					ip_addresses = [
						"10.0.10.25","10.0.10.0/24"
					]
					depends_on = [databricks_workspace_conf.features]
				}
				  `, randomName1, randomName2),

				//   resource "databricks_ip_access_list" "%[2]s" {
				// 	label = "%[2]s"
				// 	list_type = "ALLOW"
				// 	ip_addresses = [
				// 	  "0.0.0.0/0"
				// 	]
				// 	depends_on = [databricks_workspace_conf.features]
				//   }
				//   `, randomName1, randomName2),
				Check: resource.ComposeTestCheckFunc(
					acceptance.ResourceCheck("databricks_workspace_conf.features",
						func(client *common.DatabricksClient, id string) error {
							ipAccessList, err := access.NewIPAccessListsAPI(client).Read(id)
							if err != nil {
								return err
							}
							assert.Equal(t, ipAccessList.Label, randomName1)
							return nil
						}),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
