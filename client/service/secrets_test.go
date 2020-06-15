package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

//go:generate easytags $GOFILE

func TestSecretsAPI_Create(t *testing.T) {
	type args struct {
		StringValue string `json:"string_value"`
		Scope       string `json:"scope"`
		Key         string `json:"key"`
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
				StringValue: "my-secret",
				Scope:       "my-scope",
				Key:         "my-key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/put", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Secrets().Create(tt.args.StringValue, tt.args.Scope, tt.args.Key)
			})
		})
	}
}

func TestSecretsAPI_Delete(t *testing.T) {
	type args struct {
		Scope string `json:"scope"`
		Key   string `json:"key"`
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
				Key:   "my-key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Secrets().Delete(tt.args.Scope, tt.args.Key)
			})
		})
	}
}

func TestSecretsAPI_ListSecrets(t *testing.T) {
	type args struct {
		Scope string `json:"scope"`
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     []model.SecretMetadata
		wantErr  bool
	}{
		{
			name: "Basic Test",
			args: args{
				Scope: "my-scope",
			},
			response: `{
						  "secrets": [
							{
							  "key": "my-string-key",
							  "last_updated_timestamp": 1520467595000
							},
							{
							  "key": "my-byte-key",
							  "last_updated_timestamp": 1520467595000
							}
						  ]
						}`,
			want: []model.SecretMetadata{
				{
					Key:                  "my-string-key",
					LastUpdatedTimestamp: 1520467595000,
				},
				{
					Key:                  "my-byte-key",
					LastUpdatedTimestamp: 1520467595000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, "/api/2.0/secrets/list?scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Secrets().List(tt.args.Scope)
			})
		})
	}
}

func TestSecretsAPI_Read(t *testing.T) {
	type args struct {
		Scope string `json:"scope"`
		Key   string `json:"key"`
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     model.SecretMetadata
		wantErr  bool
	}{
		{
			name: "Basic test",
			args: args{
				Scope: "my-scope",
				Key:   "my-string-key",
			},
			response: `{
						  "secrets": [
							{
							  "key": "my-string-key",
							  "last_updated_timestamp": 1520467595000
							},
							{
							  "key": "my-byte-key",
							  "last_updated_timestamp": 1520467595000
							}
						  ]
						}`,
			want: model.SecretMetadata{
				Key:                  "my-string-key",
				LastUpdatedTimestamp: 1520467595000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, "/api/2.0/secrets/list?scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Secrets().Read(tt.args.Scope, tt.args.Key)
			})
		})
	}
}
