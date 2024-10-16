package cluster_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const dataClusterTemplateById = `
	data "databricks_cluster "by_id" {
		cluster_id = "{env.TEST_DEFAULT_CLUSTER_ID}"
	}
`

func TestAccDataSourceClusterByID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: dataClusterTemplateById,
	})
}

func TestAccDataSourceClusterByName(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: dataClusterTemplateById + `
		data "databricks_cluster" "by_name" {
			cluster_name = data.databricks_cluster.by_id.cluster_name
		}`,
	})
}
