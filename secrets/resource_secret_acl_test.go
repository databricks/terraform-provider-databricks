package secrets

import (
	"os"
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
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
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
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
				},
			},
			// Additional GET for ApplyAndExpectData's automatic read verification
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
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
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
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
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=myapplication%2Fbranch",
				Response: workspace.AclItem{
					Permission: "MANAGE",
				},
			},
			// Additional GET for ApplyAndExpectData's automatic read verification
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=myapplication%2Fbranch",
				Response: workspace.AclItem{
					Permission: "MANAGE",
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
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
	// Add 3 failures to test that all 3 retries are attempted
	fixtures := []qa.HTTPFixture{}
	for i := 0; i < 3; i++ {
		fixtures = append(fixtures, qa.HTTPFixture{
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/put",
			Response: internalErrorResponse,
			Status:   400,
		})
	}
	
	d, err := qa.ResourceFixture{
		Fixtures: fixtures,
		Resource: ResourceSecretACL(),
		State: map[string]any{
			"permission": "MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "failed to create Secret ACL after 3 attempts")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretACLCreate_RetriesOnVerificationFailure(t *testing.T) {
	// This test verifies that the create operation retries when verification fails
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// First attempt - PUT succeeds, GET fails
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: doesNotExistResponse,
				Status:   404,
			},
			// Second attempt - PUT succeeds, GET fails
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: doesNotExistResponse,
				Status:   404,
			},
			// Third attempt - PUT succeeds, GET succeeds
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
				},
			},
			// Additional GET for ApplyAndExpectData's automatic read verification
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
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

func TestResourceSecretACLCreate_RetriesOnPermissionMismatch(t *testing.T) {
	// This test verifies that the create operation retries when permission doesn't match
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// First attempt - PUT succeeds, GET returns wrong permission
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "READ", // Wrong permission
				},
			},
			// Second attempt - PUT succeeds, GET returns correct permission
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/put",
				ExpectedRequest: workspace.PutAcl{
					Principal:  "something",
					Permission: "MANAGE",
					Scope:      "global",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
				},
			},
			// Additional GET for ApplyAndExpectData's automatic read verification
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
				Response: workspace.AclItem{
					Permission: "MANAGE",
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

func TestResourceSecretACLCreate_ExhaustsRetries(t *testing.T) {
	// This test verifies that the create operation fails after 3 attempts
	// Set retry interval to 1ms for faster testing
	os.Setenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS", "1")
	defer os.Unsetenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS")
	
	fixtures := []qa.HTTPFixture{}
	
	// Add 3 attempts, each with PUT success but GET failure
	for i := 0; i < 3; i++ {
		fixtures = append(fixtures, qa.HTTPFixture{
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/put",
			ExpectedRequest: workspace.PutAcl{
				Principal:  "something",
				Permission: "MANAGE",
				Scope:      "global",
			},
		})
		fixtures = append(fixtures, qa.HTTPFixture{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
			Response: doesNotExistResponse,
			Status:   404,
		})
	}
	
	d, err := qa.ResourceFixture{
		Fixtures: fixtures,
		Resource: ResourceSecretACL(),
		State: map[string]any{
			"permission": "MANAGE",
			"principal":  "something",
			"scope":      "global",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "secret ACL creation could not be verified after 3 attempts")
	assert.Equal(t, "", d.Id(), "Id should be empty when all retries are exhausted")
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
