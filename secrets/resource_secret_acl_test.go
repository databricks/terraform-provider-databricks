package secrets

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

var internalErrorResponse = apierr.APIError{
	ErrorCode:  "INVALID_REQUEST",
	Message:    "Internal error happened",
	StatusCode: 400,
}

var doesNotExistResponse = apierr.APIError{
	StatusCode: 404,
	ErrorCode:  "NOT_FOUND",
	Message:    "Secret Scope does not exist",
}

func TestResourceSecretACLRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Read:     true,
		ID:       "global|||something",
	}.ApplyAndExpectData(t, map[string]any{
		"permission": "MANAGE",
		"principal":  "something",
		"scope":      "global",
		"id":         "global|||something",
	})
}

func TestResourceSecretACLRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: doesNotExistResponse,
				Status:   404,
			},
		},
		Resource: ResourceSecretACL(),
		Read:     true,
		Removed:  true,
		ID:       "global|||something",
	}.ApplyNoError(t)
}

func TestResourceSecretACLRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: internalErrorResponse,
				Status:   400,
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
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
		},
		Resource: ResourceSecretACL(),
		State: map[string]any{
			"permission": "MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"permission": "MANAGE",
		"principal":  "something",
		"scope":      "global",
		"id":         "global|||something",
	})
}

func TestResourceSecretACLCreate_ScopeWithSlash(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: workspace.AclPermissionManage,
					Scope:      "myapplication/branch",
				},
			},
		},
		Resource: ResourceSecretACL(),
		State: map[string]any{
			"permission": "MANAGE",
			"principal":  "something",
			"scope":      "myapplication/branch",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"permission": "MANAGE",
		"principal":  "something",
		"scope":      "myapplication/branch",
		"id":         "myapplication/branch|||something",
	})
}

func TestResourceSecretACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				Response: internalErrorResponse,
				Status:   400,
			},
		},
		Resource: ResourceSecretACL(),
		State: map[string]any{
			"permission": "MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretACLDelete(t *testing.T) {
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "global|||something",
	})
}

func TestResourceSecretACLDelete_DoesntExist(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				ExpectedRequest: map[string]string{
					"scope":     "global",
					"principal": "something",
				},
				Response: doesNotExistResponse,
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||something",
	}.ApplyNoError(t)
}

func TestResourceSecretACLDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				Response: internalErrorResponse,
				Status:   400,
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||something",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "global|||something", d.Id())
}
