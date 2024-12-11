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
type CreateDashboardRequest struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDashboardRequest) {
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState CreateDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDashboardRequest
// only implements ToObjectValue() and Type().
func (o CreateDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard": o.Dashboard,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{ElemType: Dashboard{}.Type(ctx)},
		},
	}
}

// GetDashboard returns the value of the Dashboard field in CreateDashboardRequest as
// a Dashboard value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDashboardRequest) GetDashboard(ctx context.Context) (Dashboard, bool) {
	var e Dashboard
	if o.Dashboard.IsNull() || o.Dashboard.IsUnknown() {
		return e, false
	}
	var v []Dashboard
	d := o.Dashboard.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboard sets the value of the Dashboard field in CreateDashboardRequest.
func (o *CreateDashboardRequest) SetDashboard(ctx context.Context, v Dashboard) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard"]
	o.Dashboard = types.ListValueMust(t, vs)
}

// Create dashboard schedule
type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScheduleRequest) {
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState CreateScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScheduleRequest
// only implements ToObjectValue() and Type().
func (o CreateScheduleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule":     o.Schedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule":     basetypes.ListType{ElemType: Schedule{}.Type(ctx)},
		},
	}
}

// GetSchedule returns the value of the Schedule field in CreateScheduleRequest as
// a Schedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateScheduleRequest) GetSchedule(ctx context.Context) (Schedule, bool) {
	var e Schedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []Schedule
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in CreateScheduleRequest.
func (o *CreateScheduleRequest) SetSchedule(ctx context.Context, v Schedule) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// Create schedule subscription
type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`

	Subscription types.List `tfsdk:"subscription" tf:"optional,object"`
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSubscriptionRequest) {
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState CreateSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscription": reflect.TypeOf(Subscription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSubscriptionRequest
// only implements ToObjectValue() and Type().
func (o CreateSubscriptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule_id":  o.ScheduleId,
			"subscription": o.Subscription,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
			"subscription": basetypes.ListType{ElemType: Subscription{}.Type(ctx)},
		},
	}
}

// GetSubscription returns the value of the Subscription field in CreateSubscriptionRequest as
// a Subscription value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSubscriptionRequest) GetSubscription(ctx context.Context) (Subscription, bool) {
	var e Subscription
	if o.Subscription.IsNull() || o.Subscription.IsUnknown() {
		return e, false
	}
	var v []Subscription
	d := o.Subscription.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSubscription sets the value of the Subscription field in CreateSubscriptionRequest.
func (o *CreateSubscriptionRequest) SetSubscription(ctx context.Context, v Subscription) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscription"]
	o.Subscription = types.ListValueMust(t, vs)
}

type CronSchedule struct {
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

func (newState *CronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule) {
}

func (newState *CronSchedule) SyncEffectiveFieldsDuringRead(existingState CronSchedule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronSchedule
// only implements ToObjectValue() and Type().
func (o CronSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_expression": o.QuartzCronExpression,
			"timezone_id":            o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath types.String `tfsdk:"parent_path" tf:"computed,optional"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path types.String `tfsdk:"path" tf:"computed,optional"`
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
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Dashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard) {
}

