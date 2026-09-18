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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDirectGroupMemberProxyRequest struct {
	// Required. The group membership to create.
	DirectGroupMember types.Object `tfsdk:"direct_group_member"`
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
}

func (to *CreateDirectGroupMemberProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDirectGroupMemberProxyRequest) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				// Recursively sync the fields of DirectGroupMember
				toDirectGroupMember.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (to *CreateDirectGroupMemberProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDirectGroupMemberProxyRequest) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				toDirectGroupMember.SyncFieldsDuringRead(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (m CreateDirectGroupMemberProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direct_group_member"] = attrs["direct_group_member"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDirectGroupMemberProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_member": reflect.TypeOf(DirectGroupMember{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectGroupMemberProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateDirectGroupMemberProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_member": m.DirectGroupMember,
			"group_id":            m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDirectGroupMemberProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_member": DirectGroupMember{}.Type(ctx),
			"group_id":            types.Int64Type,
		},
	}
}

// GetDirectGroupMember returns the value of the DirectGroupMember field in CreateDirectGroupMemberProxyRequest as
// a DirectGroupMember value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDirectGroupMemberProxyRequest) GetDirectGroupMember(ctx context.Context) (DirectGroupMember, bool) {
	var e DirectGroupMember
	if m.DirectGroupMember.IsNull() || m.DirectGroupMember.IsUnknown() {
		return e, false
	}
	var v DirectGroupMember
	d := m.DirectGroupMember.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectGroupMember sets the value of the DirectGroupMember field in CreateDirectGroupMemberProxyRequest.
func (m *CreateDirectGroupMemberProxyRequest) SetDirectGroupMember(ctx context.Context, v DirectGroupMember) {
	vs := v.ToObjectValue(ctx)
	m.DirectGroupMember = vs
}

type CreateDirectGroupMemberRequest struct {
	// Required. The direct group member to be added to the group.
	DirectGroupMember types.Object `tfsdk:"direct_group_member"`
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
}

func (to *CreateDirectGroupMemberRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDirectGroupMemberRequest) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				// Recursively sync the fields of DirectGroupMember
				toDirectGroupMember.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (to *CreateDirectGroupMemberRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDirectGroupMemberRequest) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				toDirectGroupMember.SyncFieldsDuringRead(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (m CreateDirectGroupMemberRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direct_group_member"] = attrs["direct_group_member"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDirectGroupMemberRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_member": reflect.TypeOf(DirectGroupMember{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectGroupMemberRequest
// only implements ToObjectValue() and Type().
func (m CreateDirectGroupMemberRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_member": m.DirectGroupMember,
			"group_id":            m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDirectGroupMemberRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_member": DirectGroupMember{}.Type(ctx),
			"group_id":            types.Int64Type,
		},
	}
}

// GetDirectGroupMember returns the value of the DirectGroupMember field in CreateDirectGroupMemberRequest as
// a DirectGroupMember value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDirectGroupMemberRequest) GetDirectGroupMember(ctx context.Context) (DirectGroupMember, bool) {
	var e DirectGroupMember
	if m.DirectGroupMember.IsNull() || m.DirectGroupMember.IsUnknown() {
		return e, false
	}
	var v DirectGroupMember
	d := m.DirectGroupMember.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectGroupMember sets the value of the DirectGroupMember field in CreateDirectGroupMemberRequest.
func (m *CreateDirectGroupMemberRequest) SetDirectGroupMember(ctx context.Context, v DirectGroupMember) {
	vs := v.ToObjectValue(ctx)
	m.DirectGroupMember = vs
}

type CreateGroupProxyRequest struct {
	// Required. Group to be created in <Databricks>
	Group types.Object `tfsdk:"group"`
}

func (to *CreateGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGroupProxyRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				// Recursively sync the fields of Group
				toGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (to *CreateGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateGroupProxyRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m CreateGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": Group{}.Type(ctx),
		},
	}
}

// GetGroup returns the value of the Group field in CreateGroupProxyRequest as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateGroupProxyRequest) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := m.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in CreateGroupProxyRequest.
func (m *CreateGroupProxyRequest) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	m.Group = vs
}

type CreateGroupRequest struct {
	// Required. Group to be created in <Databricks>
	Group types.Object `tfsdk:"group"`
}

func (to *CreateGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGroupRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				// Recursively sync the fields of Group
				toGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (to *CreateGroupRequest) SyncFieldsDuringRead(ctx context.Context, from CreateGroupRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m CreateGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGroupRequest
// only implements ToObjectValue() and Type().
func (m CreateGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": Group{}.Type(ctx),
		},
	}
}

// GetGroup returns the value of the Group field in CreateGroupRequest as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateGroupRequest) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := m.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in CreateGroupRequest.
func (m *CreateGroupRequest) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	m.Group = vs
}

type CreateServicePrincipalProxyRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
}

func (to *CreateServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalProxyRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				// Recursively sync the fields of ServicePrincipal
				toServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (to *CreateServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalProxyRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m CreateServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": ServicePrincipal{}.Type(ctx),
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in CreateServicePrincipalProxyRequest as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServicePrincipalProxyRequest) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := m.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in CreateServicePrincipalProxyRequest.
func (m *CreateServicePrincipalProxyRequest) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	m.ServicePrincipal = vs
}

type CreateServicePrincipalRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
}

func (to *CreateServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				// Recursively sync the fields of ServicePrincipal
				toServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (to *CreateServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m CreateServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": ServicePrincipal{}.Type(ctx),
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in CreateServicePrincipalRequest as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServicePrincipalRequest) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := m.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in CreateServicePrincipalRequest.
func (m *CreateServicePrincipalRequest) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	m.ServicePrincipal = vs
}

type CreateUserProxyRequest struct {
	// Required. User to be created in <Databricks>
	User types.Object `tfsdk:"user"`
}

func (to *CreateUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateUserProxyRequest) {
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

func (to *CreateUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateUserProxyRequest) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m CreateUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateUserProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": User{}.Type(ctx),
		},
	}
}

// GetUser returns the value of the User field in CreateUserProxyRequest as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateUserProxyRequest) GetUser(ctx context.Context) (User, bool) {
	var e User
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v User
	d := m.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in CreateUserProxyRequest.
func (m *CreateUserProxyRequest) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	m.User = vs
}

type CreateUserRequest struct {
	// Required. User to be created in <Databricks>
	User types.Object `tfsdk:"user"`
}

func (to *CreateUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateUserRequest) {
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

func (to *CreateUserRequest) SyncFieldsDuringRead(ctx context.Context, from CreateUserRequest) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m CreateUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateUserRequest
// only implements ToObjectValue() and Type().
func (m CreateUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": User{}.Type(ctx),
		},
	}
}

// GetUser returns the value of the User field in CreateUserRequest as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateUserRequest) GetUser(ctx context.Context) (User, bool) {
	var e User
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v User
	d := m.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in CreateUserRequest.
func (m *CreateUserRequest) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	m.User = vs
}

type CreateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail types.Object `tfsdk:"workspace_assignment_detail"`
}

