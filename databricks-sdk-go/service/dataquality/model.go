// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataquality

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// The granularity for aggregating data into time windows based on their
// timestamp.
type AggregationGranularity string

const AggregationGranularityAggregationGranularity1Day AggregationGranularity = `AGGREGATION_GRANULARITY_1_DAY`

const AggregationGranularityAggregationGranularity1Hour AggregationGranularity = `AGGREGATION_GRANULARITY_1_HOUR`

const AggregationGranularityAggregationGranularity1Month AggregationGranularity = `AGGREGATION_GRANULARITY_1_MONTH`

const AggregationGranularityAggregationGranularity1Week AggregationGranularity = `AGGREGATION_GRANULARITY_1_WEEK`

const AggregationGranularityAggregationGranularity1Year AggregationGranularity = `AGGREGATION_GRANULARITY_1_YEAR`

const AggregationGranularityAggregationGranularity2Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_2_WEEKS`

const AggregationGranularityAggregationGranularity30Minutes AggregationGranularity = `AGGREGATION_GRANULARITY_30_MINUTES`

const AggregationGranularityAggregationGranularity3Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_3_WEEKS`

const AggregationGranularityAggregationGranularity4Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_4_WEEKS`

const AggregationGranularityAggregationGranularity5Minutes AggregationGranularity = `AGGREGATION_GRANULARITY_5_MINUTES`

// String representation for [fmt.Print]
func (f *AggregationGranularity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AggregationGranularity) Set(v string) error {
	switch v {
	case `AGGREGATION_GRANULARITY_1_DAY`, `AGGREGATION_GRANULARITY_1_HOUR`, `AGGREGATION_GRANULARITY_1_MONTH`, `AGGREGATION_GRANULARITY_1_WEEK`, `AGGREGATION_GRANULARITY_1_YEAR`, `AGGREGATION_GRANULARITY_2_WEEKS`, `AGGREGATION_GRANULARITY_30_MINUTES`, `AGGREGATION_GRANULARITY_3_WEEKS`, `AGGREGATION_GRANULARITY_4_WEEKS`, `AGGREGATION_GRANULARITY_5_MINUTES`:
		*f = AggregationGranularity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AGGREGATION_GRANULARITY_1_DAY", "AGGREGATION_GRANULARITY_1_HOUR", "AGGREGATION_GRANULARITY_1_MONTH", "AGGREGATION_GRANULARITY_1_WEEK", "AGGREGATION_GRANULARITY_1_YEAR", "AGGREGATION_GRANULARITY_2_WEEKS", "AGGREGATION_GRANULARITY_30_MINUTES", "AGGREGATION_GRANULARITY_3_WEEKS", "AGGREGATION_GRANULARITY_4_WEEKS", "AGGREGATION_GRANULARITY_5_MINUTES"`, v)
	}
}

// Values returns all possible values for AggregationGranularity.
//
// There is no guarantee on the order of the values in the slice.
func (f *AggregationGranularity) Values() []AggregationGranularity {
	return []AggregationGranularity{
		AggregationGranularityAggregationGranularity1Day,
		AggregationGranularityAggregationGranularity1Hour,
		AggregationGranularityAggregationGranularity1Month,
		AggregationGranularityAggregationGranularity1Week,
		AggregationGranularityAggregationGranularity1Year,
		AggregationGranularityAggregationGranularity2Weeks,
		AggregationGranularityAggregationGranularity30Minutes,
		AggregationGranularityAggregationGranularity3Weeks,
		AggregationGranularityAggregationGranularity4Weeks,
		AggregationGranularityAggregationGranularity5Minutes,
	}
}

// Type always returns AggregationGranularity to satisfy [pflag.Value] interface
func (f *AggregationGranularity) Type() string {
	return "AggregationGranularity"
}

