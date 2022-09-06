package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
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

func TestResourceEntitlementGroupCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: oldGroup,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
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
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: newGroup,
			},
		},
		Resource: ResourceEntitlement(),
		HCL: `
		group_id = "abc"
		allow_instance_pool_create = true
		allow_cluster_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "group/abc", d.Id())
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementGroupRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: oldGroup,
			},
		},
		Resource: ResourceEntitlement(),
		HCL:      `group_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "group/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"group_id":             "abc",
		"allow_cluster_create": true,
	})
}
func TestResourceEntitlementGroupUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: oldGroup,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
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
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: newGroup,
			},
		},
		Resource: ResourceEntitlement(),
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
	require.NoError(t, err, err)
	assert.Equal(t, "group/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, true, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}
func TestResourceEntitlementGroupDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: oldGroup,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: Group{
					Schemas:      []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName:  "Data Scientists",
					Entitlements: []ComplexValue{},
				},
				Response: Group{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlement(),
		Delete:   true,
		ID:       "group/abc",
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

func TestResourceEntitlementUserCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
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
				Response: newUser,
			},
		},
		Resource: ResourceEntitlement(),
		HCL: `
		user_id = "abc"
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "user/abc", d.Id())
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementUserRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: oldUser,
			},
		},
		Resource: ResourceEntitlement(),
		HCL:      `user_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "user/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"user_id":              "abc",
		"allow_cluster_create": true,
	})
}

func TestResourceEntitlementUserUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
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
				Response: newUser,
			},
		},
		Resource: ResourceEntitlement(),
		Update:   true,
		ID:       "user/abc",
		InstanceState: map[string]string{
			"user_id":              "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		user_id    = "abc"
		allow_cluster_create = false
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "user/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementUserDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{},
					UserName:     "me@example.com",
					Schemas:      []URN{UserSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlement(),
		Delete:   true,
		ID:       "user/abc",
	}.ApplyNoError(t)
}

func TestResourceEntitlementSPNCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
						},
					},
					UserName: "me@example.com",
					Schemas:  []URN{ServicePrincipalSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlement(),
		HCL: `
		spn_id = "abc"
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "spn/abc", d.Id())
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementSPNRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
		Resource: ResourceEntitlement(),
		HCL:      `spn_id = "abc"`,
		New:      true,
		Read:     true,
		ID:       "spn/abc",
	}.ApplyAndExpectData(t, map[string]any{
		"spn_id":               "abc",
		"allow_cluster_create": true,
	})
}

func TestResourceEntitlementSPNRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   404,
			},
		},
		Resource: ResourceEntitlement(),
		New:      true,
		Read:     true,
		Removed:  true,
		ID:       "spn/abc",
		HCL:      `spn_id = "abc"`,
	}.ApplyNoError(t)
}

func TestResourceEntitlementSPNRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   400,
				Response: common.APIErrorBody{
					ScimDetail: "Something",
					ScimStatus: "Else",
				},
			},
		},
		Resource: ResourceEntitlement(),
		New:      true,
		Read:     true,
		ID:       "spn/abc",
		HCL:      `spn_id = "abc"`,
	}.ExpectError(t, "Something")
}

func TestResourceEntitlementSPNUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{
						{
							Value: "allow-instance-pool-create",
						},
						{
							Value: "databricks-sql-access",
						},
					},
					UserName: "me@example.com",
					Schemas:  []URN{ServicePrincipalSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: newUser,
			},
		},
		Resource: ResourceEntitlement(),
		Update:   true,
		ID:       "spn/abc",
		InstanceState: map[string]string{
			"spn_id":               "abc",
			"allow_cluster_create": "true",
		},
		HCL: `
		spn_id    = "abc"
		allow_cluster_create = false
		allow_instance_pool_create = true
		databricks_sql_access = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "spn/abc", d.Id(), "Id should not be empty")
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("databricks_sql_access"))
}

func TestResourceEntitlementSPNDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: oldUser,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: User{
					Entitlements: entitlements{},
					UserName:     "me@example.com",
					Schemas:      []URN{ServicePrincipalSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
		},
		Resource: ResourceEntitlement(),
		Delete:   true,
		ID:       "spn/abc",
	}.ApplyNoError(t)
}
