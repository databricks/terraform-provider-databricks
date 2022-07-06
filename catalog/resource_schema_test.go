package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestSchemaCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSchema())
}

func TestCreateSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/schemas",
				ExpectedRequest: SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					Comment:     "c",
				},
				Response: SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				Response: SchemaInfo{
					MetastoreID: "d",
					Comment:     "c",
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

func TestCreateSchemaWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/schemas",
				ExpectedRequest: SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					Comment:     "c",
					Owner:       "administrators",
				},
				Response: SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "testers",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				Response: SchemaInfo{
					MetastoreID: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceSchema(),
		Create:   true,
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				Response: SchemaInfo{
					Name:        "a",
					MetastoreID: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"metastore_id": "d",
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}
