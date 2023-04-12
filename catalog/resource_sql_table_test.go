package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceSqlTableCreateStatement_External(t *testing.T) {
	ti := &SqlTableInfo{
		Name:             "bar",
		CatalogName:      "main",
		SchemaName:       "foo",
		TableType:        "EXTERNAL",
		DataSourceFormat: "DELTA",
		//ColumnInfos:           []SqlColumnInfo,
		StorageLocation:       "s3://ext-main/foo/bar1",
		StorageCredentialName: "somecred",
		//ViewDefinition        string            `json:"view_definition,omitempty"`
		Comment: "terraform managed",
		//Properties            map[string]string `json:"properties,omitempty"`
	}
	stmt := ti.buildTableCreateStatement()
	assert.Contains(t, stmt, "CREATE EXTERNAL TABLE main.foo.bar")
	assert.Contains(t, stmt, "USING DELTA")
	assert.Contains(t, stmt, "LOCATION 's3://ext-main/foo/bar1' WITH (CREDENTIAL `somecred`)")
	assert.Contains(t, stmt, "COMMENT 'terraform managed'")
}

func TestResourceSqlTableCreateStatement_View(t *testing.T) {
	ti := &SqlTableInfo{
		Name:             "bar",
		CatalogName:      "main",
		SchemaName:       "foo",
		TableType:        "VIEW",
		DataSourceFormat: "DELTA",
		//ColumnInfos:           []SqlColumnInfo,
		StorageLocation:       "s3://ext-main/foo/bar1",
		StorageCredentialName: "somecred",
		//ViewDefinition        string            `json:"view_definition,omitempty"`
		Comment: "terraform managed",
		//Properties            map[string]string `json:"properties,omitempty"`
	}
	stmt := ti.buildTableCreateStatement()
	assert.Contains(t, stmt, "CREATE VIEW main.foo.bar")
	assert.NotContains(t, stmt, "USING DELTA")
	assert.NotContains(t, stmt, "LOCATION 's3://ext-main/foo/bar1' WITH CREDENTIAL somecred")
	assert.Contains(t, stmt, "COMMENT 'terraform managed'")
}

func TestResourceSqlTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceTable())
}

func TestResourceSqlTableUpdate_External(t *testing.T) {
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

func TestResourceSqlTableUpdate_Managed(t *testing.T) {
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
