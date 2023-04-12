package acceptance

import (
	"testing"
)

func TestAccResourceSqlTable(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name = "foo"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = "foo"
			table_type         = "MANAGED"
			data_source_format = "DELTA"
			column {
				name      = "id"
				position  = 0
				type_name = "INT"
				type_text = "int"
			}
			column {
				name      = "name"
				position  = 1
				type_name = "STRING"
				type_text = "varchar(64)"
			}
			comment = "this table is managed by terraform"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name = "foo"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = "foo"
			table_type         = "EXTNERAL"
			data_source_format = "DELTA"
			storage_location   = "s3://ext-main/foo/bar1"
			comment = "this table is managed by terraform"
		}`,
	})
}
