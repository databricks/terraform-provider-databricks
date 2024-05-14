package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func checkTableDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		_, ok := s.Modules[0].Resources["data.databricks_table.this"]
		require.True(t, ok, "data.databricks_table.this has to be there")
		return nil
	}
}
func TestUcAccDataSourceTable(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}
		
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.RANDOM}"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}
		
		resource "databricks_sql_table" "mytable" {
			catalog_name = databricks_catalog.sandbox.id
			schema_name = databricks_schema.things.name
			name = "bar"
			table_type = "MANAGED"
			data_source_format = "DELTA"
			
			column {
				name = "id"
				type = "int"
			}
		}		

		data "databricks_table" "this" {
			name = databricks_sql_table.mytable.id
			depends_on = [
				databricks_sql_table.mytable,
			]
		}
		`,
		Check: checkTableDataSourcePopulated(t),
	})
}
