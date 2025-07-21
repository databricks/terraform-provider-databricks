package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"gotest.tools/assert"
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
