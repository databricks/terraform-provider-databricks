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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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

func (to *AccessControl_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccessControl_SdkV2) {
}

func (to *AccessControl_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AccessControl_SdkV2) {
}

func (m AccessControl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AccessControl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControl_SdkV2
// only implements ToObjectValue() and Type().
func (m AccessControl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":       m.GroupName,
			"permission_level": m.PermissionLevel,
			"user_name":        m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AccessControl_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Alert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Alert_SdkV2) {
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

func (to *Alert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Alert_SdkV2) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (m Alert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Alert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Alert_SdkV2
// only implements ToObjectValue() and Type().
func (m Alert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            m.Condition,
			"create_time":          m.CreateTime,
			"custom_body":          m.CustomBody,
			"custom_subject":       m.CustomSubject,
			"display_name":         m.DisplayName,
			"id":                   m.Id,
			"lifecycle_state":      m.LifecycleState,
			"notify_on_ok":         m.NotifyOnOk,
			"owner_user_name":      m.OwnerUserName,
			"parent_path":          m.ParentPath,
			"query_id":             m.QueryId,
			"seconds_to_retrigger": m.SecondsToRetrigger,
			"state":                m.State,
			"trigger_time":         m.TriggerTime,
			"update_time":          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Alert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Alert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if m.Condition.IsNull() || m.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := m.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in Alert_SdkV2.
func (m *Alert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	m.Condition = types.ListValueMust(t, vs)
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

func (to *AlertCondition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertCondition_SdkV2) {
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

func (to *AlertCondition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertCondition_SdkV2) {
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

func (m AlertCondition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertCondition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"operand":   reflect.TypeOf(AlertConditionOperand_SdkV2{}),
		"threshold": reflect.TypeOf(AlertConditionThreshold_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertCondition_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertCondition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"empty_result_state": m.EmptyResultState,
			"op":                 m.Op,
			"operand":            m.Operand,
			"threshold":          m.Threshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertCondition_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertCondition_SdkV2) GetOperand(ctx context.Context) (AlertConditionOperand_SdkV2, bool) {
	var e AlertConditionOperand_SdkV2
	if m.Operand.IsNull() || m.Operand.IsUnknown() {
		return e, false
	}
	var v []AlertConditionOperand_SdkV2
	d := m.Operand.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOperand sets the value of the Operand field in AlertCondition_SdkV2.
func (m *AlertCondition_SdkV2) SetOperand(ctx context.Context, v AlertConditionOperand_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["operand"]
	m.Operand = types.ListValueMust(t, vs)
}

// GetThreshold returns the value of the Threshold field in AlertCondition_SdkV2 as
// a AlertConditionThreshold_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertCondition_SdkV2) GetThreshold(ctx context.Context) (AlertConditionThreshold_SdkV2, bool) {
	var e AlertConditionThreshold_SdkV2
	if m.Threshold.IsNull() || m.Threshold.IsUnknown() {
		return e, false
	}
	var v []AlertConditionThreshold_SdkV2
	d := m.Threshold.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThreshold sets the value of the Threshold field in AlertCondition_SdkV2.
func (m *AlertCondition_SdkV2) SetThreshold(ctx context.Context, v AlertConditionThreshold_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["threshold"]
	m.Threshold = types.ListValueMust(t, vs)
}

type AlertConditionOperand_SdkV2 struct {
	Column types.List `tfsdk:"column"`
}

func (to *AlertConditionOperand_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConditionOperand_SdkV2) {
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

func (to *AlertConditionOperand_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertConditionOperand_SdkV2) {
	if !from.Column.IsNull() && !from.Column.IsUnknown() {
		if toColumn, ok := to.GetColumn(ctx); ok {
			if fromColumn, ok := from.GetColumn(ctx); ok {
				toColumn.SyncFieldsDuringRead(ctx, fromColumn)
				to.SetColumn(ctx, toColumn)
			}
		}
	}
}

func (m AlertConditionOperand_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertConditionOperand_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertOperandColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionOperand_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertConditionOperand_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": m.Column,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertConditionOperand_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertConditionOperand_SdkV2) GetColumn(ctx context.Context) (AlertOperandColumn_SdkV2, bool) {
	var e AlertOperandColumn_SdkV2
	if m.Column.IsNull() || m.Column.IsUnknown() {
		return e, false
	}
	var v []AlertOperandColumn_SdkV2
	d := m.Column.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetColumn sets the value of the Column field in AlertConditionOperand_SdkV2.
func (m *AlertConditionOperand_SdkV2) SetColumn(ctx context.Context, v AlertOperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["column"]
	m.Column = types.ListValueMust(t, vs)
}

type AlertConditionThreshold_SdkV2 struct {
	Value types.List `tfsdk:"value"`
}

func (to *AlertConditionThreshold_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertConditionThreshold_SdkV2) {
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

func (to *AlertConditionThreshold_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertConditionThreshold_SdkV2) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (m AlertConditionThreshold_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertConditionThreshold_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(AlertOperandValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertConditionThreshold_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertConditionThreshold_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertConditionThreshold_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertConditionThreshold_SdkV2) GetValue(ctx context.Context) (AlertOperandValue_SdkV2, bool) {
	var e AlertOperandValue_SdkV2
	if m.Value.IsNull() || m.Value.IsUnknown() {
		return e, false
	}
	var v []AlertOperandValue_SdkV2
	d := m.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in AlertConditionThreshold_SdkV2.
func (m *AlertConditionThreshold_SdkV2) SetValue(ctx context.Context, v AlertOperandValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	m.Value = types.ListValueMust(t, vs)
}

type AlertOperandColumn_SdkV2 struct {
	Name types.String `tfsdk:"name"`
}

func (to *AlertOperandColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOperandColumn_SdkV2) {
}

func (to *AlertOperandColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertOperandColumn_SdkV2) {
}

func (m AlertOperandColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertOperandColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertOperandColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertOperandColumn_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *AlertOperandValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOperandValue_SdkV2) {
}

func (to *AlertOperandValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertOperandValue_SdkV2) {
}

func (m AlertOperandValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertOperandValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOperandValue_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertOperandValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   m.BoolValue,
			"double_value": m.DoubleValue,
			"string_value": m.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertOperandValue_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *AlertOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertOptions_SdkV2) {
}

func (to *AlertOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertOptions_SdkV2) {
}

func (m AlertOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column":             m.Column,
			"custom_body":        m.CustomBody,
			"custom_subject":     m.CustomSubject,
			"empty_result_state": m.EmptyResultState,
			"muted":              m.Muted,
			"op":                 m.Op,
			"value":              m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertOptions_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *AlertQuery_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertQuery_SdkV2) {
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

func (to *AlertQuery_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertQuery_SdkV2) {
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

func (m AlertQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(QueryOptions_SdkV2{}),
		"tags":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertQuery_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":     m.CreatedAt,
			"data_source_id": m.DataSourceId,
			"description":    m.Description,
			"id":             m.Id,
			"is_archived":    m.IsArchived,
			"is_draft":       m.IsDraft,
			"is_safe":        m.IsSafe,
			"name":           m.Name,
			"options":        m.Options,
			"query":          m.Query,
			"tags":           m.Tags,
			"updated_at":     m.UpdatedAt,
			"user_id":        m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertQuery_SdkV2) GetOptions(ctx context.Context) (QueryOptions_SdkV2, bool) {
	var e QueryOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AlertQuery_SdkV2.
func (m *AlertQuery_SdkV2) SetOptions(ctx context.Context, v QueryOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in AlertQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in AlertQuery_SdkV2.
func (m *AlertQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
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
	// The actual identity that will be used to execute the alert. This is an
	// output-only field that shows the resolved run-as identity after applying
	// permissions and defaults.
	EffectiveRunAs types.List `tfsdk:"effective_run_as"`

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
	// Specifies the identity that will be used to run the alert. This field
	// allows you to configure alerts to run as a specific user or service
	// principal. - For user identity: Set `user_name` to the email of an active
	// workspace user. Users can only set this to their own email. - For service
	// principal: Set `service_principal_name` to the application ID. Requires
	// the `servicePrincipal/user` role. If not specified, the alert will run as
	// the request user.
	RunAs types.List `tfsdk:"run_as"`
	// The run as username or application ID of service principal. On Create and
	// Update, this field can be set to application ID of an active service
	// principal. Setting this field requires the servicePrincipal/user role.
	// Deprecated: Use `run_as` field instead. This field will be removed in a
	// future release.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`

	Schedule types.List `tfsdk:"schedule"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *AlertV2_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2_SdkV2) {
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

func (to *AlertV2_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2_SdkV2) {
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

func (m AlertV2_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["custom_description"] = attrs["custom_description"].SetOptional()
	attrs["custom_summary"] = attrs["custom_summary"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["effective_run_as"] = attrs["effective_run_as"].SetComputed()
	attrs["effective_run_as"] = attrs["effective_run_as"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["evaluation"] = attrs["evaluation"].SetRequired()
	attrs["evaluation"] = attrs["evaluation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["lifecycle_state"] = attrs["lifecycle_state"].SetComputed()
	attrs["owner_user_name"] = attrs["owner_user_name"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetRequired()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as"] = attrs["run_as"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetRequired()
	attrs["schedule"] = attrs["schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AlertV2_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_run_as": reflect.TypeOf(AlertV2RunAs_SdkV2{}),
		"evaluation":       reflect.TypeOf(AlertV2Evaluation_SdkV2{}),
		"run_as":           reflect.TypeOf(AlertV2RunAs_SdkV2{}),
		"schedule":         reflect.TypeOf(CronSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":        m.CreateTime,
			"custom_description": m.CustomDescription,
			"custom_summary":     m.CustomSummary,
			"display_name":       m.DisplayName,
			"effective_run_as":   m.EffectiveRunAs,
			"evaluation":         m.Evaluation,
			"id":                 m.Id,
			"lifecycle_state":    m.LifecycleState,
			"owner_user_name":    m.OwnerUserName,
			"parent_path":        m.ParentPath,
			"query_text":         m.QueryText,
			"run_as":             m.RunAs,
			"run_as_user_name":   m.RunAsUserName,
			"schedule":           m.Schedule,
			"update_time":        m.UpdateTime,
			"warehouse_id":       m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"custom_description": types.StringType,
			"custom_summary":     types.StringType,
			"display_name":       types.StringType,
			"effective_run_as": basetypes.ListType{
				ElemType: AlertV2RunAs_SdkV2{}.Type(ctx),
			},
			"evaluation": basetypes.ListType{
				ElemType: AlertV2Evaluation_SdkV2{}.Type(ctx),
			},
			"id":              types.StringType,
			"lifecycle_state": types.StringType,
			"owner_user_name": types.StringType,
			"parent_path":     types.StringType,
			"query_text":      types.StringType,
			"run_as": basetypes.ListType{
				ElemType: AlertV2RunAs_SdkV2{}.Type(ctx),
			},
			"run_as_user_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule_SdkV2{}.Type(ctx),
			},
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetEffectiveRunAs returns the value of the EffectiveRunAs field in AlertV2_SdkV2 as
// a AlertV2RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2_SdkV2) GetEffectiveRunAs(ctx context.Context) (AlertV2RunAs_SdkV2, bool) {
	var e AlertV2RunAs_SdkV2
	if m.EffectiveRunAs.IsNull() || m.EffectiveRunAs.IsUnknown() {
		return e, false
	}
	var v []AlertV2RunAs_SdkV2
	d := m.EffectiveRunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveRunAs sets the value of the EffectiveRunAs field in AlertV2_SdkV2.
func (m *AlertV2_SdkV2) SetEffectiveRunAs(ctx context.Context, v AlertV2RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_run_as"]
	m.EffectiveRunAs = types.ListValueMust(t, vs)
}

// GetEvaluation returns the value of the Evaluation field in AlertV2_SdkV2 as
// a AlertV2Evaluation_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2_SdkV2) GetEvaluation(ctx context.Context) (AlertV2Evaluation_SdkV2, bool) {
	var e AlertV2Evaluation_SdkV2
	if m.Evaluation.IsNull() || m.Evaluation.IsUnknown() {
		return e, false
	}
	var v []AlertV2Evaluation_SdkV2
	d := m.Evaluation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEvaluation sets the value of the Evaluation field in AlertV2_SdkV2.
func (m *AlertV2_SdkV2) SetEvaluation(ctx context.Context, v AlertV2Evaluation_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["evaluation"]
	m.Evaluation = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in AlertV2_SdkV2 as
// a AlertV2RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2_SdkV2) GetRunAs(ctx context.Context) (AlertV2RunAs_SdkV2, bool) {
	var e AlertV2RunAs_SdkV2
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v []AlertV2RunAs_SdkV2
	d := m.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in AlertV2_SdkV2.
func (m *AlertV2_SdkV2) SetRunAs(ctx context.Context, v AlertV2RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	m.RunAs = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in AlertV2_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
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

// SetSchedule sets the value of the Schedule field in AlertV2_SdkV2.
func (m *AlertV2_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	m.Schedule = types.ListValueMust(t, vs)
}

type AlertV2Evaluation_SdkV2 struct {
	// Operator used for comparison in alert evaluation.
	ComparisonOperator types.String `tfsdk:"comparison_operator"`
	// Alert state if result is empty. Please avoid setting this field to be
	// `UNKNOWN` because `UNKNOWN` state is planned to be deprecated.
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

func (to *AlertV2Evaluation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Evaluation_SdkV2) {
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

func (to *AlertV2Evaluation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2Evaluation_SdkV2) {
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

func (m AlertV2Evaluation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comparison_operator"] = attrs["comparison_operator"].SetRequired()
	attrs["empty_result_state"] = attrs["empty_result_state"].SetOptional()
	attrs["last_evaluated_at"] = attrs["last_evaluated_at"].SetComputed()
	attrs["notification"] = attrs["notification"].SetOptional()
	attrs["notification"] = attrs["notification"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].SetRequired()
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
func (m AlertV2Evaluation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notification": reflect.TypeOf(AlertV2Notification_SdkV2{}),
		"source":       reflect.TypeOf(AlertV2OperandColumn_SdkV2{}),
		"threshold":    reflect.TypeOf(AlertV2Operand_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Evaluation_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2Evaluation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comparison_operator": m.ComparisonOperator,
			"empty_result_state":  m.EmptyResultState,
			"last_evaluated_at":   m.LastEvaluatedAt,
			"notification":        m.Notification,
			"source":              m.Source,
			"state":               m.State,
			"threshold":           m.Threshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2Evaluation_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertV2Evaluation_SdkV2) GetNotification(ctx context.Context) (AlertV2Notification_SdkV2, bool) {
	var e AlertV2Notification_SdkV2
	if m.Notification.IsNull() || m.Notification.IsUnknown() {
		return e, false
	}
	var v []AlertV2Notification_SdkV2
	d := m.Notification.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotification sets the value of the Notification field in AlertV2Evaluation_SdkV2.
func (m *AlertV2Evaluation_SdkV2) SetNotification(ctx context.Context, v AlertV2Notification_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notification"]
	m.Notification = types.ListValueMust(t, vs)
}

// GetSource returns the value of the Source field in AlertV2Evaluation_SdkV2 as
// a AlertV2OperandColumn_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2Evaluation_SdkV2) GetSource(ctx context.Context) (AlertV2OperandColumn_SdkV2, bool) {
	var e AlertV2OperandColumn_SdkV2
	if m.Source.IsNull() || m.Source.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandColumn_SdkV2
	d := m.Source.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSource sets the value of the Source field in AlertV2Evaluation_SdkV2.
func (m *AlertV2Evaluation_SdkV2) SetSource(ctx context.Context, v AlertV2OperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["source"]
	m.Source = types.ListValueMust(t, vs)
}

// GetThreshold returns the value of the Threshold field in AlertV2Evaluation_SdkV2 as
// a AlertV2Operand_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2Evaluation_SdkV2) GetThreshold(ctx context.Context) (AlertV2Operand_SdkV2, bool) {
	var e AlertV2Operand_SdkV2
	if m.Threshold.IsNull() || m.Threshold.IsUnknown() {
		return e, false
	}
	var v []AlertV2Operand_SdkV2
	d := m.Threshold.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThreshold sets the value of the Threshold field in AlertV2Evaluation_SdkV2.
func (m *AlertV2Evaluation_SdkV2) SetThreshold(ctx context.Context, v AlertV2Operand_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["threshold"]
	m.Threshold = types.ListValueMust(t, vs)
}

type AlertV2Notification_SdkV2 struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk          types.Bool `tfsdk:"notify_on_ok"`
	EffectiveNotifyOnOk types.Bool `tfsdk:"effective_notify_on_ok"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	RetriggerSeconds          types.Int64 `tfsdk:"retrigger_seconds"`
	EffectiveRetriggerSeconds types.Int64 `tfsdk:"effective_retrigger_seconds"`

	Subscriptions types.Set `tfsdk:"subscriptions"`
}

func (to *AlertV2Notification_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Notification_SdkV2) {
	to.EffectiveNotifyOnOk = to.NotifyOnOk
	to.NotifyOnOk = from.NotifyOnOk
	to.EffectiveRetriggerSeconds = to.RetriggerSeconds
	to.RetriggerSeconds = from.RetriggerSeconds
	if !from.Subscriptions.IsNull() && !from.Subscriptions.IsUnknown() && to.Subscriptions.IsNull() && len(from.Subscriptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subscriptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subscriptions = from.Subscriptions
	}
}

func (to *AlertV2Notification_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2Notification_SdkV2) {
	to.EffectiveNotifyOnOk = from.EffectiveNotifyOnOk
	if from.EffectiveNotifyOnOk.ValueBool() == to.NotifyOnOk.ValueBool() {
		to.NotifyOnOk = from.NotifyOnOk
	}
	to.EffectiveRetriggerSeconds = from.EffectiveRetriggerSeconds
	if from.EffectiveRetriggerSeconds.ValueInt64() == to.RetriggerSeconds.ValueInt64() {
		to.RetriggerSeconds = from.RetriggerSeconds
	}
	if !from.Subscriptions.IsNull() && !from.Subscriptions.IsUnknown() && to.Subscriptions.IsNull() && len(from.Subscriptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subscriptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subscriptions = from.Subscriptions
	}
}

func (m AlertV2Notification_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_notify_on_ok"] = attrs["effective_notify_on_ok"].SetComputed()
	attrs["notify_on_ok"] = attrs["notify_on_ok"].SetOptional()
	attrs["effective_retrigger_seconds"] = attrs["effective_retrigger_seconds"].SetComputed()
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
func (m AlertV2Notification_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(AlertV2Subscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Notification_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2Notification_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notify_on_ok": m.NotifyOnOk, "effective_notify_on_ok": m.EffectiveNotifyOnOk,
			"retrigger_seconds": m.RetriggerSeconds, "effective_retrigger_seconds": m.EffectiveRetriggerSeconds,
			"subscriptions": m.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2Notification_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notify_on_ok":                types.BoolType,
			"effective_notify_on_ok":      types.BoolType,
			"retrigger_seconds":           types.Int64Type,
			"effective_retrigger_seconds": types.Int64Type,
			"subscriptions": basetypes.SetType{
				ElemType: AlertV2Subscription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in AlertV2Notification_SdkV2 as
// a slice of AlertV2Subscription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2Notification_SdkV2) GetSubscriptions(ctx context.Context) ([]AlertV2Subscription_SdkV2, bool) {
	if m.Subscriptions.IsNull() || m.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []AlertV2Subscription_SdkV2
	d := m.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in AlertV2Notification_SdkV2.
func (m *AlertV2Notification_SdkV2) SetSubscriptions(ctx context.Context, v []AlertV2Subscription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Subscriptions = types.SetValueMust(t, vs)
}

type AlertV2Operand_SdkV2 struct {
	Column types.List `tfsdk:"column"`

	Value types.List `tfsdk:"value"`
}

func (to *AlertV2Operand_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Operand_SdkV2) {
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

func (to *AlertV2Operand_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2Operand_SdkV2) {
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

func (m AlertV2Operand_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertV2Operand_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column": reflect.TypeOf(AlertV2OperandColumn_SdkV2{}),
		"value":  reflect.TypeOf(AlertV2OperandValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Operand_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2Operand_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column": m.Column,
			"value":  m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2Operand_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *AlertV2Operand_SdkV2) GetColumn(ctx context.Context) (AlertV2OperandColumn_SdkV2, bool) {
	var e AlertV2OperandColumn_SdkV2
	if m.Column.IsNull() || m.Column.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandColumn_SdkV2
	d := m.Column.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetColumn sets the value of the Column field in AlertV2Operand_SdkV2.
func (m *AlertV2Operand_SdkV2) SetColumn(ctx context.Context, v AlertV2OperandColumn_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["column"]
	m.Column = types.ListValueMust(t, vs)
}

// GetValue returns the value of the Value field in AlertV2Operand_SdkV2 as
// a AlertV2OperandValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AlertV2Operand_SdkV2) GetValue(ctx context.Context) (AlertV2OperandValue_SdkV2, bool) {
	var e AlertV2OperandValue_SdkV2
	if m.Value.IsNull() || m.Value.IsUnknown() {
		return e, false
	}
	var v []AlertV2OperandValue_SdkV2
	d := m.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in AlertV2Operand_SdkV2.
func (m *AlertV2Operand_SdkV2) SetValue(ctx context.Context, v AlertV2OperandValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	m.Value = types.ListValueMust(t, vs)
}

type AlertV2OperandColumn_SdkV2 struct {
	Aggregation types.String `tfsdk:"aggregation"`

	Display types.String `tfsdk:"display"`

	Name types.String `tfsdk:"name"`
}

func (to *AlertV2OperandColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2OperandColumn_SdkV2) {
}

func (to *AlertV2OperandColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2OperandColumn_SdkV2) {
}

func (m AlertV2OperandColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aggregation"] = attrs["aggregation"].SetOptional()
	attrs["display"] = attrs["display"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AlertV2OperandColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AlertV2OperandColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2OperandColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aggregation": m.Aggregation,
			"display":     m.Display,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2OperandColumn_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *AlertV2OperandValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2OperandValue_SdkV2) {
}

func (to *AlertV2OperandValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2OperandValue_SdkV2) {
}

func (m AlertV2OperandValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertV2OperandValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2OperandValue_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2OperandValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   m.BoolValue,
			"double_value": m.DoubleValue,
			"string_value": m.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2OperandValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"double_value": types.Float64Type,
			"string_value": types.StringType,
		},
	}
}

type AlertV2RunAs_SdkV2 struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Can only set this field to their
	// own email.
	UserName types.String `tfsdk:"user_name"`
}

func (to *AlertV2RunAs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2RunAs_SdkV2) {
}

func (to *AlertV2RunAs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2RunAs_SdkV2) {
}

func (m AlertV2RunAs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertV2RunAs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2RunAs_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2RunAs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2RunAs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AlertV2Subscription_SdkV2 struct {
	DestinationId types.String `tfsdk:"destination_id"`

	UserEmail types.String `tfsdk:"user_email"`
}

func (to *AlertV2Subscription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AlertV2Subscription_SdkV2) {
}

func (to *AlertV2Subscription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AlertV2Subscription_SdkV2) {
}

func (m AlertV2Subscription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AlertV2Subscription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Subscription_SdkV2
// only implements ToObjectValue() and Type().
func (m AlertV2Subscription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": m.DestinationId,
			"user_email":     m.UserEmail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AlertV2Subscription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
			"user_email":     types.StringType,
		},
	}
}

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

func (to *BaseChunkInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BaseChunkInfo_SdkV2) {
}

func (to *BaseChunkInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BaseChunkInfo_SdkV2) {
}

func (m BaseChunkInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BaseChunkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseChunkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m BaseChunkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":  m.ByteCount,
			"chunk_index": m.ChunkIndex,
			"row_count":   m.RowCount,
			"row_offset":  m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BaseChunkInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CancelExecutionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelExecutionRequest_SdkV2) {
}

func (to *CancelExecutionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CancelExecutionRequest_SdkV2) {
}

func (m CancelExecutionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CancelExecutionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelExecutionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CancelExecutionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelExecutionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
		},
	}
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel_SdkV2 struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version"`

	Name types.String `tfsdk:"name"`
}

func (to *Channel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Channel_SdkV2) {
}

func (to *Channel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Channel_SdkV2) {
}

func (m Channel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Channel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Channel_SdkV2
// only implements ToObjectValue() and Type().
func (m Channel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": m.DbsqlVersion,
			"name":          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Channel_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ChannelInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChannelInfo_SdkV2) {
}

func (to *ChannelInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ChannelInfo_SdkV2) {
}

func (m ChannelInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ChannelInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChannelInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ChannelInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbsql_version": m.DbsqlVersion,
			"name":          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ChannelInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ClientConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClientConfig_SdkV2) {
}

func (to *ClientConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClientConfig_SdkV2) {
}

func (m ClientConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ClientConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClientConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m ClientConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_custom_js_visualizations":   m.AllowCustomJsVisualizations,
			"allow_downloads":                  m.AllowDownloads,
			"allow_external_shares":            m.AllowExternalShares,
			"allow_subscriptions":              m.AllowSubscriptions,
			"date_format":                      m.DateFormat,
			"date_time_format":                 m.DateTimeFormat,
			"disable_publish":                  m.DisablePublish,
			"enable_legacy_autodetect_types":   m.EnableLegacyAutodetectTypes,
			"feature_show_permissions_control": m.FeatureShowPermissionsControl,
			"hide_plotly_mode_bar":             m.HidePlotlyModeBar,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClientConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ColumnInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (m ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":               m.Name,
			"position":           m.Position,
			"type_interval_type": m.TypeIntervalType,
			"type_name":          m.TypeName,
			"type_precision":     m.TypePrecision,
			"type_scale":         m.TypeScale,
			"type_text":          m.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CreateAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlert_SdkV2) {
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

func (to *CreateAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAlert_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m CreateAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     m.Name,
			"options":  m.Options,
			"parent":   m.Parent,
			"query_id": m.QueryId,
			"rearm":    m.Rearm,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateAlert_SdkV2.
func (m *CreateAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

type CreateAlertRequest_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`
}

func (to *CreateAlertRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertRequest_SdkV2) {
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

func (to *CreateAlertRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAlertRequest_SdkV2) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (m CreateAlertRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetOptional()
	attrs["alert"] = attrs["alert"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(CreateAlertRequestAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":                     m.Alert,
			"auto_resolve_display_name": m.AutoResolveDisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateAlertRequest_SdkV2) GetAlert(ctx context.Context) (CreateAlertRequestAlert_SdkV2, bool) {
	var e CreateAlertRequestAlert_SdkV2
	if m.Alert.IsNull() || m.Alert.IsUnknown() {
		return e, false
	}
	var v []CreateAlertRequestAlert_SdkV2
	d := m.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in CreateAlertRequest_SdkV2.
func (m *CreateAlertRequest_SdkV2) SetAlert(ctx context.Context, v CreateAlertRequestAlert_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	m.Alert = types.ListValueMust(t, vs)
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

func (to *CreateAlertRequestAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertRequestAlert_SdkV2) {
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

func (to *CreateAlertRequestAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAlertRequestAlert_SdkV2) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (m CreateAlertRequestAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateAlertRequestAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertRequestAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateAlertRequestAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            m.Condition,
			"custom_body":          m.CustomBody,
			"custom_subject":       m.CustomSubject,
			"display_name":         m.DisplayName,
			"notify_on_ok":         m.NotifyOnOk,
			"parent_path":          m.ParentPath,
			"query_id":             m.QueryId,
			"seconds_to_retrigger": m.SecondsToRetrigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAlertRequestAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateAlertRequestAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if m.Condition.IsNull() || m.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := m.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in CreateAlertRequestAlert_SdkV2.
func (m *CreateAlertRequestAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	m.Condition = types.ListValueMust(t, vs)
}

type CreateAlertV2Request_SdkV2 struct {
	Alert types.List `tfsdk:"alert"`
}

func (to *CreateAlertV2Request_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAlertV2Request_SdkV2) {
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

func (to *CreateAlertV2Request_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAlertV2Request_SdkV2) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (m CreateAlertV2Request_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetRequired()
	attrs["alert"] = attrs["alert"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAlertV2Request.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert": m.Alert,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateAlertV2Request_SdkV2) GetAlert(ctx context.Context) (AlertV2_SdkV2, bool) {
	var e AlertV2_SdkV2
	if m.Alert.IsNull() || m.Alert.IsUnknown() {
		return e, false
	}
	var v []AlertV2_SdkV2
	d := m.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in CreateAlertV2Request_SdkV2.
func (m *CreateAlertV2Request_SdkV2) SetAlert(ctx context.Context, v AlertV2_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	m.Alert = types.ListValueMust(t, vs)
}

type CreateQueryRequest_SdkV2 struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName types.Bool `tfsdk:"auto_resolve_display_name"`

	Query types.List `tfsdk:"query"`
}

func (to *CreateQueryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryRequest_SdkV2) {
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

func (to *CreateQueryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateQueryRequest_SdkV2) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (m CreateQueryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query"] = attrs["query"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(CreateQueryRequestQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_resolve_display_name": m.AutoResolveDisplayName,
			"query":                     m.Query,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateQueryRequest_SdkV2) GetQuery(ctx context.Context) (CreateQueryRequestQuery_SdkV2, bool) {
	var e CreateQueryRequestQuery_SdkV2
	if m.Query.IsNull() || m.Query.IsUnknown() {
		return e, false
	}
	var v []CreateQueryRequestQuery_SdkV2
	d := m.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in CreateQueryRequest_SdkV2.
func (m *CreateQueryRequest_SdkV2) SetQuery(ctx context.Context, v CreateQueryRequestQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	m.Query = types.ListValueMust(t, vs)
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

func (to *CreateQueryRequestQuery_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryRequestQuery_SdkV2) {
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

func (to *CreateQueryRequestQuery_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateQueryRequestQuery_SdkV2) {
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

func (m CreateQueryRequestQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateQueryRequestQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryRequestQuery_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateQueryRequestQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit": m.ApplyAutoLimit,
			"catalog":          m.Catalog,
			"description":      m.Description,
			"display_name":     m.DisplayName,
			"parameters":       m.Parameters,
			"parent_path":      m.ParentPath,
			"query_text":       m.QueryText,
			"run_as_mode":      m.RunAsMode,
			"schema":           m.Schema,
			"tags":             m.Tags,
			"warehouse_id":     m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateQueryRequestQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateQueryRequestQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in CreateQueryRequestQuery_SdkV2.
func (m *CreateQueryRequestQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateQueryRequestQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateQueryRequestQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateQueryRequestQuery_SdkV2.
func (m *CreateQueryRequestQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
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

func (to *CreateQueryVisualizationsLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQueryVisualizationsLegacyRequest_SdkV2) {
}

func (to *CreateQueryVisualizationsLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateQueryVisualizationsLegacyRequest_SdkV2) {
}

func (m CreateQueryVisualizationsLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateQueryVisualizationsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQueryVisualizationsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateQueryVisualizationsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"options":     m.Options,
			"query_id":    m.QueryId,
			"type":        m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateQueryVisualizationsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CreateVisualizationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVisualizationRequest_SdkV2) {
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

func (to *CreateVisualizationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateVisualizationRequest_SdkV2) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				toVisualization.SyncFieldsDuringRead(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (m CreateVisualizationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["visualization"] = attrs["visualization"].SetOptional()
	attrs["visualization"] = attrs["visualization"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVisualizationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(CreateVisualizationRequestVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"visualization": m.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateVisualizationRequest_SdkV2) GetVisualization(ctx context.Context) (CreateVisualizationRequestVisualization_SdkV2, bool) {
	var e CreateVisualizationRequestVisualization_SdkV2
	if m.Visualization.IsNull() || m.Visualization.IsUnknown() {
		return e, false
	}
	var v []CreateVisualizationRequestVisualization_SdkV2
	d := m.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in CreateVisualizationRequest_SdkV2.
func (m *CreateVisualizationRequest_SdkV2) SetVisualization(ctx context.Context, v CreateVisualizationRequestVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	m.Visualization = types.ListValueMust(t, vs)
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

func (to *CreateVisualizationRequestVisualization_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVisualizationRequestVisualization_SdkV2) {
}

func (to *CreateVisualizationRequestVisualization_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateVisualizationRequestVisualization_SdkV2) {
}

func (m CreateVisualizationRequestVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateVisualizationRequestVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVisualizationRequestVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateVisualizationRequestVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          m.DisplayName,
			"query_id":              m.QueryId,
			"serialized_options":    m.SerializedOptions,
			"serialized_query_plan": m.SerializedQueryPlan,
			"type":                  m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateVisualizationRequestVisualization_SdkV2) Type(ctx context.Context) attr.Type {
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

// Creates a new SQL warehouse.
type CreateWarehouseRequest_SdkV2 struct {
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
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
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
	// Configurations whether the endpoint should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
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

func (to *CreateWarehouseRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWarehouseRequest_SdkV2) {
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

func (to *CreateWarehouseRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWarehouseRequest_SdkV2) {
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

func (m CreateWarehouseRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["tags"] = attrs["tags"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel_SdkV2{}),
		"tags":    reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            m.AutoStopMins,
			"channel":                   m.Channel,
			"cluster_size":              m.ClusterSize,
			"creator_name":              m.CreatorName,
			"enable_photon":             m.EnablePhoton,
			"enable_serverless_compute": m.EnableServerlessCompute,
			"instance_profile_arn":      m.InstanceProfileArn,
			"max_num_clusters":          m.MaxNumClusters,
			"min_num_clusters":          m.MinNumClusters,
			"name":                      m.Name,
			"spot_instance_policy":      m.SpotInstancePolicy,
			"tags":                      m.Tags,
			"warehouse_type":            m.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateWarehouseRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in CreateWarehouseRequest_SdkV2.
func (m *CreateWarehouseRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateWarehouseRequest_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWarehouseRequest_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in CreateWarehouseRequest_SdkV2.
func (m *CreateWarehouseRequest_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	m.Tags = types.ListValueMust(t, vs)
}

type CreateWarehouseResponse_SdkV2 struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id"`
}

func (to *CreateWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWarehouseResponse_SdkV2) {
}

func (to *CreateWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWarehouseResponse_SdkV2) {
}

func (m CreateWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CreateWidget_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWidget_SdkV2) {
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

func (to *CreateWidget_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWidget_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m CreateWidget_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateWidget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWidget_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWidget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":     m.DashboardId,
			"options":          m.Options,
			"text":             m.Text,
			"visualization_id": m.VisualizationId,
			"width":            m.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWidget_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateWidget_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in CreateWidget_SdkV2.
func (m *CreateWidget_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
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

func (to *CronSchedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronSchedule_SdkV2) {
}

func (to *CronSchedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CronSchedule_SdkV2) {
}

func (m CronSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pause_status"] = attrs["pause_status"].SetOptional()
	attrs["quartz_cron_schedule"] = attrs["quartz_cron_schedule"].SetRequired()
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
			"pause_status":         m.PauseStatus,
			"quartz_cron_schedule": m.QuartzCronSchedule,
			"timezone_id":          m.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Dashboard_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Dashboard_SdkV2) {
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

func (to *Dashboard_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Dashboard_SdkV2) {
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

func (m Dashboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Dashboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m Dashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_edit":                  m.CanEdit,
			"created_at":                m.CreatedAt,
			"dashboard_filters_enabled": m.DashboardFiltersEnabled,
			"id":                        m.Id,
			"is_archived":               m.IsArchived,
			"is_draft":                  m.IsDraft,
			"is_favorite":               m.IsFavorite,
			"name":                      m.Name,
			"options":                   m.Options,
			"parent":                    m.Parent,
			"permission_tier":           m.PermissionTier,
			"slug":                      m.Slug,
			"tags":                      m.Tags,
			"updated_at":                m.UpdatedAt,
			"user":                      m.User,
			"user_id":                   m.UserId,
			"widgets":                   m.Widgets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Dashboard_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Dashboard_SdkV2) GetOptions(ctx context.Context) (DashboardOptions_SdkV2, bool) {
	var e DashboardOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []DashboardOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Dashboard_SdkV2.
func (m *Dashboard_SdkV2) SetOptions(ctx context.Context, v DashboardOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Dashboard_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Dashboard_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Dashboard_SdkV2.
func (m *Dashboard_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in Dashboard_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Dashboard_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := m.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in Dashboard_SdkV2.
func (m *Dashboard_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

// GetWidgets returns the value of the Widgets field in Dashboard_SdkV2 as
// a slice of Widget_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Dashboard_SdkV2) GetWidgets(ctx context.Context) ([]Widget_SdkV2, bool) {
	if m.Widgets.IsNull() || m.Widgets.IsUnknown() {
		return nil, false
	}
	var v []Widget_SdkV2
	d := m.Widgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWidgets sets the value of the Widgets field in Dashboard_SdkV2.
func (m *Dashboard_SdkV2) SetWidgets(ctx context.Context, v []Widget_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["widgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Widgets = types.ListValueMust(t, vs)
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

func (to *DashboardEditContent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardEditContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *DashboardEditContent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DashboardEditContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m DashboardEditContent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DashboardEditContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEditContent_SdkV2
// only implements ToObjectValue() and Type().
func (m DashboardEditContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
			"name":         m.Name,
			"run_as_role":  m.RunAsRole,
			"tags":         m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DashboardEditContent_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *DashboardEditContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DashboardEditContent_SdkV2.
func (m *DashboardEditContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type DashboardOptions_SdkV2 struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`
}

func (to *DashboardOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardOptions_SdkV2) {
}

func (to *DashboardOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DashboardOptions_SdkV2) {
}

func (m DashboardOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DashboardOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m DashboardOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"moved_to_trash_at": m.MovedToTrashAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DashboardOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"moved_to_trash_at": types.StringType,
		},
	}
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

func (to *DataSource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataSource_SdkV2) {
}

func (to *DataSource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataSource_SdkV2) {
}

func (m DataSource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataSource_SdkV2
// only implements ToObjectValue() and Type().
func (m DataSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":                  m.Id,
			"name":                m.Name,
			"pause_reason":        m.PauseReason,
			"paused":              m.Paused,
			"supports_auto_limit": m.SupportsAutoLimit,
			"syntax":              m.Syntax,
			"type":                m.Type_,
			"view_only":           m.ViewOnly,
			"warehouse_id":        m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataSource_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DateRange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateRange_SdkV2) {
}

func (to *DateRange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DateRange_SdkV2) {
}

func (m DateRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DateRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRange_SdkV2
// only implements ToObjectValue() and Type().
func (m DateRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end":   m.End,
			"start": m.Start,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DateRange_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DateRangeValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateRangeValue_SdkV2) {
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

func (to *DateRangeValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DateRangeValue_SdkV2) {
	if !from.DateRangeValue.IsNull() && !from.DateRangeValue.IsUnknown() {
		if toDateRangeValue, ok := to.GetDateRangeValue(ctx); ok {
			if fromDateRangeValue, ok := from.GetDateRangeValue(ctx); ok {
				toDateRangeValue.SyncFieldsDuringRead(ctx, fromDateRangeValue)
				to.SetDateRangeValue(ctx, toDateRangeValue)
			}
		}
	}
}

func (m DateRangeValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DateRangeValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"date_range_value": reflect.TypeOf(DateRange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateRangeValue_SdkV2
// only implements ToObjectValue() and Type().
func (m DateRangeValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_range_value":         m.DateRangeValue,
			"dynamic_date_range_value": m.DynamicDateRangeValue,
			"precision":                m.Precision,
			"start_day_of_week":        m.StartDayOfWeek,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DateRangeValue_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *DateRangeValue_SdkV2) GetDateRangeValue(ctx context.Context) (DateRange_SdkV2, bool) {
	var e DateRange_SdkV2
	if m.DateRangeValue.IsNull() || m.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRange_SdkV2
	d := m.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in DateRangeValue_SdkV2.
func (m *DateRangeValue_SdkV2) SetDateRangeValue(ctx context.Context, v DateRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	m.DateRangeValue = types.ListValueMust(t, vs)
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

func (to *DateValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DateValue_SdkV2) {
}

func (to *DateValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DateValue_SdkV2) {
}

func (m DateValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DateValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DateValue_SdkV2
// only implements ToObjectValue() and Type().
func (m DateValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_value":         m.DateValue,
			"dynamic_date_value": m.DynamicDateValue,
			"precision":          m.Precision,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DateValue_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteAlertsLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAlertsLegacyRequest_SdkV2) {
}

func (to *DeleteAlertsLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAlertsLegacyRequest_SdkV2) {
}

func (m DeleteAlertsLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAlertsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAlertsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAlertsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": m.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAlertsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

type DeleteDashboardRequest_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *DeleteDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardRequest_SdkV2) {
}

func (to *DeleteDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardRequest_SdkV2) {
}

func (m DeleteDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteDashboardWidgetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardWidgetRequest_SdkV2) {
}

func (to *DeleteDashboardWidgetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardWidgetRequest_SdkV2) {
}

func (m DeleteDashboardWidgetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDashboardWidgetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardWidgetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDashboardWidgetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardWidgetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *DeleteQueriesLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteQueriesLegacyRequest_SdkV2) {
}

func (to *DeleteQueriesLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteQueriesLegacyRequest_SdkV2) {
}

func (m DeleteQueriesLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": m.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteQueryVisualizationsLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteQueryVisualizationsLegacyRequest_SdkV2) {
}

func (to *DeleteQueryVisualizationsLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteQueryVisualizationsLegacyRequest_SdkV2) {
}

func (m DeleteQueryVisualizationsLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteQueryVisualizationsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQueryVisualizationsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteQueryVisualizationsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteQueryVisualizationsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (m DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteVisualizationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteVisualizationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteVisualizationRequest_SdkV2) {
}

func (to *DeleteVisualizationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteVisualizationRequest_SdkV2) {
}

func (m DeleteVisualizationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteWarehouseRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWarehouseRequest_SdkV2) {
}

func (to *DeleteWarehouseRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWarehouseRequest_SdkV2) {
}

func (m DeleteWarehouseRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWarehouseResponse_SdkV2 struct {
}

func (to *DeleteWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWarehouseResponse_SdkV2) {
}

func (to *DeleteWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWarehouseResponse_SdkV2) {
}

func (m DeleteWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *EditAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditAlert_SdkV2) {
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

func (to *EditAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EditAlert_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m EditAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EditAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m EditAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": m.AlertId,
			"name":     m.Name,
			"options":  m.Options,
			"query_id": m.QueryId,
			"rearm":    m.Rearm,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EditAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in EditAlert_SdkV2.
func (m *EditAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// This is an incremental edit functionality, so all fields except id are
// optional. If a field is set, the corresponding configuration in the SQL
// warehouse is modified. If a field is unset, the existing configuration value
// in the SQL warehouse is retained. Thus, this API is not idempotent.
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
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Required. Id of the warehouse to configure.
	Id types.String `tfsdk:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
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
	// Configurations whether the endpoint should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
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

func (to *EditWarehouseRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditWarehouseRequest_SdkV2) {
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

func (to *EditWarehouseRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EditWarehouseRequest_SdkV2) {
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

func (m EditWarehouseRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
	attrs["tags"] = attrs["tags"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EditWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel": reflect.TypeOf(Channel_SdkV2{}),
		"tags":    reflect.TypeOf(EndpointTags_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m EditWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            m.AutoStopMins,
			"channel":                   m.Channel,
			"cluster_size":              m.ClusterSize,
			"creator_name":              m.CreatorName,
			"enable_photon":             m.EnablePhoton,
			"enable_serverless_compute": m.EnableServerlessCompute,
			"id":                        m.Id,
			"instance_profile_arn":      m.InstanceProfileArn,
			"max_num_clusters":          m.MaxNumClusters,
			"min_num_clusters":          m.MinNumClusters,
			"name":                      m.Name,
			"spot_instance_policy":      m.SpotInstancePolicy,
			"tags":                      m.Tags,
			"warehouse_type":            m.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EditWarehouseRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EditWarehouseRequest_SdkV2.
func (m *EditWarehouseRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EditWarehouseRequest_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditWarehouseRequest_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EditWarehouseRequest_SdkV2.
func (m *EditWarehouseRequest_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	m.Tags = types.ListValueMust(t, vs)
}

type EditWarehouseResponse_SdkV2 struct {
}

func (to *EditWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditWarehouseResponse_SdkV2) {
}

func (to *EditWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EditWarehouseResponse_SdkV2) {
}

func (m EditWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m EditWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty_SdkV2 struct {
}

func (to *Empty_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Empty_SdkV2) {
}

func (to *Empty_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Empty_SdkV2) {
}

func (m Empty_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Empty_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty_SdkV2
// only implements ToObjectValue() and Type().
func (m Empty_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m Empty_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EndpointConfPair_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *EndpointConfPair_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointConfPair_SdkV2) {
}

func (to *EndpointConfPair_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointConfPair_SdkV2) {
}

func (m EndpointConfPair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointConfPair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointConfPair_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointConfPair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointConfPair_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Health status of the endpoint.
	Status types.String `tfsdk:"status"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary"`
}

func (to *EndpointHealth_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointHealth_SdkV2) {
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

func (to *EndpointHealth_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointHealth_SdkV2) {
	if !from.FailureReason.IsNull() && !from.FailureReason.IsUnknown() {
		if toFailureReason, ok := to.GetFailureReason(ctx); ok {
			if fromFailureReason, ok := from.GetFailureReason(ctx); ok {
				toFailureReason.SyncFieldsDuringRead(ctx, fromFailureReason)
				to.SetFailureReason(ctx, toFailureReason)
			}
		}
	}
}

func (m EndpointHealth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["failure_reason"] = attrs["failure_reason"].SetOptional()
	attrs["failure_reason"] = attrs["failure_reason"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["message"] = attrs["message"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
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
func (m EndpointHealth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failure_reason": reflect.TypeOf(TerminationReason_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointHealth_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointHealth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":        m.Details,
			"failure_reason": m.FailureReason,
			"message":        m.Message,
			"status":         m.Status,
			"summary":        m.Summary,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointHealth_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EndpointHealth_SdkV2) GetFailureReason(ctx context.Context) (TerminationReason_SdkV2, bool) {
	var e TerminationReason_SdkV2
	if m.FailureReason.IsNull() || m.FailureReason.IsUnknown() {
		return e, false
	}
	var v []TerminationReason_SdkV2
	d := m.FailureReason.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailureReason sets the value of the FailureReason field in EndpointHealth_SdkV2.
func (m *EndpointHealth_SdkV2) SetFailureReason(ctx context.Context, v TerminationReason_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failure_reason"]
	m.FailureReason = types.ListValueMust(t, vs)
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
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
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
	// Configurations whether the endpoint should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
	// state of the endpoint
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

func (to *EndpointInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointInfo_SdkV2) {
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

func (to *EndpointInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointInfo_SdkV2) {
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

func (m EndpointInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetComputed()
	attrs["health"] = attrs["health"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["id"] = attrs["id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
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
func (m EndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            m.AutoStopMins,
			"channel":                   m.Channel,
			"cluster_size":              m.ClusterSize,
			"creator_name":              m.CreatorName,
			"enable_photon":             m.EnablePhoton,
			"enable_serverless_compute": m.EnableServerlessCompute,
			"health":                    m.Health,
			"id":                        m.Id,
			"instance_profile_arn":      m.InstanceProfileArn,
			"jdbc_url":                  m.JdbcUrl,
			"max_num_clusters":          m.MaxNumClusters,
			"min_num_clusters":          m.MinNumClusters,
			"name":                      m.Name,
			"num_active_sessions":       m.NumActiveSessions,
			"num_clusters":              m.NumClusters,
			"odbc_params":               m.OdbcParams,
			"spot_instance_policy":      m.SpotInstancePolicy,
			"state":                     m.State,
			"tags":                      m.Tags,
			"warehouse_type":            m.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EndpointInfo_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in EndpointInfo_SdkV2 as
// a EndpointHealth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo_SdkV2) GetHealth(ctx context.Context) (EndpointHealth_SdkV2, bool) {
	var e EndpointHealth_SdkV2
	if m.Health.IsNull() || m.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth_SdkV2
	d := m.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetHealth(ctx context.Context, v EndpointHealth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	m.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in EndpointInfo_SdkV2 as
// a OdbcParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo_SdkV2) GetOdbcParams(ctx context.Context) (OdbcParams_SdkV2, bool) {
	var e OdbcParams_SdkV2
	if m.OdbcParams.IsNull() || m.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams_SdkV2
	d := m.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetOdbcParams(ctx context.Context, v OdbcParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	m.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EndpointInfo_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	m.Tags = types.ListValueMust(t, vs)
}

type EndpointTagPair_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *EndpointTagPair_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTagPair_SdkV2) {
}

func (to *EndpointTagPair_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointTagPair_SdkV2) {
}

func (m EndpointTagPair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointTagPair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTagPair_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointTagPair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointTagPair_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *EndpointTags_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTags_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (to *EndpointTags_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointTags_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (m EndpointTags_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointTags_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(EndpointTagPair_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTags_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointTags_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags": m.CustomTags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointTags_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EndpointTags_SdkV2) GetCustomTags(ctx context.Context) ([]EndpointTagPair_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTagPair_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EndpointTags_SdkV2.
func (m *EndpointTags_SdkV2) SetCustomTags(ctx context.Context, v []EndpointTagPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type EnumValue_SdkV2 struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions types.String `tfsdk:"enum_options"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (to *EnumValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnumValue_SdkV2) {
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

func (to *EnumValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnumValue_SdkV2) {
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

func (m EnumValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EnumValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnumValue_SdkV2
// only implements ToObjectValue() and Type().
func (m EnumValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enum_options":         m.EnumOptions,
			"multi_values_options": m.MultiValuesOptions,
			"values":               m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnumValue_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EnumValue_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if m.MultiValuesOptions.IsNull() || m.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := m.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in EnumValue_SdkV2.
func (m *EnumValue_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	m.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in EnumValue_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EnumValue_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in EnumValue_SdkV2.
func (m *EnumValue_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type ExecuteStatementRequest_SdkV2 struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explicitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog"`
	// The fetch disposition provides two modes of fetching results: `INLINE`
	// and `EXTERNAL_LINKS`.
	//
	// Statements executed with `INLINE` disposition will return result data
	// inline, in `JSON_ARRAY` format, in a series of chunks. If a given
	// statement produces a result set with a size larger than 25 MiB, that
	// statement execution is aborted, and no result set will be available.
	//
	// **NOTE** Byte limits are computed based upon internal representations of
	// the result set data, and might not match the sizes visible in JSON
	// responses.
	//
	// Statements executed with `EXTERNAL_LINKS` disposition will return result
	// data as external links: URLs that point to cloud storage internal to the
	// workspace. Using `EXTERNAL_LINKS` disposition allows statements to
	// generate arbitrarily sized result sets for fetching up to 100 GiB. The
	// resulting links have two important properties:
	//
	// 1. They point to resources _external_ to the Databricks compute;
	// therefore any associated authentication information (typically a personal
	// access token, OAuth token, or similar) _must be removed_ when fetching
	// from these links.
	//
	// 2. These are URLs with a specific expiration, indicated in the response.
	// The behavior when attempting to use an expired link is cloud specific.
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
	// ``` SELECT * FROM my_table WHERE name = :my_name AND date = :my_date ```
	//
	// The parameters can be passed in the request body as follows:
	//
	// ` { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	// `
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

func (to *ExecuteStatementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteStatementRequest_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (to *ExecuteStatementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExecuteStatementRequest_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (m ExecuteStatementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExecuteStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(StatementParameterListItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExecuteStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_limit":      m.ByteLimit,
			"catalog":         m.Catalog,
			"disposition":     m.Disposition,
			"format":          m.Format,
			"on_wait_timeout": m.OnWaitTimeout,
			"parameters":      m.Parameters,
			"row_limit":       m.RowLimit,
			"schema":          m.Schema,
			"statement":       m.Statement,
			"wait_timeout":    m.WaitTimeout,
			"warehouse_id":    m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExecuteStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ExecuteStatementRequest_SdkV2) GetParameters(ctx context.Context) ([]StatementParameterListItem_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []StatementParameterListItem_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ExecuteStatementRequest_SdkV2.
func (m *ExecuteStatementRequest_SdkV2) SetParameters(ctx context.Context, v []StatementParameterListItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
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
	// A URL pointing to a chunk of result data, hosted by an external service,
	// with a short expiration time (<= 15 minutes). As this URL contains a
	// temporary credential, it should be considered sensitive and the client
	// should not expose this URL in a log.
	ExternalLink types.String `tfsdk:"external_link"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getstatementresultchunkn request.
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

func (to *ExternalLink_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalLink_SdkV2) {
}

func (to *ExternalLink_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalLink_SdkV2) {
}

func (m ExternalLink_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalLink_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"http_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLink_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalLink_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               m.ByteCount,
			"chunk_index":              m.ChunkIndex,
			"expiration":               m.Expiration,
			"external_link":            m.ExternalLink,
			"http_headers":             m.HttpHeaders,
			"next_chunk_index":         m.NextChunkIndex,
			"next_chunk_internal_link": m.NextChunkInternalLink,
			"row_count":                m.RowCount,
			"row_offset":               m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalLink_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ExternalLink_SdkV2) GetHttpHeaders(ctx context.Context) (map[string]types.String, bool) {
	if m.HttpHeaders.IsNull() || m.HttpHeaders.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.HttpHeaders.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpHeaders sets the value of the HttpHeaders field in ExternalLink_SdkV2.
func (m *ExternalLink_SdkV2) SetHttpHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_headers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.HttpHeaders = types.MapValueMust(t, vs)
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

func (to *ExternalQuerySource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalQuerySource_SdkV2) {
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

func (to *ExternalQuerySource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalQuerySource_SdkV2) {
	if !from.JobInfo.IsNull() && !from.JobInfo.IsUnknown() {
		if toJobInfo, ok := to.GetJobInfo(ctx); ok {
			if fromJobInfo, ok := from.GetJobInfo(ctx); ok {
				toJobInfo.SyncFieldsDuringRead(ctx, fromJobInfo)
				to.SetJobInfo(ctx, toJobInfo)
			}
		}
	}
}

func (m ExternalQuerySource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalQuerySource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_info": reflect.TypeOf(ExternalQuerySourceJobInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySource_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalQuerySource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id":            m.AlertId,
			"dashboard_id":        m.DashboardId,
			"genie_space_id":      m.GenieSpaceId,
			"job_info":            m.JobInfo,
			"legacy_dashboard_id": m.LegacyDashboardId,
			"notebook_id":         m.NotebookId,
			"sql_query_id":        m.SqlQueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalQuerySource_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ExternalQuerySource_SdkV2) GetJobInfo(ctx context.Context) (ExternalQuerySourceJobInfo_SdkV2, bool) {
	var e ExternalQuerySourceJobInfo_SdkV2
	if m.JobInfo.IsNull() || m.JobInfo.IsUnknown() {
		return e, false
	}
	var v []ExternalQuerySourceJobInfo_SdkV2
	d := m.JobInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobInfo sets the value of the JobInfo field in ExternalQuerySource_SdkV2.
func (m *ExternalQuerySource_SdkV2) SetJobInfo(ctx context.Context, v ExternalQuerySourceJobInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["job_info"]
	m.JobInfo = types.ListValueMust(t, vs)
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

func (to *ExternalQuerySourceJobInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalQuerySourceJobInfo_SdkV2) {
}

func (to *ExternalQuerySourceJobInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalQuerySourceJobInfo_SdkV2) {
}

func (m ExternalQuerySourceJobInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalQuerySourceJobInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalQuerySourceJobInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalQuerySourceJobInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":          m.JobId,
			"job_run_id":      m.JobRunId,
			"job_task_run_id": m.JobTaskRunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalQuerySourceJobInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetAlertRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertRequest_SdkV2) {
}

func (to *GetAlertRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAlertRequest_SdkV2) {
}

func (m GetAlertRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAlertV2Request_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetAlertV2Request_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertV2Request_SdkV2) {
}

func (to *GetAlertV2Request_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAlertV2Request_SdkV2) {
}

func (m GetAlertV2Request_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAlertsLegacyRequest_SdkV2 struct {
	AlertId types.String `tfsdk:"-"`
}

func (to *GetAlertsLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAlertsLegacyRequest_SdkV2) {
}

func (to *GetAlertsLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAlertsLegacyRequest_SdkV2) {
}

func (m GetAlertsLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAlertsLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAlertsLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAlertsLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id": m.AlertId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAlertsLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id": types.StringType,
		},
	}
}

type GetConfigRequest_SdkV2 struct {
}

func (to *GetConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetConfigRequest_SdkV2) {
}

func (to *GetConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetConfigRequest_SdkV2) {
}

func (m GetConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetDashboardRequest_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *GetDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDashboardRequest_SdkV2) {
}

func (to *GetDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDashboardRequest_SdkV2) {
}

func (m GetDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetDbsqlPermissionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDbsqlPermissionRequest_SdkV2) {
}

func (to *GetDbsqlPermissionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDbsqlPermissionRequest_SdkV2) {
}

func (m GetDbsqlPermissionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDbsqlPermissionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDbsqlPermissionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDbsqlPermissionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDbsqlPermissionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type GetQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *GetQueriesLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetQueriesLegacyRequest_SdkV2) {
}

func (to *GetQueriesLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetQueriesLegacyRequest_SdkV2) {
}

func (m GetQueriesLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": m.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type GetQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetQueryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetQueryRequest_SdkV2) {
}

func (to *GetQueryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetQueryRequest_SdkV2) {
}

func (m GetQueryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetResponse_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *GetResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetResponse_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m GetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetResponse_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in GetResponse_SdkV2.
func (m *GetResponse_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type GetStatementRequest_SdkV2 struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (to *GetStatementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementRequest_SdkV2) {
}

func (to *GetStatementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatementRequest_SdkV2) {
}

func (m GetStatementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetStatementResultChunkNRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementResultChunkNRequest_SdkV2) {
}

func (to *GetStatementResultChunkNRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatementResultChunkNRequest_SdkV2) {
}

func (m GetStatementResultChunkNRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetStatementResultChunkNRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementResultChunkNRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatementResultChunkNRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_index":  m.ChunkIndex,
			"statement_id": m.StatementId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatementResultChunkNRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetWarehousePermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionLevelsRequest_SdkV2) {
}

func (to *GetWarehousePermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionLevelsRequest_SdkV2) {
}

func (m GetWarehousePermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWarehousePermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWarehousePermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWarehousePermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetWarehousePermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetWarehousePermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetWarehousePermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWarehousePermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WarehousePermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWarehousePermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWarehousePermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetWarehousePermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]WarehousePermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWarehousePermissionLevelsResponse_SdkV2.
func (m *GetWarehousePermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []WarehousePermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWarehousePermissionsRequest_SdkV2 struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetWarehousePermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehousePermissionsRequest_SdkV2) {
}

func (to *GetWarehousePermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWarehousePermissionsRequest_SdkV2) {
}

func (m GetWarehousePermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWarehousePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehousePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWarehousePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWarehousePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetWarehouseRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehouseRequest_SdkV2) {
}

func (to *GetWarehouseRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWarehouseRequest_SdkV2) {
}

func (m GetWarehouseRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWarehouseRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWarehouseRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWarehouseRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWarehouseRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
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
	// Configurations whether the endpoint should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy"`
	// state of the endpoint
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

func (to *GetWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWarehouseResponse_SdkV2) {
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

func (to *GetWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWarehouseResponse_SdkV2) {
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

func (m GetWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_stop_mins"] = attrs["auto_stop_mins"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_size"] = attrs["cluster_size"].SetOptional()
	attrs["creator_name"] = attrs["creator_name"].SetOptional()
	attrs["enable_photon"] = attrs["enable_photon"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
	attrs["health"] = attrs["health"].SetComputed()
	attrs["health"] = attrs["health"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["id"] = attrs["id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
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
func (m GetWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m GetWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_stop_mins":            m.AutoStopMins,
			"channel":                   m.Channel,
			"cluster_size":              m.ClusterSize,
			"creator_name":              m.CreatorName,
			"enable_photon":             m.EnablePhoton,
			"enable_serverless_compute": m.EnableServerlessCompute,
			"health":                    m.Health,
			"id":                        m.Id,
			"instance_profile_arn":      m.InstanceProfileArn,
			"jdbc_url":                  m.JdbcUrl,
			"max_num_clusters":          m.MaxNumClusters,
			"min_num_clusters":          m.MinNumClusters,
			"name":                      m.Name,
			"num_active_sessions":       m.NumActiveSessions,
			"num_clusters":              m.NumClusters,
			"odbc_params":               m.OdbcParams,
			"spot_instance_policy":      m.SpotInstancePolicy,
			"state":                     m.State,
			"tags":                      m.Tags,
			"warehouse_type":            m.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetWarehouseResponse_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWarehouseResponse_SdkV2.
func (m *GetWarehouseResponse_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in GetWarehouseResponse_SdkV2 as
// a EndpointHealth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWarehouseResponse_SdkV2) GetHealth(ctx context.Context) (EndpointHealth_SdkV2, bool) {
	var e EndpointHealth_SdkV2
	if m.Health.IsNull() || m.Health.IsUnknown() {
		return e, false
	}
	var v []EndpointHealth_SdkV2
	d := m.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in GetWarehouseResponse_SdkV2.
func (m *GetWarehouseResponse_SdkV2) SetHealth(ctx context.Context, v EndpointHealth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	m.Health = types.ListValueMust(t, vs)
}

// GetOdbcParams returns the value of the OdbcParams field in GetWarehouseResponse_SdkV2 as
// a OdbcParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWarehouseResponse_SdkV2) GetOdbcParams(ctx context.Context) (OdbcParams_SdkV2, bool) {
	var e OdbcParams_SdkV2
	if m.OdbcParams.IsNull() || m.OdbcParams.IsUnknown() {
		return e, false
	}
	var v []OdbcParams_SdkV2
	d := m.OdbcParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOdbcParams sets the value of the OdbcParams field in GetWarehouseResponse_SdkV2.
func (m *GetWarehouseResponse_SdkV2) SetOdbcParams(ctx context.Context, v OdbcParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["odbc_params"]
	m.OdbcParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in GetWarehouseResponse_SdkV2 as
// a EndpointTags_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWarehouseResponse_SdkV2) GetTags(ctx context.Context) (EndpointTags_SdkV2, bool) {
	var e EndpointTags_SdkV2
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return e, false
	}
	var v []EndpointTags_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTags sets the value of the Tags field in GetWarehouseResponse_SdkV2.
func (m *GetWarehouseResponse_SdkV2) SetTags(ctx context.Context, v EndpointTags_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	m.Tags = types.ListValueMust(t, vs)
}

type GetWorkspaceWarehouseConfigRequest_SdkV2 struct {
}

func (to *GetWorkspaceWarehouseConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceWarehouseConfigRequest_SdkV2) {
}

func (to *GetWorkspaceWarehouseConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceWarehouseConfigRequest_SdkV2) {
}

func (m GetWorkspaceWarehouseConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceWarehouseConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceWarehouseConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceWarehouseConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceWarehouseConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Enable Serverless compute for SQL warehouses
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
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
	// AWS Only: The instance profile used to pass an IAM role to the SQL
	// warehouses. This configuration is also applied to the workspace's
	// serverless compute for notebooks and jobs.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters"`
}

func (to *GetWorkspaceWarehouseConfigResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceWarehouseConfigResponse_SdkV2) {
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

func (to *GetWorkspaceWarehouseConfigResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceWarehouseConfigResponse_SdkV2) {
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

func (m GetWorkspaceWarehouseConfigResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["config_param"] = attrs["config_param"].SetOptional()
	attrs["config_param"] = attrs["config_param"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["data_access_config"] = attrs["data_access_config"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
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
func (m GetWorkspaceWarehouseConfigResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m GetWorkspaceWarehouseConfigResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel":                      m.Channel,
			"config_param":                 m.ConfigParam,
			"data_access_config":           m.DataAccessConfig,
			"enable_serverless_compute":    m.EnableServerlessCompute,
			"enabled_warehouse_types":      m.EnabledWarehouseTypes,
			"global_param":                 m.GlobalParam,
			"google_service_account":       m.GoogleServiceAccount,
			"instance_profile_arn":         m.InstanceProfileArn,
			"security_policy":              m.SecurityPolicy,
			"sql_configuration_parameters": m.SqlConfigurationParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceWarehouseConfigResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
			"enable_serverless_compute": types.BoolType,
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
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.ConfigParam.IsNull() || m.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	m.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if m.DataAccessConfig.IsNull() || m.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := m.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a slice of WarehouseTypePair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair_SdkV2, bool) {
	if m.EnabledWarehouseTypes.IsNull() || m.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair_SdkV2
	d := m.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.GlobalParam.IsNull() || m.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	m.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.SqlConfigurationParameters.IsNull() || m.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in GetWorkspaceWarehouseConfigResponse_SdkV2.
func (m *GetWorkspaceWarehouseConfigResponse_SdkV2) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	m.SqlConfigurationParameters = types.ListValueMust(t, vs)
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

func (to *LegacyAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyAlert_SdkV2) {
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

func (to *LegacyAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LegacyAlert_SdkV2) {
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

func (m LegacyAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LegacyAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AlertOptions_SdkV2{}),
		"query":   reflect.TypeOf(AlertQuery_SdkV2{}),
		"user":    reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m LegacyAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":        m.CreatedAt,
			"id":                m.Id,
			"last_triggered_at": m.LastTriggeredAt,
			"name":              m.Name,
			"options":           m.Options,
			"parent":            m.Parent,
			"query":             m.Query,
			"rearm":             m.Rearm,
			"state":             m.State,
			"updated_at":        m.UpdatedAt,
			"user":              m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LegacyAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *LegacyAlert_SdkV2) GetOptions(ctx context.Context) (AlertOptions_SdkV2, bool) {
	var e AlertOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AlertOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyAlert_SdkV2.
func (m *LegacyAlert_SdkV2) SetOptions(ctx context.Context, v AlertOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// GetQuery returns the value of the Query field in LegacyAlert_SdkV2 as
// a AlertQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyAlert_SdkV2) GetQuery(ctx context.Context) (AlertQuery_SdkV2, bool) {
	var e AlertQuery_SdkV2
	if m.Query.IsNull() || m.Query.IsUnknown() {
		return e, false
	}
	var v []AlertQuery_SdkV2
	d := m.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyAlert_SdkV2.
func (m *LegacyAlert_SdkV2) SetQuery(ctx context.Context, v AlertQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	m.Query = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyAlert_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyAlert_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := m.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyAlert_SdkV2.
func (m *LegacyAlert_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
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

func (to *LegacyQuery_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyQuery_SdkV2) {
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

func (to *LegacyQuery_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LegacyQuery_SdkV2) {
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

func (m LegacyQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LegacyQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m LegacyQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_edit":             m.CanEdit,
			"created_at":           m.CreatedAt,
			"data_source_id":       m.DataSourceId,
			"description":          m.Description,
			"id":                   m.Id,
			"is_archived":          m.IsArchived,
			"is_draft":             m.IsDraft,
			"is_favorite":          m.IsFavorite,
			"is_safe":              m.IsSafe,
			"last_modified_by":     m.LastModifiedBy,
			"last_modified_by_id":  m.LastModifiedById,
			"latest_query_data_id": m.LatestQueryDataId,
			"name":                 m.Name,
			"options":              m.Options,
			"parent":               m.Parent,
			"permission_tier":      m.PermissionTier,
			"query":                m.Query,
			"query_hash":           m.QueryHash,
			"run_as_role":          m.RunAsRole,
			"tags":                 m.Tags,
			"updated_at":           m.UpdatedAt,
			"user":                 m.User,
			"user_id":              m.UserId,
			"visualizations":       m.Visualizations,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LegacyQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *LegacyQuery_SdkV2) GetLastModifiedBy(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if m.LastModifiedBy.IsNull() || m.LastModifiedBy.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := m.LastModifiedBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLastModifiedBy sets the value of the LastModifiedBy field in LegacyQuery_SdkV2.
func (m *LegacyQuery_SdkV2) SetLastModifiedBy(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["last_modified_by"]
	m.LastModifiedBy = types.ListValueMust(t, vs)
}

// GetOptions returns the value of the Options field in LegacyQuery_SdkV2 as
// a QueryOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyQuery_SdkV2) GetOptions(ctx context.Context) (QueryOptions_SdkV2, bool) {
	var e QueryOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []QueryOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in LegacyQuery_SdkV2.
func (m *LegacyQuery_SdkV2) SetOptions(ctx context.Context, v QueryOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LegacyQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LegacyQuery_SdkV2.
func (m *LegacyQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// GetUser returns the value of the User field in LegacyQuery_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyQuery_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
	var e User_SdkV2
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v []User_SdkV2
	d := m.User.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUser sets the value of the User field in LegacyQuery_SdkV2.
func (m *LegacyQuery_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

// GetVisualizations returns the value of the Visualizations field in LegacyQuery_SdkV2 as
// a slice of LegacyVisualization_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LegacyQuery_SdkV2) GetVisualizations(ctx context.Context) ([]LegacyVisualization_SdkV2, bool) {
	if m.Visualizations.IsNull() || m.Visualizations.IsUnknown() {
		return nil, false
	}
	var v []LegacyVisualization_SdkV2
	d := m.Visualizations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVisualizations sets the value of the Visualizations field in LegacyQuery_SdkV2.
func (m *LegacyQuery_SdkV2) SetVisualizations(ctx context.Context, v []LegacyVisualization_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["visualizations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Visualizations = types.ListValueMust(t, vs)
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

func (to *LegacyVisualization_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LegacyVisualization_SdkV2) {
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

func (to *LegacyVisualization_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LegacyVisualization_SdkV2) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (m LegacyVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LegacyVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(LegacyQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LegacyVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (m LegacyVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":  m.CreatedAt,
			"description": m.Description,
			"id":          m.Id,
			"name":        m.Name,
			"options":     m.Options,
			"query":       m.Query,
			"type":        m.Type_,
			"updated_at":  m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LegacyVisualization_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *LegacyVisualization_SdkV2) GetQuery(ctx context.Context) (LegacyQuery_SdkV2, bool) {
	var e LegacyQuery_SdkV2
	if m.Query.IsNull() || m.Query.IsUnknown() {
		return e, false
	}
	var v []LegacyQuery_SdkV2
	d := m.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in LegacyVisualization_SdkV2.
func (m *LegacyVisualization_SdkV2) SetQuery(ctx context.Context, v LegacyQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	m.Query = types.ListValueMust(t, vs)
}

type ListAlertsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAlertsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsRequest_SdkV2) {
}

func (to *ListAlertsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAlertsRequest_SdkV2) {
}

func (m ListAlertsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAlertsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAlertsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAlertsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListAlertsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListAlertsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAlertsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListAlertsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAlertsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListAlertsResponseAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAlertsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"results":         m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAlertsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListAlertsResponse_SdkV2) GetResults(ctx context.Context) ([]ListAlertsResponseAlert_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []ListAlertsResponseAlert_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListAlertsResponse_SdkV2.
func (m *ListAlertsResponse_SdkV2) SetResults(ctx context.Context, v []ListAlertsResponseAlert_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
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

func (to *ListAlertsResponseAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsResponseAlert_SdkV2) {
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

func (to *ListAlertsResponseAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAlertsResponseAlert_SdkV2) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (m ListAlertsResponseAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAlertsResponseAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsResponseAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAlertsResponseAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            m.Condition,
			"create_time":          m.CreateTime,
			"custom_body":          m.CustomBody,
			"custom_subject":       m.CustomSubject,
			"display_name":         m.DisplayName,
			"id":                   m.Id,
			"lifecycle_state":      m.LifecycleState,
			"notify_on_ok":         m.NotifyOnOk,
			"owner_user_name":      m.OwnerUserName,
			"query_id":             m.QueryId,
			"seconds_to_retrigger": m.SecondsToRetrigger,
			"state":                m.State,
			"trigger_time":         m.TriggerTime,
			"update_time":          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAlertsResponseAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListAlertsResponseAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if m.Condition.IsNull() || m.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := m.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in ListAlertsResponseAlert_SdkV2.
func (m *ListAlertsResponseAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	m.Condition = types.ListValueMust(t, vs)
}

type ListAlertsV2Request_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAlertsV2Request_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsV2Request_SdkV2) {
}

func (to *ListAlertsV2Request_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAlertsV2Request_SdkV2) {
}

func (m ListAlertsV2Request_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAlertsV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAlertsV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAlertsV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAlertsV2Response_SdkV2 struct {
	Alerts types.List `tfsdk:"alerts"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListAlertsV2Response_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAlertsV2Response_SdkV2) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
}

func (to *ListAlertsV2Response_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAlertsV2Response_SdkV2) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
}

func (m ListAlertsV2Response_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alerts"] = attrs["alerts"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAlertsV2Response.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAlertsV2Response_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAlertsV2Response_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAlertsV2Response_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":          m.Alerts,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAlertsV2Response_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alerts": basetypes.ListType{
				ElemType: AlertV2_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAlerts returns the value of the Alerts field in ListAlertsV2Response_SdkV2 as
// a slice of AlertV2_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAlertsV2Response_SdkV2) GetAlerts(ctx context.Context) ([]AlertV2_SdkV2, bool) {
	if m.Alerts.IsNull() || m.Alerts.IsUnknown() {
		return nil, false
	}
	var v []AlertV2_SdkV2
	d := m.Alerts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlerts sets the value of the Alerts field in ListAlertsV2Response_SdkV2.
func (m *ListAlertsV2Response_SdkV2) SetAlerts(ctx context.Context, v []AlertV2_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Alerts = types.ListValueMust(t, vs)
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

func (to *ListDashboardsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDashboardsRequest_SdkV2) {
}

func (to *ListDashboardsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDashboardsRequest_SdkV2) {
}

func (m ListDashboardsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDashboardsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDashboardsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDashboardsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"order":     m.Order,
			"page":      m.Page,
			"page_size": m.PageSize,
			"q":         m.Q,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDashboardsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListQueriesLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesLegacyRequest_SdkV2) {
}

func (to *ListQueriesLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueriesLegacyRequest_SdkV2) {
}

func (m ListQueriesLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"order":     m.Order,
			"page":      m.Page,
			"page_size": m.PageSize,
			"q":         m.Q,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListQueriesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesRequest_SdkV2) {
}

func (to *ListQueriesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueriesRequest_SdkV2) {
}

func (m ListQueriesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQueriesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueriesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueriesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListQueriesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueriesResponse_SdkV2) {
	if !from.Res.IsNull() && !from.Res.IsUnknown() && to.Res.IsNull() && len(from.Res.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Res, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Res = from.Res
	}
}

func (to *ListQueriesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueriesResponse_SdkV2) {
	if !from.Res.IsNull() && !from.Res.IsUnknown() && to.Res.IsNull() && len(from.Res.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Res, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Res = from.Res
	}
}

func (m ListQueriesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQueriesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"res": reflect.TypeOf(QueryInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueriesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueriesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_next_page":   m.HasNextPage,
			"next_page_token": m.NextPageToken,
			"res":             m.Res,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueriesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListQueriesResponse_SdkV2) GetRes(ctx context.Context) ([]QueryInfo_SdkV2, bool) {
	if m.Res.IsNull() || m.Res.IsUnknown() {
		return nil, false
	}
	var v []QueryInfo_SdkV2
	d := m.Res.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRes sets the value of the Res field in ListQueriesResponse_SdkV2.
func (m *ListQueriesResponse_SdkV2) SetRes(ctx context.Context, v []QueryInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["res"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Res = types.ListValueMust(t, vs)
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

func (to *ListQueryHistoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryHistoryRequest_SdkV2) {
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

func (to *ListQueryHistoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueryHistoryRequest_SdkV2) {
	if !from.FilterBy.IsNull() && !from.FilterBy.IsUnknown() {
		if toFilterBy, ok := to.GetFilterBy(ctx); ok {
			if fromFilterBy, ok := from.GetFilterBy(ctx); ok {
				toFilterBy.SyncFieldsDuringRead(ctx, fromFilterBy)
				to.SetFilterBy(ctx, toFilterBy)
			}
		}
	}
}

func (m ListQueryHistoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["filter_by"] = attrs["filter_by"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m ListQueryHistoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"filter_by": reflect.TypeOf(QueryFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryHistoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueryHistoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter_by":       m.FilterBy,
			"include_metrics": m.IncludeMetrics,
			"max_results":     m.MaxResults,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueryHistoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListQueryHistoryRequest_SdkV2) GetFilterBy(ctx context.Context) (QueryFilter_SdkV2, bool) {
	var e QueryFilter_SdkV2
	if m.FilterBy.IsNull() || m.FilterBy.IsUnknown() {
		return e, false
	}
	var v []QueryFilter_SdkV2
	d := m.FilterBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilterBy sets the value of the FilterBy field in ListQueryHistoryRequest_SdkV2.
func (m *ListQueryHistoryRequest_SdkV2) SetFilterBy(ctx context.Context, v QueryFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filter_by"]
	m.FilterBy = types.ListValueMust(t, vs)
}

type ListQueryObjectsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListQueryObjectsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryObjectsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListQueryObjectsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueryObjectsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListQueryObjectsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQueryObjectsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListQueryObjectsResponseQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueryObjectsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"results":         m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueryObjectsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListQueryObjectsResponse_SdkV2) GetResults(ctx context.Context) ([]ListQueryObjectsResponseQuery_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []ListQueryObjectsResponseQuery_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListQueryObjectsResponse_SdkV2.
func (m *ListQueryObjectsResponse_SdkV2) SetResults(ctx context.Context, v []ListQueryObjectsResponseQuery_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
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

func (to *ListQueryObjectsResponseQuery_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQueryObjectsResponseQuery_SdkV2) {
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

func (to *ListQueryObjectsResponseQuery_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListQueryObjectsResponseQuery_SdkV2) {
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

func (m ListQueryObjectsResponseQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQueryObjectsResponseQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQueryObjectsResponseQuery_SdkV2
// only implements ToObjectValue() and Type().
func (m ListQueryObjectsResponseQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit":        m.ApplyAutoLimit,
			"catalog":                 m.Catalog,
			"create_time":             m.CreateTime,
			"description":             m.Description,
			"display_name":            m.DisplayName,
			"id":                      m.Id,
			"last_modifier_user_name": m.LastModifierUserName,
			"lifecycle_state":         m.LifecycleState,
			"owner_user_name":         m.OwnerUserName,
			"parameters":              m.Parameters,
			"query_text":              m.QueryText,
			"run_as_mode":             m.RunAsMode,
			"schema":                  m.Schema,
			"tags":                    m.Tags,
			"update_time":             m.UpdateTime,
			"warehouse_id":            m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQueryObjectsResponseQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListQueryObjectsResponseQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ListQueryObjectsResponseQuery_SdkV2.
func (m *ListQueryObjectsResponseQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ListQueryObjectsResponseQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListQueryObjectsResponseQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ListQueryObjectsResponseQuery_SdkV2.
func (m *ListQueryObjectsResponseQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ListRequest_SdkV2 struct {
}

func (to *ListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRequest_SdkV2) {
}

func (to *ListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRequest_SdkV2) {
}

func (m ListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Dashboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"count":     m.Count,
			"page":      m.Page,
			"page_size": m.PageSize,
			"results":   m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListResponse_SdkV2) GetResults(ctx context.Context) ([]Dashboard_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []Dashboard_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListResponse_SdkV2.
func (m *ListResponse_SdkV2) SetResults(ctx context.Context, v []Dashboard_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
}

type ListVisualizationsForQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListVisualizationsForQueryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVisualizationsForQueryRequest_SdkV2) {
}

func (to *ListVisualizationsForQueryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVisualizationsForQueryRequest_SdkV2) {
}

func (m ListVisualizationsForQueryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListVisualizationsForQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListVisualizationsForQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         m.Id,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVisualizationsForQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListVisualizationsForQueryResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVisualizationsForQueryResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListVisualizationsForQueryResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVisualizationsForQueryResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListVisualizationsForQueryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListVisualizationsForQueryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(Visualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVisualizationsForQueryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListVisualizationsForQueryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"results":         m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVisualizationsForQueryResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListVisualizationsForQueryResponse_SdkV2) GetResults(ctx context.Context) ([]Visualization_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []Visualization_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListVisualizationsForQueryResponse_SdkV2.
func (m *ListVisualizationsForQueryResponse_SdkV2) SetResults(ctx context.Context, v []Visualization_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
}

type ListWarehousesRequest_SdkV2 struct {
	// The max number of warehouses to return.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListWarehouses` call. Provide
	// this to retrieve the subsequent page; otherwise the first will be
	// retrieved.
	//
	// When paginating, all other parameters provided to `ListWarehouses` must
	// match the call that provided the page token.
	PageToken types.String `tfsdk:"-"`
	// Service Principal which will be used to fetch the list of endpoints. If
	// not specified, SQL Gateway will use the user from the session header.
	RunAsUserId types.Int64 `tfsdk:"-"`
}

func (to *ListWarehousesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWarehousesRequest_SdkV2) {
}

func (to *ListWarehousesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWarehousesRequest_SdkV2) {
}

func (m ListWarehousesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_as_user_id"] = attrs["run_as_user_id"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWarehousesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWarehousesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWarehousesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":      m.PageSize,
			"page_token":     m.PageToken,
			"run_as_user_id": m.RunAsUserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWarehousesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":      types.Int64Type,
			"page_token":     types.StringType,
			"run_as_user_id": types.Int64Type,
		},
	}
}

type ListWarehousesResponse_SdkV2 struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// A list of warehouses and their configurations.
	Warehouses types.List `tfsdk:"warehouses"`
}

func (to *ListWarehousesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWarehousesResponse_SdkV2) {
	if !from.Warehouses.IsNull() && !from.Warehouses.IsUnknown() && to.Warehouses.IsNull() && len(from.Warehouses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Warehouses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Warehouses = from.Warehouses
	}
}

func (to *ListWarehousesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWarehousesResponse_SdkV2) {
	if !from.Warehouses.IsNull() && !from.Warehouses.IsUnknown() && to.Warehouses.IsNull() && len(from.Warehouses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Warehouses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Warehouses = from.Warehouses
	}
}

func (m ListWarehousesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
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
func (m ListWarehousesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"warehouses": reflect.TypeOf(EndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWarehousesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWarehousesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"warehouses":      m.Warehouses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWarehousesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"warehouses": basetypes.ListType{
				ElemType: EndpointInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWarehouses returns the value of the Warehouses field in ListWarehousesResponse_SdkV2 as
// a slice of EndpointInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWarehousesResponse_SdkV2) GetWarehouses(ctx context.Context) ([]EndpointInfo_SdkV2, bool) {
	if m.Warehouses.IsNull() || m.Warehouses.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo_SdkV2
	d := m.Warehouses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarehouses sets the value of the Warehouses field in ListWarehousesResponse_SdkV2.
func (m *ListWarehousesResponse_SdkV2) SetWarehouses(ctx context.Context, v []EndpointInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Warehouses = types.ListValueMust(t, vs)
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

func (to *MultiValuesOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MultiValuesOptions_SdkV2) {
}

func (to *MultiValuesOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MultiValuesOptions_SdkV2) {
}

func (m MultiValuesOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m MultiValuesOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MultiValuesOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m MultiValuesOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"prefix":    m.Prefix,
			"separator": m.Separator,
			"suffix":    m.Suffix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MultiValuesOptions_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *NumericValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NumericValue_SdkV2) {
}

func (to *NumericValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NumericValue_SdkV2) {
}

func (m NumericValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NumericValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NumericValue_SdkV2
// only implements ToObjectValue() and Type().
func (m NumericValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NumericValue_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *OdbcParams_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OdbcParams_SdkV2) {
}

func (to *OdbcParams_SdkV2) SyncFieldsDuringRead(ctx context.Context, from OdbcParams_SdkV2) {
}

func (m OdbcParams_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m OdbcParams_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OdbcParams_SdkV2
// only implements ToObjectValue() and Type().
func (m OdbcParams_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hostname": m.Hostname,
			"path":     m.Path,
			"port":     m.Port,
			"protocol": m.Protocol,
		})
}

// Type implements basetypes.ObjectValuable.
func (m OdbcParams_SdkV2) Type(ctx context.Context) attr.Type {
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
	EnumOptions types.String `tfsdk:"enum_options"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions types.List `tfsdk:"multi_values_options"`
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

func (to *Parameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Parameter_SdkV2) {
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

func (to *Parameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Parameter_SdkV2) {
	if !from.MultiValuesOptions.IsNull() && !from.MultiValuesOptions.IsUnknown() {
		if toMultiValuesOptions, ok := to.GetMultiValuesOptions(ctx); ok {
			if fromMultiValuesOptions, ok := from.GetMultiValuesOptions(ctx); ok {
				toMultiValuesOptions.SyncFieldsDuringRead(ctx, fromMultiValuesOptions)
				to.SetMultiValuesOptions(ctx, toMultiValuesOptions)
			}
		}
	}
}

func (m Parameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enum_options"] = attrs["enum_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].SetOptional()
	attrs["multi_values_options"] = attrs["multi_values_options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m Parameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Parameter_SdkV2
// only implements ToObjectValue() and Type().
func (m Parameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enum_options":         m.EnumOptions,
			"multi_values_options": m.MultiValuesOptions,
			"name":                 m.Name,
			"query_id":             m.QueryId,
			"title":                m.Title,
			"type":                 m.Type_,
			"value":                m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Parameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enum_options": types.StringType,
			"multi_values_options": basetypes.ListType{
				ElemType: MultiValuesOptions_SdkV2{}.Type(ctx),
			},
			"name":     types.StringType,
			"query_id": types.StringType,
			"title":    types.StringType,
			"type":     types.StringType,
			"value":    types.ObjectType{},
		},
	}
}

// GetMultiValuesOptions returns the value of the MultiValuesOptions field in Parameter_SdkV2 as
// a MultiValuesOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Parameter_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if m.MultiValuesOptions.IsNull() || m.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := m.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in Parameter_SdkV2.
func (m *Parameter_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	m.MultiValuesOptions = types.ListValueMust(t, vs)
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

func (to *Query_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Query_SdkV2) {
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

func (to *Query_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Query_SdkV2) {
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

func (m Query_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Query_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Query_SdkV2
// only implements ToObjectValue() and Type().
func (m Query_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit":        m.ApplyAutoLimit,
			"catalog":                 m.Catalog,
			"create_time":             m.CreateTime,
			"description":             m.Description,
			"display_name":            m.DisplayName,
			"id":                      m.Id,
			"last_modifier_user_name": m.LastModifierUserName,
			"lifecycle_state":         m.LifecycleState,
			"owner_user_name":         m.OwnerUserName,
			"parameters":              m.Parameters,
			"parent_path":             m.ParentPath,
			"query_text":              m.QueryText,
			"run_as_mode":             m.RunAsMode,
			"schema":                  m.Schema,
			"tags":                    m.Tags,
			"update_time":             m.UpdateTime,
			"warehouse_id":            m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Query_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Query_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in Query_SdkV2.
func (m *Query_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Query_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Query_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Query_SdkV2.
func (m *Query_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type QueryBackedValue_SdkV2 struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.List `tfsdk:"multi_values_options"`
	// UUID of the query that provides the parameter values.
	QueryId types.String `tfsdk:"query_id"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values"`
}

func (to *QueryBackedValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryBackedValue_SdkV2) {
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

func (to *QueryBackedValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryBackedValue_SdkV2) {
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

func (m QueryBackedValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryBackedValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"multi_values_options": reflect.TypeOf(MultiValuesOptions_SdkV2{}),
		"values":               reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryBackedValue_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryBackedValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"multi_values_options": m.MultiValuesOptions,
			"query_id":             m.QueryId,
			"values":               m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryBackedValue_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryBackedValue_SdkV2) GetMultiValuesOptions(ctx context.Context) (MultiValuesOptions_SdkV2, bool) {
	var e MultiValuesOptions_SdkV2
	if m.MultiValuesOptions.IsNull() || m.MultiValuesOptions.IsUnknown() {
		return e, false
	}
	var v []MultiValuesOptions_SdkV2
	d := m.MultiValuesOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMultiValuesOptions sets the value of the MultiValuesOptions field in QueryBackedValue_SdkV2.
func (m *QueryBackedValue_SdkV2) SetMultiValuesOptions(ctx context.Context, v MultiValuesOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["multi_values_options"]
	m.MultiValuesOptions = types.ListValueMust(t, vs)
}

// GetValues returns the value of the Values field in QueryBackedValue_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryBackedValue_SdkV2) GetValues(ctx context.Context) ([]types.String, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in QueryBackedValue_SdkV2.
func (m *QueryBackedValue_SdkV2) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
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

func (to *QueryEditContent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEditContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *QueryEditContent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryEditContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m QueryEditContent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryEditContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEditContent_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryEditContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_source_id": m.DataSourceId,
			"description":    m.Description,
			"name":           m.Name,
			"options":        m.Options,
			"query":          m.Query,
			"query_id":       m.QueryId,
			"run_as_role":    m.RunAsRole,
			"tags":           m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryEditContent_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryEditContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in QueryEditContent_SdkV2.
func (m *QueryEditContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
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

func (to *QueryFilter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryFilter_SdkV2) {
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

func (to *QueryFilter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryFilter_SdkV2) {
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

func (m QueryFilter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m QueryFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_start_time_range": m.QueryStartTimeRange,
			"statement_ids":          m.StatementIds,
			"statuses":               m.Statuses,
			"user_ids":               m.UserIds,
			"warehouse_ids":          m.WarehouseIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryFilter_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryFilter_SdkV2) GetQueryStartTimeRange(ctx context.Context) (TimeRange_SdkV2, bool) {
	var e TimeRange_SdkV2
	if m.QueryStartTimeRange.IsNull() || m.QueryStartTimeRange.IsUnknown() {
		return e, false
	}
	var v []TimeRange_SdkV2
	d := m.QueryStartTimeRange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryStartTimeRange sets the value of the QueryStartTimeRange field in QueryFilter_SdkV2.
func (m *QueryFilter_SdkV2) SetQueryStartTimeRange(ctx context.Context, v TimeRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_start_time_range"]
	m.QueryStartTimeRange = types.ListValueMust(t, vs)
}

// GetStatementIds returns the value of the StatementIds field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryFilter_SdkV2) GetStatementIds(ctx context.Context) ([]types.String, bool) {
	if m.StatementIds.IsNull() || m.StatementIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.StatementIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatementIds sets the value of the StatementIds field in QueryFilter_SdkV2.
func (m *QueryFilter_SdkV2) SetStatementIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["statement_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.StatementIds = types.ListValueMust(t, vs)
}

// GetStatuses returns the value of the Statuses field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryFilter_SdkV2) GetStatuses(ctx context.Context) ([]types.String, bool) {
	if m.Statuses.IsNull() || m.Statuses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in QueryFilter_SdkV2.
func (m *QueryFilter_SdkV2) SetStatuses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Statuses = types.ListValueMust(t, vs)
}

// GetUserIds returns the value of the UserIds field in QueryFilter_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryFilter_SdkV2) GetUserIds(ctx context.Context) ([]types.Int64, bool) {
	if m.UserIds.IsNull() || m.UserIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.UserIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserIds sets the value of the UserIds field in QueryFilter_SdkV2.
func (m *QueryFilter_SdkV2) SetUserIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserIds = types.ListValueMust(t, vs)
}

// GetWarehouseIds returns the value of the WarehouseIds field in QueryFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryFilter_SdkV2) GetWarehouseIds(ctx context.Context) ([]types.String, bool) {
	if m.WarehouseIds.IsNull() || m.WarehouseIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.WarehouseIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarehouseIds sets the value of the WarehouseIds field in QueryFilter_SdkV2.
func (m *QueryFilter_SdkV2) SetWarehouseIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["warehouse_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WarehouseIds = types.ListValueMust(t, vs)
}

type QueryInfo_SdkV2 struct {
	// The ID of the cached query if this result retrieved from cache
	CacheQueryId types.String `tfsdk:"cache_query_id"`
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

func (to *QueryInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryInfo_SdkV2) {
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

func (to *QueryInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryInfo_SdkV2) {
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

func (m QueryInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cache_query_id"] = attrs["cache_query_id"].SetOptional()
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
func (m QueryInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"channel_used": reflect.TypeOf(ChannelInfo_SdkV2{}),
		"metrics":      reflect.TypeOf(QueryMetrics_SdkV2{}),
		"query_source": reflect.TypeOf(ExternalQuerySource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cache_query_id":        m.CacheQueryId,
			"channel_used":          m.ChannelUsed,
			"client_application":    m.ClientApplication,
			"duration":              m.Duration,
			"endpoint_id":           m.EndpointId,
			"error_message":         m.ErrorMessage,
			"executed_as_user_id":   m.ExecutedAsUserId,
			"executed_as_user_name": m.ExecutedAsUserName,
			"execution_end_time_ms": m.ExecutionEndTimeMs,
			"is_final":              m.IsFinal,
			"lookup_key":            m.LookupKey,
			"metrics":               m.Metrics,
			"plans_state":           m.PlansState,
			"query_end_time_ms":     m.QueryEndTimeMs,
			"query_id":              m.QueryId,
			"query_source":          m.QuerySource,
			"query_start_time_ms":   m.QueryStartTimeMs,
			"query_text":            m.QueryText,
			"rows_produced":         m.RowsProduced,
			"spark_ui_url":          m.SparkUiUrl,
			"statement_type":        m.StatementType,
			"status":                m.Status,
			"user_id":               m.UserId,
			"user_name":             m.UserName,
			"warehouse_id":          m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cache_query_id": types.StringType,
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
func (m *QueryInfo_SdkV2) GetChannelUsed(ctx context.Context) (ChannelInfo_SdkV2, bool) {
	var e ChannelInfo_SdkV2
	if m.ChannelUsed.IsNull() || m.ChannelUsed.IsUnknown() {
		return e, false
	}
	var v []ChannelInfo_SdkV2
	d := m.ChannelUsed.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannelUsed sets the value of the ChannelUsed field in QueryInfo_SdkV2.
func (m *QueryInfo_SdkV2) SetChannelUsed(ctx context.Context, v ChannelInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel_used"]
	m.ChannelUsed = types.ListValueMust(t, vs)
}

// GetMetrics returns the value of the Metrics field in QueryInfo_SdkV2 as
// a QueryMetrics_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryInfo_SdkV2) GetMetrics(ctx context.Context) (QueryMetrics_SdkV2, bool) {
	var e QueryMetrics_SdkV2
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return e, false
	}
	var v []QueryMetrics_SdkV2
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetrics sets the value of the Metrics field in QueryInfo_SdkV2.
func (m *QueryInfo_SdkV2) SetMetrics(ctx context.Context, v QueryMetrics_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	m.Metrics = types.ListValueMust(t, vs)
}

// GetQuerySource returns the value of the QuerySource field in QueryInfo_SdkV2 as
// a ExternalQuerySource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryInfo_SdkV2) GetQuerySource(ctx context.Context) (ExternalQuerySource_SdkV2, bool) {
	var e ExternalQuerySource_SdkV2
	if m.QuerySource.IsNull() || m.QuerySource.IsUnknown() {
		return e, false
	}
	var v []ExternalQuerySource_SdkV2
	d := m.QuerySource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuerySource sets the value of the QuerySource field in QueryInfo_SdkV2.
func (m *QueryInfo_SdkV2) SetQuerySource(ctx context.Context, v ExternalQuerySource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_source"]
	m.QuerySource = types.ListValueMust(t, vs)
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

func (to *QueryList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryList_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *QueryList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryList_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m QueryList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(LegacyQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryList_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"count":     m.Count,
			"page":      m.Page,
			"page_size": m.PageSize,
			"results":   m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryList_SdkV2) GetResults(ctx context.Context) ([]LegacyQuery_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []LegacyQuery_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in QueryList_SdkV2.
func (m *QueryList_SdkV2) SetResults(ctx context.Context, v []LegacyQuery_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
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
	// Total number of file bytes in all tables not read due to pruning
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
	// Total number of file bytes in all tables read
	ReadFilesBytes types.Int64 `tfsdk:"read_files_bytes"`
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
	TaskTimeOverTimeRange types.List `tfsdk:"task_time_over_time_range"`
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

func (to *QueryMetrics_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryMetrics_SdkV2) {
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

func (to *QueryMetrics_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryMetrics_SdkV2) {
	if !from.TaskTimeOverTimeRange.IsNull() && !from.TaskTimeOverTimeRange.IsUnknown() {
		if toTaskTimeOverTimeRange, ok := to.GetTaskTimeOverTimeRange(ctx); ok {
			if fromTaskTimeOverTimeRange, ok := from.GetTaskTimeOverTimeRange(ctx); ok {
				toTaskTimeOverTimeRange.SyncFieldsDuringRead(ctx, fromTaskTimeOverTimeRange)
				to.SetTaskTimeOverTimeRange(ctx, toTaskTimeOverTimeRange)
			}
		}
	}
}

func (m QueryMetrics_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["read_files_bytes"] = attrs["read_files_bytes"].SetOptional()
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
	attrs["task_time_over_time_range"] = attrs["task_time_over_time_range"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m QueryMetrics_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"task_time_over_time_range": reflect.TypeOf(TaskTimeOverRange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryMetrics_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryMetrics_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compilation_time_ms":                    m.CompilationTimeMs,
			"execution_time_ms":                      m.ExecutionTimeMs,
			"network_sent_bytes":                     m.NetworkSentBytes,
			"overloading_queue_start_timestamp":      m.OverloadingQueueStartTimestamp,
			"photon_total_time_ms":                   m.PhotonTotalTimeMs,
			"projected_remaining_task_total_time_ms": m.ProjectedRemainingTaskTotalTimeMs,
			"projected_remaining_wallclock_time_ms":  m.ProjectedRemainingWallclockTimeMs,
			"provisioning_queue_start_timestamp":     m.ProvisioningQueueStartTimestamp,
			"pruned_bytes":                           m.PrunedBytes,
			"pruned_files_count":                     m.PrunedFilesCount,
			"query_compilation_start_timestamp":      m.QueryCompilationStartTimestamp,
			"read_bytes":                             m.ReadBytes,
			"read_cache_bytes":                       m.ReadCacheBytes,
			"read_files_bytes":                       m.ReadFilesBytes,
			"read_files_count":                       m.ReadFilesCount,
			"read_partitions_count":                  m.ReadPartitionsCount,
			"read_remote_bytes":                      m.ReadRemoteBytes,
			"remaining_task_count":                   m.RemainingTaskCount,
			"result_fetch_time_ms":                   m.ResultFetchTimeMs,
			"result_from_cache":                      m.ResultFromCache,
			"rows_produced_count":                    m.RowsProducedCount,
			"rows_read_count":                        m.RowsReadCount,
			"runnable_tasks":                         m.RunnableTasks,
			"spill_to_disk_bytes":                    m.SpillToDiskBytes,
			"task_time_over_time_range":              m.TaskTimeOverTimeRange,
			"task_total_time_ms":                     m.TaskTotalTimeMs,
			"total_time_ms":                          m.TotalTimeMs,
			"work_to_be_done":                        m.WorkToBeDone,
			"write_remote_bytes":                     m.WriteRemoteBytes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryMetrics_SdkV2) Type(ctx context.Context) attr.Type {
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
			"read_files_bytes":                       types.Int64Type,
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
			"task_time_over_time_range": basetypes.ListType{
				ElemType: TaskTimeOverRange_SdkV2{}.Type(ctx),
			},
			"task_total_time_ms": types.Int64Type,
			"total_time_ms":      types.Int64Type,
			"work_to_be_done":    types.Int64Type,
			"write_remote_bytes": types.Int64Type,
		},
	}
}

// GetTaskTimeOverTimeRange returns the value of the TaskTimeOverTimeRange field in QueryMetrics_SdkV2 as
// a TaskTimeOverRange_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryMetrics_SdkV2) GetTaskTimeOverTimeRange(ctx context.Context) (TaskTimeOverRange_SdkV2, bool) {
	var e TaskTimeOverRange_SdkV2
	if m.TaskTimeOverTimeRange.IsNull() || m.TaskTimeOverTimeRange.IsUnknown() {
		return e, false
	}
	var v []TaskTimeOverRange_SdkV2
	d := m.TaskTimeOverTimeRange.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTaskTimeOverTimeRange sets the value of the TaskTimeOverTimeRange field in QueryMetrics_SdkV2.
func (m *QueryMetrics_SdkV2) SetTaskTimeOverTimeRange(ctx context.Context, v TaskTimeOverRange_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["task_time_over_time_range"]
	m.TaskTimeOverTimeRange = types.ListValueMust(t, vs)
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

func (to *QueryOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryOptions_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (to *QueryOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryOptions_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (m QueryOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(Parameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog":           m.Catalog,
			"moved_to_trash_at": m.MovedToTrashAt,
			"parameters":        m.Parameters,
			"schema":            m.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryOptions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryOptions_SdkV2) GetParameters(ctx context.Context) ([]Parameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []Parameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in QueryOptions_SdkV2.
func (m *QueryOptions_SdkV2) SetParameters(ctx context.Context, v []Parameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
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

func (to *QueryParameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryParameter_SdkV2) {
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

func (to *QueryParameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryParameter_SdkV2) {
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

func (m QueryParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m QueryParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"date_range_value":   m.DateRangeValue,
			"date_value":         m.DateValue,
			"enum_value":         m.EnumValue,
			"name":               m.Name,
			"numeric_value":      m.NumericValue,
			"query_backed_value": m.QueryBackedValue,
			"text_value":         m.TextValue,
			"title":              m.Title,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryParameter_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryParameter_SdkV2) GetDateRangeValue(ctx context.Context) (DateRangeValue_SdkV2, bool) {
	var e DateRangeValue_SdkV2
	if m.DateRangeValue.IsNull() || m.DateRangeValue.IsUnknown() {
		return e, false
	}
	var v []DateRangeValue_SdkV2
	d := m.DateRangeValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateRangeValue sets the value of the DateRangeValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetDateRangeValue(ctx context.Context, v DateRangeValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["date_range_value"]
	m.DateRangeValue = types.ListValueMust(t, vs)
}

// GetDateValue returns the value of the DateValue field in QueryParameter_SdkV2 as
// a DateValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryParameter_SdkV2) GetDateValue(ctx context.Context) (DateValue_SdkV2, bool) {
	var e DateValue_SdkV2
	if m.DateValue.IsNull() || m.DateValue.IsUnknown() {
		return e, false
	}
	var v []DateValue_SdkV2
	d := m.DateValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDateValue sets the value of the DateValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetDateValue(ctx context.Context, v DateValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["date_value"]
	m.DateValue = types.ListValueMust(t, vs)
}

// GetEnumValue returns the value of the EnumValue field in QueryParameter_SdkV2 as
// a EnumValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryParameter_SdkV2) GetEnumValue(ctx context.Context) (EnumValue_SdkV2, bool) {
	var e EnumValue_SdkV2
	if m.EnumValue.IsNull() || m.EnumValue.IsUnknown() {
		return e, false
	}
	var v []EnumValue_SdkV2
	d := m.EnumValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnumValue sets the value of the EnumValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetEnumValue(ctx context.Context, v EnumValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["enum_value"]
	m.EnumValue = types.ListValueMust(t, vs)
}

// GetNumericValue returns the value of the NumericValue field in QueryParameter_SdkV2 as
// a NumericValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryParameter_SdkV2) GetNumericValue(ctx context.Context) (NumericValue_SdkV2, bool) {
	var e NumericValue_SdkV2
	if m.NumericValue.IsNull() || m.NumericValue.IsUnknown() {
		return e, false
	}
	var v []NumericValue_SdkV2
	d := m.NumericValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNumericValue sets the value of the NumericValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetNumericValue(ctx context.Context, v NumericValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["numeric_value"]
	m.NumericValue = types.ListValueMust(t, vs)
}

// GetQueryBackedValue returns the value of the QueryBackedValue field in QueryParameter_SdkV2 as
// a QueryBackedValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryParameter_SdkV2) GetQueryBackedValue(ctx context.Context) (QueryBackedValue_SdkV2, bool) {
	var e QueryBackedValue_SdkV2
	if m.QueryBackedValue.IsNull() || m.QueryBackedValue.IsUnknown() {
		return e, false
	}
	var v []QueryBackedValue_SdkV2
	d := m.QueryBackedValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryBackedValue sets the value of the QueryBackedValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetQueryBackedValue(ctx context.Context, v QueryBackedValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_backed_value"]
	m.QueryBackedValue = types.ListValueMust(t, vs)
}

// GetTextValue returns the value of the TextValue field in QueryParameter_SdkV2 as
// a TextValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryParameter_SdkV2) GetTextValue(ctx context.Context) (TextValue_SdkV2, bool) {
	var e TextValue_SdkV2
	if m.TextValue.IsNull() || m.TextValue.IsUnknown() {
		return e, false
	}
	var v []TextValue_SdkV2
	d := m.TextValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTextValue sets the value of the TextValue field in QueryParameter_SdkV2.
func (m *QueryParameter_SdkV2) SetTextValue(ctx context.Context, v TextValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["text_value"]
	m.TextValue = types.ListValueMust(t, vs)
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

func (to *QueryPostContent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryPostContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *QueryPostContent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryPostContent_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m QueryPostContent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryPostContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryPostContent_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryPostContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_source_id": m.DataSourceId,
			"description":    m.Description,
			"name":           m.Name,
			"options":        m.Options,
			"parent":         m.Parent,
			"query":          m.Query,
			"run_as_role":    m.RunAsRole,
			"tags":           m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryPostContent_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *QueryPostContent_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in QueryPostContent_SdkV2.
func (m *QueryPostContent_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type RepeatedEndpointConfPairs_SdkV2 struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs"`
}

func (to *RepeatedEndpointConfPairs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepeatedEndpointConfPairs_SdkV2) {
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

func (to *RepeatedEndpointConfPairs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepeatedEndpointConfPairs_SdkV2) {
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

func (m RepeatedEndpointConfPairs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepeatedEndpointConfPairs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config_pair":         reflect.TypeOf(EndpointConfPair_SdkV2{}),
		"configuration_pairs": reflect.TypeOf(EndpointConfPair_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepeatedEndpointConfPairs_SdkV2
// only implements ToObjectValue() and Type().
func (m RepeatedEndpointConfPairs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_pair":         m.ConfigPair,
			"configuration_pairs": m.ConfigurationPairs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepeatedEndpointConfPairs_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *RepeatedEndpointConfPairs_SdkV2) GetConfigPair(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if m.ConfigPair.IsNull() || m.ConfigPair.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := m.ConfigPair.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigPair sets the value of the ConfigPair field in RepeatedEndpointConfPairs_SdkV2.
func (m *RepeatedEndpointConfPairs_SdkV2) SetConfigPair(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config_pair"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ConfigPair = types.ListValueMust(t, vs)
}

// GetConfigurationPairs returns the value of the ConfigurationPairs field in RepeatedEndpointConfPairs_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RepeatedEndpointConfPairs_SdkV2) GetConfigurationPairs(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if m.ConfigurationPairs.IsNull() || m.ConfigurationPairs.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := m.ConfigurationPairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfigurationPairs sets the value of the ConfigurationPairs field in RepeatedEndpointConfPairs_SdkV2.
func (m *RepeatedEndpointConfPairs_SdkV2) SetConfigurationPairs(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration_pairs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ConfigurationPairs = types.ListValueMust(t, vs)
}

type RestoreDashboardRequest_SdkV2 struct {
	DashboardId types.String `tfsdk:"-"`
}

func (to *RestoreDashboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreDashboardRequest_SdkV2) {
}

func (to *RestoreDashboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreDashboardRequest_SdkV2) {
}

func (m RestoreDashboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id": m.DashboardId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type RestoreQueriesLegacyRequest_SdkV2 struct {
	QueryId types.String `tfsdk:"-"`
}

func (to *RestoreQueriesLegacyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreQueriesLegacyRequest_SdkV2) {
}

func (to *RestoreQueriesLegacyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreQueriesLegacyRequest_SdkV2) {
}

func (m RestoreQueriesLegacyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreQueriesLegacyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreQueriesLegacyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreQueriesLegacyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": m.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreQueriesLegacyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type RestoreResponse_SdkV2 struct {
}

func (to *RestoreResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreResponse_SdkV2) {
}

func (to *RestoreResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreResponse_SdkV2) {
}

func (m RestoreResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Contains the result data of a single chunk when using `INLINE` disposition.
// When using `EXTERNAL_LINKS` disposition, the array `external_links` is used
// instead to provide URLs to the result data in cloud storage. Exactly one of
// these alternatives is used. (While the `external_links` array prepares the
// API to return multiple links in a single response. Currently only a single
// link is returned.)
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
	// with a :method:statementexecution/getstatementresultchunkn request.
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

func (to *ResultData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData_SdkV2) {
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

func (to *ResultData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultData_SdkV2) {
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

func (m ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array":     reflect.TypeOf(types.String{}),
		"external_links": reflect.TypeOf(ExternalLink_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               m.ByteCount,
			"chunk_index":              m.ChunkIndex,
			"data_array":               m.DataArray,
			"external_links":           m.ExternalLinks,
			"next_chunk_index":         m.NextChunkIndex,
			"next_chunk_internal_link": m.NextChunkInternalLink,
			"row_count":                m.RowCount,
			"row_offset":               m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultData_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.String, bool) {
	if m.DataArray.IsNull() || m.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// GetExternalLinks returns the value of the ExternalLinks field in ResultData_SdkV2 as
// a slice of ExternalLink_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData_SdkV2) GetExternalLinks(ctx context.Context) ([]ExternalLink_SdkV2, bool) {
	if m.ExternalLinks.IsNull() || m.ExternalLinks.IsUnknown() {
		return nil, false
	}
	var v []ExternalLink_SdkV2
	d := m.ExternalLinks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLinks sets the value of the ExternalLinks field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetExternalLinks(ctx context.Context, v []ExternalLink_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_links"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExternalLinks = types.ListValueMust(t, vs)
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

func (to *ResultManifest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest_SdkV2) {
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

func (to *ResultManifest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultManifest_SdkV2) {
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

func (m ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"chunks": reflect.TypeOf(BaseChunkInfo_SdkV2{}),
		"schema": reflect.TypeOf(ResultSchema_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunks":            m.Chunks,
			"format":            m.Format,
			"schema":            m.Schema,
			"total_byte_count":  m.TotalByteCount,
			"total_chunk_count": m.TotalChunkCount,
			"total_row_count":   m.TotalRowCount,
			"truncated":         m.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResultManifest_SdkV2) GetChunks(ctx context.Context) ([]BaseChunkInfo_SdkV2, bool) {
	if m.Chunks.IsNull() || m.Chunks.IsUnknown() {
		return nil, false
	}
	var v []BaseChunkInfo_SdkV2
	d := m.Chunks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChunks sets the value of the Chunks field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetChunks(ctx context.Context, v []BaseChunkInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["chunks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Chunks = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in ResultManifest_SdkV2 as
// a ResultSchema_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetSchema(ctx context.Context) (ResultSchema_SdkV2, bool) {
	var e ResultSchema_SdkV2
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return e, false
	}
	var v []ResultSchema_SdkV2
	d := m.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetSchema(ctx context.Context, v ResultSchema_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	m.Schema = types.ListValueMust(t, vs)
}

// The schema is an ordered list of column descriptions.
type ResultSchema_SdkV2 struct {
	ColumnCount types.Int64 `tfsdk:"column_count"`

	Columns types.List `tfsdk:"columns"`
}

func (to *ResultSchema_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultSchema_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *ResultSchema_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultSchema_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m ResultSchema_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultSchema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultSchema_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultSchema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": m.ColumnCount,
			"columns":      m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultSchema_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResultSchema_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultSchema_SdkV2.
func (m *ResultSchema_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type ServiceError_SdkV2 struct {
	ErrorCode types.String `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

func (to *ServiceError_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServiceError_SdkV2) {
}

func (to *ServiceError_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServiceError_SdkV2) {
}

func (m ServiceError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ServiceError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceError_SdkV2
// only implements ToObjectValue() and Type().
func (m ServiceError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_code": m.ErrorCode,
			"message":    m.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServiceError_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *SetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m SetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAccessControlList returns the value of the AccessControlList field in SetRequest_SdkV2 as
// a slice of AccessControl_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SetRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetRequest_SdkV2.
func (m *SetRequest_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type SetResponse_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type"`
}

func (to *SetResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetResponse_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SetResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetResponse_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m SetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControl_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *SetResponse_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControl_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControl_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetResponse_SdkV2.
func (m *SetResponse_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControl_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// Sets the workspace level warehouse configuration that is shared by all SQL
// warehouses in this workspace.
//
// This is idempotent.
type SetWorkspaceWarehouseConfigRequest_SdkV2 struct {
	// Optional: Channel selection details
	Channel types.List `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.List `tfsdk:"config_param"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config"`
	// Enable Serverless compute for SQL warehouses
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
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
	// AWS Only: The instance profile used to pass an IAM role to the SQL
	// warehouses. This configuration is also applied to the workspace's
	// serverless compute for notebooks and jobs.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters types.List `tfsdk:"sql_configuration_parameters"`
}

func (to *SetWorkspaceWarehouseConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetWorkspaceWarehouseConfigRequest_SdkV2) {
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

func (to *SetWorkspaceWarehouseConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetWorkspaceWarehouseConfigRequest_SdkV2) {
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

func (m SetWorkspaceWarehouseConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["channel"] = attrs["channel"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["config_param"] = attrs["config_param"].SetOptional()
	attrs["config_param"] = attrs["config_param"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["data_access_config"] = attrs["data_access_config"].SetOptional()
	attrs["enable_serverless_compute"] = attrs["enable_serverless_compute"].SetOptional()
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetWorkspaceWarehouseConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m SetWorkspaceWarehouseConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel":                      m.Channel,
			"config_param":                 m.ConfigParam,
			"data_access_config":           m.DataAccessConfig,
			"enable_serverless_compute":    m.EnableServerlessCompute,
			"enabled_warehouse_types":      m.EnabledWarehouseTypes,
			"global_param":                 m.GlobalParam,
			"google_service_account":       m.GoogleServiceAccount,
			"instance_profile_arn":         m.InstanceProfileArn,
			"security_policy":              m.SecurityPolicy,
			"sql_configuration_parameters": m.SqlConfigurationParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetWorkspaceWarehouseConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
			"enable_serverless_compute": types.BoolType,
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
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetChannel(ctx context.Context) (Channel_SdkV2, bool) {
	var e Channel_SdkV2
	if m.Channel.IsNull() || m.Channel.IsUnknown() {
		return e, false
	}
	var v []Channel_SdkV2
	d := m.Channel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetChannel sets the value of the Channel field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetChannel(ctx context.Context, v Channel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["channel"]
	m.Channel = types.ListValueMust(t, vs)
}

// GetConfigParam returns the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetConfigParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.ConfigParam.IsNull() || m.ConfigParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.ConfigParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfigParam sets the value of the ConfigParam field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetConfigParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config_param"]
	m.ConfigParam = types.ListValueMust(t, vs)
}

// GetDataAccessConfig returns the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a slice of EndpointConfPair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetDataAccessConfig(ctx context.Context) ([]EndpointConfPair_SdkV2, bool) {
	if m.DataAccessConfig.IsNull() || m.DataAccessConfig.IsUnknown() {
		return nil, false
	}
	var v []EndpointConfPair_SdkV2
	d := m.DataAccessConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataAccessConfig sets the value of the DataAccessConfig field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetDataAccessConfig(ctx context.Context, v []EndpointConfPair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_access_config"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataAccessConfig = types.ListValueMust(t, vs)
}

// GetEnabledWarehouseTypes returns the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a slice of WarehouseTypePair_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetEnabledWarehouseTypes(ctx context.Context) ([]WarehouseTypePair_SdkV2, bool) {
	if m.EnabledWarehouseTypes.IsNull() || m.EnabledWarehouseTypes.IsUnknown() {
		return nil, false
	}
	var v []WarehouseTypePair_SdkV2
	d := m.EnabledWarehouseTypes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnabledWarehouseTypes sets the value of the EnabledWarehouseTypes field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetEnabledWarehouseTypes(ctx context.Context, v []WarehouseTypePair_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["enabled_warehouse_types"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnabledWarehouseTypes = types.ListValueMust(t, vs)
}

// GetGlobalParam returns the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetGlobalParam(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.GlobalParam.IsNull() || m.GlobalParam.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.GlobalParam.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlobalParam sets the value of the GlobalParam field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetGlobalParam(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["global_param"]
	m.GlobalParam = types.ListValueMust(t, vs)
}

// GetSqlConfigurationParameters returns the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest_SdkV2 as
// a RepeatedEndpointConfPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) GetSqlConfigurationParameters(ctx context.Context) (RepeatedEndpointConfPairs_SdkV2, bool) {
	var e RepeatedEndpointConfPairs_SdkV2
	if m.SqlConfigurationParameters.IsNull() || m.SqlConfigurationParameters.IsUnknown() {
		return e, false
	}
	var v []RepeatedEndpointConfPairs_SdkV2
	d := m.SqlConfigurationParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlConfigurationParameters sets the value of the SqlConfigurationParameters field in SetWorkspaceWarehouseConfigRequest_SdkV2.
func (m *SetWorkspaceWarehouseConfigRequest_SdkV2) SetSqlConfigurationParameters(ctx context.Context, v RepeatedEndpointConfPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_configuration_parameters"]
	m.SqlConfigurationParameters = types.ListValueMust(t, vs)
}

type SetWorkspaceWarehouseConfigResponse_SdkV2 struct {
}

func (to *SetWorkspaceWarehouseConfigResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (to *SetWorkspaceWarehouseConfigResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetWorkspaceWarehouseConfigResponse_SdkV2) {
}

func (m SetWorkspaceWarehouseConfigResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetWorkspaceWarehouseConfigResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetWorkspaceWarehouseConfigResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetWorkspaceWarehouseConfigResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetWorkspaceWarehouseConfigResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetWorkspaceWarehouseConfigResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StartRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *StartRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartRequest_SdkV2) {
}

func (to *StartRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StartRequest_SdkV2) {
}

func (m StartRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m StartRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StartWarehouseResponse_SdkV2 struct {
}

func (to *StartWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartWarehouseResponse_SdkV2) {
}

func (to *StartWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StartWarehouseResponse_SdkV2) {
}

func (m StartWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StartWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StartWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m StartWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *StatementParameterListItem_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementParameterListItem_SdkV2) {
}

func (to *StatementParameterListItem_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementParameterListItem_SdkV2) {
}

func (m StatementParameterListItem_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StatementParameterListItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementParameterListItem_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementParameterListItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  m.Name,
			"type":  m.Type_,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementParameterListItem_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *StatementResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementResponse_SdkV2) {
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

func (to *StatementResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementResponse_SdkV2) {
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

func (m StatementResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StatementResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":   reflect.TypeOf(ResultData_SdkV2{}),
		"status":   reflect.TypeOf(StatementStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":     m.Manifest,
			"result":       m.Result,
			"statement_id": m.StatementId,
			"status":       m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *StatementResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := m.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	m.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in StatementResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in StatementResponse_SdkV2 as
// a StatementStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse_SdkV2) GetStatus(ctx context.Context) (StatementStatus_SdkV2, bool) {
	var e StatementStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []StatementStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetStatus(ctx context.Context, v StatementStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus_SdkV2 struct {
	Error types.List `tfsdk:"error"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accompanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State types.String `tfsdk:"state"`
}

func (to *StatementStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementStatus_SdkV2) {
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

func (to *StatementStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementStatus_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m StatementStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StatementStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(ServiceError_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error": m.Error,
			"state": m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *StatementStatus_SdkV2) GetError(ctx context.Context) (ServiceError_SdkV2, bool) {
	var e ServiceError_SdkV2
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v []ServiceError_SdkV2
	d := m.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in StatementStatus_SdkV2.
func (m *StatementStatus_SdkV2) SetError(ctx context.Context, v ServiceError_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	m.Error = types.ListValueMust(t, vs)
}

type StopRequest_SdkV2 struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (to *StopRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopRequest_SdkV2) {
}

func (to *StopRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StopRequest_SdkV2) {
}

func (m StopRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StopRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m StopRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StopWarehouseResponse_SdkV2 struct {
}

func (to *StopWarehouseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopWarehouseResponse_SdkV2) {
}

func (to *StopWarehouseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StopWarehouseResponse_SdkV2) {
}

func (m StopWarehouseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopWarehouseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StopWarehouseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopWarehouseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StopWarehouseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m StopWarehouseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Success_SdkV2 struct {
	Message types.String `tfsdk:"message"`
}

func (to *Success_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Success_SdkV2) {
}

func (to *Success_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Success_SdkV2) {
}

func (m Success_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Success_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Success_SdkV2
// only implements ToObjectValue() and Type().
func (m Success_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Success_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TaskTimeOverRange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TaskTimeOverRange_SdkV2) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
}

func (to *TaskTimeOverRange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TaskTimeOverRange_SdkV2) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
}

func (m TaskTimeOverRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TaskTimeOverRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entries": reflect.TypeOf(TaskTimeOverRangeEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRange_SdkV2
// only implements ToObjectValue() and Type().
func (m TaskTimeOverRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entries":  m.Entries,
			"interval": m.Interval,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TaskTimeOverRange_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *TaskTimeOverRange_SdkV2) GetEntries(ctx context.Context) ([]TaskTimeOverRangeEntry_SdkV2, bool) {
	if m.Entries.IsNull() || m.Entries.IsUnknown() {
		return nil, false
	}
	var v []TaskTimeOverRangeEntry_SdkV2
	d := m.Entries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntries sets the value of the Entries field in TaskTimeOverRange_SdkV2.
func (m *TaskTimeOverRange_SdkV2) SetEntries(ctx context.Context, v []TaskTimeOverRangeEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entries = types.ListValueMust(t, vs)
}

type TaskTimeOverRangeEntry_SdkV2 struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	TaskCompletedTimeMs types.Int64 `tfsdk:"task_completed_time_ms"`
}

func (to *TaskTimeOverRangeEntry_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TaskTimeOverRangeEntry_SdkV2) {
}

func (to *TaskTimeOverRangeEntry_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TaskTimeOverRangeEntry_SdkV2) {
}

func (m TaskTimeOverRangeEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TaskTimeOverRangeEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskTimeOverRangeEntry_SdkV2
// only implements ToObjectValue() and Type().
func (m TaskTimeOverRangeEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"task_completed_time_ms": m.TaskCompletedTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TaskTimeOverRangeEntry_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TerminationReason_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TerminationReason_SdkV2) {
}

func (to *TerminationReason_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TerminationReason_SdkV2) {
}

func (m TerminationReason_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TerminationReason_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationReason_SdkV2
// only implements ToObjectValue() and Type().
func (m TerminationReason_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":       m.Code,
			"parameters": m.Parameters,
			"type":       m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TerminationReason_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *TerminationReason_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in TerminationReason_SdkV2.
func (m *TerminationReason_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.MapValueMust(t, vs)
}

type TextValue_SdkV2 struct {
	Value types.String `tfsdk:"value"`
}

func (to *TextValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TextValue_SdkV2) {
}

func (to *TextValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TextValue_SdkV2) {
}

func (m TextValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TextValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TextValue_SdkV2
// only implements ToObjectValue() and Type().
func (m TextValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TextValue_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TimeRange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeRange_SdkV2) {
}

func (to *TimeRange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TimeRange_SdkV2) {
}

func (m TimeRange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TimeRange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeRange_SdkV2
// only implements ToObjectValue() and Type().
func (m TimeRange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   m.EndTimeMs,
			"start_time_ms": m.StartTimeMs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TimeRange_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TransferOwnershipObjectId_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransferOwnershipObjectId_SdkV2) {
}

func (to *TransferOwnershipObjectId_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransferOwnershipObjectId_SdkV2) {
}

func (m TransferOwnershipObjectId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TransferOwnershipObjectId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipObjectId_SdkV2
// only implements ToObjectValue() and Type().
func (m TransferOwnershipObjectId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner": m.NewOwner,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransferOwnershipObjectId_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TransferOwnershipRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransferOwnershipRequest_SdkV2) {
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

func (to *TransferOwnershipRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransferOwnershipRequest_SdkV2) {
	if !from.ObjectId.IsNull() && !from.ObjectId.IsUnknown() {
		if toObjectId, ok := to.GetObjectId(ctx); ok {
			if fromObjectId, ok := from.GetObjectId(ctx); ok {
				toObjectId.SyncFieldsDuringRead(ctx, fromObjectId)
				to.SetObjectId(ctx, toObjectId)
			}
		}
	}
}

func (m TransferOwnershipRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["new_owner"] = attrs["new_owner"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_id"] = attrs["object_id"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransferOwnershipRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransferOwnershipRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object_id": reflect.TypeOf(TransferOwnershipObjectId_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransferOwnershipRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TransferOwnershipRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"new_owner":   m.NewOwner,
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransferOwnershipRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"new_owner": types.StringType,
			"object_id": basetypes.ListType{
				ElemType: TransferOwnershipObjectId_SdkV2{}.Type(ctx),
			},
			"object_type": types.StringType,
		},
	}
}

// GetObjectId returns the value of the ObjectId field in TransferOwnershipRequest_SdkV2 as
// a TransferOwnershipObjectId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *TransferOwnershipRequest_SdkV2) GetObjectId(ctx context.Context) (TransferOwnershipObjectId_SdkV2, bool) {
	var e TransferOwnershipObjectId_SdkV2
	if m.ObjectId.IsNull() || m.ObjectId.IsUnknown() {
		return e, false
	}
	var v []TransferOwnershipObjectId_SdkV2
	d := m.ObjectId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetObjectId sets the value of the ObjectId field in TransferOwnershipRequest_SdkV2.
func (m *TransferOwnershipRequest_SdkV2) SetObjectId(ctx context.Context, v TransferOwnershipObjectId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["object_id"]
	m.ObjectId = types.ListValueMust(t, vs)
}

type TrashAlertRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashAlertRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashAlertRequest_SdkV2) {
}

func (to *TrashAlertRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TrashAlertRequest_SdkV2) {
}

func (m TrashAlertRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TrashAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TrashAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TrashAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type TrashAlertV2Request_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashAlertV2Request_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashAlertV2Request_SdkV2) {
}

func (to *TrashAlertV2Request_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TrashAlertV2Request_SdkV2) {
}

func (m TrashAlertV2Request_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TrashAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (m TrashAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TrashAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type TrashQueryRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *TrashQueryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrashQueryRequest_SdkV2) {
}

func (to *TrashQueryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TrashQueryRequest_SdkV2) {
}

func (m TrashQueryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TrashQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrashQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TrashQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TrashQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *UpdateAlertRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertRequest_SdkV2) {
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

func (to *UpdateAlertRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertRequest_SdkV2) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (m UpdateAlertRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetOptional()
	attrs["alert"] = attrs["alert"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateAlertRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(UpdateAlertRequestAlert_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAlertRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":                     m.Alert,
			"auto_resolve_display_name": m.AutoResolveDisplayName,
			"id":                        m.Id,
			"update_mask":               m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAlertRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateAlertRequest_SdkV2) GetAlert(ctx context.Context) (UpdateAlertRequestAlert_SdkV2, bool) {
	var e UpdateAlertRequestAlert_SdkV2
	if m.Alert.IsNull() || m.Alert.IsUnknown() {
		return e, false
	}
	var v []UpdateAlertRequestAlert_SdkV2
	d := m.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in UpdateAlertRequest_SdkV2.
func (m *UpdateAlertRequest_SdkV2) SetAlert(ctx context.Context, v UpdateAlertRequestAlert_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	m.Alert = types.ListValueMust(t, vs)
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

func (to *UpdateAlertRequestAlert_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertRequestAlert_SdkV2) {
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

func (to *UpdateAlertRequestAlert_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertRequestAlert_SdkV2) {
	if !from.Condition.IsNull() && !from.Condition.IsUnknown() {
		if toCondition, ok := to.GetCondition(ctx); ok {
			if fromCondition, ok := from.GetCondition(ctx); ok {
				toCondition.SyncFieldsDuringRead(ctx, fromCondition)
				to.SetCondition(ctx, toCondition)
			}
		}
	}
}

func (m UpdateAlertRequestAlert_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateAlertRequestAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition": reflect.TypeOf(AlertCondition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertRequestAlert_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAlertRequestAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":            m.Condition,
			"custom_body":          m.CustomBody,
			"custom_subject":       m.CustomSubject,
			"display_name":         m.DisplayName,
			"notify_on_ok":         m.NotifyOnOk,
			"owner_user_name":      m.OwnerUserName,
			"query_id":             m.QueryId,
			"seconds_to_retrigger": m.SecondsToRetrigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAlertRequestAlert_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateAlertRequestAlert_SdkV2) GetCondition(ctx context.Context) (AlertCondition_SdkV2, bool) {
	var e AlertCondition_SdkV2
	if m.Condition.IsNull() || m.Condition.IsUnknown() {
		return e, false
	}
	var v []AlertCondition_SdkV2
	d := m.Condition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCondition sets the value of the Condition field in UpdateAlertRequestAlert_SdkV2.
func (m *UpdateAlertRequestAlert_SdkV2) SetCondition(ctx context.Context, v AlertCondition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["condition"]
	m.Condition = types.ListValueMust(t, vs)
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

func (to *UpdateAlertV2Request_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAlertV2Request_SdkV2) {
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

func (to *UpdateAlertV2Request_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAlertV2Request_SdkV2) {
	if !from.Alert.IsNull() && !from.Alert.IsUnknown() {
		if toAlert, ok := to.GetAlert(ctx); ok {
			if fromAlert, ok := from.GetAlert(ctx); ok {
				toAlert.SyncFieldsDuringRead(ctx, fromAlert)
				to.SetAlert(ctx, toAlert)
			}
		}
	}
}

func (m UpdateAlertV2Request_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert"] = attrs["alert"].SetRequired()
	attrs["alert"] = attrs["alert"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateAlertV2Request_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert": reflect.TypeOf(AlertV2_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAlertV2Request_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAlertV2Request_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":       m.Alert,
			"id":          m.Id,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAlertV2Request_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateAlertV2Request_SdkV2) GetAlert(ctx context.Context) (AlertV2_SdkV2, bool) {
	var e AlertV2_SdkV2
	if m.Alert.IsNull() || m.Alert.IsUnknown() {
		return e, false
	}
	var v []AlertV2_SdkV2
	d := m.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in UpdateAlertV2Request_SdkV2.
func (m *UpdateAlertV2Request_SdkV2) SetAlert(ctx context.Context, v AlertV2_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	m.Alert = types.ListValueMust(t, vs)
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

func (to *UpdateQueryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateQueryRequest_SdkV2) {
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

func (to *UpdateQueryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateQueryRequest_SdkV2) {
	if !from.Query.IsNull() && !from.Query.IsUnknown() {
		if toQuery, ok := to.GetQuery(ctx); ok {
			if fromQuery, ok := from.GetQuery(ctx); ok {
				toQuery.SyncFieldsDuringRead(ctx, fromQuery)
				to.SetQuery(ctx, toQuery)
			}
		}
	}
}

func (m UpdateQueryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_resolve_display_name"] = attrs["auto_resolve_display_name"].SetOptional()
	attrs["query"] = attrs["query"].SetOptional()
	attrs["query"] = attrs["query"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateQueryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(UpdateQueryRequestQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateQueryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_resolve_display_name": m.AutoResolveDisplayName,
			"id":                        m.Id,
			"query":                     m.Query,
			"update_mask":               m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateQueryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateQueryRequest_SdkV2) GetQuery(ctx context.Context) (UpdateQueryRequestQuery_SdkV2, bool) {
	var e UpdateQueryRequestQuery_SdkV2
	if m.Query.IsNull() || m.Query.IsUnknown() {
		return e, false
	}
	var v []UpdateQueryRequestQuery_SdkV2
	d := m.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in UpdateQueryRequest_SdkV2.
func (m *UpdateQueryRequest_SdkV2) SetQuery(ctx context.Context, v UpdateQueryRequestQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	m.Query = types.ListValueMust(t, vs)
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

func (to *UpdateQueryRequestQuery_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateQueryRequestQuery_SdkV2) {
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

func (to *UpdateQueryRequestQuery_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateQueryRequestQuery_SdkV2) {
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

func (m UpdateQueryRequestQuery_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateQueryRequestQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(QueryParameter_SdkV2{}),
		"tags":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQueryRequestQuery_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateQueryRequestQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_auto_limit": m.ApplyAutoLimit,
			"catalog":          m.Catalog,
			"description":      m.Description,
			"display_name":     m.DisplayName,
			"owner_user_name":  m.OwnerUserName,
			"parameters":       m.Parameters,
			"query_text":       m.QueryText,
			"run_as_mode":      m.RunAsMode,
			"schema":           m.Schema,
			"tags":             m.Tags,
			"warehouse_id":     m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateQueryRequestQuery_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateQueryRequestQuery_SdkV2) GetParameters(ctx context.Context) ([]QueryParameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []QueryParameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in UpdateQueryRequestQuery_SdkV2.
func (m *UpdateQueryRequestQuery_SdkV2) SetParameters(ctx context.Context, v []QueryParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in UpdateQueryRequestQuery_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateQueryRequestQuery_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in UpdateQueryRequestQuery_SdkV2.
func (m *UpdateQueryRequestQuery_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
}

func (to *UpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateResponse_SdkV2) {
}

func (to *UpdateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateResponse_SdkV2) {
}

func (m UpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *UpdateVisualizationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVisualizationRequest_SdkV2) {
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

func (to *UpdateVisualizationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateVisualizationRequest_SdkV2) {
	if !from.Visualization.IsNull() && !from.Visualization.IsUnknown() {
		if toVisualization, ok := to.GetVisualization(ctx); ok {
			if fromVisualization, ok := from.GetVisualization(ctx); ok {
				toVisualization.SyncFieldsDuringRead(ctx, fromVisualization)
				to.SetVisualization(ctx, toVisualization)
			}
		}
	}
}

func (m UpdateVisualizationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["visualization"] = attrs["visualization"].SetOptional()
	attrs["visualization"] = attrs["visualization"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateVisualizationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"visualization": reflect.TypeOf(UpdateVisualizationRequestVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateVisualizationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":            m.Id,
			"update_mask":   m.UpdateMask,
			"visualization": m.Visualization,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateVisualizationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateVisualizationRequest_SdkV2) GetVisualization(ctx context.Context) (UpdateVisualizationRequestVisualization_SdkV2, bool) {
	var e UpdateVisualizationRequestVisualization_SdkV2
	if m.Visualization.IsNull() || m.Visualization.IsUnknown() {
		return e, false
	}
	var v []UpdateVisualizationRequestVisualization_SdkV2
	d := m.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in UpdateVisualizationRequest_SdkV2.
func (m *UpdateVisualizationRequest_SdkV2) SetVisualization(ctx context.Context, v UpdateVisualizationRequestVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	m.Visualization = types.ListValueMust(t, vs)
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

func (to *UpdateVisualizationRequestVisualization_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVisualizationRequestVisualization_SdkV2) {
}

func (to *UpdateVisualizationRequestVisualization_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateVisualizationRequestVisualization_SdkV2) {
}

func (m UpdateVisualizationRequestVisualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateVisualizationRequestVisualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVisualizationRequestVisualization_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateVisualizationRequestVisualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":          m.DisplayName,
			"serialized_options":    m.SerializedOptions,
			"serialized_query_plan": m.SerializedQueryPlan,
			"type":                  m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateVisualizationRequestVisualization_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *UpdateWidgetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWidgetRequest_SdkV2) {
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

func (to *UpdateWidgetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWidgetRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m UpdateWidgetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateWidgetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(WidgetOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWidgetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWidgetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":     m.DashboardId,
			"id":               m.Id,
			"options":          m.Options,
			"text":             m.Text,
			"visualization_id": m.VisualizationId,
			"width":            m.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWidgetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateWidgetRequest_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in UpdateWidgetRequest_SdkV2.
func (m *UpdateWidgetRequest_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

type User_SdkV2 struct {
	Email types.String `tfsdk:"email"`

	Id types.Int64 `tfsdk:"id"`

	Name types.String `tfsdk:"name"`
}

func (to *User_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User_SdkV2) {
}

func (to *User_SdkV2) SyncFieldsDuringRead(ctx context.Context, from User_SdkV2) {
}

func (m User_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m User_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User_SdkV2
// only implements ToObjectValue() and Type().
func (m User_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email": m.Email,
			"id":    m.Id,
			"name":  m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m User_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Visualization_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Visualization_SdkV2) {
}

func (to *Visualization_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Visualization_SdkV2) {
}

func (m Visualization_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Visualization_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Visualization_SdkV2
// only implements ToObjectValue() and Type().
func (m Visualization_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"display_name":          m.DisplayName,
			"id":                    m.Id,
			"query_id":              m.QueryId,
			"serialized_options":    m.SerializedOptions,
			"serialized_query_plan": m.SerializedQueryPlan,
			"type":                  m.Type_,
			"update_time":           m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Visualization_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *WarehouseAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseAccessControlRequest_SdkV2) {
}

func (to *WarehouseAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehouseAccessControlRequest_SdkV2) {
}

func (m WarehouseAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehouseAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehouseAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehouseAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *WarehouseAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *WarehouseAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehouseAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m WarehouseAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehouseAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WarehousePermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehouseAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehouseAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WarehouseAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]WarehousePermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WarehousePermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WarehouseAccessControlResponse_SdkV2.
func (m *WarehouseAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []WarehousePermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type WarehousePermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WarehousePermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *WarehousePermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehousePermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m WarehousePermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehousePermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermission_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehousePermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehousePermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WarehousePermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in WarehousePermission_SdkV2.
func (m *WarehousePermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type WarehousePermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *WarehousePermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WarehousePermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WarehousePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehousePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehousePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehousePermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WarehousePermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissions_SdkV2.
func (m *WarehousePermissions_SdkV2) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type WarehousePermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WarehousePermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissionsDescription_SdkV2) {
}

func (to *WarehousePermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissionsDescription_SdkV2) {
}

func (m WarehousePermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehousePermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehousePermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehousePermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *WarehousePermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehousePermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WarehousePermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehousePermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WarehousePermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehousePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WarehouseAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehousePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehousePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"warehouse_id":        m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehousePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WarehousePermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]WarehouseAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WarehouseAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WarehousePermissionsRequest_SdkV2.
func (m *WarehousePermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []WarehouseAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// * Configuration values to enable or disable the access to specific warehouse
// types in the workspace.
type WarehouseTypePair_SdkV2 struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled"`

	WarehouseType types.String `tfsdk:"warehouse_type"`
}

func (to *WarehouseTypePair_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WarehouseTypePair_SdkV2) {
}

func (to *WarehouseTypePair_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WarehouseTypePair_SdkV2) {
}

func (m WarehouseTypePair_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WarehouseTypePair_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WarehouseTypePair_SdkV2
// only implements ToObjectValue() and Type().
func (m WarehouseTypePair_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":        m.Enabled,
			"warehouse_type": m.WarehouseType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WarehouseTypePair_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Widget_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Widget_SdkV2) {
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

func (to *Widget_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Widget_SdkV2) {
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

func (m Widget_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Widget_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":       reflect.TypeOf(WidgetOptions_SdkV2{}),
		"visualization": reflect.TypeOf(LegacyVisualization_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Widget_SdkV2
// only implements ToObjectValue() and Type().
func (m Widget_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":            m.Id,
			"options":       m.Options,
			"visualization": m.Visualization,
			"width":         m.Width,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Widget_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Widget_SdkV2) GetOptions(ctx context.Context) (WidgetOptions_SdkV2, bool) {
	var e WidgetOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []WidgetOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in Widget_SdkV2.
func (m *Widget_SdkV2) SetOptions(ctx context.Context, v WidgetOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

// GetVisualization returns the value of the Visualization field in Widget_SdkV2 as
// a LegacyVisualization_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Widget_SdkV2) GetVisualization(ctx context.Context) (LegacyVisualization_SdkV2, bool) {
	var e LegacyVisualization_SdkV2
	if m.Visualization.IsNull() || m.Visualization.IsUnknown() {
		return e, false
	}
	var v []LegacyVisualization_SdkV2
	d := m.Visualization.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVisualization sets the value of the Visualization field in Widget_SdkV2.
func (m *Widget_SdkV2) SetVisualization(ctx context.Context, v LegacyVisualization_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["visualization"]
	m.Visualization = types.ListValueMust(t, vs)
}

type WidgetOptions_SdkV2 struct {
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
	Position types.List `tfsdk:"position"`
	// Custom title of the widget
	Title types.String `tfsdk:"title"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (to *WidgetOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WidgetOptions_SdkV2) {
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

func (to *WidgetOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WidgetOptions_SdkV2) {
	if !from.Position.IsNull() && !from.Position.IsUnknown() {
		if toPosition, ok := to.GetPosition(ctx); ok {
			if fromPosition, ok := from.GetPosition(ctx); ok {
				toPosition.SyncFieldsDuringRead(ctx, fromPosition)
				to.SetPosition(ctx, toPosition)
			}
		}
	}
}

func (m WidgetOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["is_hidden"] = attrs["is_hidden"].SetOptional()
	attrs["parameter_mappings"] = attrs["parameter_mappings"].SetOptional()
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
func (m WidgetOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"position": reflect.TypeOf(WidgetPosition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m WidgetOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":         m.CreatedAt,
			"description":        m.Description,
			"is_hidden":          m.IsHidden,
			"parameter_mappings": m.ParameterMappings,
			"position":           m.Position,
			"title":              m.Title,
			"updated_at":         m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WidgetOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":         types.StringType,
			"description":        types.StringType,
			"is_hidden":          types.BoolType,
			"parameter_mappings": types.ObjectType{},
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
func (m *WidgetOptions_SdkV2) GetPosition(ctx context.Context) (WidgetPosition_SdkV2, bool) {
	var e WidgetPosition_SdkV2
	if m.Position.IsNull() || m.Position.IsUnknown() {
		return e, false
	}
	var v []WidgetPosition_SdkV2
	d := m.Position.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPosition sets the value of the Position field in WidgetOptions_SdkV2.
func (m *WidgetOptions_SdkV2) SetPosition(ctx context.Context, v WidgetPosition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["position"]
	m.Position = types.ListValueMust(t, vs)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition_SdkV2 struct {
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

func (to *WidgetPosition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WidgetPosition_SdkV2) {
}

func (to *WidgetPosition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WidgetPosition_SdkV2) {
}

func (m WidgetPosition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WidgetPosition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WidgetPosition_SdkV2
// only implements ToObjectValue() and Type().
func (m WidgetPosition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_height": m.AutoHeight,
			"col":         m.Col,
			"row":         m.Row,
			"size_x":      m.SizeX,
			"size_y":      m.SizeY,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WidgetPosition_SdkV2) Type(ctx context.Context) attr.Type {
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
