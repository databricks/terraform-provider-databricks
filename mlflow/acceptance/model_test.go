package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAccMLflowModel(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_mlflow_model" "m1" {
					name = "tf-{var.RANDOM}"
					description = "tf-{var.RANDOM} description"
					
					tags {
					  key   = "key-{var.RANDOM}"
					  value = "{var.RANDOM}"
					}
					tags {
					  key   = "key-{var.RANDOM}"
					  value = "{var.RANDOM}"
					}
				}
			`,
		},
	})
}
