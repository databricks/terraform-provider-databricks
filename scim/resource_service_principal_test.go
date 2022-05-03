package scim

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccServicePrincipalOnAzure(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "azure" {
		t.Skip("Test will only run with CLOUD_ENV=azure")
	}

	client := common.NewClientFromEnvironment()
	ctx := context.Background()

	spAPI := NewServicePrincipalsAPI(ctx, client)

	sp, err := spAPI.Create(User{
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

	err = spAPI.Update(sp.ID, User{
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
				Response: User{
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
				ExpectedRequest: User{
					DisplayName: "Example Service Principal",
					Active:      true,
					Entitlements: []ComplexValue{
						{
							Value: "allow-cluster-create",
						},
					},
					Schemas: []URN{ServicePrincipalSchema},
				},
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: User{
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

func TestResourceServicePrincipalUpdate(t *testing.T) {
	newServicePrincipal := User{
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
				Response: User{
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
				Response: User{
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

func TestCreateForceOverridesManuallyAddedServicePrincipalErrorNotMatched(t *testing.T) {
	d := ResourceUser().TestResourceData()
	d.Set("force", true)
	rerr := createForceOverridesManuallyAddedServicePrincipal(
		fmt.Errorf("nonsense"), d,
		NewServicePrincipalsAPI(context.Background(), &common.DatabricksClient{}), User{})
	assert.EqualError(t, rerr, "nonsense")
}

func TestCreateForceOverwriteCannotListServicePrincipals(t *testing.T) {
	appID := "12344ca0-e1d7-45d1-951e-f4b93592f123"
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%%20eq%%20%%27%s%%27", appID),
			Status:   417,
			Response: common.APIError{
				Message: "cannot find service principal",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedServicePrincipal(
			fmt.Errorf("Service principal with application ID %s already exists.", appID),
			d, NewServicePrincipalsAPI(ctx, client), User{
				ApplicationID: appID,
			})
		assert.EqualError(t, err, "cannot find service principal")
	})
}

func TestCreateForceOverwriteCannotListAccServicePrincipals(t *testing.T) {
	appID := "12344ca0-e1d7-45d1-951e-f4b93592f123"
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%%20eq%%20%%27%s%%27", appID),
			Response: UserList{
				TotalResults: 0,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		err := createForceOverridesManuallyAddedServicePrincipal(
			fmt.Errorf("Service principal with application ID %s already exists.", appID),
			d, NewServicePrincipalsAPI(ctx, client), User{
				ApplicationID: appID,
			})
		assert.EqualError(t, err, fmt.Sprintf("cannot find SP with ID %s for force import", appID))
	})
}

func TestCreateForceOverwriteFindsAndSetsServicePrincipalID(t *testing.T) {
	appID := "12344ca0-e1d7-45d1-951e-f4b93592f123"
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%%20eq%%20%%27%s%%27", appID),
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
			Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
			Response: User{
				ID: "abc",
			},
		},
		{
			Method:   "PUT",
			Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
			ExpectedRequest: User{
				Schemas:       []URN{ServicePrincipalSchema},
				ApplicationID: appID,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().TestResourceData()
		d.Set("force", true)
		d.Set("application_id", appID)
		err := createForceOverridesManuallyAddedServicePrincipal(
			fmt.Errorf("Service principal with application ID %s already exists.", appID),
			d, NewServicePrincipalsAPI(ctx, client), User{
				ApplicationID: appID,
			})
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id())
	})
}
