// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package dashboards_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Create dashboard
type CreateDashboardRequest_SdkV2 struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
}

func (newState *CreateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDashboardRequest_SdkV2) {
}

func (newState *CreateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard": o.Dashboard,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
		},
	}
}

// GetDashboard returns the value of the Dashboard field in CreateDashboardRequest_SdkV2 as
// a Dashboard_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDashboardRequest_SdkV2) GetDashboard(ctx context.Context) (Dashboard_SdkV2, bool) {
	var e Dashboard_SdkV2
	if o.Dashboard.IsNull() || o.Dashboard.IsUnknown() {
		return e, false
	}
	var v []Dashboard_SdkV2
	d := o.Dashboard.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboard sets the value of the Dashboard field in CreateDashboardRequest_SdkV2.
func (o *CreateDashboardRequest_SdkV2) SetDashboard(ctx context.Context, v Dashboard_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard"]
	o.Dashboard = types.ListValueMust(t, vs)
}

// Create dashboard schedule
type CreateScheduleRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
}

func (newState *CreateScheduleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScheduleRequest_SdkV2) {
}

func (newState *CreateScheduleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateScheduleRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateScheduleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScheduleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateScheduleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule":     o.Schedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateScheduleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
		},
	}
}

// GetSchedule returns the value of the Schedule field in CreateScheduleRequest_SdkV2 as
// a Schedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateScheduleRequest_SdkV2) GetSchedule(ctx context.Context) (Schedule_SdkV2, bool) {
	var e Schedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []Schedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in CreateScheduleRequest_SdkV2.
func (o *CreateScheduleRequest_SdkV2) SetSchedule(ctx context.Context, v Schedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// Create schedule subscription
type CreateSubscriptionRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`

	Subscription types.List `tfsdk:"subscription" tf:"optional,object"`
}

func (newState *CreateSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSubscriptionRequest_SdkV2) {
}

func (newState *CreateSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateSubscriptionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSubscriptionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscription": reflect.TypeOf(Subscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSubscriptionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateSubscriptionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule_id":  o.ScheduleId,
			"subscription": o.Subscription,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
			"subscription": basetypes.ListType{
				ElemType: Subscription{}.Type(ctx),
			},
		},
	}
}

// GetSubscription returns the value of the Subscription field in CreateSubscriptionRequest_SdkV2 as
// a Subscription_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSubscriptionRequest_SdkV2) GetSubscription(ctx context.Context) (Subscription_SdkV2, bool) {
	var e Subscription_SdkV2
	if o.Subscription.IsNull() || o.Subscription.IsUnknown() {
		return e, false
	}
	var v []Subscription_SdkV2
	d := o.Subscription.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSubscription sets the value of the Subscription field in CreateSubscriptionRequest_SdkV2.
func (o *CreateSubscriptionRequest_SdkV2) SetSubscription(ctx context.Context, v Subscription_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscription"]
	o.Subscription = types.ListValueMust(t, vs)
}

type CronSchedule_SdkV2 struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression" tf:""`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id" tf:""`
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule_SdkV2) {
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState CronSchedule_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (o CronSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_expression": o.QuartzCronExpression,
			"timezone_id":            o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type Dashboard_SdkV2 struct {
	// The timestamp of when the dashboard was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag types.String `tfsdk:"etag" tf:"computed"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"computed"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath types.String `tfsdk:"parent_path" tf:"computed"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path types.String `tfsdk:"path" tf:"computed"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard_SdkV2) {
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState Dashboard_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dashboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dashboard_SdkV2
// only implements ToObjectValue() and Type().
func (o Dashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":          o.CreateTime,
			"dashboard_id":         o.DashboardId,
			"display_name":         o.DisplayName,
			"etag":                 o.Etag,
			"lifecycle_state":      o.LifecycleState,
			"parent_path":          o.ParentPath,
			"path":                 o.Path,
			"serialized_dashboard": o.SerializedDashboard,
			"update_time":          o.UpdateTime,
			"warehouse_id":         o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dashboard_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":          types.StringType,
			"dashboard_id":         types.StringType,
			"display_name":         types.StringType,
			"etag":                 types.StringType,
			"lifecycle_state":      types.StringType,
			"parent_path":          types.StringType,
			"path":                 types.StringType,
			"serialized_dashboard": types.StringType,
			"update_time":          types.StringType,
			"warehouse_id":         types.StringType,
		},
	}
}

// Delete dashboard schedule
type DeleteScheduleRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *DeleteScheduleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleRequest_SdkV2) {
}

