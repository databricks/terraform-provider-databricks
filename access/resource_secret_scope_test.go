package access

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

//go:generate easytags $GOFILE

func TestResourceSecretScopeRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/secrets/scopes/list",
				Response: SecretScopeList{
					Scopes: []SecretScope{
						{
							Name:        "abc",
							BackendType: "DATABRICKS",
						},
					},
				},
				Status: 200,
			},
		},
		Resource: ResourceSecretScope(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "DATABRICKS", d.Get("backend_type"))
	assert.Equal(t, "", d.Get("initial_manage_principal"))
	assert.Equal(t, "abc", d.Get("name"))
}

func TestResourceSecretScopeRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/secrets/scopes/list",
				Response: SecretScopeList{
					Scopes: []SecretScope{
						{
							Name:        "bcd",
							BackendType: "DATABRICKS",
						},
					},
				},
				Status: 200,
			},
		},
		Resource: ResourceSecretScope(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceSecretScopeRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/scopes/list",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretScope(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretScopeCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/scopes/create",
				ExpectedRequest: map[string]string{
					"scope": "Boom",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/secrets/scopes/list",
				Response: SecretScopeList{
					Scopes: []SecretScope{
						{
							Name:        "Boom",
							BackendType: "DATABRICKS",
						},
					},
				},
				Status: 200,
			},
		},
		Resource: ResourceSecretScope(),
		State: map[string]interface{}{
			"name": "Boom",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "Boom", d.Id())
}

func TestResourceSecretScopeCreate_Users(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/scopes/create",
				ExpectedRequest: map[string]string{
					"scope":                    "Boom",
					"initial_manage_principal": "users",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/secrets/scopes/list",
				Response: SecretScopeList{
					Scopes: []SecretScope{
						{
							Name:        "Boom",
							BackendType: "DATABRICKS",
						},
					},
				},
				Status: 200,
			},
		},
		Resource: ResourceSecretScope(),
		State: map[string]interface{}{
			"name":                     "Boom",
			"initial_manage_principal": "users",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "Boom", d.Id())
}

func TestResourceSecretScopeCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/scopes/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretScope(),
		State: map[string]interface{}{
			"initial_manage_principal": "groups",
			"name":                     "Boom",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretScopeDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/secrets/scopes/delete",
				ExpectedRequest: map[string]string{
					"scope": "abc",
				},
			},
		},
		Resource: ResourceSecretScope(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceSecretScopeDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/scopes/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecretScope(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
