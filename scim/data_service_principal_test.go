package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
)

func TestDataServicePrincipalReadByAppId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%20eq%20%27abc%27",
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
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"sp_id":          "abc",
		"application_id": "abc",
		"display_name":   "Example Service Principal",
		"active":         true,
		"home":           "/Users/abc",
		"repos":          "/Repos/abc",
	})
}

func TestDataServicePrincipalReadNotFound(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%20eq%20%27abc%27",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err, err)
}

func TestDataServicePrincipalReadError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=applicationId%20eq%20%27abc%27",
				Status:   500,
			},
		},
		Resource:    DataSourceServicePrincipal(),
		HCL:         `application_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err, err)
}
