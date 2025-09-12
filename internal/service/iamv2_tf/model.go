// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package iamv2_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GetWorkspaceAccessDetailLocalRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceAccessDetailLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailLocalRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAccessDetailLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"view":         o.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAccessDetailLocalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"view":         types.StringType,
		},
	}
}

type GetWorkspaceAccessDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
	// Required. The workspace ID for which the access details are being
	// requested.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAccessDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceAccessDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAccessDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"view":         o.View,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAccessDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"view":         types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// The details of a Group resource.
type Group struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Display name of the group.
	GroupName types.String `tfsdk:"group_name"`
	// Internal group ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`
}

func (toState *Group) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Group) {
}

func (toState *Group) SyncFieldsDuringRead(ctx context.Context, fromState Group) {
}

func (c Group) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()

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
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group
// only implements ToObjectValue() and Type().
func (o Group) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":  o.AccountId,
			"external_id": o.ExternalId,
			"group_name":  o.GroupName,
			"internal_id": o.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Group) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_name":  types.StringType,
			"internal_id": types.Int64Type,
		},
	}
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// groupname, and inherited parent groups.
type ResolveGroupProxyRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupProxyRequest
// only implements ToObjectValue() and Type().
func (o ResolveGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// groupname, and inherited parent groups.
type ResolveGroupRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupRequest
// only implements ToObjectValue() and Type().
func (o ResolveGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveGroupResponse struct {
	// The group that was resolved.
	Group types.Object `tfsdk:"group"`
}

func (toState *ResolveGroupResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveGroupResponse) {
	if !fromPlan.Group.IsNull() && !fromPlan.Group.IsUnknown() {
		if toStateGroup, ok := toState.GetGroup(ctx); ok {
			if fromPlanGroup, ok := fromPlan.GetGroup(ctx); ok {
				toStateGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGroup)
				toState.SetGroup(ctx, toStateGroup)
			}
		}
	}
}

func (toState *ResolveGroupResponse) SyncFieldsDuringRead(ctx context.Context, fromState ResolveGroupResponse) {
	if !fromState.Group.IsNull() && !fromState.Group.IsUnknown() {
		if toStateGroup, ok := toState.GetGroup(ctx); ok {
			if fromStateGroup, ok := fromState.GetGroup(ctx); ok {
				toStateGroup.SyncFieldsDuringRead(ctx, fromStateGroup)
				toState.SetGroup(ctx, toStateGroup)
			}
		}
	}
}

func (c ResolveGroupResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveGroupResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupResponse
// only implements ToObjectValue() and Type().
func (o ResolveGroupResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": o.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": Group{}.Type(ctx),
		},
	}
}

// GetGroup returns the value of the Group field in ResolveGroupResponse as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveGroupResponse) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if o.Group.IsNull() || o.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := o.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in ResolveGroupResponse.
func (o *ResolveGroupResponse) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	o.Group = vs
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveServicePrincipalResponse struct {
	// The service principal that was resolved.
	ServicePrincipal types.Object `tfsdk:"service_principal"`
}

func (toState *ResolveServicePrincipalResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveServicePrincipalResponse) {
	if !fromPlan.ServicePrincipal.IsNull() && !fromPlan.ServicePrincipal.IsUnknown() {
		if toStateServicePrincipal, ok := toState.GetServicePrincipal(ctx); ok {
			if fromPlanServicePrincipal, ok := fromPlan.GetServicePrincipal(ctx); ok {
				toStateServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanServicePrincipal)
				toState.SetServicePrincipal(ctx, toStateServicePrincipal)
			}
		}
	}
}

func (toState *ResolveServicePrincipalResponse) SyncFieldsDuringRead(ctx context.Context, fromState ResolveServicePrincipalResponse) {
	if !fromState.ServicePrincipal.IsNull() && !fromState.ServicePrincipal.IsUnknown() {
		if toStateServicePrincipal, ok := toState.GetServicePrincipal(ctx); ok {
			if fromStateServicePrincipal, ok := fromState.GetServicePrincipal(ctx); ok {
				toStateServicePrincipal.SyncFieldsDuringRead(ctx, fromStateServicePrincipal)
				toState.SetServicePrincipal(ctx, toStateServicePrincipal)
			}
		}
	}
}

