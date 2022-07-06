package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccRemoveScopes(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Cleaning up tests only from IDE")
	}
	t.Parallel()
	client := common.CommonEnvironmentClient()
	scopesAPI := secrets.NewSecretScopesAPI(context.Background(), client)
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
	if client.IsAzureClientSecretSet() {
		t.Skip("AKV scopes don't work for SP auth yet")
	}
	scopesAPI := secrets.NewSecretScopesAPI(context.Background(), client)
	name := qa.RandomName("tf-scope-")

	err := scopesAPI.Create(secrets.SecretScope{
		Name: name,
		KeyvaultMetadata: &secrets.KeyvaultMetadata{
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
	t.Parallel()
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	scopesAPI := secrets.NewSecretScopesAPI(context.Background(), client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(secrets.SecretScope{Name: scope})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
	acls, err := secretACLAPI.List(scope)
	require.NoError(t, err)

	usersAPI := scim.NewUsersAPI(ctx, client)
	me, err := usersAPI.Me()
	require.NoError(t, err)
	assert.Equal(t, 1, len(acls))
	assert.Equal(t, me.UserName, acls[0].Principal)
}

func TestAccInitialManagePrincipalsGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	scopesAPI := secrets.NewSecretScopesAPI(ctx, client)

	scope := fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	err := scopesAPI.Create(secrets.SecretScope{
		Name:                   scope,
		InitialManagePrincipal: "users",
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, scopesAPI.Delete(scope))
	}()

	secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
	acls, err := secretACLAPI.List(scope)
	require.NoError(t, err)
	assert.Equal(t, 1, len(acls))
	assert.Equal(t, "users", acls[0].Principal)
}

func TestAccSecretScopeResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	scope := qa.RandomName("tf-")
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "databricks_secret_scope" "my_scope" {
					name = "%s"
				}`, scope),
				Check: resource.ComposeTestCheckFunc(
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
					acceptance.ResourceCheck("databricks_secret_scope.my_scope",
						func(ctx context.Context, client *common.DatabricksClient, id string) error {
							secretACLAPI := secrets.NewSecretAclsAPI(ctx, client)
							acls, err := secretACLAPI.List(id)
							require.NoError(t, err)

							usersAPI := scim.NewUsersAPI(ctx, client)
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
					err := secrets.NewSecretScopesAPI(context.Background(), client).Delete(scope)
					assert.NoError(t, err, err)
				},
				Config: fmt.Sprintf(`
				resource "databricks_secret_scope" "my_scope" {
					name = "%s"
				}`, scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", "DATABRICKS"),
				),
			},
		},
	})
}
