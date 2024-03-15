package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecretScopeResource(t *testing.T) {
	scope := qa.RandomName("tf-")
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_secret_scope" "my_scope" {
			name = "%s"
		}`, scope),
		Check: resource.ComposeTestCheckFunc(
			// verify local values
			resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
			resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
			resourceCheck("databricks_secret_scope.my_scope",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					w, err := client.WorkspaceClient()
					require.NoError(t, err)
					acls_resp, err := w.Secrets.ListAclsByScope(ctx, id)
					require.NoError(t, err)
					acls := acls_resp.Items
					me, err := w.CurrentUser.Me(ctx)
					require.NoError(t, err)
					assert.Equal(t, 1, len(acls))
					assert.Equal(t, me.UserName, acls[0].Principal)

					err = w.Secrets.DeleteScope(context.Background(), workspace.DeleteScope{Scope: scope})
					assert.NoError(t, err)
					return nil
				}),
		),
		ExpectNonEmptyPlan: true,
	}, step{
		Template: fmt.Sprintf(`
		resource "databricks_secret_scope" "my_scope" {
			name = "%s"
		}`, scope),
		// compose a basic test, checking both remote and local values
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
			resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
		),
	})
}

func TestAccSecretScopeResourceAkvWithSp(t *testing.T) {
	if os.Getenv("ARM_CLIENT_ID") == "" {
		t.Skipf("service principal isn't defined")
	}

	workspaceLevel(t, step{
		Template: `
		resource "databricks_secret_scope" "my_scope" {
			name = "tf-{var.RANDOM}"
			keyvault_metadata {
				resource_id = "{env.TEST_KEY_VAULT_RESOURCE_ID}"
				dns_name    = "{env.TEST_KEY_VAULT_DNS_NAME}"
			}
		}`,
	})
}
