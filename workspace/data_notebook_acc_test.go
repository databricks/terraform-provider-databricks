package workspace_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func notebookDataTemplate(provider_config string) string {
	return fmt.Sprintf(`
		resource "databricks_notebook" "this" {
			source = "{var.CWD}/acceptance/testdata/acc-test-update1.dbc"
			path = "/Shared/provider-test/jupiter_{var.STICKY_RANDOM}"
		}
		data "databricks_notebook" "this" {
			%s
			path = databricks_notebook.this.path
			format = "DBC"
		}
	`, provider_config)
}

func TestAccNotebookData_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccNotebookData_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccNotebookData_ProviderConfig_MismatchedReapply(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		PlanOnly:    true,
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	}, acceptance.Step{
		Template: notebookDataTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	})
}

func TestAccNotebookData_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccNotebookData_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccNotebookData_InState(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notebookDataTemplate(""),
	}, acceptance.Step{
		Template: fmt.Sprintf(notebookDataTemplate(`
			provider_config {
				workspace_id = "%s"
			}
		`), workspaceIDStr),
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_notebook.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			id := r.Primary.Attributes["provider_config.0.workspace_id"]
			if id != workspaceIDStr {
				return fmt.Errorf("wrong workspace_id found: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
