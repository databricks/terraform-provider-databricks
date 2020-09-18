package acceptance

import (
	"fmt"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestIPACLListsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{

		Steps: []resource.TestStep{
			{
				Config: testIPACLListCreate(randomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_permissions.dummy_can_use",
						"object_type", "cluster-policy"),
					acceptance.ResourceCheck("databricks_permissions.dummy_can_use",
						func(client *common.DatabricksClient, id string) error {
							permissions, err := NewPermissionsAPI(client).Read(id)
							if err != nil {
								return err
							}
							assert.Len(t, permissions.AccessControlList, 3)
							return nil
						}),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testIPACLListCreate(name string) string {
	return fmt.Sprintf(`
	  resource "databricks_workspace_conf" "features" {
		enable_ip_access_lists = "true"
	  }	
	
	  resource "databricks_ip_access_list" "naughty_list" {
		label = "lumps_of_coal_%[1]s"
		list_type = "BLACKLIST"
		ip_addresses = [
		  "10.0.10.25","10.0.10.0/24"
		]
	  }
	`, name)
}