func (to *CreateWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentDetailProxyRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignmentDetail
				toWorkspaceAssignmentDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (to *CreateWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentDetailProxyRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentDetailProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentDetailProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentDetailProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentDetailProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentDetailProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentDetailProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment_detail": WorkspaceAssignmentDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailProxyRequest as
// a WorkspaceAssignmentDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentDetailProxyRequest) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail, bool) {
	var e WorkspaceAssignmentDetail
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignmentDetail
	d := m.WorkspaceAssignmentDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailProxyRequest.
func (m *CreateWorkspaceAssignmentDetailProxyRequest) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignmentDetail = vs
}

type CreateWorkspaceAssignmentDetailRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail types.Object `tfsdk:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *CreateWorkspaceAssignmentDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentDetailRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignmentDetail
				toWorkspaceAssignmentDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (to *CreateWorkspaceAssignmentDetailRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentDetailRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentDetailRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
			"workspace_id":                m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment_detail": WorkspaceAssignmentDetail{}.Type(ctx),
			"workspace_id":                types.Int64Type,
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailRequest as
// a WorkspaceAssignmentDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentDetailRequest) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail, bool) {
	var e WorkspaceAssignmentDetail
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignmentDetail
	d := m.WorkspaceAssignmentDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailRequest.
func (m *CreateWorkspaceAssignmentDetailRequest) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignmentDetail = vs
}

type CreateWorkspaceAssignmentProxyRequest struct {
	// Required. Workspace assignment to be created in <Databricks>.
	WorkspaceAssignment types.Object `tfsdk:"workspace_assignment"`
}

func (to *CreateWorkspaceAssignmentProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentProxyRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignment
				toWorkspaceAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (to *CreateWorkspaceAssignmentProxyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentProxyRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentProxyRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment": WorkspaceAssignment{}.Type(ctx),
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentProxyRequest as
// a WorkspaceAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentProxyRequest) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment, bool) {
	var e WorkspaceAssignment
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignment
	d := m.WorkspaceAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentProxyRequest.
func (m *CreateWorkspaceAssignmentProxyRequest) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignment = vs
}

type CreateWorkspaceAssignmentRequest struct {
	// Required. Workspace assignment to be created in <Databricks>.
	WorkspaceAssignment types.Object `tfsdk:"workspace_assignment"`
	// Required. The workspace ID for which the workspace assignment is being
	// created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *CreateWorkspaceAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignment
				toWorkspaceAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (to *CreateWorkspaceAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment": m.WorkspaceAssignment,
			"workspace_id":         m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment": WorkspaceAssignment{}.Type(ctx),
			"workspace_id":         types.Int64Type,
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentRequest as
// a WorkspaceAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentRequest) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment, bool) {
	var e WorkspaceAssignment
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignment
	d := m.WorkspaceAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentRequest.
func (m *CreateWorkspaceAssignmentRequest) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignment = vs
}

type DeleteDirectGroupMemberProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
	// Required. Internal ID of the principal to be unassigned from the group.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteDirectGroupMemberProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDirectGroupMemberProxyRequest) {
}

func (to *DeleteDirectGroupMemberProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDirectGroupMemberProxyRequest) {
}

func (m DeleteDirectGroupMemberProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDirectGroupMemberProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectGroupMemberProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteDirectGroupMemberProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":     m.GroupId,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDirectGroupMemberProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":     types.Int64Type,
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteDirectGroupMemberRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
	// Required. Internal ID of the principal to be unassigned from the group.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteDirectGroupMemberRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDirectGroupMemberRequest) {
}

func (to *DeleteDirectGroupMemberRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDirectGroupMemberRequest) {
}

func (m DeleteDirectGroupMemberRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDirectGroupMemberRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectGroupMemberRequest
// only implements ToObjectValue() and Type().
func (m DeleteDirectGroupMemberRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":     m.GroupId,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDirectGroupMemberRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":     types.Int64Type,
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
}

func (to *DeleteGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupProxyRequest) {
}

func (to *DeleteGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupProxyRequest) {
}

func (m DeleteGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id": m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id": types.StringType,
		},
	}
}

type DeleteGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
}

func (to *DeleteGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupRequest) {
}

func (to *DeleteGroupRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupRequest) {
}

func (m DeleteGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupRequest
// only implements ToObjectValue() and Type().
func (m DeleteGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id": m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id": types.StringType,
		},
	}
}

type DeleteServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalProxyRequest) {
}

func (to *DeleteServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalProxyRequest) {
}

func (m DeleteServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_id": types.StringType,
		},
	}
}

type DeleteServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalRequest) {
}

func (to *DeleteServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalRequest) {
}

func (m DeleteServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_id": types.StringType,
		},
	}
}

type DeleteUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *DeleteUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserProxyRequest) {
}

func (to *DeleteUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteUserProxyRequest) {
}

func (m DeleteUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.StringType,
		},
	}
}

type DeleteUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *DeleteUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserRequest) {
}

func (to *DeleteUserRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteUserRequest) {
}

func (m DeleteUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserRequest
// only implements ToObjectValue() and Type().
func (m DeleteUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.StringType,
		},
	}
}

type DeleteWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentDetailProxyRequest) {
}

func (to *DeleteWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentDetailProxyRequest) {
}

func (m DeleteWorkspaceAssignmentDetailProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentDetailProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentDetailProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentDetailProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentDetailProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentDetailProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentDetailRequest) {
}

func (to *DeleteWorkspaceAssignmentDetailRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentDetailRequest) {
}

func (m DeleteWorkspaceAssignmentDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentDetailRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAssignmentProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentProxyRequest) {
}

func (to *DeleteWorkspaceAssignmentProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentProxyRequest) {
}

func (m DeleteWorkspaceAssignmentProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentProxyRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAssignmentRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentRequest) {
}

func (to *DeleteWorkspaceAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentRequest) {
}

func (m DeleteWorkspaceAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

// Represents a principal that is a direct member of a group, with its source of
// membership.
type DirectGroupMember struct {
	// Display name of the principal.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the principal in Databricks.
	ExternalId types.String `tfsdk:"external_id"`
	// The internal ID of the group this member belongs to.
	GroupId types.Int64 `tfsdk:"group_id"`
	// The source of group membership (internal or from identity provider).
	MembershipSource types.String `tfsdk:"membership_source"`
	// Internal ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group).
	PrincipalType types.String `tfsdk:"principal_type"`
}

func (to *DirectGroupMember) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectGroupMember) {
}

func (to *DirectGroupMember) SyncFieldsDuringRead(ctx context.Context, from DirectGroupMember) {
}

func (m DirectGroupMember) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["membership_source"] = attrs["membership_source"].SetComputed()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["principal_type"] = attrs["principal_type"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectGroupMember.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DirectGroupMember) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectGroupMember
// only implements ToObjectValue() and Type().
func (m DirectGroupMember) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":      m.DisplayName,
			"external_id":       m.ExternalId,
			"group_id":          m.GroupId,
			"membership_source": m.MembershipSource,
			"principal_id":      m.PrincipalId,
			"principal_type":    m.PrincipalType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectGroupMember) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":      types.StringType,
			"external_id":       types.StringType,
			"group_id":          types.Int64Type,
			"membership_source": types.StringType,
			"principal_id":      types.Int64Type,
			"principal_type":    types.StringType,
		},
	}
}

// An external group from the customer's Identity Provider, resolved into
// Databricks. This is a read-only resource keyed by the IdP external ID. The
// Get method may trigger an idempotent sync from the customer's IdP to
// provision or refresh the group's data in Databricks.
type ExternalGroup struct {
	// The parent account ID, from Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// Display name of the group from the customer's IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the group in the customer's IdP.
	ExternalGroupId types.String `tfsdk:"external_group_id"`
	// Internal groupId of the group in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external group. The format depends on the API
	// that returned it: - Account-scoped:
	// accounts/{account_id}/external-groups/{external_group_id} -
	// Workspace-scoped: external-groups/{external_group_id}
	Name types.String `tfsdk:"name"`
}

func (to *ExternalGroup) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalGroup) {
}

func (to *ExternalGroup) SyncFieldsDuringRead(ctx context.Context, from ExternalGroup) {
}

func (m ExternalGroup) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_group_id"] = attrs["external_group_id"].SetComputed()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalGroup.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalGroup) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalGroup
// only implements ToObjectValue() and Type().
func (m ExternalGroup) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        m.AccountId,
			"display_name":      m.DisplayName,
			"external_group_id": m.ExternalGroupId,
			"internal_id":       m.InternalId,
			"name":              m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalGroup) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":        types.StringType,
			"display_name":      types.StringType,
			"external_group_id": types.StringType,
			"internal_id":       types.StringType,
			"name":              types.StringType,
		},
	}
}

// An external service principal from the customer's Identity Provider, resolved
// into Databricks. This is a read-only resource keyed by the IdP external ID.
// The Get method may trigger an idempotent sync from the customer's IdP to
// provision or refresh the service principal's data in Databricks.
type ExternalServicePrincipal struct {
	// The parent account ID, from Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of the service principal in the Databricks account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// Application ID of the service principal, from the customer's IdP.
	ApplicationId types.String `tfsdk:"application_id"`
	// Display name of the service principal, from the customer's IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the service principal in the customer's IdP.
	ExternalServicePrincipalId types.String `tfsdk:"external_service_principal_id"`
	// Internal servicePrincipalId of the service principal in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external service principal. The format depends
	// on the API that returned it: - Account-scoped:
	// accounts/{account_id}/external-service-principals/{external_service_principal_id}
	// - Workspace-scoped:
	// external-service-principals/{external_service_principal_id}
	Name types.String `tfsdk:"name"`
}

func (to *ExternalServicePrincipal) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalServicePrincipal) {
}

func (to *ExternalServicePrincipal) SyncFieldsDuringRead(ctx context.Context, from ExternalServicePrincipal) {
}

func (m ExternalServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_sp_status"] = attrs["account_sp_status"].SetComputed()
	attrs["application_id"] = attrs["application_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_service_principal_id"] = attrs["external_service_principal_id"].SetComputed()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalServicePrincipal
// only implements ToObjectValue() and Type().
func (m ExternalServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                    m.AccountId,
			"account_sp_status":             m.AccountSpStatus,
			"application_id":                m.ApplicationId,
			"display_name":                  m.DisplayName,
			"external_service_principal_id": m.ExternalServicePrincipalId,
			"internal_id":                   m.InternalId,
			"name":                          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                    types.StringType,
			"account_sp_status":             types.StringType,
			"application_id":                types.StringType,
			"display_name":                  types.StringType,
			"external_service_principal_id": types.StringType,
			"internal_id":                   types.StringType,
			"name":                          types.StringType,
		},
	}
}

// An external user from the customer's Identity Provider, resolved into
// Databricks. This is a read-only resource that allows customers to look up
// external user identities by their IdP external ID and retrieve the
// corresponding Databricks internal ID and metadata. The Get method may trigger
// an idempotent sync from the customer's IdP to provision or refresh the user's
// data in Databricks.
type ExternalUser struct {
	// The parent account ID, from Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of the user in the Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// Display name of the user from the customer's IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the user in the customer's IdP.
	ExternalUserId types.String `tfsdk:"external_user_id"`
	// The full name of the user, from the customer's IdP.
	FullName types.Object `tfsdk:"full_name"`
	// Internal userId of the user in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external user. The format depends on the API
	// that returned it: - Account-scoped:
	// accounts/{account_id}/external-users/{external_user_id} -
	// Workspace-scoped: external-users/{external_user_id}
	Name types.String `tfsdk:"name"`
	// Username/email of the user, from Databricks.
	Username types.String `tfsdk:"username"`
}

func (to *ExternalUser) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalUser) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				// Recursively sync the fields of FullName
				toFullName.SyncFieldsDuringCreateOrUpdate(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (to *ExternalUser) SyncFieldsDuringRead(ctx context.Context, from ExternalUser) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				toFullName.SyncFieldsDuringRead(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (m ExternalUser) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_user_id"] = attrs["external_user_id"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["username"] = attrs["username"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalUser.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalUser) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name": reflect.TypeOf(FullName{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalUser
// only implements ToObjectValue() and Type().
func (m ExternalUser) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":          m.AccountId,
			"account_user_status": m.AccountUserStatus,
			"display_name":        m.DisplayName,
			"external_user_id":    m.ExternalUserId,
			"full_name":           m.FullName,
			"internal_id":         m.InternalId,
			"name":                m.Name,
			"username":            m.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalUser) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"display_name":        types.StringType,
			"external_user_id":    types.StringType,
			"full_name":           FullName{}.Type(ctx),
			"internal_id":         types.StringType,
			"name":                types.StringType,
			"username":            types.StringType,
		},
	}
}

// GetFullName returns the value of the FullName field in ExternalUser as
// a FullName value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalUser) GetFullName(ctx context.Context) (FullName, bool) {
	var e FullName
	if m.FullName.IsNull() || m.FullName.IsUnknown() {
		return e, false
	}
	var v FullName
	d := m.FullName.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullName sets the value of the FullName field in ExternalUser.
func (m *ExternalUser) SetFullName(ctx context.Context, v FullName) {
	vs := v.ToObjectValue(ctx)
	m.FullName = vs
}

// The full name of a user.
type FullName struct {
	// The family (last) name of the user, from the customer's IdP.
	FamilyName types.String `tfsdk:"family_name"`
	// The given (first) name of the user, from the customer's IdP.
	GivenName types.String `tfsdk:"given_name"`
}

func (to *FullName) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FullName) {
}

func (to *FullName) SyncFieldsDuringRead(ctx context.Context, from FullName) {
}

func (m FullName) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["family_name"] = attrs["family_name"].SetComputed()
	attrs["given_name"] = attrs["given_name"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FullName.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FullName) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FullName
// only implements ToObjectValue() and Type().
func (m FullName) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FullName) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"family_name": types.StringType,
			"given_name":  types.StringType,
		},
	}
}

type GetDirectGroupMemberProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
	// Required. Internal ID of the principal belonging to the group in
	// Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetDirectGroupMemberProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDirectGroupMemberProxyRequest) {
}

func (to *GetDirectGroupMemberProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetDirectGroupMemberProxyRequest) {
}

func (m GetDirectGroupMemberProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDirectGroupMemberProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectGroupMemberProxyRequest
// only implements ToObjectValue() and Type().
func (m GetDirectGroupMemberProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":     m.GroupId,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDirectGroupMemberProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":     types.Int64Type,
			"principal_id": types.Int64Type,
		},
	}
}

type GetDirectGroupMemberRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.Int64 `tfsdk:"-"`
	// Required. Internal ID of the principal belonging to the group in
	// Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetDirectGroupMemberRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDirectGroupMemberRequest) {
}

func (to *GetDirectGroupMemberRequest) SyncFieldsDuringRead(ctx context.Context, from GetDirectGroupMemberRequest) {
}

func (m GetDirectGroupMemberRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDirectGroupMemberRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectGroupMemberRequest
// only implements ToObjectValue() and Type().
func (m GetDirectGroupMemberRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":     m.GroupId,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDirectGroupMemberRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":     types.Int64Type,
			"principal_id": types.Int64Type,
		},
	}
}

type GetExternalGroupProxyRequest struct {
	// Required. The resource name of the external group. Format:
	// external-groups/{external_group_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalGroupProxyRequest) {
}

func (to *GetExternalGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalGroupProxyRequest) {
}

func (m GetExternalGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m GetExternalGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalGroupRequest struct {
	// Required. The resource name of the external group. Format:
	// accounts/{account_id}/external-groups/{external_group_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalGroupRequest) {
}

func (to *GetExternalGroupRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalGroupRequest) {
}

func (m GetExternalGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalGroupRequest
// only implements ToObjectValue() and Type().
func (m GetExternalGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalServicePrincipalProxyRequest struct {
	// Required. The resource name of the external service principal. Format:
	// external-service-principals/{external_service_principal_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalServicePrincipalProxyRequest) {
}

func (to *GetExternalServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalServicePrincipalProxyRequest) {
}

func (m GetExternalServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m GetExternalServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalServicePrincipalRequest struct {
	// Required. The resource name of the external service principal. Format:
	// accounts/{account_id}/external-service-principals/{external_service_principal_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalServicePrincipalRequest) {
}

func (to *GetExternalServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalServicePrincipalRequest) {
}

func (m GetExternalServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m GetExternalServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalUserProxyRequest struct {
	// Required. The resource name of the external user. Format:
	// external-users/{external_user_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalUserProxyRequest) {
}

func (to *GetExternalUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalUserProxyRequest) {
}

func (m GetExternalUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalUserProxyRequest
// only implements ToObjectValue() and Type().
func (m GetExternalUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalUserRequest struct {
	// Required. The resource name of the external user. Format:
	// accounts/{account_id}/external-users/{external_user_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalUserRequest) {
}

func (to *GetExternalUserRequest) SyncFieldsDuringRead(ctx context.Context, from GetExternalUserRequest) {
}

func (m GetExternalUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExternalUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalUserRequest
// only implements ToObjectValue() and Type().
func (m GetExternalUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
}

func (to *GetGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupProxyRequest) {
}

func (to *GetGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetGroupProxyRequest) {
}

func (m GetGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m GetGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id": m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id": types.StringType,
		},
	}
}

type GetGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
}

func (to *GetGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupRequest) {
}

func (to *GetGroupRequest) SyncFieldsDuringRead(ctx context.Context, from GetGroupRequest) {
}

func (m GetGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupRequest
// only implements ToObjectValue() and Type().
func (m GetGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id": m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id": types.StringType,
		},
	}
}

type GetServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *GetServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalProxyRequest) {
}

func (to *GetServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalProxyRequest) {
}

func (m GetServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_id": types.StringType,
		},
	}
}

type GetServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *GetServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalRequest) {
}

func (to *GetServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalRequest) {
}

func (m GetServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_id": types.StringType,
		},
	}
}

type GetUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *GetUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserProxyRequest) {
}

func (to *GetUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetUserProxyRequest) {
}

func (m GetUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserProxyRequest
// only implements ToObjectValue() and Type().
func (m GetUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.StringType,
		},
	}
}

type GetUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *GetUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserRequest) {
}

func (to *GetUserRequest) SyncFieldsDuringRead(ctx context.Context, from GetUserRequest) {
}

func (m GetUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserRequest
// only implements ToObjectValue() and Type().
func (m GetUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user_id": m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.StringType,
		},
	}
}

type GetWorkspaceAccessDetailLocalRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAccessDetailLocalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest) {
}

func (to *GetWorkspaceAccessDetailLocalRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest) {
}

func (m GetWorkspaceAccessDetailLocalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["view"] = attrs["view"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAccessDetailLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailLocalRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAccessDetailLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"view":         m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailLocalRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetWorkspaceAccessDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailRequest) {
}

func (to *GetWorkspaceAccessDetailRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailRequest) {
}

func (m GetWorkspaceAccessDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["view"] = attrs["view"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAccessDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAccessDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAccessDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"view":         m.View,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"view":         types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentDetailProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentDetailProxyRequest) {
}

func (to *GetWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentDetailProxyRequest) {
}

func (m GetWorkspaceAssignmentDetailProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentDetailProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentDetailProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentDetailProxyRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentDetailProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentDetailProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The workspace ID for which the assignment details are being
	// requested.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentDetailRequest) {
}

func (to *GetWorkspaceAssignmentDetailRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentDetailRequest) {
}

func (m GetWorkspaceAssignmentDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentDetailRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment is being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentProxyRequest) {
}

func (to *GetWorkspaceAssignmentProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentProxyRequest) {
}

func (m GetWorkspaceAssignmentProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentProxyRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment is being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The workspace ID for which the assignment is being requested.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentRequest) {
}

func (to *GetWorkspaceAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentRequest) {
}

func (m GetWorkspaceAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetWorkspaceIdentityDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// identity details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceIdentityDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceIdentityDetailRequest) {
}

func (to *GetWorkspaceIdentityDetailRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceIdentityDetailRequest) {
}

func (m GetWorkspaceIdentityDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceIdentityDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceIdentityDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceIdentityDetailRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceIdentityDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceIdentityDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

// The details of a Group resource.
type Group struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal group ID of the group in Databricks.
	GroupId types.String `tfsdk:"group_id"`
	// Display name of the group.
	GroupName types.String `tfsdk:"group_name"`
}

func (to *Group) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Group) {
}

func (to *Group) SyncFieldsDuringRead(ctx context.Context, from Group) {
}

func (m Group) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["group_name"] = attrs["group_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Group.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Group) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group
// only implements ToObjectValue() and Type().
func (m Group) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":  m.AccountId,
			"external_id": m.ExternalId,
			"group_id":    m.GroupId,
			"group_name":  m.GroupName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Group) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_id":    types.StringType,
			"group_name":  types.StringType,
		},
	}
}

type ListDirectGroupMembersProxyRequest struct {
	// Required. Internal ID of the group in Databricks whose direct members are
	// being listed.
	GroupId types.Int64 `tfsdk:"-"`
	// The maximum number of members to return. The service may return fewer
	// than this value. If not provided, defaults to 1000, which is also the
	// maximum allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDirectGroupMembersProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersProxyRequest) {
}

func (to *ListDirectGroupMembersProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersProxyRequest) {
}

func (m ListDirectGroupMembersProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectGroupMembersProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDirectGroupMembersProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersProxyRequest
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":   m.GroupId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":   types.Int64Type,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDirectGroupMembersRequest struct {
	// Required. Internal ID of the group in Databricks whose direct members are
	// being listed.
	GroupId types.Int64 `tfsdk:"-"`
	// The maximum number of members to return. The service may return fewer
	// than this value. If not provided, defaults to 1000, which is also the
	// maximum allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListDirectGroupMembers call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDirectGroupMembersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersRequest) {
}

func (to *ListDirectGroupMembersRequest) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersRequest) {
}

func (m ListDirectGroupMembersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectGroupMembersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDirectGroupMembersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersRequest
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":   m.GroupId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":   types.Int64Type,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response message for listing direct group members.
type ListDirectGroupMembersResponse struct {
	// The list of direct group members with their membership source type.
	DirectGroupMembers types.List `tfsdk:"direct_group_members"`
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDirectGroupMembersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersResponse) {
	if !from.DirectGroupMembers.IsNull() && !from.DirectGroupMembers.IsUnknown() && to.DirectGroupMembers.IsNull() && len(from.DirectGroupMembers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DirectGroupMembers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DirectGroupMembers = from.DirectGroupMembers
	}
	if !from.DirectGroupMembers.IsNull() && !from.DirectGroupMembers.IsUnknown() {
		if toDirectGroupMembers, ok := to.GetDirectGroupMembers(ctx); ok {
			if fromDirectGroupMembers, ok := from.GetDirectGroupMembers(ctx); ok {
				// Recursively sync the fields of each DirectGroupMembers element by position.
				for i := range toDirectGroupMembers {
					if i < len(fromDirectGroupMembers) {
						toDirectGroupMembers[i].SyncFieldsDuringCreateOrUpdate(ctx, fromDirectGroupMembers[i])
					}
				}
				to.SetDirectGroupMembers(ctx, toDirectGroupMembers)
			}
		}
	}
}

func (to *ListDirectGroupMembersResponse) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersResponse) {
	if !from.DirectGroupMembers.IsNull() && !from.DirectGroupMembers.IsUnknown() && to.DirectGroupMembers.IsNull() && len(from.DirectGroupMembers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DirectGroupMembers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DirectGroupMembers = from.DirectGroupMembers
	}
	if !from.DirectGroupMembers.IsNull() && !from.DirectGroupMembers.IsUnknown() {
		if toDirectGroupMembers, ok := to.GetDirectGroupMembers(ctx); ok {
			if fromDirectGroupMembers, ok := from.GetDirectGroupMembers(ctx); ok {
				for i := range toDirectGroupMembers {
					if i < len(fromDirectGroupMembers) {
						toDirectGroupMembers[i].SyncFieldsDuringRead(ctx, fromDirectGroupMembers[i])
					}
				}
				to.SetDirectGroupMembers(ctx, toDirectGroupMembers)
			}
		}
	}
}

func (m ListDirectGroupMembersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direct_group_members"] = attrs["direct_group_members"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectGroupMembersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDirectGroupMembersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_members": reflect.TypeOf(DirectGroupMember{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersResponse
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_members": m.DirectGroupMembers,
			"next_page_token":      m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_members": basetypes.ListType{
				ElemType: DirectGroupMember{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDirectGroupMembers returns the value of the DirectGroupMembers field in ListDirectGroupMembersResponse as
// a slice of DirectGroupMember values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDirectGroupMembersResponse) GetDirectGroupMembers(ctx context.Context) ([]DirectGroupMember, bool) {
	if m.DirectGroupMembers.IsNull() || m.DirectGroupMembers.IsUnknown() {
		return nil, false
	}
	var v []DirectGroupMember
	d := m.DirectGroupMembers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectGroupMembers sets the value of the DirectGroupMembers field in ListDirectGroupMembersResponse.
func (m *ListDirectGroupMembersResponse) SetDirectGroupMembers(ctx context.Context, v []DirectGroupMember) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_group_members"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DirectGroupMembers = types.ListValueMust(t, vs)
}

type ListGroupsProxyRequest struct {
	// Optional. Allows filtering groups by group name or external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of groups to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListGroupsProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsProxyRequest) {
}

func (to *ListGroupsProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListGroupsProxyRequest) {
}

func (m ListGroupsProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGroupsProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsProxyRequest
// only implements ToObjectValue() and Type().
func (m ListGroupsProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListGroupsRequest struct {
	// Optional. Allows filtering groups by group name or external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of groups to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListGroupsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsRequest) {
}

func (to *ListGroupsRequest) SyncFieldsDuringRead(ctx context.Context, from ListGroupsRequest) {
}

func (m ListGroupsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsRequest
// only implements ToObjectValue() and Type().
func (m ListGroupsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response message containing a page of groups in the account.
type ListGroupsResponse struct {
	Groups types.List `tfsdk:"groups"`
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListGroupsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsResponse) {
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() && to.Groups.IsNull() && len(from.Groups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Groups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Groups = from.Groups
	}
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() {
		if toGroups, ok := to.GetGroups(ctx); ok {
			if fromGroups, ok := from.GetGroups(ctx); ok {
				// Recursively sync the fields of each Groups element by position.
				for i := range toGroups {
					if i < len(fromGroups) {
						toGroups[i].SyncFieldsDuringCreateOrUpdate(ctx, fromGroups[i])
					}
				}
				to.SetGroups(ctx, toGroups)
			}
		}
	}
}

func (to *ListGroupsResponse) SyncFieldsDuringRead(ctx context.Context, from ListGroupsResponse) {
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() && to.Groups.IsNull() && len(from.Groups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Groups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Groups = from.Groups
	}
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() {
		if toGroups, ok := to.GetGroups(ctx); ok {
			if fromGroups, ok := from.GetGroups(ctx); ok {
				for i := range toGroups {
					if i < len(fromGroups) {
						toGroups[i].SyncFieldsDuringRead(ctx, fromGroups[i])
					}
				}
				to.SetGroups(ctx, toGroups)
			}
		}
	}
}

func (m ListGroupsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["groups"] = attrs["groups"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"groups": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsResponse
// only implements ToObjectValue() and Type().
func (m ListGroupsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"groups":          m.Groups,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"groups": basetypes.ListType{
				ElemType: Group{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetGroups returns the value of the Groups field in ListGroupsResponse as
// a slice of Group values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListGroupsResponse) GetGroups(ctx context.Context) ([]Group, bool) {
	if m.Groups.IsNull() || m.Groups.IsUnknown() {
		return nil, false
	}
	var v []Group
	d := m.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in ListGroupsResponse.
func (m *ListGroupsResponse) SetGroups(ctx context.Context, v []Group) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Groups = types.ListValueMust(t, vs)
}

type ListServicePrincipalsProxyRequest struct {
	// Optional. Allows filtering service principals by application id or
	// external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of SPs to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsProxyRequest) {
}

func (to *ListServicePrincipalsProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsProxyRequest) {
}

func (m ListServicePrincipalsProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalsProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsProxyRequest
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListServicePrincipalsRequest struct {
	// Optional. Allows filtering service principals by application id or
	// external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of service principals to return. The service may
	// return fewer than this value. If not provided, defaults to 1000, which is
	// also the maximum allowed. Requests for more than the maximum are clamped
	// to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsRequest) {
}

func (to *ListServicePrincipalsRequest) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsRequest) {
}

func (m ListServicePrincipalsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsRequest
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response message containing a page of service principals in the account.
type ListServicePrincipalsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	ServicePrincipals types.List `tfsdk:"service_principals"`
}

func (to *ListServicePrincipalsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsResponse) {
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() && to.ServicePrincipals.IsNull() && len(from.ServicePrincipals.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServicePrincipals, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServicePrincipals = from.ServicePrincipals
	}
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() {
		if toServicePrincipals, ok := to.GetServicePrincipals(ctx); ok {
			if fromServicePrincipals, ok := from.GetServicePrincipals(ctx); ok {
				// Recursively sync the fields of each ServicePrincipals element by position.
				for i := range toServicePrincipals {
					if i < len(fromServicePrincipals) {
						toServicePrincipals[i].SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipals[i])
					}
				}
				to.SetServicePrincipals(ctx, toServicePrincipals)
			}
		}
	}
}

func (to *ListServicePrincipalsResponse) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsResponse) {
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() && to.ServicePrincipals.IsNull() && len(from.ServicePrincipals.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServicePrincipals, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServicePrincipals = from.ServicePrincipals
	}
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() {
		if toServicePrincipals, ok := to.GetServicePrincipals(ctx); ok {
			if fromServicePrincipals, ok := from.GetServicePrincipals(ctx); ok {
				for i := range toServicePrincipals {
					if i < len(fromServicePrincipals) {
						toServicePrincipals[i].SyncFieldsDuringRead(ctx, fromServicePrincipals[i])
					}
				}
				to.SetServicePrincipals(ctx, toServicePrincipals)
			}
		}
	}
}

func (m ListServicePrincipalsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["service_principals"] = attrs["service_principals"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principals": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsResponse
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":    m.NextPageToken,
			"service_principals": m.ServicePrincipals,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"service_principals": basetypes.ListType{
				ElemType: ServicePrincipal{}.Type(ctx),
			},
		},
	}
}

// GetServicePrincipals returns the value of the ServicePrincipals field in ListServicePrincipalsResponse as
// a slice of ServicePrincipal values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListServicePrincipalsResponse) GetServicePrincipals(ctx context.Context) ([]ServicePrincipal, bool) {
	if m.ServicePrincipals.IsNull() || m.ServicePrincipals.IsUnknown() {
		return nil, false
	}
	var v []ServicePrincipal
	d := m.ServicePrincipals.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipals sets the value of the ServicePrincipals field in ListServicePrincipalsResponse.
func (m *ListServicePrincipalsResponse) SetServicePrincipals(ctx context.Context, v []ServicePrincipal) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principals"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServicePrincipals = types.ListValueMust(t, vs)
}

type ListTransitiveParentGroupsProxyRequest struct {
	// The maximum number of parent groups to return. The service may return
	// fewer than this value. If not provided, defaults to 1000, which is also
	// the maximum allowed. Requests for more than the maximum are clamped to
	// 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListTransitiveParentGroups call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// Required. Internal ID of the principal in Databricks whose transitive
	// parent groups are being listed.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *ListTransitiveParentGroupsProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsProxyRequest) {
}

func (to *ListTransitiveParentGroupsProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsProxyRequest) {
}

func (m ListTransitiveParentGroupsProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitiveParentGroupsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTransitiveParentGroupsProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsProxyRequest
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"principal_id": types.Int64Type,
		},
	}
}

type ListTransitiveParentGroupsRequest struct {
	// The maximum number of parent groups to return. The service may return
	// fewer than this value. If not provided, defaults to 1000, which is also
	// the maximum allowed. Requests for more than the maximum are clamped to
	// 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListTransitiveParentGroups call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// Required. Internal ID of the principal in Databricks whose transitive
	// parent groups are being listed.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *ListTransitiveParentGroupsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsRequest) {
}

func (to *ListTransitiveParentGroupsRequest) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsRequest) {
}

func (m ListTransitiveParentGroupsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitiveParentGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTransitiveParentGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsRequest
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"principal_id": types.Int64Type,
		},
	}
}

// Response message for listing all transitive parent groups of a principal.
type ListTransitiveParentGroupsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of transitive parent groups.
	TransitiveParentGroups types.List `tfsdk:"transitive_parent_groups"`
}

func (to *ListTransitiveParentGroupsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsResponse) {
	if !from.TransitiveParentGroups.IsNull() && !from.TransitiveParentGroups.IsUnknown() && to.TransitiveParentGroups.IsNull() && len(from.TransitiveParentGroups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TransitiveParentGroups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TransitiveParentGroups = from.TransitiveParentGroups
	}
	if !from.TransitiveParentGroups.IsNull() && !from.TransitiveParentGroups.IsUnknown() {
		if toTransitiveParentGroups, ok := to.GetTransitiveParentGroups(ctx); ok {
			if fromTransitiveParentGroups, ok := from.GetTransitiveParentGroups(ctx); ok {
				// Recursively sync the fields of each TransitiveParentGroups element by position.
				for i := range toTransitiveParentGroups {
					if i < len(fromTransitiveParentGroups) {
						toTransitiveParentGroups[i].SyncFieldsDuringCreateOrUpdate(ctx, fromTransitiveParentGroups[i])
					}
				}
				to.SetTransitiveParentGroups(ctx, toTransitiveParentGroups)
			}
		}
	}
}

func (to *ListTransitiveParentGroupsResponse) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsResponse) {
	if !from.TransitiveParentGroups.IsNull() && !from.TransitiveParentGroups.IsUnknown() && to.TransitiveParentGroups.IsNull() && len(from.TransitiveParentGroups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TransitiveParentGroups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TransitiveParentGroups = from.TransitiveParentGroups
	}
	if !from.TransitiveParentGroups.IsNull() && !from.TransitiveParentGroups.IsUnknown() {
		if toTransitiveParentGroups, ok := to.GetTransitiveParentGroups(ctx); ok {
			if fromTransitiveParentGroups, ok := from.GetTransitiveParentGroups(ctx); ok {
				for i := range toTransitiveParentGroups {
					if i < len(fromTransitiveParentGroups) {
						toTransitiveParentGroups[i].SyncFieldsDuringRead(ctx, fromTransitiveParentGroups[i])
					}
				}
				to.SetTransitiveParentGroups(ctx, toTransitiveParentGroups)
			}
		}
	}
}

func (m ListTransitiveParentGroupsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["transitive_parent_groups"] = attrs["transitive_parent_groups"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitiveParentGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTransitiveParentGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"transitive_parent_groups": reflect.TypeOf(TransitiveParentGroup{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsResponse
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"transitive_parent_groups": m.TransitiveParentGroups,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"transitive_parent_groups": basetypes.ListType{
				ElemType: TransitiveParentGroup{}.Type(ctx),
			},
		},
	}
}

// GetTransitiveParentGroups returns the value of the TransitiveParentGroups field in ListTransitiveParentGroupsResponse as
// a slice of TransitiveParentGroup values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTransitiveParentGroupsResponse) GetTransitiveParentGroups(ctx context.Context) ([]TransitiveParentGroup, bool) {
	if m.TransitiveParentGroups.IsNull() || m.TransitiveParentGroups.IsUnknown() {
		return nil, false
	}
	var v []TransitiveParentGroup
	d := m.TransitiveParentGroups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTransitiveParentGroups sets the value of the TransitiveParentGroups field in ListTransitiveParentGroupsResponse.
func (m *ListTransitiveParentGroupsResponse) SetTransitiveParentGroups(ctx context.Context, v []TransitiveParentGroup) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["transitive_parent_groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TransitiveParentGroups = types.ListValueMust(t, vs)
}

type ListUsersProxyRequest struct {
	// Optional. Allows filtering users by username or external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of users to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListUsersProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersProxyRequest) {
}

func (to *ListUsersProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListUsersProxyRequest) {
}

func (m ListUsersProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListUsersProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersProxyRequest
// only implements ToObjectValue() and Type().
func (m ListUsersProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListUsersRequest struct {
	// Optional. Allows filtering users by username or external id.
	Filter types.String `tfsdk:"-"`
	// The maximum number of users to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListUsersRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersRequest) {
}

func (to *ListUsersRequest) SyncFieldsDuringRead(ctx context.Context, from ListUsersRequest) {
}

func (m ListUsersRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersRequest
// only implements ToObjectValue() and Type().
func (m ListUsersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListUsersResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Users types.List `tfsdk:"users"`
}

func (to *ListUsersResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersResponse) {
	if !from.Users.IsNull() && !from.Users.IsUnknown() && to.Users.IsNull() && len(from.Users.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Users, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Users = from.Users
	}
	if !from.Users.IsNull() && !from.Users.IsUnknown() {
		if toUsers, ok := to.GetUsers(ctx); ok {
			if fromUsers, ok := from.GetUsers(ctx); ok {
				// Recursively sync the fields of each Users element by position.
				for i := range toUsers {
					if i < len(fromUsers) {
						toUsers[i].SyncFieldsDuringCreateOrUpdate(ctx, fromUsers[i])
					}
				}
				to.SetUsers(ctx, toUsers)
			}
		}
	}
}

func (to *ListUsersResponse) SyncFieldsDuringRead(ctx context.Context, from ListUsersResponse) {
	if !from.Users.IsNull() && !from.Users.IsUnknown() && to.Users.IsNull() && len(from.Users.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Users, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Users = from.Users
	}
	if !from.Users.IsNull() && !from.Users.IsUnknown() {
		if toUsers, ok := to.GetUsers(ctx); ok {
			if fromUsers, ok := from.GetUsers(ctx); ok {
				for i := range toUsers {
					if i < len(fromUsers) {
						toUsers[i].SyncFieldsDuringRead(ctx, fromUsers[i])
					}
				}
				to.SetUsers(ctx, toUsers)
			}
		}
	}
}

func (m ListUsersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["users"] = attrs["users"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListUsersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"users": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersResponse
// only implements ToObjectValue() and Type().
func (m ListUsersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"users":           m.Users,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"users": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
		},
	}
}

// GetUsers returns the value of the Users field in ListUsersResponse as
// a slice of User values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListUsersResponse) GetUsers(ctx context.Context) ([]User, bool) {
	if m.Users.IsNull() || m.Users.IsUnknown() {
		return nil, false
	}
	var v []User
	d := m.Users.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsers sets the value of the Users field in ListUsersResponse.
func (m *ListUsersResponse) SetUsers(ctx context.Context, v []User) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["users"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Users = types.ListValueMust(t, vs)
}

type ListWorkspaceAssignmentDetailsProxyRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value. If not provided, defaults to 1000,
	// which is also the maximum allowed. Requests for more than the maximum are
	// clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentDetailsProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsProxyRequest) {
}

func (to *ListWorkspaceAssignmentDetailsProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsProxyRequest) {
}

func (m ListWorkspaceAssignmentDetailsProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentDetailsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentDetailsProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsProxyRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceAssignmentDetailsRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value. If not provided, defaults to 1000,
	// which is also the maximum allowed. Requests for more than the maximum are
	// clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListWorkspaceAssignmentDetails
	// call. Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// Required. The workspace ID for which the workspace assignment details are
	// being fetched.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentDetailsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsRequest) {
}

func (to *ListWorkspaceAssignmentDetailsRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsRequest) {
}

func (m ListWorkspaceAssignmentDetailsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentDetailsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentDetailsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Response message for listing workspace assignment details.
type ListWorkspaceAssignmentDetailsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	WorkspaceAssignmentDetails types.List `tfsdk:"workspace_assignment_details"`
}

func (to *ListWorkspaceAssignmentDetailsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsResponse) {
	if !from.WorkspaceAssignmentDetails.IsNull() && !from.WorkspaceAssignmentDetails.IsUnknown() && to.WorkspaceAssignmentDetails.IsNull() && len(from.WorkspaceAssignmentDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAssignmentDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAssignmentDetails = from.WorkspaceAssignmentDetails
	}
	if !from.WorkspaceAssignmentDetails.IsNull() && !from.WorkspaceAssignmentDetails.IsUnknown() {
		if toWorkspaceAssignmentDetails, ok := to.GetWorkspaceAssignmentDetails(ctx); ok {
			if fromWorkspaceAssignmentDetails, ok := from.GetWorkspaceAssignmentDetails(ctx); ok {
				// Recursively sync the fields of each WorkspaceAssignmentDetails element by position.
				for i := range toWorkspaceAssignmentDetails {
					if i < len(fromWorkspaceAssignmentDetails) {
						toWorkspaceAssignmentDetails[i].SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignmentDetails[i])
					}
				}
				to.SetWorkspaceAssignmentDetails(ctx, toWorkspaceAssignmentDetails)
			}
		}
	}
}

func (to *ListWorkspaceAssignmentDetailsResponse) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsResponse) {
	if !from.WorkspaceAssignmentDetails.IsNull() && !from.WorkspaceAssignmentDetails.IsUnknown() && to.WorkspaceAssignmentDetails.IsNull() && len(from.WorkspaceAssignmentDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAssignmentDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAssignmentDetails = from.WorkspaceAssignmentDetails
	}
	if !from.WorkspaceAssignmentDetails.IsNull() && !from.WorkspaceAssignmentDetails.IsUnknown() {
		if toWorkspaceAssignmentDetails, ok := to.GetWorkspaceAssignmentDetails(ctx); ok {
			if fromWorkspaceAssignmentDetails, ok := from.GetWorkspaceAssignmentDetails(ctx); ok {
				for i := range toWorkspaceAssignmentDetails {
					if i < len(fromWorkspaceAssignmentDetails) {
						toWorkspaceAssignmentDetails[i].SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetails[i])
					}
				}
				to.SetWorkspaceAssignmentDetails(ctx, toWorkspaceAssignmentDetails)
			}
		}
	}
}

func (m ListWorkspaceAssignmentDetailsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["workspace_assignment_details"] = attrs["workspace_assignment_details"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentDetailsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentDetailsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_details": reflect.TypeOf(WorkspaceAssignmentDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsResponse
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":              m.NextPageToken,
			"workspace_assignment_details": m.WorkspaceAssignmentDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_assignment_details": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignmentDetails returns the value of the WorkspaceAssignmentDetails field in ListWorkspaceAssignmentDetailsResponse as
// a slice of WorkspaceAssignmentDetail values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAssignmentDetailsResponse) GetWorkspaceAssignmentDetails(ctx context.Context) ([]WorkspaceAssignmentDetail, bool) {
	if m.WorkspaceAssignmentDetails.IsNull() || m.WorkspaceAssignmentDetails.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAssignmentDetail
	d := m.WorkspaceAssignmentDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetails sets the value of the WorkspaceAssignmentDetails field in ListWorkspaceAssignmentDetailsResponse.
func (m *ListWorkspaceAssignmentDetailsResponse) SetWorkspaceAssignmentDetails(ctx context.Context, v []WorkspaceAssignmentDetail) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceAssignmentDetails = types.ListValueMust(t, vs)
}

type ListWorkspaceAssignmentsProxyRequest struct {
	// The maximum number of workspace assignments to return. The service may
	// return fewer than this value. If not provided, defaults to 1000, which is
	// also the maximum allowed. Requests for more than the maximum are clamped
	// to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentsProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsProxyRequest) {
}

func (to *ListWorkspaceAssignmentsProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsProxyRequest) {
}

func (m ListWorkspaceAssignmentsProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentsProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsProxyRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceAssignmentsRequest struct {
	// The maximum number of workspace assignments to return. The service may
	// return fewer than this value. If not provided, defaults to 1000, which is
	// also the maximum allowed. Requests for more than the maximum are clamped
	// to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListWorkspaceAssignments call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// Required. The workspace ID for which the workspace assignments are being
	// fetched.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsRequest) {
}

func (to *ListWorkspaceAssignmentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsRequest) {
}

func (m ListWorkspaceAssignmentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Response message for listing workspace assignments.
type ListWorkspaceAssignmentsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	WorkspaceAssignments types.List `tfsdk:"workspace_assignments"`
}

func (to *ListWorkspaceAssignmentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsResponse) {
	if !from.WorkspaceAssignments.IsNull() && !from.WorkspaceAssignments.IsUnknown() && to.WorkspaceAssignments.IsNull() && len(from.WorkspaceAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAssignments = from.WorkspaceAssignments
	}
	if !from.WorkspaceAssignments.IsNull() && !from.WorkspaceAssignments.IsUnknown() {
		if toWorkspaceAssignments, ok := to.GetWorkspaceAssignments(ctx); ok {
			if fromWorkspaceAssignments, ok := from.GetWorkspaceAssignments(ctx); ok {
				// Recursively sync the fields of each WorkspaceAssignments element by position.
				for i := range toWorkspaceAssignments {
					if i < len(fromWorkspaceAssignments) {
						toWorkspaceAssignments[i].SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignments[i])
					}
				}
				to.SetWorkspaceAssignments(ctx, toWorkspaceAssignments)
			}
		}
	}
}

func (to *ListWorkspaceAssignmentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsResponse) {
	if !from.WorkspaceAssignments.IsNull() && !from.WorkspaceAssignments.IsUnknown() && to.WorkspaceAssignments.IsNull() && len(from.WorkspaceAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAssignments = from.WorkspaceAssignments
	}
	if !from.WorkspaceAssignments.IsNull() && !from.WorkspaceAssignments.IsUnknown() {
		if toWorkspaceAssignments, ok := to.GetWorkspaceAssignments(ctx); ok {
			if fromWorkspaceAssignments, ok := from.GetWorkspaceAssignments(ctx); ok {
				for i := range toWorkspaceAssignments {
					if i < len(fromWorkspaceAssignments) {
						toWorkspaceAssignments[i].SyncFieldsDuringRead(ctx, fromWorkspaceAssignments[i])
					}
				}
				to.SetWorkspaceAssignments(ctx, toWorkspaceAssignments)
			}
		}
	}
}

func (m ListWorkspaceAssignmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["workspace_assignments"] = attrs["workspace_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignments": reflect.TypeOf(WorkspaceAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsResponse
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":       m.NextPageToken,
			"workspace_assignments": m.WorkspaceAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_assignments": basetypes.ListType{
				ElemType: WorkspaceAssignment{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignments returns the value of the WorkspaceAssignments field in ListWorkspaceAssignmentsResponse as
// a slice of WorkspaceAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAssignmentsResponse) GetWorkspaceAssignments(ctx context.Context) ([]WorkspaceAssignment, bool) {
	if m.WorkspaceAssignments.IsNull() || m.WorkspaceAssignments.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAssignment
	d := m.WorkspaceAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignments sets the value of the WorkspaceAssignments field in ListWorkspaceAssignmentsResponse.
func (m *ListWorkspaceAssignmentsResponse) SetWorkspaceAssignments(ctx context.Context, v []WorkspaceAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceAssignments = types.ListValueMust(t, vs)
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// name and inherited parent groups.
type ResolveGroupProxyRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupProxyRequest) {
}

func (to *ResolveGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupProxyRequest) {
}

func (m ResolveGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m ResolveGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// name and inherited parent groups.
type ResolveGroupRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupRequest) {
}

func (to *ResolveGroupRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupRequest) {
}

func (m ResolveGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupRequest
// only implements ToObjectValue() and Type().
func (m ResolveGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupRequest) Type(ctx context.Context) attr.Type {
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

func (to *ResolveGroupResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupResponse) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				// Recursively sync the fields of Group
				toGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (to *ResolveGroupResponse) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupResponse) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m ResolveGroupResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveGroupResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupResponse
// only implements ToObjectValue() and Type().
func (m ResolveGroupResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": Group{}.Type(ctx),
		},
	}
}

// GetGroup returns the value of the Group field in ResolveGroupResponse as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResolveGroupResponse) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := m.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in ResolveGroupResponse.
func (m *ResolveGroupResponse) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	m.Group = vs
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's display name, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalProxyRequest) {
}

func (to *ResolveServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalProxyRequest) {
}

func (m ResolveServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's display name, status, and inherited parent groups.
type ResolveServicePrincipalRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalRequest) {
}

func (to *ResolveServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalRequest) {
}

func (m ResolveServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalRequest) Type(ctx context.Context) attr.Type {
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

func (to *ResolveServicePrincipalResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalResponse) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				// Recursively sync the fields of ServicePrincipal
				toServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (to *ResolveServicePrincipalResponse) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalResponse) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m ResolveServicePrincipalResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalResponse
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": ServicePrincipal{}.Type(ctx),
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in ResolveServicePrincipalResponse as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResolveServicePrincipalResponse) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := m.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in ResolveServicePrincipalResponse.
func (m *ResolveServicePrincipalResponse) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	m.ServicePrincipal = vs
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// display name, status, and inherited parent groups.
type ResolveUserProxyRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserProxyRequest) {
}

func (to *ResolveUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveUserProxyRequest) {
}

func (m ResolveUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserProxyRequest
// only implements ToObjectValue() and Type().
func (m ResolveUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// display name, status, and inherited parent groups.
type ResolveUserRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserRequest) {
}

func (to *ResolveUserRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveUserRequest) {
}

func (m ResolveUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserRequest
// only implements ToObjectValue() and Type().
func (m ResolveUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserRequest) Type(ctx context.Context) attr.Type {
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

func (to *ResolveUserResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserResponse) {
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

func (to *ResolveUserResponse) SyncFieldsDuringRead(ctx context.Context, from ResolveUserResponse) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m ResolveUserResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveUserResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserResponse
// only implements ToObjectValue() and Type().
func (m ResolveUserResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": User{}.Type(ctx),
		},
	}
}

// GetUser returns the value of the User field in ResolveUserResponse as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResolveUserResponse) GetUser(ctx context.Context) (User, bool) {
	var e User
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v User
	d := m.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in ResolveUserResponse.
func (m *ResolveUserResponse) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	m.User = vs
}

// The details of a ServicePrincipal resource.
type ServicePrincipal struct {
	// The parent account ID for the service principal in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a service principal in a Databricks account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// Application ID of the service principal. Set at creation time and cannot
	// be changed afterwards; when omitted, the server generates one.
	ApplicationId types.String `tfsdk:"application_id"`
	// Display name of the service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// ExternalId of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal service principal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"service_principal_id"`
}

