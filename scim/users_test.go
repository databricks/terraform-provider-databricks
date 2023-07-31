package scim

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersFilter(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles",

			Response: UserList{
				Resources: []User{
					{UserName: "me"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20somebody",
			Response: UserList{},
		},
	})
	require.NoError(t, err)
	defer server.Close()
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	users, err := usersAPI.Filter("", true)
	require.NoError(t, err)
	assert.Len(t, users, 1)

	users, err = usersAPI.Filter("userName eq somebody", true)
	require.NoError(t, err)
	assert.Len(t, users, 0)
}
