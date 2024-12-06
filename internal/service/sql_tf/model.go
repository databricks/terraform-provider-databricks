// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sql_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccessControl struct {
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`

	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControl) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControl) {
}

func (newState *AccessControl) SyncEffectiveFieldsDuringRead(existingState AccessControl) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControl
// only implements ToObjectValue() and Type().
func (o AccessControl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":       o.GroupName,
			"permission_level": o.PermissionLevel,
			"user_name":        o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccessControl) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":       types.StringType,
			"permission_level": types.StringType,
			"user_name":        types.StringType,
		},
	}
}

type Alert struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition" tf:"optional,object"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time" tf:"optional"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *Alert) SyncEffectiveFieldsDuringCreateOrUpdate(plan Alert) {
}

func (newState *Alert) SyncEffectiveFieldsDuringRead(existingState Alert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Alert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Alert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Alert
// only implements ToObjectValue() and Type().
func (o Alert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            o.Condition,
			"create_time":          o.CreateTime,
			"custom_body":          o.CustomBody,
			"custom_subject":       o.CustomSubject,
			"display_name":         o.DisplayName,
			"id":                   o.Id,
			"lifecycle_state":      o.LifecycleState,
			"notify_on_ok":         o.NotifyOnOk,
			"owner_user_name":      o.OwnerUserName,
			"parent_path":          o.ParentPath,
			"query_id":             o.QueryId,
			"seconds_to_retrigger": o.SecondsToRetrigger,
			"state":                o.State,
			"trigger_time":         o.TriggerTime,
			"update_time":          o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Alert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition{}.Type(ctx),
			},
			"create_time":          types.StringType,
			"custom_body":          types.StringType,
			"custom_subject":       types.StringType,
			"display_name":         types.StringType,
			"id":                   types.StringType,
			"lifecycle_state":      types.StringType,
			"notify_on_ok":         types.BoolType,
			"owner_user_name":      types.StringType,
			"parent_path":          types.StringType,
			"query_id":             types.StringType,
			"seconds_to_retrigger": types.Int64Type,
			"state":                types.StringType,
			"trigger_time":         types.StringType,
			"update_time":          types.StringType,
		},
	}
}

// GetCondition returns the value of the Condition field in Alert as
// a AlertCondition value.
// If the field is unknown or null, the boolean return value is false.
func (o *Alert) GetCondition(ctx context.Context) (AlertCondition, bool) {
	var e AlertCondition
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in Alert.
func (o *Alert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type AlertCondition struct {
	// Alert state if result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state" tf:"optional"`
	// Operator used for comparison in alert evaluation.
	Op types.String `tfsdk:"op" tf:"optional"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand types.List `tfsdk:"operand" tf:"optional,object"`
	// Threshold value used for comparison in alert evaluation.
	Threshold types.List `tfsdk:"threshold" tf:"optional,object"`
}

func (newState *AlertCondition) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertCondition) {
}

func (newState *AlertCondition) SyncEffectiveFieldsDuringRead(existingState AlertCondition) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertCondition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertCondition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"operand":   reflect.TypeOf(AlertConditionOperand{}),
		"threshold": reflect.TypeOf(AlertConditionThreshold{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertCondition
// only implements ToObjectValue() and Type().
func (o AlertCondition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"empty_result_state": o.EmptyResultState,
			"op":                 o.Op,
			"operand":            o.Operand,
			"threshold":          o.Threshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertCondition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"empty_result_state": types.StringType,
			"op":                 types.StringType,
			"operand": basetypes.ListType{
				ElemType: AlertConditionOperand{}.Type(ctx),
			},
			"threshold": basetypes.ListType{
				ElemType: AlertConditionThreshold{}.Type(ctx),
			},
		},
	}
}

// GetOperand returns the value of the Operand field in AlertCondition as
// a AlertConditionOperand value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertCondition) GetOperand(ctx context.Context) (AlertConditionOperand, bool) {
	var e AlertConditionOperand
	if o.Operand.IsNull() || o.Operand.IsUnknown() {
		return e, false
	}
	var v []AlertConditionOperand
	d := o.Operand.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOperand sets the value of the Operand field in AlertCondition.
func (o *AlertCondition) SetOperand(ctx context.Context, v AlertConditionOperand) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["operand"]
	o.Operand = types.ListValueMust(t, vs)
}

// GetThreshold returns the value of the Threshold field in AlertCondition as
// a AlertConditionThreshold value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertCondition) GetThreshold(ctx context.Context) (AlertConditionThreshold, bool) {
	var e AlertConditionThreshold
	if o.Threshold.IsNull() || o.Threshold.IsUnknown() {
		return e, false
	}
	var v []AlertConditionThreshold
	d := o.Threshold.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThreshold sets the value of the Threshold field in AlertCondition.
func (o *AlertCondition) SetThreshold(ctx context.Context, v AlertConditionThreshold) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["threshold"]
	o.Threshold = types.ListValueMust(t, vs)
}

type AlertConditionOperand struct {
	Column types.List `tfsdk:"column" tf:"optional,object"`
}

func (newState *AlertConditionOperand) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionOperand) {
}

func (newState *AlertConditionOperand) SyncEffectiveFieldsDuringRead(existingState AlertConditionOperand) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConditionOperand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConditionOperand) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertOperandColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionOperand
// only implements ToObjectValue() and Type().
func (o AlertConditionOperand) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": o.Column,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertConditionOperand) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column": basetypes.ListType{
				ElemType: AlertOperandColumn{}.Type(ctx),
			},
		},
	}
}

// GetColumn returns the value of the Column field in AlertConditionOperand as
// a AlertOperandColumn value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConditionOperand) GetColumn(ctx context.Context) (AlertOperandColumn, bool) {
	var e AlertOperandColumn
	if o.Column.IsNull() || o.Column.IsUnknown() {
		return e, false
	}
	var v []AlertOperandColumn
	d := o.Column.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetColumn sets the value of the Column field in AlertConditionOperand.
func (o *AlertConditionOperand) SetColumn(ctx context.Context, v AlertOperandColumn) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["column"]
	o.Column = types.ListValueMust(t, vs)
}

type AlertConditionThreshold struct {
	Value types.List `tfsdk:"value" tf:"optional,object"`
}

func (newState *AlertConditionThreshold) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionThreshold) {
}

func (newState *AlertConditionThreshold) SyncEffectiveFieldsDuringRead(existingState AlertConditionThreshold) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConditionThreshold.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConditionThreshold) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(AlertOperandValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionThreshold
// only implements ToObjectValue() and Type().
func (o AlertConditionThreshold) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertConditionThreshold) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: AlertOperandValue{}.Type(ctx),
			},
		},
	}
}

// GetValue returns the value of the Value field in AlertConditionThreshold as
// a AlertOperandValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConditionThreshold) GetValue(ctx context.Context) (AlertOperandValue, bool) {
	var e AlertOperandValue
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v []AlertOperandValue
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in AlertConditionThreshold.
func (o *AlertConditionThreshold) SetValue(ctx context.Context, v AlertOperandValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type AlertOperandColumn struct {
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *AlertOperandColumn) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandColumn) {
}

func (newState *AlertOperandColumn) SyncEffectiveFieldsDuringRead(existingState AlertOperandColumn) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOperandColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOperandColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandColumn
// only implements ToObjectValue() and Type().
func (o AlertOperandColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertOperandColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type AlertOperandValue struct {
	BoolValue types.Bool `tfsdk:"bool_value" tf:"optional"`

	DoubleValue types.Float64 `tfsdk:"double_value" tf:"optional"`

	StringValue types.String `tfsdk:"string_value" tf:"optional"`
}

func (newState *AlertOperandValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandValue) {
}

func (newState *AlertOperandValue) SyncEffectiveFieldsDuringRead(existingState AlertOperandValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOperandValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOperandValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandValue
// only implements ToObjectValue() and Type().
func (o AlertOperandValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"double_value": o.DoubleValue,
			"string_value": o.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertOperandValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"double_value": types.Float64Type,
			"string_value": types.StringType,
		},
	}
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	Column types.String `tfsdk:"column" tf:""`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state" tf:"optional"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted types.Bool `tfsdk:"muted" tf:"optional"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op types.String `tfsdk:"op" tf:""`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value types.Object `tfsdk:"value" tf:""`
}

func (newState *AlertOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOptions) {
}

func (newState *AlertOptions) SyncEffectiveFieldsDuringRead(existingState AlertOptions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOptions
// only implements ToObjectValue() and Type().
func (o AlertOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column":             o.Column,
			"custom_body":        o.CustomBody,
			"custom_subject":     o.CustomSubject,
			"empty_result_state": o.EmptyResultState,
			"muted":              o.Muted,
			"op":                 o.Op,
			"value":              o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column":             types.StringType,
			"custom_body":        types.StringType,
			"custom_subject":     types.StringType,
			"empty_result_state": types.StringType,
			"muted":              types.BoolType,
			"op":                 types.StringType,
			"value":              types.ObjectType{},
		},
	}
}

type AlertQuery struct {
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Query ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.List `tfsdk:"options" tf:"optional,object"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
}

func (newState *AlertQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertQuery) {
}

func (newState *AlertQuery) SyncEffectiveFieldsDuringRead(existingState AlertQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(QueryOptions{}),
		"tags":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertQuery
// only implements ToObjectValue() and Type().
func (o AlertQuery) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":     o.CreatedAt,
			"data_source_id": o.DataSourceId,
			"description":    o.Description,
			"id":             o.Id,
			"is_archived":    o.IsArchived,
			"is_draft":       o.IsDraft,
			"is_safe":        o.IsSafe,
			"name":           o.Name,
			"options":        o.Options,
			"query":          o.Query,
			"tags":           o.Tags,
			"updated_at":     o.UpdatedAt,
			"user_id":        o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertQuery) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":     types.StringType,
			"data_source_id": types.StringType,
			"description":    types.StringType,
			"id":             types.StringType,
			"is_archived":    types.BoolType,
			"is_draft":       types.BoolType,
			"is_safe":        types.BoolType,
			"name":           types.StringType,
			"options": basetypes.ListType{
				ElemType: QueryOptions{}.Type(ctx),
			},
			"query": types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user_id":    types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in AlertQuery as
// a QueryOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertQuery) GetOptions(ctx context.Context) (QueryOptions, bool) {
	var e QueryOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AlertQuery.
func (o *AlertQuery) SetOptions(ctx context.Context, v QueryOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in AlertQuery as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertQuery) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in AlertQuery.
func (o *AlertQuery) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *BaseChunkInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseChunkInfo) {
}

func (newState *BaseChunkInfo) SyncEffectiveFieldsDuringRead(existingState BaseChunkInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseChunkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseChunkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseChunkInfo
// only implements ToObjectValue() and Type().
func (o BaseChunkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":  o.ByteCount,
			"chunk_index": o.ChunkIndex,
			"row_count":   o.RowCount,
			"row_offset":  o.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BaseChunkInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":  types.Int64Type,
			"chunk_index": types.Int64Type,
			"row_count":   types.Int64Type,
			"row_offset":  types.Int64Type,
		},
	}
}

// Cancel statement execution
type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *CancelExecutionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelExecutionRequest) {
}

func (newState *CancelExecutionRequest) SyncEffectiveFieldsDuringRead(existingState CancelExecutionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelExecutionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelExecutionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelExecutionRequest
// only implements ToObjectValue() and Type().
func (o CancelExecutionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelExecutionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
		},
	}
}

type CancelExecutionResponse struct {
}

func (newState *CancelExecutionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelExecutionResponse) {
}

func (newState *CancelExecutionResponse) SyncEffectiveFieldsDuringRead(existingState CancelExecutionResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelExecutionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelExecutionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelExecutionResponse
// only implements ToObjectValue() and Type().
func (o CancelExecutionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelExecutionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version" tf:"optional"`

	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *Channel) SyncEffectiveFieldsDuringCreateOrUpdate(plan Channel) {
}

func (newState *Channel) SyncEffectiveFieldsDuringRead(existingState Channel) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Channel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Channel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Channel
// only implements ToObjectValue() and Type().
func (o Channel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": o.DbsqlVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Channel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbsql_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion types.String `tfsdk:"dbsql_version" tf:"optional"`
	// Name of the channel
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ChannelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChannelInfo) {
}

func (newState *ChannelInfo) SyncEffectiveFieldsDuringRead(existingState ChannelInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChannelInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChannelInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChannelInfo
// only implements ToObjectValue() and Type().
func (o ChannelInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": o.DbsqlVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChannelInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbsql_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ColumnInfo struct {
	// The name of the column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The ordinal position of the column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The format of the interval type.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName types.String `tfsdk:"type_name" tf:"optional"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// The full SQL type specification.
	TypeText types.String `tfsdk:"type_text" tf:"optional"`
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo
// only implements ToObjectValue() and Type().
func (o ColumnInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":               o.Name,
			"position":           o.Position,
			"type_interval_type": o.TypeIntervalType,
			"type_name":          o.TypeName,
			"type_precision":     o.TypePrecision,
			"type_scale":         o.TypeScale,
			"type_text":          o.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ColumnInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":               types.StringType,
			"position":           types.Int64Type,
			"type_interval_type": types.StringType,
			"type_name":          types.StringType,
			"type_precision":     types.Int64Type,
			"type_scale":         types.Int64Type,
			"type_text":          types.StringType,
		},
	}
}

type CreateAlert struct {
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:""`
	// Alert configuration options.
	Options types.List `tfsdk:"options" tf:"object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id" tf:""`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
}

func (newState *CreateAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlert) {
}

func (newState *CreateAlert) SyncEffectiveFieldsDuringRead(existingState CreateAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlert
// only implements ToObjectValue() and Type().
func (o CreateAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"options":  o.Options,
			"parent":   o.Parent,
			"query_id": o.QueryId,
			"rearm":    o.Rearm,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions{}.Type(ctx),
			},
			"parent":   types.StringType,
			"query_id": types.StringType,
			"rearm":    types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in CreateAlert as
// a AlertOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlert) GetOptions(ctx context.Context) (AlertOptions, bool) {
	var e AlertOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateAlert.
func (o *CreateAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type CreateAlertRequest struct {
	Alert types.List `tfsdk:"alert" tf:"optional,object"`
}

func (newState *CreateAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlertRequest) {
}

func (newState *CreateAlertRequest) SyncEffectiveFieldsDuringRead(existingState CreateAlertRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(CreateAlertRequestAlert{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequest
// only implements ToObjectValue() and Type().
func (o CreateAlertRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert": o.Alert,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: CreateAlertRequestAlert{}.Type(ctx),
			},
		},
	}
}

// GetAlert returns the value of the Alert field in CreateAlertRequest as
// a CreateAlertRequestAlert value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertRequest) GetAlert(ctx context.Context) (CreateAlertRequestAlert, bool) {
	var e CreateAlertRequestAlert
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []CreateAlertRequestAlert
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in CreateAlertRequest.
func (o *CreateAlertRequest) SetAlert(ctx context.Context, v CreateAlertRequestAlert) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition" tf:"optional,object"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
}

func (newState *CreateAlertRequestAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlertRequestAlert) {
}

func (newState *CreateAlertRequestAlert) SyncEffectiveFieldsDuringRead(existingState CreateAlertRequestAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertRequestAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertRequestAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequestAlert
// only implements ToObjectValue() and Type().
func (o CreateAlertRequestAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            o.Condition,
			"custom_body":          o.CustomBody,
			"custom_subject":       o.CustomSubject,
			"display_name":         o.DisplayName,
			"notify_on_ok":         o.NotifyOnOk,
			"parent_path":          o.ParentPath,
			"query_id":             o.QueryId,
			"seconds_to_retrigger": o.SecondsToRetrigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertRequestAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition{}.Type(ctx),
			},
			"custom_body":          types.StringType,
			"custom_subject":       types.StringType,
			"display_name":         types.StringType,
			"notify_on_ok":         types.BoolType,
			"parent_path":          types.StringType,
			"query_id":             types.StringType,
			"seconds_to_retrigger": types.Int64Type,
		},
	}
}

// GetCondition returns the value of the Condition field in CreateAlertRequestAlert as
// a AlertCondition value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertRequestAlert) GetCondition(ctx context.Context) (AlertCondition, bool) {
	var e AlertCondition
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in CreateAlertRequestAlert.
func (o *CreateAlertRequestAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type CreateQueryRequest struct {
	Query types.List `tfsdk:"query" tf:"optional,object"`
}

func (newState *CreateQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryRequest) {
}

func (newState *CreateQueryRequest) SyncEffectiveFieldsDuringRead(existingState CreateQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(CreateQueryRequestQuery{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequest
// only implements ToObjectValue() and Type().
func (o CreateQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query": o.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query": basetypes.ListType{
				ElemType: CreateQueryRequestQuery{}.Type(ctx),
			},
		},
	}
}

// GetQuery returns the value of the Query field in CreateQueryRequest as
// a CreateQueryRequestQuery value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequest) GetQuery(ctx context.Context) (CreateQueryRequestQuery, bool) {
	var e CreateQueryRequestQuery
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []CreateQueryRequestQuery
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in CreateQueryRequest.
func (o *CreateQueryRequest) SetQuery(ctx context.Context, v CreateQueryRequestQuery) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateQueryRequestQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryRequestQuery) {
}

func (newState *CreateQueryRequestQuery) SyncEffectiveFieldsDuringRead(existingState CreateQueryRequestQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryRequestQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryRequestQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequestQuery
// only implements ToObjectValue() and Type().
func (o CreateQueryRequestQuery) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit": o.ApplyAutoLimit,
			"catalog":          o.Catalog,
			"description":      o.Description,
			"display_name":     o.DisplayName,
			"parameters":       o.Parameters,
			"parent_path":      o.ParentPath,
			"query_text":       o.QueryText,
			"run_as_mode":      o.RunAsMode,
			"schema":           o.Schema,
			"tags":             o.Tags,
			"warehouse_id":     o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQueryRequestQuery) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit": types.BoolType,
			"catalog":          types.StringType,
			"description":      types.StringType,
			"display_name":     types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter{}.Type(ctx),
			},
			"parent_path": types.StringType,
			"query_text":  types.StringType,
			"run_as_mode": types.StringType,
			"schema":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in CreateQueryRequestQuery as
// a slice of QueryParameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequestQuery) GetParameters(ctx context.Context) ([]QueryParameter, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in CreateQueryRequestQuery.
func (o *CreateQueryRequestQuery) SetParameters(ctx context.Context, v []QueryParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateQueryRequestQuery as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequestQuery) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateQueryRequestQuery.
func (o *CreateQueryRequestQuery) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options types.Object `tfsdk:"options" tf:""`
	// The identifier returned by :method:queries/create
	QueryId types.String `tfsdk:"query_id" tf:""`
	// The type of visualization: chart, table, pivot table, and so on.
	Type_ types.String `tfsdk:"type" tf:""`
}

func (newState *CreateQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryVisualizationsLegacyRequest) {
}

func (newState *CreateQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState CreateQueryVisualizationsLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryVisualizationsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryVisualizationsLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryVisualizationsLegacyRequest
// only implements ToObjectValue() and Type().
func (o CreateQueryVisualizationsLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"options":     o.Options,
			"query_id":    o.QueryId,
			"type":        o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQueryVisualizationsLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"options":     types.ObjectType{},
			"query_id":    types.StringType,
			"type":        types.StringType,
		},
	}
}

type CreateVisualizationRequest struct {
	Visualization types.List `tfsdk:"visualization" tf:"optional,object"`
}

func (newState *CreateVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVisualizationRequest) {
}

func (newState *CreateVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState CreateVisualizationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVisualizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(CreateVisualizationRequestVisualization{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequest
// only implements ToObjectValue() and Type().
func (o CreateVisualizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visualization": o.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVisualizationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"visualization": basetypes.ListType{
				ElemType: CreateVisualizationRequestVisualization{}.Type(ctx),
			},
		},
	}
}

// GetVisualization returns the value of the Visualization field in CreateVisualizationRequest as
// a CreateVisualizationRequestVisualization value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVisualizationRequest) GetVisualization(ctx context.Context) (CreateVisualizationRequestVisualization, bool) {
	var e CreateVisualizationRequestVisualization
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []CreateVisualizationRequestVisualization
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in CreateVisualizationRequest.
func (o *CreateVisualizationRequest) SetVisualization(ctx context.Context, v CreateVisualizationRequestVisualization) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type CreateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *CreateVisualizationRequestVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVisualizationRequestVisualization) {
}

func (newState *CreateVisualizationRequestVisualization) SyncEffectiveFieldsDuringRead(existingState CreateVisualizationRequestVisualization) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVisualizationRequestVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVisualizationRequestVisualization) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequestVisualization
// only implements ToObjectValue() and Type().
func (o CreateVisualizationRequestVisualization) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          o.DisplayName,
			"query_id":              o.QueryId,
			"serialized_options":    o.SerializedOptions,
			"serialized_query_plan": o.SerializedQueryPlan,
			"type":                  o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVisualizationRequestVisualization) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":          types.StringType,
			"query_id":              types.StringType,
			"serialized_options":    types.StringType,
			"serialized_query_plan": types.StringType,
			"type":                  types.StringType,
		},
	}
}

type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be >= 0 mins for serverless warehouses - Must be
	// == 0 or >= 10 mins for non-serverless warehouses - 0 indicates no
	// autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *CreateWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWarehouseRequest) {
}

func (newState *CreateWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState CreateWarehouseRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWarehouseRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel{}),
		"tags":    reflect.TypeOf(EndpointTags{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseRequest
// only implements ToObjectValue() and Type().
func (o CreateWarehouseRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            o.AutoStopMins,
			"channel":                   o.Channel,
			"cluster_size":              o.ClusterSize,
			"creator_name":              o.CreatorName,
			"enable_photon":             o.EnablePhoton,
			"enable_serverless_compute": o.EnableServerlessCompute,
			"instance_profile_arn":      o.InstanceProfileArn,
			"max_num_clusters":          o.MaxNumClusters,
			"min_num_clusters":          o.MinNumClusters,
			"name":                      o.Name,
			"spot_instance_policy":      o.SpotInstancePolicy,
			"tags":                      o.Tags,
			"warehouse_type":            o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWarehouseRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"instance_profile_arn":      types.StringType,
			"max_num_clusters":          types.Int64Type,
			"min_num_clusters":          types.Int64Type,
			"name":                      types.StringType,
			"spot_instance_policy":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in CreateWarehouseRequest as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWarehouseRequest) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in CreateWarehouseRequest.
func (o *CreateWarehouseRequest) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateWarehouseRequest as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWarehouseRequest) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in CreateWarehouseRequest.
func (o *CreateWarehouseRequest) SetTags(ctx context.Context, v EndpointTags) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreateWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWarehouseResponse) {
}

func (newState *CreateWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState CreateWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseResponse
// only implements ToObjectValue() and Type().
func (o CreateWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`

	Options types.List `tfsdk:"options" tf:"object"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text" tf:"optional"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id" tf:"optional"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width" tf:""`
}

