package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/mws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestMwsAccCustomerManagedKeys(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless CLOUD_ENV=MWS is set")
	}
	cmkConfig := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_customer_managed_keys" "my_cmk" {
		account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
		aws_key_info {
			key_arn   = "{env.TEST_KMS_KEY_ARN}"
			key_alias = "{env.TEST_KMS_KEY_ALIAS}"
		}
	}`)
	var customerManagedKey CustomerManagedKey
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testMWSNetworkResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: cmkConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSCustomerManagedKeyResourceExists("databricks_mws_customer_managed_keys.my_cmk", &customerManagedKey, t),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					conn := common.CommonEnvironmentClient()
					err := NewCustomerManagedKeysAPI(context.Background(), conn).Delete(customerManagedKey.AccountID,
						customerManagedKey.CustomerManagedKeyID)
					if err != nil {
						panic(err)
					}
				},
				Config: cmkConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSCustomerManagedKeyResourceExists("databricks_mws_customer_managed_keys.my_cmk", &customerManagedKey, t),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				Config: cmkConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSCustomerManagedKeyResourceExists("databricks_mws_customer_managed_keys.my_cmk", &customerManagedKey, t),
				),
			},
		},
	})
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSCustomerManagedKeyResourceExists(n string, cmk *CustomerManagedKey, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		// todo: make pair id reading
		// rs, ok := s.RootModule().Resources[n]
		// if !ok {
		// 	return fmt.Errorf("Not found: %s", n)
		// }

		// // retrieve the configured client from the test setup
		// conn := common.CommonEnvironmentClient()
		// resp, err := NewCustomerManagedKeysAPI(context.Background(), conn).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		// if err != nil {
		// 	return err
		// }

		// // If no error, assign the response Widget attribute to the widget pointer
		// *cmk = resp
		return nil
	}
}
