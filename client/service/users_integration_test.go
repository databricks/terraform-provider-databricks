package service

import (
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"github.com/stretchr/testify/assert"
	"testing"
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
	t.Log(user)
	assert.NoError(t, err, err)

}
