package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccInitialManagePrincipals(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	scopesAPI := NewSecretScopesAPI(client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(scope, "")
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := NewSecretAclsAPI(client)
	acls, err := secretACLAPI.List(scope)
	require.NoError(t, err)

	usersAPI := identity.NewUsersAPI(client)
	me, err := usersAPI.Me()
	require.NoError(t, err)
	assert.Equal(t, 1, len(acls))
	assert.Equal(t, me.UserName, acls[0].Principal)
}

func TestAccInitialManagePrincipalsGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	scopesAPI := NewSecretScopesAPI(client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(scope, "users")
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := NewSecretAclsAPI(client)
	acls, err := secretACLAPI.List(scope)
	require.NoError(t, err)
	assert.Equal(t, 1, len(acls))
	assert.Equal(t, "users", acls[0].Principal)
}

func TestAccSecretScopeResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var secretScope SecretScope

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testSecretScopeResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(ScopeBackendTypeDatabricks)),
					acceptance.ResourceCheck("databricks_secret_scope.my_scope",
						func(client *common.DatabricksClient, id string) error {
							secretACLAPI := NewSecretAclsAPI(client)
							acls, err := secretACLAPI.List(id)
							require.NoError(t, err)

							usersAPI := identity.NewUsersAPI(client)
							me, err := usersAPI.Me()
							require.NoError(t, err)
							assert.Equal(t, 1, len(acls))
							assert.Equal(t, me.UserName, acls[0].Principal)
							return nil
						}),
				),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretScopesAPI(client).Delete(scope)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(ScopeBackendTypeDatabricks)),
				),
			},
		},
	})
}

func testSecretScopeResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := NewSecretScopesAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource token is not cleaned up")
	}
	return nil
}

func testSecretScopeValues(t *testing.T, secretScope *SecretScope, scope string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secretScope.Name == scope)
		assert.True(t, secretScope.BackendType == ScopeBackendTypeDatabricks)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretScopeResourceExists(n string, secretScope *SecretScope, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewSecretScopesAPI(conn).Read(rs.Primary.ID)
		//t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*secretScope = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testAccTokenResource returns an configuration for an Example Widget with the provided name
func testSecretScopeResource(scopeName string) string {
	return fmt.Sprintf(`
		resource "databricks_secret_scope" "my_scope" {
			name = "%s"
		}
		`, scopeName)
}
