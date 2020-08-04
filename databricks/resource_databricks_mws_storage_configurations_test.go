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
	var bucket model.MWSStorageConfigurations
	config := EnvironmentTemplate(t, `
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
	bucketName := FirstKeyValue(t, config, "bucket_name")
	configName := FirstKeyValue(t, config, "storage_configuration_name")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
					conn := service.CommonEnvironmentClient()
					err := conn.MWSStorageConfigurations().Delete(bucket.AccountID, bucket.StorageConfigurationID)
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
	client := service.CommonEnvironmentClient()
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
		return errors.New("resource is not cleaned up")
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
		conn := service.CommonEnvironmentClient()
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
