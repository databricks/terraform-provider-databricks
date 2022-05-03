package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestViewsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/?catalog_name=a&schema_name=b",
				Response: Tables{
					Tables: []TableInfo{
						{
							CatalogName: "a",
							SchemaName:  "b",
							Name:        "c",
							TableType:   "MANAGED",
						},
						{
							CatalogName: "a",
							SchemaName:  "b",
							Name:        "d",
							TableType:   "VIEW",
						},
					},
				},
			},
		},
		Resource: DataSourceViews(),
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": []string{"a.b.d"},
	})
}

func TestViewsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceViews(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