// Anomaly Detection Configurations.
type AnomalyDetectionConfig struct {
	// List of fully qualified table names to exclude from anomaly detection.
	// Format: `catalog.schema.table_name`.
	ExcludedTableFullNames []string `json:"excluded_table_full_names,omitempty"`
	// Who to alert when anomaly detection finds unhealthy tables in this
	// schema. Applicable to `schema` monitors only. This is the same schema
	// alert you manage from **Manage alerts** in the Data Quality Monitoring
	// UI, so a change made in either place shows in both.
	//
	// Set the recipients in `on_unhealthy_table`; `on_failure` is not supported
	// here. Only `destination_ids` and `destination_user_ids` are supported.
	//
	// On update, an update mask of `*`, `anomaly_detection_config`, or
	// `anomaly_detection_config.notification_settings` replaces the alert's
	// recipients with the ones in the request; omitting the field, or leaving
	// both recipient lists empty, removes the alert. Other masks leave the
	// alert unchanged.
	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
}

func (s *AnomalyDetectionConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to cancel a refresh.
type CancelRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

func (s *CancelRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response to cancelling a refresh.
type CancelRefreshResponse struct {
	// The refresh to cancel.
	Refresh *Refresh `json:"refresh,omitempty"`
}

func (s *CancelRefreshResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateMonitorRequest struct {
	// The monitor to create.
	Monitor Monitor `json:"monitor"`
}

func (s *CreateMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// The refresh to create
	Refresh Refresh `json:"refresh"`
}

func (s *CreateRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The data quality monitoring workflow cron schedule.
type CronSchedule struct {
	// Read only field that indicates whether the schedule is paused or not.
	PauseStatus CronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// The Java timezone ID in which the cron expression is evaluated. See [Java
	// TimeZone] for the supported IDs.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/8/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The data quality monitoring workflow cron schedule pause status.
type CronSchedulePauseStatus string

const CronSchedulePauseStatusCronSchedulePauseStatusPaused CronSchedulePauseStatus = `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`

const CronSchedulePauseStatusCronSchedulePauseStatusUnpaused CronSchedulePauseStatus = `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`

// String representation for [fmt.Print]
func (f *CronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`, `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`:
		*f = CronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CRON_SCHEDULE_PAUSE_STATUS_PAUSED", "CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED"`, v)
	}
}

// Values returns all possible values for CronSchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *CronSchedulePauseStatus) Values() []CronSchedulePauseStatus {
	return []CronSchedulePauseStatus{
		CronSchedulePauseStatusCronSchedulePauseStatusPaused,
		CronSchedulePauseStatusCronSchedulePauseStatusUnpaused,
	}
}

// Type always returns CronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *CronSchedulePauseStatus) Type() string {
	return "CronSchedulePauseStatus"
}

// Data Profiling Configurations.
type DataProfilingConfig struct {
	// The absolute workspace path of the directory where the monitor's assets,
	// such as the dashboard, are stored. The UI and the Python APIs normally
	// prepopulate a default user location.
	AssetsDir string `json:"assets_dir,omitempty"`
	// The fully qualified name of a baseline table. Format:
	// `catalog.schema.table_name`. Drift metrics compare the monitored table
	// against the baseline data, so the two tables must have the same schema.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute in addition to the built-in ones. See [custom
	// metrics].
	//
	// [custom metrics]: https://docs.databricks.com/aws/en/data-governance/unity-catalog/data-quality-monitoring/data-profiling/custom-metrics
	CustomMetrics []DataProfilingCustomMetric `json:"custom_metrics,omitempty"`
	// The ID of the dashboard that visualizes the computed metrics. It is empty
	// while the monitor is pending, when the dashboard was skipped with
	// `skip_builtin_dashboard`, or when the caller is in a different workspace
	// from the one where the monitor was created.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Table that stores drift metrics data. Format:
	// `catalog.schema.table_name`.
	DriftMetricsTableName string `json:"drift_metrics_table_name,omitempty"`
	// The warehouse for dashboard creation
	EffectiveWarehouseId string `json:"effective_warehouse_id,omitempty"`
	// Profiles an inference log table, which stores the requests and
	// predictions of a model. Metrics are computed per time window and include
	// model quality metrics.
	InferenceLog *InferenceLogConfig `json:"inference_log,omitempty"`
	// The latest error message for a monitor failure.
	LatestMonitorFailureMessage string `json:"latest_monitor_failure_message,omitempty"`
	// Represents the current monitor configuration version in use. The version
	// will be represented in a numeric fashion (1,2,3...). The field has
	// flexibility to take on negative values, which can indicate corrupted
	// monitor_version numbers.
	MonitorVersion int64 `json:"monitor_version,omitempty"`
	// The fully qualified name of the monitored table. Format:
	// `catalog.schema.table_name`.
	MonitoredTableName string `json:"monitored_table_name,omitempty"`
	// Who to notify about this monitor's refreshes. Only `on_failure` with
	// `email_addresses` is supported here.
	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
	// The UUID of the Unity Catalog schema where the profile and drift metrics
	// tables are created.
	OutputSchemaId string `json:"output_schema_id"`
	// Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	ProfileMetricsTableName string `json:"profile_metrics_table_name,omitempty"`
	// The schedule on which the monitor refreshes its metrics. Refreshes can
	// also be triggered on demand through the Refresh resource.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For example
	// `slicing_exprs=["col_1", "col_2 > 10"]` will generate the following
	// slices: two slices for `col_2 > 10` (True and False), and one slice per
	// unique value in `col_1`. For high-cardinality columns, only the top 100
	// unique values by frequency will generate slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Profiles the full contents of the table. Metrics are computed over all
	// rows at each refresh.
	Snapshot *SnapshotConfig `json:"snapshot,omitempty"`
	// The status of the monitor. `DATA_PROFILING_STATUS_PENDING` means the
	// monitor's assets are being created or updated.
	// `DATA_PROFILING_STATUS_ACTIVE` means the monitor is ready and can be
	// refreshed. `DATA_PROFILING_STATUS_FAILED` and
	// `DATA_PROFILING_STATUS_ERROR` mean the monitor is not operational, and
	// `latest_monitor_failure_message` describes the cause when one was
	// recorded. `DATA_PROFILING_STATUS_DELETE_PENDING` means the monitor is
	// being deleted.
	Status DataProfilingStatus `json:"status,omitempty"`
	// Profiles a time series table. Metrics are computed per time window, as
	// defined by `timestamp_column` and `granularities`.
	TimeSeries *TimeSeriesConfig `json:"time_series,omitempty"`
	// The ID of the SQL warehouse used to create the dashboard. If not
	// specified, the first running warehouse the caller can access is used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataProfilingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataProfilingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom metric definition.
type DataProfilingCustomMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use `":table"` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	Name string `json:"name"`
	// Specifies the Spark SQL data type of the metric value, either as a plain
	// type name or as a JSON-encoded StructField.
	OutputDataType string `json:"output_data_type"`
	// The type of the custom metric.
	// `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE` metrics depend only on the
	// columns of the input table, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`
	// metrics depend only on previously computed aggregate metrics, and
	// `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT` metrics depend on previously
	// computed aggregate or derived metrics.
	Type DataProfilingCustomMetricType `json:"type"`
}

func (s *DataProfilingCustomMetric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The custom metric type.
type DataProfilingCustomMetricType string

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeAggregate DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDerived DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDrift DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`

// String representation for [fmt.Print]
func (f *DataProfilingCustomMetricType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataProfilingCustomMetricType) Set(v string) error {
	switch v {
	case `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`:
		*f = DataProfilingCustomMetricType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE", "DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED", "DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT"`, v)
	}
}

