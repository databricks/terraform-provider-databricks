package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoreDataVerify(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/testaccount/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						Name:        "xyz",
						MetastoreId: "abc",
						Owner:       "pqr",
					},
				},
			},
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		metastore_id = "abc"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"metastore_info.0.name":         "xyz",
		"metastore_info.0.owner":        "pqr",
		"metastore_info.0.metastore_id": "abc",
	})
}

func TestMetastoreDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "I'm a teapot")
}
