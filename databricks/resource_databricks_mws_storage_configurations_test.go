package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
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

func TestResourceMWSStorageConfigurationsCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/storage-configurations",
			ExpectedRequest: model.MWSStorageConfigurations{
				StorageConfigurationName: "Main Storage",
				RootBucketInfo: &model.RootBucketInfo{
					BucketName: "bucket",
				},
			},
			Response: model.MWSStorageConfigurations{
				StorageConfigurationID: "scid",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			Response: model.MWSStorageConfigurations{
				StorageConfigurationID:   "scid",
				StorageConfigurationName: "Main Storage",
				RootBucketInfo: &model.RootBucketInfo{
					BucketName: "bucket",
				},
			},
		},
	}, resourceMWSStorageConfigurations, map[string]interface{}{
		"account_id":                 "abc",
		"bucket_name":                "bucket",
		"storage_configuration_name": "Main Storage",
	}, resourceMWSStorageConfigurationsCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/scid", d.Id())
}

func TestResourceMWSStorageConfigurationsCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/storage-configurations",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSStorageConfigurations, map[string]interface{}{
		"account_id":                 "abc",
		"bucket_name":                "bucket",
		"storage_configuration_name": "Main Storage",
	}, resourceMWSStorageConfigurationsCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceMWSStorageConfigurationsRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			Response: model.MWSStorageConfigurations{
				StorageConfigurationID:   "scid",
				StorageConfigurationName: "Main Storage",
				RootBucketInfo: &model.RootBucketInfo{
					BucketName: "bucket",
				},
			},
		},
	}, resourceMWSStorageConfigurations, nil, actionWithID("abc/scid", resourceMWSStorageConfigurationsRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/scid", d.Id(), "Id should not be empty")
	assert.Equal(t, "bucket", d.Get("bucket_name"))
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "scid", d.Get("storage_configuration_id"))
	assert.Equal(t, "Main Storage", d.Get("storage_configuration_name"))
}

func TestResourceMWSStorageConfigurationsRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceMWSStorageConfigurations, nil, actionWithID("abc/scid", resourceMWSStorageConfigurationsRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceMWSStorageConfigurationsRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for correct url...
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSStorageConfigurations, nil, actionWithID("abc/scid", resourceMWSStorageConfigurationsRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/scid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceMWSStorageConfigurationsDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
		},
	}, resourceMWSStorageConfigurations, nil, actionWithID("abc/scid", resourceMWSStorageConfigurationsDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/scid", d.Id())
}

func TestResourceMWSStorageConfigurationsDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSStorageConfigurations, nil, actionWithID("abc/scid", resourceMWSStorageConfigurationsDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/scid", d.Id())
}
