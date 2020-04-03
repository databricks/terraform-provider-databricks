package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	group, err := client.Groups().Create("my-test-group", nil)
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
	//t.Log(group.ID)
	err = client.Groups().Update(group.ID, []string{user.ID, user2.ID}, nil)
	assert.NoError(t, err, err)

	err = client.Groups().Update(group.ID, nil, []string{user.ID})
	assert.NoError(t, err, err)

	group, err = client.Groups().Read(group.ID)
	assert.True(t, len(group.Members) == 1)
	assert.True(t, group.Members[0].Value == user2.ID)
}
