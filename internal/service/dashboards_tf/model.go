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
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sql"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateScheduleRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateScheduleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateSubscriptionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateSubscriptionRequest) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CronSchedule{}

// Equal implements basetypes.ObjectValuable.
func (o CronSchedule) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CronSchedule) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CronSchedule) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CronSchedule) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CronSchedule) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CronSchedule) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Dashboard{}

// Equal implements basetypes.ObjectValuable.
func (o Dashboard) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Dashboard) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Dashboard) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Dashboard) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Dashboard) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Dashboard) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteScheduleRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteScheduleRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteScheduleResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteScheduleResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteSubscriptionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteSubscriptionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteSubscriptionResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteSubscriptionResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieAttachment{}

// Equal implements basetypes.ObjectValuable.
func (o GenieAttachment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieAttachment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieAttachment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieAttachment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieAttachment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieAttachment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GenieAttachment) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieConversation{}

// Equal implements basetypes.ObjectValuable.
func (o GenieConversation) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieConversation) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieConversation) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieConversation) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieConversation) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieConversation) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieCreateConversationMessageRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieCreateConversationMessageRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieExecuteMessageQueryRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieExecuteMessageQueryRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieGetConversationMessageRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieGetConversationMessageRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieGetMessageQueryResultRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	StatementResponse sql.StatementResponse `tfsdk:"statement_response" tf:"optional,object"`
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
		"statement_response": reflect.TypeOf(sql.StatementResponse{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieGetMessageQueryResultResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageQueryResultResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_response": basetypes.ListType{
				ElemType: sql_tf.StatementResponse{}.Type(ctx),
			},
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieMessage{}

// Equal implements basetypes.ObjectValuable.
func (o GenieMessage) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieMessage) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieMessage) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieMessage) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieMessage) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieMessage) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GenieMessage) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieStartConversationMessageRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieStartConversationMessageRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GenieStartConversationResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GenieStartConversationResponse) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPublishedDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPublishedDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetScheduleRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetScheduleRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetScheduleRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetScheduleRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetScheduleRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetScheduleRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetScheduleRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetSubscriptionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetSubscriptionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListDashboardsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListDashboardsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboards": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSchedulesRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSchedulesRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSchedulesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListSchedulesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schedules": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSubscriptionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSubscriptionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSubscriptionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListSubscriptionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"subscriptions": basetypes.ListType{
				ElemType: Subscription{}.Type(ctx),
			},
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = MessageError{}

// Equal implements basetypes.ObjectValuable.
func (o MessageError) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o MessageError) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o MessageError) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o MessageError) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o MessageError) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o MessageError) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = MigrateDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PublishRequest{}

// Equal implements basetypes.ObjectValuable.
func (o PublishRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PublishRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PublishRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PublishRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PublishRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PublishRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PublishedDashboard{}

// Equal implements basetypes.ObjectValuable.
func (o PublishedDashboard) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PublishedDashboard) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PublishedDashboard) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PublishedDashboard) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PublishedDashboard) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PublishedDashboard) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = QueryAttachment{}

// Equal implements basetypes.ObjectValuable.
func (o QueryAttachment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o QueryAttachment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o QueryAttachment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o QueryAttachment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o QueryAttachment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o QueryAttachment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o QueryAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Result{}

// Equal implements basetypes.ObjectValuable.
func (o Result) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Result) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Result) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Result) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Result) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Result) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Schedule{}

// Equal implements basetypes.ObjectValuable.
func (o Schedule) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Schedule) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Schedule) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Schedule) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Schedule) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Schedule) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Schedule) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Subscriber{}

// Equal implements basetypes.ObjectValuable.
func (o Subscriber) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Subscriber) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Subscriber) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Subscriber) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Subscriber) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Subscriber) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Subscriber) Type(ctx context.Context) attr.Type {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Subscription{}

// Equal implements basetypes.ObjectValuable.
func (o Subscription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Subscription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Subscription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Subscription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Subscription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Subscription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
			"subscriber": basetypes.ListType{
				ElemType: Subscriber{}.Type(ctx),
			},
			"subscription_id": types.StringType,
			"update_time":     types.StringType,
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SubscriptionSubscriberDestination{}

// Equal implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberDestination) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SubscriptionSubscriberUser{}

// Equal implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SubscriptionSubscriberUser) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TextAttachment{}

// Equal implements basetypes.ObjectValuable.
func (o TextAttachment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TextAttachment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TextAttachment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TextAttachment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TextAttachment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TextAttachment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TrashDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TrashDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TrashDashboardResponse{}

// Equal implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TrashDashboardResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UnpublishDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UnpublishDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UnpublishDashboardResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UnpublishDashboardResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateDashboardRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
			"dashboard_id": types.StringType,
		},
	}
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateScheduleRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateScheduleRequest) Type(ctx context.Context) attr.Type {
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

// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart context
// step to determine relevant context. * `ASKING_AI`: Waiting for the LLM to
// respond to the users question. * `EXECUTING_QUERY`: Executing AI provided SQL
// query. Get the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
// **Important: The message status will stay in the `EXECUTING_QUERY` until a
// client calls [getMessageQueryResult](:method:genie/getMessageQueryResult)**.
// * `FAILED`: Generating a response or the executing the query failed. Please
// see `error` field. * `COMPLETED`: Message processing is completed. Results
// are in the `attachments` field. Get the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API. *
// `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL result
// is not available anymore. The user needs to execute the query again. *
// `CANCELLED`: Message has been cancelled.
