package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWarehouseData(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&sql.GetWarehouseResponse{
				Name:        "foo",
				ClusterSize: "Small",
				Id:          "abc",
				State:       "RUNNING",
			}, nil)
			addDataSourceListHttpFixture(w)
		},
		Resource:    DataSourceWarehouse(),
		HCL:         `id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "foo", d.Get("name"))
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestWarehouseData_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(nil, qa.ErrImATeapot)
			addDataSourceListHttpFixture(w)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `id = "abc"`,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestWarehouseDataById_DataSourceListError(t *testing.T) {
	// When DataSources.List fails with ID lookup, we still get the warehouse but without data_source_id
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return(nil, qa.ErrImATeapot)
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&sql.GetWarehouseResponse{
				Name:        "foo",
				ClusterSize: "Small",
				Id:          "abc",
				State:       "RUNNING",
			}, nil)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `id = "abc"`,
		ID:          "abc",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "foo", d.Get("name"))
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "", d.Get("data_source_id"))
}

func TestWarehouseDataByName_DataSourceListError(t *testing.T) {
	// When DataSources.List fails, we fall back to Warehouses.ListAll and don't populate data_source_id
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return(nil, qa.ErrImATeapot)
			w.GetMockWarehousesAPI().EXPECT().ListAll(mock.Anything, sql.ListWarehousesRequest{}).Return([]sql.EndpointInfo{
				{
					Id:   "abc",
					Name: "abc",
				},
			}, nil)
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&sql.GetWarehouseResponse{
				Name:        "abc",
				ClusterSize: "Small",
				Id:          "abc",
				State:       "RUNNING",
			}, nil)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "", d.Get("data_source_id"))
}

func TestWarehouseDataByName_NotFoundError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad401",
					WarehouseId: "def",
					Name:        "test",
				},
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad40c",
					WarehouseId: "abc",
					Name:        "abc2",
				},
			}, nil)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.ExpectError(t, "can't find SQL warehouse with the name 'abc'")
}

func TestWarehouseDataByName_DuplicatesError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad401",
					WarehouseId: "def",
					Name:        "abc",
				},
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad40c",
					WarehouseId: "abc",
					Name:        "abc",
				},
			}, nil)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "abc"`,
		ID:          "_",
	}.ExpectError(t, "there are multiple SQL warehouses with the name 'abc'")
}

func TestWarehouseDataByName(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad401",
					WarehouseId: "def",
					Name:        "abc",
				},
				{
					Id:          "d7c9d05c-7496-4c69-b089-48823edad40c",
					WarehouseId: "abc",
					Name:        "test",
				},
			}, nil)
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&sql.GetWarehouseResponse{
				Name:        "test",
				ClusterSize: "Small",
				Id:          "abc",
				State:       "RUNNING",
			}, nil)
		},
		Resource:    DataSourceWarehouse(),
		Read:        true,
		NonWritable: true,
		HCL:         `name = "test"`,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}