func (newState *CreateWidget) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWidget) {
}

func (newState *CreateWidget) SyncEffectiveFieldsDuringRead(existingState CreateWidget) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWidget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWidget) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWidget
// only implements ToObjectValue() and Type().
func (o CreateWidget) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":     o.DashboardId,
			"id":               o.Id,
			"options":          o.Options,
			"text":             o.Text,
			"visualization_id": o.VisualizationId,
			"width":            o.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWidget) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"options": basetypes.ListType{
				ElemType: WidgetOptions{}.Type(ctx),
			},
			"text":             types.StringType,
			"visualization_id": types.StringType,
			"width":            types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in CreateWidget as
// a WidgetOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWidget) GetOptions(ctx context.Context) (WidgetOptions, bool) {
	var e WidgetOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateWidget.
func (o *CreateWidget) SetOptions(ctx context.Context, v WidgetOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit types.Bool `tfsdk:"can_edit" tf:"optional"`
	// Timestamp when this dashboard was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled" tf:"optional"`
	// The ID for this dashboard.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.List `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier" tf:"optional"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug types.String `tfsdk:"slug" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.List `tfsdk:"user" tf:"optional,object"`
	// The ID of the user who owns the dashboard.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`

	Widgets types.List `tfsdk:"widgets" tf:"optional"`
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
	return map[string]reflect.Type{
		"options": reflect.TypeOf(DashboardOptions{}),
		"tags":    reflect.TypeOf(types.String{}),
		"user":    reflect.TypeOf(User{}),
		"widgets": reflect.TypeOf(Widget{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dashboard
// only implements ToObjectValue() and Type().
func (o Dashboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_edit":                  o.CanEdit,
			"created_at":                o.CreatedAt,
			"dashboard_filters_enabled": o.DashboardFiltersEnabled,
			"id":                        o.Id,
			"is_archived":               o.IsArchived,
			"is_draft":                  o.IsDraft,
			"is_favorite":               o.IsFavorite,
			"name":                      o.Name,
			"options":                   o.Options,
			"parent":                    o.Parent,
			"permission_tier":           o.PermissionTier,
			"slug":                      o.Slug,
			"tags":                      o.Tags,
			"updated_at":                o.UpdatedAt,
			"user":                      o.User,
			"user_id":                   o.UserId,
			"widgets":                   o.Widgets,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dashboard) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_edit":                  types.BoolType,
			"created_at":                types.StringType,
			"dashboard_filters_enabled": types.BoolType,
			"id":                        types.StringType,
			"is_archived":               types.BoolType,
			"is_draft":                  types.BoolType,
			"is_favorite":               types.BoolType,
			"name":                      types.StringType,
			"options": basetypes.ListType{
				ElemType: DashboardOptions{}.Type(ctx),
			},
			"parent":          types.StringType,
			"permission_tier": types.StringType,
			"slug":            types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
			"user_id": types.Int64Type,
			"widgets": basetypes.ListType{
				ElemType: Widget{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in Dashboard as
// a DashboardOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard) GetOptions(ctx context.Context) (DashboardOptions, bool) {
	var e DashboardOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []DashboardOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Dashboard.
func (o *Dashboard) SetOptions(ctx context.Context, v DashboardOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Dashboard as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Dashboard.
func (o *Dashboard) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in Dashboard as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard) GetUser(ctx context.Context) (User, bool) {
	var e User
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in Dashboard.
func (o *Dashboard) SetUser(ctx context.Context, v User) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

// GetWidgets returns the value of the Widgets field in Dashboard as
// a slice of Widget values.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard) GetWidgets(ctx context.Context) ([]Widget, bool) {
	if o.Widgets.IsNull() || o.Widgets.IsUnknown() {
		return nil, false
	}
	var v []Widget
	d := o.Widgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWidgets sets the value of the Widgets field in Dashboard.
func (o *Dashboard) SetWidgets(ctx context.Context, v []Widget) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["widgets"]
	o.Widgets = types.ListValueMust(t, vs)
}

type DashboardEditContent struct {
	DashboardId types.String `tfsdk:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DashboardEditContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardEditContent) {
}

func (newState *DashboardEditContent) SyncEffectiveFieldsDuringRead(existingState DashboardEditContent) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardEditContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardEditContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEditContent
// only implements ToObjectValue() and Type().
func (o DashboardEditContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
			"name":         o.Name,
			"run_as_role":  o.RunAsRole,
			"tags":         o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DashboardEditContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"name":         types.StringType,
			"run_as_role":  types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTags returns the value of the Tags field in DashboardEditContent as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DashboardEditContent) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DashboardEditContent.
func (o *DashboardEditContent) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at" tf:"optional"`
}

func (newState *DashboardOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardOptions) {
}

func (newState *DashboardOptions) SyncEffectiveFieldsDuringRead(existingState DashboardOptions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardOptions
// only implements ToObjectValue() and Type().
func (o DashboardOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"moved_to_trash_at": o.MovedToTrashAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DashboardOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"moved_to_trash_at": types.StringType,
		},
	}
}

type DashboardPostContent struct {
	// Indicates whether the dashboard filters are enabled
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled" tf:"optional"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:""`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DashboardPostContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardPostContent) {
}

func (newState *DashboardPostContent) SyncEffectiveFieldsDuringRead(existingState DashboardPostContent) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardPostContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardPostContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardPostContent
// only implements ToObjectValue() and Type().
func (o DashboardPostContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_filters_enabled": o.DashboardFiltersEnabled,
			"is_favorite":               o.IsFavorite,
			"name":                      o.Name,
			"parent":                    o.Parent,
			"run_as_role":               o.RunAsRole,
			"tags":                      o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DashboardPostContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_filters_enabled": types.BoolType,
			"is_favorite":               types.BoolType,
			"name":                      types.StringType,
			"parent":                    types.StringType,
			"run_as_role":               types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTags returns the value of the Tags field in DashboardPostContent as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DashboardPostContent) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DashboardPostContent.
func (o *DashboardPostContent) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id types.String `tfsdk:"id" tf:"optional"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Reserved for internal use.
	PauseReason types.String `tfsdk:"pause_reason" tf:"optional"`
	// Reserved for internal use.
	Paused types.Int64 `tfsdk:"paused" tf:"optional"`
	// Reserved for internal use.
	SupportsAutoLimit types.Bool `tfsdk:"supports_auto_limit" tf:"optional"`
	// Reserved for internal use.
	Syntax types.String `tfsdk:"syntax" tf:"optional"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type_ types.String `tfsdk:"type" tf:"optional"`
	// Reserved for internal use.
	ViewOnly types.Bool `tfsdk:"view_only" tf:"optional"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *DataSource) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataSource) {
}

func (newState *DataSource) SyncEffectiveFieldsDuringRead(existingState DataSource) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataSource
// only implements ToObjectValue() and Type().
func (o DataSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":                  o.Id,
			"name":                o.Name,
			"pause_reason":        o.PauseReason,
			"paused":              o.Paused,
			"supports_auto_limit": o.SupportsAutoLimit,
			"syntax":              o.Syntax,
			"type":                o.Type_,
			"view_only":           o.ViewOnly,
			"warehouse_id":        o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                  types.StringType,
			"name":                types.StringType,
			"pause_reason":        types.StringType,
			"paused":              types.Int64Type,
			"supports_auto_limit": types.BoolType,
			"syntax":              types.StringType,
			"type":                types.StringType,
			"view_only":           types.BoolType,
			"warehouse_id":        types.StringType,
		},
	}
}

type DateRange struct {
	End types.String `tfsdk:"end" tf:""`

	Start types.String `tfsdk:"start" tf:""`
}

func (newState *DateRange) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRange) {
}

func (newState *DateRange) SyncEffectiveFieldsDuringRead(existingState DateRange) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateRange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRange
// only implements ToObjectValue() and Type().
func (o DateRange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end":   o.End,
			"start": o.Start,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DateRange) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end":   types.StringType,
			"start": types.StringType,
		},
	}
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	DateRangeValue types.List `tfsdk:"date_range_value" tf:"optional,object"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue types.String `tfsdk:"dynamic_date_range_value" tf:"optional"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision" tf:"optional"`

	StartDayOfWeek types.Int64 `tfsdk:"start_day_of_week" tf:"optional"`
}

func (newState *DateRangeValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRangeValue) {
}

func (newState *DateRangeValue) SyncEffectiveFieldsDuringRead(existingState DateRangeValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateRangeValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateRangeValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"date_range_value": reflect.TypeOf(DateRange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRangeValue
// only implements ToObjectValue() and Type().
func (o DateRangeValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_range_value":         o.DateRangeValue,
			"dynamic_date_range_value": o.DynamicDateRangeValue,
			"precision":                o.Precision,
			"start_day_of_week":        o.StartDayOfWeek,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DateRangeValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_range_value": basetypes.ListType{
				ElemType: DateRange{}.Type(ctx),
			},
			"dynamic_date_range_value": types.StringType,
			"precision":                types.StringType,
			"start_day_of_week":        types.Int64Type,
		},
	}
}

// GetDateRangeValue returns the value of the DateRangeValue field in DateRangeValue as
// a DateRange value.
// If the field is unknown or null, the boolean return value is false.
func (o *DateRangeValue) GetDateRangeValue(ctx context.Context) (DateRange, bool) {
	var e DateRange
	if o.DateRangeValue.IsNull() || o.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRange
	d := o.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in DateRangeValue.
func (o *DateRangeValue) SetDateRangeValue(ctx context.Context, v DateRange) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	o.DateRangeValue = types.ListValueMust(t, vs)
}

type DateValue struct {
	// Manually specified date-time value.
	DateValue types.String `tfsdk:"date_value" tf:"optional"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue types.String `tfsdk:"dynamic_date_value" tf:"optional"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision" tf:"optional"`
}

func (newState *DateValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateValue) {
}

func (newState *DateValue) SyncEffectiveFieldsDuringRead(existingState DateValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateValue
// only implements ToObjectValue() and Type().
func (o DateValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_value":         o.DateValue,
			"dynamic_date_value": o.DynamicDateValue,
			"precision":          o.Precision,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DateValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_value":         types.StringType,
			"dynamic_date_value": types.StringType,
			"precision":          types.StringType,
		},
	}
}

// Delete an alert
type DeleteAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (newState *DeleteAlertsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAlertsLegacyRequest) {
}

func (newState *DeleteAlertsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAlertsLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAlertsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAlertsLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAlertsLegacyRequest
// only implements ToObjectValue() and Type().
func (o DeleteAlertsLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": o.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAlertsLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

// Remove a dashboard
type DeleteDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (newState *DeleteDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDashboardRequest) {
}

func (newState *DeleteDashboardRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardRequest
// only implements ToObjectValue() and Type().
func (o DeleteDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Remove widget
type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteDashboardWidgetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDashboardWidgetRequest) {
}

func (newState *DeleteDashboardWidgetRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDashboardWidgetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardWidgetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDashboardWidgetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardWidgetRequest
// only implements ToObjectValue() and Type().
func (o DeleteDashboardWidgetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardWidgetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a query
type DeleteQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *DeleteQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQueriesLegacyRequest) {
}

func (newState *DeleteQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQueriesLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQueriesLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueriesLegacyRequest
// only implements ToObjectValue() and Type().
func (o DeleteQueriesLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQueriesLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

// Remove visualization
type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvizualisations/create
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQueryVisualizationsLegacyRequest) {
}

func (newState *DeleteQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQueryVisualizationsLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQueryVisualizationsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQueryVisualizationsLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueryVisualizationsLegacyRequest
// only implements ToObjectValue() and Type().
func (o DeleteQueryVisualizationsLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQueryVisualizationsLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (o DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Remove a visualization
type DeleteVisualizationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVisualizationRequest) {
}

func (newState *DeleteVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVisualizationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVisualizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVisualizationRequest
// only implements ToObjectValue() and Type().
func (o DeleteVisualizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVisualizationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a warehouse
type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWarehouseRequest) {
}

func (newState *DeleteWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWarehouseRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWarehouseRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseRequest
// only implements ToObjectValue() and Type().
func (o DeleteWarehouseRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWarehouseRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWarehouseResponse struct {
}

func (newState *DeleteWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWarehouseResponse) {
}

func (newState *DeleteWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseResponse
// only implements ToObjectValue() and Type().
func (o DeleteWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditAlert struct {
	AlertId types.String `tfsdk:"-"`
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:""`
	// Alert configuration options.
	Options types.List `tfsdk:"options" tf:"object"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id" tf:""`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
}

func (newState *EditAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditAlert) {
}

func (newState *EditAlert) SyncEffectiveFieldsDuringRead(existingState EditAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditAlert
// only implements ToObjectValue() and Type().
func (o EditAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": o.AlertId,
			"name":     o.Name,
			"options":  o.Options,
			"query_id": o.QueryId,
			"rearm":    o.Rearm,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
			"name":     types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions{}.Type(ctx),
			},
			"query_id": types.StringType,
			"rearm":    types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in EditAlert as
// a AlertOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditAlert) GetOptions(ctx context.Context) (AlertOptions, bool) {
	var e AlertOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in EditAlert.
func (o *EditAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Required. Id of the warehouse to configure.
	Id types.String `tfsdk:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *EditWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditWarehouseRequest) {
}

func (newState *EditWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState EditWarehouseRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditWarehouseRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel{}),
		"tags":    reflect.TypeOf(EndpointTags{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseRequest
// only implements ToObjectValue() and Type().
func (o EditWarehouseRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            o.AutoStopMins,
			"channel":                   o.Channel,
			"cluster_size":              o.ClusterSize,
			"creator_name":              o.CreatorName,
			"enable_photon":             o.EnablePhoton,
			"enable_serverless_compute": o.EnableServerlessCompute,
			"id":                        o.Id,
			"instance_profile_arn":      o.InstanceProfileArn,
			"max_num_clusters":          o.MaxNumClusters,
			"min_num_clusters":          o.MinNumClusters,
			"name":                      o.Name,
			"spot_instance_policy":      o.SpotInstancePolicy,
			"tags":                      o.Tags,
			"warehouse_type":            o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditWarehouseRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"id":                        types.StringType,
			"instance_profile_arn":      types.StringType,
			"max_num_clusters":          types.Int64Type,
			"min_num_clusters":          types.Int64Type,
			"name":                      types.StringType,
			"spot_instance_policy":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in EditWarehouseRequest as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditWarehouseRequest) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EditWarehouseRequest.
func (o *EditWarehouseRequest) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EditWarehouseRequest as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditWarehouseRequest) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EditWarehouseRequest.
func (o *EditWarehouseRequest) SetTags(ctx context.Context, v EndpointTags) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type EditWarehouseResponse struct {
}

func (newState *EditWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditWarehouseResponse) {
}

func (newState *EditWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState EditWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseResponse
// only implements ToObjectValue() and Type().
func (o EditWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty struct {
}

func (newState *Empty) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty) {
}

func (newState *Empty) SyncEffectiveFieldsDuringRead(existingState Empty) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Empty) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty
// only implements ToObjectValue() and Type().
func (o Empty) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o Empty) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EndpointConfPair struct {
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointConfPair) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointConfPair) {
}

func (newState *EndpointConfPair) SyncEffectiveFieldsDuringRead(existingState EndpointConfPair) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointConfPair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointConfPair) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointConfPair
// only implements ToObjectValue() and Type().
func (o EndpointConfPair) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointConfPair) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details types.String `tfsdk:"details" tf:"optional"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason types.List `tfsdk:"failure_reason" tf:"optional,object"`
	// Deprecated. split into summary and details for security
	Message types.String `tfsdk:"message" tf:"optional"`
	// Health status of the warehouse.
	Status types.String `tfsdk:"status" tf:"optional"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary" tf:"optional"`
}

func (newState *EndpointHealth) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointHealth) {
}

func (newState *EndpointHealth) SyncEffectiveFieldsDuringRead(existingState EndpointHealth) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointHealth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointHealth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failure_reason": reflect.TypeOf(TerminationReason{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointHealth
// only implements ToObjectValue() and Type().
func (o EndpointHealth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":        o.Details,
			"failure_reason": o.FailureReason,
			"message":        o.Message,
			"status":         o.Status,
			"summary":        o.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointHealth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": types.StringType,
			"failure_reason": basetypes.ListType{
				ElemType: TerminationReason{}.Type(ctx),
			},
			"message": types.StringType,
			"status":  types.StringType,
			"summary": types.StringType,
		},
	}
}

// GetFailureReason returns the value of the FailureReason field in EndpointHealth as
// a TerminationReason value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointHealth) GetFailureReason(ctx context.Context) (TerminationReason, bool) {
	var e TerminationReason
	if o.FailureReason.IsNull() || o.FailureReason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason
	d := o.FailureReason.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailureReason sets the value of the FailureReason field in EndpointHealth.
func (o *EndpointHealth) SetFailureReason(ctx context.Context, v TerminationReason) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failure_reason"]
	o.FailureReason = types.ListValueMust(t, vs)
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions" tf:"optional"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters" tf:"optional"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.List `tfsdk:"odbc_params" tf:"optional,object"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// State of the warehouse
	State types.String `tfsdk:"state" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointInfo) {
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringRead(existingState EndpointInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":     reflect.TypeOf(Channel{}),
		"health":      reflect.TypeOf(EndpointHealth{}),
		"odbc_params": reflect.TypeOf(OdbcParams{}),
		"tags":        reflect.TypeOf(EndpointTags{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo
// only implements ToObjectValue() and Type().
func (o EndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            o.AutoStopMins,
			"channel":                   o.Channel,
			"cluster_size":              o.ClusterSize,
			"creator_name":              o.CreatorName,
			"enable_photon":             o.EnablePhoton,
			"enable_serverless_compute": o.EnableServerlessCompute,
			"health":                    o.Health,
			"id":                        o.Id,
			"instance_profile_arn":      o.InstanceProfileArn,
			"jdbc_url":                  o.JdbcUrl,
			"max_num_clusters":          o.MaxNumClusters,
			"min_num_clusters":          o.MinNumClusters,
			"name":                      o.Name,
			"num_active_sessions":       o.NumActiveSessions,
			"num_clusters":              o.NumClusters,
			"odbc_params":               o.OdbcParams,
			"spot_instance_policy":      o.SpotInstancePolicy,
			"state":                     o.State,
			"tags":                      o.Tags,
			"warehouse_type":            o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health": basetypes.ListType{
				ElemType: EndpointHealth{}.Type(ctx),
			},
			"id":                   types.StringType,
			"instance_profile_arn": types.StringType,
			"jdbc_url":             types.StringType,
			"max_num_clusters":     types.Int64Type,
			"min_num_clusters":     types.Int64Type,
			"name":                 types.StringType,
			"num_active_sessions":  types.Int64Type,
			"num_clusters":         types.Int64Type,
			"odbc_params": basetypes.ListType{
				ElemType: OdbcParams{}.Type(ctx),
			},
			"spot_instance_policy": types.StringType,
			"state":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in EndpointInfo as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EndpointInfo.
func (o *EndpointInfo) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in EndpointInfo as
// a EndpointHealth value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetHealth(ctx context.Context) (EndpointHealth, bool) {
	var e EndpointHealth
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in EndpointInfo.
func (o *EndpointInfo) SetHealth(ctx context.Context, v EndpointHealth) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in EndpointInfo as
// a OdbcParams value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetOdbcParams(ctx context.Context) (OdbcParams, bool) {
	var e OdbcParams
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams
	d := o.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in EndpointInfo.
func (o *EndpointInfo) SetOdbcParams(ctx context.Context, v OdbcParams) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	o.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EndpointInfo as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EndpointInfo.
func (o *EndpointInfo) SetTags(ctx context.Context, v EndpointTags) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type EndpointTagPair struct {
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointTagPair) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTagPair) {
}

func (newState *EndpointTagPair) SyncEffectiveFieldsDuringRead(existingState EndpointTagPair) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTagPair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointTagPair) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTagPair
// only implements ToObjectValue() and Type().
func (o EndpointTagPair) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTagPair) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointTags struct {
	CustomTags types.List `tfsdk:"custom_tags" tf:"optional"`
}

func (newState *EndpointTags) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTags) {
}

func (newState *EndpointTags) SyncEffectiveFieldsDuringRead(existingState EndpointTags) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTags.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointTags) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(EndpointTagPair{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTags
// only implements ToObjectValue() and Type().
func (o EndpointTags) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags": o.CustomTags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTags) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: EndpointTagPair{}.Type(ctx),
			},
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EndpointTags as
// a slice of EndpointTagPair values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointTags) GetCustomTags(ctx context.Context) ([]EndpointTagPair, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTagPair
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EndpointTags.
func (o *EndpointTags) SetCustomTags(ctx context.Context, v []EndpointTagPair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	o.CustomTags = types.ListValueMust(t, vs)
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions types.String `tfsdk:"enum_options" tf:"optional"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options" tf:"optional,object"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *EnumValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnumValue) {
}

func (newState *EnumValue) SyncEffectiveFieldsDuringRead(existingState EnumValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnumValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnumValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnumValue
// only implements ToObjectValue() and Type().
func (o EnumValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enum_options":         o.EnumOptions,
			"multi_values_options": o.MultiValuesOptions,
			"values":               o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnumValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enum_options": types.StringType,
			"multi_values_options": basetypes.ListType{
				ElemType: MultiValuesOptions{}.Type(ctx),
			},
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in EnumValue as
// a MultiValuesOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *EnumValue) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions, bool) {
	var e MultiValuesOptions
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in EnumValue.
func (o *EnumValue) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in EnumValue as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EnumValue) GetValues(ctx context.Context) ([]types.String, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in EnumValue.
func (o *EnumValue) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	o.Values = types.ListValueMust(t, vs)
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit" tf:"optional"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog" tf:"optional"`

	Disposition types.String `tfsdk:"disposition" tf:"optional"`
	// Statement execution supports three result formats: `JSON_ARRAY`
	// (default), `ARROW_STREAM`, and `CSV`.
	//
	// Important: The formats `ARROW_STREAM` and `CSV` are supported only with
	// `EXTERNAL_LINKS` disposition. `JSON_ARRAY` is supported in `INLINE` and
	// `EXTERNAL_LINKS` disposition.
	//
	// When specifying `format=JSON_ARRAY`, result data will be formatted as an
	// array of arrays of values, where each value is either the *string
	// representation* of a value, or `null`. For example, the output of `SELECT
	// concat('id-', id) AS strCol, id AS intCol, null AS nullCol FROM range(3)`
	// would look like this:
	//
	// ``` [ [ "id-1", "1", null ], [ "id-2", "2", null ], [ "id-3", "3", null
	// ], ] ```
	//
	// When specifying `format=JSON_ARRAY` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result contains compact JSON with no indentation or
	// extra whitespace.
	//
	// When specifying `format=ARROW_STREAM` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result will be formatted as Apache Arrow Stream. See
	// the [Apache Arrow streaming format].
	//
	// When specifying `format=CSV` and `disposition=EXTERNAL_LINKS`, each chunk
	// in the result will be a CSV according to [RFC 4180] standard. All the
	// columns values will have *string representation* similar to the
	// `JSON_ARRAY` format, and `null` values will be encoded as “null”.
	// Only the first chunk in the result would contain a header row with column
	// names. For example, the output of `SELECT concat('id-', id) AS strCol, id
	// AS intCol, null as nullCol FROM range(3)` would look like this:
	//
	// ``` strCol,intCol,nullCol id-1,1,null id-2,2,null id-3,3,null ```
	//
	// [Apache Arrow streaming format]: https://arrow.apache.org/docs/format/Columnar.html#ipc-streaming-format
	// [RFC 4180]: https://www.rfc-editor.org/rfc/rfc4180
	Format types.String `tfsdk:"format" tf:"optional"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	OnWaitTimeout types.String `tfsdk:"on_wait_timeout" tf:"optional"`
	// A list of parameters to pass into a SQL statement containing parameter
	// markers. A parameter consists of a name, a value, and optionally a type.
	// To represent a NULL value, the `value` field may be omitted or set to
	// `null` explicitly. If the `type` field is omitted, the value is
	// interpreted as a string.
	//
	// If the type is given, parameters will be checked for type correctness
	// according to the given type. A value is correct if the provided string
	// can be converted to the requested type using the `cast` function. The
	// exact semantics are described in the section [`cast` function] of the SQL
	// language reference.
	//
	// For example, the following statement contains two parameters, `my_name`
	// and `my_date`:
	//
	// SELECT * FROM my_table WHERE name = :my_name AND date = :my_date
	//
	// The parameters can be passed in the request body as follows:
	//
	// { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	//
	// Currently, positional parameters denoted by a `?` marker are not
	// supported by the Databricks SQL Statement Execution API.
	//
	// Also see the section [Parameter markers] of the SQL language reference.
	//
	// [Parameter markers]: https://docs.databricks.com/sql/language-manual/sql-ref-parameter-marker.html
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit types.Int64 `tfsdk:"row_limit" tf:"optional"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
	Statement types.String `tfsdk:"statement" tf:""`
	// The time in seconds the call will wait for the statement's result set as
	// `Ns`, where `N` can be set to 0 or to a value between 5 and 50.
	//
	// When set to `0s`, the statement will execute in asynchronous mode and the
	// call will not wait for the execution to finish. In this case, the call
	// returns directly with `PENDING` state and a statement ID which can be
	// used for polling with :method:statementexecution/getStatement.
	//
	// When set between 5 and 50 seconds, the call will behave synchronously up
	// to this timeout and wait for the statement execution to finish. If the
	// execution finishes within this time, the call returns immediately with a
	// manifest and result data (or a `FAILED` state in case of an execution
	// error). If the statement takes longer to execute, `on_wait_timeout`
	// determines what should happen after the timeout is reached.
	WaitTimeout types.String `tfsdk:"wait_timeout" tf:"optional"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId types.String `tfsdk:"warehouse_id" tf:""`
}

func (newState *ExecuteStatementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExecuteStatementRequest) {
}

func (newState *ExecuteStatementRequest) SyncEffectiveFieldsDuringRead(existingState ExecuteStatementRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExecuteStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExecuteStatementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(StatementParameterListItem{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteStatementRequest
// only implements ToObjectValue() and Type().
func (o ExecuteStatementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_limit":      o.ByteLimit,
			"catalog":         o.Catalog,
			"disposition":     o.Disposition,
			"format":          o.Format,
			"on_wait_timeout": o.OnWaitTimeout,
			"parameters":      o.Parameters,
			"row_limit":       o.RowLimit,
			"schema":          o.Schema,
			"statement":       o.Statement,
			"wait_timeout":    o.WaitTimeout,
			"warehouse_id":    o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExecuteStatementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_limit":      types.Int64Type,
			"catalog":         types.StringType,
			"disposition":     types.StringType,
			"format":          types.StringType,
			"on_wait_timeout": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: StatementParameterListItem{}.Type(ctx),
			},
			"row_limit":    types.Int64Type,
			"schema":       types.StringType,
			"statement":    types.StringType,
			"wait_timeout": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in ExecuteStatementRequest as
// a slice of StatementParameterListItem values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExecuteStatementRequest) GetParameters(ctx context.Context) ([]StatementParameterListItem, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []StatementParameterListItem
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ExecuteStatementRequest.
func (o *ExecuteStatementRequest) SetParameters(ctx context.Context, v []StatementParameterListItem) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration types.String `tfsdk:"expiration" tf:"optional"`

	ExternalLink types.String `tfsdk:"external_link" tf:"optional"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers" tf:"optional"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index" tf:"optional"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *ExternalLink) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLink) {
}

func (newState *ExternalLink) SyncEffectiveFieldsDuringRead(existingState ExternalLink) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalLink.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalLink) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"http_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLink
// only implements ToObjectValue() and Type().
func (o ExternalLink) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               o.ByteCount,
			"chunk_index":              o.ChunkIndex,
			"expiration":               o.Expiration,
			"external_link":            o.ExternalLink,
			"http_headers":             o.HttpHeaders,
			"next_chunk_index":         o.NextChunkIndex,
			"next_chunk_internal_link": o.NextChunkInternalLink,
			"row_count":                o.RowCount,
			"row_offset":               o.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalLink) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":    types.Int64Type,
			"chunk_index":   types.Int64Type,
			"expiration":    types.StringType,
			"external_link": types.StringType,
			"http_headers": basetypes.MapType{
				ElemType: types.StringType,
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetHttpHeaders returns the value of the HttpHeaders field in ExternalLink as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalLink) GetHttpHeaders(ctx context.Context) (map[string]types.String, bool) {
	if o.HttpHeaders.IsNull() || o.HttpHeaders.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.HttpHeaders.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpHeaders sets the value of the HttpHeaders field in ExternalLink.
func (o *ExternalLink) SetHttpHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["http_headers"]
	o.HttpHeaders = types.MapValueMust(t, vs)
}

// Get an alert
type GetAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAlertRequest) {
}

func (newState *GetAlertRequest) SyncEffectiveFieldsDuringRead(existingState GetAlertRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertRequest
// only implements ToObjectValue() and Type().
func (o GetAlertRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get an alert
type GetAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (newState *GetAlertsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAlertsLegacyRequest) {
}

func (newState *GetAlertsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState GetAlertsLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertsLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertsLegacyRequest
// only implements ToObjectValue() and Type().
func (o GetAlertsLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": o.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertsLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

// Retrieve a definition
type GetDashboardRequest struct {
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

// Get object ACL
type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permissions to check.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *GetDbsqlPermissionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDbsqlPermissionRequest) {
}

func (newState *GetDbsqlPermissionRequest) SyncEffectiveFieldsDuringRead(existingState GetDbsqlPermissionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDbsqlPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDbsqlPermissionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDbsqlPermissionRequest
// only implements ToObjectValue() and Type().
func (o GetDbsqlPermissionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objectId":   o.ObjectId,
			"objectType": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDbsqlPermissionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"objectId":   types.StringType,
			"objectType": types.StringType,
		},
	}
}

// Get a query definition.
type GetQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *GetQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQueriesLegacyRequest) {
}

func (newState *GetQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState GetQueriesLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQueriesLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueriesLegacyRequest
// only implements ToObjectValue() and Type().
func (o GetQueriesLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQueriesLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

// Get a query
type GetQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQueryRequest) {
}

func (newState *GetQueryRequest) SyncEffectiveFieldsDuringRead(existingState GetQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueryRequest
// only implements ToObjectValue() and Type().
func (o GetQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id" tf:"optional"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *GetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetResponse) {
}

func (newState *GetResponse) SyncEffectiveFieldsDuringRead(existingState GetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResponse
// only implements ToObjectValue() and Type().
func (o GetResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in GetResponse as
// a slice of AccessControl values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetResponse) GetAccessControlList(ctx context.Context) ([]AccessControl, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in GetResponse.
func (o *GetResponse) SetAccessControlList(ctx context.Context, v []AccessControl) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	o.AccessControlList = types.ListValueMust(t, vs)
}

// Get status, manifest, and result first chunk
type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *GetStatementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatementRequest) {
}

func (newState *GetStatementRequest) SyncEffectiveFieldsDuringRead(existingState GetStatementRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementRequest
// only implements ToObjectValue() and Type().
func (o GetStatementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
		},
	}
}

// Get result chunk by index
type GetStatementResultChunkNRequest struct {
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *GetStatementResultChunkNRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatementResultChunkNRequest) {
}

func (newState *GetStatementResultChunkNRequest) SyncEffectiveFieldsDuringRead(existingState GetStatementResultChunkNRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatementResultChunkNRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatementResultChunkNRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementResultChunkNRequest
// only implements ToObjectValue() and Type().
func (o GetStatementResultChunkNRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_index":  o.ChunkIndex,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatementResultChunkNRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunk_index":  types.Int64Type,
			"statement_id": types.StringType,
		},
	}
}

// Get SQL warehouse permission levels
type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *GetWarehousePermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionLevelsRequest) {
}

func (newState *GetWarehousePermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionLevelsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
		},
	}
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetWarehousePermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionLevelsResponse) {
}

func (newState *GetWarehousePermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WarehousePermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: WarehousePermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetWarehousePermissionLevelsResponse as
// a slice of WarehousePermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehousePermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]WarehousePermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWarehousePermissionLevelsResponse.
func (o *GetWarehousePermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []WarehousePermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get SQL warehouse permissions
type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *GetWarehousePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionsRequest) {
}

func (newState *GetWarehousePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
		},
	}
}

// Get warehouse info
type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *GetWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehouseRequest) {
}

func (newState *GetWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehouseRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehouseRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehouseRequest
// only implements ToObjectValue() and Type().
func (o GetWarehouseRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehouseRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions" tf:"optional"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters" tf:"optional"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.List `tfsdk:"odbc_params" tf:"optional,object"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// State of the warehouse
	State types.String `tfsdk:"state" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *GetWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehouseResponse) {
}

func (newState *GetWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState GetWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":     reflect.TypeOf(Channel{}),
		"health":      reflect.TypeOf(EndpointHealth{}),
		"odbc_params": reflect.TypeOf(OdbcParams{}),
		"tags":        reflect.TypeOf(EndpointTags{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehouseResponse
// only implements ToObjectValue() and Type().
func (o GetWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            o.AutoStopMins,
			"channel":                   o.Channel,
			"cluster_size":              o.ClusterSize,
			"creator_name":              o.CreatorName,
			"enable_photon":             o.EnablePhoton,
			"enable_serverless_compute": o.EnableServerlessCompute,
			"health":                    o.Health,
			"id":                        o.Id,
			"instance_profile_arn":      o.InstanceProfileArn,
			"jdbc_url":                  o.JdbcUrl,
			"max_num_clusters":          o.MaxNumClusters,
			"min_num_clusters":          o.MinNumClusters,
			"name":                      o.Name,
			"num_active_sessions":       o.NumActiveSessions,
			"num_clusters":              o.NumClusters,
			"odbc_params":               o.OdbcParams,
			"spot_instance_policy":      o.SpotInstancePolicy,
			"state":                     o.State,
			"tags":                      o.Tags,
			"warehouse_type":            o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health": basetypes.ListType{
				ElemType: EndpointHealth{}.Type(ctx),
			},
			"id":                   types.StringType,
			"instance_profile_arn": types.StringType,
			"jdbc_url":             types.StringType,
			"max_num_clusters":     types.Int64Type,
			"min_num_clusters":     types.Int64Type,
			"name":                 types.StringType,
			"num_active_sessions":  types.Int64Type,
			"num_clusters":         types.Int64Type,
			"odbc_params": basetypes.ListType{
				ElemType: OdbcParams{}.Type(ctx),
			},
			"spot_instance_policy": types.StringType,
			"state":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in GetWarehouseResponse as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in GetWarehouseResponse as
// a EndpointHealth value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetHealth(ctx context.Context) (EndpointHealth, bool) {
	var e EndpointHealth
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetHealth(ctx context.Context, v EndpointHealth) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in GetWarehouseResponse as
// a OdbcParams value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetOdbcParams(ctx context.Context) (OdbcParams, bool) {
	var e OdbcParams
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams
	d := o.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetOdbcParams(ctx context.Context, v OdbcParams) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	o.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in GetWarehouseResponse as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetTags(ctx context.Context, v EndpointTags) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.List `tfsdk:"config_param" tf:"optional,object"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config" tf:"optional"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types" tf:"optional"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.List `tfsdk:"global_param" tf:"optional,object"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy" tf:"optional"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters" tf:"optional,object"`
}

func (newState *GetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceWarehouseConfigResponse) {
}

func (newState *GetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceWarehouseConfigResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceWarehouseConfigResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceWarehouseConfigResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":                      reflect.TypeOf(Channel{}),
		"config_param":                 reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"data_access_config":           reflect.TypeOf(EndpointConfPair{}),
		"enabled_warehouse_types":      reflect.TypeOf(WarehouseTypePair{}),
		"global_param":                 reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"sql_configuration_parameters": reflect.TypeOf(RepeatedEndpointConfPairs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceWarehouseConfigResponse
// only implements ToObjectValue() and Type().
func (o GetWorkspaceWarehouseConfigResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel":                      o.Channel,
			"config_param":                 o.ConfigParam,
			"data_access_config":           o.DataAccessConfig,
			"enabled_warehouse_types":      o.EnabledWarehouseTypes,
			"global_param":                 o.GlobalParam,
			"google_service_account":       o.GoogleServiceAccount,
			"instance_profile_arn":         o.InstanceProfileArn,
			"security_policy":              o.SecurityPolicy,
			"sql_configuration_parameters": o.SqlConfigurationParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceWarehouseConfigResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"config_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair{}.Type(ctx),
			},
			"global_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
			"google_service_account": types.StringType,
			"instance_profile_arn":   types.StringType,
			"security_policy":        types.StringType,
			"sql_configuration_parameters": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
		},
	}
}

// GetChannel returns the value of the Channel field in GetWorkspaceWarehouseConfigResponse as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	o.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse as
// a slice of EndpointConfPair values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair, bool) {
	if o.DataAccessConfig.IsNull() || o.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair
	d := o.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	o.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse as
// a slice of WarehouseTypePair values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair, bool) {
	if o.EnabledWarehouseTypes.IsNull() || o.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair
	d := o.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	o.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.GlobalParam.IsNull() || o.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	o.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	o.SqlConfigurationParameters = types.ListValueMust(t, vs)
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Alert ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt types.String `tfsdk:"last_triggered_at" tf:"optional"`
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Alert configuration options.
	Options types.List `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`

	Query types.List `tfsdk:"query" tf:"optional,object"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.List `tfsdk:"user" tf:"optional,object"`
}

func (newState *LegacyAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyAlert) {
}

func (newState *LegacyAlert) SyncEffectiveFieldsDuringRead(existingState LegacyAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions{}),
		"query":   reflect.TypeOf(AlertQuery{}),
		"user":    reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyAlert
// only implements ToObjectValue() and Type().
func (o LegacyAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":        o.CreatedAt,
			"id":                o.Id,
			"last_triggered_at": o.LastTriggeredAt,
			"name":              o.Name,
			"options":           o.Options,
			"parent":            o.Parent,
			"query":             o.Query,
			"rearm":             o.Rearm,
			"state":             o.State,
			"updated_at":        o.UpdatedAt,
			"user":              o.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LegacyAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":        types.StringType,
			"id":                types.StringType,
			"last_triggered_at": types.StringType,
			"name":              types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions{}.Type(ctx),
			},
			"parent": types.StringType,
			"query": basetypes.ListType{
				ElemType: AlertQuery{}.Type(ctx),
			},
			"rearm":      types.Int64Type,
			"state":      types.StringType,
			"updated_at": types.StringType,
			"user": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in LegacyAlert as
// a AlertOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert) GetOptions(ctx context.Context) (AlertOptions, bool) {
	var e AlertOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyAlert.
func (o *LegacyAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetQuery returns the value of the Query field in LegacyAlert as
// a AlertQuery value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert) GetQuery(ctx context.Context) (AlertQuery, bool) {
	var e AlertQuery
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []AlertQuery
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyAlert.
func (o *LegacyAlert) SetQuery(ctx context.Context, v AlertQuery) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyAlert as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert) GetUser(ctx context.Context) (User, bool) {
	var e User
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyAlert.
func (o *LegacyAlert) SetUser(ctx context.Context, v User) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

type LegacyQuery struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit types.Bool `tfsdk:"can_edit" tf:"optional"`
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Query ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe" tf:"optional"`

	LastModifiedBy types.List `tfsdk:"last_modified_by" tf:"optional,object"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById types.Int64 `tfsdk:"last_modified_by_id" tf:"optional"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId types.String `tfsdk:"latest_query_data_id" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.List `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash types.String `tfsdk:"query_hash" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.List `tfsdk:"user" tf:"optional,object"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`

	Visualizations types.List `tfsdk:"visualizations" tf:"optional"`
}

func (newState *LegacyQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyQuery) {
}

func (newState *LegacyQuery) SyncEffectiveFieldsDuringRead(existingState LegacyQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"last_modified_by": reflect.TypeOf(User{}),
		"options":          reflect.TypeOf(QueryOptions{}),
		"tags":             reflect.TypeOf(types.String{}),
		"user":             reflect.TypeOf(User{}),
		"visualizations":   reflect.TypeOf(LegacyVisualization{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyQuery
// only implements ToObjectValue() and Type().
func (o LegacyQuery) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_edit":             o.CanEdit,
			"created_at":           o.CreatedAt,
			"data_source_id":       o.DataSourceId,
			"description":          o.Description,
			"id":                   o.Id,
			"is_archived":          o.IsArchived,
			"is_draft":             o.IsDraft,
			"is_favorite":          o.IsFavorite,
			"is_safe":              o.IsSafe,
			"last_modified_by":     o.LastModifiedBy,
			"last_modified_by_id":  o.LastModifiedById,
			"latest_query_data_id": o.LatestQueryDataId,
			"name":                 o.Name,
			"options":              o.Options,
			"parent":               o.Parent,
			"permission_tier":      o.PermissionTier,
			"query":                o.Query,
			"query_hash":           o.QueryHash,
			"run_as_role":          o.RunAsRole,
			"tags":                 o.Tags,
			"updated_at":           o.UpdatedAt,
			"user":                 o.User,
			"user_id":              o.UserId,
			"visualizations":       o.Visualizations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LegacyQuery) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_edit":       types.BoolType,
			"created_at":     types.StringType,
			"data_source_id": types.StringType,
			"description":    types.StringType,
			"id":             types.StringType,
			"is_archived":    types.BoolType,
			"is_draft":       types.BoolType,
			"is_favorite":    types.BoolType,
			"is_safe":        types.BoolType,
			"last_modified_by": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
			"last_modified_by_id":  types.Int64Type,
			"latest_query_data_id": types.StringType,
			"name":                 types.StringType,
			"options": basetypes.ListType{
				ElemType: QueryOptions{}.Type(ctx),
			},
			"parent":          types.StringType,
			"permission_tier": types.StringType,
			"query":           types.StringType,
			"query_hash":      types.StringType,
			"run_as_role":     types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
			"user_id": types.Int64Type,
			"visualizations": basetypes.ListType{
				ElemType: LegacyVisualization{}.Type(ctx),
			},
		},
	}
}

// GetLastModifiedBy returns the value of the LastModifiedBy field in LegacyQuery as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetLastModifiedBy(ctx context.Context) (User, bool) {
	var e User
	if o.LastModifiedBy.IsNull() || o.LastModifiedBy.IsUnknown() {
		return e, false
	}
	var v []User
	d := o.LastModifiedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLastModifiedBy sets the value of the LastModifiedBy field in LegacyQuery.
func (o *LegacyQuery) SetLastModifiedBy(ctx context.Context, v User) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["last_modified_by"]
	o.LastModifiedBy = types.ListValueMust(t, vs)
}

// GetOptions returns the value of the Options field in LegacyQuery as
// a QueryOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetOptions(ctx context.Context) (QueryOptions, bool) {
	var e QueryOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyQuery.
func (o *LegacyQuery) SetOptions(ctx context.Context, v QueryOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LegacyQuery as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LegacyQuery.
func (o *LegacyQuery) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyQuery as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetUser(ctx context.Context) (User, bool) {
	var e User
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyQuery.
func (o *LegacyQuery) SetUser(ctx context.Context, v User) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

// GetVisualizations returns the value of the Visualizations field in LegacyQuery as
// a slice of LegacyVisualization values.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetVisualizations(ctx context.Context) ([]LegacyVisualization, bool) {
	if o.Visualizations.IsNull() || o.Visualizations.IsUnknown() {
		return nil, false
	}
	var v []LegacyVisualization
	d := o.Visualizations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualizations sets the value of the Visualizations field in LegacyQuery.
func (o *LegacyQuery) SetVisualizations(ctx context.Context, v []LegacyVisualization) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualizations"]
	o.Visualizations = types.ListValueMust(t, vs)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The UUID for this visualization.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options types.Object `tfsdk:"options" tf:"optional"`

	Query types.List `tfsdk:"query" tf:"optional,object"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type_ types.String `tfsdk:"type" tf:"optional"`

	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
}

func (newState *LegacyVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyVisualization) {
}

func (newState *LegacyVisualization) SyncEffectiveFieldsDuringRead(existingState LegacyVisualization) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyVisualization) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(LegacyQuery{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyVisualization
// only implements ToObjectValue() and Type().
func (o LegacyVisualization) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":  o.CreatedAt,
			"description": o.Description,
			"id":          o.Id,
			"name":        o.Name,
			"options":     o.Options,
			"query":       o.Query,
			"type":        o.Type_,
			"updated_at":  o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LegacyVisualization) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":  types.StringType,
			"description": types.StringType,
			"id":          types.StringType,
			"name":        types.StringType,
			"options":     types.ObjectType{},
			"query": basetypes.ListType{
				ElemType: LegacyQuery{}.Type(ctx),
			},
			"type":       types.StringType,
			"updated_at": types.StringType,
		},
	}
}

// GetQuery returns the value of the Query field in LegacyVisualization as
// a LegacyQuery value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyVisualization) GetQuery(ctx context.Context) (LegacyQuery, bool) {
	var e LegacyQuery
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []LegacyQuery
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyVisualization.
func (o *LegacyVisualization) SetQuery(ctx context.Context, v LegacyQuery) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

// List alerts
type ListAlertsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAlertsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsRequest) {
}

func (newState *ListAlertsRequest) SyncEffectiveFieldsDuringRead(existingState ListAlertsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsRequest
// only implements ToObjectValue() and Type().
func (o ListAlertsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAlertsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListAlertsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponse) {
}

func (newState *ListAlertsResponse) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListAlertsResponseAlert{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponse
// only implements ToObjectValue() and Type().
func (o ListAlertsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListAlertsResponseAlert{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListAlertsResponse as
// a slice of ListAlertsResponseAlert values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsResponse) GetResults(ctx context.Context) ([]ListAlertsResponseAlert, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListAlertsResponseAlert
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListAlertsResponse.
func (o *ListAlertsResponse) SetResults(ctx context.Context, v []ListAlertsResponseAlert) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition" tf:"optional,object"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time" tf:"optional"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *ListAlertsResponseAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponseAlert) {
}

func (newState *ListAlertsResponseAlert) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponseAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsResponseAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsResponseAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponseAlert
// only implements ToObjectValue() and Type().
func (o ListAlertsResponseAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            o.Condition,
			"create_time":          o.CreateTime,
			"custom_body":          o.CustomBody,
			"custom_subject":       o.CustomSubject,
			"display_name":         o.DisplayName,
			"id":                   o.Id,
			"lifecycle_state":      o.LifecycleState,
			"notify_on_ok":         o.NotifyOnOk,
			"owner_user_name":      o.OwnerUserName,
			"query_id":             o.QueryId,
			"seconds_to_retrigger": o.SecondsToRetrigger,
			"state":                o.State,
			"trigger_time":         o.TriggerTime,
			"update_time":          o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsResponseAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition{}.Type(ctx),
			},
			"create_time":          types.StringType,
			"custom_body":          types.StringType,
			"custom_subject":       types.StringType,
			"display_name":         types.StringType,
			"id":                   types.StringType,
			"lifecycle_state":      types.StringType,
			"notify_on_ok":         types.BoolType,
			"owner_user_name":      types.StringType,
			"query_id":             types.StringType,
			"seconds_to_retrigger": types.Int64Type,
			"state":                types.StringType,
			"trigger_time":         types.StringType,
			"update_time":          types.StringType,
		},
	}
}

// GetCondition returns the value of the Condition field in ListAlertsResponseAlert as
// a AlertCondition value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsResponseAlert) GetCondition(ctx context.Context) (AlertCondition, bool) {
	var e AlertCondition
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in ListAlertsResponseAlert.
func (o *ListAlertsResponseAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

// Get dashboard objects
type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order types.String `tfsdk:"-"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-"`
	// Number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// Full text search term.
	Q types.String `tfsdk:"-"`
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
			"order":     o.Order,
			"page":      o.Page,
			"page_size": o.PageSize,
			"q":         o.Q,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"order":     types.StringType,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"q":         types.StringType,
		},
	}
}

// Get a list of queries
type ListQueriesLegacyRequest struct {
	// Name of query attribute to order by. Default sort order is ascending.
	// Append a dash (`-`) to order descending instead.
	//
	// - `name`: The name of the query.
	//
	// - `created_at`: The timestamp the query was created.
	//
	// - `runtime`: The time it took to run this query. This is blank for
	// parameterized queries. A blank value is treated as the highest value for
	// sorting.
	//
	// - `executed_at`: The timestamp when the query was last run.
	//
	// - `created_by`: The user name of the user that created the query.
	Order types.String `tfsdk:"-"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-"`
	// Number of queries to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// Full text search term
	Q types.String `tfsdk:"-"`
}

func (newState *ListQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesLegacyRequest) {
}

func (newState *ListQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState ListQueriesLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesLegacyRequest
// only implements ToObjectValue() and Type().
func (o ListQueriesLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"order":     o.Order,
			"page":      o.Page,
			"page_size": o.PageSize,
			"q":         o.Q,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueriesLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"order":     types.StringType,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"q":         types.StringType,
		},
	}
}

