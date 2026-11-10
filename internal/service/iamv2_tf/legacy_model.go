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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDirectGroupMemberProxyRequest_SdkV2 struct {
	// The member to add. Set its `principal_id`.
	DirectGroupMember types.List `tfsdk:"direct_group_member"`
	// The resource name of the group to add the member to, in the format
	// `groups/{group_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *CreateDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDirectGroupMemberProxyRequest_SdkV2) {
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

func (to *CreateDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDirectGroupMemberProxyRequest_SdkV2) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				toDirectGroupMember.SyncFieldsDuringRead(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (m CreateDirectGroupMemberProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direct_group_member"] = attrs["direct_group_member"].SetRequired()
	attrs["direct_group_member"] = attrs["direct_group_member"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDirectGroupMemberProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_member": reflect.TypeOf(DirectGroupMember_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectGroupMemberProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDirectGroupMemberProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_member": m.DirectGroupMember,
			"parent":              m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDirectGroupMemberProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_member": basetypes.ListType{
				ElemType: DirectGroupMember_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
		},
	}
}

// GetDirectGroupMember returns the value of the DirectGroupMember field in CreateDirectGroupMemberProxyRequest_SdkV2 as
// a DirectGroupMember_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDirectGroupMemberProxyRequest_SdkV2) GetDirectGroupMember(ctx context.Context) (DirectGroupMember_SdkV2, bool) {
	var e DirectGroupMember_SdkV2
	if m.DirectGroupMember.IsNull() || m.DirectGroupMember.IsUnknown() {
		return e, false
	}
	var v []DirectGroupMember_SdkV2
	d := m.DirectGroupMember.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectGroupMember sets the value of the DirectGroupMember field in CreateDirectGroupMemberProxyRequest_SdkV2.
func (m *CreateDirectGroupMemberProxyRequest_SdkV2) SetDirectGroupMember(ctx context.Context, v DirectGroupMember_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_group_member"]
	m.DirectGroupMember = types.ListValueMust(t, vs)
}

type CreateDirectGroupMemberRequest_SdkV2 struct {
	// The member to add. Set its `principal_id`.
	DirectGroupMember types.List `tfsdk:"direct_group_member"`
	// The resource name of the group to add the member to, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *CreateDirectGroupMemberRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDirectGroupMemberRequest_SdkV2) {
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

func (to *CreateDirectGroupMemberRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDirectGroupMemberRequest_SdkV2) {
	if !from.DirectGroupMember.IsNull() && !from.DirectGroupMember.IsUnknown() {
		if toDirectGroupMember, ok := to.GetDirectGroupMember(ctx); ok {
			if fromDirectGroupMember, ok := from.GetDirectGroupMember(ctx); ok {
				toDirectGroupMember.SyncFieldsDuringRead(ctx, fromDirectGroupMember)
				to.SetDirectGroupMember(ctx, toDirectGroupMember)
			}
		}
	}
}

func (m CreateDirectGroupMemberRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["direct_group_member"] = attrs["direct_group_member"].SetRequired()
	attrs["direct_group_member"] = attrs["direct_group_member"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDirectGroupMemberRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_member": reflect.TypeOf(DirectGroupMember_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectGroupMemberRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDirectGroupMemberRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_member": m.DirectGroupMember,
			"parent":              m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDirectGroupMemberRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_member": basetypes.ListType{
				ElemType: DirectGroupMember_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
		},
	}
}

// GetDirectGroupMember returns the value of the DirectGroupMember field in CreateDirectGroupMemberRequest_SdkV2 as
// a DirectGroupMember_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDirectGroupMemberRequest_SdkV2) GetDirectGroupMember(ctx context.Context) (DirectGroupMember_SdkV2, bool) {
	var e DirectGroupMember_SdkV2
	if m.DirectGroupMember.IsNull() || m.DirectGroupMember.IsUnknown() {
		return e, false
	}
	var v []DirectGroupMember_SdkV2
	d := m.DirectGroupMember.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectGroupMember sets the value of the DirectGroupMember field in CreateDirectGroupMemberRequest_SdkV2.
func (m *CreateDirectGroupMemberRequest_SdkV2) SetDirectGroupMember(ctx context.Context, v DirectGroupMember_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_group_member"]
	m.DirectGroupMember = types.ListValueMust(t, vs)
}

type CreateGroupProxyRequest_SdkV2 struct {
	// The group to create.
	Group types.List `tfsdk:"group"`
	// The resource name of the account to create the group in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
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
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"group":  m.Group,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
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
	// The group to create.
	Group types.List `tfsdk:"group"`
	// The resource name of the account to create the group in, in the format
	// `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
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
	attrs["parent"] = attrs["parent"].SetRequired()

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
			"group":  m.Group,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group": basetypes.ListType{
				ElemType: Group_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
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

type CreateIdentityVisibilityFilterRequest_SdkV2 struct {
	// The filter to create. Set `principal_type` to `GROUP` and exactly one of
	// `exact`, `prefix`, or `regex`. Databricks sets `name`, so any value you
	// set is ignored.
	IdentityVisibilityFilter types.List `tfsdk:"identity_visibility_filter"`
	// The ID to give the filter, which becomes the last part of its resource
	// name. If you leave it empty, Databricks generates one. It can't contain
	// `/`.
	IdentityVisibilityFilterId types.String `tfsdk:"-"`
	// The resource name of the account, in the format `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *CreateIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIdentityVisibilityFilterRequest_SdkV2) {
	if !from.IdentityVisibilityFilter.IsNull() && !from.IdentityVisibilityFilter.IsUnknown() {
		if toIdentityVisibilityFilter, ok := to.GetIdentityVisibilityFilter(ctx); ok {
			if fromIdentityVisibilityFilter, ok := from.GetIdentityVisibilityFilter(ctx); ok {
				// Recursively sync the fields of IdentityVisibilityFilter
				toIdentityVisibilityFilter.SyncFieldsDuringCreateOrUpdate(ctx, fromIdentityVisibilityFilter)
				to.SetIdentityVisibilityFilter(ctx, toIdentityVisibilityFilter)
			}
		}
	}
}

func (to *CreateIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateIdentityVisibilityFilterRequest_SdkV2) {
	if !from.IdentityVisibilityFilter.IsNull() && !from.IdentityVisibilityFilter.IsUnknown() {
		if toIdentityVisibilityFilter, ok := to.GetIdentityVisibilityFilter(ctx); ok {
			if fromIdentityVisibilityFilter, ok := from.GetIdentityVisibilityFilter(ctx); ok {
				toIdentityVisibilityFilter.SyncFieldsDuringRead(ctx, fromIdentityVisibilityFilter)
				to.SetIdentityVisibilityFilter(ctx, toIdentityVisibilityFilter)
			}
		}
	}
}

func (m CreateIdentityVisibilityFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["identity_visibility_filter"] = attrs["identity_visibility_filter"].SetRequired()
	attrs["identity_visibility_filter"] = attrs["identity_visibility_filter"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["identity_visibility_filter_id"] = attrs["identity_visibility_filter_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIdentityVisibilityFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIdentityVisibilityFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"identity_visibility_filter": reflect.TypeOf(IdentityVisibilityFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIdentityVisibilityFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateIdentityVisibilityFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"identity_visibility_filter":    m.IdentityVisibilityFilter,
			"identity_visibility_filter_id": m.IdentityVisibilityFilterId,
			"parent":                        m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIdentityVisibilityFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"identity_visibility_filter": basetypes.ListType{
				ElemType: IdentityVisibilityFilter_SdkV2{}.Type(ctx),
			},
			"identity_visibility_filter_id": types.StringType,
			"parent":                        types.StringType,
		},
	}
}

// GetIdentityVisibilityFilter returns the value of the IdentityVisibilityFilter field in CreateIdentityVisibilityFilterRequest_SdkV2 as
// a IdentityVisibilityFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIdentityVisibilityFilterRequest_SdkV2) GetIdentityVisibilityFilter(ctx context.Context) (IdentityVisibilityFilter_SdkV2, bool) {
	var e IdentityVisibilityFilter_SdkV2
	if m.IdentityVisibilityFilter.IsNull() || m.IdentityVisibilityFilter.IsUnknown() {
		return e, false
	}
	var v []IdentityVisibilityFilter_SdkV2
	d := m.IdentityVisibilityFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIdentityVisibilityFilter sets the value of the IdentityVisibilityFilter field in CreateIdentityVisibilityFilterRequest_SdkV2.
func (m *CreateIdentityVisibilityFilterRequest_SdkV2) SetIdentityVisibilityFilter(ctx context.Context, v IdentityVisibilityFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["identity_visibility_filter"]
	m.IdentityVisibilityFilter = types.ListValueMust(t, vs)
}

type CreateServicePrincipalProxyRequest_SdkV2 struct {
	// The resource name of the account to create the service principal in, in
	// the format `accounts/{account_id}`. Defaults to the account that contains
	// this workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
	// The service principal to create.
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
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"parent":            m.Parent,
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
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
	// The resource name of the account to create the service principal in, in
	// the format `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
	// The service principal to create.
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
	attrs["parent"] = attrs["parent"].SetRequired()

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
			"parent":            m.Parent,
			"service_principal": m.ServicePrincipal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
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
	// The resource name of the account to create the user in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
	// The user to create.
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
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"parent": m.Parent,
			"user":   m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
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
	// The resource name of the account to create the user in, in the format
	// `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
	// The user to create.
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
	attrs["parent"] = attrs["parent"].SetRequired()

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
			"parent": m.Parent,
			"user":   m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
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

type CreateWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. Workspace assignment detail to be created in Databricks.
	WorkspaceAssignmentDetail types.List `tfsdk:"workspace_assignment_detail"`
}

func (to *CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) {
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

func (to *CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentDetailProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentDetailProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment_detail": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailProxyRequest_SdkV2 as
// a WorkspaceAssignmentDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail_SdkV2, bool) {
	var e WorkspaceAssignmentDetail_SdkV2
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignmentDetail_SdkV2
	d := m.WorkspaceAssignmentDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailProxyRequest_SdkV2.
func (m *CreateWorkspaceAssignmentDetailProxyRequest_SdkV2) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_detail"]
	m.WorkspaceAssignmentDetail = types.ListValueMust(t, vs)
}

type CreateWorkspaceAssignmentDetailRequest_SdkV2 struct {
	// Required. Workspace assignment detail to be created in Databricks.
	WorkspaceAssignmentDetail types.List `tfsdk:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being created.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *CreateWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentDetailRequest_SdkV2) {
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

func (to *CreateWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentDetailRequest_SdkV2) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateWorkspaceAssignmentDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
			"workspace_id":                m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_assignment_detail": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailRequest_SdkV2 as
// a WorkspaceAssignmentDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentDetailRequest_SdkV2) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail_SdkV2, bool) {
	var e WorkspaceAssignmentDetail_SdkV2
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignmentDetail_SdkV2
	d := m.WorkspaceAssignmentDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in CreateWorkspaceAssignmentDetailRequest_SdkV2.
func (m *CreateWorkspaceAssignmentDetailRequest_SdkV2) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_detail"]
	m.WorkspaceAssignmentDetail = types.ListValueMust(t, vs)
}

type CreateWorkspaceAssignmentProxyRequest_SdkV2 struct {
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Defaults to this
	// workspace. Any other workspace is rejected.
	Parent types.String `tfsdk:"-"`
	// The assignment to create. Set its `principal_id` and `entitlements`.
	WorkspaceAssignment types.List `tfsdk:"workspace_assignment"`
}

func (to *CreateWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentProxyRequest_SdkV2) {
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

func (to *CreateWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentProxyRequest_SdkV2) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["workspace_assignment"] = attrs["workspace_assignment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":               m.Parent,
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"workspace_assignment": basetypes.ListType{
				ElemType: WorkspaceAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentProxyRequest_SdkV2 as
// a WorkspaceAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentProxyRequest_SdkV2) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment_SdkV2, bool) {
	var e WorkspaceAssignment_SdkV2
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignment_SdkV2
	d := m.WorkspaceAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentProxyRequest_SdkV2.
func (m *CreateWorkspaceAssignmentProxyRequest_SdkV2) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment"]
	m.WorkspaceAssignment = types.ListValueMust(t, vs)
}

type CreateWorkspaceAssignmentRequest_SdkV2 struct {
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`.
	Parent types.String `tfsdk:"-"`
	// The assignment to create. Set its `principal_id` and `entitlements`.
	WorkspaceAssignment types.List `tfsdk:"workspace_assignment"`
}

func (to *CreateWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceAssignmentRequest_SdkV2) {
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

func (to *CreateWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceAssignmentRequest_SdkV2) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m CreateWorkspaceAssignmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["workspace_assignment"] = attrs["workspace_assignment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":               m.Parent,
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"workspace_assignment": basetypes.ListType{
				ElemType: WorkspaceAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentRequest_SdkV2 as
// a WorkspaceAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceAssignmentRequest_SdkV2) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment_SdkV2, bool) {
	var e WorkspaceAssignment_SdkV2
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignment_SdkV2
	d := m.WorkspaceAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in CreateWorkspaceAssignmentRequest_SdkV2.
func (m *CreateWorkspaceAssignmentRequest_SdkV2) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment"]
	m.WorkspaceAssignment = types.ListValueMust(t, vs)
}

type DeleteDirectGroupMemberProxyRequest_SdkV2 struct {
	// The resource name of the membership, in the format
	// `groups/{group_id}/direct-members/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDirectGroupMemberProxyRequest_SdkV2) {
}

func (to *DeleteDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDirectGroupMemberProxyRequest_SdkV2) {
}

func (m DeleteDirectGroupMemberProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDirectGroupMemberProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectGroupMemberProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDirectGroupMemberProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDirectGroupMemberProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDirectGroupMemberRequest_SdkV2 struct {
	// The resource name of the membership, in the format
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDirectGroupMemberRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDirectGroupMemberRequest_SdkV2) {
}

func (to *DeleteDirectGroupMemberRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDirectGroupMemberRequest_SdkV2) {
}

func (m DeleteDirectGroupMemberRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDirectGroupMemberRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectGroupMemberRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDirectGroupMemberRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDirectGroupMemberRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteGroupProxyRequest_SdkV2 struct {
	// The resource name of the group, in the format `groups/{group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupProxyRequest_SdkV2) {
}

func (to *DeleteGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupProxyRequest_SdkV2) {
}

func (m DeleteGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteGroupRequest_SdkV2 struct {
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGroupRequest_SdkV2) {
}

func (to *DeleteGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteGroupRequest_SdkV2) {
}

func (m DeleteGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteIdentityVisibilityFilterRequest_SdkV2 struct {
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIdentityVisibilityFilterRequest_SdkV2) {
}

func (to *DeleteIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteIdentityVisibilityFilterRequest_SdkV2) {
}

func (m DeleteIdentityVisibilityFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIdentityVisibilityFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIdentityVisibilityFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIdentityVisibilityFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteIdentityVisibilityFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIdentityVisibilityFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteServicePrincipalProxyRequest_SdkV2 struct {
	// The resource name of the service principal, in the format
	// `service-principals/{service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalProxyRequest_SdkV2) {
}

func (to *DeleteServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalProxyRequest_SdkV2) {
}

func (m DeleteServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteServicePrincipalRequest_SdkV2 struct {
	// The resource name of the service principal, in the format
	// `accounts/{account_id}/service-principals/{service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalRequest_SdkV2) {
}

func (to *DeleteServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalRequest_SdkV2) {
}

func (m DeleteServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteUserProxyRequest_SdkV2 struct {
	// The resource name of the user, in the format `users/{user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserProxyRequest_SdkV2) {
}

func (to *DeleteUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteUserProxyRequest_SdkV2) {
}

func (m DeleteUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteUserRequest_SdkV2 struct {
	// The resource name of the user, in the format
	// `accounts/{account_id}/users/{user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteUserRequest_SdkV2) {
}

func (to *DeleteUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteUserRequest_SdkV2) {
}

func (m DeleteUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) {
}

func (to *DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) {
}

func (m DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentDetailProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAssignmentDetailRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentDetailRequest_SdkV2) {
}

func (to *DeleteWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentDetailRequest_SdkV2) {
}

func (m DeleteWorkspaceAssignmentDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWorkspaceAssignmentDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspaceAssignmentProxyRequest_SdkV2 struct {
	// The resource name of the assignment, in the format
	// `workspace-assignments/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentProxyRequest_SdkV2) {
}

func (to *DeleteWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentProxyRequest_SdkV2) {
}

func (m DeleteWorkspaceAssignmentProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteWorkspaceAssignmentRequest_SdkV2 struct {
	// The resource name of the assignment, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceAssignmentRequest_SdkV2) {
}

func (to *DeleteWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceAssignmentRequest_SdkV2) {
}

func (m DeleteWorkspaceAssignmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A user, service principal, or group that's a direct member of a group.
type DirectGroupMember_SdkV2 struct {
	// The member's display name.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the member in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The Databricks ID of the group.
	GroupId types.Int64 `tfsdk:"group_id"`
	// Where the membership comes from.
	MembershipSource types.String `tfsdk:"membership_source"`
	// The resource name of the membership, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}` -
	// Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the member.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// Whether the member is a user, service principal, or group.
	PrincipalType types.String `tfsdk:"principal_type"`
}

func (to *DirectGroupMember_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectGroupMember_SdkV2) {
}

func (to *DirectGroupMember_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DirectGroupMember_SdkV2) {
}

func (m DirectGroupMember_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["membership_source"] = attrs["membership_source"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
func (m DirectGroupMember_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectGroupMember_SdkV2
// only implements ToObjectValue() and Type().
func (m DirectGroupMember_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":      m.DisplayName,
			"external_id":       m.ExternalId,
			"group_id":          m.GroupId,
			"membership_source": m.MembershipSource,
			"name":              m.Name,
			"parent":            m.Parent,
			"principal_id":      m.PrincipalId,
			"principal_type":    m.PrincipalType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectGroupMember_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":      types.StringType,
			"external_id":       types.StringType,
			"group_id":          types.Int64Type,
			"membership_source": types.StringType,
			"name":              types.StringType,
			"parent":            types.StringType,
			"principal_id":      types.Int64Type,
			"principal_type":    types.StringType,
		},
	}
}

// A rule that matches a group by its ID in your identity provider (IdP).
type ExactMatchRule_SdkV2 struct {
	// The ID of the group in your IdP, not its Databricks ID. It's matched
	// exactly, including case.
	ExternalId types.String `tfsdk:"external_id"`
}

func (to *ExactMatchRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExactMatchRule_SdkV2) {
}

func (to *ExactMatchRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExactMatchRule_SdkV2) {
}

func (m ExactMatchRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["external_id"] = attrs["external_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExactMatchRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExactMatchRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExactMatchRule_SdkV2
// only implements ToObjectValue() and Type().
func (m ExactMatchRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": m.ExternalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExactMatchRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
		},
	}
}

// A group from your identity provider (IdP), identified by its ID in your IdP.
// This resource is read-only. Getting an external group syncs its name and
// parent groups from your IdP, and adds it to your account if it isn't already
// there. If your account has group visibility filters, the group must match
// them.
type ExternalGroup_SdkV2 struct {
	// The ID of the account that contains the group.
	AccountId types.String `tfsdk:"account_id"`
	// The name of the group, from your IdP. It's the same as the group's
	// `group_name`.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the group in your IdP.
	ExternalGroupId types.String `tfsdk:"external_group_id"`
	// The Databricks ID of the group. Use it as the `group_id` in the group
	// APIs.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external group, in one of these formats: -
	// Account APIs: `accounts/{account_id}/external-groups/{external_group_id}`
	// - Workspace APIs: `external-groups/{external_group_id}`
	Name types.String `tfsdk:"name"`
}

func (to *ExternalGroup_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalGroup_SdkV2) {
}

func (to *ExternalGroup_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalGroup_SdkV2) {
}

func (m ExternalGroup_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalGroup_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalGroup_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalGroup_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExternalGroup_SdkV2) Type(ctx context.Context) attr.Type {
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

// A service principal from your identity provider (IdP), identified by its ID
// in your IdP. This resource is read-only and requires Microsoft Entra ID as
// your IdP. Getting an external service principal syncs its details and group
// memberships from your IdP, and adds it to your account if it isn't already
// there.
type ExternalServicePrincipal_SdkV2 struct {
	// The ID of the account that contains the service principal.
	AccountId types.String `tfsdk:"account_id"`
	// Whether the service principal is active in the account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// The service principal's application ID, from your IdP.
	ApplicationId types.String `tfsdk:"application_id"`
	// The service principal's display name, from your IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the service principal in your IdP.
	ExternalServicePrincipalId types.String `tfsdk:"external_service_principal_id"`
	// The Databricks ID of the service principal. Use it as the
	// `service_principal_id` in the service principal APIs.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external service principal, in one of these
	// formats: - Account APIs:
	// `accounts/{account_id}/external-service-principals/{external_service_principal_id}`
	// - Workspace APIs:
	// `external-service-principals/{external_service_principal_id}`
	Name types.String `tfsdk:"name"`
}

func (to *ExternalServicePrincipal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalServicePrincipal_SdkV2) {
}

func (to *ExternalServicePrincipal_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalServicePrincipal_SdkV2) {
}

func (m ExternalServicePrincipal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalServicePrincipal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalServicePrincipal_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalServicePrincipal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExternalServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
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

// A user from your identity provider (IdP), identified by their ID in your IdP.
// This resource is read-only. Getting an external user syncs the user's details
// and group memberships from your IdP, and adds the user to your account if
// they aren't already there.
type ExternalUser_SdkV2 struct {
	// The ID of the account that contains the user.
	AccountId types.String `tfsdk:"account_id"`
	// Whether the user is active in the account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// The user's given and family names from your IdP, separated by a space.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the user in your IdP.
	ExternalUserId types.String `tfsdk:"external_user_id"`
	// The user's name, from your IdP.
	FullName types.List `tfsdk:"full_name"`
	// The Databricks ID of the user. Use it as the `user_id` in the user APIs.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external user, in one of these formats: -
	// Account APIs: `accounts/{account_id}/external-users/{external_user_id}` -
	// Workspace APIs: `external-users/{external_user_id}`
	Name types.String `tfsdk:"name"`
	// The user's username, which is their email address.
	Username types.String `tfsdk:"username"`
}

func (to *ExternalUser_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalUser_SdkV2) {
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

func (to *ExternalUser_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalUser_SdkV2) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				toFullName.SyncFieldsDuringRead(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (m ExternalUser_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_user_id"] = attrs["external_user_id"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["full_name"] = attrs["full_name"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m ExternalUser_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name": reflect.TypeOf(FullName_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalUser_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalUser_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExternalUser_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"display_name":        types.StringType,
			"external_user_id":    types.StringType,
			"full_name": basetypes.ListType{
				ElemType: FullName_SdkV2{}.Type(ctx),
			},
			"internal_id": types.StringType,
			"name":        types.StringType,
			"username":    types.StringType,
		},
	}
}

// GetFullName returns the value of the FullName field in ExternalUser_SdkV2 as
// a FullName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalUser_SdkV2) GetFullName(ctx context.Context) (FullName_SdkV2, bool) {
	var e FullName_SdkV2
	if m.FullName.IsNull() || m.FullName.IsUnknown() {
		return e, false
	}
	var v []FullName_SdkV2
	d := m.FullName.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFullName sets the value of the FullName field in ExternalUser_SdkV2.
func (m *ExternalUser_SdkV2) SetFullName(ctx context.Context, v FullName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_name"]
	m.FullName = types.ListValueMust(t, vs)
}

// A user's name, from your identity provider (IdP).
type FullName_SdkV2 struct {
	// The user's family (last) name.
	FamilyName types.String `tfsdk:"family_name"`
	// The user's given (first) name.
	GivenName types.String `tfsdk:"given_name"`
}

func (to *FullName_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FullName_SdkV2) {
}

func (to *FullName_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FullName_SdkV2) {
}

func (m FullName_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FullName_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FullName_SdkV2
// only implements ToObjectValue() and Type().
func (m FullName_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FullName_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"family_name": types.StringType,
			"given_name":  types.StringType,
		},
	}
}

type GetDirectGroupMemberProxyRequest_SdkV2 struct {
	// The resource name of the membership, in the format
	// `groups/{group_id}/direct-members/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDirectGroupMemberProxyRequest_SdkV2) {
}

func (to *GetDirectGroupMemberProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDirectGroupMemberProxyRequest_SdkV2) {
}

func (m GetDirectGroupMemberProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectGroupMemberProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDirectGroupMemberProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectGroupMemberProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDirectGroupMemberProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDirectGroupMemberProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDirectGroupMemberRequest_SdkV2 struct {
	// The resource name of the membership, in the format
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetDirectGroupMemberRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDirectGroupMemberRequest_SdkV2) {
}

func (to *GetDirectGroupMemberRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDirectGroupMemberRequest_SdkV2) {
}

func (m GetDirectGroupMemberRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectGroupMemberRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDirectGroupMemberRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectGroupMemberRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDirectGroupMemberRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDirectGroupMemberRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalGroupProxyRequest_SdkV2 struct {
	// The resource name of the external group, in the format
	// `external-groups/{external_group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalGroupProxyRequest_SdkV2) {
}

func (to *GetExternalGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalGroupProxyRequest_SdkV2) {
}

func (m GetExternalGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalGroupProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalGroupProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalGroupProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalGroupRequest_SdkV2 struct {
	// The resource name of the external group, in the format
	// `accounts/{account_id}/external-groups/{external_group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalGroupRequest_SdkV2) {
}

func (to *GetExternalGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalGroupRequest_SdkV2) {
}

func (m GetExternalGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalGroupRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalGroupRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalGroupRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalServicePrincipalProxyRequest_SdkV2 struct {
	// The resource name of the external service principal, in the format
	// `external-service-principals/{external_service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalServicePrincipalProxyRequest_SdkV2) {
}

func (to *GetExternalServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalServicePrincipalProxyRequest_SdkV2) {
}

func (m GetExternalServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalServicePrincipalProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalServicePrincipalProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalServicePrincipalProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalServicePrincipalRequest_SdkV2 struct {
	// The resource name of the external service principal, in the format
	// `accounts/{account_id}/external-service-principals/{external_service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalServicePrincipalRequest_SdkV2) {
}

func (to *GetExternalServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalServicePrincipalRequest_SdkV2) {
}

func (m GetExternalServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalServicePrincipalRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalServicePrincipalRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalServicePrincipalRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalUserProxyRequest_SdkV2 struct {
	// The resource name of the external user, in the format
	// `external-users/{external_user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalUserProxyRequest_SdkV2) {
}

func (to *GetExternalUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalUserProxyRequest_SdkV2) {
}

func (m GetExternalUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalUserProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalUserProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalUserProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetExternalUserRequest_SdkV2 struct {
	// The resource name of the external user, in the format
	// `accounts/{account_id}/external-users/{external_user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetExternalUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExternalUserRequest_SdkV2) {
}

func (to *GetExternalUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExternalUserRequest_SdkV2) {
}

func (m GetExternalUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExternalUserRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalUserRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExternalUserRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExternalUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetGroupProxyRequest_SdkV2 struct {
	// The resource name of the group, in the format `groups/{group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetGroupProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupProxyRequest_SdkV2) {
}

func (to *GetGroupProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetGroupProxyRequest_SdkV2) {
}

func (m GetGroupProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetGroupRequest_SdkV2 struct {
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetGroupRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGroupRequest_SdkV2) {
}

func (to *GetGroupRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetGroupRequest_SdkV2) {
}

func (m GetGroupRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGroupRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetIdentityVisibilityFilterRequest_SdkV2 struct {
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIdentityVisibilityFilterRequest_SdkV2) {
}

func (to *GetIdentityVisibilityFilterRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIdentityVisibilityFilterRequest_SdkV2) {
}

func (m GetIdentityVisibilityFilterRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIdentityVisibilityFilterRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIdentityVisibilityFilterRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIdentityVisibilityFilterRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIdentityVisibilityFilterRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIdentityVisibilityFilterRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetServicePrincipalProxyRequest_SdkV2 struct {
	// The resource name of the service principal, in the format
	// `service-principals/{service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalProxyRequest_SdkV2) {
}

func (to *GetServicePrincipalProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalProxyRequest_SdkV2) {
}

func (m GetServicePrincipalProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetServicePrincipalRequest_SdkV2 struct {
	// The resource name of the service principal, in the format
	// `accounts/{account_id}/service-principals/{service_principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetServicePrincipalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalRequest_SdkV2) {
}

func (to *GetServicePrincipalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalRequest_SdkV2) {
}

func (m GetServicePrincipalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetUserProxyRequest_SdkV2 struct {
	// The resource name of the user, in the format `users/{user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetUserProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserProxyRequest_SdkV2) {
}

func (to *GetUserProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUserProxyRequest_SdkV2) {
}

func (m GetUserProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetUserRequest_SdkV2 struct {
	// The resource name of the user, in the format
	// `accounts/{account_id}/users/{user_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetUserRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUserRequest_SdkV2) {
}

func (to *GetUserRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUserRequest_SdkV2) {
}

func (m GetUserRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

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
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetWorkspaceAccessDetailLocalRequest_SdkV2 struct {
	// The resource name of the workspace access detail. Format:
	// workspace-access-details/{principal_id}
	Name types.String `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (to *GetWorkspaceAccessDetailLocalRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailLocalRequest_SdkV2) {
}

func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
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
			"name": m.Name,
			"view": m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailLocalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"view": types.StringType,
		},
	}
}

type GetWorkspaceAccessDetailRequest_SdkV2 struct {
	// The resource name of the workspace access detail. Format:
	// accounts/{account_id}/workspaces/{workspace_id}/workspace-access-details/{principal_id}
	Name types.String `tfsdk:"-"`
	// Controls what fields are returned.
	View types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAccessDetailRequest_SdkV2) {
}

func (to *GetWorkspaceAccessDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAccessDetailRequest_SdkV2) {
}

func (m GetWorkspaceAccessDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
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
			"name": m.Name,
			"view": m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAccessDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"view": types.StringType,
		},
	}
}

type GetWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentDetailProxyRequest_SdkV2) {
}

func (to *GetWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentDetailProxyRequest_SdkV2) {
}

func (m GetWorkspaceAssignmentDetailProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceAssignmentDetailProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentDetailProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentDetailProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentDetailProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentDetailRequest_SdkV2 struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The workspace ID for which the assignment details are being
	// requested.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentDetailRequest_SdkV2) {
}

func (to *GetWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentDetailRequest_SdkV2) {
}

func (m GetWorkspaceAssignmentDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceAssignmentDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id": m.PrincipalId,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type GetWorkspaceAssignmentProxyRequest_SdkV2 struct {
	// The resource name of the assignment, in the format
	// `workspace-assignments/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentProxyRequest_SdkV2) {
}

func (to *GetWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentProxyRequest_SdkV2) {
}

func (m GetWorkspaceAssignmentProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetWorkspaceAssignmentRequest_SdkV2 struct {
	// The resource name of the assignment, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceAssignmentRequest_SdkV2) {
}

func (to *GetWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceAssignmentRequest_SdkV2) {
}

func (m GetWorkspaceAssignmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetWorkspaceIdentityDetailRequest_SdkV2 struct {
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name types.String `tfsdk:"-"`
}

func (to *GetWorkspaceIdentityDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceIdentityDetailRequest_SdkV2) {
}

func (to *GetWorkspaceIdentityDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceIdentityDetailRequest_SdkV2) {
}

func (m GetWorkspaceIdentityDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceIdentityDetailRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceIdentityDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceIdentityDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceIdentityDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceIdentityDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A group in your Databricks account.
type Group_SdkV2 struct {
	// The ID of the account that contains the group.
	AccountId types.String `tfsdk:"account_id"`
	// The ID of the group in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The Databricks ID of the group.
	GroupId types.String `tfsdk:"group_id"`
	// The name of the group.
	GroupName types.String `tfsdk:"group_name"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the account that contains the group, in the format
	// `accounts/{account_id}`. Workspace APIs return the same format.
	Parent types.String `tfsdk:"parent"`
}

func (to *Group_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Group_SdkV2) {
}

func (to *Group_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Group_SdkV2) {
}

func (m Group_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

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
			"group_id":    m.GroupId,
			"group_name":  m.GroupName,
			"name":        m.Name,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Group_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_id":    types.StringType,
			"group_name":  types.StringType,
			"name":        types.StringType,
			"parent":      types.StringType,
		},
	}
}

// An account-level rule that controls which groups from your identity provider
// (IdP) can be found in directory search and group pickers when automatic
// identity management is on. Filters apply to every workspace in the account.
//
// If you have no group filters, all groups are visible. Otherwise, a group is
// visible if it or any of its parent groups matches at least one filter. A
// group is also visible if it or any of its parent groups is already in your
// account. Local Databricks groups are always visible.
//
// You can only create group filters. Filters don't affect provisioning or
// sign-in for users and service principals. You can't change a filter: to
// change a rule, delete it and create a new one.
type IdentityVisibilityFilter_SdkV2 struct {
	// Matches a group by its ID in your IdP.
	Exact types.List `tfsdk:"exact"`
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name types.String `tfsdk:"name"`
	// Matches groups whose display name starts with a prefix.
	Prefix types.List `tfsdk:"prefix"`
	// The type of principal the filter applies to. Set it to `GROUP`. Other
	// types aren't supported.
	PrincipalType types.String `tfsdk:"principal_type"`
	// Matches groups whose display name matches a regular expression.
	Regex types.List `tfsdk:"regex"`
}

func (to *IdentityVisibilityFilter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IdentityVisibilityFilter_SdkV2) {
	if !from.Exact.IsNull() && !from.Exact.IsUnknown() {
		if toExact, ok := to.GetExact(ctx); ok {
			if fromExact, ok := from.GetExact(ctx); ok {
				// Recursively sync the fields of Exact
				toExact.SyncFieldsDuringCreateOrUpdate(ctx, fromExact)
				to.SetExact(ctx, toExact)
			}
		}
	}
	if !from.Prefix.IsNull() && !from.Prefix.IsUnknown() {
		if toPrefix, ok := to.GetPrefix(ctx); ok {
			if fromPrefix, ok := from.GetPrefix(ctx); ok {
				// Recursively sync the fields of Prefix
				toPrefix.SyncFieldsDuringCreateOrUpdate(ctx, fromPrefix)
				to.SetPrefix(ctx, toPrefix)
			}
		}
	}
	if !from.Regex.IsNull() && !from.Regex.IsUnknown() {
		if toRegex, ok := to.GetRegex(ctx); ok {
			if fromRegex, ok := from.GetRegex(ctx); ok {
				// Recursively sync the fields of Regex
				toRegex.SyncFieldsDuringCreateOrUpdate(ctx, fromRegex)
				to.SetRegex(ctx, toRegex)
			}
		}
	}
}

func (to *IdentityVisibilityFilter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IdentityVisibilityFilter_SdkV2) {
	if !from.Exact.IsNull() && !from.Exact.IsUnknown() {
		if toExact, ok := to.GetExact(ctx); ok {
			if fromExact, ok := from.GetExact(ctx); ok {
				toExact.SyncFieldsDuringRead(ctx, fromExact)
				to.SetExact(ctx, toExact)
			}
		}
	}
	if !from.Prefix.IsNull() && !from.Prefix.IsUnknown() {
		if toPrefix, ok := to.GetPrefix(ctx); ok {
			if fromPrefix, ok := from.GetPrefix(ctx); ok {
				toPrefix.SyncFieldsDuringRead(ctx, fromPrefix)
				to.SetPrefix(ctx, toPrefix)
			}
		}
	}
	if !from.Regex.IsNull() && !from.Regex.IsUnknown() {
		if toRegex, ok := to.GetRegex(ctx); ok {
			if fromRegex, ok := from.GetRegex(ctx); ok {
				toRegex.SyncFieldsDuringRead(ctx, fromRegex)
				to.SetRegex(ctx, toRegex)
			}
		}
	}
}

func (m IdentityVisibilityFilter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exact"] = attrs["exact"].SetOptional()
	attrs["exact"] = attrs["exact"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["exact"] = attrs["exact"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["prefix"] = attrs["prefix"].SetOptional()
	attrs["prefix"] = attrs["prefix"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["prefix"] = attrs["prefix"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["principal_type"] = attrs["principal_type"].SetRequired()
	attrs["principal_type"] = attrs["principal_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["regex"] = attrs["regex"].SetOptional()
	attrs["regex"] = attrs["regex"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IdentityVisibilityFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IdentityVisibilityFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exact":  reflect.TypeOf(ExactMatchRule_SdkV2{}),
		"prefix": reflect.TypeOf(PrefixMatchRule_SdkV2{}),
		"regex":  reflect.TypeOf(RegexMatchRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IdentityVisibilityFilter_SdkV2
// only implements ToObjectValue() and Type().
func (m IdentityVisibilityFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exact":          m.Exact,
			"name":           m.Name,
			"prefix":         m.Prefix,
			"principal_type": m.PrincipalType,
			"regex":          m.Regex,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IdentityVisibilityFilter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exact": basetypes.ListType{
				ElemType: ExactMatchRule_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"prefix": basetypes.ListType{
				ElemType: PrefixMatchRule_SdkV2{}.Type(ctx),
			},
			"principal_type": types.StringType,
			"regex": basetypes.ListType{
				ElemType: RegexMatchRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetExact returns the value of the Exact field in IdentityVisibilityFilter_SdkV2 as
// a ExactMatchRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IdentityVisibilityFilter_SdkV2) GetExact(ctx context.Context) (ExactMatchRule_SdkV2, bool) {
	var e ExactMatchRule_SdkV2
	if m.Exact.IsNull() || m.Exact.IsUnknown() {
		return e, false
	}
	var v []ExactMatchRule_SdkV2
	d := m.Exact.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExact sets the value of the Exact field in IdentityVisibilityFilter_SdkV2.
func (m *IdentityVisibilityFilter_SdkV2) SetExact(ctx context.Context, v ExactMatchRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exact"]
	m.Exact = types.ListValueMust(t, vs)
}

// GetPrefix returns the value of the Prefix field in IdentityVisibilityFilter_SdkV2 as
// a PrefixMatchRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IdentityVisibilityFilter_SdkV2) GetPrefix(ctx context.Context) (PrefixMatchRule_SdkV2, bool) {
	var e PrefixMatchRule_SdkV2
	if m.Prefix.IsNull() || m.Prefix.IsUnknown() {
		return e, false
	}
	var v []PrefixMatchRule_SdkV2
	d := m.Prefix.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrefix sets the value of the Prefix field in IdentityVisibilityFilter_SdkV2.
func (m *IdentityVisibilityFilter_SdkV2) SetPrefix(ctx context.Context, v PrefixMatchRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["prefix"]
	m.Prefix = types.ListValueMust(t, vs)
}

// GetRegex returns the value of the Regex field in IdentityVisibilityFilter_SdkV2 as
// a RegexMatchRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IdentityVisibilityFilter_SdkV2) GetRegex(ctx context.Context) (RegexMatchRule_SdkV2, bool) {
	var e RegexMatchRule_SdkV2
	if m.Regex.IsNull() || m.Regex.IsUnknown() {
		return e, false
	}
	var v []RegexMatchRule_SdkV2
	d := m.Regex.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegex sets the value of the Regex field in IdentityVisibilityFilter_SdkV2.
func (m *IdentityVisibilityFilter_SdkV2) SetRegex(ctx context.Context, v RegexMatchRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["regex"]
	m.Regex = types.ListValueMust(t, vs)
}

type ListDirectGroupMembersProxyRequest_SdkV2 struct {
	// The maximum number of members to return, up to 1000. The default is 1000.
	// A page can contain fewer members.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the group, in the format `groups/{group_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListDirectGroupMembersProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersProxyRequest_SdkV2) {
}

func (to *ListDirectGroupMembersProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersProxyRequest_SdkV2) {
}

func (m ListDirectGroupMembersProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
func (m ListDirectGroupMembersProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListDirectGroupMembersRequest_SdkV2 struct {
	// The maximum number of members to return, up to 1000. The default is 1000.
	// A page can contain fewer members.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListDirectGroupMembersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersRequest_SdkV2) {
}

func (to *ListDirectGroupMembersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersRequest_SdkV2) {
}

func (m ListDirectGroupMembersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
func (m ListDirectGroupMembersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of direct members of a group.
type ListDirectGroupMembersResponse_SdkV2 struct {
	// The members on this page.
	DirectGroupMembers types.List `tfsdk:"direct_group_members"`
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDirectGroupMembersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectGroupMembersResponse_SdkV2) {
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

func (to *ListDirectGroupMembersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDirectGroupMembersResponse_SdkV2) {
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

func (m ListDirectGroupMembersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDirectGroupMembersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_members": reflect.TypeOf(DirectGroupMember_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectGroupMembersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDirectGroupMembersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"direct_group_members": m.DirectGroupMembers,
			"next_page_token":      m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectGroupMembersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"direct_group_members": basetypes.ListType{
				ElemType: DirectGroupMember_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDirectGroupMembers returns the value of the DirectGroupMembers field in ListDirectGroupMembersResponse_SdkV2 as
// a slice of DirectGroupMember_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDirectGroupMembersResponse_SdkV2) GetDirectGroupMembers(ctx context.Context) ([]DirectGroupMember_SdkV2, bool) {
	if m.DirectGroupMembers.IsNull() || m.DirectGroupMembers.IsUnknown() {
		return nil, false
	}
	var v []DirectGroupMember_SdkV2
	d := m.DirectGroupMembers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectGroupMembers sets the value of the DirectGroupMembers field in ListDirectGroupMembersResponse_SdkV2.
func (m *ListDirectGroupMembersResponse_SdkV2) SetDirectGroupMembers(ctx context.Context, v []DirectGroupMember_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_group_members"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DirectGroupMembers = types.ListValueMust(t, vs)
}

type ListGroupsProxyRequest_SdkV2 struct {
	// Filters groups by `group_name` or `external_id`, for example `group_name
	// eq "Data engineers"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list groups in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
}

func (to *ListGroupsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsProxyRequest_SdkV2) {
}

func (to *ListGroupsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsProxyRequest_SdkV2) {
}

func (m ListGroupsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListGroupsRequest_SdkV2 struct {
	// Filters groups by `group_name` or `external_id`, for example `group_name
	// eq "Data engineers"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list groups in, in the format
	// `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListGroupsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsRequest_SdkV2) {
}

func (to *ListGroupsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsRequest_SdkV2) {
}

func (m ListGroupsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGroupsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of groups.
type ListGroupsResponse_SdkV2 struct {
	// The groups on this page.
	Groups types.List `tfsdk:"groups"`
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListGroupsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGroupsResponse_SdkV2) {
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

func (to *ListGroupsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListGroupsResponse_SdkV2) {
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

type ListIdentityVisibilityFiltersRequest_SdkV2 struct {
	// The maximum number of filters to return, up to 1000. The default is 100.
	// A page can contain fewer filters.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account, in the format `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListIdentityVisibilityFiltersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIdentityVisibilityFiltersRequest_SdkV2) {
}

func (to *ListIdentityVisibilityFiltersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIdentityVisibilityFiltersRequest_SdkV2) {
}

func (m ListIdentityVisibilityFiltersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIdentityVisibilityFiltersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIdentityVisibilityFiltersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIdentityVisibilityFiltersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIdentityVisibilityFiltersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIdentityVisibilityFiltersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of group visibility filters.
type ListIdentityVisibilityFiltersResponse_SdkV2 struct {
	// The filters on this page. These are the rules, not the groups they match.
	IdentityVisibilityFilters types.List `tfsdk:"identity_visibility_filters"`
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListIdentityVisibilityFiltersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIdentityVisibilityFiltersResponse_SdkV2) {
	if !from.IdentityVisibilityFilters.IsNull() && !from.IdentityVisibilityFilters.IsUnknown() && to.IdentityVisibilityFilters.IsNull() && len(from.IdentityVisibilityFilters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IdentityVisibilityFilters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IdentityVisibilityFilters = from.IdentityVisibilityFilters
	}
	if !from.IdentityVisibilityFilters.IsNull() && !from.IdentityVisibilityFilters.IsUnknown() {
		if toIdentityVisibilityFilters, ok := to.GetIdentityVisibilityFilters(ctx); ok {
			if fromIdentityVisibilityFilters, ok := from.GetIdentityVisibilityFilters(ctx); ok {
				// Recursively sync the fields of each IdentityVisibilityFilters element by position.
				for i := range toIdentityVisibilityFilters {
					if i < len(fromIdentityVisibilityFilters) {
						toIdentityVisibilityFilters[i].SyncFieldsDuringCreateOrUpdate(ctx, fromIdentityVisibilityFilters[i])
					}
				}
				to.SetIdentityVisibilityFilters(ctx, toIdentityVisibilityFilters)
			}
		}
	}
}

func (to *ListIdentityVisibilityFiltersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIdentityVisibilityFiltersResponse_SdkV2) {
	if !from.IdentityVisibilityFilters.IsNull() && !from.IdentityVisibilityFilters.IsUnknown() && to.IdentityVisibilityFilters.IsNull() && len(from.IdentityVisibilityFilters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IdentityVisibilityFilters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IdentityVisibilityFilters = from.IdentityVisibilityFilters
	}
	if !from.IdentityVisibilityFilters.IsNull() && !from.IdentityVisibilityFilters.IsUnknown() {
		if toIdentityVisibilityFilters, ok := to.GetIdentityVisibilityFilters(ctx); ok {
			if fromIdentityVisibilityFilters, ok := from.GetIdentityVisibilityFilters(ctx); ok {
				for i := range toIdentityVisibilityFilters {
					if i < len(fromIdentityVisibilityFilters) {
						toIdentityVisibilityFilters[i].SyncFieldsDuringRead(ctx, fromIdentityVisibilityFilters[i])
					}
				}
				to.SetIdentityVisibilityFilters(ctx, toIdentityVisibilityFilters)
			}
		}
	}
}

func (m ListIdentityVisibilityFiltersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["identity_visibility_filters"] = attrs["identity_visibility_filters"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIdentityVisibilityFiltersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIdentityVisibilityFiltersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"identity_visibility_filters": reflect.TypeOf(IdentityVisibilityFilter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIdentityVisibilityFiltersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIdentityVisibilityFiltersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"identity_visibility_filters": m.IdentityVisibilityFilters,
			"next_page_token":             m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIdentityVisibilityFiltersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"identity_visibility_filters": basetypes.ListType{
				ElemType: IdentityVisibilityFilter_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetIdentityVisibilityFilters returns the value of the IdentityVisibilityFilters field in ListIdentityVisibilityFiltersResponse_SdkV2 as
// a slice of IdentityVisibilityFilter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListIdentityVisibilityFiltersResponse_SdkV2) GetIdentityVisibilityFilters(ctx context.Context) ([]IdentityVisibilityFilter_SdkV2, bool) {
	if m.IdentityVisibilityFilters.IsNull() || m.IdentityVisibilityFilters.IsUnknown() {
		return nil, false
	}
	var v []IdentityVisibilityFilter_SdkV2
	d := m.IdentityVisibilityFilters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIdentityVisibilityFilters sets the value of the IdentityVisibilityFilters field in ListIdentityVisibilityFiltersResponse_SdkV2.
func (m *ListIdentityVisibilityFiltersResponse_SdkV2) SetIdentityVisibilityFilters(ctx context.Context, v []IdentityVisibilityFilter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["identity_visibility_filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IdentityVisibilityFilters = types.ListValueMust(t, vs)
}

type ListServicePrincipalsProxyRequest_SdkV2 struct {
	// Filters service principals by `application_id` or `external_id`, for
	// example `application_id eq "<application-id>"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of service principals to return, up to 1000. The
	// default is 1000. A page can contain fewer service principals.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list service principals in, in the
	// format `accounts/{account_id}`. Defaults to the account that contains
	// this workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsProxyRequest_SdkV2) {
}

func (to *ListServicePrincipalsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsProxyRequest_SdkV2) {
}

func (m ListServicePrincipalsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListServicePrincipalsRequest_SdkV2 struct {
	// Filters service principals by `application_id` or `external_id`, for
	// example `application_id eq "<application-id>"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of service principals to return, up to 1000. The
	// default is 1000. A page can contain fewer service principals.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list service principals in, in the
	// format `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsRequest_SdkV2) {
}

func (to *ListServicePrincipalsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsRequest_SdkV2) {
}

func (m ListServicePrincipalsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of service principals.
type ListServicePrincipalsResponse_SdkV2 struct {
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The service principals on this page.
	ServicePrincipals types.List `tfsdk:"service_principals"`
}

func (to *ListServicePrincipalsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalsResponse_SdkV2) {
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

func (to *ListServicePrincipalsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalsResponse_SdkV2) {
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

type ListTransitiveParentGroupsProxyRequest_SdkV2 struct {
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The user, service principal, or group whose parent groups to list, in the
	// format `principals/{principal_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListTransitiveParentGroupsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsProxyRequest_SdkV2) {
}

func (to *ListTransitiveParentGroupsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsProxyRequest_SdkV2) {
}

func (m ListTransitiveParentGroupsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
func (m ListTransitiveParentGroupsProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListTransitiveParentGroupsRequest_SdkV2 struct {
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The user, service principal, or group whose parent groups to list, in the
	// format `accounts/{account_id}/principals/{principal_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListTransitiveParentGroupsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsRequest_SdkV2) {
}

func (to *ListTransitiveParentGroupsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsRequest_SdkV2) {
}

func (m ListTransitiveParentGroupsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
func (m ListTransitiveParentGroupsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of the groups a principal belongs to.
type ListTransitiveParentGroupsResponse_SdkV2 struct {
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The groups on this page.
	TransitiveParentGroups types.List `tfsdk:"transitive_parent_groups"`
}

func (to *ListTransitiveParentGroupsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitiveParentGroupsResponse_SdkV2) {
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

func (to *ListTransitiveParentGroupsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTransitiveParentGroupsResponse_SdkV2) {
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

func (m ListTransitiveParentGroupsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTransitiveParentGroupsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"transitive_parent_groups": reflect.TypeOf(TransitiveParentGroup_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitiveParentGroupsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTransitiveParentGroupsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"transitive_parent_groups": m.TransitiveParentGroups,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitiveParentGroupsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"transitive_parent_groups": basetypes.ListType{
				ElemType: TransitiveParentGroup_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTransitiveParentGroups returns the value of the TransitiveParentGroups field in ListTransitiveParentGroupsResponse_SdkV2 as
// a slice of TransitiveParentGroup_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTransitiveParentGroupsResponse_SdkV2) GetTransitiveParentGroups(ctx context.Context) ([]TransitiveParentGroup_SdkV2, bool) {
	if m.TransitiveParentGroups.IsNull() || m.TransitiveParentGroups.IsUnknown() {
		return nil, false
	}
	var v []TransitiveParentGroup_SdkV2
	d := m.TransitiveParentGroups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTransitiveParentGroups sets the value of the TransitiveParentGroups field in ListTransitiveParentGroupsResponse_SdkV2.
func (m *ListTransitiveParentGroupsResponse_SdkV2) SetTransitiveParentGroups(ctx context.Context, v []TransitiveParentGroup_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["transitive_parent_groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TransitiveParentGroups = types.ListValueMust(t, vs)
}

type ListUsersProxyRequest_SdkV2 struct {
	// Filters users by `username` or `external_id`, for example `username eq
	// "someone@example.com"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of users to return, up to 1000. The default is 1000. A
	// page can contain fewer users.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list users in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent types.String `tfsdk:"-"`
}

func (to *ListUsersProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersProxyRequest_SdkV2) {
}

func (to *ListUsersProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersProxyRequest_SdkV2) {
}

func (m ListUsersProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()

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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListUsersRequest_SdkV2 struct {
	// Filters users by `username` or `external_id`, for example `username eq
	// "someone@example.com"`. Put the value in quotes.
	Filter types.String `tfsdk:"-"`
	// The maximum number of users to return, up to 1000. The default is 1000. A
	// page can contain fewer users.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the account to list users in, in the format
	// `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListUsersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersRequest_SdkV2) {
}

func (to *ListUsersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersRequest_SdkV2) {
}

func (m ListUsersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUsersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of users.
type ListUsersResponse_SdkV2 struct {
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The users on this page.
	Users types.List `tfsdk:"users"`
}

func (to *ListUsersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUsersResponse_SdkV2) {
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

func (to *ListUsersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUsersResponse_SdkV2) {
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

type ListWorkspaceAssignmentDetailsProxyRequest_SdkV2 struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value. If not provided, defaults to 1000,
	// which is also the maximum allowed. Requests for more than the maximum are
	// clamped to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) {
}

func (to *ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) {
}

func (m ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceAssignmentDetailsRequest_SdkV2 struct {
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

func (to *ListWorkspaceAssignmentDetailsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsRequest_SdkV2) {
}

func (to *ListWorkspaceAssignmentDetailsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsRequest_SdkV2) {
}

func (m ListWorkspaceAssignmentDetailsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAssignmentDetailsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Response message for listing workspace assignment details.
type ListWorkspaceAssignmentDetailsResponse_SdkV2 struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	WorkspaceAssignmentDetails types.List `tfsdk:"workspace_assignment_details"`
}

func (to *ListWorkspaceAssignmentDetailsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentDetailsResponse_SdkV2) {
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

func (to *ListWorkspaceAssignmentDetailsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsResponse_SdkV2) {
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

func (m ListWorkspaceAssignmentDetailsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAssignmentDetailsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_details": reflect.TypeOf(WorkspaceAssignmentDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentDetailsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentDetailsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":              m.NextPageToken,
			"workspace_assignment_details": m.WorkspaceAssignmentDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentDetailsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_assignment_details": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignmentDetails returns the value of the WorkspaceAssignmentDetails field in ListWorkspaceAssignmentDetailsResponse_SdkV2 as
// a slice of WorkspaceAssignmentDetail_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAssignmentDetailsResponse_SdkV2) GetWorkspaceAssignmentDetails(ctx context.Context) ([]WorkspaceAssignmentDetail_SdkV2, bool) {
	if m.WorkspaceAssignmentDetails.IsNull() || m.WorkspaceAssignmentDetails.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAssignmentDetail_SdkV2
	d := m.WorkspaceAssignmentDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignmentDetails sets the value of the WorkspaceAssignmentDetails field in ListWorkspaceAssignmentDetailsResponse_SdkV2.
func (m *ListWorkspaceAssignmentDetailsResponse_SdkV2) SetWorkspaceAssignmentDetails(ctx context.Context, v []WorkspaceAssignmentDetail_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceAssignmentDetails = types.ListValueMust(t, vs)
}

type ListWorkspaceAssignmentsProxyRequest_SdkV2 struct {
	// The maximum number of assignments to return, up to 1000. The default is
	// 1000. A page can contain fewer assignments.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Defaults to this
	// workspace. Any other workspace is rejected.
	Parent types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentsProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsProxyRequest_SdkV2) {
}

func (to *ListWorkspaceAssignmentsProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsProxyRequest_SdkV2) {
}

func (m ListWorkspaceAssignmentsProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["parent"] = attrs["parent"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentsProxyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceAssignmentsProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListWorkspaceAssignmentsRequest_SdkV2 struct {
	// The maximum number of assignments to return, up to 1000. The default is
	// 1000. A page can contain fewer assignments.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListWorkspaceAssignmentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsRequest_SdkV2) {
}

func (to *ListWorkspaceAssignmentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsRequest_SdkV2) {
}

func (m ListWorkspaceAssignmentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
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
func (m ListWorkspaceAssignmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A page of workspace assignments.
type ListWorkspaceAssignmentsResponse_SdkV2 struct {
	// The token for the next page. Not set on the last page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The assignments on this page.
	WorkspaceAssignments types.List `tfsdk:"workspace_assignments"`
}

func (to *ListWorkspaceAssignmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceAssignmentsResponse_SdkV2) {
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

func (to *ListWorkspaceAssignmentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentsResponse_SdkV2) {
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

func (m ListWorkspaceAssignmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceAssignmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignments": reflect.TypeOf(WorkspaceAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceAssignmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceAssignmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":       m.NextPageToken,
			"workspace_assignments": m.WorkspaceAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceAssignmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_assignments": basetypes.ListType{
				ElemType: WorkspaceAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignments returns the value of the WorkspaceAssignments field in ListWorkspaceAssignmentsResponse_SdkV2 as
// a slice of WorkspaceAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceAssignmentsResponse_SdkV2) GetWorkspaceAssignments(ctx context.Context) ([]WorkspaceAssignment_SdkV2, bool) {
	if m.WorkspaceAssignments.IsNull() || m.WorkspaceAssignments.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceAssignment_SdkV2
	d := m.WorkspaceAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceAssignments sets the value of the WorkspaceAssignments field in ListWorkspaceAssignmentsResponse_SdkV2.
func (m *ListWorkspaceAssignmentsResponse_SdkV2) SetWorkspaceAssignments(ctx context.Context, v []WorkspaceAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceAssignments = types.ListValueMust(t, vs)
}

// A rule that matches groups whose display name starts with a prefix. Matching
// is case-sensitive. For example, the prefix `data-` matches `data-platform`.
type PrefixMatchRule_SdkV2 struct {
	// The prefix to match. It must contain at least one non-whitespace
	// character. Leading and trailing spaces are kept.
	DisplayNamePrefix types.String `tfsdk:"display_name_prefix"`
}

func (to *PrefixMatchRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrefixMatchRule_SdkV2) {
}

func (to *PrefixMatchRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PrefixMatchRule_SdkV2) {
}

func (m PrefixMatchRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name_prefix"] = attrs["display_name_prefix"].SetRequired()
	attrs["display_name_prefix"] = attrs["display_name_prefix"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrefixMatchRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrefixMatchRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrefixMatchRule_SdkV2
// only implements ToObjectValue() and Type().
func (m PrefixMatchRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name_prefix": m.DisplayNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrefixMatchRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name_prefix": types.StringType,
		},
	}
}

// A rule that matches groups whose whole display name matches an RE2 regular
// expression. Matching is case-sensitive. To match part of a name, add `.*`,
// for example `.*-admins`.
type RegexMatchRule_SdkV2 struct {
	// The RE2 regular expression to match against the whole display name. At
	// most 255 characters.
	DisplayNameRegex types.String `tfsdk:"display_name_regex"`
}

func (to *RegexMatchRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegexMatchRule_SdkV2) {
}

func (to *RegexMatchRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegexMatchRule_SdkV2) {
}

func (m RegexMatchRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name_regex"] = attrs["display_name_regex"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegexMatchRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegexMatchRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegexMatchRule_SdkV2
// only implements ToObjectValue() and Type().
func (m RegexMatchRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name_regex": m.DisplayNameRegex,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegexMatchRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name_regex": types.StringType,
		},
	}
}

// Request to resolve a group by its ID in your identity provider (IdP).
type ResolveGroupProxyRequest_SdkV2 struct {
	// The ID of the group in your IdP.
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

// Request to resolve a group by its ID in your identity provider (IdP).
type ResolveGroupRequest_SdkV2 struct {
	// The ID of the group in your IdP.
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

// The response to resolving a group.
type ResolveGroupResponse_SdkV2 struct {
	// The resolved group.
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

// Request to resolve a service principal by its ID in your identity provider
// (IdP).
type ResolveServicePrincipalProxyRequest_SdkV2 struct {
	// The ID of the service principal in your IdP.
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

// Request to resolve a service principal by its ID in your identity provider
// (IdP).
type ResolveServicePrincipalRequest_SdkV2 struct {
	// The ID of the service principal in your IdP.
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

// The response to resolving a service principal.
type ResolveServicePrincipalResponse_SdkV2 struct {
	// The resolved service principal.
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

// Request to resolve a user by their ID in your identity provider (IdP).
type ResolveUserProxyRequest_SdkV2 struct {
	// The ID of the user in your IdP.
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

// Request to resolve a user by their ID in your identity provider (IdP).
type ResolveUserRequest_SdkV2 struct {
	// The ID of the user in your IdP.
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

// The response to resolving a user.
type ResolveUserResponse_SdkV2 struct {
	// The resolved user.
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

// A service principal in your Databricks account.
type ServicePrincipal_SdkV2 struct {
	// The ID of the account that contains the service principal.
	AccountId types.String `tfsdk:"account_id"`
	// Whether the service principal is active in the account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// The service principal's application ID, a UUID. If you don't set it when
	// you create the service principal, Databricks generates one. It must be
	// unique in the account and can't be changed.
	ApplicationId types.String `tfsdk:"application_id"`
	// The service principal's display name.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the service principal in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the account that contains the service principal, in
	// the format `accounts/{account_id}`. Workspace APIs return the same
	// format.
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the service principal.
	ServicePrincipalId types.String `tfsdk:"service_principal_id"`
}

func (to *ServicePrincipal_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServicePrincipal_SdkV2) {
}

func (to *ServicePrincipal_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServicePrincipal_SdkV2) {
}

func (m ServicePrincipal_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["account_sp_status"] = attrs["account_sp_status"].SetRequired()
	attrs["application_id"] = attrs["application_id"].SetComputed()
	attrs["application_id"] = attrs["application_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

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
			"account_id":           m.AccountId,
			"account_sp_status":    m.AccountSpStatus,
			"application_id":       m.ApplicationId,
			"display_name":         m.DisplayName,
			"external_id":          m.ExternalId,
			"name":                 m.Name,
			"parent":               m.Parent,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":           types.StringType,
			"account_sp_status":    types.StringType,
			"application_id":       types.StringType,
			"display_name":         types.StringType,
			"external_id":          types.StringType,
			"name":                 types.StringType,
			"parent":               types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

// A group that a principal belongs to, either directly or through another
// group.
type TransitiveParentGroup_SdkV2 struct {
	// The ID of the account that contains the group.
	AccountId types.String `tfsdk:"account_id"`
	// The ID of the group in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The Databricks ID of the group.
	GroupId types.String `tfsdk:"group_id"`
	// The resource name of the parent group, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/principals/{principal_id}/transitive-parent-groups/{group_id}`
	// - Workspace APIs:
	// `principals/{principal_id}/transitive-parent-groups/{group_id}`
	Name types.String `tfsdk:"name"`
}

func (to *TransitiveParentGroup_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitiveParentGroup_SdkV2) {
}

func (to *TransitiveParentGroup_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransitiveParentGroup_SdkV2) {
}

func (m TransitiveParentGroup_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitiveParentGroup.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitiveParentGroup_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitiveParentGroup_SdkV2
// only implements ToObjectValue() and Type().
func (m TransitiveParentGroup_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":  m.AccountId,
			"external_id": m.ExternalId,
			"group_id":    m.GroupId,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitiveParentGroup_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":  types.StringType,
			"external_id": types.StringType,
			"group_id":    types.StringType,
			"name":        types.StringType,
		},
	}
}

type UpdateGroupProxyRequest_SdkV2 struct {
	// The group to update. Its `name` identifies the group.
	Group types.List `tfsdk:"group"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update, separated by commas. You can update `group_name`
	// and `external_id`.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":        m.Name,
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
			"name":        types.StringType,
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
	// The group to update. Its `name` identifies the group.
	Group types.List `tfsdk:"group"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update, separated by commas. You can update `group_name`
	// and `external_id`.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":        m.Name,
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
			"name":        types.StringType,
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
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name types.String `tfsdk:"-"`
	// The service principal to update. Its `name` identifies the service
	// principal.
	ServicePrincipal types.List `tfsdk:"service_principal"`
	// The fields to update, separated by commas. You can update `display_name`,
	// `account_sp_status`, and `external_id`.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":              m.Name,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
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
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name types.String `tfsdk:"-"`
	// The service principal to update. Its `name` identifies the service
	// principal.
	ServicePrincipal types.List `tfsdk:"service_principal"`
	// The fields to update, separated by commas. You can update `display_name`,
	// `account_sp_status`, and `external_id`.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":              m.Name,
			"service_principal": m.ServicePrincipal,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
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
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update, separated by commas. You can update
	// `full_name.given_name`, `full_name.family_name`, `account_user_status`,
	// and `external_id`.
	UpdateMask types.String `tfsdk:"-"`
	// The user to update. Its `name` identifies the user.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":        m.Name,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
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
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update, separated by commas. You can update
	// `full_name.given_name`, `full_name.family_name`, `account_user_status`,
	// and `external_id`.
	UpdateMask types.String `tfsdk:"-"`
	// The user to update. Its `name` identifies the user.
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
	attrs["name"] = attrs["name"].SetRequired()
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
			"name":        m.Name,
			"update_mask": m.UpdateMask,
			"user":        m.User,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateUserRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
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

type UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment detail to be updated in Databricks.
	WorkspaceAssignmentDetail types.List `tfsdk:"workspace_assignment_detail"`
}

func (to *UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) {
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

func (to *UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal_id":                m.PrincipalId,
			"update_mask":                 m.UpdateMask,
			"workspace_assignment_detail": m.WorkspaceAssignmentDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"update_mask":  types.StringType,
			"workspace_assignment_detail": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2 as
// a WorkspaceAssignmentDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail_SdkV2, bool) {
	var e WorkspaceAssignmentDetail_SdkV2
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignmentDetail_SdkV2
	d := m.WorkspaceAssignmentDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2.
func (m *UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_detail"]
	m.WorkspaceAssignmentDetail = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAssignmentDetailRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment detail to be updated in Databricks.
	WorkspaceAssignmentDetail types.List `tfsdk:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being updated.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentDetailRequest_SdkV2) {
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

func (to *UpdateWorkspaceAssignmentDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentDetailRequest_SdkV2) {
	if !from.WorkspaceAssignmentDetail.IsNull() && !from.WorkspaceAssignmentDetail.IsUnknown() {
		if toWorkspaceAssignmentDetail, ok := to.GetWorkspaceAssignmentDetail(ctx); ok {
			if fromWorkspaceAssignmentDetail, ok := from.GetWorkspaceAssignmentDetail(ctx); ok {
				toWorkspaceAssignmentDetail.SyncFieldsDuringRead(ctx, fromWorkspaceAssignmentDetail)
				to.SetWorkspaceAssignmentDetail(ctx, toWorkspaceAssignmentDetail)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].SetRequired()
	attrs["workspace_assignment_detail"] = attrs["workspace_assignment_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateWorkspaceAssignmentDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment_detail": reflect.TypeOf(WorkspaceAssignmentDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateWorkspaceAssignmentDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"update_mask":  types.StringType,
			"workspace_assignment_detail": basetypes.ListType{
				ElemType: WorkspaceAssignmentDetail_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetWorkspaceAssignmentDetail returns the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailRequest_SdkV2 as
// a WorkspaceAssignmentDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentDetailRequest_SdkV2) GetWorkspaceAssignmentDetail(ctx context.Context) (WorkspaceAssignmentDetail_SdkV2, bool) {
	var e WorkspaceAssignmentDetail_SdkV2
	if m.WorkspaceAssignmentDetail.IsNull() || m.WorkspaceAssignmentDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignmentDetail_SdkV2
	d := m.WorkspaceAssignmentDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignmentDetail sets the value of the WorkspaceAssignmentDetail field in UpdateWorkspaceAssignmentDetailRequest_SdkV2.
func (m *UpdateWorkspaceAssignmentDetailRequest_SdkV2) SetWorkspaceAssignmentDetail(ctx context.Context, v WorkspaceAssignmentDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment_detail"]
	m.WorkspaceAssignmentDetail = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAssignmentProxyRequest_SdkV2 struct {
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update. The only supported field is `entitlements`.
	UpdateMask types.String `tfsdk:"-"`
	// The assignment to update. Its `name` identifies the assignment.
	WorkspaceAssignment types.List `tfsdk:"workspace_assignment"`
}

func (to *UpdateWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentProxyRequest_SdkV2) {
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

func (to *UpdateWorkspaceAssignmentProxyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentProxyRequest_SdkV2) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentProxyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["workspace_assignment"] = attrs["workspace_assignment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
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
func (m UpdateWorkspaceAssignmentProxyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentProxyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentProxyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                 m.Name,
			"update_mask":          m.UpdateMask,
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentProxyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"update_mask": types.StringType,
			"workspace_assignment": basetypes.ListType{
				ElemType: WorkspaceAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentProxyRequest_SdkV2 as
// a WorkspaceAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentProxyRequest_SdkV2) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment_SdkV2, bool) {
	var e WorkspaceAssignment_SdkV2
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignment_SdkV2
	d := m.WorkspaceAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentProxyRequest_SdkV2.
func (m *UpdateWorkspaceAssignmentProxyRequest_SdkV2) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment"]
	m.WorkspaceAssignment = types.ListValueMust(t, vs)
}

type UpdateWorkspaceAssignmentRequest_SdkV2 struct {
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name types.String `tfsdk:"-"`
	// The fields to update. The only supported field is `entitlements`.
	UpdateMask types.String `tfsdk:"-"`
	// The assignment to update. Its `name` identifies the assignment.
	WorkspaceAssignment types.List `tfsdk:"workspace_assignment"`
}

func (to *UpdateWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceAssignmentRequest_SdkV2) {
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

func (to *UpdateWorkspaceAssignmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceAssignmentRequest_SdkV2) {
	if !from.WorkspaceAssignment.IsNull() && !from.WorkspaceAssignment.IsUnknown() {
		if toWorkspaceAssignment, ok := to.GetWorkspaceAssignment(ctx); ok {
			if fromWorkspaceAssignment, ok := from.GetWorkspaceAssignment(ctx); ok {
				toWorkspaceAssignment.SyncFieldsDuringRead(ctx, fromWorkspaceAssignment)
				to.SetWorkspaceAssignment(ctx, toWorkspaceAssignment)
			}
		}
	}
}

func (m UpdateWorkspaceAssignmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_assignment"] = attrs["workspace_assignment"].SetRequired()
	attrs["workspace_assignment"] = attrs["workspace_assignment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
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
func (m UpdateWorkspaceAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignment": reflect.TypeOf(WorkspaceAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                 m.Name,
			"update_mask":          m.UpdateMask,
			"workspace_assignment": m.WorkspaceAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"update_mask": types.StringType,
			"workspace_assignment": basetypes.ListType{
				ElemType: WorkspaceAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceAssignment returns the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentRequest_SdkV2 as
// a WorkspaceAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceAssignmentRequest_SdkV2) GetWorkspaceAssignment(ctx context.Context) (WorkspaceAssignment_SdkV2, bool) {
	var e WorkspaceAssignment_SdkV2
	if m.WorkspaceAssignment.IsNull() || m.WorkspaceAssignment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceAssignment_SdkV2
	d := m.WorkspaceAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceAssignment sets the value of the WorkspaceAssignment field in UpdateWorkspaceAssignmentRequest_SdkV2.
func (m *UpdateWorkspaceAssignmentRequest_SdkV2) SetWorkspaceAssignment(ctx context.Context, v WorkspaceAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_assignment"]
	m.WorkspaceAssignment = types.ListValueMust(t, vs)
}

type UpdateWorkspaceIdentityDetailRequest_SdkV2 struct {
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name types.String `tfsdk:"-"`
	// The fields to update. The only supported field is
	// `workspace_identity_status`.
	UpdateMask types.String `tfsdk:"-"`
	// The workspace identity detail to update. Its `name` identifies the
	// principal.
	WorkspaceIdentityDetail types.List `tfsdk:"workspace_identity_detail"`
}

func (to *UpdateWorkspaceIdentityDetailRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceIdentityDetailRequest_SdkV2) {
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

func (to *UpdateWorkspaceIdentityDetailRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceIdentityDetailRequest_SdkV2) {
	if !from.WorkspaceIdentityDetail.IsNull() && !from.WorkspaceIdentityDetail.IsUnknown() {
		if toWorkspaceIdentityDetail, ok := to.GetWorkspaceIdentityDetail(ctx); ok {
			if fromWorkspaceIdentityDetail, ok := from.GetWorkspaceIdentityDetail(ctx); ok {
				toWorkspaceIdentityDetail.SyncFieldsDuringRead(ctx, fromWorkspaceIdentityDetail)
				to.SetWorkspaceIdentityDetail(ctx, toWorkspaceIdentityDetail)
			}
		}
	}
}

func (m UpdateWorkspaceIdentityDetailRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_identity_detail"] = attrs["workspace_identity_detail"].SetRequired()
	attrs["workspace_identity_detail"] = attrs["workspace_identity_detail"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
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
func (m UpdateWorkspaceIdentityDetailRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_identity_detail": reflect.TypeOf(WorkspaceIdentityDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceIdentityDetailRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceIdentityDetailRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                      m.Name,
			"update_mask":               m.UpdateMask,
			"workspace_identity_detail": m.WorkspaceIdentityDetail,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceIdentityDetailRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"update_mask": types.StringType,
			"workspace_identity_detail": basetypes.ListType{
				ElemType: WorkspaceIdentityDetail_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceIdentityDetail returns the value of the WorkspaceIdentityDetail field in UpdateWorkspaceIdentityDetailRequest_SdkV2 as
// a WorkspaceIdentityDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceIdentityDetailRequest_SdkV2) GetWorkspaceIdentityDetail(ctx context.Context) (WorkspaceIdentityDetail_SdkV2, bool) {
	var e WorkspaceIdentityDetail_SdkV2
	if m.WorkspaceIdentityDetail.IsNull() || m.WorkspaceIdentityDetail.IsUnknown() {
		return e, false
	}
	var v []WorkspaceIdentityDetail_SdkV2
	d := m.WorkspaceIdentityDetail.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceIdentityDetail sets the value of the WorkspaceIdentityDetail field in UpdateWorkspaceIdentityDetailRequest_SdkV2.
func (m *UpdateWorkspaceIdentityDetailRequest_SdkV2) SetWorkspaceIdentityDetail(ctx context.Context, v WorkspaceIdentityDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_identity_detail"]
	m.WorkspaceIdentityDetail = types.ListValueMust(t, vs)
}

// A user in your Databricks account.
type User_SdkV2 struct {
	// The ID of the account that contains the user.
	AccountId types.String `tfsdk:"account_id"`
	// Whether the user is active in the account. New users are active unless
	// you set this to `INACTIVE`.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// The ID of the user in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The user's name.
	FullName types.List `tfsdk:"full_name"`
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the account that contains the user, in the format
	// `accounts/{account_id}`. Workspace APIs return the same format.
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the user.
	UserId types.String `tfsdk:"user_id"`
	// The user's username, which is their email address. It can't be changed
	// after the user is created.
	Username types.String `tfsdk:"username"`
}

func (to *User_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from User_SdkV2) {
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

func (to *User_SdkV2) SyncFieldsDuringRead(ctx context.Context, from User_SdkV2) {
	if !from.FullName.IsNull() && !from.FullName.IsUnknown() {
		if toFullName, ok := to.GetFullName(ctx); ok {
			if fromFullName, ok := from.GetFullName(ctx); ok {
				toFullName.SyncFieldsDuringRead(ctx, fromFullName)
				to.SetFullName(ctx, toFullName)
			}
		}
	}
}

func (m User_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["account_user_status"] = attrs["account_user_status"].SetRequired()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["user_id"] = attrs["user_id"].SetComputed()
	attrs["user_id"] = attrs["user_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
func (m User_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name": reflect.TypeOf(UserFullName_SdkV2{}),
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
			"full_name":           m.FullName,
			"name":                m.Name,
			"parent":              m.Parent,
			"user_id":             m.UserId,
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
			"full_name": basetypes.ListType{
				ElemType: UserFullName_SdkV2{}.Type(ctx),
			},
			"name":     types.StringType,
			"parent":   types.StringType,
			"user_id":  types.StringType,
			"username": types.StringType,
		},
	}
}

// GetFullName returns the value of the FullName field in User_SdkV2 as
// a UserFullName_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *User_SdkV2) GetFullName(ctx context.Context) (UserFullName_SdkV2, bool) {
	var e UserFullName_SdkV2
	if m.FullName.IsNull() || m.FullName.IsUnknown() {
		return e, false
	}
	var v []UserFullName_SdkV2
	d := m.FullName.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFullName sets the value of the FullName field in User_SdkV2.
func (m *User_SdkV2) SetFullName(ctx context.Context, v UserFullName_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_name"]
	m.FullName = types.ListValueMust(t, vs)
}

// A user's name.
type UserFullName_SdkV2 struct {
	// The user's family (last) name.
	FamilyName types.String `tfsdk:"family_name"`
	// The user's given (first) name.
	GivenName types.String `tfsdk:"given_name"`
}

func (to *UserFullName_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UserFullName_SdkV2) {
}

func (to *UserFullName_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UserFullName_SdkV2) {
}

func (m UserFullName_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UserFullName_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserFullName_SdkV2
// only implements ToObjectValue() and Type().
func (m UserFullName_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"family_name": m.FamilyName,
			"given_name":  m.GivenName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UserFullName_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The resource name of the workspace access detail. The format depends on
	// the API that returned it: - Account-scoped:
	// accounts/{account_id}/workspaces/{workspace_id}/workspace-access-details/{principal_id}
	// - Workspace-scoped: workspace-access-details/{principal_id}
	Name types.String `tfsdk:"name"`
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
	attrs["name"] = attrs["name"].SetOptional()
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
			"name":           m.Name,
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
			"name":        types.StringType,
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

// A principal's direct assignment to a workspace, and the entitlements it
// grants there.
//
// Only principals assigned directly to the workspace have an assignment. A
// principal that gets access only through a group doesn't. Creating an
// assignment gives the principal access to the workspace, and deleting it
// removes that direct access. These APIs replace managing workspace assignments
// through SCIM and the permission assignment APIs. They require a workspace
// with identity federation enabled. You can't assign, change, or unassign
// principals that Databricks manages.
//
// `entitlements` is the only field you set. `effective_entitlements` also
// includes the entitlements the principal gets through its groups. An
// assignment must grant at least one entitlement. To remove a principal's
// direct access, delete the assignment.
type WorkspaceAssignment_SdkV2 struct {
	// The ID of the account that contains the workspace.
	AccountId types.String `tfsdk:"account_id"`
	// Every entitlement the principal has in the workspace, granted directly or
	// through a group. List responses leave it empty.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// The entitlements granted directly by this assignment. Create and update
	// set exactly this list, even if the principal also gets some of these
	// entitlements through a group. List responses leave it empty.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return
	// the same format.
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the assigned user, service principal, or group. It
	// can't be changed after the assignment is created.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// Whether the assigned principal is a user, service principal, or group.
	PrincipalType types.String `tfsdk:"principal_type"`
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceAssignment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignment_SdkV2) {
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

func (to *WorkspaceAssignment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignment_SdkV2) {
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

func (m WorkspaceAssignment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_entitlements"] = attrs["effective_entitlements"].SetComputed()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":             m.AccountId,
			"effective_entitlements": m.EffectiveEntitlements,
			"entitlements":           m.Entitlements,
			"name":                   m.Name,
			"parent":                 m.Parent,
			"principal_id":           m.PrincipalId,
			"principal_type":         m.PrincipalType,
			"workspace_id":           m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"effective_entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"name":           types.StringType,
			"parent":         types.StringType,
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

// GetEffectiveEntitlements returns the value of the EffectiveEntitlements field in WorkspaceAssignment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment_SdkV2) GetEffectiveEntitlements(ctx context.Context) ([]types.String, bool) {
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

// SetEffectiveEntitlements sets the value of the EffectiveEntitlements field in WorkspaceAssignment_SdkV2.
func (m *WorkspaceAssignment_SdkV2) SetEffectiveEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveEntitlements = types.SetValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in WorkspaceAssignment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment_SdkV2) GetEntitlements(ctx context.Context) ([]types.String, bool) {
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

// SetEntitlements sets the value of the Entitlements field in WorkspaceAssignment_SdkV2.
func (m *WorkspaceAssignment_SdkV2) SetEntitlements(ctx context.Context, v []types.String) {
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
type WorkspaceAssignmentDetail_SdkV2 struct {
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

func (to *WorkspaceAssignmentDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignmentDetail_SdkV2) {
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

func (to *WorkspaceAssignmentDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignmentDetail_SdkV2) {
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

func (m WorkspaceAssignmentDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceAssignmentDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignmentDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignmentDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m WorkspaceAssignmentDetail_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetEffectiveEntitlements returns the value of the EffectiveEntitlements field in WorkspaceAssignmentDetail_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignmentDetail_SdkV2) GetEffectiveEntitlements(ctx context.Context) ([]types.String, bool) {
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

// SetEffectiveEntitlements sets the value of the EffectiveEntitlements field in WorkspaceAssignmentDetail_SdkV2.
func (m *WorkspaceAssignmentDetail_SdkV2) SetEffectiveEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveEntitlements = types.SetValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in WorkspaceAssignmentDetail_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignmentDetail_SdkV2) GetEntitlements(ctx context.Context) ([]types.String, bool) {
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

// SetEntitlements sets the value of the Entitlements field in WorkspaceAssignmentDetail_SdkV2.
func (m *WorkspaceAssignmentDetail_SdkV2) SetEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entitlements = types.SetValueMust(t, vs)
}

// A user, service principal, or group's status in a workspace, and whether it
// has access to the workspace directly or through a group. Update it to
// activate or deactivate a user or service principal in the workspace.
type WorkspaceIdentityDetail_SdkV2 struct {
	// Whether the principal is assigned to the workspace directly or has access
	// through a group.
	AssignmentType types.String `tfsdk:"assignment_type"`
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name types.String `tfsdk:"name"`
	// The Databricks ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// Whether the principal is a user, service principal, or group.
	PrincipalType types.String `tfsdk:"principal_type"`
	// Whether the principal is active in the workspace. Groups are always
	// `ACTIVE`.
	WorkspaceIdentityStatus types.String `tfsdk:"workspace_identity_status"`
}

func (to *WorkspaceIdentityDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceIdentityDetail_SdkV2) {
}

func (to *WorkspaceIdentityDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceIdentityDetail_SdkV2) {
}

func (m WorkspaceIdentityDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assignment_type"] = attrs["assignment_type"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
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
func (m WorkspaceIdentityDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceIdentityDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceIdentityDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assignment_type":           m.AssignmentType,
			"name":                      m.Name,
			"principal_id":              m.PrincipalId,
			"principal_type":            m.PrincipalType,
			"workspace_identity_status": m.WorkspaceIdentityStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceIdentityDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assignment_type":           types.StringType,
			"name":                      types.StringType,
			"principal_id":              types.Int64Type,
			"principal_type":            types.StringType,
			"workspace_identity_status": types.StringType,
		},
	}
}