// Values returns all possible values for DataProfilingCustomMetricType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataProfilingCustomMetricType) Values() []DataProfilingCustomMetricType {
	return []DataProfilingCustomMetricType{
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeAggregate,
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDerived,
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDrift,
	}
}

// Type always returns DataProfilingCustomMetricType to satisfy [pflag.Value] interface
func (f *DataProfilingCustomMetricType) Type() string {
	return "DataProfilingCustomMetricType"
}

// The status of the data profiling monitor.
type DataProfilingStatus string

const DataProfilingStatusDataProfilingStatusActive DataProfilingStatus = `DATA_PROFILING_STATUS_ACTIVE`

const DataProfilingStatusDataProfilingStatusDeletePending DataProfilingStatus = `DATA_PROFILING_STATUS_DELETE_PENDING`

const DataProfilingStatusDataProfilingStatusError DataProfilingStatus = `DATA_PROFILING_STATUS_ERROR`

const DataProfilingStatusDataProfilingStatusFailed DataProfilingStatus = `DATA_PROFILING_STATUS_FAILED`

const DataProfilingStatusDataProfilingStatusPending DataProfilingStatus = `DATA_PROFILING_STATUS_PENDING`

// String representation for [fmt.Print]
func (f *DataProfilingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataProfilingStatus) Set(v string) error {
	switch v {
	case `DATA_PROFILING_STATUS_ACTIVE`, `DATA_PROFILING_STATUS_DELETE_PENDING`, `DATA_PROFILING_STATUS_ERROR`, `DATA_PROFILING_STATUS_FAILED`, `DATA_PROFILING_STATUS_PENDING`:
		*f = DataProfilingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_PROFILING_STATUS_ACTIVE", "DATA_PROFILING_STATUS_DELETE_PENDING", "DATA_PROFILING_STATUS_ERROR", "DATA_PROFILING_STATUS_FAILED", "DATA_PROFILING_STATUS_PENDING"`, v)
	}
}

