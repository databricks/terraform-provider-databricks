package catalog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestExternalLocationCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Owner:       "efg",
					MetastoreID: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: map[string]any{
					"owner": "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Owner:       "administrators",
					MetastoreID: "fgh",
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

func TestCreateExternalLocationReadOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					ReadOnly:       true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Owner:       "efg",
					MetastoreID: "fgh",
					ReadOnly:    true,
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		read_only = true
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: map[string]any{
					"credential_name": "bcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocation_skipValidationSuppressDiff(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				Response: ExternalLocationInfo{
					Name:           "abc",
					URL:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"skip_validation": "false",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "abc"
		comment = "def"
		skip_validation = true
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.False(t, d.HasChanges("skip_validation"))
}
