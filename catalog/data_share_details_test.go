package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestShareDetailsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/shares/a?include_shared_data=true",
				Response: ShareInfo{
					Name: "a",
					Objects: []SharedDataObject{
						{
							Name:           "a",
							DataObjectType: "TABLE",
							Comment:        "c",
							SharedAs:       "",
							AddedAt:        0,
							AddedBy:        "",
						},
					},
				},
			},
		},
		Resource:    DataSourceShareDetails(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "a"
		`,
	}.ApplyNoError(t)
}

func TestShareDetailsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceShareDetails(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
