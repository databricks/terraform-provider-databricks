package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccRegisteredModel(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t,
		acceptance.Step{
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
		acceptance.Step{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				comment = "new comment"
			}
		`,
		},
		acceptance.Step{
			Template: `
			resource "databricks_registered_model" "model" {
				name = "terraform-test-registered-model-update-{var.STICKY_RANDOM}"
				catalog_name = "main"
				schema_name = "default"
				owner = "account users"
				comment = "new comment"
			}
			data "databricks_registered_model" "model" {
				full_name = databricks_registered_model.model.id
			}
			data "databricks_registered_model_versions" "model_versions" {
				full_name = databricks_registered_model.model.id
			}
			output "model" {
				value = data.databricks_registered_model.model
			}
			output "model_versions" {
				value = data.databricks_registered_model_versions.model_versions
			}
		`,
		},
	)
}
