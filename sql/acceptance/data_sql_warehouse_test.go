package acceptance

import (
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"testing"
)

func TestAccDataSourceWarehouse(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_sql_endpoint" "this" {
				name             = "Endpoint of ${data.databricks_current_user.me.alphanumeric}"
				cluster_size     = "Small"
				max_num_clusters = 1
			
			}

			data "databricks_sql_warehouse" "this" {
				id = databricks_sql_endpoint.this.id
			}
			
			output "warehouse_info" {
				value = data.databricks_sql_warehouse.this.name
            }`,
		},
	})
}
