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

type CreateGroupProxyRequest_SdkV2 struct {
	// Required. Group to be created in <Databricks>
	Group types.List `tfsdk:"group"`
}

func (to *CreateGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGroupProxyRequest_SdkV2) {
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

func (to *CreateGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateGroupProxyRequest_SdkV2) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m CreateGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["group"] = attrs["group"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGroup returns the value of the Group field in CreateGroupProxyRequest_SdkV2 as
// a Group_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateGroupProxyRequest_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := m.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in CreateGroupProxyRequest_SdkV2.
func (m *CreateGroupProxyRequest_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	m.Group = types.ListValueMust(t, vs)
}

type CreateGroupRequest_SdkV2 struct {
	// Required. Group to be created in <Databricks>
	Group types.List `tfsdk:"group"`
}

func (to *CreateGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGroupRequest_SdkV2) {
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

func (to *CreateGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateGroupRequest_SdkV2) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m CreateGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["group"] = attrs["group"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGroup returns the value of the Group field in CreateGroupRequest_SdkV2 as
// a Group_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateGroupRequest_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := m.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in CreateGroupRequest_SdkV2.
func (m *CreateGroupRequest_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	m.Group = types.ListValueMust(t, vs)
}

type CreateServicePrincipalProxyRequest_SdkV2 struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal types.List `tfsdk:"service_principal"`
}

func (to *CreateServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalProxyRequest_SdkV2) {
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

func (to *CreateServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalProxyRequest_SdkV2) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m CreateServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["service_principal"] = attrs["service_principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in CreateServicePrincipalProxyRequest_SdkV2 as
// a ServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServicePrincipalProxyRequest_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in CreateServicePrincipalProxyRequest_SdkV2.
func (m *CreateServicePrincipalProxyRequest_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	m.ServicePrincipal = types.ListValueMust(t, vs)
}

type CreateServicePrincipalRequest_SdkV2 struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal types.List `tfsdk:"service_principal"`
}

func (to *CreateServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalRequest_SdkV2) {
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

func (to *CreateServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalRequest_SdkV2) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m CreateServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["service_principal"] = attrs["service_principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in CreateServicePrincipalRequest_SdkV2 as
// a ServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServicePrincipalRequest_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in CreateServicePrincipalRequest_SdkV2.
func (m *CreateServicePrincipalRequest_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	m.ServicePrincipal = types.ListValueMust(t, vs)
}

type CreateUserProxyRequest_SdkV2 struct {
	// Required. User to be created in <Databricks>
	User types.List `tfsdk:"user"`
}

func (to *CreateUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateUserProxyRequest_SdkV2) {
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

func (to *CreateUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateUserProxyRequest_SdkV2) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m CreateUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUser returns the value of the User field in CreateUserProxyRequest_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateUserProxyRequest_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in CreateUserProxyRequest_SdkV2.
func (m *CreateUserProxyRequest_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

type CreateUserRequest_SdkV2 struct {
	// Required. User to be created in <Databricks>
	User types.List `tfsdk:"user"`
}

func (to *CreateUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateUserRequest_SdkV2) {
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

func (to *CreateUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateUserRequest_SdkV2) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m CreateUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUser returns the value of the User field in CreateUserRequest_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateUserRequest_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in CreateUserRequest_SdkV2.
func (m *CreateUserRequest_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

type CreateWorkspaceAccessDetailLocalRequest_SdkV2 struct {
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail types.List `tfsdk:"workspace_access_detail"`
}

func (to *CreateWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAccessDetailLocalRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAccessDetail
				toWorkspaceAccessDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (to *CreateWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAccessDetailLocalRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m CreateWorkspaceAccessDetailLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAccessDetailLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAccessDetailLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAccessDetailLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_access_detail": basetypes.ListType{
				ElemType: WorkspaceAccessDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailLocalRequest_SdkV2 as
// a WorkspaceAccessDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAccessDetailLocalRequest_SdkV2) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail_SdkV2, bool) {
	var e WorkspaceAccessDetail_SdkV2
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAccessDetail_SdkV2
	d := m.WorkspaceAccessDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailLocalRequest_SdkV2.
func (m *CreateWorkspaceAccessDetailLocalRequest_SdkV2) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_access_detail"]
	m.WorkspaceAccessDetail = types.ListValueMust(t, vs)
}

type CreateWorkspaceAccessDetailRequest_SdkV2 struct {
	// Required. The parent path for workspace access detail.
	Parent types.String `tfsdk:"-"`
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail types.List `tfsdk:"workspace_access_detail"`
}

func (to *CreateWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAccessDetailRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAccessDetail
				toWorkspaceAccessDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (to *CreateWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAccessDetailRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m CreateWorkspaceAccessDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAccessDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAccessDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAccessDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAccessDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":                  m.Parent,
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"workspace_access_detail": basetypes.ListType{
				ElemType: WorkspaceAccessDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailRequest_SdkV2 as
// a WorkspaceAccessDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAccessDetailRequest_SdkV2) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail_SdkV2, bool) {
	var e WorkspaceAccessDetail_SdkV2
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAccessDetail_SdkV2
	d := m.WorkspaceAccessDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailRequest_SdkV2.
func (m *CreateWorkspaceAccessDetailRequest_SdkV2) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_access_detail"]
	m.WorkspaceAccessDetail = types.ListValueMust(t, vs)
}

type DeleteGroupProxyRequest_SdkV2 struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupProxyRequest_SdkV2) {
}

func (to *DeleteGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupProxyRequest_SdkV2) {
}

func (m DeleteGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteGroupRequest_SdkV2 struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupRequest_SdkV2) {
}

func (to *DeleteGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupRequest_SdkV2) {
}

func (m DeleteGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalProxyRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalProxyRequest_SdkV2) {
}

func (to *DeleteServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalProxyRequest_SdkV2) {
}

func (m DeleteServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalRequest_SdkV2) {
}

func (to *DeleteServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalRequest_SdkV2) {
}

func (m DeleteServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteUserProxyRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserProxyRequest_SdkV2) {
}

func (to *DeleteUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteUserProxyRequest_SdkV2) {
}

func (m DeleteUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteUserRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserRequest_SdkV2) {
}

func (to *DeleteUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteUserRequest_SdkV2) {
}

func (m DeleteUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAccessDetailLocalRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (to *DeleteWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (m DeleteWorkspaceAccessDetailLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAccessDetailLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAccessDetailLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAccessDetailLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAccessDetailRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks to delete workspace access
	// for.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAccessDetailRequest_SdkV2) {
}

func (to *DeleteWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAccessDetailRequest_SdkV2) {
}

func (m DeleteWorkspaceAccessDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAccessDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAccessDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAccessDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAccessDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetGroupProxyRequest_SdkV2 struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupProxyRequest_SdkV2) {
}

func (to *GetGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetGroupProxyRequest_SdkV2) {
}

func (m GetGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetGroupRequest_SdkV2 struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupRequest_SdkV2) {
}

func (to *GetGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetGroupRequest_SdkV2) {
}

func (m GetGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetServicePrincipalProxyRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalProxyRequest_SdkV2) {
}

func (to *GetServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalProxyRequest_SdkV2) {
}

func (m GetServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetServicePrincipalRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalRequest_SdkV2) {
}

func (to *GetServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalRequest_SdkV2) {
}

func (m GetServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetUserProxyRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserProxyRequest_SdkV2) {
}

func (to *GetUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUserProxyRequest_SdkV2) {
}

func (m GetUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetUserRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserRequest_SdkV2) {
}

func (to *GetUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUserRequest_SdkV2) {
}

func (m GetUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAccessDetailLocalRequest_SdkV2 struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (to *GetWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"view":         m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailRequest_SdkV2) {
}

func (to *GetWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailRequest_SdkV2) {
}

func (m GetWorkspaceAccessDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceAccessDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAccessDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAccessDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"view":         m.View,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Group_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Group_SdkV2) {
}

func (to *Group_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Group_SdkV2) {
}

func (m Group_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Group_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Group_SdkV2
// only implements ToObjectValue() and Type().
func (m Group_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":  m.AccountId,
			"external_id": m.ExternalId,
			"group_name":  m.GroupName,
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Group_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_name":  types.StringType,
			"internal_id": types.Int64Type,
		},
	}
}

type ListGroupsProxyRequest_SdkV2 struct {
	// The maximum number of groups to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListGroupsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsProxyRequest_SdkV2) {
}

func (to *ListGroupsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsProxyRequest_SdkV2) {
}

func (m ListGroupsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGroupsProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListGroupsProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListGroupsRequest_SdkV2 struct {
	// The maximum number of groups to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListGroupsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsRequest_SdkV2) {
}

func (to *ListGroupsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsRequest_SdkV2) {
}

func (m ListGroupsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGroupsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListGroupsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
type ListGroupsResponse_SdkV2 struct {
	Groups types.List `tfsdk:"groups"`
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListGroupsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsResponse_SdkV2) {
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() && to.Groups.IsNull() && len(from.Groups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Groups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Groups = from.Groups
	}
}

func (to *ListGroupsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsResponse_SdkV2) {
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() && to.Groups.IsNull() && len(from.Groups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Groups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Groups = from.Groups
	}
}

func (m ListGroupsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListGroupsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"groups": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGroupsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListGroupsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"groups":          m.Groups,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"groups": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetGroups returns the value of the Groups field in ListGroupsResponse_SdkV2 as
// a slice of Group_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListGroupsResponse_SdkV2) GetGroups(ctx context.Context) ([]Group_SdkV2, bool) {
	if m.Groups.IsNull() || m.Groups.IsUnknown() {
		return nil, false
	}
	var v []Group_SdkV2
	d := m.Groups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGroups sets the value of the Groups field in ListGroupsResponse_SdkV2.
func (m *ListGroupsResponse_SdkV2) SetGroups(ctx context.Context, v []Group_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Groups = types.ListValueMust(t, vs)
}

type ListServicePrincipalsProxyRequest_SdkV2 struct {
	// The maximum number of SPs to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsProxyRequest_SdkV2) {
}

func (to *ListServicePrincipalsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsProxyRequest_SdkV2) {
}

func (m ListServicePrincipalsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalsProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListServicePrincipalsRequest_SdkV2 struct {
	// The maximum number of service principals to return. The service may
	// return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsRequest_SdkV2) {
}

func (to *ListServicePrincipalsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsRequest_SdkV2) {
}

func (m ListServicePrincipalsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
type ListServicePrincipalsResponse_SdkV2 struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	ServicePrincipals types.List `tfsdk:"service_principals"`
}

func (to *ListServicePrincipalsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsResponse_SdkV2) {
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() && to.ServicePrincipals.IsNull() && len(from.ServicePrincipals.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServicePrincipals, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServicePrincipals = from.ServicePrincipals
	}
}

func (to *ListServicePrincipalsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsResponse_SdkV2) {
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() && to.ServicePrincipals.IsNull() && len(from.ServicePrincipals.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServicePrincipals, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServicePrincipals = from.ServicePrincipals
	}
}

func (m ListServicePrincipalsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListServicePrincipalsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principals": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":    m.NextPageToken,
			"service_principals": m.ServicePrincipals,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"service_principals": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServicePrincipals returns the value of the ServicePrincipals field in ListServicePrincipalsResponse_SdkV2 as
// a slice of ServicePrincipal_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListServicePrincipalsResponse_SdkV2) GetServicePrincipals(ctx context.Context) ([]ServicePrincipal_SdkV2, bool) {
	if m.ServicePrincipals.IsNull() || m.ServicePrincipals.IsUnknown() {
		return nil, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipals.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServicePrincipals sets the value of the ServicePrincipals field in ListServicePrincipalsResponse_SdkV2.
func (m *ListServicePrincipalsResponse_SdkV2) SetServicePrincipals(ctx context.Context, v []ServicePrincipal_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principals"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServicePrincipals = types.ListValueMust(t, vs)
}

type ListUsersProxyRequest_SdkV2 struct {
	// The maximum number of users to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListUsersProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersProxyRequest_SdkV2) {
}

func (to *ListUsersProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersProxyRequest_SdkV2) {
}

func (m ListUsersProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListUsersProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListUsersProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListUsersRequest_SdkV2 struct {
	// The maximum number of users to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListUsersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersRequest_SdkV2) {
}

func (to *ListUsersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersRequest_SdkV2) {
}

func (m ListUsersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListUsersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListUsersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
type ListUsersResponse_SdkV2 struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Users types.List `tfsdk:"users"`
}

func (to *ListUsersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersResponse_SdkV2) {
	if !from.Users.IsNull() && !from.Users.IsUnknown() && to.Users.IsNull() && len(from.Users.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Users, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Users = from.Users
	}
}

func (to *ListUsersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersResponse_SdkV2) {
	if !from.Users.IsNull() && !from.Users.IsUnknown() && to.Users.IsNull() && len(from.Users.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Users, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Users = from.Users
	}
}

func (m ListUsersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListUsersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"users": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUsersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListUsersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"users":           m.Users,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"users": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUsers returns the value of the Users field in ListUsersResponse_SdkV2 as
// a slice of User_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListUsersResponse_SdkV2) GetUsers(ctx context.Context) ([]User_SdkV2, bool) {
	if m.Users.IsNull() || m.Users.IsUnknown() {
		return nil, false
	}
	var v []User_SdkV2
	d := m.Users.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsers sets the value of the Users field in ListUsersResponse_SdkV2.
func (m *ListUsersResponse_SdkV2) SetUsers(ctx context.Context, v []User_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["users"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Users = types.ListValueMust(t, vs)
}

type ListWorkspaceAccessDetailsLocalRequest_SdkV2 struct {
	// The maximum number of workspace access details to return. The service may
	// return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListWorkspaceAccessDetails call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAccessDetailsLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsLocalRequest_SdkV2) {
}

func (to *ListWorkspaceAccessDetailsLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsLocalRequest_SdkV2) {
}

func (m ListWorkspaceAccessDetailsLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAccessDetailsLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAccessDetailsLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceAccessDetailsRequest_SdkV2 struct {
	// The maximum number of workspace access details to return. The service may
	// return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListWorkspaceAccessDetails call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The workspace ID for which the workspace access details are being
	// fetched.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *ListWorkspaceAccessDetailsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsRequest_SdkV2) {
}

func (to *ListWorkspaceAccessDetailsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsRequest_SdkV2) {
}

func (m ListWorkspaceAccessDetailsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAccessDetailsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAccessDetailsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// TODO: Write description later when this method is implemented
type ListWorkspaceAccessDetailsResponse_SdkV2 struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	WorkspaceAccessDetails types.List `tfsdk:"workspace_access_details"`
}

func (to *ListWorkspaceAccessDetailsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsResponse_SdkV2) {
	if !from.WorkspaceAccessDetails.IsNull() && !from.WorkspaceAccessDetails.IsUnknown() && to.WorkspaceAccessDetails.IsNull() && len(from.WorkspaceAccessDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAccessDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAccessDetails = from.WorkspaceAccessDetails
	}
}

func (to *ListWorkspaceAccessDetailsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsResponse_SdkV2) {
	if !from.WorkspaceAccessDetails.IsNull() && !from.WorkspaceAccessDetails.IsUnknown() && to.WorkspaceAccessDetails.IsNull() && len(from.WorkspaceAccessDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAccessDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAccessDetails = from.WorkspaceAccessDetails
	}
}

func (m ListWorkspaceAccessDetailsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["workspace_access_details"] = attrs["workspace_access_details"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAccessDetailsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAccessDetailsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_details": reflect.TypeOf(WorkspaceAccessDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"workspace_access_details": m.WorkspaceAccessDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_access_details": basetypes.ListType{
				ElemType: WorkspaceAccessDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAccessDetails returns the value of the WorkspaceAccessDetails field in ListWorkspaceAccessDetailsResponse_SdkV2 as
// a slice of WorkspaceAccessDetail_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAccessDetailsResponse_SdkV2) GetWorkspaceAccessDetails(ctx context.Context) ([]WorkspaceAccessDetail_SdkV2, bool) {
	if m.WorkspaceAccessDetails.IsNull() || m.WorkspaceAccessDetails.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAccessDetail_SdkV2
	d := m.WorkspaceAccessDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetails sets the value of the WorkspaceAccessDetails field in ListWorkspaceAccessDetailsResponse_SdkV2.
func (m *ListWorkspaceAccessDetailsResponse_SdkV2) SetWorkspaceAccessDetails(ctx context.Context, v []WorkspaceAccessDetail_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_access_details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceAccessDetails = types.ListValueMust(t, vs)
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// groupname, and inherited parent groups.
type ResolveGroupProxyRequest_SdkV2 struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupProxyRequest_SdkV2) {
}

func (to *ResolveGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupProxyRequest_SdkV2) {
}

func (m ResolveGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupRequest_SdkV2) {
}

func (to *ResolveGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupRequest_SdkV2) {
}

func (m ResolveGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveGroupResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveGroupResponse_SdkV2) {
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

func (to *ResolveGroupResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveGroupResponse_SdkV2) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m ResolveGroupResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveGroupResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveGroupResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveGroupResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group": m.Group,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveGroupResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResolveGroupResponse_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := m.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in ResolveGroupResponse_SdkV2.
func (m *ResolveGroupResponse_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	m.Group = types.ListValueMust(t, vs)
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest_SdkV2 struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalProxyRequest_SdkV2) {
}

func (to *ResolveServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalProxyRequest_SdkV2) {
}

func (m ResolveServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalRequest_SdkV2) {
}

func (to *ResolveServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalRequest_SdkV2) {
}

func (m ResolveServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveServicePrincipalResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveServicePrincipalResponse_SdkV2) {
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

func (to *ResolveServicePrincipalResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveServicePrincipalResponse_SdkV2) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m ResolveServicePrincipalResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveServicePrincipalResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveServicePrincipalResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveServicePrincipalResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveServicePrincipalResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResolveServicePrincipalResponse_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in ResolveServicePrincipalResponse_SdkV2.
func (m *ResolveServicePrincipalResponse_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	m.ServicePrincipal = types.ListValueMust(t, vs)
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserProxyRequest_SdkV2 struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ResolveUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserProxyRequest_SdkV2) {
}

func (to *ResolveUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveUserProxyRequest_SdkV2) {
}

func (m ResolveUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserRequest_SdkV2) {
}

func (to *ResolveUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveUserRequest_SdkV2) {
}

func (m ResolveUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ResolveUserResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveUserResponse_SdkV2) {
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

func (to *ResolveUserResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveUserResponse_SdkV2) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m ResolveUserResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveUserResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveUserResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveUserResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"user": m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveUserResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ResolveUserResponse_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in ResolveUserResponse_SdkV2.
func (m *ResolveUserResponse_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
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

func (to *ServicePrincipal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServicePrincipal_SdkV2) {
}

func (to *ServicePrincipal_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServicePrincipal_SdkV2) {
}

func (m ServicePrincipal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ServicePrincipal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal_SdkV2
// only implements ToObjectValue() and Type().
func (m ServicePrincipal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        m.AccountId,
			"account_sp_status": m.AccountSpStatus,
			"application_id":    m.ApplicationId,
			"display_name":      m.DisplayName,
			"external_id":       m.ExternalId,
			"internal_id":       m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
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

type UpdateGroupProxyRequest_SdkV2 struct {
	// Required. Group to be updated in <Databricks>
	Group types.List `tfsdk:"group"`
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateGroupProxyRequest_SdkV2) {
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

func (to *UpdateGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateGroupProxyRequest_SdkV2) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m UpdateGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["group"] = attrs["group"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group":       m.Group,
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
		},
	}
}

// GetGroup returns the value of the Group field in UpdateGroupProxyRequest_SdkV2 as
// a Group_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateGroupProxyRequest_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := m.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in UpdateGroupProxyRequest_SdkV2.
func (m *UpdateGroupProxyRequest_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	m.Group = types.ListValueMust(t, vs)
}

type UpdateGroupRequest_SdkV2 struct {
	// Required. Group to be updated in <Databricks>
	Group types.List `tfsdk:"group"`
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateGroupRequest_SdkV2) {
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

func (to *UpdateGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateGroupRequest_SdkV2) {
	if !from.Group.IsNull() && !from.Group.IsUnknown() {
		if toGroup, ok := to.GetGroup(ctx); ok {
			if fromGroup, ok := from.GetGroup(ctx); ok {
				toGroup.SyncFieldsDuringRead(ctx, fromGroup)
				to.SetGroup(ctx, toGroup)
			}
		}
	}
}

func (m UpdateGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group"] = attrs["group"].SetRequired()
	attrs["group"] = attrs["group"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"group": reflect.TypeOf(Group_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group":       m.Group,
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
		},
	}
}

// GetGroup returns the value of the Group field in UpdateGroupRequest_SdkV2 as
// a Group_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateGroupRequest_SdkV2) GetGroup(ctx context.Context) (Group_SdkV2, bool) {
	var e Group_SdkV2
	if m.Group.IsNull() || m.Group.IsUnknown() {
		return e, false
	}
	var v []Group_SdkV2
	d := m.Group.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGroup sets the value of the Group field in UpdateGroupRequest_SdkV2.
func (m *UpdateGroupRequest_SdkV2) SetGroup(ctx context.Context, v Group_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["group"]
	m.Group = types.ListValueMust(t, vs)
}

type UpdateServicePrincipalProxyRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Required. Service principal to be updated in <Databricks>
	ServicePrincipal types.List `tfsdk:"service_principal"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalProxyRequest_SdkV2) {
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

func (to *UpdateServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalProxyRequest_SdkV2) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m UpdateServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["service_principal"] = attrs["service_principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id":       m.InternalId,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"service_principal": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in UpdateServicePrincipalProxyRequest_SdkV2 as
// a ServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateServicePrincipalProxyRequest_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in UpdateServicePrincipalProxyRequest_SdkV2.
func (m *UpdateServicePrincipalProxyRequest_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	m.ServicePrincipal = types.ListValueMust(t, vs)
}

type UpdateServicePrincipalRequest_SdkV2 struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Required. Service Principal to be updated in <Databricks>
	ServicePrincipal types.List `tfsdk:"service_principal"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalRequest_SdkV2) {
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

func (to *UpdateServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalRequest_SdkV2) {
	if !from.ServicePrincipal.IsNull() && !from.ServicePrincipal.IsUnknown() {
		if toServicePrincipal, ok := to.GetServicePrincipal(ctx); ok {
			if fromServicePrincipal, ok := from.GetServicePrincipal(ctx); ok {
				toServicePrincipal.SyncFieldsDuringRead(ctx, fromServicePrincipal)
				to.SetServicePrincipal(ctx, toServicePrincipal)
			}
		}
	}
}

func (m UpdateServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal"] = attrs["service_principal"].SetRequired()
	attrs["service_principal"] = attrs["service_principal"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principal": reflect.TypeOf(ServicePrincipal_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id":       m.InternalId,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"service_principal": basetypes.ListType{
				ElemType: ServicePrincipal_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetServicePrincipal returns the value of the ServicePrincipal field in UpdateServicePrincipalRequest_SdkV2 as
// a ServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateServicePrincipalRequest_SdkV2) GetServicePrincipal(ctx context.Context) (ServicePrincipal_SdkV2, bool) {
	var e ServicePrincipal_SdkV2
	if m.ServicePrincipal.IsNull() || m.ServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []ServicePrincipal_SdkV2
	d := m.ServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetServicePrincipal sets the value of the ServicePrincipal field in UpdateServicePrincipalRequest_SdkV2.
func (m *UpdateServicePrincipalRequest_SdkV2) SetServicePrincipal(ctx context.Context, v ServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["service_principal"]
	m.ServicePrincipal = types.ListValueMust(t, vs)
}

type UpdateUserProxyRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.List `tfsdk:"user"`
}

func (to *UpdateUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateUserProxyRequest_SdkV2) {
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

func (to *UpdateUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateUserProxyRequest_SdkV2) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m UpdateUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUser returns the value of the User field in UpdateUserProxyRequest_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateUserProxyRequest_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in UpdateUserProxyRequest_SdkV2.
func (m *UpdateUserProxyRequest_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

type UpdateUserRequest_SdkV2 struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.List `tfsdk:"user"`
}

func (to *UpdateUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateUserRequest_SdkV2) {
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

func (to *UpdateUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateUserRequest_SdkV2) {
	if !from.User.IsNull() && !from.User.IsUnknown() {
		if toUser, ok := to.GetUser(ctx); ok {
			if fromUser, ok := from.GetUser(ctx); ok {
				toUser.SyncFieldsDuringRead(ctx, fromUser)
				to.SetUser(ctx, toUser)
			}
		}
	}
}

func (m UpdateUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["user"] = attrs["user"].SetRequired()
	attrs["user"] = attrs["user"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["internal_id"] = attrs["internal_id"].SetRequired()
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
func (m UpdateUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"user": reflect.TypeOf(User_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
			"user": basetypes.ListType{
				ElemType: User_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUser returns the value of the User field in UpdateUserRequest_SdkV2 as
// a User_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateUserRequest_SdkV2) GetUser(ctx context.Context) (User_SdkV2, bool) {
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

// SetUser sets the value of the User field in UpdateUserRequest_SdkV2.
func (m *UpdateUserRequest_SdkV2) SetUser(ctx context.Context, v User_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user"]
	m.User = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAccessDetailLocalRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. WorkspaceAccessDetail to be updated in <Databricks>
	WorkspaceAccessDetail types.List `tfsdk:"workspace_access_detail"`
}

func (to *UpdateWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAccessDetailLocalRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAccessDetail
				toWorkspaceAccessDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (to *UpdateWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAccessDetailLocalRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAccessDetailLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAccessDetailLocalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAccessDetailLocalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAccessDetailLocalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":            m.PrincipalId,
			"update_mask":             m.UpdateMask,
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"update_mask":  types.StringType,
			"workspace_access_detail": basetypes.ListType{
				ElemType: WorkspaceAccessDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailLocalRequest_SdkV2 as
// a WorkspaceAccessDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAccessDetailLocalRequest_SdkV2) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail_SdkV2, bool) {
	var e WorkspaceAccessDetail_SdkV2
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAccessDetail_SdkV2
	d := m.WorkspaceAccessDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailLocalRequest_SdkV2.
func (m *UpdateWorkspaceAccessDetailLocalRequest_SdkV2) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_access_detail"]
	m.WorkspaceAccessDetail = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAccessDetailRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace access detail to be updated in <Databricks>
	WorkspaceAccessDetail types.List `tfsdk:"workspace_access_detail"`
	// Required. The workspace ID for which the workspace access detail is being
	// updated.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAccessDetailRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				// Recursively sync the fields of WorkspaceAccessDetail
				toWorkspaceAccessDetail.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (to *UpdateWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAccessDetailRequest_SdkV2) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAccessDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAccessDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceAccessDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAccessDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAccessDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":            m.PrincipalId,
			"update_mask":             m.UpdateMask,
			"workspace_access_detail": m.WorkspaceAccessDetail,
			"workspace_id":            m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"update_mask":  types.StringType,
			"workspace_access_detail": basetypes.ListType{
				ElemType: WorkspaceAccessDetail_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailRequest_SdkV2 as
// a WorkspaceAccessDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAccessDetailRequest_SdkV2) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail_SdkV2, bool) {
	var e WorkspaceAccessDetail_SdkV2
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAccessDetail_SdkV2
	d := m.WorkspaceAccessDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailRequest_SdkV2.
func (m *UpdateWorkspaceAccessDetailRequest_SdkV2) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_access_detail"]
	m.WorkspaceAccessDetail = types.ListValueMust(t, vs)
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

func (to *User_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User_SdkV2) {
	if !from.Name.IsNull() && !from.Name.IsUnknown() {
		if toName, ok := to.GetName(ctx); ok {
			if fromName, ok := from.GetName(ctx); ok {
				// Recursively sync the fields of Name
				toName.SyncFieldsDuringCreateOrUpdate(ctx, fromName)
				to.SetName(ctx, toName)
			}
		}
	}
}

func (to *User_SdkV2) SyncFieldsDuringRead(ctx context.Context, from User_SdkV2) {
	if !from.Name.IsNull() && !from.Name.IsUnknown() {
		if toName, ok := to.GetName(ctx); ok {
			if fromName, ok := from.GetName(ctx); ok {
				toName.SyncFieldsDuringRead(ctx, fromName)
				to.SetName(ctx, toName)
			}
		}
	}
}

func (m User_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m User_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"name": reflect.TypeOf(UserName_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, User_SdkV2
// only implements ToObjectValue() and Type().
func (m User_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":          m.AccountId,
			"account_user_status": m.AccountUserStatus,
			"external_id":         m.ExternalId,
			"internal_id":         m.InternalId,
			"name":                m.Name,
			"username":            m.Username,
		})
}

// Type implements basetypes.ObjectValuable.
func (m User_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *User_SdkV2) GetName(ctx context.Context) (UserName_SdkV2, bool) {
	var e UserName_SdkV2
	if m.Name.IsNull() || m.Name.IsUnknown() {
		return e, false
	}
	var v []UserName_SdkV2
	d := m.Name.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetName sets the value of the Name field in User_SdkV2.
func (m *User_SdkV2) SetName(ctx context.Context, v UserName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["name"]
	m.Name = types.ListValueMust(t, vs)
}

type UserName_SdkV2 struct {
	FamilyName types.String `tfsdk:"family_name"`

	GivenName types.String `tfsdk:"given_name"`
}

func (to *UserName_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UserName_SdkV2) {
}

func (to *UserName_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UserName_SdkV2) {
}

func (m UserName_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UserName_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserName_SdkV2
// only implements ToObjectValue() and Type().
func (m UserName_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UserName_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *WorkspaceAccessDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAccessDetail_SdkV2) {
	if !from.Permissions.IsNull() && !from.Permissions.IsUnknown() && to.Permissions.IsNull() && len(from.Permissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Permissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Permissions = from.Permissions
	}
}

func (to *WorkspaceAccessDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAccessDetail_SdkV2) {
	if !from.Permissions.IsNull() && !from.Permissions.IsUnknown() && to.Permissions.IsNull() && len(from.Permissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Permissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Permissions = from.Permissions
	}
}

func (m WorkspaceAccessDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceAccessDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAccessDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceAccessDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m WorkspaceAccessDetail_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceAccessDetail_SdkV2) GetPermissions(ctx context.Context) ([]types.String, bool) {
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

// SetPermissions sets the value of the Permissions field in WorkspaceAccessDetail_SdkV2.
func (m *WorkspaceAccessDetail_SdkV2) SetPermissions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Permissions = types.ListValueMust(t, vs)
}
