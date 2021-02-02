package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceUserRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: ScimUser{
					ID:          "abc",
					DisplayName: "Example user",
					UserName:    "me@example.com",
					Groups: []groupsListItem{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
				},
			},
		},
		Resource: ResourceUser(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "me@example.com", d.Get("user_name"))
	assert.Equal(t, "Example user", d.Get("display_name"))
	assert.Equal(t, false, d.Get("allow_cluster_create"))
}

func TestResourceUserRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Status:   404,
			},
		},
		Resource: ResourceUser(),
		New:      true,
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceUserRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Status:   400,
				Response: common.APIErrorBody{
					ScimDetail: "Something",
					ScimStatus: "Else",
				},
			},
		},
		Resource: ResourceUser(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	require.Error(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceUserCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Users",
				ExpectedRequest: ScimUser{
					DisplayName: "Example user",
					Active:      true,
					Entitlements: []entitlementsListItem{
						{
							Value: "allow-cluster-create",
						},
					},
					UserName: "me@example.com",
					Schemas:  []URN{UserSchema},
				},
				Response: ScimUser{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: ScimUser{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							Value: AllowClusterCreateEntitlement,
						},
					},
					Groups: []groupsListItem{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
				},
			},
		},
		Resource: ResourceUser(),
		Create:   true,
		HCL: `
		user_name    = "me@example.com"
		display_name = "Example user"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "me@example.com", d.Get("user_name"))
	assert.Equal(t, "Example user", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
}

func TestResourceUserCreate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Users",
				Status:   400,
			},
		},
		Resource: ResourceUser(),
		Create:   true,
		HCL: `
		user_name    = "me@example.com"
		display_name = "Example user"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceUserUpdate(t *testing.T) {
	newUser := ScimUser{
		Schemas:     []URN{UserSchema},
		DisplayName: "Changed Name",
		UserName:    "me@example.com",
		Active:      true,
		Entitlements: []entitlementsListItem{
			{
				Value: AllowInstancePoolCreateEntitlement,
			},
		},
		Groups: []groupsListItem{
			{
				Display: "admins",
				Value:   "4567",
			},
			{
				Display: "ds",
				Value:   "9877",
			},
		},
		Roles: []roleListItem{
			{
				Value: "a",
			},
			{
				Value: "b",
			},
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: ScimUser{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							Value: AllowClusterCreateEntitlement,
						},
					},
					Groups: []groupsListItem{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
					Roles: []roleListItem{
						{
							Value: "a",
						},
						{
							Value: "b",
						},
					},
				},
			},
			{
				Method:          "PUT",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: newUser,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: newUser,
			},
		},
		Resource: ResourceUser(),
		Update:   true,
		ID:       "abc",
		HCL: `
		user_name    = "me@example.com"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "me@example.com", d.Get("user_name"))
	assert.Equal(t, "Changed Name", d.Get("display_name"))
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
}

func TestResourceUserUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Status:   400,
			},
		},
		Resource: ResourceUser(),
		Update:   true,
		ID:       "abc",
		HCL: `
		user_name    = "me@example.com"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceUserUpdate_ErrorPut(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: ScimUser{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							Value: AllowClusterCreateEntitlement,
						},
					},
					Groups: []groupsListItem{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
					Roles: []roleListItem{
						{
							Value: "a",
						},
						{
							Value: "b",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Status:   400,
			},
		},
		Resource: ResourceUser(),
		Update:   true,
		ID:       "abc",
		HCL: `
		user_name    = "me@example.com"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceUserDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
}

func TestResourceUserDelete_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Status:   400,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	require.Error(t, err, err)
}
