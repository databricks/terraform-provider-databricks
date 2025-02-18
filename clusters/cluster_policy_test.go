package clusters_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_cluster_policy" "external_metastore" {
			name = "Terraform policy {var.RANDOM}"
			definition = jsonencode({
				"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
					"type": "fixed",
					"value": "jdbc:sqlserver://<jdbc-url>"
				}
			})
		}`,
	}, acceptance.Step{
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

func TestAccClusterPolicyResourceOverrideBuiltIn(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_cluster_policy" "personal_vm" {
			name = "Personal Compute"
			policy_family_id = "personal-vm"
			policy_family_definition_overrides = jsonencode({
				"node_type_id": {
				  "type": "fixed",
				  "value": "Standard_DS3_v2"
				}
			  })
		}
		`,
	})
}

func TestAccClusterPolicyResourceOverrideNew(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_cluster_policy" "policyoverrideempty" {
			policy_family_id = "personal-vm"
			name             = "Policy Override {var.RANDOM}"
		  }
		  `,
	})
}