func (newState *DeleteScheduleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScheduleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteScheduleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"etag":         o.Etag,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"etag":         types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type DeleteScheduleResponse_SdkV2 struct {
}

func (newState *DeleteScheduleResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleResponse_SdkV2) {
}

func (newState *DeleteScheduleResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScheduleResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteScheduleResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete schedule subscription
type DeleteSubscriptionRequest_SdkV2 struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-"`
}

func (newState *DeleteSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionRequest_SdkV2) {
}

func (newState *DeleteSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSubscriptionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSubscriptionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":    o.DashboardId,
			"etag":            o.Etag,
			"schedule_id":     o.ScheduleId,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"etag":            types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type DeleteSubscriptionResponse_SdkV2 struct {
}

func (newState *DeleteSubscriptionResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionResponse_SdkV2) {
}

func (newState *DeleteSubscriptionResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSubscriptionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSubscriptionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Genie AI Response
type GenieAttachment_SdkV2 struct {
	Query types.List `tfsdk:"query" tf:"optional,object"`

	Text types.List `tfsdk:"text" tf:"optional,object"`
}

func (newState *GenieAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieAttachment_SdkV2) {
}

func (newState *GenieAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieAttachment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieAttachment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(QueryAttachment_SdkV2{}),
		"text":  reflect.TypeOf(TextAttachment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAttachment_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieAttachment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query": o.Query,
			"text":  o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieAttachment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query": basetypes.ListType{
				ElemType: QueryAttachment{}.Type(ctx),
			},
			"text": basetypes.ListType{
				ElemType: TextAttachment{}.Type(ctx),
			},
		},
	}
}

// GetQuery returns the value of the Query field in GenieAttachment_SdkV2 as
// a QueryAttachment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment_SdkV2) GetQuery(ctx context.Context) (QueryAttachment_SdkV2, bool) {
	var e QueryAttachment_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []QueryAttachment_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in GenieAttachment_SdkV2.
func (o *GenieAttachment_SdkV2) SetQuery(ctx context.Context, v QueryAttachment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

// GetText returns the value of the Text field in GenieAttachment_SdkV2 as
// a TextAttachment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment_SdkV2) GetText(ctx context.Context) (TextAttachment_SdkV2, bool) {
	var e TextAttachment_SdkV2
	if o.Text.IsNull() || o.Text.IsUnknown() {
		return e, false
	}
	var v []TextAttachment_SdkV2
	d := o.Text.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetText sets the value of the Text field in GenieAttachment_SdkV2.
func (o *GenieAttachment_SdkV2) SetText(ctx context.Context, v TextAttachment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["text"]
	o.Text = types.ListValueMust(t, vs)
}

type GenieConversation_SdkV2 struct {
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp" tf:"optional"`
	// Conversation ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id" tf:""`
	// Conversation title
	Title types.String `tfsdk:"title" tf:""`
	// ID of the user who created the conversation
	UserId types.Int64 `tfsdk:"user_id" tf:""`
}

func (newState *GenieConversation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieConversation_SdkV2) {
}

func (newState *GenieConversation_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieConversation_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieConversation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieConversation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieConversation_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieConversation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_timestamp":      o.CreatedTimestamp,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"space_id":               o.SpaceId,
			"title":                  o.Title,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieConversation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_timestamp":      types.Int64Type,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"space_id":               types.StringType,
			"title":                  types.StringType,
			"user_id":                types.Int64Type,
		},
	}
}

type GenieCreateConversationMessageRequest_SdkV2 struct {
	// User message content.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieCreateConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieCreateConversationMessageRequest_SdkV2) {
}

func (newState *GenieCreateConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieCreateConversationMessageRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieCreateConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieCreateConversationMessageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieCreateConversationMessageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieCreateConversationMessageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":         o.Content,
			"conversation_id": o.ConversationId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":         types.StringType,
			"conversation_id": types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest_SdkV2 struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieExecuteMessageQueryRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieExecuteMessageQueryRequest_SdkV2) {
}

func (newState *GenieExecuteMessageQueryRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieExecuteMessageQueryRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieExecuteMessageQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieExecuteMessageQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieExecuteMessageQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieExecuteMessageQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Get conversation message
type GenieGetConversationMessageRequest_SdkV2 struct {
	// The ID associated with the target conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetConversationMessageRequest_SdkV2) {
}

func (newState *GenieGetConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGetConversationMessageRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetConversationMessageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetConversationMessageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetConversationMessageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Get conversation message SQL query result
type GenieGetMessageQueryResultRequest_SdkV2 struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetMessageQueryResultRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultRequest_SdkV2) {
}

func (newState *GenieGetMessageQueryResultRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageQueryResultRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetMessageQueryResultRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGetMessageQueryResultResponse_SdkV2 struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse types.List `tfsdk:"statement_response" tf:"optional,object"`
}

func (newState *GenieGetMessageQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultResponse_SdkV2) {
}

func (newState *GenieGetMessageQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statement_response": reflect.TypeOf(sql_tf.StatementResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageQueryResultResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetMessageQueryResultResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_response": o.StatementResponse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_response": basetypes.ListType{
				ElemType: sql_tf.StatementResponse{}.Type(ctx),
			},
		},
	}
}

