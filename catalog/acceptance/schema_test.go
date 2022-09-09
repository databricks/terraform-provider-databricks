package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccSchema(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_catalog" "sandbox" {
				name         = "sandbox{var.RANDOM}"
				comment      = "this catalog is managed by terraform"
				properties = {
					purpose = "testing"
				}
			}

			data "databricks_catalogs" "all" {
				depends_on = [databricks_catalog.sandbox]
			}

			resource "databricks_grants" "sandbox" {
				catalog = databricks_catalog.sandbox.name
				grant {
					principal  = "{env.TEST_DATA_SCI_GROUP}"
					privileges = ["USAGE", "CREATE"]
				}
				grant {
					principal  = "{env.TEST_DATA_ENG_GROUP}"
					privileges = ["USAGE"]
				}
			}
			
			resource "databricks_schema" "things" {
				catalog_name = databricks_catalog.sandbox.id
				name         = "things{var.RANDOM}"
				comment      = "this database is managed by terraform"
				properties = {
					kind = "various"
				}
			}

			data "databricks_schemas" "sandbox" {
				catalog_name = databricks_catalog.sandbox.id
				depends_on = [databricks_schema.things]
			}			  

			resource "databricks_grants" "things" {
				schema = databricks_schema.things.id
				grant {
				  principal  = "{env.TEST_DATA_ENG_GROUP}"
				  privileges = ["USAGE"]
				}
			}`,
		},
	})
}
