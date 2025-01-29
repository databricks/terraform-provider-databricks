package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccDataSourceClustersNoFilter(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		data "databricks_clusters" "this" {
		} `,
	})
}

func TestAccDataSourceClustersWithFilter(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		data "databricks_clusters" "this" {
			cluster_name_contains = "Default"
		}`,
	})
}

func checkFirstCluster(t *testing.T, f func(*compute.ClusterDetails)) func(*terraform.State) error {
	return func(s *terraform.State) error {
		w := databricks.Must(databricks.NewWorkspaceClient())
		firstClusterId, ok := s.RootModule().Resources["data.databricks_clusters.this"].Primary.Attributes["ids.0"]
		if ok {
			firstCluster, err := w.Clusters.GetByClusterId(context.Background(), firstClusterId)
			assert.NoError(t, err)
			f(firstCluster)
		}
		return nil
	}
}

func TestAccDataSourceClusters_FilterBy(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		data "databricks_clusters" "this" {
			filter_by {
				cluster_sources = ["UI", "API"]
			}
		}`,
		Check: checkFirstCluster(t, func(c *compute.ClusterDetails) {
			assert.Contains(t, []compute.ClusterSource{"UI", "API"}, c.ClusterSource)
		}),
	}, Step{
		Template: `
		data "databricks_clusters" "this" {
			filter_by {
				cluster_states = ["RUNNING", "RESIZING"]
			}
		}`,
		Check: checkFirstCluster(t, func(c *compute.ClusterDetails) {
			assert.Contains(t, []compute.State{"RUNNING", "RESIZING"}, c.State)
		}),
	}, Step{
		Template: `
		data "databricks_clusters" "this" {
			filter_by {
				is_pinned = true
			}
		}`,
		// Not possible to get whether a cluster is pinned or not
	}, Step{
		Template: `
		resource "databricks_cluster_policy" "this" {
			name = "test {var.RANDOM}"
			definition = jsonencode({
				"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
					"type": "fixed",
					"value": "jdbc:sqlserver://<jdbc-url>"
				}
			})
		}
		data "databricks_clusters" "this" {
			filter_by {
				policy_id = databricks_cluster_policy.this.id
			}
		}`,
		Check: checkFirstCluster(t, func(c *compute.ClusterDetails) {
			assert.Equal(t, "abc-123", c.PolicyId)
		}),
	})
}
