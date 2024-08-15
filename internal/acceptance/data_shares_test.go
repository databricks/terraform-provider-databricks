package acceptance

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkSharesDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		_, ok := s.Modules[0].Resources["data.databricks_shares.this"]
		require.True(t, ok, "data.databricks_shares.this has to be there")
		num_shares, _ := strconv.Atoi(s.Modules[0].Outputs["shares"].Value.(string))
		assert.GreaterOrEqual(t, num_shares, 1)
		return nil
	}
}
func TestUcAccDataSourceShares(t *testing.T) {
	unityWorkspaceLevel(t, LegacyStep{
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

		resource "databricks_table" "mytable_2" {
			catalog_name = databricks_catalog.sandbox.id
			schema_name = databricks_schema.things.name
			name = "bar_2"
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
			object {
				name = databricks_table.mytable_2.id
				cdf_enabled = false
				comment = "c"
				data_object_type = "TABLE"
			}						
		}

		data "databricks_shares" "this" {
			depends_on = [databricks_share.myshare]
		}
		output "shares" {
			value = length(data.databricks_shares.this.shares)
		}
		`,
		Check: checkSharesDataSourcePopulated(t),
	})
}
