// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package iam_tf

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

type AccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *AccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AccessControlRequest_SdkV2) {
}

func (toState *AccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AccessControlRequest_SdkV2) {
}

func (c AccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AccessControlResponse_SdkV2 struct {
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

func (toState *AccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AccessControlResponse_SdkV2) {
}

func (toState *AccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AccessControlResponse_SdkV2) {
}

func (c AccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(Permission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: Permission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in AccessControlResponse_SdkV2 as
// a slice of Permission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]Permission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []Permission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in AccessControlResponse_SdkV2.
func (o *AccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []Permission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

// represents an identity trying to access a resource - user or a service
// principal group can be a principal of a permission set assignment but an
// actor is always a user or a service principal
type Actor_SdkV2 struct {
	ActorId types.Int64 `tfsdk:"actor_id"`
}

func (toState *Actor_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Actor_SdkV2) {
}

func (toState *Actor_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Actor_SdkV2) {
}

func (c Actor_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["actor_id"] = attrs["actor_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Actor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Actor_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Actor_SdkV2
// only implements ToObjectValue() and Type().
func (o Actor_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id": o.ActorId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Actor_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id": types.Int64Type,
		},
	}
}

type CheckPolicyResponse_SdkV2 struct {
	ConsistencyToken types.List `tfsdk:"consistency_token"`

	IsPermitted types.Bool `tfsdk:"is_permitted"`
}

func (toState *CheckPolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CheckPolicyResponse_SdkV2) {
	if !fromPlan.ConsistencyToken.IsNull() && !fromPlan.ConsistencyToken.IsUnknown() {
		if toStateConsistencyToken, ok := toState.GetConsistencyToken(ctx); ok {
			if fromPlanConsistencyToken, ok := fromPlan.GetConsistencyToken(ctx); ok {
				toStateConsistencyToken.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanConsistencyToken)
				toState.SetConsistencyToken(ctx, toStateConsistencyToken)
			}
		}
	}
}

func (toState *CheckPolicyResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CheckPolicyResponse_SdkV2) {
	if !fromState.ConsistencyToken.IsNull() && !fromState.ConsistencyToken.IsUnknown() {
		if toStateConsistencyToken, ok := toState.GetConsistencyToken(ctx); ok {
			if fromStateConsistencyToken, ok := fromState.GetConsistencyToken(ctx); ok {
				toStateConsistencyToken.SyncFieldsDuringRead(ctx, fromStateConsistencyToken)
				toState.SetConsistencyToken(ctx, toStateConsistencyToken)
			}
		}
	}
}

func (c CheckPolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["consistency_token"] = attrs["consistency_token"].SetRequired()
	attrs["consistency_token"] = attrs["consistency_token"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["is_permitted"] = attrs["is_permitted"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CheckPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CheckPolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"consistency_token": reflect.TypeOf(ConsistencyToken_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CheckPolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CheckPolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"consistency_token": o.ConsistencyToken,
			"is_permitted":      o.IsPermitted,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CheckPolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"consistency_token": basetypes.ListType{
				ElemType: ConsistencyToken_SdkV2{}.Type(ctx),
			},
			"is_permitted": types.BoolType,
		},
	}
}

// GetConsistencyToken returns the value of the ConsistencyToken field in CheckPolicyResponse_SdkV2 as
// a ConsistencyToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CheckPolicyResponse_SdkV2) GetConsistencyToken(ctx context.Context) (ConsistencyToken_SdkV2, bool) {
	var e ConsistencyToken_SdkV2
	if o.ConsistencyToken.IsNull() || o.ConsistencyToken.IsUnknown() {
		return e, false
	}
	var v []ConsistencyToken_SdkV2
	d := o.ConsistencyToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConsistencyToken sets the value of the ConsistencyToken field in CheckPolicyResponse_SdkV2.
func (o *CheckPolicyResponse_SdkV2) SetConsistencyToken(ctx context.Context, v ConsistencyToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["consistency_token"]
	o.ConsistencyToken = types.ListValueMust(t, vs)
}

type ComplexValue_SdkV2 struct {
	Display types.String `tfsdk:"display"`

	Primary types.Bool `tfsdk:"primary"`

	Ref types.String `tfsdk:"$ref"`

	Type_ types.String `tfsdk:"type"`

	Value types.String `tfsdk:"value"`
}

func (toState *ComplexValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ComplexValue_SdkV2) {
}

func (toState *ComplexValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ComplexValue_SdkV2) {
}

func (c ComplexValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display"] = attrs["display"].SetOptional()
	attrs["primary"] = attrs["primary"].SetOptional()
	attrs["$ref"] = attrs["$ref"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplexValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplexValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplexValue_SdkV2
// only implements ToObjectValue() and Type().
func (o ComplexValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display": o.Display,
			"primary": o.Primary,
			"$ref":    o.Ref,
			"type":    o.Type_,
			"value":   o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplexValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display": types.StringType,
			"primary": types.BoolType,
			"$ref":    types.StringType,
			"type":    types.StringType,
			"value":   types.StringType,
		},
	}
}

type ConsistencyToken_SdkV2 struct {
	Value types.String `tfsdk:"value"`
}

func (toState *ConsistencyToken_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ConsistencyToken_SdkV2) {
}

func (toState *ConsistencyToken_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ConsistencyToken_SdkV2) {
}

func (c ConsistencyToken_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConsistencyToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConsistencyToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConsistencyToken_SdkV2
// only implements ToObjectValue() and Type().
func (o ConsistencyToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConsistencyToken_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type DeleteAccountGroupRequest_SdkV2 struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteAccountServicePrincipalRequest_SdkV2 struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteAccountUserRequest_SdkV2 struct {
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteGroupRequest_SdkV2 struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
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

type DeleteServicePrincipalRequest_SdkV2 struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteUserRequest_SdkV2 struct {
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWorkspaceAssignmentRequest_SdkV2 struct {
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspacePermissionAssignmentResponse_SdkV2 struct {
}

func (toState *DeleteWorkspacePermissionAssignmentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteWorkspacePermissionAssignmentResponse_SdkV2) {
}

func (toState *DeleteWorkspacePermissionAssignmentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteWorkspacePermissionAssignmentResponse_SdkV2) {
}

func (c DeleteWorkspacePermissionAssignmentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspacePermissionAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspacePermissionAssignmentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspacePermissionAssignmentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWorkspacePermissionAssignmentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetAccountGroupRequest_SdkV2 struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAccountServicePrincipalRequest_SdkV2 struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetAccountUserRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"id":                 o.Id,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type GetAssignableRolesForResourceRequest_SdkV2 struct {
	// The resource name for which assignable roles will be listed.
	//
	// Examples | Summary :--- | :--- `resource=accounts/<ACCOUNT_ID>` | A
	// resource name for the account.
	// `resource=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>` | A resource name for
	// the group. `resource=accounts/<ACCOUNT_ID>/servicePrincipals/<SP_ID>` | A
	// resource name for the service principal.
	Resource types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAssignableRolesForResourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAssignableRolesForResourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource": o.Resource,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource": types.StringType,
		},
	}
}

type GetAssignableRolesForResourceResponse_SdkV2 struct {
	Roles types.List `tfsdk:"roles"`
}

func (toState *GetAssignableRolesForResourceResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetAssignableRolesForResourceResponse_SdkV2) {
}

func (toState *GetAssignableRolesForResourceResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetAssignableRolesForResourceResponse_SdkV2) {
}

func (c GetAssignableRolesForResourceResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["roles"] = attrs["roles"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAssignableRolesForResourceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAssignableRolesForResourceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"roles": o.Roles,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"roles": basetypes.ListType{
				ElemType: Role_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRoles returns the value of the Roles field in GetAssignableRolesForResourceResponse_SdkV2 as
// a slice of Role_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetAssignableRolesForResourceResponse_SdkV2) GetRoles(ctx context.Context) ([]Role_SdkV2, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []Role_SdkV2
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in GetAssignableRolesForResourceResponse_SdkV2.
func (o *GetAssignableRolesForResourceResponse_SdkV2) SetRoles(ctx context.Context, v []Role_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

type GetGroupRequest_SdkV2 struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPasswordPermissionLevelsRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPasswordPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPasswordPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPasswordPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPasswordPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetPasswordPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (toState *GetPasswordPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPasswordPermissionLevelsResponse_SdkV2) {
}

func (toState *GetPasswordPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetPasswordPermissionLevelsResponse_SdkV2) {
}

func (c GetPasswordPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPasswordPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPasswordPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PasswordPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPasswordPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPasswordPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PasswordPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPasswordPermissionLevelsResponse_SdkV2 as
// a slice of PasswordPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPasswordPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]PasswordPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PasswordPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPasswordPermissionLevelsResponse_SdkV2.
func (o *GetPasswordPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []PasswordPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetPasswordPermissionsRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPasswordPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPasswordPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPasswordPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPasswordPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetPasswordPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetPermissionLevelsRequest_SdkV2 struct {
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

type GetPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (toState *GetPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPermissionLevelsResponse_SdkV2) {
}

func (toState *GetPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetPermissionLevelsResponse_SdkV2) {
}

func (c GetPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPermissionLevelsResponse_SdkV2 as
// a slice of PermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]PermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPermissionLevelsResponse_SdkV2.
func (o *GetPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []PermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetPermissionRequest_SdkV2 struct {
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPermissionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

type GetRuleSetRequest_SdkV2 struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	//
	// Examples | Summary :--- | :--- `etag=` | An empty etag can only be used
	// in GET to indicate no freshness requirements.
	// `etag=RENUAAABhSweA4NvVmmUYdiU717H3Tgy0UJdor3gE4a+mq/oj9NjAf8ZsQ==` | An
	// etag encoded a specific version of the rule set to get or to be updated.
	Etag types.String `tfsdk:"-"`
	// The ruleset name associated with the request.
	//
	// Examples | Summary :--- | :---
	// `name=accounts/<ACCOUNT_ID>/ruleSets/default` | A name for a rule set on
	// the account.
	// `name=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>/ruleSets/default` | A name
	// for a rule set on the group.
	// `name=accounts/<ACCOUNT_ID>/servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>/ruleSets/default`
	// | A name for a rule set on the service principal.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRuleSetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRuleSetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRuleSetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRuleSetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

type GetServicePrincipalRequest_SdkV2 struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetUserRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"id":                 o.Id,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentRequest_SdkV2 struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type GrantRule_SdkV2 struct {
	// Principals this grant rule applies to. A principal can be a user (for end
	// users), a service principal (for applications and compute workloads), or
	// an account group. Each principal has its own identifier format: *
	// users/<USERNAME> * groups/<GROUP_NAME> *
	// servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>
	Principals types.List `tfsdk:"principals"`
	// Role that is assigned to the list of principals.
	Role types.String `tfsdk:"role"`
}

func (toState *GrantRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GrantRule_SdkV2) {
}

func (toState *GrantRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GrantRule_SdkV2) {
}

func (c GrantRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principals"] = attrs["principals"].SetOptional()
	attrs["role"] = attrs["role"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GrantRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GrantRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"principals": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GrantRule_SdkV2
// only implements ToObjectValue() and Type().
func (o GrantRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principals": o.Principals,
			"role":       o.Role,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GrantRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principals": basetypes.ListType{
				ElemType: types.StringType,
			},
			"role": types.StringType,
		},
	}
}

// GetPrincipals returns the value of the Principals field in GrantRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GrantRule_SdkV2) GetPrincipals(ctx context.Context) ([]types.String, bool) {
	if o.Principals.IsNull() || o.Principals.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Principals.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrincipals sets the value of the Principals field in GrantRule_SdkV2.
func (o *GrantRule_SdkV2) SetPrincipals(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["principals"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Principals = types.ListValueMust(t, vs)
}

type Group_SdkV2 struct {
	// String that represents a human-readable group name
	DisplayName types.String `tfsdk:"displayName"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements"`
	// external_id should be unique for identifying groups
	ExternalId types.String `tfsdk:"externalId"`

	Groups types.List `tfsdk:"groups"`
	// Databricks group ID
	Id types.String `tfsdk:"id"`

	Members types.List `tfsdk:"members"`
	// Container for the group identifier. Workspace local versus account.
	Meta types.List `tfsdk:"meta"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles"`
	// The schema of the group.
	Schemas types.List `tfsdk:"schemas"`
}

func (toState *Group_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Group_SdkV2) {
	if !fromPlan.Meta.IsNull() && !fromPlan.Meta.IsUnknown() {
		if toStateMeta, ok := toState.GetMeta(ctx); ok {
			if fromPlanMeta, ok := fromPlan.GetMeta(ctx); ok {
				toStateMeta.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanMeta)
				toState.SetMeta(ctx, toStateMeta)
			}
		}
	}
}

func (toState *Group_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Group_SdkV2) {
	if !fromState.Meta.IsNull() && !fromState.Meta.IsUnknown() {
		if toStateMeta, ok := toState.GetMeta(ctx); ok {
			if fromStateMeta, ok := fromState.GetMeta(ctx); ok {
				toStateMeta.SyncFieldsDuringRead(ctx, fromStateMeta)
				toState.SetMeta(ctx, toStateMeta)
			}
		}
	}
}

func (c Group_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["displayName"] = attrs["displayName"].SetOptional()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["externalId"] = attrs["externalId"].SetOptional()
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["members"] = attrs["members"].SetOptional()
	attrs["meta"] = attrs["meta"].SetOptional()
	attrs["meta"] = attrs["meta"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["roles"] = attrs["roles"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Group.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Group_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue_SdkV2{}),
		"groups":       reflect.TypeOf(ComplexValue_SdkV2{}),
		"members":      reflect.TypeOf(ComplexValue_SdkV2{}),
		"meta":         reflect.TypeOf(ResourceMeta_SdkV2{}),
		"roles":        reflect.TypeOf(ComplexValue_SdkV2{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group_SdkV2
// only implements ToObjectValue() and Type().
func (o Group_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"displayName":  o.DisplayName,
			"entitlements": o.Entitlements,
			"externalId":   o.ExternalId,
			"groups":       o.Groups,
			"id":           o.Id,
			"members":      o.Members,
			"meta":         o.Meta,
			"roles":        o.Roles,
			"schemas":      o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Group_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"displayName": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"members": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"meta": basetypes.ListType{
				ElemType: ResourceMeta_SdkV2{}.Type(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEntitlements returns the value of the Entitlements field in Group_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetEntitlements(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in Group_SdkV2.
func (o *Group_SdkV2) SetEntitlements(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in Group_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetGroups(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in Group_SdkV2.
func (o *Group_SdkV2) SetGroups(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetMembers returns the value of the Members field in Group_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetMembers(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Members.IsNull() || o.Members.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Members.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMembers sets the value of the Members field in Group_SdkV2.
func (o *Group_SdkV2) SetMembers(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["members"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Members = types.ListValueMust(t, vs)
}

// GetMeta returns the value of the Meta field in Group_SdkV2 as
// a ResourceMeta_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetMeta(ctx context.Context) (ResourceMeta_SdkV2, bool) {
	var e ResourceMeta_SdkV2
	if o.Meta.IsNull() || o.Meta.IsUnknown() {
		return e, false
	}
	var v []ResourceMeta_SdkV2
	d := o.Meta.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMeta sets the value of the Meta field in Group_SdkV2.
func (o *Group_SdkV2) SetMeta(ctx context.Context, v ResourceMeta_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["meta"]
	o.Meta = types.ListValueMust(t, vs)
}

// GetRoles returns the value of the Roles field in Group_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetRoles(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in Group_SdkV2.
func (o *Group_SdkV2) SetRoles(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in Group_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in Group_SdkV2.
func (o *Group_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListAccountGroupsRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountGroupsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountGroupsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountGroupsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListAccountServicePrincipalsRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountServicePrincipalsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountServicePrincipalsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountServicePrincipalsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListAccountUsersRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountUsersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountUsersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountUsersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListGroupsRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListGroupsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListGroupsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListGroupsResponse_SdkV2 struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources"`
	// The schema of the service principal.
	Schemas types.List `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults"`
}

func (toState *ListGroupsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListGroupsResponse_SdkV2) {
}

func (toState *ListGroupsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListGroupsResponse_SdkV2) {
}

func (c ListGroupsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["itemsPerPage"] = attrs["itemsPerPage"].SetOptional()
	attrs["Resources"] = attrs["Resources"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()
	attrs["startIndex"] = attrs["startIndex"].SetOptional()
	attrs["totalResults"] = attrs["totalResults"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(Group_SdkV2{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListGroupsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"itemsPerPage": o.ItemsPerPage,
			"Resources":    o.Resources,
			"schemas":      o.Schemas,
			"startIndex":   o.StartIndex,
			"totalResults": o.TotalResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListGroupsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListGroupsResponse_SdkV2 as
// a slice of Group_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListGroupsResponse_SdkV2) GetResources(ctx context.Context) ([]Group_SdkV2, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []Group_SdkV2
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListGroupsResponse_SdkV2.
func (o *ListGroupsResponse_SdkV2) SetResources(ctx context.Context, v []Group_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListGroupsResponse_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListGroupsResponse_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListGroupsResponse_SdkV2.
func (o *ListGroupsResponse_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListServicePrincipalResponse_SdkV2 struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults"`
}

func (toState *ListServicePrincipalResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListServicePrincipalResponse_SdkV2) {
}

func (toState *ListServicePrincipalResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListServicePrincipalResponse_SdkV2) {
}

func (c ListServicePrincipalResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["itemsPerPage"] = attrs["itemsPerPage"].SetOptional()
	attrs["Resources"] = attrs["Resources"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()
	attrs["startIndex"] = attrs["startIndex"].SetOptional()
	attrs["totalResults"] = attrs["totalResults"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(ServicePrincipal_SdkV2{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"itemsPerPage": o.ItemsPerPage,
			"Resources":    o.Resources,
			"schemas":      o.Schemas,
			"startIndex":   o.StartIndex,
			"totalResults": o.TotalResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListServicePrincipalResponse_SdkV2 as
// a slice of ServicePrincipal_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListServicePrincipalResponse_SdkV2) GetResources(ctx context.Context) ([]ServicePrincipal_SdkV2, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []ServicePrincipal_SdkV2
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListServicePrincipalResponse_SdkV2.
func (o *ListServicePrincipalResponse_SdkV2) SetResources(ctx context.Context, v []ServicePrincipal_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListServicePrincipalResponse_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListServicePrincipalResponse_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListServicePrincipalResponse_SdkV2.
func (o *ListServicePrincipalResponse_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListServicePrincipalsRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListUsersRequest_SdkV2 struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListUsersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":         o.Attributes,
			"count":              o.Count,
			"excludedAttributes": o.ExcludedAttributes,
			"filter":             o.Filter,
			"sortBy":             o.SortBy,
			"sortOrder":          o.SortOrder,
			"startIndex":         o.StartIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListUsersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListUsersResponse_SdkV2 struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults"`
}

func (toState *ListUsersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListUsersResponse_SdkV2) {
}

func (toState *ListUsersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListUsersResponse_SdkV2) {
}

func (c ListUsersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["itemsPerPage"] = attrs["itemsPerPage"].SetOptional()
	attrs["Resources"] = attrs["Resources"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()
	attrs["startIndex"] = attrs["startIndex"].SetOptional()
	attrs["totalResults"] = attrs["totalResults"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(User_SdkV2{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListUsersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"itemsPerPage": o.ItemsPerPage,
			"Resources":    o.Resources,
			"schemas":      o.Schemas,
			"startIndex":   o.StartIndex,
			"totalResults": o.TotalResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListUsersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListUsersResponse_SdkV2 as
// a slice of User_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUsersResponse_SdkV2) GetResources(ctx context.Context) ([]User_SdkV2, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []User_SdkV2
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListUsersResponse_SdkV2.
func (o *ListUsersResponse_SdkV2) SetResources(ctx context.Context, v []User_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListUsersResponse_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUsersResponse_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListUsersResponse_SdkV2.
func (o *ListUsersResponse_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListWorkspaceAssignmentRequest_SdkV2 struct {
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type MeRequest_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o MeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type MigratePermissionsRequest_SdkV2 struct {
	// The name of the workspace group that permissions will be migrated from.
	FromWorkspaceGroupName types.String `tfsdk:"from_workspace_group_name"`
	// The maximum number of permissions that will be migrated.
	Size types.Int64 `tfsdk:"size"`
	// The name of the account group that permissions will be migrated to.
	ToAccountGroupName types.String `tfsdk:"to_account_group_name"`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigratePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o MigratePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"from_workspace_group_name": o.FromWorkspaceGroupName,
			"size":                      o.Size,
			"to_account_group_name":     o.ToAccountGroupName,
			"workspace_id":              o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"from_workspace_group_name": types.StringType,
			"size":                      types.Int64Type,
			"to_account_group_name":     types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

type MigratePermissionsResponse_SdkV2 struct {
	// Number of permissions migrated.
	PermissionsMigrated types.Int64 `tfsdk:"permissions_migrated"`
}

func (toState *MigratePermissionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan MigratePermissionsResponse_SdkV2) {
}

func (toState *MigratePermissionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState MigratePermissionsResponse_SdkV2) {
}

func (c MigratePermissionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permissions_migrated"] = attrs["permissions_migrated"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigratePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o MigratePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions_migrated": o.PermissionsMigrated,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions_migrated": types.Int64Type,
		},
	}
}

type Name_SdkV2 struct {
	// Family name of the Databricks user.
	FamilyName types.String `tfsdk:"familyName"`
	// Given name of the Databricks user.
	GivenName types.String `tfsdk:"givenName"`
}

func (toState *Name_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Name_SdkV2) {
}

func (toState *Name_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Name_SdkV2) {
}

func (c Name_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["familyName"] = attrs["familyName"].SetOptional()
	attrs["givenName"] = attrs["givenName"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Name.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Name_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Name_SdkV2
// only implements ToObjectValue() and Type().
func (o Name_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"familyName": o.FamilyName,
			"givenName":  o.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Name_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"familyName": types.StringType,
			"givenName":  types.StringType,
		},
	}
}

type ObjectPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *ObjectPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ObjectPermissions_SdkV2) {
}

func (toState *ObjectPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ObjectPermissions_SdkV2) {
}

func (c ObjectPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ObjectPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o ObjectPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ObjectPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ObjectPermissions_SdkV2 as
// a slice of AccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ObjectPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ObjectPermissions_SdkV2.
func (o *ObjectPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PartialUpdate_SdkV2 struct {
	// Unique ID in the Databricks workspace.
	Id types.String `tfsdk:"-"`

	Operations types.List `tfsdk:"Operations"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas types.List `tfsdk:"schemas"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartialUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartialUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operations": reflect.TypeOf(Patch_SdkV2{}),
		"schemas":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartialUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (o PartialUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"Operations": o.Operations,
			"schemas":    o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartialUpdate_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"Operations": basetypes.ListType{
				ElemType: Patch_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetOperations returns the value of the Operations field in PartialUpdate_SdkV2 as
// a slice of Patch_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartialUpdate_SdkV2) GetOperations(ctx context.Context) ([]Patch_SdkV2, bool) {
	if o.Operations.IsNull() || o.Operations.IsUnknown() {
		return nil, false
	}
	var v []Patch_SdkV2
	d := o.Operations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOperations sets the value of the Operations field in PartialUpdate_SdkV2.
func (o *PartialUpdate_SdkV2) SetOperations(ctx context.Context, v []Patch_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Operations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Operations = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in PartialUpdate_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartialUpdate_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in PartialUpdate_SdkV2.
func (o *PartialUpdate_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type PasswordAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *PasswordAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PasswordAccessControlRequest_SdkV2) {
}

func (toState *PasswordAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PasswordAccessControlRequest_SdkV2) {
}

func (c PasswordAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PasswordAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordAccessControlResponse_SdkV2 struct {
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

func (toState *PasswordAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PasswordAccessControlResponse_SdkV2) {
}

func (toState *PasswordAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PasswordAccessControlResponse_SdkV2) {
}

func (c PasswordAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PasswordPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PasswordAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PasswordPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in PasswordAccessControlResponse_SdkV2 as
// a slice of PasswordPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]PasswordPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PasswordPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PasswordAccessControlResponse_SdkV2.
func (o *PasswordAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []PasswordPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type PasswordPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PasswordPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PasswordPermission_SdkV2) {
}

func (toState *PasswordPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PasswordPermission_SdkV2) {
}

func (c PasswordPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in PasswordPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in PasswordPermission_SdkV2.
func (o *PasswordPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type PasswordPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *PasswordPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PasswordPermissions_SdkV2) {
}

func (toState *PasswordPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PasswordPermissions_SdkV2) {
}

func (c PasswordPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PasswordPermissions_SdkV2 as
// a slice of PasswordAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]PasswordAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PasswordAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PasswordPermissions_SdkV2.
func (o *PasswordPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []PasswordAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PasswordPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PasswordPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PasswordPermissionsDescription_SdkV2) {
}

func (toState *PasswordPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PasswordPermissionsDescription_SdkV2) {
}

func (c PasswordPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PasswordPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlRequest_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PasswordPermissionsRequest_SdkV2 as
// a slice of PasswordAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]PasswordAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PasswordAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PasswordPermissionsRequest_SdkV2.
func (o *PasswordPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []PasswordAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type Patch_SdkV2 struct {
	// Type of patch operation.
	Op types.String `tfsdk:"op"`
	// Selection of patch operation
	Path types.String `tfsdk:"path"`
	// Value to modify
	Value types.Object `tfsdk:"value"`
}

func (toState *Patch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Patch_SdkV2) {
}

func (toState *Patch_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Patch_SdkV2) {
}

func (c Patch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["op"] = attrs["op"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Patch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Patch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Patch_SdkV2
// only implements ToObjectValue() and Type().
func (o Patch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"op":    o.Op,
			"path":  o.Path,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Patch_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"op":    types.StringType,
			"path":  types.StringType,
			"value": types.ObjectType{},
		},
	}
}

type PatchResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PatchResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Permission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *Permission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Permission_SdkV2) {
}

func (toState *Permission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Permission_SdkV2) {
}

func (c Permission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Permission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Permission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Permission_SdkV2
// only implements ToObjectValue() and Type().
func (o Permission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Permission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in Permission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Permission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in Permission_SdkV2.
func (o *Permission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment_SdkV2 struct {
	// Error response associated with a workspace permission assignment, if any.
	Error types.String `tfsdk:"error"`
	// The permissions level of the principal.
	Permissions types.List `tfsdk:"permissions"`
	// Information about the principal assigned to the workspace.
	Principal types.List `tfsdk:"principal"`
}

func (toState *PermissionAssignment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PermissionAssignment_SdkV2) {
	if !fromPlan.Principal.IsNull() && !fromPlan.Principal.IsUnknown() {
		if toStatePrincipal, ok := toState.GetPrincipal(ctx); ok {
			if fromPlanPrincipal, ok := fromPlan.GetPrincipal(ctx); ok {
				toStatePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPrincipal)
				toState.SetPrincipal(ctx, toStatePrincipal)
			}
		}
	}
}

func (toState *PermissionAssignment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PermissionAssignment_SdkV2) {
	if !fromState.Principal.IsNull() && !fromState.Principal.IsUnknown() {
		if toStatePrincipal, ok := toState.GetPrincipal(ctx); ok {
			if fromStatePrincipal, ok := fromState.GetPrincipal(ctx); ok {
				toStatePrincipal.SyncFieldsDuringRead(ctx, fromStatePrincipal)
				toState.SetPrincipal(ctx, toStatePrincipal)
			}
		}
	}
}

func (c PermissionAssignment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["permissions"] = attrs["permissions"].SetOptional()
	attrs["principal"] = attrs["principal"].SetOptional()
	attrs["principal"] = attrs["principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
		"principal":   reflect.TypeOf(PrincipalOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":       o.Error,
			"permissions": o.Permissions,
			"principal":   o.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": basetypes.ListType{
				ElemType: PrincipalOutput_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissions returns the value of the Permissions field in PermissionAssignment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignment_SdkV2) GetPermissions(ctx context.Context) ([]types.String, bool) {
	if o.Permissions.IsNull() || o.Permissions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Permissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissions sets the value of the Permissions field in PermissionAssignment_SdkV2.
func (o *PermissionAssignment_SdkV2) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}

// GetPrincipal returns the value of the Principal field in PermissionAssignment_SdkV2 as
// a PrincipalOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignment_SdkV2) GetPrincipal(ctx context.Context) (PrincipalOutput_SdkV2, bool) {
	var e PrincipalOutput_SdkV2
	if o.Principal.IsNull() || o.Principal.IsUnknown() {
		return e, false
	}
	var v []PrincipalOutput_SdkV2
	d := o.Principal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrincipal sets the value of the Principal field in PermissionAssignment_SdkV2.
func (o *PermissionAssignment_SdkV2) SetPrincipal(ctx context.Context, v PrincipalOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["principal"]
	o.Principal = types.ListValueMust(t, vs)
}

type PermissionAssignments_SdkV2 struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments types.List `tfsdk:"permission_assignments"`
}

func (toState *PermissionAssignments_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PermissionAssignments_SdkV2) {
}

func (toState *PermissionAssignments_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PermissionAssignments_SdkV2) {
}

func (c PermissionAssignments_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_assignments"] = attrs["permission_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignments_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(PermissionAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionAssignments_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionAssignments_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_assignments": o.PermissionAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignments_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_assignments": basetypes.ListType{
				ElemType: PermissionAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionAssignments returns the value of the PermissionAssignments field in PermissionAssignments_SdkV2 as
// a slice of PermissionAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignments_SdkV2) GetPermissionAssignments(ctx context.Context) ([]PermissionAssignment_SdkV2, bool) {
	if o.PermissionAssignments.IsNull() || o.PermissionAssignments.IsUnknown() {
		return nil, false
	}
	var v []PermissionAssignment_SdkV2
	d := o.PermissionAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionAssignments sets the value of the PermissionAssignments field in PermissionAssignments_SdkV2.
func (o *PermissionAssignments_SdkV2) SetPermissionAssignments(ctx context.Context, v []PermissionAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionAssignments = types.ListValueMust(t, vs)
}

type PermissionOutput_SdkV2 struct {
	// The results of a permissions query.
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PermissionOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PermissionOutput_SdkV2) {
}

func (toState *PermissionOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PermissionOutput_SdkV2) {
}

func (c PermissionOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PermissionsDescription_SdkV2) {
}

func (toState *PermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PermissionsDescription_SdkV2) {
}

func (c PermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

// Information about the principal assigned to the workspace.
type PrincipalOutput_SdkV2 struct {
	// The display name of the principal.
	DisplayName types.String `tfsdk:"display_name"`
	// The group name of the group. Present only if the principal is a group.
	GroupName types.String `tfsdk:"group_name"`
	// The unique, opaque id of the principal.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The username of the user. Present only if the principal is a user.
	UserName types.String `tfsdk:"user_name"`
}

func (toState *PrincipalOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PrincipalOutput_SdkV2) {
}

func (toState *PrincipalOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PrincipalOutput_SdkV2) {
}

func (c PrincipalOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrincipalOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrincipalOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrincipalOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o PrincipalOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"principal_id":           o.PrincipalId,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrincipalOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"principal_id":           types.Int64Type,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ResourceMeta_SdkV2 struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType types.String `tfsdk:"resourceType"`
}

func (toState *ResourceMeta_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResourceMeta_SdkV2) {
}

func (toState *ResourceMeta_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ResourceMeta_SdkV2) {
}

func (c ResourceMeta_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resourceType"] = attrs["resourceType"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResourceMeta.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResourceMeta_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResourceMeta_SdkV2
// only implements ToObjectValue() and Type().
func (o ResourceMeta_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resourceType": o.ResourceType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResourceMeta_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resourceType": types.StringType,
		},
	}
}

type Role_SdkV2 struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name types.String `tfsdk:"name"`
}

func (toState *Role_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Role_SdkV2) {
}

func (toState *Role_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Role_SdkV2) {
}

func (c Role_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Role.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Role_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Role_SdkV2
// only implements ToObjectValue() and Type().
func (o Role_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Role_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RuleSetResponse_SdkV2 struct {
	// Identifies the version of the rule set returned. Etag used for
	// versioning. The response is at least as fresh as the eTag provided. Etag
	// is used for optimistic concurrency control as a way to help prevent
	// simultaneous updates of a rule set from overwriting each other. It is
	// strongly suggested that systems make use of the etag in the read ->
	// modify -> write pattern to perform rule set updates in order to avoid
	// race conditions that is get an etag from a GET rule set request, and pass
	// it with the PUT update request to identify the rule set version you are
	// updating.
	Etag types.String `tfsdk:"etag"`

	GrantRules types.List `tfsdk:"grant_rules"`
	// Name of the rule set.
	Name types.String `tfsdk:"name"`
}

func (toState *RuleSetResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RuleSetResponse_SdkV2) {
}

func (toState *RuleSetResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RuleSetResponse_SdkV2) {
}

func (c RuleSetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()
	attrs["grant_rules"] = attrs["grant_rules"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleSetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RuleSetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        o.Etag,
			"grant_rules": o.GrantRules,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in RuleSetResponse_SdkV2 as
// a slice of GrantRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RuleSetResponse_SdkV2) GetGrantRules(ctx context.Context) ([]GrantRule_SdkV2, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []GrantRule_SdkV2
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in RuleSetResponse_SdkV2.
func (o *RuleSetResponse_SdkV2) SetGrantRules(ctx context.Context, v []GrantRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type RuleSetUpdateRequest_SdkV2 struct {
	// Identifies the version of the rule set returned. Etag used for
	// versioning. The response is at least as fresh as the eTag provided. Etag
	// is used for optimistic concurrency control as a way to help prevent
	// simultaneous updates of a rule set from overwriting each other. It is
	// strongly suggested that systems make use of the etag in the read ->
	// modify -> write pattern to perform rule set updates in order to avoid
	// race conditions that is get an etag from a GET rule set request, and pass
	// it with the PUT update request to identify the rule set version you are
	// updating.
	Etag types.String `tfsdk:"etag"`

	GrantRules types.List `tfsdk:"grant_rules"`
	// Name of the rule set.
	Name types.String `tfsdk:"name"`
}

func (toState *RuleSetUpdateRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RuleSetUpdateRequest_SdkV2) {
}

func (toState *RuleSetUpdateRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RuleSetUpdateRequest_SdkV2) {
}

func (c RuleSetUpdateRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()
	attrs["grant_rules"] = attrs["grant_rules"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetUpdateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleSetUpdateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RuleSetUpdateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        o.Etag,
			"grant_rules": o.GrantRules,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in RuleSetUpdateRequest_SdkV2 as
// a slice of GrantRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RuleSetUpdateRequest_SdkV2) GetGrantRules(ctx context.Context) ([]GrantRule_SdkV2, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []GrantRule_SdkV2
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in RuleSetUpdateRequest_SdkV2.
func (o *RuleSetUpdateRequest_SdkV2) SetGrantRules(ctx context.Context, v []GrantRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type ServicePrincipal_SdkV2 struct {
	// If this user is active
	Active types.Bool `tfsdk:"active"`
	// UUID relating to the service principal
	ApplicationId types.String `tfsdk:"applicationId"`
	// String that represents a concatenation of given and family names.
	DisplayName types.String `tfsdk:"displayName"`
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements"`

	ExternalId types.String `tfsdk:"externalId"`

	Groups types.List `tfsdk:"groups"`
	// Databricks service principal ID.
	Id types.String `tfsdk:"id"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas"`
}

func (toState *ServicePrincipal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ServicePrincipal_SdkV2) {
}

func (toState *ServicePrincipal_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ServicePrincipal_SdkV2) {
}

func (c ServicePrincipal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active"] = attrs["active"].SetOptional()
	attrs["applicationId"] = attrs["applicationId"].SetOptional()
	attrs["displayName"] = attrs["displayName"].SetOptional()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["externalId"] = attrs["externalId"].SetOptional()
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["roles"] = attrs["roles"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServicePrincipal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue_SdkV2{}),
		"groups":       reflect.TypeOf(ComplexValue_SdkV2{}),
		"roles":        reflect.TypeOf(ComplexValue_SdkV2{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal_SdkV2
// only implements ToObjectValue() and Type().
func (o ServicePrincipal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active":        o.Active,
			"applicationId": o.ApplicationId,
			"displayName":   o.DisplayName,
			"entitlements":  o.Entitlements,
			"externalId":    o.ExternalId,
			"groups":        o.Groups,
			"id":            o.Id,
			"roles":         o.Roles,
			"schemas":       o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":        types.BoolType,
			"applicationId": types.StringType,
			"displayName":   types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"roles": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEntitlements returns the value of the Entitlements field in ServicePrincipal_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal_SdkV2) GetEntitlements(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in ServicePrincipal_SdkV2.
func (o *ServicePrincipal_SdkV2) SetEntitlements(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in ServicePrincipal_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal_SdkV2) GetGroups(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in ServicePrincipal_SdkV2.
func (o *ServicePrincipal_SdkV2) SetGroups(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetRoles returns the value of the Roles field in ServicePrincipal_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal_SdkV2) GetRoles(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in ServicePrincipal_SdkV2.
func (o *ServicePrincipal_SdkV2) SetRoles(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ServicePrincipal_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ServicePrincipal_SdkV2.
func (o *ServicePrincipal_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type SetObjectPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetObjectPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetObjectPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o SetObjectPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetObjectPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest_SdkV2{}.Type(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetObjectPermissions_SdkV2 as
// a slice of AccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetObjectPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetObjectPermissions_SdkV2.
func (o *SetObjectPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type UpdateObjectPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateObjectPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateObjectPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateObjectPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateObjectPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest_SdkV2{}.Type(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in UpdateObjectPermissions_SdkV2 as
// a slice of AccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateObjectPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]AccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in UpdateObjectPermissions_SdkV2.
func (o *UpdateObjectPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []AccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
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

type UpdateRuleSetRequest_SdkV2 struct {
	// Name of the rule set.
	Name types.String `tfsdk:"name"`

	RuleSet types.List `tfsdk:"rule_set"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRuleSetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_set": reflect.TypeOf(RuleSetUpdateRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRuleSetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRuleSetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"rule_set": o.RuleSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rule_set": basetypes.ListType{
				ElemType: RuleSetUpdateRequest_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRuleSet returns the value of the RuleSet field in UpdateRuleSetRequest_SdkV2 as
// a RuleSetUpdateRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRuleSetRequest_SdkV2) GetRuleSet(ctx context.Context) (RuleSetUpdateRequest_SdkV2, bool) {
	var e RuleSetUpdateRequest_SdkV2
	if o.RuleSet.IsNull() || o.RuleSet.IsUnknown() {
		return e, false
	}
	var v []RuleSetUpdateRequest_SdkV2
	d := o.RuleSet.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRuleSet sets the value of the RuleSet field in UpdateRuleSetRequest_SdkV2.
func (o *UpdateRuleSetRequest_SdkV2) SetRuleSet(ctx context.Context, v RuleSetUpdateRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rule_set"]
	o.RuleSet = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAssignments_SdkV2 struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	Permissions types.List `tfsdk:"permissions"`
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceAssignments_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignments_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceAssignments_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions":  o.Permissions,
			"principal_id": o.PrincipalId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

// GetPermissions returns the value of the Permissions field in UpdateWorkspaceAssignments_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceAssignments_SdkV2) GetPermissions(ctx context.Context) ([]types.String, bool) {
	if o.Permissions.IsNull() || o.Permissions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Permissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissions sets the value of the Permissions field in UpdateWorkspaceAssignments_SdkV2.
func (o *UpdateWorkspaceAssignments_SdkV2) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}

type User_SdkV2 struct {
	// If this user is active
	Active types.Bool `tfsdk:"active"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	DisplayName types.String `tfsdk:"displayName"`
	// All the emails associated with the Databricks user.
	Emails types.List `tfsdk:"emails"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId types.String `tfsdk:"externalId"`

	Groups types.List `tfsdk:"groups"`
	// Databricks user ID.
	Id types.String `tfsdk:"id"`

	Name types.List `tfsdk:"name"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles"`
	// The schema of the user.
	Schemas types.List `tfsdk:"schemas"`
	// Email address of the Databricks user.
	UserName types.String `tfsdk:"userName"`
}

func (toState *User_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan User_SdkV2) {
	if !fromPlan.Name.IsNull() && !fromPlan.Name.IsUnknown() {
		if toStateName, ok := toState.GetName(ctx); ok {
			if fromPlanName, ok := fromPlan.GetName(ctx); ok {
				toStateName.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanName)
				toState.SetName(ctx, toStateName)
			}
		}
	}
}

func (toState *User_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState User_SdkV2) {
	if !fromState.Name.IsNull() && !fromState.Name.IsUnknown() {
		if toStateName, ok := toState.GetName(ctx); ok {
			if fromStateName, ok := fromState.GetName(ctx); ok {
				toStateName.SyncFieldsDuringRead(ctx, fromStateName)
				toState.SetName(ctx, toStateName)
			}
		}
	}
}

func (c User_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active"] = attrs["active"].SetOptional()
	attrs["displayName"] = attrs["displayName"].SetOptional()
	attrs["emails"] = attrs["emails"].SetOptional()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["externalId"] = attrs["externalId"].SetOptional()
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["roles"] = attrs["roles"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()
	attrs["userName"] = attrs["userName"].SetOptional()

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
	return map[string]reflect.Type{
		"emails":       reflect.TypeOf(ComplexValue_SdkV2{}),
		"entitlements": reflect.TypeOf(ComplexValue_SdkV2{}),
		"groups":       reflect.TypeOf(ComplexValue_SdkV2{}),
		"name":         reflect.TypeOf(Name_SdkV2{}),
		"roles":        reflect.TypeOf(ComplexValue_SdkV2{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User_SdkV2
// only implements ToObjectValue() and Type().
func (o User_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active":       o.Active,
			"displayName":  o.DisplayName,
			"emails":       o.Emails,
			"entitlements": o.Entitlements,
			"externalId":   o.ExternalId,
			"groups":       o.Groups,
			"id":           o.Id,
			"name":         o.Name,
			"roles":        o.Roles,
			"schemas":      o.Schemas,
			"userName":     o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o User_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":      types.BoolType,
			"displayName": types.StringType,
			"emails": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"name": basetypes.ListType{
				ElemType: Name_SdkV2{}.Type(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue_SdkV2{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"userName": types.StringType,
		},
	}
}

// GetEmails returns the value of the Emails field in User_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetEmails(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Emails.IsNull() || o.Emails.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Emails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmails sets the value of the Emails field in User_SdkV2.
func (o *User_SdkV2) SetEmails(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["emails"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Emails = types.ListValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in User_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetEntitlements(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in User_SdkV2.
func (o *User_SdkV2) SetEntitlements(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in User_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetGroups(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in User_SdkV2.
func (o *User_SdkV2) SetGroups(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetName returns the value of the Name field in User_SdkV2 as
// a Name_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetName(ctx context.Context) (Name_SdkV2, bool) {
	var e Name_SdkV2
	if o.Name.IsNull() || o.Name.IsUnknown() {
		return e, false
	}
	var v []Name_SdkV2
	d := o.Name.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetName sets the value of the Name field in User_SdkV2.
func (o *User_SdkV2) SetName(ctx context.Context, v Name_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["name"]
	o.Name = types.ListValueMust(t, vs)
}

// GetRoles returns the value of the Roles field in User_SdkV2 as
// a slice of ComplexValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetRoles(ctx context.Context) ([]ComplexValue_SdkV2, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue_SdkV2
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in User_SdkV2.
func (o *User_SdkV2) SetRoles(ctx context.Context, v []ComplexValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in User_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetSchemas(ctx context.Context) ([]types.String, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in User_SdkV2.
func (o *User_SdkV2) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type WorkspacePermissions_SdkV2 struct {
	// Array of permissions defined for a workspace.
	Permissions types.List `tfsdk:"permissions"`
}

func (toState *WorkspacePermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WorkspacePermissions_SdkV2) {
}

func (toState *WorkspacePermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState WorkspacePermissions_SdkV2) {
}

func (c WorkspacePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permissions"] = attrs["permissions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspacePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspacePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(PermissionOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspacePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspacePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions": o.Permissions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspacePermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: PermissionOutput_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissions returns the value of the Permissions field in WorkspacePermissions_SdkV2 as
// a slice of PermissionOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspacePermissions_SdkV2) GetPermissions(ctx context.Context) ([]PermissionOutput_SdkV2, bool) {
	if o.Permissions.IsNull() || o.Permissions.IsUnknown() {
		return nil, false
	}
	var v []PermissionOutput_SdkV2
	d := o.Permissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissions sets the value of the Permissions field in WorkspacePermissions_SdkV2.
func (o *WorkspacePermissions_SdkV2) SetPermissions(ctx context.Context, v []PermissionOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}
