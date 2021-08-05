package identity

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAwsAccOboFlow(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	tmAPI := NewTokenManagementAPI(ctx, client)
	spAPI := NewServicePrincipalsAPI(ctx, client)
	groupsAPI := NewGroupsAPI(ctx, client)

	sp, err := spAPI.Create(ScimUser{
		DisplayName: qa.RandomName("tf"),
	})
	require.NoError(t, err)
	defer spAPI.Delete(sp.ID)

	admins, err := groupsAPI.ReadByDisplayName("admins")
	require.NoError(t, err)

	// add new SP to admins, as it's the easiest way to give it permissions to do so,
	// because importing access package will create circular import dependency
	err = groupsAPI.Patch(admins.ID, scimPatchRequest("add", "members", sp.ID))
	require.NoError(t, err)

	oboToken, err := tmAPI.CreateTokenOnBehalfOfServicePrincipal(OboToken{
		ApplicationID:   sp.ApplicationID,
		Comment:         qa.RandomLongName(),
		LifetimeSeconds: 60,
	})
	require.NoError(t, err)
	defer tmAPI.Delete(oboToken.TokenInfo.TokenID)

	spClient := &common.DatabricksClient{
		Host:  client.Host,
		Token: oboToken.TokenValue,
	}
	err = spClient.Configure()
	require.NoError(t, err)

	newMe, err := NewUsersAPI(ctx, spClient).Me()
	require.NoError(t, err)
	assert.Equal(t, newMe.DisplayName, sp.DisplayName)

	r, err := tmAPI.Read(oboToken.TokenInfo.TokenID)
	require.NoError(t, err)
	assert.Equal(t, r.TokenInfo.TokenID, oboToken.TokenInfo.TokenID)
}

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
	assert.NoError(t, err, err)
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
				Response: common.APIError{
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
				Response: common.APIError{
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
	assert.NoError(t, err, err)
	assert.Equal(t, "bcd", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
}

func TestResourceOboTokenDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/token-management/tokens/abc",
			},
		},
		Resource: ResourceOboToken(),
		Delete:   true,
		New:      true,
		ID:       "abc",
	}.ApplyNoError(t)
}