// GetStatementResponse returns the value of the StatementResponse field in GenieGetMessageQueryResultResponse_SdkV2 as
// a sql_tf.StatementResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieGetMessageQueryResultResponse_SdkV2) GetStatementResponse(ctx context.Context) (sql_tf.StatementResponse_SdkV2, bool) {
	var e sql_tf.StatementResponse_SdkV2
	if o.StatementResponse.IsNull() || o.StatementResponse.IsUnknown() {
		return e, false
	}
	var v []sql_tf.StatementResponse_SdkV2
	d := o.StatementResponse.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatementResponse sets the value of the StatementResponse field in GenieGetMessageQueryResultResponse_SdkV2.
func (o *GenieGetMessageQueryResultResponse_SdkV2) SetStatementResponse(ctx context.Context, v sql_tf.StatementResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_response"]
	o.StatementResponse = types.ListValueMust(t, vs)
}

type GenieMessage_SdkV2 struct {
	// AI produced response to the message
	Attachments types.List `tfsdk:"attachments" tf:"optional"`
	// User message content
	Content types.String `tfsdk:"content" tf:""`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp" tf:"optional"`
	// Error message if AI failed to respond to the message
	Error types.List `tfsdk:"error" tf:"optional,object"`
	// Message ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The result of SQL query if the message has a query attachment
	QueryResult types.List `tfsdk:"query_result" tf:"optional,object"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id" tf:""`
	// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart
	// context step to determine relevant context. * `ASKING_AI`: Waiting for
	// the LLM to respond to the users question. * `EXECUTING_QUERY`: Executing
	// AI provided SQL query. Get the SQL query result by calling
	// [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
	// **Important: The message status will stay in the `EXECUTING_QUERY` until
	// a client calls
	// [getMessageQueryResult](:method:genie/getMessageQueryResult)**. *
	// `FAILED`: Generating a response or the executing the query failed. Please
	// see `error` field. * `COMPLETED`: Message processing is completed.
	// Results are in the `attachments` field. Get the SQL query result by
	// calling [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
	// * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL
	// result is not available anymore. The user needs to execute the query
	// again. * `CANCELLED`: Message has been cancelled.
	Status types.String `tfsdk:"status" tf:"optional"`
	// ID of the user who created the message
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
}

func (newState *GenieMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieMessage_SdkV2) {
}

func (newState *GenieMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieMessage_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attachments":  reflect.TypeOf(GenieAttachment_SdkV2{}),
		"error":        reflect.TypeOf(MessageError_SdkV2{}),
		"query_result": reflect.TypeOf(Result_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachments":            o.Attachments,
			"content":                o.Content,
			"conversation_id":        o.ConversationId,
			"created_timestamp":      o.CreatedTimestamp,
			"error":                  o.Error,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"query_result":           o.QueryResult,
			"space_id":               o.SpaceId,
			"status":                 o.Status,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachments": basetypes.ListType{
				ElemType: GenieAttachment{}.Type(ctx),
			},
			"content":           types.StringType,
			"conversation_id":   types.StringType,
			"created_timestamp": types.Int64Type,
			"error": basetypes.ListType{
				ElemType: MessageError{}.Type(ctx),
			},
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query_result": basetypes.ListType{
				ElemType: Result{}.Type(ctx),
			},
			"space_id": types.StringType,
			"status":   types.StringType,
			"user_id":  types.Int64Type,
		},
	}
}

// GetAttachments returns the value of the Attachments field in GenieMessage_SdkV2 as
// a slice of GenieAttachment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage_SdkV2) GetAttachments(ctx context.Context) ([]GenieAttachment_SdkV2, bool) {
	if o.Attachments.IsNull() || o.Attachments.IsUnknown() {
		return nil, false
	}
	var v []GenieAttachment_SdkV2
	d := o.Attachments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAttachments sets the value of the Attachments field in GenieMessage_SdkV2.
func (o *GenieMessage_SdkV2) SetAttachments(ctx context.Context, v []GenieAttachment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["attachments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Attachments = types.ListValueMust(t, vs)
}

// GetError returns the value of the Error field in GenieMessage_SdkV2 as
// a MessageError_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage_SdkV2) GetError(ctx context.Context) (MessageError_SdkV2, bool) {
	var e MessageError_SdkV2
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []MessageError_SdkV2
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in GenieMessage_SdkV2.
func (o *GenieMessage_SdkV2) SetError(ctx context.Context, v MessageError_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

// GetQueryResult returns the value of the QueryResult field in GenieMessage_SdkV2 as
// a Result_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage_SdkV2) GetQueryResult(ctx context.Context) (Result_SdkV2, bool) {
	var e Result_SdkV2
	if o.QueryResult.IsNull() || o.QueryResult.IsUnknown() {
		return e, false
	}
	var v []Result_SdkV2
	d := o.QueryResult.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryResult sets the value of the QueryResult field in GenieMessage_SdkV2.
func (o *GenieMessage_SdkV2) SetQueryResult(ctx context.Context, v Result_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_result"]
	o.QueryResult = types.ListValueMust(t, vs)
}

type GenieStartConversationMessageRequest_SdkV2 struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieStartConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationMessageRequest_SdkV2) {
}

func (newState *GenieStartConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationMessageRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationMessageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieStartConversationMessageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieStartConversationMessageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":  o.Content,
			"space_id": o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":  types.StringType,
			"space_id": types.StringType,
		},
	}
}

