package acceptance

import (
	"testing"
)

func TestUcAccRegisteredModel(t *testing.T) {
	UnityWorkspaceLevel(t,
		Step{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
			}
			
			resource "databricks_grants" "model_grants" {
				model = databricks_registered_model.model.id
			  
				grant {
				  principal = "account users"
				  privileges = ["EXECUTE"]
				}
			}
		`,
		},
		Step{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				comment = "new comment"
			}
		`,
		},
		Step{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-update-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				owner = "account users"
				comment = "new comment"
			}
			data "databricks_registered_model" "model" {
				full_name = databricks_registered_model.model.full_name
				include_model_versions = true
			}
			output "model" {
				value = data.databricks_registered_model.model
			}
		`,
		},
	)
}