// List queries
type ListQueriesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQueriesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesRequest) {
}

func (newState *ListQueriesRequest) SyncEffectiveFieldsDuringRead(existingState ListQueriesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesRequest
// only implements ToObjectValue() and Type().
func (o ListQueriesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueriesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	HasNextPage types.Bool `tfsdk:"has_next_page" tf:"optional"`
	// A token that can be used to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Res types.List `tfsdk:"res" tf:"optional"`
}

func (newState *ListQueriesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesResponse) {
}

func (newState *ListQueriesResponse) SyncEffectiveFieldsDuringRead(existingState ListQueriesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"res": reflect.TypeOf(QueryInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesResponse
// only implements ToObjectValue() and Type().
func (o ListQueriesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_next_page":   o.HasNextPage,
			"next_page_token": o.NextPageToken,
			"res":             o.Res,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueriesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_next_page":   types.BoolType,
			"next_page_token": types.StringType,
			"res": basetypes.ListType{
				ElemType: QueryInfo{}.Type(ctx),
			},
		},
	}
}

// GetRes returns the value of the Res field in ListQueriesResponse as
// a slice of QueryInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueriesResponse) GetRes(ctx context.Context) ([]QueryInfo, bool) {
	if o.Res.IsNull() || o.Res.IsUnknown() {
		return nil, false
	}
	var v []QueryInfo
	d := o.Res.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRes sets the value of the Res field in ListQueriesResponse.
func (o *ListQueriesResponse) SetRes(ctx context.Context, v []QueryInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["res"]
	o.Res = types.ListValueMust(t, vs)
}

// List Queries
type ListQueryHistoryRequest struct {
	// A filter to limit query history results. This field is optional.
	FilterBy types.List `tfsdk:"-"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics types.Bool `tfsdk:"-"`
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	MaxResults types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQueryHistoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryHistoryRequest) {
}

func (newState *ListQueryHistoryRequest) SyncEffectiveFieldsDuringRead(existingState ListQueryHistoryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryHistoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryHistoryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(QueryFilter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryHistoryRequest
// only implements ToObjectValue() and Type().
func (o ListQueryHistoryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":       o.FilterBy,
			"include_metrics": o.IncludeMetrics,
			"max_results":     o.MaxResults,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueryHistoryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by": basetypes.ListType{
				ElemType: QueryFilter{}.Type(ctx),
			},
			"include_metrics": types.BoolType,
			"max_results":     types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListQueryHistoryRequest as
// a QueryFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryHistoryRequest) GetFilterBy(ctx context.Context) (QueryFilter, bool) {
	var e QueryFilter
	if o.FilterBy.IsNull() || o.FilterBy.IsUnknown() {
		return e, false
	}
	var v []QueryFilter
	d := o.FilterBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListQueryHistoryRequest.
func (o *ListQueryHistoryRequest) SetFilterBy(ctx context.Context, v QueryFilter) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter_by"]
	o.FilterBy = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListQueryObjectsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponse) {
}

func (newState *ListQueryObjectsResponse) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryObjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryObjectsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListQueryObjectsResponseQuery{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponse
// only implements ToObjectValue() and Type().
func (o ListQueryObjectsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueryObjectsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListQueryObjectsResponseQuery{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListQueryObjectsResponse as
// a slice of ListQueryObjectsResponseQuery values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponse) GetResults(ctx context.Context) ([]ListQueryObjectsResponseQuery, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListQueryObjectsResponseQuery
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListQueryObjectsResponse.
func (o *ListQueryObjectsResponse) SetResults(ctx context.Context, v []ListQueryObjectsResponseQuery) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name" tf:"optional"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *ListQueryObjectsResponseQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponseQuery) {
}

func (newState *ListQueryObjectsResponseQuery) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponseQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryObjectsResponseQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryObjectsResponseQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponseQuery
// only implements ToObjectValue() and Type().
func (o ListQueryObjectsResponseQuery) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit":        o.ApplyAutoLimit,
			"catalog":                 o.Catalog,
			"create_time":             o.CreateTime,
			"description":             o.Description,
			"display_name":            o.DisplayName,
			"id":                      o.Id,
			"last_modifier_user_name": o.LastModifierUserName,
			"lifecycle_state":         o.LifecycleState,
			"owner_user_name":         o.OwnerUserName,
			"parameters":              o.Parameters,
			"query_text":              o.QueryText,
			"run_as_mode":             o.RunAsMode,
			"schema":                  o.Schema,
			"tags":                    o.Tags,
			"update_time":             o.UpdateTime,
			"warehouse_id":            o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueryObjectsResponseQuery) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit":        types.BoolType,
			"catalog":                 types.StringType,
			"create_time":             types.StringType,
			"description":             types.StringType,
			"display_name":            types.StringType,
			"id":                      types.StringType,
			"last_modifier_user_name": types.StringType,
			"lifecycle_state":         types.StringType,
			"owner_user_name":         types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter{}.Type(ctx),
			},
			"query_text":  types.StringType,
			"run_as_mode": types.StringType,
			"schema":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in ListQueryObjectsResponseQuery as
// a slice of QueryParameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponseQuery) GetParameters(ctx context.Context) ([]QueryParameter, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ListQueryObjectsResponseQuery.
func (o *ListQueryObjectsResponseQuery) SetParameters(ctx context.Context, v []QueryParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListQueryObjectsResponseQuery as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponseQuery) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListQueryObjectsResponseQuery.
func (o *ListQueryObjectsResponseQuery) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type ListResponse struct {
	// The total number of dashboards.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// The current page being displayed.
	Page types.Int64 `tfsdk:"page" tf:"optional"`
	// The number of dashboards per page.
	PageSize types.Int64 `tfsdk:"page_size" tf:"optional"`
	// List of dashboards returned.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListResponse) {
}

func (newState *ListResponse) SyncEffectiveFieldsDuringRead(existingState ListResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Dashboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse
// only implements ToObjectValue() and Type().
func (o ListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"count":     o.Count,
			"page":      o.Page,
			"page_size": o.PageSize,
			"results":   o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":     types.Int64Type,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"results": basetypes.ListType{
				ElemType: Dashboard{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListResponse as
// a slice of Dashboard values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListResponse) GetResults(ctx context.Context) ([]Dashboard, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []Dashboard
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListResponse.
func (o *ListResponse) SetResults(ctx context.Context, v []Dashboard) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

// List visualizations on a query
type ListVisualizationsForQueryRequest struct {
	Id types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListVisualizationsForQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVisualizationsForQueryRequest) {
}

func (newState *ListVisualizationsForQueryRequest) SyncEffectiveFieldsDuringRead(existingState ListVisualizationsForQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVisualizationsForQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVisualizationsForQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryRequest
// only implements ToObjectValue() and Type().
func (o ListVisualizationsForQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVisualizationsForQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListVisualizationsForQueryResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListVisualizationsForQueryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVisualizationsForQueryResponse) {
}

func (newState *ListVisualizationsForQueryResponse) SyncEffectiveFieldsDuringRead(existingState ListVisualizationsForQueryResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVisualizationsForQueryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVisualizationsForQueryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Visualization{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryResponse
// only implements ToObjectValue() and Type().
func (o ListVisualizationsForQueryResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVisualizationsForQueryResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: Visualization{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListVisualizationsForQueryResponse as
// a slice of Visualization values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListVisualizationsForQueryResponse) GetResults(ctx context.Context) ([]Visualization, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []Visualization
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListVisualizationsForQueryResponse.
func (o *ListVisualizationsForQueryResponse) SetResults(ctx context.Context, v []Visualization) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

// List warehouses
type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId types.Int64 `tfsdk:"-"`
}

func (newState *ListWarehousesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWarehousesRequest) {
}

func (newState *ListWarehousesRequest) SyncEffectiveFieldsDuringRead(existingState ListWarehousesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWarehousesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWarehousesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesRequest
// only implements ToObjectValue() and Type().
func (o ListWarehousesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_as_user_id": o.RunAsUserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWarehousesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_as_user_id": types.Int64Type,
		},
	}
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	Warehouses types.List `tfsdk:"warehouses" tf:"optional"`
}

func (newState *ListWarehousesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWarehousesResponse) {
}

func (newState *ListWarehousesResponse) SyncEffectiveFieldsDuringRead(existingState ListWarehousesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWarehousesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWarehousesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"warehouses": reflect.TypeOf(EndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesResponse
// only implements ToObjectValue() and Type().
func (o ListWarehousesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouses": o.Warehouses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWarehousesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouses": basetypes.ListType{
				ElemType: EndpointInfo{}.Type(ctx),
			},
		},
	}
}

// GetWarehouses returns the value of the Warehouses field in ListWarehousesResponse as
// a slice of EndpointInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWarehousesResponse) GetWarehouses(ctx context.Context) ([]EndpointInfo, bool) {
	if o.Warehouses.IsNull() || o.Warehouses.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo
	d := o.Warehouses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarehouses sets the value of the Warehouses field in ListWarehousesResponse.
func (o *ListWarehousesResponse) SetWarehouses(ctx context.Context, v []EndpointInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouses"]
	o.Warehouses = types.ListValueMust(t, vs)
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	Prefix types.String `tfsdk:"prefix" tf:"optional"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator types.String `tfsdk:"separator" tf:"optional"`
	// Character that suffixes each selected parameter value.
	Suffix types.String `tfsdk:"suffix" tf:"optional"`
}

func (newState *MultiValuesOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan MultiValuesOptions) {
}

func (newState *MultiValuesOptions) SyncEffectiveFieldsDuringRead(existingState MultiValuesOptions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MultiValuesOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MultiValuesOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MultiValuesOptions
// only implements ToObjectValue() and Type().
func (o MultiValuesOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"prefix":    o.Prefix,
			"separator": o.Separator,
			"suffix":    o.Suffix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MultiValuesOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"prefix":    types.StringType,
			"separator": types.StringType,
			"suffix":    types.StringType,
		},
	}
}