func (newState *Dashboard) SyncEffectiveFieldsDuringRead(existingState Dashboard) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dashboard
// only implements ToObjectValue() and Type().
func (o Dashboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Dashboard) Type(ctx context.Context) attr.Type {
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
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleRequest) {
}

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScheduleRequest
// only implements ToObjectValue() and Type().
func (o DeleteScheduleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"etag":         o.Etag,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"etag":         types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type DeleteScheduleResponse struct {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleResponse) {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScheduleResponse
// only implements ToObjectValue() and Type().
func (o DeleteScheduleResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
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

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionRequest) {
}

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSubscriptionRequest
// only implements ToObjectValue() and Type().
func (o DeleteSubscriptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DeleteSubscriptionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"etag":            types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type DeleteSubscriptionResponse struct {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionResponse) {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSubscriptionResponse
// only implements ToObjectValue() and Type().
func (o DeleteSubscriptionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Genie AI Response
type GenieAttachment struct {
	Query types.List `tfsdk:"query" tf:"optional,object"`

	Text types.List `tfsdk:"text" tf:"optional,object"`
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieAttachment) {
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringRead(existingState GenieAttachment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(QueryAttachment{}),
		"text":  reflect.TypeOf(TextAttachment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAttachment
// only implements ToObjectValue() and Type().
func (o GenieAttachment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query": o.Query,
			"text":  o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query": basetypes.ListType{ElemType: QueryAttachment{}.Type(ctx)},
			"text":  basetypes.ListType{ElemType: TextAttachment{}.Type(ctx)},
		},
	}
}

// GetQuery returns the value of the Query field in GenieAttachment as
// a QueryAttachment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment) GetQuery(ctx context.Context) (QueryAttachment, bool) {
	var e QueryAttachment
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []QueryAttachment
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in GenieAttachment.
func (o *GenieAttachment) SetQuery(ctx context.Context, v QueryAttachment) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

// GetText returns the value of the Text field in GenieAttachment as
// a TextAttachment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment) GetText(ctx context.Context) (TextAttachment, bool) {
	var e TextAttachment
	if o.Text.IsNull() || o.Text.IsUnknown() {
		return e, false
	}
	var v []TextAttachment
	d := o.Text.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetText sets the value of the Text field in GenieAttachment.
func (o *GenieAttachment) SetText(ctx context.Context, v TextAttachment) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["text"]
	o.Text = types.ListValueMust(t, vs)
}

type GenieConversation struct {
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

func (newState *GenieConversation) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieConversation) {
}

func (newState *GenieConversation) SyncEffectiveFieldsDuringRead(existingState GenieConversation) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieConversation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieConversation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieConversation
// only implements ToObjectValue() and Type().
func (o GenieConversation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieConversation) Type(ctx context.Context) attr.Type {
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

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieCreateConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieCreateConversationMessageRequest) {
}

func (newState *GenieCreateConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieCreateConversationMessageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieCreateConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieCreateConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieCreateConversationMessageRequest
// only implements ToObjectValue() and Type().
func (o GenieCreateConversationMessageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":         o.Content,
			"conversation_id": o.ConversationId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":         types.StringType,
			"conversation_id": types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieExecuteMessageQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieExecuteMessageQueryRequest) {
}

func (newState *GenieExecuteMessageQueryRequest) SyncEffectiveFieldsDuringRead(existingState GenieExecuteMessageQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieExecuteMessageQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieExecuteMessageQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieExecuteMessageQueryRequest
// only implements ToObjectValue() and Type().
func (o GenieExecuteMessageQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Get conversation message
type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetConversationMessageRequest) {
}

func (newState *GenieGetConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieGetConversationMessageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetConversationMessageRequest
// only implements ToObjectValue() and Type().
func (o GenieGetConversationMessageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Get conversation message SQL query result
type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetMessageQueryResultRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultRequest) {
}

func (newState *GenieGetMessageQueryResultRequest) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageQueryResultRequest
// only implements ToObjectValue() and Type().
func (o GenieGetMessageQueryResultRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse types.List `tfsdk:"statement_response" tf:"optional,object"`
}

func (newState *GenieGetMessageQueryResultResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultResponse) {
}

func (newState *GenieGetMessageQueryResultResponse) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statement_response": reflect.TypeOf(sql_tf.StatementResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageQueryResultResponse
// only implements ToObjectValue() and Type().
func (o GenieGetMessageQueryResultResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_response": o.StatementResponse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_response": basetypes.ListType{ElemType: sql_tf.StatementResponse{}.Type(ctx)},
		},
	}
}

// GetStatementResponse returns the value of the StatementResponse field in GenieGetMessageQueryResultResponse as
// a sql_tf.StatementResponse value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieGetMessageQueryResultResponse) GetStatementResponse(ctx context.Context) (sql_tf.StatementResponse, bool) {
	var e sql_tf.StatementResponse
	if o.StatementResponse.IsNull() || o.StatementResponse.IsUnknown() {
		return e, false
	}
	var v []sql_tf.StatementResponse
	d := o.StatementResponse.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatementResponse sets the value of the StatementResponse field in GenieGetMessageQueryResultResponse.
func (o *GenieGetMessageQueryResultResponse) SetStatementResponse(ctx context.Context, v sql_tf.StatementResponse) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_response"]
	o.StatementResponse = types.ListValueMust(t, vs)
}

type GenieMessage struct {
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

func (newState *GenieMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieMessage) {
}

func (newState *GenieMessage) SyncEffectiveFieldsDuringRead(existingState GenieMessage) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attachments":  reflect.TypeOf(GenieAttachment{}),
		"error":        reflect.TypeOf(MessageError{}),
		"query_result": reflect.TypeOf(Result{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieMessage
// only implements ToObjectValue() and Type().
func (o GenieMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachments":            basetypes.ListType{ElemType: basetypes.ListType{ElemType: GenieAttachment{}.Type(ctx)}},
			"content":                types.StringType,
			"conversation_id":        types.StringType,
			"created_timestamp":      types.Int64Type,
			"error":                  basetypes.ListType{ElemType: MessageError{}.Type(ctx)},
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query_result":           basetypes.ListType{ElemType: Result{}.Type(ctx)},
			"space_id":               types.StringType,
			"status":                 types.StringType,
			"user_id":                types.Int64Type,
		},
	}
}

// GetAttachments returns the value of the Attachments field in GenieMessage as
// a slice of GenieAttachment values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage) GetAttachments(ctx context.Context) ([]GenieAttachment, bool) {
	if o.Attachments.IsNull() || o.Attachments.IsUnknown() {
		return nil, false
	}
	var v []GenieAttachment
	d := o.Attachments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAttachments sets the value of the Attachments field in GenieMessage.
func (o *GenieMessage) SetAttachments(ctx context.Context, v []GenieAttachment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["attachments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Attachments = types.ListValueMust(t, vs)
}

// GetError returns the value of the Error field in GenieMessage as
// a MessageError value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage) GetError(ctx context.Context) (MessageError, bool) {
	var e MessageError
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []MessageError
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in GenieMessage.
func (o *GenieMessage) SetError(ctx context.Context, v MessageError) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

// GetQueryResult returns the value of the QueryResult field in GenieMessage as
// a Result value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage) GetQueryResult(ctx context.Context) (Result, bool) {
	var e Result
	if o.QueryResult.IsNull() || o.QueryResult.IsUnknown() {
		return e, false
	}
	var v []Result
	d := o.QueryResult.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryResult sets the value of the QueryResult field in GenieMessage.
func (o *GenieMessage) SetQueryResult(ctx context.Context, v Result) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_result"]
	o.QueryResult = types.ListValueMust(t, vs)
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieStartConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationMessageRequest) {
}

func (newState *GenieStartConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationMessageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieStartConversationMessageRequest
// only implements ToObjectValue() and Type().
func (o GenieStartConversationMessageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":  o.Content,
			"space_id": o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":  types.StringType,
			"space_id": types.StringType,
		},
	}
}

type GenieStartConversationResponse struct {
	Conversation types.List `tfsdk:"conversation" tf:"optional,object"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`

	Message types.List `tfsdk:"message" tf:"optional,object"`
	// Message ID
	MessageId types.String `tfsdk:"message_id" tf:""`
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationResponse) {
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"conversation": reflect.TypeOf(GenieConversation{}),
		"message":      reflect.TypeOf(GenieMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieStartConversationResponse
// only implements ToObjectValue() and Type().
func (o GenieStartConversationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieStartConversationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation":    basetypes.ListType{ElemType: GenieConversation{}.Type(ctx)},
			"conversation_id": types.StringType,
			"message":         basetypes.ListType{ElemType: GenieMessage{}.Type(ctx)},
			"message_id":      types.StringType,
		},
	}
}

// GetConversation returns the value of the Conversation field in GenieStartConversationResponse as
// a GenieConversation value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieStartConversationResponse) GetConversation(ctx context.Context) (GenieConversation, bool) {
	var e GenieConversation
	if o.Conversation.IsNull() || o.Conversation.IsUnknown() {
		return e, false
	}
	var v []GenieConversation
	d := o.Conversation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConversation sets the value of the Conversation field in GenieStartConversationResponse.
func (o *GenieStartConversationResponse) SetConversation(ctx context.Context, v GenieConversation) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["conversation"]
	o.Conversation = types.ListValueMust(t, vs)
}

// GetMessage returns the value of the Message field in GenieStartConversationResponse as
// a GenieMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieStartConversationResponse) GetMessage(ctx context.Context) (GenieMessage, bool) {
	var e GenieMessage
	if o.Message.IsNull() || o.Message.IsUnknown() {
		return e, false
	}
	var v []GenieMessage
	d := o.Message.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMessage sets the value of the Message field in GenieStartConversationResponse.
func (o *GenieStartConversationResponse) SetMessage(ctx context.Context, v GenieMessage) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["message"]
	o.Message = types.ListValueMust(t, vs)
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDashboardRequest) {
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardRequest
// only implements ToObjectValue() and Type().
func (o GetDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedDashboardRequest) {
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetPublishedDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardRequest
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetScheduleRequest) {
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringRead(existingState GetScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetScheduleRequest
// only implements ToObjectValue() and Type().
func (o GetScheduleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetScheduleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-"`
}

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSubscriptionRequest) {
}

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState GetSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSubscriptionRequest
// only implements ToObjectValue() and Type().
func (o GetSubscriptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":    o.DashboardId,
			"schedule_id":     o.ScheduleId,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

// List dashboards
type ListDashboardsRequest struct {
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

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsRequest) {
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringRead(existingState ListDashboardsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDashboardsRequest
// only implements ToObjectValue() and Type().
func (o ListDashboardsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListDashboardsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"show_trashed": types.BoolType,
			"view":         types.StringType,
		},
	}
}

type ListDashboardsResponse struct {
	Dashboards types.List `tfsdk:"dashboards" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsResponse) {
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringRead(existingState ListDashboardsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboards": reflect.TypeOf(Dashboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDashboardsResponse
// only implements ToObjectValue() and Type().
func (o ListDashboardsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboards":      o.Dashboards,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboards":      basetypes.ListType{ElemType: basetypes.ListType{ElemType: Dashboard{}.Type(ctx)}},
			"next_page_token": types.StringType,
		},
	}
}

// GetDashboards returns the value of the Dashboards field in ListDashboardsResponse as
// a slice of Dashboard values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDashboardsResponse) GetDashboards(ctx context.Context) ([]Dashboard, bool) {
	if o.Dashboards.IsNull() || o.Dashboards.IsUnknown() {
		return nil, false
	}
	var v []Dashboard
	d := o.Dashboards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDashboards sets the value of the Dashboards field in ListDashboardsResponse.
func (o *ListDashboardsResponse) SetDashboards(ctx context.Context, v []Dashboard) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dashboards = types.ListValueMust(t, vs)
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesRequest) {
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringRead(existingState ListSchedulesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchedulesRequest
// only implements ToObjectValue() and Type().
func (o ListSchedulesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"page_size":    o.PageSize,
			"page_token":   o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
		},
	}
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`

	Schedules types.List `tfsdk:"schedules" tf:"optional"`
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesResponse) {
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringRead(existingState ListSchedulesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedules": reflect.TypeOf(Schedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchedulesResponse
// only implements ToObjectValue() and Type().
func (o ListSchedulesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schedules":       o.Schedules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schedules":       basetypes.ListType{ElemType: basetypes.ListType{ElemType: Schedule{}.Type(ctx)}},
		},
	}
}

// GetSchedules returns the value of the Schedules field in ListSchedulesResponse as
// a slice of Schedule values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSchedulesResponse) GetSchedules(ctx context.Context) ([]Schedule, bool) {
	if o.Schedules.IsNull() || o.Schedules.IsUnknown() {
		return nil, false
	}
	var v []Schedule
	d := o.Schedules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedules sets the value of the Schedules field in ListSchedulesResponse.
func (o *ListSchedulesResponse) SetSchedules(ctx context.Context, v []Schedule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schedules = types.ListValueMust(t, vs)
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
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

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsRequest) {
}

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (o ListSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListSubscriptionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`

	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsResponse) {
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(Subscription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSubscriptionsResponse
// only implements ToObjectValue() and Type().
func (o ListSubscriptionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"subscriptions":   o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"subscriptions":   basetypes.ListType{ElemType: basetypes.ListType{ElemType: Subscription{}.Type(ctx)}},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in ListSubscriptionsResponse as
// a slice of Subscription values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSubscriptionsResponse) GetSubscriptions(ctx context.Context) ([]Subscription, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []Subscription
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in ListSubscriptionsResponse.
func (o *ListSubscriptionsResponse) SetSubscriptions(ctx context.Context, v []Subscription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type MessageError struct {
	Error types.String `tfsdk:"error" tf:"optional"`

	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *MessageError) SyncEffectiveFieldsDuringCreateOrUpdate(plan MessageError) {
}

func (newState *MessageError) SyncEffectiveFieldsDuringRead(existingState MessageError) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MessageError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MessageError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MessageError
// only implements ToObjectValue() and Type().
func (o MessageError) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error": o.Error,
			"type":  o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MessageError) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"type":  types.StringType,
		},
	}
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId types.String `tfsdk:"source_dashboard_id" tf:""`
}

func (newState *MigrateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigrateDashboardRequest) {
}

func (newState *MigrateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState MigrateDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigrateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrateDashboardRequest
// only implements ToObjectValue() and Type().
func (o MigrateDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":        o.DisplayName,
			"parent_path":         o.ParentPath,
			"source_dashboard_id": o.SourceDashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":        types.StringType,
			"parent_path":         types.StringType,
			"source_dashboard_id": types.StringType,
		},
	}
}

type PublishRequest struct {
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

func (newState *PublishRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishRequest) {
}

func (newState *PublishRequest) SyncEffectiveFieldsDuringRead(existingState PublishRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishRequest
// only implements ToObjectValue() and Type().
func (o PublishRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":      o.DashboardId,
			"embed_credentials": o.EmbedCredentials,
			"warehouse_id":      o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":      types.StringType,
			"embed_credentials": types.BoolType,
			"warehouse_id":      types.StringType,
		},
	}
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"computed,optional"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time" tf:"computed,optional"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedDashboard) {
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringRead(existingState PublishedDashboard) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishedDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishedDashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishedDashboard
// only implements ToObjectValue() and Type().
func (o PublishedDashboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PublishedDashboard) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":         types.StringType,
			"embed_credentials":    types.BoolType,
			"revision_create_time": types.StringType,
			"warehouse_id":         types.StringType,
		},
	}
}

type QueryAttachment struct {
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

func (newState *QueryAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryAttachment) {
}

func (newState *QueryAttachment) SyncEffectiveFieldsDuringRead(existingState QueryAttachment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cached_query_schema": reflect.TypeOf(QuerySchema{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryAttachment
// only implements ToObjectValue() and Type().
func (o QueryAttachment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cached_query_schema":    basetypes.ListType{ElemType: QuerySchema{}.Type(ctx)},
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

// GetCachedQuerySchema returns the value of the CachedQuerySchema field in QueryAttachment as
// a QuerySchema value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryAttachment) GetCachedQuerySchema(ctx context.Context) (QuerySchema, bool) {
	var e QuerySchema
	if o.CachedQuerySchema.IsNull() || o.CachedQuerySchema.IsUnknown() {
		return e, false
	}
	var v []QuerySchema
	d := o.CachedQuerySchema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCachedQuerySchema sets the value of the CachedQuerySchema field in QueryAttachment.
func (o *QueryAttachment) SetCachedQuerySchema(ctx context.Context, v QuerySchema) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cached_query_schema"]
	o.CachedQuerySchema = types.ListValueMust(t, vs)
}

type QuerySchema struct {
	Columns types.List `tfsdk:"columns" tf:"optional"`
	// Used to determine if the stored query schema is compatible with the
	// latest run. The service should always clear the schema when the query is
	// re-executed.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

func (newState *QuerySchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuerySchema) {
}

func (newState *QuerySchema) SyncEffectiveFieldsDuringRead(existingState QuerySchema) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuerySchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuerySchema) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(QuerySchemaColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuerySchema
// only implements ToObjectValue() and Type().
func (o QuerySchema) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":      o.Columns,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QuerySchema) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns":      basetypes.ListType{ElemType: basetypes.ListType{ElemType: QuerySchemaColumn{}.Type(ctx)}},
			"statement_id": types.StringType,
		},
	}
}

// GetColumns returns the value of the Columns field in QuerySchema as
// a slice of QuerySchemaColumn values.
// If the field is unknown or null, the boolean return value is false.
func (o *QuerySchema) GetColumns(ctx context.Context) ([]QuerySchemaColumn, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []QuerySchemaColumn
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in QuerySchema.
func (o *QuerySchema) SetColumns(ctx context.Context, v []QuerySchemaColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type QuerySchemaColumn struct {
	// Populated from
	// https://docs.databricks.com/sql/language-manual/sql-ref-datatypes.html
	DataType types.String `tfsdk:"data_type" tf:""`

	Name types.String `tfsdk:"name" tf:""`
	// Corresponds to type desc
	TypeText types.String `tfsdk:"type_text" tf:""`
}

func (newState *QuerySchemaColumn) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuerySchemaColumn) {
}

func (newState *QuerySchemaColumn) SyncEffectiveFieldsDuringRead(existingState QuerySchemaColumn) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuerySchemaColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuerySchemaColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuerySchemaColumn
// only implements ToObjectValue() and Type().
func (o QuerySchemaColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_type": o.DataType,
			"name":      o.Name,
			"type_text": o.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QuerySchemaColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_type": types.StringType,
			"name":      types.StringType,
			"type_text": types.StringType,
		},
	}
}

type Result struct {
	// If result is truncated
	IsTruncated types.Bool `tfsdk:"is_truncated" tf:"optional"`
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

func (newState *Result) SyncEffectiveFieldsDuringCreateOrUpdate(plan Result) {
}

func (newState *Result) SyncEffectiveFieldsDuringRead(existingState Result) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Result.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Result) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Result
// only implements ToObjectValue() and Type().
func (o Result) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_truncated": o.IsTruncated,
			"row_count":    o.RowCount,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Result) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated": types.BoolType,
			"row_count":    types.Int64Type,
			"statement_id": types.StringType,
		},
	}
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule types.List `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed,optional"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Schedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan Schedule) {
}

func (newState *Schedule) SyncEffectiveFieldsDuringRead(existingState Schedule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Schedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Schedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron_schedule": reflect.TypeOf(CronSchedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Schedule
// only implements ToObjectValue() and Type().
func (o Schedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Schedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":   types.StringType,
			"cron_schedule": basetypes.ListType{ElemType: CronSchedule{}.Type(ctx)},
			"dashboard_id":  types.StringType,
			"display_name":  types.StringType,
			"etag":          types.StringType,
			"pause_status":  types.StringType,
			"schedule_id":   types.StringType,
			"update_time":   types.StringType,
			"warehouse_id":  types.StringType,
		},
	}
}

// GetCronSchedule returns the value of the CronSchedule field in Schedule as
// a CronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *Schedule) GetCronSchedule(ctx context.Context) (CronSchedule, bool) {
	var e CronSchedule
	if o.CronSchedule.IsNull() || o.CronSchedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule
	d := o.CronSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCronSchedule sets the value of the CronSchedule field in Schedule.
func (o *Schedule) SetCronSchedule(ctx context.Context, v CronSchedule) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cron_schedule"]
	o.CronSchedule = types.ListValueMust(t, vs)
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber types.List `tfsdk:"destination_subscriber" tf:"optional,object"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber types.List `tfsdk:"user_subscriber" tf:"optional,object"`
}

func (newState *Subscriber) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscriber) {
}

func (newState *Subscriber) SyncEffectiveFieldsDuringRead(existingState Subscriber) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscriber.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscriber) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destination_subscriber": reflect.TypeOf(SubscriptionSubscriberDestination{}),
		"user_subscriber":        reflect.TypeOf(SubscriptionSubscriberUser{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Subscriber
// only implements ToObjectValue() and Type().
func (o Subscriber) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_subscriber": o.DestinationSubscriber,
			"user_subscriber":        o.UserSubscriber,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Subscriber) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_subscriber": basetypes.ListType{ElemType: SubscriptionSubscriberDestination{}.Type(ctx)},
			"user_subscriber":        basetypes.ListType{ElemType: SubscriptionSubscriberUser{}.Type(ctx)},
		},
	}
}

// GetDestinationSubscriber returns the value of the DestinationSubscriber field in Subscriber as
// a SubscriptionSubscriberDestination value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscriber) GetDestinationSubscriber(ctx context.Context) (SubscriptionSubscriberDestination, bool) {
	var e SubscriptionSubscriberDestination
	if o.DestinationSubscriber.IsNull() || o.DestinationSubscriber.IsUnknown() {
		return e, false
	}
	var v []SubscriptionSubscriberDestination
	d := o.DestinationSubscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDestinationSubscriber sets the value of the DestinationSubscriber field in Subscriber.
func (o *Subscriber) SetDestinationSubscriber(ctx context.Context, v SubscriptionSubscriberDestination) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["destination_subscriber"]
	o.DestinationSubscriber = types.ListValueMust(t, vs)
}

// GetUserSubscriber returns the value of the UserSubscriber field in Subscriber as
// a SubscriptionSubscriberUser value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscriber) GetUserSubscriber(ctx context.Context) (SubscriptionSubscriberUser, bool) {
	var e SubscriptionSubscriberUser
	if o.UserSubscriber.IsNull() || o.UserSubscriber.IsUnknown() {
		return e, false
	}
	var v []SubscriptionSubscriberUser
	d := o.UserSubscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUserSubscriber sets the value of the UserSubscriber field in Subscriber.
func (o *Subscriber) SetUserSubscriber(ctx context.Context, v SubscriptionSubscriberUser) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_subscriber"]
	o.UserSubscriber = types.ListValueMust(t, vs)
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id" tf:"computed,optional"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed,optional"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber types.List `tfsdk:"subscriber" tf:"object"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"computed,optional"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
}

func (newState *Subscription) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscription) {
}

func (newState *Subscription) SyncEffectiveFieldsDuringRead(existingState Subscription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriber": reflect.TypeOf(Subscriber{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Subscription
// only implements ToObjectValue() and Type().
func (o Subscription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Subscription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"created_by_user_id": types.Int64Type,
			"dashboard_id":       types.StringType,
			"etag":               types.StringType,
			"schedule_id":        types.StringType,
			"subscriber":         basetypes.ListType{ElemType: Subscriber{}.Type(ctx)},
			"subscription_id":    types.StringType,
			"update_time":        types.StringType,
		},
	}
}

// GetSubscriber returns the value of the Subscriber field in Subscription as
// a Subscriber value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscription) GetSubscriber(ctx context.Context) (Subscriber, bool) {
	var e Subscriber
	if o.Subscriber.IsNull() || o.Subscriber.IsUnknown() {
		return e, false
	}
	var v []Subscriber
	d := o.Subscriber.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSubscriber sets the value of the Subscriber field in Subscription.
func (o *Subscription) SetSubscriber(ctx context.Context, v Subscriber) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriber"]
	o.Subscriber = types.ListValueMust(t, vs)
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberDestination) {
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberDestination) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubscriptionSubscriberDestination
// only implements ToObjectValue() and Type().
func (o SubscriptionSubscriberDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": o.DestinationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
		},
	}
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId types.Int64 `tfsdk:"user_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberUser) {
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberUser) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberUser.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberUser) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubscriptionSubscriberUser
// only implements ToObjectValue() and Type().
func (o SubscriptionSubscriberUser) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.Int64Type,
		},
	}
}

type TextAttachment struct {
	// AI generated message
	Content types.String `tfsdk:"content" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *TextAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextAttachment) {
}

