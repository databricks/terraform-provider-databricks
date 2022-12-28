package acceptance

import (
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"testing"
)

func TestAccDataSourceCluster(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_spark_version" "latest" {}
			data "databricks_node_type" "smallest" {
				local_disk = true
			}

			resource "databricks_cluster" "this" {
				cluster_name            = "job-datasource-acceptance-test"
				spark_version           = data.databricks_spark_version.latest_lts.id
				node_type_id            = data.databricks_node_type.smallest.id
				autotermination_minutes = 20
				autoscale {
					min_workers = 1
					max_workers = 50
				}
			}

			data "databricks_cluster" "this" {
				cluster_id = databricks_cluster.this.id
			}
			
			output "cluster_info" {
				value = data.databricks_cluster.this.cluster_info
            }`,
		},
	})
}
