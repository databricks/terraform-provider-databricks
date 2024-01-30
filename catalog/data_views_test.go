package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestViewsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables?catalog_name=a&schema_name=b",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							CatalogName: "a",
							SchemaName:  "b",
							Name:        "c",
							FullName:    "a.b.c",
							TableType:   "MANAGED",
						},
						{
							CatalogName: "a",
							SchemaName:  "b",
							Name:        "d",
							FullName:    "a.b.d",
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
	}.ApplyAndExpectData(t, map[string]any{
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
	}.ExpectError(t, "i'm a teapot")
}
