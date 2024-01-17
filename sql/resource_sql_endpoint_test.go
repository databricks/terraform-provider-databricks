package sql

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
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

func TestResourceSQLEndpointCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses",
				ExpectedRequest: sql.CreateWarehouseRequest{
					Name:               "foo",
					ClusterSize:        "Small",
					MaxNumClusters:     1,
					AutoStopMins:       120,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
				Response: sql.CreateWarehouseResponse{
					Id: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc?",
				ReuseRequest: true,
				Response: sql.GetWarehouseResponse{
					Name:           "foo",
					ClusterSize:    "Small",
					Id:             "abc",
					State:          "RUNNING",
					Tags:           &sql.EndpointTags{},
					MaxNumClusters: 1,
					NumClusters:    1,
				},
			},
			dataSourceListHTTPFixture,
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

func TestResourceSQLEndpointCreate_NoServerless(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses",
				ExpectedRequest: sql.CreateWarehouseRequest{
					Name:               "foo",
					ClusterSize:        "Small",
					MaxNumClusters:     1,
					AutoStopMins:       120,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
					ForceSendFields:    []string{"EnableServerlessCompute"},
				},
				Response: sql.CreateWarehouseResponse{
					Id: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc?",
				ReuseRequest: true,
				Response: sql.GetWarehouseResponse{
					Name:           "foo",
					ClusterSize:    "Small",
					Id:             "abc",
					State:          "RUNNING",
					Tags:           &sql.EndpointTags{},
					MaxNumClusters: 1,
					NumClusters:    1,
				},
			},
			dataSourceListHTTPFixture,
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		enable_serverless_compute = false
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointCreateNoAutoTermination(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses",
				ExpectedRequest: sql.CreateWarehouseRequest{
					Name:               "foo",
					ClusterSize:        "Small",
					MaxNumClusters:     1,
					AutoStopMins:       0,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
				Response: sql.CreateWarehouseResponse{
					Id: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc?",
				ReuseRequest: true,
				Response: sql.GetWarehouseResponse{
					Name:           "foo",
					ClusterSize:    "Small",
					Id:             "abc",
					State:          "RUNNING",
					Tags:           &sql.EndpointTags{},
					MaxNumClusters: 1,
					NumClusters:    1,
				},
			},
			dataSourceListHTTPFixture,
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		auto_stop_mins = 0
		`,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointCreate_ErrorDisabled(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses",
				Status:   404,
				Response: apierr.APIError{
					ErrorCode: "FEATURE_DISABLED",
					Message:   "Databricks SQL is not supported",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc?",
				ReuseRequest: true,
				Response: sql.GetWarehouseResponse{
					Name:        "foo",
					ClusterSize: "Small",
					Id:          "abc",
					State:       "RUNNING",
				},
			},
			dataSourceListHTTPFixture,
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses/abc/edit",
				ExpectedRequest: sql.EditWarehouseRequest{
					Id:                 "abc",
					Name:               "foo",
					ClusterSize:        "Small",
					AutoStopMins:       120,
					MaxNumClusters:     1,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc?",
				ReuseRequest: true,
				Response: sql.GetWarehouseResponse{
					Name:        "foo",
					ClusterSize: "Small",
					Id:          "abc",
					State:       "RUNNING",
					NumClusters: 1,
				},
			},
			dataSourceListHTTPFixture,
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/sql/warehouses/abc?",
			},
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
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/data_sources",
			Response: map[string]any{},
			Status:   404,
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
	})
}

func TestResolveDataSourceIDNotFound(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/data_sources",
			Response: []any{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
	})
}
