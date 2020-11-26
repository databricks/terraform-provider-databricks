package acceptance

import (
	"context"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSecretAclResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: qa.EnvironmentTemplate(t, `
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
					usersAPI := identity.NewUsersAPI(ctx, client)
					me, err := usersAPI.Me()
					require.NoError(t, err)

					secretACLAPI := NewSecretAclsAPI(client)
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
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: qa.EnvironmentTemplate(t, `
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
					func(client *common.DatabricksClient, id string) error {
						secretACLAPI := NewSecretAclsAPI(client)
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
