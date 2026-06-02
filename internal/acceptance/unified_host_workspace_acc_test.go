package acceptance

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	databricks "github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// This file (and its unified_host_*_acc_test.go siblings) adds apply-based
// acceptance coverage for the SDKv2 "raw HTTP" resources whose CRUD calls were
// switched to inject the X-Databricks-Workspace-Id routing header per-callsite
// (PR #5759). Each resource is exercised through a provider_config { workspace_id }
// block so the create/read round-trip actually traverses the header path:
//
//   - TestMwsAccUnifiedHost*  : account-level provider pointed at the UNIFIED host;
//     the only way the request can reach the right workspace is the routing header.
//   - TestAcc*_WorkspaceLevel / TestUcAcc*_WorkspaceLevel : workspace-level provider on
//     a normal host — one per resource (except databricks_permission_assignment, which
//     needs an account principal a workspace-level provider cannot create; see below).
//
// The existing *_provider_config_test.go tests do not count for this purpose —
// they are PlanOnly and never call the API.
//
// Intentionally NOT covered here, with rationale:
//   - databricks_sql_dashboard / databricks_sql_widget: legacy dashboard creation
//     is disabled server-side ("Legacy dashboards can no longer be created via the
//     API", see sql/sql_dashboard_test.go), so an apply test cannot succeed.
//   - databricks_mount (+ aws_s3/azure_blob/adls_gen1/adls_gen2 variants): already
//     covered by TestAccCreateDatabricksMountWithProviderConfig in
//     storage/mounts_acc_test.go; mounts execute Python on a cluster (no SDK GET)
//     and need live object-store credentials.
//   - databricks_sql_permissions: requires a table-ACL-enabled cluster; the command
//     executor path it shares is already exercised by the mount test above.
//   - databricks_obo_token: requires the on-behalf-of service principal to be a
//     workspace admin; reliable cross-level setup through an account/unified provider
//     is not available in the standard acc env. Covered by TestAccAwsOboTokenResource.
//   - databricks_sql_global_config: a workspace-wide singleton that mutates shared
//     state and needs the cross-test lock from sql/sql_global_config_test.go; a
//     unified-host variant would risk clobbering parallel runs. Covered by
//     TestAccSQLGlobalConfig.

// pcBlock returns a provider_config block pinning a resource to workspaceID.
// Inject it into every resource AND data-source block in a template, because on a
// unified host every workspace-scoped request needs the routing header.
func pcBlock(workspaceID string) string {
	return fmt.Sprintf("provider_config {\n\t\t\t\tworkspace_id = \"%s\"\n\t\t\t}", workspaceID)
}

// currentWorkspaceID returns the numeric ID of the workspace the default
// (workspace-level) provider points at. Used by the *_WorkspaceLevel variants to
// pin provider_config.workspace_id to the current workspace, so no extra env var
// is required to run them.
func currentWorkspaceID(t *testing.T) string {
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	id, err := w.CurrentWorkspaceID(ctx)
	if err != nil {
		t.Fatalf("failed to get current workspace ID: %s", err)
	}
	return strconv.FormatInt(id, 10)
}

// ==========================================
// databricks_token (tokens/resource_token.go)
// ==========================================

func createTokenWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	comment := "tf-" + RandomName() + "-token"
	step := Step{
		Template: `
		resource "databricks_token" "this" {
			comment          = "` + comment + `"
			lifetime_seconds = 3600
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_token.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			tokens, err := w.Tokens.ListAll(ctx)
			if err != nil {
				return err
			}
			for _, tok := range tokens {
				if tok.TokenId == id {
					if tok.Comment != comment {
						return fmt.Errorf("expected token comment %q, got %q", comment, tok.Comment)
					}
					return nil
				}
			}
			return fmt.Errorf("token %s not found in workspace %s", id, workspaceID)
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateToken(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createTokenWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestAccToken_WorkspaceLevel(t *testing.T) {
	LoadWorkspaceEnv(t)
	createTokenWithProviderConfig(t, currentWorkspaceID(t), nil)
}

// ==========================================
// databricks_dbfs_file (storage/dbfs.go)
// ==========================================

func createDbfsFileWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	path := "/tmp/tf-acc-" + RandomName()
	step := Step{
		Template: `
		resource "databricks_dbfs_file" "this" {
			path           = "` + path + `"
			content_base64 = base64encode("hello unified host")
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: ResourceCheck("databricks_dbfs_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			info, err := w.Dbfs.GetStatusByPath(ctx, id)
			if err != nil {
				return err
			}
			if info.IsDir {
				return fmt.Errorf("expected a file at %q, got a directory", id)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateDbfsFile(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createDbfsFileWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestAccDbfsFile_WorkspaceLevel(t *testing.T) {
	LoadWorkspaceEnv(t)
	createDbfsFileWithProviderConfig(t, currentWorkspaceID(t), nil)
}

// ==========================================
// databricks_permission_assignment (access/resource_permission_assignment.go)
//
// Assigns an account-level principal (a service principal created via the base
// account/unified provider) into the workspace targeted by provider_config. There
// is no workspace-client SDK service for permission assignments, so the check is a
// state assertion — the create+read round-trip already traverses the routing header
// (Read re-fetches the assignment via the workspace-scoped client), so a successful
// apply proves the principal landed in the right workspace. Only the unified-host
// variant is added: a workspace-level provider cannot create an account principal
// to assign, so a "normal workspace" variant has no self-contained form.
// ==========================================

func createPermissionAssignmentWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	spName := "tf-" + RandomName() + "-sp"
	step := Step{
		Template: `
		resource "databricks_service_principal" "this" {
			display_name = "` + spName + `"
		}
		resource "databricks_permission_assignment" "this" {
			principal_id = databricks_service_principal.this.id
			permissions  = ["USER"]
			` + pcBlock(workspaceID) + `
		}
		`,
		Check: resource.TestCheckResourceAttrSet("databricks_permission_assignment.this", "id"),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreatePermissionAssignment(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createPermissionAssignmentWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}
