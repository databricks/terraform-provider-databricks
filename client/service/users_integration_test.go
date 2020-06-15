package service

import (
	"log"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	user, err := client.Users().Create("testuser@databricks.com", "Display Name", nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	defer func() {
		err := client.Users().Delete(idToDelete)
		assert.NoError(t, err, err)
	}()

	user, err = client.Users().Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	err = client.Users().Update(user.ID, "newtestuser@databricks.com", "Test User", []string{string(model.AllowClusterCreateEntitlement)}, nil)
	//t.Log(user)
	assert.NoError(t, err, err)
}

func TestCreateAdminUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	user, err := client.Users().Create("testusersriterraform@databricks.com", "Display Name", nil, nil)
	assert.NoError(t, err, err)
	assert.True(t, len(user.ID) > 0, "User id is empty")
	idToDelete := user.ID
	//defer func() {
	//	err := client.Users().Delete(idToDelete)
	//	assert.NoError(t, err, err)
	//}()
	log.Println(idToDelete)

	user, err = client.Users().Read(user.ID)
	t.Log(user)
	assert.NoError(t, err, err)

	group, err := client.Groups().GetAdminGroup()
	assert.NoError(t, err, err)

	adminGroupID := group.ID

	err = client.Users().SetUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err := client.Users().VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == true)
	log.Println(userIsAdmin)

	err = client.Users().RemoveUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)

	userIsAdmin, err = client.Users().VerifyUserAsAdmin(user.ID, adminGroupID)
	assert.NoError(t, err, err)
	assert.True(t, userIsAdmin == false)
	log.Println(userIsAdmin)
}

// user id 101354
func TestRoleDifferences(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()
	//user, err := client.Users().Create("testusersriterraform@databricks.com", "Display Name", nil, nil)
	//assert.NoError(t, err, err)
	//assert.True(t, len(user.ID) > 0, "User id is empty")
	//idToDelete := user.ID
	//log.Println(idToDelete)

	user, err := client.Users().Read("101354")
	assert.NoError(t, err, err)
	t.Log(user.Roles)
	t.Log(user.Groups)
	t.Log(user.InheritedRoles)
	t.Log(user.UnInheritedRoles)
}
