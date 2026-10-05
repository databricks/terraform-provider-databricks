package tokens

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
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

func TestResourceOboTokenRead_NoExpire(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",

				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment:    "Hello, world!",
						ExpiryTime: -1,
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

func TestResourceOboTokenCreateWithScopesAndAutoscopeDisabled(t *testing.T) {
	autoscopeEnabled := false
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/token-management/on-behalf-of/tokens",
				ExpectedRequest: oboTokenRequest{
					OboToken: OboToken{
						ApplicationID:   "abc",
						LifetimeSeconds: 60,
						Scopes:          []string{"sql", "unity-catalog"},
					},
					AutoscopeEnabled: &autoscopeEnabled,
				},
				Response: TokenResponse{
					TokenValue: "s#Cr3t!11",
					TokenInfo: &TokenInfo{
						TokenID:    "bcd",
						ExpiryTime: time.Now().UnixMilli() + 1000,
						Scopes:     []string{"sql", "unity-catalog"},
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
						Scopes:     []string{"sql", "unity-catalog"},
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Create:   true,
		HCL: `
		application_id = "abc"
		lifetime_seconds = 60
		scopes = ["sql", "unity-catalog"]
		autoscope_enabled = false
		`,
		New: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "bcd", d.Id())
	assert.ElementsMatch(t, []any{"sql", "unity-catalog"}, d.Get("scopes").(*schema.Set).List())
	assert.Equal(t, false, d.Get("autoscope_enabled"))
}

func TestResourceOboTokenRead_ScopesChanged(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",
				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						ExpiryTime: time.Now().UnixMilli() + 1000,
						Scopes:     []string{"sql"},
					},
				},
			},
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []any{"sql"}, d.Get("scopes").(*schema.Set).List())
}

func TestResourceOboTokenScopesChangeRequiresNew(t *testing.T) {
	r := ResourceOboToken().ToResource()
	r.CustomizeDiff = nil
	state := &terraform.InstanceState{
		ID: "bcd",
		Attributes: map[string]string{
			"id":             "bcd",
			"application_id": "abc",
			"scopes.#":       "1",
			"scopes.0":       "sql",
		},
	}
	config := terraform.NewResourceConfigRaw(map[string]any{
		"application_id": "abc",
		"scopes":         []any{"sql", "unity-catalog"},
	})
	diff, err := r.SimpleDiff(context.Background(), state, config, nil)
	assert.NoError(t, err)
	assert.True(t, diff.RequiresNew())
}
