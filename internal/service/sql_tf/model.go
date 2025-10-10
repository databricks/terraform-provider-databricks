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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccessControl struct {
	GroupName types.String `tfsdk:"group_name"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel types.String `tfsdk:"permission_level"`

	UserName types.String `tfsdk:"user_name"`
}

func (to *AccessControl) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccessControl) {
}

func (to *AccessControl) SyncFieldsDuringRead(ctx context.Context, from AccessControl) {
}

func (c AccessControl) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Condition types.Object `tfsdk:"condition"`
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

func (to *Alert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Alert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				// Recursively sync the fields of Condition
				toCondition.SyncFieldsDuringCreateOrUpdate(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (to *Alert) SyncFieldsDuringRead(ctx context.Context, from Alert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (c Alert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
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
			"condition":            AlertCondition{}.Type(ctx),
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
	var v AlertCondition
	d := o.Condition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCondition sets the value of the Condition field in Alert.
func (o *Alert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := v.ToObjectValue(ctx)
	o.Condition = vs
}

type AlertCondition struct {
	// Alert state if result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state"`
	// Operator used for comparison in alert evaluation.
	Op types.String `tfsdk:"op"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand types.Object `tfsdk:"operand"`
	// Threshold value used for comparison in alert evaluation.
	Threshold types.Object `tfsdk:"threshold"`
}

func (to *AlertCondition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertCondition) {
	if !from.Operand.IsNull() && !from.Operand.IsUnknown() {
		if toOperand, ok := to.GetOperand(ctx); ok {
			if fromOperand, ok := from.GetOperand(ctx); ok {
				// Recursively sync the fields of Operand
				toOperand.SyncFieldsDuringCreateOrUpdate(ctx, fromOperand)
				to.SetOperand(ctx, toOperand)
			}
		}
	}
	if !from.Threshold.IsNull() && !from.Threshold.IsUnknown() {
		if toThreshold, ok := to.GetThreshold(ctx); ok {
			if fromThreshold, ok := from.GetThreshold(ctx); ok {
				// Recursively sync the fields of Threshold
				toThreshold.SyncFieldsDuringCreateOrUpdate(ctx, fromThreshold)
				to.SetThreshold(ctx, toThreshold)
			}
		}
	}
}

func (to *AlertCondition) SyncFieldsDuringRead(ctx context.Context, from AlertCondition) {
	if !from.Operand.IsNull() && !from.Operand.IsUnknown() {
		if toOperand, ok := to.GetOperand(ctx); ok {
			if fromOperand, ok := from.GetOperand(ctx); ok {
				toOperand.SyncFieldsDuringRead(ctx, fromOperand)
				to.SetOperand(ctx, toOperand)
			}
		}
	}
	if !from.Threshold.IsNull() && !from.Threshold.IsUnknown() {
		if toThreshold, ok := to.GetThreshold(ctx); ok {
			if fromThreshold, ok := from.GetThreshold(ctx); ok {
				toThreshold.SyncFieldsDuringRead(ctx, fromThreshold)
				to.SetThreshold(ctx, toThreshold)
			}
		}
	}
}

func (c AlertCondition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["op"] = attrs["op"].SetOptional()
	attrs["operand"] = attrs["operand"].SetOptional()
	attrs["threshold"] = attrs["threshold"].SetOptional()

	return attrs
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
			"operand":            AlertConditionOperand{}.Type(ctx),
			"threshold":          AlertConditionThreshold{}.Type(ctx),
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
	var v AlertConditionOperand
	d := o.Operand.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOperand sets the value of the Operand field in AlertCondition.
func (o *AlertCondition) SetOperand(ctx context.Context, v AlertConditionOperand) {
	vs := v.ToObjectValue(ctx)
	o.Operand = vs
}

// GetThreshold returns the value of the Threshold field in AlertCondition as
// a AlertConditionThreshold value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertCondition) GetThreshold(ctx context.Context) (AlertConditionThreshold, bool) {
	var e AlertConditionThreshold
	if o.Threshold.IsNull() || o.Threshold.IsUnknown() {
		return e, false
	}
	var v AlertConditionThreshold
	d := o.Threshold.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetThreshold sets the value of the Threshold field in AlertCondition.
func (o *AlertCondition) SetThreshold(ctx context.Context, v AlertConditionThreshold) {
	vs := v.ToObjectValue(ctx)
	o.Threshold = vs
}

type AlertConditionOperand struct {
	Column types.Object `tfsdk:"column"`
}

func (to *AlertConditionOperand) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConditionOperand) {
	if !from.Column.IsNull() && !from.Column.IsUnknown() {
		if toColumn, ok := to.GetColumn(ctx); ok {
			if fromColumn, ok := from.GetColumn(ctx); ok {
				// Recursively sync the fields of Column
				toColumn.SyncFieldsDuringCreateOrUpdate(ctx, fromColumn)
				to.SetColumn(ctx, toColumn)
			}
		}
	}
}

func (to *AlertConditionOperand) SyncFieldsDuringRead(ctx context.Context, from AlertConditionOperand) {
	if !from.Column.IsNull() && !from.Column.IsUnknown() {
		if toColumn, ok := to.GetColumn(ctx); ok {
			if fromColumn, ok := from.GetColumn(ctx); ok {
				toColumn.SyncFieldsDuringRead(ctx, fromColumn)
				to.SetColumn(ctx, toColumn)
			}
		}
	}
}

func (c AlertConditionOperand) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetOptional()

	return attrs
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
			"column": AlertOperandColumn{}.Type(ctx),
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
	var v AlertOperandColumn
	d := o.Column.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumn sets the value of the Column field in AlertConditionOperand.
func (o *AlertConditionOperand) SetColumn(ctx context.Context, v AlertOperandColumn) {
	vs := v.ToObjectValue(ctx)
	o.Column = vs
}

type AlertConditionThreshold struct {
	Value types.Object `tfsdk:"value"`
}

func (to *AlertConditionThreshold) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConditionThreshold) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				// Recursively sync the fields of Value
				toValue.SyncFieldsDuringCreateOrUpdate(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (to *AlertConditionThreshold) SyncFieldsDuringRead(ctx context.Context, from AlertConditionThreshold) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (c AlertConditionThreshold) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
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
			"value": AlertOperandValue{}.Type(ctx),
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
	var v AlertOperandValue
	d := o.Value.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValue sets the value of the Value field in AlertConditionThreshold.
func (o *AlertConditionThreshold) SetValue(ctx context.Context, v AlertOperandValue) {
	vs := v.ToObjectValue(ctx)
	o.Value = vs
}

type AlertOperandColumn struct {
	Name types.String `tfsdk:"name"`
}

func (to *AlertOperandColumn) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOperandColumn) {
}

func (to *AlertOperandColumn) SyncFieldsDuringRead(ctx context.Context, from AlertOperandColumn) {
}

func (c AlertOperandColumn) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	BoolValue types.Bool `tfsdk:"bool_value"`

	DoubleValue types.Float64 `tfsdk:"double_value"`

	StringValue types.String `tfsdk:"string_value"`
}

func (to *AlertOperandValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOperandValue) {
}

func (to *AlertOperandValue) SyncFieldsDuringRead(ctx context.Context, from AlertOperandValue) {
}

func (c AlertOperandValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *AlertOptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOptions) {
}

func (to *AlertOptions) SyncFieldsDuringRead(ctx context.Context, from AlertOptions) {
}

func (c AlertOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

	Options types.Object `tfsdk:"options"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`

	Tags types.List `tfsdk:"tags"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`
}

func (to *AlertQuery) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertQuery) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *AlertQuery) SyncFieldsDuringRead(ctx context.Context, from AlertQuery) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c AlertQuery) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["data_source_id"] = attrs["data_source_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_archived"] = attrs["is_archived"].SetOptional()
	attrs["is_draft"] = attrs["is_draft"].SetOptional()
	attrs["is_safe"] = attrs["is_safe"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
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
			"options":        QueryOptions{}.Type(ctx),
			"query":          types.StringType,
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
	var v QueryOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in AlertQuery.
func (o *AlertQuery) SetOptions(ctx context.Context, v QueryOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom description for the alert. support mustache template.
	CustomDescription types.String `tfsdk:"custom_description"`
	// Custom summary for the alert. support mustache template.
	CustomSummary types.String `tfsdk:"custom_summary"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name"`
	// The actual identity that will be used to execute the alert. This is an
	// output-only field that shows the resolved run-as identity after applying
	// permissions and defaults.
	EffectiveRunAs types.Object `tfsdk:"effective_run_as"`

	Evaluation types.Object `tfsdk:"evaluation"`
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
	// Specifies the identity that will be used to run the alert. This field
	// allows you to configure alerts to run as a specific user or service
	// principal. - For user identity: Set `user_name` to the email of an active
	// workspace user. Users can only set this to their own email. - For service
	// principal: Set `service_principal_name` to the application ID. Requires
	// the `servicePrincipal/user` role. If not specified, the alert will run as
	// the request user.
	RunAs types.Object `tfsdk:"run_as"`
	// The run as username or application ID of service principal. On Create and
	// Update, this field can be set to application ID of an active service
	// principal. Setting this field requires the servicePrincipal/user role.
	// Deprecated: Use `run_as` field instead. This field will be removed in a
	// future release.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`

	Schedule types.Object `tfsdk:"schedule"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *AlertV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2) {
	if !from.EffectiveRunAs.IsNull() && !from.EffectiveRunAs.IsUnknown() {
		if toEffectiveRunAs, ok := to.GetEffectiveRunAs(ctx); ok {
			if fromEffectiveRunAs, ok := from.GetEffectiveRunAs(ctx); ok {
				// Recursively sync the fields of EffectiveRunAs
				toEffectiveRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveRunAs)
				to.SetEffectiveRunAs(ctx, toEffectiveRunAs)
			}
		}
	}
	if !from.Evaluation.IsNull() && !from.Evaluation.IsUnknown() {
		if toEvaluation, ok := to.GetEvaluation(ctx); ok {
			if fromEvaluation, ok := from.GetEvaluation(ctx); ok {
				// Recursively sync the fields of Evaluation
				toEvaluation.SyncFieldsDuringCreateOrUpdate(ctx, fromEvaluation)
				to.SetEvaluation(ctx, toEvaluation)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				// Recursively sync the fields of RunAs
				toRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
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
}

func (to *AlertV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2) {
	if !from.EffectiveRunAs.IsNull() && !from.EffectiveRunAs.IsUnknown() {
		if toEffectiveRunAs, ok := to.GetEffectiveRunAs(ctx); ok {
			if fromEffectiveRunAs, ok := from.GetEffectiveRunAs(ctx); ok {
				toEffectiveRunAs.SyncFieldsDuringRead(ctx, fromEffectiveRunAs)
				to.SetEffectiveRunAs(ctx, toEffectiveRunAs)
			}
		}
	}
	if !from.Evaluation.IsNull() && !from.Evaluation.IsUnknown() {
		if toEvaluation, ok := to.GetEvaluation(ctx); ok {
			if fromEvaluation, ok := from.GetEvaluation(ctx); ok {
				toEvaluation.SyncFieldsDuringRead(ctx, fromEvaluation)
				to.SetEvaluation(ctx, toEvaluation)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				toRunAs.SyncFieldsDuringRead(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
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
}

func (c AlertV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["custom_description"] = attrs["custom_description"].SetOptional()
	attrs["custom_summary"] = attrs["custom_summary"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["effective_run_as"] = attrs["effective_run_as"].SetComputed()
	attrs["evaluation"] = attrs["evaluation"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetOptional()
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
func (a AlertV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_run_as": reflect.TypeOf(AlertV2RunAs{}),
		"evaluation":       reflect.TypeOf(AlertV2Evaluation{}),
		"run_as":           reflect.TypeOf(AlertV2RunAs{}),
		"schedule":         reflect.TypeOf(CronSchedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2
// only implements ToObjectValue() and Type().
func (o AlertV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":        o.CreateTime,
			"custom_description": o.CustomDescription,
			"custom_summary":     o.CustomSummary,
			"display_name":       o.DisplayName,
			"effective_run_as":   o.EffectiveRunAs,
			"evaluation":         o.Evaluation,
			"id":                 o.Id,
			"lifecycle_state":    o.LifecycleState,
			"owner_user_name":    o.OwnerUserName,
			"parent_path":        o.ParentPath,
			"query_text":         o.QueryText,
			"run_as":             o.RunAs,
			"run_as_user_name":   o.RunAsUserName,
			"schedule":           o.Schedule,
			"update_time":        o.UpdateTime,
			"warehouse_id":       o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"custom_description": types.StringType,
			"custom_summary":     types.StringType,
			"display_name":       types.StringType,
			"effective_run_as":   AlertV2RunAs{}.Type(ctx),
			"evaluation":         AlertV2Evaluation{}.Type(ctx),
			"id":                 types.StringType,
			"lifecycle_state":    types.StringType,
			"owner_user_name":    types.StringType,
			"parent_path":        types.StringType,
			"query_text":         types.StringType,
			"run_as":             AlertV2RunAs{}.Type(ctx),
			"run_as_user_name":   types.StringType,
			"schedule":           CronSchedule{}.Type(ctx),
			"update_time":        types.StringType,
			"warehouse_id":       types.StringType,
		},
	}
}

// GetEffectiveRunAs returns the value of the EffectiveRunAs field in AlertV2 as
// a AlertV2RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2) GetEffectiveRunAs(ctx context.Context) (AlertV2RunAs, bool) {
	var e AlertV2RunAs
	if o.EffectiveRunAs.IsNull() || o.EffectiveRunAs.IsUnknown() {
		return e, false
	}
	var v AlertV2RunAs
	d := o.EffectiveRunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveRunAs sets the value of the EffectiveRunAs field in AlertV2.
func (o *AlertV2) SetEffectiveRunAs(ctx context.Context, v AlertV2RunAs) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveRunAs = vs
}

// GetEvaluation returns the value of the Evaluation field in AlertV2 as
// a AlertV2Evaluation value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2) GetEvaluation(ctx context.Context) (AlertV2Evaluation, bool) {
	var e AlertV2Evaluation
	if o.Evaluation.IsNull() || o.Evaluation.IsUnknown() {
		return e, false
	}
	var v AlertV2Evaluation
	d := o.Evaluation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvaluation sets the value of the Evaluation field in AlertV2.
func (o *AlertV2) SetEvaluation(ctx context.Context, v AlertV2Evaluation) {
	vs := v.ToObjectValue(ctx)
	o.Evaluation = vs
}

// GetRunAs returns the value of the RunAs field in AlertV2 as
// a AlertV2RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2) GetRunAs(ctx context.Context) (AlertV2RunAs, bool) {
	var e AlertV2RunAs
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v AlertV2RunAs
	d := o.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in AlertV2.
func (o *AlertV2) SetRunAs(ctx context.Context, v AlertV2RunAs) {
	vs := v.ToObjectValue(ctx)
	o.RunAs = vs
}

// GetSchedule returns the value of the Schedule field in AlertV2 as
// a CronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2) GetSchedule(ctx context.Context) (CronSchedule, bool) {
	var e CronSchedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v CronSchedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in AlertV2.
func (o *AlertV2) SetSchedule(ctx context.Context, v CronSchedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation.
	ComparisonOperator types.String `tfsdk:"comparison_operator"`
	// Alert state if result is empty. Please avoid setting this field to be
	// `UNKNOWN` because `UNKNOWN` state is planned to be deprecated.
	EmptyResultState types.String `tfsdk:"empty_result_state"`
	// Timestamp of the last evaluation.
	LastEvaluatedAt types.String `tfsdk:"last_evaluated_at"`
	// User or Notification Destination to notify when alert is triggered.
	Notification types.Object `tfsdk:"notification"`
	// Source column from result to use to evaluate alert
	Source types.Object `tfsdk:"source"`
	// Latest state of alert evaluation.
	State types.String `tfsdk:"state"`
	// Threshold to user for alert evaluation, can be a column or a value.
	Threshold types.Object `tfsdk:"threshold"`
}

func (to *AlertV2Evaluation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Evaluation) {
	if !from.Notification.IsNull() && !from.Notification.IsUnknown() {
		if toNotification, ok := to.GetNotification(ctx); ok {
			if fromNotification, ok := from.GetNotification(ctx); ok {
				// Recursively sync the fields of Notification
				toNotification.SyncFieldsDuringCreateOrUpdate(ctx, fromNotification)
				to.SetNotification(ctx, toNotification)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				// Recursively sync the fields of Source
				toSource.SyncFieldsDuringCreateOrUpdate(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.Threshold.IsNull() && !from.Threshold.IsUnknown() {
		if toThreshold, ok := to.GetThreshold(ctx); ok {
			if fromThreshold, ok := from.GetThreshold(ctx); ok {
				// Recursively sync the fields of Threshold
				toThreshold.SyncFieldsDuringCreateOrUpdate(ctx, fromThreshold)
				to.SetThreshold(ctx, toThreshold)
			}
		}
	}
}

func (to *AlertV2Evaluation) SyncFieldsDuringRead(ctx context.Context, from AlertV2Evaluation) {
	if !from.Notification.IsNull() && !from.Notification.IsUnknown() {
		if toNotification, ok := to.GetNotification(ctx); ok {
			if fromNotification, ok := from.GetNotification(ctx); ok {
				toNotification.SyncFieldsDuringRead(ctx, fromNotification)
				to.SetNotification(ctx, toNotification)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				toSource.SyncFieldsDuringRead(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.Threshold.IsNull() && !from.Threshold.IsUnknown() {
		if toThreshold, ok := to.GetThreshold(ctx); ok {
			if fromThreshold, ok := from.GetThreshold(ctx); ok {
				toThreshold.SyncFieldsDuringRead(ctx, fromThreshold)
				to.SetThreshold(ctx, toThreshold)
			}
		}
	}
}

func (c AlertV2Evaluation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comparison_operator"] = attrs["comparison_operator"].SetOptional()
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["last_evaluated_at"] = attrs["last_evaluated_at"].SetComputed()
	attrs["notification"] = attrs["notification"].SetOptional()
	attrs["source"] = attrs["source"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["threshold"] = attrs["threshold"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Evaluation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Evaluation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notification": reflect.TypeOf(AlertV2Notification{}),
		"source":       reflect.TypeOf(AlertV2OperandColumn{}),
		"threshold":    reflect.TypeOf(AlertV2Operand{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Evaluation
// only implements ToObjectValue() and Type().
func (o AlertV2Evaluation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AlertV2Evaluation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comparison_operator": types.StringType,
			"empty_result_state":  types.StringType,
			"last_evaluated_at":   types.StringType,
			"notification":        AlertV2Notification{}.Type(ctx),
			"source":              AlertV2OperandColumn{}.Type(ctx),
			"state":               types.StringType,
			"threshold":           AlertV2Operand{}.Type(ctx),
		},
	}
}

// GetNotification returns the value of the Notification field in AlertV2Evaluation as
// a AlertV2Notification value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation) GetNotification(ctx context.Context) (AlertV2Notification, bool) {
	var e AlertV2Notification
	if o.Notification.IsNull() || o.Notification.IsUnknown() {
		return e, false
	}
	var v AlertV2Notification
	d := o.Notification.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotification sets the value of the Notification field in AlertV2Evaluation.
func (o *AlertV2Evaluation) SetNotification(ctx context.Context, v AlertV2Notification) {
	vs := v.ToObjectValue(ctx)
	o.Notification = vs
}

// GetSource returns the value of the Source field in AlertV2Evaluation as
// a AlertV2OperandColumn value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation) GetSource(ctx context.Context) (AlertV2OperandColumn, bool) {
	var e AlertV2OperandColumn
	if o.Source.IsNull() || o.Source.IsUnknown() {
		return e, false
	}
	var v AlertV2OperandColumn
	d := o.Source.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSource sets the value of the Source field in AlertV2Evaluation.
func (o *AlertV2Evaluation) SetSource(ctx context.Context, v AlertV2OperandColumn) {
	vs := v.ToObjectValue(ctx)
	o.Source = vs
}

// GetThreshold returns the value of the Threshold field in AlertV2Evaluation as
// a AlertV2Operand value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Evaluation) GetThreshold(ctx context.Context) (AlertV2Operand, bool) {
	var e AlertV2Operand
	if o.Threshold.IsNull() || o.Threshold.IsUnknown() {
		return e, false
	}
	var v AlertV2Operand
	d := o.Threshold.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetThreshold sets the value of the Threshold field in AlertV2Evaluation.
func (o *AlertV2Evaluation) SetThreshold(ctx context.Context, v AlertV2Operand) {
	vs := v.ToObjectValue(ctx)
	o.Threshold = vs
}

type AlertV2Notification struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	RetriggerSeconds types.Int64 `tfsdk:"retrigger_seconds"`

	Subscriptions types.List `tfsdk:"subscriptions"`
}

func (to *AlertV2Notification) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Notification) {
	if !from.Subscriptions.IsNull() && !from.Subscriptions.IsUnknown() && to.Subscriptions.IsNull() && len(from.Subscriptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subscriptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subscriptions = from.Subscriptions
	}
}

func (to *AlertV2Notification) SyncFieldsDuringRead(ctx context.Context, from AlertV2Notification) {
	if !from.Subscriptions.IsNull() && !from.Subscriptions.IsUnknown() && to.Subscriptions.IsNull() && len(from.Subscriptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subscriptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subscriptions = from.Subscriptions
	}
}

func (c AlertV2Notification) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AlertV2Notification) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(AlertV2Subscription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Notification
// only implements ToObjectValue() and Type().
func (o AlertV2Notification) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notify_on_ok":      o.NotifyOnOk,
			"retrigger_seconds": o.RetriggerSeconds,
			"subscriptions":     o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Notification) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notify_on_ok":      types.BoolType,
			"retrigger_seconds": types.Int64Type,
			"subscriptions": basetypes.ListType{
				ElemType: AlertV2Subscription{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in AlertV2Notification as
// a slice of AlertV2Subscription values.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Notification) GetSubscriptions(ctx context.Context) ([]AlertV2Subscription, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []AlertV2Subscription
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in AlertV2Notification.
func (o *AlertV2Notification) SetSubscriptions(ctx context.Context, v []AlertV2Subscription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type AlertV2Operand struct {
	Column types.Object `tfsdk:"column"`

	Value types.Object `tfsdk:"value"`
}

func (to *AlertV2Operand) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Operand) {
	if !from.Column.IsNull() && !from.Column.IsUnknown() {
		if toColumn, ok := to.GetColumn(ctx); ok {
			if fromColumn, ok := from.GetColumn(ctx); ok {
				// Recursively sync the fields of Column
				toColumn.SyncFieldsDuringCreateOrUpdate(ctx, fromColumn)
				to.SetColumn(ctx, toColumn)
			}
		}
	}
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				// Recursively sync the fields of Value
				toValue.SyncFieldsDuringCreateOrUpdate(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (to *AlertV2Operand) SyncFieldsDuringRead(ctx context.Context, from AlertV2Operand) {
	if !from.Column.IsNull() && !from.Column.IsUnknown() {
		if toColumn, ok := to.GetColumn(ctx); ok {
			if fromColumn, ok := from.GetColumn(ctx); ok {
				toColumn.SyncFieldsDuringRead(ctx, fromColumn)
				to.SetColumn(ctx, toColumn)
			}
		}
	}
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (c AlertV2Operand) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2Operand.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2Operand) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertV2OperandColumn{}),
		"value":  reflect.TypeOf(AlertV2OperandValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Operand
// only implements ToObjectValue() and Type().
func (o AlertV2Operand) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": o.Column,
			"value":  o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Operand) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column": AlertV2OperandColumn{}.Type(ctx),
			"value":  AlertV2OperandValue{}.Type(ctx),
		},
	}
}

// GetColumn returns the value of the Column field in AlertV2Operand as
// a AlertV2OperandColumn value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Operand) GetColumn(ctx context.Context) (AlertV2OperandColumn, bool) {
	var e AlertV2OperandColumn
	if o.Column.IsNull() || o.Column.IsUnknown() {
		return e, false
	}
	var v AlertV2OperandColumn
	d := o.Column.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumn sets the value of the Column field in AlertV2Operand.
func (o *AlertV2Operand) SetColumn(ctx context.Context, v AlertV2OperandColumn) {
	vs := v.ToObjectValue(ctx)
	o.Column = vs
}

// GetValue returns the value of the Value field in AlertV2Operand as
// a AlertV2OperandValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *AlertV2Operand) GetValue(ctx context.Context) (AlertV2OperandValue, bool) {
	var e AlertV2OperandValue
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v AlertV2OperandValue
	d := o.Value.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValue sets the value of the Value field in AlertV2Operand.
func (o *AlertV2Operand) SetValue(ctx context.Context, v AlertV2OperandValue) {
	vs := v.ToObjectValue(ctx)
	o.Value = vs
}

type AlertV2OperandColumn struct {
	Aggregation types.String `tfsdk:"aggregation"`

	Display types.String `tfsdk:"display"`

	Name types.String `tfsdk:"name"`
}

func (to *AlertV2OperandColumn) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2OperandColumn) {
}

func (to *AlertV2OperandColumn) SyncFieldsDuringRead(ctx context.Context, from AlertV2OperandColumn) {
}

func (c AlertV2OperandColumn) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AlertV2OperandColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandColumn
// only implements ToObjectValue() and Type().
func (o AlertV2OperandColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aggregation": o.Aggregation,
			"display":     o.Display,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2OperandColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aggregation": types.StringType,
			"display":     types.StringType,
			"name":        types.StringType,
		},
	}
}

type AlertV2OperandValue struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	DoubleValue types.Float64 `tfsdk:"double_value"`

	StringValue types.String `tfsdk:"string_value"`
}

func (to *AlertV2OperandValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2OperandValue) {
}

func (to *AlertV2OperandValue) SyncFieldsDuringRead(ctx context.Context, from AlertV2OperandValue) {
}

func (c AlertV2OperandValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AlertV2OperandValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandValue
// only implements ToObjectValue() and Type().
func (o AlertV2OperandValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"double_value": o.DoubleValue,
			"string_value": o.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2OperandValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"double_value": types.Float64Type,
			"string_value": types.StringType,
		},
	}
}

type AlertV2RunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Can only set this field to their
	// own email.
	UserName types.String `tfsdk:"user_name"`
}

func (to *AlertV2RunAs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2RunAs) {
}

func (to *AlertV2RunAs) SyncFieldsDuringRead(ctx context.Context, from AlertV2RunAs) {
}

func (c AlertV2RunAs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2RunAs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AlertV2RunAs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2RunAs
// only implements ToObjectValue() and Type().
func (o AlertV2RunAs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2RunAs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AlertV2Subscription struct {
	DestinationId types.String `tfsdk:"destination_id"`

	UserEmail types.String `tfsdk:"user_email"`
}

func (to *AlertV2Subscription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Subscription) {
}

func (to *AlertV2Subscription) SyncFieldsDuringRead(ctx context.Context, from AlertV2Subscription) {
}

func (c AlertV2Subscription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AlertV2Subscription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Subscription
// only implements ToObjectValue() and Type().
func (o AlertV2Subscription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": o.DestinationId,
			"user_email":     o.UserEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AlertV2Subscription) Type(ctx context.Context) attr.Type {
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
type BaseChunkInfo struct {
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

func (to *BaseChunkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BaseChunkInfo) {
}

func (to *BaseChunkInfo) SyncFieldsDuringRead(ctx context.Context, from BaseChunkInfo) {
}

func (c BaseChunkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (to *CancelExecutionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelExecutionRequest) {
}

func (to *CancelExecutionRequest) SyncFieldsDuringRead(ctx context.Context, from CancelExecutionRequest) {
}

func (c CancelExecutionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_id"] = attrs["statement_id"].SetRequired()

	return attrs
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

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version"`

	Name types.String `tfsdk:"name"`
}

