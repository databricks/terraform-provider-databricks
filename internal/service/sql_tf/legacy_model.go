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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccessControl_SdkV2 struct {
	GroupName types.String `tfsdk:"group_name"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel types.String `tfsdk:"permission_level"`

	UserName types.String `tfsdk:"user_name"`
}

func (newState *AccessControl_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControl_SdkV2) {
}

func (newState *AccessControl_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccessControl_SdkV2) {
}

func (c AccessControl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControl_SdkV2
// only implements ToObjectValue() and Type().
func (o AccessControl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":       o.GroupName,
			"permission_level": o.PermissionLevel,
			"user_name":        o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccessControl_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":       types.StringType,
			"permission_level": types.StringType,
			"user_name":        types.StringType,
		},
	}
}

type Alert_SdkV2 struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *Alert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Alert_SdkV2) {
}

func (newState *Alert_SdkV2) SyncEffectiveFieldsDuringRead(existingState Alert_SdkV2) {
}

func (c Alert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
	attrs["condition"] = attrs["condition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["custom_body"] = attrs["custom_body"].SetOptional()
	attrs["custom_subject"] = attrs["custom_subject"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetOptional()
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["seconds_to_retrigger"] = attrs["seconds_to_retrigger"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["trigger_time"] = attrs["trigger_time"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Alert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Alert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Alert_SdkV2
// only implements ToObjectValue() and Type().
func (o Alert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Alert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition_SdkV2{}.Type(ctx),
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

// GetCondition returns the value of the Condition field in Alert_SdkV2 as
// a AlertCondition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Alert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in Alert_SdkV2.
func (o *Alert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type AlertCondition_SdkV2 struct {
	// Alert state if result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state"`
	// Operator used for comparison in alert evaluation.
	Op types.String `tfsdk:"op"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand types.List `tfsdk:"operand"`
	// Threshold value used for comparison in alert evaluation.
	Threshold types.List `tfsdk:"threshold"`
}

func (newState *AlertCondition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertCondition_SdkV2) {
}

func (newState *AlertCondition_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertCondition_SdkV2) {
}

func (c AlertCondition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["op"] = attrs["op"].SetOptional()
	attrs["operand"] = attrs["operand"].SetOptional()
	attrs["operand"] = attrs["operand"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["threshold"] = attrs["threshold"].SetOptional()
	attrs["threshold"] = attrs["threshold"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertCondition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertCondition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"operand":   reflect.TypeOf(AlertConditionOperand_SdkV2{}),
		"threshold": reflect.TypeOf(AlertConditionThreshold_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertCondition_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertCondition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AlertCondition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"empty_result_state": types.StringType,
			"op":                 types.StringType,
			"operand": basetypes.ListType{
				ElemType: AlertConditionOperand_SdkV2{}.Type(ctx),
			},
			"threshold": basetypes.ListType{
				ElemType: AlertConditionThreshold_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOperand returns the value of the Operand field in AlertCondition_SdkV2 as
// a AlertConditionOperand_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertCondition_SdkV2) GetOperand(ctx context.Context) (AlertConditionOperand_SdkV2, bool) {
	var e AlertConditionOperand_SdkV2
	if o.Operand.IsNull() || o.Operand.IsUnknown() {
		return e, false
	}
	var v []AlertConditionOperand_SdkV2
	d := o.Operand.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOperand sets the value of the Operand field in AlertCondition_SdkV2.
func (o *AlertCondition_SdkV2) SetOperand(ctx context.Context, v AlertConditionOperand_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["operand"]
	o.Operand = types.ListValueMust(t, vs)
}

// GetThreshold returns the value of the Threshold field in AlertCondition_SdkV2 as
// a AlertConditionThreshold_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertCondition_SdkV2) GetThreshold(ctx context.Context) (AlertConditionThreshold_SdkV2, bool) {
	var e AlertConditionThreshold_SdkV2
	if o.Threshold.IsNull() || o.Threshold.IsUnknown() {
		return e, false
	}
	var v []AlertConditionThreshold_SdkV2
	d := o.Threshold.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThreshold sets the value of the Threshold field in AlertCondition_SdkV2.
func (o *AlertCondition_SdkV2) SetThreshold(ctx context.Context, v AlertConditionThreshold_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["threshold"]
	o.Threshold = types.ListValueMust(t, vs)
}

type AlertConditionOperand_SdkV2 struct {
	Column types.List `tfsdk:"column"`
}

func (newState *AlertConditionOperand_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionOperand_SdkV2) {
}

func (newState *AlertConditionOperand_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertConditionOperand_SdkV2) {
}

func (c AlertConditionOperand_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetOptional()
	attrs["column"] = attrs["column"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConditionOperand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConditionOperand_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertOperandColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionOperand_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertConditionOperand_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": o.Column,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertConditionOperand_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column": basetypes.ListType{
				ElemType: AlertOperandColumn_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumn returns the value of the Column field in AlertConditionOperand_SdkV2 as
// a AlertOperandColumn_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConditionOperand_SdkV2) GetColumn(ctx context.Context) (AlertOperandColumn_SdkV2, bool) {
	var e AlertOperandColumn_SdkV2
	if o.Column.IsNull() || o.Column.IsUnknown() {
		return e, false
	}
	var v []AlertOperandColumn_SdkV2
	d := o.Column.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetColumn sets the value of the Column field in AlertConditionOperand_SdkV2.
func (o *AlertConditionOperand_SdkV2) SetColumn(ctx context.Context, v AlertOperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["column"]
	o.Column = types.ListValueMust(t, vs)
}

type AlertConditionThreshold_SdkV2 struct {
	Value types.List `tfsdk:"value"`
}

func (newState *AlertConditionThreshold_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionThreshold_SdkV2) {
}

func (newState *AlertConditionThreshold_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertConditionThreshold_SdkV2) {
}

func (c AlertConditionThreshold_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value"] = attrs["value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertConditionThreshold.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertConditionThreshold_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(AlertOperandValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionThreshold_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertConditionThreshold_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertConditionThreshold_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: AlertOperandValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValue returns the value of the Value field in AlertConditionThreshold_SdkV2 as
// a AlertOperandValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertConditionThreshold_SdkV2) GetValue(ctx context.Context) (AlertOperandValue_SdkV2, bool) {
	var e AlertOperandValue_SdkV2
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v []AlertOperandValue_SdkV2
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in AlertConditionThreshold_SdkV2.
func (o *AlertConditionThreshold_SdkV2) SetValue(ctx context.Context, v AlertOperandValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type AlertOperandColumn_SdkV2 struct {
	Name types.String `tfsdk:"name"`
}

func (newState *AlertOperandColumn_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandColumn_SdkV2) {
}

func (newState *AlertOperandColumn_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertOperandColumn_SdkV2) {
}

func (c AlertOperandColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOperandColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOperandColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandColumn_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertOperandColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertOperandColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type AlertOperandValue_SdkV2 struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	DoubleValue types.Float64 `tfsdk:"double_value"`

	StringValue types.String `tfsdk:"string_value"`
}

func (newState *AlertOperandValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandValue_SdkV2) {
}

func (newState *AlertOperandValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertOperandValue_SdkV2) {
}

func (c AlertOperandValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bool_value"] = attrs["bool_value"].SetOptional()
	attrs["double_value"] = attrs["double_value"].SetOptional()
	attrs["string_value"] = attrs["string_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOperandValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOperandValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandValue_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertOperandValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"double_value": o.DoubleValue,
			"string_value": o.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertOperandValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"double_value": types.Float64Type,
			"string_value": types.StringType,
		},
	}
}

// Alert configuration options.
type AlertOptions_SdkV2 struct {
	// Name of column in the query result to compare in alert evaluation.
	Column types.String `tfsdk:"column"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted types.Bool `tfsdk:"muted"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op types.String `tfsdk:"op"`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value types.Object `tfsdk:"value"`
}

func (newState *AlertOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOptions_SdkV2) {
}

func (newState *AlertOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertOptions_SdkV2) {
}

func (c AlertOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetRequired()
	attrs["custom_body"] = attrs["custom_body"].SetOptional()
	attrs["custom_subject"] = attrs["custom_subject"].SetOptional()
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["muted"] = attrs["muted"].SetOptional()
	attrs["op"] = attrs["op"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AlertOptions_SdkV2) Type(ctx context.Context) attr.Type {
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

type AlertQuery_SdkV2 struct {
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Query ID.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`

	Options types.List `tfsdk:"options"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`

	Tags types.List `tfsdk:"tags"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`
}

func (newState *AlertQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertQuery_SdkV2) {
}

func (newState *AlertQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertQuery_SdkV2) {
}

