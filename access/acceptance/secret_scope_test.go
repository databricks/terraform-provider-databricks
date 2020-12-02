package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRemoveScopes(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Cleaning up tests only from IDE")
	}
	client := common.CommonEnvironmentClient()
	scopesAPI := NewSecretScopesAPI(context.Background(), client)
	scopeList, err := scopesAPI.List()
	require.NoError(t, err)
	for _, scope := range scopeList {
		assert.NoError(t, scopesAPI.Delete(scope.Name))
	}
}

func TestAzureAccKeyVaultSimple(t *testing.T) {
	resourceID := qa.GetEnvOrSkipTest(t, "TEST_KEY_VAULT_RESOURCE_ID")
	DNSName := qa.GetEnvOrSkipTest(t, "TEST_KEY_VAULT_DNS_NAME")

	client := common.CommonEnvironmentClient()
	if client.AzureAuth.IsClientSecretSet() {
		t.Skip("AKV scopes don't work for SP auth yet")
	}
	scopesAPI := NewSecretScopesAPI(context.Background(), client)
	name := qa.RandomName("tf-scope-")

	err := scopesAPI.Create(SecretScope{
		Name: name,
		KeyvaultMetadata: &KeyvaultMetadata{
			ResourceID: resourceID,
			DNSName:    DNSName,
		},
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(name))
	}()

	scope, err := scopesAPI.Read(name)
	require.NoError(t, err)
	require.Equal(t, "AZURE_KEYVAULT", scope.BackendType)
	assert.Equal(t, resourceID, scope.KeyvaultMetadata.ResourceID)
	assert.Equal(t, DNSName, scope.KeyvaultMetadata.DNSName)
}

func TestAccInitialManagePrincipals(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	scopesAPI := NewSecretScopesAPI(context.Background(), client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(SecretScope{Name: scope})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := NewSecretAclsAPI(ctx, client)
	acls, err := secretACLAPI.List(scope)
	require.NoError(t, err)

	usersAPI := identity.NewUsersAPI(ctx, client)
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
	ctx := context.Background()
	scopesAPI := NewSecretScopesAPI(ctx, client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(SecretScope{
		Name:                   scope,
		InitialManagePrincipal: "users",
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := NewSecretAclsAPI(ctx, client)
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
	scope := qa.RandomName("tf-")
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
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
					acceptance.ResourceCheck("databricks_secret_scope.my_scope",
						func(client *common.DatabricksClient, id string) error {
							ctx := context.Background()
							secretACLAPI := NewSecretAclsAPI(ctx, client)
							acls, err := secretACLAPI.List(id)
							require.NoError(t, err)

							usersAPI := identity.NewUsersAPI(ctx, client)
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
					err := NewSecretScopesAPI(context.Background(), client).Delete(scope)
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
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
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
		_, err := NewSecretScopesAPI(context.Background(), client).Read(rs.Primary.ID)
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
		assert.True(t, secretScope.BackendType == "DATABRICKS")
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
		resp, err := NewSecretScopesAPI(context.Background(), conn).Read(rs.Primary.ID)
		//t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*secretScope = resp
		return nil
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
