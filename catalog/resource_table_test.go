package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceTable())
}

func TestTableCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/tables",
				ExpectedRequest: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "EXTERNAL",
					DataSourceFormat: "JSON",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "int",
							TypeText: "int",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/main.foo.bar",
			},
		},
		Resource: ResourceTable(),
		Create:   true,
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		table_type = "EXTERNAL"
		data_source_format = "JSON"
		storage_location = "s3://ext-main/foo/bar"
		
		column {
			name = "id"
			type_text = "int"
			type_name = "int"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}

func TestTableCreateWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/tables",
				ExpectedRequest: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "EXTERNAL",
					DataSourceFormat: "JSON",
					Owner:            "administrators",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "int",
							TypeText: "int",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/tables/main.foo.bar",
				Response: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "EXTERNAL",
					DataSourceFormat: "JSON",
					Owner:            "testers",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "int",
							TypeText: "int",
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
				},
			},
		},
		Resource: ResourceTable(),
		Create:   true,
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		owner = "administrators"
		table_type = "EXTERNAL"
		data_source_format = "JSON"
		storage_location = "s3://ext-main/foo/bar"
		column {
			name = "id"
			type_text = "int"
			type_name = "int"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}