func (to *ServicePrincipal) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServicePrincipal) {
}

func (to *ServicePrincipal) SyncFieldsDuringRead(ctx context.Context, from ServicePrincipal) {
}

func (m ServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_sp_status"] = attrs["account_sp_status"].SetRequired()
	attrs["application_id"] = attrs["application_id"].SetComputed()
	attrs["application_id"] = attrs["application_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal
// only implements ToObjectValue() and Type().
func (m ServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":           m.AccountId,
			"account_sp_status":    m.AccountSpStatus,
			"application_id":       m.ApplicationId,
			"display_name":         m.DisplayName,
			"external_id":          m.ExternalId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":           types.StringType,
			"account_sp_status":    types.StringType,
			"application_id":       types.StringType,
			"display_name":         types.StringType,
			"external_id":          types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

// Represents a group that is a transitive parent of a principal.
type TransitiveParentGroup struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal group ID of the group in Databricks.
	GroupId types.String `tfsdk:"group_id"`
}

func (to *TransitiveParentGroup) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitiveParentGroup) {
}

func (to *TransitiveParentGroup) SyncFieldsDuringRead(ctx context.Context, from TransitiveParentGroup) {
}

func (m TransitiveParentGroup) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitiveParentGroup.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitiveParentGroup) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitiveParentGroup
// only implements ToObjectValue() and Type().
func (m TransitiveParentGroup) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":  m.AccountId,
			"external_id": m.ExternalId,
			"group_id":    m.GroupId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitiveParentGroup) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_id":    types.StringType,
		},
	}
}

