package acceptance

import (
	"testing"
)

func TestAccDataSourceMlflowModel(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_mlflow_model" "this" {
		  name = "My MLflow Model"
		
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
		  name       = "My MLflow Model"
		}
		
		output "model" {
		  value = data.databricks_mlflow_model.this
		}`,
	})
}
