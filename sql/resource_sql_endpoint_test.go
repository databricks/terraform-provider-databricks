package sql

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
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
			}).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
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
			}).Return(&sql.WaitGetWarehouseRunning[struct{}]{Poll: poll.Simple(getResponse)}, nil)
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
			"odbc_params.#":             {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_clusters":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"num_active_sessions":       {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"jdbc_url":                  {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"id":                        {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"enable_serverless_compute": {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"data_source_id":            {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
			"creator_name":              {Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false},
		},
		HCL: `
		name = "foo"
  		cluster_size = "Small"
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
		_, err = resolveDataSourceID(ctx, w, "any")
		require.Error(t, err)
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
