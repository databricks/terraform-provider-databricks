package mlflow_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccMLflowExperiment(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mlflow_experiment" "e1" {
			name = "/Shared/tf-{var.RANDOM}"
			artifact_location = "dbfs:/tmp/tf-{var.RANDOM}"
			description = "tf-{var.RANDOM} description"
		}
		`,
	})
}

func TestAccMLflowExperimentWithDiff(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mlflow_experiment" "e1" {
			name = "/Workspace/Shared/tf-{var.RANDOM}/"
			artifact_location = "dbfs:/tmp/tf-{var.RANDOM}"
			description = "tf-{var.RANDOM} description"
		}
		`,
	})
}
