package tokens

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
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
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
	assert.Equal(t, 10, d.Get("creation_time"))
	assert.Equal(t, 20, d.Get("expiry_time"))
	assert.Equal(t, "", d.Get("token_value"))
}

func TestResourceTokenRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
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
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
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
		State: map[string]any{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
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
		State: map[string]any{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceTokenCreate_NoExpiration(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/token/create",
				ExpectedRequest: TokenRequest{},
				Response: TokenResponse{
					TokenValue: "dapi...",
					TokenInfo: &TokenInfo{
						TokenID:    "abc",
						Comment:    "",
						ExpiryTime: -1,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token/list",
				Response: TokenList{
					TokenInfos: []TokenInfo{
						{
							Comment:      "",
							CreationTime: 10,
							ExpiryTime:   -1,
							TokenID:      "abc",
						},
					},
				},
			},
		},
		Resource: ResourceToken(),
		State:    map[string]any{},
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, -1, d.Get("expiry_time")) // tokens without expiration have expiry_time set to -1 as listed in the examples on https://docs.databricks.com/dev-tools/api/latest/tokens.html#list
	assert.Equal(t, "dapi...", d.Get("token_value"))
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
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceTokenDelete_NotFoundNoError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token/delete",
				Response: apierr.NotFound("RESOURCE_DOES_NOT_EXIST"), // per documentation this is the error response
				Status:   404,
			},
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
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
