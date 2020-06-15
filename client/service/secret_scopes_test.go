package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

//go:generate easytags $GOFILE

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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/scopes/create", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.SecretScopes().Create(tt.args.Scope, tt.args.InitialManagePrincipal)
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/scopes/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.SecretScopes().Delete(tt.args.Scope)
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
		want     model.SecretScope
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
			want: model.SecretScope{
				Name:        "my-databricks-scope",
				BackendType: model.ScopeBackendTypeDatabricks,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/secrets/scopes/list?", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.SecretScopes().Read(tt.args.ScopeName)
			})
		})
	}
}

func TestSecretScopesAPI_List(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     []model.SecretScope
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
			want: []model.SecretScope{
				{
					Name:        "my-databricks-scope",
					BackendType: model.ScopeBackendTypeDatabricks,
				},
				{
					Name:        "mount-points",
					BackendType: model.ScopeBackendTypeDatabricks,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertRequestWithMockServer(t, nil, http.MethodGet, "/api/2.0/secrets/scopes/list?", nil, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.SecretScopes().List()
			})
		})
	}
}
