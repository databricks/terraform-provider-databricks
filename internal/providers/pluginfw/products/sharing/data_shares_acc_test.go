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

const sharesDataTemplate = `
	resource "databricks_catalog" "sandbox" {
		name         = "sandbox{var.STICKY_RANDOM}"
		comment      = "this catalog is managed by terraform"
		properties = {
			purpose = "testing"
		}
	}

	resource "databricks_schema" "things" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "things{var.STICKY_RANDOM}"
		comment      = "this database is managed by terraform"
		properties = {
			kind = "various"
		}
	}

	resource "databricks_sql_table" "mytable" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar"
		table_type = "MANAGED"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"

		column {
			name = "id"
			type = "int"
		}
	}

	resource "databricks_sql_table" "mytable_2" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar_2"
		table_type = "MANAGED"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"

		column {
			name = "id"
			type = "int"
		}
	}

	resource "databricks_share_pluginframework" "myshare" {
		name = "{var.STICKY_RANDOM}-terraform-delta-share"
		object {
			name = databricks_sql_table.mytable.id
			comment = "c"
			data_object_type = "TABLE"
			history_data_sharing_status = "ENABLED"
		}
		object {
			name = databricks_sql_table.mytable_2.id
			comment = "c"
			data_object_type = "TABLE"
			history_data_sharing_status = "ENABLED"
		}
	}
`

func TestUcAccDataSourceShares(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: sharesDataTemplate + `
		data "databricks_shares_pluginframework" "this" {
			depends_on = [databricks_share_pluginframework.myshare]
		}
		`,
		Check: checkSharesDataSourcePopulated(t),
	}, acceptance.Step{
		Template: sharesDataTemplate + `
		data "databricks_shares_pluginframework" "this" {
			depends_on = [databricks_share_pluginframework.myshare]
			provider_config = {
				workspace_id = "{env.THIS_WORKSPACE_ID}"
			}
		}
		`,
		Check: checkSharesDataSourcePopulated(t),
	})
}