// Values returns all possible values for DataProfilingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataProfilingStatus) Values() []DataProfilingStatus {
	return []DataProfilingStatus{
		DataProfilingStatusDataProfilingStatusActive,
		DataProfilingStatusDataProfilingStatusDeletePending,
		DataProfilingStatusDataProfilingStatusError,
		DataProfilingStatusDataProfilingStatusFailed,
		DataProfilingStatusDataProfilingStatusPending,
	}
}

// Type always returns DataProfilingStatus to satisfy [pflag.Value] interface
func (f *DataProfilingStatus) Type() string {
	return "DataProfilingStatus"
}

type DeleteMonitorRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
}

func (s *DeleteMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

func (s *DeleteRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMonitorRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
}

func (s *GetMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

func (s *GetRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Inference log configuration.
type InferenceLogConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities []AggregationGranularity `json:"granularities"`
	// Column containing the ground truth for the model's predictions.
	LabelColumn string `json:"label_column,omitempty"`
	// Column containing the ID of the model that produced each prediction. A
	// slice is created automatically for each distinct model ID.
	ModelIdColumn string `json:"model_id_column"`
	// Column containing the model's predicted values.
	PredictionColumn string `json:"prediction_column"`
	// The type of problem the model solves. It determines which model quality
	// metrics are computed.
	ProblemType InferenceProblemType `json:"problem_type"`
	// Column containing the timestamp of each inference request. Its type must
	// be TIMESTAMP or convertible to a timestamp.
	TimestampColumn string `json:"timestamp_column"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InferenceLogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InferenceLogConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Inference problem type the model aims to solve.
type InferenceProblemType string

const InferenceProblemTypeInferenceProblemTypeClassification InferenceProblemType = `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`

const InferenceProblemTypeInferenceProblemTypeRegression InferenceProblemType = `INFERENCE_PROBLEM_TYPE_REGRESSION`

// String representation for [fmt.Print]
func (f *InferenceProblemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InferenceProblemType) Set(v string) error {
	switch v {
	case `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`, `INFERENCE_PROBLEM_TYPE_REGRESSION`:
		*f = InferenceProblemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INFERENCE_PROBLEM_TYPE_CLASSIFICATION", "INFERENCE_PROBLEM_TYPE_REGRESSION"`, v)
	}
}

// Values returns all possible values for InferenceProblemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *InferenceProblemType) Values() []InferenceProblemType {
	return []InferenceProblemType{
		InferenceProblemTypeInferenceProblemTypeClassification,
		InferenceProblemTypeInferenceProblemTypeRegression,
	}
}

// Type always returns InferenceProblemType to satisfy [pflag.Value] interface
func (f *InferenceProblemType) Type() string {
	return "InferenceProblemType"
}