type NumericValue struct {
	Value types.Float64 `tfsdk:"value" tf:"optional"`
}

func (newState *NumericValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan NumericValue) {
}

func (newState *NumericValue) SyncEffectiveFieldsDuringRead(existingState NumericValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NumericValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NumericValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NumericValue
// only implements ToObjectValue() and Type().
func (o NumericValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NumericValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Float64Type,
		},
	}
}

type OdbcParams struct {
	Hostname types.String `tfsdk:"hostname" tf:"optional"`

	Path types.String `tfsdk:"path" tf:"optional"`

	Port types.Int64 `tfsdk:"port" tf:"optional"`

	Protocol types.String `tfsdk:"protocol" tf:"optional"`
}

func (newState *OdbcParams) SyncEffectiveFieldsDuringCreateOrUpdate(plan OdbcParams) {
}

func (newState *OdbcParams) SyncEffectiveFieldsDuringRead(existingState OdbcParams) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OdbcParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OdbcParams) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OdbcParams
// only implements ToObjectValue() and Type().
func (o OdbcParams) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hostname": o.Hostname,
			"path":     o.Path,
			"port":     o.Port,
			"protocol": o.Protocol,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OdbcParams) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hostname": types.StringType,
			"path":     types.StringType,
			"port":     types.Int64Type,
			"protocol": types.StringType,
		},
	}
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions types.String `tfsdk:"enumOptions" tf:"optional"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions types.List `tfsdk:"multiValuesOptions" tf:"optional,object"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId types.String `tfsdk:"queryId" tf:"optional"`
	// The text displayed in a parameter picking widget.
	Title types.String `tfsdk:"title" tf:"optional"`
	// Parameters can have several different types.
	Type_ types.String `tfsdk:"type" tf:"optional"`
	// The default value for this parameter.
	Value types.Object `tfsdk:"value" tf:"optional"`
}