type UpdateGroupProxyRequest struct {
	// Required. Group to be updated in <Databricks>
	Group types.Object `tfsdk:"group"`
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateGroupProxyRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				// Recursively sync the fields of Group
				toGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (to *UpdateGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateGroupProxyRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m UpdateGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateGroupProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateGroupProxyRequest
// only implements ToObjectValue() and Type().
func (m UpdateGroupProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group":       m.Group,
			"group_id":    m.GroupId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group":       Group{}.Type(ctx),
			"group_id":    types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetGroup returns the value of the Group field in UpdateGroupProxyRequest as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateGroupProxyRequest) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := m.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in UpdateGroupProxyRequest.
func (m *UpdateGroupProxyRequest) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	m.Group = vs
}

type UpdateGroupRequest struct {
	// Required. Group to be updated in <Databricks>
	Group types.Object `tfsdk:"group"`
	// Required. Internal ID of the group in Databricks.
	GroupId types.String `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateGroupRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				// Recursively sync the fields of Group
				toGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (to *UpdateGroupRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateGroupRequest) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m UpdateGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateGroupRequest
// only implements ToObjectValue() and Type().
func (m UpdateGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group":       m.Group,
			"group_id":    m.GroupId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group":       Group{}.Type(ctx),
			"group_id":    types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetGroup returns the value of the Group field in UpdateGroupRequest as
// a Group value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateGroupRequest) GetGroup(ctx context.Context) (Group, bool) {
	var e Group
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v Group
	d := m.Group.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroup sets the value of the Group field in UpdateGroupRequest.
func (m *UpdateGroupRequest) SetGroup(ctx context.Context, v Group) {
	vs := v.ToObjectValue(ctx)
	m.Group = vs
}

type UpdateServicePrincipalProxyRequest struct {
	// Required. Service principal to be updated in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalProxyRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				// Recursively sync the fields of ServicePrincipal
				toServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (to *UpdateServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalProxyRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m UpdateServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateServicePrincipalProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalProxyRequest
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal":    m.ServicePrincipal,
			"service_principal_id": m.ServicePrincipalId,
			"update_mask":          m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal":    ServicePrincipal{}.Type(ctx),
			"service_principal_id": types.StringType,
			"update_mask":          types.StringType,
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in UpdateServicePrincipalProxyRequest as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateServicePrincipalProxyRequest) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := m.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in UpdateServicePrincipalProxyRequest.
func (m *UpdateServicePrincipalProxyRequest) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	m.ServicePrincipal = vs
}

type UpdateServicePrincipalRequest struct {
	// Required. Service Principal to be updated in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId types.String `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				// Recursively sync the fields of ServicePrincipal
				toServicePrincipal.SyncFieldsDuringCreateOrUpdate(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (to *UpdateServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalRequest) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m UpdateServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalRequest
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal":    m.ServicePrincipal,
			"service_principal_id": m.ServicePrincipalId,
			"update_mask":          m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal":    ServicePrincipal{}.Type(ctx),
			"service_principal_id": types.StringType,
			"update_mask":          types.StringType,
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in UpdateServicePrincipalRequest as
// a ServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateServicePrincipalRequest) GetServicePrincipal(ctx context.Context) (ServicePrincipal, bool) {
	var e ServicePrincipal
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v ServicePrincipal
	d := m.ServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in UpdateServicePrincipalRequest.
func (m *UpdateServicePrincipalRequest) SetServicePrincipal(ctx context.Context, v ServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	m.ServicePrincipal = vs
}

type UpdateUserProxyRequest struct {
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.Object `tfsdk:"user"`
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *UpdateUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateUserProxyRequest) {
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

func (to *UpdateUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateUserProxyRequest) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m UpdateUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateUserProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateUserProxyRequest
// only implements ToObjectValue() and Type().
func (m UpdateUserProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_mask": m.UpdateMask,
			"user":        m.User,
			"user_id":     m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update_mask": types.StringType,
			"user":        User{}.Type(ctx),
			"user_id":     types.StringType,
		},
	}
}

// GetUser returns the value of the User field in UpdateUserProxyRequest as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateUserProxyRequest) GetUser(ctx context.Context) (User, bool) {
	var e User
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v User
	d := m.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in UpdateUserProxyRequest.
func (m *UpdateUserProxyRequest) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	m.User = vs
}

type UpdateUserRequest struct {
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.Object `tfsdk:"user"`
	// Required. Internal ID of the user in Databricks.
	UserId types.String `tfsdk:"-"`
}

func (to *UpdateUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateUserRequest) {
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

func (to *UpdateUserRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateUserRequest) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m UpdateUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateUserRequest
// only implements ToObjectValue() and Type().
func (m UpdateUserRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_mask": m.UpdateMask,
			"user":        m.User,
			"user_id":     m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update_mask": types.StringType,
			"user":        User{}.Type(ctx),
			"user_id":     types.StringType,
		},
	}
}

