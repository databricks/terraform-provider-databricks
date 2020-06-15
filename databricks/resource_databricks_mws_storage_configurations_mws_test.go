package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccMWSStorageConfigurations(t *testing.T) {
	var MWSStorageConfigurations model.MWSStorageConfigurations
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mwsAcctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	mwsHost := os.Getenv("DATABRICKS_MWS_HOST")
	storageConfigName := "test-mws-storage-configurations-tf"
	bucketName := "terraform-test-bucket"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testMWSStorageConfigurationsResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSStorageConfigurationsCreate(mwsAcctID, mwsHost, storageConfigName, bucketName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.my_mws_storage_configurations", &MWSStorageConfigurations, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "storage_configuration_name", storageConfigName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "bucket_name", bucketName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSStorageConfigurationsCreate(mwsAcctID, mwsHost, storageConfigName, bucketName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.my_mws_storage_configurations", &MWSStorageConfigurations, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "storage_configuration_name", storageConfigName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "bucket_name", bucketName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := getMWSClient()
					err := conn.MWSStorageConfigurations().Delete(MWSStorageConfigurations.AccountID, MWSStorageConfigurations.StorageConfigurationID)
					if err != nil {
						panic(err)
					}
				},
				// use a dynamic configuration with the random name from above
				Config: testMWSStorageConfigurationsCreate(mwsAcctID, mwsHost, storageConfigName, bucketName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.my_mws_storage_configurations", &MWSStorageConfigurations, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "storage_configuration_name", storageConfigName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.my_mws_storage_configurations", "bucket_name", bucketName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
		},
	})
}

func testMWSStorageConfigurationsResourceDestroy(s *terraform.State) error {
	client := getMWSClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_storage_configurations" {
			continue
		}
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = client.MWSStorageConfigurations().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSStorageConfigurationsResourceExists(n string, mwsCreds *model.MWSStorageConfigurations, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := getMWSClient()
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := conn.MWSStorageConfigurations().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*mwsCreds = resp
		return nil
	}
}

func testMWSStorageConfigurationsCreate(mwsAcctID, mwsHost, storageConfigName, bucketName string) string {
	return fmt.Sprintf(`
								provider "databricks" {
								  host = "%s"
								  basic_auth {}
								}
								resource "databricks_mws_storage_configurations" "my_mws_storage_configurations" {
								  account_id = "%s"
								  storage_configuration_name = "%s"
								  bucket_name         = "%s"
								}
								`, mwsHost, mwsAcctID, storageConfigName, bucketName)
}
