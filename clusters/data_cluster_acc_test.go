package clusters_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceCluster(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_cluster" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}

		output "cluster_info" {
			value = data.databricks_cluster.this.cluster_info
		}`,
	})
}
