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

func TestAccGroup(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	ctx := context.Background()
	usersAPI := NewUsersAPI(ctx, client)
	groupsAPI := NewGroupsAPI(ctx, client)

	user, err := usersAPI.Create(UserEntity{UserName: qa.RandomEmail()})
	require.NoError(t, err, err)

	user2, err := usersAPI.Create(UserEntity{UserName: qa.RandomEmail()})
	require.NoError(t, err, err)

	//Create empty group
	group, err := groupsAPI.Create(qa.RandomName("tf-"), nil, nil, nil)
	require.NoError(t, err, err)

	defer func() {
		err := groupsAPI.Delete(group.ID)
		assert.NoError(t, err, err)
		err = usersAPI.Delete(user.ID)
		assert.NoError(t, err, err)
		err = usersAPI.Delete(user2.ID)
		assert.NoError(t, err, err)
	}()

	group, err = groupsAPI.Read(group.ID)
	require.NoError(t, err, err)

	err = groupsAPI.Patch(group.ID, []string{user.ID, user2.ID}, nil, GroupMembersPath)
	assert.NoError(t, err, err)

	err = groupsAPI.Patch(group.ID, nil, []string{user.ID}, GroupMembersPath)
	assert.NoError(t, err, err)

	group, err = groupsAPI.Read(group.ID)
	assert.NoError(t, err, err)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}

func TestAccFilterGroup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := common.NewClientFromEnvironment()
	ctx := context.Background()
	groupsAPI := NewGroupsAPI(ctx, client)
	groupList, err := groupsAPI.Filter("displayName eq admins")
	assert.NoError(t, err, err)
	assert.NotNil(t, groupList)
	assert.Len(t, groupList.Resources, 1)
}
