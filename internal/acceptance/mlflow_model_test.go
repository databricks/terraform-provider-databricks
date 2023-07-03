package acceptance

import (
	"testing"
)

func TestAccMLflowModel(t *testing.T) {
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
