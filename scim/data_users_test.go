package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceDataUsers(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// TODO: run this test to get fixtures
		},
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