type GenieStartConversationResponse_SdkV2 struct {
	Conversation types.List `tfsdk:"conversation" tf:"optional,object"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`

	Message types.List `tfsdk:"message" tf:"optional,object"`
	// Message ID
	MessageId types.String `tfsdk:"message_id" tf:""`
}

func (newState *GenieStartConversationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationResponse_SdkV2) {
}

func (newState *GenieStartConversationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"conversation": reflect.TypeOf(GenieConversation_SdkV2{}),
		"message":      reflect.TypeOf(GenieMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieStartConversationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieStartConversationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation":    o.Conversation,
			"conversation_id": o.ConversationId,
			"message":         o.Message,
			"message_id":      o.MessageId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation": basetypes.ListType{
				ElemType: GenieConversation{}.Type(ctx),
			},
			"conversation_id": types.StringType,
			"message": basetypes.ListType{
				ElemType: GenieMessage{}.Type(ctx),
			},
			"message_id": types.StringType,
		},
	}
}

// GetConversation returns the value of the Conversation field in GenieStartConversationResponse_SdkV2 as
// a GenieConversation_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieStartConversationResponse_SdkV2) GetConversation(ctx context.Context) (GenieConversation_SdkV2, bool) {
	var e GenieConversation_SdkV2
	if o.Conversation.IsNull() || o.Conversation.IsUnknown() {
		return e, false
	}
	var v []GenieConversation_SdkV2
	d := o.Conversation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConversation sets the value of the Conversation field in GenieStartConversationResponse_SdkV2.
func (o *GenieStartConversationResponse_SdkV2) SetConversation(ctx context.Context, v GenieConversation_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["conversation"]
	o.Conversation = types.ListValueMust(t, vs)
}

// GetMessage returns the value of the Message field in GenieStartConversationResponse_SdkV2 as
// a GenieMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieStartConversationResponse_SdkV2) GetMessage(ctx context.Context) (GenieMessage_SdkV2, bool) {
	var e GenieMessage_SdkV2
	if o.Message.IsNull() || o.Message.IsUnknown() {
		return e, false
	}
	var v []GenieMessage_SdkV2
	d := o.Message.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMessage sets the value of the Message field in GenieStartConversationResponse_SdkV2.
func (o *GenieStartConversationResponse_SdkV2) SetMessage(ctx context.Context, v GenieMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["message"]
	o.Message = types.ListValueMust(t, vs)
}

// Get dashboard
type GetDashboardRequest_SdkV2 struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDashboardRequest_SdkV2) {
}

func (newState *GetDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get published dashboard
type GetPublishedDashboardRequest_SdkV2 struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetPublishedDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedDashboardRequest_SdkV2) {
}

func (newState *GetPublishedDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPublishedDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get dashboard schedule
type GetScheduleRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *GetScheduleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetScheduleRequest_SdkV2) {
}

func (newState *GetScheduleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetScheduleRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetScheduleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetScheduleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetScheduleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetScheduleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

// Get schedule subscription
type GetSubscriptionRequest_SdkV2 struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-"`
}

func (newState *GetSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSubscriptionRequest_SdkV2) {
}

func (newState *GetSubscriptionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetSubscriptionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSubscriptionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSubscriptionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSubscriptionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":    o.DashboardId,
			"schedule_id":     o.ScheduleId,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

// List dashboards
type ListDashboardsRequest_SdkV2 struct {
	// The number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed types.Bool `tfsdk:"-"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View types.String `tfsdk:"-"`
}

func (newState *ListDashboardsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsRequest_SdkV2) {
}

func (newState *ListDashboardsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListDashboardsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDashboardsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDashboardsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    o.PageSize,
			"page_token":   o.PageToken,
			"show_trashed": o.ShowTrashed,
			"view":         o.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"show_trashed": types.BoolType,
			"view":         types.StringType,
		},
	}
}

type ListDashboardsResponse_SdkV2 struct {
	Dashboards types.List `tfsdk:"dashboards" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed"`
}

func (newState *ListDashboardsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsResponse_SdkV2) {
}

func (newState *ListDashboardsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListDashboardsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboards": reflect.TypeOf(Dashboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDashboardsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDashboardsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboards":      o.Dashboards,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboards": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDashboards returns the value of the Dashboards field in ListDashboardsResponse_SdkV2 as
// a slice of Dashboard_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDashboardsResponse_SdkV2) GetDashboards(ctx context.Context) ([]Dashboard_SdkV2, bool) {
	if o.Dashboards.IsNull() || o.Dashboards.IsUnknown() {
		return nil, false
	}
	var v []Dashboard_SdkV2
	d := o.Dashboards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDashboards sets the value of the Dashboards field in ListDashboardsResponse_SdkV2.
func (o *ListDashboardsResponse_SdkV2) SetDashboards(ctx context.Context, v []Dashboard_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dashboards = types.ListValueMust(t, vs)
}

// List dashboard schedules
type ListSchedulesRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSchedulesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesRequest_SdkV2) {
}

