package catalog

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestResourceSqlTableCreateStatement_IfNotExists(t *testing.T) {
	for _, tc := range []struct {
		name        string
		tableType   string
		ifNotExists bool
		expected    string
	}{
		{"foreign", "FOREIGN", true, "CREATE TABLE IF NOT EXISTS `main`.`test`.`existing`\nUSING DELTA;"},
		{"external", "EXTERNAL", true, "CREATE TABLE IF NOT EXISTS `main`.`test`.`existing`\nUSING DELTA;"},
		{"managed", "MANAGED", true, "CREATE TABLE IF NOT EXISTS `main`.`test`.`existing`\nUSING DELTA;"},
		{"view", "VIEW", true, "CREATE VIEW IF NOT EXISTS `main`.`test`.`existing`\nAS SELECT 1 AS id;"},
		{"foreign_default", "FOREIGN", false, "CREATE TABLE `main`.`test`.`existing`\nUSING DELTA;"},
		{"managed_default", "MANAGED", false, "CREATE OR REPLACE TABLE `main`.`test`.`existing`\nUSING DELTA;"},
		{"view_default", "VIEW", false, "CREATE VIEW `main`.`test`.`existing`\nAS SELECT 1 AS id;"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ti := &SqlTableInfo{
				Name: "existing", CatalogName: "main", SchemaName: "test",
				TableType: tc.tableType, DataSourceFormat: "DELTA", IfNotExists: tc.ifNotExists,
				ViewDefinition: "SELECT 1 AS id",
			}
			assert.Equal(t, tc.expected, ti.buildTableCreateStatement())
		})
	}
}

func TestResourceSqlTableCreate_IfNotExistsRetryAfterOwnerFailure(t *testing.T) {
	// Mock only the remote SQL and Tables APIs; both attempts execute the real resource Create.
	qa.MockWorkspaceApply(t, func(w *mocks.MockWorkspaceClient) {
		w.GetMockStatementExecutionAPI().EXPECT().ExecuteStatement(mock.Anything, sql.ExecuteStatementRequest{
			Statement:   "CREATE TABLE IF NOT EXISTS `main`.`test`.`existing`\nUSING DELTA;",
			WarehouseId: "warehouse", WaitTimeout: "50s", OnWaitTimeout: sql.ExecuteStatementRequestOnWaitTimeoutCancel,
		}).Return(&sql.StatementResponse{Status: &sql.StatementStatus{State: sql.StatementStateSucceeded}}, nil).Twice()
		update := catalog.UpdateTableRequest{FullName: "main.test.existing", Owner: "account users"}
		w.GetMockTablesAPI().EXPECT().Update(mock.Anything, update).
			Return(errors.New("owner update failed")).Once()
		w.GetMockTablesAPI().EXPECT().Update(mock.Anything, update).
			Return(nil).Once()
	}, func(ctx context.Context, client *common.DatabricksClient) {
		client.WithCommandMock(func(string) common.CommandResults { return common.CommandResults{} })
		resource := ResourceSqlTable()
		config := map[string]any{
			"name": "existing", "catalog_name": "main", "schema_name": "test",
			"table_type": "FOREIGN", "data_source_format": "DELTA", "warehouse_id": "warehouse",
			"if_not_exists": true, "owner": "account users",
		}
		first := schema.TestResourceDataRaw(t, resource.Schema, config)
		require.EqualError(t, resource.Create(ctx, first, client), "owner update failed")
		assert.Empty(t, first.Id())
		retry := schema.TestResourceDataRaw(t, resource.Schema, config)
		require.NoError(t, resource.Create(ctx, retry, client))
		assert.Equal(t, "main.test.existing", retry.Id())
		assert.Equal(t, true, retry.Get("if_not_exists"))
	})
}

func TestResourceSqlTableCreate_IfNotExistsSqlFailure(t *testing.T) {
	for _, status := range []sql.StatementState{sql.StatementStateFailed, sql.StatementStateCanceled} {
		t.Run(string(status), func(t *testing.T) {
			qa.MockWorkspaceApply(t, func(w *mocks.MockWorkspaceClient) {
				w.GetMockStatementExecutionAPI().EXPECT().ExecuteStatement(mock.Anything, mock.Anything).
					Return(&sql.StatementResponse{Status: &sql.StatementStatus{State: status}}, nil)
			}, func(ctx context.Context, client *common.DatabricksClient) {
				client.WithCommandMock(func(string) common.CommandResults { return common.CommandResults{} })
				resource := ResourceSqlTable()
				data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
					"name": "existing", "catalog_name": "main", "schema_name": "test",
					"table_type": "FOREIGN", "warehouse_id": "warehouse", "if_not_exists": true,
				})
				err := resource.Create(ctx, data, client)
				assert.EqualError(t, err, "statement failed to execute: "+string(status))
				assert.Empty(t, data.Id())
			})
		})
	}
}

func TestResourceSqlTableRead_PreservesIfNotExists(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceSqlTable(), Read: true, ID: "main.test.existing",
		HCL: `name = "existing"
catalog_name = "main"
schema_name = "test"
table_type = "FOREIGN"
if_not_exists = true`,
		Fixtures: []qa.HTTPFixture{
			{
				Method: "GET", Resource: "/api/2.1/unity-catalog/tables/main.test.existing",
				Response: SqlTableInfo{Name: "existing", CatalogName: "main", SchemaName: "test", TableType: "FOREIGN"},
			},
			{
				Method: "GET", Resource: "/api/2.1/unity-catalog/tables/main.test.existing?",
				Response: catalog.TableInfo{},
			},
		},
	}.ApplyAndExpectData(t, map[string]any{"id": "main.test.existing", "if_not_exists": true})
}
