package access

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

//go:generate easytags $GOFILE

func TestSecretsScopesAclsIntegration(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	testScope := "my-test-scope"
	testKey := "my-test-key"
	testSecret := "my-test-secret"
	initialManagePrincipal := ""
	// TODO: on random group
	testPrincipal := "users"

	err := NewSecretScopesAPI(client).Create(testScope, initialManagePrincipal)
	assert.NoError(t, err, err)

	defer func() {
		//	Deleting scope deletes everything else
		err := NewSecretScopesAPI(client).Delete(testScope)
		assert.NoError(t, err, err)
	}()

	scopes, err := NewSecretScopesAPI(client).List()
	assert.NoError(t, err, err)
	assert.True(t, len(scopes) >= 1, "Scopes are empty list")

	scope, err := NewSecretScopesAPI(client).Read(testScope)
	assert.NoError(t, err, err)
	assert.Equal(t, testScope, scope.Name, "Scope lookup does not yield same scope")

	err = NewSecretsAPI(client).Create(testSecret, testScope, testKey)
	assert.NoError(t, err, err)

	secrets, err := NewSecretsAPI(client).List(testScope)
	assert.NoError(t, err, err)
	assert.True(t, len(secrets) > 0, "Secrets are empty list")

	secret, err := NewSecretsAPI(client).Read(testScope, testKey)
	assert.NoError(t, err, err)
	assert.Equal(t, testKey, secret.Key, "Secret lookup does not yield same key")

	err = NewSecretAclsAPI(client).Create(testScope, testPrincipal, ACLPermissionManage)
	assert.NoError(t, err, err)

	secretAcls, err := NewSecretAclsAPI(client).List(testScope)
	assert.NoError(t, err, err)
	assert.True(t, len(secretAcls) > 0, "Secrets acls are empty list")

	secretACL, err := NewSecretAclsAPI(client).Read(testScope, testPrincipal)
	assert.NoError(t, err, err)
	assert.Equal(t, testPrincipal, secretACL.Principal, "Secret lookup does not yield same key")
	assert.Equal(t, ACLPermissionManage, secretACL.Permission, "Secret lookup does not yield same key")

	err = NewSecretsAPI(client).Delete(testScope, testKey)
	assert.NoError(t, err, err)

	err = NewSecretAclsAPI(client).Delete(testScope, testPrincipal)
	assert.NoError(t, err, err)
}

func TestResourceSecretACLRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: ACLItem{
					Permission: "CAN_MANAGE",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Read:     true,
		ID:       "global|||something",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id(), "Id should not be empty")
	assert.Equal(t, "CAN_MANAGE", d.Get("permission"))
	assert.Equal(t, "something", d.Get("principal"))
	assert.Equal(t, "global", d.Get("scope"))
}

func TestResourceSecretACLRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceSecretACL(),
		Read:     true,
		ID:       "global|||something",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceSecretACLRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretACL(),
		Read:     true,
		ID:       "global|||something",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "global|||something", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretACLCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: SecretACLRequest{
					Principal:  "something",
					Permission: "CAN_MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: ACLItem{
					Permission: "CAN_MANAGE",
				},
			},
		},
		Resource: ResourceSecretACL(),
		State: map[string]interface{}{
			"permission": "CAN_MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id())
}

func TestResourceSecretACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretACL(),
		State: map[string]interface{}{
			"permission": "CAN_MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretACLDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				ExpectedRequest: map[string]string{
					"scope":     "global",
					"principal": "something",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||something",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id())
}

func TestResourceSecretACLDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||something",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "global|||something", d.Id())
}
