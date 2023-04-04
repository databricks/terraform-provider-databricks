package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/secrets"

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
					secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
					acls, err := secretACLAPI.List(id)
					require.NoError(t, err)

					w, err := client.WorkspaceClient()
					require.NoError(t, err)

					me, err := w.CurrentUser.Me(ctx)
					require.NoError(t, err)
					assert.Equal(t, 1, len(acls))
					assert.Equal(t, me.UserName, acls[0].Principal)

					err = secrets.NewSecretScopesAPI(context.Background(), client).Delete(scope)
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
