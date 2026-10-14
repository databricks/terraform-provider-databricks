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
			{name: "omitted"},
		} {
			t.Run(operation+"/"+tc.name, func(t *testing.T) {
				id := ""
				if update {
					id = "abc"
				}
				var instanceState map[string]string
				if update && tc.name == "zero" {
					instanceState = map[string]string{"statement_timeout": "3600"}
				}
				response := getResponse
				response.StatementTimeout = tc.timeout
				if tc.name == "omitted" {
					response.StatementTimeout = 7200
				}
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

func TestSqlEndpointStatementTimeoutOmittedNoDiff(t *testing.T) {
	response := getResponse
	response.AutoStopMins = 120
	response.EnablePhoton = true
	response.SpotInstancePolicy = "COST_OPTIMIZED"
	response.StatementTimeout = 3600
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWarehousesAPI().EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
			addDataSourceListHttpFixture(w)
		},
		Resource: ResourceSqlEndpoint(), ID: "abc", Read: true,
	}.Apply(t)
	require.NoError(t, err)
	state := d.State().Attributes
	state["name"] = "foo"
	state["cluster_size"] = "Small"
	state["odbc_params.#"] = "0"
	qa.ResourceFixture{
		Resource: ResourceSqlEndpoint(), ID: "abc",
		InstanceState: state,
		ExpectedDiff:  map[string]*terraform.ResourceAttrDiff{},
		HCL: `name = "foo"
cluster_size = "Small"`,
	}.ApplyNoError(t)
}
