package access_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func sqlPermissionsProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_sql_permissions" "this" {
		database = "default"
		%s
	}
	`, providerConfig)
}

func TestAccSqlPermissions_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(""),
	}, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_sql_permissions.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccSqlPermissions_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(""),
	}, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_sql_permissions.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}
