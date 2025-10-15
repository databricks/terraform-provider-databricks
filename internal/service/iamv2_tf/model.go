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

type CreateWorkspaceAccessDetailLocalRequest struct {
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail types.Object `tfsdk:"workspace_access_detail"`
}

func (to *CreateWorkspaceAccessDetailLocalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAccessDetailLocalRequest) {
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

func (to *CreateWorkspaceAccessDetailLocalRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAccessDetailLocalRequest) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m CreateWorkspaceAccessDetailLocalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAccessDetailLocalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAccessDetailLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAccessDetailLocalRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAccessDetailLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAccessDetailLocalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_access_detail": WorkspaceAccessDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailLocalRequest as
// a WorkspaceAccessDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAccessDetailLocalRequest) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail, bool) {
	var e WorkspaceAccessDetail
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAccessDetail
	d := m.WorkspaceAccessDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailLocalRequest.
func (m *CreateWorkspaceAccessDetailLocalRequest) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAccessDetail = vs
}

type CreateWorkspaceAccessDetailRequest struct {
	// Required. The parent path for workspace access detail.
	Parent types.String `tfsdk:"-"`
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail types.Object `tfsdk:"workspace_access_detail"`
}

func (to *CreateWorkspaceAccessDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAccessDetailRequest) {
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

func (to *CreateWorkspaceAccessDetailRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAccessDetailRequest) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m CreateWorkspaceAccessDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
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
func (m CreateWorkspaceAccessDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAccessDetailRequest
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAccessDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":                  m.Parent,
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAccessDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent":                  types.StringType,
			"workspace_access_detail": WorkspaceAccessDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailRequest as
// a WorkspaceAccessDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAccessDetailRequest) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail, bool) {
	var e WorkspaceAccessDetail
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAccessDetail
	d := m.WorkspaceAccessDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in CreateWorkspaceAccessDetailRequest.
func (m *CreateWorkspaceAccessDetailRequest) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAccessDetail = vs
}

type DeleteGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupProxyRequest) {
}

func (to *DeleteGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupProxyRequest) {
}

func (m DeleteGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupRequest) {
}

func (to *DeleteGroupRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupRequest) {
}

func (m DeleteGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalProxyRequest) {
}

func (to *DeleteServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalProxyRequest) {
}

func (m DeleteServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalRequest) {
}

func (to *DeleteServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalRequest) {
}

func (m DeleteServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserProxyRequest) {
}

func (to *DeleteUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteUserProxyRequest) {
}

func (m DeleteUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserRequest) {
}

func (to *DeleteUserRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteUserRequest) {
}

func (m DeleteUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAccessDetailLocalRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAccessDetailLocalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAccessDetailLocalRequest) {
}

func (to *DeleteWorkspaceAccessDetailLocalRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAccessDetailLocalRequest) {
}

func (m DeleteWorkspaceAccessDetailLocalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWorkspaceAccessDetailLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAccessDetailLocalRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAccessDetailLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAccessDetailLocalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAccessDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace access
	// for.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAccessDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAccessDetailRequest) {
}

func (to *DeleteWorkspaceAccessDetailRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAccessDetailRequest) {
}

func (m DeleteWorkspaceAccessDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWorkspaceAccessDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAccessDetailRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAccessDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAccessDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetGroupProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupProxyRequest) {
}

func (to *GetGroupProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetGroupProxyRequest) {
}

func (m GetGroupProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupRequest) {
}

func (to *GetGroupRequest) SyncFieldsDuringRead(ctx context.Context, from GetGroupRequest) {
}

func (m GetGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalProxyRequest) {
}

func (to *GetServicePrincipalProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalProxyRequest) {
}

func (m GetServicePrincipalProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalRequest) {
}

func (to *GetServicePrincipalRequest) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalRequest) {
}

func (m GetServicePrincipalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetUserProxyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserProxyRequest) {
}

func (to *GetUserProxyRequest) SyncFieldsDuringRead(ctx context.Context, from GetUserProxyRequest) {
}

