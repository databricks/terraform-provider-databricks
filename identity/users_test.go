package identity

import (
	"fmt"
	"log"
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
	me, err := NewUsersAPI(client).Me()
	assert.NoError(t, err, err)

	if strings.Contains(me.UserName, "@") {
		// let's assume that service principals do not look like emails
		ru, err := NewUsersAPI(client).ReadR(me.ID)
		assert.NoError(t, err, err)
		assert.NotNil(t, ru)
	}
}

func TestAccCreateRUserNonAdmin(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	given := UserEntity{
		DisplayName:        "Mr " + randomName,
		UserName:           fmt.Sprintf("test+%s@example.com", randomName),
		AllowClusterCreate: true,
	}
	meh, err := NewUsersAPI(client).CreateR(given)
	assert.NoError(t, err, err)

	ru, err := NewUsersAPI(client).ReadR(meh.ID)
	assert.NoError(t, err, err)
	assert.NotNil(t, ru)

	assert.Equal(t, given.UserName, ru.UserName)
	assert.Equal(t, given.DisplayName, ru.DisplayName)
	assert.True(t, ru.AllowClusterCreate)
}

func TestAccCreateUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	user, err := NewUsersAPI(client).Create(
		fmt.Sprintf("test+%s@example.com", randomName), "Display Name", nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := NewUsersAPI(client).Delete(idToDelete)
		assert.NoError(t, err, err)
	}()

	user, err = NewUsersAPI(client).Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	err = NewUsersAPI(client).Update(user.ID, fmt.Sprintf("updated+%s@example.com", randomName),
		"Test User", []string{string(AllowClusterCreateEntitlement)}, nil)
	//t.Log(user)
	assert.NoError(t, err, err)
}

func TestAccCreateAdminUser(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	user, err := NewUsersAPI(client).Create(
		fmt.Sprintf("terraform+%s@example.com", randomName),
		"Terra "+randomName, nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := NewUsersAPI(client).Delete(idToDelete)
		assert.NoError(t, err, err)
	}()
	log.Println(idToDelete)

	user, err = NewUsersAPI(client).Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	group, err := NewGroupsAPI(client).GetAdminGroup()
	assert.NoError(t, err, err)

	adminGroupID := group.ID

	err = NewUsersAPI(client).SetUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err := NewUsersAPI(client).VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == true)
	log.Println(userIsAdmin)

	err = NewUsersAPI(client).RemoveUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err = NewUsersAPI(client).VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == false)
	log.Println(userIsAdmin)
}

func TestAccRoleDifferences(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	user, err := NewUsersAPI(client).Create(
		fmt.Sprintf("terraform+%s@example.com", randomName),
		"Terra "+randomName, nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID

	user, err = NewUsersAPI(client).Read(idToDelete)
	assert.NoError(t, err, err)
	t.Log(user.Roles)
	t.Log(user.Groups)
	t.Log(user.InheritedRoles)
	t.Log(user.UnInheritedRoles)

	err = NewUsersAPI(client).Delete(idToDelete)
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
	users, err := NewUsersAPI(client).Filter("")
	require.NoError(t, err)
	assert.Len(t, users, 1)

	users, err = NewUsersAPI(client).Filter("userName eq somebody")
	require.NoError(t, err)
	assert.Len(t, users, 0)
}
