package acceptance

import (
	"testing"
)

func TestAccMLflowExperiment(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_mlflow_experiment" "e1" {
			name = "/Shared/tf-{var.RANDOM}"
			artifact_location = "dbfs:/tmp/tf-{var.RANDOM}"
			description = "tf-{var.RANDOM} description"
		}
		`,
	})
}
