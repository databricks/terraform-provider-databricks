package library_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccLibraryCreationPluginFramework(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "test-library-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
			autotermination_minutes = 10
			num_workers = 0
			spark_conf = {
				"spark.databricks.cluster.profile" = "singleNode"
				"spark.master" = "local[*]"
			}
			custom_tags = {
				"ResourceClass" = "SingleNode"
			}
		}
		resource "databricks_library_pluginframework" "new_library" {
			cluster_id = databricks_cluster.this.id
			pypi {
				repo = "https://pypi.org/dummy"
				package = "databricks-sdk"
			}
		}
		`,
	})
}

func TestAccLibraryUpdatePluginFramework(t *testing.T) {
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: `data "databricks_spark_version" "latest" {
				}
				resource "databricks_cluster" "this" {
					cluster_name = "cluster-{var.STICKY_RANDOM}"
					spark_version = data.databricks_spark_version.latest.id
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					autotermination_minutes = 10
					num_workers = 0
					spark_conf = {
						"spark.databricks.cluster.profile" = "singleNode"
						"spark.master" = "local[*]"
					}
					custom_tags = {
						"ResourceClass" = "SingleNode"
					}
				}
				resource "databricks_library_pluginframework" "new_library" {
					cluster_id = databricks_cluster.this.id
					pypi {
						repo = "https://pypi.org/simple"
						package = "databricks-sdk"
					}
				}
				`,
		},
		acceptance.Step{
			Template: `data "databricks_spark_version" "latest" {
			}
			resource "databricks_cluster" "this" {
				cluster_name = "cluster-{var.STICKY_RANDOM}"
				spark_version = data.databricks_spark_version.latest.id
				instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
				autotermination_minutes = 10
				num_workers = 0
				spark_conf = {
					"spark.databricks.cluster.profile" = "singleNode"
					"spark.master" = "local[*]"
				}
				custom_tags = {
					"ResourceClass" = "SingleNode"
				}
			}
			resource "databricks_library_pluginframework" "new_library" {
				cluster_id = databricks_cluster.this.id
				pypi {
					package = "networkx"
				}
			}
			`,
		},
	)
}
