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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
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
	// Required. Workspace assignment detail to be created in <Databricks>.
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

// The details of a Group resource.
type Group_SdkV2 struct {
	// The parent account ID for group in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// ExternalId of the group in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`
	// Display name of the group.
	GroupName types.String `tfsdk:"group_name"`
	// Internal group ID of the group in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
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
			"internal_id": types.StringType,
		},
	}
}

type ListWorkspaceAssignmentDetailsProxyRequest_SdkV2 struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// ListWorkspaceAssignmentDetailsProxy call. Provide this to retrieve the
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
	// may return fewer than this value.
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
}

func (to *ListWorkspaceAssignmentDetailsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceAssignmentDetailsResponse_SdkV2) {
	if !from.WorkspaceAssignmentDetails.IsNull() && !from.WorkspaceAssignmentDetails.IsUnknown() && to.WorkspaceAssignmentDetails.IsNull() && len(from.WorkspaceAssignmentDetails.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceAssignmentDetails, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceAssignmentDetails = from.WorkspaceAssignmentDetails
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
	InternalId types.String `tfsdk:"internal_id"`
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
			"internal_id":       types.StringType,
		},
	}
}

type UpdateWorkspaceAssignmentDetailProxyRequest_SdkV2 struct {
	// Required. ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"-"`
	// Required. The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
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
	// Required. Workspace assignment detail to be updated in <Databricks>.
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

// The details of a User resource.
type User_SdkV2 struct {
	// The accountId parent of the user in Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// ExternalId of the user in the customer's IdP.
	ExternalId types.String `tfsdk:"external_id"`

	FullName types.List `tfsdk:"full_name"`
	// Internal userId of the user in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// Username/email of the user.
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
	attrs["account_user_status"] = attrs["account_user_status"].SetOptional()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["full_name"] = attrs["full_name"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
			"internal_id":         m.InternalId,
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
			"internal_id": types.StringType,
			"username":    types.StringType,
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

type UserFullName_SdkV2 struct {
	FamilyName types.String `tfsdk:"family_name"`

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

// The details of a principal's assignment to a workspace.
type WorkspaceAssignmentDetail_SdkV2 struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId types.String `tfsdk:"account_id"`
	// Entitlements granted directly to the principal on this workspace. The
	// only client-settable field: create and update manage exactly this set
	// (including entitlements the principal also holds via a group). Not
	// populated by ListWorkspaceAssignmentDetails (omitted for scalability);
	// call GetWorkspaceAssignmentDetail to read the entitlements for a single
	// principal.
	Entitlements types.List `tfsdk:"entitlements"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`

	PrincipalType types.String `tfsdk:"principal_type"`
	// The workspace ID where the principal is assigned
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceAssignmentDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignmentDetail_SdkV2) {
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (to *WorkspaceAssignmentDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignmentDetail_SdkV2) {
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (m WorkspaceAssignmentDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
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
		"entitlements": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignmentDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignmentDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":     m.AccountId,
			"entitlements":   m.Entitlements,
			"principal_id":   m.PrincipalId,
			"principal_type": m.PrincipalType,
			"workspace_id":   m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceAssignmentDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
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
	m.Entitlements = types.ListValueMust(t, vs)
}
