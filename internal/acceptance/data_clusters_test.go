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
