package monitoring

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestLakehouseMonitorCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceLakehouseMonitor())
}

func TestLakehouseMonitorCreateTimeseries(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor",
				ExpectedRequest: &catalog.CreateMonitor{
					FullName:         "test_table",
					OutputSchemaName: "output.schema",
					AssetsDir:        "sample.dir",
					TimeSeries: &catalog.MonitorTimeSeriesProfileType{
						Granularities: []string{"1 day"},
						TimestampCol:  "timestamp",
					},
				},
				Response: &catalog.MonitorInfo{
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor?",
				Response: &catalog.MonitorInfo{
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
				},
			},
		},
		Resource: ResourceLakehouseMonitor(),
		HCL: `
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			time_series = {
				granularities = ["1 day"],
				timestamp_col = "timestamp"
			} 
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestLakehouseMonitorCreateInference(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor",
				ExpectedRequest: &catalog.CreateMonitor{
					FullName:         "test_table",
					OutputSchemaName: "output.schema",
					AssetsDir:        "sample.dir",
					InferenceLog: &catalog.MonitorInferenceLogProfileType{
						Granularities: []string{"1 day"},
						TimestampCol:  "timestamp",
						PredictionCol: "prediction",
						ModelIdCol:    "model_id",
					},
				},
				Response: &catalog.MonitorInfo{
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor?",
				Response: &catalog.MonitorInfo{
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
				},
			},
		},
		Resource: ResourceLakehouseMonitor(),
		HCL: `
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			inference_log = {
				granularities = ["1 day"],
				timestamp_col = "timestamp",
				prediction_col = "prediction",
				model_id_col = "model_id"
			} 
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestLakehouseMonitorCreateSnapshot(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor",
				ExpectedRequest: &catalog.CreateMonitor{
					FullName:         "test_table",
					OutputSchemaName: "output.schema",
					AssetsDir:        "sample.dir",
					Snapshot:         struct{}{},
				},
				Response: &catalog.MonitorInfo{
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor?",
				Response: &catalog.MonitorInfo{
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
					AssetsDir:        "sample.dir",
					OutputSchemaName: "output.schema",
				},
			},
		},
		Resource: ResourceLakehouseMonitor(),
		HCL: `
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			snapshot = {}
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestLakehouseMonitorUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor",
				ExpectedRequest: &catalog.UpdateMonitor{
					AssetsDir:        "new_assets.dir",
					OutputSchemaName: "output.schema",
					InferenceLog: &catalog.MonitorInferenceLogProfileType{
						Granularities: []string{"1 week"},
						TimestampCol:  "timestamp",
						PredictionCol: "prediction",
						ModelIdCol:    "model_id",
					},
				},
				Response: catalog.MonitorInfo{
					AssetsDir:        "new_assets.dir",
					OutputSchemaName: "output.schema",
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
					InferenceLog: &catalog.MonitorInferenceLogProfileType{
						Granularities: []string{"1 week"},
						TimestampCol:  "timestamp",
						PredictionCol: "prediction",
						ModelIdCol:    "model_id",
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor?",
				Response: &catalog.MonitorInfo{
					TableName:        "test_table",
					Status:           catalog.MonitorInfoStatusMonitorStatusActive,
					AssetsDir:        "new_assets.dir",
					OutputSchemaName: "output.schema",
					InferenceLog: &catalog.MonitorInferenceLogProfileType{
						Granularities: []string{"1 week"},
						TimestampCol:  "timestamp",
						PredictionCol: "prediction",
						ModelIdCol:    "model_id",
					},
				},
			},
		},
		Resource: ResourceLakehouseMonitor(),
		Update:   true,
		ID:       "test_table",
		InstanceState: map[string]string{
			"table_name": "test_table",
		},
		HCL: `
			table_name = "test_table",
			assets_dir = "new_assets.dir",
			output_schema_name = "output.schema",
			inference_log = {
				granularities = ["1 week"],
				timestamp_col = "timestamp"
				prediction_col = "prediction",
				model_id_col = "model_id"
			} 
		`,
	}.ApplyNoError(t)
}

func TestLakehouseMonitorDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/tables/test_table/monitor?",
				Response: "",
			},
		},
		Resource: ResourceLakehouseMonitor(),
		Delete:   true,
		ID:       "test_table",
	}.ApplyNoError(t)
}
