package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestQualityMonitorCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceQualityMonitor())
}

func TestQualityMonitorCreateTimeseries(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
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
		},
		Resource: ResourceQualityMonitor(),
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

func TestQualityMonitorCreateInference(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateMonitor{
				TableName:        "test_table",
				OutputSchemaName: "output.schema",
				AssetsDir:        "sample.dir",
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 day"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
					ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
				},
			}).Return(&catalog.MonitorInfo{
				AssetsDir:        "sample.dir",
				OutputSchemaName: "output.schema",
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
			}, nil)
			e.GetByTableName(mock.Anything, "test_table").Return(&catalog.MonitorInfo{
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				AssetsDir:        "sample.dir",
				OutputSchemaName: "output.schema",
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 day"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
					ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
				},
			}, nil)
		},
		Resource: ResourceQualityMonitor(),
		HCL: `
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			inference_log = {
				granularities = ["1 day"],
				timestamp_col = "timestamp",
				prediction_col = "prediction",
				model_id_col = "model_id",
				problem_type = "PROBLEM_TYPE_REGRESSION"
			} 
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestQualityMonitorCreateSnapshot(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateMonitor{
				TableName:        "test_table",
				OutputSchemaName: "output.schema",
				AssetsDir:        "sample.dir",
				Snapshot:         &catalog.MonitorSnapshot{},
			}).Return(&catalog.MonitorInfo{
				AssetsDir:        "sample.dir",
				OutputSchemaName: "output.schema",
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
			}, nil)
			e.GetByTableName(mock.Anything, "test_table").Return(&catalog.MonitorInfo{
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				AssetsDir:        "sample.dir",
				OutputSchemaName: "output.schema",
				Snapshot:         &catalog.MonitorSnapshot{},
			}, nil)
		},
		Resource: ResourceQualityMonitor(),
		HCL: `
			table_name = "test_table",
			assets_dir = "sample.dir",
			output_schema_name = "output.schema",
			snapshot = {}
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestQualityMonitorGet(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.GetByTableName(mock.Anything, "test_table").Return(&catalog.MonitorInfo{
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				AssetsDir:        "new_assets.dir",
				OutputSchemaName: "output.schema",
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 week"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
				},
				DriftMetricsTableName: "test_table_drift"}, nil)
		},
		Resource: ResourceQualityMonitor(),
		Read:     true,
		ID:       "test_table",
	}.ApplyNoError(t)
}

func TestQualityMonitorUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateMonitor{
				TableName:        "test_table",
				OutputSchemaName: "output.schema",
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 week"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
					ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
				},
			}).Return(&catalog.MonitorInfo{
				AssetsDir:        "new_assets.dir",
				OutputSchemaName: "output.schema",
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 week"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
				},
				DriftMetricsTableName: "test_table_drift",
			}, nil)
			e.GetByTableName(mock.Anything, "test_table").Return(&catalog.MonitorInfo{
				TableName:        "test_table",
				Status:           catalog.MonitorInfoStatusMonitorStatusActive,
				AssetsDir:        "new_assets.dir",
				OutputSchemaName: "output.schema",
				InferenceLog: &catalog.MonitorInferenceLog{
					Granularities: []string{"1 week"},
					TimestampCol:  "timestamp",
					PredictionCol: "prediction",
					ModelIdCol:    "model_id",
				},
				DriftMetricsTableName: "test_table_drift",
			}, nil)
		},
		Resource: ResourceQualityMonitor(),
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
				model_id_col = "model_id",
				problem_type = "PROBLEM_TYPE_REGRESSION"
			} 
		`,
	}.ApplyNoError(t)
}

func TestQualityMonitorDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.DeleteByTableName(mock.Anything, "test_table").Return(nil)
		},
		Resource: ResourceQualityMonitor(),
		Delete:   true,
		ID:       "test_table",
	}.ApplyNoError(t)
}