func (to *Channel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Channel) {
}

func (to *Channel) SyncFieldsDuringRead(ctx context.Context, from Channel) {
}

func (c Channel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	DbsqlVersion types.String `tfsdk:"dbsql_version"`
	// Name of the channel
	Name types.String `tfsdk:"name"`
}

func (to *ChannelInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChannelInfo) {
}

func (to *ChannelInfo) SyncFieldsDuringRead(ctx context.Context, from ChannelInfo) {
}

func (c ChannelInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type ClientConfig struct {
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

func (to *ClientConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClientConfig) {
}

func (to *ClientConfig) SyncFieldsDuringRead(ctx context.Context, from ClientConfig) {
}

func (c ClientConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClientConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientConfig
// only implements ToObjectValue() and Type().
func (o ClientConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClientConfig) Type(ctx context.Context) attr.Type {
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

type ColumnInfo struct {
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

func (to *ColumnInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo) {
}

func (to *ColumnInfo) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo) {
}

func (c ColumnInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.Object `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

func (to *CreateAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *CreateAlert) SyncFieldsDuringRead(ctx context.Context, from CreateAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (c CreateAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetRequired()
	attrs["rearm"] = attrs["rearm"].SetOptional()

	return attrs
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
			"name":     types.StringType,
			"options":  AlertOptions{}.Type(ctx),
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
	var v AlertOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in CreateAlert.
func (o *CreateAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

type CreateAlertRequest struct {
	Alert types.Object `tfsdk:"alert"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`
}

func (to *CreateAlertRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertRequest) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				// Recursively sync the fields of Alert
				toAlert.SyncFieldsDuringCreateOrUpdate(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (to *CreateAlertRequest) SyncFieldsDuringRead(ctx context.Context, from CreateAlertRequest) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (c CreateAlertRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetOptional()
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()

	return attrs
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
			"alert":                     o.Alert,
			"auto_resolve_display_name": o.AutoResolveDisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert":                     CreateAlertRequestAlert{}.Type(ctx),
			"auto_resolve_display_name": types.BoolType,
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
	var v CreateAlertRequestAlert
	d := o.Alert.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlert sets the value of the Alert field in CreateAlertRequest.
func (o *CreateAlertRequest) SetAlert(ctx context.Context, v CreateAlertRequestAlert) {
	vs := v.ToObjectValue(ctx)
	o.Alert = vs
}

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition"`
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

func (to *CreateAlertRequestAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertRequestAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				// Recursively sync the fields of Condition
				toCondition.SyncFieldsDuringCreateOrUpdate(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (to *CreateAlertRequestAlert) SyncFieldsDuringRead(ctx context.Context, from CreateAlertRequestAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (c CreateAlertRequestAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
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
			"condition":            AlertCondition{}.Type(ctx),
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
	var v AlertCondition
	d := o.Condition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCondition sets the value of the Condition field in CreateAlertRequestAlert.
func (o *CreateAlertRequestAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := v.ToObjectValue(ctx)
	o.Condition = vs
}

type CreateAlertV2Request struct {
	Alert types.Object `tfsdk:"alert"`
}

func (to *CreateAlertV2Request) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertV2Request) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				// Recursively sync the fields of Alert
				toAlert.SyncFieldsDuringCreateOrUpdate(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (to *CreateAlertV2Request) SyncFieldsDuringRead(ctx context.Context, from CreateAlertV2Request) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (c CreateAlertV2Request) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAlertV2Request) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertV2Request
// only implements ToObjectValue() and Type().
func (o CreateAlertV2Request) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert": o.Alert,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAlertV2Request) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": AlertV2{}.Type(ctx),
		},
	}
}

// GetAlert returns the value of the Alert field in CreateAlertV2Request as
// a AlertV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAlertV2Request) GetAlert(ctx context.Context) (AlertV2, bool) {
	var e AlertV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v AlertV2
	d := o.Alert.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlert sets the value of the Alert field in CreateAlertV2Request.
func (o *CreateAlertV2Request) SetAlert(ctx context.Context, v AlertV2) {
	vs := v.ToObjectValue(ctx)
	o.Alert = vs
}

type CreateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Query types.Object `tfsdk:"query"`
}

func (to *CreateQueryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryRequest) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				// Recursively sync the fields of Query
				toQuery.SyncFieldsDuringCreateOrUpdate(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (to *CreateQueryRequest) SyncFieldsDuringRead(ctx context.Context, from CreateQueryRequest) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (c CreateQueryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()

	return attrs
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
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"query":                     o.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_resolve_display_name": types.BoolType,
			"query":                     CreateQueryRequestQuery{}.Type(ctx),
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
	var v CreateQueryRequestQuery
	d := o.Query.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuery sets the value of the Query field in CreateQueryRequest.
func (o *CreateQueryRequest) SetQuery(ctx context.Context, v CreateQueryRequestQuery) {
	vs := v.ToObjectValue(ctx)
	o.Query = vs
}

type CreateQueryRequestQuery struct {
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

func (to *CreateQueryRequestQuery) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryRequestQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateQueryRequestQuery) SyncFieldsDuringRead(ctx context.Context, from CreateQueryRequestQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c CreateQueryRequestQuery) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
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

func (to *CreateQueryVisualizationsLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryVisualizationsLegacyRequest) {
}

func (to *CreateQueryVisualizationsLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateQueryVisualizationsLegacyRequest) {
}

func (c CreateQueryVisualizationsLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["query_id"] = attrs["query_id"].SetRequired()
	attrs["type"] = attrs["type"].SetRequired()

	return attrs
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
	Visualization types.Object `tfsdk:"visualization"`
}

func (to *CreateVisualizationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVisualizationRequest) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				// Recursively sync the fields of Visualization
				toVisualization.SyncFieldsDuringCreateOrUpdate(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (to *CreateVisualizationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateVisualizationRequest) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				toVisualization.SyncFieldsDuringRead(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (c CreateVisualizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["visualization"] = attrs["visualization"].SetOptional()

	return attrs
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
			"visualization": CreateVisualizationRequestVisualization{}.Type(ctx),
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
	var v CreateVisualizationRequestVisualization
	d := o.Visualization.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualization sets the value of the Visualization field in CreateVisualizationRequest.
func (o *CreateVisualizationRequest) SetVisualization(ctx context.Context, v CreateVisualizationRequestVisualization) {
	vs := v.ToObjectValue(ctx)
	o.Visualization = vs
}

type CreateVisualizationRequestVisualization struct {
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

func (to *CreateVisualizationRequestVisualization) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVisualizationRequestVisualization) {
}

func (to *CreateVisualizationRequestVisualization) SyncFieldsDuringRead(ctx context.Context, from CreateVisualizationRequestVisualization) {
}

func (c CreateVisualizationRequestVisualization) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.Object `tfsdk:"channel"`
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
	Tags types.Object `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *CreateWarehouseRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWarehouseRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				// Recursively sync the fields of Tags
				toTags.SyncFieldsDuringCreateOrUpdate(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (to *CreateWarehouseRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWarehouseRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				toTags.SyncFieldsDuringRead(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (c CreateWarehouseRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["warehouse_type"] = attrs["warehouse_type"].SetOptional()

	return attrs
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
			"auto_stop_mins":            types.Int64Type,
			"channel":                   Channel{}.Type(ctx),
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"instance_profile_arn":      types.StringType,
			"max_num_clusters":          types.Int64Type,
			"min_num_clusters":          types.Int64Type,
			"name":                      types.StringType,
			"spot_instance_policy":      types.StringType,
			"tags":                      EndpointTags{}.Type(ctx),
			"warehouse_type":            types.StringType,
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in CreateWarehouseRequest.
func (o *CreateWarehouseRequest) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetTags returns the value of the Tags field in CreateWarehouseRequest as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWarehouseRequest) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v EndpointTags
	d := o.Tags.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateWarehouseRequest.
func (o *CreateWarehouseRequest) SetTags(ctx context.Context, v EndpointTags) {
	vs := v.ToObjectValue(ctx)
	o.Tags = vs
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id"`
}

func (to *CreateWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWarehouseResponse) {
}

func (to *CreateWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from CreateWarehouseResponse) {
}

func (c CreateWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	DashboardId types.String `tfsdk:"dashboard_id"`

	Options types.Object `tfsdk:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width"`
}

func (to *CreateWidget) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWidget) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *CreateWidget) SyncFieldsDuringRead(ctx context.Context, from CreateWidget) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (c CreateWidget) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["text"] = attrs["text"].SetOptional()
	attrs["visualization_id"] = attrs["visualization_id"].SetOptional()
	attrs["width"] = attrs["width"].SetRequired()

	return attrs
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
			"dashboard_id":     types.StringType,
			"options":          WidgetOptions{}.Type(ctx),
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
	var v WidgetOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in CreateWidget.
func (o *CreateWidget) SetOptions(ctx context.Context, v WidgetOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

type CronSchedule struct {
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

func (to *CronSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronSchedule) {
}

func (to *CronSchedule) SyncFieldsDuringRead(ctx context.Context, from CronSchedule) {
}

func (c CronSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"pause_status":         o.PauseStatus,
			"quartz_cron_schedule": o.QuartzCronSchedule,
			"timezone_id":          o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronSchedule) Type(ctx context.Context) attr.Type {
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
type Dashboard struct {
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

	Options types.Object `tfsdk:"options"`
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

	User types.Object `tfsdk:"user"`
	// The ID of the user who owns the dashboard.
	UserId types.Int64 `tfsdk:"user_id"`

	Widgets types.List `tfsdk:"widgets"`
}

func (to *Dashboard) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Dashboard) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				// Recursively sync the fields of User
				toUser.SyncFieldsDuringCreateOrUpdate(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
	if !from.Widgets.IsNull() && !from.Widgets.IsUnknown() && to.Widgets.IsNull() && len(from.Widgets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Widgets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Widgets = from.Widgets
	}
}

func (to *Dashboard) SyncFieldsDuringRead(ctx context.Context, from Dashboard) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
	if !from.Widgets.IsNull() && !from.Widgets.IsUnknown() && to.Widgets.IsNull() && len(from.Widgets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Widgets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Widgets = from.Widgets
	}
}

func (c Dashboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_edit"] = attrs["can_edit"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["dashboard_filters_enabled"] = attrs["dashboard_filters_enabled"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["is_archived"] = attrs["is_archived"].SetOptional()
	attrs["is_draft"] = attrs["is_draft"].SetOptional()
	attrs["is_favorite"] = attrs["is_favorite"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["permission_tier"] = attrs["permission_tier"].SetOptional()
	attrs["slug"] = attrs["slug"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()
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
			"options":                   DashboardOptions{}.Type(ctx),
			"parent":                    types.StringType,
			"permission_tier":           types.StringType,
			"slug":                      types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user":       User{}.Type(ctx),
			"user_id":    types.Int64Type,
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
	var v DashboardOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in Dashboard.
func (o *Dashboard) SetOptions(ctx context.Context, v DashboardOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	var v User
	d := o.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in Dashboard.
func (o *Dashboard) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	o.User = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Widgets = types.ListValueMust(t, vs)
}

type DashboardEditContent struct {
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

func (to *DashboardEditContent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardEditContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *DashboardEditContent) SyncFieldsDuringRead(ctx context.Context, from DashboardEditContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c DashboardEditContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["run_as_role"] = attrs["run_as_role"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()

	return attrs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`
}

func (to *DashboardOptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardOptions) {
}

func (to *DashboardOptions) SyncFieldsDuringRead(ctx context.Context, from DashboardOptions) {
}

func (c DashboardOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
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

func (to *DataSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataSource) {
}

func (to *DataSource) SyncFieldsDuringRead(ctx context.Context, from DataSource) {
}

func (c DataSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	End types.String `tfsdk:"end"`

	Start types.String `tfsdk:"start"`
}

func (to *DateRange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateRange) {
}

func (to *DateRange) SyncFieldsDuringRead(ctx context.Context, from DateRange) {
}

func (c DateRange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	DateRangeValue types.Object `tfsdk:"date_range_value"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue types.String `tfsdk:"dynamic_date_range_value"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision"`

	StartDayOfWeek types.Int64 `tfsdk:"start_day_of_week"`
}

func (to *DateRangeValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateRangeValue) {
	if !from.DateRangeValue.IsNull() && !from.DateRangeValue.IsUnknown() {
		if toDateRangeValue, ok := to.GetDateRangeValue(ctx); ok {
			if fromDateRangeValue, ok := from.GetDateRangeValue(ctx); ok {
				// Recursively sync the fields of DateRangeValue
				toDateRangeValue.SyncFieldsDuringCreateOrUpdate(ctx, fromDateRangeValue)
				to.SetDateRangeValue(ctx, toDateRangeValue)
			}
		}
	}
}

func (to *DateRangeValue) SyncFieldsDuringRead(ctx context.Context, from DateRangeValue) {
	if !from.DateRangeValue.IsNull() && !from.DateRangeValue.IsUnknown() {
		if toDateRangeValue, ok := to.GetDateRangeValue(ctx); ok {
			if fromDateRangeValue, ok := from.GetDateRangeValue(ctx); ok {
				toDateRangeValue.SyncFieldsDuringRead(ctx, fromDateRangeValue)
				to.SetDateRangeValue(ctx, toDateRangeValue)
			}
		}
	}
}

func (c DateRangeValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["date_range_value"] = attrs["date_range_value"].SetOptional()
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
			"date_range_value":         DateRange{}.Type(ctx),
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
	var v DateRange
	d := o.DateRangeValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDateRangeValue sets the value of the DateRangeValue field in DateRangeValue.
func (o *DateRangeValue) SetDateRangeValue(ctx context.Context, v DateRange) {
	vs := v.ToObjectValue(ctx)
	o.DateRangeValue = vs
}

type DateValue struct {
	// Manually specified date-time value.
	DateValue types.String `tfsdk:"date_value"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue types.String `tfsdk:"dynamic_date_value"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision"`
}

func (to *DateValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateValue) {
}

func (to *DateValue) SyncFieldsDuringRead(ctx context.Context, from DateValue) {
}

func (c DateValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (to *DeleteAlertsLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAlertsLegacyRequest) {
}

func (to *DeleteAlertsLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAlertsLegacyRequest) {
}

func (c DeleteAlertsLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_id"] = attrs["alert_id"].SetRequired()

	return attrs
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

type DeleteDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *DeleteDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardRequest) {
}

func (to *DeleteDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardRequest) {
}

func (c DeleteDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()

	return attrs
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

type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`
}

func (to *DeleteDashboardWidgetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardWidgetRequest) {
}

func (to *DeleteDashboardWidgetRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardWidgetRequest) {
}

func (c DeleteDashboardWidgetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

type DeleteQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *DeleteQueriesLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteQueriesLegacyRequest) {
}

func (to *DeleteQueriesLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteQueriesLegacyRequest) {
}

func (c DeleteQueriesLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_id"] = attrs["query_id"].SetRequired()

	return attrs
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

type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvisualizations/create
	Id types.String `tfsdk:"-"`
}

func (to *DeleteQueryVisualizationsLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteQueryVisualizationsLegacyRequest) {
}

func (to *DeleteQueryVisualizationsLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteQueryVisualizationsLegacyRequest) {
}

func (c DeleteQueryVisualizationsLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse) {
}

func (to *DeleteResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse) {
}

func (c DeleteResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteVisualizationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteVisualizationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteVisualizationRequest) {
}

func (to *DeleteVisualizationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteVisualizationRequest) {
}

func (c DeleteVisualizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *DeleteWarehouseRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWarehouseRequest) {
}

func (to *DeleteWarehouseRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWarehouseRequest) {
}

func (c DeleteWarehouseRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *DeleteWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWarehouseResponse) {
}

func (to *DeleteWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteWarehouseResponse) {
}

func (c DeleteWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.Object `tfsdk:"options"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

func (to *EditAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *EditAlert) SyncFieldsDuringRead(ctx context.Context, from EditAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (c EditAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["query_id"] = attrs["query_id"].SetRequired()
	attrs["rearm"] = attrs["rearm"].SetOptional()
	attrs["alert_id"] = attrs["alert_id"].SetRequired()

	return attrs
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
			"options":  AlertOptions{}.Type(ctx),
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
	var v AlertOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in EditAlert.
func (o *EditAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.Object `tfsdk:"channel"`
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
	Tags types.Object `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *EditWarehouseRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditWarehouseRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				// Recursively sync the fields of Tags
				toTags.SyncFieldsDuringCreateOrUpdate(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (to *EditWarehouseRequest) SyncFieldsDuringRead(ctx context.Context, from EditWarehouseRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				toTags.SyncFieldsDuringRead(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (c EditWarehouseRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["warehouse_type"] = attrs["warehouse_type"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
			"auto_stop_mins":            types.Int64Type,
			"channel":                   Channel{}.Type(ctx),
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
			"tags":                      EndpointTags{}.Type(ctx),
			"warehouse_type":            types.StringType,
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in EditWarehouseRequest.
func (o *EditWarehouseRequest) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetTags returns the value of the Tags field in EditWarehouseRequest as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditWarehouseRequest) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v EndpointTags
	d := o.Tags.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EditWarehouseRequest.
func (o *EditWarehouseRequest) SetTags(ctx context.Context, v EndpointTags) {
	vs := v.ToObjectValue(ctx)
	o.Tags = vs
}

type EditWarehouseResponse struct {
}

func (to *EditWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditWarehouseResponse) {
}

func (to *EditWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from EditWarehouseResponse) {
}

func (c EditWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (to *Empty) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Empty) {
}

func (to *Empty) SyncFieldsDuringRead(ctx context.Context, from Empty) {
}

func (c Empty) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *EndpointConfPair) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointConfPair) {
}

func (to *EndpointConfPair) SyncFieldsDuringRead(ctx context.Context, from EndpointConfPair) {
}

func (c EndpointConfPair) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Details types.String `tfsdk:"details"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason types.Object `tfsdk:"failure_reason"`
	// Deprecated. split into summary and details for security
	Message types.String `tfsdk:"message"`

	Status types.String `tfsdk:"status"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary"`
}

func (to *EndpointHealth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointHealth) {
	if !from.FailureReason.IsNull() && !from.FailureReason.IsUnknown() {
		if toFailureReason, ok := to.GetFailureReason(ctx); ok {
			if fromFailureReason, ok := from.GetFailureReason(ctx); ok {
				// Recursively sync the fields of FailureReason
				toFailureReason.SyncFieldsDuringCreateOrUpdate(ctx, fromFailureReason)
				to.SetFailureReason(ctx, toFailureReason)
			}
		}
	}
}

func (to *EndpointHealth) SyncFieldsDuringRead(ctx context.Context, from EndpointHealth) {
	if !from.FailureReason.IsNull() && !from.FailureReason.IsUnknown() {
		if toFailureReason, ok := to.GetFailureReason(ctx); ok {
			if fromFailureReason, ok := from.GetFailureReason(ctx); ok {
				toFailureReason.SyncFieldsDuringRead(ctx, fromFailureReason)
				to.SetFailureReason(ctx, toFailureReason)
			}
		}
	}
}

func (c EndpointHealth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["failure_reason"] = attrs["failure_reason"].SetOptional()
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
			"details":        types.StringType,
			"failure_reason": TerminationReason{}.Type(ctx),
			"message":        types.StringType,
			"status":         types.StringType,
			"summary":        types.StringType,
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
	var v TerminationReason
	d := o.FailureReason.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailureReason sets the value of the FailureReason field in EndpointHealth.
func (o *EndpointHealth) SetFailureReason(ctx context.Context, v TerminationReason) {
	vs := v.ToObjectValue(ctx)
	o.FailureReason = vs
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.Object `tfsdk:"channel"`
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
	Health types.Object `tfsdk:"health"`
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
	OdbcParams types.Object `tfsdk:"odbc_params"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`

	State types.String `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *EndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointInfo) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Health.IsNull() && !from.Health.IsUnknown() {
		if toHealth, ok := to.GetHealth(ctx); ok {
			if fromHealth, ok := from.GetHealth(ctx); ok {
				// Recursively sync the fields of Health
				toHealth.SyncFieldsDuringCreateOrUpdate(ctx, fromHealth)
				to.SetHealth(ctx, toHealth)
			}
		}
	}
	if !from.OdbcParams.IsNull() && !from.OdbcParams.IsUnknown() {
		if toOdbcParams, ok := to.GetOdbcParams(ctx); ok {
			if fromOdbcParams, ok := from.GetOdbcParams(ctx); ok {
				// Recursively sync the fields of OdbcParams
				toOdbcParams.SyncFieldsDuringCreateOrUpdate(ctx, fromOdbcParams)
				to.SetOdbcParams(ctx, toOdbcParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				// Recursively sync the fields of Tags
				toTags.SyncFieldsDuringCreateOrUpdate(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (to *EndpointInfo) SyncFieldsDuringRead(ctx context.Context, from EndpointInfo) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Health.IsNull() && !from.Health.IsUnknown() {
		if toHealth, ok := to.GetHealth(ctx); ok {
			if fromHealth, ok := from.GetHealth(ctx); ok {
				toHealth.SyncFieldsDuringRead(ctx, fromHealth)
				to.SetHealth(ctx, toHealth)
			}
		}
	}
	if !from.OdbcParams.IsNull() && !from.OdbcParams.IsUnknown() {
		if toOdbcParams, ok := to.GetOdbcParams(ctx); ok {
			if fromOdbcParams, ok := from.GetOdbcParams(ctx); ok {
				toOdbcParams.SyncFieldsDuringRead(ctx, fromOdbcParams)
				to.SetOdbcParams(ctx, toOdbcParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				toTags.SyncFieldsDuringRead(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (c EndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["jdbc_url"] = attrs["jdbc_url"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["num_active_sessions"] = attrs["num_active_sessions"].SetOptional()
	attrs["num_clusters"] = attrs["num_clusters"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].SetOptional()
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
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
			"auto_stop_mins":            types.Int64Type,
			"channel":                   Channel{}.Type(ctx),
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health":                    EndpointHealth{}.Type(ctx),
			"id":                        types.StringType,
			"instance_profile_arn":      types.StringType,
			"jdbc_url":                  types.StringType,
			"max_num_clusters":          types.Int64Type,
			"min_num_clusters":          types.Int64Type,
			"name":                      types.StringType,
			"num_active_sessions":       types.Int64Type,
			"num_clusters":              types.Int64Type,
			"odbc_params":               OdbcParams{}.Type(ctx),
			"spot_instance_policy":      types.StringType,
			"state":                     types.StringType,
			"tags":                      EndpointTags{}.Type(ctx),
			"warehouse_type":            types.StringType,
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in EndpointInfo.
func (o *EndpointInfo) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetHealth returns the value of the Health field in EndpointInfo as
// a EndpointHealth value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetHealth(ctx context.Context) (EndpointHealth, bool) {
	var e EndpointHealth
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v EndpointHealth
	d := o.Health.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHealth sets the value of the Health field in EndpointInfo.
func (o *EndpointInfo) SetHealth(ctx context.Context, v EndpointHealth) {
	vs := v.ToObjectValue(ctx)
	o.Health = vs
}

// GetOdbcParams returns the value of the OdbcParams field in EndpointInfo as
// a OdbcParams value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetOdbcParams(ctx context.Context) (OdbcParams, bool) {
	var e OdbcParams
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v OdbcParams
	d := o.OdbcParams.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOdbcParams sets the value of the OdbcParams field in EndpointInfo.
func (o *EndpointInfo) SetOdbcParams(ctx context.Context, v OdbcParams) {
	vs := v.ToObjectValue(ctx)
	o.OdbcParams = vs
}

// GetTags returns the value of the Tags field in EndpointInfo as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v EndpointTags
	d := o.Tags.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EndpointInfo.
func (o *EndpointInfo) SetTags(ctx context.Context, v EndpointTags) {
	vs := v.ToObjectValue(ctx)
	o.Tags = vs
}

type EndpointTagPair struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *EndpointTagPair) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTagPair) {
}

func (to *EndpointTagPair) SyncFieldsDuringRead(ctx context.Context, from EndpointTagPair) {
}

func (c EndpointTagPair) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	CustomTags types.List `tfsdk:"custom_tags"`
}

func (to *EndpointTags) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTags) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (to *EndpointTags) SyncFieldsDuringRead(ctx context.Context, from EndpointTags) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (c EndpointTags) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.ListValueMust(t, vs)
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions types.String `tfsdk:"enum_options"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.Object `tfsdk:"multi_values_options"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (to *EnumValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnumValue) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				// Recursively sync the fields of MultiValuesOptions
				toMultiValuesOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *EnumValue) SyncFieldsDuringRead(ctx context.Context, from EnumValue) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				toMultiValuesOptions.SyncFieldsDuringRead(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (c EnumValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enum_options"] = attrs["enum_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
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
			"enum_options":         types.StringType,
			"multi_values_options": MultiValuesOptions{}.Type(ctx),
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
	var v MultiValuesOptions
	d := o.MultiValuesOptions.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in EnumValue.
func (o *EnumValue) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := v.ToObjectValue(ctx)
	o.MultiValuesOptions = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type ExecuteStatementRequest struct {
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

func (to *ExecuteStatementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteStatementRequest) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (to *ExecuteStatementRequest) SyncFieldsDuringRead(ctx context.Context, from ExecuteStatementRequest) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (c ExecuteStatementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_limit"] = attrs["byte_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["disposition"] = attrs["disposition"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["on_wait_timeout"] = attrs["on_wait_timeout"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["row_limit"] = attrs["row_limit"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["statement"] = attrs["statement"].SetRequired()
	attrs["wait_timeout"] = attrs["wait_timeout"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type ExternalLink struct {
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

func (to *ExternalLink) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalLink) {
}

func (to *ExternalLink) SyncFieldsDuringRead(ctx context.Context, from ExternalLink) {
}

func (c ExternalLink) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.HttpHeaders = types.MapValueMust(t, vs)
}

type ExternalQuerySource struct {
	// The canonical identifier for this SQL alert
	AlertId types.String `tfsdk:"alert_id"`
	// The canonical identifier for this Lakeview dashboard
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The canonical identifier for this Genie space
	GenieSpaceId types.String `tfsdk:"genie_space_id"`

	JobInfo types.Object `tfsdk:"job_info"`
	// The canonical identifier for this legacy dashboard
	LegacyDashboardId types.String `tfsdk:"legacy_dashboard_id"`
	// The canonical identifier for this notebook
	NotebookId types.String `tfsdk:"notebook_id"`
	// The canonical identifier for this SQL query
	SqlQueryId types.String `tfsdk:"sql_query_id"`
}

func (to *ExternalQuerySource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalQuerySource) {
	if !from.JobInfo.IsNull() && !from.JobInfo.IsUnknown() {
		if toJobInfo, ok := to.GetJobInfo(ctx); ok {
			if fromJobInfo, ok := from.GetJobInfo(ctx); ok {
				// Recursively sync the fields of JobInfo
				toJobInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromJobInfo)
				to.SetJobInfo(ctx, toJobInfo)
			}
		}
	}
}

func (to *ExternalQuerySource) SyncFieldsDuringRead(ctx context.Context, from ExternalQuerySource) {
	if !from.JobInfo.IsNull() && !from.JobInfo.IsUnknown() {
		if toJobInfo, ok := to.GetJobInfo(ctx); ok {
			if fromJobInfo, ok := from.GetJobInfo(ctx); ok {
				toJobInfo.SyncFieldsDuringRead(ctx, fromJobInfo)
				to.SetJobInfo(ctx, toJobInfo)
			}
		}
	}
}

func (c ExternalQuerySource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_id"] = attrs["alert_id"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["genie_space_id"] = attrs["genie_space_id"].SetOptional()
	attrs["job_info"] = attrs["job_info"].SetOptional()
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
func (a ExternalQuerySource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_info": reflect.TypeOf(ExternalQuerySourceJobInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySource
// only implements ToObjectValue() and Type().
func (o ExternalQuerySource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExternalQuerySource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id":            types.StringType,
			"dashboard_id":        types.StringType,
			"genie_space_id":      types.StringType,
			"job_info":            ExternalQuerySourceJobInfo{}.Type(ctx),
			"legacy_dashboard_id": types.StringType,
			"notebook_id":         types.StringType,
			"sql_query_id":        types.StringType,
		},
	}
}

// GetJobInfo returns the value of the JobInfo field in ExternalQuerySource as
// a ExternalQuerySourceJobInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalQuerySource) GetJobInfo(ctx context.Context) (ExternalQuerySourceJobInfo, bool) {
	var e ExternalQuerySourceJobInfo
	if o.JobInfo.IsNull() || o.JobInfo.IsUnknown() {
		return e, false
	}
	var v ExternalQuerySourceJobInfo
	d := o.JobInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobInfo sets the value of the JobInfo field in ExternalQuerySource.
func (o *ExternalQuerySource) SetJobInfo(ctx context.Context, v ExternalQuerySourceJobInfo) {
	vs := v.ToObjectValue(ctx)
	o.JobInfo = vs
}

type ExternalQuerySourceJobInfo struct {
	// The canonical identifier for this job.
	JobId types.String `tfsdk:"job_id"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	JobRunId types.String `tfsdk:"job_run_id"`
	// The canonical identifier of the task run.
	JobTaskRunId types.String `tfsdk:"job_task_run_id"`
}

func (to *ExternalQuerySourceJobInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalQuerySourceJobInfo) {
}

func (to *ExternalQuerySourceJobInfo) SyncFieldsDuringRead(ctx context.Context, from ExternalQuerySourceJobInfo) {
}

func (c ExternalQuerySourceJobInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExternalQuerySourceJobInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySourceJobInfo
// only implements ToObjectValue() and Type().
func (o ExternalQuerySourceJobInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":          o.JobId,
			"job_run_id":      o.JobRunId,
			"job_task_run_id": o.JobTaskRunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalQuerySourceJobInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":          types.StringType,
			"job_run_id":      types.StringType,
			"job_task_run_id": types.StringType,
		},
	}
}

type GetAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetAlertRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertRequest) {
}

func (to *GetAlertRequest) SyncFieldsDuringRead(ctx context.Context, from GetAlertRequest) {
}

func (c GetAlertRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

type GetAlertV2Request struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetAlertV2Request) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertV2Request) {
}

func (to *GetAlertV2Request) SyncFieldsDuringRead(ctx context.Context, from GetAlertV2Request) {
}

func (c GetAlertV2Request) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAlertV2Request) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertV2Request
// only implements ToObjectValue() and Type().
func (o GetAlertV2Request) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAlertV2Request) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (to *GetAlertsLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertsLegacyRequest) {
}

func (to *GetAlertsLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from GetAlertsLegacyRequest) {
}

func (c GetAlertsLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_id"] = attrs["alert_id"].SetRequired()

	return attrs
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

type GetConfigRequest struct {
}

func (to *GetConfigRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetConfigRequest) {
}

func (to *GetConfigRequest) SyncFieldsDuringRead(ctx context.Context, from GetConfigRequest) {
}

func (c GetConfigRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetConfigRequest
// only implements ToObjectValue() and Type().
func (o GetConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *GetDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDashboardRequest) {
}

func (to *GetDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from GetDashboardRequest) {
}

func (c GetDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()

	return attrs
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

type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permissions to check.
	ObjectType types.String `tfsdk:"-"`
}

func (to *GetDbsqlPermissionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDbsqlPermissionRequest) {
}

func (to *GetDbsqlPermissionRequest) SyncFieldsDuringRead(ctx context.Context, from GetDbsqlPermissionRequest) {
}

func (c GetDbsqlPermissionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
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
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDbsqlPermissionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type GetQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *GetQueriesLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetQueriesLegacyRequest) {
}

func (to *GetQueriesLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from GetQueriesLegacyRequest) {
}

func (c GetQueriesLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_id"] = attrs["query_id"].SetRequired()

	return attrs
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

type GetQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetQueryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetQueryRequest) {
}

func (to *GetQueryRequest) SyncFieldsDuringRead(ctx context.Context, from GetQueryRequest) {
}

func (c GetQueryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type"`
}

func (to *GetResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetResponse) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *GetResponse) SyncFieldsDuringRead(ctx context.Context, from GetResponse) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c GetResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (to *GetStatementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementRequest) {
}

func (to *GetStatementRequest) SyncFieldsDuringRead(ctx context.Context, from GetStatementRequest) {
}

func (c GetStatementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_id"] = attrs["statement_id"].SetRequired()

	return attrs
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

type GetStatementResultChunkNRequest struct {
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (to *GetStatementResultChunkNRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementResultChunkNRequest) {
}

func (to *GetStatementResultChunkNRequest) SyncFieldsDuringRead(ctx context.Context, from GetStatementResultChunkNRequest) {
}

func (c GetStatementResultChunkNRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["statement_id"] = attrs["statement_id"].SetRequired()
	attrs["chunk_index"] = attrs["chunk_index"].SetRequired()

	return attrs
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

type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetWarehousePermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionLevelsRequest) {
}

func (to *GetWarehousePermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionLevelsRequest) {
}

func (c GetWarehousePermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
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
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetWarehousePermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetWarehousePermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (c GetWarehousePermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetWarehousePermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionsRequest) {
}

func (to *GetWarehousePermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionsRequest) {
}

func (c GetWarehousePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
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

type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *GetWarehouseRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehouseRequest) {
}

func (to *GetWarehouseRequest) SyncFieldsDuringRead(ctx context.Context, from GetWarehouseRequest) {
}

func (c GetWarehouseRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel types.Object `tfsdk:"channel"`
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
	Health types.Object `tfsdk:"health"`
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
	OdbcParams types.Object `tfsdk:"odbc_params"`

	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`

	State types.String `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *GetWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehouseResponse) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Health.IsNull() && !from.Health.IsUnknown() {
		if toHealth, ok := to.GetHealth(ctx); ok {
			if fromHealth, ok := from.GetHealth(ctx); ok {
				// Recursively sync the fields of Health
				toHealth.SyncFieldsDuringCreateOrUpdate(ctx, fromHealth)
				to.SetHealth(ctx, toHealth)
			}
		}
	}
	if !from.OdbcParams.IsNull() && !from.OdbcParams.IsUnknown() {
		if toOdbcParams, ok := to.GetOdbcParams(ctx); ok {
			if fromOdbcParams, ok := from.GetOdbcParams(ctx); ok {
				// Recursively sync the fields of OdbcParams
				toOdbcParams.SyncFieldsDuringCreateOrUpdate(ctx, fromOdbcParams)
				to.SetOdbcParams(ctx, toOdbcParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				// Recursively sync the fields of Tags
				toTags.SyncFieldsDuringCreateOrUpdate(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (to *GetWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from GetWarehouseResponse) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.Health.IsNull() && !from.Health.IsUnknown() {
		if toHealth, ok := to.GetHealth(ctx); ok {
			if fromHealth, ok := from.GetHealth(ctx); ok {
				toHealth.SyncFieldsDuringRead(ctx, fromHealth)
				to.SetHealth(ctx, toHealth)
			}
		}
	}
	if !from.OdbcParams.IsNull() && !from.OdbcParams.IsUnknown() {
		if toOdbcParams, ok := to.GetOdbcParams(ctx); ok {
			if fromOdbcParams, ok := from.GetOdbcParams(ctx); ok {
				toOdbcParams.SyncFieldsDuringRead(ctx, fromOdbcParams)
				to.SetOdbcParams(ctx, toOdbcParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() {
		if toTags, ok := to.GetTags(ctx); ok {
			if fromTags, ok := from.GetTags(ctx); ok {
				toTags.SyncFieldsDuringRead(ctx, fromTags)
				to.SetTags(ctx, toTags)
			}
		}
	}
}

func (c GetWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["jdbc_url"] = attrs["jdbc_url"].SetOptional()
	attrs["max_num_clusters"] = attrs["max_num_clusters"].SetOptional()
	attrs["min_num_clusters"] = attrs["min_num_clusters"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["num_active_sessions"] = attrs["num_active_sessions"].SetOptional()
	attrs["num_clusters"] = attrs["num_clusters"].SetOptional()
	attrs["odbc_params"] = attrs["odbc_params"].SetOptional()
	attrs["spot_instance_policy"] = attrs["spot_instance_policy"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
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
			"auto_stop_mins":            types.Int64Type,
			"channel":                   Channel{}.Type(ctx),
			"cluster_size":              types.StringType,
			"creator_name":              types.StringType,
			"enable_photon":             types.BoolType,
			"enable_serverless_compute": types.BoolType,
			"health":                    EndpointHealth{}.Type(ctx),
			"id":                        types.StringType,
			"instance_profile_arn":      types.StringType,
			"jdbc_url":                  types.StringType,
			"max_num_clusters":          types.Int64Type,
			"min_num_clusters":          types.Int64Type,
			"name":                      types.StringType,
			"num_active_sessions":       types.Int64Type,
			"num_clusters":              types.Int64Type,
			"odbc_params":               OdbcParams{}.Type(ctx),
			"spot_instance_policy":      types.StringType,
			"state":                     types.StringType,
			"tags":                      EndpointTags{}.Type(ctx),
			"warehouse_type":            types.StringType,
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetHealth returns the value of the Health field in GetWarehouseResponse as
// a EndpointHealth value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetHealth(ctx context.Context) (EndpointHealth, bool) {
	var e EndpointHealth
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v EndpointHealth
	d := o.Health.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHealth sets the value of the Health field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetHealth(ctx context.Context, v EndpointHealth) {
	vs := v.ToObjectValue(ctx)
	o.Health = vs
}

// GetOdbcParams returns the value of the OdbcParams field in GetWarehouseResponse as
// a OdbcParams value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetOdbcParams(ctx context.Context) (OdbcParams, bool) {
	var e OdbcParams
	if o.OdbcParams.IsNull() || o.OdbcParams.IsUnknown() {
		return e, false
	}
	var v OdbcParams
	d := o.OdbcParams.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOdbcParams sets the value of the OdbcParams field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetOdbcParams(ctx context.Context, v OdbcParams) {
	vs := v.ToObjectValue(ctx)
	o.OdbcParams = vs
}

// GetTags returns the value of the Tags field in GetWarehouseResponse as
// a EndpointTags value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWarehouseResponse) GetTags(ctx context.Context) (EndpointTags, bool) {
	var e EndpointTags
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return e, false
	}
	var v EndpointTags
	d := o.Tags.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in GetWarehouseResponse.
func (o *GetWarehouseResponse) SetTags(ctx context.Context, v EndpointTags) {
	vs := v.ToObjectValue(ctx)
	o.Tags = vs
}

type GetWorkspaceWarehouseConfigRequest struct {
}

func (to *GetWorkspaceWarehouseConfigRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceWarehouseConfigRequest) {
}

func (to *GetWorkspaceWarehouseConfigRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceWarehouseConfigRequest) {
}

func (c GetWorkspaceWarehouseConfigRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceWarehouseConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceWarehouseConfigRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceWarehouseConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceWarehouseConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel types.Object `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.Object `tfsdk:"config_param"`
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
	GlobalParam types.Object `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.Object `tfsdk:"sql_configuration_parameters"`
}

func (to *GetWorkspaceWarehouseConfigResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceWarehouseConfigResponse) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.ConfigParam.IsNull() && !from.ConfigParam.IsUnknown() {
		if toConfigParam, ok := to.GetConfigParam(ctx); ok {
			if fromConfigParam, ok := from.GetConfigParam(ctx); ok {
				// Recursively sync the fields of ConfigParam
				toConfigParam.SyncFieldsDuringCreateOrUpdate(ctx, fromConfigParam)
				to.SetConfigParam(ctx, toConfigParam)
			}
		}
	}
	if !from.DataAccessConfig.IsNull() && !from.DataAccessConfig.IsUnknown() && to.DataAccessConfig.IsNull() && len(from.DataAccessConfig.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataAccessConfig, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataAccessConfig = from.DataAccessConfig
	}
	if !from.EnabledWarehouseTypes.IsNull() && !from.EnabledWarehouseTypes.IsUnknown() && to.EnabledWarehouseTypes.IsNull() && len(from.EnabledWarehouseTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnabledWarehouseTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnabledWarehouseTypes = from.EnabledWarehouseTypes
	}
	if !from.GlobalParam.IsNull() && !from.GlobalParam.IsUnknown() {
		if toGlobalParam, ok := to.GetGlobalParam(ctx); ok {
			if fromGlobalParam, ok := from.GetGlobalParam(ctx); ok {
				// Recursively sync the fields of GlobalParam
				toGlobalParam.SyncFieldsDuringCreateOrUpdate(ctx, fromGlobalParam)
				to.SetGlobalParam(ctx, toGlobalParam)
			}
		}
	}
	if !from.SqlConfigurationParameters.IsNull() && !from.SqlConfigurationParameters.IsUnknown() {
		if toSqlConfigurationParameters, ok := to.GetSqlConfigurationParameters(ctx); ok {
			if fromSqlConfigurationParameters, ok := from.GetSqlConfigurationParameters(ctx); ok {
				// Recursively sync the fields of SqlConfigurationParameters
				toSqlConfigurationParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromSqlConfigurationParameters)
				to.SetSqlConfigurationParameters(ctx, toSqlConfigurationParameters)
			}
		}
	}
}

func (to *GetWorkspaceWarehouseConfigResponse) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceWarehouseConfigResponse) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.ConfigParam.IsNull() && !from.ConfigParam.IsUnknown() {
		if toConfigParam, ok := to.GetConfigParam(ctx); ok {
			if fromConfigParam, ok := from.GetConfigParam(ctx); ok {
				toConfigParam.SyncFieldsDuringRead(ctx, fromConfigParam)
				to.SetConfigParam(ctx, toConfigParam)
			}
		}
	}
	if !from.DataAccessConfig.IsNull() && !from.DataAccessConfig.IsUnknown() && to.DataAccessConfig.IsNull() && len(from.DataAccessConfig.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataAccessConfig, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataAccessConfig = from.DataAccessConfig
	}
	if !from.EnabledWarehouseTypes.IsNull() && !from.EnabledWarehouseTypes.IsUnknown() && to.EnabledWarehouseTypes.IsNull() && len(from.EnabledWarehouseTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnabledWarehouseTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnabledWarehouseTypes = from.EnabledWarehouseTypes
	}
	if !from.GlobalParam.IsNull() && !from.GlobalParam.IsUnknown() {
		if toGlobalParam, ok := to.GetGlobalParam(ctx); ok {
			if fromGlobalParam, ok := from.GetGlobalParam(ctx); ok {
				toGlobalParam.SyncFieldsDuringRead(ctx, fromGlobalParam)
				to.SetGlobalParam(ctx, toGlobalParam)
			}
		}
	}
	if !from.SqlConfigurationParameters.IsNull() && !from.SqlConfigurationParameters.IsUnknown() {
		if toSqlConfigurationParameters, ok := to.GetSqlConfigurationParameters(ctx); ok {
			if fromSqlConfigurationParameters, ok := from.GetSqlConfigurationParameters(ctx); ok {
				toSqlConfigurationParameters.SyncFieldsDuringRead(ctx, fromSqlConfigurationParameters)
				to.SetSqlConfigurationParameters(ctx, toSqlConfigurationParameters)
			}
		}
	}
}

func (c GetWorkspaceWarehouseConfigResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["config_param"] = attrs["config_param"].SetOptional()
	attrs["data_access_config"] = attrs["data_access_config"].SetOptional()
	attrs["enabled_warehouse_types"] = attrs["enabled_warehouse_types"].SetOptional()
	attrs["global_param"] = attrs["global_param"].SetOptional()
	attrs["google_service_account"] = attrs["google_service_account"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["security_policy"] = attrs["security_policy"].SetOptional()
	attrs["sql_configuration_parameters"] = attrs["sql_configuration_parameters"].SetOptional()

	return attrs
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
			"channel":      Channel{}.Type(ctx),
			"config_param": RepeatedEndpointConfPairs{}.Type(ctx),
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair{}.Type(ctx),
			},
			"global_param":                 RepeatedEndpointConfPairs{}.Type(ctx),
			"google_service_account":       types.StringType,
			"instance_profile_arn":         types.StringType,
			"security_policy":              types.StringType,
			"sql_configuration_parameters": RepeatedEndpointConfPairs{}.Type(ctx),
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetConfigParam returns the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v RepeatedEndpointConfPairs
	d := o.ConfigParam.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigParam sets the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.ConfigParam = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	var v RepeatedEndpointConfPairs
	d := o.GlobalParam.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlobalParam sets the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.GlobalParam = vs
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceWarehouseConfigResponse) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v RepeatedEndpointConfPairs
	d := o.SqlConfigurationParameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse.
func (o *GetWorkspaceWarehouseConfigResponse) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.SqlConfigurationParameters = vs
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Alert ID.
	Id types.String `tfsdk:"id"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt types.String `tfsdk:"last_triggered_at"`
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options types.Object `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`

	Query types.Object `tfsdk:"query"`
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

	User types.Object `tfsdk:"user"`
}

func (to *LegacyAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				// Recursively sync the fields of Query
				toQuery.SyncFieldsDuringCreateOrUpdate(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				// Recursively sync the fields of User
				toUser.SyncFieldsDuringCreateOrUpdate(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (to *LegacyAlert) SyncFieldsDuringRead(ctx context.Context, from LegacyAlert) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (c LegacyAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_triggered_at"] = attrs["last_triggered_at"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["rearm"] = attrs["rearm"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()

	return attrs
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
			"options":           AlertOptions{}.Type(ctx),
			"parent":            types.StringType,
			"query":             AlertQuery{}.Type(ctx),
			"rearm":             types.Int64Type,
			"state":             types.StringType,
			"updated_at":        types.StringType,
			"user":              User{}.Type(ctx),
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
	var v AlertOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in LegacyAlert.
func (o *LegacyAlert) SetOptions(ctx context.Context, v AlertOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

// GetQuery returns the value of the Query field in LegacyAlert as
// a AlertQuery value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert) GetQuery(ctx context.Context) (AlertQuery, bool) {
	var e AlertQuery
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v AlertQuery
	d := o.Query.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuery sets the value of the Query field in LegacyAlert.
func (o *LegacyAlert) SetQuery(ctx context.Context, v AlertQuery) {
	vs := v.ToObjectValue(ctx)
	o.Query = vs
}

// GetUser returns the value of the User field in LegacyAlert as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyAlert) GetUser(ctx context.Context) (User, bool) {
	var e User
	if o.User.IsNull() || o.User.IsUnknown() {
		return e, false
	}
	var v User
	d := o.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in LegacyAlert.
func (o *LegacyAlert) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	o.User = vs
}

type LegacyQuery struct {
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

	LastModifiedBy types.Object `tfsdk:"last_modified_by"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById types.Int64 `tfsdk:"last_modified_by_id"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId types.String `tfsdk:"latest_query_data_id"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`

	Options types.Object `tfsdk:"options"`
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

	User types.Object `tfsdk:"user"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`

	Visualizations types.List `tfsdk:"visualizations"`
}

func (to *LegacyQuery) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyQuery) {
	if !from.LastModifiedBy.IsNull() && !from.LastModifiedBy.IsUnknown() {
		if toLastModifiedBy, ok := to.GetLastModifiedBy(ctx); ok {
			if fromLastModifiedBy, ok := from.GetLastModifiedBy(ctx); ok {
				// Recursively sync the fields of LastModifiedBy
				toLastModifiedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromLastModifiedBy)
				to.SetLastModifiedBy(ctx, toLastModifiedBy)
			}
		}
	}
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				// Recursively sync the fields of User
				toUser.SyncFieldsDuringCreateOrUpdate(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
	if !from.Visualizations.IsNull() && !from.Visualizations.IsUnknown() && to.Visualizations.IsNull() && len(from.Visualizations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Visualizations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Visualizations = from.Visualizations
	}
}

func (to *LegacyQuery) SyncFieldsDuringRead(ctx context.Context, from LegacyQuery) {
	if !from.LastModifiedBy.IsNull() && !from.LastModifiedBy.IsUnknown() {
		if toLastModifiedBy, ok := to.GetLastModifiedBy(ctx); ok {
			if fromLastModifiedBy, ok := from.GetLastModifiedBy(ctx); ok {
				toLastModifiedBy.SyncFieldsDuringRead(ctx, fromLastModifiedBy)
				to.SetLastModifiedBy(ctx, toLastModifiedBy)
			}
		}
	}
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
	if !from.Visualizations.IsNull() && !from.Visualizations.IsUnknown() && to.Visualizations.IsNull() && len(from.Visualizations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Visualizations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Visualizations = from.Visualizations
	}
}

func (c LegacyQuery) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["last_modified_by_id"] = attrs["last_modified_by_id"].SetOptional()
	attrs["latest_query_data_id"] = attrs["latest_query_data_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["permission_tier"] = attrs["permission_tier"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query_hash"] = attrs["query_hash"].SetOptional()
	attrs["run_as_role"] = attrs["run_as_role"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["user"] = attrs["user"].SetOptional()
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
			"can_edit":             types.BoolType,
			"created_at":           types.StringType,
			"data_source_id":       types.StringType,
			"description":          types.StringType,
			"id":                   types.StringType,
			"is_archived":          types.BoolType,
			"is_draft":             types.BoolType,
			"is_favorite":          types.BoolType,
			"is_safe":              types.BoolType,
			"last_modified_by":     User{}.Type(ctx),
			"last_modified_by_id":  types.Int64Type,
			"latest_query_data_id": types.StringType,
			"name":                 types.StringType,
			"options":              QueryOptions{}.Type(ctx),
			"parent":               types.StringType,
			"permission_tier":      types.StringType,
			"query":                types.StringType,
			"query_hash":           types.StringType,
			"run_as_role":          types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"updated_at": types.StringType,
			"user":       User{}.Type(ctx),
			"user_id":    types.Int64Type,
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
	var v User
	d := o.LastModifiedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLastModifiedBy sets the value of the LastModifiedBy field in LegacyQuery.
func (o *LegacyQuery) SetLastModifiedBy(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	o.LastModifiedBy = vs
}

// GetOptions returns the value of the Options field in LegacyQuery as
// a QueryOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *LegacyQuery) GetOptions(ctx context.Context) (QueryOptions, bool) {
	var e QueryOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v QueryOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in LegacyQuery.
func (o *LegacyQuery) SetOptions(ctx context.Context, v QueryOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	var v User
	d := o.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in LegacyQuery.
func (o *LegacyQuery) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	o.User = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Visualizations = types.ListValueMust(t, vs)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {
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

	Query types.Object `tfsdk:"query"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type_ types.String `tfsdk:"type"`

	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (to *LegacyVisualization) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyVisualization) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				// Recursively sync the fields of Query
				toQuery.SyncFieldsDuringCreateOrUpdate(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (to *LegacyVisualization) SyncFieldsDuringRead(ctx context.Context, from LegacyVisualization) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (c LegacyVisualization) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
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
			"query":       LegacyQuery{}.Type(ctx),
			"type":        types.StringType,
			"updated_at":  types.StringType,
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
	var v LegacyQuery
	d := o.Query.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuery sets the value of the Query field in LegacyVisualization.
func (o *LegacyVisualization) SetQuery(ctx context.Context, v LegacyQuery) {
	vs := v.ToObjectValue(ctx)
	o.Query = vs
}

type ListAlertsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAlertsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsRequest) {
}

func (to *ListAlertsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAlertsRequest) {
}

func (c ListAlertsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListAlertsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListAlertsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAlertsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c ListAlertsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition"`
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

func (to *ListAlertsResponseAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsResponseAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				// Recursively sync the fields of Condition
				toCondition.SyncFieldsDuringCreateOrUpdate(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (to *ListAlertsResponseAlert) SyncFieldsDuringRead(ctx context.Context, from ListAlertsResponseAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (c ListAlertsResponseAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
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
			"condition":            AlertCondition{}.Type(ctx),
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
	var v AlertCondition
	d := o.Condition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCondition sets the value of the Condition field in ListAlertsResponseAlert.
func (o *ListAlertsResponseAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := v.ToObjectValue(ctx)
	o.Condition = vs
}

type ListAlertsV2Request struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAlertsV2Request) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsV2Request) {
}

func (to *ListAlertsV2Request) SyncFieldsDuringRead(ctx context.Context, from ListAlertsV2Request) {
}

func (c ListAlertsV2Request) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAlertsV2Request) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Request
// only implements ToObjectValue() and Type().
func (o ListAlertsV2Request) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsV2Request) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAlertsV2Response struct {
	Alerts types.List `tfsdk:"alerts"`

	NextPageToken types.String `tfsdk:"next_page_token"`
	// Deprecated. Use `alerts` instead.
	Results types.List `tfsdk:"results"`
}

func (to *ListAlertsV2Response) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsV2Response) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListAlertsV2Response) SyncFieldsDuringRead(ctx context.Context, from ListAlertsV2Response) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c ListAlertsV2Response) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alerts"] = attrs["alerts"].SetOptional()
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
func (a ListAlertsV2Response) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":  reflect.TypeOf(AlertV2{}),
		"results": reflect.TypeOf(AlertV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Response
// only implements ToObjectValue() and Type().
func (o ListAlertsV2Response) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":          o.Alerts,
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAlertsV2Response) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alerts": basetypes.ListType{
				ElemType: AlertV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: AlertV2{}.Type(ctx),
			},
		},
	}
}

// GetAlerts returns the value of the Alerts field in ListAlertsV2Response as
// a slice of AlertV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsV2Response) GetAlerts(ctx context.Context) ([]AlertV2, bool) {
	if o.Alerts.IsNull() || o.Alerts.IsUnknown() {
		return nil, false
	}
	var v []AlertV2
	d := o.Alerts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlerts sets the value of the Alerts field in ListAlertsV2Response.
func (o *ListAlertsV2Response) SetAlerts(ctx context.Context, v []AlertV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Alerts = types.ListValueMust(t, vs)
}

// GetResults returns the value of the Results field in ListAlertsV2Response as
// a slice of AlertV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAlertsV2Response) GetResults(ctx context.Context) ([]AlertV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []AlertV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListAlertsV2Response.
func (o *ListAlertsV2Response) SetResults(ctx context.Context, v []AlertV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

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

func (to *ListDashboardsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDashboardsRequest) {
}

func (to *ListDashboardsRequest) SyncFieldsDuringRead(ctx context.Context, from ListDashboardsRequest) {
}

func (c ListDashboardsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["order"] = attrs["order"].SetOptional()
	attrs["page"] = attrs["page"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["q"] = attrs["q"].SetOptional()

	return attrs
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

func (to *ListQueriesLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesLegacyRequest) {
}

func (to *ListQueriesLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from ListQueriesLegacyRequest) {
}

func (c ListQueriesLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["order"] = attrs["order"].SetOptional()
	attrs["page"] = attrs["page"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["q"] = attrs["q"].SetOptional()

	return attrs
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

type ListQueriesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListQueriesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesRequest) {
}

func (to *ListQueriesRequest) SyncFieldsDuringRead(ctx context.Context, from ListQueriesRequest) {
}

func (c ListQueriesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	HasNextPage types.Bool `tfsdk:"has_next_page"`
	// A token that can be used to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Res types.List `tfsdk:"res"`
}

func (to *ListQueriesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesResponse) {
	if !from.Res.IsNull() && !from.Res.IsUnknown() && to.Res.IsNull() && len(from.Res.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Res, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Res = from.Res
	}
}

func (to *ListQueriesResponse) SyncFieldsDuringRead(ctx context.Context, from ListQueriesResponse) {
	if !from.Res.IsNull() && !from.Res.IsUnknown() && to.Res.IsNull() && len(from.Res.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Res, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Res = from.Res
	}
}

func (c ListQueriesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Res = types.ListValueMust(t, vs)
}

type ListQueryHistoryRequest struct {
	// An optional filter object to limit query history results. Accepts
	// parameters such as user IDs, endpoint IDs, and statuses to narrow the
	// returned data. In a URL, the parameters of this filter are specified with
	// dot notation. For example: `filter_by.statement_ids`.
	FilterBy types.Object `tfsdk:"-"`
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

func (to *ListQueryHistoryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryHistoryRequest) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				// Recursively sync the fields of FilterBy
				toFilterBy.SyncFieldsDuringCreateOrUpdate(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
}

func (to *ListQueryHistoryRequest) SyncFieldsDuringRead(ctx context.Context, from ListQueryHistoryRequest) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				toFilterBy.SyncFieldsDuringRead(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
}

func (c ListQueryHistoryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["include_metrics"] = attrs["include_metrics"].SetOptional()

	return attrs
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
			"filter_by":       QueryFilter{}.Type(ctx),
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
	var v QueryFilter
	d := o.FilterBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilterBy sets the value of the FilterBy field in ListQueryHistoryRequest.
func (o *ListQueryHistoryRequest) SetFilterBy(ctx context.Context, v QueryFilter) {
	vs := v.ToObjectValue(ctx)
	o.FilterBy = vs
}

type ListQueryObjectsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListQueryObjectsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryObjectsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListQueryObjectsResponse) SyncFieldsDuringRead(ctx context.Context, from ListQueryObjectsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c ListQueryObjectsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponseQuery struct {
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

func (to *ListQueryObjectsResponseQuery) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryObjectsResponseQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ListQueryObjectsResponseQuery) SyncFieldsDuringRead(ctx context.Context, from ListQueryObjectsResponseQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c ListQueryObjectsResponseQuery) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ListRequest struct {
}

func (to *ListRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRequest) {
}

func (to *ListRequest) SyncFieldsDuringRead(ctx context.Context, from ListRequest) {
}

func (c ListRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRequest
// only implements ToObjectValue() and Type().
func (o ListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListResponse struct {
	// The total number of dashboards.
	Count types.Int64 `tfsdk:"count"`
	// The current page being displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of dashboards per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of dashboards returned.
	Results types.List `tfsdk:"results"`
}

func (to *ListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListResponse) SyncFieldsDuringRead(ctx context.Context, from ListResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c ListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListVisualizationsForQueryRequest struct {
	Id types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListVisualizationsForQueryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVisualizationsForQueryRequest) {
}

func (to *ListVisualizationsForQueryRequest) SyncFieldsDuringRead(ctx context.Context, from ListVisualizationsForQueryRequest) {
}

func (c ListVisualizationsForQueryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListVisualizationsForQueryResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVisualizationsForQueryResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListVisualizationsForQueryResponse) SyncFieldsDuringRead(ctx context.Context, from ListVisualizationsForQueryResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c ListVisualizationsForQueryResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId types.Int64 `tfsdk:"-"`
}

func (to *ListWarehousesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWarehousesRequest) {
}

func (to *ListWarehousesRequest) SyncFieldsDuringRead(ctx context.Context, from ListWarehousesRequest) {
}

func (c ListWarehousesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_as_user_id"] = attrs["run_as_user_id"].SetOptional()

	return attrs
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
	Warehouses types.List `tfsdk:"warehouses"`
}

func (to *ListWarehousesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWarehousesResponse) {
	if !from.Warehouses.IsNull() && !from.Warehouses.IsUnknown() && to.Warehouses.IsNull() && len(from.Warehouses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Warehouses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Warehouses = from.Warehouses
	}
}

func (to *ListWarehousesResponse) SyncFieldsDuringRead(ctx context.Context, from ListWarehousesResponse) {
	if !from.Warehouses.IsNull() && !from.Warehouses.IsUnknown() && to.Warehouses.IsNull() && len(from.Warehouses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Warehouses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Warehouses = from.Warehouses
	}
}

func (c ListWarehousesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Warehouses = types.ListValueMust(t, vs)
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	Prefix types.String `tfsdk:"prefix"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator types.String `tfsdk:"separator"`
	// Character that suffixes each selected parameter value.
	Suffix types.String `tfsdk:"suffix"`
}

func (to *MultiValuesOptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MultiValuesOptions) {
}

func (to *MultiValuesOptions) SyncFieldsDuringRead(ctx context.Context, from MultiValuesOptions) {
}

func (c MultiValuesOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Value types.Float64 `tfsdk:"value"`
}

func (to *NumericValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NumericValue) {
}

func (to *NumericValue) SyncFieldsDuringRead(ctx context.Context, from NumericValue) {
}

func (c NumericValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Hostname types.String `tfsdk:"hostname"`

	Path types.String `tfsdk:"path"`

	Port types.Int64 `tfsdk:"port"`

	Protocol types.String `tfsdk:"protocol"`
}

func (to *OdbcParams) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OdbcParams) {
}

func (to *OdbcParams) SyncFieldsDuringRead(ctx context.Context, from OdbcParams) {
}

func (c OdbcParams) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	EnumOptions types.String `tfsdk:"enum_options"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions types.Object `tfsdk:"multi_values_options"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name types.String `tfsdk:"name"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId types.String `tfsdk:"query_id"`
	// The text displayed in a parameter picking widget.
	Title types.String `tfsdk:"title"`
	// Parameters can have several different types.
	Type_ types.String `tfsdk:"type"`
	// The default value for this parameter.
	Value types.Object `tfsdk:"value"`
}

func (to *Parameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Parameter) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				// Recursively sync the fields of MultiValuesOptions
				toMultiValuesOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
}

func (to *Parameter) SyncFieldsDuringRead(ctx context.Context, from Parameter) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				toMultiValuesOptions.SyncFieldsDuringRead(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
}

func (c Parameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enum_options"] = attrs["enum_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
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
func (a Parameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Parameter
// only implements ToObjectValue() and Type().
func (o Parameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enum_options":         o.EnumOptions,
			"multi_values_options": o.MultiValuesOptions,
			"name":                 o.Name,
			"query_id":             o.QueryId,
			"title":                o.Title,
			"type":                 o.Type_,
			"value":                o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Parameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enum_options":         types.StringType,
			"multi_values_options": MultiValuesOptions{}.Type(ctx),
			"name":                 types.StringType,
			"query_id":             types.StringType,
			"title":                types.StringType,
			"type":                 types.StringType,
			"value":                types.ObjectType{},
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
	var v MultiValuesOptions
	d := o.MultiValuesOptions.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in Parameter.
func (o *Parameter) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := v.ToObjectValue(ctx)
	o.MultiValuesOptions = vs
}

type Query struct {
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

func (to *Query) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Query) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Query) SyncFieldsDuringRead(ctx context.Context, from Query) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c Query) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.Object `tfsdk:"multi_values_options"`
	// UUID of the query that provides the parameter values.
	QueryId types.String `tfsdk:"query_id"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (to *QueryBackedValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryBackedValue) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				// Recursively sync the fields of MultiValuesOptions
				toMultiValuesOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *QueryBackedValue) SyncFieldsDuringRead(ctx context.Context, from QueryBackedValue) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				toMultiValuesOptions.SyncFieldsDuringRead(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (c QueryBackedValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
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
			"multi_values_options": MultiValuesOptions{}.Type(ctx),
			"query_id":             types.StringType,
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
	var v MultiValuesOptions
	d := o.MultiValuesOptions.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in QueryBackedValue.
func (o *QueryBackedValue) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions) {
	vs := v.ToObjectValue(ctx)
	o.MultiValuesOptions = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type QueryEditContent struct {
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

func (to *QueryEditContent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEditContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *QueryEditContent) SyncFieldsDuringRead(ctx context.Context, from QueryEditContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c QueryEditContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_source_id"] = attrs["data_source_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["run_as_role"] = attrs["run_as_role"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetRequired()

	return attrs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be less than
	// or equal to 30 days.
	QueryStartTimeRange types.Object `tfsdk:"query_start_time_range"`
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

func (to *QueryFilter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryFilter) {
	if !from.QueryStartTimeRange.IsNull() && !from.QueryStartTimeRange.IsUnknown() {
		if toQueryStartTimeRange, ok := to.GetQueryStartTimeRange(ctx); ok {
			if fromQueryStartTimeRange, ok := from.GetQueryStartTimeRange(ctx); ok {
				// Recursively sync the fields of QueryStartTimeRange
				toQueryStartTimeRange.SyncFieldsDuringCreateOrUpdate(ctx, fromQueryStartTimeRange)
				to.SetQueryStartTimeRange(ctx, toQueryStartTimeRange)
			}
		}
	}
	if !from.StatementIds.IsNull() && !from.StatementIds.IsUnknown() && to.StatementIds.IsNull() && len(from.StatementIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StatementIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StatementIds = from.StatementIds
	}
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
	if !from.UserIds.IsNull() && !from.UserIds.IsUnknown() && to.UserIds.IsNull() && len(from.UserIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserIds = from.UserIds
	}
	if !from.WarehouseIds.IsNull() && !from.WarehouseIds.IsUnknown() && to.WarehouseIds.IsNull() && len(from.WarehouseIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarehouseIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarehouseIds = from.WarehouseIds
	}
}

func (to *QueryFilter) SyncFieldsDuringRead(ctx context.Context, from QueryFilter) {
	if !from.QueryStartTimeRange.IsNull() && !from.QueryStartTimeRange.IsUnknown() {
		if toQueryStartTimeRange, ok := to.GetQueryStartTimeRange(ctx); ok {
			if fromQueryStartTimeRange, ok := from.GetQueryStartTimeRange(ctx); ok {
				toQueryStartTimeRange.SyncFieldsDuringRead(ctx, fromQueryStartTimeRange)
				to.SetQueryStartTimeRange(ctx, toQueryStartTimeRange)
			}
		}
	}
	if !from.StatementIds.IsNull() && !from.StatementIds.IsUnknown() && to.StatementIds.IsNull() && len(from.StatementIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StatementIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StatementIds = from.StatementIds
	}
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
	if !from.UserIds.IsNull() && !from.UserIds.IsUnknown() && to.UserIds.IsNull() && len(from.UserIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserIds = from.UserIds
	}
	if !from.WarehouseIds.IsNull() && !from.WarehouseIds.IsUnknown() && to.WarehouseIds.IsNull() && len(from.WarehouseIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarehouseIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarehouseIds = from.WarehouseIds
	}
}

func (c QueryFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_start_time_range"] = attrs["query_start_time_range"].SetOptional()
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
			"query_start_time_range": TimeRange{}.Type(ctx),
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
	var v TimeRange
	d := o.QueryStartTimeRange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryStartTimeRange sets the value of the QueryStartTimeRange field in QueryFilter.
func (o *QueryFilter) SetQueryStartTimeRange(ctx context.Context, v TimeRange) {
	vs := v.ToObjectValue(ctx)
	o.QueryStartTimeRange = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarehouseIds = types.ListValueMust(t, vs)
}

type QueryInfo struct {
	// The ID of the cached query if this result retrieved from cache
	CacheQueryId types.String `tfsdk:"cache_query_id"`
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed types.Object `tfsdk:"channel_used"`
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
	Metrics types.Object `tfsdk:"metrics"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState types.String `tfsdk:"plans_state"`
	// The time the query ended.
	QueryEndTimeMs types.Int64 `tfsdk:"query_end_time_ms"`
	// The query ID.
	QueryId types.String `tfsdk:"query_id"`
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	QuerySource types.Object `tfsdk:"query_source"`
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

func (to *QueryInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryInfo) {
	if !from.ChannelUsed.IsNull() && !from.ChannelUsed.IsUnknown() {
		if toChannelUsed, ok := to.GetChannelUsed(ctx); ok {
			if fromChannelUsed, ok := from.GetChannelUsed(ctx); ok {
				// Recursively sync the fields of ChannelUsed
				toChannelUsed.SyncFieldsDuringCreateOrUpdate(ctx, fromChannelUsed)
				to.SetChannelUsed(ctx, toChannelUsed)
			}
		}
	}
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() {
		if toMetrics, ok := to.GetMetrics(ctx); ok {
			if fromMetrics, ok := from.GetMetrics(ctx); ok {
				// Recursively sync the fields of Metrics
				toMetrics.SyncFieldsDuringCreateOrUpdate(ctx, fromMetrics)
				to.SetMetrics(ctx, toMetrics)
			}
		}
	}
	if !from.QuerySource.IsNull() && !from.QuerySource.IsUnknown() {
		if toQuerySource, ok := to.GetQuerySource(ctx); ok {
			if fromQuerySource, ok := from.GetQuerySource(ctx); ok {
				// Recursively sync the fields of QuerySource
				toQuerySource.SyncFieldsDuringCreateOrUpdate(ctx, fromQuerySource)
				to.SetQuerySource(ctx, toQuerySource)
			}
		}
	}
}

func (to *QueryInfo) SyncFieldsDuringRead(ctx context.Context, from QueryInfo) {
	if !from.ChannelUsed.IsNull() && !from.ChannelUsed.IsUnknown() {
		if toChannelUsed, ok := to.GetChannelUsed(ctx); ok {
			if fromChannelUsed, ok := from.GetChannelUsed(ctx); ok {
				toChannelUsed.SyncFieldsDuringRead(ctx, fromChannelUsed)
				to.SetChannelUsed(ctx, toChannelUsed)
			}
		}
	}
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() {
		if toMetrics, ok := to.GetMetrics(ctx); ok {
			if fromMetrics, ok := from.GetMetrics(ctx); ok {
				toMetrics.SyncFieldsDuringRead(ctx, fromMetrics)
				to.SetMetrics(ctx, toMetrics)
			}
		}
	}
	if !from.QuerySource.IsNull() && !from.QuerySource.IsUnknown() {
		if toQuerySource, ok := to.GetQuerySource(ctx); ok {
			if fromQuerySource, ok := from.GetQuerySource(ctx); ok {
				toQuerySource.SyncFieldsDuringRead(ctx, fromQuerySource)
				to.SetQuerySource(ctx, toQuerySource)
			}
		}
	}
}

func (c QueryInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cache_query_id"] = attrs["cache_query_id"].SetOptional()
	attrs["channel_used"] = attrs["channel_used"].SetOptional()
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
	attrs["plans_state"] = attrs["plans_state"].SetOptional()
	attrs["query_end_time_ms"] = attrs["query_end_time_ms"].SetOptional()
	attrs["query_id"] = attrs["query_id"].SetOptional()
	attrs["query_source"] = attrs["query_source"].SetOptional()
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
func (a QueryInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel_used": reflect.TypeOf(ChannelInfo{}),
		"metrics":      reflect.TypeOf(QueryMetrics{}),
		"query_source": reflect.TypeOf(ExternalQuerySource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryInfo
// only implements ToObjectValue() and Type().
func (o QueryInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cache_query_id":        o.CacheQueryId,
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
func (o QueryInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cache_query_id":        types.StringType,
			"channel_used":          ChannelInfo{}.Type(ctx),
			"client_application":    types.StringType,
			"duration":              types.Int64Type,
			"endpoint_id":           types.StringType,
			"error_message":         types.StringType,
			"executed_as_user_id":   types.Int64Type,
			"executed_as_user_name": types.StringType,
			"execution_end_time_ms": types.Int64Type,
			"is_final":              types.BoolType,
			"lookup_key":            types.StringType,
			"metrics":               QueryMetrics{}.Type(ctx),
			"plans_state":           types.StringType,
			"query_end_time_ms":     types.Int64Type,
			"query_id":              types.StringType,
			"query_source":          ExternalQuerySource{}.Type(ctx),
			"query_start_time_ms":   types.Int64Type,
			"query_text":            types.StringType,
			"rows_produced":         types.Int64Type,
			"spark_ui_url":          types.StringType,
			"statement_type":        types.StringType,
			"status":                types.StringType,
			"user_id":               types.Int64Type,
			"user_name":             types.StringType,
			"warehouse_id":          types.StringType,
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
	var v ChannelInfo
	d := o.ChannelUsed.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannelUsed sets the value of the ChannelUsed field in QueryInfo.
func (o *QueryInfo) SetChannelUsed(ctx context.Context, v ChannelInfo) {
	vs := v.ToObjectValue(ctx)
	o.ChannelUsed = vs
}

// GetMetrics returns the value of the Metrics field in QueryInfo as
// a QueryMetrics value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo) GetMetrics(ctx context.Context) (QueryMetrics, bool) {
	var e QueryMetrics
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return e, false
	}
	var v QueryMetrics
	d := o.Metrics.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in QueryInfo.
func (o *QueryInfo) SetMetrics(ctx context.Context, v QueryMetrics) {
	vs := v.ToObjectValue(ctx)
	o.Metrics = vs
}

// GetQuerySource returns the value of the QuerySource field in QueryInfo as
// a ExternalQuerySource value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryInfo) GetQuerySource(ctx context.Context) (ExternalQuerySource, bool) {
	var e ExternalQuerySource
	if o.QuerySource.IsNull() || o.QuerySource.IsUnknown() {
		return e, false
	}
	var v ExternalQuerySource
	d := o.QuerySource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuerySource sets the value of the QuerySource field in QueryInfo.
func (o *QueryInfo) SetQuerySource(ctx context.Context, v ExternalQuerySource) {
	vs := v.ToObjectValue(ctx)
	o.QuerySource = vs
}

type QueryList struct {
	// The total number of queries.
	Count types.Int64 `tfsdk:"count"`
	// The page number that is currently displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of queries per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of queries returned.
	Results types.List `tfsdk:"results"`
}

func (to *QueryList) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryList) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *QueryList) SyncFieldsDuringRead(ctx context.Context, from QueryList) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (c QueryList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
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
	// projected remaining work to be done aggregated across all stages in the
	// query, in milliseconds
	ProjectedRemainingTaskTotalTimeMs types.Int64 `tfsdk:"projected_remaining_task_total_time_ms"`
	// projected lower bound on remaining total task time based on
	// projected_remaining_task_total_time_ms / maximum concurrency
	ProjectedRemainingWallclockTimeMs types.Int64 `tfsdk:"projected_remaining_wallclock_time_ms"`
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
	// number of remaining tasks to complete this is based on the current status
	// and could be bigger or smaller in the future based on future updates
	RemainingTaskCount types.Int64 `tfsdk:"remaining_task_count"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs types.Int64 `tfsdk:"result_fetch_time_ms"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache types.Bool `tfsdk:"result_from_cache"`
	// Total number of rows returned by the query.
	RowsProducedCount types.Int64 `tfsdk:"rows_produced_count"`
	// Total number of rows read by the query.
	RowsReadCount types.Int64 `tfsdk:"rows_read_count"`
	// number of remaining tasks to complete, calculated by autoscaler
	// StatementAnalysis.scala deprecated: use remaining_task_count instead
	RunnableTasks types.Int64 `tfsdk:"runnable_tasks"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes types.Int64 `tfsdk:"spill_to_disk_bytes"`
	// sum of task times completed in a range of wall clock time, approximated
	// to a configurable number of points aggregated over all stages and jobs in
	// the query (based on task_total_time_ms)
	TaskTimeOverTimeRange types.Object `tfsdk:"task_time_over_time_range"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs types.Int64 `tfsdk:"task_total_time_ms"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs types.Int64 `tfsdk:"total_time_ms"`
	// remaining work to be done across all stages in the query, calculated by
	// autoscaler StatementAnalysis.scala, in milliseconds deprecated: using
	// projected_remaining_task_total_time_ms instead
	WorkToBeDone types.Int64 `tfsdk:"work_to_be_done"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes types.Int64 `tfsdk:"write_remote_bytes"`
}

func (to *QueryMetrics) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryMetrics) {
	if !from.TaskTimeOverTimeRange.IsNull() && !from.TaskTimeOverTimeRange.IsUnknown() {
		if toTaskTimeOverTimeRange, ok := to.GetTaskTimeOverTimeRange(ctx); ok {
			if fromTaskTimeOverTimeRange, ok := from.GetTaskTimeOverTimeRange(ctx); ok {
				// Recursively sync the fields of TaskTimeOverTimeRange
				toTaskTimeOverTimeRange.SyncFieldsDuringCreateOrUpdate(ctx, fromTaskTimeOverTimeRange)
				to.SetTaskTimeOverTimeRange(ctx, toTaskTimeOverTimeRange)
			}
		}
	}
}

func (to *QueryMetrics) SyncFieldsDuringRead(ctx context.Context, from QueryMetrics) {
	if !from.TaskTimeOverTimeRange.IsNull() && !from.TaskTimeOverTimeRange.IsUnknown() {
		if toTaskTimeOverTimeRange, ok := to.GetTaskTimeOverTimeRange(ctx); ok {
			if fromTaskTimeOverTimeRange, ok := from.GetTaskTimeOverTimeRange(ctx); ok {
				toTaskTimeOverTimeRange.SyncFieldsDuringRead(ctx, fromTaskTimeOverTimeRange)
				to.SetTaskTimeOverTimeRange(ctx, toTaskTimeOverTimeRange)
			}
		}
	}
}

func (c QueryMetrics) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compilation_time_ms"] = attrs["compilation_time_ms"].SetOptional()
	attrs["execution_time_ms"] = attrs["execution_time_ms"].SetOptional()
	attrs["network_sent_bytes"] = attrs["network_sent_bytes"].SetOptional()
	attrs["overloading_queue_start_timestamp"] = attrs["overloading_queue_start_timestamp"].SetOptional()
	attrs["photon_total_time_ms"] = attrs["photon_total_time_ms"].SetOptional()
	attrs["projected_remaining_task_total_time_ms"] = attrs["projected_remaining_task_total_time_ms"].SetOptional()
	attrs["projected_remaining_wallclock_time_ms"] = attrs["projected_remaining_wallclock_time_ms"].SetOptional()
	attrs["provisioning_queue_start_timestamp"] = attrs["provisioning_queue_start_timestamp"].SetOptional()
	attrs["pruned_bytes"] = attrs["pruned_bytes"].SetOptional()
	attrs["pruned_files_count"] = attrs["pruned_files_count"].SetOptional()
	attrs["query_compilation_start_timestamp"] = attrs["query_compilation_start_timestamp"].SetOptional()
	attrs["read_bytes"] = attrs["read_bytes"].SetOptional()
	attrs["read_cache_bytes"] = attrs["read_cache_bytes"].SetOptional()
	attrs["read_files_count"] = attrs["read_files_count"].SetOptional()
	attrs["read_partitions_count"] = attrs["read_partitions_count"].SetOptional()
	attrs["read_remote_bytes"] = attrs["read_remote_bytes"].SetOptional()
	attrs["remaining_task_count"] = attrs["remaining_task_count"].SetOptional()
	attrs["result_fetch_time_ms"] = attrs["result_fetch_time_ms"].SetOptional()
	attrs["result_from_cache"] = attrs["result_from_cache"].SetOptional()
	attrs["rows_produced_count"] = attrs["rows_produced_count"].SetOptional()
	attrs["rows_read_count"] = attrs["rows_read_count"].SetOptional()
	attrs["runnable_tasks"] = attrs["runnable_tasks"].SetOptional()
	attrs["spill_to_disk_bytes"] = attrs["spill_to_disk_bytes"].SetOptional()
	attrs["task_time_over_time_range"] = attrs["task_time_over_time_range"].SetOptional()
	attrs["task_total_time_ms"] = attrs["task_total_time_ms"].SetOptional()
	attrs["total_time_ms"] = attrs["total_time_ms"].SetOptional()
	attrs["work_to_be_done"] = attrs["work_to_be_done"].SetOptional()
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
func (a QueryMetrics) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"task_time_over_time_range": reflect.TypeOf(TaskTimeOverRange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryMetrics
// only implements ToObjectValue() and Type().
func (o QueryMetrics) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compilation_time_ms":                    o.CompilationTimeMs,
			"execution_time_ms":                      o.ExecutionTimeMs,
			"network_sent_bytes":                     o.NetworkSentBytes,
			"overloading_queue_start_timestamp":      o.OverloadingQueueStartTimestamp,
			"photon_total_time_ms":                   o.PhotonTotalTimeMs,
			"projected_remaining_task_total_time_ms": o.ProjectedRemainingTaskTotalTimeMs,
			"projected_remaining_wallclock_time_ms":  o.ProjectedRemainingWallclockTimeMs,
			"provisioning_queue_start_timestamp":     o.ProvisioningQueueStartTimestamp,
			"pruned_bytes":                           o.PrunedBytes,
			"pruned_files_count":                     o.PrunedFilesCount,
			"query_compilation_start_timestamp":      o.QueryCompilationStartTimestamp,
			"read_bytes":                             o.ReadBytes,
			"read_cache_bytes":                       o.ReadCacheBytes,
			"read_files_count":                       o.ReadFilesCount,
			"read_partitions_count":                  o.ReadPartitionsCount,
			"read_remote_bytes":                      o.ReadRemoteBytes,
			"remaining_task_count":                   o.RemainingTaskCount,
			"result_fetch_time_ms":                   o.ResultFetchTimeMs,
			"result_from_cache":                      o.ResultFromCache,
			"rows_produced_count":                    o.RowsProducedCount,
			"rows_read_count":                        o.RowsReadCount,
			"runnable_tasks":                         o.RunnableTasks,
			"spill_to_disk_bytes":                    o.SpillToDiskBytes,
			"task_time_over_time_range":              o.TaskTimeOverTimeRange,
			"task_total_time_ms":                     o.TaskTotalTimeMs,
			"total_time_ms":                          o.TotalTimeMs,
			"work_to_be_done":                        o.WorkToBeDone,
			"write_remote_bytes":                     o.WriteRemoteBytes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryMetrics) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compilation_time_ms":                    types.Int64Type,
			"execution_time_ms":                      types.Int64Type,
			"network_sent_bytes":                     types.Int64Type,
			"overloading_queue_start_timestamp":      types.Int64Type,
			"photon_total_time_ms":                   types.Int64Type,
			"projected_remaining_task_total_time_ms": types.Int64Type,
			"projected_remaining_wallclock_time_ms":  types.Int64Type,
			"provisioning_queue_start_timestamp":     types.Int64Type,
			"pruned_bytes":                           types.Int64Type,
			"pruned_files_count":                     types.Int64Type,
			"query_compilation_start_timestamp":      types.Int64Type,
			"read_bytes":                             types.Int64Type,
			"read_cache_bytes":                       types.Int64Type,
			"read_files_count":                       types.Int64Type,
			"read_partitions_count":                  types.Int64Type,
			"read_remote_bytes":                      types.Int64Type,
			"remaining_task_count":                   types.Int64Type,
			"result_fetch_time_ms":                   types.Int64Type,
			"result_from_cache":                      types.BoolType,
			"rows_produced_count":                    types.Int64Type,
			"rows_read_count":                        types.Int64Type,
			"runnable_tasks":                         types.Int64Type,
			"spill_to_disk_bytes":                    types.Int64Type,
			"task_time_over_time_range":              TaskTimeOverRange{}.Type(ctx),
			"task_total_time_ms":                     types.Int64Type,
			"total_time_ms":                          types.Int64Type,
			"work_to_be_done":                        types.Int64Type,
			"write_remote_bytes":                     types.Int64Type,
		},
	}
}

// GetTaskTimeOverTimeRange returns the value of the TaskTimeOverTimeRange field in QueryMetrics as
// a TaskTimeOverRange value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryMetrics) GetTaskTimeOverTimeRange(ctx context.Context) (TaskTimeOverRange, bool) {
	var e TaskTimeOverRange
	if o.TaskTimeOverTimeRange.IsNull() || o.TaskTimeOverTimeRange.IsUnknown() {
		return e, false
	}
	var v TaskTimeOverRange
	d := o.TaskTimeOverTimeRange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTaskTimeOverTimeRange sets the value of the TaskTimeOverTimeRange field in QueryMetrics.
func (o *QueryMetrics) SetTaskTimeOverTimeRange(ctx context.Context, v TaskTimeOverRange) {
	vs := v.ToObjectValue(ctx)
	o.TaskTimeOverTimeRange = vs
}

type QueryOptions struct {
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

func (to *QueryOptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryOptions) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (to *QueryOptions) SyncFieldsDuringRead(ctx context.Context, from QueryOptions) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (c QueryOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue types.Object `tfsdk:"date_range_value"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue types.Object `tfsdk:"date_value"`
	// Dropdown query parameter value.
	EnumValue types.Object `tfsdk:"enum_value"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name types.String `tfsdk:"name"`
	// Numeric query parameter value.
	NumericValue types.Object `tfsdk:"numeric_value"`
	// Query-based dropdown query parameter value.
	QueryBackedValue types.Object `tfsdk:"query_backed_value"`
	// Text query parameter value.
	TextValue types.Object `tfsdk:"text_value"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title types.String `tfsdk:"title"`
}

func (to *QueryParameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryParameter) {
	if !from.DateRangeValue.IsNull() && !from.DateRangeValue.IsUnknown() {
		if toDateRangeValue, ok := to.GetDateRangeValue(ctx); ok {
			if fromDateRangeValue, ok := from.GetDateRangeValue(ctx); ok {
				// Recursively sync the fields of DateRangeValue
				toDateRangeValue.SyncFieldsDuringCreateOrUpdate(ctx, fromDateRangeValue)
				to.SetDateRangeValue(ctx, toDateRangeValue)
			}
		}
	}
	if !from.DateValue.IsNull() && !from.DateValue.IsUnknown() {
		if toDateValue, ok := to.GetDateValue(ctx); ok {
			if fromDateValue, ok := from.GetDateValue(ctx); ok {
				// Recursively sync the fields of DateValue
				toDateValue.SyncFieldsDuringCreateOrUpdate(ctx, fromDateValue)
				to.SetDateValue(ctx, toDateValue)
			}
		}
	}
	if !from.EnumValue.IsNull() && !from.EnumValue.IsUnknown() {
		if toEnumValue, ok := to.GetEnumValue(ctx); ok {
			if fromEnumValue, ok := from.GetEnumValue(ctx); ok {
				// Recursively sync the fields of EnumValue
				toEnumValue.SyncFieldsDuringCreateOrUpdate(ctx, fromEnumValue)
				to.SetEnumValue(ctx, toEnumValue)
			}
		}
	}
	if !from.NumericValue.IsNull() && !from.NumericValue.IsUnknown() {
		if toNumericValue, ok := to.GetNumericValue(ctx); ok {
			if fromNumericValue, ok := from.GetNumericValue(ctx); ok {
				// Recursively sync the fields of NumericValue
				toNumericValue.SyncFieldsDuringCreateOrUpdate(ctx, fromNumericValue)
				to.SetNumericValue(ctx, toNumericValue)
			}
		}
	}
	if !from.QueryBackedValue.IsNull() && !from.QueryBackedValue.IsUnknown() {
		if toQueryBackedValue, ok := to.GetQueryBackedValue(ctx); ok {
			if fromQueryBackedValue, ok := from.GetQueryBackedValue(ctx); ok {
				// Recursively sync the fields of QueryBackedValue
				toQueryBackedValue.SyncFieldsDuringCreateOrUpdate(ctx, fromQueryBackedValue)
				to.SetQueryBackedValue(ctx, toQueryBackedValue)
			}
		}
	}
	if !from.TextValue.IsNull() && !from.TextValue.IsUnknown() {
		if toTextValue, ok := to.GetTextValue(ctx); ok {
			if fromTextValue, ok := from.GetTextValue(ctx); ok {
				// Recursively sync the fields of TextValue
				toTextValue.SyncFieldsDuringCreateOrUpdate(ctx, fromTextValue)
				to.SetTextValue(ctx, toTextValue)
			}
		}
	}
}

func (to *QueryParameter) SyncFieldsDuringRead(ctx context.Context, from QueryParameter) {
	if !from.DateRangeValue.IsNull() && !from.DateRangeValue.IsUnknown() {
		if toDateRangeValue, ok := to.GetDateRangeValue(ctx); ok {
			if fromDateRangeValue, ok := from.GetDateRangeValue(ctx); ok {
				toDateRangeValue.SyncFieldsDuringRead(ctx, fromDateRangeValue)
				to.SetDateRangeValue(ctx, toDateRangeValue)
			}
		}
	}
	if !from.DateValue.IsNull() && !from.DateValue.IsUnknown() {
		if toDateValue, ok := to.GetDateValue(ctx); ok {
			if fromDateValue, ok := from.GetDateValue(ctx); ok {
				toDateValue.SyncFieldsDuringRead(ctx, fromDateValue)
				to.SetDateValue(ctx, toDateValue)
			}
		}
	}
	if !from.EnumValue.IsNull() && !from.EnumValue.IsUnknown() {
		if toEnumValue, ok := to.GetEnumValue(ctx); ok {
			if fromEnumValue, ok := from.GetEnumValue(ctx); ok {
				toEnumValue.SyncFieldsDuringRead(ctx, fromEnumValue)
				to.SetEnumValue(ctx, toEnumValue)
			}
		}
	}
	if !from.NumericValue.IsNull() && !from.NumericValue.IsUnknown() {
		if toNumericValue, ok := to.GetNumericValue(ctx); ok {
			if fromNumericValue, ok := from.GetNumericValue(ctx); ok {
				toNumericValue.SyncFieldsDuringRead(ctx, fromNumericValue)
				to.SetNumericValue(ctx, toNumericValue)
			}
		}
	}
	if !from.QueryBackedValue.IsNull() && !from.QueryBackedValue.IsUnknown() {
		if toQueryBackedValue, ok := to.GetQueryBackedValue(ctx); ok {
			if fromQueryBackedValue, ok := from.GetQueryBackedValue(ctx); ok {
				toQueryBackedValue.SyncFieldsDuringRead(ctx, fromQueryBackedValue)
				to.SetQueryBackedValue(ctx, toQueryBackedValue)
			}
		}
	}
	if !from.TextValue.IsNull() && !from.TextValue.IsUnknown() {
		if toTextValue, ok := to.GetTextValue(ctx); ok {
			if fromTextValue, ok := from.GetTextValue(ctx); ok {
				toTextValue.SyncFieldsDuringRead(ctx, fromTextValue)
				to.SetTextValue(ctx, toTextValue)
			}
		}
	}
}

func (c QueryParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["date_range_value"] = attrs["date_range_value"].SetOptional()
	attrs["date_value"] = attrs["date_value"].SetOptional()
	attrs["enum_value"] = attrs["enum_value"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["numeric_value"] = attrs["numeric_value"].SetOptional()
	attrs["query_backed_value"] = attrs["query_backed_value"].SetOptional()
	attrs["text_value"] = attrs["text_value"].SetOptional()
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
			"date_range_value":   DateRangeValue{}.Type(ctx),
			"date_value":         DateValue{}.Type(ctx),
			"enum_value":         EnumValue{}.Type(ctx),
			"name":               types.StringType,
			"numeric_value":      NumericValue{}.Type(ctx),
			"query_backed_value": QueryBackedValue{}.Type(ctx),
			"text_value":         TextValue{}.Type(ctx),
			"title":              types.StringType,
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
	var v DateRangeValue
	d := o.DateRangeValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDateRangeValue sets the value of the DateRangeValue field in QueryParameter.
func (o *QueryParameter) SetDateRangeValue(ctx context.Context, v DateRangeValue) {
	vs := v.ToObjectValue(ctx)
	o.DateRangeValue = vs
}

// GetDateValue returns the value of the DateValue field in QueryParameter as
// a DateValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetDateValue(ctx context.Context) (DateValue, bool) {
	var e DateValue
	if o.DateValue.IsNull() || o.DateValue.IsUnknown() {
		return e, false
	}
	var v DateValue
	d := o.DateValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDateValue sets the value of the DateValue field in QueryParameter.
func (o *QueryParameter) SetDateValue(ctx context.Context, v DateValue) {
	vs := v.ToObjectValue(ctx)
	o.DateValue = vs
}

// GetEnumValue returns the value of the EnumValue field in QueryParameter as
// a EnumValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetEnumValue(ctx context.Context) (EnumValue, bool) {
	var e EnumValue
	if o.EnumValue.IsNull() || o.EnumValue.IsUnknown() {
		return e, false
	}
	var v EnumValue
	d := o.EnumValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnumValue sets the value of the EnumValue field in QueryParameter.
func (o *QueryParameter) SetEnumValue(ctx context.Context, v EnumValue) {
	vs := v.ToObjectValue(ctx)
	o.EnumValue = vs
}

// GetNumericValue returns the value of the NumericValue field in QueryParameter as
// a NumericValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetNumericValue(ctx context.Context) (NumericValue, bool) {
	var e NumericValue
	if o.NumericValue.IsNull() || o.NumericValue.IsUnknown() {
		return e, false
	}
	var v NumericValue
	d := o.NumericValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNumericValue sets the value of the NumericValue field in QueryParameter.
func (o *QueryParameter) SetNumericValue(ctx context.Context, v NumericValue) {
	vs := v.ToObjectValue(ctx)
	o.NumericValue = vs
}

// GetQueryBackedValue returns the value of the QueryBackedValue field in QueryParameter as
// a QueryBackedValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetQueryBackedValue(ctx context.Context) (QueryBackedValue, bool) {
	var e QueryBackedValue
	if o.QueryBackedValue.IsNull() || o.QueryBackedValue.IsUnknown() {
		return e, false
	}
	var v QueryBackedValue
	d := o.QueryBackedValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryBackedValue sets the value of the QueryBackedValue field in QueryParameter.
func (o *QueryParameter) SetQueryBackedValue(ctx context.Context, v QueryBackedValue) {
	vs := v.ToObjectValue(ctx)
	o.QueryBackedValue = vs
}

// GetTextValue returns the value of the TextValue field in QueryParameter as
// a TextValue value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryParameter) GetTextValue(ctx context.Context) (TextValue, bool) {
	var e TextValue
	if o.TextValue.IsNull() || o.TextValue.IsUnknown() {
		return e, false
	}
	var v TextValue
	d := o.TextValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTextValue sets the value of the TextValue field in QueryParameter.
func (o *QueryParameter) SetTextValue(ctx context.Context, v TextValue) {
	vs := v.ToObjectValue(ctx)
	o.TextValue = vs
}

type QueryPostContent struct {
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

func (to *QueryPostContent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryPostContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *QueryPostContent) SyncFieldsDuringRead(ctx context.Context, from QueryPostContent) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c QueryPostContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_source_id"] = attrs["data_source_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["run_as_role"] = attrs["run_as_role"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs"`
}

func (to *RepeatedEndpointConfPairs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepeatedEndpointConfPairs) {
	if !from.ConfigPair.IsNull() && !from.ConfigPair.IsUnknown() && to.ConfigPair.IsNull() && len(from.ConfigPair.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ConfigPair, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ConfigPair = from.ConfigPair
	}
	if !from.ConfigurationPairs.IsNull() && !from.ConfigurationPairs.IsUnknown() && to.ConfigurationPairs.IsNull() && len(from.ConfigurationPairs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ConfigurationPairs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ConfigurationPairs = from.ConfigurationPairs
	}
}

func (to *RepeatedEndpointConfPairs) SyncFieldsDuringRead(ctx context.Context, from RepeatedEndpointConfPairs) {
	if !from.ConfigPair.IsNull() && !from.ConfigPair.IsUnknown() && to.ConfigPair.IsNull() && len(from.ConfigPair.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ConfigPair, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ConfigPair = from.ConfigPair
	}
	if !from.ConfigurationPairs.IsNull() && !from.ConfigurationPairs.IsUnknown() && to.ConfigurationPairs.IsNull() && len(from.ConfigurationPairs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ConfigurationPairs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ConfigurationPairs = from.ConfigurationPairs
	}
}

func (c RepeatedEndpointConfPairs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.ConfigurationPairs = types.ListValueMust(t, vs)
}

type RestoreDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *RestoreDashboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreDashboardRequest) {
}

func (to *RestoreDashboardRequest) SyncFieldsDuringRead(ctx context.Context, from RestoreDashboardRequest) {
}

func (c RestoreDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()

	return attrs
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

type RestoreQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *RestoreQueriesLegacyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreQueriesLegacyRequest) {
}

func (to *RestoreQueriesLegacyRequest) SyncFieldsDuringRead(ctx context.Context, from RestoreQueriesLegacyRequest) {
}

func (c RestoreQueriesLegacyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_id"] = attrs["query_id"].SetRequired()

	return attrs
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

func (to *RestoreResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreResponse) {
}

func (to *RestoreResponse) SyncFieldsDuringRead(ctx context.Context, from RestoreResponse) {
}

func (c RestoreResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (to *ResultData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() && to.ExternalLinks.IsNull() && len(from.ExternalLinks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExternalLinks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExternalLinks = from.ExternalLinks
	}
}

func (to *ResultData) SyncFieldsDuringRead(ctx context.Context, from ResultData) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() && to.ExternalLinks.IsNull() && len(from.ExternalLinks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExternalLinks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExternalLinks = from.ExternalLinks
	}
}

func (c ResultData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExternalLinks = types.ListValueMust(t, vs)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks"`

	Format types.String `tfsdk:"format"`

	Schema types.Object `tfsdk:"schema"`
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

func (to *ResultManifest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest) {
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() && to.Chunks.IsNull() && len(from.Chunks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Chunks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Chunks = from.Chunks
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				// Recursively sync the fields of Schema
				toSchema.SyncFieldsDuringCreateOrUpdate(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
}

func (to *ResultManifest) SyncFieldsDuringRead(ctx context.Context, from ResultManifest) {
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() && to.Chunks.IsNull() && len(from.Chunks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Chunks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Chunks = from.Chunks
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				toSchema.SyncFieldsDuringRead(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
}

func (c ResultManifest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["chunks"] = attrs["chunks"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
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
			"format":            types.StringType,
			"schema":            ResultSchema{}.Type(ctx),
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	var v ResultSchema
	d := o.Schema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchema sets the value of the Schema field in ResultManifest.
func (o *ResultManifest) SetSchema(ctx context.Context, v ResultSchema) {
	vs := v.ToObjectValue(ctx)
	o.Schema = vs
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	ColumnCount types.Int64 `tfsdk:"column_count"`

	Columns types.List `tfsdk:"columns"`
}

func (to *ResultSchema) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultSchema) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *ResultSchema) SyncFieldsDuringRead(ctx context.Context, from ResultSchema) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (c ResultSchema) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type ServiceError struct {
	ErrorCode types.String `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

func (to *ServiceError) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServiceError) {
}

func (to *ServiceError) SyncFieldsDuringRead(ctx context.Context, from ServiceError) {
}

func (c ServiceError) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AccessControlList types.List `tfsdk:"access_control_list"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permission to set.
	ObjectType types.String `tfsdk:"-"`
}

func (to *SetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SetRequest) SyncFieldsDuringRead(ctx context.Context, from SetRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c SetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
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
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetRequest) Type(ctx context.Context) attr.Type {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type"`
}

func (to *SetResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetResponse) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SetResponse) SyncFieldsDuringRead(ctx context.Context, from SetResponse) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c SetResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel types.Object `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.Object `tfsdk:"config_param"`
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
	GlobalParam types.Object `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.Object `tfsdk:"sql_configuration_parameters"`
}

func (to *SetWorkspaceWarehouseConfigRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetWorkspaceWarehouseConfigRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				// Recursively sync the fields of Channel
				toChannel.SyncFieldsDuringCreateOrUpdate(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.ConfigParam.IsNull() && !from.ConfigParam.IsUnknown() {
		if toConfigParam, ok := to.GetConfigParam(ctx); ok {
			if fromConfigParam, ok := from.GetConfigParam(ctx); ok {
				// Recursively sync the fields of ConfigParam
				toConfigParam.SyncFieldsDuringCreateOrUpdate(ctx, fromConfigParam)
				to.SetConfigParam(ctx, toConfigParam)
			}
		}
	}
	if !from.DataAccessConfig.IsNull() && !from.DataAccessConfig.IsUnknown() && to.DataAccessConfig.IsNull() && len(from.DataAccessConfig.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataAccessConfig, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataAccessConfig = from.DataAccessConfig
	}
	if !from.EnabledWarehouseTypes.IsNull() && !from.EnabledWarehouseTypes.IsUnknown() && to.EnabledWarehouseTypes.IsNull() && len(from.EnabledWarehouseTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnabledWarehouseTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnabledWarehouseTypes = from.EnabledWarehouseTypes
	}
	if !from.GlobalParam.IsNull() && !from.GlobalParam.IsUnknown() {
		if toGlobalParam, ok := to.GetGlobalParam(ctx); ok {
			if fromGlobalParam, ok := from.GetGlobalParam(ctx); ok {
				// Recursively sync the fields of GlobalParam
				toGlobalParam.SyncFieldsDuringCreateOrUpdate(ctx, fromGlobalParam)
				to.SetGlobalParam(ctx, toGlobalParam)
			}
		}
	}
	if !from.SqlConfigurationParameters.IsNull() && !from.SqlConfigurationParameters.IsUnknown() {
		if toSqlConfigurationParameters, ok := to.GetSqlConfigurationParameters(ctx); ok {
			if fromSqlConfigurationParameters, ok := from.GetSqlConfigurationParameters(ctx); ok {
				// Recursively sync the fields of SqlConfigurationParameters
				toSqlConfigurationParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromSqlConfigurationParameters)
				to.SetSqlConfigurationParameters(ctx, toSqlConfigurationParameters)
			}
		}
	}
}

func (to *SetWorkspaceWarehouseConfigRequest) SyncFieldsDuringRead(ctx context.Context, from SetWorkspaceWarehouseConfigRequest) {
	if !from.Channel.IsNull() && !from.Channel.IsUnknown() {
		if toChannel, ok := to.GetChannel(ctx); ok {
			if fromChannel, ok := from.GetChannel(ctx); ok {
				toChannel.SyncFieldsDuringRead(ctx, fromChannel)
				to.SetChannel(ctx, toChannel)
			}
		}
	}
	if !from.ConfigParam.IsNull() && !from.ConfigParam.IsUnknown() {
		if toConfigParam, ok := to.GetConfigParam(ctx); ok {
			if fromConfigParam, ok := from.GetConfigParam(ctx); ok {
				toConfigParam.SyncFieldsDuringRead(ctx, fromConfigParam)
				to.SetConfigParam(ctx, toConfigParam)
			}
		}
	}
	if !from.DataAccessConfig.IsNull() && !from.DataAccessConfig.IsUnknown() && to.DataAccessConfig.IsNull() && len(from.DataAccessConfig.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataAccessConfig, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataAccessConfig = from.DataAccessConfig
	}
	if !from.EnabledWarehouseTypes.IsNull() && !from.EnabledWarehouseTypes.IsUnknown() && to.EnabledWarehouseTypes.IsNull() && len(from.EnabledWarehouseTypes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnabledWarehouseTypes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnabledWarehouseTypes = from.EnabledWarehouseTypes
	}
	if !from.GlobalParam.IsNull() && !from.GlobalParam.IsUnknown() {
		if toGlobalParam, ok := to.GetGlobalParam(ctx); ok {
			if fromGlobalParam, ok := from.GetGlobalParam(ctx); ok {
				toGlobalParam.SyncFieldsDuringRead(ctx, fromGlobalParam)
				to.SetGlobalParam(ctx, toGlobalParam)
			}
		}
	}
	if !from.SqlConfigurationParameters.IsNull() && !from.SqlConfigurationParameters.IsUnknown() {
		if toSqlConfigurationParameters, ok := to.GetSqlConfigurationParameters(ctx); ok {
			if fromSqlConfigurationParameters, ok := from.GetSqlConfigurationParameters(ctx); ok {
				toSqlConfigurationParameters.SyncFieldsDuringRead(ctx, fromSqlConfigurationParameters)
				to.SetSqlConfigurationParameters(ctx, toSqlConfigurationParameters)
			}
		}
	}
}

func (c SetWorkspaceWarehouseConfigRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["config_param"] = attrs["config_param"].SetOptional()
	attrs["data_access_config"] = attrs["data_access_config"].SetOptional()
	attrs["enabled_warehouse_types"] = attrs["enabled_warehouse_types"].SetOptional()
	attrs["global_param"] = attrs["global_param"].SetOptional()
	attrs["google_service_account"] = attrs["google_service_account"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["security_policy"] = attrs["security_policy"].SetOptional()
	attrs["sql_configuration_parameters"] = attrs["sql_configuration_parameters"].SetOptional()

	return attrs
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
			"channel":      Channel{}.Type(ctx),
			"config_param": RepeatedEndpointConfPairs{}.Type(ctx),
			"data_access_config": basetypes.ListType{
				ElemType: EndpointConfPair{}.Type(ctx),
			},
			"enabled_warehouse_types": basetypes.ListType{
				ElemType: WarehouseTypePair{}.Type(ctx),
			},
			"global_param":                 RepeatedEndpointConfPairs{}.Type(ctx),
			"google_service_account":       types.StringType,
			"instance_profile_arn":         types.StringType,
			"security_policy":              types.StringType,
			"sql_configuration_parameters": RepeatedEndpointConfPairs{}.Type(ctx),
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
	var v Channel
	d := o.Channel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChannel sets the value of the Channel field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetChannel(ctx context.Context, v Channel) {
	vs := v.ToObjectValue(ctx)
	o.Channel = vs
}

// GetConfigParam returns the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.ConfigParam.IsNull() || o.ConfigParam.IsUnknown() {
		return e, false
	}
	var v RepeatedEndpointConfPairs
	d := o.ConfigParam.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigParam sets the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.ConfigParam = vs
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	var v RepeatedEndpointConfPairs
	d := o.GlobalParam.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlobalParam sets the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.GlobalParam = vs
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest as
// a RepeatedEndpointConfPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *SetWorkspaceWarehouseConfigRequest) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs, bool) {
	var e RepeatedEndpointConfPairs
	if o.SqlConfigurationParameters.IsNull() || o.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v RepeatedEndpointConfPairs
	d := o.SqlConfigurationParameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest.
func (o *SetWorkspaceWarehouseConfigRequest) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs) {
	vs := v.ToObjectValue(ctx)
	o.SqlConfigurationParameters = vs
}

type SetWorkspaceWarehouseConfigResponse struct {
}

func (to *SetWorkspaceWarehouseConfigResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetWorkspaceWarehouseConfigResponse) {
}

func (to *SetWorkspaceWarehouseConfigResponse) SyncFieldsDuringRead(ctx context.Context, from SetWorkspaceWarehouseConfigResponse) {
}

func (c SetWorkspaceWarehouseConfigResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *StartRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartRequest) {
}

func (to *StartRequest) SyncFieldsDuringRead(ctx context.Context, from StartRequest) {
}

func (c StartRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *StartWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartWarehouseResponse) {
}

func (to *StartWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from StartWarehouseResponse) {
}

func (c StartWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (to *StatementParameterListItem) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementParameterListItem) {
}

func (to *StatementParameterListItem) SyncFieldsDuringRead(ctx context.Context, from StatementParameterListItem) {
}

func (c StatementParameterListItem) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Manifest types.Object `tfsdk:"manifest"`

	Result types.Object `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`

	Status types.Object `tfsdk:"status"`
}

func (to *StatementResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementResponse) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *StatementResponse) SyncFieldsDuringRead(ctx context.Context, from StatementResponse) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (c StatementResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
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
			"manifest":     ResultManifest{}.Type(ctx),
			"result":       ResultData{}.Type(ctx),
			"statement_id": types.StringType,
			"status":       StatementStatus{}.Type(ctx),
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
	var v ResultManifest
	d := o.Manifest.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManifest sets the value of the Manifest field in StatementResponse.
func (o *StatementResponse) SetManifest(ctx context.Context, v ResultManifest) {
	vs := v.ToObjectValue(ctx)
	o.Manifest = vs
}

// GetResult returns the value of the Result field in StatementResponse as
// a ResultData value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse) GetResult(ctx context.Context) (ResultData, bool) {
	var e ResultData
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v ResultData
	d := o.Result.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResult sets the value of the Result field in StatementResponse.
func (o *StatementResponse) SetResult(ctx context.Context, v ResultData) {
	vs := v.ToObjectValue(ctx)
	o.Result = vs
}

// GetStatus returns the value of the Status field in StatementResponse as
// a StatementStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *StatementResponse) GetStatus(ctx context.Context) (StatementStatus, bool) {
	var e StatementStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v StatementStatus
	d := o.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in StatementResponse.
func (o *StatementResponse) SetStatus(ctx context.Context, v StatementStatus) {
	vs := v.ToObjectValue(ctx)
	o.Status = vs
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	Error types.Object `tfsdk:"error"`

	State types.String `tfsdk:"state"`
}

func (to *StatementStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementStatus) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *StatementStatus) SyncFieldsDuringRead(ctx context.Context, from StatementStatus) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (c StatementStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
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
			"error": ServiceError{}.Type(ctx),
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
	var v ServiceError
	d := o.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in StatementStatus.
func (o *StatementStatus) SetError(ctx context.Context, v ServiceError) {
	vs := v.ToObjectValue(ctx)
	o.Error = vs
}

type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *StopRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopRequest) {
}

func (to *StopRequest) SyncFieldsDuringRead(ctx context.Context, from StopRequest) {
}

func (c StopRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

func (to *StopWarehouseResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopWarehouseResponse) {
}

func (to *StopWarehouseResponse) SyncFieldsDuringRead(ctx context.Context, from StopWarehouseResponse) {
}

func (c StopWarehouseResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Message types.String `tfsdk:"message"`
}

func (to *Success) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Success) {
}

func (to *Success) SyncFieldsDuringRead(ctx context.Context, from Success) {
}

func (c Success) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type TaskTimeOverRange struct {
	Entries types.List `tfsdk:"entries"`
	// interval length for all entries (difference in start time and end time of
	// an entry range) the same for all entries start time of first interval is
	// query_start_time_ms
	Interval types.Int64 `tfsdk:"interval"`
}

func (to *TaskTimeOverRange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TaskTimeOverRange) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
}

func (to *TaskTimeOverRange) SyncFieldsDuringRead(ctx context.Context, from TaskTimeOverRange) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
}

func (c TaskTimeOverRange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TaskTimeOverRange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entries": reflect.TypeOf(TaskTimeOverRangeEntry{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRange
// only implements ToObjectValue() and Type().
func (o TaskTimeOverRange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entries":  o.Entries,
			"interval": o.Interval,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskTimeOverRange) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entries": basetypes.ListType{
				ElemType: TaskTimeOverRangeEntry{}.Type(ctx),
			},
			"interval": types.Int64Type,
		},
	}
}

// GetEntries returns the value of the Entries field in TaskTimeOverRange as
// a slice of TaskTimeOverRangeEntry values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskTimeOverRange) GetEntries(ctx context.Context) ([]TaskTimeOverRangeEntry, bool) {
	if o.Entries.IsNull() || o.Entries.IsUnknown() {
		return nil, false
	}
	var v []TaskTimeOverRangeEntry
	d := o.Entries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntries sets the value of the Entries field in TaskTimeOverRange.
func (o *TaskTimeOverRange) SetEntries(ctx context.Context, v []TaskTimeOverRangeEntry) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entries = types.ListValueMust(t, vs)
}

type TaskTimeOverRangeEntry struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	TaskCompletedTimeMs types.Int64 `tfsdk:"task_completed_time_ms"`
}

func (to *TaskTimeOverRangeEntry) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TaskTimeOverRangeEntry) {
}

func (to *TaskTimeOverRangeEntry) SyncFieldsDuringRead(ctx context.Context, from TaskTimeOverRangeEntry) {
}

func (c TaskTimeOverRangeEntry) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TaskTimeOverRangeEntry) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRangeEntry
// only implements ToObjectValue() and Type().
func (o TaskTimeOverRangeEntry) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"task_completed_time_ms": o.TaskCompletedTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskTimeOverRangeEntry) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"task_completed_time_ms": types.Int64Type,
		},
	}
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters"`
	// type of the termination
	Type_ types.String `tfsdk:"type"`
}

func (to *TerminationReason) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TerminationReason) {
}

func (to *TerminationReason) SyncFieldsDuringRead(ctx context.Context, from TerminationReason) {
}

func (c TerminationReason) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type TextValue struct {
	Value types.String `tfsdk:"value"`
}

func (to *TextValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TextValue) {
}

func (to *TextValue) SyncFieldsDuringRead(ctx context.Context, from TextValue) {
}

func (c TextValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// The start time in milliseconds.
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
}

func (to *TimeRange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeRange) {
}

func (to *TimeRange) SyncFieldsDuringRead(ctx context.Context, from TimeRange) {
}

func (c TimeRange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	NewOwner types.String `tfsdk:"new_owner"`
}

func (to *TransferOwnershipObjectId) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransferOwnershipObjectId) {
}

func (to *TransferOwnershipObjectId) SyncFieldsDuringRead(ctx context.Context, from TransferOwnershipObjectId) {
}

func (c TransferOwnershipObjectId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner"`
	// The ID of the object on which to change ownership.
	ObjectId types.Object `tfsdk:"-"`
	// The type of object on which to change ownership.
	ObjectType types.String `tfsdk:"-"`
}

func (to *TransferOwnershipRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransferOwnershipRequest) {
	if !from.ObjectId.IsNull() && !from.ObjectId.IsUnknown() {
		if toObjectId, ok := to.GetObjectId(ctx); ok {
			if fromObjectId, ok := from.GetObjectId(ctx); ok {
				// Recursively sync the fields of ObjectId
				toObjectId.SyncFieldsDuringCreateOrUpdate(ctx, fromObjectId)
				to.SetObjectId(ctx, toObjectId)
			}
		}
	}
}

func (to *TransferOwnershipRequest) SyncFieldsDuringRead(ctx context.Context, from TransferOwnershipRequest) {
	if !from.ObjectId.IsNull() && !from.ObjectId.IsUnknown() {
		if toObjectId, ok := to.GetObjectId(ctx); ok {
			if fromObjectId, ok := from.GetObjectId(ctx); ok {
				toObjectId.SyncFieldsDuringRead(ctx, fromObjectId)
				to.SetObjectId(ctx, toObjectId)
			}
		}
	}
}

func (c TransferOwnershipRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["new_owner"] = attrs["new_owner"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
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
		"object_id": reflect.TypeOf(TransferOwnershipObjectId{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipRequest
// only implements ToObjectValue() and Type().
func (o TransferOwnershipRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner":   o.NewOwner,
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransferOwnershipRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner":   types.StringType,
			"object_id":   TransferOwnershipObjectId{}.Type(ctx),
			"object_type": types.StringType,
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
	var v TransferOwnershipObjectId
	d := o.ObjectId.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjectId sets the value of the ObjectId field in TransferOwnershipRequest.
func (o *TransferOwnershipRequest) SetObjectId(ctx context.Context, v TransferOwnershipObjectId) {
	vs := v.ToObjectValue(ctx)
	o.ObjectId = vs
}

type TrashAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashAlertRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashAlertRequest) {
}

func (to *TrashAlertRequest) SyncFieldsDuringRead(ctx context.Context, from TrashAlertRequest) {
}

func (c TrashAlertRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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

type TrashAlertV2Request struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashAlertV2Request) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashAlertV2Request) {
}

func (to *TrashAlertV2Request) SyncFieldsDuringRead(ctx context.Context, from TrashAlertV2Request) {
}

func (c TrashAlertV2Request) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashAlertV2Request) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertV2Request
// only implements ToObjectValue() and Type().
func (o TrashAlertV2Request) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrashAlertV2Request) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type TrashQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashQueryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashQueryRequest) {
}

func (to *TrashQueryRequest) SyncFieldsDuringRead(ctx context.Context, from TrashQueryRequest) {
}

func (c TrashQueryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
	Alert types.Object `tfsdk:"alert"`
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

func (to *UpdateAlertRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertRequest) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				// Recursively sync the fields of Alert
				toAlert.SyncFieldsDuringCreateOrUpdate(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (to *UpdateAlertRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertRequest) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (c UpdateAlertRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetOptional()
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
			"alert":                     o.Alert,
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"id":                        o.Id,
			"update_mask":               o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert":                     UpdateAlertRequestAlert{}.Type(ctx),
			"auto_resolve_display_name": types.BoolType,
			"id":                        types.StringType,
			"update_mask":               types.StringType,
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
	var v UpdateAlertRequestAlert
	d := o.Alert.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlert sets the value of the Alert field in UpdateAlertRequest.
func (o *UpdateAlertRequest) SetAlert(ctx context.Context, v UpdateAlertRequestAlert) {
	vs := v.ToObjectValue(ctx)
	o.Alert = vs
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition"`
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

func (to *UpdateAlertRequestAlert) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertRequestAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				// Recursively sync the fields of Condition
				toCondition.SyncFieldsDuringCreateOrUpdate(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (to *UpdateAlertRequestAlert) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertRequestAlert) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (c UpdateAlertRequestAlert) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["condition"] = attrs["condition"].SetOptional()
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
			"condition":            AlertCondition{}.Type(ctx),
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
	var v AlertCondition
	d := o.Condition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCondition sets the value of the Condition field in UpdateAlertRequestAlert.
func (o *UpdateAlertRequestAlert) SetCondition(ctx context.Context, v AlertCondition) {
	vs := v.ToObjectValue(ctx)
	o.Condition = vs
}

type UpdateAlertV2Request struct {
	Alert types.Object `tfsdk:"alert"`
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

func (to *UpdateAlertV2Request) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertV2Request) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				// Recursively sync the fields of Alert
				toAlert.SyncFieldsDuringCreateOrUpdate(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (to *UpdateAlertV2Request) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertV2Request) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (c UpdateAlertV2Request) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetRequired()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAlertV2Request) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertV2Request
// only implements ToObjectValue() and Type().
func (o UpdateAlertV2Request) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":       o.Alert,
			"id":          o.Id,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAlertV2Request) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert":       AlertV2{}.Type(ctx),
			"id":          types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetAlert returns the value of the Alert field in UpdateAlertV2Request as
// a AlertV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAlertV2Request) GetAlert(ctx context.Context) (AlertV2, bool) {
	var e AlertV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v AlertV2
	d := o.Alert.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlert sets the value of the Alert field in UpdateAlertV2Request.
func (o *UpdateAlertV2Request) SetAlert(ctx context.Context, v AlertV2) {
	vs := v.ToObjectValue(ctx)
	o.Alert = vs
}

type UpdateQueryRequest struct {
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Id types.String `tfsdk:"-"`

	Query types.Object `tfsdk:"query"`
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

func (to *UpdateQueryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateQueryRequest) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				// Recursively sync the fields of Query
				toQuery.SyncFieldsDuringCreateOrUpdate(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (to *UpdateQueryRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateQueryRequest) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (c UpdateQueryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
			"auto_resolve_display_name": o.AutoResolveDisplayName,
			"id":                        o.Id,
			"query":                     o.Query,
			"update_mask":               o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQueryRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_resolve_display_name": types.BoolType,
			"id":                        types.StringType,
			"query":                     UpdateQueryRequestQuery{}.Type(ctx),
			"update_mask":               types.StringType,
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
	var v UpdateQueryRequestQuery
	d := o.Query.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuery sets the value of the Query field in UpdateQueryRequest.
func (o *UpdateQueryRequest) SetQuery(ctx context.Context, v UpdateQueryRequestQuery) {
	vs := v.ToObjectValue(ctx)
	o.Query = vs
}

type UpdateQueryRequestQuery struct {
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

func (to *UpdateQueryRequestQuery) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateQueryRequestQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *UpdateQueryRequestQuery) SyncFieldsDuringRead(ctx context.Context, from UpdateQueryRequestQuery) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c UpdateQueryRequestQuery) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type UpdateResponse struct {
}

func (to *UpdateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateResponse) {
}

func (to *UpdateResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateResponse) {
}

func (c UpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

	Visualization types.Object `tfsdk:"visualization"`
}

func (to *UpdateVisualizationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVisualizationRequest) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				// Recursively sync the fields of Visualization
				toVisualization.SyncFieldsDuringCreateOrUpdate(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (to *UpdateVisualizationRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateVisualizationRequest) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				toVisualization.SyncFieldsDuringRead(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (c UpdateVisualizationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["visualization"] = attrs["visualization"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
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
			"id":            types.StringType,
			"update_mask":   types.StringType,
			"visualization": UpdateVisualizationRequestVisualization{}.Type(ctx),
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
	var v UpdateVisualizationRequestVisualization
	d := o.Visualization.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualization sets the value of the Visualization field in UpdateVisualizationRequest.
func (o *UpdateVisualizationRequest) SetVisualization(ctx context.Context, v UpdateVisualizationRequestVisualization) {
	vs := v.ToObjectValue(ctx)
	o.Visualization = vs
}

type UpdateVisualizationRequestVisualization struct {
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

func (to *UpdateVisualizationRequestVisualization) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVisualizationRequestVisualization) {
}

func (to *UpdateVisualizationRequestVisualization) SyncFieldsDuringRead(ctx context.Context, from UpdateVisualizationRequestVisualization) {
}

func (c UpdateVisualizationRequestVisualization) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type UpdateWidgetRequest struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`

	Options types.Object `tfsdk:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width"`
}

func (to *UpdateWidgetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWidgetRequest) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *UpdateWidgetRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWidgetRequest) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (c UpdateWidgetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["text"] = attrs["text"].SetOptional()
	attrs["visualization_id"] = attrs["visualization_id"].SetOptional()
	attrs["width"] = attrs["width"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWidgetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWidgetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWidgetRequest
// only implements ToObjectValue() and Type().
func (o UpdateWidgetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateWidgetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":     types.StringType,
			"id":               types.StringType,
			"options":          WidgetOptions{}.Type(ctx),
			"text":             types.StringType,
			"visualization_id": types.StringType,
			"width":            types.Int64Type,
		},
	}
}

// GetOptions returns the value of the Options field in UpdateWidgetRequest as
// a WidgetOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWidgetRequest) GetOptions(ctx context.Context) (WidgetOptions, bool) {
	var e WidgetOptions
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return e, false
	}
	var v WidgetOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in UpdateWidgetRequest.
func (o *UpdateWidgetRequest) SetOptions(ctx context.Context, v WidgetOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

type User struct {
	Email types.String `tfsdk:"email"`

	Id types.Int64 `tfsdk:"id"`

	Name types.String `tfsdk:"name"`
}

func (to *User) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User) {
}

func (to *User) SyncFieldsDuringRead(ctx context.Context, from User) {
}

func (c User) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *Visualization) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Visualization) {
}

func (to *Visualization) SyncFieldsDuringRead(ctx context.Context, from Visualization) {
}

func (c Visualization) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *WarehouseAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseAccessControlRequest) {
}

func (to *WarehouseAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from WarehouseAccessControlRequest) {
}

func (c WarehouseAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *WarehouseAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *WarehouseAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from WarehouseAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (c WarehouseAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type WarehousePermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WarehousePermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *WarehousePermission) SyncFieldsDuringRead(ctx context.Context, from WarehousePermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (c WarehousePermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type WarehousePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *WarehousePermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WarehousePermissions) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c WarehousePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehousePermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WarehousePermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissionsDescription) {
}

func (to *WarehousePermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissionsDescription) {
}

func (c WarehousePermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *WarehousePermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WarehousePermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c WarehousePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
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
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *WarehouseTypePair) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseTypePair) {
}

func (to *WarehouseTypePair) SyncFieldsDuringRead(ctx context.Context, from WarehouseTypePair) {
}

func (c WarehouseTypePair) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Id types.String `tfsdk:"id"`

	Options types.Object `tfsdk:"options"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization types.Object `tfsdk:"visualization"`
	// Unused field.
	Width types.Int64 `tfsdk:"width"`
}

func (to *Widget) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Widget) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				// Recursively sync the fields of Visualization
				toVisualization.SyncFieldsDuringCreateOrUpdate(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (to *Widget) SyncFieldsDuringRead(ctx context.Context, from Widget) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				toVisualization.SyncFieldsDuringRead(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (c Widget) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["visualization"] = attrs["visualization"].SetOptional()
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
			"id":            types.StringType,
			"options":       WidgetOptions{}.Type(ctx),
			"visualization": LegacyVisualization{}.Type(ctx),
			"width":         types.Int64Type,
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
	var v WidgetOptions
	d := o.Options.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in Widget.
func (o *Widget) SetOptions(ctx context.Context, v WidgetOptions) {
	vs := v.ToObjectValue(ctx)
	o.Options = vs
}

// GetVisualization returns the value of the Visualization field in Widget as
// a LegacyVisualization value.
// If the field is unknown or null, the boolean return value is false.
func (o *Widget) GetVisualization(ctx context.Context) (LegacyVisualization, bool) {
	var e LegacyVisualization
	if o.Visualization.IsNull() || o.Visualization.IsUnknown() {
		return e, false
	}
	var v LegacyVisualization
	d := o.Visualization.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualization sets the value of the Visualization field in Widget.
func (o *Widget) SetVisualization(ctx context.Context, v LegacyVisualization) {
	vs := v.ToObjectValue(ctx)
	o.Visualization = vs
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt types.String `tfsdk:"created_at"`
	// Custom description of the widget
	Description types.String `tfsdk:"description"`
	// Whether this widget is hidden on the dashboard.
	IsHidden types.Bool `tfsdk:"is_hidden"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings types.Object `tfsdk:"parameter_mappings"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position types.Object `tfsdk:"position"`
	// Custom title of the widget
	Title types.String `tfsdk:"title"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (to *WidgetOptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WidgetOptions) {
	if !from.Position.IsNull() && !from.Position.IsUnknown() {
		if toPosition, ok := to.GetPosition(ctx); ok {
			if fromPosition, ok := from.GetPosition(ctx); ok {
				// Recursively sync the fields of Position
				toPosition.SyncFieldsDuringCreateOrUpdate(ctx, fromPosition)
				to.SetPosition(ctx, toPosition)
			}
		}
	}
}

func (to *WidgetOptions) SyncFieldsDuringRead(ctx context.Context, from WidgetOptions) {
	if !from.Position.IsNull() && !from.Position.IsUnknown() {
		if toPosition, ok := to.GetPosition(ctx); ok {
			if fromPosition, ok := from.GetPosition(ctx); ok {
				toPosition.SyncFieldsDuringRead(ctx, fromPosition)
				to.SetPosition(ctx, toPosition)
			}
		}
	}
}

func (c WidgetOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["is_hidden"] = attrs["is_hidden"].SetOptional()
	attrs["parameter_mappings"] = attrs["parameter_mappings"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
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
			"created_at":         o.CreatedAt,
			"description":        o.Description,
			"is_hidden":          o.IsHidden,
			"parameter_mappings": o.ParameterMappings,
			"position":           o.Position,
			"title":              o.Title,
			"updated_at":         o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WidgetOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":         types.StringType,
			"description":        types.StringType,
			"is_hidden":          types.BoolType,
			"parameter_mappings": types.ObjectType{},
			"position":           WidgetPosition{}.Type(ctx),
			"title":              types.StringType,
			"updated_at":         types.StringType,
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
	var v WidgetPosition
	d := o.Position.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPosition sets the value of the Position field in WidgetOptions.
func (o *WidgetOptions) SetPosition(ctx context.Context, v WidgetPosition) {
	vs := v.ToObjectValue(ctx)
	o.Position = vs
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	AutoHeight types.Bool `tfsdk:"auto_height"`
	// column in the dashboard grid. Values start with 0
	Col types.Int64 `tfsdk:"col"`
	// row in the dashboard grid. Values start with 0
	Row types.Int64 `tfsdk:"row"`
	// width of the widget measured in dashboard grid cells
	SizeX types.Int64 `tfsdk:"size_x"`
	// height of the widget measured in dashboard grid cells
	SizeY types.Int64 `tfsdk:"size_y"`
}

func (to *WidgetPosition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WidgetPosition) {
}

func (to *WidgetPosition) SyncFieldsDuringRead(ctx context.Context, from WidgetPosition) {
}

func (c WidgetPosition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_height"] = attrs["auto_height"].SetOptional()
	attrs["col"] = attrs["col"].SetOptional()
	attrs["row"] = attrs["row"].SetOptional()
	attrs["size_x"] = attrs["size_x"].SetOptional()
	attrs["size_y"] = attrs["size_y"].SetOptional()

	return attrs
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
			"auto_height": o.AutoHeight,
			"col":         o.Col,
			"row":         o.Row,
			"size_x":      o.SizeX,
			"size_y":      o.SizeY,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WidgetPosition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_height": types.BoolType,
			"col":         types.Int64Type,
			"row":         types.Int64Type,
			"size_x":      types.Int64Type,
			"size_y":      types.Int64Type,
		},
	}
}
