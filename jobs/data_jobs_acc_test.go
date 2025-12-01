package jobs_test

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

func TestAccDataSourcesJob_InvalidID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
				workspace_id = "invalid"
			}
		}`,
<<<<<<< HEAD
		ExpectError: regexp.MustCompile(`workspace_id must be a valid integer`),
=======
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
>>>>>>> d3264a686497fd3bff26572b29e7db25ef11673c
		PlanOnly:    true,
	})
}

func TestAccDataSourcesJob_MismatchedID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
				workspace_id = "123"
			}
		}`,
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccDataSourcesJob_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
				workspace_id = ""
			}
		}`,
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
	})
}

func TestAccDataSourcesJob_EmptyBlock(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
			}
		}`,
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
	})
}

func TestAccDataSourcesJob(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
		}`,
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
				workspace_id = "%s"
			}
		}`, workspaceIDStr),
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_jobs.all"]
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
