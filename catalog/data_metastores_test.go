package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoresData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores",
				Response: Metastores{
					Metastores: []MetastoreInfo{
						{
							Name: "a",
							StorageRoot: "",
							DefaultDacID: "",
							Owner: "",
							MetastoreID: "",
							Region: "",
							Cloud: "",
							GlobalMetastoreId: "",
							CreatedAt: 0,
							CreatedBy: "",
							UpdatedAt: 0,
							UpdatedBy: "",
							DeltaSharingScope: "",
							DeltaSharingRecipientTokenLifetimeInSeconds: 0,
							DeltaSharingOrganizationName: "",
						},
					},
				},
			},
		},
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestMetastoresData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
