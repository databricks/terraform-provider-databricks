package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSQLEndpoint(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_sql_endpoint" "this" {
			name = "tf-{var.RANDOM}"
			cluster_size = "2X-Small"
			max_num_clusters = 1

			tags {
				custom_tags {
					key   = "Owner"
					value = "eng-dev-ecosystem-team@databricks.com"
				}
			}
		}`,
	}, step{
		Template: `
		resource "databricks_sql_endpoint" "that" {
			name = "tf-{var.RANDOM}"
			cluster_size = "2X-Small"
			max_num_clusters = 1

			enable_serverless_compute = false

			tags {
				custom_tags {
					key   = "Owner"
					value = "eng-dev-ecosystem-team@databricks.com"
				}
			}
		}`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			require.NoError(t, err)
			warehouseId := s.RootModule().Resources["databricks_sql_endpoint.that"].Primary.ID
			warehouse, err := w.Warehouses.GetById(context.Background(), warehouseId)
			require.NoError(t, err)
			assert.False(t, warehouse.EnableServerlessCompute)
			return nil
		},
	})
}
