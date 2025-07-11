package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceTable())
}

func TestTableCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/tables",
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
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
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
				Resource: "/api/2.1/unity-catalog/tables",
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
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
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
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]any{
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

func TestTableUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]interface{}{
					"columns": map[string]interface{}{
						"comment":            "",
						"name":               "id",
						"nullable":           true,
						"partition_index":    0,
						"position":           0,
						"type_interval_type": "",
						"type_json":          "",
						"type_name":          "string",
						"type_precision":     0,
						"type_scale":         0,
						"type_text":          "string",
					},
					"storage_location": "s3://ext-main/foo/bar1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: TableInfo{
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
							TypeName: "string",
							TypeText: "string",
						},
					},
				},
			},
		},
		Resource: ResourceTable(),
		Update:   true,
		ID:       "main.foo.bar",
		InstanceState: map[string]string{
			"catalog_name":       "main",
			"schema_name":        "foo",
			"name":               "bar",
			"table_type":         "EXTERNAL",
			"data_source_format": "JSON",
			"storage_location":   "s3://ext-main/foo/bar",
			"column": `[{
				\"comment\": \"\",
				\"name\": \"id\",
				\"nullable\": true,
				\"partition_index\": 0,
				\"position\": 0,
				\"type_interval_type\": \"\",
				\"type_json\": \"\",
				\"type_name\": \"int\",
				\"type_precision\": 0,
				\"type_scale\": 0,
				\"type_text\": \"int\"
			}]`,
		},
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		table_type = "EXTERNAL"
		data_source_format = "JSON"
		storage_location = "s3://ext-main/foo/bar1"

		column {
			name = "id"
			type_text = "string"
			type_name = "string"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}

func TestManagedTableUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				ExpectedRequest: map[string]interface{}{
					"columns": map[string]interface{}{
						"comment":            "",
						"name":               "id",
						"nullable":           true,
						"partition_index":    0,
						"position":           0,
						"type_interval_type": "",
						"type_json":          "",
						"type_name":          "string",
						"type_precision":     0,
						"type_scale":         0,
						"type_text":          "string",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables/main.foo.bar",
				Response: TableInfo{
					StorageLocation:  "s3://ext-main/foo/bar",
					Name:             "bar",
					CatalogName:      "main",
					SchemaName:       "foo",
					TableType:        "MANAGED",
					DataSourceFormat: "JSON",
					ColumnInfos: []ColumnInfo{
						{
							Name:     "id",
							Nullable: true,
							Position: 0,
							TypeName: "string",
							TypeText: "string",
						},
					},
				},
			},
		},
		Resource: ResourceTable(),
		Update:   true,
		ID:       "main.foo.bar",
		InstanceState: map[string]string{
			"catalog_name":       "main",
			"schema_name":        "foo",
			"name":               "bar",
			"table_type":         "MANAGED",
			"data_source_format": "JSON",
			"storage_location":   "s3://ext-main/foo/bar",
			"column": `[{
				\"comment\": \"\",
				\"name\": \"id\",
				\"nullable\": true,
				\"partition_index\": 0,
				\"position\": 0,
				\"type_interval_type\": \"\",
				\"type_json\": \"\",
				\"type_name\": \"int\",
				\"type_precision\": 0,
				\"type_scale\": 0,
				\"type_text\": \"int\"
			}]`,
		},
		HCL: `
		catalog_name = "main"
		schema_name = "foo"
		name = "bar"
		table_type = "MANAGED"
		data_source_format = "JSON"

		column {
			name = "id"
			type_text = "string"
			type_name = "string"
			position = 0
		}
		`,
	}.ApplyNoError(t)
}
