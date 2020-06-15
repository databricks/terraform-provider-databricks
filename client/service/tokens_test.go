package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestTokensAPI_Create(t *testing.T) {
	type args struct {
		LifetimeSeconds int32  `json:"lifetime_seconds,omitempty"`
		Comment         string `json:"comment,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		want           model.TokenResponse
		wantErr        bool
	}{
		{
			name: "Create Test",
			response: `{
						  "token_value":"dapideadbeefdeadbeefdeadbeefdeadbeef",
						  "token_info": {
							"token_id":"5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
							"creation_time":1513120516294,
							"expiry_time":1513120616294,
							"comment":"this is an example token"
						  }
						}`,
			responseStatus: http.StatusOK,
			want: model.TokenResponse{
				TokenValue: "dapideadbeefdeadbeefdeadbeefdeadbeef",
				TokenInfo: &model.TokenInfo{
					TokenID:      "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
					CreationTime: 1513120516294,
					ExpiryTime:   1513120616294,
					Comment:      "this is an example token",
				},
			},
			args: args{
				LifetimeSeconds: 1000,
				Comment:         "this is an example token",
			},
			wantErr: false,
		},
		{
			name:           "Create Test Failure",
			response:       ``,
			want:           model.TokenResponse{},
			responseStatus: http.StatusBadRequest,
			args: args{
				LifetimeSeconds: 1000,
				Comment:         "this is an example token",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/token/create", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Tokens().Create(tt.args.LifetimeSeconds, tt.args.Comment)
			})
		})
	}
}

func TestTokensAPI_Delete(t *testing.T) {
	type args struct {
		TokenID string `json:"token_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name:           "Delete test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				TokenID: "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/token/delete", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Tokens().Delete(tt.args.TokenID)
			})
		})
	}
}

func TestTokensAPI_List(t *testing.T) {
	type args struct{}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           *args
		wantURI        string
		want           []model.TokenInfo
		wantErr        bool
	}{
		{
			name: "List test",
			response: `{
						  "token_infos": [
							{
							  "token_id":"5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
							  "creation_time":1513120516294,
							  "expiry_time":1513120616294,
							  "comment":"this is an example token"
							},
							{
							  "token_id":"902eb9ac42c9bef80d0097a2d1746533103c88593add482a331500187946ceb5",
							  "creation_time":1512684023036,
							  "expiry_time":-1,
							  "comment":"this is another example token"
							}
						  ]
						}`,
			responseStatus: http.StatusOK,
			args:           nil,
			wantURI:        "/api/2.0/token/list?",
			want: []model.TokenInfo{
				{
					TokenID:      "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
					CreationTime: 1513120516294,
					ExpiryTime:   1513120616294,
					Comment:      "this is an example token",
				},
				{
					TokenID:      "902eb9ac42c9bef80d0097a2d1746533103c88593add482a331500187946ceb5",
					CreationTime: 1512684023036,
					ExpiryTime:   -1,
					Comment:      "this is another example token",
				},
			},
			wantErr: false,
		},
		{
			name:           "List test failure",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args:           nil,
			wantURI:        "/api/2.0/token/list?",
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Tokens().List()
			})
		})
	}
}

func TestTokensAPI_Read(t *testing.T) {
	type args struct {
		TokenID string `json:"token_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		wantURI        string
		want           model.TokenInfo
		wantErr        bool
	}{
		{
			name: "Read test happy path",
			response: `{
						  "token_infos": [
							{
							  "token_id":"5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
							  "creation_time":1513120516294,
							  "expiry_time":1513120616294,
							  "comment":"this is an example token"
							},
							{
							  "token_id":"902eb9ac42c9bef80d0097a2d1746533103c88593add482a331500187946ceb5",
							  "creation_time":1512684023036,
							  "expiry_time":-1,
							  "comment":"this is another example token"
							}
						  ]
						}`,
			args: args{
				TokenID: "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
			},
			responseStatus: http.StatusOK,
			want: model.TokenInfo{
				TokenID:      "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
				CreationTime: 1513120516294,
				ExpiryTime:   1513120616294,
				Comment:      "this is an example token",
			},
			wantURI: "/api/2.0/token/list?",
			wantErr: false,
		},
		{
			name:     "Read list fails",
			response: ``,
			args: args{
				TokenID: "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
			},
			responseStatus: http.StatusBadRequest,
			want:           model.TokenInfo{},
			wantURI:        "/api/2.0/token/list?",
			wantErr:        true,
		},
		{
			name: "Read list fails",
			response: `{
						  "token_infos": [
							{
							  "token_id":"902eb9ac42c9bef80d0097a2d1746533103c88593add482a33aaaaaaa946ceb5",
							  "creation_time":1513120516294,
							  "expiry_time":1513120616294,
							  "comment":"this is an example token"
							},
							{
							  "token_id":"902eb9ac42c9bef80d0097a2d1746533103c88593add482a331500187946ceb5",
							  "creation_time":1512684023036,
							  "expiry_time":-1,
							  "comment":"this is another example token"
							}
						  ]
						}`,
			args: args{
				TokenID: "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
			},
			responseStatus: http.StatusOK,
			want:           model.TokenInfo{},
			wantURI:        "/api/2.0/token/list?",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Tokens().Read(tt.args.TokenID)
			})
		})
	}
}
