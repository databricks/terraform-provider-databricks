package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccSQLEndpoint(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_sql_endpoint" "this" {
				name = "tf-{var.RANDOM}"
				cluster_size = "XX-Small"
				max_num_clusters = 1
			}`,
		},
	})
}
