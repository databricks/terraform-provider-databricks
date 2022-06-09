package scim

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
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

func TestDataServicePrincipalsReadByDisplayName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20eq%20%27def%27",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc1",
							DisplayName:   "def",
							Active:        true,
							ApplicationID: "123",
						},
						{
							ID:            "abc2",
							DisplayName:   "def",
							Active:        true,
							ApplicationID: "124",
						},
					},
				},
			},
		},
		Resource:    DataSourceServicePrincipals(),
		HCL:         `display_name = "def"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"application_ids": []string{"123", "124"},
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

func TestDataServicePrincipalsReadNotFound(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20eq%20%27def%27",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipals(),
		HCL:         `display_name = "def"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err, err)
}
