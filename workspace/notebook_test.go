package workspace_test

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

func TestAccNotebookResourceScalability(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccNotebookResourceDbcUpdate(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update2.dbc"
			path = "/Shared/provider-test/dbc_{var.STICKY_RANDOM}"
		}`,
	})
}

func TestAccNotebookResourceJupiterUpdate(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update2.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}`,
	})
}

func notebookTemplate(provider_config string) string {
	return fmt.Sprintf(`
		resource "databricks_notebook" "this" {
			%s
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.ipynb"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}
	`, provider_config)
}

func TestAccNotebook_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccNotebook_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}

func TestAccNotebook_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccNotebook_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccNotebook_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(""),
	}, acceptance.Step{
		Template: notebookTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_notebook.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccNotebook_ProviderConfig_Recreate(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(""),
	}, acceptance.Step{
		Template: notebookTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: notebookTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PostApplyPreRefresh: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_notebook.this", plancheck.ResourceActionDestroyBeforeCreate),
			},
		},
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	})
}

func TestAccNotebook_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookTemplate(""),
	}, acceptance.Step{
		Template: notebookTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: notebookTemplate(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_notebook.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}
