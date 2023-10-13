package acceptance

import (
	"testing"
)

func TestAccMlflowModel(t *testing.T) {
	workspaceLevel(t,
		step{
			Template: `
			resource "databricks_mlflow_model" "this" {
			  name = "model-{var.RANDOM}"
			
			  description = "My MLflow model description"
			
			  tags {
				key   = "key1"
				value = "value1"
			  }
			  tags {
				key   = "key2"
				value = "value2"
			  }
			}`,
		},
		step{
			Template: `
			resource "databricks_mlflow_model" "this" {
			  name = "model-{var.RANDOM}"
			
			  description = "My MLflow model description"
			
			  tags {
				key   = "key1"
				value = "value1"
			  }
			  tags {
				key   = "key2"
				value = "value2"
			  }
			}			

			data "databricks_mlflow_model" "this" {
			  depends_on = [databricks_mlflow_model.this]
			  name       = databricks_mlflow_model.this.name
			}
			
			output "model" {
			  value = data.databricks_mlflow_model.this
			}`,
		})
}
