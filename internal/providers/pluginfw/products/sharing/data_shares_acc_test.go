package sharing_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkSharesDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_shares.this"]
		require.True(t, ok, "data.databricks_shares.this has to be there")
		num_shares, _ := strconv.Atoi(ds.Primary.Attributes["shares.#"])
		assert.GreaterOrEqual(t, num_shares, 1)
		return nil
	}
}
func TestUcAccDataSourceShares(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
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

		resource "databricks_share" "myshare" {
			name = "{var.RANDOM}-terraform-delta-share"
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

		data "databricks_shares" "this" {
			depends_on = [databricks_share.myshare]
		}
		`,
		Check: checkSharesDataSourcePopulated(t),
	})
}

func dataSourceSharesTemplate(provider_config string) string {
	return fmt.Sprintf(`
	resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-share-config"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
	}
	data "databricks_shares" "this" {
		%s
		depends_on = [databricks_share.myshare]
	}
`, provider_config)
}

func TestAccSharesData_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceSharesTemplate(`
			provider_config = {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(
			`Attribute provider_config\.workspace_id\s+workspace_id must be a valid integer`,
		),
		PlanOnly: true,
	})
}

func TestAccSharesData_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceSharesTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client.*workspace_id mismatch` +
				`.*please check the workspace_id provided in ` +
				`provider_config`,
		),
	})
}

func TestAccSharesData_ProviderConfig_MismatchReapply(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceSharesTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client.*workspace_id mismatch` +
				`.*please check the workspace_id provided in ` +
				`provider_config`,
		),
	}, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceSharesTemplate(fmt.Sprintf(`
			provider_config = {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	})
}
