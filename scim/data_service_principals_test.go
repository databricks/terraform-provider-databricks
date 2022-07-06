package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
)

func TestDataServicePrincipalsReadByDisplayName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20co%20%27def%27",
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
		HCL:         `display_name_contains = "def"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"application_ids": []string{"123", "124"},
	})
}

func TestDataServicePrincipalsReadNotFound(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20co%20%27def%27",
				Response: UserList{},
			},
		},
		Resource:    DataSourceServicePrincipals(),
		HCL:         `display_name_contains = "def"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err, err)
}

func TestDataServicePrincipalsReadNoFilter(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20co%20%27%27",
				Response: UserList{
					Resources: []User{
						{
							ID:            "abc1",
							DisplayName:   "def1",
							Active:        true,
							ApplicationID: "124",
						},
						{
							ID:            "abc2",
							DisplayName:   "def2",
							Active:        true,
							ApplicationID: "123",
						},
					},
				},
			},
		},
		Resource:    DataSourceServicePrincipals(),
		HCL:         ``,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"application_ids": []string{"123", "124"},
	})
}

func TestDataServicePrincipalsReadError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals?filter=displayName%20co%20%27def%27",
				Status:   500,
			},
		},
		Resource:    DataSourceServicePrincipals(),
		HCL:         `display_name_contains = "def"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err, err)
}
