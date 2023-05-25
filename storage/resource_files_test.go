package storage

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestVolumesCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateFiles(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/files",
				ExpectedRequest: CreateFileRequestContent{
					Name:    "testName",
					Comment: "This is a test comment.",
					Owner:   "testOwner",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/files",
				Response: FileInfo{
					Name:        "testName",
					Comment:     "This is a test comment.",
					Owner: "testOwner".
				},
			},
		},
		Resource: ResourceFiles(),
		Create:   true,
		HCL: `
		name = "testName"
		owner = "testOwner"
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testName", d.Get("name"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
	assert.Equal(t, "testOwner", d.Get("owner"))

}
