package sharing

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestSharesData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares",
				Response: Shares{
					Shares: []ShareInfo{
						{
							Name: "a",
							Objects: []SharedDataObject{
								{
									Name:           "a",
									DataObjectType: "TABLE",
									Comment:        "c",
								},
							},
							CreatedAt: 0,
							CreatedBy: "",
						},
					},
				},
			},
		},
		Resource:    DataSourceShares(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestSharesData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceShares(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
