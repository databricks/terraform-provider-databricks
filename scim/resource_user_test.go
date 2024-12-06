package scim

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceUserRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
	assert.Equal(t, "/Users/me@example.com", d.Get("home"))
	assert.Equal(t, "/Repos/me@example.com", d.Get("repos"))
	assert.Equal(t, "users/me@example.com", d.Get("acl_principal_id"))
}

func TestResourceUserRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
	assert.Equal(t, "/Users/me@example.com", d.Get("home"))
	assert.Equal(t, "/Repos/me@example.com", d.Get("repos"))
}

func TestResourceUserCreateInactive(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Users",
				ExpectedRequest: User{
					DisplayName: "Example user",
					Active:      false,
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					DisplayName: "Example user",
					Active:      false,
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
		active = false
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=groups,roles",
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
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

func TestResourceUserDelete_NoError(t *testing.T) {
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
		HCL: `
			user_name = "abc",
			force_delete_repos = false,
			force_delete_home_dir = false 
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: qa.HTTPFailures,
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
	}.ExpectError(t, "i'm a teapot")
}

func TestResourceUserDelete_NoErrorEmtpyParams(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Repos/abc",
					Recursive: true,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Users/abc",
					Recursive: true,
				},
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
		`,
	}.ApplyNoError(t)
}

func TestResourceUserforce_delete_reposError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Repos/abc",
					Recursive: true,
				},
				Status: 400,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			force_delete_repos = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}
func TestResourceUserDelete_NonExistingRepo(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Repos/abc",
					Recursive: true,
				},
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Path (/Repos/abc) doesn't exist.",
				},
				Status: 400,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			force_delete_repos = true
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDelete_DirError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Users/abc",
					Recursive: true,
				},
				Status: 400,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			force_delete_home_dir = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}
func TestResourceUserDelete_NonExistingDir(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Users/abc",
					Recursive: true,
				},
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Path (/Users/abc) doesn't exist.",
				},
				Status: 400,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			force_delete_home_dir = true
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDelete_ForceDeleteHomeDir(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: workspace.DeletePath{
					Path:      "/Users/abc",
					Recursive: true,
				},
				Status: 200,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			force_delete_home_dir = true
		`,
	}.ApplyNoError(t)
}

func TestCreateForceOverridesManuallyAddedUserErrorNotMatched(t *testing.T) {
	d := ResourceUser().ToResource().TestResourceData()
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
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22me%40example.com%22",
			Status:   417,
			Response: common.APIErrorBody{
				Message: "cannot find user",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().ToResource().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedUser(
			errors.New(userExistsErrorMessage("me@example.com", false)),
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
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22me%40example.com%22",
			Response: UserList{
				TotalResults: 0,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().ToResource().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedUser(
			errors.New(userExistsErrorMessage("me@example.com", true)),
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
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22me%40example.com%22",
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
			Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=groups,roles",
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
		d := ResourceUser().ToResource().TestResourceData()
		d.Set("force", true)
		d.Set("user_name", "me@example.com")
		err := createForceOverridesManuallyAddedUser(
			errors.New(userExistsErrorMessage("Me@Example.Com", false)),
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
			Resource: "/api/2.0/preview/scim/v2/Users?excludedAttributes=roles&filter=userName%20eq%20%22me%40example.com%22",
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
			Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=groups,roles",
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
		d := ResourceUser().ToResource().TestResourceData()
		d.Set("force", true)
		d.Set("user_name", "me@example.com")
		err := createForceOverridesManuallyAddedUser(
			errors.New(userExistsErrorMessage("me@example.com", true)),
			d, NewUsersAPI(ctx, client), User{
				UserName: "me@example.com",
			})
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id())
	})
}

func TestUserResource_SparkConfDiffSuppress(t *testing.T) {
	jr := ResourceUser()
	scs := jr.Schema["user_name"]
	assert.True(t, scs.DiffSuppressFunc("user_name", "abcdef@example.com", "AbcDef@example.com", nil))
	assert.False(t, scs.DiffSuppressFunc("user_name", "abcdef@example.com", "abcdef2@example.com", nil))
}
