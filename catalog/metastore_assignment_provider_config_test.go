package catalog_test

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

func metastoreAssignmentProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_metastore_assignment" "this" {
		workspace_id = 123456789
		metastore_id = "fake-metastore-id"
		%s
	}
	`, providerConfig)
}

func TestAccMetastoreAssignment_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreAssignmentProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreAssignment_ProviderConfig_Required(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreAssignmentProviderConfigTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreAssignment_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreAssignmentProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreAssignment_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreAssignmentProviderConfigTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreAssignment_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	currentWorkspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(currentWorkspaceID, 10)

	currentMetastore, err := w.Metastores.Current(ctx)
	require.NoError(t, err)

	template := func(providerConfig string) string {
		return fmt.Sprintf(`
		resource "databricks_metastore_assignment" "this" {
			workspace_id = %d
			metastore_id = "%s"
			%s
		}
		`, currentWorkspaceID, currentMetastore.MetastoreId, providerConfig)
	}
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: template(""),
	}, acceptance.Step{
		Template: template(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_metastore_assignment.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccMetastoreAssignment_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	currentWorkspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(currentWorkspaceID, 10)

	currentMetastore, err := w.Metastores.Current(ctx)
	require.NoError(t, err)

	template := func(providerConfig string) string {
		return fmt.Sprintf(`
		resource "databricks_metastore_assignment" "this" {
			workspace_id = %d
			metastore_id = "%s"
			%s
		}
		`, currentWorkspaceID, currentMetastore.MetastoreId, providerConfig)
	}
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: template(""),
	}, acceptance.Step{
		Template: template(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: template(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_metastore_assignment.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}
