package sql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestUcAccResourceSqlTable_IfNotExists(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	warehouseID := acceptance.GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID")
	ctx := context.Background()
	execute := func(statement string) (*sql.StatementResponse, error) {
		response, err := w.StatementExecution.ExecuteStatement(ctx, sql.ExecuteStatementRequest{
			Statement: statement, WarehouseId: warehouseID, WaitTimeout: "50s",
			OnWaitTimeout: sql.ExecuteStatementRequestOnWaitTimeoutCancel,
		})
		if err != nil {
			return nil, err
		}
		if response.Status.State != sql.StatementStateSucceeded {
			return nil, fmt.Errorf("SQL statement failed: %+v", response.Status)
		}
		return response, nil
	}
	const schemaTemplate = `
	resource "databricks_schema" "this" {
		name          = "{var.STICKY_RANDOM}"
		catalog_name  = "main"
		force_destroy = true
	}`
	const tableTemplate = schemaTemplate + `
	resource "databricks_sql_table" "this" {
		name               = "existing"
		catalog_name       = "main"
		schema_name        = databricks_schema.this.name
		table_type         = "MANAGED"
		data_source_format = "DELTA"
		warehouse_id       = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		if_not_exists      = true
	}`
	var quotedName, tableID string
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: schemaTemplate,
		Check: func(state *terraform.State) error {
			schemaName := state.RootModule().Resources["databricks_schema.this"].Primary.Attributes["name"]
			quotedName = fmt.Sprintf("`main`.`%s`.`existing`", schemaName)
			for _, statement := range []string{
				"CREATE TABLE " + quotedName + " (id INT) USING DELTA",
				"INSERT INTO " + quotedName + " VALUES (7)",
			} {
				if _, err := execute(statement); err != nil {
					return err
				}
			}
			info, err := w.Tables.GetByFullName(ctx, "main."+schemaName+".existing")
			if err != nil {
				return err
			}
			tableID = info.TableId
			return nil
		},
	}, acceptance.Step{
		Template: tableTemplate,
		Check: func(state *terraform.State) error {
			attributes := state.RootModule().Resources["databricks_sql_table.this"].Primary.Attributes
			if attributes["table_id"] != tableID || attributes["if_not_exists"] != "true" {
				return fmt.Errorf("existing table or creation option was not preserved in state")
			}
			response, err := execute("SELECT id FROM " + quotedName)
			if err != nil {
				return err
			}
			if response.Result == nil || len(response.Result.DataArray) != 1 ||
				len(response.Result.DataArray[0]) != 1 || response.Result.DataArray[0][0] != "7" {
				return fmt.Errorf("existing table data was not preserved")
			}
			return nil
		},
	}, acceptance.Step{
		Template: tableTemplate,
		PlanOnly: true,
	})
}
