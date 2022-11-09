package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccClusterResource_CreateClusterWithLibraries(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_spark_version" "latest" {
			}
			resource "databricks_cluster" "this" {
				cluster_name = "libs-{var.RANDOM}"
				spark_version = data.databricks_spark_version.latest.id
				instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
				autotermination_minutes = 10
				num_workers = 1
				spark_conf = {
					"spark.databricks.repl.allowedLanguages" = "sql,python,r"
					"spark.databricks.cluster.profile" = "serverless"
				}
				custom_tags = {
					"ResourceClass" = "Serverless"
				}
				library {
					maven {
						coordinates = "org.jsoup:jsoup:1.7.2"
						repo = "https://mavencentral.org"
						exclusions = ["slf4j:slf4j"]
					}
				}
				library {
					pypi {
						repo = "https://pypi.org/simple"
						package = "Faker"
					}
				}
				library {
					pypi {
						package = "networkx"
					}
				}
				library {
					maven {
						coordinates = "com.microsoft.azure:azure-eventhubs-spark_2.11:2.3.7"
					}
				}
			}`,
		},
	})
}

func TestAccClusterResource_CreateSingleNodeCluster(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_spark_version" "latest" {
			}
			resource "databricks_cluster" "this" {
				cluster_name = "singlenode-{var.RANDOM}"
				spark_version = data.databricks_spark_version.latest.id
				instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
				num_workers = 0
				autotermination_minutes = 10
				spark_conf = {
					"spark.databricks.cluster.profile" = "singleNode"
					"spark.master" = "local[*]"
				}
				custom_tags = {
					"ResourceClass" = "SingleNode"
				}
			}`,
		},
	})
}
