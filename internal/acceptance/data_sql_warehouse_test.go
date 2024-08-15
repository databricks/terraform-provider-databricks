package acceptance

import (
	"testing"
)

func TestAccDataSourceWarehouse(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `
		data "databricks_sql_warehouse" "this" {
			id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		}
		
		output "warehouse_info" {
			value = data.databricks_sql_warehouse.this.name
		}`,
	})
}
