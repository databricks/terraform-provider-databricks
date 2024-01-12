package sql

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Define fixture for retrieving all data sources.
// Shared between tests that end up performing a read operation.
var dataSourceListHTTPFixture = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/preview/sql/data_sources",
	ReuseRequest: true,
	Response: json.RawMessage(`
		[
			{
				"id": "2f47f0f9-b4b7-40e2-b130-43103151864c",
				"warehouse_id": "def"
			},
			{
				"id": "d7c9d05c-7496-4c69-b089-48823edad40c",
				"warehouse_id": "abc"
			}
		]
	`),
}

func addDataSourceListHttpFixture(mw *mocks.MockWorkspaceClient) {
	mw.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{
		{
			Id:          "2f47f0f9-b4b7-40e2-b130-43103151864c",
			WarehouseId: "def",
		},
		{
			Id:          "d7c9d05c-7496-4c69-b089-48823edad40c",
			WarehouseId: "abc",
		},
	}, nil)
}

func makeResponseWaiter[R any](resp R, warehouse *sql.GetWarehouseResponse) *sql.WaitGetWarehouseRunning[R] {
	return &sql.WaitGetWarehouseRunning[R]{
		Response: &resp,
		Poll: func(_ time.Duration, _ func(*sql.GetWarehouseResponse)) (*sql.GetWarehouseResponse, error) {
			return warehouse, nil
		},
	}
}

var createRequest = sql.CreateWarehouseRequest{
	Name:               "foo",
	ClusterSize:        "Small",
	MaxNumClusters:     1,
	AutoStopMins:       120,
	EnablePhoton:       true,
	SpotInstancePolicy: "COST_OPTIMIZED",
}
var getResponse = sql.GetWarehouseResponse{
	Name:           "foo",
	ClusterSize:    "Small",
	Id:             "abc",
	State:          "RUNNING",
	Tags:           &sql.EndpointTags{},
	MaxNumClusters: 1,
	NumClusters:    1,
}

func TestResourceSQLEndpointCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockWarehousesAPI()
			api.EXPECT().Create(mock.Anything, createRequest).Return(makeResponseWaiter(sql.CreateWarehouseResponse{
				Id: "abc",
			}, &getResponse), nil)
			api.EXPECT().GetById(mock.Anything, "abc").Return(&getResponse, nil)
			addDataSourceListHttpFixture(w)
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointCreate_ErrorDisabled(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			api := mwc.GetMockWarehousesAPI()
			api.EXPECT().
				Create(mock.Anything, createRequest).
				Return(nil, errors.New("Databricks SQL is not supported"))
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.ExpectError(t, "failed creating warehouse: Databricks SQL is not supported")
}

func TestResourceSQLEndpointRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			api := mwc.GetMockWarehousesAPI()
			api.EXPECT().GetById(mock.Anything, "abc").Return(&getResponse, nil)
			addDataSourceListHttpFixture(mwc)
		},
		Resource: ResourceSqlEndpoint(),
		ID:       "abc",
		Read:     true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			api := mwc.GetMockWarehousesAPI()
			api.EXPECT().Edit(mock.Anything, sql.EditWarehouseRequest{
				Id:                 "abc",
				Name:               "foo",
				ClusterSize:        "Small",
				AutoStopMins:       120,
				MaxNumClusters:     1,
				EnablePhoton:       true,
				SpotInstancePolicy: "COST_OPTIMIZED",
			}).Return(makeResponseWaiter(struct{}{}, &getResponse), nil)
			api.EXPECT().GetById(mock.Anything, "abc").Return(&getResponse, nil)
			addDataSourceListHttpFixture(mwc)
		},
		Resource: ResourceSqlEndpoint(),
		ID:       "abc",
		Update:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			api := mwc.GetMockWarehousesAPI()
			api.EXPECT().DeleteById(mock.Anything, "abc").Return(nil)
		},
		Resource: ResourceSqlEndpoint(),
		ID:       "abc",
		Delete:   true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
}

func TestResourceSQLEndpoint_CornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlEndpoint())
}

func TestResolveDataSourceIDError(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return(nil, apierr.ErrNotFound)
	}, func(ctx context.Context, client common.DatabricksAPI) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
	})
}

func TestResolveDataSourceIDNotFound(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{}, nil)
	}, func(ctx context.Context, client common.DatabricksAPI) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
	})
}
