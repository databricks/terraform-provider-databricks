package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestSqlEndpointStatementTimeout(t *testing.T) {
	for _, warehouseType := range []string{"PRO", "CLASSIC"} {
		t.Run(warehouseType, func(t *testing.T) {
			request := createRequest
			request.WarehouseType = sql.CreateWarehouseRequestWarehouseType(warehouseType)
			request.StatementTimeout = 600
			response := getResponse
			response.WarehouseType = sql.GetWarehouseResponseWarehouseType(warehouseType)
			response.StatementTimeout = 600
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(workspace *mocks.MockWorkspaceClient) {
					api := workspace.GetMockWarehousesAPI()
					api.EXPECT().Create(mock.Anything, request).Return(&sql.WaitGetWarehouseRunning[sql.CreateWarehouseResponse]{
						Id: "abc", Poll: poll.Simple(response),
					}, nil)
					api.EXPECT().GetById(mock.Anything, "abc").Return(&response, nil)
					addDataSourceListHttpFixture(workspace)
				},
				Resource: ResourceSqlEndpoint(),
				Create:   true,
				State: map[string]any{
					"name": "foo", "cluster_size": "Small",
					"warehouse_type": warehouseType, "statement_timeout": 600,
				},
			}.ApplyAndExpectData(t, map[string]any{"statement_timeout": 600})
		})
	}
}
