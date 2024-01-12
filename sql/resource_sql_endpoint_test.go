package sql

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
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
				"endpoint_id": "def"
			},
			{
				"id": "d7c9d05c-7496-4c69-b089-48823edad40c",
				"endpoint_id": "abc"
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
				ExpectedRequest: SQLEndpoint{
					Name:               "foo",
					ClusterSize:        "Small",
					MaxNumClusters:     1,
					AutoStopMinutes:    120,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
				Response: SQLEndpoint{
					ID: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:           "foo",
					ClusterSize:    "Small",
					ID:             "abc",
					State:          "RUNNING",
					Tags:           &Tags{},
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

func TestResourceSQLEndpointCreateNoAutoTermination(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/warehouses",
				ExpectedRequest: SQLEndpoint{
					Name:               "foo",
					ClusterSize:        "Small",
					MaxNumClusters:     1,
					AutoStopMinutes:    0,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
				Response: SQLEndpoint{
					ID: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:           "foo",
					ClusterSize:    "Small",
					ID:             "abc",
					State:          "RUNNING",
					Tags:           &Tags{},
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
	}.ExpectError(t, "Databricks SQL is not supported")
}

func TestResourceSQLEndpointRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:        "foo",
					ClusterSize: "Small",
					ID:          "abc",
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
				ExpectedRequest: SQLEndpoint{
					ID:                 "abc",
					Name:               "foo",
					ClusterSize:        "Small",
					AutoStopMinutes:    120,
					MaxNumClusters:     1,
					EnablePhoton:       true,
					SpotInstancePolicy: "COST_OPTIMIZED",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/warehouses/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:        "foo",
					ClusterSize: "Small",
					ID:          "abc",
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

func TestSQLEnpointAPI(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/warehouses",
			Response: map[string]any{
				"warehouses": []SQLEndpoint{
					{
						ID:   "foo",
						Name: "bar",
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/warehouses/failstart/start",
			Status:   404,
			Response: apierr.NotFound("nope"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/warehouses/deleting/start",
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/warehouses/deleting",
			Response: SQLEndpoint{
				State: "STARTING",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/warehouses/deleting",
			Response: SQLEndpoint{
				State: "DELETED",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/warehouses/cantwait",
			Status:   500,
			Response: apierr.APIError{
				Message: "does not compute",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/warehouses/stopping/stop",
		},
	}, func(ctx context.Context, client common.DatabricksAPI) {
		a := NewSQLEndpointsAPI(ctx, client)
		list, err := a.List()
		require.NoError(t, err)
		assert.Len(t, list.Endpoints, 1)

		err = a.Start("failstart", 5*time.Minute)
		assert.EqualError(t, err, "nope")

		err = a.Start("deleting", 5*time.Minute)
		assert.EqualError(t, err, "endpoint got deleted during creation")

		err = a.waitForRunning("cantwait", 5*time.Minute)
		assert.EqualError(t, err, "does not compute")

		err = a.Stop("stopping")
		require.NoError(t, err)
	})
}

func TestResolveDataSourceIDError(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/data_sources",
			Response: map[string]any{},
			Status:   404,
		},
	}, func(ctx context.Context, client common.DatabricksAPI) {
		_, err := NewSQLEndpointsAPI(ctx, client).ResolveDataSourceID("any")
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
	}, func(ctx context.Context, client common.DatabricksAPI) {
		_, err := NewSQLEndpointsAPI(ctx, client).ResolveDataSourceID("any")
		require.Error(t, err)
	})
}
