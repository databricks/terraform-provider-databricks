package acceptance

import (
	"testing"
)

func TestAccMLflowModel(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_mlflow_model" "m1" {
			name = "tf-{var.RANDOM}"
			description = "tf-{var.RANDOM} description"
			
			tags {
				key   = "key-{var.RANDOM}"
				value = "{var.RANDOM}"
			}
		}

		resource "databricks_permissions" "mlflow_model_permissions" {
			registered_model_id = databricks_mlflow_model.m1.registered_model_id
			access_control {
			  group_name       = "users"
			  permission_level = "CAN_READ"
			}
		  }
		`,
	})
}
