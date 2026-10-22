package sql

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Define fixture for retrieving all data sources.
// Shared between tests that end up performing a read operation.
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
			api.EXPECT().Create(mock.Anything, createRequest).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
				Id:   "abc", // Set the ID in the waiter
				Poll: poll.Simple(getResponse),
			}, nil)
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

func TestResourceSQLEndpointCreate_ForceSendFields(t *testing.T) {
	type forceSendFieldTestCase struct {
		hcl                             string
		expectedEnableServerlessCompute bool
		expectedPhoton                  bool
		expectedForceSendFields         []string
	}
	cases := []forceSendFieldTestCase{
		{
			hcl:                             "enable_serverless_compute = true",
			expectedEnableServerlessCompute: true,
			expectedPhoton:                  true,
			expectedForceSendFields:         nil,
		},
		{
			hcl:                             "enable_serverless_compute = false",
			expectedEnableServerlessCompute: false,
			expectedPhoton:                  true,
			expectedForceSendFields:         []string{"EnableServerlessCompute"},
		},
		{
			hcl:                             "enable_photon = false",
			expectedEnableServerlessCompute: false,
			expectedPhoton:                  false,
			expectedForceSendFields:         []string{"EnablePhoton"},
		},
	}
	for _, c := range cases {
		t.Run(c.hcl, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					api := w.GetMockWarehousesAPI()
					response := sql.GetWarehouseResponse{
						Id:                      "abc",
						Name:                    "foo",
						ClusterSize:             "Small",
						AutoStopMins:            120,
						EnablePhoton:            c.expectedPhoton,
						EnableServerlessCompute: c.expectedEnableServerlessCompute,
						ForceSendFields:         c.expectedForceSendFields,
					}
					api.EXPECT().Create(mock.Anything, sql.CreateWarehouseRequest{
						Name:                    "foo",
						ClusterSize:             "Small",
						AutoStopMins:            120,
						EnablePhoton:            c.expectedPhoton,
						EnableServerlessCompute: c.expectedEnableServerlessCompute,
						MaxNumClusters:          1,
						SpotInstancePolicy:      "COST_OPTIMIZED",
						ForceSendFields:         c.expectedForceSendFields,
					}).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
						Id:   "abc",
						Poll: poll.Simple(response),
					}, nil)
					api.EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
					addDataSourceListHttpFixture(w)
				},
				Resource: ResourceSqlEndpoint(),
				Create:   true,
				HCL: `
				name = "foo"
				cluster_size = "Small"
				` + c.hcl,
			}.Apply(t)
			require.NoError(t, err)
			assert.Equal(t, "abc", d.Id(), "Id should not be empty")
			assert.Equal(t, "d7c9d05c-7496-4c69-b089-48823edad40c", d.Get("data_source_id"))
		})
	}

}

func TestResourceSQLEndpointCreateNoAutoTermination(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockWarehousesAPI().EXPECT()
			e.Create(mock.Anything, sql.CreateWarehouseRequest{
				Name:               "foo",
				ClusterSize:        "Small",
				MaxNumClusters:     1,
				AutoStopMins:       0,
				EnablePhoton:       true,
				SpotInstancePolicy: "COST_OPTIMIZED",
				ForceSendFields:    []string{"AutoStopMins"},
			}).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
				Id:   "abc",
				Poll: poll.Simple(getResponse),
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&getResponse, nil)
			addDataSourceListHttpFixture(w)
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

// this is pending https://github.com/databricks/databricks-sdk-go/pull/1155
func simpleError[R any](err error) poll.PollFunc[R] {
	return func(_ time.Duration, _ func(*R)) (*R, error) {
		return nil, err
	}
}

