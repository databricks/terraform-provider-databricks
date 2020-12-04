package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceServicePrincipalRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					ID:            "abc",
					DisplayName:   "Example Service Principal",
					ApplicationID: "00000000-0000-0000-0000-000000000000",
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
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", d.Get("application_id"))
	assert.Equal(t, "Example Service Principal", d.Get("display_name"))
	assert.Equal(t, false, d.Get("allow_cluster_create"))
}

func TestResourceServicePrincipalRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   404,
			},
		},
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
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
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	require.Error(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceServicePrincipalCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals",
				ExpectedRequest: ScimUser{
					DisplayName: "Example Service Principal",
					Active:      true,
					Entitlements: []entitlementsListItem{
						{
							Value: "allow-cluster-create",
						},
					},
					ApplicationID: "00000000-0000-0000-0000-000000000000",
					Schemas:       []URN{ServicePrincipalSchema},
				},
				Response: ScimUser{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					DisplayName:   "Example Service Principal",
					Active:        true,
					ApplicationID: "00000000-0000-0000-0000-000000000000",
					ID:            "abc",
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
		Resource: ResourceServicePrincipal(),
		Create:   true,
		HCL: `
		application_id    = "00000000-0000-0000-0000-000000000000"
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", d.Get("application_id"))
	assert.Equal(t, "Example Service Principal", d.Get("display_name"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
}

func TestResourceServicePrincipalCreate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals",
				Status:   400,
			},
		},
		Resource: ResourceServicePrincipal(),
		Create:   true,
		HCL: `
		application_id    = "00000000-0000-0000-0000-000000000000"
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalUpdate(t *testing.T) {
	newServicePrincipal := ScimUser{
		Schemas:       []URN{ServicePrincipalSchema},
		DisplayName:   "Changed Name",
		ApplicationID: "00000000-0000-0000-0000-000000000000",
		Active:        true,
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
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					DisplayName:   "Example Service Principal",
					Active:        true,
					ApplicationID: "00000000-0000-0000-0000-000000000000",
					ID:            "abc",
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
			{
				Method:          "PUT",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: newServicePrincipal,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: newServicePrincipal,
			},
		},
		Resource: ResourceServicePrincipal(),
		Update:   true,
		ID:       "abc",
		HCL: `
		application_id    = "00000000-0000-0000-0000-000000000000"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", d.Get("application_id"))
	assert.Equal(t, "Changed Name", d.Get("display_name"))
	assert.Equal(t, false, d.Get("allow_cluster_create"))
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
}

func TestResourceServicePrincipalUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   400,
			},
		},
		Resource: ResourceServicePrincipal(),
		Update:   true,
		ID:       "abc",
		HCL: `
		application_id    = "00000000-0000-0000-0000-000000000000"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalUpdate_ErrorPut(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					DisplayName:   "Example Service Principal",
					Active:        true,
					ApplicationID: "00000000-0000-0000-0000-000000000000",
					ID:            "abc",
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
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   400,
			},
		},
		Resource: ResourceServicePrincipal(),
		Update:   true,
		ID:       "abc",
		HCL: `
		application_id    = "00000000-0000-0000-0000-000000000000"
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
}

func TestResourceServicePrincipalDelete_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Status:   400,
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	require.Error(t, err, err)
}