func (newState *ListSchedulesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSchedulesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchedulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSchedulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"page_size":    o.PageSize,
			"page_token":   o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchedulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
		},
	}
}

type ListSchedulesResponse_SdkV2 struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed"`

	Schedules types.List `tfsdk:"schedules" tf:"optional"`
}

func (newState *ListSchedulesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesResponse_SdkV2) {
}

func (newState *ListSchedulesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSchedulesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedules": reflect.TypeOf(Schedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchedulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSchedulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schedules":       o.Schedules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchedulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schedules": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
		},
	}
}

// GetSchedules returns the value of the Schedules field in ListSchedulesResponse_SdkV2 as
// a slice of Schedule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSchedulesResponse_SdkV2) GetSchedules(ctx context.Context) ([]Schedule_SdkV2, bool) {
	if o.Schedules.IsNull() || o.Schedules.IsUnknown() {
		return nil, false
	}
	var v []Schedule_SdkV2
	d := o.Schedules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedules sets the value of the Schedules field in ListSchedulesResponse_SdkV2.
func (o *ListSchedulesResponse_SdkV2) SetSchedules(ctx context.Context, v []Schedule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schedules = types.ListValueMust(t, vs)
}

// List schedule subscriptions
type ListSubscriptionsRequest_SdkV2 struct {
	// UUID identifying the dashboard which the subscriptions belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of subscriptions to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscriptions belongs.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *ListSubscriptionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsRequest_SdkV2) {
}

func (newState *ListSubscriptionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSubscriptionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSubscriptionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"page_size":    o.PageSize,
			"page_token":   o.PageToken,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type ListSubscriptionsResponse_SdkV2 struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed"`

	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *ListSubscriptionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsResponse_SdkV2) {
}

func (newState *ListSubscriptionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(Subscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSubscriptionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSubscriptionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"subscriptions":   o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"subscriptions": basetypes.ListType{
				ElemType: Subscription{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in ListSubscriptionsResponse_SdkV2 as
// a slice of Subscription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSubscriptionsResponse_SdkV2) GetSubscriptions(ctx context.Context) ([]Subscription_SdkV2, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []Subscription_SdkV2
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in ListSubscriptionsResponse_SdkV2.
func (o *ListSubscriptionsResponse_SdkV2) SetSubscriptions(ctx context.Context, v []Subscription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type MessageError_SdkV2 struct {
	Error types.String `tfsdk:"error" tf:"optional"`

	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *MessageError_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MessageError_SdkV2) {
}

func (newState *MessageError_SdkV2) SyncEffectiveFieldsDuringRead(existingState MessageError_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MessageError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MessageError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MessageError_SdkV2
// only implements ToObjectValue() and Type().
func (o MessageError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error": o.Error,
			"type":  o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MessageError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"type":  types.StringType,
		},
	}
}

type MigrateDashboardRequest_SdkV2 struct {
	// Display name for the new Lakeview dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId types.String `tfsdk:"source_dashboard_id" tf:""`
}

func (newState *MigrateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigrateDashboardRequest_SdkV2) {
}

func (newState *MigrateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState MigrateDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigrateDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrateDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o MigrateDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":        o.DisplayName,
			"parent_path":         o.ParentPath,
			"source_dashboard_id": o.SourceDashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":        types.StringType,
			"parent_path":         types.StringType,
			"source_dashboard_id": types.StringType,
		},
	}
}

type PublishRequest_SdkV2 struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishRequest_SdkV2) {
}

func (newState *PublishRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PublishRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PublishRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":      o.DashboardId,
			"embed_credentials": o.EmbedCredentials,
			"warehouse_id":      o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":      types.StringType,
			"embed_credentials": types.BoolType,
			"warehouse_id":      types.StringType,
		},
	}
}

type PublishedDashboard_SdkV2 struct {
	// The display name of the published dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"computed"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time" tf:"computed"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishedDashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedDashboard_SdkV2) {
}

func (newState *PublishedDashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState PublishedDashboard_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishedDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishedDashboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishedDashboard_SdkV2
// only implements ToObjectValue() and Type().
func (o PublishedDashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":         o.DisplayName,
			"embed_credentials":    o.EmbedCredentials,
			"revision_create_time": o.RevisionCreateTime,
			"warehouse_id":         o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishedDashboard_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":         types.StringType,
			"embed_credentials":    types.BoolType,
			"revision_create_time": types.StringType,
			"warehouse_id":         types.StringType,
		},
	}
}

