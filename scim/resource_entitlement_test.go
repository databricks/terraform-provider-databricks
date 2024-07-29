package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var oldGroup = Group{
	Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
	DisplayName: "Data Scientists",
	ID:          "abc",
	Entitlements: []ComplexValue{
		{
			Value: "allow-cluster-create",
		},
	},
}

var newGroup = Group{
	Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
	DisplayName: "Data Scientists",
	ID:          "abc",
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
}

var emptyGroup = Group{
	Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
	DisplayName: "Data Scientists",
	ID:          "abc",
}

var addRequest = PatchRequestComplexValue([]patchOperation{
	{
		"replace", "entitlements", []ComplexValue{
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
})

var emptyAddRequest = PatchRequestComplexValue([]patchOperation{
	{
		"replace", "entitlements", []ComplexValue{
			{
				Value: "",
			},
		},
	},
})

var updateRequest = PatchRequestComplexValue([]patchOperation{
	{
		"replace", "entitlements", []ComplexValue{
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
})

var deleteRequest = PatchRequestComplexValue([]patchOperation{
	{
		"remove", "entitlements", []ComplexValue{
			{
				Value: "allow-cluster-create",
			},
		},
	},
})

func TestResourceEntitlementsGroupCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: addRequest,
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: newGroup,
			},
		},
		Resource: ResourceEntitlements(),
		HCL: `
		group_id = "abc"
		allow_instance_pool_create = true
		allow_cluster_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "group/abc", d.Id())
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsGroupRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: oldGroup,
			},
		},
		Resource: ResourceEntitlements(),
		HCL:      `group_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "group/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"group_id":             "abc",
		"allow_cluster_create": true,
	})
}

func TestResourceEntitlementsGroupReadEmpty(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: emptyGroup,
			},
		},
		Resource: ResourceEntitlements(),
		HCL:      `group_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "group/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"group_id":                   "abc",
		"allow_cluster_create":       false,
		"workspace_access":           false,
		"allow_instance_pool_create": false,
		"databricks_sql_access":      false,
	})
}

func TestResourceEntitlementsGroupRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Status:   400,
				Response: apierr.APIError{
					Message:   "Something",
					ErrorCode: "SCIM_Else",
				},
			},
		},
		Resource: ResourceEntitlements(),
		New:      true,
		Read:     true,
		ID:       "group/abc",
		HCL:      `group_id = "abc"`,
	}.ExpectError(t, "Something")
}

func TestResourceEntitlementsGroupUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: updateRequest,
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: newGroup,
			},
		},
		Resource: ResourceEntitlements(),
		Update:   true,
		ID:       "group/abc",
		InstanceState: map[string]string{
			"group_id":             "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		group_id    = "abc"
		allow_cluster_create = true
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "group/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsGroupDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: oldGroup,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: deleteRequest,
				Response: Group{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlements(),
		Delete:   true,
		ID:       "group/abc",
		InstanceState: map[string]string{
			"group_id":             "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		group_id    = "abc"
		allow_cluster_create = true
		`,
	}.Apply(t)
}

