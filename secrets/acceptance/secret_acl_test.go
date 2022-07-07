package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecretAclResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: acceptance.EnvironmentTemplate(t, `
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
					}`),
				Check: func(s *terraform.State) error {
					client := common.CommonEnvironmentClient()

					ctx := context.Background()
					usersAPI := scim.NewUsersAPI(ctx, client)
					me, err := usersAPI.Me()
					require.NoError(t, err)

					secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
					scope := s.RootModule().Resources["databricks_secret_scope.app"].Primary.ID
					acls, err := secretACLAPI.List(scope)
					require.NoError(t, err)
					assert.Equal(t, 2, len(acls))
					m := map[string]string{}
					for _, acl := range acls {
						m[acl.Principal] = string(acl.Permission)
					}

					group := s.RootModule().Resources["databricks_group.ds"].Primary.Attributes["display_name"]
					require.Contains(t, m, group)
					assert.Equal(t, "READ", m[group])
					assert.Equal(t, "MANAGE", m[me.UserName])
					return nil
				},
			},
		},
	})
}

func TestAccSecretAclResourceDefaultPrincipal(t *testing.T) {
	t.Parallel()
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: acceptance.EnvironmentTemplate(t, `
					resource "databricks_secret_scope" "app" {
						name = "app-{var.RANDOM}"
						initial_manage_principal = "users"
					}
					resource "databricks_secret_acl" "ds_can_read_app" {
						principal = "users"
						permission = "READ"
						scope = databricks_secret_scope.app.name
					}`),
				Check: acceptance.ResourceCheck("databricks_secret_scope.app",
					func(ctx context.Context, client *common.DatabricksClient, id string) error {
						secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
						acls, err := secretACLAPI.List(id)
						require.NoError(t, err)
						assert.Equal(t, 1, len(acls))
						assert.Equal(t, "users", acls[0].Principal)
						assert.Equal(t, "READ", string(acls[0].Permission))
						return nil
					}),
			},
		},
	})
}
