package tokens

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceOboTokenRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",

				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment:    "Hello, world!",
						ExpiryTime: time.Now().UnixMilli() + 1000,
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "Hello, world!",
		"id":      "abc",
	})
}

func TestResourceOboTokenRead_Expired(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",

				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment:    "Hello, world!",
						ExpiryTime: time.Now().UnixMilli() - 1000,
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "",
	})
}

func TestResourceOboTokenRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",
				Status:   500,
				Response: apierr.APIError{
					Message: "nope",
				},
			},
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.ExpectError(t, "nope")

}
func TestResourceOboTokenRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Token does not exist",
				},
				Status: 404,
			},
		},
		Resource: ResourceOboToken(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "",
	})
}

func TestResourceOboTokenCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token-management/on-behalf-of/tokens",
				Status:   500,
				Response: apierr.APIError{
					Message: "nope",
				},
			},
		},
		Resource: ResourceOboToken(),
		Create:   true,
		New:      true,
	}.ExpectError(t, "nope")
}

func TestResourceOboTokenCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token-management/on-behalf-of/tokens",
				ExpectedRequest: OboToken{
					ApplicationID:   "abc",
					LifetimeSeconds: 60,
					Comment:         "e",
				},
				Response: TokenResponse{
					TokenValue: "s#Cr3t!11",
					TokenInfo: &TokenInfo{
						TokenID:    "bcd",
						ExpiryTime: time.Now().UnixMilli() + 1000,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/bcd",
				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment:    "Hello, world!",
						ExpiryTime: time.Now().UnixMilli() + 1000,
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Create:   true,
		HCL: `
		application_id = "abc"
		comment = "e"
		lifetime_seconds = 60
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "Hello, world!",
		"id":      "bcd",
	})
}

func TestResourceOboTokenDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/token-management/tokens/abc?",
			},
		},
		Resource: ResourceOboToken(),
		Delete:   true,
		New:      true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceOboTokenCreateNoLifetimeOrComment(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token-management/on-behalf-of/tokens",
				ExpectedRequest: OboToken{
					ApplicationID: "abc",
				},
				Response: TokenResponse{
					TokenValue: "s#Cr3t!11",
					TokenInfo: &TokenInfo{
						TokenID:    "bcd",
						ExpiryTime: time.Now().UnixMilli() + 1000,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/bcd",
				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						TokenID:    "bcd",
						ExpiryTime: time.Now().UnixMilli() + 1000,
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Create:   true,
		HCL: `
		application_id = "abc"
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "bcd",
	})
}