type QueryAttachment_SdkV2 struct {
	CachedQuerySchema types.List `tfsdk:"cached_query_schema" tf:"optional,object"`
	// Description of the query
	Description types.String `tfsdk:"description" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// If the query was created on an instruction (trusted asset) we link to the
	// id
	InstructionId types.String `tfsdk:"instruction_id" tf:"optional"`
	// Always store the title next to the id in case the original instruction
	// title changes or the instruction is deleted.
	InstructionTitle types.String `tfsdk:"instruction_title" tf:"optional"`
	// Time when the user updated the query last
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// AI generated SQL query
	Query types.String `tfsdk:"query" tf:"optional"`
	// Name of the query
	Title types.String `tfsdk:"title" tf:"optional"`
}

func (newState *QueryAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryAttachment_SdkV2) {
}

func (newState *QueryAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryAttachment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryAttachment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cached_query_schema": reflect.TypeOf(QuerySchema_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryAttachment_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryAttachment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cached_query_schema":    o.CachedQuerySchema,
			"description":            o.Description,
			"id":                     o.Id,
			"instruction_id":         o.InstructionId,
			"instruction_title":      o.InstructionTitle,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"query":                  o.Query,
			"title":                  o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryAttachment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cached_query_schema": basetypes.ListType{
				ElemType: QuerySchema{}.Type(ctx),
			},
			"description":            types.StringType,
			"id":                     types.StringType,
			"instruction_id":         types.StringType,
			"instruction_title":      types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query":                  types.StringType,
			"title":                  types.StringType,
		},
	}
}

// GetCachedQuerySchema returns the value of the CachedQuerySchema field in QueryAttachment_SdkV2 as
// a QuerySchema_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryAttachment_SdkV2) GetCachedQuerySchema(ctx context.Context) (QuerySchema_SdkV2, bool) {
	var e QuerySchema_SdkV2
	if o.CachedQuerySchema.IsNull() || o.CachedQuerySchema.IsUnknown() {
		return e, false
	}
	var v []QuerySchema_SdkV2
	d := o.CachedQuerySchema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCachedQuerySchema sets the value of the CachedQuerySchema field in QueryAttachment_SdkV2.
func (o *QueryAttachment_SdkV2) SetCachedQuerySchema(ctx context.Context, v QuerySchema_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cached_query_schema"]
	o.CachedQuerySchema = types.ListValueMust(t, vs)
}

type QuerySchema_SdkV2 struct {
	Columns types.List `tfsdk:"columns" tf:"optional"`
	// Used to determine if the stored query schema is compatible with the
	// latest run. The service should always clear the schema when the query is
	// re-executed.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

func (newState *QuerySchema_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuerySchema_SdkV2) {
}

func (newState *QuerySchema_SdkV2) SyncEffectiveFieldsDuringRead(existingState QuerySchema_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuerySchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuerySchema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(QuerySchemaColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuerySchema_SdkV2
// only implements ToObjectValue() and Type().
func (o QuerySchema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":      o.Columns,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QuerySchema_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: QuerySchemaColumn{}.Type(ctx),
			},
			"statement_id": types.StringType,
		},
	}
}

// GetColumns returns the value of the Columns field in QuerySchema_SdkV2 as
// a slice of QuerySchemaColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QuerySchema_SdkV2) GetColumns(ctx context.Context) ([]QuerySchemaColumn_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []QuerySchemaColumn_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in QuerySchema_SdkV2.
func (o *QuerySchema_SdkV2) SetColumns(ctx context.Context, v []QuerySchemaColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type QuerySchemaColumn_SdkV2 struct {
	// Populated from
	// https://docs.databricks.com/sql/language-manual/sql-ref-datatypes.html
	DataType types.String `tfsdk:"data_type" tf:""`

	Name types.String `tfsdk:"name" tf:""`
	// Corresponds to type desc
	TypeText types.String `tfsdk:"type_text" tf:""`
}

func (newState *QuerySchemaColumn_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuerySchemaColumn_SdkV2) {
}

func (newState *QuerySchemaColumn_SdkV2) SyncEffectiveFieldsDuringRead(existingState QuerySchemaColumn_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuerySchemaColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuerySchemaColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuerySchemaColumn_SdkV2
// only implements ToObjectValue() and Type().
func (o QuerySchemaColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_type": o.DataType,
			"name":      o.Name,
			"type_text": o.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QuerySchemaColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_type": types.StringType,
			"name":      types.StringType,
			"type_text": types.StringType,
		},
	}
}

type Result_SdkV2 struct {
	// If result is truncated
	IsTruncated types.Bool `tfsdk:"is_truncated" tf:"optional"`
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

func (newState *Result_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Result_SdkV2) {
}

func (newState *Result_SdkV2) SyncEffectiveFieldsDuringRead(existingState Result_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Result.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Result_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Result_SdkV2
// only implements ToObjectValue() and Type().
func (o Result_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_truncated": o.IsTruncated,
			"row_count":    o.RowCount,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Result_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated": types.BoolType,
			"row_count":    types.Int64Type,
			"statement_id": types.StringType,
		},
	}
}

type Schedule_SdkV2 struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule types.List `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag" tf:"computed"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Schedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Schedule_SdkV2) {
}

