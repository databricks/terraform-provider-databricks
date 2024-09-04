package acceptance

import (
	"testing"
)

func TestAccDataSourceClustersNoFilter(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_clusters" "this" {
		} `,
	})
}

func TestAccDataSourceClustersWithFilter(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_clusters" "this" {
			cluster_name_contains = "Default"
		}`,
	})
}

func TestAccDataSourceClustersNoFilterPluginFramework(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
		} `,
	})
}

func TestAccDataSourceClustersWithFilterPluginFramework(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
			cluster_name_contains = "Default"
		}`,
	})
}
