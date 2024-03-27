package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecretAclResource(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_group" "ds" {
			display_name = "data-scientists-{var.RANDOM}"
		}
		resource "databricks_secret_scope" "app" {
			name = "app-{var.RANDOM}"
		}
		resource "databricks_secret_acl" "ds_can_read_app" {
			principal = databricks_group.ds.display_name
			permission = "READ"
			scope = databricks_secret_scope.app.name
		}`,
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())

			ctx := context.Background()
			me, err := w.CurrentUser.Me(ctx)
			require.NoError(t, err)

			scope := s.RootModule().Resources["databricks_secret_scope.app"].Primary.ID
			acls, err := w.Secrets.ListAclsByScope(ctx, scope)
			require.NoError(t, err)
			assert.Equal(t, 2, len(acls.Items))
			m := map[string]string{}
			for _, acl := range acls.Items {
				m[acl.Principal] = string(acl.Permission)
			}

			group := s.RootModule().Resources["databricks_group.ds"].Primary.Attributes["display_name"]
			require.Contains(t, m, group)
			assert.Equal(t, "READ", m[group])
			assert.Equal(t, "MANAGE", m[me.UserName])
			return nil
		},
	})
}

func TestAccSecretAclResourceDefaultPrincipal(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_secret_scope" "app" {
			name = "app-{var.RANDOM}"
			initial_manage_principal = "users"
		}
		resource "databricks_secret_acl" "ds_can_read_app" {
			principal = "users"
			permission = "READ"
			scope = databricks_secret_scope.app.name
		}`,
		Check: resourceCheck("databricks_secret_scope.app",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				acls_resp, err := w.Secrets.ListAclsByScope(ctx, id)
				require.NoError(t, err)
				acls := acls_resp.Items
				assert.Equal(t, 1, len(acls))
				assert.Equal(t, "users", acls[0].Principal)
				assert.Equal(t, "READ", string(acls[0].Permission))
				return nil
			}),
	})
}