func (c AlertQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["data_source_id"] = attrs["data_source_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_archived"] = attrs["is_archived"].SetOptional()
	attrs["is_draft"] = attrs["is_draft"].SetOptional()
	attrs["is_safe"] = attrs["is_safe"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["query"] = attrs["query"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(QueryOptions_SdkV2{}),
		"tags":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AlertQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: QueryOptions_SdkV2{}.Type(ctx),
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

// GetOptions returns the value of the Options field in AlertQuery_SdkV2 as
// a QueryOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertQuery_SdkV2) GetOptions(ctx context.Context) (QueryOptions_SdkV2, bool) {
	var e QueryOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AlertQuery_SdkV2.
func (o *AlertQuery_SdkV2) SetOptions(ctx context.Context, v QueryOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in AlertQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in AlertQuery_SdkV2.
func (o *AlertQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type AlertV2_SdkV2 struct {
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom description for the alert. support mustache template.
	CustomDescription types.String `tfsdk:"custom_description"`
	// Custom summary for the alert. support mustache template.
	CustomSummary types.String `tfsdk:"custom_summary"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`

	Evaluation types.List `tfsdk:"evaluation"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	ParentPath types.String `tfsdk:"parent_path"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// The run as username or application ID of service principal. This field is
	// set to "Unavailable" if the user has been deleted. On Create and Update,
	// this field can be set to application ID of an active service principal.
	// Setting this field requires the servicePrincipal/user role. If not
	// specified it'll default to be request user.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`

	Schedule types.List `tfsdk:"schedule"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *AlertV2_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2_SdkV2) {
}

func (newState *AlertV2_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2_SdkV2) {
}

func (c AlertV2_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["custom_description"] = attrs["custom_description"].SetOptional()
	attrs["custom_summary"] = attrs["custom_summary"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["evaluation"] = attrs["evaluation"].SetOptional()
	attrs["evaluation"] = attrs["evaluation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["schedule"] = attrs["schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"evaluation": reflect.TypeOf(AlertV2Evaluation_SdkV2{}),
		"schedule":   reflect.TypeOf(CronSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":        o.CreateTime,
			"custom_description": o.CustomDescription,
			"custom_summary":     o.CustomSummary,
			"display_name":       o.DisplayName,
			"evaluation":         o.Evaluation,
			"id":                 o.Id,
			"lifecycle_state":    o.LifecycleState,
			"owner_user_name":    o.OwnerUserName,
			"parent_path":        o.ParentPath,
			"query_text":         o.QueryText,
			"run_as_user_name":   o.RunAsUserName,
			"schedule":           o.Schedule,
			"update_time":        o.UpdateTime,
			"warehouse_id":       o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"custom_description": types.StringType,
			"custom_summary":     types.StringType,
			"display_name":       types.StringType,
			"evaluation": basetypes.ListType{
				ElemType: AlertV2Evaluation_SdkV2{}.Type(ctx),
			},
			"id":               types.StringType,
			"lifecycle_state":  types.StringType,
			"owner_user_name":  types.StringType,
			"parent_path":      types.StringType,
			"query_text":       types.StringType,
			"run_as_user_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule_SdkV2{}.Type(ctx),
			},
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetEvaluation returns the value of the Evaluation field in AlertV2_SdkV2 as
// a AlertV2Evaluation_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2_SdkV2) GetEvaluation(ctx context.Context) (AlertV2Evaluation_SdkV2, bool) {
	var e AlertV2Evaluation_SdkV2
	if o.Evaluation.IsNull() || o.Evaluation.IsUnknown() {
		return e, false
	}
	var v []AlertV2Evaluation_SdkV2
	d := o.Evaluation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEvaluation sets the value of the Evaluation field in AlertV2_SdkV2.
func (o *AlertV2_SdkV2) SetEvaluation(ctx context.Context, v AlertV2Evaluation_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["evaluation"]
	o.Evaluation = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in AlertV2_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in AlertV2_SdkV2.
func (o *AlertV2_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

type AlertV2Evaluation_SdkV2 struct {
	// Operator used for comparison in alert evaluation.
	ComparisonOperator types.String `tfsdk:"comparison_operator"`
	// Alert state if result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state"`
	// Timestamp of the last evaluation.
	LastEvaluatedAt types.String `tfsdk:"last_evaluated_at"`
	// User or Notification Destination to notify when alert is triggered.
	Notification types.List `tfsdk:"notification"`
	// Source column from result to use to evaluate alert
	Source types.List `tfsdk:"source"`
	// Latest state of alert evaluation.
	State types.String `tfsdk:"state"`
	// Threshold to user for alert evaluation, can be a column or a value.
	Threshold types.List `tfsdk:"threshold"`
}

func (newState *AlertV2Evaluation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2Evaluation_SdkV2) {
}

func (newState *AlertV2Evaluation_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2Evaluation_SdkV2) {
}

func (c AlertV2Evaluation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comparison_operator"] = attrs["comparison_operator"].SetOptional()
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["last_evaluated_at"] = attrs["last_evaluated_at"].SetComputed()
	attrs["notification"] = attrs["notification"].SetOptional()
	attrs["notification"] = attrs["notification"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].SetOptional()
	attrs["source"] = attrs["source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetComputed()
	attrs["threshold"] = attrs["threshold"].SetOptional()
	attrs["threshold"] = attrs["threshold"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Evaluation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Evaluation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notification": reflect.TypeOf(AlertV2Notification_SdkV2{}),
		"source":       reflect.TypeOf(AlertV2OperandColumn_SdkV2{}),
		"threshold":    reflect.TypeOf(AlertV2Operand_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Evaluation_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2Evaluation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comparison_operator": o.ComparisonOperator,
			"empty_result_state":  o.EmptyResultState,
			"last_evaluated_at":   o.LastEvaluatedAt,
			"notification":        o.Notification,
			"source":              o.Source,
			"state":               o.State,
			"threshold":           o.Threshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Evaluation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comparison_operator": types.StringType,
			"empty_result_state":  types.StringType,
			"last_evaluated_at":   types.StringType,
			"notification": basetypes.ListType{
				ElemType: AlertV2Notification_SdkV2{}.Type(ctx),
			},
			"source": basetypes.ListType{
				ElemType: AlertV2OperandColumn_SdkV2{}.Type(ctx),
			},
			"state": types.StringType,
			"threshold": basetypes.ListType{
				ElemType: AlertV2Operand_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNotification returns the value of the Notification field in AlertV2Evaluation_SdkV2 as
// a AlertV2Notification_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation_SdkV2) GetNotification(ctx context.Context) (AlertV2Notification_SdkV2, bool) {
	var e AlertV2Notification_SdkV2
	if o.Notification.IsNull() || o.Notification.IsUnknown() {
		return e, false
	}
	var v []AlertV2Notification_SdkV2
	d := o.Notification.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotification sets the value of the Notification field in AlertV2Evaluation_SdkV2.
func (o *AlertV2Evaluation_SdkV2) SetNotification(ctx context.Context, v AlertV2Notification_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification"]
	o.Notification = types.ListValueMust(t, vs)
}

// GetSource returns the value of the Source field in AlertV2Evaluation_SdkV2 as
// a AlertV2OperandColumn_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation_SdkV2) GetSource(ctx context.Context) (AlertV2OperandColumn_SdkV2, bool) {
	var e AlertV2OperandColumn_SdkV2
	if o.Source.IsNull() || o.Source.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandColumn_SdkV2
	d := o.Source.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSource sets the value of the Source field in AlertV2Evaluation_SdkV2.
func (o *AlertV2Evaluation_SdkV2) SetSource(ctx context.Context, v AlertV2OperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["source"]
	o.Source = types.ListValueMust(t, vs)
}

// GetThreshold returns the value of the Threshold field in AlertV2Evaluation_SdkV2 as
// a AlertV2Operand_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation_SdkV2) GetThreshold(ctx context.Context) (AlertV2Operand_SdkV2, bool) {
	var e AlertV2Operand_SdkV2
	if o.Threshold.IsNull() || o.Threshold.IsUnknown() {
		return e, false
	}
	var v []AlertV2Operand_SdkV2
	d := o.Threshold.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThreshold sets the value of the Threshold field in AlertV2Evaluation_SdkV2.
func (o *AlertV2Evaluation_SdkV2) SetThreshold(ctx context.Context, v AlertV2Operand_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["threshold"]
	o.Threshold = types.ListValueMust(t, vs)
}

type AlertV2Notification_SdkV2 struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	RetriggerSeconds types.Int64 `tfsdk:"retrigger_seconds"`

	Subscriptions types.List `tfsdk:"subscriptions"`
}

func (newState *AlertV2Notification_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2Notification_SdkV2) {
}

func (newState *AlertV2Notification_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2Notification_SdkV2) {
}

func (c AlertV2Notification_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["retrigger_seconds"] = attrs["retrigger_seconds"].SetOptional()
	attrs["subscriptions"] = attrs["subscriptions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Notification.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Notification_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(AlertV2Subscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Notification_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2Notification_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notify_on_ok":      o.NotifyOnOk,
			"retrigger_seconds": o.RetriggerSeconds,
			"subscriptions":     o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Notification_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notify_on_ok":      types.BoolType,
			"retrigger_seconds": types.Int64Type,
			"subscriptions": basetypes.ListType{
				ElemType: AlertV2Subscription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in AlertV2Notification_SdkV2 as
// a slice of AlertV2Subscription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Notification_SdkV2) GetSubscriptions(ctx context.Context) ([]AlertV2Subscription_SdkV2, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []AlertV2Subscription_SdkV2
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in AlertV2Notification_SdkV2.
func (o *AlertV2Notification_SdkV2) SetSubscriptions(ctx context.Context, v []AlertV2Subscription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type AlertV2Operand_SdkV2 struct {
	Column types.List `tfsdk:"column"`

	Value types.List `tfsdk:"value"`
}

func (newState *AlertV2Operand_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2Operand_SdkV2) {
}

func (newState *AlertV2Operand_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2Operand_SdkV2) {
}

func (c AlertV2Operand_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetOptional()
	attrs["column"] = attrs["column"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value"] = attrs["value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Operand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Operand_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertV2OperandColumn_SdkV2{}),
		"value":  reflect.TypeOf(AlertV2OperandValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Operand_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2Operand_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": o.Column,
			"value":  o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Operand_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column": basetypes.ListType{
				ElemType: AlertV2OperandColumn_SdkV2{}.Type(ctx),
			},
			"value": basetypes.ListType{
				ElemType: AlertV2OperandValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumn returns the value of the Column field in AlertV2Operand_SdkV2 as
// a AlertV2OperandColumn_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Operand_SdkV2) GetColumn(ctx context.Context) (AlertV2OperandColumn_SdkV2, bool) {
	var e AlertV2OperandColumn_SdkV2
	if o.Column.IsNull() || o.Column.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandColumn_SdkV2
	d := o.Column.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetColumn sets the value of the Column field in AlertV2Operand_SdkV2.
func (o *AlertV2Operand_SdkV2) SetColumn(ctx context.Context, v AlertV2OperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["column"]
	o.Column = types.ListValueMust(t, vs)
}

// GetValue returns the value of the Value field in AlertV2Operand_SdkV2 as
// a AlertV2OperandValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Operand_SdkV2) GetValue(ctx context.Context) (AlertV2OperandValue_SdkV2, bool) {
	var e AlertV2OperandValue_SdkV2
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandValue_SdkV2
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in AlertV2Operand_SdkV2.
func (o *AlertV2Operand_SdkV2) SetValue(ctx context.Context, v AlertV2OperandValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type AlertV2OperandColumn_SdkV2 struct {
	Aggregation types.String `tfsdk:"aggregation"`

	Display types.String `tfsdk:"display"`

	Name types.String `tfsdk:"name"`
}

func (newState *AlertV2OperandColumn_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2OperandColumn_SdkV2) {
}

func (newState *AlertV2OperandColumn_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2OperandColumn_SdkV2) {
}

func (c AlertV2OperandColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aggregation"] = attrs["aggregation"].SetOptional()
	attrs["display"] = attrs["display"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2OperandColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2OperandColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandColumn_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2OperandColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aggregation": o.Aggregation,
			"display":     o.Display,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2OperandColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aggregation": types.StringType,
			"display":     types.StringType,
			"name":        types.StringType,
		},
	}
}

type AlertV2OperandValue_SdkV2 struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	DoubleValue types.Float64 `tfsdk:"double_value"`

	StringValue types.String `tfsdk:"string_value"`
}

func (newState *AlertV2OperandValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2OperandValue_SdkV2) {
}

func (newState *AlertV2OperandValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2OperandValue_SdkV2) {
}

func (c AlertV2OperandValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bool_value"] = attrs["bool_value"].SetOptional()
	attrs["double_value"] = attrs["double_value"].SetOptional()
	attrs["string_value"] = attrs["string_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2OperandValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2OperandValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandValue_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2OperandValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"double_value": o.DoubleValue,
			"string_value": o.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2OperandValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"double_value": types.Float64Type,
			"string_value": types.StringType,
		},
	}
}

type AlertV2Subscription_SdkV2 struct {
	DestinationId types.String `tfsdk:"destination_id"`

	UserEmail types.String `tfsdk:"user_email"`
}

func (newState *AlertV2Subscription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertV2Subscription_SdkV2) {
}

func (newState *AlertV2Subscription_SdkV2) SyncEffectiveFieldsDuringRead(existingState AlertV2Subscription_SdkV2) {
}

func (c AlertV2Subscription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_id"] = attrs["destination_id"].SetOptional()
	attrs["user_email"] = attrs["user_email"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Subscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Subscription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Subscription_SdkV2
// only implements ToObjectValue() and Type().
func (o AlertV2Subscription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": o.DestinationId,
			"user_email":     o.UserEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Subscription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
			"user_email":     types.StringType,
		},
	}
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (newState *BaseChunkInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseChunkInfo_SdkV2) {
}

func (newState *BaseChunkInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState BaseChunkInfo_SdkV2) {
}

func (c BaseChunkInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseChunkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseChunkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseChunkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o BaseChunkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o BaseChunkInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":  types.Int64Type,
			"chunk_index": types.Int64Type,
			"row_count":   types.Int64Type,
			"row_offset":  types.Int64Type,
		},
	}
}

type CancelExecutionRequest_SdkV2 struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelExecutionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelExecutionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelExecutionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelExecutionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelExecutionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
		},
	}
}

type CancelExecutionResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelExecutionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelExecutionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelExecutionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelExecutionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelExecutionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel_SdkV2 struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version"`

	Name types.String `tfsdk:"name"`
}

func (newState *Channel_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Channel_SdkV2) {
}

func (newState *Channel_SdkV2) SyncEffectiveFieldsDuringRead(existingState Channel_SdkV2) {
}

func (c Channel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dbsql_version"] = attrs["dbsql_version"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Channel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Channel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Channel_SdkV2
// only implements ToObjectValue() and Type().
func (o Channel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": o.DbsqlVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Channel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbsql_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

// Details about a Channel.
type ChannelInfo_SdkV2 struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion types.String `tfsdk:"dbsql_version"`
	// Name of the channel
	Name types.String `tfsdk:"name"`
}

func (newState *ChannelInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChannelInfo_SdkV2) {
}

func (newState *ChannelInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ChannelInfo_SdkV2) {
}

func (c ChannelInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dbsql_version"] = attrs["dbsql_version"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChannelInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChannelInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChannelInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ChannelInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": o.DbsqlVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChannelInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbsql_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ClientConfig_SdkV2 struct {
	AllowCustomJsVisualizations types.Bool `tfsdk:"allow_custom_js_visualizations"`

	AllowDownloads types.Bool `tfsdk:"allow_downloads"`

	AllowExternalShares types.Bool `tfsdk:"allow_external_shares"`

	AllowSubscriptions types.Bool `tfsdk:"allow_subscriptions"`

	DateFormat types.String `tfsdk:"date_format"`

	DateTimeFormat types.String `tfsdk:"date_time_format"`

	DisablePublish types.Bool `tfsdk:"disable_publish"`

	EnableLegacyAutodetectTypes types.Bool `tfsdk:"enable_legacy_autodetect_types"`

	FeatureShowPermissionsControl types.Bool `tfsdk:"feature_show_permissions_control"`

	HidePlotlyModeBar types.Bool `tfsdk:"hide_plotly_mode_bar"`
}

func (newState *ClientConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClientConfig_SdkV2) {
}

func (newState *ClientConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClientConfig_SdkV2) {
}

func (c ClientConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_custom_js_visualizations"] = attrs["allow_custom_js_visualizations"].SetOptional()
	attrs["allow_downloads"] = attrs["allow_downloads"].SetOptional()
	attrs["allow_external_shares"] = attrs["allow_external_shares"].SetOptional()
	attrs["allow_subscriptions"] = attrs["allow_subscriptions"].SetOptional()
	attrs["date_format"] = attrs["date_format"].SetOptional()
	attrs["date_time_format"] = attrs["date_time_format"].SetOptional()
	attrs["disable_publish"] = attrs["disable_publish"].SetOptional()
	attrs["enable_legacy_autodetect_types"] = attrs["enable_legacy_autodetect_types"].SetOptional()
	attrs["feature_show_permissions_control"] = attrs["feature_show_permissions_control"].SetOptional()
	attrs["hide_plotly_mode_bar"] = attrs["hide_plotly_mode_bar"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClientConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClientConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o ClientConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_custom_js_visualizations":   o.AllowCustomJsVisualizations,
			"allow_downloads":                  o.AllowDownloads,
			"allow_external_shares":            o.AllowExternalShares,
			"allow_subscriptions":              o.AllowSubscriptions,
			"date_format":                      o.DateFormat,
			"date_time_format":                 o.DateTimeFormat,
			"disable_publish":                  o.DisablePublish,
			"enable_legacy_autodetect_types":   o.EnableLegacyAutodetectTypes,
			"feature_show_permissions_control": o.FeatureShowPermissionsControl,
			"hide_plotly_mode_bar":             o.HidePlotlyModeBar,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClientConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_custom_js_visualizations":   types.BoolType,
			"allow_downloads":                  types.BoolType,
			"allow_external_shares":            types.BoolType,
			"allow_subscriptions":              types.BoolType,
			"date_format":                      types.StringType,
			"date_time_format":                 types.StringType,
			"disable_publish":                  types.BoolType,
			"enable_legacy_autodetect_types":   types.BoolType,
			"feature_show_permissions_control": types.BoolType,
			"hide_plotly_mode_bar":             types.BoolType,
		},
	}
}

type ColumnInfo_SdkV2 struct {
	// The name of the column.
	Name types.String `tfsdk:"name"`
	// The ordinal position of the column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// The format of the interval type.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName types.String `tfsdk:"type_name"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// The full SQL type specification.
	TypeText types.String `tfsdk:"type_text"`
}

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo_SdkV2) {
}

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ColumnInfo_SdkV2) {
}

func (c ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["type_interval_type"] = attrs["type_interval_type"].SetOptional()
	attrs["type_name"] = attrs["type_name"].SetOptional()
	attrs["type_precision"] = attrs["type_precision"].SetOptional()
	attrs["type_scale"] = attrs["type_scale"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateAlert_SdkV2 struct {
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.List `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions_SdkV2{}.Type(ctx),
			},
			"parent":   types.StringType,
			"query_id": types.StringType,
			"rearm":    types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in CreateAlert_SdkV2 as
// a AlertOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateAlert_SdkV2.
func (o *CreateAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type CreateAlertRequest_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(CreateAlertRequestAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":                     o.Alert,
			"auto_resolve_display_name": o.AutoResolveDisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: CreateAlertRequestAlert_SdkV2{}.Type(ctx),
			},
			"auto_resolve_display_name": types.BoolType,
		},
	}
}

// GetAlert returns the value of the Alert field in CreateAlertRequest_SdkV2 as
// a CreateAlertRequestAlert_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertRequest_SdkV2) GetAlert(ctx context.Context) (CreateAlertRequestAlert_SdkV2, bool) {
	var e CreateAlertRequestAlert_SdkV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []CreateAlertRequestAlert_SdkV2
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in CreateAlertRequest_SdkV2.
func (o *CreateAlertRequest_SdkV2) SetAlert(ctx context.Context, v CreateAlertRequestAlert_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type CreateAlertRequestAlert_SdkV2 struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger"`
}

func (newState *CreateAlertRequestAlert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlertRequestAlert_SdkV2) {
}

func (newState *CreateAlertRequestAlert_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateAlertRequestAlert_SdkV2) {
}

func (c CreateAlertRequestAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
	attrs["condition"] = attrs["condition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_body"] = attrs["custom_body"].SetOptional()
	attrs["custom_subject"] = attrs["custom_subject"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["seconds_to_retrigger"] = attrs["seconds_to_retrigger"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertRequestAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertRequestAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequestAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAlertRequestAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateAlertRequestAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition_SdkV2{}.Type(ctx),
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

// GetCondition returns the value of the Condition field in CreateAlertRequestAlert_SdkV2 as
// a AlertCondition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertRequestAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in CreateAlertRequestAlert_SdkV2.
func (o *CreateAlertRequestAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type CreateAlertV2Request_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert": o.Alert,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: AlertV2_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAlert returns the value of the Alert field in CreateAlertV2Request_SdkV2 as
// a AlertV2_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertV2Request_SdkV2) GetAlert(ctx context.Context) (AlertV2_SdkV2, bool) {
	var e AlertV2_SdkV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []AlertV2_SdkV2
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in CreateAlertV2Request_SdkV2.
func (o *CreateAlertV2Request_SdkV2) SetAlert(ctx context.Context, v AlertV2_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type CreateQueryRequest_SdkV2 struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Query types.List `tfsdk:"query"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(CreateQueryRequestQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"query":                     o.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_resolve_display_name": types.BoolType,
			"query": basetypes.ListType{
				ElemType: CreateQueryRequestQuery_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQuery returns the value of the Query field in CreateQueryRequest_SdkV2 as
// a CreateQueryRequestQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequest_SdkV2) GetQuery(ctx context.Context) (CreateQueryRequestQuery_SdkV2, bool) {
	var e CreateQueryRequestQuery_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []CreateQueryRequestQuery_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in CreateQueryRequest_SdkV2.
func (o *CreateQueryRequest_SdkV2) SetQuery(ctx context.Context, v CreateQueryRequestQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type CreateQueryRequestQuery_SdkV2 struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema"`

	Tags types.List `tfsdk:"tags"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *CreateQueryRequestQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryRequestQuery_SdkV2) {
}

func (newState *CreateQueryRequestQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateQueryRequestQuery_SdkV2) {
}

func (c CreateQueryRequestQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_auto_limit"] = attrs["apply_auto_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as_mode"] = attrs["run_as_mode"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryRequestQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryRequestQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequestQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateQueryRequestQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateQueryRequestQuery_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit": types.BoolType,
			"catalog":          types.StringType,
			"description":      types.StringType,
			"display_name":     types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter_SdkV2{}.Type(ctx),
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

// GetParameters returns the value of the Parameters field in CreateQueryRequestQuery_SdkV2 as
// a slice of QueryParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequestQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in CreateQueryRequestQuery_SdkV2.
func (o *CreateQueryRequestQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateQueryRequestQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQueryRequestQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in CreateQueryRequestQuery_SdkV2.
func (o *CreateQueryRequestQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest_SdkV2 struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options types.Object `tfsdk:"options"`
	// The identifier returned by :method:queries/create
	QueryId types.String `tfsdk:"query_id"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type_ types.String `tfsdk:"type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryVisualizationsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQueryVisualizationsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryVisualizationsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateQueryVisualizationsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateQueryVisualizationsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateVisualizationRequest_SdkV2 struct {
	Visualization types.List `tfsdk:"visualization"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(CreateVisualizationRequestVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visualization": o.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"visualization": basetypes.ListType{
				ElemType: CreateVisualizationRequestVisualization_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVisualization returns the value of the Visualization field in CreateVisualizationRequest_SdkV2 as
// a CreateVisualizationRequestVisualization_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVisualizationRequest_SdkV2) GetVisualization(ctx context.Context) (CreateVisualizationRequestVisualization_SdkV2, bool) {
	var e CreateVisualizationRequestVisualization_SdkV2
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []CreateVisualizationRequestVisualization_SdkV2
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in CreateVisualizationRequest_SdkV2.
func (o *CreateVisualizationRequest_SdkV2) SetVisualization(ctx context.Context, v CreateVisualizationRequestVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type CreateVisualizationRequestVisualization_SdkV2 struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type"`
}

func (newState *CreateVisualizationRequestVisualization_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVisualizationRequestVisualization_SdkV2) {
}

func (newState *CreateVisualizationRequestVisualization_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateVisualizationRequestVisualization_SdkV2) {
}

func (c CreateVisualizationRequestVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["serialized_options"] = attrs["serialized_options"].SetOptional()
	attrs["serialized_query_plan"] = attrs["serialized_query_plan"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVisualizationRequestVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVisualizationRequestVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequestVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVisualizationRequestVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateVisualizationRequestVisualization_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateWarehouseRequest_SdkV2 struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be >= 0 mins for serverless warehouses - Must be
	// == 0 or >= 10 mins for non-serverless warehouses - 0 indicates no
	// autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.List `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel_SdkV2{}),
		"tags":    reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
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
				ElemType: EndpointTags_SdkV2{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in CreateWarehouseRequest_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWarehouseRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in CreateWarehouseRequest_SdkV2.
func (o *CreateWarehouseRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateWarehouseRequest_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWarehouseRequest_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in CreateWarehouseRequest_SdkV2.
func (o *CreateWarehouseRequest_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type CreateWarehouseResponse_SdkV2 struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id"`
}

func (newState *CreateWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWarehouseResponse_SdkV2) {
}

func (newState *CreateWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateWarehouseResponse_SdkV2) {
}

func (c CreateWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateWidget_SdkV2 struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id"`

	Options types.List `tfsdk:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWidget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWidget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWidget_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWidget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":     o.DashboardId,
			"options":          o.Options,
			"text":             o.Text,
			"visualization_id": o.VisualizationId,
			"width":            o.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWidget_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"options": basetypes.ListType{
				ElemType: WidgetOptions_SdkV2{}.Type(ctx),
			},
			"text":             types.StringType,
			"visualization_id": types.StringType,
			"width":            types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in CreateWidget_SdkV2 as
// a WidgetOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWidget_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateWidget_SdkV2.
func (o *CreateWidget_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type CronSchedule_SdkV2 struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status"`
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule_SdkV2) {
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState CronSchedule_SdkV2) {
}

func (c CronSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pause_status"] = attrs["pause_status"].SetOptional()
	attrs["quartz_cron_schedule"] = attrs["quartz_cron_schedule"].SetOptional()
	attrs["timezone_id"] = attrs["timezone_id"].SetOptional()

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
			"pause_status":         o.PauseStatus,
			"quartz_cron_schedule": o.QuartzCronSchedule,
			"timezone_id":          o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":         types.StringType,
			"quartz_cron_schedule": types.StringType,
			"timezone_id":          types.StringType,
		},
	}
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard_SdkV2 struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit types.Bool `tfsdk:"can_edit"`
	// Timestamp when this dashboard was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled"`
	// The ID for this dashboard.
	Id types.String `tfsdk:"id"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`

	Options types.List `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug types.String `tfsdk:"slug"`

	Tags types.List `tfsdk:"tags"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User types.List `tfsdk:"user"`
	// The ID of the user who owns the dashboard.
	UserId types.Int64 `tfsdk:"user_id"`

	Widgets types.List `tfsdk:"widgets"`
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard_SdkV2) {
}

func (newState *Dashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState Dashboard_SdkV2) {
}

func (c Dashboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_edit"] = attrs["can_edit"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["dashboard_filters_enabled"] = attrs["dashboard_filters_enabled"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_archived"] = attrs["is_archived"].SetOptional()
	attrs["is_draft"] = attrs["is_draft"].SetOptional()
	attrs["is_favorite"] = attrs["is_favorite"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["permission_tier"] = attrs["permission_tier"].SetOptional()
	attrs["slug"] = attrs["slug"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["user_id"] = attrs["user_id"].SetOptional()
	attrs["widgets"] = attrs["widgets"].SetOptional()

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
	return map[string]reflect.Type{
		"options": reflect.TypeOf(DashboardOptions_SdkV2{}),
		"tags":    reflect.TypeOf(types.String{}),
		"user":    reflect.TypeOf(User_SdkV2{}),
		"widgets": reflect.TypeOf(Widget_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dashboard_SdkV2
// only implements ToObjectValue() and Type().
func (o Dashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Dashboard_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: DashboardOptions_SdkV2{}.Type(ctx),
			},
			"parent":          types.StringType,
			"permission_tier": types.StringType,
			"slug":            types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
			"user_id": types.Int64Type,
			"widgets": basetypes.ListType{
				ElemType: Widget_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in Dashboard_SdkV2 as
// a DashboardOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard_SdkV2) GetOptions(ctx context.Context) (DashboardOptions_SdkV2, bool) {
	var e DashboardOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []DashboardOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Dashboard_SdkV2.
func (o *Dashboard_SdkV2) SetOptions(ctx context.Context, v DashboardOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Dashboard_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in Dashboard_SdkV2.
func (o *Dashboard_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in Dashboard_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in Dashboard_SdkV2.
func (o *Dashboard_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

// GetWidgets returns the value of the Widgets field in Dashboard_SdkV2 as
// a slice of Widget_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Dashboard_SdkV2) GetWidgets(ctx context.Context) ([]Widget_SdkV2, bool) {
	if o.Widgets.IsNull() || o.Widgets.IsUnknown() {
		return nil, false
	}
	var v []Widget_SdkV2
	d := o.Widgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWidgets sets the value of the Widgets field in Dashboard_SdkV2.
func (o *Dashboard_SdkV2) SetWidgets(ctx context.Context, v []Widget_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["widgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Widgets = types.ListValueMust(t, vs)
}

type DashboardEditContent_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardEditContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardEditContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEditContent_SdkV2
// only implements ToObjectValue() and Type().
func (o DashboardEditContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DashboardEditContent_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetTags returns the value of the Tags field in DashboardEditContent_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DashboardEditContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in DashboardEditContent_SdkV2.
func (o *DashboardEditContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type DashboardOptions_SdkV2 struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`
}

func (newState *DashboardOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardOptions_SdkV2) {
}

func (newState *DashboardOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState DashboardOptions_SdkV2) {
}

func (c DashboardOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["moved_to_trash_at"] = attrs["moved_to_trash_at"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o DashboardOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"moved_to_trash_at": o.MovedToTrashAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DashboardOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"moved_to_trash_at": types.StringType,
		},
	}
}

type DashboardPostContent_SdkV2 struct {
	// Indicates whether the dashboard filters are enabled
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardPostContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DashboardPostContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardPostContent_SdkV2
// only implements ToObjectValue() and Type().
func (o DashboardPostContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DashboardPostContent_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetTags returns the value of the Tags field in DashboardPostContent_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DashboardPostContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in DashboardPostContent_SdkV2.
func (o *DashboardPostContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource_SdkV2 struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id types.String `tfsdk:"id"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name types.String `tfsdk:"name"`
	// Reserved for internal use.
	PauseReason types.String `tfsdk:"pause_reason"`
	// Reserved for internal use.
	Paused types.Int64 `tfsdk:"paused"`
	// Reserved for internal use.
	SupportsAutoLimit types.Bool `tfsdk:"supports_auto_limit"`
	// Reserved for internal use.
	Syntax types.String `tfsdk:"syntax"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type_ types.String `tfsdk:"type"`
	// Reserved for internal use.
	ViewOnly types.Bool `tfsdk:"view_only"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *DataSource_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataSource_SdkV2) {
}

func (newState *DataSource_SdkV2) SyncEffectiveFieldsDuringRead(existingState DataSource_SdkV2) {
}

func (c DataSource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pause_reason"] = attrs["pause_reason"].SetOptional()
	attrs["paused"] = attrs["paused"].SetOptional()
	attrs["supports_auto_limit"] = attrs["supports_auto_limit"].SetOptional()
	attrs["syntax"] = attrs["syntax"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["view_only"] = attrs["view_only"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataSource_SdkV2
// only implements ToObjectValue() and Type().
func (o DataSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DataSource_SdkV2) Type(ctx context.Context) attr.Type {
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

type DateRange_SdkV2 struct {
	End types.String `tfsdk:"end"`

	Start types.String `tfsdk:"start"`
}

func (newState *DateRange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRange_SdkV2) {
}

func (newState *DateRange_SdkV2) SyncEffectiveFieldsDuringRead(existingState DateRange_SdkV2) {
}

func (c DateRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end"] = attrs["end"].SetRequired()
	attrs["start"] = attrs["start"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRange_SdkV2
// only implements ToObjectValue() and Type().
func (o DateRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end":   o.End,
			"start": o.Start,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DateRange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end":   types.StringType,
			"start": types.StringType,
		},
	}
}

type DateRangeValue_SdkV2 struct {
	// Manually specified date-time range value.
	DateRangeValue types.List `tfsdk:"date_range_value"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue types.String `tfsdk:"dynamic_date_range_value"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision"`

	StartDayOfWeek types.Int64 `tfsdk:"start_day_of_week"`
}

func (newState *DateRangeValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRangeValue_SdkV2) {
}

func (newState *DateRangeValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState DateRangeValue_SdkV2) {
}

func (c DateRangeValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["date_range_value"] = attrs["date_range_value"].SetOptional()
	attrs["date_range_value"] = attrs["date_range_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["dynamic_date_range_value"] = attrs["dynamic_date_range_value"].SetOptional()
	attrs["precision"] = attrs["precision"].SetOptional()
	attrs["start_day_of_week"] = attrs["start_day_of_week"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateRangeValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateRangeValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"date_range_value": reflect.TypeOf(DateRange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRangeValue_SdkV2
// only implements ToObjectValue() and Type().
func (o DateRangeValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DateRangeValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_range_value": basetypes.ListType{
				ElemType: DateRange_SdkV2{}.Type(ctx),
			},
			"dynamic_date_range_value": types.StringType,
			"precision":                types.StringType,
			"start_day_of_week":        types.Int64Type,
		},
	}
}

// GetDateRangeValue returns the value of the DateRangeValue field in DateRangeValue_SdkV2 as
// a DateRange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *DateRangeValue_SdkV2) GetDateRangeValue(ctx context.Context) (DateRange_SdkV2, bool) {
	var e DateRange_SdkV2
	if o.DateRangeValue.IsNull() || o.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRange_SdkV2
	d := o.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in DateRangeValue_SdkV2.
func (o *DateRangeValue_SdkV2) SetDateRangeValue(ctx context.Context, v DateRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	o.DateRangeValue = types.ListValueMust(t, vs)
}

type DateValue_SdkV2 struct {
	// Manually specified date-time value.
	DateValue types.String `tfsdk:"date_value"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue types.String `tfsdk:"dynamic_date_value"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision"`
}

func (newState *DateValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateValue_SdkV2) {
}

func (newState *DateValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState DateValue_SdkV2) {
}

func (c DateValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["date_value"] = attrs["date_value"].SetOptional()
	attrs["dynamic_date_value"] = attrs["dynamic_date_value"].SetOptional()
	attrs["precision"] = attrs["precision"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DateValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DateValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateValue_SdkV2
// only implements ToObjectValue() and Type().
func (o DateValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_value":         o.DateValue,
			"dynamic_date_value": o.DynamicDateValue,
			"precision":          o.Precision,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DateValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_value":         types.StringType,
			"dynamic_date_value": types.StringType,
			"precision":          types.StringType,
		},
	}
}

type DeleteAlertsLegacyRequest_SdkV2 struct {
	AlertId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAlertsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAlertsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAlertsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAlertsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": o.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAlertsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

type DeleteDashboardRequest_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type DeleteDashboardWidgetRequest_SdkV2 struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardWidgetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDashboardWidgetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardWidgetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDashboardWidgetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardWidgetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type DeleteQueryVisualizationsLegacyRequest_SdkV2 struct {
	// Widget ID returned by :method:queryvisualizations/create
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQueryVisualizationsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQueryVisualizationsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueryVisualizationsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteQueryVisualizationsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQueryVisualizationsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse_SdkV2) {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteResponse_SdkV2) {
}

func (c DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteVisualizationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWarehouseRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWarehouseResponse_SdkV2 struct {
}

func (newState *DeleteWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWarehouseResponse_SdkV2) {
}

func (newState *DeleteWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteWarehouseResponse_SdkV2) {
}

func (c DeleteWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditAlert_SdkV2 struct {
	AlertId types.String `tfsdk:"-"`
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.List `tfsdk:"options"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o EditAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EditAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
			"name":     types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions_SdkV2{}.Type(ctx),
			},
			"query_id": types.StringType,
			"rearm":    types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in EditAlert_SdkV2 as
// a AlertOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in EditAlert_SdkV2.
func (o *EditAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type EditWarehouseRequest_SdkV2 struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.List `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Required. Id of the warehouse to configure.
	Id types.String `tfsdk:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel_SdkV2{}),
		"tags":    reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o EditWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EditWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
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
				ElemType: EndpointTags_SdkV2{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in EditWarehouseRequest_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditWarehouseRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EditWarehouseRequest_SdkV2.
func (o *EditWarehouseRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EditWarehouseRequest_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditWarehouseRequest_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EditWarehouseRequest_SdkV2.
func (o *EditWarehouseRequest_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type EditWarehouseResponse_SdkV2 struct {
}

func (newState *EditWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditWarehouseResponse_SdkV2) {
}

func (newState *EditWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState EditWarehouseResponse_SdkV2) {
}

func (c EditWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty_SdkV2 struct {
}

func (newState *Empty_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty_SdkV2) {
}

func (newState *Empty_SdkV2) SyncEffectiveFieldsDuringRead(existingState Empty_SdkV2) {
}

func (c Empty_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Empty_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty_SdkV2
// only implements ToObjectValue() and Type().
func (o Empty_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o Empty_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EndpointConfPair_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (newState *EndpointConfPair_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointConfPair_SdkV2) {
}

func (newState *EndpointConfPair_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointConfPair_SdkV2) {
}

func (c EndpointConfPair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointConfPair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointConfPair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointConfPair_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointConfPair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointConfPair_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointHealth_SdkV2 struct {
	// Details about errors that are causing current degraded/failed status.
	Details types.String `tfsdk:"details"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason types.List `tfsdk:"failure_reason"`
	// Deprecated. split into summary and details for security
	Message types.String `tfsdk:"message"`

	Status types.String `tfsdk:"status"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary"`
}

func (newState *EndpointHealth_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointHealth_SdkV2) {
}

func (newState *EndpointHealth_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointHealth_SdkV2) {
}

func (c EndpointHealth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["failure_reason"] = attrs["failure_reason"].SetOptional()
	attrs["failure_reason"] = attrs["failure_reason"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["message"] = attrs["message"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["summary"] = attrs["summary"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointHealth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointHealth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failure_reason": reflect.TypeOf(TerminationReason_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointHealth_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointHealth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EndpointHealth_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": types.StringType,
			"failure_reason": basetypes.ListType{
				ElemType: TerminationReason_SdkV2{}.Type(ctx),
			},
			"message": types.StringType,
			"status":  types.StringType,
			"summary": types.StringType,
		},
	}
}

// GetFailureReason returns the value of the FailureReason field in EndpointHealth_SdkV2 as
// a TerminationReason_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointHealth_SdkV2) GetFailureReason(ctx context.Context) (TerminationReason_SdkV2, bool) {
	var e TerminationReason_SdkV2
	if o.FailureReason.IsNull() || o.FailureReason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason_SdkV2
	d := o.FailureReason.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailureReason sets the value of the FailureReason field in EndpointHealth_SdkV2.
func (o *EndpointHealth_SdkV2) SetFailureReason(ctx context.Context, v TerminationReason_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failure_reason"]
	o.FailureReason = types.ListValueMust(t, vs)
}

type EndpointInfo_SdkV2 struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.List `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.List `tfsdk:"health"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.List `tfsdk:"odbc_params"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`

	State types.String `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (newState *EndpointInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointInfo_SdkV2) {
}

func (newState *EndpointInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointInfo_SdkV2) {
}

func (c EndpointInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["health"] = attrs["health"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["jdbc_url"] = attrs["jdbc_url"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["num_active_sessions"] = attrs["num_active_sessions"].SetOptional()
	attrs["num_clusters"] = attrs["num_clusters"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["tags"] = attrs["tags"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["warehouse_type"] = attrs["warehouse_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":     reflect.TypeOf(Channel_SdkV2{}),
		"health":      reflect.TypeOf(EndpointHealth_SdkV2{}),
		"odbc_params": reflect.TypeOf(OdbcParams_SdkV2{}),
		"tags":        reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health": basetypes.ListType{
				ElemType: EndpointHealth_SdkV2{}.Type(ctx),
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
				ElemType: OdbcParams_SdkV2{}.Type(ctx),
			},
			"spot_instance_policy": types.StringType,
			"state":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags_SdkV2{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in EndpointInfo_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EndpointInfo_SdkV2.
func (o *EndpointInfo_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in EndpointInfo_SdkV2 as
// a EndpointHealth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo_SdkV2) GetHealth(ctx context.Context) (EndpointHealth_SdkV2, bool) {
	var e EndpointHealth_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in EndpointInfo_SdkV2.
func (o *EndpointInfo_SdkV2) SetHealth(ctx context.Context, v EndpointHealth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in EndpointInfo_SdkV2 as
// a OdbcParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo_SdkV2) GetOdbcParams(ctx context.Context) (OdbcParams_SdkV2, bool) {
	var e OdbcParams_SdkV2
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams_SdkV2
	d := o.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in EndpointInfo_SdkV2.
func (o *EndpointInfo_SdkV2) SetOdbcParams(ctx context.Context, v OdbcParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	o.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EndpointInfo_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EndpointInfo_SdkV2.
func (o *EndpointInfo_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type EndpointTagPair_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (newState *EndpointTagPair_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTagPair_SdkV2) {
}

func (newState *EndpointTagPair_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointTagPair_SdkV2) {
}

func (c EndpointTagPair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTagPair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointTagPair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTagPair_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointTagPair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTagPair_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointTags_SdkV2 struct {
	CustomTags types.List `tfsdk:"custom_tags"`
}

func (newState *EndpointTags_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTags_SdkV2) {
}

func (newState *EndpointTags_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointTags_SdkV2) {
}

func (c EndpointTags_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTags.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointTags_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(EndpointTagPair_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTags_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointTags_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags": o.CustomTags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTags_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: EndpointTagPair_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EndpointTags_SdkV2 as
// a slice of EndpointTagPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointTags_SdkV2) GetCustomTags(ctx context.Context) ([]EndpointTagPair_SdkV2, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTagPair_SdkV2
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EndpointTags_SdkV2.
func (o *EndpointTags_SdkV2) SetCustomTags(ctx context.Context, v []EndpointTagPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.ListValueMust(t, vs)
}

type EnumValue_SdkV2 struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions types.String `tfsdk:"enum_options"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (newState *EnumValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnumValue_SdkV2) {
}

func (newState *EnumValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnumValue_SdkV2) {
}

func (c EnumValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enum_options"] = attrs["enum_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnumValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnumValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnumValue_SdkV2
// only implements ToObjectValue() and Type().
func (o EnumValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enum_options":         o.EnumOptions,
			"multi_values_options": o.MultiValuesOptions,
			"values":               o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnumValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enum_options": types.StringType,
			"multi_values_options": basetypes.ListType{
				ElemType: MultiValuesOptions_SdkV2{}.Type(ctx),
			},
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in EnumValue_SdkV2 as
// a MultiValuesOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EnumValue_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in EnumValue_SdkV2.
func (o *EnumValue_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in EnumValue_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EnumValue_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
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

// SetValues sets the value of the Values field in EnumValue_SdkV2.
func (o *EnumValue_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type ExecuteStatementRequest_SdkV2 struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog"`

	Disposition types.String `tfsdk:"disposition"`
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
	Format types.String `tfsdk:"format"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	OnWaitTimeout types.String `tfsdk:"on_wait_timeout"`
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
	Parameters types.List `tfsdk:"parameters"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit types.Int64 `tfsdk:"row_limit"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema types.String `tfsdk:"schema"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	Statement types.String `tfsdk:"statement"`
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
	WaitTimeout types.String `tfsdk:"wait_timeout"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExecuteStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExecuteStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(StatementParameterListItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExecuteStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExecuteStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_limit":      types.Int64Type,
			"catalog":         types.StringType,
			"disposition":     types.StringType,
			"format":          types.StringType,
			"on_wait_timeout": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: StatementParameterListItem_SdkV2{}.Type(ctx),
			},
			"row_limit":    types.Int64Type,
			"schema":       types.StringType,
			"statement":    types.StringType,
			"wait_timeout": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in ExecuteStatementRequest_SdkV2 as
// a slice of StatementParameterListItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExecuteStatementRequest_SdkV2) GetParameters(ctx context.Context) ([]StatementParameterListItem_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []StatementParameterListItem_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ExecuteStatementRequest_SdkV2.
func (o *ExecuteStatementRequest_SdkV2) SetParameters(ctx context.Context, v []StatementParameterListItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type ExternalLink_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration types.String `tfsdk:"expiration"`

	ExternalLink types.String `tfsdk:"external_link"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (newState *ExternalLink_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLink_SdkV2) {
}

func (newState *ExternalLink_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalLink_SdkV2) {
}

func (c ExternalLink_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["expiration"] = attrs["expiration"].SetOptional()
	attrs["external_link"] = attrs["external_link"].SetOptional()
	attrs["http_headers"] = attrs["http_headers"].SetOptional()
	attrs["next_chunk_index"] = attrs["next_chunk_index"].SetOptional()
	attrs["next_chunk_internal_link"] = attrs["next_chunk_internal_link"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalLink.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalLink_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"http_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLink_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalLink_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExternalLink_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetHttpHeaders returns the value of the HttpHeaders field in ExternalLink_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalLink_SdkV2) GetHttpHeaders(ctx context.Context) (map[string]types.String, bool) {
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

// SetHttpHeaders sets the value of the HttpHeaders field in ExternalLink_SdkV2.
func (o *ExternalLink_SdkV2) SetHttpHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["http_headers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.HttpHeaders = types.MapValueMust(t, vs)
}

type ExternalQuerySource_SdkV2 struct {
	// The canonical identifier for this SQL alert
	AlertId types.String `tfsdk:"alert_id"`
	// The canonical identifier for this Lakeview dashboard
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The canonical identifier for this Genie space
	GenieSpaceId types.String `tfsdk:"genie_space_id"`

	JobInfo types.List `tfsdk:"job_info"`
	// The canonical identifier for this legacy dashboard
	LegacyDashboardId types.String `tfsdk:"legacy_dashboard_id"`
	// The canonical identifier for this notebook
	NotebookId types.String `tfsdk:"notebook_id"`
	// The canonical identifier for this SQL query
	SqlQueryId types.String `tfsdk:"sql_query_id"`
}

func (newState *ExternalQuerySource_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalQuerySource_SdkV2) {
}

func (newState *ExternalQuerySource_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalQuerySource_SdkV2) {
}

func (c ExternalQuerySource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_id"] = attrs["alert_id"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["genie_space_id"] = attrs["genie_space_id"].SetOptional()
	attrs["job_info"] = attrs["job_info"].SetOptional()
	attrs["job_info"] = attrs["job_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["legacy_dashboard_id"] = attrs["legacy_dashboard_id"].SetOptional()
	attrs["notebook_id"] = attrs["notebook_id"].SetOptional()
	attrs["sql_query_id"] = attrs["sql_query_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalQuerySource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalQuerySource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_info": reflect.TypeOf(ExternalQuerySourceJobInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySource_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalQuerySource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id":            o.AlertId,
			"dashboard_id":        o.DashboardId,
			"genie_space_id":      o.GenieSpaceId,
			"job_info":            o.JobInfo,
			"legacy_dashboard_id": o.LegacyDashboardId,
			"notebook_id":         o.NotebookId,
			"sql_query_id":        o.SqlQueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalQuerySource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id":       types.StringType,
			"dashboard_id":   types.StringType,
			"genie_space_id": types.StringType,
			"job_info": basetypes.ListType{
				ElemType: ExternalQuerySourceJobInfo_SdkV2{}.Type(ctx),
			},
			"legacy_dashboard_id": types.StringType,
			"notebook_id":         types.StringType,
			"sql_query_id":        types.StringType,
		},
	}
}

// GetJobInfo returns the value of the JobInfo field in ExternalQuerySource_SdkV2 as
// a ExternalQuerySourceJobInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalQuerySource_SdkV2) GetJobInfo(ctx context.Context) (ExternalQuerySourceJobInfo_SdkV2, bool) {
	var e ExternalQuerySourceJobInfo_SdkV2
	if o.JobInfo.IsNull() || o.JobInfo.IsUnknown() {
		return e, false
	}
	var v []ExternalQuerySourceJobInfo_SdkV2
	d := o.JobInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobInfo sets the value of the JobInfo field in ExternalQuerySource_SdkV2.
func (o *ExternalQuerySource_SdkV2) SetJobInfo(ctx context.Context, v ExternalQuerySourceJobInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_info"]
	o.JobInfo = types.ListValueMust(t, vs)
}

type ExternalQuerySourceJobInfo_SdkV2 struct {
	// The canonical identifier for this job.
	JobId types.String `tfsdk:"job_id"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	JobRunId types.String `tfsdk:"job_run_id"`
	// The canonical identifier of the task run.
	JobTaskRunId types.String `tfsdk:"job_task_run_id"`
}

func (newState *ExternalQuerySourceJobInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalQuerySourceJobInfo_SdkV2) {
}

func (newState *ExternalQuerySourceJobInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalQuerySourceJobInfo_SdkV2) {
}

func (c ExternalQuerySourceJobInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["job_id"] = attrs["job_id"].SetOptional()
	attrs["job_run_id"] = attrs["job_run_id"].SetOptional()
	attrs["job_task_run_id"] = attrs["job_task_run_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalQuerySourceJobInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalQuerySourceJobInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySourceJobInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalQuerySourceJobInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":          o.JobId,
			"job_run_id":      o.JobRunId,
			"job_task_run_id": o.JobTaskRunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalQuerySourceJobInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":          types.StringType,
			"job_run_id":      types.StringType,
			"job_task_run_id": types.StringType,
		},
	}
}

type GetAlertRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAlertV2Request_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAlertsLegacyRequest_SdkV2 struct {
	AlertId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertsLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAlertsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": o.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

type GetConfigRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetDashboardRequest_SdkV2 struct {
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

type GetDbsqlPermissionRequest_SdkV2 struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permissions to check.
	ObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDbsqlPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDbsqlPermissionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDbsqlPermissionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDbsqlPermissionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objectId":   o.ObjectId,
			"objectType": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDbsqlPermissionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"objectId":   types.StringType,
			"objectType": types.StringType,
		},
	}
}

type GetQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type GetQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetResponse_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *GetResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetResponse_SdkV2) {
}

func (newState *GetResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetResponse_SdkV2) {
}

func (c GetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in GetResponse_SdkV2 as
// a slice of AccessControl_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetResponse_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in GetResponse_SdkV2.
func (o *GetResponse_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type GetStatementRequest_SdkV2 struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
		},
	}
}

type GetStatementResultChunkNRequest_SdkV2 struct {
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatementResultChunkNRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatementResultChunkNRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementResultChunkNRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStatementResultChunkNRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_index":  o.ChunkIndex,
			"statement_id": o.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatementResultChunkNRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunk_index":  types.Int64Type,
			"statement_id": types.StringType,
		},
	}
}

type GetWarehousePermissionLevelsRequest_SdkV2 struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
		},
	}
}

type GetWarehousePermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetWarehousePermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionLevelsResponse_SdkV2) {
}

func (newState *GetWarehousePermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionLevelsResponse_SdkV2) {
}

func (c GetWarehousePermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WarehousePermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: WarehousePermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetWarehousePermissionLevelsResponse_SdkV2 as
// a slice of WarehousePermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehousePermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]WarehousePermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWarehousePermissionLevelsResponse_SdkV2.
func (o *GetWarehousePermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []WarehousePermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWarehousePermissionsRequest_SdkV2 struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehousePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehousePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWarehousePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehousePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
		},
	}
}

type GetWarehouseRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehouseRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetWarehouseResponse_SdkV2 struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.List `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.List `tfsdk:"health"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.List `tfsdk:"odbc_params"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`

	State types.String `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.List `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (newState *GetWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehouseResponse_SdkV2) {
}

func (newState *GetWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetWarehouseResponse_SdkV2) {
}

func (c GetWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["health"] = attrs["health"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["jdbc_url"] = attrs["jdbc_url"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["num_active_sessions"] = attrs["num_active_sessions"].SetOptional()
	attrs["num_clusters"] = attrs["num_clusters"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["tags"] = attrs["tags"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["warehouse_type"] = attrs["warehouse_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":     reflect.TypeOf(Channel_SdkV2{}),
		"health":      reflect.TypeOf(EndpointHealth_SdkV2{}),
		"odbc_params": reflect.TypeOf(OdbcParams_SdkV2{}),
		"tags":        reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_stop_mins": types.Int64Type,
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
			},
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health": basetypes.ListType{
				ElemType: EndpointHealth_SdkV2{}.Type(ctx),
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
				ElemType: OdbcParams_SdkV2{}.Type(ctx),
			},
			"spot_instance_policy": types.StringType,
			"state":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTags_SdkV2{}.Type(ctx),
			},
			"warehouse_type": types.StringType,
		},
	}
}

// GetChannel returns the value of the Channel field in GetWarehouseResponse_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWarehouseResponse_SdkV2.
func (o *GetWarehouseResponse_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in GetWarehouseResponse_SdkV2 as
// a EndpointHealth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse_SdkV2) GetHealth(ctx context.Context) (EndpointHealth_SdkV2, bool) {
	var e EndpointHealth_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in GetWarehouseResponse_SdkV2.
func (o *GetWarehouseResponse_SdkV2) SetHealth(ctx context.Context, v EndpointHealth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in GetWarehouseResponse_SdkV2 as
// a OdbcParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse_SdkV2) GetOdbcParams(ctx context.Context) (OdbcParams_SdkV2, bool) {
	var e OdbcParams_SdkV2
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams_SdkV2
	d := o.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in GetWarehouseResponse_SdkV2.
func (o *GetWarehouseResponse_SdkV2) SetOdbcParams(ctx context.Context, v OdbcParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	o.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in GetWarehouseResponse_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in GetWarehouseResponse_SdkV2.
func (o *GetWarehouseResponse_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	o.Tags = types.ListValueMust(t, vs)
}

type GetWorkspaceWarehouseConfigRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceWarehouseConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceWarehouseConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceWarehouseConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceWarehouseConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetWorkspaceWarehouseConfigResponse_SdkV2 struct {
	// Optional: Channel selection details
	Channel types.List `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.List `tfsdk:"config_param"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.List `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters"`
}

func (newState *GetWorkspaceWarehouseConfigResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (newState *GetWorkspaceWarehouseConfigResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (c GetWorkspaceWarehouseConfigResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["config_param"] = attrs["config_param"].SetOptional()
	attrs["config_param"] = attrs["config_param"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["data_access_config"] = attrs["data_access_config"].SetOptional()
	attrs["enabled_warehouse_types"] = attrs["enabled_warehouse_types"].SetOptional()
	attrs["global_param"] = attrs["global_param"].SetOptional()
	attrs["global_param"] = attrs["global_param"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["google_service_account"] = attrs["google_service_account"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["security_policy"] = attrs["security_policy"].SetOptional()
	attrs["sql_configuration_parameters"] = attrs["sql_configuration_parameters"].SetOptional()
	attrs["sql_configuration_parameters"] = attrs["sql_configuration_parameters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceWarehouseConfigResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceWarehouseConfigResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":                      reflect.TypeOf(Channel_SdkV2{}),
		"config_param":                 reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
		"data_access_config":           reflect.TypeOf(EndpointConfPair_SdkV2{}),
		"enabled_warehouse_types":      reflect.TypeOf(WarehouseTypePair_SdkV2{}),
		"global_param":                 reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
		"sql_configuration_parameters": reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceWarehouseConfigResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceWarehouseConfigResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetWorkspaceWarehouseConfigResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
			},
			"config_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair_SdkV2{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair_SdkV2{}.Type(ctx),
			},
			"global_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
			"google_service_account": types.StringType,
			"instance_profile_arn":   types.StringType,
			"security_policy":        types.StringType,
			"sql_configuration_parameters": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetChannel returns the value of the Channel field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	o.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if o.DataAccessConfig.IsNull() || o.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := o.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a slice of WarehouseTypePair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair_SdkV2, bool) {
	if o.EnabledWarehouseTypes.IsNull() || o.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair_SdkV2
	d := o.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.GlobalParam.IsNull() || o.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	o.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (o *GetWorkspaceWarehouseConfigResponse_SdkV2) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	o.SqlConfigurationParameters = types.ListValueMust(t, vs)
}

type LegacyAlert_SdkV2 struct {
	// Timestamp when the alert was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Alert ID.
	Id types.String `tfsdk:"id"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt types.String `tfsdk:"last_triggered_at"`
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.List `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`

	Query types.List `tfsdk:"query"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State types.String `tfsdk:"state"`
	// Timestamp when the alert was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User types.List `tfsdk:"user"`
}

func (newState *LegacyAlert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyAlert_SdkV2) {
}

func (newState *LegacyAlert_SdkV2) SyncEffectiveFieldsDuringRead(existingState LegacyAlert_SdkV2) {
}

func (c LegacyAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_triggered_at"] = attrs["last_triggered_at"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query"] = attrs["query"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["rearm"] = attrs["rearm"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
		"query":   reflect.TypeOf(AlertQuery_SdkV2{}),
		"user":    reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o LegacyAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LegacyAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":        types.StringType,
			"id":                types.StringType,
			"last_triggered_at": types.StringType,
			"name":              types.StringType,
			"options": basetypes.ListType{
				ElemType: AlertOptions_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
			"query": basetypes.ListType{
				ElemType: AlertQuery_SdkV2{}.Type(ctx),
			},
			"rearm":      types.Int64Type,
			"state":      types.StringType,
			"updated_at": types.StringType,
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in LegacyAlert_SdkV2 as
// a AlertOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyAlert_SdkV2.
func (o *LegacyAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetQuery returns the value of the Query field in LegacyAlert_SdkV2 as
// a AlertQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert_SdkV2) GetQuery(ctx context.Context) (AlertQuery_SdkV2, bool) {
	var e AlertQuery_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []AlertQuery_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyAlert_SdkV2.
func (o *LegacyAlert_SdkV2) SetQuery(ctx context.Context, v AlertQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyAlert_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyAlert_SdkV2.
func (o *LegacyAlert_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

type LegacyQuery_SdkV2 struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit types.Bool `tfsdk:"can_edit"`
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Query ID.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe"`

	LastModifiedBy types.List `tfsdk:"last_modified_by"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById types.Int64 `tfsdk:"last_modified_by_id"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId types.String `tfsdk:"latest_query_data_id"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`

	Options types.List `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash types.String `tfsdk:"query_hash"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User types.List `tfsdk:"user"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`

	Visualizations types.List `tfsdk:"visualizations"`
}

func (newState *LegacyQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyQuery_SdkV2) {
}

func (newState *LegacyQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState LegacyQuery_SdkV2) {
}

func (c LegacyQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_edit"] = attrs["can_edit"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["data_source_id"] = attrs["data_source_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_archived"] = attrs["is_archived"].SetOptional()
	attrs["is_draft"] = attrs["is_draft"].SetOptional()
	attrs["is_favorite"] = attrs["is_favorite"].SetOptional()
	attrs["is_safe"] = attrs["is_safe"].SetOptional()
	attrs["last_modified_by"] = attrs["last_modified_by"].SetOptional()
	attrs["last_modified_by"] = attrs["last_modified_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_modified_by_id"] = attrs["last_modified_by_id"].SetOptional()
	attrs["latest_query_data_id"] = attrs["latest_query_data_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["permission_tier"] = attrs["permission_tier"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query_hash"] = attrs["query_hash"].SetOptional()
	attrs["run_as_role"] = attrs["run_as_role"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["user_id"] = attrs["user_id"].SetOptional()
	attrs["visualizations"] = attrs["visualizations"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"last_modified_by": reflect.TypeOf(User_SdkV2{}),
		"options":          reflect.TypeOf(QueryOptions_SdkV2{}),
		"tags":             reflect.TypeOf(types.String{}),
		"user":             reflect.TypeOf(User_SdkV2{}),
		"visualizations":   reflect.TypeOf(LegacyVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o LegacyQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LegacyQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: User_SdkV2{}.Type(ctx),
			},
			"last_modified_by_id":  types.Int64Type,
			"latest_query_data_id": types.StringType,
			"name":                 types.StringType,
			"options": basetypes.ListType{
				ElemType: QueryOptions_SdkV2{}.Type(ctx),
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
				ElemType: User_SdkV2{}.Type(ctx),
			},
			"user_id": types.Int64Type,
			"visualizations": basetypes.ListType{
				ElemType: LegacyVisualization_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetLastModifiedBy returns the value of the LastModifiedBy field in LegacyQuery_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery_SdkV2) GetLastModifiedBy(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if o.LastModifiedBy.IsNull() || o.LastModifiedBy.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := o.LastModifiedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLastModifiedBy sets the value of the LastModifiedBy field in LegacyQuery_SdkV2.
func (o *LegacyQuery_SdkV2) SetLastModifiedBy(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["last_modified_by"]
	o.LastModifiedBy = types.ListValueMust(t, vs)
}

// GetOptions returns the value of the Options field in LegacyQuery_SdkV2 as
// a QueryOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery_SdkV2) GetOptions(ctx context.Context) (QueryOptions_SdkV2, bool) {
	var e QueryOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyQuery_SdkV2.
func (o *LegacyQuery_SdkV2) SetOptions(ctx context.Context, v QueryOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LegacyQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in LegacyQuery_SdkV2.
func (o *LegacyQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyQuery_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := o.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyQuery_SdkV2.
func (o *LegacyQuery_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

// GetVisualizations returns the value of the Visualizations field in LegacyQuery_SdkV2 as
// a slice of LegacyVisualization_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery_SdkV2) GetVisualizations(ctx context.Context) ([]LegacyVisualization_SdkV2, bool) {
	if o.Visualizations.IsNull() || o.Visualizations.IsUnknown() {
		return nil, false
	}
	var v []LegacyVisualization_SdkV2
	d := o.Visualizations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualizations sets the value of the Visualizations field in LegacyQuery_SdkV2.
func (o *LegacyQuery_SdkV2) SetVisualizations(ctx context.Context, v []LegacyVisualization_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualizations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Visualizations = types.ListValueMust(t, vs)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization_SdkV2 struct {
	CreatedAt types.String `tfsdk:"created_at"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description"`
	// The UUID for this visualization.
	Id types.String `tfsdk:"id"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options types.Object `tfsdk:"options"`

	Query types.List `tfsdk:"query"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type_ types.String `tfsdk:"type"`

	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (newState *LegacyVisualization_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyVisualization_SdkV2) {
}

func (newState *LegacyVisualization_SdkV2) SyncEffectiveFieldsDuringRead(existingState LegacyVisualization_SdkV2) {
}

func (c LegacyVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query"] = attrs["query"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["type"] = attrs["type"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LegacyVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LegacyVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(LegacyQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (o LegacyVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LegacyVisualization_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":  types.StringType,
			"description": types.StringType,
			"id":          types.StringType,
			"name":        types.StringType,
			"options":     types.ObjectType{},
			"query": basetypes.ListType{
				ElemType: LegacyQuery_SdkV2{}.Type(ctx),
			},
			"type":       types.StringType,
			"updated_at": types.StringType,
		},
	}
}

// GetQuery returns the value of the Query field in LegacyVisualization_SdkV2 as
// a LegacyQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyVisualization_SdkV2) GetQuery(ctx context.Context) (LegacyQuery_SdkV2, bool) {
	var e LegacyQuery_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []LegacyQuery_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyVisualization_SdkV2.
func (o *LegacyVisualization_SdkV2) SetQuery(ctx context.Context, v LegacyQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type ListAlertsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAlertsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAlertsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (newState *ListAlertsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponse_SdkV2) {
}

func (newState *ListAlertsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponse_SdkV2) {
}

func (c ListAlertsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListAlertsResponseAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAlertsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListAlertsResponseAlert_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListAlertsResponse_SdkV2 as
// a slice of ListAlertsResponseAlert_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsResponse_SdkV2) GetResults(ctx context.Context) ([]ListAlertsResponseAlert_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListAlertsResponseAlert_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListAlertsResponse_SdkV2.
func (o *ListAlertsResponse_SdkV2) SetResults(ctx context.Context, v []ListAlertsResponseAlert_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListAlertsResponseAlert_SdkV2 struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *ListAlertsResponseAlert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponseAlert_SdkV2) {
}

func (newState *ListAlertsResponseAlert_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponseAlert_SdkV2) {
}

func (c ListAlertsResponseAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
	attrs["condition"] = attrs["condition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["custom_body"] = attrs["custom_body"].SetOptional()
	attrs["custom_subject"] = attrs["custom_subject"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetOptional()
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["seconds_to_retrigger"] = attrs["seconds_to_retrigger"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["trigger_time"] = attrs["trigger_time"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsResponseAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsResponseAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponseAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAlertsResponseAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListAlertsResponseAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition_SdkV2{}.Type(ctx),
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

// GetCondition returns the value of the Condition field in ListAlertsResponseAlert_SdkV2 as
// a AlertCondition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsResponseAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in ListAlertsResponseAlert_SdkV2.
func (o *ListAlertsResponseAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type ListAlertsV2Request_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAlertsV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAlertsV2Response_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (newState *ListAlertsV2Response_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsV2Response_SdkV2) {
}

func (newState *ListAlertsV2Response_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAlertsV2Response_SdkV2) {
}

func (c ListAlertsV2Response_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsV2Response.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsV2Response_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Response_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAlertsV2Response_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsV2Response_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: AlertV2_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListAlertsV2Response_SdkV2 as
// a slice of AlertV2_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsV2Response_SdkV2) GetResults(ctx context.Context) ([]AlertV2_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []AlertV2_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListAlertsV2Response_SdkV2.
func (o *ListAlertsV2Response_SdkV2) SetResults(ctx context.Context, v []AlertV2_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListDashboardsRequest_SdkV2 struct {
	// Name of dashboard attribute to order by.
	Order types.String `tfsdk:"-"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-"`
	// Number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// Full text search term.
	Q types.String `tfsdk:"-"`
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
			"order":     o.Order,
			"page":      o.Page,
			"page_size": o.PageSize,
			"q":         o.Q,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDashboardsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"order":     types.StringType,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"q":         types.StringType,
		},
	}
}

type ListQueriesLegacyRequest_SdkV2 struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"order":     types.StringType,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"q":         types.StringType,
		},
	}
}

type ListQueriesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueriesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueriesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListQueriesResponse_SdkV2 struct {
	// Whether there is another page of results.
	HasNextPage types.Bool `tfsdk:"has_next_page"`
	// A token that can be used to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Res types.List `tfsdk:"res"`
}

func (newState *ListQueriesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesResponse_SdkV2) {
}

func (newState *ListQueriesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListQueriesResponse_SdkV2) {
}

func (c ListQueriesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["has_next_page"] = attrs["has_next_page"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["res"] = attrs["res"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueriesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueriesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"res": reflect.TypeOf(QueryInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueriesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_next_page":   o.HasNextPage,
			"next_page_token": o.NextPageToken,
			"res":             o.Res,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueriesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_next_page":   types.BoolType,
			"next_page_token": types.StringType,
			"res": basetypes.ListType{
				ElemType: QueryInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRes returns the value of the Res field in ListQueriesResponse_SdkV2 as
// a slice of QueryInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueriesResponse_SdkV2) GetRes(ctx context.Context) ([]QueryInfo_SdkV2, bool) {
	if o.Res.IsNull() || o.Res.IsUnknown() {
		return nil, false
	}
	var v []QueryInfo_SdkV2
	d := o.Res.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRes sets the value of the Res field in ListQueriesResponse_SdkV2.
func (o *ListQueriesResponse_SdkV2) SetRes(ctx context.Context, v []QueryInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["res"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Res = types.ListValueMust(t, vs)
}

type ListQueryHistoryRequest_SdkV2 struct {
	// An optional filter object to limit query history results. Accepts
	// parameters such as user IDs, endpoint IDs, and statuses to narrow the
	// returned data. In a URL, the parameters of this filter are specified with
	// dot notation. For example: `filter_by.statement_ids`.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryHistoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryHistoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(QueryFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryHistoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueryHistoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListQueryHistoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter_by": basetypes.ListType{
				ElemType: QueryFilter_SdkV2{}.Type(ctx),
			},
			"include_metrics": types.BoolType,
			"max_results":     types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

// GetFilterBy returns the value of the FilterBy field in ListQueryHistoryRequest_SdkV2 as
// a QueryFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryHistoryRequest_SdkV2) GetFilterBy(ctx context.Context) (QueryFilter_SdkV2, bool) {
	var e QueryFilter_SdkV2
	if o.FilterBy.IsNull() || o.FilterBy.IsUnknown() {
		return e, false
	}
	var v []QueryFilter_SdkV2
	d := o.FilterBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListQueryHistoryRequest_SdkV2.
func (o *ListQueryHistoryRequest_SdkV2) SetFilterBy(ctx context.Context, v QueryFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filter_by"]
	o.FilterBy = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (newState *ListQueryObjectsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponse_SdkV2) {
}

func (newState *ListQueryObjectsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponse_SdkV2) {
}

func (c ListQueryObjectsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryObjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryObjectsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListQueryObjectsResponseQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueryObjectsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQueryObjectsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListQueryObjectsResponseQuery_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListQueryObjectsResponse_SdkV2 as
// a slice of ListQueryObjectsResponseQuery_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponse_SdkV2) GetResults(ctx context.Context) ([]ListQueryObjectsResponseQuery_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListQueryObjectsResponseQuery_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListQueryObjectsResponse_SdkV2.
func (o *ListQueryObjectsResponse_SdkV2) SetResults(ctx context.Context, v []ListQueryObjectsResponseQuery_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponseQuery_SdkV2 struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema"`

	Tags types.List `tfsdk:"tags"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *ListQueryObjectsResponseQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponseQuery_SdkV2) {
}

func (newState *ListQueryObjectsResponseQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponseQuery_SdkV2) {
}

func (c ListQueryObjectsResponseQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_auto_limit"] = attrs["apply_auto_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_modifier_user_name"] = attrs["last_modifier_user_name"].SetOptional()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as_mode"] = attrs["run_as_mode"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQueryObjectsResponseQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQueryObjectsResponseQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponseQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQueryObjectsResponseQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListQueryObjectsResponseQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: QueryParameter_SdkV2{}.Type(ctx),
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

// GetParameters returns the value of the Parameters field in ListQueryObjectsResponseQuery_SdkV2 as
// a slice of QueryParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponseQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ListQueryObjectsResponseQuery_SdkV2.
func (o *ListQueryObjectsResponseQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListQueryObjectsResponseQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQueryObjectsResponseQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in ListQueryObjectsResponseQuery_SdkV2.
func (o *ListQueryObjectsResponseQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ListRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListResponse_SdkV2 struct {
	// The total number of dashboards.
	Count types.Int64 `tfsdk:"count"`
	// The current page being displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of dashboards per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of dashboards returned.
	Results types.List `tfsdk:"results"`
}

func (newState *ListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListResponse_SdkV2) {
}

func (newState *ListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListResponse_SdkV2) {
}

func (c ListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["count"] = attrs["count"].SetOptional()
	attrs["page"] = attrs["page"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Dashboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":     types.Int64Type,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"results": basetypes.ListType{
				ElemType: Dashboard_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListResponse_SdkV2 as
// a slice of Dashboard_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListResponse_SdkV2) GetResults(ctx context.Context) ([]Dashboard_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []Dashboard_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListResponse_SdkV2.
func (o *ListResponse_SdkV2) SetResults(ctx context.Context, v []Dashboard_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListVisualizationsForQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVisualizationsForQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVisualizationsForQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVisualizationsForQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVisualizationsForQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListVisualizationsForQueryResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (newState *ListVisualizationsForQueryResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVisualizationsForQueryResponse_SdkV2) {
}

func (newState *ListVisualizationsForQueryResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListVisualizationsForQueryResponse_SdkV2) {
}

func (c ListVisualizationsForQueryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVisualizationsForQueryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVisualizationsForQueryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Visualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVisualizationsForQueryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVisualizationsForQueryResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: Visualization_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListVisualizationsForQueryResponse_SdkV2 as
// a slice of Visualization_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListVisualizationsForQueryResponse_SdkV2) GetResults(ctx context.Context) ([]Visualization_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []Visualization_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListVisualizationsForQueryResponse_SdkV2.
func (o *ListVisualizationsForQueryResponse_SdkV2) SetResults(ctx context.Context, v []Visualization_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListWarehousesRequest_SdkV2 struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWarehousesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWarehousesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWarehousesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_as_user_id": o.RunAsUserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWarehousesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_as_user_id": types.Int64Type,
		},
	}
}

type ListWarehousesResponse_SdkV2 struct {
	// A list of warehouses and their configurations.
	Warehouses types.List `tfsdk:"warehouses"`
}

func (newState *ListWarehousesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWarehousesResponse_SdkV2) {
}

func (newState *ListWarehousesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListWarehousesResponse_SdkV2) {
}

func (c ListWarehousesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouses"] = attrs["warehouses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWarehousesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWarehousesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"warehouses": reflect.TypeOf(EndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWarehousesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouses": o.Warehouses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWarehousesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouses": basetypes.ListType{
				ElemType: EndpointInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWarehouses returns the value of the Warehouses field in ListWarehousesResponse_SdkV2 as
// a slice of EndpointInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWarehousesResponse_SdkV2) GetWarehouses(ctx context.Context) ([]EndpointInfo_SdkV2, bool) {
	if o.Warehouses.IsNull() || o.Warehouses.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo_SdkV2
	d := o.Warehouses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarehouses sets the value of the Warehouses field in ListWarehousesResponse_SdkV2.
func (o *ListWarehousesResponse_SdkV2) SetWarehouses(ctx context.Context, v []EndpointInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Warehouses = types.ListValueMust(t, vs)
}

type MultiValuesOptions_SdkV2 struct {
	// Character that prefixes each selected parameter value.
	Prefix types.String `tfsdk:"prefix"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator types.String `tfsdk:"separator"`
	// Character that suffixes each selected parameter value.
	Suffix types.String `tfsdk:"suffix"`
}

func (newState *MultiValuesOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MultiValuesOptions_SdkV2) {
}

func (newState *MultiValuesOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState MultiValuesOptions_SdkV2) {
}

func (c MultiValuesOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["prefix"] = attrs["prefix"].SetOptional()
	attrs["separator"] = attrs["separator"].SetOptional()
	attrs["suffix"] = attrs["suffix"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MultiValuesOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MultiValuesOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MultiValuesOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o MultiValuesOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"prefix":    o.Prefix,
			"separator": o.Separator,
			"suffix":    o.Suffix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MultiValuesOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"prefix":    types.StringType,
			"separator": types.StringType,
			"suffix":    types.StringType,
		},
	}
}

type NumericValue_SdkV2 struct {
	Value types.Float64 `tfsdk:"value"`
}

func (newState *NumericValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NumericValue_SdkV2) {
}

func (newState *NumericValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState NumericValue_SdkV2) {
}

func (c NumericValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NumericValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NumericValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NumericValue_SdkV2
// only implements ToObjectValue() and Type().
func (o NumericValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NumericValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Float64Type,
		},
	}
}

type OdbcParams_SdkV2 struct {
	Hostname types.String `tfsdk:"hostname"`

	Path types.String `tfsdk:"path"`

	Port types.Int64 `tfsdk:"port"`

	Protocol types.String `tfsdk:"protocol"`
}

func (newState *OdbcParams_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OdbcParams_SdkV2) {
}

func (newState *OdbcParams_SdkV2) SyncEffectiveFieldsDuringRead(existingState OdbcParams_SdkV2) {
}

func (c OdbcParams_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["hostname"] = attrs["hostname"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["port"] = attrs["port"].SetOptional()
	attrs["protocol"] = attrs["protocol"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OdbcParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OdbcParams_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OdbcParams_SdkV2
// only implements ToObjectValue() and Type().
func (o OdbcParams_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OdbcParams_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hostname": types.StringType,
			"path":     types.StringType,
			"port":     types.Int64Type,
			"protocol": types.StringType,
		},
	}
}

type Parameter_SdkV2 struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions types.String `tfsdk:"enumOptions"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions types.List `tfsdk:"multiValuesOptions"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name types.String `tfsdk:"name"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId types.String `tfsdk:"queryId"`
	// The text displayed in a parameter picking widget.
	Title types.String `tfsdk:"title"`
	// Parameters can have several different types.
	Type_ types.String `tfsdk:"type"`
	// The default value for this parameter.
	Value types.Object `tfsdk:"value"`
}

func (newState *Parameter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Parameter_SdkV2) {
}

func (newState *Parameter_SdkV2) SyncEffectiveFieldsDuringRead(existingState Parameter_SdkV2) {
}

func (c Parameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enumOptions"] = attrs["enumOptions"].SetOptional()
	attrs["multiValuesOptions"] = attrs["multiValuesOptions"].SetOptional()
	attrs["multiValuesOptions"] = attrs["multiValuesOptions"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["queryId"] = attrs["queryId"].SetOptional()
	attrs["title"] = attrs["title"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Parameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Parameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multiValuesOptions": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Parameter_SdkV2
// only implements ToObjectValue() and Type().
func (o Parameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Parameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enumOptions": types.StringType,
			"multiValuesOptions": basetypes.ListType{
				ElemType: MultiValuesOptions_SdkV2{}.Type(ctx),
			},
			"name":    types.StringType,
			"queryId": types.StringType,
			"title":   types.StringType,
			"type":    types.StringType,
			"value":   types.ObjectType{},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in Parameter_SdkV2 as
// a MultiValuesOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Parameter_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in Parameter_SdkV2.
func (o *Parameter_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multiValuesOptions"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

type Query_SdkV2 struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema"`

	Tags types.List `tfsdk:"tags"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *Query_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Query_SdkV2) {
}

func (newState *Query_SdkV2) SyncEffectiveFieldsDuringRead(existingState Query_SdkV2) {
}

func (c Query_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_auto_limit"] = attrs["apply_auto_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_modifier_user_name"] = attrs["last_modifier_user_name"].SetOptional()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as_mode"] = attrs["run_as_mode"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Query.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Query_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Query_SdkV2
// only implements ToObjectValue() and Type().
func (o Query_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Query_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: QueryParameter_SdkV2{}.Type(ctx),
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

// GetParameters returns the value of the Parameters field in Query_SdkV2 as
// a slice of QueryParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Query_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in Query_SdkV2.
func (o *Query_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Query_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Query_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in Query_SdkV2.
func (o *Query_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type QueryBackedValue_SdkV2 struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options"`
	// UUID of the query that provides the parameter values.
	QueryId types.String `tfsdk:"query_id"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (newState *QueryBackedValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryBackedValue_SdkV2) {
}

func (newState *QueryBackedValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryBackedValue_SdkV2) {
}

func (c QueryBackedValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryBackedValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryBackedValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryBackedValue_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryBackedValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"multi_values_options": o.MultiValuesOptions,
			"query_id":             o.QueryId,
			"values":               o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryBackedValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"multi_values_options": basetypes.ListType{
				ElemType: MultiValuesOptions_SdkV2{}.Type(ctx),
			},
			"query_id": types.StringType,
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in QueryBackedValue_SdkV2 as
// a MultiValuesOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryBackedValue_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if o.MultiValuesOptions.IsNull() || o.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := o.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in QueryBackedValue_SdkV2.
func (o *QueryBackedValue_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	o.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in QueryBackedValue_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryBackedValue_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
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

// SetValues sets the value of the Values field in QueryBackedValue_SdkV2.
func (o *QueryBackedValue_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type QueryEditContent_SdkV2 struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options types.Object `tfsdk:"options"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`

	QueryId types.String `tfsdk:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEditContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryEditContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEditContent_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryEditContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryEditContent_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetTags returns the value of the Tags field in QueryEditContent_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEditContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in QueryEditContent_SdkV2.
func (o *QueryEditContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type QueryFilter_SdkV2 struct {
	// A range filter for query submitted time. The time range must be less than
	// or equal to 30 days.
	QueryStartTimeRange types.List `tfsdk:"query_start_time_range"`
	// A list of statement IDs.
	StatementIds types.List `tfsdk:"statement_ids"`
	// A list of statuses (QUEUED, RUNNING, CANCELED, FAILED, FINISHED) to match
	// query results. Corresponds to the `status` field in the response.
	// Filtering for multiple statuses is not recommended. Instead, opt to
	// filter by a single status multiple times and then combine the results.
	Statuses types.List `tfsdk:"statuses"`
	// A list of user IDs who ran the queries.
	UserIds types.List `tfsdk:"user_ids"`
	// A list of warehouse IDs.
	WarehouseIds types.List `tfsdk:"warehouse_ids"`
}

func (newState *QueryFilter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryFilter_SdkV2) {
}

func (newState *QueryFilter_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryFilter_SdkV2) {
}

func (c QueryFilter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_start_time_range"] = attrs["query_start_time_range"].SetOptional()
	attrs["query_start_time_range"] = attrs["query_start_time_range"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["statement_ids"] = attrs["statement_ids"].SetOptional()
	attrs["statuses"] = attrs["statuses"].SetOptional()
	attrs["user_ids"] = attrs["user_ids"].SetOptional()
	attrs["warehouse_ids"] = attrs["warehouse_ids"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_start_time_range": reflect.TypeOf(TimeRange_SdkV2{}),
		"statement_ids":          reflect.TypeOf(types.String{}),
		"statuses":               reflect.TypeOf(types.String{}),
		"user_ids":               reflect.TypeOf(types.Int64{}),
		"warehouse_ids":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryFilter_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryFilter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_start_time_range": basetypes.ListType{
				ElemType: TimeRange_SdkV2{}.Type(ctx),
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

// GetQueryStartTimeRange returns the value of the QueryStartTimeRange field in QueryFilter_SdkV2 as
// a TimeRange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter_SdkV2) GetQueryStartTimeRange(ctx context.Context) (TimeRange_SdkV2, bool) {
	var e TimeRange_SdkV2
	if o.QueryStartTimeRange.IsNull() || o.QueryStartTimeRange.IsUnknown() {
		return e, false
	}
	var v []TimeRange_SdkV2
	d := o.QueryStartTimeRange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryStartTimeRange sets the value of the QueryStartTimeRange field in QueryFilter_SdkV2.
func (o *QueryFilter_SdkV2) SetQueryStartTimeRange(ctx context.Context, v TimeRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_start_time_range"]
	o.QueryStartTimeRange = types.ListValueMust(t, vs)
}

// GetStatementIds returns the value of the StatementIds field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter_SdkV2) GetStatementIds(ctx context.Context) ([]types.String, bool) {
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

// SetStatementIds sets the value of the StatementIds field in QueryFilter_SdkV2.
func (o *QueryFilter_SdkV2) SetStatementIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.StatementIds = types.ListValueMust(t, vs)
}

// GetStatuses returns the value of the Statuses field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter_SdkV2) GetStatuses(ctx context.Context) ([]types.String, bool) {
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

// SetStatuses sets the value of the Statuses field in QueryFilter_SdkV2.
func (o *QueryFilter_SdkV2) SetStatuses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Statuses = types.ListValueMust(t, vs)
}

// GetUserIds returns the value of the UserIds field in QueryFilter_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter_SdkV2) GetUserIds(ctx context.Context) ([]types.Int64, bool) {
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

// SetUserIds sets the value of the UserIds field in QueryFilter_SdkV2.
func (o *QueryFilter_SdkV2) SetUserIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UserIds = types.ListValueMust(t, vs)
}

// GetWarehouseIds returns the value of the WarehouseIds field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryFilter_SdkV2) GetWarehouseIds(ctx context.Context) ([]types.String, bool) {
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

// SetWarehouseIds sets the value of the WarehouseIds field in QueryFilter_SdkV2.
func (o *QueryFilter_SdkV2) SetWarehouseIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouse_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarehouseIds = types.ListValueMust(t, vs)
}

type QueryInfo_SdkV2 struct {
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed types.List `tfsdk:"channel_used"`
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	ClientApplication types.String `tfsdk:"client_application"`
	// Total time of the statement execution. This value does not include the
	// time taken to retrieve the results, which can result in a discrepancy
	// between this value and the start-to-finish wall-clock time.
	Duration types.Int64 `tfsdk:"duration"`
	// Alias for `warehouse_id`.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// Message describing why the query could not complete.
	ErrorMessage types.String `tfsdk:"error_message"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId types.Int64 `tfsdk:"executed_as_user_id"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName types.String `tfsdk:"executed_as_user_name"`
	// The time execution of the query ended.
	ExecutionEndTimeMs types.Int64 `tfsdk:"execution_end_time_ms"`
	// Whether more updates for the query are expected.
	IsFinal types.Bool `tfsdk:"is_final"`
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key"`
	// Metrics about query execution.
	Metrics types.List `tfsdk:"metrics"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState types.String `tfsdk:"plans_state"`
	// The time the query ended.
	QueryEndTimeMs types.Int64 `tfsdk:"query_end_time_ms"`
	// The query ID.
	QueryId types.String `tfsdk:"query_id"`
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	QuerySource types.List `tfsdk:"query_source"`
	// The time the query started.
	QueryStartTimeMs types.Int64 `tfsdk:"query_start_time_ms"`
	// The text of the query.
	QueryText types.String `tfsdk:"query_text"`
	// The number of results returned by the query.
	RowsProduced types.Int64 `tfsdk:"rows_produced"`
	// URL to the Spark UI query plan.
	SparkUiUrl types.String `tfsdk:"spark_ui_url"`
	// Type of statement for this query
	StatementType types.String `tfsdk:"statement_type"`
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status types.String `tfsdk:"status"`
	// The ID of the user who ran the query.
	UserId types.Int64 `tfsdk:"user_id"`
	// The email address or username of the user who ran the query.
	UserName types.String `tfsdk:"user_name"`
	// Warehouse ID.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *QueryInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryInfo_SdkV2) {
}

func (newState *QueryInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryInfo_SdkV2) {
}

func (c QueryInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel_used"] = attrs["channel_used"].SetOptional()
	attrs["channel_used"] = attrs["channel_used"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["client_application"] = attrs["client_application"].SetOptional()
	attrs["duration"] = attrs["duration"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["executed_as_user_id"] = attrs["executed_as_user_id"].SetOptional()
	attrs["executed_as_user_name"] = attrs["executed_as_user_name"].SetOptional()
	attrs["execution_end_time_ms"] = attrs["execution_end_time_ms"].SetOptional()
	attrs["is_final"] = attrs["is_final"].SetOptional()
	attrs["lookup_key"] = attrs["lookup_key"].SetOptional()
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["metrics"] = attrs["metrics"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["plans_state"] = attrs["plans_state"].SetOptional()
	attrs["query_end_time_ms"] = attrs["query_end_time_ms"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["query_source"] = attrs["query_source"].SetOptional()
	attrs["query_source"] = attrs["query_source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["query_start_time_ms"] = attrs["query_start_time_ms"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["rows_produced"] = attrs["rows_produced"].SetOptional()
	attrs["spark_ui_url"] = attrs["spark_ui_url"].SetOptional()
	attrs["statement_type"] = attrs["statement_type"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel_used": reflect.TypeOf(ChannelInfo_SdkV2{}),
		"metrics":      reflect.TypeOf(QueryMetrics_SdkV2{}),
		"query_source": reflect.TypeOf(ExternalQuerySource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel_used":          o.ChannelUsed,
			"client_application":    o.ClientApplication,
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
			"query_source":          o.QuerySource,
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
func (o QueryInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel_used": basetypes.ListType{
				ElemType: ChannelInfo_SdkV2{}.Type(ctx),
			},
			"client_application":    types.StringType,
			"duration":              types.Int64Type,
			"endpoint_id":           types.StringType,
			"error_message":         types.StringType,
			"executed_as_user_id":   types.Int64Type,
			"executed_as_user_name": types.StringType,
			"execution_end_time_ms": types.Int64Type,
			"is_final":              types.BoolType,
			"lookup_key":            types.StringType,
			"metrics": basetypes.ListType{
				ElemType: QueryMetrics_SdkV2{}.Type(ctx),
			},
			"plans_state":       types.StringType,
			"query_end_time_ms": types.Int64Type,
			"query_id":          types.StringType,
			"query_source": basetypes.ListType{
				ElemType: ExternalQuerySource_SdkV2{}.Type(ctx),
			},
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

// GetChannelUsed returns the value of the ChannelUsed field in QueryInfo_SdkV2 as
// a ChannelInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo_SdkV2) GetChannelUsed(ctx context.Context) (ChannelInfo_SdkV2, bool) {
	var e ChannelInfo_SdkV2
	if o.ChannelUsed.IsNull() || o.ChannelUsed.IsUnknown() {
		return e, false
	}
	var v []ChannelInfo_SdkV2
	d := o.ChannelUsed.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannelUsed sets the value of the ChannelUsed field in QueryInfo_SdkV2.
func (o *QueryInfo_SdkV2) SetChannelUsed(ctx context.Context, v ChannelInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel_used"]
	o.ChannelUsed = types.ListValueMust(t, vs)
}

// GetMetrics returns the value of the Metrics field in QueryInfo_SdkV2 as
// a QueryMetrics_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo_SdkV2) GetMetrics(ctx context.Context) (QueryMetrics_SdkV2, bool) {
	var e QueryMetrics_SdkV2
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return e, false
	}
	var v []QueryMetrics_SdkV2
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetrics sets the value of the Metrics field in QueryInfo_SdkV2.
func (o *QueryInfo_SdkV2) SetMetrics(ctx context.Context, v QueryMetrics_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	o.Metrics = types.ListValueMust(t, vs)
}

// GetQuerySource returns the value of the QuerySource field in QueryInfo_SdkV2 as
// a ExternalQuerySource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo_SdkV2) GetQuerySource(ctx context.Context) (ExternalQuerySource_SdkV2, bool) {
	var e ExternalQuerySource_SdkV2
	if o.QuerySource.IsNull() || o.QuerySource.IsUnknown() {
		return e, false
	}
	var v []ExternalQuerySource_SdkV2
	d := o.QuerySource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuerySource sets the value of the QuerySource field in QueryInfo_SdkV2.
func (o *QueryInfo_SdkV2) SetQuerySource(ctx context.Context, v ExternalQuerySource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_source"]
	o.QuerySource = types.ListValueMust(t, vs)
}

type QueryList_SdkV2 struct {
	// The total number of queries.
	Count types.Int64 `tfsdk:"count"`
	// The page number that is currently displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of queries per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of queries returned.
	Results types.List `tfsdk:"results"`
}

func (newState *QueryList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryList_SdkV2) {
}

func (newState *QueryList_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryList_SdkV2) {
}

func (c QueryList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["count"] = attrs["count"].SetOptional()
	attrs["page"] = attrs["page"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(LegacyQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryList_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":     types.Int64Type,
			"page":      types.Int64Type,
			"page_size": types.Int64Type,
			"results": basetypes.ListType{
				ElemType: LegacyQuery_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in QueryList_SdkV2 as
// a slice of LegacyQuery_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryList_SdkV2) GetResults(ctx context.Context) ([]LegacyQuery_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []LegacyQuery_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in QueryList_SdkV2.
func (o *QueryList_SdkV2) SetResults(ctx context.Context, v []LegacyQuery_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics_SdkV2 struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs types.Int64 `tfsdk:"compilation_time_ms"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs types.Int64 `tfsdk:"execution_time_ms"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes types.Int64 `tfsdk:"network_sent_bytes"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp types.Int64 `tfsdk:"overloading_queue_start_timestamp"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs types.Int64 `tfsdk:"photon_total_time_ms"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp types.Int64 `tfsdk:"provisioning_queue_start_timestamp"`
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes types.Int64 `tfsdk:"pruned_bytes"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount types.Int64 `tfsdk:"pruned_files_count"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp types.Int64 `tfsdk:"query_compilation_start_timestamp"`
	// Total size of data read by the query, in bytes.
	ReadBytes types.Int64 `tfsdk:"read_bytes"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes types.Int64 `tfsdk:"read_cache_bytes"`
	// Number of files read after pruning
	ReadFilesCount types.Int64 `tfsdk:"read_files_count"`
	// Number of partitions read after pruning.
	ReadPartitionsCount types.Int64 `tfsdk:"read_partitions_count"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes types.Int64 `tfsdk:"read_remote_bytes"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs types.Int64 `tfsdk:"result_fetch_time_ms"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache types.Bool `tfsdk:"result_from_cache"`
	// Total number of rows returned by the query.
	RowsProducedCount types.Int64 `tfsdk:"rows_produced_count"`
	// Total number of rows read by the query.
	RowsReadCount types.Int64 `tfsdk:"rows_read_count"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes types.Int64 `tfsdk:"spill_to_disk_bytes"`
	// sum of task times completed in a range of wall clock time, approximated
	// to a configurable number of points aggregated over all stages and jobs in
	// the query (based on task_total_time_ms)
	TaskTimeOverTimeRange types.List `tfsdk:"task_time_over_time_range"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs types.Int64 `tfsdk:"task_total_time_ms"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs types.Int64 `tfsdk:"total_time_ms"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes types.Int64 `tfsdk:"write_remote_bytes"`
}

func (newState *QueryMetrics_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryMetrics_SdkV2) {
}

func (newState *QueryMetrics_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryMetrics_SdkV2) {
}

func (c QueryMetrics_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compilation_time_ms"] = attrs["compilation_time_ms"].SetOptional()
	attrs["execution_time_ms"] = attrs["execution_time_ms"].SetOptional()
	attrs["network_sent_bytes"] = attrs["network_sent_bytes"].SetOptional()
	attrs["overloading_queue_start_timestamp"] = attrs["overloading_queue_start_timestamp"].SetOptional()
	attrs["photon_total_time_ms"] = attrs["photon_total_time_ms"].SetOptional()
	attrs["provisioning_queue_start_timestamp"] = attrs["provisioning_queue_start_timestamp"].SetOptional()
	attrs["pruned_bytes"] = attrs["pruned_bytes"].SetOptional()
	attrs["pruned_files_count"] = attrs["pruned_files_count"].SetOptional()
	attrs["query_compilation_start_timestamp"] = attrs["query_compilation_start_timestamp"].SetOptional()
	attrs["read_bytes"] = attrs["read_bytes"].SetOptional()
	attrs["read_cache_bytes"] = attrs["read_cache_bytes"].SetOptional()
	attrs["read_files_count"] = attrs["read_files_count"].SetOptional()
	attrs["read_partitions_count"] = attrs["read_partitions_count"].SetOptional()
	attrs["read_remote_bytes"] = attrs["read_remote_bytes"].SetOptional()
	attrs["result_fetch_time_ms"] = attrs["result_fetch_time_ms"].SetOptional()
	attrs["result_from_cache"] = attrs["result_from_cache"].SetOptional()
	attrs["rows_produced_count"] = attrs["rows_produced_count"].SetOptional()
	attrs["rows_read_count"] = attrs["rows_read_count"].SetOptional()
	attrs["spill_to_disk_bytes"] = attrs["spill_to_disk_bytes"].SetOptional()
	attrs["task_time_over_time_range"] = attrs["task_time_over_time_range"].SetOptional()
	attrs["task_time_over_time_range"] = attrs["task_time_over_time_range"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["task_total_time_ms"] = attrs["task_total_time_ms"].SetOptional()
	attrs["total_time_ms"] = attrs["total_time_ms"].SetOptional()
	attrs["write_remote_bytes"] = attrs["write_remote_bytes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryMetrics.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryMetrics_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"task_time_over_time_range": reflect.TypeOf(TaskTimeOverRange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryMetrics_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryMetrics_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"task_time_over_time_range":          o.TaskTimeOverTimeRange,
			"task_total_time_ms":                 o.TaskTotalTimeMs,
			"total_time_ms":                      o.TotalTimeMs,
			"write_remote_bytes":                 o.WriteRemoteBytes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryMetrics_SdkV2) Type(ctx context.Context) attr.Type {
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
			"task_time_over_time_range": basetypes.ListType{
				ElemType: TaskTimeOverRange_SdkV2{}.Type(ctx),
			},
			"task_total_time_ms": types.Int64Type,
			"total_time_ms":      types.Int64Type,
			"write_remote_bytes": types.Int64Type,
		},
	}
}

// GetTaskTimeOverTimeRange returns the value of the TaskTimeOverTimeRange field in QueryMetrics_SdkV2 as
// a TaskTimeOverRange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryMetrics_SdkV2) GetTaskTimeOverTimeRange(ctx context.Context) (TaskTimeOverRange_SdkV2, bool) {
	var e TaskTimeOverRange_SdkV2
	if o.TaskTimeOverTimeRange.IsNull() || o.TaskTimeOverTimeRange.IsUnknown() {
		return e, false
	}
	var v []TaskTimeOverRange_SdkV2
	d := o.TaskTimeOverTimeRange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTaskTimeOverTimeRange sets the value of the TaskTimeOverTimeRange field in QueryMetrics_SdkV2.
func (o *QueryMetrics_SdkV2) SetTaskTimeOverTimeRange(ctx context.Context, v TaskTimeOverRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task_time_over_time_range"]
	o.TaskTimeOverTimeRange = types.ListValueMust(t, vs)
}

type QueryOptions_SdkV2 struct {
	// The name of the catalog to execute this query in.
	Catalog types.String `tfsdk:"catalog"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`

	Parameters types.List `tfsdk:"parameters"`
	// The name of the schema to execute this query in.
	Schema types.String `tfsdk:"schema"`
}

func (newState *QueryOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryOptions_SdkV2) {
}

func (newState *QueryOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryOptions_SdkV2) {
}

func (c QueryOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["moved_to_trash_at"] = attrs["moved_to_trash_at"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(Parameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog":           types.StringType,
			"moved_to_trash_at": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: Parameter_SdkV2{}.Type(ctx),
			},
			"schema": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in QueryOptions_SdkV2 as
// a slice of Parameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryOptions_SdkV2) GetParameters(ctx context.Context) ([]Parameter_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []Parameter_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in QueryOptions_SdkV2.
func (o *QueryOptions_SdkV2) SetParameters(ctx context.Context, v []Parameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type QueryParameter_SdkV2 struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue types.List `tfsdk:"date_range_value"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue types.List `tfsdk:"date_value"`
	// Dropdown query parameter value.
	EnumValue types.List `tfsdk:"enum_value"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name types.String `tfsdk:"name"`
	// Numeric query parameter value.
	NumericValue types.List `tfsdk:"numeric_value"`
	// Query-based dropdown query parameter value.
	QueryBackedValue types.List `tfsdk:"query_backed_value"`
	// Text query parameter value.
	TextValue types.List `tfsdk:"text_value"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title types.String `tfsdk:"title"`
}

func (newState *QueryParameter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryParameter_SdkV2) {
}

func (newState *QueryParameter_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryParameter_SdkV2) {
}

func (c QueryParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["date_range_value"] = attrs["date_range_value"].SetOptional()
	attrs["date_range_value"] = attrs["date_range_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["date_value"] = attrs["date_value"].SetOptional()
	attrs["date_value"] = attrs["date_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["enum_value"] = attrs["enum_value"].SetOptional()
	attrs["enum_value"] = attrs["enum_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["numeric_value"] = attrs["numeric_value"].SetOptional()
	attrs["numeric_value"] = attrs["numeric_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["query_backed_value"] = attrs["query_backed_value"].SetOptional()
	attrs["query_backed_value"] = attrs["query_backed_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["text_value"] = attrs["text_value"].SetOptional()
	attrs["text_value"] = attrs["text_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["title"] = attrs["title"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"date_range_value":   reflect.TypeOf(DateRangeValue_SdkV2{}),
		"date_value":         reflect.TypeOf(DateValue_SdkV2{}),
		"enum_value":         reflect.TypeOf(EnumValue_SdkV2{}),
		"numeric_value":      reflect.TypeOf(NumericValue_SdkV2{}),
		"query_backed_value": reflect.TypeOf(QueryBackedValue_SdkV2{}),
		"text_value":         reflect.TypeOf(TextValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryParameter_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryParameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"date_range_value": basetypes.ListType{
				ElemType: DateRangeValue_SdkV2{}.Type(ctx),
			},
			"date_value": basetypes.ListType{
				ElemType: DateValue_SdkV2{}.Type(ctx),
			},
			"enum_value": basetypes.ListType{
				ElemType: EnumValue_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"numeric_value": basetypes.ListType{
				ElemType: NumericValue_SdkV2{}.Type(ctx),
			},
			"query_backed_value": basetypes.ListType{
				ElemType: QueryBackedValue_SdkV2{}.Type(ctx),
			},
			"text_value": basetypes.ListType{
				ElemType: TextValue_SdkV2{}.Type(ctx),
			},
			"title": types.StringType,
		},
	}
}

// GetDateRangeValue returns the value of the DateRangeValue field in QueryParameter_SdkV2 as
// a DateRangeValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetDateRangeValue(ctx context.Context) (DateRangeValue_SdkV2, bool) {
	var e DateRangeValue_SdkV2
	if o.DateRangeValue.IsNull() || o.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRangeValue_SdkV2
	d := o.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetDateRangeValue(ctx context.Context, v DateRangeValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	o.DateRangeValue = types.ListValueMust(t, vs)
}

// GetDateValue returns the value of the DateValue field in QueryParameter_SdkV2 as
// a DateValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetDateValue(ctx context.Context) (DateValue_SdkV2, bool) {
	var e DateValue_SdkV2
	if o.DateValue.IsNull() || o.DateValue.IsUnknown() {
		return e, false
	}
	var v []DateValue_SdkV2
	d := o.DateValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateValue sets the value of the DateValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetDateValue(ctx context.Context, v DateValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["date_value"]
	o.DateValue = types.ListValueMust(t, vs)
}

// GetEnumValue returns the value of the EnumValue field in QueryParameter_SdkV2 as
// a EnumValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetEnumValue(ctx context.Context) (EnumValue_SdkV2, bool) {
	var e EnumValue_SdkV2
	if o.EnumValue.IsNull() || o.EnumValue.IsUnknown() {
		return e, false
	}
	var v []EnumValue_SdkV2
	d := o.EnumValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnumValue sets the value of the EnumValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetEnumValue(ctx context.Context, v EnumValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enum_value"]
	o.EnumValue = types.ListValueMust(t, vs)
}

// GetNumericValue returns the value of the NumericValue field in QueryParameter_SdkV2 as
// a NumericValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetNumericValue(ctx context.Context) (NumericValue_SdkV2, bool) {
	var e NumericValue_SdkV2
	if o.NumericValue.IsNull() || o.NumericValue.IsUnknown() {
		return e, false
	}
	var v []NumericValue_SdkV2
	d := o.NumericValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNumericValue sets the value of the NumericValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetNumericValue(ctx context.Context, v NumericValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["numeric_value"]
	o.NumericValue = types.ListValueMust(t, vs)
}

// GetQueryBackedValue returns the value of the QueryBackedValue field in QueryParameter_SdkV2 as
// a QueryBackedValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetQueryBackedValue(ctx context.Context) (QueryBackedValue_SdkV2, bool) {
	var e QueryBackedValue_SdkV2
	if o.QueryBackedValue.IsNull() || o.QueryBackedValue.IsUnknown() {
		return e, false
	}
	var v []QueryBackedValue_SdkV2
	d := o.QueryBackedValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryBackedValue sets the value of the QueryBackedValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetQueryBackedValue(ctx context.Context, v QueryBackedValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_backed_value"]
	o.QueryBackedValue = types.ListValueMust(t, vs)
}

// GetTextValue returns the value of the TextValue field in QueryParameter_SdkV2 as
// a TextValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter_SdkV2) GetTextValue(ctx context.Context) (TextValue_SdkV2, bool) {
	var e TextValue_SdkV2
	if o.TextValue.IsNull() || o.TextValue.IsUnknown() {
		return e, false
	}
	var v []TextValue_SdkV2
	d := o.TextValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTextValue sets the value of the TextValue field in QueryParameter_SdkV2.
func (o *QueryParameter_SdkV2) SetTextValue(ctx context.Context, v TextValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["text_value"]
	o.TextValue = types.ListValueMust(t, vs)
}

type QueryPostContent_SdkV2 struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options types.Object `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryPostContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryPostContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryPostContent_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryPostContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QueryPostContent_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetTags returns the value of the Tags field in QueryPostContent_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryPostContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in QueryPostContent_SdkV2.
func (o *QueryPostContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type RepeatedEndpointConfPairs_SdkV2 struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs"`
}

func (newState *RepeatedEndpointConfPairs_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepeatedEndpointConfPairs_SdkV2) {
}

func (newState *RepeatedEndpointConfPairs_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepeatedEndpointConfPairs_SdkV2) {
}

func (c RepeatedEndpointConfPairs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config_pair"] = attrs["config_pair"].SetOptional()
	attrs["configuration_pairs"] = attrs["configuration_pairs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepeatedEndpointConfPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepeatedEndpointConfPairs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config_pair":         reflect.TypeOf(EndpointConfPair_SdkV2{}),
		"configuration_pairs": reflect.TypeOf(EndpointConfPair_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepeatedEndpointConfPairs_SdkV2
// only implements ToObjectValue() and Type().
func (o RepeatedEndpointConfPairs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_pair":         o.ConfigPair,
			"configuration_pairs": o.ConfigurationPairs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepeatedEndpointConfPairs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_pair": basetypes.ListType{
				ElemType: EndpointConfPair_SdkV2{}.Type(ctx),
			},
			"configuration_pairs": basetypes.ListType{
				ElemType: EndpointConfPair_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetConfigPair returns the value of the ConfigPair field in RepeatedEndpointConfPairs_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepeatedEndpointConfPairs_SdkV2) GetConfigPair(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if o.ConfigPair.IsNull() || o.ConfigPair.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := o.ConfigPair.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigPair sets the value of the ConfigPair field in RepeatedEndpointConfPairs_SdkV2.
func (o *RepeatedEndpointConfPairs_SdkV2) SetConfigPair(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_pair"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ConfigPair = types.ListValueMust(t, vs)
}

// GetConfigurationPairs returns the value of the ConfigurationPairs field in RepeatedEndpointConfPairs_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepeatedEndpointConfPairs_SdkV2) GetConfigurationPairs(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if o.ConfigurationPairs.IsNull() || o.ConfigurationPairs.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := o.ConfigurationPairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigurationPairs sets the value of the ConfigurationPairs field in RepeatedEndpointConfPairs_SdkV2.
func (o *RepeatedEndpointConfPairs_SdkV2) SetConfigurationPairs(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration_pairs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ConfigurationPairs = types.ListValueMust(t, vs)
}

type RestoreDashboardRequest_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": o.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type RestoreQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreQueriesLegacyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type RestoreResponse_SdkV2 struct {
}

func (newState *RestoreResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreResponse_SdkV2) {
}

func (newState *RestoreResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RestoreResponse_SdkV2) {
}

func (c RestoreResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResultData_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray types.List `tfsdk:"data_array"`

	ExternalLinks types.List `tfsdk:"external_links"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (newState *ResultData_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultData_SdkV2) {
}

func (newState *ResultData_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResultData_SdkV2) {
}

func (c ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["data_array"] = attrs["data_array"].SetOptional()
	attrs["external_links"] = attrs["external_links"].SetOptional()
	attrs["next_chunk_index"] = attrs["next_chunk_index"].SetOptional()
	attrs["next_chunk_internal_link"] = attrs["next_chunk_internal_link"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array":     reflect.TypeOf(types.String{}),
		"external_links": reflect.TypeOf(ExternalLink_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (o ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ResultData_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: ExternalLink_SdkV2{}.Type(ctx),
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.String, bool) {
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

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (o *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataArray = types.ListValueMust(t, vs)
}

// GetExternalLinks returns the value of the ExternalLinks field in ResultData_SdkV2 as
// a slice of ExternalLink_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultData_SdkV2) GetExternalLinks(ctx context.Context) ([]ExternalLink_SdkV2, bool) {
	if o.ExternalLinks.IsNull() || o.ExternalLinks.IsUnknown() {
		return nil, false
	}
	var v []ExternalLink_SdkV2
	d := o.ExternalLinks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLinks sets the value of the ExternalLinks field in ResultData_SdkV2.
func (o *ResultData_SdkV2) SetExternalLinks(ctx context.Context, v []ExternalLink_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_links"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExternalLinks = types.ListValueMust(t, vs)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest_SdkV2 struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks"`

	Format types.String `tfsdk:"format"`

	Schema types.List `tfsdk:"schema"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount types.Int64 `tfsdk:"total_byte_count"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount types.Int64 `tfsdk:"total_chunk_count"`
	// The total number of rows in the result set.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated types.Bool `tfsdk:"truncated"`
}

func (newState *ResultManifest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultManifest_SdkV2) {
}

func (newState *ResultManifest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResultManifest_SdkV2) {
}

func (c ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["chunks"] = attrs["chunks"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["schema"] = attrs["schema"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["total_byte_count"] = attrs["total_byte_count"].SetOptional()
	attrs["total_chunk_count"] = attrs["total_chunk_count"].SetOptional()
	attrs["total_row_count"] = attrs["total_row_count"].SetOptional()
	attrs["truncated"] = attrs["truncated"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"chunks": reflect.TypeOf(BaseChunkInfo_SdkV2{}),
		"schema": reflect.TypeOf(ResultSchema_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunks": basetypes.ListType{
				ElemType: BaseChunkInfo_SdkV2{}.Type(ctx),
			},
			"format": types.StringType,
			"schema": basetypes.ListType{
				ElemType: ResultSchema_SdkV2{}.Type(ctx),
			},
			"total_byte_count":  types.Int64Type,
			"total_chunk_count": types.Int64Type,
			"total_row_count":   types.Int64Type,
			"truncated":         types.BoolType,
		},
	}
}

// GetChunks returns the value of the Chunks field in ResultManifest_SdkV2 as
// a slice of BaseChunkInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultManifest_SdkV2) GetChunks(ctx context.Context) ([]BaseChunkInfo_SdkV2, bool) {
	if o.Chunks.IsNull() || o.Chunks.IsUnknown() {
		return nil, false
	}
	var v []BaseChunkInfo_SdkV2
	d := o.Chunks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChunks sets the value of the Chunks field in ResultManifest_SdkV2.
func (o *ResultManifest_SdkV2) SetChunks(ctx context.Context, v []BaseChunkInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["chunks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Chunks = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in ResultManifest_SdkV2 as
// a ResultSchema_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultManifest_SdkV2) GetSchema(ctx context.Context) (ResultSchema_SdkV2, bool) {
	var e ResultSchema_SdkV2
	if o.Schema.IsNull() || o.Schema.IsUnknown() {
		return e, false
	}
	var v []ResultSchema_SdkV2
	d := o.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in ResultManifest_SdkV2.
func (o *ResultManifest_SdkV2) SetSchema(ctx context.Context, v ResultSchema_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	o.Schema = types.ListValueMust(t, vs)
}

// The schema is an ordered list of column descriptions.
type ResultSchema_SdkV2 struct {
	ColumnCount types.Int64 `tfsdk:"column_count"`

	Columns types.List `tfsdk:"columns"`
}

func (newState *ResultSchema_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultSchema_SdkV2) {
}

func (newState *ResultSchema_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResultSchema_SdkV2) {
}

func (c ResultSchema_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_count"] = attrs["column_count"].SetOptional()
	attrs["columns"] = attrs["columns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultSchema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultSchema_SdkV2
// only implements ToObjectValue() and Type().
func (o ResultSchema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": o.ColumnCount,
			"columns":      o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultSchema_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in ResultSchema_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResultSchema_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultSchema_SdkV2.
func (o *ResultSchema_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type ServiceError_SdkV2 struct {
	ErrorCode types.String `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

func (newState *ServiceError_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServiceError_SdkV2) {
}

func (newState *ServiceError_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServiceError_SdkV2) {
}

func (c ServiceError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServiceError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServiceError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceError_SdkV2
// only implements ToObjectValue() and Type().
func (o ServiceError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_code": o.ErrorCode,
			"message":    o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServiceError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_code": types.StringType,
			"message":    types.StringType,
		},
	}
}

// Set object ACL
type SetRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permission to set.
	ObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"objectId":            o.ObjectId,
			"objectType":          o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl_SdkV2{}.Type(ctx),
			},
			"objectId":   types.StringType,
			"objectType": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetRequest_SdkV2 as
// a slice of AccessControl_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetRequest_SdkV2.
func (o *SetRequest_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetResponse_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *SetResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetResponse_SdkV2) {
}

func (newState *SetResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState SetResponse_SdkV2) {
}

func (c SetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControl_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetResponse_SdkV2 as
// a slice of AccessControl_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetResponse_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetResponse_SdkV2.
func (o *SetResponse_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigRequest_SdkV2 struct {
	// Optional: Channel selection details
	Channel types.List `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.List `tfsdk:"config_param"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.List `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetWorkspaceWarehouseConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel":                      reflect.TypeOf(Channel_SdkV2{}),
		"config_param":                 reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
		"data_access_config":           reflect.TypeOf(EndpointConfPair_SdkV2{}),
		"enabled_warehouse_types":      reflect.TypeOf(WarehouseTypePair_SdkV2{}),
		"global_param":                 reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
		"sql_configuration_parameters": reflect.TypeOf(RepeatedEndpointConfPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetWorkspaceWarehouseConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetWorkspaceWarehouseConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SetWorkspaceWarehouseConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel": basetypes.ListType{
				ElemType: Channel_SdkV2{}.Type(ctx),
			},
			"config_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair_SdkV2{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair_SdkV2{}.Type(ctx),
			},
			"global_param": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
			"google_service_account": types.StringType,
			"instance_profile_arn":   types.StringType,
			"security_policy":        types.StringType,
			"sql_configuration_parameters": basetypes.ListType{
				ElemType: RepeatedEndpointConfPairs_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetChannel returns the value of the Channel field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a Channel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if o.Channel.IsNull() || o.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := o.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	o.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	o.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if o.DataAccessConfig.IsNull() || o.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := o.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a slice of WarehouseTypePair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair_SdkV2, bool) {
	if o.EnabledWarehouseTypes.IsNull() || o.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair_SdkV2
	d := o.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.GlobalParam.IsNull() || o.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	o.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := o.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (o *SetWorkspaceWarehouseConfigRequest_SdkV2) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	o.SqlConfigurationParameters = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigResponse_SdkV2 struct {
}

func (newState *SetWorkspaceWarehouseConfigResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (newState *SetWorkspaceWarehouseConfigResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState SetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (c SetWorkspaceWarehouseConfigResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetWorkspaceWarehouseConfigResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetWorkspaceWarehouseConfigResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetWorkspaceWarehouseConfigResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetWorkspaceWarehouseConfigResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StartRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StartRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StartWarehouseResponse_SdkV2 struct {
}

func (newState *StartWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartWarehouseResponse_SdkV2) {
}

func (newState *StartWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState StartWarehouseResponse_SdkV2) {
}

func (c StartWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StartWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StartWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StatementParameterListItem_SdkV2 struct {
	// The name of a parameter marker to be substituted in the statement.
	Name types.String `tfsdk:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type_ types.String `tfsdk:"type"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value types.String `tfsdk:"value"`
}

func (newState *StatementParameterListItem_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementParameterListItem_SdkV2) {
}

func (newState *StatementParameterListItem_SdkV2) SyncEffectiveFieldsDuringRead(existingState StatementParameterListItem_SdkV2) {
}

func (c StatementParameterListItem_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementParameterListItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementParameterListItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementParameterListItem_SdkV2
// only implements ToObjectValue() and Type().
func (o StatementParameterListItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  o.Name,
			"type":  o.Type_,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StatementParameterListItem_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"type":  types.StringType,
			"value": types.StringType,
		},
	}
}

type StatementResponse_SdkV2 struct {
	Manifest types.List `tfsdk:"manifest"`

	Result types.List `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`

	Status types.List `tfsdk:"status"`
}

func (newState *StatementResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementResponse_SdkV2) {
}

func (newState *StatementResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState StatementResponse_SdkV2) {
}

func (c StatementResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["manifest"] = attrs["manifest"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["result"] = attrs["result"].SetOptional()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":   reflect.TypeOf(ResultData_SdkV2{}),
		"status":   reflect.TypeOf(StatementStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StatementResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StatementResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest": basetypes.ListType{
				ElemType: ResultManifest_SdkV2{}.Type(ctx),
			},
			"result": basetypes.ListType{
				ElemType: ResultData_SdkV2{}.Type(ctx),
			},
			"statement_id": types.StringType,
			"status": basetypes.ListType{
				ElemType: StatementStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetManifest returns the value of the Manifest field in StatementResponse_SdkV2 as
// a ResultManifest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if o.Manifest.IsNull() || o.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := o.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in StatementResponse_SdkV2.
func (o *StatementResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	o.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in StatementResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := o.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in StatementResponse_SdkV2.
func (o *StatementResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	o.Result = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in StatementResponse_SdkV2 as
// a StatementStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse_SdkV2) GetStatus(ctx context.Context) (StatementStatus_SdkV2, bool) {
	var e StatementStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []StatementStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in StatementResponse_SdkV2.
func (o *StatementResponse_SdkV2) SetStatus(ctx context.Context, v StatementStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus_SdkV2 struct {
	Error types.List `tfsdk:"error"`

	State types.String `tfsdk:"state"`
}

func (newState *StatementStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementStatus_SdkV2) {
}

func (newState *StatementStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState StatementStatus_SdkV2) {
}

func (c StatementStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StatementStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(ServiceError_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o StatementStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error": o.Error,
			"state": o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StatementStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": basetypes.ListType{
				ElemType: ServiceError_SdkV2{}.Type(ctx),
			},
			"state": types.StringType,
		},
	}
}

// GetError returns the value of the Error field in StatementStatus_SdkV2 as
// a ServiceError_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementStatus_SdkV2) GetError(ctx context.Context) (ServiceError_SdkV2, bool) {
	var e ServiceError_SdkV2
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []ServiceError_SdkV2
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in StatementStatus_SdkV2.
func (o *StatementStatus_SdkV2) SetError(ctx context.Context, v ServiceError_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

type StopRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StopRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StopWarehouseResponse_SdkV2 struct {
}

func (newState *StopWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopWarehouseResponse_SdkV2) {
}

func (newState *StopWarehouseResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState StopWarehouseResponse_SdkV2) {
}

func (c StopWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StopWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StopWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Success_SdkV2 struct {
	Message types.String `tfsdk:"message"`
}

func (newState *Success_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Success_SdkV2) {
}

func (newState *Success_SdkV2) SyncEffectiveFieldsDuringRead(existingState Success_SdkV2) {
}

func (c Success_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Success.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Success_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Success_SdkV2
// only implements ToObjectValue() and Type().
func (o Success_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Success_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
		},
	}
}

type TaskTimeOverRange_SdkV2 struct {
	Entries types.List `tfsdk:"entries"`
	// interval length for all entries (difference in start time and end time of
	// an entry range) the same for all entries start time of first interval is
	// query_start_time_ms
	Interval types.Int64 `tfsdk:"interval"`
}

func (newState *TaskTimeOverRange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskTimeOverRange_SdkV2) {
}

func (newState *TaskTimeOverRange_SdkV2) SyncEffectiveFieldsDuringRead(existingState TaskTimeOverRange_SdkV2) {
}

func (c TaskTimeOverRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entries"] = attrs["entries"].SetOptional()
	attrs["interval"] = attrs["interval"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskTimeOverRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskTimeOverRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entries": reflect.TypeOf(TaskTimeOverRangeEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRange_SdkV2
// only implements ToObjectValue() and Type().
func (o TaskTimeOverRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entries":  o.Entries,
			"interval": o.Interval,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskTimeOverRange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entries": basetypes.ListType{
				ElemType: TaskTimeOverRangeEntry_SdkV2{}.Type(ctx),
			},
			"interval": types.Int64Type,
		},
	}
}

// GetEntries returns the value of the Entries field in TaskTimeOverRange_SdkV2 as
// a slice of TaskTimeOverRangeEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskTimeOverRange_SdkV2) GetEntries(ctx context.Context) ([]TaskTimeOverRangeEntry_SdkV2, bool) {
	if o.Entries.IsNull() || o.Entries.IsUnknown() {
		return nil, false
	}
	var v []TaskTimeOverRangeEntry_SdkV2
	d := o.Entries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntries sets the value of the Entries field in TaskTimeOverRange_SdkV2.
func (o *TaskTimeOverRange_SdkV2) SetEntries(ctx context.Context, v []TaskTimeOverRangeEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entries = types.ListValueMust(t, vs)
}

type TaskTimeOverRangeEntry_SdkV2 struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	TaskCompletedTimeMs types.Int64 `tfsdk:"task_completed_time_ms"`
}

func (newState *TaskTimeOverRangeEntry_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskTimeOverRangeEntry_SdkV2) {
}

func (newState *TaskTimeOverRangeEntry_SdkV2) SyncEffectiveFieldsDuringRead(existingState TaskTimeOverRangeEntry_SdkV2) {
}

func (c TaskTimeOverRangeEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["task_completed_time_ms"] = attrs["task_completed_time_ms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskTimeOverRangeEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskTimeOverRangeEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRangeEntry_SdkV2
// only implements ToObjectValue() and Type().
func (o TaskTimeOverRangeEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"task_completed_time_ms": o.TaskCompletedTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskTimeOverRangeEntry_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"task_completed_time_ms": types.Int64Type,
		},
	}
}

type TerminationReason_SdkV2 struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters"`
	// type of the termination
	Type_ types.String `tfsdk:"type"`
}

func (newState *TerminationReason_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationReason_SdkV2) {
}

func (newState *TerminationReason_SdkV2) SyncEffectiveFieldsDuringRead(existingState TerminationReason_SdkV2) {
}

func (c TerminationReason_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["code"] = attrs["code"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TerminationReason.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TerminationReason_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason_SdkV2
// only implements ToObjectValue() and Type().
func (o TerminationReason_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       o.Code,
			"parameters": o.Parameters,
			"type":       o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TerminationReason_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetParameters returns the value of the Parameters field in TerminationReason_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TerminationReason_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
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

// SetParameters sets the value of the Parameters field in TerminationReason_SdkV2.
func (o *TerminationReason_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type TextValue_SdkV2 struct {
	Value types.String `tfsdk:"value"`
}

func (newState *TextValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextValue_SdkV2) {
}

func (newState *TextValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState TextValue_SdkV2) {
}

func (c TextValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TextValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TextValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TextValue_SdkV2
// only implements ToObjectValue() and Type().
func (o TextValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TextValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TimeRange_SdkV2 struct {
	// The end time in milliseconds.
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// The start time in milliseconds.
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
}

func (newState *TimeRange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TimeRange_SdkV2) {
}

func (newState *TimeRange_SdkV2) SyncEffectiveFieldsDuringRead(existingState TimeRange_SdkV2) {
}

func (c TimeRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time_ms"] = attrs["end_time_ms"].SetOptional()
	attrs["start_time_ms"] = attrs["start_time_ms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TimeRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TimeRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeRange_SdkV2
// only implements ToObjectValue() and Type().
func (o TimeRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   o.EndTimeMs,
			"start_time_ms": o.StartTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TimeRange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time_ms":   types.Int64Type,
			"start_time_ms": types.Int64Type,
		},
	}
}

type TransferOwnershipObjectId_SdkV2 struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner"`
}

func (newState *TransferOwnershipObjectId_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransferOwnershipObjectId_SdkV2) {
}

func (newState *TransferOwnershipObjectId_SdkV2) SyncEffectiveFieldsDuringRead(existingState TransferOwnershipObjectId_SdkV2) {
}

func (c TransferOwnershipObjectId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["new_owner"] = attrs["new_owner"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransferOwnershipObjectId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TransferOwnershipObjectId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipObjectId_SdkV2
// only implements ToObjectValue() and Type().
func (o TransferOwnershipObjectId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner": o.NewOwner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransferOwnershipObjectId_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner": types.StringType,
		},
	}
}

type TransferOwnershipRequest_SdkV2 struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner"`
	// The ID of the object on which to change ownership.
	ObjectId types.List `tfsdk:"-"`
	// The type of object on which to change ownership.
	ObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransferOwnershipRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TransferOwnershipRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objectId": reflect.TypeOf(TransferOwnershipObjectId_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TransferOwnershipRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner":  o.NewOwner,
			"objectId":   o.ObjectId,
			"objectType": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransferOwnershipRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner": types.StringType,
			"objectId": basetypes.ListType{
				ElemType: TransferOwnershipObjectId_SdkV2{}.Type(ctx),
			},
			"objectType": types.StringType,
		},
	}
}

// GetObjectId returns the value of the ObjectId field in TransferOwnershipRequest_SdkV2 as
// a TransferOwnershipObjectId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TransferOwnershipRequest_SdkV2) GetObjectId(ctx context.Context) (TransferOwnershipObjectId_SdkV2, bool) {
	var e TransferOwnershipObjectId_SdkV2
	if o.ObjectId.IsNull() || o.ObjectId.IsUnknown() {
		return e, false
	}
	var v []TransferOwnershipObjectId_SdkV2
	d := o.ObjectId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetObjectId sets the value of the ObjectId field in TransferOwnershipRequest_SdkV2.
func (o *TransferOwnershipRequest_SdkV2) SetObjectId(ctx context.Context, v TransferOwnershipObjectId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objectId"]
	o.ObjectId = types.ListValueMust(t, vs)
}

type TrashAlertRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TrashAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type TrashAlertV2Request_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (o TrashAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type TrashQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TrashQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type UpdateAlertRequest_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Id types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"update_mask"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(UpdateAlertRequestAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":                     o.Alert,
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"id":                        o.Id,
			"update_mask":               o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: UpdateAlertRequestAlert_SdkV2{}.Type(ctx),
			},
			"auto_resolve_display_name": types.BoolType,
			"id":                        types.StringType,
			"update_mask":               types.StringType,
		},
	}
}

// GetAlert returns the value of the Alert field in UpdateAlertRequest_SdkV2 as
// a UpdateAlertRequestAlert_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertRequest_SdkV2) GetAlert(ctx context.Context) (UpdateAlertRequestAlert_SdkV2, bool) {
	var e UpdateAlertRequestAlert_SdkV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []UpdateAlertRequestAlert_SdkV2
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in UpdateAlertRequest_SdkV2.
func (o *UpdateAlertRequest_SdkV2) SetAlert(ctx context.Context, v UpdateAlertRequestAlert_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type UpdateAlertRequestAlert_SdkV2 struct {
	// Trigger conditions of the alert.
	Condition types.List `tfsdk:"condition"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger"`
}

func (newState *UpdateAlertRequestAlert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAlertRequestAlert_SdkV2) {
}

func (newState *UpdateAlertRequestAlert_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAlertRequestAlert_SdkV2) {
}

func (c UpdateAlertRequestAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
	attrs["condition"] = attrs["condition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_body"] = attrs["custom_body"].SetOptional()
	attrs["custom_subject"] = attrs["custom_subject"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["seconds_to_retrigger"] = attrs["seconds_to_retrigger"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertRequestAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertRequestAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequestAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAlertRequestAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateAlertRequestAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition": basetypes.ListType{
				ElemType: AlertCondition_SdkV2{}.Type(ctx),
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

// GetCondition returns the value of the Condition field in UpdateAlertRequestAlert_SdkV2 as
// a AlertCondition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertRequestAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if o.Condition.IsNull() || o.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := o.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in UpdateAlertRequestAlert_SdkV2.
func (o *UpdateAlertRequestAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	o.Condition = types.ListValueMust(t, vs)
}

type UpdateAlertV2Request_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":       o.Alert,
			"id":          o.Id,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: AlertV2_SdkV2{}.Type(ctx),
			},
			"id":          types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetAlert returns the value of the Alert field in UpdateAlertV2Request_SdkV2 as
// a AlertV2_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertV2Request_SdkV2) GetAlert(ctx context.Context) (AlertV2_SdkV2, bool) {
	var e AlertV2_SdkV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []AlertV2_SdkV2
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in UpdateAlertV2Request_SdkV2.
func (o *UpdateAlertV2Request_SdkV2) SetAlert(ctx context.Context, v AlertV2_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

type UpdateQueryRequest_SdkV2 struct {
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Id types.String `tfsdk:"-"`

	Query types.List `tfsdk:"query"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"update_mask"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(UpdateQueryRequestQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"id":                        o.Id,
			"query":                     o.Query,
			"update_mask":               o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_resolve_display_name": types.BoolType,
			"id":                        types.StringType,
			"query": basetypes.ListType{
				ElemType: UpdateQueryRequestQuery_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetQuery returns the value of the Query field in UpdateQueryRequest_SdkV2 as
// a UpdateQueryRequestQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequest_SdkV2) GetQuery(ctx context.Context) (UpdateQueryRequestQuery_SdkV2, bool) {
	var e UpdateQueryRequestQuery_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []UpdateQueryRequestQuery_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in UpdateQueryRequest_SdkV2.
func (o *UpdateQueryRequest_SdkV2) SetQuery(ctx context.Context, v UpdateQueryRequestQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type UpdateQueryRequestQuery_SdkV2 struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema"`

	Tags types.List `tfsdk:"tags"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *UpdateQueryRequestQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateQueryRequestQuery_SdkV2) {
}

func (newState *UpdateQueryRequestQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateQueryRequestQuery_SdkV2) {
}

func (c UpdateQueryRequestQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_auto_limit"] = attrs["apply_auto_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as_mode"] = attrs["run_as_mode"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQueryRequestQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQueryRequestQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequestQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateQueryRequestQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateQueryRequestQuery_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_auto_limit": types.BoolType,
			"catalog":          types.StringType,
			"description":      types.StringType,
			"display_name":     types.StringType,
			"owner_user_name":  types.StringType,
			"parameters": basetypes.ListType{
				ElemType: QueryParameter_SdkV2{}.Type(ctx),
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

// GetParameters returns the value of the Parameters field in UpdateQueryRequestQuery_SdkV2 as
// a slice of QueryParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequestQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in UpdateQueryRequestQuery_SdkV2.
func (o *UpdateQueryRequestQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in UpdateQueryRequestQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQueryRequestQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in UpdateQueryRequestQuery_SdkV2.
func (o *UpdateQueryRequestQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse_SdkV2) {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateResponse_SdkV2) {
}

func (c UpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateVisualizationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"update_mask"`

	Visualization types.List `tfsdk:"visualization"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(UpdateVisualizationRequestVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":            o.Id,
			"update_mask":   o.UpdateMask,
			"visualization": o.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":          types.StringType,
			"update_mask": types.StringType,
			"visualization": basetypes.ListType{
				ElemType: UpdateVisualizationRequestVisualization_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVisualization returns the value of the Visualization field in UpdateVisualizationRequest_SdkV2 as
// a UpdateVisualizationRequestVisualization_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateVisualizationRequest_SdkV2) GetVisualization(ctx context.Context) (UpdateVisualizationRequestVisualization_SdkV2, bool) {
	var e UpdateVisualizationRequestVisualization_SdkV2
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []UpdateVisualizationRequestVisualization_SdkV2
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in UpdateVisualizationRequest_SdkV2.
func (o *UpdateVisualizationRequest_SdkV2) SetVisualization(ctx context.Context, v UpdateVisualizationRequestVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type UpdateVisualizationRequestVisualization_SdkV2 struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type"`
}

func (newState *UpdateVisualizationRequestVisualization_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVisualizationRequestVisualization_SdkV2) {
}

func (newState *UpdateVisualizationRequestVisualization_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateVisualizationRequestVisualization_SdkV2) {
}

func (c UpdateVisualizationRequestVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["serialized_options"] = attrs["serialized_options"].SetOptional()
	attrs["serialized_query_plan"] = attrs["serialized_query_plan"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVisualizationRequestVisualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVisualizationRequestVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequestVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateVisualizationRequestVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateVisualizationRequestVisualization_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":          types.StringType,
			"serialized_options":    types.StringType,
			"serialized_query_plan": types.StringType,
			"type":                  types.StringType,
		},
	}
}

type UpdateWidgetRequest_SdkV2 struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`

	Options types.List `tfsdk:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWidgetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWidgetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWidgetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWidgetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateWidgetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"id":           types.StringType,
			"options": basetypes.ListType{
				ElemType: WidgetOptions_SdkV2{}.Type(ctx),
			},
			"text":             types.StringType,
			"visualization_id": types.StringType,
			"width":            types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in UpdateWidgetRequest_SdkV2 as
// a WidgetOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWidgetRequest_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in UpdateWidgetRequest_SdkV2.
func (o *UpdateWidgetRequest_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

type User_SdkV2 struct {
	Email types.String `tfsdk:"email"`

	Id types.Int64 `tfsdk:"id"`

	Name types.String `tfsdk:"name"`
}

func (newState *User_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan User_SdkV2) {
}

func (newState *User_SdkV2) SyncEffectiveFieldsDuringRead(existingState User_SdkV2) {
}

func (c User_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email"] = attrs["email"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in User.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a User_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User_SdkV2
// only implements ToObjectValue() and Type().
func (o User_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email": o.Email,
			"id":    o.Id,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o User_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.Int64Type,
			"name":  types.StringType,
		},
	}
}

type Visualization_SdkV2 struct {
	// The timestamp indicating when the visualization was created.
	CreateTime types.String `tfsdk:"create_time"`
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying the visualization.
	Id types.String `tfsdk:"id"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan"`
	// The type of visualization: counter, table, funnel, and so on.
	Type_ types.String `tfsdk:"type"`
	// The timestamp indicating when the visualization was updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *Visualization_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Visualization_SdkV2) {
}

func (newState *Visualization_SdkV2) SyncEffectiveFieldsDuringRead(existingState Visualization_SdkV2) {
}

func (c Visualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["serialized_options"] = attrs["serialized_options"].SetOptional()
	attrs["serialized_query_plan"] = attrs["serialized_query_plan"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Visualization.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Visualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Visualization_SdkV2
// only implements ToObjectValue() and Type().
func (o Visualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Visualization_SdkV2) Type(ctx context.Context) attr.Type {
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

type WarehouseAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *WarehouseAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlRequest_SdkV2) {
}

func (newState *WarehouseAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlRequest_SdkV2) {
}

func (c WarehouseAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehouseAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WarehouseAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type WarehouseAccessControlResponse_SdkV2 struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *WarehouseAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlResponse_SdkV2) {
}

func (newState *WarehouseAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlResponse_SdkV2) {
}

func (c WarehouseAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WarehousePermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehouseAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WarehouseAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: WarehousePermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in WarehouseAccessControlResponse_SdkV2 as
// a slice of WarehousePermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehouseAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]WarehousePermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WarehouseAccessControlResponse_SdkV2.
func (o *WarehouseAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []WarehousePermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type WarehousePermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *WarehousePermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermission_SdkV2) {
}

func (newState *WarehousePermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehousePermission_SdkV2) {
}

func (c WarehousePermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermission_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehousePermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in WarehousePermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in WarehousePermission_SdkV2.
func (o *WarehousePermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type WarehousePermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *WarehousePermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissions_SdkV2) {
}

func (newState *WarehousePermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehousePermissions_SdkV2) {
}

func (c WarehousePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehousePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WarehouseAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WarehousePermissions_SdkV2 as
// a slice of WarehouseAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissions_SdkV2.
func (o *WarehousePermissions_SdkV2) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehousePermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *WarehousePermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissionsDescription_SdkV2) {
}

func (newState *WarehousePermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehousePermissionsDescription_SdkV2) {
}

func (c WarehousePermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehousePermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type WarehousePermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehousePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehousePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehousePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"warehouse_id":        o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehousePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WarehouseAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WarehousePermissionsRequest_SdkV2 as
// a slice of WarehouseAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WarehousePermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissionsRequest_SdkV2.
func (o *WarehousePermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehouseTypePair_SdkV2 struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (newState *WarehouseTypePair_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseTypePair_SdkV2) {
}

func (newState *WarehouseTypePair_SdkV2) SyncEffectiveFieldsDuringRead(existingState WarehouseTypePair_SdkV2) {
}

func (c WarehouseTypePair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["warehouse_type"] = attrs["warehouse_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WarehouseTypePair.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WarehouseTypePair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseTypePair_SdkV2
// only implements ToObjectValue() and Type().
func (o WarehouseTypePair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":        o.Enabled,
			"warehouse_type": o.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WarehouseTypePair_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":        types.BoolType,
			"warehouse_type": types.StringType,
		},
	}
}

type Widget_SdkV2 struct {
	// The unique ID for this widget.
	Id types.String `tfsdk:"id"`

	Options types.List `tfsdk:"options"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization types.List `tfsdk:"visualization"`
	// Unused field.
	Width types.Int64 `tfsdk:"width"`
}

func (newState *Widget_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Widget_SdkV2) {
}

func (newState *Widget_SdkV2) SyncEffectiveFieldsDuringRead(existingState Widget_SdkV2) {
}

func (c Widget_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["visualization"] = attrs["visualization"].SetOptional()
	attrs["visualization"] = attrs["visualization"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["width"] = attrs["width"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Widget.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Widget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":       reflect.TypeOf(WidgetOptions_SdkV2{}),
		"visualization": reflect.TypeOf(LegacyVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Widget_SdkV2
// only implements ToObjectValue() and Type().
func (o Widget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Widget_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"options": basetypes.ListType{
				ElemType: WidgetOptions_SdkV2{}.Type(ctx),
			},
			"visualization": basetypes.ListType{
				ElemType: LegacyVisualization_SdkV2{}.Type(ctx),
			},
			"width": types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in Widget_SdkV2 as
// a WidgetOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Widget_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Widget_SdkV2.
func (o *Widget_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	o.Options = types.ListValueMust(t, vs)
}

// GetVisualization returns the value of the Visualization field in Widget_SdkV2 as
// a LegacyVisualization_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Widget_SdkV2) GetVisualization(ctx context.Context) (LegacyVisualization_SdkV2, bool) {
	var e LegacyVisualization_SdkV2
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v []LegacyVisualization_SdkV2
	d := o.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in Widget_SdkV2.
func (o *Widget_SdkV2) SetVisualization(ctx context.Context, v LegacyVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	o.Visualization = types.ListValueMust(t, vs)
}

type WidgetOptions_SdkV2 struct {
	// Timestamp when this object was created
	CreatedAt types.String `tfsdk:"created_at"`
	// Custom description of the widget
	Description types.String `tfsdk:"description"`
	// Whether this widget is hidden on the dashboard.
	IsHidden types.Bool `tfsdk:"isHidden"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings types.Object `tfsdk:"parameterMappings"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position types.List `tfsdk:"position"`
	// Custom title of the widget
	Title types.String `tfsdk:"title"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (newState *WidgetOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetOptions_SdkV2) {
}

func (newState *WidgetOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState WidgetOptions_SdkV2) {
}

func (c WidgetOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["isHidden"] = attrs["isHidden"].SetOptional()
	attrs["parameterMappings"] = attrs["parameterMappings"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["position"] = attrs["position"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["title"] = attrs["title"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WidgetOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WidgetOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"position": reflect.TypeOf(WidgetPosition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o WidgetOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WidgetOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":        types.StringType,
			"description":       types.StringType,
			"isHidden":          types.BoolType,
			"parameterMappings": types.ObjectType{},
			"position": basetypes.ListType{
				ElemType: WidgetPosition_SdkV2{}.Type(ctx),
			},
			"title":      types.StringType,
			"updated_at": types.StringType,
		},
	}
}

// GetPosition returns the value of the Position field in WidgetOptions_SdkV2 as
// a WidgetPosition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *WidgetOptions_SdkV2) GetPosition(ctx context.Context) (WidgetPosition_SdkV2, bool) {
	var e WidgetPosition_SdkV2
	if o.Position.IsNull() || o.Position.IsUnknown() {
		return e, false
	}
	var v []WidgetPosition_SdkV2
	d := o.Position.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPosition sets the value of the Position field in WidgetOptions_SdkV2.
func (o *WidgetOptions_SdkV2) SetPosition(ctx context.Context, v WidgetPosition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["position"]
	o.Position = types.ListValueMust(t, vs)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition_SdkV2 struct {
	// reserved for internal use
	AutoHeight types.Bool `tfsdk:"autoHeight"`
	// column in the dashboard grid. Values start with 0
	Col types.Int64 `tfsdk:"col"`
	// row in the dashboard grid. Values start with 0
	Row types.Int64 `tfsdk:"row"`
	// width of the widget measured in dashboard grid cells
	SizeX types.Int64 `tfsdk:"sizeX"`
	// height of the widget measured in dashboard grid cells
	SizeY types.Int64 `tfsdk:"sizeY"`
}

func (newState *WidgetPosition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetPosition_SdkV2) {
}

func (newState *WidgetPosition_SdkV2) SyncEffectiveFieldsDuringRead(existingState WidgetPosition_SdkV2) {
}

func (c WidgetPosition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoHeight"] = attrs["autoHeight"].SetOptional()
	attrs["col"] = attrs["col"].SetOptional()
	attrs["row"] = attrs["row"].SetOptional()
	attrs["sizeX"] = attrs["sizeX"].SetOptional()
	attrs["sizeY"] = attrs["sizeY"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WidgetPosition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WidgetPosition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetPosition_SdkV2
// only implements ToObjectValue() and Type().
func (o WidgetPosition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WidgetPosition_SdkV2) Type(ctx context.Context) attr.Type {
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