func (newState *Parameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan Parameter) {
}

func (newState *Parameter) SyncEffectiveFieldsDuringRead(existingState Parameter) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Parameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Parameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multiValuesOptions": reflect.TypeOf(MultiValuesOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Parameter
// only implements ToObjectValue() and Type().
func (o Parameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enumOptions":        o.EnumOptions,
			"multiValuesOptions": o.MultiValuesOptions,
			"name":               o.Name,
			"queryId":            o.QueryId,
			"title":              o.Title,
			"type":               o.Type_,
			"value":              o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Parameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enumOptions": types.StringType,
			"multiValuesOptions": basetypes.ListType{
				ElemType: MultiValuesOptions{}.Type(ctx),
			},
			"name":    types.StringType,
			"queryId": types.StringType,
			"title":   types.StringType,
			"type":    types.StringType,
			"value":   types.ObjectType{},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in Parameter as
// a MultiValuesOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *Parameter) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions, bool) {
	var e MultiValuesOptions
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in Parameter.
func (o *Parameter) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multiValuesOptions"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name" tf:"optional"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Query) SyncEffectiveFieldsDuringCreateOrUpdate(plan Query) {
}

func (newState *Query) SyncEffectiveFieldsDuringRead(existingState Query) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Query.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Query) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Query
// only implements ToObjectValue() and Type().
func (o Query) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit":        o.ApplyAutoLimit,
			"catalog":                 o.Catalog,
			"create_time":             o.CreateTime,
			"description":             o.Description,
			"display_name":            o.DisplayName,
			"id":                      o.Id,
			"last_modifier_user_name": o.LastModifierUserName,
			"lifecycle_state":         o.LifecycleState,
			"owner_user_name":         o.OwnerUserName,
			"parameters":              o.Parameters,
			"parent_path":             o.ParentPath,
			"query_text":              o.QueryText,
			"run_as_mode":             o.RunAsMode,
			"schema":                  o.Schema,
			"tags":                    o.Tags,
			"update_time":             o.UpdateTime,
			"warehouse_id":            o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Query) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit":        types.BoolType,
			"catalog":                 types.StringType,
			"create_time":             types.StringType,
			"description":             types.StringType,
			"display_name":            types.StringType,
			"id":                      types.StringType,
			"last_modifier_user_name": types.StringType,
			"lifecycle_state":         types.StringType,
			"owner_user_name":         types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter{}.Type(ctx),
			},
			"parent_path": types.StringType,
			"query_text":  types.StringType,
			"run_as_mode": types.StringType,
			"schema":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in Query as
// a slice of QueryParameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *Query) GetParameters(ctx context.Context) ([]QueryParameter, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in Query.
func (o *Query) SetParameters(ctx context.Context, v []QueryParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Query as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Query) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Query.
func (o *Query) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options" tf:"optional,object"`
	// UUID of the query that provides the parameter values.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *QueryBackedValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryBackedValue) {
}

