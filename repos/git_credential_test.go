package repos_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccGitCredentials(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_git_credential" "this" {
			git_username = "test"
			git_provider = "gitHub"
			personal_access_token = "test_token"
			force = true
		}`,
		Check: acceptance.ResourceCheck("databricks_git_credential.this",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				assert.NoError(t, err)
				// Look up this credential by ID, as other tests may create credentials in parallel
				credID, err := strconv.ParseInt(id, 10, 64)
				assert.NoError(t, err)
				cred, err := w.GitCredentials.Get(ctx, workspace.GetCredentialsRequest{CredentialId: credID})
				assert.NoError(t, err)
				assert.Equal(t, "test", cred.GitUsername)
				return nil
			}),
	})
}

func gitCredentialWriteOnlyTemplate(token string, version int) string {
	return fmt.Sprintf(`resource "databricks_git_credential" "this" {
		git_username                     = "test"
		git_provider                     = "gitHub"
		personal_access_token_wo         = "%s"
		personal_access_token_wo_version = %d
		force                            = true
	}`, token, version)
}

// checkGitCredentialTokenNotInState verifies that neither the write-only token nor the
// token sourced from the environment ends up in state.
func checkGitCredentialTokenNotInState(version string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources["databricks_git_credential.this"]
		if !ok {
			return fmt.Errorf("databricks_git_credential.this not found in state")
		}
		attrs := r.Primary.Attributes
		if v := attrs["personal_access_token"]; v != "" {
			return fmt.Errorf("personal_access_token must not be stored in state")
		}
		if attrs["personal_access_token_wo"] != "" {
			return fmt.Errorf("personal_access_token_wo must not be stored in state")
		}
		if attrs["personal_access_token_wo_version"] != version {
			return fmt.Errorf("expected personal_access_token_wo_version %s, got %s",
				version, attrs["personal_access_token_wo_version"])
		}
		return nil
	}
}

func TestAccGitCredentialsWriteOnly(t *testing.T) {
	// The provider runs in-process, so this token is picked up by the `personal_access_token`
	// default. It must be ignored when `personal_access_token_wo` is configured.
	// t.Setenv can't be used, as acceptance tests run in parallel.
	old, hadOld := os.LookupEnv("GITHUB_TOKEN")
	t.Cleanup(func() {
		if hadOld {
			os.Setenv("GITHUB_TOKEN", old)
		} else {
			os.Unsetenv("GITHUB_TOKEN")
		}
	})
	acceptance.WorkspaceLevel(t, acceptance.Step{
		PreConfig: func() { os.Setenv("GITHUB_TOKEN", "token_from_env") },
		Template:  gitCredentialWriteOnlyTemplate("write_only_token_v1", 1),
		Check:     checkGitCredentialTokenNotInState("1"),
	}, acceptance.Step{
		// Bumping the version sends the new token without storing it
		Template: gitCredentialWriteOnlyTemplate("write_only_token_v2", 2),
		Check:    checkGitCredentialTokenNotInState("2"),
	})
}
