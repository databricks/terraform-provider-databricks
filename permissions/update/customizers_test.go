package update

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCurrentUserAsManage_CaseInsensitiveMatch(t *testing.T) {
	acl, err := AddCurrentUserAsManage(ACLCustomizerContext{
		GetCurrentUser: func() (string, error) {
			return "User@Example.com", nil
		},
	}, []iam.AccessControlRequest{
		{
			UserName:        "user@example.com",
			PermissionLevel: "CAN_USE",
		},
	})

	require.NoError(t, err)
	require.Len(t, acl, 1)
	assert.Equal(t, "user@example.com", acl[0].UserName)
	assert.Equal(t, iam.PermissionLevelCanUse, acl[0].PermissionLevel)
}

func TestAddCurrentUserAsManage_AppendsWhenMissing(t *testing.T) {
	acl, err := AddCurrentUserAsManage(ACLCustomizerContext{
		GetCurrentUser: func() (string, error) {
			return "user@example.com", nil
		},
	}, []iam.AccessControlRequest{
		{
			UserName:        "another@example.com",
			PermissionLevel: "CAN_USE",
		},
	})

	require.NoError(t, err)
	require.Len(t, acl, 2)
	assert.Equal(t, "user@example.com", acl[1].UserName)
	assert.Equal(t, iam.PermissionLevelCanManage, acl[1].PermissionLevel)
}
