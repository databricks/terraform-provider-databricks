package service

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestSecretsScopesAclsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()

	testScope := "my-test-scope"
	testKey := "my-test-key"
	testSecret := "my-test-secret"
	initialManagePrincipal := "users"
	//TODO: Please replace it with users api and get random user
	testPrincipal := "users"

	err := client.SecretScopes().Create(testScope, initialManagePrincipal)
	assert.NoError(t, err, err)

	defer func() {
		//	Deleting scope deletes everything else
		err := client.SecretScopes().Delete(testScope)
		assert.NoError(t, err, err)
	}()

	scopes, err := client.SecretScopes().List()
	assert.NoError(t, err, err)
	assert.True(t, len(scopes) >= 1, "Scopes are empty list")

	scope, err := client.SecretScopes().Read(testScope)
	assert.NoError(t, err, err)
	assert.Equal(t, testScope, scope.Name, "Scope lookup does not yield same scope")

	err = client.Secrets().Create(testSecret, testScope, testKey)
	assert.NoError(t, err, err)

	secrets, err := client.Secrets().List(testScope)
	assert.NoError(t, err, err)
	assert.True(t, len(secrets) > 0, "Secrets are empty list")

	secret, err := client.Secrets().Read(testScope, testKey)
	assert.NoError(t, err, err)
	assert.Equal(t, testKey, secret.Key, "Secret lookup does not yield same key")

	err = client.SecretAcls().Create(testScope, testPrincipal, model.ACLPermissionManage)
	assert.NoError(t, err, err)

	secretAcls, err := client.SecretAcls().List(testScope)
	assert.NoError(t, err, err)
	assert.True(t, len(secretAcls) > 0, "Secrets acls are empty list")

	secretACL, err := client.SecretAcls().Read(testScope, testPrincipal)
	assert.NoError(t, err, err)
	assert.Equal(t, testPrincipal, secretACL.Principal, "Secret lookup does not yield same key")
	assert.Equal(t, model.ACLPermissionManage, secretACL.Permission, "Secret lookup does not yield same key")

	err = client.Secrets().Delete(testScope, testKey)
	assert.NoError(t, err, err)

	err = client.SecretAcls().Delete(testScope, testPrincipal)
	assert.NoError(t, err, err)
}
