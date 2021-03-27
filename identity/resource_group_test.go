package identity

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				ExpectedRequest: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Entitlements: []entitlementsListItem{
						{
							AllowClusterCreateEntitlement,
						},
						{
							AllowSQLAnalyticsAccessEntitlement,
						},
						{
							AllowInstancePoolCreateEntitlement,
						},
					},
				},
				Response: ScimGroup{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							AllowClusterCreateEntitlement,
						},
						{
							AllowSQLAnalyticsAccessEntitlement,
						},
						{
							AllowInstancePoolCreateEntitlement,
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
		allow_sql_analytics_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("allow_sql_analytics_access"))
}

func TestResourceGroupCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		State: map[string]interface{}{
			"display_name": "Data Scientists",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceGroupRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							AllowSQLAnalyticsAccessEntitlement,
						},
						{
							AllowClusterCreateEntitlement,
						},
						{
							AllowInstancePoolCreateEntitlement,
						},
					},
				},
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("allow_sql_analytics_access"))
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
}

func TestResourceGroupRead_NoEntitlements(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, false, d.Get("allow_instance_pool_create"))
	assert.Equal(t, false, d.Get("allow_sql_analytics_access"))
	assert.Equal(t, "Data Scientists", d.Get("display_name"))
}

func TestResourceGroupRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceGroupUpdate_AddPerms(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: GroupPatchRequest{
					Schemas: []URN{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
					Operations: []GroupPatchOperations{
						{
							Op:   "add",
							Path: "entitlements",
							Value: []ValueListItem{
								{
									Value: "allow-cluster-create",
								},
								{
									Value: "sql-analytics-access",
								},
								{
									Value: "allow-instance-pool-create",
								},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Ninjas",
					ID:          "abc",
					Entitlements: []entitlementsListItem{
						{
							AllowSQLAnalyticsAccessEntitlement,
						},
						{
							AllowClusterCreateEntitlement,
						},
						{
							AllowInstancePoolCreateEntitlement,
						},
					},
				},
			},
		},
		Resource: ResourceGroup(),
		InstanceState: map[string]string{
			"display_name":               "Data Ninjas",
			"allow_instance_pool_create": "false",
			"allow_cluster_create":       "false",
			"allow_sql_analytics_access": "false",
		},
		HCL: `
		display_name = "Data Ninjas"
		allow_instance_pool_create = true
		allow_cluster_create = true
		allow_sql_analytics_access = true
		`,
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
	assert.Equal(t, "Data Ninjas", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("allow_sql_analytics_access"))
}

func TestResourceGroupUpdate_RemovePerms(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: GroupPatchRequest{
					Schemas: []URN{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
					Operations: []GroupPatchOperations{
						{
							Op:   "remove",
							Path: "entitlements[value eq \"allow-cluster-create\"]",
						},
						{
							Op:   "remove",
							Path: "entitlements[value eq \"sql-analytics-access\"]",
						},
						{
							Op:   "remove",
							Path: "entitlements[value eq \"allow-instance-pool-create\"]",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:      []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName:  "Data Ninjas",
					ID:           "abc",
					Entitlements: []entitlementsListItem{},
				},
			},
		},
		Resource: ResourceGroup(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"display_name":               "Data Ninjas",
			"allow_instance_pool_create": "true",
			"allow_cluster_create":       "true",
			"allow_sql_analytics_access": "true",
		},
		HCL: `
		display_name = "Data Ninjas"
		allow_instance_pool_create = false
		allow_cluster_create = false
		allow_sql_analytics_access = false
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Data Ninjas", d.Get("display_name"))
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, false, d.Get("allow_instance_pool_create"))
	assert.Equal(t, false, d.Get("allow_sql_analytics_access"))
}

func TestResourceGroupUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		State: map[string]interface{}{
			"display_name":               "Data Ninjas",
			"allow_instance_pool_create": true,
		},
		Update: true,
		ID:     "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceGroupDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
			},
		},
		Resource: ResourceGroup(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceGroupDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroup(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
