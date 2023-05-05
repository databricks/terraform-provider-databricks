package tokens

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceOboTokenRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/abc",

				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment: "Hello, world!",
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
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
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
	d, err := qa.ResourceFixture{
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
						TokenID: "bcd",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/bcd",
				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						Comment: "Hello, world!",
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
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "bcd", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
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
	d, err := qa.ResourceFixture{
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
						TokenID: "bcd",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/token-management/tokens/bcd",
				Response: TokenResponse{
					TokenInfo: &TokenInfo{
						TokenID: "bcd",
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
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "bcd", d.Id(), "Id should not be empty")
}
