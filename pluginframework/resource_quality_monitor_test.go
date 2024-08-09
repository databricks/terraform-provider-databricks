package pluginframework

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginFrameworkResource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQualityMonitorCreateTimeseriesPluginFramework(t *testing.T) {
	// Setup mock workspace or account client whichever is applicable
	// This will be use to mock the databricks-sdk-go calls
	MockWorkspaceClientFunc := func(w *mocks.MockWorkspaceClient) {
		e := w.GetMockQualityMonitorsAPI().EXPECT()
		e.Create(mock.Anything, catalog.CreateMonitor{
			TableName:        "test_table",
			OutputSchemaName: "output.schema",
			AssetsDir:        "sample.dir",
			TimeSeries: &catalog.MonitorTimeSeries{
				Granularities: []string{"1 day"},
				TimestampCol:  "timestamp",
			},
		}).Return(&catalog.MonitorInfo{
			AssetsDir:             "sample.dir",
			OutputSchemaName:      "output.schema",
			TableName:             "test_table",
			Status:                catalog.MonitorInfoStatusMonitorStatusPending,
			DriftMetricsTableName: "test_table_drift",
		}, nil)
		e.GetByTableName(mock.Anything, "test_table").Return(&catalog.MonitorInfo{
			TableName:             "test_table",
			Status:                catalog.MonitorInfoStatusMonitorStatusActive,
			AssetsDir:             "sample.dir",
			OutputSchemaName:      "output.schema",
			DriftMetricsTableName: "test_table_drift",
		}, nil)
	}
	_ = MockWorkspaceClientFunc

	// Start the server
	// tanmaytodo

	// Create resource using the above mock workspace client
	qualityMonitorResource := QualityMonitorResource{
		Client: &common.DatabricksClient{},
	}

	// Create request and response for the respective CRUD operation and pass the config
	createRequest := pluginFrameworkResource.CreateRequest{
		Config: GetPluginFrameworkConfig(`
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			time_series = {
				granularities = ["1 day"],
				timestamp_col = "timestamp"
			} 
		`),
	}
	createResponse := pluginFrameworkResource.CreateResponse{}
	ctx := context.Background()

	// Call the method to unit test
	qualityMonitorResource.Create(ctx, createRequest, &createResponse)

	// Add assertions
	assert.False(t, createResponse.Diagnostics.HasError())
}
