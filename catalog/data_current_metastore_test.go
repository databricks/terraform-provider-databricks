package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestCurrentMetastoreDataVerify(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().
				Summary(mock.Anything).
				Return(&catalog.GetMetastoreSummaryResponse{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Cloud:       "aws",
				}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().
				Summary(mock.Anything).
				Return(nil, &apierr.APIError{
					ErrorCode: "BAD_REQUEST",
					Message:   "Bad request: unable to get metastore summary",
				})
		},
		Resource:    DataSourceCurrentMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "Bad request: unable to get metastore summary")
}
