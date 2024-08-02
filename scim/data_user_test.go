package scim

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceUser(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22ds%22",
				Response: UserList{
					Resources: []User{
						{
							ID:       "123",
							UserName: "mr.test@example.com",
							Active:   true,
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceUser(),
		ID:          ".",
		State: map[string]any{
			"user_name": "ds",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "123", d.Id())
	assert.Equal(t, d.Get("user_name"), "mr.test@example.com")
	assert.Equal(t, d.Get("home"), "/Users/mr.test@example.com")
	assert.Equal(t, d.Get("acl_principal_id"), "users/mr.test@example.com")
	assert.Equal(t, d.Get("alphanumeric"), "mr_test")
	assert.Equal(t, d.Get("active"), true)
}

func TestDataSourceUserGerUser(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users/a?attributes=userName,displayName,externalId,applicationId",
			Response: User{
				ID: "a",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22searching_error%22",
			Status:   404,
			Response: apierr.APIError{
				Message: "searching_error",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22empty_search%22",
			Response: UserList{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		usersAPI := NewUsersAPI(ctx, client)
		user, err := getUser(usersAPI, "a", "")
		assert.NoError(t, err)
		assert.Equal(t, "a", user.ID)

		_, err = getUser(usersAPI, "", "searching_error")
		assert.EqualError(t, err, "searching_error")

		_, err = getUser(usersAPI, "", "empty_search")
		assert.EqualError(t, err, "cannot find user empty_search")
	})
}
