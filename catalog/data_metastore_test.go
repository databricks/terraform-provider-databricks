package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoreDataApply(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/testaccount/metastores/abc?",
				Response: catalog.MetastoreInfo{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
				},
			},
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "abc",
		AccountID:   "testaccount",
	}.ApplyNoError(t)
}

func TestMetastoreDataVerify(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/testaccount/metastores/abc?",
				Response: catalog.MetastoreInfo{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
				},
			},
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "abc",
		AccountID:   "testaccount",
	}.ApplyAndExpectData(t, map[string]any{})
}

func TestMetastoreData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "I'm a teapot")
}
