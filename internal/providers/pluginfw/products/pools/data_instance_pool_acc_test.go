package pools_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceInstancePool(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_node_type" "smallest" {
			local_disk = true
		}

		resource "databricks_instance_pool" "this" {
			instance_pool_name = "tf-pool-{var.RANDOM}"
			min_idle_instances = 0
			max_capacity       = 10
			node_type_id       = data.databricks_node_type.smallest.id
			idle_instance_autotermination_minutes = 10
		}

		data "databricks_instance_pool" "this" {
			name = databricks_instance_pool.this.instance_pool_name
            depends_on = [databricks_instance_pool.this]
		}
		`,
	})
}
