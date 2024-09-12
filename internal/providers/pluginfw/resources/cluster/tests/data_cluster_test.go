package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/require"
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
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("cluster1", clusterName),
	}, acceptance.Step{
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("cluster1", clusterName) + fmt.Sprintf(`
		data "databricks_cluster_pluginframework" "this" {
			cluster_name = "%s"
		}`, clusterName),
	})
}

func TestAccDataSourceClusterByNameError(t *testing.T) {
	clusterName := acceptance.RandomName("non-existent-cluster-")
	defer func() {
		if r := recover(); r != nil {
			// Convert the recovered value (which is typically an error) to a string
			errMsg := fmt.Sprintf("%v", r)
			// Assert that the error message contains "cluster not found"
			require.True(t, strings.Contains(errMsg, "there is no cluster with name"),
				"Expected error message to contain 'there is no cluster with name', but got: %s", errMsg)
		}
	}()
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
		data "databricks_cluster_pluginframework" "this" {
			cluster_name = "%s"
		}`, clusterName),
	})
}

func TestAccDataSourceClusterMultipleNamesError(t *testing.T) {
	clusterName := acceptance.RandomName("cluster-test-name-")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("cluster1", clusterName) + clusterResourceTemplate("cluster2", clusterName),
	}, acceptance.Step{
		Template: sparkVersionDataSourceTemplate + clusterResourceTemplate("cluster1", clusterName) + clusterResourceTemplate("cluster2", clusterName) + fmt.Sprintf(`
		data "databricks_cluster_pluginframework" "this" {
			cluster_name = "%s"
		}`, clusterName),
	})
}
