package repos_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccRepo_WithGitCredentialID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_git_credential" "this" {
			git_username          = "test"
			git_provider          = "gitHub"
			personal_access_token = "fake-personal-access-token"
			force                 = true
		}

		resource "databricks_repo" "this" {
			url               = "https://github.com/databricks/databricks-sdk-go"
			git_credential_id = databricks_git_credential.this.id
		}
		`,
	})
}
