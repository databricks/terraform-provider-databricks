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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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

// The details of a Group resource.
type Group struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Display name of the group.
	GroupName types.String `tfsdk:"group_name"`
	// Internal group ID of the group in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
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
			"internal_id": types.StringType,
		},
	}
}

type ListWorkspaceAssignmentDetailsProxyRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// ListWorkspaceAssignmentDetailsProxy call. Provide this to retrieve the
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
	// may return fewer than this value.
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
	InternalId types.String `tfsdk:"internal_id"`
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
			"internal_id":       types.StringType,
		},
	}
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
	InternalId types.String `tfsdk:"internal_id"`
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
	attrs["account_user_status"] = attrs["account_user_status"].SetOptional()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["username"] = attrs["username"].SetOptional()
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
			"internal_id":         m.InternalId,
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
			"internal_id":         types.StringType,
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

// The details of a principal's assignment to a workspace.
type WorkspaceAssignmentDetail struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId types.String `tfsdk:"account_id"`
	// The principal's full effective entitlements granted in this workspace:
	// every entitlement it holds whether granted directly or via group
	// membership. Populated on Get; empty on List.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// Entitlements granted directly to the principal on this workspace. The
	// only client-settable field: create and update manage exactly this set
	// (including entitlements the principal also holds via a group). Not
	// populated by ListWorkspaceAssignmentDetails (omitted for scalability);
	// call GetWorkspaceAssignmentDetail to read the entitlements for a single
	// principal.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`

	PrincipalType types.String `tfsdk:"principal_type"`
	// The workspace ID where the principal is assigned
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
