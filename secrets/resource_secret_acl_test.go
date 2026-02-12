package secrets

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

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
					Permission: workspace.AclPermissionManage,
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
	qa.AssertErrorStartsWith(t, err, "failed to create Secret ACL")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

type getACLResponse struct {
	resp *workspace.AclItem
	err  error
}

type mockSecretsClient struct {
	putDelay time.Duration // delay before returning the put error
	putError []error
	getResp  []getACLResponse
}

func (m *mockSecretsClient) PutAcl(ctx context.Context, req workspace.PutAcl) error {
	time.Sleep(m.putDelay)
	err := m.putError[0]
	m.putError = m.putError[1:] // pop
	return err
}

func (m *mockSecretsClient) GetAcl(ctx context.Context, req workspace.GetAclRequest) (*workspace.AclItem, error) {
	r := m.getResp[0]
	m.getResp = m.getResp[1:] // pop
	return r.resp, r.err
}

func TestResourceSecretACLDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				ExpectedRequest: workspace.DeleteAcl{
					Scope:     "global",
					Principal: "something",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||something",
		State: map[string]any{
			"scope":      "global",
			"principal":  "something",
			"permission": "READ",
		},
	}.ApplyNoError(t)
}

func TestResourceSecretACLDelete_SkipsCurrentUserWithManagePermission(t *testing.T) {
	// Test that delete is skipped when the principal is the current user with MANAGE permission.
	// This prevents users from accidentally locking themselves out of managing the secret scope.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: map[string]any{
					"id":       "123",
					"userName": "current_user@example.com",
				},
			},
			// No delete call expected since we skip deletion for current user with MANAGE
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||current_user@example.com",
		State: map[string]any{
			"scope":      "global",
			"principal":  "current_user@example.com",
			"permission": "MANAGE",
		},
	}.ApplyNoError(t)
}

func TestResourceSecretACLDelete_DoesNotSkipCurrentUserWithReadPermission(t *testing.T) {
	// Test that delete proceeds when the principal is the current user but has READ permission (not MANAGE).
	// No Me API call is made since permission is not MANAGE.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				ExpectedRequest: workspace.DeleteAcl{
					Scope:     "global",
					Principal: "current_user@example.com",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||current_user@example.com",
		State: map[string]any{
			"scope":      "global",
			"principal":  "current_user@example.com",
			"permission": "READ",
		},
	}.ApplyNoError(t)
}

func TestResourceSecretACLDelete_DeletesOtherUserWithManagePermission(t *testing.T) {
	// Test that delete proceeds when the principal is a different user, even with MANAGE permission.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: map[string]any{
					"id":       "123",
					"userName": "current_user@example.com",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/acls/delete",
				ExpectedRequest: workspace.DeleteAcl{
					Scope:     "global",
					Principal: "other_user@example.com",
				},
			},
		},
		Resource: ResourceSecretACL(),
		Delete:   true,
		ID:       "global|||other_user@example.com",
		State: map[string]any{
			"scope":      "global",
			"principal":  "other_user@example.com",
			"permission": "MANAGE",
		},
	}.ApplyNoError(t)
}

func TestRobustPutACL(t *testing.T) {
	testCases := []struct {
		name          string
		timeout       time.Duration
		sc            secretsClient
		req           workspace.PutAcl
		wantErrPrefix string
	}{
		{
			name:    "success",
			timeout: defaultTimeout,
			sc: &mockSecretsClient{
				putError: []error{nil},
				getResp:  []getACLResponse{{resp: &workspace.AclItem{Permission: "MANAGE"}, err: nil}},
			},
			req: workspace.PutAcl{
				Permission: "MANAGE",
			},
		},
		{
			name:    "timeout",
			timeout: 10 * time.Millisecond,
			sc: &mockSecretsClient{
				putDelay: 1 * time.Second, // more than timeout
				putError: []error{nil},
				getResp:  []getACLResponse{{resp: &workspace.AclItem{Permission: "OTHER"}, err: nil}},
			},
			req:           workspace.PutAcl{Permission: "MANAGE"},
			wantErrPrefix: "secret ACL permission mismatch",
		},
		{
			name:    "put error",
			timeout: defaultTimeout,
			sc: &mockSecretsClient{
				putError: []error{errors.New("test error")},
			},
			req:           workspace.PutAcl{Permission: "MANAGE"},
			wantErrPrefix: "failed to create Secret ACL: test error",
		},
		{
			name:    "retry on get error",
			timeout: defaultTimeout,
			sc: &mockSecretsClient{
				putError: []error{
					nil,
					nil,
				},
				getResp: []getACLResponse{
					{err: errors.New("test error")},
					{resp: &workspace.AclItem{Permission: "MANAGE"}},
				},
			},
			req: workspace.PutAcl{Permission: "MANAGE"},
		},
		{
			name:    "retry on permission mismatch",
			timeout: defaultTimeout,
			sc: &mockSecretsClient{
				putError: []error{
					nil,
					nil,
					nil,
				},
				getResp: []getACLResponse{
					{err: errors.New("test error")},
					{resp: &workspace.AclItem{Permission: "OTHER"}},
					{resp: &workspace.AclItem{Permission: "MANAGE"}},
				},
			},
			req: workspace.PutAcl{Permission: "MANAGE"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := robustPutACL(tc.sc, context.Background(), tc.req, tc.timeout)

			if err == nil && tc.wantErrPrefix != "" {
				t.Errorf("expected error, got nil")
			}
			if err != nil && tc.wantErrPrefix == "" {
				t.Errorf("expected no error, got %v", err)
			}
			if err != nil && tc.wantErrPrefix != "" {
				if !strings.HasPrefix(err.Error(), tc.wantErrPrefix) {
					t.Errorf("expected error to start with %q, got %v", tc.wantErrPrefix, err)
				}
			}
		})
	}
}
