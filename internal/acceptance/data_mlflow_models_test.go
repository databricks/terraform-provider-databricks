package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataMlflowModels(t *testing.T) {
	WorkspaceLevel(t,
		Step{
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

			data "databricks_mlflow_models" "this" {
			  depends_on = [databricks_mlflow_model.this]
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mlflow_models.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				names := r.Primary.Attributes["names.#"]
				if names == "" {
					return fmt.Errorf("names are empty: %v", r.Primary.Attributes)
				}
				return nil
			},
		})
}
