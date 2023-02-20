package acceptance

import (
	"testing"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_cluster_policy" "external_metastore" {
			name = "Terraform policy {var.RANDOM}"
			definition = jsonencode({
				"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
					"type": "fixed",
					"value": "jdbc:sqlserver://<jdbc-url>"
				}
			})
		}`,
	}, step{
		// renaming to a new random name
		Template: `resource "databricks_cluster_policy" "external_metastore" {
			name = "Terraform policy {var.RANDOM}"
			definition = jsonencode({
				"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
					"type": "fixed",
					"value": "jdbc:sqlserver://<jdbc-url>"
				}
			})
		}`,
	})
}
