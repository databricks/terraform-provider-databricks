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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AuthorizationDetails struct {
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

func (toState *AuthorizationDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AuthorizationDetails) {
}

func (toState *AuthorizationDetails) SyncFieldsDuringRead(ctx context.Context, fromState AuthorizationDetails) {
}

func (c AuthorizationDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AuthorizationDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(AuthorizationDetailsGrantRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AuthorizationDetails
// only implements ToObjectValue() and Type().
func (o AuthorizationDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AuthorizationDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"grant_rules": basetypes.ListType{
				ElemType: AuthorizationDetailsGrantRule{}.Type(ctx),
			},
			"resource_legacy_acl_path": types.StringType,
			"resource_name":            types.StringType,
			"type":                     types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in AuthorizationDetails as
// a slice of AuthorizationDetailsGrantRule values.
// If the field is unknown or null, the boolean return value is false.
func (o *AuthorizationDetails) GetGrantRules(ctx context.Context) ([]AuthorizationDetailsGrantRule, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []AuthorizationDetailsGrantRule
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in AuthorizationDetails.
func (o *AuthorizationDetails) SetGrantRules(ctx context.Context, v []AuthorizationDetailsGrantRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type AuthorizationDetailsGrantRule struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	PermissionSet types.String `tfsdk:"permission_set"`
}

func (toState *AuthorizationDetailsGrantRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AuthorizationDetailsGrantRule) {
}

func (toState *AuthorizationDetailsGrantRule) SyncFieldsDuringRead(ctx context.Context, fromState AuthorizationDetailsGrantRule) {
}

func (c AuthorizationDetailsGrantRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AuthorizationDetailsGrantRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AuthorizationDetailsGrantRule
// only implements ToObjectValue() and Type().
func (o AuthorizationDetailsGrantRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": o.PermissionSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AuthorizationDetailsGrantRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
		},
	}
}

type CreateDashboardRequest struct {
	Dashboard types.Object `tfsdk:"dashboard"`
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
			"dashboard": Dashboard{}.Type(ctx),
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
	var v Dashboard
	d := o.Dashboard.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDashboard sets the value of the Dashboard field in CreateDashboardRequest.
func (o *CreateDashboardRequest) SetDashboard(ctx context.Context, v Dashboard) {
	vs := v.ToObjectValue(ctx)
	o.Dashboard = vs
}

type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The schedule to create. A dashboard is limited to 10 schedules.
	Schedule types.Object `tfsdk:"schedule"`
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
			"schedule":     Schedule{}.Type(ctx),
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
	var v Schedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in CreateScheduleRequest.
func (o *CreateScheduleRequest) SetSchedule(ctx context.Context, v Schedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// The subscription to create. A schedule is limited to 100 subscriptions.
	Subscription types.Object `tfsdk:"subscription"`
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
			"subscription": Subscription{}.Type(ctx),
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
	var v Subscription
	d := o.Subscription.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscription sets the value of the Subscription field in CreateSubscriptionRequest.
func (o *CreateSubscriptionRequest) SetSubscription(ctx context.Context, v Subscription) {
	vs := v.ToObjectValue(ctx)
	o.Subscription = vs
}

type CronSchedule struct {
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

func (toState *CronSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CronSchedule) {
}

func (toState *CronSchedule) SyncFieldsDuringRead(ctx context.Context, fromState CronSchedule) {
}

func (c CronSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (toState *Dashboard) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Dashboard) {
}

func (toState *Dashboard) SyncFieldsDuringRead(ctx context.Context, fromState Dashboard) {
}

func (c Dashboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteScheduleRequest struct {
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

// Genie AI Response
type GenieAttachment struct {
	// Attachment ID
	AttachmentId types.String `tfsdk:"attachment_id"`
	// Query Attachment if Genie responds with a SQL query
	Query types.Object `tfsdk:"query"`
	// Text Attachment if Genie responds with text
	Text types.Object `tfsdk:"text"`
}

func (toState *GenieAttachment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieAttachment) {
	if !fromPlan.Query.IsNull() && !fromPlan.Query.IsUnknown() {
		if toStateQuery, ok := toState.GetQuery(ctx); ok {
			if fromPlanQuery, ok := fromPlan.GetQuery(ctx); ok {
				toStateQuery.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanQuery)
				toState.SetQuery(ctx, toStateQuery)
			}
		}
	}
	if !fromPlan.Text.IsNull() && !fromPlan.Text.IsUnknown() {
		if toStateText, ok := toState.GetText(ctx); ok {
			if fromPlanText, ok := fromPlan.GetText(ctx); ok {
				toStateText.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanText)
				toState.SetText(ctx, toStateText)
			}
		}
	}
}

func (toState *GenieAttachment) SyncFieldsDuringRead(ctx context.Context, fromState GenieAttachment) {
	if !fromState.Query.IsNull() && !fromState.Query.IsUnknown() {
		if toStateQuery, ok := toState.GetQuery(ctx); ok {
			if fromStateQuery, ok := fromState.GetQuery(ctx); ok {
				toStateQuery.SyncFieldsDuringRead(ctx, fromStateQuery)
				toState.SetQuery(ctx, toStateQuery)
			}
		}
	}
	if !fromState.Text.IsNull() && !fromState.Text.IsUnknown() {
		if toStateText, ok := toState.GetText(ctx); ok {
			if fromStateText, ok := fromState.GetText(ctx); ok {
				toStateText.SyncFieldsDuringRead(ctx, fromStateText)
				toState.SetText(ctx, toStateText)
			}
		}
	}
}

func (c GenieAttachment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attachment_id"] = attrs["attachment_id"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["text"] = attrs["text"].SetOptional()

	return attrs
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
		"query": reflect.TypeOf(GenieQueryAttachment{}),
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
			"attachment_id": o.AttachmentId,
			"query":         o.Query,
			"text":          o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id": types.StringType,
			"query":         GenieQueryAttachment{}.Type(ctx),
			"text":          TextAttachment{}.Type(ctx),
		},
	}
}

// GetQuery returns the value of the Query field in GenieAttachment as
// a GenieQueryAttachment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment) GetQuery(ctx context.Context) (GenieQueryAttachment, bool) {
	var e GenieQueryAttachment
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v GenieQueryAttachment
	d := o.Query.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuery sets the value of the Query field in GenieAttachment.
func (o *GenieAttachment) SetQuery(ctx context.Context, v GenieQueryAttachment) {
	vs := v.ToObjectValue(ctx)
	o.Query = vs
}

// GetText returns the value of the Text field in GenieAttachment as
// a TextAttachment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieAttachment) GetText(ctx context.Context) (TextAttachment, bool) {
	var e TextAttachment
	if o.Text.IsNull() || o.Text.IsUnknown() {
		return e, false
	}
	var v TextAttachment
	d := o.Text.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetText sets the value of the Text field in GenieAttachment.
func (o *GenieAttachment) SetText(ctx context.Context, v TextAttachment) {
	vs := v.ToObjectValue(ctx)
	o.Text = vs
}

type GenieConversation struct {
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

func (toState *GenieConversation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieConversation) {
}

func (toState *GenieConversation) SyncFieldsDuringRead(ctx context.Context, fromState GenieConversation) {
}

func (c GenieConversation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (o GenieConversation) Type(ctx context.Context) attr.Type {
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

type GenieConversationSummary struct {
	ConversationId types.String `tfsdk:"conversation_id"`

	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`

	Title types.String `tfsdk:"title"`
}

func (toState *GenieConversationSummary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieConversationSummary) {
}

func (toState *GenieConversationSummary) SyncFieldsDuringRead(ctx context.Context, fromState GenieConversationSummary) {
}

func (c GenieConversationSummary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetRequired()
	attrs["title"] = attrs["title"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieConversationSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieConversationSummary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieConversationSummary
// only implements ToObjectValue() and Type().
func (o GenieConversationSummary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id":   o.ConversationId,
			"created_timestamp": o.CreatedTimestamp,
			"title":             o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieConversationSummary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id":   types.StringType,
			"created_timestamp": types.Int64Type,
			"title":             types.StringType,
		},
	}
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content types.String `tfsdk:"content"`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
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

type GenieDeleteConversationMessageRequest struct {
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the message to delete.
	MessageId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the message is located.
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieDeleteConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieDeleteConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieDeleteConversationMessageRequest
// only implements ToObjectValue() and Type().
func (o GenieDeleteConversationMessageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieDeleteConversationMessageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieDeleteConversationRequest struct {
	// The ID of the conversation to delete.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is located.
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieDeleteConversationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieDeleteConversationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieDeleteConversationRequest
// only implements ToObjectValue() and Type().
func (o GenieDeleteConversationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieDeleteConversationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieExecuteMessageAttachmentQueryRequest struct {
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
func (a GenieExecuteMessageAttachmentQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieExecuteMessageAttachmentQueryRequest
// only implements ToObjectValue() and Type().
func (o GenieExecuteMessageAttachmentQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieExecuteMessageAttachmentQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieExecuteMessageQueryRequest struct {
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

// Feedback containing rating and optional comment
type GenieFeedback struct {
	// Optional feedback comment text
	Comment types.String `tfsdk:"comment"`
	// The feedback rating
	Rating types.String `tfsdk:"rating"`
}

func (toState *GenieFeedback) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieFeedback) {
}

func (toState *GenieFeedback) SyncFieldsDuringRead(ctx context.Context, fromState GenieFeedback) {
}

func (c GenieFeedback) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["rating"] = attrs["rating"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieFeedback.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieFeedback) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieFeedback
// only implements ToObjectValue() and Type().
func (o GenieFeedback) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"rating":  o.Rating,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieFeedback) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"rating":  types.StringType,
		},
	}
}

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

type GenieGetMessageAttachmentQueryResultRequest struct {
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
func (a GenieGetMessageAttachmentQueryResultRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetMessageAttachmentQueryResultRequest
// only implements ToObjectValue() and Type().
func (o GenieGetMessageAttachmentQueryResultRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieGetMessageAttachmentQueryResultRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGetMessageQueryResultRequest struct {
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
	StatementResponse types.Object `tfsdk:"statement_response"`
}

func (toState *GenieGetMessageQueryResultResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieGetMessageQueryResultResponse) {
	if !fromPlan.StatementResponse.IsNull() && !fromPlan.StatementResponse.IsUnknown() {
		if toStateStatementResponse, ok := toState.GetStatementResponse(ctx); ok {
			if fromPlanStatementResponse, ok := fromPlan.GetStatementResponse(ctx); ok {
				toStateStatementResponse.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanStatementResponse)
				toState.SetStatementResponse(ctx, toStateStatementResponse)
			}
		}
	}
}

func (toState *GenieGetMessageQueryResultResponse) SyncFieldsDuringRead(ctx context.Context, fromState GenieGetMessageQueryResultResponse) {
	if !fromState.StatementResponse.IsNull() && !fromState.StatementResponse.IsUnknown() {
		if toStateStatementResponse, ok := toState.GetStatementResponse(ctx); ok {
			if fromStateStatementResponse, ok := fromState.GetStatementResponse(ctx); ok {
				toStateStatementResponse.SyncFieldsDuringRead(ctx, fromStateStatementResponse)
				toState.SetStatementResponse(ctx, toStateStatementResponse)
			}
		}
	}
}

func (c GenieGetMessageQueryResultResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_response"] = attrs["statement_response"].SetOptional()

	return attrs
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
			"statement_response": sql_tf.StatementResponse{}.Type(ctx),
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
	var v sql_tf.StatementResponse
	d := o.StatementResponse.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatementResponse sets the value of the StatementResponse field in GenieGetMessageQueryResultResponse.
func (o *GenieGetMessageQueryResultResponse) SetStatementResponse(ctx context.Context, v sql_tf.StatementResponse) {
	vs := v.ToObjectValue(ctx)
	o.StatementResponse = vs
}

type GenieGetQueryResultByAttachmentRequest struct {
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
func (a GenieGetQueryResultByAttachmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetQueryResultByAttachmentRequest
// only implements ToObjectValue() and Type().
func (o GenieGetQueryResultByAttachmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieGetQueryResultByAttachmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachment_id":   types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieGetSpaceRequest struct {
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
func (a GenieGetSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieGetSpaceRequest
// only implements ToObjectValue() and Type().
func (o GenieGetSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"space_id": o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieGetSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"space_id": types.StringType,
		},
	}
}

type GenieListConversationMessagesRequest struct {
	// The ID of the conversation to list messages from
	ConversationId types.String `tfsdk:"-"`
	// Maximum number of messages to return per page
	PageSize types.Int64 `tfsdk:"-"`
	// Token to get the next page of results
	PageToken types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is located
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListConversationMessagesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListConversationMessagesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListConversationMessagesRequest
// only implements ToObjectValue() and Type().
func (o GenieListConversationMessagesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversation_id": o.ConversationId,
			"page_size":       o.PageSize,
			"page_token":      o.PageToken,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListConversationMessagesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieListConversationMessagesResponse struct {
	// List of messages in the conversation.
	Messages types.List `tfsdk:"messages"`
	// The token to use for retrieving the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *GenieListConversationMessagesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieListConversationMessagesResponse) {
}

func (toState *GenieListConversationMessagesResponse) SyncFieldsDuringRead(ctx context.Context, fromState GenieListConversationMessagesResponse) {
}

func (c GenieListConversationMessagesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["messages"] = attrs["messages"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListConversationMessagesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListConversationMessagesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"messages": reflect.TypeOf(GenieMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListConversationMessagesResponse
// only implements ToObjectValue() and Type().
func (o GenieListConversationMessagesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"messages":        o.Messages,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListConversationMessagesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"messages": basetypes.ListType{
				ElemType: GenieMessage{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMessages returns the value of the Messages field in GenieListConversationMessagesResponse as
// a slice of GenieMessage values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieListConversationMessagesResponse) GetMessages(ctx context.Context) ([]GenieMessage, bool) {
	if o.Messages.IsNull() || o.Messages.IsUnknown() {
		return nil, false
	}
	var v []GenieMessage
	d := o.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in GenieListConversationMessagesResponse.
func (o *GenieListConversationMessagesResponse) SetMessages(ctx context.Context, v []GenieMessage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Messages = types.ListValueMust(t, vs)
}

type GenieListConversationsRequest struct {
	// Include all conversations in the space across all users. Requires at
	// least CAN MANAGE permission on the space.
	IncludeAll types.Bool `tfsdk:"-"`
	// Maximum number of conversations to return per page
	PageSize types.Int64 `tfsdk:"-"`
	// Token to get the next page of results
	PageToken types.String `tfsdk:"-"`
	// The ID of the Genie space to retrieve conversations from.
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListConversationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListConversationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListConversationsRequest
// only implements ToObjectValue() and Type().
func (o GenieListConversationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_all": o.IncludeAll,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
			"space_id":    o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListConversationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_all": types.BoolType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
			"space_id":    types.StringType,
		},
	}
}

type GenieListConversationsResponse struct {
	// List of conversations in the Genie space
	Conversations types.List `tfsdk:"conversations"`
	// Token to get the next page of results
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *GenieListConversationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieListConversationsResponse) {
}

func (toState *GenieListConversationsResponse) SyncFieldsDuringRead(ctx context.Context, fromState GenieListConversationsResponse) {
}

func (c GenieListConversationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["conversations"] = attrs["conversations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieListConversationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieListConversationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"conversations": reflect.TypeOf(GenieConversationSummary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListConversationsResponse
// only implements ToObjectValue() and Type().
func (o GenieListConversationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"conversations":   o.Conversations,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListConversationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversations": basetypes.ListType{
				ElemType: GenieConversationSummary{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetConversations returns the value of the Conversations field in GenieListConversationsResponse as
// a slice of GenieConversationSummary values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieListConversationsResponse) GetConversations(ctx context.Context) ([]GenieConversationSummary, bool) {
	if o.Conversations.IsNull() || o.Conversations.IsUnknown() {
		return nil, false
	}
	var v []GenieConversationSummary
	d := o.Conversations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConversations sets the value of the Conversations field in GenieListConversationsResponse.
func (o *GenieListConversationsResponse) SetConversations(ctx context.Context, v []GenieConversationSummary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["conversations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Conversations = types.ListValueMust(t, vs)
}

type GenieListSpacesRequest struct {
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
func (a GenieListSpacesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListSpacesRequest
// only implements ToObjectValue() and Type().
func (o GenieListSpacesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListSpacesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type GenieListSpacesResponse struct {
	// Token to get the next page of results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Genie spaces
	Spaces types.List `tfsdk:"spaces"`
}

func (toState *GenieListSpacesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieListSpacesResponse) {
}

func (toState *GenieListSpacesResponse) SyncFieldsDuringRead(ctx context.Context, fromState GenieListSpacesResponse) {
}

func (c GenieListSpacesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenieListSpacesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spaces": reflect.TypeOf(GenieSpace{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieListSpacesResponse
// only implements ToObjectValue() and Type().
func (o GenieListSpacesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"spaces":          o.Spaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieListSpacesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"spaces": basetypes.ListType{
				ElemType: GenieSpace{}.Type(ctx),
			},
		},
	}
}

// GetSpaces returns the value of the Spaces field in GenieListSpacesResponse as
// a slice of GenieSpace values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieListSpacesResponse) GetSpaces(ctx context.Context) ([]GenieSpace, bool) {
	if o.Spaces.IsNull() || o.Spaces.IsUnknown() {
		return nil, false
	}
	var v []GenieSpace
	d := o.Spaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpaces sets the value of the Spaces field in GenieListSpacesResponse.
func (o *GenieListSpacesResponse) SetSpaces(ctx context.Context, v []GenieSpace) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Spaces = types.ListValueMust(t, vs)
}

type GenieMessage struct {
	// AI-generated response to the message
	Attachments types.List `tfsdk:"attachments"`
	// User message content
	Content types.String `tfsdk:"content"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp"`
	// Error message if Genie failed to respond to the message
	Error types.Object `tfsdk:"error"`
	// User feedback for the message if provided
	Feedback types.Object `tfsdk:"feedback"`
	// Message ID. Legacy identifier, use message_id instead
	Id types.String `tfsdk:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Message ID
	MessageId types.String `tfsdk:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult types.Object `tfsdk:"query_result"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id"`

	Status types.String `tfsdk:"status"`
	// ID of the user who created the message
	UserId types.Int64 `tfsdk:"user_id"`
}

func (toState *GenieMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieMessage) {
	if !fromPlan.Error.IsNull() && !fromPlan.Error.IsUnknown() {
		if toStateError, ok := toState.GetError(ctx); ok {
			if fromPlanError, ok := fromPlan.GetError(ctx); ok {
				toStateError.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanError)
				toState.SetError(ctx, toStateError)
			}
		}
	}
	if !fromPlan.Feedback.IsNull() && !fromPlan.Feedback.IsUnknown() {
		if toStateFeedback, ok := toState.GetFeedback(ctx); ok {
			if fromPlanFeedback, ok := fromPlan.GetFeedback(ctx); ok {
				toStateFeedback.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFeedback)
				toState.SetFeedback(ctx, toStateFeedback)
			}
		}
	}
	if !fromPlan.QueryResult.IsNull() && !fromPlan.QueryResult.IsUnknown() {
		if toStateQueryResult, ok := toState.GetQueryResult(ctx); ok {
			if fromPlanQueryResult, ok := fromPlan.GetQueryResult(ctx); ok {
				toStateQueryResult.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanQueryResult)
				toState.SetQueryResult(ctx, toStateQueryResult)
			}
		}
	}
}

func (toState *GenieMessage) SyncFieldsDuringRead(ctx context.Context, fromState GenieMessage) {
	if !fromState.Error.IsNull() && !fromState.Error.IsUnknown() {
		if toStateError, ok := toState.GetError(ctx); ok {
			if fromStateError, ok := fromState.GetError(ctx); ok {
				toStateError.SyncFieldsDuringRead(ctx, fromStateError)
				toState.SetError(ctx, toStateError)
			}
		}
	}
	if !fromState.Feedback.IsNull() && !fromState.Feedback.IsUnknown() {
		if toStateFeedback, ok := toState.GetFeedback(ctx); ok {
			if fromStateFeedback, ok := fromState.GetFeedback(ctx); ok {
				toStateFeedback.SyncFieldsDuringRead(ctx, fromStateFeedback)
				toState.SetFeedback(ctx, toStateFeedback)
			}
		}
	}
	if !fromState.QueryResult.IsNull() && !fromState.QueryResult.IsUnknown() {
		if toStateQueryResult, ok := toState.GetQueryResult(ctx); ok {
			if fromStateQueryResult, ok := fromState.GetQueryResult(ctx); ok {
				toStateQueryResult.SyncFieldsDuringRead(ctx, fromStateQueryResult)
				toState.SetQueryResult(ctx, toStateQueryResult)
			}
		}
	}
}

func (c GenieMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attachments"] = attrs["attachments"].SetOptional()
	attrs["content"] = attrs["content"].SetRequired()
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["created_timestamp"] = attrs["created_timestamp"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["feedback"] = attrs["feedback"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["message_id"] = attrs["message_id"].SetRequired()
	attrs["query_result"] = attrs["query_result"].SetOptional()
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
func (a GenieMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attachments":  reflect.TypeOf(GenieAttachment{}),
		"error":        reflect.TypeOf(MessageError{}),
		"feedback":     reflect.TypeOf(GenieFeedback{}),
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
			"feedback":               o.Feedback,
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
func (o GenieMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachments": basetypes.ListType{
				ElemType: GenieAttachment{}.Type(ctx),
			},
			"content":                types.StringType,
			"conversation_id":        types.StringType,
			"created_timestamp":      types.Int64Type,
			"error":                  MessageError{}.Type(ctx),
			"feedback":               GenieFeedback{}.Type(ctx),
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"message_id":             types.StringType,
			"query_result":           Result{}.Type(ctx),
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
	var v MessageError
	d := o.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in GenieMessage.
func (o *GenieMessage) SetError(ctx context.Context, v MessageError) {
	vs := v.ToObjectValue(ctx)
	o.Error = vs
}

// GetFeedback returns the value of the Feedback field in GenieMessage as
// a GenieFeedback value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage) GetFeedback(ctx context.Context) (GenieFeedback, bool) {
	var e GenieFeedback
	if o.Feedback.IsNull() || o.Feedback.IsUnknown() {
		return e, false
	}
	var v GenieFeedback
	d := o.Feedback.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeedback sets the value of the Feedback field in GenieMessage.
func (o *GenieMessage) SetFeedback(ctx context.Context, v GenieFeedback) {
	vs := v.ToObjectValue(ctx)
	o.Feedback = vs
}

// GetQueryResult returns the value of the QueryResult field in GenieMessage as
// a Result value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieMessage) GetQueryResult(ctx context.Context) (Result, bool) {
	var e Result
	if o.QueryResult.IsNull() || o.QueryResult.IsUnknown() {
		return e, false
	}
	var v Result
	d := o.QueryResult.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryResult sets the value of the QueryResult field in GenieMessage.
func (o *GenieMessage) SetQueryResult(ctx context.Context, v Result) {
	vs := v.ToObjectValue(ctx)
	o.QueryResult = vs
}

type GenieQueryAttachment struct {
	// Description of the query
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
	// Time when the user updated the query last
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// AI generated SQL query
	Query types.String `tfsdk:"query"`
	// Metadata associated with the query result.
	QueryResultMetadata types.Object `tfsdk:"query_result_metadata"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id"`
	// Name of the query
	Title types.String `tfsdk:"title"`
}

func (toState *GenieQueryAttachment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieQueryAttachment) {
	if !fromPlan.QueryResultMetadata.IsNull() && !fromPlan.QueryResultMetadata.IsUnknown() {
		if toStateQueryResultMetadata, ok := toState.GetQueryResultMetadata(ctx); ok {
			if fromPlanQueryResultMetadata, ok := fromPlan.GetQueryResultMetadata(ctx); ok {
				toStateQueryResultMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanQueryResultMetadata)
				toState.SetQueryResultMetadata(ctx, toStateQueryResultMetadata)
			}
		}
	}
}

func (toState *GenieQueryAttachment) SyncFieldsDuringRead(ctx context.Context, fromState GenieQueryAttachment) {
	if !fromState.QueryResultMetadata.IsNull() && !fromState.QueryResultMetadata.IsUnknown() {
		if toStateQueryResultMetadata, ok := toState.GetQueryResultMetadata(ctx); ok {
			if fromStateQueryResultMetadata, ok := fromState.GetQueryResultMetadata(ctx); ok {
				toStateQueryResultMetadata.SyncFieldsDuringRead(ctx, fromStateQueryResultMetadata)
				toState.SetQueryResultMetadata(ctx, toStateQueryResultMetadata)
			}
		}
	}
}

func (c GenieQueryAttachment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query_result_metadata"] = attrs["query_result_metadata"].SetOptional()
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
func (a GenieQueryAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_result_metadata": reflect.TypeOf(GenieResultMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieQueryAttachment
// only implements ToObjectValue() and Type().
func (o GenieQueryAttachment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenieQueryAttachment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query":                  types.StringType,
			"query_result_metadata":  GenieResultMetadata{}.Type(ctx),
			"statement_id":           types.StringType,
			"title":                  types.StringType,
		},
	}
}

// GetQueryResultMetadata returns the value of the QueryResultMetadata field in GenieQueryAttachment as
// a GenieResultMetadata value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieQueryAttachment) GetQueryResultMetadata(ctx context.Context) (GenieResultMetadata, bool) {
	var e GenieResultMetadata
	if o.QueryResultMetadata.IsNull() || o.QueryResultMetadata.IsUnknown() {
		return e, false
	}
	var v GenieResultMetadata
	d := o.QueryResultMetadata.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryResultMetadata sets the value of the QueryResultMetadata field in GenieQueryAttachment.
func (o *GenieQueryAttachment) SetQueryResultMetadata(ctx context.Context, v GenieResultMetadata) {
	vs := v.ToObjectValue(ctx)
	o.QueryResultMetadata = vs
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	IsTruncated types.Bool `tfsdk:"is_truncated"`
	// The number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (toState *GenieResultMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieResultMetadata) {
}

func (toState *GenieResultMetadata) SyncFieldsDuringRead(ctx context.Context, fromState GenieResultMetadata) {
}

func (c GenieResultMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenieResultMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieResultMetadata
// only implements ToObjectValue() and Type().
func (o GenieResultMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_truncated": o.IsTruncated,
			"row_count":    o.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieResultMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated": types.BoolType,
			"row_count":    types.Int64Type,
		},
	}
}

type GenieSendMessageFeedbackRequest struct {
	// Optional text feedback that will be stored as a comment.
	Comment types.String `tfsdk:"comment"`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the message to provide feedback for.
	MessageId types.String `tfsdk:"-"`
	// The rating (POSITIVE, NEGATIVE, or NONE).
	Rating types.String `tfsdk:"rating"`
	// The ID associated with the Genie space where the message is located.
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieSendMessageFeedbackRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieSendMessageFeedbackRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieSendMessageFeedbackRequest
// only implements ToObjectValue() and Type().
func (o GenieSendMessageFeedbackRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":         o.Comment,
			"conversation_id": o.ConversationId,
			"message_id":      o.MessageId,
			"rating":          o.Rating,
			"space_id":        o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieSendMessageFeedbackRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"rating":          types.StringType,
			"space_id":        types.StringType,
		},
	}
}

type GenieSpace struct {
	// Description of the Genie Space
	Description types.String `tfsdk:"description"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id"`
	// Title of the Genie Space
	Title types.String `tfsdk:"title"`
}

func (toState *GenieSpace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieSpace) {
}

func (toState *GenieSpace) SyncFieldsDuringRead(ctx context.Context, fromState GenieSpace) {
}

func (c GenieSpace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenieSpace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieSpace
// only implements ToObjectValue() and Type().
func (o GenieSpace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"space_id":    o.SpaceId,
			"title":       o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieSpace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"space_id":    types.StringType,
			"title":       types.StringType,
		},
	}
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
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
	Conversation types.Object `tfsdk:"conversation"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id"`

	Message types.Object `tfsdk:"message"`
	// Message ID
	MessageId types.String `tfsdk:"message_id"`
}

func (toState *GenieStartConversationResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GenieStartConversationResponse) {
	if !fromPlan.Conversation.IsNull() && !fromPlan.Conversation.IsUnknown() {
		if toStateConversation, ok := toState.GetConversation(ctx); ok {
			if fromPlanConversation, ok := fromPlan.GetConversation(ctx); ok {
				toStateConversation.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanConversation)
				toState.SetConversation(ctx, toStateConversation)
			}
		}
	}
	if !fromPlan.Message.IsNull() && !fromPlan.Message.IsUnknown() {
		if toStateMessage, ok := toState.GetMessage(ctx); ok {
			if fromPlanMessage, ok := fromPlan.GetMessage(ctx); ok {
				toStateMessage.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanMessage)
				toState.SetMessage(ctx, toStateMessage)
			}
		}
	}
}

func (toState *GenieStartConversationResponse) SyncFieldsDuringRead(ctx context.Context, fromState GenieStartConversationResponse) {
	if !fromState.Conversation.IsNull() && !fromState.Conversation.IsUnknown() {
		if toStateConversation, ok := toState.GetConversation(ctx); ok {
			if fromStateConversation, ok := fromState.GetConversation(ctx); ok {
				toStateConversation.SyncFieldsDuringRead(ctx, fromStateConversation)
				toState.SetConversation(ctx, toStateConversation)
			}
		}
	}
	if !fromState.Message.IsNull() && !fromState.Message.IsUnknown() {
		if toStateMessage, ok := toState.GetMessage(ctx); ok {
			if fromStateMessage, ok := fromState.GetMessage(ctx); ok {
				toStateMessage.SyncFieldsDuringRead(ctx, fromStateMessage)
				toState.SetMessage(ctx, toStateMessage)
			}
		}
	}
}

func (c GenieStartConversationResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["conversation"] = attrs["conversation"].SetOptional()
	attrs["conversation_id"] = attrs["conversation_id"].SetRequired()
	attrs["message"] = attrs["message"].SetOptional()
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
			"conversation":    GenieConversation{}.Type(ctx),
			"conversation_id": types.StringType,
			"message":         GenieMessage{}.Type(ctx),
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
	var v GenieConversation
	d := o.Conversation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConversation sets the value of the Conversation field in GenieStartConversationResponse.
func (o *GenieStartConversationResponse) SetConversation(ctx context.Context, v GenieConversation) {
	vs := v.ToObjectValue(ctx)
	o.Conversation = vs
}

// GetMessage returns the value of the Message field in GenieStartConversationResponse as
// a GenieMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenieStartConversationResponse) GetMessage(ctx context.Context) (GenieMessage, bool) {
	var e GenieMessage
	if o.Message.IsNull() || o.Message.IsUnknown() {
		return e, false
	}
	var v GenieMessage
	d := o.Message.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessage sets the value of the Message field in GenieStartConversationResponse.
func (o *GenieStartConversationResponse) SetMessage(ctx context.Context, v GenieMessage) {
	vs := v.ToObjectValue(ctx)
	o.Message = vs
}

type GenieTrashSpaceRequest struct {
	// The ID associated with the Genie space to be sent to the trash.
	SpaceId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieTrashSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieTrashSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieTrashSpaceRequest
// only implements ToObjectValue() and Type().
func (o GenieTrashSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"space_id": o.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenieTrashSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"space_id": types.StringType,
		},
	}
}

type GetDashboardRequest struct {
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

type GetPublishedDashboardRequest struct {
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

type GetPublishedDashboardTokenInfoRequest struct {
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
func (a GetPublishedDashboardTokenInfoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardTokenInfoRequest
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardTokenInfoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":       o.DashboardId,
			"external_value":     o.ExternalValue,
			"external_viewer_id": o.ExternalViewerId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardTokenInfoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":       types.StringType,
			"external_value":     types.StringType,
			"external_viewer_id": types.StringType,
		},
	}
}

type GetPublishedDashboardTokenInfoResponse struct {
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

func (toState *GetPublishedDashboardTokenInfoResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPublishedDashboardTokenInfoResponse) {
}

func (toState *GetPublishedDashboardTokenInfoResponse) SyncFieldsDuringRead(ctx context.Context, fromState GetPublishedDashboardTokenInfoResponse) {
}

func (c GetPublishedDashboardTokenInfoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetPublishedDashboardTokenInfoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"authorization_details": reflect.TypeOf(AuthorizationDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedDashboardTokenInfoResponse
// only implements ToObjectValue() and Type().
func (o GetPublishedDashboardTokenInfoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization_details": o.AuthorizationDetails,
			"custom_claim":          o.CustomClaim,
			"scope":                 o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedDashboardTokenInfoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization_details": basetypes.ListType{
				ElemType: AuthorizationDetails{}.Type(ctx),
			},
			"custom_claim": types.StringType,
			"scope":        types.StringType,
		},
	}
}

// GetAuthorizationDetails returns the value of the AuthorizationDetails field in GetPublishedDashboardTokenInfoResponse as
// a slice of AuthorizationDetails values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPublishedDashboardTokenInfoResponse) GetAuthorizationDetails(ctx context.Context) ([]AuthorizationDetails, bool) {
	if o.AuthorizationDetails.IsNull() || o.AuthorizationDetails.IsUnknown() {
		return nil, false
	}
	var v []AuthorizationDetails
	d := o.AuthorizationDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAuthorizationDetails sets the value of the AuthorizationDetails field in GetPublishedDashboardTokenInfoResponse.
func (o *GetPublishedDashboardTokenInfoResponse) SetAuthorizationDetails(ctx context.Context, v []AuthorizationDetails) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["authorization_details"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AuthorizationDetails = types.ListValueMust(t, vs)
}

type GetScheduleRequest struct {
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

type GetSubscriptionRequest struct {
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
	Dashboards types.List `tfsdk:"dashboards"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListDashboardsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDashboardsResponse) {
}

func (toState *ListDashboardsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListDashboardsResponse) {
}

func (c ListDashboardsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"dashboards": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
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

type ListSchedulesRequest struct {
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Schedules types.List `tfsdk:"schedules"`
}

func (toState *ListSchedulesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListSchedulesResponse) {
}

func (toState *ListSchedulesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListSchedulesResponse) {
}

func (c ListSchedulesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"schedules": basetypes.ListType{
				ElemType: Schedule{}.Type(ctx),
			},
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Subscriptions types.List `tfsdk:"subscriptions"`
}

func (toState *ListSubscriptionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListSubscriptionsResponse) {
}

func (toState *ListSubscriptionsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListSubscriptionsResponse) {
}

func (c ListSubscriptionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"subscriptions": basetypes.ListType{
				ElemType: Subscription{}.Type(ctx),
			},
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
	Error types.String `tfsdk:"error"`

	Type_ types.String `tfsdk:"type"`
}

func (toState *MessageError) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan MessageError) {
}

func (toState *MessageError) SyncFieldsDuringRead(ctx context.Context, fromState MessageError) {
}

func (c MessageError) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"display_name":            o.DisplayName,
			"parent_path":             o.ParentPath,
			"source_dashboard_id":     o.SourceDashboardId,
			"update_parameter_syntax": o.UpdateParameterSyntax,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigrateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":            types.StringType,
			"parent_path":             types.StringType,
			"source_dashboard_id":     types.StringType,
			"update_parameter_syntax": types.BoolType,
		},
	}
}

type PublishRequest struct {
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
	DisplayName types.String `tfsdk:"display_name"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (toState *PublishedDashboard) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PublishedDashboard) {
}

func (toState *PublishedDashboard) SyncFieldsDuringRead(ctx context.Context, fromState PublishedDashboard) {
}

func (c PublishedDashboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type Result struct {
	// If result is truncated
	IsTruncated types.Bool `tfsdk:"is_truncated"`
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id"`
	// JWT corresponding to the statement contained in this result
	StatementIdSignature types.String `tfsdk:"statement_id_signature"`
}

func (toState *Result) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Result) {
}

func (toState *Result) SyncFieldsDuringRead(ctx context.Context, fromState Result) {
}

func (c Result) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_truncated"] = attrs["is_truncated"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["statement_id_signature"] = attrs["statement_id_signature"].SetOptional()

	return attrs
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
			"is_truncated":           o.IsTruncated,
			"row_count":              o.RowCount,
			"statement_id":           o.StatementId,
			"statement_id_signature": o.StatementIdSignature,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Result) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated":           types.BoolType,
			"row_count":              types.Int64Type,
			"statement_id":           types.StringType,
			"statement_id_signature": types.StringType,
		},
	}
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule types.Object `tfsdk:"cron_schedule"`
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

func (toState *Schedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Schedule) {
	if !fromPlan.CronSchedule.IsNull() && !fromPlan.CronSchedule.IsUnknown() {
		if toStateCronSchedule, ok := toState.GetCronSchedule(ctx); ok {
			if fromPlanCronSchedule, ok := fromPlan.GetCronSchedule(ctx); ok {
				toStateCronSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCronSchedule)
				toState.SetCronSchedule(ctx, toStateCronSchedule)
			}
		}
	}
}

func (toState *Schedule) SyncFieldsDuringRead(ctx context.Context, fromState Schedule) {
	if !fromState.CronSchedule.IsNull() && !fromState.CronSchedule.IsUnknown() {
		if toStateCronSchedule, ok := toState.GetCronSchedule(ctx); ok {
			if fromStateCronSchedule, ok := fromState.GetCronSchedule(ctx); ok {
				toStateCronSchedule.SyncFieldsDuringRead(ctx, fromStateCronSchedule)
				toState.SetCronSchedule(ctx, toStateCronSchedule)
			}
		}
	}
}

func (c Schedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["cron_schedule"] = attrs["cron_schedule"].SetRequired()
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
			"cron_schedule": CronSchedule{}.Type(ctx),
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
	var v CronSchedule
	d := o.CronSchedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCronSchedule sets the value of the CronSchedule field in Schedule.
func (o *Schedule) SetCronSchedule(ctx context.Context, v CronSchedule) {
	vs := v.ToObjectValue(ctx)
	o.CronSchedule = vs
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber types.Object `tfsdk:"destination_subscriber"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber types.Object `tfsdk:"user_subscriber"`
}

func (toState *Subscriber) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Subscriber) {
	if !fromPlan.DestinationSubscriber.IsNull() && !fromPlan.DestinationSubscriber.IsUnknown() {
		if toStateDestinationSubscriber, ok := toState.GetDestinationSubscriber(ctx); ok {
			if fromPlanDestinationSubscriber, ok := fromPlan.GetDestinationSubscriber(ctx); ok {
				toStateDestinationSubscriber.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDestinationSubscriber)
				toState.SetDestinationSubscriber(ctx, toStateDestinationSubscriber)
			}
		}
	}
	if !fromPlan.UserSubscriber.IsNull() && !fromPlan.UserSubscriber.IsUnknown() {
		if toStateUserSubscriber, ok := toState.GetUserSubscriber(ctx); ok {
			if fromPlanUserSubscriber, ok := fromPlan.GetUserSubscriber(ctx); ok {
				toStateUserSubscriber.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUserSubscriber)
				toState.SetUserSubscriber(ctx, toStateUserSubscriber)
			}
		}
	}
}

func (toState *Subscriber) SyncFieldsDuringRead(ctx context.Context, fromState Subscriber) {
	if !fromState.DestinationSubscriber.IsNull() && !fromState.DestinationSubscriber.IsUnknown() {
		if toStateDestinationSubscriber, ok := toState.GetDestinationSubscriber(ctx); ok {
			if fromStateDestinationSubscriber, ok := fromState.GetDestinationSubscriber(ctx); ok {
				toStateDestinationSubscriber.SyncFieldsDuringRead(ctx, fromStateDestinationSubscriber)
				toState.SetDestinationSubscriber(ctx, toStateDestinationSubscriber)
			}
		}
	}
	if !fromState.UserSubscriber.IsNull() && !fromState.UserSubscriber.IsUnknown() {
		if toStateUserSubscriber, ok := toState.GetUserSubscriber(ctx); ok {
			if fromStateUserSubscriber, ok := fromState.GetUserSubscriber(ctx); ok {
				toStateUserSubscriber.SyncFieldsDuringRead(ctx, fromStateUserSubscriber)
				toState.SetUserSubscriber(ctx, toStateUserSubscriber)
			}
		}
	}
}

func (c Subscriber) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_subscriber"] = attrs["destination_subscriber"].SetOptional()
	attrs["user_subscriber"] = attrs["user_subscriber"].SetOptional()

	return attrs
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
			"destination_subscriber": SubscriptionSubscriberDestination{}.Type(ctx),
			"user_subscriber":        SubscriptionSubscriberUser{}.Type(ctx),
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
	var v SubscriptionSubscriberDestination
	d := o.DestinationSubscriber.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinationSubscriber sets the value of the DestinationSubscriber field in Subscriber.
func (o *Subscriber) SetDestinationSubscriber(ctx context.Context, v SubscriptionSubscriberDestination) {
	vs := v.ToObjectValue(ctx)
	o.DestinationSubscriber = vs
}

// GetUserSubscriber returns the value of the UserSubscriber field in Subscriber as
// a SubscriptionSubscriberUser value.
// If the field is unknown or null, the boolean return value is false.
func (o *Subscriber) GetUserSubscriber(ctx context.Context) (SubscriptionSubscriberUser, bool) {
	var e SubscriptionSubscriberUser
	if o.UserSubscriber.IsNull() || o.UserSubscriber.IsUnknown() {
		return e, false
	}
	var v SubscriptionSubscriberUser
	d := o.UserSubscriber.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserSubscriber sets the value of the UserSubscriber field in Subscriber.
func (o *Subscriber) SetUserSubscriber(ctx context.Context, v SubscriptionSubscriberUser) {
	vs := v.ToObjectValue(ctx)
	o.UserSubscriber = vs
}

type Subscription struct {
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
	Subscriber types.Object `tfsdk:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (toState *Subscription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Subscription) {
	if !fromPlan.Subscriber.IsNull() && !fromPlan.Subscriber.IsUnknown() {
		if toStateSubscriber, ok := toState.GetSubscriber(ctx); ok {
			if fromPlanSubscriber, ok := fromPlan.GetSubscriber(ctx); ok {
				toStateSubscriber.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSubscriber)
				toState.SetSubscriber(ctx, toStateSubscriber)
			}
		}
	}
}

func (toState *Subscription) SyncFieldsDuringRead(ctx context.Context, fromState Subscription) {
	if !fromState.Subscriber.IsNull() && !fromState.Subscriber.IsUnknown() {
		if toStateSubscriber, ok := toState.GetSubscriber(ctx); ok {
			if fromStateSubscriber, ok := fromState.GetSubscriber(ctx); ok {
				toStateSubscriber.SyncFieldsDuringRead(ctx, fromStateSubscriber)
				toState.SetSubscriber(ctx, toStateSubscriber)
			}
		}
	}
}

func (c Subscription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by_user_id"] = attrs["created_by_user_id"].SetComputed()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["schedule_id"] = attrs["schedule_id"].SetComputed()
	attrs["subscriber"] = attrs["subscriber"].SetRequired()
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
			"subscriber":         Subscriber{}.Type(ctx),
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
	var v Subscriber
	d := o.Subscriber.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriber sets the value of the Subscriber field in Subscription.
func (o *Subscription) SetSubscriber(ctx context.Context, v Subscriber) {
	vs := v.ToObjectValue(ctx)
	o.Subscriber = vs
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id"`
}

func (toState *SubscriptionSubscriberDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SubscriptionSubscriberDestination) {
}

func (toState *SubscriptionSubscriberDestination) SyncFieldsDuringRead(ctx context.Context, fromState SubscriptionSubscriberDestination) {
}

func (c SubscriptionSubscriberDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	UserId types.Int64 `tfsdk:"user_id"`
}

func (toState *SubscriptionSubscriberUser) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SubscriptionSubscriberUser) {
}

func (toState *SubscriptionSubscriberUser) SyncFieldsDuringRead(ctx context.Context, fromState SubscriptionSubscriberUser) {
}

func (c SubscriptionSubscriberUser) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Content types.String `tfsdk:"content"`

	Id types.String `tfsdk:"id"`
}

func (toState *TextAttachment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TextAttachment) {
}

func (toState *TextAttachment) SyncFieldsDuringRead(ctx context.Context, fromState TextAttachment) {
}

func (c TextAttachment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type TrashDashboardRequest struct {
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

func (toState *TrashDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TrashDashboardResponse) {
}

func (toState *TrashDashboardResponse) SyncFieldsDuringRead(ctx context.Context, fromState TrashDashboardResponse) {
}

func (c TrashDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type UnpublishDashboardRequest struct {
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

func (toState *UnpublishDashboardResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UnpublishDashboardResponse) {
}

func (toState *UnpublishDashboardResponse) SyncFieldsDuringRead(ctx context.Context, fromState UnpublishDashboardResponse) {
}

func (c UnpublishDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type UpdateDashboardRequest struct {
	Dashboard types.Object `tfsdk:"dashboard"`
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
			"dashboard":    Dashboard{}.Type(ctx),
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
	var v Dashboard
	d := o.Dashboard.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDashboard sets the value of the Dashboard field in UpdateDashboardRequest.
func (o *UpdateDashboardRequest) SetDashboard(ctx context.Context, v Dashboard) {
	vs := v.ToObjectValue(ctx)
	o.Dashboard = vs
}

type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The schedule to update.
	Schedule types.Object `tfsdk:"schedule"`
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
			"schedule":     Schedule{}.Type(ctx),
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
	var v Schedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in UpdateScheduleRequest.
func (o *UpdateScheduleRequest) SetSchedule(ctx context.Context, v Schedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}
