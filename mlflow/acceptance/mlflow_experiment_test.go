package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccMLflowExperiment(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_mlflow_experiment" "e1" {
					name = "/Shared/tf-{var.RANDOM}"
					artifact_location = "dbfs:/tmp/tf-{var.RANDOM}"
					description = "tf-{var.RANDOM} description"
				}
			`,
		},
	})
}
