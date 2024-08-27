package acceptance

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateQualityMonitorUnit(t *testing.T) {
	err := ResourceFixturePluginFramework{
		Steps: []Step{
			{
				// CREATE
				Template: `
				resource "databricks_lakehouse_monitor_pluginframework" "testMonitorInference" {
					table_name = "test_table"
					assets_dir = "new_assets.dir"
					output_schema_name = "output.schema"
					inference_log = {
						granularities = ["1 week"]
						timestamp_col = "timestamp"
						prediction_col = "prediction"
						model_id_col = "model_id"
						problem_type = "PROBLEM_TYPE_REGRESSION"
					}
				}`,
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					e := w.GetMockQualityMonitorsAPI().EXPECT()
					e.Create(mock.Anything, catalog.CreateMonitor{
						AssetsDir:        "new_assets.dir",
						TableName:        "test_table",
						OutputSchemaName: "output.schema",
						InferenceLog: &catalog.MonitorInferenceLog{
							Granularities:   []string{"1 week"},
							TimestampCol:    "timestamp",
							PredictionCol:   "prediction",
							ModelIdCol:      "model_id",
							ProblemType:     catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
							ForceSendFields: []string{"ModelIdCol", "PredictionCol", "TimestampCol"},
						},
						ForceSendFields: []string{"AssetsDir", "OutputSchemaName", "TableName"},
					}).Return(&catalog.MonitorInfo{
						AssetsDir:        "new_assets.dir",
						TableName:        "test_table",
						OutputSchemaName: "output.schema",
						Status:           catalog.MonitorInfoStatusMonitorStatusActive,
						InferenceLog: &catalog.MonitorInferenceLog{
							Granularities: []string{"1 week"},
							TimestampCol:  "timestamp",
							PredictionCol: "prediction",
							ModelIdCol:    "model_id",
							ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
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
							ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
						},
						DriftMetricsTableName: "test_table_drift",
					}, nil)
					e.DeleteByTableName(mock.Anything, "test_table").Return(nil)
				},
			},
		},
	}.RunUnitTest(t)
	assert.NoError(t, err)
}

func TestUpdateQualityMonitorUnit(t *testing.T) {
	err := ResourceFixturePluginFramework{
		Steps: []Step{
			{
				// CREATE
				Template: `
				resource "databricks_lakehouse_monitor_pluginframework" "testMonitorInference" {
					table_name = "test_table"
					assets_dir = "new_assets.dir"
					output_schema_name = "output.schema"
					inference_log = {
						granularities = ["1 week"]
						timestamp_col = "timestamp"
						prediction_col = "prediction"
						model_id_col = "model_id"
						problem_type = "PROBLEM_TYPE_REGRESSION"
					}
				}`,
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					e := w.GetMockQualityMonitorsAPI().EXPECT()
					e.Create(mock.Anything, catalog.CreateMonitor{
						AssetsDir:        "new_assets.dir",
						TableName:        "test_table",
						OutputSchemaName: "output.schema",
						InferenceLog: &catalog.MonitorInferenceLog{
							Granularities:   []string{"1 week"},
							TimestampCol:    "timestamp",
							PredictionCol:   "prediction",
							ModelIdCol:      "model_id",
							ProblemType:     catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
							ForceSendFields: []string{"ModelIdCol", "PredictionCol", "TimestampCol"},
						},
						ForceSendFields: []string{"AssetsDir", "OutputSchemaName", "TableName"},
					}).Return(&catalog.MonitorInfo{
						AssetsDir:        "new_assets.dir",
						TableName:        "test_table",
						OutputSchemaName: "output.schema",
						Status:           catalog.MonitorInfoStatusMonitorStatusActive,
						InferenceLog: &catalog.MonitorInferenceLog{
							Granularities: []string{"1 week"},
							TimestampCol:  "timestamp",
							PredictionCol: "prediction",
							ModelIdCol:    "model_id",
							ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
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
							ProblemType:   catalog.MonitorInferenceLogProblemTypeProblemTypeRegression,
						},
						DriftMetricsTableName: "test_table_drift",
					}, nil)
				},
			},
			{
				// UPDATE: Inference log granularity from 1 week to 1 day
				Template: `
				resource "databricks_lakehouse_monitor_pluginframework" "testMonitorInference" {
					table_name = "test_table"
					assets_dir = "new_assets.dir"
					output_schema_name = "output.schema"
					inference_log = {
						granularities = ["1 day"] # ==== UPDATE ====
						timestamp_col = "timestamp"
						prediction_col = "prediction"
						model_id_col = "model_id"
						problem_type = "PROBLEM_TYPE_REGRESSION"
					}
				}`,
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					e := w.GetMockQualityMonitorsAPI().EXPECT()
					e.Update(mock.Anything, catalog.UpdateMonitor{
						TableName:        "test_table",
						OutputSchemaName: "output.schema",
						InferenceLog: &catalog.MonitorInferenceLog{
							Granularities: []string{"1 day"},
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
							Granularities: []string{"1 day"},
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
							Granularities: []string{"1 day"},
							TimestampCol:  "timestamp",
							PredictionCol: "prediction",
							ModelIdCol:    "model_id",
						},
						DriftMetricsTableName: "test_table_drift",
					}, nil)
				},
			},
		},
	}.RunUnitTest(t)
	assert.NoError(t, err)
}
