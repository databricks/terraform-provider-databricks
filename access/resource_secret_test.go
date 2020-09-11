package access

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/put", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewSecretsAPI(client).Create(tt.args.StringValue, tt.args.Scope, tt.args.Key)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/secrets/delete", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewSecretsAPI(client).Delete(tt.args.Scope, tt.args.Key)
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
		want     []SecretMetadata
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
			want: []SecretMetadata{
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
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, "/api/2.0/secrets/list?scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewSecretsAPI(client).List(tt.args.Scope)
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
		want     SecretMetadata
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
			want: SecretMetadata{
				Key:                  "my-string-key",
				LastUpdatedTimestamp: 1520467595000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, "/api/2.0/secrets/list?scope=my-scope", &input, tt.response, http.StatusOK, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewSecretsAPI(client).Read(tt.args.Scope, tt.args.Key)
			})
		})
	}
}

func TestResourceSecretRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		ID:       "foo|||bar",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
	assert.Equal(t, "bar", d.Get("key"))
	assert.Equal(t, 12345678, d.Get("last_updated_timestamp"))
	assert.Equal(t, "foo", d.Get("scope"))
	assert.Equal(t, "", d.Get("string_value"))
}

func TestResourceSecretRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		ID:       "foo|||missing",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceSecretRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		ID:       "foo|||bar",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/put",
				ExpectedRequest: SecretsRequest{
					StringValue: "SparkIsTh3Be$t",
					Scope:       "foo",
					Key:         "bar",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		State: map[string]interface{}{
			"scope":        "foo",
			"key":          "bar",
			"string_value": "SparkIsTh3Be$t",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/put",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		State: map[string]interface{}{
			"key":          "...",
			"scope":        "...",
			"string_value": "...",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/secrets/delete",
				ExpectedRequest: map[string]string{
					"scope": "foo",
					"key":   "bar",
				},
			},
		},
		Resource: ResourceSecret(),
		Delete:   true,
		ID:       "foo|||bar",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		Delete:   true,
		ID:       "foo|||bar",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id())
}