type ListMonitorRequest struct {
	// The maximum number of monitors to return in one page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token from a previous response, used to fetch the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMonitorRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing Monitors.
type ListMonitorResponse struct {
	// The monitors on this page.
	Monitors []Monitor `json:"monitors,omitempty"`
	// Pagination token to pass as `page_token` to fetch the next page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMonitorResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMonitorResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRefreshRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing refreshes.
type ListRefreshResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Refreshes []Refresh `json:"refreshes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRefreshResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRefreshResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A data quality monitor on a Unity Catalog object. A `schema` monitor runs
// anomaly detection on the tables in the schema and is configured through
// `anomaly_detection_config`. A `table` monitor profiles a single table and is
// configured through `data_profiling_config`.
type Monitor struct {
	// Anomaly detection settings. Required when `object_type` is `schema`. When
	// reading a `table` monitor, this holds the settings of the parent schema's
	// monitor if the schema has one.
	AnomalyDetectionConfig *AnomalyDetectionConfig `json:"anomaly_detection_config,omitempty"`
	// Data profiling settings. Required when `object_type` is `table`. Exactly
	// one of `inference_log`, `time_series` or `snapshot` must be set.
	DataProfilingConfig *DataProfilingConfig `json:"data_profiling_config,omitempty"`
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"object_type"`
}

func (s *Monitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Destination of the data quality monitoring notification.
type NotificationDestination struct {
	// The notification-destination IDs (Slack / Teams / webhook / Asana /
	// ServiceNow / AWS EventBridge) to send the notification to. The
	// destinations must already exist in the workspace; a workspace admin
	// creates them. A maximum of 25 destination IDs is supported. Only
	// supported on schema monitors, via `anomaly_detection_config`; rejected on
	// table monitors.
	DestinationIds []string `json:"destination_ids,omitempty"`
	// The IDs of the workspace users to notify by email. A maximum of 25 user
	// IDs is supported. Only supported on schema monitors, via
	// `anomaly_detection_config`; rejected on table monitors.
	DestinationUserIds []int64 `json:"destination_user_ids,omitempty"`
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported. Only supported on table monitors, via
	// `data_profiling_config`; rejected on schema monitors.
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func (s *NotificationDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Settings for sending notifications on the data quality monitoring.
type NotificationSettings struct {
	// Recipients to notify when a refresh of the monitor fails or times out.
	// Only supported on table monitors, via `data_profiling_config`, and only
	// `email_addresses` is accepted there.
	OnFailure *NotificationDestination `json:"on_failure,omitempty"`
	// Recipients to notify when anomaly detection finds unhealthy tables in the
	// schema. Each recipient gets one notification per important unhealthy
	// table. Only supported on schema monitors, via `anomaly_detection_config`.
	OnUnhealthyTable *NotificationDestination `json:"on_unhealthy_table,omitempty"`
}

func (s *NotificationSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The Refresh object gives information on a refresh of the data quality
// monitoring pipeline.
type Refresh struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message string `json:"message,omitempty"`
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"object_type"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"refresh_id,omitempty"`
	// Time when the refresh started (milliseconds since 1/1/1970 UTC).
	StartTimeMs int64 `json:"start_time_ms,omitempty"`
	// The current state of the refresh.
	State RefreshState `json:"state,omitempty"`
	// What triggered the refresh.
	Trigger RefreshTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Refresh) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Refresh) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the refresh.
type RefreshState string

const RefreshStateMonitorRefreshStateCanceled RefreshState = `MONITOR_REFRESH_STATE_CANCELED`

const RefreshStateMonitorRefreshStateFailed RefreshState = `MONITOR_REFRESH_STATE_FAILED`

const RefreshStateMonitorRefreshStatePending RefreshState = `MONITOR_REFRESH_STATE_PENDING`

const RefreshStateMonitorRefreshStateRunning RefreshState = `MONITOR_REFRESH_STATE_RUNNING`

const RefreshStateMonitorRefreshStateSuccess RefreshState = `MONITOR_REFRESH_STATE_SUCCESS`

const RefreshStateMonitorRefreshStateUnknown RefreshState = `MONITOR_REFRESH_STATE_UNKNOWN`

// String representation for [fmt.Print]
func (f *RefreshState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RefreshState) Set(v string) error {
	switch v {
	case `MONITOR_REFRESH_STATE_CANCELED`, `MONITOR_REFRESH_STATE_FAILED`, `MONITOR_REFRESH_STATE_PENDING`, `MONITOR_REFRESH_STATE_RUNNING`, `MONITOR_REFRESH_STATE_SUCCESS`, `MONITOR_REFRESH_STATE_UNKNOWN`:
		*f = RefreshState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_REFRESH_STATE_CANCELED", "MONITOR_REFRESH_STATE_FAILED", "MONITOR_REFRESH_STATE_PENDING", "MONITOR_REFRESH_STATE_RUNNING", "MONITOR_REFRESH_STATE_SUCCESS", "MONITOR_REFRESH_STATE_UNKNOWN"`, v)
	}
}

// Values returns all possible values for RefreshState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RefreshState) Values() []RefreshState {
	return []RefreshState{
		RefreshStateMonitorRefreshStateCanceled,
		RefreshStateMonitorRefreshStateFailed,
		RefreshStateMonitorRefreshStatePending,
		RefreshStateMonitorRefreshStateRunning,
		RefreshStateMonitorRefreshStateSuccess,
		RefreshStateMonitorRefreshStateUnknown,
	}
}

// Type always returns RefreshState to satisfy [pflag.Value] interface
func (f *RefreshState) Type() string {
	return "RefreshState"
}

// The trigger of the refresh.
type RefreshTrigger string

const RefreshTriggerMonitorRefreshTriggerDataChange RefreshTrigger = `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`

const RefreshTriggerMonitorRefreshTriggerManual RefreshTrigger = `MONITOR_REFRESH_TRIGGER_MANUAL`

const RefreshTriggerMonitorRefreshTriggerSchedule RefreshTrigger = `MONITOR_REFRESH_TRIGGER_SCHEDULE`

const RefreshTriggerMonitorRefreshTriggerUnknown RefreshTrigger = `MONITOR_REFRESH_TRIGGER_UNKNOWN`

// String representation for [fmt.Print]
func (f *RefreshTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RefreshTrigger) Set(v string) error {
	switch v {
	case `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`, `MONITOR_REFRESH_TRIGGER_MANUAL`, `MONITOR_REFRESH_TRIGGER_SCHEDULE`, `MONITOR_REFRESH_TRIGGER_UNKNOWN`:
		*f = RefreshTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_REFRESH_TRIGGER_DATA_CHANGE", "MONITOR_REFRESH_TRIGGER_MANUAL", "MONITOR_REFRESH_TRIGGER_SCHEDULE", "MONITOR_REFRESH_TRIGGER_UNKNOWN"`, v)
	}
}

// Values returns all possible values for RefreshTrigger.
//
// There is no guarantee on the order of the values in the slice.
func (f *RefreshTrigger) Values() []RefreshTrigger {
	return []RefreshTrigger{
		RefreshTriggerMonitorRefreshTriggerDataChange,
		RefreshTriggerMonitorRefreshTriggerManual,
		RefreshTriggerMonitorRefreshTriggerSchedule,
		RefreshTriggerMonitorRefreshTriggerUnknown,
	}
}

// Type always returns RefreshTrigger to satisfy [pflag.Value] interface
func (f *RefreshTrigger) Type() string {
	return "RefreshTrigger"
}

// Snapshot analysis configuration.
type SnapshotConfig struct {
}

func (s *SnapshotConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Time series analysis configuration.
type TimeSeriesConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities []AggregationGranularity `json:"granularities"`
	// Column containing the timestamp of each row. Its type must be TIMESTAMP
	// or convertible to a timestamp.
	TimestampColumn string `json:"timestamp_column"`
}

func (s *TimeSeriesConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateMonitorRequest struct {
	// The new values for the monitor. Only the fields named in `update_mask`
	// are applied.
	Monitor Monitor `json:"monitor"`
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// The field mask to specify which fields to update as a comma-separated
	// list. Example value:
	// `data_profiling_config.custom_metrics,data_profiling_config.schedule.quartz_cron_expression`
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateRefreshRequest struct {
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// The refresh to update.
	Refresh Refresh `json:"refresh"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
	// The field mask to specify which fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
