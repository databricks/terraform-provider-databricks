package pluginframework

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	pluginFrameworkResource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQualityMonitorCreateTimeseriesPluginFramework(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
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
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)
	createRequest := pluginFrameworkResource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw: MapToTfTypesValue(map[string]any{
				"table_name":         "test_table",
				"assets_dir":         "sample.dir",
				"output_schema_name": "output.schema",
				"time_series":        `{"granularities":["1 day"],"timestamp_col":"timestamp"}`,
			}),
			Schema: schema.Schema,
		},
	}
	createResponse := pluginFrameworkResource.CreateResponse{}

	qualityMonitorResource.Create(ctx, createRequest, &createResponse)

	assert.False(t, createResponse.Diagnostics.HasError())
}

func TestQualityMonitorCreateInferencePluginFramework(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
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
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)
	createRequest := pluginFrameworkResource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw: MapToTfTypesValue(map[string]any{
				"table_name":         "test_table",
				"assets_dir":         "sample.dir",
				"output_schema_name": "output.schema",
				"inference_log": `{
					granularities: ["1 day"],
					timestamp_col: "timestamp",
					prediction_col: "prediction",
					model_id_col: "model_id",
					problem_type: "PROBLEM_TYPE_REGRESSION"
				}`,
			}),
			Schema: schema.Schema,
		},
	}
	createResponse := pluginFrameworkResource.CreateResponse{}

	qualityMonitorResource.Create(ctx, createRequest, &createResponse)

	assert.False(t, createResponse.Diagnostics.HasError())
}

func TestQualityMonitorCreateSnapshot(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
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
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)

	createRequest := pluginFrameworkResource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw: MapToTfTypesValue(map[string]any{
				"table_name":         "test_table",
				"assets_dir":         "sample.dir",
				"output_schema_name": "output.schema",
				"snapshot":           `{}`,
			}),
			Schema: schema.Schema,
		},
	}
	createResponse := pluginFrameworkResource.CreateResponse{}

	qualityMonitorResource.Create(ctx, createRequest, &createResponse)

	assert.False(t, createResponse.Diagnostics.HasError())
}

func TestQualityMonitorGet(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
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
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)
}

func TestQualityMonitorUpdate(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
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
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)
}

func TestQualityMonitorDelete(t *testing.T) {
	ctx := context.Background()
	client, err := ResourceFixturePluginFramework{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQualityMonitorsAPI().EXPECT()
			e.DeleteByTableName(mock.Anything, "test_table").Return(nil)
		},
	}.Start(t)
	assert.NoError(t, err)

	qualityMonitorResource := QualityMonitorResource{
		Client: client,
	}

	schema := pluginFrameworkResource.SchemaResponse{}
	qualityMonitorResource.Schema(ctx, pluginFrameworkResource.SchemaRequest{}, &schema)
}
