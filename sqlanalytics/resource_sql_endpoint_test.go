package sqlanalytics

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPreviewAccSQLEndpoints(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.NewClientFromEnvironment()
	endpoitsAPI := NewSQLEndpointsAPI(ctx, client)
	se := SQLEndpoint{
		Name:            qa.RandomName("tf-"),
		ClusterSize:     "Small",
		AutoStopMinutes: 10,
		MaxNumClusters:  1,
		Tags: &Tags{
			CustomTags: []Tag{
				{"Country", "Netherlands"},
				{"City", "Amsterdam"},
			},
		},
	}
	err := endpoitsAPI.Create(&se, 20*time.Minute)
	require.NoError(t, err)
	defer func() {
		err = endpoitsAPI.Delete(se.ID)
		assert.NoError(t, err)
	}()

	se.Name = "renamed-" + se.Name
	err = endpoitsAPI.Edit(se)
	require.NoError(t, err)
}

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
				Resource: "/api/2.0/sql/endpoints",
				ExpectedRequest: SQLEndpoint{
					Name:           "foo",
					ClusterSize:    "Small",
					MaxNumClusters: 1,
				},
				Response: SQLEndpoint{
					ID: "abc",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/endpoints/abc",
				ReuseRequest: true,
				Response: SQLEndpoint{
					Name:           "foo",
					ClusterSize:    "Small",
					ID:             "abc",
					State:          "RUNNING",
					MaxNumClusters: 1,
				},
			},
			dataSourceListHTTPFixture,
		},
		Resource: ResourceSQLEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointCreate_ErrorDisabled(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/endpoints",
				ExpectedRequest: SQLEndpoint{
					Name:           "foo",
					ClusterSize:    "Small",
					MaxNumClusters: 1,
				},
				Status: 404,
				Response: common.APIError{
					ErrorCode: "FEATURE_DISABLED",
					Message:   "SQL Analytics is not supported",
				},
			},
		},
		Resource: ResourceSQLEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.ExpectError(t, "SQL Analytics is not supported")
}

func TestResourceSQLEndpointRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/endpoints/abc",
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
		Resource: ResourceSQLEndpoint(),
		ID:       "abc",
		Read:     true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/endpoints/abc/edit",
				ExpectedRequest: SQLEndpoint{
					ID:             "abc",
					Name:           "foo",
					ClusterSize:    "Small",
					MaxNumClusters: 1,
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/sql/endpoints/abc",
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
		Resource: ResourceSQLEndpoint(),
		ID:       "abc",
		Update:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
}

func TestResourceSQLEndpointDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/sql/endpoints/abc",
			},
		},
		Resource: ResourceSQLEndpoint(),
		ID:       "abc",
		Delete:   true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
}

func TestResourceSQLEndpoint_CornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSQLEndpoint())
}

func TestSQLEnpointAPI(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/endpoints",
			Response: EndpointList{
				Endpoints: []SQLEndpoint{
					{
						ID:   "foo",
						Name: "bar",
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/endpoints/failstart/start",
			Status:   404,
			Response: common.NotFound("nope"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/endpoints/deleting/start",
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/endpoints/deleting",
			Response: SQLEndpoint{
				State: "STARTING",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/endpoints/deleting",
			Response: SQLEndpoint{
				State: "DELETED",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/sql/endpoints/cantwait",
			Status:   500,
			Response: common.APIError{
				Message: "does not compute",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/endpoints/stopping/stop",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
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
