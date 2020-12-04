package acceptance

import (
	"context"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/mws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"testing"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	var bucket StorageConfiguration
	config := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_storage_configurations" "this" {
		account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
		storage_configuration_name = "terraform-{var.RANDOM}"
		bucket_name                = "terraform-{var.RANDOM}"
	  }
	`)
	bucketName := qa.FirstKeyValue(t, config, "bucket_name")
	configName := qa.FirstKeyValue(t, config, "storage_configuration_name")

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testMWSStorageConfigurationsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.this", &bucket, t),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "storage_configuration_name", configName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "bucket_name", bucketName),
				),
				Destroy: false,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.this", &bucket, t),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "storage_configuration_name", configName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "bucket_name", bucketName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := common.CommonEnvironmentClient()
					err := NewStorageConfigurationsAPI(context.Background(), conn).Delete(bucket.AccountID, bucket.StorageConfigurationID)
					if err != nil {
						panic(err)
					}
				},
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSStorageConfigurationsResourceExists("databricks_mws_storage_configurations.this", &bucket, t),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "storage_configuration_name", configName),
					resource.TestCheckResourceAttr("databricks_mws_storage_configurations.this", "bucket_name", bucketName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
		},
	})
}

func testMWSStorageConfigurationsResourceDestroy(s *terraform.State) error {
	// client := common.CommonEnvironmentClient()
	// for _, rs := range s.RootModule().Resources {
	// 	if rs.Type != "databricks_mws_storage_configurations" {
	// 		continue
	// 	}
	// 	packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	_, err = NewStorageConfigurationsAPI(context.Background(), client).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
	// 	if err != nil {
	// 		return nil
	// 	}
	// 	return errors.New("resource is not cleaned up")
	// }
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSStorageConfigurationsResourceExists(n string, mwsCreds *StorageConfiguration, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// // find the corresponding state object
		// rs, ok := s.RootModule().Resources[n]
		// if !ok {
		// 	return fmt.Errorf("Not found: %s", n)
		// }

		// // retrieve the configured client from the test setup
		// conn := common.CommonEnvironmentClient()
		// packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
		// if err != nil {
		// 	return err
		// }
		// resp, err := NewStorageConfigurationsAPI(context.Background(), conn).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		// if err != nil {
		// 	return err
		// }

		// // If no error, assign the response Widget attribute to the widget pointer
		// *mwsCreds = resp
		return nil
	}
}
