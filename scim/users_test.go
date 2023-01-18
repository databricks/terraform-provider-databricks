package scim

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccReadUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	client := common.NewClientFromEnvironment()
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	me, err := usersAPI.Me()
	assert.NoError(t, err)

	if strings.Contains(me.UserName, "@") {
		// let's assume that service principals do not look like emails
		ru, err := usersAPI.Read(me.ID)
		assert.NoError(t, err)
		assert.NotNil(t, ru)
	}
}

func TestAccCreateUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	client := common.NewClientFromEnvironment()
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	user, err := usersAPI.Create(User{
		UserName: qa.RandomEmail(),
	})
	assert.NoError(t, err)
	require.True(t, len(user.ID) > 0, "User id is empty")
	defer usersAPI.Delete(user.ID)

	user, err = usersAPI.Read(user.ID)
	t.Log(user)
	assert.NoError(t, err)

	err = usersAPI.Update(user.ID, User{
		UserName:    qa.RandomEmail(),
		DisplayName: "TestAccCreateUser",
	})
	assert.NoError(t, err)
}

func TestUsersFilter(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?",

			Response: UserList{
				Resources: []User{
					{UserName: "me"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20somebody",
			Response: UserList{},
		},
	})
	require.NoError(t, err)
	defer server.Close()
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	users, err := usersAPI.Filter("")
	require.NoError(t, err)
	assert.Len(t, users, 1)

	users, err = usersAPI.Filter("userName eq somebody")
	require.NoError(t, err)
	assert.Len(t, users, 0)
}
