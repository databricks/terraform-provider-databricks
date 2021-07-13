package identity

import (
	"testing"

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
					Entitlements: []valueItem{
						{"allow-cluster-create"},
						{"allow-instance-pool-create"},
						{"databricks-sql-access"},
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
					Entitlements: []valueItem{
						{"allow-cluster-create"},
						{"databricks-sql-access"},
						{"allow-instance-pool-create"},
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
	qa.ResourceFixture{
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
	}.ExpectError(t, "Internal error happened")
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
					Entitlements: []valueItem{
						{"databricks-sql-access"},
						{"allow-cluster-create"},
						{"allow-instance-pool-create"},
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
	qa.ResourceFixture{
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
	}.ExpectError(t, "Internal error happened")
}

func TestResourceGroupUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Members: []GroupMember{
						{
							Display: "scotchmo",
						},
					},
					Roles: []valueItem{
						{"reader"},
					},
					Groups: []GroupMember{
						{
							Display: "Rangers",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: ScimGroup{
					DisplayName: "Data Ninjas",
					Entitlements: entitlements{
						{"allow-cluster-create"},
						{"allow-instance-pool-create"},
						{"databricks-sql-access"},
					},
					Members: []GroupMember{
						{
							Display: "scotchmo",
						},
					},
					Roles: []valueItem{
						{"reader"},
					},
					Groups: []GroupMember{
						{
							Display: "Rangers",
						},
					},
					Schemas: []URN{GroupSchema},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					DisplayName: "Data Ninjas",
					Entitlements: entitlements{
						{"allow-cluster-create"},
						{"allow-instance-pool-create"},
						{"databricks-sql-access"},
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

func TestResourceGroupUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
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
		State: map[string]interface{}{
			"display_name":               "Data Ninjas",
			"allow_instance_pool_create": true,
		},
		Update: true,
		ID:     "abc",
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
	}.ExpectError(t, "Internal error happened")
}
