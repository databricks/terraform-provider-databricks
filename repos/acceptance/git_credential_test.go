package acceptance

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/repos"
	"github.com/stretchr/testify/assert"
)

func TestAccGitCredentials(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_git_credential" "this" {
				git_username = "test"
				git_provider = "gitHub"
				personal_access_token = "test_token"
				force = true
			}`,
			Check: acceptance.ResourceCheck("databricks_git_credential.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					creds, err := repos.NewGitCredentialsAPI(ctx, client).List()
					assert.NoError(t, err)
					assert.Len(t, creds, 1)
					assert.Equal(t, creds[0].UserName, "test")
					return nil
				}),
		},
	})
}