// GetUser returns the value of the User field in UpdateUserRequest as
// a User value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateUserRequest) GetUser(ctx context.Context) (User, bool) {
	var e User
	if m.User.IsNull() || m.User.IsUnknown() {
		return e, false
	}
	var v User
	d := m.User.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUser sets the value of the User field in UpdateUserRequest.
func (m *UpdateUserRequest) SetUser(ctx context.Context, v User) {
	vs := v.ToObjectValue(ctx)
	m.User = vs
}

type UpdateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
	WorkspaceAssignmentDetail types.Object `tfsdk:"workspace_assignment_detail"`
}

func (to *UpdateWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentDetailProxyRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignmentDetail
				toWorkspaceAssignmentDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (to *UpdateWorkspaceAssignmentDetailProxyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentDetailProxyRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentDetailProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignmentDetailProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAssignmentDetailProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentDetailProxyRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentDetailProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":                m.PrincipalId,
			"update_mask":                 m.UpdateMask,
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentDetailProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":                types.Int64Type,
			"update_mask":                 types.StringType,
			"workspace_assignment_detail": WorkspaceAssignmentDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailProxyRequest as
// a WorkspaceAssignmentDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentDetailProxyRequest) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail, bool) {
	var e WorkspaceAssignmentDetail
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignmentDetail
	d := m.WorkspaceAssignmentDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailProxyRequest.
func (m *UpdateWorkspaceAssignmentDetailProxyRequest) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignmentDetail = vs
}

type UpdateWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
	WorkspaceAssignmentDetail types.Object `tfsdk:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being updated.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceAssignmentDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentDetailRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignmentDetail
				toWorkspaceAssignmentDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (to *UpdateWorkspaceAssignmentDetailRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentDetailRequest) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignmentDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAssignmentDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentDetailRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":                m.PrincipalId,
			"update_mask":                 m.UpdateMask,
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
			"workspace_id":                m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":                types.Int64Type,
			"update_mask":                 types.StringType,
			"workspace_assignment_detail": WorkspaceAssignmentDetail{}.Type(ctx),
			"workspace_id":                types.Int64Type,
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailRequest as
// a WorkspaceAssignmentDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentDetailRequest) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail, bool) {
	var e WorkspaceAssignmentDetail
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignmentDetail
	d := m.WorkspaceAssignmentDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailRequest.
func (m *UpdateWorkspaceAssignmentDetailRequest) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignmentDetail = vs
}

type UpdateWorkspaceAssignmentProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment to be updated in <Databricks>.
	WorkspaceAssignment types.Object `tfsdk:"workspace_assignment"`
}

func (to *UpdateWorkspaceAssignmentProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentProxyRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignment
				toWorkspaceAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (to *UpdateWorkspaceAssignmentProxyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentProxyRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAssignmentProxyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentProxyRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentProxyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":         m.PrincipalId,
			"update_mask":          m.UpdateMask,
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":         types.Int64Type,
			"update_mask":          types.StringType,
			"workspace_assignment": WorkspaceAssignment{}.Type(ctx),
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentProxyRequest as
// a WorkspaceAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentProxyRequest) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment, bool) {
	var e WorkspaceAssignment
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignment
	d := m.WorkspaceAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentProxyRequest.
func (m *UpdateWorkspaceAssignmentProxyRequest) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignment = vs
}

type UpdateWorkspaceAssignmentRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment to be updated in <Databricks>.
	WorkspaceAssignment types.Object `tfsdk:"workspace_assignment"`
	// Required. The workspace ID for which the workspace assignment is being
	// updated.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				// Recursively sync the fields of WorkspaceAssignment
				toWorkspaceAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (to *UpdateWorkspaceAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentRequest) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":         m.PrincipalId,
			"update_mask":          m.UpdateMask,
			"workspace_assignment": m.WorkspaceAssignment,
			"workspace_id":         m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":         types.Int64Type,
			"update_mask":          types.StringType,
			"workspace_assignment": WorkspaceAssignment{}.Type(ctx),
			"workspace_id":         types.Int64Type,
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentRequest as
// a WorkspaceAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentRequest) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment, bool) {
	var e WorkspaceAssignment
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v WorkspaceAssignment
	d := m.WorkspaceAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentRequest.
func (m *UpdateWorkspaceAssignmentRequest) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAssignment = vs
}

type UpdateWorkspaceIdentityDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace identity detail to be updated in <Databricks>.
	WorkspaceIdentityDetail types.Object `tfsdk:"workspace_identity_detail"`
}

func (to *UpdateWorkspaceIdentityDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceIdentityDetailRequest) {
	if !from.WorkspaceIdentityDetail.IsNull() && !from.WorkspaceIdentityDetail.IsUnknown() {
		if toWorkspaceIdentityDetail, ok := to.GetWorkspaceIdentityDetail(ctx); ok {
			if fromWorkspaceIdentityDetail, ok := from.GetWorkspaceIdentityDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceIdentityDetail
				toWorkspaceIdentityDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceIdentityDetail)
				to.SetWorkspaceIdentityDetail(ctx, toWorkspaceIdentityDetail)
			}
		}
	}
}

func (to *UpdateWorkspaceIdentityDetailRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceIdentityDetailRequest) {
	if !from.WorkspaceIdentityDetail.IsNull() && !from.WorkspaceIdentityDetail.IsUnknown() {
		if toWorkspaceIdentityDetail, ok := to.GetWorkspaceIdentityDetail(ctx); ok {
			if fromWorkspaceIdentityDetail, ok := from.GetWorkspaceIdentityDetail(ctx); ok {
				toWorkspaceIdentityDetail.SyncFieldsDuringRead(ctx, fromWorkspaceIdentityDetail)
				to.SetWorkspaceIdentityDetail(ctx, toWorkspaceIdentityDetail)
			}
		}
	}
}

func (m UpdateWorkspaceIdentityDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_identity_detail"] = attrs["workspace_identity_detail"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceIdentityDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceIdentityDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_identity_detail": reflect.TypeOf(WorkspaceIdentityDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceIdentityDetailRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceIdentityDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":              m.PrincipalId,
			"update_mask":               m.UpdateMask,
			"workspace_identity_detail": m.WorkspaceIdentityDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceIdentityDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":              types.Int64Type,
			"update_mask":               types.StringType,
			"workspace_identity_detail": WorkspaceIdentityDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceIdentityDetail returns the value of the WorkspaceIdentityDetail field in UpdateWorkspaceIdentityDetailRequest as
// a WorkspaceIdentityDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceIdentityDetailRequest) GetWorkspaceIdentityDetail(ctx context.Context) (WorkspaceIdentityDetail, bool) {
	var e WorkspaceIdentityDetail
	if m.WorkspaceIdentityDetail.IsNull() || m.WorkspaceIdentityDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceIdentityDetail
	d := m.WorkspaceIdentityDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIdentityDetail sets the value of the WorkspaceIdentityDetail field in UpdateWorkspaceIdentityDetailRequest.
func (m *UpdateWorkspaceIdentityDetailRequest) SetWorkspaceIdentityDetail(ctx context.Context, v WorkspaceIdentityDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceIdentityDetail = vs
}

// The details of a User resource.
type User struct {
	// The accountId parent of the user in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// ExternalId of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`

	FullName types.Object `tfsdk:"full_name"`
	// Internal userId of the user in Databricks.
	UserId types.String `tfsdk:"user_id"`
	// Username/email of the user.
	Username types.String `tfsdk:"username"`
}

