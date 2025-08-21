package sharing_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkSharesDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_shares_pluginframework.this"]
		require.True(t, ok, "data.databricks_shares_pluginframework.this has to be there")
		num_shares, _ := strconv.Atoi(ds.Primary.Attributes["shares.#"])
		assert.GreaterOrEqual(t, num_shares, 1)
		return nil
	}
}
func TestUcAccDataSourceShares(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}_data"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}

		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.RANDOM}_data"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		resource "databricks_table" "mytable" {
			catalog_name = databricks_catalog.sandbox.id
			schema_name = databricks_schema.things.name
			name = "bar{var.RANDOM}_data"
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
			name = "bar_2{var.RANDOM}_data"
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

		resource "databricks_share_pluginframework" "myshare" {
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

		data "databricks_shares_pluginframework" "this" {
			depends_on = [databricks_share_pluginframework.myshare]
		}
		`,
		Check: checkSharesDataSourcePopulated(t),
	})
}
