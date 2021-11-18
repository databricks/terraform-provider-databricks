package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAccMLflowExperiment(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_mlflow_experiment" "e1" {
					name = "/Shared/tf-{var.RANDOM}"
					artifact_location = "dbfs:/tmp/tf-{var.RANDOM}"
					description = "tf-{var.RANDOM} description"
					
					tags {
					  key   = "k1-{var.RANDOM}"
					  value = "v1-{var.RANDOM}"
					}
					tags {
					  key   = "k2-{var.RANDOM}"
					  value = "v2-{var.RANDOM}"
					}
				}
			`,
		},
	})
}
