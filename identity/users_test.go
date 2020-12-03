package identity

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccReadUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	me, err := usersAPI.Me()
	assert.NoError(t, err, err)

	if strings.Contains(me.UserName, "@") {
		// let's assume that service principals do not look like emails
		ru, err := usersAPI.Read(me.ID)
		assert.NoError(t, err, err)
		assert.NotNil(t, ru)
	}
}

func TestAccCreateUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	user, err := usersAPI.Create(UserEntity{
		UserName: fmt.Sprintf("test+%s@example.com", randomName),
	})
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := usersAPI.Delete(idToDelete)
		assert.NoError(t, err, err)
	}()

	user, err = usersAPI.read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	err = usersAPI.Update(user.ID, UserEntity{
		UserName:           fmt.Sprintf("updated+%s@example.com", randomName),
		AllowClusterCreate: true,
		DisplayName:        "TU",
	})
	//t.Log(user)
	assert.NoError(t, err, err)
}

func TestUsersFilter(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?",

			Response: UserList{
				Resources: []ScimUser{
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
