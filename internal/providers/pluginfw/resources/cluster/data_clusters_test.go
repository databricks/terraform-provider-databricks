package cluster_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceClustersNoFilterPluginFramework(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
		} `,
	})
}

func TestAccDataSourceClustersWithFilterPluginFramework(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
			cluster_name_contains = "Default"
		}`,
	})
}