func (m GetUserProxyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
		},
	}
}

type GetUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
}

func (to *GetUserRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserRequest) {
}

func (to *GetUserRequest) SyncFieldsDuringRead(ctx context.Context, from GetUserRequest) {
}

func (m GetUserRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
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

func (to *Group) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Group) {
}

func (to *Group) SyncFieldsDuringRead(ctx context.Context, from Group) {
}

func (m Group) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"group_name":  m.GroupName,
			"internal_id": m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Group) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_name":  types.StringType,
			"internal_id": types.Int64Type,
		},
	}
}

type ListGroupsProxyRequest struct {
	// The maximum number of groups to return. The service may return fewer than
	// this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListGroupsRequest struct {
	// The maximum number of groups to return. The service may return fewer than
	// this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
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
}

func (to *ListGroupsResponse) SyncFieldsDuringRead(ctx context.Context, from ListGroupsResponse) {
	if !from.Groups.IsNull() && !from.Groups.IsUnknown() && to.Groups.IsNull() && len(from.Groups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Groups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Groups = from.Groups
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
	// The maximum number of SPs to return. The service may return fewer than
	// this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListServicePrincipalsRequest struct {
	// The maximum number of service principals to return. The service may
	// return fewer than this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
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
}

func (to *ListServicePrincipalsResponse) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsResponse) {
	if !from.ServicePrincipals.IsNull() && !from.ServicePrincipals.IsUnknown() && to.ServicePrincipals.IsNull() && len(from.ServicePrincipals.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServicePrincipals, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServicePrincipals = from.ServicePrincipals
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

type ListUsersProxyRequest struct {
	// The maximum number of users to return. The service may return fewer than
	// this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListUsersRequest struct {
	// The maximum number of users to return. The service may return fewer than
	// this value.
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
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// TODO: Write description later when this method is implemented
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
}

func (to *ListUsersResponse) SyncFieldsDuringRead(ctx context.Context, from ListUsersResponse) {
	if !from.Users.IsNull() && !from.Users.IsUnknown() && to.Users.IsNull() && len(from.Users.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Users, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Users = from.Users
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

type ListWorkspaceAccessDetailsLocalRequest struct {
	// The maximum number of workspace access details to return. The service may
	// return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous ListWorkspaceAccessDetails call.
	// Provide this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAccessDetailsLocalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsLocalRequest) {
}

func (to *ListWorkspaceAccessDetailsLocalRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsLocalRequest) {
}

func (m ListWorkspaceAccessDetailsLocalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAccessDetailsLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsLocalRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsLocalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceAccessDetailsRequest struct {
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

func (to *ListWorkspaceAccessDetailsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsRequest) {
}

func (to *ListWorkspaceAccessDetailsRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsRequest) {
}

func (m ListWorkspaceAccessDetailsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAccessDetailsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// TODO: Write description later when this method is implemented
type ListWorkspaceAccessDetailsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	WorkspaceAccessDetails types.List `tfsdk:"workspace_access_details"`
}

func (to *ListWorkspaceAccessDetailsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAccessDetailsResponse) {
	if !from.WorkspaceAccessDetails.IsNull() && !from.WorkspaceAccessDetails.IsUnknown() && to.WorkspaceAccessDetails.IsNull() && len(from.WorkspaceAccessDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAccessDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAccessDetails = from.WorkspaceAccessDetails
	}
}

func (to *ListWorkspaceAccessDetailsResponse) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAccessDetailsResponse) {
	if !from.WorkspaceAccessDetails.IsNull() && !from.WorkspaceAccessDetails.IsUnknown() && to.WorkspaceAccessDetails.IsNull() && len(from.WorkspaceAccessDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAccessDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAccessDetails = from.WorkspaceAccessDetails
	}
}

func (m ListWorkspaceAccessDetailsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAccessDetailsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_details": reflect.TypeOf(WorkspaceAccessDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAccessDetailsResponse
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAccessDetailsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"workspace_access_details": m.WorkspaceAccessDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAccessDetailsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_access_details": basetypes.ListType{
				ElemType: WorkspaceAccessDetail{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAccessDetails returns the value of the WorkspaceAccessDetails field in ListWorkspaceAccessDetailsResponse as
// a slice of WorkspaceAccessDetail values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAccessDetailsResponse) GetWorkspaceAccessDetails(ctx context.Context) ([]WorkspaceAccessDetail, bool) {
	if m.WorkspaceAccessDetails.IsNull() || m.WorkspaceAccessDetails.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAccessDetail
	d := m.WorkspaceAccessDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetails sets the value of the WorkspaceAccessDetails field in ListWorkspaceAccessDetailsResponse.
func (m *ListWorkspaceAccessDetailsResponse) SetWorkspaceAccessDetails(ctx context.Context, v []WorkspaceAccessDetail) {
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
// groupname, and inherited parent groups.
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
// service principal's displayname, status, and inherited parent groups.
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
// service principal's displayname, status, and inherited parent groups.
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
// displayname, status, and inherited parent groups.
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
// displayname, status, and inherited parent groups.
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
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id"`
	// Display name of the service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// ExternalId of the service principal in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Internal service principal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"internal_id"`
}

func (to *ServicePrincipal) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServicePrincipal) {
}

func (to *ServicePrincipal) SyncFieldsDuringRead(ctx context.Context, from ServicePrincipal) {
}

func (m ServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"account_id":        m.AccountId,
			"account_sp_status": m.AccountSpStatus,
			"application_id":    m.ApplicationId,
			"display_name":      m.DisplayName,
			"external_id":       m.ExternalId,
			"internal_id":       m.InternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServicePrincipal) Type(ctx context.Context) attr.Type {
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

type UpdateGroupProxyRequest struct {
	// Required. Group to be updated in <Databricks>
	Group types.Object `tfsdk:"group"`
	// Required. Internal ID of the group in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
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
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group":       Group{}.Type(ctx),
			"internal_id": types.Int64Type,
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
	InternalId types.Int64 `tfsdk:"-"`
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
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group":       Group{}.Type(ctx),
			"internal_id": types.Int64Type,
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
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Required. Service principal to be updated in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
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
			"internal_id":       m.InternalId,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id":       types.Int64Type,
			"service_principal": ServicePrincipal{}.Type(ctx),
			"update_mask":       types.StringType,
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
	// Required. Internal ID of the service principal in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Required. Service Principal to be updated in <Databricks>
	ServicePrincipal types.Object `tfsdk:"service_principal"`
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
			"internal_id":       m.InternalId,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id":       types.Int64Type,
			"service_principal": ServicePrincipal{}.Type(ctx),
			"update_mask":       types.StringType,
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
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.Object `tfsdk:"user"`
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
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserProxyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
			"user":        User{}.Type(ctx),
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
	// Required. Internal ID of the user in Databricks.
	InternalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. User to be updated in <Databricks>
	User types.Object `tfsdk:"user"`
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
			"internal_id": m.InternalId,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_id": types.Int64Type,
			"update_mask": types.StringType,
			"user":        User{}.Type(ctx),
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

type UpdateWorkspaceAccessDetailLocalRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. WorkspaceAccessDetail to be updated in <Databricks>
	WorkspaceAccessDetail types.Object `tfsdk:"workspace_access_detail"`
}

func (to *UpdateWorkspaceAccessDetailLocalRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAccessDetailLocalRequest) {
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

func (to *UpdateWorkspaceAccessDetailLocalRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAccessDetailLocalRequest) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAccessDetailLocalRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
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
func (m UpdateWorkspaceAccessDetailLocalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAccessDetailLocalRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAccessDetailLocalRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":            m.PrincipalId,
			"update_mask":             m.UpdateMask,
			"workspace_access_detail": m.WorkspaceAccessDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAccessDetailLocalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":            types.Int64Type,
			"update_mask":             types.StringType,
			"workspace_access_detail": WorkspaceAccessDetail{}.Type(ctx),
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailLocalRequest as
// a WorkspaceAccessDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAccessDetailLocalRequest) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail, bool) {
	var e WorkspaceAccessDetail
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAccessDetail
	d := m.WorkspaceAccessDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailLocalRequest.
func (m *UpdateWorkspaceAccessDetailLocalRequest) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAccessDetail = vs
}

type UpdateWorkspaceAccessDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Optional. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace access detail to be updated in <Databricks>
	WorkspaceAccessDetail types.Object `tfsdk:"workspace_access_detail"`
	// Required. The workspace ID for which the workspace access detail is being
	// updated.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceAccessDetailRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAccessDetailRequest) {
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

func (to *UpdateWorkspaceAccessDetailRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAccessDetailRequest) {
	if !from.WorkspaceAccessDetail.IsNull() && !from.WorkspaceAccessDetail.IsUnknown() {
		if toWorkspaceAccessDetail, ok := to.GetWorkspaceAccessDetail(ctx); ok {
			if fromWorkspaceAccessDetail, ok := from.GetWorkspaceAccessDetail(ctx); ok {
				toWorkspaceAccessDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAccessDetail)
				to.SetWorkspaceAccessDetail(ctx, toWorkspaceAccessDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAccessDetailRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_access_detail"] = attrs["workspace_access_detail"].SetRequired()
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
func (m UpdateWorkspaceAccessDetailRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_access_detail": reflect.TypeOf(WorkspaceAccessDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAccessDetailRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAccessDetailRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateWorkspaceAccessDetailRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id":            types.Int64Type,
			"update_mask":             types.StringType,
			"workspace_access_detail": WorkspaceAccessDetail{}.Type(ctx),
			"workspace_id":            types.Int64Type,
		},
	}
}

// GetWorkspaceAccessDetail returns the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailRequest as
// a WorkspaceAccessDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAccessDetailRequest) GetWorkspaceAccessDetail(ctx context.Context) (WorkspaceAccessDetail, bool) {
	var e WorkspaceAccessDetail
	if m.WorkspaceAccessDetail.IsNull() || m.WorkspaceAccessDetail.IsUnknown() {
		return e, false
	}
	var v WorkspaceAccessDetail
	d := m.WorkspaceAccessDetail.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAccessDetail sets the value of the WorkspaceAccessDetail field in UpdateWorkspaceAccessDetailRequest.
func (m *UpdateWorkspaceAccessDetailRequest) SetWorkspaceAccessDetail(ctx context.Context, v WorkspaceAccessDetail) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceAccessDetail = vs
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

func (to *User) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User) {
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

func (to *User) SyncFieldsDuringRead(ctx context.Context, from User) {
	if !from.Name.IsNull() && !from.Name.IsUnknown() {
		if toName, ok := to.GetName(ctx); ok {
			if fromName, ok := from.GetName(ctx); ok {
				toName.SyncFieldsDuringRead(ctx, fromName)
				to.SetName(ctx, toName)
			}
		}
	}
}

func (m User) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"name": reflect.TypeOf(UserName{}),
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
			"internal_id":         m.InternalId,
			"name":                m.Name,
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
			"internal_id":         types.Int64Type,
			"name":                UserName{}.Type(ctx),
			"username":            types.StringType,
		},
	}
}

// GetName returns the value of the Name field in User as
// a UserName value.
// If the field is unknown or null, the boolean return value is false.
func (m *User) GetName(ctx context.Context) (UserName, bool) {
	var e UserName
	if m.Name.IsNull() || m.Name.IsUnknown() {
		return e, false
	}
	var v UserName
	d := m.Name.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetName sets the value of the Name field in User.
func (m *User) SetName(ctx context.Context, v UserName) {
	vs := v.ToObjectValue(ctx)
	m.Name = vs
}

type UserName struct {
	FamilyName types.String `tfsdk:"family_name"`

	GivenName types.String `tfsdk:"given_name"`
}

func (to *UserName) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UserName) {
}

func (to *UserName) SyncFieldsDuringRead(ctx context.Context, from UserName) {
}

func (m UserName) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UserName) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserName
// only implements ToObjectValue() and Type().
func (m UserName) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UserName) Type(ctx context.Context) attr.Type {
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
