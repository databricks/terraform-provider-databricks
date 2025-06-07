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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AuthorizationDetails_SdkV2 struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	GrantRules types.List `tfsdk:"grant_rules"`
	// The acl path of the tree store resource resource.
	ResourceLegacyAclPath types.String `tfsdk:"resource_legacy_acl_path"`
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	ResourceName types.String `tfsdk:"resource_name"`
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	Type_ types.String `tfsdk:"type"`
}

func (newState *AuthorizationDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AuthorizationDetails_SdkV2) {
}

func (newState *AuthorizationDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState AuthorizationDetails_SdkV2) {
}

func (c AuthorizationDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["grant_rules"] = attrs["grant_rules"].SetOptional()
	attrs["resource_legacy_acl_path"] = attrs["resource_legacy_acl_path"].SetOptional()
	attrs["resource_name"] = attrs["resource_name"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AuthorizationDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AuthorizationDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(AuthorizationDetailsGrantRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AuthorizationDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o AuthorizationDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"grant_rules":              o.GrantRules,
			"resource_legacy_acl_path": o.ResourceLegacyAclPath,
			"resource_name":            o.ResourceName,
			"type":                     o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AuthorizationDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"grant_rules": basetypes.ListType{
				ElemType: AuthorizationDetailsGrantRule_SdkV2{}.Type(ctx),
			},
			"resource_legacy_acl_path": types.StringType,
			"resource_name":            types.StringType,
			"type":                     types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in AuthorizationDetails_SdkV2 as
// a slice of AuthorizationDetailsGrantRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AuthorizationDetails_SdkV2) GetGrantRules(ctx context.Context) ([]AuthorizationDetailsGrantRule_SdkV2, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []AuthorizationDetailsGrantRule_SdkV2
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in AuthorizationDetails_SdkV2.
func (o *AuthorizationDetails_SdkV2) SetGrantRules(ctx context.Context, v []AuthorizationDetailsGrantRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type AuthorizationDetailsGrantRule_SdkV2 struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	PermissionSet types.String `tfsdk:"permission_set"`
}

func (newState *AuthorizationDetailsGrantRule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AuthorizationDetailsGrantRule_SdkV2) {
}

func (newState *AuthorizationDetailsGrantRule_SdkV2) SyncEffectiveFieldsDuringRead(existingState AuthorizationDetailsGrantRule_SdkV2) {
}

func (c AuthorizationDetailsGrantRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_set"] = attrs["permission_set"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AuthorizationDetailsGrantRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AuthorizationDetailsGrantRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AuthorizationDetailsGrantRule_SdkV2
// only implements ToObjectValue() and Type().
func (o AuthorizationDetailsGrantRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": o.PermissionSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AuthorizationDetailsGrantRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
		},
	}
}

// Create dashboard
type CreateDashboardRequest_SdkV2 struct {
	Dashboard types.List `tfsdk:"dashboard"`
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
				ElemType: Dashboard_SdkV2{}.Type(ctx),
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

	Schedule types.List `tfsdk:"schedule"`
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
				ElemType: Schedule_SdkV2{}.Type(ctx),
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

	Subscription types.List `tfsdk:"subscription"`
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
				ElemType: Subscription_SdkV2{}.Type(ctx),
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
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression"`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule_SdkV2) {
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState CronSchedule_SdkV2) {
}

func (c CronSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	CreateTime types.String `tfsdk:"create_time"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag types.String `tfsdk:"etag"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath types.String `tfsdk:"parent_path"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path types.String `tfsdk:"path"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard types.String `tfsdk:"serialized_dashboard"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime types.String `tfsdk:"update_time"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard_SdkV2) {
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState Dashboard_SdkV2) {
}

func (c Dashboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["path"] = attrs["path"].SetComputed()
	attrs["serialized_dashboard"] = attrs["serialized_dashboard"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
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
	// Attachment ID
	AttachmentId types.String `tfsdk:"attachment_id"`
	// Query Attachment if Genie responds with a SQL query
	Query types.List `tfsdk:"query"`
	// Text Attachment if Genie responds with text
	Text types.List `tfsdk:"text"`
}

func (newState *GenieAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieAttachment_SdkV2) {
}

func (newState *GenieAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieAttachment_SdkV2) {
}

func (c GenieAttachment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attachment_id"] = attrs["attachment_id"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query"] = attrs["query"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["text"] = attrs["text"].SetOptional()
	attrs["text"] = attrs["text"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
		"query": reflect.TypeOf(GenieQueryAttachment_SdkV2{}),
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
			"attachment_id": o.AttachmentId,
			"query":         o.Query,
			"text":          o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieAttachment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id": types.StringType,
			"query": basetypes.ListType{
				ElemType: GenieQueryAttachment_SdkV2{}.Type(ctx),
			},
			"text": basetypes.ListType{
				ElemType: TextAttachment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQuery returns the value of the Query field in GenieAttachment_SdkV2 as
// a GenieQueryAttachment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment_SdkV2) GetQuery(ctx context.Context) (GenieQueryAttachment_SdkV2, bool) {
	var e GenieQueryAttachment_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []GenieQueryAttachment_SdkV2
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
func (o *GenieAttachment_SdkV2) SetQuery(ctx context.Context, v GenieQueryAttachment_SdkV2) {
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
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`
	// Conversation ID. Legacy identifier, use conversation_id instead
	Id types.String `tfsdk:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id"`
	// Conversation title
	Title types.String `tfsdk:"title"`
	// ID of the user who created the conversation
	UserId types.Int64 `tfsdk:"user_id"`
}

func (newState *GenieConversation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieConversation_SdkV2) {
}

func (newState *GenieConversation_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieConversation_SdkV2) {
}

func (c GenieConversation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["space_id"] = attrs["space_id"].SetRequired()
	attrs["title"] = attrs["title"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
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
			"conversation_id":        o.ConversationId,
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
			"conversation_id":        types.StringType,
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
	Content types.String `tfsdk:"content"`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieCreateConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieCreateConversationMessageRequest_SdkV2) {
}

func (newState *GenieCreateConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieCreateConversationMessageRequest_SdkV2) {
}

func (c GenieCreateConversationMessageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["space_id"] = attrs["space_id"].SetRequired()

	return attrs
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

// Execute message attachment SQL query
type GenieExecuteMessageAttachmentQueryRequest_SdkV2 struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"-"`
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieExecuteMessageAttachmentQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieExecuteMessageAttachmentQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieExecuteMessageAttachmentQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieExecuteMessageAttachmentQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachment_id":   o.AttachmentId,
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieExecuteMessageAttachmentQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// [Deprecated] Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest_SdkV2 struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
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

// Generate full query result download
type GenieGenerateDownloadFullQueryResultRequest_SdkV2 struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"-"`
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGenerateDownloadFullQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGenerateDownloadFullQueryResultRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGenerateDownloadFullQueryResultRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGenerateDownloadFullQueryResultRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachment_id":   o.AttachmentId,
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGenerateDownloadFullQueryResultRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGenerateDownloadFullQueryResultResponse_SdkV2 struct {
	// Download ID. Use this ID to track the download request in subsequent
	// polling calls
	DownloadId types.String `tfsdk:"download_id"`
}

func (newState *GenieGenerateDownloadFullQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGenerateDownloadFullQueryResultResponse_SdkV2) {
}

func (newState *GenieGenerateDownloadFullQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGenerateDownloadFullQueryResultResponse_SdkV2) {
}

func (c GenieGenerateDownloadFullQueryResultResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["download_id"] = attrs["download_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGenerateDownloadFullQueryResultResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGenerateDownloadFullQueryResultResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGenerateDownloadFullQueryResultResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGenerateDownloadFullQueryResultResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"download_id": o.DownloadId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGenerateDownloadFullQueryResultResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"download_id": types.StringType,
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

// Get download full query result
type GenieGetDownloadFullQueryResultRequest_SdkV2 struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"-"`
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	DownloadId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetDownloadFullQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetDownloadFullQueryResultRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetDownloadFullQueryResultRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetDownloadFullQueryResultRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachment_id":   o.AttachmentId,
			"conversation_id": o.ConversationId,
			"download_id":     o.DownloadId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetDownloadFullQueryResultRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"download_id":     types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGetDownloadFullQueryResultResponse_SdkV2 struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse types.List `tfsdk:"statement_response"`
}

func (newState *GenieGetDownloadFullQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetDownloadFullQueryResultResponse_SdkV2) {
}

func (newState *GenieGetDownloadFullQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGetDownloadFullQueryResultResponse_SdkV2) {
}

func (c GenieGetDownloadFullQueryResultResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_response"] = attrs["statement_response"].SetOptional()
	attrs["statement_response"] = attrs["statement_response"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetDownloadFullQueryResultResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetDownloadFullQueryResultResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statement_response": reflect.TypeOf(sql_tf.StatementResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetDownloadFullQueryResultResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetDownloadFullQueryResultResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_response": o.StatementResponse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetDownloadFullQueryResultResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_response": basetypes.ListType{
				ElemType: sql_tf.StatementResponse_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStatementResponse returns the value of the StatementResponse field in GenieGetDownloadFullQueryResultResponse_SdkV2 as
// a sql_tf.StatementResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieGetDownloadFullQueryResultResponse_SdkV2) GetStatementResponse(ctx context.Context) (sql_tf.StatementResponse_SdkV2, bool) {
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

// SetStatementResponse sets the value of the StatementResponse field in GenieGetDownloadFullQueryResultResponse_SdkV2.
func (o *GenieGetDownloadFullQueryResultResponse_SdkV2) SetStatementResponse(ctx context.Context, v sql_tf.StatementResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_response"]
	o.StatementResponse = types.ListValueMust(t, vs)
}

// Get message attachment SQL query result
type GenieGetMessageAttachmentQueryResultRequest_SdkV2 struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"-"`
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageAttachmentQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageAttachmentQueryResultRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageAttachmentQueryResultRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetMessageAttachmentQueryResultRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachment_id":   o.AttachmentId,
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetMessageAttachmentQueryResultRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// [Deprecated] Get conversation message SQL query result
type GenieGetMessageQueryResultRequest_SdkV2 struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
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
	StatementResponse types.List `tfsdk:"statement_response"`
}

func (newState *GenieGetMessageQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultResponse_SdkV2) {
}

func (newState *GenieGetMessageQueryResultResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultResponse_SdkV2) {
}

func (c GenieGetMessageQueryResultResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_response"] = attrs["statement_response"].SetOptional()
	attrs["statement_response"] = attrs["statement_response"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: sql_tf.StatementResponse_SdkV2{}.Type(ctx),
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

// [Deprecated] Get conversation message SQL query result
type GenieGetQueryResultByAttachmentRequest_SdkV2 struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"-"`
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetQueryResultByAttachmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetQueryResultByAttachmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetQueryResultByAttachmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetQueryResultByAttachmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attachment_id":   o.AttachmentId,
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetQueryResultByAttachmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Get Genie Space
type GenieGetSpaceRequest_SdkV2 struct {
	// The ID associated with the Genie space
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetSpaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetSpaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieGetSpaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"space_id": o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetSpaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"space_id": types.StringType,
		},
	}
}

// List Genie spaces
type GenieListSpacesRequest_SdkV2 struct {
	// Maximum number of spaces to return per page
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token for getting the next page of results
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListSpacesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListSpacesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListSpacesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieListSpacesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListSpacesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GenieListSpacesResponse_SdkV2 struct {
	// Token to get the next page of results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Genie spaces
	Spaces types.List `tfsdk:"spaces"`
}

func (newState *GenieListSpacesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieListSpacesResponse_SdkV2) {
}

func (newState *GenieListSpacesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieListSpacesResponse_SdkV2) {
}

func (c GenieListSpacesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["spaces"] = attrs["spaces"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListSpacesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListSpacesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spaces": reflect.TypeOf(GenieSpace_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListSpacesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieListSpacesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"spaces":          o.Spaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListSpacesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"spaces": basetypes.ListType{
				ElemType: GenieSpace_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSpaces returns the value of the Spaces field in GenieListSpacesResponse_SdkV2 as
// a slice of GenieSpace_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieListSpacesResponse_SdkV2) GetSpaces(ctx context.Context) ([]GenieSpace_SdkV2, bool) {
	if o.Spaces.IsNull() || o.Spaces.IsUnknown() {
		return nil, false
	}
	var v []GenieSpace_SdkV2
	d := o.Spaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpaces sets the value of the Spaces field in GenieListSpacesResponse_SdkV2.
func (o *GenieListSpacesResponse_SdkV2) SetSpaces(ctx context.Context, v []GenieSpace_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Spaces = types.ListValueMust(t, vs)
}

type GenieMessage_SdkV2 struct {
	// AI-generated response to the message
	Attachments types.List `tfsdk:"attachments"`
	// User message content
	Content types.String `tfsdk:"content"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`
	// Error message if Genie failed to respond to the message
	Error types.List `tfsdk:"error"`
	// Message ID. Legacy identifier, use message_id instead
	Id types.String `tfsdk:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Message ID
	MessageId types.String `tfsdk:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult types.List `tfsdk:"query_result"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id"`
	// MessageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart
	// context step to determine relevant context. * `ASKING_AI`: Waiting for
	// the LLM to respond to the user's question. * `PENDING_WAREHOUSE`: Waiting
	// for warehouse before the SQL query can start executing. *
	// `EXECUTING_QUERY`: Executing a generated SQL query. Get the SQL query
	// result by calling
	// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
	// API. * `FAILED`: The response generation or query execution failed. See
	// `error` field. * `COMPLETED`: Message processing is completed. Results
	// are in the `attachments` field. Get the SQL query result by calling
	// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
	// API. * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`:
	// SQL result is not available anymore. The user needs to rerun the query.
	// Rerun the SQL query result by calling
	// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
	// API. * `CANCELLED`: Message has been cancelled.
	Status types.String `tfsdk:"status"`
	// ID of the user who created the message
	UserId types.Int64 `tfsdk:"user_id"`
}

func (newState *GenieMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieMessage_SdkV2) {
}

func (newState *GenieMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieMessage_SdkV2) {
}

func (c GenieMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attachments"] = attrs["attachments"].SetOptional()
	attrs["content"] = attrs["content"].SetRequired()
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["message_id"] = attrs["message_id"].SetRequired()
	attrs["query_result"] = attrs["query_result"].SetOptional()
	attrs["query_result"] = attrs["query_result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["space_id"] = attrs["space_id"].SetRequired()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
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
			"message_id":             o.MessageId,
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
				ElemType: GenieAttachment_SdkV2{}.Type(ctx),
			},
			"content":           types.StringType,
			"conversation_id":   types.StringType,
			"created_timestamp": types.Int64Type,
			"error": basetypes.ListType{
				ElemType: MessageError_SdkV2{}.Type(ctx),
			},
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"message_id":             types.StringType,
			"query_result": basetypes.ListType{
				ElemType: Result_SdkV2{}.Type(ctx),
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

type GenieQueryAttachment_SdkV2 struct {
	// Description of the query
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
	// Time when the user updated the query last
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// AI generated SQL query
	Query types.String `tfsdk:"query"`
	// Metadata associated with the query result.
	QueryResultMetadata types.List `tfsdk:"query_result_metadata"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id"`
	// Name of the query
	Title types.String `tfsdk:"title"`
}

func (newState *GenieQueryAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieQueryAttachment_SdkV2) {
}

func (newState *GenieQueryAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieQueryAttachment_SdkV2) {
}

func (c GenieQueryAttachment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query_result_metadata"] = attrs["query_result_metadata"].SetOptional()
	attrs["query_result_metadata"] = attrs["query_result_metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["title"] = attrs["title"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieQueryAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieQueryAttachment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_result_metadata": reflect.TypeOf(GenieResultMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieQueryAttachment_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieQueryAttachment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":            o.Description,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"query":                  o.Query,
			"query_result_metadata":  o.QueryResultMetadata,
			"statement_id":           o.StatementId,
			"title":                  o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieQueryAttachment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query":                  types.StringType,
			"query_result_metadata": basetypes.ListType{
				ElemType: GenieResultMetadata_SdkV2{}.Type(ctx),
			},
			"statement_id": types.StringType,
			"title":        types.StringType,
		},
	}
}

// GetQueryResultMetadata returns the value of the QueryResultMetadata field in GenieQueryAttachment_SdkV2 as
// a GenieResultMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieQueryAttachment_SdkV2) GetQueryResultMetadata(ctx context.Context) (GenieResultMetadata_SdkV2, bool) {
	var e GenieResultMetadata_SdkV2
	if o.QueryResultMetadata.IsNull() || o.QueryResultMetadata.IsUnknown() {
		return e, false
	}
	var v []GenieResultMetadata_SdkV2
	d := o.QueryResultMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryResultMetadata sets the value of the QueryResultMetadata field in GenieQueryAttachment_SdkV2.
func (o *GenieQueryAttachment_SdkV2) SetQueryResultMetadata(ctx context.Context, v GenieResultMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_result_metadata"]
	o.QueryResultMetadata = types.ListValueMust(t, vs)
}

type GenieResultMetadata_SdkV2 struct {
	// Indicates whether the result set is truncated.
	IsTruncated types.Bool `tfsdk:"is_truncated"`
	// The number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (newState *GenieResultMetadata_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieResultMetadata_SdkV2) {
}

func (newState *GenieResultMetadata_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieResultMetadata_SdkV2) {
}

func (c GenieResultMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_truncated"] = attrs["is_truncated"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieResultMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieResultMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieResultMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieResultMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_truncated": o.IsTruncated,
			"row_count":    o.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieResultMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated": types.BoolType,
			"row_count":    types.Int64Type,
		},
	}
}

type GenieSpace_SdkV2 struct {
	// Description of the Genie Space
	Description types.String `tfsdk:"description"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id"`
	// Title of the Genie Space
	Title types.String `tfsdk:"title"`
}

func (newState *GenieSpace_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieSpace_SdkV2) {
}

func (newState *GenieSpace_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieSpace_SdkV2) {
}

func (c GenieSpace_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["space_id"] = attrs["space_id"].SetRequired()
	attrs["title"] = attrs["title"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieSpace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieSpace_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieSpace_SdkV2
// only implements ToObjectValue() and Type().
func (o GenieSpace_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"space_id":    o.SpaceId,
			"title":       o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieSpace_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"space_id":    types.StringType,
			"title":       types.StringType,
		},
	}
}

type GenieStartConversationMessageRequest_SdkV2 struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieStartConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationMessageRequest_SdkV2) {
}

func (newState *GenieStartConversationMessageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationMessageRequest_SdkV2) {
}

func (c GenieStartConversationMessageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["space_id"] = attrs["space_id"].SetRequired()

	return attrs
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
	Conversation types.List `tfsdk:"conversation"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id"`

	Message types.List `tfsdk:"message"`
	// Message ID
	MessageId types.String `tfsdk:"message_id"`
}

func (newState *GenieStartConversationResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationResponse_SdkV2) {
}

func (newState *GenieStartConversationResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationResponse_SdkV2) {
}

func (c GenieStartConversationResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["conversation"] = attrs["conversation"].SetOptional()
	attrs["conversation"] = attrs["conversation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["message"] = attrs["message"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["message_id"] = attrs["message_id"].SetRequired()

	return attrs
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
				ElemType: GenieConversation_SdkV2{}.Type(ctx),
			},
			"conversation_id": types.StringType,
			"message": basetypes.ListType{
				ElemType: GenieMessage_SdkV2{}.Type(ctx),
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

// Read an information of a published dashboard to mint an OAuth token.
type GetPublishedDashboardTokenInfoRequest_SdkV2 struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
	// Provided external value to be included in the custom claim.
	ExternalValue types.String `tfsdk:"-"`
	// Provided external viewer id to be included in the custom claim.
	ExternalViewerId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedDashboardTokenInfoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedDashboardTokenInfoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardTokenInfoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardTokenInfoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":       o.DashboardId,
			"external_value":     o.ExternalValue,
			"external_viewer_id": o.ExternalViewerId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardTokenInfoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":       types.StringType,
			"external_value":     types.StringType,
			"external_viewer_id": types.StringType,
		},
	}
}

type GetPublishedDashboardTokenInfoResponse_SdkV2 struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	AuthorizationDetails types.List `tfsdk:"authorization_details"`
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	CustomClaim types.String `tfsdk:"custom_claim"`
	// Scope defining access permissions.
	Scope types.String `tfsdk:"scope"`
}

func (newState *GetPublishedDashboardTokenInfoResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedDashboardTokenInfoResponse_SdkV2) {
}

func (newState *GetPublishedDashboardTokenInfoResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPublishedDashboardTokenInfoResponse_SdkV2) {
}

func (c GetPublishedDashboardTokenInfoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authorization_details"] = attrs["authorization_details"].SetOptional()
	attrs["custom_claim"] = attrs["custom_claim"].SetOptional()
	attrs["scope"] = attrs["scope"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedDashboardTokenInfoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedDashboardTokenInfoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"authorization_details": reflect.TypeOf(AuthorizationDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardTokenInfoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardTokenInfoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization_details": o.AuthorizationDetails,
			"custom_claim":          o.CustomClaim,
			"scope":                 o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardTokenInfoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization_details": basetypes.ListType{
				ElemType: AuthorizationDetails_SdkV2{}.Type(ctx),
			},
			"custom_claim": types.StringType,
			"scope":        types.StringType,
		},
	}
}

// GetAuthorizationDetails returns the value of the AuthorizationDetails field in GetPublishedDashboardTokenInfoResponse_SdkV2 as
// a slice of AuthorizationDetails_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPublishedDashboardTokenInfoResponse_SdkV2) GetAuthorizationDetails(ctx context.Context) ([]AuthorizationDetails_SdkV2, bool) {
	if o.AuthorizationDetails.IsNull() || o.AuthorizationDetails.IsUnknown() {
		return nil, false
	}
	var v []AuthorizationDetails_SdkV2
	d := o.AuthorizationDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAuthorizationDetails sets the value of the AuthorizationDetails field in GetPublishedDashboardTokenInfoResponse_SdkV2.
func (o *GetPublishedDashboardTokenInfoResponse_SdkV2) SetAuthorizationDetails(ctx context.Context, v []AuthorizationDetails_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["authorization_details"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AuthorizationDetails = types.ListValueMust(t, vs)
}

// Get dashboard schedule
type GetScheduleRequest_SdkV2 struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
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
	Dashboards types.List `tfsdk:"dashboards"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListDashboardsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsResponse_SdkV2) {
}

func (newState *ListDashboardsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListDashboardsResponse_SdkV2) {
}

func (c ListDashboardsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboards"] = attrs["dashboards"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
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
				ElemType: Dashboard_SdkV2{}.Type(ctx),
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Schedules types.List `tfsdk:"schedules"`
}

func (newState *ListSchedulesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesResponse_SdkV2) {
}

func (newState *ListSchedulesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSchedulesResponse_SdkV2) {
}

func (c ListSchedulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["schedules"] = attrs["schedules"].SetOptional()

	return attrs
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
				ElemType: Schedule_SdkV2{}.Type(ctx),
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Subscriptions types.List `tfsdk:"subscriptions"`
}

func (newState *ListSubscriptionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsResponse_SdkV2) {
}

func (newState *ListSubscriptionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsResponse_SdkV2) {
}

func (c ListSubscriptionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["subscriptions"] = attrs["subscriptions"].SetOptional()

	return attrs
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
				ElemType: Subscription_SdkV2{}.Type(ctx),
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
	Error types.String `tfsdk:"error"`

	Type_ types.String `tfsdk:"type"`
}

func (newState *MessageError_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MessageError_SdkV2) {
}

func (newState *MessageError_SdkV2) SyncEffectiveFieldsDuringRead(existingState MessageError_SdkV2) {
}

func (c MessageError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
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
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath types.String `tfsdk:"parent_path"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId types.String `tfsdk:"source_dashboard_id"`
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	UpdateParameterSyntax types.Bool `tfsdk:"update_parameter_syntax"`
}

func (newState *MigrateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigrateDashboardRequest_SdkV2) {
}

func (newState *MigrateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState MigrateDashboardRequest_SdkV2) {
}

func (c MigrateDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["source_dashboard_id"] = attrs["source_dashboard_id"].SetRequired()
	attrs["update_parameter_syntax"] = attrs["update_parameter_syntax"].SetOptional()

	return attrs
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
			"display_name":            o.DisplayName,
			"parent_path":             o.ParentPath,
			"source_dashboard_id":     o.SourceDashboardId,
			"update_parameter_syntax": o.UpdateParameterSyntax,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":            types.StringType,
			"parent_path":             types.StringType,
			"source_dashboard_id":     types.StringType,
			"update_parameter_syntax": types.BoolType,
		},
	}
}

type PublishRequest_SdkV2 struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *PublishRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishRequest_SdkV2) {
}

func (newState *PublishRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PublishRequest_SdkV2) {
}

func (c PublishRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["embed_credentials"] = attrs["embed_credentials"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
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
	DisplayName types.String `tfsdk:"display_name"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *PublishedDashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedDashboard_SdkV2) {
}

func (newState *PublishedDashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState PublishedDashboard_SdkV2) {
}

func (c PublishedDashboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["embed_credentials"] = attrs["embed_credentials"].SetOptional()
	attrs["revision_create_time"] = attrs["revision_create_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
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

type Result_SdkV2 struct {
	// If result is truncated
	IsTruncated types.Bool `tfsdk:"is_truncated"`
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id"`
}

func (newState *Result_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Result_SdkV2) {
}

func (newState *Result_SdkV2) SyncEffectiveFieldsDuringRead(existingState Result_SdkV2) {
}

func (c Result_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_truncated"] = attrs["is_truncated"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["statement_id"] = attrs["statement_id"].SetOptional()

	return attrs
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
	CreateTime types.String `tfsdk:"create_time"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule types.List `tfsdk:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *Schedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Schedule_SdkV2) {
}

func (newState *Schedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState Schedule_SdkV2) {
}

func (c Schedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["cron_schedule"] = attrs["cron_schedule"].SetRequired()
	attrs["cron_schedule"] = attrs["cron_schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["pause_status"] = attrs["pause_status"].SetOptional()
	attrs["schedule_id"] = attrs["schedule_id"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
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
				ElemType: CronSchedule_SdkV2{}.Type(ctx),
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
	DestinationSubscriber types.List `tfsdk:"destination_subscriber"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber types.List `tfsdk:"user_subscriber"`
}

func (newState *Subscriber_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscriber_SdkV2) {
}

func (newState *Subscriber_SdkV2) SyncEffectiveFieldsDuringRead(existingState Subscriber_SdkV2) {
}

func (c Subscriber_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_subscriber"] = attrs["destination_subscriber"].SetOptional()
	attrs["destination_subscriber"] = attrs["destination_subscriber"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["user_subscriber"] = attrs["user_subscriber"].SetOptional()
	attrs["user_subscriber"] = attrs["user_subscriber"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: SubscriptionSubscriberDestination_SdkV2{}.Type(ctx),
			},
			"user_subscriber": basetypes.ListType{
				ElemType: SubscriptionSubscriberUser_SdkV2{}.Type(ctx),
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
	CreateTime types.String `tfsdk:"create_time"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber types.List `tfsdk:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *Subscription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscription_SdkV2) {
}

func (newState *Subscription_SdkV2) SyncEffectiveFieldsDuringRead(existingState Subscription_SdkV2) {
}

func (c Subscription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by_user_id"] = attrs["created_by_user_id"].SetComputed()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["schedule_id"] = attrs["schedule_id"].SetComputed()
	attrs["subscriber"] = attrs["subscriber"].SetRequired()
	attrs["subscriber"] = attrs["subscriber"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["subscription_id"] = attrs["subscription_id"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
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
				ElemType: Subscriber_SdkV2{}.Type(ctx),
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
	DestinationId types.String `tfsdk:"destination_id"`
}

func (newState *SubscriptionSubscriberDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberDestination_SdkV2) {
}

func (newState *SubscriptionSubscriberDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberDestination_SdkV2) {
}

func (c SubscriptionSubscriberDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_id"] = attrs["destination_id"].SetComputed()

	return attrs
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
	UserId types.Int64 `tfsdk:"user_id"`
}

func (newState *SubscriptionSubscriberUser_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberUser_SdkV2) {
}

func (newState *SubscriptionSubscriberUser_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberUser_SdkV2) {
}

func (c SubscriptionSubscriberUser_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user_id"] = attrs["user_id"].SetComputed()

	return attrs
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
	Content types.String `tfsdk:"content"`

	Id types.String `tfsdk:"id"`
}

func (newState *TextAttachment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextAttachment_SdkV2) {
}

func (newState *TextAttachment_SdkV2) SyncEffectiveFieldsDuringRead(existingState TextAttachment_SdkV2) {
}

func (c TextAttachment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
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

func (c TrashDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (c UnpublishDashboardResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Dashboard types.List `tfsdk:"dashboard"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
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
				ElemType: Dashboard_SdkV2{}.Type(ctx),
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

	Schedule types.List `tfsdk:"schedule"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
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
				ElemType: Schedule_SdkV2{}.Type(ctx),
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
