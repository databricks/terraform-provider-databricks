package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceCurrentConfig(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": false,
		"cloud_type": "aws",
	})
}

func TestDataSourceCurrentConfigAccAzure(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		AccountID:   "123456",
		Azure:       true,
	}.ApplyAndExpectData(t, map[string]any{
		"account_id": "123456",
		"is_account": true,
		"cloud_type": "azure",
	})
}

func TestDataSourceCurrentConfigCloudOverride(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         `cloud = "gcp"`,
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": false,
		"cloud_type": "gcp",
		"cloud":      "gcp",
	})
}

func TestDataSourceCurrentConfigCloudOverrideAzure(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         `cloud = "azure"`,
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": false,
		"cloud_type": "azure",
		"cloud":      "azure",
	})
}

func TestDataSourceCurrentConfigCloudInvalidValue(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         `cloud = "invalid"`,
	}.ExpectError(t, "invalid config supplied. [cloud] expected cloud to be one of [aws azure gcp], got invalid")
}

func TestDataSourceCurrentConfigCloudEmptyValue(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         `cloud = ""`,
	}.ExpectError(t, "invalid config supplied. [cloud] expected cloud to be one of [aws azure gcp], got ")
}

func TestDataSourceCurrentConfigCloudOverrideAccountLevel(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		AccountID:   "acc-123",
		HCL:         `cloud = "aws"`,
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": true,
		"account_id": "acc-123",
		"cloud_type": "aws",
		"cloud":      "aws",
	})
}

func TestDataSourceCurrentConfig_ApiFieldAccount(t *testing.T) {
	// api = "account" without AccountID → is_account = true, account_id = ""
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         `api = "account"`,
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": true,
		"account_id": "",
	})
}

func TestDataSourceCurrentConfig_ApiFieldWorkspace(t *testing.T) {
	// Cross-override: api = "workspace" WITH AccountID → is_account = false
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		AccountID:   "acc-123",
		HCL:         `api = "workspace"`,
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": false,
	})
}

func TestDataSourceCurrentConfig_ApiFieldAccountWithCloud(t *testing.T) {
	// api = "account" + cloud = "gcp" → is_account = true, cloud_type = "gcp"
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentConfiguration(),
		ID:          ".",
		HCL:         "api = \"account\"\ncloud = \"gcp\"",
	}.ApplyAndExpectData(t, map[string]any{
		"is_account": true,
		"cloud_type": "gcp",
		"cloud":      "gcp",
	})
}
