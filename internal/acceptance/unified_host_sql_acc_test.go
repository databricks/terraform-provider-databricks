package acceptance

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ==========================================
// databricks_sql_query (sql/resource_sql_query.go)
//
// Legacy DBSQL query. Needs a SQL-warehouse data source id
// (TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID — the data-source id, NOT the warehouse id).
// ==========================================

func createSqlQueryWithProviderConfig(t *testing.T, workspaceID, dataSourceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	queryName := "tf-" + RandomName() + "-query"
	step := Step{
		Template: `
		resource "databricks_sql_query" "this" {
			data_source_id = "` + dataSourceID + `"
			name           = "` + queryName + `"
			query          = "SELECT 1"
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_sql_query.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			q, err := w.QueriesLegacy.GetByQueryId(ctx, id)
			if err != nil {
				return err
			}
			if q.Name != queryName {
				return fmt.Errorf("expected query name %q, got %q (query may be in the wrong workspace)", queryName, q.Name)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateSqlQuery(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	dataSourceID := GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID")
	createSqlQueryWithProviderConfig(t, workspaceID, dataSourceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

// ==========================================
// databricks_sql_visualization (sql/resource_sql_visualization.go)
//
// Depends on a query. The resource ID is composite "query_id/visualization_id".
// There is no direct SDK GET for a visualization, so it is verified by scanning
// the parent query's Visualizations.
// ==========================================

func createSqlVisualizationWithProviderConfig(t *testing.T, workspaceID, dataSourceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	queryName := "tf-" + RandomName() + "-query"
	vizName := "tf-" + RandomName() + "-viz"
	step := Step{
		Template: `
		resource "databricks_sql_query" "this" {
			data_source_id = "` + dataSourceID + `"
			name           = "` + queryName + `"
			query          = "SELECT 1"
			` + pcBlock(workspaceID) + `
		}
		resource "databricks_sql_visualization" "this" {
			query_id = databricks_sql_query.this.id
			type     = "table"
			name     = "` + vizName + `"
			options  = jsonencode({})
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_sql_visualization.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			parts := strings.SplitN(id, "/", 2)
			if len(parts) != 2 {
				return fmt.Errorf("unexpected visualization id %q, want query_id/visualization_id", id)
			}
			q, err := w.QueriesLegacy.GetByQueryId(ctx, parts[0])
			if err != nil {
				return err
			}
			for _, v := range q.Visualizations {
				if v.Id == parts[1] {
					return nil
				}
			}
			return fmt.Errorf("visualization %q not found on query %q in workspace %s", parts[1], parts[0], workspaceID)
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateSqlVisualization(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	dataSourceID := GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID")
	createSqlVisualizationWithProviderConfig(t, workspaceID, dataSourceID, unifiedHostProviderFactories(unifiedHost, accountID))
}
