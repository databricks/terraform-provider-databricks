package identity

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureAccSP(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "azure" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	ctx := context.Background()

	spAPI := NewServicePrincipalsAPI(ctx, client)

	sp, err := spAPI.Create(ScimUser{
		ApplicationID: "00000000-0000-0000-0000-000000000001",
		Entitlements: entitlements{
			{
				Value: "allow-cluster-create",
			},
		},
		DisplayName: "ABC SP",
		Active:      true,
	})
	require.NoError(t, err)
	defer func() {
		err := spAPI.Delete(sp.ID)
		require.NoError(t, err)
	}()

	err = spAPI.Update(sp.ID, ScimUser{
		ApplicationID: sp.ApplicationID,
		Entitlements: entitlements{
			{
				Value: "allow-instance-pool-create",
			},
		},
		DisplayName: "BCD",
	})
	require.NoError(t, err)
}

func TestResourceServicePrincipalRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					ID:          "abc",
					DisplayName: "Example Service Principal",
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
		Resource: ResourceServicePrincipal(),
		HCL:      `display_name = "Sylens"`,
		New:      true,
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
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
		HCL:      `display_name = "Scotchmo"`,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalRead_Invalid_Azure(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		Azure:    true,
		ID:       "abc",
	}.ExpectError(t, "application_id is required for service principals in Azure Databricks")
}

func TestResourceServicePrincipalRead_Invalid_AWS(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		ID:       "abc",
	}.ExpectError(t, "display_name is required for service principals in Databricks on AWS")
}

func TestResourceServicePrincipalRead_Error(t *testing.T) {
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
		Resource: ResourceServicePrincipal(),
		New:      true,
		Read:     true,
		ID:       "abc",
		HCL:      `display_name = "Nightly Runner"`,
	}.ExpectError(t, "Something")
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
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
					},
					Schemas: []URN{ServicePrincipalSchema},
				},
				Response: ScimUser{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					DisplayName: "Example Service Principal",
					Active:      true,
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
		Resource: ResourceServicePrincipal(),
		Create:   true,
		HCL: `
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
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
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalCreate_Error_AzureNoApplicationID(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		Create:   true,
		Azure:    true,
		HCL: `
		display_name = "abc"
		`,
	}.ExpectError(t, "application_id is required for service principals in Azure Databricks")
}

func TestResourceServicePrincipalUpdate(t *testing.T) {
	newServicePrincipal := ScimUser{
		Schemas:     []URN{ServicePrincipalSchema},
		DisplayName: "Changed Name",
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
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: ScimUser{
					DisplayName: "Example Service Principal",
					Active:      true,
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
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
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
					DisplayName: "Example Service Principal",
					Active:      true,
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
		HCL:      `display_name = "Squanchy"`,
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
