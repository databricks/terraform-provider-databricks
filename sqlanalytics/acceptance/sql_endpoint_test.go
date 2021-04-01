package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestPreviewAccSQLEndpoint(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_sql_endpoint" "this" {
				name = "tf-{var.RANDOM}"
				cluster_size = "Small"
				max_num_clusters = 1
			}`,
		},
	})
}
