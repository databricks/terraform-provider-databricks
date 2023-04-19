package acceptance

import (
	"testing"
)

func TestAccResourceSqlTable_Managed(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "foom"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				this      = "that"
				something = "else"
			}

			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "varchar(64)"
			}
			comment = "this table is managed by terraform"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "foom"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				that      = "this"
				something = "else2"
			}
			
			column {
				name      = "id"
				typ       = "int"
			}
			column {
				name      = "name"
				type      = "string"
			}
			comment = "this table is managed by terraform..."
		}`,
	})
}

func TestAccResourceSqlTable_External(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "fooe"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "EXTERNAL"
			data_source_format = "DELTA"
			storage_location   = "s3://ext-main/foo/bar1"
			comment 		   = "this table is managed by terraform"
		}`,
	})
}

func TestAccResourceSqlTable_View(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "foov"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			comment 		   = "this table is managed by terraform"
			properties         = {
				this      = "that"
				something = "else"
			}

			column {
				name      = "id"
				type      = "string"
			}
		}
		
		resource "databricks_sql_table" "view" {
			name               = "bar_view"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "VIEW"
			comment 		   = "this view is managed by terraform"
			view_definition    = format("SELECT id FROM %s", databricks_sql_table.this.id)

			properties         = {
				this      = "that"
				something = "else"
			}
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "foov"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			data_source_format = "DELTA"
			comment 		   = "this table is managed by terraform..."
			properties         = {
				that = "this"
				else = "something"
			}

			column {
				name      = "id"
				type      = "string"
			}
		}
		
		resource "databricks_sql_table" "view" {
			name               = "bar_view"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "VIEW"
			comment 		   = "this view is managed by terraform..."
			view_definition    = format("SELECT id, 1 FROM %s", databricks_sql_table.this.id)
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "foov"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			data_source_format = "DELTA"
			comment 		   = "this table is managed by terraform..."

			column {
				name      = "id"
				type      = "string"
			}

			column {
				name      = "name"
				type      = "string"
			}
		}
		
		resource "databricks_sql_table" "view" {
			name               = "bar_view"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "VIEW"
			comment 		   = "this view is managed by terraform..."
			view_definition    = format("SELECT id, name FROM %s", databricks_sql_table.this.id)
		}`,
	})
}
