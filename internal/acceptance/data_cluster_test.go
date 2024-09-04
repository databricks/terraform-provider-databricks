package acceptance

import (
	"testing"
)

func TestAccDataSourceCluster(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_cluster" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}
		
		output "cluster_info" {
			value = data.databricks_cluster.this.cluster_info
		}`,
	})
}

func TestAccDataSourceClusterPluginFramework(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_cluster_pluginframework" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}
		
		output "cluster_info" {
			value = data.databricks_cluster.this.cluster_info
		}`,
	})
}