func (to *User) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				// Recursively sync the fields of FullName
				toFullName.SyncFieldsDuringCreateOrUpdate(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (to *User) SyncFieldsDuringRead(ctx context.Context, from User) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				toFullName.SyncFieldsDuringRead(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (m User) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetRequired()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["user_id"] = attrs["user_id"].SetComputed()
	attrs["username"] = attrs["username"].SetRequired()
	attrs["username"] = attrs["username"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in User.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name": reflect.TypeOf(UserFullName{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User
// only implements ToObjectValue() and Type().
func (m User) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":          m.AccountId,
			"account_user_status": m.AccountUserStatus,
			"external_id":         m.ExternalId,
			"full_name":           m.FullName,
			"user_id":             m.UserId,
			"username":            m.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (m User) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"external_id":         types.StringType,
			"full_name":           UserFullName{}.Type(ctx),
			"user_id":             types.StringType,
			"username":            types.StringType,
		},
	}
}

// GetFullName returns the value of the FullName field in User as
// a UserFullName value.
// If the field is unknown or null, the boolean return value is false.
func (m *User) GetFullName(ctx context.Context) (UserFullName, bool) {
	var e UserFullName
	if m.FullName.IsNull() || m.FullName.IsUnknown() {
		return e, false
	}
	var v UserFullName
	d := m.FullName.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullName sets the value of the FullName field in User.
func (m *User) SetFullName(ctx context.Context, v UserFullName) {
	vs := v.ToObjectValue(ctx)
	m.FullName = vs
}

type UserFullName struct {
	FamilyName types.String `tfsdk:"family_name"`

	GivenName types.String `tfsdk:"given_name"`
}

func (to *UserFullName) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UserFullName) {
}