func TestResourceEntitlementsGroupDeleteEmptyEntitlement(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: emptyGroup,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: deleteRequest,
				Response: apierr.APIError{
					ErrorCode: "INVALID_PATH",
					Message:   "invalidPath No such attribute with the name : entitlements in the current resource",
				},
				Status: 400,
			},
		},
		Resource: ResourceEntitlements(),
		Delete:   true,
		ID:       "group/abc",
		InstanceState: map[string]string{
			"group_id":             "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		group_id    = "abc"
		allow_cluster_create = true
		`,
	}.Apply(t)
}

var oldUser = User{
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
}

var newUser = User{
	DisplayName: "Example user",
	Active:      true,
	UserName:    "me@example.com",
	ID:          "abc",
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

func TestResourceEntitlementsUserCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: addRequest,
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlements(),
		HCL: `
		user_id = "abc"
		allow_instance_pool_create = true
		allow_cluster_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "user/abc", d.Id())
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsUserRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Response: oldUser,
			},
		},
		Resource: ResourceEntitlements(),
		HCL:      `user_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "user/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"user_id":              "abc",
		"allow_cluster_create": true,
	})
}

func TestResourceEntitlementsUserRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Status:   400,
				Response: apierr.APIError{
					Message:   "Something",
					ErrorCode: "SCIM_Else",
				},
			},
		},
		Resource: ResourceEntitlements(),
		New:      true,
		Read:     true,
		ID:       "user/abc",
		HCL:      `user_id = "abc"`,
	}.ExpectError(t, "Something")
}

func TestResourceEntitlementsUserUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Status:   400,
				Response: apierr.APIError{
					Message:   "Something",
					ErrorCode: "SCIM_Else",
				},
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: updateRequest,
				Status:          400,
				Response: apierr.APIError{
					Message:   "Something",
					ErrorCode: "SCIM_Else",
				},
			},
		},
		Resource: ResourceEntitlements(),
		Update:   true,
		ID:       "user/abc",
		InstanceState: map[string]string{
			"user_id":              "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		user_id    = "abc"
		allow_cluster_create = true
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
	}.ExpectError(t, "Something")
}

func TestResourceEntitlementsUserUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: updateRequest,
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlements(),
		Update:   true,
		ID:       "user/abc",
		InstanceState: map[string]string{
			"user_id":              "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		user_id    = "abc"
		allow_cluster_create = true
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "user/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsUserDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=entitlements",
				Response: oldUser,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: deleteRequest,
				Response: User{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlements(),
		Delete:   true,
		ID:       "user/abc",
		InstanceState: map[string]string{
			"user_id":              "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		user_id    = "abc"
		allow_cluster_create = true
		`,
	}.ApplyNoError(t)
}

func TestResourceEntitlementsSPNCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: addRequest,
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlements(),
		HCL: `
		service_principal_id = "abc"
		allow_cluster_create = true
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "spn/abc", d.Id())
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsSPNRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Response: User{
					ID:            "abc",
					ApplicationID: "bcd",
					DisplayName:   "Example Service Principal",
					Active:        true,
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
					},
				},
			},
		},
		Resource: ResourceEntitlements(),
		HCL:      `service_principal_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "spn/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"service_principal_id": "abc",
		"allow_cluster_create": true,
	})
}

func TestResourceEntitlementsSPNRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Status:   404,
			},
		},
		Resource: ResourceEntitlements(),
		New:      true,
		Read:     true,
		Removed:  true,
		ID:       "spn/abc",
		HCL:      `service_principal_id = "abc"`,
	}.ApplyNoError(t)
}

func TestResourceEntitlementsSPNRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Status:   400,
				Response: apierr.APIError{
					Message:   "Something",
					ErrorCode: "SCIM_Else",
				},
			},
		},
		Resource: ResourceEntitlements(),
		New:      true,
		Read:     true,
		ID:       "spn/abc",
		HCL:      `service_principal_id = "abc"`,
	}.ExpectError(t, "Something")
}

func TestResourceEntitlementsSPNUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: updateRequest,
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlements(),
		Update:   true,
		ID:       "spn/abc",
		InstanceState: map[string]string{
			"service_principal_id": "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		service_principal_id       = "abc"
		allow_cluster_create       = true
		allow_instance_pool_create = true
		databricks_sql_access      = true
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "spn/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementsSPNDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=entitlements",
				Response: oldUser,
			},
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: deleteRequest,
				Response: User{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlements(),
		Delete:   true,
		ID:       "spn/abc",
		InstanceState: map[string]string{
			"service_principal_id": "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		service_principal_id = "abc"
		allow_cluster_create = true
		`,
	}.ApplyNoError(t)
}

func TestResourceEntitlementsGroupCreateEmpty(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: emptyAddRequest,
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=entitlements",
				Response: emptyGroup,
			},
		},
		Resource: ResourceEntitlements(),
		HCL: `
		group_id = "abc"
		allow_instance_pool_create = false
		allow_cluster_create = false
		databricks_sql_access = false
		workspace_access = false
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "group/abc", d.Id())
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, false, d.Get("allow_instance_pool_create"))
	assert.Equal(t, false, d.Get("databricks_sql_access"))
	assert.Equal(t, false, d.Get("workspace_access"))
}

func TestResourceEntitlementsCreate_AccountLevelShouldError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource:  ResourceEntitlements(),
		HCL:       `group_id = "abc"`,
		Create:    true,
		AccountID: "abc-123",
	}.Apply(t)
	assert.Contains(t, "entitlements can only be managed with a provider configured at the workspace-level", err.Error())
}
