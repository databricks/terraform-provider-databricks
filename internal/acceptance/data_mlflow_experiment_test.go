package acceptance

import (
	"testing"
)

func TestAccDataSourceMlflowExperiment(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data "databricks_current_user" "me" {}
		
		resource "databricks_mlflow_experiment" "this" {
		  name              = "${data.databricks_current_user.me.home}/Sample"
		  artifact_location = "dbfs:/tmp/my-experiment"
		  description       = "My MLflow experiment description"
		}
		
		data "databricks_mlflow_experiment" "this" {
		  depends_on = [databricks_mlflow_model.this]
		  name       = "${data.databricks_current_user.me.home}/Sample"
		}
		
		output "experiment" {
		  value = data.databricks_mlflow_experiment.this
		}`,
	})
}