func (to *UserFullName) SyncFieldsDuringRead(ctx context.Context, from UserFullName) {
}

func (m UserFullName) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["family_name"] = attrs["family_name"].SetOptional()
	attrs["given_name"] = attrs["given_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UserFullName.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UserFullName) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserFullName
// only implements ToObjectValue() and Type().
func (m UserFullName) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UserFullName) Type(ctx context.Context) attr.Type {
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

func (to *WorkspaceAccessDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAccessDetail) {
	if !from.Permissions.IsNull() && !from.Permissions.IsUnknown() && to.Permissions.IsNull() && len(from.Permissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Permissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Permissions = from.Permissions
	}
}

func (to *WorkspaceAccessDetail) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAccessDetail) {
	if !from.Permissions.IsNull() && !from.Permissions.IsUnknown() && to.Permissions.IsNull() && len(from.Permissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Permissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Permissions = from.Permissions
	}
}

func (m WorkspaceAccessDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceAccessDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAccessDetail
// only implements ToObjectValue() and Type().
func (m WorkspaceAccessDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_type":    m.AccessType,
			"account_id":     m.AccountId,
			"permissions":    m.Permissions,
			"principal_id":   m.PrincipalId,
			"principal_type": m.PrincipalType,
			"status":         m.Status,
			"workspace_id":   m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceAccessDetail) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceAccessDetail) GetPermissions(ctx context.Context) ([]types.String, bool) {
	if m.Permissions.IsNull() || m.Permissions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Permissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissions sets the value of the Permissions field in WorkspaceAccessDetail.
func (m *WorkspaceAccessDetail) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Permissions = types.ListValueMust(t, vs)
}

