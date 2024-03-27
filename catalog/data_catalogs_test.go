package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCatalogsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/catalogs?",
				Response: catalog.ListCatalogsResponse{
					Catalogs: []catalog.CatalogInfo{
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
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a", "b"},
	})
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCatalogs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
