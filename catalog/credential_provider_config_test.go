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

func credentialProviderConfigTemplate(name string, providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_credential" "this" {
		name    = "%s"
		purpose = "SERVICE"
		aws_iam_role {
			role_arn = "arn:aws:iam::123456789012:role/tf-test"
		}
		%s
	}
	`, name, providerConfig)
}

func TestAccCredential_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate("tf-test-cred-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccCredential_ProviderConfig_Required(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate("tf-test-cred-{var.STICKY_RANDOM}", `
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccCredential_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate("tf-test-cred-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccCredential_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate("tf-test-cred-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}

func TestAccCredential_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	name := "tf-test-cred-{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, ""),
	}, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_credential.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccCredential_ProviderConfig_Recreate(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	name := "tf-test-cred-{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, ""),
	}, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, `
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}

func TestAccCredential_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	name := "tf-test-cred-{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, ""),
	}, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: credentialProviderConfigTemplate(name, ""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_credential.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}
