package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestWarehousesData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWarehousesAPI().EXPECT().
				ListAll(mock.Anything, sql.ListWarehousesRequest{}).
				Return([]sql.EndpointInfo{
					{
						Id:   "1",
						Name: "bar",
					},
					{
						Id:   "2",
						Name: "bar",
					},
				}, nil)
		},
		Resource:    DataSourceWarehouses(),
		HCL:         ``,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": []string{"1", "2"},
	})
}

func TestWarehousesDataContains(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWarehousesAPI().EXPECT().
				ListAll(mock.Anything, sql.ListWarehousesRequest{}).
				Return([]sql.EndpointInfo{
					{
						Id:   "111",
						Name: "bar",
					},
					{
						Id:   "2",
						Name: "br",
					},
				}, nil)
		},
		Resource:    DataSourceWarehouses(),
		HCL:         `warehouse_name_contains = "ba"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": []string{"111"},
	})
}

func TestWarehousesData_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockWarehousesAPI().EXPECT().
				ListAll(mock.Anything, sql.ListWarehousesRequest{}).
				Return(nil, qa.ErrImATeapot)
		},
		Resource:    DataSourceWarehouses(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
