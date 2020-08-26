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
	assert.Equal(t, "users", d.Get("initial_manage_principal"))
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
					"scope":                    "Boom",
					"initial_manage_principal": "groups",
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
			"initial_manage_principal": "groups",
			"name":                     "Boom",
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

func TestSecretScopesAPI_Create(t *testing.T) {
	type args struct {
		Scope                  string `json:"scope"`
		InitialManagePrincipal string `json:"initial_manage_principal"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Basic test",
			response: "",
			args: args{
				Scope:                  "my-scope",
				InitialManagePrincipal: "users",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/scopes/create", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewSecretScopesAPI(&client).Create(tt.args.Scope, tt.args.InitialManagePrincipal)
			})
		})
	}
}

func TestSecretScopesAPI_Delete(t *testing.T) {
	type args struct {
		Scope string `json:"scope"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Basic test",
			response: "",
			args: args{
				Scope: "my-scope",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/scopes/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewSecretScopesAPI(&client).Delete(tt.args.Scope)
			})
		})
	}
}

func TestSecretScopesAPI_Read(t *testing.T) {
	type args struct {
		ScopeName string `json:"scope_name"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     SecretScope
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "scopes": [{
							  "name": "my-databricks-scope",
							  "backend_type": "DATABRICKS"
						  },{
							  "name": "mount-points",
							  "backend_type": "DATABRICKS"
						  }]
						}`,
			args: args{
				ScopeName: "my-databricks-scope",
			},
			want: SecretScope{
				Name:        "my-databricks-scope",
				BackendType: ScopeBackendTypeDatabricks,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/secrets/scopes/list", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewSecretScopesAPI(&client).Read(tt.args.ScopeName)
			})
		})
	}
}

func TestSecretScopesAPI_List(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     []SecretScope
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "scopes": [{
							  "name": "my-databricks-scope",
							  "backend_type": "DATABRICKS"
						  },{
							  "name": "mount-points",
							  "backend_type": "DATABRICKS"
						  }]
						}`,
			want: []SecretScope{
				{
					Name:        "my-databricks-scope",
					BackendType: ScopeBackendTypeDatabricks,
				},
				{
					Name:        "mount-points",
					BackendType: ScopeBackendTypeDatabricks,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertRequestWithMockServer(t, nil, http.MethodGet, "/api/2.0/secrets/scopes/list", nil, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewSecretScopesAPI(&client).List()
			})
		})
	}
}
