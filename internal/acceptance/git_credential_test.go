package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestAccGitCredentials(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_git_credential" "this" {
			git_username = "test"
			git_provider = "gitHub"
			personal_access_token = "test_token"
			force = true
		}`,
		Check: resourceCheck("databricks_git_credential.this",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				assert.NoError(t, err)
				creds, err := w.GitCredentials.ListAll(ctx)
				assert.NoError(t, err)
				assert.Len(t, creds, 1)
				assert.Equal(t, creds[0].GitUsername, "test")
				return nil
			}),
	})
}
