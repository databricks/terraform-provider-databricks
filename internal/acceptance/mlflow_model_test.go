package acceptance

import (
	"os"
	"testing"
)

func TestAccMLflowModel(t *testing.T) {
	os.Setenv("CLOUD_ENV", "AZURE")
	workspaceLevel(t, step{
		Template: `
		resource "databricks_mlflow_model" "m1" {
			name = "tf-{var.RANDOM}"
			description = "tf-{var.RANDOM} description"
			
			tags {
				key   = "key-{var.RANDOM}"
				value = "{var.RANDOM}"
			}
		}
		`,
	})
}
