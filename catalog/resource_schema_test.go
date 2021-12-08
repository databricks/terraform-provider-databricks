package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestSchemaCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSchema())
}

func TestCreateSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/schemas",
				ExpectedRequest: SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					Comment:     "c",
				},
				Response: SchemaInfo{
					FullName: "b.a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/schemas/b.a",
				Response: SchemaInfo{
					MetastoreID: "d",
					Owner:       "e",
				},
			},
		},
		Resource: ResourceSchema(),
		Create:   true,
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		`,
	}.ApplyNoError(t)
}
