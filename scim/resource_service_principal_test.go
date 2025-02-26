package scim

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceServicePrincipalRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
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
		Resource: ResourceServicePrincipal(),
		HCL:      `display_name = "Sylens"`,
		New:      true,
		Read:     true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":         "Example Service Principal",
		"application_id":       "bcd",
		"allow_cluster_create": true,
		"home":                 "/Users/bcd",
		"repos":                "/Repos/bcd",
		"acl_principal_id":     "servicePrincipals/bcd",
	})
}

func TestResourceServicePrincipalRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
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
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
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
	qa.ResourceFixture{
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
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
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
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalCreate_ErrorNoDisplayNameOrAppId(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		Create:   true,
		HCL: `
		allow_cluster_create = true
		`,
	}.ExpectError(t, "invalid config supplied. [application_id] Missing required argument. [display_name] Missing required argument")
}

func TestResourceServicePrincipalCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: qa.HTTPFailures,
		Resource: ResourceServicePrincipal(),
		Create:   true,
		HCL: `
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.ExpectError(t, "i'm a teapot")
}

func TestResourceServicePrincipalUpdateOnAWS(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=groups,roles",
				Response: User{
					// application ID is created by platform on AWS
					ApplicationID: "existing-application-id",

					DisplayName: "Existing Service Principal Display Name",
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
				ExpectedRequest: User{
					// application ID is not allowed to be modified by client side on AWS

					Schemas:     []URN{ServicePrincipalSchema},
					DisplayName: "New Service Principal Display Name",
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
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					Schemas:       []URN{ServicePrincipalSchema},
					ApplicationID: "existing-application-id",
					DisplayName:   "New Service Principal Display Name",
					Active:        true,
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
				},
			},
		},
		Resource: ResourceServicePrincipal(),
		InstanceState: map[string]string{},
		Update: true,
		ID:     "abc",
		HCL: `
		display_name = "New Service Principal Display Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":               "New Service Principal Display Name",
		"allow_cluster_create":       false,
		"allow_instance_pool_create": true,
	})
}

// https://github.com/databricks/terraform-provider-databricks/issues/1319
func TestResourceServicePrincipalUpdateOnAzure(t *testing.T) {
	qa.ResourceFixture{
		Azure: true,
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=groups,roles",
				Response: User{
					// application id is specified by user on Azure
					ApplicationID: "existing-application-id",

					DisplayName: "Example Service Principal",
					Active:      true,
					ID:          "abc",
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: User{
					// application id is specified by user on Azure and also must be part of modification
					ApplicationID: "existing-application-id",

					Schemas:     []URN{ServicePrincipalSchema},
					DisplayName: "Example Service Principal",
					Entitlements: entitlements{
						{
							Value: "allow-cluster-create",
						},
					},
					Active: true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					Schemas:       []URN{ServicePrincipalSchema},
					ApplicationID: "existing-application-id",
					DisplayName:   "Example Service Principal",
					Entitlements: entitlements{
						{
							Value: "allow-cluster-create",
						},
					},
					Active: true,
				},
			},
		},
		Resource: ResourceServicePrincipal(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"application_id": "existing-application-id",
			"display_name":   "Example Service Principal",
		},
		HCL: `
		application_id = "existing-application-id"
		display_name = "Example Service Principal"
		allow_cluster_create = true
		`,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: qa.HTTPFailures,
		Resource: ResourceServicePrincipal(),
		InstanceState: map[string]string{
			"display_name": "Changed Name",
		},
		Update: true,
		ID:     "abc",
		HCL: `
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.ExpectError(t, "i'm a teapot")
}

func TestResourceServicePrincipalUpdate_ErrorPut(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					DisplayName: "Example Service Principal",
					Active:      true,
					ID:          "abc",
					Entitlements: entitlements{
						{
							Value: "allow-cluster-create",
						},
					},
				},
			},
			qa.HTTPFailures[0],
		},
		Resource: ResourceServicePrincipal(),
		InstanceState: map[string]string{
			"display_name": "Changed Name",
		},
		Update: true,
		ID:     "abc",
		HCL: `
		display_name = "Changed Name"
		allow_cluster_create = false
		allow_instance_pool_create = true
		`,
	}.ExpectError(t, "i'm a teapot")
}

func TestResourceServicePrincipalDelete(t *testing.T) {
	qa.ResourceFixture{
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
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: qa.HTTPFailures,
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
	}.ExpectError(t, "i'm a teapot")
}

func TestResourceServicePrincipalDelete_NoErrorEmtpyParams(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceServicePrinicpalforce_delete_reposError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id = "abc"
			force_delete_repos = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalDelete_NonExistingRepo(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
				Status: 404,
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id = "abc"
			force_delete_repos = true	
		`,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalDelete_DirError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id = "abc"
			force_delete_home_dir = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalDelete_NonExistingDir(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
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
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
		 	application_id = "abc"
			force_delete_home_dir = true	
		`,
	}.ApplyNoError(t)
}

func TestCreateForceOverridesManuallyAddedServicePrincipalErrorNotMatched(t *testing.T) {
	d := ResourceServicePrincipal().ToResource().TestResourceData()
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
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%%20eq%%20%%22%s%%22", appID),
			Status:   417,
			Response: apierr.APIError{
				Message: "cannot find service principal",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceServicePrincipal().ToResource().TestResourceData()
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
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%%20eq%%20%%22%s%%22", appID),
			Response: UserList{
				TotalResults: 0,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceUser().ToResource().TestResourceData()
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
			Resource: fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%%20eq%%20%%22%s%%22", appID),
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
			Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=groups,roles",
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
		d := ResourceServicePrincipal().ToResource().TestResourceData()
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
