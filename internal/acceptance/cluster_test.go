package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClusterResource_CreateClusterWithLibraries(t *testing.T) {
	workspaceLevel(t, step{
		Template: `data "databricks_spark_version" "latest" {
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
					package = "databricks-sdk"
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
	})
}

func singleNodeClusterTemplate(autoTerminationMinutes string) string {
	return fmt.Sprintf(`
		data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "singlenode-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
			num_workers = 0
			autotermination_minutes = %s
			spark_conf = {
				"spark.databricks.cluster.profile" = "singleNode"
				"spark.master" = "local[*]"
			}
			custom_tags = {
				"ResourceClass" = "SingleNode"
			}
		}
	`, autoTerminationMinutes)
}

func TestAccClusterResource_CreateSingleNodeCluster(t *testing.T) {
	workspaceLevel(t, step{
		Template: singleNodeClusterTemplate("10"),
	}, step{
		Template: singleNodeClusterTemplate("20"),
	})
}

func awsClusterTemplate(availability string) string {
	return fmt.Sprintf(`
		data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "aws-cluster-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			num_workers = 1
			autotermination_minutes = 10
			aws_attributes {
				availability = "%s"
			}
			custom_tags = {
				"Owner" = "eng-dev-ecosystem-team@databricks.com"
			}
			node_type_id = "i3.xlarge"
		}
	`, availability)
}

func TestAccClusterResource_CreateAndUpdateAwsAttributes(t *testing.T) {
	loadWorkspaceEnv(t)
	if isAws(t) {
		workspaceLevel(t, step{
			Template: awsClusterTemplate("SPOT"),
		}, step{
			Template: awsClusterTemplate("SPOT_WITH_FALLBACK"),
		})
	}
}

func TestAccClusterResource_CreateAndNoWait(t *testing.T) {
	workspaceLevel(t, step{
		Template: `data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "nowait-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
			num_workers = 1
			autotermination_minutes = 10
			spark_conf = {
				"spark.databricks.cluster.profile" = "serverless"
			}
			custom_tags = {
				"ResourceClass" = "Serverless"
			}
			no_wait = true
		}`,
	})
}

func TestAccClusterResource_WorkloadType(t *testing.T) {
	workspaceLevel(t, step{
		Template: testAccClusterResourceWorkloadTypeTemplate(""),
	}, step{
		Template: testAccClusterResourceWorkloadTypeTemplate(`
		workload_type {
		    clients {
				jobs = true
				notebooks = true
			}
		}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.jobs", "true"),
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.notebooks", "true"),
		),
	}, step{
		Template: testAccClusterResourceWorkloadTypeTemplate(`
		workload_type {
		    clients {
				jobs = false
				notebooks = false
			}
		}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.jobs", "false"),
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.notebooks", "false"),
		),
	}, step{
		Template: testAccClusterResourceWorkloadTypeTemplate(`
		workload_type {
		    clients { }
		}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.jobs", "true"),
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.0.clients.0.notebooks", "true"),
		),
	}, step{
		Template: testAccClusterResourceWorkloadTypeTemplate(``),
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_cluster.this", "workload_type.#", "0"),
		),
	})
}

func testAccClusterResourceWorkloadTypeTemplate(workloadType string) string {
	return fmt.Sprintf(`
data "databricks_spark_version" "latest" {}
resource "databricks_cluster" "this" {
	cluster_name = "workload-{var.RANDOM}"
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
	%s
}`, workloadType)
}