func (newState *Schedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState Schedule_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Schedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Schedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron_schedule": reflect.TypeOf(CronSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Schedule_SdkV2
// only implements ToObjectValue() and Type().
func (o Schedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":   o.CreateTime,
			"cron_schedule": o.CronSchedule,
			"dashboard_id":  o.DashboardId,
			"display_name":  o.DisplayName,
			"etag":          o.Etag,
			"pause_status":  o.PauseStatus,
			"schedule_id":   o.ScheduleId,
			"update_time":   o.UpdateTime,
			"warehouse_id":  o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Schedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"cron_schedule": basetypes.ListType{
				ElemType: CronSchedule{}.Type(ctx),
			},
			"dashboard_id": types.StringType,
			"display_name": types.StringType,
			"etag":         types.StringType,
			"pause_status": types.StringType,
			"schedule_id":  types.StringType,
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetCronSchedule returns the value of the CronSchedule field in Schedule_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Schedule_SdkV2) GetCronSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.CronSchedule.IsNull() || o.CronSchedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.CronSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCronSchedule sets the value of the CronSchedule field in Schedule_SdkV2.
func (o *Schedule_SdkV2) SetCronSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cron_schedule"]
	o.CronSchedule = types.ListValueMust(t, vs)
}

type Subscriber_SdkV2 struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber types.List `tfsdk:"destination_subscriber" tf:"optional,object"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber types.List `tfsdk:"user_subscriber" tf:"optional,object"`
}

func (newState *Subscriber_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscriber_SdkV2) {
}

func (newState *Subscriber_SdkV2) SyncEffectiveFieldsDuringRead(existingState Subscriber_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscriber.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscriber_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destination_subscriber": reflect.TypeOf(SubscriptionSubscriberDestination_SdkV2{}),
		"user_subscriber":        reflect.TypeOf(SubscriptionSubscriberUser_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Subscriber_SdkV2
// only implements ToObjectValue() and Type().
func (o Subscriber_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_subscriber": o.DestinationSubscriber,
			"user_subscriber":        o.UserSubscriber,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Subscriber_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_subscriber": basetypes.ListType{
				ElemType: SubscriptionSubscriberDestination{}.Type(ctx),
			},
			"user_subscriber": basetypes.ListType{
				ElemType: SubscriptionSubscriberUser{}.Type(ctx),
			},
		},
	}
}

// GetDestinationSubscriber returns the value of the DestinationSubscriber field in Subscriber_SdkV2 as
// a SubscriptionSubscriberDestination_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscriber_SdkV2) GetDestinationSubscriber(ctx context.Context) (SubscriptionSubscriberDestination_SdkV2, bool) {
	var e SubscriptionSubscriberDestination_SdkV2
	if o.DestinationSubscriber.IsNull() || o.DestinationSubscriber.IsUnknown() {
		return e, false
	}
	var v []SubscriptionSubscriberDestination_SdkV2
	d := o.DestinationSubscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDestinationSubscriber sets the value of the DestinationSubscriber field in Subscriber_SdkV2.
func (o *Subscriber_SdkV2) SetDestinationSubscriber(ctx context.Context, v SubscriptionSubscriberDestination_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["destination_subscriber"]
	o.DestinationSubscriber = types.ListValueMust(t, vs)
}

// GetUserSubscriber returns the value of the UserSubscriber field in Subscriber_SdkV2 as
// a SubscriptionSubscriberUser_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscriber_SdkV2) GetUserSubscriber(ctx context.Context) (SubscriptionSubscriberUser_SdkV2, bool) {
	var e SubscriptionSubscriberUser_SdkV2
	if o.UserSubscriber.IsNull() || o.UserSubscriber.IsUnknown() {
		return e, false
	}
	var v []SubscriptionSubscriberUser_SdkV2
	d := o.UserSubscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUserSubscriber sets the value of the UserSubscriber field in Subscriber_SdkV2.
func (o *Subscriber_SdkV2) SetUserSubscriber(ctx context.Context, v SubscriptionSubscriberUser_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_subscriber"]
	o.UserSubscriber = types.ListValueMust(t, vs)
}

type Subscription_SdkV2 struct {
	// A timestamp indicating when the subscription was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id" tf:"computed"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag" tf:"computed"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber types.List `tfsdk:"subscriber" tf:"object"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"computed"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed"`
}

func (newState *Subscription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscription_SdkV2) {
}

func (newState *Subscription_SdkV2) SyncEffectiveFieldsDuringRead(existingState Subscription_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriber": reflect.TypeOf(Subscriber_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Subscription_SdkV2
// only implements ToObjectValue() and Type().
func (o Subscription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":        o.CreateTime,
			"created_by_user_id": o.CreatedByUserId,
			"dashboard_id":       o.DashboardId,
			"etag":               o.Etag,
			"schedule_id":        o.ScheduleId,
			"subscriber":         o.Subscriber,
			"subscription_id":    o.SubscriptionId,
			"update_time":        o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Subscription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"created_by_user_id": types.Int64Type,
			"dashboard_id":       types.StringType,
			"etag":               types.StringType,
			"schedule_id":        types.StringType,
			"subscriber": basetypes.ListType{
				ElemType: Subscriber{}.Type(ctx),
			},
			"subscription_id": types.StringType,
			"update_time":     types.StringType,
		},
	}
}

