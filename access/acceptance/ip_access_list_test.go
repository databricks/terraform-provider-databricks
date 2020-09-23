package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestIPACLListsIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	assert.Equal(t, "NOTHING", cloudEnv, "Got cloud env: "+cloudEnv)
	config := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_credentials" "my_e2_credentials" {
		account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
		credentials_name = "creds-test-{var.RANDOM}"
		role_arn         = "arn:aws:iam::999999999999:role/tf-test-{var.RANDOM}"
	}`)
}

// func TestIPACLListsResourceFullLifecycle(t *testing.T) {
// 	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	acceptance.AccTest(t, resource.TestCase{

// 		Steps: []resource.TestStep{
// 			{
// 				Config: testIPACLListCreate(randomName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("databricks_permissions.dummy_can_use",
// 						"object_type", "cluster-policy"),
// 					acceptance.ResourceCheck("databricks_permissions.dummy_can_use",
// 						func(client *common.DatabricksClient, id string) error {
// 							permissions, err := NewPermissionsAPI(client).Read(id)
// 							if err != nil {
// 								return err
// 							}
// 							assert.Len(t, permissions.AccessControlList, 3)
// 							return nil
// 						}),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// func testIPACLListCreate(name string) string {
// 	return fmt.Sprintf(`
// 	  resource "databricks_workspace_conf" "features" {
// 		enable_ip_access_lists = "true"
// 	  }

// 	  resource "databricks_ip_access_list" "naughty_list" {
// 		label = "lumps_of_coal_%[1]s"
// 		list_type = "BLACKLIST"
// 		ip_addresses = [
// 		  "10.0.10.25","10.0.10.0/24"
// 		]
// 	  }
// 	`, name)
// }