func (newState *QueryBackedValue) SyncEffectiveFieldsDuringRead(existingState QueryBackedValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryBackedValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryBackedValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryBackedValue
// only implements ToObjectValue() and Type().
func (o QueryBackedValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"multi_values_options": o.MultiValuesOptions,
			"query_id":             o.QueryId,
			"values":               o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryBackedValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"multi_values_options": basetypes.ListType{
				ElemType: MultiValuesOptions{}.Type(ctx),
			},
			"query_id": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in QueryBackedValue as
// a MultiValuesOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryBackedValue) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions, bool) {
	var e MultiValuesOptions
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in QueryBackedValue.
func (o *QueryBackedValue) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in QueryBackedValue as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryBackedValue) GetValues(ctx context.Context) ([]types.String, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in QueryBackedValue.
func (o *QueryBackedValue) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	o.Values = types.ListValueMust(t, vs)
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options types.Object `tfsdk:"options" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`

	QueryId types.String `tfsdk:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *QueryEditContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEditContent) {
}

func (newState *QueryEditContent) SyncEffectiveFieldsDuringRead(existingState QueryEditContent) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEditContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryEditContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEditContent
// only implements ToObjectValue() and Type().
func (o QueryEditContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_source_id": o.DataSourceId,
			"description":    o.Description,
			"name":           o.Name,
			"options":        o.Options,
			"query":          o.Query,
			"query_id":       o.QueryId,
			"run_as_role":    o.RunAsRole,
			"tags":           o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryEditContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_source_id": types.StringType,
			"description":    types.StringType,
			"name":           types.StringType,
			"options":        types.ObjectType{},
			"query":          types.StringType,
			"query_id":       types.StringType,
			"run_as_role":    types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTags returns the value of the Tags field in QueryEditContent as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEditContent) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in QueryEditContent.
func (o *QueryEditContent) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be <= 30
	// days.
	QueryStartTimeRange types.List `tfsdk:"query_start_time_range" tf:"optional,object"`
	// A list of statement IDs.
	StatementIds types.List `tfsdk:"statement_ids" tf:"optional"`

	Statuses types.List `tfsdk:"statuses" tf:"optional"`
	// A list of user IDs who ran the queries.
	UserIds types.List `tfsdk:"user_ids" tf:"optional"`
	// A list of warehouse IDs.
	WarehouseIds types.List `tfsdk:"warehouse_ids" tf:"optional"`
}

func (newState *QueryFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryFilter) {
}

func (newState *QueryFilter) SyncEffectiveFieldsDuringRead(existingState QueryFilter) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_start_time_range": reflect.TypeOf(TimeRange{}),
		"statement_ids":          reflect.TypeOf(types.String{}),
		"statuses":               reflect.TypeOf(types.String{}),
		"user_ids":               reflect.TypeOf(types.Int64{}),
		"warehouse_ids":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryFilter
// only implements ToObjectValue() and Type().
func (o QueryFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_start_time_range": o.QueryStartTimeRange,
			"statement_ids":          o.StatementIds,
			"statuses":               o.Statuses,
			"user_ids":               o.UserIds,
			"warehouse_ids":          o.WarehouseIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryFilter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_start_time_range": basetypes.ListType{
				ElemType: TimeRange{}.Type(ctx),
			},
			"statement_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"statuses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"user_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"warehouse_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetQueryStartTimeRange returns the value of the QueryStartTimeRange field in QueryFilter as
// a TimeRange value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter) GetQueryStartTimeRange(ctx context.Context) (TimeRange, bool) {
	var e TimeRange
	if o.QueryStartTimeRange.IsNull() || o.QueryStartTimeRange.IsUnknown() {
		return e, false
	}
	var v []TimeRange
	d := o.QueryStartTimeRange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryStartTimeRange sets the value of the QueryStartTimeRange field in QueryFilter.
func (o *QueryFilter) SetQueryStartTimeRange(ctx context.Context, v TimeRange) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_start_time_range"]
	o.QueryStartTimeRange = types.ListValueMust(t, vs)
}

// GetStatementIds returns the value of the StatementIds field in QueryFilter as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter) GetStatementIds(ctx context.Context) ([]types.String, bool) {
	if o.StatementIds.IsNull() || o.StatementIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.StatementIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatementIds sets the value of the StatementIds field in QueryFilter.
func (o *QueryFilter) SetStatementIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_ids"]
	o.StatementIds = types.ListValueMust(t, vs)
}

// GetStatuses returns the value of the Statuses field in QueryFilter as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter) GetStatuses(ctx context.Context) ([]types.String, bool) {
	if o.Statuses.IsNull() || o.Statuses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in QueryFilter.
func (o *QueryFilter) SetStatuses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	o.Statuses = types.ListValueMust(t, vs)
}

// GetUserIds returns the value of the UserIds field in QueryFilter as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter) GetUserIds(ctx context.Context) ([]types.Int64, bool) {
	if o.UserIds.IsNull() || o.UserIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.UserIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserIds sets the value of the UserIds field in QueryFilter.
func (o *QueryFilter) SetUserIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_ids"]
	o.UserIds = types.ListValueMust(t, vs)
}

// GetWarehouseIds returns the value of the WarehouseIds field in QueryFilter as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter) GetWarehouseIds(ctx context.Context) ([]types.String, bool) {
	if o.WarehouseIds.IsNull() || o.WarehouseIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.WarehouseIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarehouseIds sets the value of the WarehouseIds field in QueryFilter.
func (o *QueryFilter) SetWarehouseIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouse_ids"]
	o.WarehouseIds = types.ListValueMust(t, vs)
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed types.List `tfsdk:"channel_used" tf:"optional,object"`
	// Total execution time of the statement ( excluding result fetch time ).
	Duration types.Int64 `tfsdk:"duration" tf:"optional"`
	// Alias for `warehouse_id`.
	EndpointId types.String `tfsdk:"endpoint_id" tf:"optional"`
	// Message describing why the query could not complete.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId types.Int64 `tfsdk:"executed_as_user_id" tf:"optional"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName types.String `tfsdk:"executed_as_user_name" tf:"optional"`
	// The time execution of the query ended.
	ExecutionEndTimeMs types.Int64 `tfsdk:"execution_end_time_ms" tf:"optional"`
	// Whether more updates for the query are expected.
	IsFinal types.Bool `tfsdk:"is_final" tf:"optional"`
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key" tf:"optional"`
	// Metrics about query execution.
	Metrics types.List `tfsdk:"metrics" tf:"optional,object"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState types.String `tfsdk:"plans_state" tf:"optional"`
	// The time the query ended.
	QueryEndTimeMs types.Int64 `tfsdk:"query_end_time_ms" tf:"optional"`
	// The query ID.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The time the query started.
	QueryStartTimeMs types.Int64 `tfsdk:"query_start_time_ms" tf:"optional"`
	// The text of the query.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// The number of results returned by the query.
	RowsProduced types.Int64 `tfsdk:"rows_produced" tf:"optional"`
	// URL to the Spark UI query plan.
	SparkUiUrl types.String `tfsdk:"spark_ui_url" tf:"optional"`
	// Type of statement for this query
	StatementType types.String `tfsdk:"statement_type" tf:"optional"`
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The ID of the user who ran the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
	// The email address or username of the user who ran the query.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
	// Warehouse ID.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *QueryInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryInfo) {
}

func (newState *QueryInfo) SyncEffectiveFieldsDuringRead(existingState QueryInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel_used": reflect.TypeOf(ChannelInfo{}),
		"metrics":      reflect.TypeOf(QueryMetrics{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryInfo
// only implements ToObjectValue() and Type().
func (o QueryInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel_used":          o.ChannelUsed,
			"duration":              o.Duration,
			"endpoint_id":           o.EndpointId,
			"error_message":         o.ErrorMessage,
			"executed_as_user_id":   o.ExecutedAsUserId,
			"executed_as_user_name": o.ExecutedAsUserName,
			"execution_end_time_ms": o.ExecutionEndTimeMs,
			"is_final":              o.IsFinal,
			"lookup_key":            o.LookupKey,
			"metrics":               o.Metrics,
			"plans_state":           o.PlansState,
			"query_end_time_ms":     o.QueryEndTimeMs,
			"query_id":              o.QueryId,
			"query_start_time_ms":   o.QueryStartTimeMs,
			"query_text":            o.QueryText,
			"rows_produced":         o.RowsProduced,
			"spark_ui_url":          o.SparkUiUrl,
			"statement_type":        o.StatementType,
			"status":                o.Status,
			"user_id":               o.UserId,
			"user_name":             o.UserName,
			"warehouse_id":          o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel_used": basetypes.ListType{
				ElemType: ChannelInfo{}.Type(ctx),
			},
			"duration":              types.Int64Type,
			"endpoint_id":           types.StringType,
			"error_message":         types.StringType,
			"executed_as_user_id":   types.Int64Type,
			"executed_as_user_name": types.StringType,
			"execution_end_time_ms": types.Int64Type,
			"is_final":              types.BoolType,
			"lookup_key":            types.StringType,
			"metrics": basetypes.ListType{
				ElemType: QueryMetrics{}.Type(ctx),
			},
			"plans_state":         types.StringType,
			"query_end_time_ms":   types.Int64Type,
			"query_id":            types.StringType,
			"query_start_time_ms": types.Int64Type,
			"query_text":          types.StringType,
			"rows_produced":       types.Int64Type,
			"spark_ui_url":        types.StringType,
			"statement_type":      types.StringType,
			"status":              types.StringType,
			"user_id":             types.Int64Type,
			"user_name":           types.StringType,
			"warehouse_id":        types.StringType,
		},
	}
}

// GetChannelUsed returns the value of the ChannelUsed field in QueryInfo as
// a ChannelInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo) GetChannelUsed(ctx context.Context) (ChannelInfo, bool) {
	var e ChannelInfo
	if o.ChannelUsed.IsNull() || o.ChannelUsed.IsUnknown() {
		return e, false
	}
	var v []ChannelInfo
	d := o.ChannelUsed.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannelUsed sets the value of the ChannelUsed field in QueryInfo.
func (o *QueryInfo) SetChannelUsed(ctx context.Context, v ChannelInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel_used"]
	o.ChannelUsed = types.ListValueMust(t, vs)
}

// GetMetrics returns the value of the Metrics field in QueryInfo as
// a QueryMetrics value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo) GetMetrics(ctx context.Context) (QueryMetrics, bool) {
	var e QueryMetrics
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return e, false
	}
	var v []QueryMetrics
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetrics sets the value of the Metrics field in QueryInfo.
func (o *QueryInfo) SetMetrics(ctx context.Context, v QueryMetrics) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	o.Metrics = types.ListValueMust(t, vs)
}

type QueryList struct {
	// The total number of queries.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// The page number that is currently displayed.
	Page types.Int64 `tfsdk:"page" tf:"optional"`
	// The number of queries per page.
	PageSize types.Int64 `tfsdk:"page_size" tf:"optional"`
	// List of queries returned.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *QueryList) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryList) {
}

func (newState *QueryList) SyncEffectiveFieldsDuringRead(existingState QueryList) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(LegacyQuery{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryList
// only implements ToObjectValue() and Type().
func (o QueryList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"count":     o.Count,
			"page":      o.Page,
			"page_size": o.PageSize,
			"results":   o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":     types.Int64Type,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"results": basetypes.ListType{
				ElemType: LegacyQuery{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in QueryList as
// a slice of LegacyQuery values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryList) GetResults(ctx context.Context) ([]LegacyQuery, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []LegacyQuery
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in QueryList.
func (o *QueryList) SetResults(ctx context.Context, v []LegacyQuery) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	o.Results = types.ListValueMust(t, vs)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs types.Int64 `tfsdk:"compilation_time_ms" tf:"optional"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs types.Int64 `tfsdk:"execution_time_ms" tf:"optional"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes types.Int64 `tfsdk:"network_sent_bytes" tf:"optional"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp types.Int64 `tfsdk:"overloading_queue_start_timestamp" tf:"optional"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs types.Int64 `tfsdk:"photon_total_time_ms" tf:"optional"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp types.Int64 `tfsdk:"provisioning_queue_start_timestamp" tf:"optional"`
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes types.Int64 `tfsdk:"pruned_bytes" tf:"optional"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount types.Int64 `tfsdk:"pruned_files_count" tf:"optional"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp types.Int64 `tfsdk:"query_compilation_start_timestamp" tf:"optional"`
	// Total size of data read by the query, in bytes.
	ReadBytes types.Int64 `tfsdk:"read_bytes" tf:"optional"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes types.Int64 `tfsdk:"read_cache_bytes" tf:"optional"`
	// Number of files read after pruning
	ReadFilesCount types.Int64 `tfsdk:"read_files_count" tf:"optional"`
	// Number of partitions read after pruning.
	ReadPartitionsCount types.Int64 `tfsdk:"read_partitions_count" tf:"optional"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes types.Int64 `tfsdk:"read_remote_bytes" tf:"optional"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs types.Int64 `tfsdk:"result_fetch_time_ms" tf:"optional"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache types.Bool `tfsdk:"result_from_cache" tf:"optional"`
	// Total number of rows returned by the query.
	RowsProducedCount types.Int64 `tfsdk:"rows_produced_count" tf:"optional"`
	// Total number of rows read by the query.
	RowsReadCount types.Int64 `tfsdk:"rows_read_count" tf:"optional"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes types.Int64 `tfsdk:"spill_to_disk_bytes" tf:"optional"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs types.Int64 `tfsdk:"task_total_time_ms" tf:"optional"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs types.Int64 `tfsdk:"total_time_ms" tf:"optional"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes types.Int64 `tfsdk:"write_remote_bytes" tf:"optional"`
}

func (newState *QueryMetrics) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryMetrics) {
}

func (newState *QueryMetrics) SyncEffectiveFieldsDuringRead(existingState QueryMetrics) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryMetrics.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryMetrics) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryMetrics
// only implements ToObjectValue() and Type().
func (o QueryMetrics) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compilation_time_ms":                o.CompilationTimeMs,
			"execution_time_ms":                  o.ExecutionTimeMs,
			"network_sent_bytes":                 o.NetworkSentBytes,
			"overloading_queue_start_timestamp":  o.OverloadingQueueStartTimestamp,
			"photon_total_time_ms":               o.PhotonTotalTimeMs,
			"provisioning_queue_start_timestamp": o.ProvisioningQueueStartTimestamp,
			"pruned_bytes":                       o.PrunedBytes,
			"pruned_files_count":                 o.PrunedFilesCount,
			"query_compilation_start_timestamp":  o.QueryCompilationStartTimestamp,
			"read_bytes":                         o.ReadBytes,
			"read_cache_bytes":                   o.ReadCacheBytes,
			"read_files_count":                   o.ReadFilesCount,
			"read_partitions_count":              o.ReadPartitionsCount,
			"read_remote_bytes":                  o.ReadRemoteBytes,
			"result_fetch_time_ms":               o.ResultFetchTimeMs,
			"result_from_cache":                  o.ResultFromCache,
			"rows_produced_count":                o.RowsProducedCount,
			"rows_read_count":                    o.RowsReadCount,
			"spill_to_disk_bytes":                o.SpillToDiskBytes,
			"task_total_time_ms":                 o.TaskTotalTimeMs,
			"total_time_ms":                      o.TotalTimeMs,
			"write_remote_bytes":                 o.WriteRemoteBytes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryMetrics) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compilation_time_ms":                types.Int64Type,
			"execution_time_ms":                  types.Int64Type,
			"network_sent_bytes":                 types.Int64Type,
			"overloading_queue_start_timestamp":  types.Int64Type,
			"photon_total_time_ms":               types.Int64Type,
			"provisioning_queue_start_timestamp": types.Int64Type,
			"pruned_bytes":                       types.Int64Type,
			"pruned_files_count":                 types.Int64Type,
			"query_compilation_start_timestamp":  types.Int64Type,
			"read_bytes":                         types.Int64Type,
			"read_cache_bytes":                   types.Int64Type,
			"read_files_count":                   types.Int64Type,
			"read_partitions_count":              types.Int64Type,
			"read_remote_bytes":                  types.Int64Type,
			"result_fetch_time_ms":               types.Int64Type,
			"result_from_cache":                  types.BoolType,
			"rows_produced_count":                types.Int64Type,
			"rows_read_count":                    types.Int64Type,
			"spill_to_disk_bytes":                types.Int64Type,
			"task_total_time_ms":                 types.Int64Type,
			"total_time_ms":                      types.Int64Type,
			"write_remote_bytes":                 types.Int64Type,
		},
	}
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at" tf:"optional"`

	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// The name of the schema to execute this query in.
	Schema types.String `tfsdk:"schema" tf:"optional"`
}

func (newState *QueryOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryOptions) {
}

