package identity

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
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
		ru, err := NewUsersAPI(client).Read(me.ID)
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
	meh, err := NewUsersAPI(client).Create(given)
	assert.NoError(t, err, err)

	ru, err := NewUsersAPI(client).Read(meh.ID)
	assert.NoError(t, err, err)
	assert.NotNil(t, ru)

	assert.Equal(t, given.UserName, ru.UserName)
	assert.Equal(t, given.DisplayName, ru.DisplayName)
	assert.True(t, ru.AllowClusterCreate)
}
