package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestCatalogCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalog())
}

func TestCatalogCreateAlsoDeletesDefaultSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/catalogs",
				ExpectedRequest: CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/unity-catalog/schemas/a.default",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/catalogs/a",
				Response: CatalogInfo{
					Name:    "a",
					Comment: "b",
					Properties: map[string]string{
						"c": "d",
					},
					MetastoreID: "e",
					Owner:       "f",
				},
			},
		},
		Resource: ResourceCatalog(),
		Create:   true,
		HCL: `
		name = "a"
		comment = "b"
		properties = {
			c = "d"
		}
		`,
	}.ApplyNoError(t)
}
