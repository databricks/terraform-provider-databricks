package sql_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceWarehouse(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_sql_warehouse" "this" {
			id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		}
		
		output "warehouse_info" {
			value = data.databricks_sql_warehouse.this.name
		}`,
	})
}
