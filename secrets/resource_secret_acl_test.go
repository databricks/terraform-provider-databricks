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