func TestResourceSQLEndpointCreateRollback(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockWarehousesAPI().EXPECT()
			e.Create(mock.Anything, sql.CreateWarehouseRequest{
				Name:               "foo",
				ClusterSize:        "Small",
				MaxNumClusters:     1,
				AutoStopMins:       0,
				EnablePhoton:       true,
				SpotInstancePolicy: "COST_OPTIMIZED",
				ForceSendFields:    []string{"AutoStopMins"},
			}).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
				Id:   "warehouse_id",
				Poll: simpleError[sql.GetWarehouseResponse](errors.New("Clusters are failing to launch")),
			}, nil)
			e.DeleteById(mock.Anything, "warehouse_id").Return(nil)
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		auto_stop_mins = 0
		`,
	}.ExpectError(t, "failed waiting for warehouse to start: Clusters are failing to launch")
}

func TestResourceSQLEndpointCreateRollbackFailed(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockWarehousesAPI().EXPECT()
			e.Create(mock.Anything, sql.CreateWarehouseRequest{
				Name:               "foo",
				ClusterSize:        "Small",
				MaxNumClusters:     1,
				AutoStopMins:       0,
				EnablePhoton:       true,
				SpotInstancePolicy: "COST_OPTIMIZED",
				ForceSendFields:    []string{"AutoStopMins"},
			}).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
				Id:   "warehouse_id",
				Poll: simpleError[sql.GetWarehouseResponse](errors.New("Clusters are failing to launch")),
			}, nil)
			e.DeleteById(mock.Anything, "warehouse_id").Return(errors.New("Cannot delete warehouse"))
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		auto_stop_mins = 0
		`,
	}.ExpectError(t, "failed waiting for warehouse to start: Clusters are failing to launch. when rolling back, also failed: Cannot delete warehouse")
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
			}).Return(&sql.WaitGetWarehouseRunning[struct{}]{
				Id:   "abc",
				Poll: poll.Simple(getResponse),
			}, nil)
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