func (newState *QueryOptions) SyncEffectiveFieldsDuringRead(existingState QueryOptions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(Parameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryOptions
// only implements ToObjectValue() and Type().
func (o QueryOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog":           o.Catalog,
			"moved_to_trash_at": o.MovedToTrashAt,
			"parameters":        o.Parameters,
			"schema":            o.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog":           types.StringType,
			"moved_to_trash_at": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: Parameter{}.Type(ctx),
			},
			"schema": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in QueryOptions as
// a slice of Parameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryOptions) GetParameters(ctx context.Context) ([]Parameter, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []Parameter
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in QueryOptions.
func (o *QueryOptions) SetParameters(ctx context.Context, v []Parameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue types.List `tfsdk:"date_range_value" tf:"optional,object"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue types.List `tfsdk:"date_value" tf:"optional,object"`
	// Dropdown query parameter value.
	EnumValue types.List `tfsdk:"enum_value" tf:"optional,object"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Numeric query parameter value.
	NumericValue types.List `tfsdk:"numeric_value" tf:"optional,object"`
	// Query-based dropdown query parameter value.
	QueryBackedValue types.List `tfsdk:"query_backed_value" tf:"optional,object"`
	// Text query parameter value.
	TextValue types.List `tfsdk:"text_value" tf:"optional,object"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title types.String `tfsdk:"title" tf:"optional"`
}

func (newState *QueryParameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryParameter) {
}

func (newState *QueryParameter) SyncEffectiveFieldsDuringRead(existingState QueryParameter) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"date_range_value":   reflect.TypeOf(DateRangeValue{}),
		"date_value":         reflect.TypeOf(DateValue{}),
		"enum_value":         reflect.TypeOf(EnumValue{}),
		"numeric_value":      reflect.TypeOf(NumericValue{}),
		"query_backed_value": reflect.TypeOf(QueryBackedValue{}),
		"text_value":         reflect.TypeOf(TextValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryParameter
// only implements ToObjectValue() and Type().
func (o QueryParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_range_value":   o.DateRangeValue,
			"date_value":         o.DateValue,
			"enum_value":         o.EnumValue,
			"name":               o.Name,
			"numeric_value":      o.NumericValue,
			"query_backed_value": o.QueryBackedValue,
			"text_value":         o.TextValue,
			"title":              o.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryParameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_range_value": basetypes.ListType{
				ElemType: DateRangeValue{}.Type(ctx),
			},
			"date_value": basetypes.ListType{
				ElemType: DateValue{}.Type(ctx),
			},
			"enum_value": basetypes.ListType{
				ElemType: EnumValue{}.Type(ctx),
			},
			"name": types.StringType,
			"numeric_value": basetypes.ListType{
				ElemType: NumericValue{}.Type(ctx),
			},
			"query_backed_value": basetypes.ListType{
				ElemType: QueryBackedValue{}.Type(ctx),
			},
			"text_value": basetypes.ListType{
				ElemType: TextValue{}.Type(ctx),
			},
			"title": types.StringType,
		},
	}
}

// GetDateRangeValue returns the value of the DateRangeValue field in QueryParameter as
// a DateRangeValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetDateRangeValue(ctx context.Context) (DateRangeValue, bool) {
	var e DateRangeValue
	if o.DateRangeValue.IsNull() || o.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRangeValue
	d := o.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in QueryParameter.
func (o *QueryParameter) SetDateRangeValue(ctx context.Context, v DateRangeValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	o.DateRangeValue = types.ListValueMust(t, vs)
}

// GetDateValue returns the value of the DateValue field in QueryParameter as
// a DateValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetDateValue(ctx context.Context) (DateValue, bool) {
	var e DateValue
	if o.DateValue.IsNull() || o.DateValue.IsUnknown() {
		return e, false
	}
	var v []DateValue
	d := o.DateValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateValue sets the value of the DateValue field in QueryParameter.
func (o *QueryParameter) SetDateValue(ctx context.Context, v DateValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_value"]
	o.DateValue = types.ListValueMust(t, vs)
}

// GetEnumValue returns the value of the EnumValue field in QueryParameter as
// a EnumValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetEnumValue(ctx context.Context) (EnumValue, bool) {
	var e EnumValue
	if o.EnumValue.IsNull() || o.EnumValue.IsUnknown() {
		return e, false
	}
	var v []EnumValue
	d := o.EnumValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnumValue sets the value of the EnumValue field in QueryParameter.
func (o *QueryParameter) SetEnumValue(ctx context.Context, v EnumValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enum_value"]
	o.EnumValue = types.ListValueMust(t, vs)
}

// GetNumericValue returns the value of the NumericValue field in QueryParameter as
// a NumericValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetNumericValue(ctx context.Context) (NumericValue, bool) {
	var e NumericValue
	if o.NumericValue.IsNull() || o.NumericValue.IsUnknown() {
		return e, false
	}
	var v []NumericValue
	d := o.NumericValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNumericValue sets the value of the NumericValue field in QueryParameter.
func (o *QueryParameter) SetNumericValue(ctx context.Context, v NumericValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["numeric_value"]
	o.NumericValue = types.ListValueMust(t, vs)
}

// GetQueryBackedValue returns the value of the QueryBackedValue field in QueryParameter as
// a QueryBackedValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetQueryBackedValue(ctx context.Context) (QueryBackedValue, bool) {
	var e QueryBackedValue
	if o.QueryBackedValue.IsNull() || o.QueryBackedValue.IsUnknown() {
		return e, false
	}
	var v []QueryBackedValue
	d := o.QueryBackedValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryBackedValue sets the value of the QueryBackedValue field in QueryParameter.
func (o *QueryParameter) SetQueryBackedValue(ctx context.Context, v QueryBackedValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_backed_value"]
	o.QueryBackedValue = types.ListValueMust(t, vs)
}

// GetTextValue returns the value of the TextValue field in QueryParameter as
// a TextValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetTextValue(ctx context.Context) (TextValue, bool) {
	var e TextValue
	if o.TextValue.IsNull() || o.TextValue.IsUnknown() {
		return e, false
	}
	var v []TextValue
	d := o.TextValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTextValue sets the value of the TextValue field in QueryParameter.
func (o *QueryParameter) SetTextValue(ctx context.Context, v TextValue) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["text_value"]
	o.TextValue = types.ListValueMust(t, vs)
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options types.Object `tfsdk:"options" tf:"optional"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *QueryPostContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryPostContent) {
}

func (newState *QueryPostContent) SyncEffectiveFieldsDuringRead(existingState QueryPostContent) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryPostContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryPostContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryPostContent
// only implements ToObjectValue() and Type().
func (o QueryPostContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_source_id": o.DataSourceId,
			"description":    o.Description,
			"name":           o.Name,
			"options":        o.Options,
			"parent":         o.Parent,
			"query":          o.Query,
			"run_as_role":    o.RunAsRole,
			"tags":           o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryPostContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_source_id": types.StringType,
			"description":    types.StringType,
			"name":           types.StringType,
			"options":        types.ObjectType{},
			"parent":         types.StringType,
			"query":          types.StringType,
			"run_as_role":    types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetTags returns the value of the Tags field in QueryPostContent as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryPostContent) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in QueryPostContent.
func (o *QueryPostContent) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair" tf:"optional"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs" tf:"optional"`
}

func (newState *RepeatedEndpointConfPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepeatedEndpointConfPairs) {
}

func (newState *RepeatedEndpointConfPairs) SyncEffectiveFieldsDuringRead(existingState RepeatedEndpointConfPairs) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepeatedEndpointConfPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepeatedEndpointConfPairs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config_pair":         reflect.TypeOf(EndpointConfPair{}),
		"configuration_pairs": reflect.TypeOf(EndpointConfPair{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepeatedEndpointConfPairs
// only implements ToObjectValue() and Type().
func (o RepeatedEndpointConfPairs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_pair":         o.ConfigPair,
			"configuration_pairs": o.ConfigurationPairs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepeatedEndpointConfPairs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_pair": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
			"configuration_pairs": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
		},
	}
}

// GetConfigPair returns the value of the ConfigPair field in RepeatedEndpointConfPairs as
// a slice of EndpointConfPair values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepeatedEndpointConfPairs) GetConfigPair(ctx context.Context) ([]EndpointConfPair, bool) {
	if o.ConfigPair.IsNull() || o.ConfigPair.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair
	d := o.ConfigPair.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigPair sets the value of the ConfigPair field in RepeatedEndpointConfPairs.
func (o *RepeatedEndpointConfPairs) SetConfigPair(ctx context.Context, v []EndpointConfPair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_pair"]
	o.ConfigPair = types.ListValueMust(t, vs)
}

// GetConfigurationPairs returns the value of the ConfigurationPairs field in RepeatedEndpointConfPairs as
// a slice of EndpointConfPair values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepeatedEndpointConfPairs) GetConfigurationPairs(ctx context.Context) ([]EndpointConfPair, bool) {
	if o.ConfigurationPairs.IsNull() || o.ConfigurationPairs.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair
	d := o.ConfigurationPairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigurationPairs sets the value of the ConfigurationPairs field in RepeatedEndpointConfPairs.
func (o *RepeatedEndpointConfPairs) SetConfigurationPairs(ctx context.Context, v []EndpointConfPair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration_pairs"]
	o.ConfigurationPairs = types.ListValueMust(t, vs)
}

// Restore a dashboard
type RestoreDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (newState *RestoreDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreDashboardRequest) {
}

func (newState *RestoreDashboardRequest) SyncEffectiveFieldsDuringRead(existingState RestoreDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreDashboardRequest
// only implements ToObjectValue() and Type().
func (o RestoreDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Restore a query
type RestoreQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *RestoreQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreQueriesLegacyRequest) {
}

func (newState *RestoreQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState RestoreQueriesLegacyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreQueriesLegacyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreQueriesLegacyRequest
// only implements ToObjectValue() and Type().
func (o RestoreQueriesLegacyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreQueriesLegacyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type RestoreResponse struct {
}

func (newState *RestoreResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreResponse) {
}

func (newState *RestoreResponse) SyncEffectiveFieldsDuringRead(existingState RestoreResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreResponse
// only implements ToObjectValue() and Type().
func (o RestoreResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray types.List `tfsdk:"data_array" tf:"optional"`

	ExternalLinks types.List `tfsdk:"external_links" tf:"optional"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index" tf:"optional"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *ResultData) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultData) {
}

func (newState *ResultData) SyncEffectiveFieldsDuringRead(existingState ResultData) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array":     reflect.TypeOf(types.String{}),
		"external_links": reflect.TypeOf(ExternalLink{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData
// only implements ToObjectValue() and Type().
func (o ResultData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               o.ByteCount,
			"chunk_index":              o.ChunkIndex,
			"data_array":               o.DataArray,
			"external_links":           o.ExternalLinks,
			"next_chunk_index":         o.NextChunkIndex,
			"next_chunk_internal_link": o.NextChunkInternalLink,
			"row_count":                o.RowCount,
			"row_offset":               o.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":  types.Int64Type,
			"chunk_index": types.Int64Type,
			"data_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: types.StringType,
				},
			},
			"external_links": basetypes.ListType{
				ElemType: ExternalLink{}.Type(ctx),
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultData) GetDataArray(ctx context.Context) ([]types.String, bool) {
	if o.DataArray.IsNull() || o.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData.
func (o *ResultData) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	o.DataArray = types.ListValueMust(t, vs)
}

// GetExternalLinks returns the value of the ExternalLinks field in ResultData as
// a slice of ExternalLink values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultData) GetExternalLinks(ctx context.Context) ([]ExternalLink, bool) {
	if o.ExternalLinks.IsNull() || o.ExternalLinks.IsUnknown() {
		return nil, false
	}
	var v []ExternalLink
	d := o.ExternalLinks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLinks sets the value of the ExternalLinks field in ResultData.
func (o *ResultData) SetExternalLinks(ctx context.Context, v []ExternalLink) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_links"]
	o.ExternalLinks = types.ListValueMust(t, vs)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks" tf:"optional"`

	Format types.String `tfsdk:"format" tf:"optional"`
	// The schema is an ordered list of column descriptions.
	Schema types.List `tfsdk:"schema" tf:"optional,object"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount types.Int64 `tfsdk:"total_byte_count" tf:"optional"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount types.Int64 `tfsdk:"total_chunk_count" tf:"optional"`
	// The total number of rows in the result set.
	TotalRowCount types.Int64 `tfsdk:"total_row_count" tf:"optional"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated types.Bool `tfsdk:"truncated" tf:"optional"`
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultManifest) {
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringRead(existingState ResultManifest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultManifest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"chunks": reflect.TypeOf(BaseChunkInfo{}),
		"schema": reflect.TypeOf(ResultSchema{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest
// only implements ToObjectValue() and Type().
func (o ResultManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunks":            o.Chunks,
			"format":            o.Format,
			"schema":            o.Schema,
			"total_byte_count":  o.TotalByteCount,
			"total_chunk_count": o.TotalChunkCount,
			"total_row_count":   o.TotalRowCount,
			"truncated":         o.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultManifest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunks": basetypes.ListType{
				ElemType: BaseChunkInfo{}.Type(ctx),
			},
			"format": types.StringType,
			"schema": basetypes.ListType{
				ElemType: ResultSchema{}.Type(ctx),
			},
			"total_byte_count":  types.Int64Type,
			"total_chunk_count": types.Int64Type,
			"total_row_count":   types.Int64Type,
			"truncated":         types.BoolType,
		},
	}
}

// GetChunks returns the value of the Chunks field in ResultManifest as
// a slice of BaseChunkInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultManifest) GetChunks(ctx context.Context) ([]BaseChunkInfo, bool) {
	if o.Chunks.IsNull() || o.Chunks.IsUnknown() {
		return nil, false
	}
	var v []BaseChunkInfo
	d := o.Chunks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChunks sets the value of the Chunks field in ResultManifest.
func (o *ResultManifest) SetChunks(ctx context.Context, v []BaseChunkInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["chunks"]
	o.Chunks = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in ResultManifest as
// a ResultSchema value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultManifest) GetSchema(ctx context.Context) (ResultSchema, bool) {
	var e ResultSchema
	if o.Schema.IsNull() || o.Schema.IsUnknown() {
		return e, false
	}
	var v []ResultSchema
	d := o.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in ResultManifest.
func (o *ResultManifest) SetSchema(ctx context.Context, v ResultSchema) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	o.Schema = types.ListValueMust(t, vs)
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	ColumnCount types.Int64 `tfsdk:"column_count" tf:"optional"`

	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *ResultSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultSchema) {
}

func (newState *ResultSchema) SyncEffectiveFieldsDuringRead(existingState ResultSchema) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultSchema) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultSchema
// only implements ToObjectValue() and Type().
func (o ResultSchema) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": o.ColumnCount,
			"columns":      o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultSchema) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in ResultSchema as
// a slice of ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultSchema) GetColumns(ctx context.Context) ([]ColumnInfo, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultSchema.
func (o *ResultSchema) SetColumns(ctx context.Context, v []ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	o.Columns = types.ListValueMust(t, vs)
}

type ServiceError struct {
	ErrorCode types.String `tfsdk:"error_code" tf:"optional"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *ServiceError) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServiceError) {
}

func (newState *ServiceError) SyncEffectiveFieldsDuringRead(existingState ServiceError) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServiceError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServiceError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceError
// only implements ToObjectValue() and Type().
func (o ServiceError) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_code": o.ErrorCode,
			"message":    o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServiceError) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_code": types.StringType,
			"message":    types.StringType,
		},
	}
}

// Set object ACL
type SetRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permission to set.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *SetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRequest) {
}

func (newState *SetRequest) SyncEffectiveFieldsDuringRead(existingState SetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetRequest
// only implements ToObjectValue() and Type().
func (o SetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"objectId":            o.ObjectId,
			"objectType":          o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl{}.Type(ctx),
			},
			"objectId":   types.StringType,
			"objectType": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetRequest as
// a slice of AccessControl values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetRequest) GetAccessControlList(ctx context.Context) ([]AccessControl, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetRequest.
func (o *SetRequest) SetAccessControlList(ctx context.Context, v []AccessControl) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id" tf:"optional"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *SetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetResponse) {
}

func (newState *SetResponse) SyncEffectiveFieldsDuringRead(existingState SetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetResponse
// only implements ToObjectValue() and Type().
func (o SetResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetResponse as
// a slice of AccessControl values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetResponse) GetAccessControlList(ctx context.Context) ([]AccessControl, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetResponse.
func (o *SetResponse) SetAccessControlList(ctx context.Context, v []AccessControl) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel types.List `tfsdk:"channel" tf:"optional,object"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.List `tfsdk:"config_param" tf:"optional,object"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config" tf:"optional"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types" tf:"optional"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.List `tfsdk:"global_param" tf:"optional,object"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy" tf:"optional"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters" tf:"optional,object"`
}

func (newState *SetWorkspaceWarehouseConfigRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetWorkspaceWarehouseConfigRequest) {
}

func (newState *SetWorkspaceWarehouseConfigRequest) SyncEffectiveFieldsDuringRead(existingState SetWorkspaceWarehouseConfigRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetWorkspaceWarehouseConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":                      reflect.TypeOf(Channel{}),
		"config_param":                 reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"data_access_config":           reflect.TypeOf(EndpointConfPair{}),
		"enabled_warehouse_types":      reflect.TypeOf(WarehouseTypePair{}),
		"global_param":                 reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"sql_configuration_parameters": reflect.TypeOf(RepeatedEndpointConfPairs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetWorkspaceWarehouseConfigRequest
// only implements ToObjectValue() and Type().
func (o SetWorkspaceWarehouseConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel":                      o.Channel,
			"config_param":                 o.ConfigParam,
			"data_access_config":           o.DataAccessConfig,
			"enabled_warehouse_types":      o.EnabledWarehouseTypes,
			"global_param":                 o.GlobalParam,
			"google_service_account":       o.GoogleServiceAccount,
			"instance_profile_arn":         o.InstanceProfileArn,
			"security_policy":              o.SecurityPolicy,
			"sql_configuration_parameters": o.SqlConfigurationParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetWorkspaceWarehouseConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel": basetypes.ListType{
				ElemType: Channel{}.Type(ctx),
			},
			"config_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair{}.Type(ctx),
			},
			"global_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
			"google_service_account": types.StringType,
			"instance_profile_arn":   types.StringType,
			"security_policy":        types.StringType,
			"sql_configuration_parameters": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs{}.Type(ctx),
			},
		},
	}
}

// GetChannel returns the value of the Channel field in SetWorkspaceWarehouseConfigRequest as
// a Channel value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetChannel(ctx context.Context) (Channel, bool) {
	var e Channel
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetChannel(ctx context.Context, v Channel) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	o.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest as
// a slice of EndpointConfPair values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair, bool) {
	if o.DataAccessConfig.IsNull() || o.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair
	d := o.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	o.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest as
// a slice of WarehouseTypePair values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair, bool) {
	if o.EnabledWarehouseTypes.IsNull() || o.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair
	d := o.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	o.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.GlobalParam.IsNull() || o.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	o.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs
	d := o.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	o.SqlConfigurationParameters = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigResponse struct {
}

func (newState *SetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetWorkspaceWarehouseConfigResponse) {
}

func (newState *SetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringRead(existingState SetWorkspaceWarehouseConfigResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetWorkspaceWarehouseConfigResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetWorkspaceWarehouseConfigResponse
// only implements ToObjectValue() and Type().
func (o SetWorkspaceWarehouseConfigResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetWorkspaceWarehouseConfigResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Start a warehouse
type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *StartRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartRequest) {
}

func (newState *StartRequest) SyncEffectiveFieldsDuringRead(existingState StartRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartRequest
// only implements ToObjectValue() and Type().
func (o StartRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StartWarehouseResponse struct {
}

func (newState *StartWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartWarehouseResponse) {
}

func (newState *StartWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState StartWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartWarehouseResponse
// only implements ToObjectValue() and Type().
func (o StartWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StartWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	Name types.String `tfsdk:"name" tf:""`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type_ types.String `tfsdk:"type" tf:"optional"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *StatementParameterListItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementParameterListItem) {
}

func (newState *StatementParameterListItem) SyncEffectiveFieldsDuringRead(existingState StatementParameterListItem) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementParameterListItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementParameterListItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementParameterListItem
// only implements ToObjectValue() and Type().
func (o StatementParameterListItem) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  o.Name,
			"type":  o.Type_,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StatementParameterListItem) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"type":  types.StringType,
			"value": types.StringType,
		},
	}
}

type StatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest types.List `tfsdk:"manifest" tf:"optional,object"`

	Result types.List `tfsdk:"result" tf:"optional,object"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
	// The status response includes execution state and if relevant, error
	// information.
	Status types.List `tfsdk:"status" tf:"optional,object"`
}

func (newState *StatementResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementResponse) {
}

func (newState *StatementResponse) SyncEffectiveFieldsDuringRead(existingState StatementResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest{}),
		"result":   reflect.TypeOf(ResultData{}),
		"status":   reflect.TypeOf(StatementStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementResponse
// only implements ToObjectValue() and Type().
func (o StatementResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":     o.Manifest,
			"result":       o.Result,
			"statement_id": o.StatementId,
			"status":       o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StatementResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest": basetypes.ListType{
				ElemType: ResultManifest{}.Type(ctx),
			},
			"result": basetypes.ListType{
				ElemType: ResultData{}.Type(ctx),
			},
			"statement_id": types.StringType,
			"status": basetypes.ListType{
				ElemType: StatementStatus{}.Type(ctx),
			},
		},
	}
}

// GetManifest returns the value of the Manifest field in StatementResponse as
// a ResultManifest value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse) GetManifest(ctx context.Context) (ResultManifest, bool) {
	var e ResultManifest
	if o.Manifest.IsNull() || o.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest
	d := o.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in StatementResponse.
func (o *StatementResponse) SetManifest(ctx context.Context, v ResultManifest) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	o.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in StatementResponse as
// a ResultData value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse) GetResult(ctx context.Context) (ResultData, bool) {
	var e ResultData
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData
	d := o.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in StatementResponse.
func (o *StatementResponse) SetResult(ctx context.Context, v ResultData) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	o.Result = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in StatementResponse as
// a StatementStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse) GetStatus(ctx context.Context) (StatementStatus, bool) {
	var e StatementStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []StatementStatus
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in StatementResponse.
func (o *StatementResponse) SetStatus(ctx context.Context, v StatementStatus) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	Error types.List `tfsdk:"error" tf:"optional,object"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *StatementStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementStatus) {
}

func (newState *StatementStatus) SyncEffectiveFieldsDuringRead(existingState StatementStatus) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(ServiceError{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementStatus
// only implements ToObjectValue() and Type().
func (o StatementStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error": o.Error,
			"state": o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StatementStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": basetypes.ListType{
				ElemType: ServiceError{}.Type(ctx),
			},
			"state": types.StringType,
		},
	}
}

// GetError returns the value of the Error field in StatementStatus as
// a ServiceError value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementStatus) GetError(ctx context.Context) (ServiceError, bool) {
	var e ServiceError
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []ServiceError
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in StatementStatus.
func (o *StatementStatus) SetError(ctx context.Context, v ServiceError) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

// Stop a warehouse
type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *StopRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopRequest) {
}

func (newState *StopRequest) SyncEffectiveFieldsDuringRead(existingState StopRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest
// only implements ToObjectValue() and Type().
func (o StopRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StopWarehouseResponse struct {
}

func (newState *StopWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopWarehouseResponse) {
}

func (newState *StopWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState StopWarehouseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopWarehouseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopWarehouseResponse
// only implements ToObjectValue() and Type().
func (o StopWarehouseResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StopWarehouseResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Success struct {
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *Success) SyncEffectiveFieldsDuringCreateOrUpdate(plan Success) {
}

func (newState *Success) SyncEffectiveFieldsDuringRead(existingState Success) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Success.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Success) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Success
// only implements ToObjectValue() and Type().
func (o Success) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Success) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
		},
	}
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code" tf:"optional"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
	// type of the termination
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationReason) {
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringRead(existingState TerminationReason) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TerminationReason.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TerminationReason) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason
// only implements ToObjectValue() and Type().
func (o TerminationReason) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       o.Code,
			"parameters": o.Parameters,
			"type":       o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TerminationReason) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code": types.StringType,
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"type": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in TerminationReason as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TerminationReason) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in TerminationReason.
func (o *TerminationReason) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.MapValueMust(t, vs)
}

type TextValue struct {
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *TextValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextValue) {
}

func (newState *TextValue) SyncEffectiveFieldsDuringRead(existingState TextValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TextValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TextValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TextValue
// only implements ToObjectValue() and Type().
func (o TextValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TextValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TimeRange struct {
	// The end time in milliseconds.
	EndTimeMs types.Int64 `tfsdk:"end_time_ms" tf:"optional"`
	// The start time in milliseconds.
	StartTimeMs types.Int64 `tfsdk:"start_time_ms" tf:"optional"`
}

func (newState *TimeRange) SyncEffectiveFieldsDuringCreateOrUpdate(plan TimeRange) {
}

func (newState *TimeRange) SyncEffectiveFieldsDuringRead(existingState TimeRange) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TimeRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TimeRange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeRange
// only implements ToObjectValue() and Type().
func (o TimeRange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   o.EndTimeMs,
			"start_time_ms": o.StartTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TimeRange) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time_ms":   types.Int64Type,
			"start_time_ms": types.Int64Type,
		},
	}
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner" tf:"optional"`
}

func (newState *TransferOwnershipObjectId) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransferOwnershipObjectId) {
}

func (newState *TransferOwnershipObjectId) SyncEffectiveFieldsDuringRead(existingState TransferOwnershipObjectId) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransferOwnershipObjectId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TransferOwnershipObjectId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipObjectId
// only implements ToObjectValue() and Type().
func (o TransferOwnershipObjectId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner": o.NewOwner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransferOwnershipObjectId) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner": types.StringType,
		},
	}
}

