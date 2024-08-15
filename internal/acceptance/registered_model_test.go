package acceptance

import (
	"testing"
)

func TestUcAccRegisteredModel(t *testing.T) {
	unityWorkspaceLevel(t,
		LegacyStep{
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
		LegacyStep{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				comment = "new comment"
			}
		`,
		},
		LegacyStep{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-update-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				owner = "account users"
				comment = "new comment"
			}
		`,
		},
	)
}
