package acceptance

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var sparkVersionDataSourceTemplate = `
	data "databricks_spark_version" "latest" {}
`

func clusterResourceTemplate(resourceName string, clusterName string) string {
	return fmt.Sprintf(`
		resource "databricks_cluster" "%s" {
			cluster_name = "%s"
			instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
			spark_version = data.databricks_spark_version.latest.id
			num_workers = 1
			autotermination_minutes = 10
		}
	`, resourceName, clusterName)
}

func TestAccDataSourceClusterByID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_cluster_pluginframework" "this" {
			cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
		}`,
	})
}

func TestAccDataSourceClusterByName(t *testing.T) {
	clusterName := acceptance.RandomName("cluster-test-name-")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("this", clusterName),
	}, acceptance.Step{
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("this", clusterName) + fmt.Sprintf(`
		data "databricks_cluster_pluginframework" "this" {
			cluster_name = "%s"
		}`, clusterName),
	})
}
