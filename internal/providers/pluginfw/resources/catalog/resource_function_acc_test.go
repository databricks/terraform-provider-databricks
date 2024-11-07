package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func CheckFunctionResourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		_, ok := s.Modules[0].Resources["databricks_function.function"]
		require.True(t, ok, "databricks_function.function has to be in the Terraform state")
		return nil
	}
}

func TestFunctionResource(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
			resource "databricks_catalog" "sandbox" {
				name			= "sandbox-${var.STICKY_RANDOM}"
			}

			resource "databricks_schema" "functions" {
				catalog_name	= databricks_catalog.sandbox.id
				name			= "functions-${var.STICKY_RANDOM}"
			}

			resource "databricks_function" "function" {
				name 			  = "function-${var.STICKY_RANDOM}"
				catalog_name 	  = databricks_catalog.sandbox.id
				schema_name 	  = databricks_schema.functions.name
				input_params      = [
					{
						name = "weight"
						type = "DOUBLE"
					},
					{
						name = "height"
						type = "DOUBLE"
					}
				]
				data_type		  = "DOUBLE"
				routine_body	  = "SQL"
				routine_defintion = "weight / (height * height)"
				language		  = "SQL"
				is_deterministic  = true
				sql_data_access   = "CONTAINS_SQL"
				security_type 	  = "DEFINER"
			} 
		`,
		Check: CheckFunctionResourcePopulated(t),
	})
}
