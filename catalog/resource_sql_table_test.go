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
	qa.ResourceCornerCases(t, ResourceSqlTable())
}
