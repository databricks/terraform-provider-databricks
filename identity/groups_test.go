package identity

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/stretchr/testify/assert"
)

func TestAccGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	user, err := NewUsersAPI(client).Create("test-acc@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	user2, err := NewUsersAPI(client).Create("test-acc2@databricks.com", "test account", nil, nil)
	assert.NoError(t, err, err)

	//Create empty group
	group, err := NewGroupsAPI(client).Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(group.ID)
		assert.NoError(t, err, err)
		err = NewUsersAPI(client).Delete(user.ID)
		assert.NoError(t, err, err)
		err = NewUsersAPI(client).Delete(user2.ID)
		assert.NoError(t, err, err)
	}()

	group, err = NewGroupsAPI(client).Read(group.ID)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(group.ID, []string{user.ID, user2.ID}, nil, GroupMembersPath)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(group.ID, nil, []string{user.ID}, GroupMembersPath)
	assert.NoError(t, err, err)

	group, err = NewGroupsAPI(client).Read(group.ID)
	assert.NoError(t, err, err)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}

func TestAccFilterGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := common.NewClientFromEnvironment()
	groupList, err := NewGroupsAPI(client).Filter("displayName eq admins")
	assert.NoError(t, err, err)
	assert.NotNil(t, groupList)
	assert.Len(t, groupList.Resources, 1)
}

func TestAccGetAdminGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := common.NewClientFromEnvironment()
	grp, err := NewGroupsAPI(client).GetAdminGroup()
	assert.NoError(t, err, err)
	assert.NotNil(t, grp)
	assert.True(t, len(grp.ID) > 0)
}

func TestAwsAccReadInheritedRolesFromGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := common.NewClientFromEnvironment()
	// TODO: pass IAM role with ENV variable
	myTestRole := "arn:aws:iam::123456789012:instance-profile/go-sdk-integeration-testing"
	err := NewInstanceProfilesAPI(client).Create(myTestRole, true)
	assert.NoError(t, err, err)
	defer func() {
		err := NewInstanceProfilesAPI(client).Delete(myTestRole)
		assert.NoError(t, err, err)
	}()

	myTestGroup, err := NewGroupsAPI(client).Create("my-test-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(myTestGroup.ID)
		assert.NoError(t, err, err)
	}()

	myTestSubGroup, err := NewGroupsAPI(client).Create("my-test-sub-group", nil, nil, nil)
	assert.NoError(t, err, err)

	defer func() {
		err := NewGroupsAPI(client).Delete(myTestSubGroup.ID)
		assert.NoError(t, err, err)
	}()

	err = NewGroupsAPI(client).Patch(myTestGroup.ID, []string{myTestRole}, nil, GroupRolesPath)
	assert.NoError(t, err, err)

	err = NewGroupsAPI(client).Patch(myTestGroup.ID, []string{myTestSubGroup.ID}, nil, GroupMembersPath)
	assert.NoError(t, err, err)

	myTestGroupInfo, err := NewGroupsAPI(client).Read(myTestSubGroup.ID)
	assert.NoError(t, err, err)

	assert.True(t, len(myTestGroupInfo.InheritedRoles) > 0)
	assert.True(t, func(roles []RoleListItem, testRole string) bool {
		for _, role := range roles {
			if role.Value == testRole {
				return true
			}
		}
		return false
	}(myTestGroupInfo.InheritedRoles, myTestRole))
}
