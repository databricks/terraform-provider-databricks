package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataMlflowModel(t *testing.T) {
	workspaceLevel(t,
		step{
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
			}`,
		},
		step{
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

			data "databricks_mlflow_model" "this" {
			  depends_on = [databricks_mlflow_model.this]
			  name       = databricks_mlflow_model.this.name
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mlflow_model.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				id := r.Primary.Attributes["id"]
				if id == "" {
					return fmt.Errorf("id is empty: %v", r.Primary.Attributes)
				}
				expect := "My MLflow model description"
				description := r.Primary.Attributes["description"]
				if description != expect {
					return fmt.Errorf("incorrect description. expected: %v, received: %v",
						expect, description)
				}
				return nil
			},
		})
}