// Testing the customizeDiff on clearing "health" diff is working as expected.
func TestResourceSQLEndpointUpdateHealthNoDiff(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceSqlEndpoint(),
		ID:       "abc",
		InstanceState: map[string]string{
			"name":                 "foo",
			"cluster_size":         "Small",
			"auto_stop_mins":       "120",
			"enable_photon":        "true",
			"max_num_clusters":     "1",
			"spot_instance_policy": "COST_OPTIMIZED",
		},
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"state":                     {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"provider_config.#":         {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"odbc_params.#":             {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_clusters":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_active_sessions":       {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"jdbc_url":                  {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"id":                        {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"enable_serverless_compute": {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"data_source_id":            {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"creator_name":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"customer_version_label":    {NewComputed: true},
			"statement_timeout":         {NewComputed: true},
			"max_dbu":                   {NewComputed: true},
		},
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.ApplyNoError(t)
}

func TestResourceSQLEndpointUpdateNoAutoTermination(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceSqlEndpoint(),
		ID:       "abc",
		InstanceState: map[string]string{
			"name":                 "foo",
			"cluster_size":         "Small",
			"auto_stop_mins":       "120",
			"enable_photon":        "true",
			"max_num_clusters":     "1",
			"spot_instance_policy": "COST_OPTIMIZED",
		},
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"auto_stop_mins":            {Old: "120", New: "0", NewComputed: false, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"state":                     {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"provider_config.#":         {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"odbc_params.#":             {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_clusters":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_active_sessions":       {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"jdbc_url":                  {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"id":                        {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"enable_serverless_compute": {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"data_source_id":            {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"creator_name":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"customer_version_label":    {NewComputed: true},
			"statement_timeout":         {NewComputed: true},
			"max_dbu":                   {NewComputed: true},
		},
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		auto_stop_mins = 0
		`,
	}.ApplyNoError(t)
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
		mwc.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return(nil, &apierr.APIError{
			ErrorCode: "RESOURCE_DOES_NOT_EXIST",
		})
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		dataSourceId, err := resolveDataSourceID(ctx, w, "any")
		require.NoError(t, err)
		assert.Equal(t, "", dataSourceId)
	})
}

func TestResolveDataSourceIDNotFound(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mwc *mocks.MockWorkspaceClient) {
		mwc.GetMockDataSourcesAPI().EXPECT().List(mock.Anything).Return([]sql.DataSource{}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		require.NoError(t, err)
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
	})
}

func TestResourceSQLEndpointCreateNoWait(t *testing.T) {
	// Create a specific response for the no_wait test that shows a starting state
	getResponseStarting := sql.GetWarehouseResponse{
		Name:           "foo",
		ClusterSize:    "Small",
		Id:             "abc",
		State:          "STARTING", // Important: Warehouse is still starting
		Tags:           &sql.EndpointTags{},
		MaxNumClusters: 1,
		NumClusters:    0, // No clusters running yet
	}

	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockWarehousesAPI()
			// The Create should return immediately without waiting due to no_wait=true
			api.EXPECT().Create(mock.Anything, createRequest).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
				Id:   "abc",                            // Set the ID directly in the waiter
				Poll: poll.Simple(getResponseStarting), // Use starting state response
			}, nil)

			// Read operation will still be called after create to refresh state
			api.EXPECT().GetById(mock.Anything, "abc").Return(&getResponseStarting, nil)
			addDataSourceListHttpFixture(w)
		},
		Resource: ResourceSqlEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
		cluster_size = "Small"
		no_wait = true
		`,
	}.ApplyAndExpectData(t, map[string]interface{}{
		"id":             "abc",
		"no_wait":        true,
		"state":          "STARTING",
		"data_source_id": "d7c9d05c-7496-4c69-b089-48823edad40c",
	})
}

func TestResourceSQLEndpointCreateRealtime(t *testing.T) {
	for _, tc := range []struct {
		config                    string
		timeout, maxDbu, autoStop int
	}{
		{"", 0, 0, 10},
		{"statement_timeout = 0\nmax_dbu = 1\nauto_stop_mins = 0", 0, 1, 0},
		{"statement_timeout = 1\nmax_dbu = 1000\nauto_stop_mins = 5", 1, 1000, 5},
	} {
		t.Run(fmt.Sprint(tc.timeout, tc.maxDbu, tc.autoStop), func(t *testing.T) {
			request := map[string]any{
				"name": "foo", "cluster_size": "Small", "warehouse_type": "REALTIME",
				"min_num_clusters": 1, "max_num_clusters": 1, "enable_photon": true,
				"enable_serverless_compute": true, "spot_instance_policy": "COST_OPTIMIZED",
				"auto_stop_mins": tc.autoStop,
			}
			if tc.config != "" {
				request["statement_timeout"], request["max_dbu"] = tc.timeout, tc.maxDbu
			}
			response := maps.Clone(request)
			response["id"], response["state"], response["warehouse_type"] = "abc", "STOPPED", "REYDEN"
			response["customer_version_label"] = "2026.10"
			// HTTP fixtures verify zero-value serialization and omitted fields.
			qa.ResourceFixture{
				Fixtures: []qa.HTTPFixture{
					{Method: "POST", Resource: "/api/2.0/sql/warehouses", ExpectedRequest: request, Response: `{"id":"abc"}`},
					{Method: "GET", Resource: "/api/2.0/sql/warehouses/abc?", Response: response},
				},
				Resource: ResourceSqlEndpoint(), Create: true,
				HCL: `name = "foo"
cluster_size = "Small"
warehouse_type = "REALTIME"
` + tc.config,
			}.ApplyAndExpectData(t, map[string]any{
				"id": "abc", "state": "STOPPED", "warehouse_type": "REALTIME",
				"auto_stop_mins": tc.autoStop, "enable_serverless_compute": true,
				"statement_timeout": tc.timeout, "customer_version_label": "2026.10", "max_dbu": tc.maxDbu,
			})
		})
	}
}

func TestResourceSQLEndpointReadRealtimeWithoutTimeout(t *testing.T) {
	response := `{"id":"abc","name":"foo","cluster_size":"Small","warehouse_type":"REALTIME"}`
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{Method: "GET", Resource: "/api/2.0/sql/warehouses/abc?", Response: response},
		},
		Resource: ResourceSqlEndpoint(), Read: true, ID: "abc",
	}.ApplyAndExpectData(t, map[string]any{"statement_timeout": 0})
}

func TestResourceSQLEndpointRealtimeValidation(t *testing.T) {
	for _, tc := range []struct {
		field string
		value any
		error string
	}{
		{"auto_stop_mins", 4, "auto_stop_mins must be 0 or at least 5 for REALTIME warehouses"},
		{"auto_stop_mins", -1, "auto_stop_mins must be 0 or at least 5 for REALTIME warehouses"},
		{"statement_timeout", -1, "invalid config supplied. [statement_timeout] expected statement_timeout to be at least (0), got -1"},
		{"min_num_clusters", 0, "min_num_clusters must be 1 for REALTIME warehouses"},
		{"max_num_clusters", 2, "max_num_clusters must be 1 for REALTIME warehouses"},
		{"enable_photon", false, "enable_photon must be true for REALTIME warehouses"},
		{"enable_serverless_compute", false, "enable_serverless_compute must be true for REALTIME warehouses"},
		{"spot_instance_policy", "RELIABILITY_OPTIMIZED", "spot_instance_policy cannot be configured for REALTIME warehouses"},
	} {
		t.Run(tc.field, func(t *testing.T) {
			config := map[string]any{"name": "foo", "cluster_size": "Small", "warehouse_type": "REALTIME"}
			config[tc.field] = tc.value
			qa.ResourceFixture{Resource: ResourceSqlEndpoint(), Create: true, State: config}.ExpectError(t, tc.error)
		})
	}
}

func TestResourceSQLEndpointUpdateRealtime(t *testing.T) {
	assert.NotContains(t, ResourceSqlEndpoint().Schema, "reyden_cluster_size_multiplier")
	for _, tc := range []struct {
		config          string
		maxDbu, timeout int
	}{{"max_dbu = 192", 192, 3600}, {"", 96, 3600}, {"statement_timeout = 0", 96, 0}, {"statement_timeout = 600", 96, 600}} {
		t.Run(fmt.Sprint(tc.maxDbu, tc.timeout), func(t *testing.T) {
			request := map[string]any{
				"name": "foo", "cluster_size": "Small", "warehouse_type": "REALTIME",
				"min_num_clusters": 1, "max_num_clusters": 1, "enable_photon": true,
				"enable_serverless_compute": true, "spot_instance_policy": "COST_OPTIMIZED",
				"auto_stop_mins": 10,
			}
			if tc.maxDbu == 192 {
				request["max_dbu"] = tc.maxDbu
			}
			if tc.timeout != 3600 {
				request["statement_timeout"] = tc.timeout
			}
			response := maps.Clone(request)
			response["max_dbu"], response["statement_timeout"] = tc.maxDbu, tc.timeout
			response["reyden_cluster_size_multiplier"], response["customer_version_label"] = 2, "2026.10"
			qa.ResourceFixture{
				Fixtures: []qa.HTTPFixture{
					{Method: "POST", Resource: "/api/2.0/sql/warehouses/abc/edit", ExpectedRequest: request},
					{Method: "GET", Resource: "/api/2.0/sql/warehouses/abc?", Response: response},
				},
				Resource: ResourceSqlEndpoint(), Update: true, ID: "abc",
				HCL: `name = "foo"
cluster_size = "Small"
warehouse_type = "REALTIME"
` + tc.config,
				InstanceState: map[string]string{"name": "old", "warehouse_type": "REALTIME", "max_dbu": "96", "statement_timeout": "3600", "customer_version_label": "2026.10"},
			}.ApplyAndExpectData(t, map[string]any{"max_dbu": tc.maxDbu, "statement_timeout": tc.timeout, "customer_version_label": "2026.10"})
		})
	}
}

func TestResourceSQLEndpointRealtimeOmittedFieldsNoDiff(t *testing.T) {
	for _, tc := range []struct{ timeout, maxDbu string }{{"0", "0"}, {"3600", "96"}} {
		t.Run(tc.timeout, func(t *testing.T) {
			qa.ResourceFixture{
				Resource: ResourceSqlEndpoint(),
				HCL: `name = "foo"
cluster_size = "Small"
warehouse_type = "REALTIME"`,
				InstanceState: map[string]string{
					"name": "foo", "cluster_size": "Small", "warehouse_type": "REALTIME",
					"auto_stop_mins": "10", "min_num_clusters": "1", "max_num_clusters": "1",
					"enable_photon": "true", "enable_serverless_compute": "true",
					"spot_instance_policy": "COST_OPTIMIZED", "no_wait": "false",
					"statement_timeout": tc.timeout, "max_dbu": tc.maxDbu,
					"id": "abc", "state": "STOPPED", "creator_name": "owner",
					"customer_version_label": "2026.10", "data_source_id": "", "jdbc_url": "",
					"num_active_sessions": "0", "num_clusters": "0",
					"odbc_params.#": "0", "provider_config.#": "0",
				},
				ExpectedDiff: map[string]*terraform.ResourceAttrDiff{},
			}.ApplyNoError(t)
		})
	}
}

func TestResourceSQLEndpointRealtimeFieldValidation(t *testing.T) {
	for _, value := range []any{-3, 0, 1.5} {
		config := terraform.NewResourceConfigRaw(map[string]any{
			"name": "foo", "cluster_size": "Small", "warehouse_type": "REALTIME", "max_dbu": value,
		})
		assert.True(t, ResourceSqlEndpoint().ToResource().Validate(config).HasError(), "max_dbu=%v", value)
	}
	for _, warehouseType := range []string{"PRO", "CLASSIC"} {
		for _, field := range []string{"max_dbu", "statement_timeout"} {
			qa.ResourceFixture{
				Resource: ResourceSqlEndpoint(), Create: true,
				State: map[string]any{"name": "foo", "cluster_size": "Small", "warehouse_type": warehouseType, field: 96},
			}.ExpectError(t, field+" is only supported for REALTIME warehouses")
		}
	}
}

func TestResourceSQLEndpointRejectRealtimeTypeMigration(t *testing.T) {
	for _, tc := range []struct{ from, to string }{
		{"PRO", "REALTIME"}, {"CLASSIC", "REALTIME"}, {"REALTIME", "PRO"}, {"REALTIME", "CLASSIC"},
	} {
		t.Run(tc.from+"-to-"+tc.to, func(t *testing.T) {
			r := ResourceSqlEndpoint().ToResource()
			raw, err := r.CoreConfigSchema().CoerceValue(cty.ObjectVal(map[string]cty.Value{
				"name": cty.StringVal("foo"), "cluster_size": cty.StringVal("Small"), "warehouse_type": cty.StringVal(tc.to),
			}))
			require.NoError(t, err)
			// ResourceFixture does not supply the existing ID during planning.
			state := &terraform.InstanceState{ID: "abc", Attributes: map[string]string{"warehouse_type": tc.from}, RawConfig: raw}
			qa.MockWorkspaceApply(t, func(*mocks.MockWorkspaceClient) {}, func(ctx context.Context, client *common.DatabricksClient) {
				_, err := r.Diff(ctx, state, terraform.NewResourceConfigShimmed(raw, r.CoreConfigSchema()), client)
				assert.EqualError(t, err, "changing warehouse_type to or from REALTIME is not supported")
			})
		})
	}
}