// The direct assignment of a provisioned account-level principal (user, service
// principal, or group) to a workspace, together with the entitlements that
// assignment grants in the workspace.
//
// This resource covers only principals assigned directly to the workspace.
// Principals that inherit workspace access through a group are not represented
// here. See WorkspaceAccessDetail and WorkspaceIdentityDetail for the
// effective, direct-or-indirect view. Creating the resource assigns the
// principal to the workspace, and deleting it removes the assignment.
//
// `entitlements` is the only client-settable field. It holds the entitlements
// granted directly on this assignment, including any the principal also holds
// through a group. `effective_entitlements` is the read-only union of those and
// any granted through group membership.
//
// A direct assignment always carries at least one directly-assigned
// entitlement, because the assignment is what grants it. Create and update both
// reject an empty `entitlements` set. To remove a principal's assignment
// entirely, delete the resource.
//
// This resource replaces workspace assignment previously managed through the
// workspace SCIM and permission-assignment APIs, and is intended for account
// and workspace admins.
type WorkspaceAssignment struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId types.String `tfsdk:"account_id"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group) that is
	// assigned.
	PrincipalType types.String `tfsdk:"principal_type"`
	// The workspace ID where the principal is assigned.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceAssignment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignment) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (to *WorkspaceAssignment) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignment) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (m WorkspaceAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["effective_entitlements"] = attrs["effective_entitlements"].SetComputed()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignment
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":             m.AccountId,
			"effective_entitlements": m.EffectiveEntitlements,
			"entitlements":           m.Entitlements,
			"principal_id":           m.PrincipalId,
			"principal_type":         m.PrincipalType,
			"workspace_id":           m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"effective_entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

// GetEffectiveEntitlements returns the value of the EffectiveEntitlements field in WorkspaceAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment) GetEffectiveEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.EffectiveEntitlements.IsNull() || m.EffectiveEntitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EffectiveEntitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveEntitlements sets the value of the EffectiveEntitlements field in WorkspaceAssignment.
func (m *WorkspaceAssignment) SetEffectiveEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveEntitlements = types.SetValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in WorkspaceAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment) GetEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.Entitlements.IsNull() || m.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in WorkspaceAssignment.
func (m *WorkspaceAssignment) SetEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entitlements = types.SetValueMust(t, vs)
}

// The direct assignment of a provisioned account-level principal (user, service
// principal, or group) to a workspace, together with the entitlements that
// assignment grants in the workspace.
//
// This resource covers only principals assigned directly to the workspace.
// Principals that inherit workspace access through a group are not represented
// here. See WorkspaceAccessDetail and WorkspaceIdentityDetail for the
// effective, direct-or-indirect view. Creating the resource assigns the
// principal to the workspace, and deleting it removes the assignment.
//
// `entitlements` is the only client-settable field. It holds the entitlements
// granted directly on this assignment, including any the principal also holds
// through a group. `effective_entitlements` is the read-only union of those and
// any granted through group membership.
//
// A direct assignment always carries at least one directly-assigned
// entitlement, because the assignment is what grants it. Create and update both
// reject an empty `entitlements` set. To remove a principal's assignment
// entirely, delete the resource.
//
// This resource replaces workspace assignment previously managed through the
// workspace SCIM and permission-assignment APIs, and is intended for account
// and workspace admins.
type WorkspaceAssignmentDetail struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId types.String `tfsdk:"account_id"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`

	PrincipalType types.String `tfsdk:"principal_type"`
	// The workspace ID where the principal is assigned.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceAssignmentDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignmentDetail) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (to *WorkspaceAssignmentDetail) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignmentDetail) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (m WorkspaceAssignmentDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["effective_entitlements"] = attrs["effective_entitlements"].SetComputed()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceAssignmentDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceAssignmentDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignmentDetail
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignmentDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":             m.AccountId,
			"effective_entitlements": m.EffectiveEntitlements,
			"entitlements":           m.Entitlements,
			"principal_id":           m.PrincipalId,
			"principal_type":         m.PrincipalType,
			"workspace_id":           m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceAssignmentDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"effective_entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

// GetEffectiveEntitlements returns the value of the EffectiveEntitlements field in WorkspaceAssignmentDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignmentDetail) GetEffectiveEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.EffectiveEntitlements.IsNull() || m.EffectiveEntitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EffectiveEntitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveEntitlements sets the value of the EffectiveEntitlements field in WorkspaceAssignmentDetail.
func (m *WorkspaceAssignmentDetail) SetEffectiveEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveEntitlements = types.SetValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in WorkspaceAssignmentDetail as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignmentDetail) GetEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.Entitlements.IsNull() || m.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in WorkspaceAssignmentDetail.
func (m *WorkspaceAssignmentDetail) SetEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entitlements = types.SetValueMust(t, vs)
}

// The details of a directly or indirectly assigned principal's details in a
// workspace.
type WorkspaceIdentityDetail struct {
	// The type of assignment the principal has to the workspace (direct or
	// indirect).
	AssignmentType types.String `tfsdk:"assignment_type"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group).
	PrincipalType types.String `tfsdk:"principal_type"`
	// The activity status of an identity in a Databricks workspace.
	WorkspaceIdentityStatus types.String `tfsdk:"workspace_identity_status"`
}

func (to *WorkspaceIdentityDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceIdentityDetail) {
}

func (to *WorkspaceIdentityDetail) SyncFieldsDuringRead(ctx context.Context, from WorkspaceIdentityDetail) {
}

func (m WorkspaceIdentityDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assignment_type"] = attrs["assignment_type"].SetComputed()
	attrs["principal_id"] = attrs["principal_id"].SetComputed()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_identity_status"] = attrs["workspace_identity_status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceIdentityDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceIdentityDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceIdentityDetail
// only implements ToObjectValue() and Type().
func (m WorkspaceIdentityDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assignment_type":           m.AssignmentType,
			"principal_id":              m.PrincipalId,
			"principal_type":            m.PrincipalType,
			"workspace_identity_status": m.WorkspaceIdentityStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceIdentityDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assignment_type":           types.StringType,
			"principal_id":              types.Int64Type,
			"principal_type":            types.StringType,
			"workspace_identity_status": types.StringType,
		},
	}
}
