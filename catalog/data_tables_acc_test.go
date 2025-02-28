package catalog_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkTablesDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_tables.this"]
		require.True(t, ok, "data.databricks_shares.this has to be there")

		attr := r.Primary.Attributes

		assert.Equal(t, s.Modules[0].Resources["databricks_sql_table.mytable"].Primary.ID, attr["ids.0"])
		assert.Equal(t, s.Modules[0].Resources["databricks_sql_table.mytable_2"].Primary.ID, attr["ids.1"])

		num_tables, _ := strconv.Atoi(s.Modules[0].Outputs["tables"].Value.(string))
		assert.Equal(t, num_tables, 2)
		return nil
	}
}
func TestUcAccDataSourceTables(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
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
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			
			column {
				name = "id"
				type = "int"
			}
		}

		resource "databricks_sql_table" "mytable_2" {
			catalog_name = databricks_catalog.sandbox.id
			schema_name = databricks_schema.things.name
			name = "bar_2"
			table_type = "MANAGED"
			data_source_format = "DELTA"
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			
			column {
				name = "id"
				type = "int"
			}
		}			

		data "databricks_tables" "this" {
			catalog_name = databricks_catalog.sandbox.id
			schema_name = databricks_schema.things.name			
			depends_on = [
				databricks_sql_table.mytable,
				databricks_sql_table.mytable_2
			]
		}
		output "tables" {
			value = length(data.databricks_tables.this.ids)
		}
		`,
		Check: checkTablesDataSourcePopulated(t),
	})
}
