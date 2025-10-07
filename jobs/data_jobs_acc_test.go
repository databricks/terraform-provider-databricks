package jobs_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
		ExpectError: regexp.MustCompile(`failed to parse workspace_id.*invalid syntax`),
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
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
		}`,
	}, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
			provider_config {
				workspace_id = "1142582526922259"
			}
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_jobs.all"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			id := r.Primary.Attributes["provider_config.0.workspace_id"]
			if id != "1142582526922259" {
				return fmt.Errorf("wrong workspace_id found: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
