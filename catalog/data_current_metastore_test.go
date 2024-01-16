package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCurrentMetastoreDataVerify(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Cloud:       "aws",
				},
			},
		},
		Resource:    DataSourceCurrentMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"metastore_info.0.name":         "xyz",
		"metastore_info.0.owner":        "pqr",
		"metastore_info.0.metastore_id": "abc",
		"metastore_info.0.cloud":        "aws",
	})
}

func TestCurrentMetastoreDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCurrentMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