func (c ResolveServicePrincipalResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalResponse
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": o.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": ServicePrincipal{}.Type(ctx),
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in ResolveServicePrincipalResponse as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveServicePrincipalResponse) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if o.ServicePrincipal.IsNull() || o.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := o.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in ResolveServicePrincipalResponse.
func (o *ResolveServicePrincipalResponse) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.ServicePrincipal = vs
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserProxyRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserProxyRequest
// only implements ToObjectValue() and Type().
func (o ResolveUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserRequest
// only implements ToObjectValue() and Type().
func (o ResolveUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveUserResponse struct {
	// The user that was resolved.
	User types.Object `tfsdk:"user"`
}

func (toState *ResolveUserResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveUserResponse) {
	if !fromPlan.User.IsNull() && !fromPlan.User.IsUnknown() {
		if toStateUser, ok := toState.GetUser(ctx); ok {
			if fromPlanUser, ok := fromPlan.GetUser(ctx); ok {
				toStateUser.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUser)
				toState.SetUser(ctx, toStateUser)
			}
		}
	}
}

func (toState *ResolveUserResponse) SyncFieldsDuringRead(ctx context.Context, fromState ResolveUserResponse) {
	if !fromState.User.IsNull() && !fromState.User.IsUnknown() {
		if toStateUser, ok := toState.GetUser(ctx); ok {
			if fromStateUser, ok := fromState.GetUser(ctx); ok {
				toStateUser.SyncFieldsDuringRead(ctx, fromStateUser)
				toState.SetUser(ctx, toStateUser)
			}
		}
	}
}

func (c ResolveUserResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveUserResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserResponse
// only implements ToObjectValue() and Type().
func (o ResolveUserResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": o.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": User{}.Type(ctx),
		},
	}
}

// GetUser returns the value of the User field in ResolveUserResponse as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveUserResponse) GetUser(ctx context.Context) (User, bool) {
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

// SetUser sets the value of the User field in ResolveUserResponse.
func (o *ResolveUserResponse) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	o.User = vs
}

// The details of a ServicePrincipal resource.
type ServicePrincipal struct {
	// The parent account ID for the service principal in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a service principal in a Databricks account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id"`
	// Display name of the service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// ExternalId of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal service principal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`
}

func (toState *ServicePrincipal) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ServicePrincipal) {
}

func (toState *ServicePrincipal) SyncFieldsDuringRead(ctx context.Context, fromState ServicePrincipal) {
}

func (c ServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_sp_status"] = attrs["account_sp_status"].SetOptional()
	attrs["application_id"] = attrs["application_id"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()

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
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal
// only implements ToObjectValue() and Type().
func (o ServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        o.AccountId,
			"account_sp_status": o.AccountSpStatus,
			"application_id":    o.ApplicationId,
			"display_name":      o.DisplayName,
			"external_id":       o.ExternalId,
			"internal_id":       o.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":        types.StringType,
			"account_sp_status": types.StringType,
			"application_id":    types.StringType,
			"display_name":      types.StringType,
			"external_id":       types.StringType,
			"internal_id":       types.Int64Type,
		},
	}
}

// The details of a User resource.
type User struct {
	// The accountId parent of the user in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// ExternalId of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal userId of the user in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`

	Name types.Object `tfsdk:"name"`
	// Username/email of the user.
	Username types.String `tfsdk:"username"`
}

func (toState *User) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan User) {
	if !fromPlan.Name.IsNull() && !fromPlan.Name.IsUnknown() {
		if toStateName, ok := toState.GetName(ctx); ok {
			if fromPlanName, ok := fromPlan.GetName(ctx); ok {
				toStateName.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanName)
				toState.SetName(ctx, toStateName)
			}
		}
	}
}

