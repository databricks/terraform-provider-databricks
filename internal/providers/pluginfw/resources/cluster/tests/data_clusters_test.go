package tests

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDataSourceClustersNoFilter(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
		} `,
	})
}

func TestAccDataSourceClustersWithFilter(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_clusters_pluginframework" "this" {
			cluster_name_contains = "Default"
		}`,
	})
}
