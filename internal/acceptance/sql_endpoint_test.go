package acceptance

import (
	"testing"
)

func TestAccSQLEndpoint(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_sql_endpoint" "this" {
			name = "tf-{var.RANDOM}"
			cluster_size = "2X-Small"
			max_num_clusters = 1
		}`,
	})
}
