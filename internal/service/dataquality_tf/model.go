// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package dataquality_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Anomaly Detection Configurations.
type AnomalyDetectionConfig struct {
}

func (to *AnomalyDetectionConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AnomalyDetectionConfig) {
}

func (to *AnomalyDetectionConfig) SyncFieldsDuringRead(ctx context.Context, from AnomalyDetectionConfig) {
}

func (m AnomalyDetectionConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AnomalyDetectionConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AnomalyDetectionConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnomalyDetectionConfig
// only implements ToObjectValue() and Type().
func (m AnomalyDetectionConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m AnomalyDetectionConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *CancelRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelRefreshRequest) {
}

func (to *CancelRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from CancelRefreshRequest) {
}

func (m CancelRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshRequest
// only implements ToObjectValue() and Type().
func (m CancelRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
}

// Response to cancelling a refresh.
type CancelRefreshResponse struct {
	// The refresh to cancel.
	Refresh types.Object `tfsdk:"refresh"`
}

func (to *CancelRefreshResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelRefreshResponse) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				// Recursively sync the fields of Refresh
				toRefresh.SyncFieldsDuringCreateOrUpdate(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (to *CancelRefreshResponse) SyncFieldsDuringRead(ctx context.Context, from CancelRefreshResponse) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m CancelRefreshResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelRefreshResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshResponse
// only implements ToObjectValue() and Type().
func (m CancelRefreshResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh": m.Refresh,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelRefreshResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh": Refresh{}.Type(ctx),
		},
	}
}

// GetRefresh returns the value of the Refresh field in CancelRefreshResponse as
// a Refresh value.
// If the field is unknown or null, the boolean return value is false.
func (m *CancelRefreshResponse) GetRefresh(ctx context.Context) (Refresh, bool) {
	var e Refresh
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v Refresh
	d := m.Refresh.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefresh sets the value of the Refresh field in CancelRefreshResponse.
func (m *CancelRefreshResponse) SetRefresh(ctx context.Context, v Refresh) {
	vs := v.ToObjectValue(ctx)
	m.Refresh = vs
}

type CreateMonitorRequest struct {
	// The monitor to create.
	Monitor types.Object `tfsdk:"monitor"`
}

func (to *CreateMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateMonitorRequest) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				// Recursively sync the fields of Monitor
				toMonitor.SyncFieldsDuringCreateOrUpdate(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (to *CreateMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from CreateMonitorRequest) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				toMonitor.SyncFieldsDuringRead(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (m CreateMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["monitor"] = attrs["monitor"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitor": reflect.TypeOf(Monitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMonitorRequest
// only implements ToObjectValue() and Type().
func (m CreateMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"monitor": m.Monitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitor": Monitor{}.Type(ctx),
		},
	}
}

// GetMonitor returns the value of the Monitor field in CreateMonitorRequest as
// a Monitor value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateMonitorRequest) GetMonitor(ctx context.Context) (Monitor, bool) {
	var e Monitor
	if m.Monitor.IsNull() || m.Monitor.IsUnknown() {
		return e, false
	}
	var v Monitor
	d := m.Monitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMonitor sets the value of the Monitor field in CreateMonitorRequest.
func (m *CreateMonitorRequest) SetMonitor(ctx context.Context, v Monitor) {
	vs := v.ToObjectValue(ctx)
	m.Monitor = vs
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType types.String `tfsdk:"-"`
	// The refresh to create
	Refresh types.Object `tfsdk:"refresh"`
}

func (to *CreateRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRefreshRequest) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				// Recursively sync the fields of Refresh
				toRefresh.SyncFieldsDuringCreateOrUpdate(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (to *CreateRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from CreateRefreshRequest) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m CreateRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRefreshRequest
// only implements ToObjectValue() and Type().
func (m CreateRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh":     m.Refresh,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh":     Refresh{}.Type(ctx),
		},
	}
}

// GetRefresh returns the value of the Refresh field in CreateRefreshRequest as
// a Refresh value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRefreshRequest) GetRefresh(ctx context.Context) (Refresh, bool) {
	var e Refresh
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v Refresh
	d := m.Refresh.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefresh sets the value of the Refresh field in CreateRefreshRequest.
func (m *CreateRefreshRequest) SetRefresh(ctx context.Context, v Refresh) {
	vs := v.ToObjectValue(ctx)
	m.Refresh = vs
}

// The data quality monitoring workflow cron schedule.
type CronSchedule struct {
	// Read only field that indicates whether the schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression"`
	// A Java timezone id. The schedule for a job will be resolved with respect
	// to this timezone. See `Java TimeZone
	// <http://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html>`_ for
	// details. The timezone id (e.g., ``America/Los_Angeles``) in which to
	// evaluate the quartz expression.
	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (to *CronSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronSchedule) {
}

func (to *CronSchedule) SyncFieldsDuringRead(ctx context.Context, from CronSchedule) {
}

func (m CronSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pause_status"] = attrs["pause_status"].SetComputed()
	attrs["quartz_cron_expression"] = attrs["quartz_cron_expression"].SetRequired()
	attrs["timezone_id"] = attrs["timezone_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CronSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronSchedule
// only implements ToObjectValue() and Type().
func (m CronSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status":           m.PauseStatus,
			"quartz_cron_expression": m.QuartzCronExpression,
			"timezone_id":            m.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CronSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

// Data Profiling Configurations.
type DataProfilingConfig struct {
	// Field for specifying the absolute path to a custom directory to store
	// data-monitoring assets. Normally prepopulated to a default user location
	// via UI and Python APIs.
	AssetsDir types.String `tfsdk:"assets_dir"`
	// Baseline table name. Baseline data is used to compute drift from the data
	// in the monitored `table_name`. The baseline table and the monitored table
	// shall have the same schema.
	BaselineTableName types.String `tfsdk:"baseline_table_name"`
	// Custom metrics.
	CustomMetrics types.List `tfsdk:"custom_metrics"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// Table that stores drift metrics data. Format:
	// `catalog.schema.table_name`.
	DriftMetricsTableName types.String `tfsdk:"drift_metrics_table_name"`
	// The warehouse for dashboard creation
	EffectiveWarehouseId types.String `tfsdk:"effective_warehouse_id"`
	// `Analysis Configuration` for monitoring inference log tables.
	InferenceLog types.Object `tfsdk:"inference_log"`
	// The latest error message for a monitor failure.
	LatestMonitorFailureMessage types.String `tfsdk:"latest_monitor_failure_message"`
	// Represents the current monitor configuration version in use. The version
	// will be represented in a numeric fashion (1,2,3...). The field has
	// flexibility to take on negative values, which can indicate corrupted
	// monitor_version numbers.
	MonitorVersion types.Int64 `tfsdk:"monitor_version"`
	// Unity Catalog table to monitor. Format: `catalog.schema.table_name`
	MonitoredTableName types.String `tfsdk:"monitored_table_name"`
	// Field for specifying notification settings.
	NotificationSettings types.Object `tfsdk:"notification_settings"`
	// ID of the schema where output tables are created.
	OutputSchemaId types.String `tfsdk:"output_schema_id"`
	// Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	ProfileMetricsTableName types.String `tfsdk:"profile_metrics_table_name"`
	// The cron schedule.
	Schedule types.Object `tfsdk:"schedule"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard types.Bool `tfsdk:"skip_builtin_dashboard"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For example
	// `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs"`
	// `Analysis Configuration` for monitoring snapshot tables.
	Snapshot types.Object `tfsdk:"snapshot"`
	// The data profiling monitor status.
	Status types.String `tfsdk:"status"`
	// `Analysis Configuration` for monitoring time series tables.
	TimeSeries types.Object `tfsdk:"time_series"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *DataProfilingConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataProfilingConfig) {
	if !from.CustomMetrics.IsNull() && !from.CustomMetrics.IsUnknown() && to.CustomMetrics.IsNull() && len(from.CustomMetrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomMetrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomMetrics = from.CustomMetrics
	}
	if !from.InferenceLog.IsNull() && !from.InferenceLog.IsUnknown() {
		if toInferenceLog, ok := to.GetInferenceLog(ctx); ok {
			if fromInferenceLog, ok := from.GetInferenceLog(ctx); ok {
				// Recursively sync the fields of InferenceLog
				toInferenceLog.SyncFieldsDuringCreateOrUpdate(ctx, fromInferenceLog)
				to.SetInferenceLog(ctx, toInferenceLog)
			}
		}
	}
	if !from.NotificationSettings.IsNull() && !from.NotificationSettings.IsUnknown() {
		if toNotificationSettings, ok := to.GetNotificationSettings(ctx); ok {
			if fromNotificationSettings, ok := from.GetNotificationSettings(ctx); ok {
				// Recursively sync the fields of NotificationSettings
				toNotificationSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromNotificationSettings)
				to.SetNotificationSettings(ctx, toNotificationSettings)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				// Recursively sync the fields of Schedule
				toSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.SlicingExprs.IsNull() && !from.SlicingExprs.IsUnknown() && to.SlicingExprs.IsNull() && len(from.SlicingExprs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SlicingExprs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SlicingExprs = from.SlicingExprs
	}
	if !from.Snapshot.IsNull() && !from.Snapshot.IsUnknown() {
		if toSnapshot, ok := to.GetSnapshot(ctx); ok {
			if fromSnapshot, ok := from.GetSnapshot(ctx); ok {
				// Recursively sync the fields of Snapshot
				toSnapshot.SyncFieldsDuringCreateOrUpdate(ctx, fromSnapshot)
				to.SetSnapshot(ctx, toSnapshot)
			}
		}
	}
	if !from.TimeSeries.IsNull() && !from.TimeSeries.IsUnknown() {
		if toTimeSeries, ok := to.GetTimeSeries(ctx); ok {
			if fromTimeSeries, ok := from.GetTimeSeries(ctx); ok {
				// Recursively sync the fields of TimeSeries
				toTimeSeries.SyncFieldsDuringCreateOrUpdate(ctx, fromTimeSeries)
				to.SetTimeSeries(ctx, toTimeSeries)
			}
		}
	}
}

func (to *DataProfilingConfig) SyncFieldsDuringRead(ctx context.Context, from DataProfilingConfig) {
	if !from.CustomMetrics.IsNull() && !from.CustomMetrics.IsUnknown() && to.CustomMetrics.IsNull() && len(from.CustomMetrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomMetrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomMetrics = from.CustomMetrics
	}
	if !from.InferenceLog.IsNull() && !from.InferenceLog.IsUnknown() {
		if toInferenceLog, ok := to.GetInferenceLog(ctx); ok {
			if fromInferenceLog, ok := from.GetInferenceLog(ctx); ok {
				toInferenceLog.SyncFieldsDuringRead(ctx, fromInferenceLog)
				to.SetInferenceLog(ctx, toInferenceLog)
			}
		}
	}
	if !from.NotificationSettings.IsNull() && !from.NotificationSettings.IsUnknown() {
		if toNotificationSettings, ok := to.GetNotificationSettings(ctx); ok {
			if fromNotificationSettings, ok := from.GetNotificationSettings(ctx); ok {
				toNotificationSettings.SyncFieldsDuringRead(ctx, fromNotificationSettings)
				to.SetNotificationSettings(ctx, toNotificationSettings)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				toSchedule.SyncFieldsDuringRead(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
			}
		}
	}
	if !from.SlicingExprs.IsNull() && !from.SlicingExprs.IsUnknown() && to.SlicingExprs.IsNull() && len(from.SlicingExprs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SlicingExprs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SlicingExprs = from.SlicingExprs
	}
	if !from.Snapshot.IsNull() && !from.Snapshot.IsUnknown() {
		if toSnapshot, ok := to.GetSnapshot(ctx); ok {
			if fromSnapshot, ok := from.GetSnapshot(ctx); ok {
				toSnapshot.SyncFieldsDuringRead(ctx, fromSnapshot)
				to.SetSnapshot(ctx, toSnapshot)
			}
		}
	}
	if !from.TimeSeries.IsNull() && !from.TimeSeries.IsUnknown() {
		if toTimeSeries, ok := to.GetTimeSeries(ctx); ok {
			if fromTimeSeries, ok := from.GetTimeSeries(ctx); ok {
				toTimeSeries.SyncFieldsDuringRead(ctx, fromTimeSeries)
				to.SetTimeSeries(ctx, toTimeSeries)
			}
		}
	}
}

func (m DataProfilingConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets_dir"] = attrs["assets_dir"].SetOptional()
	attrs["assets_dir"] = attrs["assets_dir"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["baseline_table_name"] = attrs["baseline_table_name"].SetOptional()
	attrs["custom_metrics"] = attrs["custom_metrics"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["drift_metrics_table_name"] = attrs["drift_metrics_table_name"].SetComputed()
	attrs["effective_warehouse_id"] = attrs["effective_warehouse_id"].SetComputed()
	attrs["inference_log"] = attrs["inference_log"].SetOptional()
	attrs["latest_monitor_failure_message"] = attrs["latest_monitor_failure_message"].SetComputed()
	attrs["monitor_version"] = attrs["monitor_version"].SetComputed()
	attrs["monitored_table_name"] = attrs["monitored_table_name"].SetComputed()
	attrs["notification_settings"] = attrs["notification_settings"].SetOptional()
	attrs["output_schema_id"] = attrs["output_schema_id"].SetRequired()
	attrs["profile_metrics_table_name"] = attrs["profile_metrics_table_name"].SetComputed()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["skip_builtin_dashboard"] = attrs["skip_builtin_dashboard"].SetOptional()
	attrs["slicing_exprs"] = attrs["slicing_exprs"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["time_series"] = attrs["time_series"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataProfilingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataProfilingConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":        reflect.TypeOf(DataProfilingCustomMetric{}),
		"inference_log":         reflect.TypeOf(InferenceLogConfig{}),
		"notification_settings": reflect.TypeOf(NotificationSettings{}),
		"schedule":              reflect.TypeOf(CronSchedule{}),
		"slicing_exprs":         reflect.TypeOf(types.String{}),
		"snapshot":              reflect.TypeOf(SnapshotConfig{}),
		"time_series":           reflect.TypeOf(TimeSeriesConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataProfilingConfig
// only implements ToObjectValue() and Type().
func (m DataProfilingConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets_dir":                     m.AssetsDir,
			"baseline_table_name":            m.BaselineTableName,
			"custom_metrics":                 m.CustomMetrics,
			"dashboard_id":                   m.DashboardId,
			"drift_metrics_table_name":       m.DriftMetricsTableName,
			"effective_warehouse_id":         m.EffectiveWarehouseId,
			"inference_log":                  m.InferenceLog,
			"latest_monitor_failure_message": m.LatestMonitorFailureMessage,
			"monitor_version":                m.MonitorVersion,
			"monitored_table_name":           m.MonitoredTableName,
			"notification_settings":          m.NotificationSettings,
			"output_schema_id":               m.OutputSchemaId,
			"profile_metrics_table_name":     m.ProfileMetricsTableName,
			"schedule":                       m.Schedule,
			"skip_builtin_dashboard":         m.SkipBuiltinDashboard,
			"slicing_exprs":                  m.SlicingExprs,
			"snapshot":                       m.Snapshot,
			"status":                         m.Status,
			"time_series":                    m.TimeSeries,
			"warehouse_id":                   m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataProfilingConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: DataProfilingCustomMetric{}.Type(ctx),
			},
			"dashboard_id":                   types.StringType,
			"drift_metrics_table_name":       types.StringType,
			"effective_warehouse_id":         types.StringType,
			"inference_log":                  InferenceLogConfig{}.Type(ctx),
			"latest_monitor_failure_message": types.StringType,
			"monitor_version":                types.Int64Type,
			"monitored_table_name":           types.StringType,
			"notification_settings":          NotificationSettings{}.Type(ctx),
			"output_schema_id":               types.StringType,
			"profile_metrics_table_name":     types.StringType,
			"schedule":                       CronSchedule{}.Type(ctx),
			"skip_builtin_dashboard":         types.BoolType,
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot":     SnapshotConfig{}.Type(ctx),
			"status":       types.StringType,
			"time_series":  TimeSeriesConfig{}.Type(ctx),
			"warehouse_id": types.StringType,
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in DataProfilingConfig as
// a slice of DataProfilingCustomMetric values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetCustomMetrics(ctx context.Context) ([]DataProfilingCustomMetric, bool) {
	if m.CustomMetrics.IsNull() || m.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []DataProfilingCustomMetric
	d := m.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in DataProfilingConfig.
func (m *DataProfilingConfig) SetCustomMetrics(ctx context.Context, v []DataProfilingCustomMetric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomMetrics = types.ListValueMust(t, vs)
}

// GetInferenceLog returns the value of the InferenceLog field in DataProfilingConfig as
// a InferenceLogConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetInferenceLog(ctx context.Context) (InferenceLogConfig, bool) {
	var e InferenceLogConfig
	if m.InferenceLog.IsNull() || m.InferenceLog.IsUnknown() {
		return e, false
	}
	var v InferenceLogConfig
	d := m.InferenceLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInferenceLog sets the value of the InferenceLog field in DataProfilingConfig.
func (m *DataProfilingConfig) SetInferenceLog(ctx context.Context, v InferenceLogConfig) {
	vs := v.ToObjectValue(ctx)
	m.InferenceLog = vs
}

// GetNotificationSettings returns the value of the NotificationSettings field in DataProfilingConfig as
// a NotificationSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetNotificationSettings(ctx context.Context) (NotificationSettings, bool) {
	var e NotificationSettings
	if m.NotificationSettings.IsNull() || m.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v NotificationSettings
	d := m.NotificationSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotificationSettings sets the value of the NotificationSettings field in DataProfilingConfig.
func (m *DataProfilingConfig) SetNotificationSettings(ctx context.Context, v NotificationSettings) {
	vs := v.ToObjectValue(ctx)
	m.NotificationSettings = vs
}

// GetSchedule returns the value of the Schedule field in DataProfilingConfig as
// a CronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetSchedule(ctx context.Context) (CronSchedule, bool) {
	var e CronSchedule
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v CronSchedule
	d := m.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in DataProfilingConfig.
func (m *DataProfilingConfig) SetSchedule(ctx context.Context, v CronSchedule) {
	vs := v.ToObjectValue(ctx)
	m.Schedule = vs
}

// GetSlicingExprs returns the value of the SlicingExprs field in DataProfilingConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
	if m.SlicingExprs.IsNull() || m.SlicingExprs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SlicingExprs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlicingExprs sets the value of the SlicingExprs field in DataProfilingConfig.
func (m *DataProfilingConfig) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in DataProfilingConfig as
// a SnapshotConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetSnapshot(ctx context.Context) (SnapshotConfig, bool) {
	var e SnapshotConfig
	if m.Snapshot.IsNull() || m.Snapshot.IsUnknown() {
		return e, false
	}
	var v SnapshotConfig
	d := m.Snapshot.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSnapshot sets the value of the Snapshot field in DataProfilingConfig.
func (m *DataProfilingConfig) SetSnapshot(ctx context.Context, v SnapshotConfig) {
	vs := v.ToObjectValue(ctx)
	m.Snapshot = vs
}

// GetTimeSeries returns the value of the TimeSeries field in DataProfilingConfig as
// a TimeSeriesConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig) GetTimeSeries(ctx context.Context) (TimeSeriesConfig, bool) {
	var e TimeSeriesConfig
	if m.TimeSeries.IsNull() || m.TimeSeries.IsUnknown() {
		return e, false
	}
	var v TimeSeriesConfig
	d := m.TimeSeries.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeSeries sets the value of the TimeSeries field in DataProfilingConfig.
func (m *DataProfilingConfig) SetTimeSeries(ctx context.Context, v TimeSeriesConfig) {
	vs := v.ToObjectValue(ctx)
	m.TimeSeries = vs
}

// Custom metric definition.
type DataProfilingCustomMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition types.String `tfsdk:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns types.List `tfsdk:"input_columns"`
	// Name of the metric in the output tables.
	Name types.String `tfsdk:"name"`
	// The output type of the custom metric.
	OutputDataType types.String `tfsdk:"output_data_type"`
	// The type of the custom metric.
	Type_ types.String `tfsdk:"type"`
}

func (to *DataProfilingCustomMetric) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataProfilingCustomMetric) {
}

func (to *DataProfilingCustomMetric) SyncFieldsDuringRead(ctx context.Context, from DataProfilingCustomMetric) {
}

func (m DataProfilingCustomMetric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["definition"] = attrs["definition"].SetRequired()
	attrs["input_columns"] = attrs["input_columns"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["output_data_type"] = attrs["output_data_type"].SetRequired()
	attrs["type"] = attrs["type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataProfilingCustomMetric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataProfilingCustomMetric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataProfilingCustomMetric
// only implements ToObjectValue() and Type().
func (m DataProfilingCustomMetric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":       m.Definition,
			"input_columns":    m.InputColumns,
			"name":             m.Name,
			"output_data_type": m.OutputDataType,
			"type":             m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataProfilingCustomMetric) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition": types.StringType,
			"input_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name":             types.StringType,
			"output_data_type": types.StringType,
			"type":             types.StringType,
		},
	}
}

// GetInputColumns returns the value of the InputColumns field in DataProfilingCustomMetric as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingCustomMetric) GetInputColumns(ctx context.Context) ([]types.String, bool) {
	if m.InputColumns.IsNull() || m.InputColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InputColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputColumns sets the value of the InputColumns field in DataProfilingCustomMetric.
func (m *DataProfilingCustomMetric) SetInputColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["input_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InputColumns = types.ListValueMust(t, vs)
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
}

func (to *DeleteMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteMonitorRequest) {
}

func (to *DeleteMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteMonitorRequest) {
}

func (m DeleteMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteMonitorRequest
// only implements ToObjectValue() and Type().
func (m DeleteMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *DeleteRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRefreshRequest) {
}

func (to *DeleteRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteRefreshRequest) {
}

func (m DeleteRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRefreshRequest
// only implements ToObjectValue() and Type().
func (m DeleteRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
}

func (to *GetMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMonitorRequest) {
}

func (to *GetMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from GetMonitorRequest) {
}

func (m GetMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMonitorRequest
// only implements ToObjectValue() and Type().
func (m GetMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *GetRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRefreshRequest) {
}

func (to *GetRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from GetRefreshRequest) {
}

func (m GetRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRefreshRequest
// only implements ToObjectValue() and Type().
func (m GetRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
}

// Inference log configuration.
type InferenceLogConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities types.List `tfsdk:"granularities"`
	// Column for the label.
	LabelColumn types.String `tfsdk:"label_column"`
	// Column for the model identifier.
	ModelIdColumn types.String `tfsdk:"model_id_column"`
	// Column for the prediction.
	PredictionColumn types.String `tfsdk:"prediction_column"`
	// Problem type the model aims to solve.
	ProblemType types.String `tfsdk:"problem_type"`
	// Column for the timestamp.
	TimestampColumn types.String `tfsdk:"timestamp_column"`
}

func (to *InferenceLogConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InferenceLogConfig) {
}

func (to *InferenceLogConfig) SyncFieldsDuringRead(ctx context.Context, from InferenceLogConfig) {
}

func (m InferenceLogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["granularities"] = attrs["granularities"].SetRequired()
	attrs["label_column"] = attrs["label_column"].SetOptional()
	attrs["model_id_column"] = attrs["model_id_column"].SetRequired()
	attrs["prediction_column"] = attrs["prediction_column"].SetRequired()
	attrs["problem_type"] = attrs["problem_type"].SetRequired()
	attrs["timestamp_column"] = attrs["timestamp_column"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InferenceLogConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InferenceLogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InferenceLogConfig
// only implements ToObjectValue() and Type().
func (m InferenceLogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities":     m.Granularities,
			"label_column":      m.LabelColumn,
			"model_id_column":   m.ModelIdColumn,
			"prediction_column": m.PredictionColumn,
			"problem_type":      m.ProblemType,
			"timestamp_column":  m.TimestampColumn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InferenceLogConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label_column":      types.StringType,
			"model_id_column":   types.StringType,
			"prediction_column": types.StringType,
			"problem_type":      types.StringType,
			"timestamp_column":  types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in InferenceLogConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InferenceLogConfig) GetGranularities(ctx context.Context) ([]types.String, bool) {
	if m.Granularities.IsNull() || m.Granularities.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Granularities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGranularities sets the value of the Granularities field in InferenceLogConfig.
func (m *InferenceLogConfig) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Granularities = types.ListValueMust(t, vs)
}

type ListMonitorRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMonitorRequest) {
}

func (to *ListMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from ListMonitorRequest) {
}

func (m ListMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMonitorRequest
// only implements ToObjectValue() and Type().
func (m ListMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response for listing Monitors.
type ListMonitorResponse struct {
	Monitors types.List `tfsdk:"monitors"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListMonitorResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMonitorResponse) {
	if !from.Monitors.IsNull() && !from.Monitors.IsUnknown() && to.Monitors.IsNull() && len(from.Monitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Monitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Monitors = from.Monitors
	}
}

func (to *ListMonitorResponse) SyncFieldsDuringRead(ctx context.Context, from ListMonitorResponse) {
	if !from.Monitors.IsNull() && !from.Monitors.IsUnknown() && to.Monitors.IsNull() && len(from.Monitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Monitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Monitors = from.Monitors
	}
}

func (m ListMonitorResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["monitors"] = attrs["monitors"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMonitorResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListMonitorResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitors": reflect.TypeOf(Monitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMonitorResponse
// only implements ToObjectValue() and Type().
func (m ListMonitorResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"monitors":        m.Monitors,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMonitorResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitors": basetypes.ListType{
				ElemType: Monitor{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMonitors returns the value of the Monitors field in ListMonitorResponse as
// a slice of Monitor values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListMonitorResponse) GetMonitors(ctx context.Context) ([]Monitor, bool) {
	if m.Monitors.IsNull() || m.Monitors.IsUnknown() {
		return nil, false
	}
	var v []Monitor
	d := m.Monitors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMonitors sets the value of the Monitors field in ListMonitorResponse.
func (m *ListMonitorResponse) SetMonitors(ctx context.Context, v []Monitor) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["monitors"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Monitors = types.ListValueMust(t, vs)
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRefreshRequest) {
}

func (to *ListRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from ListRefreshRequest) {
}

func (m ListRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshRequest
// only implements ToObjectValue() and Type().
func (m ListRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

// Response for listing refreshes.
type ListRefreshResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Refreshes types.List `tfsdk:"refreshes"`
}

func (to *ListRefreshResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRefreshResponse) {
	if !from.Refreshes.IsNull() && !from.Refreshes.IsUnknown() && to.Refreshes.IsNull() && len(from.Refreshes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Refreshes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Refreshes = from.Refreshes
	}
}

func (to *ListRefreshResponse) SyncFieldsDuringRead(ctx context.Context, from ListRefreshResponse) {
	if !from.Refreshes.IsNull() && !from.Refreshes.IsUnknown() && to.Refreshes.IsNull() && len(from.Refreshes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Refreshes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Refreshes = from.Refreshes
	}
}

func (m ListRefreshResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["refreshes"] = attrs["refreshes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRefreshResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRefreshResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refreshes": reflect.TypeOf(Refresh{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshResponse
// only implements ToObjectValue() and Type().
func (m ListRefreshResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"refreshes":       m.Refreshes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRefreshResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"refreshes": basetypes.ListType{
				ElemType: Refresh{}.Type(ctx),
			},
		},
	}
}

// GetRefreshes returns the value of the Refreshes field in ListRefreshResponse as
// a slice of Refresh values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRefreshResponse) GetRefreshes(ctx context.Context) ([]Refresh, bool) {
	if m.Refreshes.IsNull() || m.Refreshes.IsUnknown() {
		return nil, false
	}
	var v []Refresh
	d := m.Refreshes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshes sets the value of the Refreshes field in ListRefreshResponse.
func (m *ListRefreshResponse) SetRefreshes(ctx context.Context, v []Refresh) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refreshes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Refreshes = types.ListValueMust(t, vs)
}

// Monitor for the data quality of unity catalog entities such as schema or
// table.
type Monitor struct {
	// Anomaly Detection Configuration, applicable to `schema` object types.
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// Data Profiling Configuration, applicable to `table` object types. Exactly
	// one `Analysis Configuration` must be present.
	DataProfilingConfig types.Object `tfsdk:"data_profiling_config"`
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
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"object_type"`
}

func (to *Monitor) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Monitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				// Recursively sync the fields of AnomalyDetectionConfig
				toAnomalyDetectionConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
	if !from.DataProfilingConfig.IsNull() && !from.DataProfilingConfig.IsUnknown() {
		if toDataProfilingConfig, ok := to.GetDataProfilingConfig(ctx); ok {
			if fromDataProfilingConfig, ok := from.GetDataProfilingConfig(ctx); ok {
				// Recursively sync the fields of DataProfilingConfig
				toDataProfilingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromDataProfilingConfig)
				to.SetDataProfilingConfig(ctx, toDataProfilingConfig)
			}
		}
	}
}

func (to *Monitor) SyncFieldsDuringRead(ctx context.Context, from Monitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				toAnomalyDetectionConfig.SyncFieldsDuringRead(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
	if !from.DataProfilingConfig.IsNull() && !from.DataProfilingConfig.IsUnknown() {
		if toDataProfilingConfig, ok := to.GetDataProfilingConfig(ctx); ok {
			if fromDataProfilingConfig, ok := from.GetDataProfilingConfig(ctx); ok {
				toDataProfilingConfig.SyncFieldsDuringRead(ctx, fromDataProfilingConfig)
				to.SetDataProfilingConfig(ctx, toDataProfilingConfig)
			}
		}
	}
}

func (m Monitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetOptional()
	attrs["data_profiling_config"] = attrs["data_profiling_config"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Monitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Monitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(AnomalyDetectionConfig{}),
		"data_profiling_config":    reflect.TypeOf(DataProfilingConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Monitor
// only implements ToObjectValue() and Type().
func (m Monitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anomaly_detection_config": m.AnomalyDetectionConfig,
			"data_profiling_config":    m.DataProfilingConfig,
			"object_id":                m.ObjectId,
			"object_type":              m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Monitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anomaly_detection_config": AnomalyDetectionConfig{}.Type(ctx),
			"data_profiling_config":    DataProfilingConfig{}.Type(ctx),
			"object_id":                types.StringType,
			"object_type":              types.StringType,
		},
	}
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in Monitor as
// a AnomalyDetectionConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor) GetAnomalyDetectionConfig(ctx context.Context) (AnomalyDetectionConfig, bool) {
	var e AnomalyDetectionConfig
	if m.AnomalyDetectionConfig.IsNull() || m.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v AnomalyDetectionConfig
	d := m.AnomalyDetectionConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in Monitor.
func (m *Monitor) SetAnomalyDetectionConfig(ctx context.Context, v AnomalyDetectionConfig) {
	vs := v.ToObjectValue(ctx)
	m.AnomalyDetectionConfig = vs
}

// GetDataProfilingConfig returns the value of the DataProfilingConfig field in Monitor as
// a DataProfilingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor) GetDataProfilingConfig(ctx context.Context) (DataProfilingConfig, bool) {
	var e DataProfilingConfig
	if m.DataProfilingConfig.IsNull() || m.DataProfilingConfig.IsUnknown() {
		return e, false
	}
	var v DataProfilingConfig
	d := m.DataProfilingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataProfilingConfig sets the value of the DataProfilingConfig field in Monitor.
func (m *Monitor) SetDataProfilingConfig(ctx context.Context, v DataProfilingConfig) {
	vs := v.ToObjectValue(ctx)
	m.DataProfilingConfig = vs
}

// Destination of the data quality monitoring notification.
type NotificationDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses types.List `tfsdk:"email_addresses"`
}

func (to *NotificationDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationDestination) {
	if !from.EmailAddresses.IsNull() && !from.EmailAddresses.IsUnknown() && to.EmailAddresses.IsNull() && len(from.EmailAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailAddresses = from.EmailAddresses
	}
}

func (to *NotificationDestination) SyncFieldsDuringRead(ctx context.Context, from NotificationDestination) {
	if !from.EmailAddresses.IsNull() && !from.EmailAddresses.IsUnknown() && to.EmailAddresses.IsNull() && len(from.EmailAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailAddresses = from.EmailAddresses
	}
}

func (m NotificationDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_addresses"] = attrs["email_addresses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotificationDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NotificationDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination
// only implements ToObjectValue() and Type().
func (m NotificationDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_addresses": m.EmailAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotificationDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEmailAddresses returns the value of the EmailAddresses field in NotificationDestination as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationDestination) GetEmailAddresses(ctx context.Context) ([]types.String, bool) {
	if m.EmailAddresses.IsNull() || m.EmailAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EmailAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailAddresses sets the value of the EmailAddresses field in NotificationDestination.
func (m *NotificationDestination) SetEmailAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmailAddresses = types.ListValueMust(t, vs)
}

// Settings for sending notifications on the data quality monitoring.
type NotificationSettings struct {
	// Destinations to send notifications on failure/timeout.
	OnFailure types.Object `tfsdk:"on_failure"`
}

func (to *NotificationSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationSettings) {
	if !from.OnFailure.IsNull() && !from.OnFailure.IsUnknown() {
		if toOnFailure, ok := to.GetOnFailure(ctx); ok {
			if fromOnFailure, ok := from.GetOnFailure(ctx); ok {
				// Recursively sync the fields of OnFailure
				toOnFailure.SyncFieldsDuringCreateOrUpdate(ctx, fromOnFailure)
				to.SetOnFailure(ctx, toOnFailure)
			}
		}
	}
}

func (to *NotificationSettings) SyncFieldsDuringRead(ctx context.Context, from NotificationSettings) {
	if !from.OnFailure.IsNull() && !from.OnFailure.IsUnknown() {
		if toOnFailure, ok := to.GetOnFailure(ctx); ok {
			if fromOnFailure, ok := from.GetOnFailure(ctx); ok {
				toOnFailure.SyncFieldsDuringRead(ctx, fromOnFailure)
				to.SetOnFailure(ctx, toOnFailure)
			}
		}
	}
}

func (m NotificationSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["on_failure"] = attrs["on_failure"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NotificationSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_failure": reflect.TypeOf(NotificationDestination{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationSettings
// only implements ToObjectValue() and Type().
func (m NotificationSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_failure": m.OnFailure,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotificationSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_failure": NotificationDestination{}.Type(ctx),
		},
	}
}

// GetOnFailure returns the value of the OnFailure field in NotificationSettings as
// a NotificationDestination value.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationSettings) GetOnFailure(ctx context.Context) (NotificationDestination, bool) {
	var e NotificationDestination
	if m.OnFailure.IsNull() || m.OnFailure.IsUnknown() {
		return e, false
	}
	var v NotificationDestination
	d := m.OnFailure.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnFailure sets the value of the OnFailure field in NotificationSettings.
func (m *NotificationSettings) SetOnFailure(ctx context.Context, v NotificationDestination) {
	vs := v.ToObjectValue(ctx)
	m.OnFailure = vs
}

// The Refresh object gives information on a refresh of the data quality
// monitoring pipeline.
type Refresh struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message"`
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
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType types.String `tfsdk:"object_type"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id"`
	// Time when the refresh started (milliseconds since 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
	// The current state of the refresh.
	State types.String `tfsdk:"state"`
	// What triggered the refresh.
	Trigger types.String `tfsdk:"trigger"`
}

func (to *Refresh) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Refresh) {
}

func (to *Refresh) SyncFieldsDuringRead(ctx context.Context, from Refresh) {
}

func (m Refresh) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time_ms"] = attrs["end_time_ms"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetComputed()
	attrs["start_time_ms"] = attrs["start_time_ms"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["trigger"] = attrs["trigger"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Refresh.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Refresh) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Refresh
// only implements ToObjectValue() and Type().
func (m Refresh) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   m.EndTimeMs,
			"message":       m.Message,
			"object_id":     m.ObjectId,
			"object_type":   m.ObjectType,
			"refresh_id":    m.RefreshId,
			"start_time_ms": m.StartTimeMs,
			"state":         m.State,
			"trigger":       m.Trigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Refresh) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time_ms":   types.Int64Type,
			"message":       types.StringType,
			"object_id":     types.StringType,
			"object_type":   types.StringType,
			"refresh_id":    types.Int64Type,
			"start_time_ms": types.Int64Type,
			"state":         types.StringType,
			"trigger":       types.StringType,
		},
	}
}

// Snapshot analysis configuration.
type SnapshotConfig struct {
}

func (to *SnapshotConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SnapshotConfig) {
}

func (to *SnapshotConfig) SyncFieldsDuringRead(ctx context.Context, from SnapshotConfig) {
}

func (m SnapshotConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SnapshotConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SnapshotConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SnapshotConfig
// only implements ToObjectValue() and Type().
func (m SnapshotConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SnapshotConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Time series analysis configuration.
type TimeSeriesConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities types.List `tfsdk:"granularities"`
	// Column for the timestamp.
	TimestampColumn types.String `tfsdk:"timestamp_column"`
}

func (to *TimeSeriesConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeSeriesConfig) {
}

func (to *TimeSeriesConfig) SyncFieldsDuringRead(ctx context.Context, from TimeSeriesConfig) {
}

func (m TimeSeriesConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["granularities"] = attrs["granularities"].SetRequired()
	attrs["timestamp_column"] = attrs["timestamp_column"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TimeSeriesConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TimeSeriesConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeSeriesConfig
// only implements ToObjectValue() and Type().
func (m TimeSeriesConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities":    m.Granularities,
			"timestamp_column": m.TimestampColumn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TimeSeriesConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"timestamp_column": types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in TimeSeriesConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TimeSeriesConfig) GetGranularities(ctx context.Context) ([]types.String, bool) {
	if m.Granularities.IsNull() || m.Granularities.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Granularities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGranularities sets the value of the Granularities field in TimeSeriesConfig.
func (m *TimeSeriesConfig) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Granularities = types.ListValueMust(t, vs)
}

type UpdateMonitorRequest struct {
	// The monitor to update.
	Monitor types.Object `tfsdk:"monitor"`
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// The field mask to specify which fields to update as a comma-separated
	// list. Example value:
	// `data_profiling_config.custom_metrics,data_profiling_config.schedule.quartz_cron_expression`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateMonitorRequest) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				// Recursively sync the fields of Monitor
				toMonitor.SyncFieldsDuringCreateOrUpdate(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (to *UpdateMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateMonitorRequest) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				toMonitor.SyncFieldsDuringRead(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (m UpdateMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["monitor"] = attrs["monitor"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitor": reflect.TypeOf(Monitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMonitorRequest
// only implements ToObjectValue() and Type().
func (m UpdateMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"monitor":     m.Monitor,
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitor":     Monitor{}.Type(ctx),
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetMonitor returns the value of the Monitor field in UpdateMonitorRequest as
// a Monitor value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateMonitorRequest) GetMonitor(ctx context.Context) (Monitor, bool) {
	var e Monitor
	if m.Monitor.IsNull() || m.Monitor.IsUnknown() {
		return e, false
	}
	var v Monitor
	d := m.Monitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMonitor sets the value of the Monitor field in UpdateMonitorRequest.
func (m *UpdateMonitorRequest) SetMonitor(ctx context.Context, v Monitor) {
	vs := v.ToObjectValue(ctx)
	m.Monitor = vs
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
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// The refresh to update.
	Refresh types.Object `tfsdk:"refresh"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
	// The field mask to specify which fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateRefreshRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRefreshRequest) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				// Recursively sync the fields of Refresh
				toRefresh.SyncFieldsDuringCreateOrUpdate(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (to *UpdateRefreshRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateRefreshRequest) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m UpdateRefreshRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRefreshRequest
// only implements ToObjectValue() and Type().
func (m UpdateRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh":     m.Refresh,
			"refresh_id":  m.RefreshId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh":     Refresh{}.Type(ctx),
			"refresh_id":  types.Int64Type,
			"update_mask": types.StringType,
		},
	}
}

// GetRefresh returns the value of the Refresh field in UpdateRefreshRequest as
// a Refresh value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRefreshRequest) GetRefresh(ctx context.Context) (Refresh, bool) {
	var e Refresh
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v Refresh
	d := m.Refresh.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefresh sets the value of the Refresh field in UpdateRefreshRequest.
func (m *UpdateRefreshRequest) SetRefresh(ctx context.Context, v Refresh) {
	vs := v.ToObjectValue(ctx)
	m.Refresh = vs
}
