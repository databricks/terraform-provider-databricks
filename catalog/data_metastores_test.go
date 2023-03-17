package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoresData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores",
				Response: Metastores{
					Metastores: []unitycatalog.MetastoreInfo{
						{
							Name:                      "a",
							StorageRoot:               "",
							DefaultDataAccessConfigId: "sth",
							Owner:                     "John.Doe@example.com",
							MetastoreId:               "abc",
							Region:                    "",
							Cloud:                     "",
							GlobalMetastoreId:         "",
							CreatedAt:                 0,
							CreatedBy:                 "",
							UpdatedAt:                 0,
							UpdatedBy:                 "",
							DeltaSharingScope:         "",
							DeltaSharingRecipientTokenLifetimeInSeconds: 0,
							DeltaSharingOrganizationName:                "",
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

func TestMetastoresDataContainsName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores",
				Response: Metastores{
					Metastores: []unitycatalog.MetastoreInfo{
						{
							Name:                      "a",
							StorageRoot:               "abc",
							DefaultDataAccessConfigId: "sth",
							Owner:                     "John.Doe@example.com",
							MetastoreId:               "abc",
						},
						{
							Name:                      "b",
							StorageRoot:               "dcw",
							DefaultDataAccessConfigId: "sth",
							Owner:                     "John.Doe@example.com",
							MetastoreId:               "ded",
						},
					},
				},
			},
		},
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `metastores {
			name = "a"
			storage_root = "abc"
		}`,
	}.ApplyAndExpectData(t, map[string]any{
		"metastores.0.name": "a",
	})
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
