package scim

import (
	"context"
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceUserRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: User{
					ID:          "abc",
					DisplayName: "Example user",
					UserName:    "me@example.com",
					Groups: []ComplexValue{
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
				ExpectedRequest: User{
					DisplayName: "Example user",
					Active:      true,
					Entitlements: entitlements{
						{
							Value: "allow-cluster-create",
						},
					},
					UserName: "me@example.com",
					Schemas:  []URN{UserSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: User{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: entitlements{
						{
							Value: "allow-cluster-create",
						},
					},
					Groups: []ComplexValue{
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
	newUser := User{
		Schemas:     []URN{UserSchema},
		DisplayName: "Changed Name",
		UserName:    "me@example.com",
		Active:      true,
		Entitlements: entitlements{
			{
				Value: "allow-instance-pool-create",
			},
		},
		Groups: []ComplexValue{
			{
				Display: "admins",
				Value:   "4567",
			},
			{
				Display: "ds",
				Value:   "9877",
			},
		},
		Roles: []ComplexValue{
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
				Response: User{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
					},
					Groups: []ComplexValue{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
					Roles: []ComplexValue{
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
		InstanceState: map[string]string{
			"user_name":    "me@example.com",
			"display_name": "Old Name",
		},
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
				Response: User{
					DisplayName: "Example user",
					Active:      true,
					UserName:    "me@example.com",
					ID:          "abc",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
					},
					Groups: []ComplexValue{
						{
							Display: "admins",
							Value:   "4567",
						},
						{
							Display: "ds",
							Value:   "9877",
						},
					},
					Roles: []ComplexValue{
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
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
	}.ApplyNoError(t)
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

func TestCreateForceOverridesManuallyAddedUserErrorNotMatched(t *testing.T) {
	d := ResourceUser().TestResourceData()
	d.Set("force", true)
	rerr := createForceOverridesManuallyAddedUser(
		fmt.Errorf("nonsense"), d,
		NewUsersAPI(context.Background(), &common.DatabricksClient{}), User{})
	assert.EqualError(t, rerr, "nonsense")
}

func TestCreateForceOverwriteCannotListUsers(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27me%40example.com%27",
			Status:   417,
			Response: common.APIError{
				Message: "cannot find user",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedUser(
			fmt.Errorf("User with username me@example.com already exists."),
			d, NewUsersAPI(ctx, client), User{
				UserName: "me@example.com",
			})
		assert.EqualError(t, err, "cannot find user")
	})
}

func TestCreateForceOverwriteCannotListAccUsers(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27me%40example.com%27",
			Response: UserList{
				TotalResults: 0,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedUser(
			fmt.Errorf("User already exists in another account"),
			d, NewUsersAPI(ctx, client), User{
				UserName: "me@example.com",
			})
		assert.EqualError(t, err, "cannot find me@example.com for force import")
	})
}

func TestCreateForceOverwriteFindsAndSetsID(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27me%40example.com%27",
			Response: UserList{
				Resources: []User{
					{
						ID: "abc",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users/abc",
			Response: User{
				ID: "abc",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/preview/scim/v2/Users/abc",
			ExpectedRequest: User{
				Schemas:  []URN{UserSchema},
				UserName: "me@example.com",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		d.Set("user_name", "me@example.com")
		err := createForceOverridesManuallyAddedUser(
			fmt.Errorf("User with username me@example.com already exists."),
			d, NewUsersAPI(ctx, client), User{
				UserName: "me@example.com",
			})
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id())
	})
}

func TestCreateForceOverwriteFindsAndSetsAccID(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27me%40example.com%27",
			Response: UserList{
				Resources: []User{
					{
						ID: "abc",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users/abc",
			Response: User{
				ID: "abc",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/preview/scim/v2/Users/abc",
			ExpectedRequest: User{
				Schemas:  []URN{UserSchema},
				UserName: "me@example.com",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		d.Set("user_name", "me@example.com")
		err := createForceOverridesManuallyAddedUser(
			fmt.Errorf("User already exists in another account"),
			d, NewUsersAPI(ctx, client), User{
				UserName: "me@example.com",
			})
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id())
	})
}
