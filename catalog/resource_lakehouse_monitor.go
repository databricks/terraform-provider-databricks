package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func removeSnapshotField(originalSchema map[string]*schema.Schema) map[string]*schema.Schema {
	tmpSchema := make(map[string]*schema.Schema)
	for k, v := range originalSchema {
		tmpSchema[k] = v
	}
	delete(tmpSchema, "snapshot")
	return tmpSchema
}

type Snapshot struct {
}
type Monitor struct {
	AssetsDir                string                                   `json:"assets_dir"`
	BaselineTableName        string                                   `json:"baseline_table_name,omitempty"`
	CustomMetrics            []catalog.MonitorCustomMetric            `json:"custom_metrics,omitempty"`
	DataClassificationConfig *catalog.MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	FullName                 string                                   `json:"table_name"`
	InferenceLog             *catalog.MonitorInferenceLogProfileType  `json:"inference_log,omitempty"`
	Notifications            []catalog.MonitorNotificationsConfig     `json:"notifications,omitempty"`
	OutputSchemaName         string                                   `json:"output_schema_name"`
	Schedule                 *catalog.MonitorCronSchedule             `json:"schedule,omitempty"`
	Snapshot                 *Snapshot                                `json:"snapshot,omitempty"`
	SkipBuiltinDashboard     bool                                     `json:"skip_builtin_dashboard,omitempty"`
	SlicingExprs             []string                                 `json:"slicing_exprs,omitempty"`
	TimeSeries               *catalog.MonitorTimeSeriesProfileType    `json:"time_series,omitempty"`
	WarehouseId              string                                   `json:"warehouse_id,omitempty"`
}

type MonitorInfo struct {
	Monitor
	DriftMetricsTableName   string                    `json:"drift_metrics_table_name,omitempty"`
	ProfileMetricsTableName string                    `json:"profile_metrics_table_name,omitempty"`
	Status                  catalog.MonitorInfoStatus `json:"status,omitempty"`
	TableName               string                    `json:"table_name,omitempty"`
}

func ResourceLakehouseMonitor() common.Resource {
	monitorSchema := common.StructToSchema(
		MonitorInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m, "output_schema_name").SetRequired()

			common.CustomizeSchemaPath(m, "inference_log", "granularities").SetOptional()
			common.CustomizeSchemaPath(m, "inference_log", "problem_type").SetOptional()
			common.CustomizeSchemaPath(m, "inference_log", "timestamp_col").SetOptional()
			common.CustomizeSchemaPath(m, "inference_log", "prediction_col").SetDefault("prediction")

			common.CustomizeSchemaPath(m, "time_series", "granularities").SetOptional()
			common.CustomizeSchemaPath(m, "time_series", "timestamp_col").SetOptional()

			common.CustomizeSchemaPath(m, "drift_metrics_table_name").SetComputed()
			common.CustomizeSchemaPath(m, "profile_metrics_table_name").SetComputed()
			common.CustomizeSchemaPath(m, "status").SetComputed()

			return m
		},
	)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			tmpSchema := removeSnapshotField(monitorSchema)

			var create catalog.CreateMonitor
			common.DataToStructPointer(d, tmpSchema, &create)
			if _, ok := d.GetOk("snapshot"); ok {
				create.Snapshot = struct{}{}
			}
			create.FullName = d.Get("table_name").(string)

			endpoint, err := w.LakehouseMonitors.Create(ctx, create)
			if err != nil {
				return err
			}
			d.SetId(endpoint.TableName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			endpoint, err := w.LakehouseMonitors.GetByFullName(ctx, d.Id())
			if err != nil {
				return err

			}
			err = common.StructToData(endpoint, monitorSchema, d)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.UpdateMonitor
			common.DataToStructPointer(d, monitorSchema, &update)
			if _, ok := d.GetOk("snapshot"); ok {
				update.Snapshot = struct{}{}
			}
			update.FullName = d.Get("table_name").(string)
			_, err = w.LakehouseMonitors.Update(ctx, update)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.LakehouseMonitors.DeleteByFullName(ctx, d.Id())
		},
		Schema: monitorSchema,
	}
}
