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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GetWorkspaceAccessDetailLocalRequest_SdkV2 struct {
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
func (a GetWorkspaceAccessDetailLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAccessDetailLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"view":         o.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"view":         types.StringType,
		},
	}
}

type GetWorkspaceAccessDetailRequest_SdkV2 struct {
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
func (a GetWorkspaceAccessDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceAccessDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": o.PrincipalId,
			"view":         o.View,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"view":         types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// The details of a Group resource.
type Group_SdkV2 struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Display name of the group.
	GroupName types.String `tfsdk:"group_name"`
	// Internal group ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`
}

func (toState *Group_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Group_SdkV2) {
}

func (toState *Group_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Group_SdkV2) {
}

func (c Group_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Group_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group_SdkV2
// only implements ToObjectValue() and Type().
func (o Group_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Group_SdkV2) Type(ctx context.Context) attr.Type {
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
type ResolveGroupProxyRequest_SdkV2 struct {
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
func (a ResolveGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// groupname, and inherited parent groups.
type ResolveGroupRequest_SdkV2 struct {
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
func (a ResolveGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveGroupResponse_SdkV2 struct {
	// The group that was resolved.
	Group types.List `tfsdk:"group"`
}

func (toState *ResolveGroupResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveGroupResponse_SdkV2) {
	if !fromPlan.Group.IsNull() && !fromPlan.Group.IsUnknown() {
		if toStateGroup, ok := toState.GetGroup(ctx); ok {
			if fromPlanGroup, ok := fromPlan.GetGroup(ctx); ok {
				toStateGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGroup)
				toState.SetGroup(ctx, toStateGroup)
			}
		}
	}
}

func (toState *ResolveGroupResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ResolveGroupResponse_SdkV2) {
	if !fromState.Group.IsNull() && !fromState.Group.IsUnknown() {
		if toStateGroup, ok := toState.GetGroup(ctx); ok {
			if fromStateGroup, ok := fromState.GetGroup(ctx); ok {
				toStateGroup.SyncFieldsDuringRead(ctx, fromStateGroup)
				toState.SetGroup(ctx, toStateGroup)
			}
		}
	}
}

func (c ResolveGroupResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetOptional()
	attrs["group"] = attrs["group"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveGroupResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveGroupResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": o.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveGroupResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGroup returns the value of the Group field in ResolveGroupResponse_SdkV2 as
// a Group_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveGroupResponse_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if o.Group.IsNull() || o.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := o.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in ResolveGroupResponse_SdkV2.
func (o *ResolveGroupResponse_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	o.Group = types.ListValueMust(t, vs)
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest_SdkV2 struct {
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
func (a ResolveServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalRequest_SdkV2 struct {
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
func (a ResolveServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveServicePrincipalResponse_SdkV2 struct {
	// The service principal that was resolved.
	ServicePrincipal types.List `tfsdk:"service_principal"`
}

func (toState *ResolveServicePrincipalResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveServicePrincipalResponse_SdkV2) {
	if !fromPlan.ServicePrincipal.IsNull() && !fromPlan.ServicePrincipal.IsUnknown() {
		if toStateServicePrincipal, ok := toState.GetServicePrincipal(ctx); ok {
			if fromPlanServicePrincipal, ok := fromPlan.GetServicePrincipal(ctx); ok {
				toStateServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanServicePrincipal)
				toState.SetServicePrincipal(ctx, toStateServicePrincipal)
			}
		}
	}
}

func (toState *ResolveServicePrincipalResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ResolveServicePrincipalResponse_SdkV2) {
	if !fromState.ServicePrincipal.IsNull() && !fromState.ServicePrincipal.IsUnknown() {
		if toStateServicePrincipal, ok := toState.GetServicePrincipal(ctx); ok {
			if fromStateServicePrincipal, ok := fromState.GetServicePrincipal(ctx); ok {
				toStateServicePrincipal.SyncFieldsDuringRead(ctx, fromStateServicePrincipal)
				toState.SetServicePrincipal(ctx, toStateServicePrincipal)
			}
		}
	}
}

func (c ResolveServicePrincipalResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetOptional()
	attrs["service_principal"] = attrs["service_principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveServicePrincipalResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveServicePrincipalResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": o.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveServicePrincipalResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in ResolveServicePrincipalResponse_SdkV2 as
// a ServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveServicePrincipalResponse_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if o.ServicePrincipal.IsNull() || o.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := o.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in ResolveServicePrincipalResponse_SdkV2.
func (o *ResolveServicePrincipalResponse_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	o.ServicePrincipal = types.ListValueMust(t, vs)
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserProxyRequest_SdkV2 struct {
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
func (a ResolveUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserRequest_SdkV2 struct {
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
func (a ResolveUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

type ResolveUserResponse_SdkV2 struct {
	// The user that was resolved.
	User types.List `tfsdk:"user"`
}

func (toState *ResolveUserResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ResolveUserResponse_SdkV2) {
	if !fromPlan.User.IsNull() && !fromPlan.User.IsUnknown() {
		if toStateUser, ok := toState.GetUser(ctx); ok {
			if fromPlanUser, ok := fromPlan.GetUser(ctx); ok {
				toStateUser.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUser)
				toState.SetUser(ctx, toStateUser)
			}
		}
	}
}

func (toState *ResolveUserResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ResolveUserResponse_SdkV2) {
	if !fromState.User.IsNull() && !fromState.User.IsUnknown() {
		if toStateUser, ok := toState.GetUser(ctx); ok {
			if fromStateUser, ok := fromState.GetUser(ctx); ok {
				toStateUser.SyncFieldsDuringRead(ctx, fromStateUser)
				toState.SetUser(ctx, toStateUser)
			}
		}
	}
}

func (c ResolveUserResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetOptional()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolveUserResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolveUserResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": o.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolveUserResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUser returns the value of the User field in ResolveUserResponse_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolveUserResponse_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in ResolveUserResponse_SdkV2.
func (o *ResolveUserResponse_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	o.User = types.ListValueMust(t, vs)
}

// The details of a ServicePrincipal resource.
type ServicePrincipal_SdkV2 struct {
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

func (toState *ServicePrincipal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ServicePrincipal_SdkV2) {
}

func (toState *ServicePrincipal_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ServicePrincipal_SdkV2) {
}

func (c ServicePrincipal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServicePrincipal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal_SdkV2
// only implements ToObjectValue() and Type().
func (o ServicePrincipal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
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
type User_SdkV2 struct {
	// The accountId parent of the user in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// ExternalId of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal userId of the user in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`

	Name types.List `tfsdk:"name"`
	// Username/email of the user.
	Username types.String `tfsdk:"username"`
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
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetOptional()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a User_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"name": reflect.TypeOf(UserName_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User_SdkV2
// only implements ToObjectValue() and Type().
func (o User_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o User_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"external_id":         types.StringType,
			"internal_id":         types.Int64Type,
			"name": basetypes.ListType{
				ElemType: UserName_SdkV2{}.Type(ctx),
			},
			"username": types.StringType,
		},
	}
}

// GetName returns the value of the Name field in User_SdkV2 as
// a UserName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *User_SdkV2) GetName(ctx context.Context) (UserName_SdkV2, bool) {
	var e UserName_SdkV2
	if o.Name.IsNull() || o.Name.IsUnknown() {
		return e, false
	}
	var v []UserName_SdkV2
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
func (o *User_SdkV2) SetName(ctx context.Context, v UserName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["name"]
	o.Name = types.ListValueMust(t, vs)
}

type UserName_SdkV2 struct {
	FamilyName types.String `tfsdk:"family_name"`

	GivenName types.String `tfsdk:"given_name"`
}

func (toState *UserName_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UserName_SdkV2) {
}

func (toState *UserName_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UserName_SdkV2) {
}

func (c UserName_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UserName_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserName_SdkV2
// only implements ToObjectValue() and Type().
func (o UserName_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": o.FamilyName,
			"given_name":  o.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UserName_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"family_name": types.StringType,
			"given_name":  types.StringType,
		},
	}
}

// The details of a principal's access to a workspace.
type WorkspaceAccessDetail_SdkV2 struct {
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

func (toState *WorkspaceAccessDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan WorkspaceAccessDetail_SdkV2) {
}

func (toState *WorkspaceAccessDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState WorkspaceAccessDetail_SdkV2) {
}

func (c WorkspaceAccessDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceAccessDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAccessDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceAccessDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WorkspaceAccessDetail_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetPermissions returns the value of the Permissions field in WorkspaceAccessDetail_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceAccessDetail_SdkV2) GetPermissions(ctx context.Context) ([]types.String, bool) {
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

// SetPermissions sets the value of the Permissions field in WorkspaceAccessDetail_SdkV2.
func (o *WorkspaceAccessDetail_SdkV2) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Permissions = types.ListValueMust(t, vs)
}
