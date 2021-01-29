package acceptance

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
)

func TestAccClusterResource_CreateClusterWithLibraries(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_spark_version" "latest" {
			}
			data "databricks_node_type" "smallest" {
				local_disk = true
			}
			resource "databricks_cluster" "this" {
				cluster_name = "libs-{var.RANDOM}"
				spark_version = data.databricks_spark_version.latest.id
				node_type_id = data.databricks_node_type.smallest.id
				autotermination_minutes = 10
				autoscale {
					min_workers = 1
					max_workers = 2
				}
				spark_conf = {
					"spark.databricks.cluster.profile" = "serverless"
					"spark.databricks.repl.allowedLanguages" = "sql,python,r"
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
						package = "Faker"
						repo = "https://pypi.org/simple"
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
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_spark_version" "latest" {
			}
			data "databricks_node_type" "smallest" {
				local_disk = true
			}
			resource "databricks_cluster" "this" {
				cluster_name = "singlenode-{var.RANDOM}"
				spark_version = data.databricks_spark_version.latest.id
				node_type_id = data.databricks_node_type.smallest.id
				num_workers = 0
				autotermination_minutes = 10
				spark_conf = {
					"spark.databricks.cluster.profile" = "singleNode"
					"spark.master" = "local[*]"
				}
			}`,
		},
	})
}