// Transfer object ownership
type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner" tf:"optional"`
	// The ID of the object on which to change ownership.
	ObjectId types.List `tfsdk:"-"`
	// The type of object on which to change ownership.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *TransferOwnershipRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransferOwnershipRequest) {
}

func (newState *TransferOwnershipRequest) SyncEffectiveFieldsDuringRead(existingState TransferOwnershipRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransferOwnershipRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TransferOwnershipRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objectId": reflect.TypeOf(TransferOwnershipObjectId{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipRequest
// only implements ToObjectValue() and Type().
func (o TransferOwnershipRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner":  o.NewOwner,
			"objectId":   o.ObjectId,
			"objectType": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransferOwnershipRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner": types.StringType,
			"objectId": basetypes.ListType{
				ElemType: TransferOwnershipObjectId{}.Type(ctx),
			},
			"objectType": types.StringType,
		},
	}
}

// GetObjectId returns the value of the ObjectId field in TransferOwnershipRequest as
// a TransferOwnershipObjectId value.
// If the field is unknown or null, the boolean return value is false.
func (o *TransferOwnershipRequest) GetObjectId(ctx context.Context) (TransferOwnershipObjectId, bool) {
	var e TransferOwnershipObjectId
	if o.ObjectId.IsNull() || o.ObjectId.IsUnknown() {
		return e, false
	}
	var v []TransferOwnershipObjectId
	d := o.ObjectId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetObjectId sets the value of the ObjectId field in TransferOwnershipRequest.
func (o *TransferOwnershipRequest) SetObjectId(ctx context.Context, v TransferOwnershipObjectId) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objectId"]
	o.ObjectId = types.ListValueMust(t, vs)
}

// Delete an alert
type TrashAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *TrashAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashAlertRequest) {
}

func (newState *TrashAlertRequest) SyncEffectiveFieldsDuringRead(existingState TrashAlertRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashAlertRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertRequest
// only implements ToObjectValue() and Type().
func (o TrashAlertRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a query
type TrashQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *TrashQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashQueryRequest) {
}

func (newState *TrashQueryRequest) SyncEffectiveFieldsDuringRead(existingState TrashQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashQueryRequest
// only implements ToObjectValue() and Type().
func (o TrashQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type UpdateAlertRequest struct {
	Alert types.List `tfsdk:"alert" tf:"optional,object"`

	Id types.String `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`
}

func (newState *UpdateAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAlertRequest) {
}

func (newState *UpdateAlertRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAlertRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(UpdateAlertRequestAlert{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequest
// only implements ToObjectValue() and Type().
func (o UpdateAlertRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":       o.Alert,
			"id":          o.Id,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: UpdateAlertRequestAlert{}.Type(ctx),
			},
			"id":          types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetAlert returns the value of the Alert field in UpdateAlertRequest as
// a UpdateAlertRequestAlert value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertRequest) GetAlert(ctx context.Context) (UpdateAlertRequestAlert, bool) {
	var e UpdateAlertRequestAlert
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []UpdateAlertRequestAlert
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in UpdateAlertRequest.
func (o *UpdateAlertRequest) SetAlert(ctx context.Context, v UpdateAlertRequestAlert) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition" tf:"optional,object"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
}

func (newState *UpdateAlertRequestAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAlertRequestAlert) {
}

func (newState *UpdateAlertRequestAlert) SyncEffectiveFieldsDuringRead(existingState UpdateAlertRequestAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertRequestAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertRequestAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequestAlert
// only implements ToObjectValue() and Type().
func (o UpdateAlertRequestAlert) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            o.Condition,
			"custom_body":          o.CustomBody,
			"custom_subject":       o.CustomSubject,
			"display_name":         o.DisplayName,
			"notify_on_ok":         o.NotifyOnOk,
			"owner_user_name":      o.OwnerUserName,
			"query_id":             o.QueryId,
			"seconds_to_retrigger": o.SecondsToRetrigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertRequestAlert) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition{}.Type(ctx),
			},
			"custom_body":          types.StringType,
			"custom_subject":       types.StringType,
			"display_name":         types.StringType,
			"notify_on_ok":         types.BoolType,
			"owner_user_name":      types.StringType,
			"query_id":             types.StringType,
			"seconds_to_retrigger": types.Int64Type,
		},
	}
}

// GetCondition returns the value of the Condition field in UpdateAlertRequestAlert as
// a AlertCondition value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertRequestAlert) GetCondition(ctx context.Context) (AlertCondition, bool) {
	var e AlertCondition
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in UpdateAlertRequestAlert.
func (o *UpdateAlertRequestAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type UpdateQueryRequest struct {
	Id types.String `tfsdk:"-"`

	Query types.List `tfsdk:"query" tf:"optional,object"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`
}

func (newState *UpdateQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateQueryRequest) {
}

func (newState *UpdateQueryRequest) SyncEffectiveFieldsDuringRead(existingState UpdateQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(UpdateQueryRequestQuery{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequest
// only implements ToObjectValue() and Type().
func (o UpdateQueryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":          o.Id,
			"query":       o.Query,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"query": basetypes.ListType{
				ElemType: UpdateQueryRequestQuery{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetQuery returns the value of the Query field in UpdateQueryRequest as
// a UpdateQueryRequestQuery value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequest) GetQuery(ctx context.Context) (UpdateQueryRequestQuery, bool) {
	var e UpdateQueryRequestQuery
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []UpdateQueryRequestQuery
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in UpdateQueryRequest.
func (o *UpdateQueryRequest) SetQuery(ctx context.Context, v UpdateQueryRequestQuery) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *UpdateQueryRequestQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateQueryRequestQuery) {
}

func (newState *UpdateQueryRequestQuery) SyncEffectiveFieldsDuringRead(existingState UpdateQueryRequestQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQueryRequestQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQueryRequestQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequestQuery
// only implements ToObjectValue() and Type().
func (o UpdateQueryRequestQuery) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit": o.ApplyAutoLimit,
			"catalog":          o.Catalog,
			"description":      o.Description,
			"display_name":     o.DisplayName,
			"owner_user_name":  o.OwnerUserName,
			"parameters":       o.Parameters,
			"query_text":       o.QueryText,
			"run_as_mode":      o.RunAsMode,
			"schema":           o.Schema,
			"tags":             o.Tags,
			"warehouse_id":     o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQueryRequestQuery) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit": types.BoolType,
			"catalog":          types.StringType,
			"description":      types.StringType,
			"display_name":     types.StringType,
			"owner_user_name":  types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter{}.Type(ctx),
			},
			"query_text":  types.StringType,
			"run_as_mode": types.StringType,
			"schema":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in UpdateQueryRequestQuery as
// a slice of QueryParameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequestQuery) GetParameters(ctx context.Context) ([]QueryParameter, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in UpdateQueryRequestQuery.
func (o *UpdateQueryRequestQuery) SetParameters(ctx context.Context, v []QueryParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in UpdateQueryRequestQuery as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequestQuery) GetTags(ctx context.Context) ([]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in UpdateQueryRequestQuery.
func (o *UpdateQueryRequestQuery) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse
// only implements ToObjectValue() and Type().
func (o UpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateVisualizationRequest struct {
	Id types.String `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`

	Visualization types.List `tfsdk:"visualization" tf:"optional,object"`
}

func (newState *UpdateVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVisualizationRequest) {
}

func (newState *UpdateVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateVisualizationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVisualizationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(UpdateVisualizationRequestVisualization{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequest
// only implements ToObjectValue() and Type().
func (o UpdateVisualizationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":            o.Id,
			"update_mask":   o.UpdateMask,
			"visualization": o.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateVisualizationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":          types.StringType,
			"update_mask": types.StringType,
			"visualization": basetypes.ListType{
				ElemType: UpdateVisualizationRequestVisualization{}.Type(ctx),
			},
		},
	}
}

// GetVisualization returns the value of the Visualization field in UpdateVisualizationRequest as
// a UpdateVisualizationRequestVisualization value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateVisualizationRequest) GetVisualization(ctx context.Context) (UpdateVisualizationRequestVisualization, bool) {
	var e UpdateVisualizationRequestVisualization
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []UpdateVisualizationRequestVisualization
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in UpdateVisualizationRequest.
func (o *UpdateVisualizationRequest) SetVisualization(ctx context.Context, v UpdateVisualizationRequestVisualization) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type UpdateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *UpdateVisualizationRequestVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVisualizationRequestVisualization) {
}

func (newState *UpdateVisualizationRequestVisualization) SyncEffectiveFieldsDuringRead(existingState UpdateVisualizationRequestVisualization) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVisualizationRequestVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVisualizationRequestVisualization) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequestVisualization
// only implements ToObjectValue() and Type().
func (o UpdateVisualizationRequestVisualization) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          o.DisplayName,
			"serialized_options":    o.SerializedOptions,
			"serialized_query_plan": o.SerializedQueryPlan,
			"type":                  o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateVisualizationRequestVisualization) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":          types.StringType,
			"serialized_options":    types.StringType,
			"serialized_query_plan": types.StringType,
			"type":                  types.StringType,
		},
	}
}

type User struct {
	Email types.String `tfsdk:"email" tf:"optional"`

	Id types.Int64 `tfsdk:"id" tf:"optional"`

	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *User) SyncEffectiveFieldsDuringCreateOrUpdate(plan User) {
}

func (newState *User) SyncEffectiveFieldsDuringRead(existingState User) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in User.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User
// only implements ToObjectValue() and Type().
func (o User) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email": o.Email,
			"id":    o.Id,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o User) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.Int64Type,
			"name":  types.StringType,
		},
	}
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the visualization.
	Id types.String `tfsdk:"id" tf:"optional"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type" tf:"optional"`
	// The timestamp indicating when the visualization was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *Visualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan Visualization) {
}

func (newState *Visualization) SyncEffectiveFieldsDuringRead(existingState Visualization) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Visualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Visualization) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Visualization
// only implements ToObjectValue() and Type().
func (o Visualization) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           o.CreateTime,
			"display_name":          o.DisplayName,
			"id":                    o.Id,
			"query_id":              o.QueryId,
			"serialized_options":    o.SerializedOptions,
			"serialized_query_plan": o.SerializedQueryPlan,
			"type":                  o.Type_,
			"update_time":           o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Visualization) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":           types.StringType,
			"display_name":          types.StringType,
			"id":                    types.StringType,
			"query_id":              types.StringType,
			"serialized_options":    types.StringType,
			"serialized_query_plan": types.StringType,
			"type":                  types.StringType,
			"update_time":           types.StringType,
		},
	}
}

type WarehouseAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *WarehouseAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlRequest) {
}

func (newState *WarehouseAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlRequest
// only implements ToObjectValue() and Type().
func (o WarehouseAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehouseAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *WarehouseAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlResponse) {
}

func (newState *WarehouseAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WarehousePermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlResponse
// only implements ToObjectValue() and Type().
func (o WarehouseAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehouseAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: WarehousePermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in WarehouseAccessControlResponse as
// a slice of WarehousePermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehouseAccessControlResponse) GetAllPermissions(ctx context.Context) ([]WarehousePermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WarehouseAccessControlResponse.
func (o *WarehouseAccessControlResponse) SetAllPermissions(ctx context.Context, v []WarehousePermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	o.AllPermissions = types.ListValueMust(t, vs)
}

type WarehousePermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WarehousePermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermission) {
}

func (newState *WarehousePermission) SyncEffectiveFieldsDuringRead(existingState WarehousePermission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermission
// only implements ToObjectValue() and Type().
func (o WarehousePermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in WarehousePermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in WarehousePermission.
func (o *WarehousePermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type WarehousePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *WarehousePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissions) {
}

func (newState *WarehousePermissions) SyncEffectiveFieldsDuringRead(existingState WarehousePermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissions
// only implements ToObjectValue() and Type().
func (o WarehousePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WarehouseAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WarehousePermissions as
// a slice of WarehouseAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermissions) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissions.
func (o *WarehousePermissions) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehousePermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WarehousePermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissionsDescription) {
}

func (newState *WarehousePermissionsDescription) SyncEffectiveFieldsDuringRead(existingState WarehousePermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsDescription
// only implements ToObjectValue() and Type().
func (o WarehousePermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type WarehousePermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *WarehousePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissionsRequest) {
}

func (newState *WarehousePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState WarehousePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsRequest
// only implements ToObjectValue() and Type().
func (o WarehousePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"warehouse_id":        o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WarehouseAccessControlRequest{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WarehousePermissionsRequest as
// a slice of WarehouseAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermissionsRequest) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissionsRequest.
func (o *WarehousePermissionsRequest) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *WarehouseTypePair) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseTypePair) {
}

func (newState *WarehouseTypePair) SyncEffectiveFieldsDuringRead(existingState WarehouseTypePair) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseTypePair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseTypePair) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseTypePair
// only implements ToObjectValue() and Type().
func (o WarehouseTypePair) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":        o.Enabled,
			"warehouse_type": o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehouseTypePair) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":        types.BoolType,
			"warehouse_type": types.StringType,
		},
	}
}

type Widget struct {
	// The unique ID for this widget.
	Id types.String `tfsdk:"id" tf:"optional"`

	Options types.List `tfsdk:"options" tf:"optional,object"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization types.List `tfsdk:"visualization" tf:"optional,object"`
	// Unused field.
	Width types.Int64 `tfsdk:"width" tf:"optional"`
}

func (newState *Widget) SyncEffectiveFieldsDuringCreateOrUpdate(plan Widget) {
}

func (newState *Widget) SyncEffectiveFieldsDuringRead(existingState Widget) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Widget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Widget) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":       reflect.TypeOf(WidgetOptions{}),
		"visualization": reflect.TypeOf(LegacyVisualization{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Widget
// only implements ToObjectValue() and Type().
func (o Widget) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":            o.Id,
			"options":       o.Options,
			"visualization": o.Visualization,
			"width":         o.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Widget) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"options": basetypes.ListType{
				ElemType: WidgetOptions{}.Type(ctx),
			},
			"visualization": basetypes.ListType{
				ElemType: LegacyVisualization{}.Type(ctx),
			},
			"width": types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in Widget as
// a WidgetOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *Widget) GetOptions(ctx context.Context) (WidgetOptions, bool) {
	var e WidgetOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Widget.
func (o *Widget) SetOptions(ctx context.Context, v WidgetOptions) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetVisualization returns the value of the Visualization field in Widget as
// a LegacyVisualization value.
// If the field is unknown or null, the boolean return value is false.
func (o *Widget) GetVisualization(ctx context.Context) (LegacyVisualization, bool) {
	var e LegacyVisualization
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []LegacyVisualization
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in Widget.
func (o *Widget) SetVisualization(ctx context.Context, v LegacyVisualization) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Custom description of the widget
	Description types.String `tfsdk:"description" tf:"optional"`
	// Whether this widget is hidden on the dashboard.
	IsHidden types.Bool `tfsdk:"isHidden" tf:"optional"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings types.Object `tfsdk:"parameterMappings" tf:"optional"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position types.List `tfsdk:"position" tf:"optional,object"`
	// Custom title of the widget
	Title types.String `tfsdk:"title" tf:"optional"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
}

func (newState *WidgetOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetOptions) {
}

func (newState *WidgetOptions) SyncEffectiveFieldsDuringRead(existingState WidgetOptions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WidgetOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WidgetOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"position": reflect.TypeOf(WidgetPosition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetOptions
// only implements ToObjectValue() and Type().
func (o WidgetOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":        o.CreatedAt,
			"description":       o.Description,
			"isHidden":          o.IsHidden,
			"parameterMappings": o.ParameterMappings,
			"position":          o.Position,
			"title":             o.Title,
			"updated_at":        o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WidgetOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":        types.StringType,
			"description":       types.StringType,
			"isHidden":          types.BoolType,
			"parameterMappings": types.ObjectType{},
			"position": basetypes.ListType{
				ElemType: WidgetPosition{}.Type(ctx),
			},
			"title":      types.StringType,
			"updated_at": types.StringType,
		},
	}
}

// GetPosition returns the value of the Position field in WidgetOptions as
// a WidgetPosition value.
// If the field is unknown or null, the boolean return value is false.
func (o *WidgetOptions) GetPosition(ctx context.Context) (WidgetPosition, bool) {
	var e WidgetPosition
	if o.Position.IsNull() || o.Position.IsUnknown() {
		return e, false
	}
	var v []WidgetPosition
	d := o.Position.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPosition sets the value of the Position field in WidgetOptions.
func (o *WidgetOptions) SetPosition(ctx context.Context, v WidgetPosition) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["position"]
	o.Position = types.ListValueMust(t, vs)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	AutoHeight types.Bool `tfsdk:"autoHeight" tf:"optional"`
	// column in the dashboard grid. Values start with 0
	Col types.Int64 `tfsdk:"col" tf:"optional"`
	// row in the dashboard grid. Values start with 0
	Row types.Int64 `tfsdk:"row" tf:"optional"`
	// width of the widget measured in dashboard grid cells
	SizeX types.Int64 `tfsdk:"sizeX" tf:"optional"`
	// height of the widget measured in dashboard grid cells
	SizeY types.Int64 `tfsdk:"sizeY" tf:"optional"`
}

func (newState *WidgetPosition) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetPosition) {
}

func (newState *WidgetPosition) SyncEffectiveFieldsDuringRead(existingState WidgetPosition) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WidgetPosition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WidgetPosition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetPosition
// only implements ToObjectValue() and Type().
func (o WidgetPosition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoHeight": o.AutoHeight,
			"col":        o.Col,
			"row":        o.Row,
			"sizeX":      o.SizeX,
			"sizeY":      o.SizeY,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WidgetPosition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoHeight": types.BoolType,
			"col":        types.Int64Type,
			"row":        types.Int64Type,
			"sizeX":      types.Int64Type,
			"sizeY":      types.Int64Type,
		},
	}
}
