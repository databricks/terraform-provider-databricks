package jobs_test

import (
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

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
				workspace_id = ""
			}
		}`,
	}, acceptance.Step{
		Template: `
		data "databricks_jobs" "all" {
			key = "id"
		}`,
	})
}

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
