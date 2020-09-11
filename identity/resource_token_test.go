package identity

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceTokenRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token/list",
				Response: TokenList{
					TokenInfos: []TokenInfo{
						{
							Comment:      "Hello, world!",
							CreationTime: 10,
							ExpiryTime:   20,
							TokenID:      "abc",
						},
					},
				},
			},
		},
		Resource: ResourceToken(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
	assert.Equal(t, 10, d.Get("creation_time"))
	assert.Equal(t, 20, d.Get("expiry_time"))
	assert.Equal(t, "", d.Get("token_value"))
}

func TestResourceTokenRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token/list",
				Response: TokenList{
					TokenInfos: []TokenInfo{
						{
							Comment:      "Hello, world!",
							CreationTime: 10,
							ExpiryTime:   20,
							TokenID:      "bcd",
						},
					},
				},
			},
		},
		Resource: ResourceToken(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceTokenRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token/list",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceToken(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceTokenCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token/create",
				ExpectedRequest: TokenRequest{
					LifetimeSeconds: 300,
					Comment:         "Hello world!",
				},
				Response: TokenResponse{
					TokenValue: "dapi...",
					TokenInfo: &TokenInfo{
						TokenID: "abc",
						// other fields may be irrelevant...
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token/list",
				Response: TokenList{
					TokenInfos: []TokenInfo{
						{
							Comment:      "Hello, world!",
							CreationTime: 10,
							ExpiryTime:   20,
							TokenID:      "abc",
						},
					},
				},
			},
		},
		Resource: ResourceToken(),
		State: map[string]interface{}{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "dapi...", d.Get("token_value"))
}

func TestResourceTokenCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceToken(),
		State: map[string]interface{}{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceTokenDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/token/delete",
				ExpectedRequest: map[string]string{
					"token_id": "abc",
				},
			},
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceTokenDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

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
		want           TokenResponse
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
			want: TokenResponse{
				TokenValue: "dapideadbeefdeadbeefdeadbeefdeadbeef",
				TokenInfo: &TokenInfo{
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
			want:           TokenResponse{},
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/token/create", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewTokensAPI(client).Create(time.Duration(tt.args.LifetimeSeconds)*time.Second, tt.args.Comment)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/token/delete", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewTokensAPI(client).Delete(tt.args.TokenID)
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
		want           []TokenInfo
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
			wantURI:        "/api/2.0/token/list",
			want: []TokenInfo{
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
			wantURI:        "/api/2.0/token/list",
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewTokensAPI(client).List()
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
		want           TokenInfo
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
			want: TokenInfo{
				TokenID:      "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
				CreationTime: 1513120516294,
				ExpiryTime:   1513120616294,
				Comment:      "this is an example token",
			},
			wantURI: "/api/2.0/token/list",
			wantErr: false,
		},
		{
			name:     "Read list fails",
			response: ``,
			args: args{
				TokenID: "5715498424f15ee0213be729257b53fc35a47d5953e3bdfd8ed22a0b93b339f4",
			},
			responseStatus: http.StatusBadRequest,
			want:           TokenInfo{},
			wantURI:        "/api/2.0/token/list",
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
			want:           TokenInfo{},
			wantURI:        "/api/2.0/token/list",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewTokensAPI(client).Read(tt.args.TokenID)
			})
		})
	}
}

func TestAccCreateToken(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()

	lifeTimeSeconds := time.Duration(30) * time.Second
	comment := "Hello world"

	token, err := NewTokensAPI(client).Create(lifeTimeSeconds, comment)
	assert.NoError(t, err, err)
	assert.True(t, len(token.TokenValue) > 0, "Token value is empty")

	defer func() {
		err := NewTokensAPI(client).Delete(token.TokenInfo.TokenID)
		assert.NoError(t, err, err)
	}()

	_, err = NewTokensAPI(client).Read(token.TokenInfo.TokenID)
	assert.NoError(t, err, err)

	tokenList, err := NewTokensAPI(client).List()
	assert.NoError(t, err, err)
	assert.True(t, len(tokenList) > 0, "Token list is empty")
}
