package settings

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceNotificationDestinations(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// TODO: run this test to get fixtures
		},
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceNotificationDestinations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
