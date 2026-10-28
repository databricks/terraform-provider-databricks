package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestSqlEndpointStatementTimeoutSchema(t *testing.T) {
	resource := ResourceSqlEndpoint()
	timeout := resource.Schema["statement_timeout"]
	require.True(t, timeout.Optional)
	require.False(t, timeout.Computed)
	require.Equal(t, 0, timeout.Default)

	dataSource := DataSourceWarehouse()
	timeout = dataSource.Schema["statement_timeout"]
	require.True(t, timeout.Computed)
	require.Nil(t, timeout.Default)
}

func TestSqlEndpointStatementTimeout(t *testing.T) {
	for _, update := range []bool{false, true} {
		operation := "create"
		if update {
			operation = "update"
		}
		for _, tc := range []struct {
			name    string
			hcl     string
			timeout int
			force   []string
		}{
			{name: "positive", hcl: "statement_timeout = 3600", timeout: 3600},
			{name: "zero", hcl: "statement_timeout = 0", force: []string{"StatementTimeout"}},
			{name: "omitted", force: []string{"StatementTimeout"}},
		} {
			t.Run(operation+"/"+tc.name, func(t *testing.T) {
				id := ""
				if update {
					id = "abc"
				}
				var instanceState map[string]string
				if update && (tc.name == "zero" || tc.name == "omitted") {
					instanceState = map[string]string{"statement_timeout": "3600"}
				}
				response := getResponse
				response.StatementTimeout = tc.timeout
				qa.ResourceFixture{
					MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
						api := w.GetMockWarehousesAPI()
						if update {
							api.EXPECT().Edit(mock.Anything, sql.EditWarehouseRequest{
								Id: "abc", Name: "foo", ClusterSize: "Small",
								AutoStopMins: 120, MaxNumClusters: 1, EnablePhoton: true,
								SpotInstancePolicy: "COST_OPTIMIZED",
								StatementTimeout:   tc.timeout, ForceSendFields: tc.force,
							}).Return(&sql.WaitGetWarehouseRunning[struct{}]{Id: "abc"}, nil)
						} else {
							request := createRequest
							request.StatementTimeout = tc.timeout
							request.ForceSendFields = tc.force
							api.EXPECT().Create(mock.Anything, request).Return(
								&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
									Id: "abc", Poll: poll.Simple(response),
								}, nil)
						}
						api.EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
						addDataSourceListHttpFixture(w)
					},
					Resource: ResourceSqlEndpoint(), ID: id,
					InstanceState: instanceState,
					Create:        !update, Update: update,
					HCL: `name = "foo"
cluster_size = "Small"
` + tc.hcl,
				}.ApplyAndExpectData(t, map[string]interface{}{
					"statement_timeout": response.StatementTimeout,
				})
			})
		}
	}
}

func TestSqlEndpointStatementTimeoutRead(t *testing.T) {
	for _, dataSource := range []bool{false, true} {
		name := "import"
		resource := ResourceSqlEndpoint()
		hcl := ""
		if dataSource {
			name = "data_source"
			resource = DataSourceWarehouse()
			hcl = `id = "abc"`
		}
		t.Run(name, func(t *testing.T) {
			response := getResponse
			response.StatementTimeout = 3600
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
					addDataSourceListHttpFixture(w)
				},
				Resource: resource, ID: "abc", HCL: hcl,
				Read: true, NonWritable: dataSource,
			}.ApplyAndExpectData(t, map[string]interface{}{"statement_timeout": 3600})
		})
	}
}

func TestSqlEndpointStatementTimeoutPlan(t *testing.T) {
	for _, tc := range []struct {
		name    string
		timeout int
		hcl     string
		reset   bool
	}{
		{name: "omitted_after_reset"},
		{name: "explicit_zero", hcl: "statement_timeout = 0"},
		{name: "configured", timeout: 3600, hcl: "statement_timeout = 3600"},
		{name: "removed", timeout: 3600, reset: true},
		{name: "clear", timeout: 3600, hcl: "statement_timeout = 0", reset: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response := getResponse
			response.AutoStopMins = 120
			response.EnablePhoton = true
			response.SpotInstancePolicy = "COST_OPTIMIZED"
			response.StatementTimeout = tc.timeout
			data, err := qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
					addDataSourceListHttpFixture(w)
				},
				Resource: ResourceSqlEndpoint(), ID: "abc", Read: true,
			}.Apply(t)
			require.NoError(t, err)
			state := data.State().Attributes
			state["name"] = "foo"
			state["cluster_size"] = "Small"
			state["odbc_params.#"] = "0"
			expectedDiff := map[string]*terraform.ResourceAttrDiff{}
			if tc.reset {
				expectedDiff["statement_timeout"] = &terraform.ResourceAttrDiff{Old: "3600", New: "0"}
			}
			qa.ResourceFixture{
				Resource: ResourceSqlEndpoint(), ID: "abc",
				InstanceState: state,
				ExpectedDiff:  expectedDiff,
				HCL: `name = "foo"
cluster_size = "Small"
` + tc.hcl,
			}.ApplyNoError(t)
		})
	}
}
