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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlRequest) {
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringRead(existingState AccessControlRequest) {
}

func (c AccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControlRequest
// only implements ToObjectValue() and Type().
func (o AccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AccessControlResponse struct {
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

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlResponse) {
}

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringRead(existingState AccessControlResponse) {
}

func (c AccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(Permission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessControlResponse
// only implements ToObjectValue() and Type().
func (o AccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: Permission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in AccessControlResponse as
// a slice of Permission values.
// If the field is unknown or null, the boolean return value is false.
func (o *AccessControlResponse) GetAllPermissions(ctx context.Context) ([]Permission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []Permission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in AccessControlResponse.
func (o *AccessControlResponse) SetAllPermissions(ctx context.Context, v []Permission) {
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
type Actor struct {
	ActorId types.Int64 `tfsdk:"actor_id"`
}

func (newState *Actor) SyncEffectiveFieldsDuringCreateOrUpdate(plan Actor) {
}

func (newState *Actor) SyncEffectiveFieldsDuringRead(existingState Actor) {
}

func (c Actor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Actor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Actor
// only implements ToObjectValue() and Type().
func (o Actor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id": o.ActorId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Actor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id": types.Int64Type,
		},
	}
}

type CheckPolicyResponse struct {
	ConsistencyToken types.Object `tfsdk:"consistency_token"`

	IsPermitted types.Bool `tfsdk:"is_permitted"`
}

func (newState *CheckPolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CheckPolicyResponse) {
}

func (newState *CheckPolicyResponse) SyncEffectiveFieldsDuringRead(existingState CheckPolicyResponse) {
}

func (c CheckPolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["consistency_token"] = attrs["consistency_token"].SetRequired()
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
func (a CheckPolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"consistency_token": reflect.TypeOf(ConsistencyToken{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CheckPolicyResponse
// only implements ToObjectValue() and Type().
func (o CheckPolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"consistency_token": o.ConsistencyToken,
			"is_permitted":      o.IsPermitted,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CheckPolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"consistency_token": ConsistencyToken{}.Type(ctx),
			"is_permitted":      types.BoolType,
		},
	}
}

// GetConsistencyToken returns the value of the ConsistencyToken field in CheckPolicyResponse as
// a ConsistencyToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *CheckPolicyResponse) GetConsistencyToken(ctx context.Context) (ConsistencyToken, bool) {
	var e ConsistencyToken
	if o.ConsistencyToken.IsNull() || o.ConsistencyToken.IsUnknown() {
		return e, false
	}
	var v []ConsistencyToken
	d := o.ConsistencyToken.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConsistencyToken sets the value of the ConsistencyToken field in CheckPolicyResponse.
func (o *CheckPolicyResponse) SetConsistencyToken(ctx context.Context, v ConsistencyToken) {
	vs := v.ToObjectValue(ctx)
	o.ConsistencyToken = vs
}

type ComplexValue struct {
	Display types.String `tfsdk:"display"`

	Primary types.Bool `tfsdk:"primary"`

	Ref types.String `tfsdk:"$ref"`

	Type_ types.String `tfsdk:"type"`

	Value types.String `tfsdk:"value"`
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplexValue) {
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringRead(existingState ComplexValue) {
}

func (c ComplexValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplexValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplexValue
// only implements ToObjectValue() and Type().
func (o ComplexValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ComplexValue) Type(ctx context.Context) attr.Type {
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

type ConsistencyToken struct {
	Value types.String `tfsdk:"value"`
}

func (newState *ConsistencyToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConsistencyToken) {
}

func (newState *ConsistencyToken) SyncEffectiveFieldsDuringRead(existingState ConsistencyToken) {
}

func (c ConsistencyToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ConsistencyToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConsistencyToken
// only implements ToObjectValue() and Type().
func (o ConsistencyToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConsistencyToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

// Delete a group.
type DeleteAccountGroupRequest struct {
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
func (a DeleteAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountGroupRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a service principal.
type DeleteAccountServicePrincipalRequest struct {
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
func (a DeleteAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user.
type DeleteAccountUserRequest struct {
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
func (a DeleteAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountUserRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a group.
type DeleteGroupRequest struct {
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
func (a DeleteGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupRequest
// only implements ToObjectValue() and Type().
func (o DeleteGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse struct {
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

// Delete a service principal.
type DeleteServicePrincipalRequest struct {
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
func (a DeleteServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (o DeleteServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user.
type DeleteUserRequest struct {
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
func (a DeleteUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserRequest
// only implements ToObjectValue() and Type().
func (o DeleteUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
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
func (a DeleteWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspacePermissionAssignmentResponse struct {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspacePermissionAssignmentResponse) {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspacePermissionAssignmentResponse) {
}

func (c DeleteWorkspacePermissionAssignmentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspacePermissionAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspacePermissionAssignmentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspacePermissionAssignmentResponse
// only implements ToObjectValue() and Type().
func (o DeleteWorkspacePermissionAssignmentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get group details.
type GetAccountGroupRequest struct {
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
func (a GetAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountGroupRequest
// only implements ToObjectValue() and Type().
func (o GetAccountGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get service principal details.
type GetAccountServicePrincipalRequest struct {
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
func (a GetAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (o GetAccountServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details.
type GetAccountUserRequest struct {
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
func (a GetAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountUserRequest
// only implements ToObjectValue() and Type().
func (o GetAccountUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetAccountUserRequest) Type(ctx context.Context) attr.Type {
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

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
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
func (a GetAssignableRolesForResourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAssignableRolesForResourceRequest
// only implements ToObjectValue() and Type().
func (o GetAssignableRolesForResourceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource": o.Resource,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource": types.StringType,
		},
	}
}

type GetAssignableRolesForResourceResponse struct {
	Roles types.List `tfsdk:"roles"`
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAssignableRolesForResourceResponse) {
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringRead(existingState GetAssignableRolesForResourceResponse) {
}

func (c GetAssignableRolesForResourceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetAssignableRolesForResourceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAssignableRolesForResourceResponse
// only implements ToObjectValue() and Type().
func (o GetAssignableRolesForResourceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"roles": o.Roles,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"roles": basetypes.ListType{
				ElemType: Role{}.Type(ctx),
			},
		},
	}
}

// GetRoles returns the value of the Roles field in GetAssignableRolesForResourceResponse as
// a slice of Role values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetAssignableRolesForResourceResponse) GetRoles(ctx context.Context) ([]Role, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []Role
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in GetAssignableRolesForResourceResponse.
func (o *GetAssignableRolesForResourceResponse) SetRoles(ctx context.Context, v []Role) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// Get group details.
type GetGroupRequest struct {
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
func (a GetGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupRequest
// only implements ToObjectValue() and Type().
func (o GetGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPasswordPermissionLevelsResponse) {
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPasswordPermissionLevelsResponse) {
}

func (c GetPasswordPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetPasswordPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PasswordPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPasswordPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetPasswordPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PasswordPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPasswordPermissionLevelsResponse as
// a slice of PasswordPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPasswordPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]PasswordPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PasswordPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPasswordPermissionLevelsResponse.
func (o *GetPasswordPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []PasswordPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get object permission levels
type GetPermissionLevelsRequest struct {
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionLevelsResponse) {
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPermissionLevelsResponse) {
}

func (c GetPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPermissionLevelsResponse as
// a slice of PermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]PermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPermissionLevelsResponse.
func (o *GetPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []PermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get object permissions
type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPermissionRequest
// only implements ToObjectValue() and Type().
func (o GetPermissionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// Get a rule set
type GetRuleSetRequest struct {
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
func (a GetRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRuleSetRequest
// only implements ToObjectValue() and Type().
func (o GetRuleSetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

// Get service principal details.
type GetServicePrincipalRequest struct {
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
func (a GetServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (o GetServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details.
type GetUserRequest struct {
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
func (a GetUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserRequest
// only implements ToObjectValue() and Type().
func (o GetUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetUserRequest) Type(ctx context.Context) attr.Type {
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

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
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
func (a GetWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type GrantRule struct {
	// Principals this grant rule applies to. A principal can be a user (for end
	// users), a service principal (for applications and compute workloads), or
	// an account group. Each principal has its own identifier format: *
	// users/<USERNAME> * groups/<GROUP_NAME> *
	// servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>
	Principals types.List `tfsdk:"principals"`
	// Role that is assigned to the list of principals.
	Role types.String `tfsdk:"role"`
}

func (newState *GrantRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan GrantRule) {
}

func (newState *GrantRule) SyncEffectiveFieldsDuringRead(existingState GrantRule) {
}

func (c GrantRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GrantRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"principals": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GrantRule
// only implements ToObjectValue() and Type().
func (o GrantRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principals": o.Principals,
			"role":       o.Role,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GrantRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principals": basetypes.ListType{
				ElemType: types.StringType,
			},
			"role": types.StringType,
		},
	}
}

// GetPrincipals returns the value of the Principals field in GrantRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GrantRule) GetPrincipals(ctx context.Context) ([]types.String, bool) {
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

// SetPrincipals sets the value of the Principals field in GrantRule.
func (o *GrantRule) SetPrincipals(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["principals"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Principals = types.ListValueMust(t, vs)
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName types.String `tfsdk:"displayName"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements"`

	ExternalId types.String `tfsdk:"externalId"`

	Groups types.List `tfsdk:"groups"`
	// Databricks group ID
	Id types.String `tfsdk:"id"`

	Members types.List `tfsdk:"members"`
	// Container for the group identifier. Workspace local versus account.
	Meta types.Object `tfsdk:"meta"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles"`
	// The schema of the group.
	Schemas types.List `tfsdk:"schemas"`
}

func (newState *Group) SyncEffectiveFieldsDuringCreateOrUpdate(plan Group) {
}

func (newState *Group) SyncEffectiveFieldsDuringRead(existingState Group) {
}

func (c Group) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["displayName"] = attrs["displayName"].SetOptional()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["externalId"] = attrs["externalId"].SetOptional()
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["members"] = attrs["members"].SetOptional()
	attrs["meta"] = attrs["meta"].SetOptional()
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
func (a Group) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"members":      reflect.TypeOf(ComplexValue{}),
		"meta":         reflect.TypeOf(ResourceMeta{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group
// only implements ToObjectValue() and Type().
func (o Group) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Group) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"displayName": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id": types.StringType,
			"members": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"meta": ResourceMeta{}.Type(ctx),
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEntitlements returns the value of the Entitlements field in Group as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetEntitlements(ctx context.Context) ([]ComplexValue, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in Group.
func (o *Group) SetEntitlements(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in Group as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetGroups(ctx context.Context) ([]ComplexValue, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in Group.
func (o *Group) SetGroups(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetMembers returns the value of the Members field in Group as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetMembers(ctx context.Context) ([]ComplexValue, bool) {
	if o.Members.IsNull() || o.Members.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Members.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMembers sets the value of the Members field in Group.
func (o *Group) SetMembers(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["members"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Members = types.ListValueMust(t, vs)
}

// GetMeta returns the value of the Meta field in Group as
// a ResourceMeta value.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetMeta(ctx context.Context) (ResourceMeta, bool) {
	var e ResourceMeta
	if o.Meta.IsNull() || o.Meta.IsUnknown() {
		return e, false
	}
	var v []ResourceMeta
	d := o.Meta.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMeta sets the value of the Meta field in Group.
func (o *Group) SetMeta(ctx context.Context, v ResourceMeta) {
	vs := v.ToObjectValue(ctx)
	o.Meta = vs
}

// GetRoles returns the value of the Roles field in Group as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetRoles(ctx context.Context) ([]ComplexValue, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in Group.
func (o *Group) SetRoles(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in Group as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Group) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in Group.
func (o *Group) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

// List group details.
type ListAccountGroupsRequest struct {
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
func (a ListAccountGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountGroupsRequest
// only implements ToObjectValue() and Type().
func (o ListAccountGroupsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListAccountGroupsRequest) Type(ctx context.Context) attr.Type {
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

// List service principals.
type ListAccountServicePrincipalsRequest struct {
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
func (a ListAccountServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountServicePrincipalsRequest
// only implements ToObjectValue() and Type().
func (o ListAccountServicePrincipalsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListAccountServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
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

// List users.
type ListAccountUsersRequest struct {
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
func (a ListAccountUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountUsersRequest
// only implements ToObjectValue() and Type().
func (o ListAccountUsersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListAccountUsersRequest) Type(ctx context.Context) attr.Type {
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

// List group details.
type ListGroupsRequest struct {
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
func (a ListGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsRequest
// only implements ToObjectValue() and Type().
func (o ListGroupsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListGroupsRequest) Type(ctx context.Context) attr.Type {
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

type ListGroupsResponse struct {
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

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGroupsResponse) {
}

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringRead(existingState ListGroupsResponse) {
}

func (c ListGroupsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(Group{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsResponse
// only implements ToObjectValue() and Type().
func (o ListGroupsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListGroupsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: Group{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListGroupsResponse as
// a slice of Group values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListGroupsResponse) GetResources(ctx context.Context) ([]Group, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []Group
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListGroupsResponse.
func (o *ListGroupsResponse) SetResources(ctx context.Context, v []Group) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListGroupsResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListGroupsResponse) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in ListGroupsResponse.
func (o *ListGroupsResponse) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListServicePrincipalResponse struct {
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

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalResponse) {
}

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalResponse) {
}

func (c ListServicePrincipalResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(ServicePrincipal{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalResponse
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListServicePrincipalResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: ServicePrincipal{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListServicePrincipalResponse as
// a slice of ServicePrincipal values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListServicePrincipalResponse) GetResources(ctx context.Context) ([]ServicePrincipal, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []ServicePrincipal
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListServicePrincipalResponse.
func (o *ListServicePrincipalResponse) SetResources(ctx context.Context, v []ServicePrincipal) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListServicePrincipalResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListServicePrincipalResponse) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in ListServicePrincipalResponse.
func (o *ListServicePrincipalResponse) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

// List service principals.
type ListServicePrincipalsRequest struct {
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
func (a ListServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsRequest
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
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

// List users.
type ListUsersRequest struct {
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
func (a ListUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersRequest
// only implements ToObjectValue() and Type().
func (o ListUsersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListUsersRequest) Type(ctx context.Context) attr.Type {
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

type ListUsersResponse struct {
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

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUsersResponse) {
}

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringRead(existingState ListUsersResponse) {
}

func (c ListUsersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListUsersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(User{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersResponse
// only implements ToObjectValue() and Type().
func (o ListUsersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListUsersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// GetResources returns the value of the Resources field in ListUsersResponse as
// a slice of User values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUsersResponse) GetResources(ctx context.Context) ([]User, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []User
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListUsersResponse.
func (o *ListUsersResponse) SetResources(ctx context.Context, v []User) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ListUsersResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUsersResponse) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in ListUsersResponse.
func (o *ListUsersResponse) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
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
func (a ListWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type MigratePermissionsRequest struct {
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

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsRequest) {
}

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsRequest) {
}

func (c MigratePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["from_workspace_group_name"] = attrs["from_workspace_group_name"].SetRequired()
	attrs["size"] = attrs["size"].SetOptional()
	attrs["to_account_group_name"] = attrs["to_account_group_name"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigratePermissionsRequest
// only implements ToObjectValue() and Type().
func (o MigratePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MigratePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"from_workspace_group_name": types.StringType,
			"size":                      types.Int64Type,
			"to_account_group_name":     types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	PermissionsMigrated types.Int64 `tfsdk:"permissions_migrated"`
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsResponse) {
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsResponse) {
}

func (c MigratePermissionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MigratePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigratePermissionsResponse
// only implements ToObjectValue() and Type().
func (o MigratePermissionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions_migrated": o.PermissionsMigrated,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions_migrated": types.Int64Type,
		},
	}
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName types.String `tfsdk:"familyName"`
	// Given name of the Databricks user.
	GivenName types.String `tfsdk:"givenName"`
}

func (newState *Name) SyncEffectiveFieldsDuringCreateOrUpdate(plan Name) {
}

func (newState *Name) SyncEffectiveFieldsDuringRead(existingState Name) {
}

func (c Name) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Name) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Name
// only implements ToObjectValue() and Type().
func (o Name) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"familyName": o.FamilyName,
			"givenName":  o.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Name) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"familyName": types.StringType,
			"givenName":  types.StringType,
		},
	}
}

type ObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ObjectPermissions) {
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringRead(existingState ObjectPermissions) {
}

func (c ObjectPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectPermissions
// only implements ToObjectValue() and Type().
func (o ObjectPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ObjectPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ObjectPermissions as
// a slice of AccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *ObjectPermissions) GetAccessControlList(ctx context.Context) ([]AccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ObjectPermissions.
func (o *ObjectPermissions) SetAccessControlList(ctx context.Context, v []AccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PartialUpdate struct {
	// Unique ID in the Databricks workspace.
	Id types.String `tfsdk:"-"`

	Operations types.List `tfsdk:"Operations"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas types.List `tfsdk:"schemas"`
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartialUpdate) {
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringRead(existingState PartialUpdate) {
}

func (c PartialUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["Operations"] = attrs["Operations"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartialUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartialUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operations": reflect.TypeOf(Patch{}),
		"schemas":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartialUpdate
// only implements ToObjectValue() and Type().
func (o PartialUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         o.Id,
			"Operations": o.Operations,
			"schemas":    o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartialUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"Operations": basetypes.ListType{
				ElemType: Patch{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetOperations returns the value of the Operations field in PartialUpdate as
// a slice of Patch values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartialUpdate) GetOperations(ctx context.Context) ([]Patch, bool) {
	if o.Operations.IsNull() || o.Operations.IsUnknown() {
		return nil, false
	}
	var v []Patch
	d := o.Operations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOperations sets the value of the Operations field in PartialUpdate.
func (o *PartialUpdate) SetOperations(ctx context.Context, v []Patch) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["Operations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Operations = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in PartialUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartialUpdate) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in PartialUpdate.
func (o *PartialUpdate) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlRequest) {
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlRequest) {
}

func (c PasswordAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PasswordAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordAccessControlRequest
// only implements ToObjectValue() and Type().
func (o PasswordAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PasswordAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordAccessControlResponse struct {
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

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlResponse) {
}

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlResponse) {
}

func (c PasswordAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PasswordAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PasswordPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordAccessControlResponse
// only implements ToObjectValue() and Type().
func (o PasswordAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PasswordAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PasswordPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in PasswordAccessControlResponse as
// a slice of PasswordPermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordAccessControlResponse) GetAllPermissions(ctx context.Context) ([]PasswordPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PasswordPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PasswordAccessControlResponse.
func (o *PasswordAccessControlResponse) SetAllPermissions(ctx context.Context, v []PasswordPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type PasswordPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermission) {
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringRead(existingState PasswordPermission) {
}

func (c PasswordPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PasswordPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermission
// only implements ToObjectValue() and Type().
func (o PasswordPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in PasswordPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in PasswordPermission.
func (o *PasswordPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type PasswordPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissions) {
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringRead(existingState PasswordPermissions) {
}

func (c PasswordPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PasswordPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissions
// only implements ToObjectValue() and Type().
func (o PasswordPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PasswordPermissions as
// a slice of PasswordAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermissions) GetAccessControlList(ctx context.Context) ([]PasswordAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PasswordAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PasswordPermissions.
func (o *PasswordPermissions) SetAccessControlList(ctx context.Context, v []PasswordAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PasswordPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsDescription) {
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsDescription) {
}

func (c PasswordPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PasswordPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissionsDescription
// only implements ToObjectValue() and Type().
func (o PasswordPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsRequest) {
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsRequest) {
}

func (c PasswordPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PasswordPermissionsRequest
// only implements ToObjectValue() and Type().
func (o PasswordPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlRequest{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PasswordPermissionsRequest as
// a slice of PasswordAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *PasswordPermissionsRequest) GetAccessControlList(ctx context.Context) ([]PasswordAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PasswordAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PasswordPermissionsRequest.
func (o *PasswordPermissionsRequest) SetAccessControlList(ctx context.Context, v []PasswordAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type Patch struct {
	// Type of patch operation.
	Op types.String `tfsdk:"op"`
	// Selection of patch operation
	Path types.String `tfsdk:"path"`
	// Value to modify
	Value types.Object `tfsdk:"value"`
}

func (newState *Patch) SyncEffectiveFieldsDuringCreateOrUpdate(plan Patch) {
}

func (newState *Patch) SyncEffectiveFieldsDuringRead(existingState Patch) {
}

func (c Patch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Patch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Patch
// only implements ToObjectValue() and Type().
func (o Patch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"op":    o.Op,
			"path":  o.Path,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Patch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"op":    types.StringType,
			"path":  types.StringType,
			"value": types.ObjectType{},
		},
	}
}

type PatchResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchResponse
// only implements ToObjectValue() and Type().
func (o PatchResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PatchResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Permission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *Permission) SyncEffectiveFieldsDuringCreateOrUpdate(plan Permission) {
}

func (newState *Permission) SyncEffectiveFieldsDuringRead(existingState Permission) {
}

func (c Permission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Permission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Permission
// only implements ToObjectValue() and Type().
func (o Permission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Permission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in Permission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Permission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in Permission.
func (o *Permission) SetInheritedFromObject(ctx context.Context, v []types.String) {
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
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error types.String `tfsdk:"error"`
	// The permissions level of the principal.
	Permissions types.List `tfsdk:"permissions"`
	// Information about the principal assigned to the workspace.
	Principal types.Object `tfsdk:"principal"`
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignment) {
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringRead(existingState PermissionAssignment) {
}

func (c PermissionAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["permissions"] = attrs["permissions"].SetOptional()
	attrs["principal"] = attrs["principal"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
		"principal":   reflect.TypeOf(PrincipalOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionAssignment
// only implements ToObjectValue() and Type().
func (o PermissionAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":       o.Error,
			"permissions": o.Permissions,
			"principal":   o.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": PrincipalOutput{}.Type(ctx),
		},
	}
}

// GetPermissions returns the value of the Permissions field in PermissionAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignment) GetPermissions(ctx context.Context) ([]types.String, bool) {
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

// SetPermissions sets the value of the Permissions field in PermissionAssignment.
func (o *PermissionAssignment) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}

// GetPrincipal returns the value of the Principal field in PermissionAssignment as
// a PrincipalOutput value.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignment) GetPrincipal(ctx context.Context) (PrincipalOutput, bool) {
	var e PrincipalOutput
	if o.Principal.IsNull() || o.Principal.IsUnknown() {
		return e, false
	}
	var v []PrincipalOutput
	d := o.Principal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrincipal sets the value of the Principal field in PermissionAssignment.
func (o *PermissionAssignment) SetPrincipal(ctx context.Context, v PrincipalOutput) {
	vs := v.ToObjectValue(ctx)
	o.Principal = vs
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments types.List `tfsdk:"permission_assignments"`
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignments) {
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringRead(existingState PermissionAssignments) {
}

func (c PermissionAssignments) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PermissionAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(PermissionAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionAssignments
// only implements ToObjectValue() and Type().
func (o PermissionAssignments) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_assignments": o.PermissionAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignments) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_assignments": basetypes.ListType{
				ElemType: PermissionAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPermissionAssignments returns the value of the PermissionAssignments field in PermissionAssignments as
// a slice of PermissionAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionAssignments) GetPermissionAssignments(ctx context.Context) ([]PermissionAssignment, bool) {
	if o.PermissionAssignments.IsNull() || o.PermissionAssignments.IsUnknown() {
		return nil, false
	}
	var v []PermissionAssignment
	d := o.PermissionAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionAssignments sets the value of the PermissionAssignments field in PermissionAssignments.
func (o *PermissionAssignments) SetPermissionAssignments(ctx context.Context, v []PermissionAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionAssignments = types.ListValueMust(t, vs)
}

type PermissionOutput struct {
	// The results of a permissions query.
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionOutput) {
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringRead(existingState PermissionOutput) {
}

func (c PermissionOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PermissionOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionOutput
// only implements ToObjectValue() and Type().
func (o PermissionOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsDescription) {
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PermissionsDescription) {
}

func (c PermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsDescription
// only implements ToObjectValue() and Type().
func (o PermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
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

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrincipalOutput) {
}

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringRead(existingState PrincipalOutput) {
}

func (c PrincipalOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PrincipalOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrincipalOutput
// only implements ToObjectValue() and Type().
func (o PrincipalOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PrincipalOutput) Type(ctx context.Context) attr.Type {
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

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType types.String `tfsdk:"resourceType"`
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResourceMeta) {
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringRead(existingState ResourceMeta) {
}

func (c ResourceMeta) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ResourceMeta) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResourceMeta
// only implements ToObjectValue() and Type().
func (o ResourceMeta) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resourceType": o.ResourceType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResourceMeta) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resourceType": types.StringType,
		},
	}
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name types.String `tfsdk:"name"`
}

func (newState *Role) SyncEffectiveFieldsDuringCreateOrUpdate(plan Role) {
}

func (newState *Role) SyncEffectiveFieldsDuringRead(existingState Role) {
}

func (c Role) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Role) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Role
// only implements ToObjectValue() and Type().
func (o Role) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Role) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RuleSetResponse struct {
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

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetResponse) {
}

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringRead(existingState RuleSetResponse) {
}

func (c RuleSetResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RuleSetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleSetResponse
// only implements ToObjectValue() and Type().
func (o RuleSetResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        o.Etag,
			"grant_rules": o.GrantRules,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in RuleSetResponse as
// a slice of GrantRule values.
// If the field is unknown or null, the boolean return value is false.
func (o *RuleSetResponse) GetGrantRules(ctx context.Context) ([]GrantRule, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []GrantRule
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in RuleSetResponse.
func (o *RuleSetResponse) SetGrantRules(ctx context.Context, v []GrantRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type RuleSetUpdateRequest struct {
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

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetUpdateRequest) {
}

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringRead(existingState RuleSetUpdateRequest) {
}

func (c RuleSetUpdateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RuleSetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleSetUpdateRequest
// only implements ToObjectValue() and Type().
func (o RuleSetUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        o.Etag,
			"grant_rules": o.GrantRules,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetGrantRules returns the value of the GrantRules field in RuleSetUpdateRequest as
// a slice of GrantRule values.
// If the field is unknown or null, the boolean return value is false.
func (o *RuleSetUpdateRequest) GetGrantRules(ctx context.Context) ([]GrantRule, bool) {
	if o.GrantRules.IsNull() || o.GrantRules.IsUnknown() {
		return nil, false
	}
	var v []GrantRule
	d := o.GrantRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGrantRules sets the value of the GrantRules field in RuleSetUpdateRequest.
func (o *RuleSetUpdateRequest) SetGrantRules(ctx context.Context, v []GrantRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["grant_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.GrantRules = types.ListValueMust(t, vs)
}

type ServicePrincipal struct {
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

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServicePrincipal) {
}

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringRead(existingState ServicePrincipal) {
}

func (c ServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal
// only implements ToObjectValue() and Type().
func (o ServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":        types.BoolType,
			"applicationId": types.StringType,
			"displayName":   types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id": types.StringType,
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEntitlements returns the value of the Entitlements field in ServicePrincipal as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal) GetEntitlements(ctx context.Context) ([]ComplexValue, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in ServicePrincipal.
func (o *ServicePrincipal) SetEntitlements(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in ServicePrincipal as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal) GetGroups(ctx context.Context) ([]ComplexValue, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in ServicePrincipal.
func (o *ServicePrincipal) SetGroups(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetRoles returns the value of the Roles field in ServicePrincipal as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal) GetRoles(ctx context.Context) ([]ComplexValue, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in ServicePrincipal.
func (o *ServicePrincipal) SetRoles(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in ServicePrincipal as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServicePrincipal) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in ServicePrincipal.
func (o *ServicePrincipal) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type SetObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *SetObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetObjectPermissions) {
}

func (newState *SetObjectPermissions) SyncEffectiveFieldsDuringRead(existingState SetObjectPermissions) {
}

func (c SetObjectPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["request_object_id"] = attrs["request_object_id"].SetRequired()
	attrs["request_object_type"] = attrs["request_object_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetObjectPermissions
// only implements ToObjectValue() and Type().
func (o SetObjectPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetObjectPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest{}.Type(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SetObjectPermissions as
// a slice of AccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetObjectPermissions) GetAccessControlList(ctx context.Context) ([]AccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SetObjectPermissions.
func (o *SetObjectPermissions) SetAccessControlList(ctx context.Context, v []AccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type UpdateObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *UpdateObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateObjectPermissions) {
}

func (newState *UpdateObjectPermissions) SyncEffectiveFieldsDuringRead(existingState UpdateObjectPermissions) {
}

func (c UpdateObjectPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["request_object_id"] = attrs["request_object_id"].SetRequired()
	attrs["request_object_type"] = attrs["request_object_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateObjectPermissions
// only implements ToObjectValue() and Type().
func (o UpdateObjectPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"request_object_id":   o.RequestObjectId,
			"request_object_type": o.RequestObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateObjectPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest{}.Type(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in UpdateObjectPermissions as
// a slice of AccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateObjectPermissions) GetAccessControlList(ctx context.Context) ([]AccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in UpdateObjectPermissions.
func (o *UpdateObjectPermissions) SetAccessControlList(ctx context.Context, v []AccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type UpdateResponse struct {
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

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name types.String `tfsdk:"name"`

	RuleSet types.Object `tfsdk:"rule_set"`
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRuleSetRequest) {
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRuleSetRequest) {
}

func (c UpdateRuleSetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["rule_set"] = attrs["rule_set"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_set": reflect.TypeOf(RuleSetUpdateRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRuleSetRequest
// only implements ToObjectValue() and Type().
func (o UpdateRuleSetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"rule_set": o.RuleSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":     types.StringType,
			"rule_set": RuleSetUpdateRequest{}.Type(ctx),
		},
	}
}

// GetRuleSet returns the value of the RuleSet field in UpdateRuleSetRequest as
// a RuleSetUpdateRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRuleSetRequest) GetRuleSet(ctx context.Context) (RuleSetUpdateRequest, bool) {
	var e RuleSetUpdateRequest
	if o.RuleSet.IsNull() || o.RuleSet.IsUnknown() {
		return e, false
	}
	var v []RuleSetUpdateRequest
	d := o.RuleSet.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRuleSet sets the value of the RuleSet field in UpdateRuleSetRequest.
func (o *UpdateRuleSetRequest) SetRuleSet(ctx context.Context, v RuleSetUpdateRequest) {
	vs := v.ToObjectValue(ctx)
	o.RuleSet = vs
}

type UpdateWorkspaceAssignments struct {
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

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceAssignments) {
}

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceAssignments) {
}

func (c UpdateWorkspaceAssignments) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permissions"] = attrs["permissions"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignments
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceAssignments) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions":  o.Permissions,
			"principal_id": o.PrincipalId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) Type(ctx context.Context) attr.Type {
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

// GetPermissions returns the value of the Permissions field in UpdateWorkspaceAssignments as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceAssignments) GetPermissions(ctx context.Context) ([]types.String, bool) {
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

// SetPermissions sets the value of the Permissions field in UpdateWorkspaceAssignments.
func (o *UpdateWorkspaceAssignments) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}

type User struct {
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

	Name types.Object `tfsdk:"name"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles"`
	// The schema of the user.
	Schemas types.List `tfsdk:"schemas"`
	// Email address of the Databricks user.
	UserName types.String `tfsdk:"userName"`
}

func (newState *User) SyncEffectiveFieldsDuringCreateOrUpdate(plan User) {
}

func (newState *User) SyncEffectiveFieldsDuringRead(existingState User) {
}

func (c User) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active"] = attrs["active"].SetOptional()
	attrs["displayName"] = attrs["displayName"].SetOptional()
	attrs["emails"] = attrs["emails"].SetOptional()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["externalId"] = attrs["externalId"].SetOptional()
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
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
func (a User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"emails":       reflect.TypeOf(ComplexValue{}),
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"name":         reflect.TypeOf(Name{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User
// only implements ToObjectValue() and Type().
func (o User) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o User) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":      types.BoolType,
			"displayName": types.StringType,
			"emails": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id":   types.StringType,
			"name": Name{}.Type(ctx),
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"userName": types.StringType,
		},
	}
}

// GetEmails returns the value of the Emails field in User as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetEmails(ctx context.Context) ([]ComplexValue, bool) {
	if o.Emails.IsNull() || o.Emails.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Emails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmails sets the value of the Emails field in User.
func (o *User) SetEmails(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["emails"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Emails = types.ListValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in User as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetEntitlements(ctx context.Context) ([]ComplexValue, bool) {
	if o.Entitlements.IsNull() || o.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in User.
func (o *User) SetEntitlements(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Entitlements = types.ListValueMust(t, vs)
}

// GetGroups returns the value of the Groups field in User as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetGroups(ctx context.Context) ([]ComplexValue, bool) {
	if o.Groups.IsNull() || o.Groups.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in User.
func (o *User) SetGroups(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Groups = types.ListValueMust(t, vs)
}

// GetName returns the value of the Name field in User as
// a Name value.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetName(ctx context.Context) (Name, bool) {
	var e Name
	if o.Name.IsNull() || o.Name.IsUnknown() {
		return e, false
	}
	var v []Name
	d := o.Name.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetName sets the value of the Name field in User.
func (o *User) SetName(ctx context.Context, v Name) {
	vs := v.ToObjectValue(ctx)
	o.Name = vs
}

// GetRoles returns the value of the Roles field in User as
// a slice of ComplexValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetRoles(ctx context.Context) ([]ComplexValue, bool) {
	if o.Roles.IsNull() || o.Roles.IsUnknown() {
		return nil, false
	}
	var v []ComplexValue
	d := o.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in User.
func (o *User) SetRoles(ctx context.Context, v []ComplexValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Roles = types.ListValueMust(t, vs)
}

// GetSchemas returns the value of the Schemas field in User as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetSchemas(ctx context.Context) ([]types.String, bool) {
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

// SetSchemas sets the value of the Schemas field in User.
func (o *User) SetSchemas(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions types.List `tfsdk:"permissions"`
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspacePermissions) {
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringRead(existingState WorkspacePermissions) {
}

func (c WorkspacePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspacePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(PermissionOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspacePermissions
// only implements ToObjectValue() and Type().
func (o WorkspacePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permissions": o.Permissions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspacePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: PermissionOutput{}.Type(ctx),
			},
		},
	}
}

// GetPermissions returns the value of the Permissions field in WorkspacePermissions as
// a slice of PermissionOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspacePermissions) GetPermissions(ctx context.Context) ([]PermissionOutput, bool) {
	if o.Permissions.IsNull() || o.Permissions.IsUnknown() {
		return nil, false
	}
	var v []PermissionOutput
	d := o.Permissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissions sets the value of the Permissions field in WorkspacePermissions.
func (o *WorkspacePermissions) SetPermissions(ctx context.Context, v []PermissionOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}