func (newState *TextAttachment) SyncEffectiveFieldsDuringRead(existingState TextAttachment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TextAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TextAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TextAttachment
// only implements ToObjectValue() and Type().
func (o TextAttachment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": o.Content,
			"id":      o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TextAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"id":      types.StringType,
		},
	}
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *TrashDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardRequest) {
}

func (newState *TrashDashboardRequest) SyncEffectiveFieldsDuringRead(existingState TrashDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashDashboardRequest
// only implements ToObjectValue() and Type().
func (o TrashDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type TrashDashboardResponse struct {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardResponse) {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringRead(existingState TrashDashboardResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashDashboardResponse
// only implements ToObjectValue() and Type().
func (o TrashDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardRequest) {
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpublishDashboardRequest
// only implements ToObjectValue() and Type().
func (o UnpublishDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type UnpublishDashboardResponse struct {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardResponse) {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnpublishDashboardResponse
// only implements ToObjectValue() and Type().
func (o UnpublishDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Update dashboard
type UpdateDashboardRequest struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDashboardRequest) {
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDashboardRequest
// only implements ToObjectValue() and Type().
func (o UpdateDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard":    o.Dashboard,
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard":    basetypes.ListType{ElemType: Dashboard{}.Type(ctx)},
			"dashboard_id": types.StringType,
		},
	}
}

// GetDashboard returns the value of the Dashboard field in UpdateDashboardRequest as
// a Dashboard value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDashboardRequest) GetDashboard(ctx context.Context) (Dashboard, bool) {
	var e Dashboard
	if o.Dashboard.IsNull() || o.Dashboard.IsUnknown() {
		return e, false
	}
	var v []Dashboard
	d := o.Dashboard.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboard sets the value of the Dashboard field in UpdateDashboardRequest.
func (o *UpdateDashboardRequest) SetDashboard(ctx context.Context, v Dashboard) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard"]
	o.Dashboard = types.ListValueMust(t, vs)
}

// Update dashboard schedule
type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateScheduleRequest) {
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState UpdateScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateScheduleRequest
// only implements ToObjectValue() and Type().
func (o UpdateScheduleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"schedule":     o.Schedule,
			"schedule_id":  o.ScheduleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule":     basetypes.ListType{ElemType: Schedule{}.Type(ctx)},
			"schedule_id":  types.StringType,
		},
	}
}

// GetSchedule returns the value of the Schedule field in UpdateScheduleRequest as
// a Schedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateScheduleRequest) GetSchedule(ctx context.Context) (Schedule, bool) {
	var e Schedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []Schedule
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in UpdateScheduleRequest.
func (o *UpdateScheduleRequest) SetSchedule(ctx context.Context, v Schedule) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}
