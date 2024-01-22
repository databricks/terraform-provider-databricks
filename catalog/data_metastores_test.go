package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoresDataContainsName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/testaccount/metastores",
				Response: catalog.ListMetastoresResponse{
					Metastores: []catalog.MetastoreInfo{
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
		AccountID:   "testaccount",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]interface{}{"a": "abc", "b": "ded"},
	})
}

func TestMetastoresData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "i'm a teapot")
}
