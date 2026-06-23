package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestConnectionDataVerify(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockConnectionsAPI().EXPECT()
			e.GetByName(mock.Anything, "test_conn").Return(
				&catalog.ConnectionInfo{
					Name:           "test_conn",
					ConnectionType: "POSTGRESQL",
					Owner:          "admin",
					Comment:        "test connection",
					FullName:       "test_conn",
					MetastoreId:    "abc",
				},
				nil)
		},
		Resource:    DataSourceConnection(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "test_conn"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                                "test_conn",
		"connection_info.0.owner":           "admin",
		"connection_info.0.connection_type": "POSTGRESQL",
		"connection_info.0.comment":         "test connection",
	})
}

func TestConnectionDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceConnection(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
