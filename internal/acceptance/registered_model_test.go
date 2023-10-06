package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccRegisteredModel(t *testing.T) {
	name := fmt.Sprintf("terraform-test-registered-model-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	workspaceLevel(t,
		step{
			Template: fmt.Sprintf(`
			resource "databricks_registered_model" "model" {
				name = "%[1]s"
				catalog_name = "main"
				schema_name = "default"
			}
			
			resource "databricks_grants" "model_grants" {
				model = databricks_registered_model.model.id
			  
				grant {
				  principal = "account users"
				  privileges = ["EXECUTE"]
				}
			
		`, name),
		},
		step{
			Template: fmt.Sprintf(`
			resource "databricks_registered_model" "model" {
				name = "%[1]s"
				catalog_name = "main"
				schema_name = "default"
				comment = "new comment"
			}
		`, name),
		},
	)
}
