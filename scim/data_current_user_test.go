package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceCurrentUser(t *testing.T) {
	userName := "mr.test@example.com"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: User{
					ID:       "123",
					UserName: userName,
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentUser(),
		ID:          ".",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "123", d.Id())
	assert.Equal(t, d.Get("user_name"), userName)
	assert.Equal(t, d.Get("home"), "/Users/"+userName)
	assert.Equal(t, d.Get("repos"), "/Repos/"+userName)
	assert.Equal(t, d.Get("acl_principal_id"), "users/"+userName)
	assert.Equal(t, d.Get("alphanumeric"), "mr_test")
}

func TestDataSourceCurrentUserAsSP(t *testing.T) {
	spId := "11111111-2222-3333-4444-555666777888"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: User{
					ID:       "123",
					UserName: spId,
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceCurrentUser(),
		ID:          ".",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "123", d.Id())
	assert.Equal(t, d.Get("user_name"), spId)
	assert.Equal(t, d.Get("home"), "/Users/"+spId)
	assert.Equal(t, d.Get("repos"), "/Repos/"+spId)
	assert.Equal(t, d.Get("acl_principal_id"), "servicePrincipals/"+spId)
}