// GetSubscriber returns the value of the Subscriber field in Subscription_SdkV2 as
// a Subscriber_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscription_SdkV2) GetSubscriber(ctx context.Context) (Subscriber_SdkV2, bool) {
	var e Subscriber_SdkV2
	if o.Subscriber.IsNull() || o.Subscriber.IsUnknown() {
		return e, false
	}
	var v []Subscriber_SdkV2
	d := o.Subscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSubscriber sets the value of the Subscriber field in Subscription_SdkV2.
func (o *Subscription_SdkV2) SetSubscriber(ctx context.Context, v Subscriber_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriber"]
	o.Subscriber = types.ListValueMust(t, vs)
}

type SubscriptionSubscriberDestination_SdkV2 struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id" tf:"computed"`
}

func (newState *SubscriptionSubscriberDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberDestination_SdkV2) {
}

func (newState *SubscriptionSubscriberDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberDestination_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubscriptionSubscriberDestination_SdkV2
// only implements ToObjectValue() and Type().
func (o SubscriptionSubscriberDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": o.DestinationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
		},
	}
}

type SubscriptionSubscriberUser_SdkV2 struct {
	// UserId of the subscriber.
	UserId types.Int64 `tfsdk:"user_id" tf:"computed"`
}

func (newState *SubscriptionSubscriberUser_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberUser_SdkV2) {
}

func (newState *SubscriptionSubscriberUser_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberUser_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberUser.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberUser_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubscriptionSubscriberUser_SdkV2
// only implements ToObjectValue() and Type().
func (o SubscriptionSubscriberUser_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.Int64Type,
		},
	}
}

type TextAttachment_SdkV2 struct {
	// AI generated message
	Content types.String `tfsdk:"content" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *TextAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextAttachment_SdkV2) {
}

func (newState *TextAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState TextAttachment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TextAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TextAttachment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TextAttachment_SdkV2
// only implements ToObjectValue() and Type().
func (o TextAttachment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": o.Content,
			"id":      o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TextAttachment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"id":      types.StringType,
		},
	}
}

// Trash dashboard
type TrashDashboardRequest_SdkV2 struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *TrashDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardRequest_SdkV2) {
}

func (newState *TrashDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState TrashDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TrashDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type TrashDashboardResponse_SdkV2 struct {
}

func (newState *TrashDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardResponse_SdkV2) {
}

func (newState *TrashDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState TrashDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o TrashDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o TrashDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Unpublish dashboard
type UnpublishDashboardRequest_SdkV2 struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UnpublishDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardRequest_SdkV2) {
}

func (newState *UnpublishDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpublishDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UnpublishDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type UnpublishDashboardResponse_SdkV2 struct {
}

func (newState *UnpublishDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardResponse_SdkV2) {
}

func (newState *UnpublishDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpublishDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UnpublishDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Update dashboard
type UpdateDashboardRequest_SdkV2 struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UpdateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDashboardRequest_SdkV2) {
}

func (newState *UpdateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard":    o.Dashboard,
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
			"dashboard_id": types.StringType,
		},
	}
}

// GetDashboard returns the value of the Dashboard field in UpdateDashboardRequest_SdkV2 as
// a Dashboard_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDashboardRequest_SdkV2) GetDashboard(ctx context.Context) (Dashboard_SdkV2, bool) {
	var e Dashboard_SdkV2
	if o.Dashboard.IsNull() || o.Dashboard.IsUnknown() {
		return e, false
	}
	var v []Dashboard_SdkV2
	d := o.Dashboard.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboard sets the value of the Dashboard field in UpdateDashboardRequest_SdkV2.
func (o *UpdateDashboardRequest_SdkV2) SetDashboard(ctx context.Context, v Dashboard_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard"]
	o.Dashboard = types.ListValueMust(t, vs)
}

// Update dashboard schedule
type UpdateScheduleRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *UpdateScheduleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateScheduleRequest_SdkV2) {
}

func (newState *UpdateScheduleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateScheduleRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateScheduleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateScheduleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateScheduleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule":     o.Schedule,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
			"schedule_id": types.StringType,
		},
	}
}

// GetSchedule returns the value of the Schedule field in UpdateScheduleRequest_SdkV2 as
// a Schedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateScheduleRequest_SdkV2) GetSchedule(ctx context.Context) (Schedule_SdkV2, bool) {
	var e Schedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []Schedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in UpdateScheduleRequest_SdkV2.
func (o *UpdateScheduleRequest_SdkV2) SetSchedule(ctx context.Context, v Schedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}
