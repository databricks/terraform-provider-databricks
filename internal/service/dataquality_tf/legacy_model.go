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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Anomaly Detection Configurations.
type AnomalyDetectionConfig_SdkV2 struct {
}

func (to *AnomalyDetectionConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AnomalyDetectionConfig_SdkV2) {
}

func (to *AnomalyDetectionConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AnomalyDetectionConfig_SdkV2) {
}

func (m AnomalyDetectionConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AnomalyDetectionConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AnomalyDetectionConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnomalyDetectionConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AnomalyDetectionConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m AnomalyDetectionConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Request to cancel a refresh.
type CancelRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *CancelRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelRefreshRequest_SdkV2) {
}

func (to *CancelRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CancelRefreshRequest_SdkV2) {
}

func (m CancelRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CancelRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CancelRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
}

// Response to cancelling a refresh.
type CancelRefreshResponse_SdkV2 struct {
	// The refresh to cancel.
	Refresh types.List `tfsdk:"refresh"`
}

func (to *CancelRefreshResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelRefreshResponse_SdkV2) {
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

func (to *CancelRefreshResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CancelRefreshResponse_SdkV2) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m CancelRefreshResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetComputed()
	attrs["refresh"] = attrs["refresh"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelRefreshResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CancelRefreshResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh": m.Refresh,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelRefreshResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh": basetypes.ListType{
				ElemType: Refresh_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRefresh returns the value of the Refresh field in CancelRefreshResponse_SdkV2 as
// a Refresh_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CancelRefreshResponse_SdkV2) GetRefresh(ctx context.Context) (Refresh_SdkV2, bool) {
	var e Refresh_SdkV2
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v []Refresh_SdkV2
	d := m.Refresh.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRefresh sets the value of the Refresh field in CancelRefreshResponse_SdkV2.
func (m *CancelRefreshResponse_SdkV2) SetRefresh(ctx context.Context, v Refresh_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh"]
	m.Refresh = types.ListValueMust(t, vs)
}

type CreateMonitorRequest_SdkV2 struct {
	// The monitor to create.
	Monitor types.List `tfsdk:"monitor"`
}

func (to *CreateMonitorRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateMonitorRequest_SdkV2) {
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

func (to *CreateMonitorRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateMonitorRequest_SdkV2) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				toMonitor.SyncFieldsDuringRead(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (m CreateMonitorRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["monitor"] = attrs["monitor"].SetRequired()
	attrs["monitor"] = attrs["monitor"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitor": reflect.TypeOf(Monitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"monitor": m.Monitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitor": basetypes.ListType{
				ElemType: Monitor_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMonitor returns the value of the Monitor field in CreateMonitorRequest_SdkV2 as
// a Monitor_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateMonitorRequest_SdkV2) GetMonitor(ctx context.Context) (Monitor_SdkV2, bool) {
	var e Monitor_SdkV2
	if m.Monitor.IsNull() || m.Monitor.IsUnknown() {
		return e, false
	}
	var v []Monitor_SdkV2
	d := m.Monitor.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMonitor sets the value of the Monitor field in CreateMonitorRequest_SdkV2.
func (m *CreateMonitorRequest_SdkV2) SetMonitor(ctx context.Context, v Monitor_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["monitor"]
	m.Monitor = types.ListValueMust(t, vs)
}

type CreateRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, table id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType types.String `tfsdk:"-"`
	// The refresh to create
	Refresh types.List `tfsdk:"refresh"`
}

func (to *CreateRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRefreshRequest_SdkV2) {
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

func (to *CreateRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRefreshRequest_SdkV2) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m CreateRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetRequired()
	attrs["refresh"] = attrs["refresh"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh":     m.Refresh,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh": basetypes.ListType{
				ElemType: Refresh_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRefresh returns the value of the Refresh field in CreateRefreshRequest_SdkV2 as
// a Refresh_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRefreshRequest_SdkV2) GetRefresh(ctx context.Context) (Refresh_SdkV2, bool) {
	var e Refresh_SdkV2
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v []Refresh_SdkV2
	d := m.Refresh.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRefresh sets the value of the Refresh field in CreateRefreshRequest_SdkV2.
func (m *CreateRefreshRequest_SdkV2) SetRefresh(ctx context.Context, v Refresh_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh"]
	m.Refresh = types.ListValueMust(t, vs)
}

// The data quality monitoring workflow cron schedule.
type CronSchedule_SdkV2 struct {
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

func (to *CronSchedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronSchedule_SdkV2) {
}

func (to *CronSchedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CronSchedule_SdkV2) {
}

func (m CronSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CronSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (m CronSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status":           m.PauseStatus,
			"quartz_cron_expression": m.QuartzCronExpression,
			"timezone_id":            m.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

// Data Profiling Configurations.
type DataProfilingConfig_SdkV2 struct {
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
	// Configuration for monitoring inference log tables.
	InferenceLog types.List `tfsdk:"inference_log"`
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
	NotificationSettings types.List `tfsdk:"notification_settings"`
	// ID of the schema where output tables are created.
	OutputSchemaId types.String `tfsdk:"output_schema_id"`
	// Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	ProfileMetricsTableName types.String `tfsdk:"profile_metrics_table_name"`
	// The cron schedule.
	Schedule types.List `tfsdk:"schedule"`
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
	// Configuration for monitoring snapshot tables.
	Snapshot types.List `tfsdk:"snapshot"`
	// The data profiling monitor status.
	Status types.String `tfsdk:"status"`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *DataProfilingConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataProfilingConfig_SdkV2) {
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

func (to *DataProfilingConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataProfilingConfig_SdkV2) {
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

func (m DataProfilingConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets_dir"] = attrs["assets_dir"].SetOptional()
	attrs["assets_dir"] = attrs["assets_dir"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["baseline_table_name"] = attrs["baseline_table_name"].SetOptional()
	attrs["custom_metrics"] = attrs["custom_metrics"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["drift_metrics_table_name"] = attrs["drift_metrics_table_name"].SetComputed()
	attrs["effective_warehouse_id"] = attrs["effective_warehouse_id"].SetComputed()
	attrs["inference_log"] = attrs["inference_log"].SetOptional()
	attrs["inference_log"] = attrs["inference_log"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["latest_monitor_failure_message"] = attrs["latest_monitor_failure_message"].SetComputed()
	attrs["monitor_version"] = attrs["monitor_version"].SetComputed()
	attrs["monitored_table_name"] = attrs["monitored_table_name"].SetComputed()
	attrs["notification_settings"] = attrs["notification_settings"].SetOptional()
	attrs["notification_settings"] = attrs["notification_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["output_schema_id"] = attrs["output_schema_id"].SetRequired()
	attrs["profile_metrics_table_name"] = attrs["profile_metrics_table_name"].SetComputed()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["schedule"] = attrs["schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["skip_builtin_dashboard"] = attrs["skip_builtin_dashboard"].SetOptional()
	attrs["slicing_exprs"] = attrs["slicing_exprs"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["time_series"] = attrs["time_series"].SetOptional()
	attrs["time_series"] = attrs["time_series"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DataProfilingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":        reflect.TypeOf(DataProfilingCustomMetric_SdkV2{}),
		"inference_log":         reflect.TypeOf(InferenceLogConfig_SdkV2{}),
		"notification_settings": reflect.TypeOf(NotificationSettings_SdkV2{}),
		"schedule":              reflect.TypeOf(CronSchedule_SdkV2{}),
		"slicing_exprs":         reflect.TypeOf(types.String{}),
		"snapshot":              reflect.TypeOf(SnapshotConfig_SdkV2{}),
		"time_series":           reflect.TypeOf(TimeSeriesConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataProfilingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m DataProfilingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DataProfilingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: DataProfilingCustomMetric_SdkV2{}.Type(ctx),
			},
			"dashboard_id":             types.StringType,
			"drift_metrics_table_name": types.StringType,
			"effective_warehouse_id":   types.StringType,
			"inference_log": basetypes.ListType{
				ElemType: InferenceLogConfig_SdkV2{}.Type(ctx),
			},
			"latest_monitor_failure_message": types.StringType,
			"monitor_version":                types.Int64Type,
			"monitored_table_name":           types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: NotificationSettings_SdkV2{}.Type(ctx),
			},
			"output_schema_id":           types.StringType,
			"profile_metrics_table_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule_SdkV2{}.Type(ctx),
			},
			"skip_builtin_dashboard": types.BoolType,
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: SnapshotConfig_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: TimeSeriesConfig_SdkV2{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in DataProfilingConfig_SdkV2 as
// a slice of DataProfilingCustomMetric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetCustomMetrics(ctx context.Context) ([]DataProfilingCustomMetric_SdkV2, bool) {
	if m.CustomMetrics.IsNull() || m.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []DataProfilingCustomMetric_SdkV2
	d := m.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetCustomMetrics(ctx context.Context, v []DataProfilingCustomMetric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomMetrics = types.ListValueMust(t, vs)
}

// GetInferenceLog returns the value of the InferenceLog field in DataProfilingConfig_SdkV2 as
// a InferenceLogConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetInferenceLog(ctx context.Context) (InferenceLogConfig_SdkV2, bool) {
	var e InferenceLogConfig_SdkV2
	if m.InferenceLog.IsNull() || m.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []InferenceLogConfig_SdkV2
	d := m.InferenceLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceLog sets the value of the InferenceLog field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetInferenceLog(ctx context.Context, v InferenceLogConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_log"]
	m.InferenceLog = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in DataProfilingConfig_SdkV2 as
// a NotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetNotificationSettings(ctx context.Context) (NotificationSettings_SdkV2, bool) {
	var e NotificationSettings_SdkV2
	if m.NotificationSettings.IsNull() || m.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []NotificationSettings_SdkV2
	d := m.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetNotificationSettings(ctx context.Context, v NotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	m.NotificationSettings = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in DataProfilingConfig_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := m.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	m.Schedule = types.ListValueMust(t, vs)
}

// GetSlicingExprs returns the value of the SlicingExprs field in DataProfilingConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
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

// SetSlicingExprs sets the value of the SlicingExprs field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in DataProfilingConfig_SdkV2 as
// a SnapshotConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetSnapshot(ctx context.Context) (SnapshotConfig_SdkV2, bool) {
	var e SnapshotConfig_SdkV2
	if m.Snapshot.IsNull() || m.Snapshot.IsUnknown() {
		return e, false
	}
	var v []SnapshotConfig_SdkV2
	d := m.Snapshot.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSnapshot sets the value of the Snapshot field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetSnapshot(ctx context.Context, v SnapshotConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["snapshot"]
	m.Snapshot = types.ListValueMust(t, vs)
}

// GetTimeSeries returns the value of the TimeSeries field in DataProfilingConfig_SdkV2 as
// a TimeSeriesConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingConfig_SdkV2) GetTimeSeries(ctx context.Context) (TimeSeriesConfig_SdkV2, bool) {
	var e TimeSeriesConfig_SdkV2
	if m.TimeSeries.IsNull() || m.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []TimeSeriesConfig_SdkV2
	d := m.TimeSeries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimeSeries sets the value of the TimeSeries field in DataProfilingConfig_SdkV2.
func (m *DataProfilingConfig_SdkV2) SetTimeSeries(ctx context.Context, v TimeSeriesConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["time_series"]
	m.TimeSeries = types.ListValueMust(t, vs)
}

// Custom metric definition.
type DataProfilingCustomMetric_SdkV2 struct {
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

func (to *DataProfilingCustomMetric_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataProfilingCustomMetric_SdkV2) {
}

func (to *DataProfilingCustomMetric_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataProfilingCustomMetric_SdkV2) {
}

func (m DataProfilingCustomMetric_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataProfilingCustomMetric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataProfilingCustomMetric_SdkV2
// only implements ToObjectValue() and Type().
func (m DataProfilingCustomMetric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DataProfilingCustomMetric_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInputColumns returns the value of the InputColumns field in DataProfilingCustomMetric_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataProfilingCustomMetric_SdkV2) GetInputColumns(ctx context.Context) ([]types.String, bool) {
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

// SetInputColumns sets the value of the InputColumns field in DataProfilingCustomMetric_SdkV2.
func (m *DataProfilingCustomMetric_SdkV2) SetInputColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["input_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InputColumns = types.ListValueMust(t, vs)
}

type DeleteMonitorRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
}

func (to *DeleteMonitorRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteMonitorRequest_SdkV2) {
}

func (to *DeleteMonitorRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteMonitorRequest_SdkV2) {
}

func (m DeleteMonitorRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type DeleteRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *DeleteRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRefreshRequest_SdkV2) {
}

func (to *DeleteRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRefreshRequest_SdkV2) {
}

func (m DeleteRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
}

type GetMonitorRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
}

func (to *GetMonitorRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMonitorRequest_SdkV2) {
}

func (to *GetMonitorRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetMonitorRequest_SdkV2) {
}

func (m GetMonitorRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type GetRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
}

func (to *GetRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRefreshRequest_SdkV2) {
}

func (to *GetRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRefreshRequest_SdkV2) {
}

func (m GetRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"refresh_id":  m.RefreshId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh_id":  types.Int64Type,
		},
	}
}

// Inference log configuration.
type InferenceLogConfig_SdkV2 struct {
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

func (to *InferenceLogConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InferenceLogConfig_SdkV2) {
}

func (to *InferenceLogConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from InferenceLogConfig_SdkV2) {
}

func (m InferenceLogConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m InferenceLogConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InferenceLogConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m InferenceLogConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m InferenceLogConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetGranularities returns the value of the Granularities field in InferenceLogConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *InferenceLogConfig_SdkV2) GetGranularities(ctx context.Context) ([]types.String, bool) {
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

// SetGranularities sets the value of the Granularities field in InferenceLogConfig_SdkV2.
func (m *InferenceLogConfig_SdkV2) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Granularities = types.ListValueMust(t, vs)
}

type ListMonitorRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListMonitorRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMonitorRequest_SdkV2) {
}

func (to *ListMonitorRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListMonitorRequest_SdkV2) {
}

func (m ListMonitorRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response for listing Monitors.
type ListMonitorResponse_SdkV2 struct {
	Monitors types.List `tfsdk:"monitors"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListMonitorResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMonitorResponse_SdkV2) {
	if !from.Monitors.IsNull() && !from.Monitors.IsUnknown() && to.Monitors.IsNull() && len(from.Monitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Monitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Monitors = from.Monitors
	}
}

func (to *ListMonitorResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListMonitorResponse_SdkV2) {
	if !from.Monitors.IsNull() && !from.Monitors.IsUnknown() && to.Monitors.IsNull() && len(from.Monitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Monitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Monitors = from.Monitors
	}
}

func (m ListMonitorResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListMonitorResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitors": reflect.TypeOf(Monitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMonitorResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListMonitorResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"monitors":        m.Monitors,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMonitorResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitors": basetypes.ListType{
				ElemType: Monitor_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMonitors returns the value of the Monitors field in ListMonitorResponse_SdkV2 as
// a slice of Monitor_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListMonitorResponse_SdkV2) GetMonitors(ctx context.Context) ([]Monitor_SdkV2, bool) {
	if m.Monitors.IsNull() || m.Monitors.IsUnknown() {
		return nil, false
	}
	var v []Monitor_SdkV2
	d := m.Monitors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMonitors sets the value of the Monitors field in ListMonitorResponse_SdkV2.
func (m *ListMonitorResponse_SdkV2) SetMonitors(ctx context.Context, v []Monitor_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["monitors"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Monitors = types.ListValueMust(t, vs)
}

type ListRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRefreshRequest_SdkV2) {
}

func (to *ListRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRefreshRequest_SdkV2) {
}

func (m ListRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type ListRefreshResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Refreshes types.List `tfsdk:"refreshes"`
}

func (to *ListRefreshResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRefreshResponse_SdkV2) {
	if !from.Refreshes.IsNull() && !from.Refreshes.IsUnknown() && to.Refreshes.IsNull() && len(from.Refreshes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Refreshes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Refreshes = from.Refreshes
	}
}

func (to *ListRefreshResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRefreshResponse_SdkV2) {
	if !from.Refreshes.IsNull() && !from.Refreshes.IsUnknown() && to.Refreshes.IsNull() && len(from.Refreshes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Refreshes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Refreshes = from.Refreshes
	}
}

func (m ListRefreshResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListRefreshResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refreshes": reflect.TypeOf(Refresh_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRefreshResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"refreshes":       m.Refreshes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRefreshResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"refreshes": basetypes.ListType{
				ElemType: Refresh_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRefreshes returns the value of the Refreshes field in ListRefreshResponse_SdkV2 as
// a slice of Refresh_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRefreshResponse_SdkV2) GetRefreshes(ctx context.Context) ([]Refresh_SdkV2, bool) {
	if m.Refreshes.IsNull() || m.Refreshes.IsUnknown() {
		return nil, false
	}
	var v []Refresh_SdkV2
	d := m.Refreshes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshes sets the value of the Refreshes field in ListRefreshResponse_SdkV2.
func (m *ListRefreshResponse_SdkV2) SetRefreshes(ctx context.Context, v []Refresh_SdkV2) {
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
type Monitor_SdkV2 struct {
	// Anomaly Detection Configuration, applicable to `schema` object types.
	AnomalyDetectionConfig types.List `tfsdk:"anomaly_detection_config"`
	// Data Profiling Configuration, applicable to `table` object types
	DataProfilingConfig types.List `tfsdk:"data_profiling_config"`
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"object_type"`
}

func (to *Monitor_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Monitor_SdkV2) {
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

func (to *Monitor_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Monitor_SdkV2) {
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

func (m Monitor_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetOptional()
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["data_profiling_config"] = attrs["data_profiling_config"].SetOptional()
	attrs["data_profiling_config"] = attrs["data_profiling_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m Monitor_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(AnomalyDetectionConfig_SdkV2{}),
		"data_profiling_config":    reflect.TypeOf(DataProfilingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Monitor_SdkV2
// only implements ToObjectValue() and Type().
func (m Monitor_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Monitor_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anomaly_detection_config": basetypes.ListType{
				ElemType: AnomalyDetectionConfig_SdkV2{}.Type(ctx),
			},
			"data_profiling_config": basetypes.ListType{
				ElemType: DataProfilingConfig_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in Monitor_SdkV2 as
// a AnomalyDetectionConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor_SdkV2) GetAnomalyDetectionConfig(ctx context.Context) (AnomalyDetectionConfig_SdkV2, bool) {
	var e AnomalyDetectionConfig_SdkV2
	if m.AnomalyDetectionConfig.IsNull() || m.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v []AnomalyDetectionConfig_SdkV2
	d := m.AnomalyDetectionConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in Monitor_SdkV2.
func (m *Monitor_SdkV2) SetAnomalyDetectionConfig(ctx context.Context, v AnomalyDetectionConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["anomaly_detection_config"]
	m.AnomalyDetectionConfig = types.ListValueMust(t, vs)
}

// GetDataProfilingConfig returns the value of the DataProfilingConfig field in Monitor_SdkV2 as
// a DataProfilingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor_SdkV2) GetDataProfilingConfig(ctx context.Context) (DataProfilingConfig_SdkV2, bool) {
	var e DataProfilingConfig_SdkV2
	if m.DataProfilingConfig.IsNull() || m.DataProfilingConfig.IsUnknown() {
		return e, false
	}
	var v []DataProfilingConfig_SdkV2
	d := m.DataProfilingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataProfilingConfig sets the value of the DataProfilingConfig field in Monitor_SdkV2.
func (m *Monitor_SdkV2) SetDataProfilingConfig(ctx context.Context, v DataProfilingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_profiling_config"]
	m.DataProfilingConfig = types.ListValueMust(t, vs)
}

// Destination of the data quality monitoring notification.
type NotificationDestination_SdkV2 struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses types.List `tfsdk:"email_addresses"`
}

func (to *NotificationDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationDestination_SdkV2) {
	if !from.EmailAddresses.IsNull() && !from.EmailAddresses.IsUnknown() && to.EmailAddresses.IsNull() && len(from.EmailAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailAddresses = from.EmailAddresses
	}
}

func (to *NotificationDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotificationDestination_SdkV2) {
	if !from.EmailAddresses.IsNull() && !from.EmailAddresses.IsUnknown() && to.EmailAddresses.IsNull() && len(from.EmailAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailAddresses = from.EmailAddresses
	}
}

func (m NotificationDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NotificationDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m NotificationDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_addresses": m.EmailAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotificationDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEmailAddresses returns the value of the EmailAddresses field in NotificationDestination_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationDestination_SdkV2) GetEmailAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetEmailAddresses sets the value of the EmailAddresses field in NotificationDestination_SdkV2.
func (m *NotificationDestination_SdkV2) SetEmailAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmailAddresses = types.ListValueMust(t, vs)
}

// Settings for sending notifications on the data quality monitoring.
type NotificationSettings_SdkV2 struct {
	// Destinations to send notifications on failure/timeout.
	OnFailure types.List `tfsdk:"on_failure"`
}

func (to *NotificationSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationSettings_SdkV2) {
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

func (to *NotificationSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotificationSettings_SdkV2) {
	if !from.OnFailure.IsNull() && !from.OnFailure.IsUnknown() {
		if toOnFailure, ok := to.GetOnFailure(ctx); ok {
			if fromOnFailure, ok := from.GetOnFailure(ctx); ok {
				toOnFailure.SyncFieldsDuringRead(ctx, fromOnFailure)
				to.SetOnFailure(ctx, toOnFailure)
			}
		}
	}
}

func (m NotificationSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["on_failure"] = attrs["on_failure"].SetOptional()
	attrs["on_failure"] = attrs["on_failure"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NotificationSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_failure": reflect.TypeOf(NotificationDestination_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m NotificationSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_failure": m.OnFailure,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotificationSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_failure": basetypes.ListType{
				ElemType: NotificationDestination_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOnFailure returns the value of the OnFailure field in NotificationSettings_SdkV2 as
// a NotificationDestination_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationSettings_SdkV2) GetOnFailure(ctx context.Context) (NotificationDestination_SdkV2, bool) {
	var e NotificationDestination_SdkV2
	if m.OnFailure.IsNull() || m.OnFailure.IsUnknown() {
		return e, false
	}
	var v []NotificationDestination_SdkV2
	d := m.OnFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnFailure sets the value of the OnFailure field in NotificationSettings_SdkV2.
func (m *NotificationSettings_SdkV2) SetOnFailure(ctx context.Context, v NotificationDestination_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["on_failure"]
	m.OnFailure = types.ListValueMust(t, vs)
}

// The Refresh object gives information on a refresh of the data quality
// monitoring pipeline.
type Refresh_SdkV2 struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message"`
	// The UUID of the request object. For example, table id.
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

func (to *Refresh_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Refresh_SdkV2) {
}

func (to *Refresh_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Refresh_SdkV2) {
}

func (m Refresh_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Refresh_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Refresh_SdkV2
// only implements ToObjectValue() and Type().
func (m Refresh_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Refresh_SdkV2) Type(ctx context.Context) attr.Type {
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
type SnapshotConfig_SdkV2 struct {
}

func (to *SnapshotConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SnapshotConfig_SdkV2) {
}

func (to *SnapshotConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SnapshotConfig_SdkV2) {
}

func (m SnapshotConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SnapshotConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SnapshotConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SnapshotConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m SnapshotConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SnapshotConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Time series analysis configuration.
type TimeSeriesConfig_SdkV2 struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities types.List `tfsdk:"granularities"`
	// Column for the timestamp.
	TimestampColumn types.String `tfsdk:"timestamp_column"`
}

func (to *TimeSeriesConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeSeriesConfig_SdkV2) {
}

func (to *TimeSeriesConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TimeSeriesConfig_SdkV2) {
}

func (m TimeSeriesConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TimeSeriesConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeSeriesConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m TimeSeriesConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities":    m.Granularities,
			"timestamp_column": m.TimestampColumn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TimeSeriesConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"timestamp_column": types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in TimeSeriesConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TimeSeriesConfig_SdkV2) GetGranularities(ctx context.Context) ([]types.String, bool) {
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

// SetGranularities sets the value of the Granularities field in TimeSeriesConfig_SdkV2.
func (m *TimeSeriesConfig_SdkV2) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Granularities = types.ListValueMust(t, vs)
}

type UpdateMonitorRequest_SdkV2 struct {
	// The monitor to update.
	Monitor types.List `tfsdk:"monitor"`
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// The field mask to specify which fields to update as a comma-separated
	// list. Example value:
	// `data_profiling_config.custom_metrics,data_profiling_config.schedule.quartz_cron_expression`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateMonitorRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateMonitorRequest_SdkV2) {
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

func (to *UpdateMonitorRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateMonitorRequest_SdkV2) {
	if !from.Monitor.IsNull() && !from.Monitor.IsUnknown() {
		if toMonitor, ok := to.GetMonitor(ctx); ok {
			if fromMonitor, ok := from.GetMonitor(ctx); ok {
				toMonitor.SyncFieldsDuringRead(ctx, fromMonitor)
				to.SetMonitor(ctx, toMonitor)
			}
		}
	}
}

func (m UpdateMonitorRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["monitor"] = attrs["monitor"].SetRequired()
	attrs["monitor"] = attrs["monitor"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitor": reflect.TypeOf(Monitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"monitor": basetypes.ListType{
				ElemType: Monitor_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetMonitor returns the value of the Monitor field in UpdateMonitorRequest_SdkV2 as
// a Monitor_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateMonitorRequest_SdkV2) GetMonitor(ctx context.Context) (Monitor_SdkV2, bool) {
	var e Monitor_SdkV2
	if m.Monitor.IsNull() || m.Monitor.IsUnknown() {
		return e, false
	}
	var v []Monitor_SdkV2
	d := m.Monitor.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMonitor sets the value of the Monitor field in UpdateMonitorRequest_SdkV2.
func (m *UpdateMonitorRequest_SdkV2) SetMonitor(ctx context.Context, v Monitor_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["monitor"]
	m.Monitor = types.ListValueMust(t, vs)
}

type UpdateRefreshRequest_SdkV2 struct {
	// The UUID of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"-"`
	// The refresh to update.
	Refresh types.List `tfsdk:"refresh"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"-"`
	// The field mask to specify which fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateRefreshRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRefreshRequest_SdkV2) {
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

func (to *UpdateRefreshRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRefreshRequest_SdkV2) {
	if !from.Refresh.IsNull() && !from.Refresh.IsUnknown() {
		if toRefresh, ok := to.GetRefresh(ctx); ok {
			if fromRefresh, ok := from.GetRefresh(ctx); ok {
				toRefresh.SyncFieldsDuringRead(ctx, fromRefresh)
				to.SetRefresh(ctx, toRefresh)
			}
		}
	}
}

func (m UpdateRefreshRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refresh"] = attrs["refresh"].SetRequired()
	attrs["refresh"] = attrs["refresh"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refresh": reflect.TypeOf(Refresh_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"refresh": basetypes.ListType{
				ElemType: Refresh_SdkV2{}.Type(ctx),
			},
			"refresh_id":  types.Int64Type,
			"update_mask": types.StringType,
		},
	}
}

// GetRefresh returns the value of the Refresh field in UpdateRefreshRequest_SdkV2 as
// a Refresh_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRefreshRequest_SdkV2) GetRefresh(ctx context.Context) (Refresh_SdkV2, bool) {
	var e Refresh_SdkV2
	if m.Refresh.IsNull() || m.Refresh.IsUnknown() {
		return e, false
	}
	var v []Refresh_SdkV2
	d := m.Refresh.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRefresh sets the value of the Refresh field in UpdateRefreshRequest_SdkV2.
func (m *UpdateRefreshRequest_SdkV2) SetRefresh(ctx context.Context, v Refresh_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh"]
	m.Refresh = types.ListValueMust(t, vs)
}
