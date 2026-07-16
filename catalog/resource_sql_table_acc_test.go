package catalog_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const sqlTableColumnCommentResource = "databricks_sql_table.this"

// sqlTableColumnCommentTemplate renders a catalog + schema + databricks_sql_table
// whose single column `id` carries the given comment. tableType is "VIEW" or a
// table type such as "MANAGED"; typeSpecificBody supplies the type-specific
// attributes (view_definition for views, data_source_format for tables); columnType
// is the column's HCL `type` line (empty for views, which derive it from the query).
func sqlTableColumnCommentTemplate(tableType, typeSpecificBody, columnType, columnComment string) string {
	return fmt.Sprintf(`
	resource "databricks_catalog" "sandbox" {
		name    = "sandbox{var.STICKY_RANDOM}"
		comment = "this catalog is managed by terraform"
	}

	resource "databricks_schema" "things" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "things{var.STICKY_RANDOM}"
		comment      = "this schema is managed by terraform"
	}

	resource "databricks_sql_table" "this" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name  = databricks_schema.things.name
		name         = "t{var.STICKY_RANDOM}"
		table_type   = "%s"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		%s

		column {
			name    = "id"
			%s
			comment = "%s"
		}
	}
	`, tableType, typeSpecificBody, columnType, columnComment)
}

// TestUcAccSqlTableViewColumnCommentUpdate is the regression test for the bug
// where updating a column comment on a VIEW emitted the invalid
// `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a
// PARSE_SYNTAX_ERROR, leaving the change stuck as a perpetual, un-appliable diff.
//
// It asserts the full behaviour the fix promises against a live workspace:
//   - Step 2 changes ONLY the column comment. This apply is exactly the
//     operation that errored on the released provider; here it must succeed,
//     the plan must be an in-place Update (NOT a replacement — so grants and
//     dependent objects are preserved), and the new comment must be persisted
//     server-side (read back into state).
//   - Step 3 re-applies the identical config and asserts a Noop plan, proving
//     the change converges with no perpetual diff.
func TestUcAccSqlTableViewColumnCommentUpdate(t *testing.T) {
	viewBody := `view_definition = "SELECT nullif(1, 2) AS id"`
	acceptance.UnityWorkspaceLevel(t,
		// Step 1: create the view with an initial column comment.
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("VIEW", viewBody, "", "initial comment"),
			Check:    resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "initial comment"),
		},
		// Step 2: change ONLY the column comment. Must be an in-place update,
		// not a replacement, and the value must land server-side.
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("VIEW", viewBody, "", "updated comment"),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(sqlTableColumnCommentResource, plancheck.ResourceActionUpdate),
				},
			},
			Check: resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "updated comment"),
		},
		// Step 3: re-apply the SAME config. The plan must be a no-op, proving
		// the comment change converges (no perpetual diff).
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("VIEW", viewBody, "", "updated comment"),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(sqlTableColumnCommentResource, plancheck.ResourceActionNoop),
				},
			},
			Check: resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "updated comment"),
		},
	)
}

// TestUcAccSqlTableColumnCommentUpdate is the control for the table code path:
// it exercises the `else` branch (ALTER TABLE ... ALTER COLUMN ... COMMENT) to
// ensure the view fix did not regress in-place column-comment updates on tables.
func TestUcAccSqlTableColumnCommentUpdate(t *testing.T) {
	tableBody := `data_source_format = "DELTA"`
	acceptance.UnityWorkspaceLevel(t,
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("MANAGED", tableBody, `type = "int"`, "initial comment"),
			Check:    resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "initial comment"),
		},
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("MANAGED", tableBody, `type = "int"`, "updated comment"),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(sqlTableColumnCommentResource, plancheck.ResourceActionUpdate),
				},
			},
			Check: resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "updated comment"),
		},
		acceptance.Step{
			Template: sqlTableColumnCommentTemplate("MANAGED", tableBody, `type = "int"`, "updated comment"),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(sqlTableColumnCommentResource, plancheck.ResourceActionNoop),
				},
			},
			Check: resource.TestCheckResourceAttr(sqlTableColumnCommentResource, "column.0.comment", "updated comment"),
		},
	)
}
