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
		Fixtures:    []qa.HTTPFixture{},
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