func (toState *User) SyncFieldsDuringRead(ctx context.Context, fromState User) {
	if !fromState.Name.IsNull() && !fromState.Name.IsUnknown() {
		if toStateName, ok := toState.GetName(ctx); ok {
			if fromStateName, ok := fromState.GetName(ctx); ok {
				toStateName.SyncFieldsDuringRead(ctx, fromStateName)
				toState.SetName(ctx, toStateName)
			}
		}
	}
}

func (c User) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetOptional()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["username"] = attrs["username"].SetRequired()

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
		"name": reflect.TypeOf(UserName{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User
// only implements ToObjectValue() and Type().
func (o User) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":          o.AccountId,
			"account_user_status": o.AccountUserStatus,
			"external_id":         o.ExternalId,
			"internal_id":         o.InternalId,
			"name":                o.Name,
			"username":            o.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (o User) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"external_id":         types.StringType,
			"internal_id":         types.Int64Type,
			"name":                UserName{}.Type(ctx),
			"username":            types.StringType,
		},
	}
}

// GetName returns the value of the Name field in User as
// a UserName value.
// If the field is unknown or null, the boolean return value is false.
func (o *User) GetName(ctx context.Context) (UserName, bool) {
	var e UserName
	if o.Name.IsNull() || o.Name.IsUnknown() {
		return e, false
	}
	var v UserName
	d := o.Name.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetName sets the value of the Name field in User.
func (o *User) SetName(ctx context.Context, v UserName) {
	vs := v.ToObjectValue(ctx)
	o.Name = vs
}

type UserName struct {
	FamilyName types.String `tfsdk:"family_name"`

	GivenName types.String `tfsdk:"given_name"`
}

func (toState *UserName) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UserName) {
}

func (toState *UserName) SyncFieldsDuringRead(ctx context.Context, fromState UserName) {
}

func (c UserName) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["family_name"] = attrs["family_name"].SetOptional()
	attrs["given_name"] = attrs["given_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UserName.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UserName) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserName
// only implements ToObjectValue() and Type().
func (o UserName) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": o.FamilyName,
			"given_name":  o.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UserName) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"family_name": types.StringType,
			"given_name":  types.StringType,
		},
	}
}

// The details of a principal's access to a workspace.
type WorkspaceAccessDetail struct {
	AccessType types.String `tfsdk:"access_type"`
	// The account ID parent of the workspace where the principal has access.
	AccountId types.String `tfsdk:"account_id"`
	// The permissions granted to the principal in the workspace.
	Permissions types.List `tfsdk:"permissions"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`

	PrincipalType types.String `tfsdk:"principal_type"`
	// The activity status of the principal in the workspace. Not applicable for
	// groups at the moment.
	Status types.String `tfsdk:"status"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (toState *WorkspaceAccessDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WorkspaceAccessDetail) {
}

func (toState *WorkspaceAccessDetail) SyncFieldsDuringRead(ctx context.Context, fromState WorkspaceAccessDetail) {
}

func (c WorkspaceAccessDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_type"] = attrs["access_type"].SetComputed()
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["permissions"] = attrs["permissions"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetComputed()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceAccessDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceAccessDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAccessDetail
// only implements ToObjectValue() and Type().
func (o WorkspaceAccessDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_type":    o.AccessType,
			"account_id":     o.AccountId,
			"permissions":    o.Permissions,
			"principal_id":   o.PrincipalId,
			"principal_type": o.PrincipalType,
			"status":         o.Status,
			"workspace_id":   o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceAccessDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_type": types.StringType,
			"account_id":  types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"status":         types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

// GetPermissions returns the value of the Permissions field in WorkspaceAccessDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceAccessDetail) GetPermissions(ctx context.Context) ([]types.String, bool) {
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

// SetPermissions sets the value of the Permissions field in WorkspaceAccessDetail.
func (o *WorkspaceAccessDetail) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}
