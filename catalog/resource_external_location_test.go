package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestExternalLocationCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/external-locations",
				ExpectedRequest: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Owner:          "administrators",
					Comment:        "def",
				},
				Response: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Owner:          "wrong_owner",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Owner:       "efg",
					MetastoreID: "fgh",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/external-locations/abc",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		owner = "administrators"
		comment = "def"
		`,
	}.ApplyNoError(t)
}
