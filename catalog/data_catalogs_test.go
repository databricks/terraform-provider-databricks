package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCatalogsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs",
				Response: Catalogs{
					Catalogs: []CatalogInfo{
						{
							Name: "b",
						},
						{
							Name: "a",
						},
					},
				},
			},
		},
		Resource:    DataSourceCatalogs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCatalogs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
