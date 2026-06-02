package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// initUnifiedHostUcacctEnv is the Unity-Catalog account-level counterpart of
// initUnifiedHostAccountEnv: it loads the ucacct environment (which has a
// metastore-assigned workspace) and requires UNIFIED_HOST.
func initUnifiedHostUcacctEnv(t *testing.T) {
	LoadUcacctEnv(t)
	if os.Getenv("UNIFIED_HOST") == "" {
		Skipf(t)("UNIFIED_HOST environment variable is missing")
	}
}

// ==========================================
// databricks_sql_table (catalog/resource_sql_table.go)
//
// Unity Catalog managed table created via a SQL warehouse. Assumes the target
// workspace has a metastore with the default "main" catalog and a warehouse
// (TEST_DEFAULT_WAREHOUSE_ID).
// ==========================================

func createSqlTableWithProviderConfig(t *testing.T, workspaceID, warehouseID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	schemaName := "tf_" + RandomName()
	tableName := "bar"
	step := Step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "` + schemaName + `"
			catalog_name = "main"
			` + pcBlock(workspaceID) + `
		}
		resource "databricks_sql_table" "this" {
			name         = "` + tableName + `"
			catalog_name = "main"
			schema_name  = databricks_schema.this.name
			table_type   = "MANAGED"
			warehouse_id = "` + warehouseID + `"
			column {
				name = "id"
				type = "int"
			}
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_sql_table.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			ti, err := w.Tables.GetByFullName(ctx, id)
			if err != nil {
				return err
			}
			if ti.Name != tableName {
				return fmt.Errorf("expected table name %q, got %q (table may be in the wrong workspace)", tableName, ti.Name)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateSqlTable(t *testing.T) {
	initUnifiedHostUcacctEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	warehouseID := GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID")
	createSqlTableWithProviderConfig(t, workspaceID, warehouseID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestUcAccSqlTable_WorkspaceLevel(t *testing.T) {
	LoadUcwsEnv(t)
	warehouseID := GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID")
	createSqlTableWithProviderConfig(t, currentWorkspaceID(t), warehouseID, nil)
}

// ==========================================
// databricks_table (catalog/resource_table.go) — LEGACY/deprecated.
//
// The sole caller of catalog/update.go's Patch+AddWorkspaceIdHeader path, and it
// had no real (apply) acceptance test before, so both a normal-workspace and a
// unified-host variant are added. Uses the legacy POST /unity-catalog/tables path
// (no SQL warehouse needed).
// ==========================================

func createTableWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	schemaName := "tf_" + RandomName()
	tableName := "test_table"
	step := Step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "` + schemaName + `"
			catalog_name = "main"
			` + pcBlock(workspaceID) + `
		}
		resource "databricks_table" "this" {
			name               = "` + tableName + `"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			data_source_format = "DELTA"
			column {
				name      = "id"
				position  = 0
				type_name = "INT"
				type_text = "int"
			}
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_table.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			ti, err := w.Tables.GetByFullName(ctx, id)
			if err != nil {
				return err
			}
			if ti.Name != tableName {
				return fmt.Errorf("expected table name %q, got %q (table may be in the wrong workspace)", tableName, ti.Name)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateTable(t *testing.T) {
	initUnifiedHostUcacctEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createTableWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestUcAccTable_WorkspaceLevel(t *testing.T) {
	LoadUcwsEnv(t)
	createTableWithProviderConfig(t, currentWorkspaceID(t), nil)
}

// ==========================================
// databricks_provider (sharing/resource_provider.go) — Delta Sharing provider.
//
// Self-contained: uses a dummy recipient profile (reserved example.com endpoint +
// a fake bearer token). For a TOKEN provider the profile is only stored at create
// time — it is not dialed until shares are read — so a placeholder endpoint is
// sufficient to exercise create/read/delete.
// ==========================================

func createSharingProviderWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	name := "tf_" + RandomName() + "_provider"
	step := Step{
		Template: `
		resource "databricks_provider" "this" {
			name                = "` + name + `"
			comment             = "made by terraform"
			authentication_type = "TOKEN"
			recipient_profile_str = jsonencode({
				shareCredentialsVersion = 1
				bearerToken             = "fake-bearer-token-for-acc-test"
				endpoint                = "https://example.com/delta-sharing/"
			})
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_provider.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			pi, err := w.Providers.GetByName(ctx, id)
			if err != nil {
				return err
			}
			if pi.Name != name {
				return fmt.Errorf("expected provider name %q, got %q (provider may be in the wrong workspace)", name, pi.Name)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateSharingProvider(t *testing.T) {
	initUnifiedHostUcacctEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createSharingProviderWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestUcAccSharingProvider_WorkspaceLevel(t *testing.T) {
	LoadUcwsEnv(t)
	createSharingProviderWithProviderConfig(t, currentWorkspaceID(t), nil)
}
