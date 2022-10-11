package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccCreateShare(t *testing.T) {
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
			resource "databricks_schema" "things" {
				catalog_name = databricks_catalog.sandbox.id
				name         = "things{var.RANDOM}"
				comment      = "this database is managed by terraform"
				properties = {
					kind = "various"
				}
			}
			
			resource "databricks_table" "mytable" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar"
				table_type = "MANAGED"
				data_source_format = "DELTA"
				
				column {
					name      = "id"
					position  = 0
					type_name = "INT"
					type_text = "int"
					type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
				}
			}
			
			resource "databricks_share" "myshare" {
				name = "{var.RANDOM}-terraform-delta-share"
				object {
					name = databricks_table.mytable.id
					comment = "c"
					data_object_type = "TABLE"
				}				
			}
			`,
		},
	})
}
