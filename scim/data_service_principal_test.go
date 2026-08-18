package scim

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataServicePrincipalReadByAppId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%20eq%20%22abc%22",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc",
							DisplayName:   "Example Service Principal",
							Active:        true,
							ApplicationID: "abc",
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
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"scim_id":          "abc",
		"id":               "abc",
		"application_id":   "abc",
		"display_name":     "Example Service Principal",
		"active":           true,
		"home":             "/Users/abc",
		"repos":            "/Repos/abc",
		"acl_principal_id": "servicePrincipals/abc",
	})
}

func TestDataServicePrincipalReadBySpId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=id%20eq%20%22abc%22",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc",
							DisplayName:   "Example Service Principal",
							Active:        true,
							ApplicationID: "abc",
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
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `scim_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"scim_id":          "abc",
		"id":               "abc",
		"application_id":   "abc",
		"display_name":     "Example Service Principal",
		"active":           true,
		"home":             "/Users/abc",
		"repos":            "/Repos/abc",
		"acl_principal_id": "servicePrincipals/abc",
	})
}

func TestDataServicePrincipalReadByDisplayName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=displayName%20eq%20%22testsp%22",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc",
							DisplayName:   "testsp",
							Active:        true,
							ApplicationID: "abc",
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
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `display_name = "testsp"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"scim_id":          "abc",
		"id":               "abc",
		"application_id":   "abc",
		"display_name":     "testsp",
		"active":           true,
		"home":             "/Users/abc",
		"repos":            "/Repos/abc",
		"acl_principal_id": "servicePrincipals/abc",
	})
}

func TestDataServicePrincipalReadByAppIdNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%20eq%20%22abc%22",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "cannot find SP with an application ID abc")
}

func TestDataServicePrincipalReadByIdNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=id%20eq%20%22abc%22",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `scim_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "cannot find SP with an ID abc")
}

func TestDataServicePrincipalReadByNameNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=displayName%20eq%20%22abc%22",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `display_name = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "cannot find SP with name abc")
}

func TestDataServicePrincipalReadError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=applicationId%20eq%20%22abc%22",
				Status:   500,
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	assert.Error(t, err)
}

func TestDataServicePrincipalReadByNameDuplicates(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=displayName%20eq%20%22abc%22",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc1",
							DisplayName:   "abc",
							Active:        true,
							ApplicationID: "abc1",
						},
						{
							ID:            "abc2",
							DisplayName:   "abc",
							Active:        true,
							ApplicationID: "abc2",
						},
					},
				},
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `display_name = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.ExpectError(t, "there are 2 Service Principals with name abc")
}

func TestDataServicePrincipalReadNoParams(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceServicePrincipal(),
		HCL:         ``,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "please specify either application_id, display_name, or scim_id")
}

// TestDataServicePrincipalRead_AccountLevelNoHookFailure is a regression test
// for https://github.com/databricks/terraform-provider-databricks/issues/5664.
// When `databricks_service_principal` data source is used with an account-level
// provider, the post-Read provider_config hook used to call CurrentWorkspaceID
// and fail with:
//
//	Error: cannot populate provider_config for service principal:
//	failed to resolve workspace_id: failed to get the workspace_id:
//	strconv.ParseInt: parsing "": invalid syntax
//
// The fix in main adds `api` to the data source schema and marks it
// `IsDual: true` so the post-Read hook is short-circuited at account level.
//
// Fixtures cover BOTH the workspace SCIM path (used by v1.114.0 where the
// `api` field does not exist and HostType falls back to WorkspaceHost) and
// the account SCIM path (used on main once `api = "account"` is set). The
// test asserts the hook does not fire by deliberately NOT stubbing
// /api/2.0/preview/scim/v2/Me — if the hook runs it will hit the test server
// with no matching fixture and the assertion fails with a missing-stub
// error referencing /Me.
func TestDataServicePrincipalRead_AccountLevelNoHookFailure(t *testing.T) {
	spResources := UserList{
		Resources: []User{
			{
				ID:            "spid",
				DisplayName:   "testsp",
				Active:        true,
				ApplicationID: "appid",
			},
		},
	}
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/accounts/abc/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=displayName%20eq%20%22testsp%22",
			ReuseRequest: true,
			Response:     spResources,
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/ServicePrincipals?excludedAttributes=roles&filter=displayName%20eq%20%22testsp%22",
			ReuseRequest: true,
			Response:     spResources,
		},
		// Deliberately NO fixture for /api/2.0/preview/scim/v2/Me — if the
		// post-Read provider_config hook fires, it will try to call /Me and
		// the test will fail with a missing-stub error.
	})
	assert.NoError(t, err)
	defer server.Close()

	// Simulate a real account-level provider: AccountID set and no cached
	// workspace ID so the post-Read hook (if it fires) is forced to call /Me.
	client.Config.AccountID = "abc"
	client.SetCachedWorkspaceID(0)

	r := DataSourceServicePrincipal().ToResource()
	d := r.TestResourceData()
	// On main the `api` field exists and routes to account-level SCIM; on
	// v1.114.0 the schema has no such field and Set returns an error which
	// we deliberately ignore (the data source then falls back to workspace
	// SCIM, but the bug still surfaces because the hook runs unconditionally).
	_ = d.Set("api", "account")
	assert.NoError(t, d.Set("display_name", "testsp"))
	ctx := context.WithValue(context.Background(), common.ResourceName, "service_principal")
	diags := r.ReadContext(ctx, d, client)

	assert.False(t, diags.HasError(),
		"post-Read provider_config hook must be skipped for service_principal data source at account level; got: %v", diags)
	assert.Equal(t, "spid", d.Id())
}

func TestDataServicePrincipalReadInvalidConfig(t *testing.T) {
	qa.ResourceFixture{
		Resource: DataSourceServicePrincipal(),
		HCL: `display_name = "abc"
		application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "invalid config supplied. [application_id] Invalid combination of arguments. [display_name] Invalid combination of arguments. [scim_id] Invalid combination of arguments")
}
