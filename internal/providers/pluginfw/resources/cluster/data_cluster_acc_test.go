package cluster_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceClusterByID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_cluster_pluginframework" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}`,
	})
}

func TestAccDataSourceClusterByName(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_cluster_pluginframework" "this" {
			cluster_name = "DEFAULT Test Cluster"
		}`,
	})
}
