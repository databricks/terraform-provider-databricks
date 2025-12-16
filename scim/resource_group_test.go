package scim

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				ExpectedRequest: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
						},
					},
				},
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
						{
							Value: "databricks-sql-access",
						},
						{
							Value: "allow-instance-pool-create",
						},
					},
				},
			},
		},
		Resource: ResourceGroup(),
		HCL: `
		display_name = "Data Scientists"
		allow_instance_pool_create = true
		allow_cluster_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
	assert.Equal(t, "groups/Data Scientists", d.Get("acl_principal_id"))
}

func TestResourceGroupCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		State: map[string]any{
			"display_name": "Data Scientists",
		},
		Create: true,
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
					Entitlements: []ComplexValue{
						{
							Value: "databricks-sql-access",
						},
						{
							Value: "allow-cluster-create",
						},
						{
							Value: "allow-instance-pool-create",
						},
					},
				},
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
}

func TestResourceGroupRead_NoEntitlements(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroup(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, false, d.Get("allow_instance_pool_create"))
	assert.Equal(t, false, d.Get("databricks_sql_access"))
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
}

func TestResourceGroupRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceGroupRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		ID:       "abc",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,entitlements,groups,members,externalId",
				Response: Group{
					Members: []ComplexValue{
						{
							Display: "scotchmo",
						},
					},
					Roles: []ComplexValue{
						{
							Value: "reader",
						},
					},
					Groups: []ComplexValue{
						{
							Display: "Rangers",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: Group{
					DisplayName: "Data Ninjas",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
						},
					},
					Members: []ComplexValue{
						{
							Display: "scotchmo",
						},
					},
					Roles: []ComplexValue{
						{
							Value: "reader",
						},
					},
					Groups: []ComplexValue{
						{
							Display: "Rangers",
						},
					},
					Schemas: []URN{GroupSchema},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					DisplayName: "Data Ninjas",
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
						},
					},
					// we don't care about other fields in this response
				},
			},
		},
		Resource: ResourceGroup(),
		HCL: `
		display_name = "Data Ninjas"
		allow_instance_pool_create = true
		allow_cluster_create = true
		databricks_sql_access = true
		`,
		RequiresNew: false,
		Update:      true,
		ID:          "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Data Ninjas", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceGroupUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,entitlements,groups,members,externalId",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		State: map[string]any{
			"display_name":               "Data Ninjas",
			"allow_instance_pool_create": true,
		},
		Update:      true,
		RequiresNew: true,
		ID:          "abc",
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
			},
		},
		Resource: ResourceGroup(),
		Delete:   true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceGroupDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		Delete:   true,
		ID:       "abc",
	}.ExpectError(t, "Internal error happened")
}

func TestCreateForceOverwriteCannotListGroups(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22abc%22",
			Status:   417,
			Response: apierr.APIError{
				Message: "cannot find group",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceGroup().ToResource().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedGroup(
			fmt.Errorf("Group with name abc already exists."),
			d, NewGroupsAPI(ctx, client), Group{
				DisplayName: "abc",
			})
		assert.EqualError(t, err, "cannot find group")
	})
}

func TestCreateForceOverwriteFindsAndSetsGroupID(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22abc%22",
			Response: GroupList{
				Resources: []Group{
					{
						ID: "123",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups/123?attributes=displayName,entitlements,groups,members,externalId",
			Response: Group{
				ID: "123",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/preview/scim/v2/Groups/123",
			ExpectedRequest: Group{
				Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
				DisplayName: "abc",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceGroup().ToResource().TestResourceData()
		d.Set("force", true)
		d.Set("display_name", "abc")
		err := createForceOverridesManuallyAddedGroup(
			fmt.Errorf("Group with name abc already exists."),
			d, NewGroupsAPI(ctx, client), Group{
				DisplayName: "abc",
			})
		assert.NoError(t, err)
		assert.Equal(t, "123", d.Id())
	})
}
