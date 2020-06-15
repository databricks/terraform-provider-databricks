package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

//go:generate easytags $GOFILE

func TestSecretAclsAPI_Create(t *testing.T) {
	type args struct {
		Scope      string              `json:"scope"`
		Principal  string              `json:"principal"`
		Permission model.ACLPermission `json:"permission"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Basic Test",
			response: "",
			args: args{
				Scope:      "my-scope",
				Principal:  "my-principal",
				Permission: model.ACLPermissionManage,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/acls/put", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.SecretAcls().Create(tt.args.Scope, tt.args.Principal, tt.args.Permission)
			})
		})
	}
}

func TestSecretAclsAPI_Delete(t *testing.T) {
	type args struct {
		Scope     string `json:"scope"`
		Principal string `json:"principal"`
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
				Scope:     "my-scope",
				Principal: "my-principal",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/acls/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.SecretAcls().Delete(tt.args.Scope, tt.args.Principal)
			})
		})
	}
}

func TestSecretAclsAPI_List(t *testing.T) {
	type args struct {
		Scope string `json:"scope"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     []model.ACLItem
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "items": [
							{
								"principal": "admins",
								"permission": "MANAGE"
							},{
								"principal": "data-scientists",
								"permission": "READ"
							}
						  ]
						}`,
			args: args{
				Scope: "my-scope",
			},
			want: []model.ACLItem{
				{
					Principal:  "admins",
					Permission: model.ACLPermissionManage,
				},
				{
					Principal:  "data-scientists",
					Permission: model.ACLPermissionRead,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/secrets/acls/list?scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.SecretAcls().List(tt.args.Scope)
			})
		})
	}
}

func TestSecretAclsAPI_Read(t *testing.T) {
	type args struct {
		Scope     string `json:"scope"`
		Principal string `json:"principal"`
	}
	tests := []struct {
		name     string
		response string
		args     args
		want     model.ACLItem
		wantErr  bool
	}{
		{
			name: "Basic test",
			response: `{
						  "principal": "data-scientists",
						  "permission": "READ"
						}`,
			args: args{
				Scope:     "my-scope",
				Principal: "my-principal",
			},
			want: model.ACLItem{
				Principal:  "data-scientists",
				Permission: model.ACLPermissionRead,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, "/api/2.0/secrets/acls/get?principal=my-principal&scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.SecretAcls().Read(tt.args.Scope, tt.args.Principal)
			})
		})
	}
}
