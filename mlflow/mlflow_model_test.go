package mlflow_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccMLflowModel(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
