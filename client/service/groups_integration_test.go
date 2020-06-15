package service

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestFmt(t *testing.T) {
	str := fmt.Sprintf("members[value eq \"%s\"]", "1000")
	t.Log(str)
}

func TestGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()

	user, err := client.Users().Create("test-acc@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	user2, err := client.Users().Create("test-acc2@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	//Create empty group
	group, err := client.Groups().Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(group.ID)
		assert.NoError(t, err, err)
		err = client.Users().Delete(user.ID)
		assert.NoError(t, err, err)
		err = client.Users().Delete(user2.ID)
		assert.NoError(t, err, err)
	}()

	group, err = client.Groups().Read(group.ID)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(group.ID, []string{user.ID, user2.ID}, nil, model.GroupMembersPath)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(group.ID, nil, []string{user.ID}, model.GroupMembersPath)
	assert.NoError(t, err, err)

	group, err = client.Groups().Read(group.ID)
	assert.NoError(t, err, err)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}

func TestGetAdminGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	grp, err := client.Groups().GetAdminGroup()
	assert.NoError(t, err, err)
	assert.NotNil(t, grp)
	assert.True(t, len(grp.ID) > 0)
}

func TestReadInheritedRolesFromGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := GetIntegrationDBAPIClient()
	myTestRole := "arn:aws:iam::123456789012:instance-profile/go-sdk-integeration-testing"
	err := client.InstanceProfiles().Create(myTestRole, true)
	assert.NoError(t, err, err)
	defer func() {
		err := client.InstanceProfiles().Delete(myTestRole)
		assert.NoError(t, err, err)
	}()

	myTestGroup, err := client.Groups().Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(myTestGroup.ID)
		assert.NoError(t, err, err)
	}()

	myTestSubGroup, err := client.Groups().Create("my-test-sub-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := client.Groups().Delete(myTestSubGroup.ID)
		assert.NoError(t, err, err)
	}()

	err = client.Groups().Patch(myTestGroup.ID, []string{myTestRole}, nil, model.GroupRolesPath)
	assert.NoError(t, err, err)

	err = client.Groups().Patch(myTestGroup.ID, []string{myTestSubGroup.ID}, nil, model.GroupMembersPath)
	assert.NoError(t, err, err)

	myTestGroupInfo, err := client.Groups().Read(myTestSubGroup.ID)
	assert.NoError(t, err, err)

	assert.True(t, len(myTestGroupInfo.InheritedRoles) > 0)
	assert.True(t, func(roles []model.RoleListItem, testRole string) bool {
		for _, role := range roles {
			if role.Value == testRole {
				return true
			}
		}
		return false
	}(myTestGroupInfo.InheritedRoles, myTestRole))
}
