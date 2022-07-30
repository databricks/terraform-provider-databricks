package jobs

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceQueryableJob(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// TODO: run this test to get fixtures
		},
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
