package acceptance

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestIPACLListsResourceFullLifecycle(t *testing.T) {
	randomName1 := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
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
				  `, randomName1),
			},
		},
	})
}
