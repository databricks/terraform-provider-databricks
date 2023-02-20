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
							Name:              "a",
							StorageRoot:       "",
							DefaultDacID:      "sth",
							Owner:             "John.Doe@example.com",
							MetastoreID:       "abc",
							Region:            "",
							Cloud:             "",
							GlobalMetastoreId: "",
							CreatedAt:         0,
							CreatedBy:         "",
							UpdatedAt:         0,
							UpdatedBy:         "",
							DeltaSharingScope: "",
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
					Metastores: []MetastoreInfo{
						{
							Name:         "a",
							StorageRoot:  "abc",
							DefaultDacID: "sth",
							Owner:        "John.Doe@example.com",
							MetastoreID:  "abc",
						},
						{
							Name:         "b",
							StorageRoot:  "dcw",
							DefaultDacID: "sth",
							Owner:        "John.Doe@example.com",
							MetastoreID:  "ded",
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
