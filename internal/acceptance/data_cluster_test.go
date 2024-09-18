package acceptance

import (
	"testing"
)

func TestAccDataSourceCluster(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		data "databricks_cluster" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}
		
		output "cluster_info" {
			value = data.databricks_cluster.this.cluster_info
		}`,
	})
}
