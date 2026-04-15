// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package environments_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// A unique identifier for this request. A random UUID is recommended. This
	// request is only idempotent if a request_id is provided.
	RequestId types.String `tfsdk:"-"`
	// Required. The workspace base environment to create.
	WorkspaceBaseEnvironment types.List `tfsdk:"workspace_base_environment"`
	// The ID to use for the workspace base environment, which will become the
	// final component of the resource name. This value should be 4-63
	// characters, and valid characters are /[a-z][0-9]-/.
	WorkspaceBaseEnvironmentId types.String `tfsdk:"-"`
}

func (to *CreateWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.WorkspaceBaseEnvironment.IsNull() && !from.WorkspaceBaseEnvironment.IsUnknown() {
		if toWorkspaceBaseEnvironment, ok := to.GetWorkspaceBaseEnvironment(ctx); ok {
			if fromWorkspaceBaseEnvironment, ok := from.GetWorkspaceBaseEnvironment(ctx); ok {
				// Recursively sync the fields of WorkspaceBaseEnvironment
				toWorkspaceBaseEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceBaseEnvironment)
				to.SetWorkspaceBaseEnvironment(ctx, toWorkspaceBaseEnvironment)
			}
		}
	}
}

func (to *CreateWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.WorkspaceBaseEnvironment.IsNull() && !from.WorkspaceBaseEnvironment.IsUnknown() {
		if toWorkspaceBaseEnvironment, ok := to.GetWorkspaceBaseEnvironment(ctx); ok {
			if fromWorkspaceBaseEnvironment, ok := from.GetWorkspaceBaseEnvironment(ctx); ok {
				toWorkspaceBaseEnvironment.SyncFieldsDuringRead(ctx, fromWorkspaceBaseEnvironment)
				to.SetWorkspaceBaseEnvironment(ctx, toWorkspaceBaseEnvironment)
			}
		}
	}
}

func (m CreateWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_base_environment"] = attrs["workspace_base_environment"].SetRequired()
	attrs["workspace_base_environment"] = attrs["workspace_base_environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workspace_base_environment_id"] = attrs["workspace_base_environment_id"].SetOptional()
	attrs["request_id"] = attrs["request_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_base_environment": reflect.TypeOf(WorkspaceBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_id":                    m.RequestId,
			"workspace_base_environment":    m.WorkspaceBaseEnvironment,
			"workspace_base_environment_id": m.WorkspaceBaseEnvironmentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_id": types.StringType,
			"workspace_base_environment": basetypes.ListType{
				ElemType: WorkspaceBaseEnvironment_SdkV2{}.Type(ctx),
			},
			"workspace_base_environment_id": types.StringType,
		},
	}
}

// GetWorkspaceBaseEnvironment returns the value of the WorkspaceBaseEnvironment field in CreateWorkspaceBaseEnvironmentRequest_SdkV2 as
// a WorkspaceBaseEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceBaseEnvironmentRequest_SdkV2) GetWorkspaceBaseEnvironment(ctx context.Context) (WorkspaceBaseEnvironment_SdkV2, bool) {
	var e WorkspaceBaseEnvironment_SdkV2
	if m.WorkspaceBaseEnvironment.IsNull() || m.WorkspaceBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceBaseEnvironment_SdkV2
	d := m.WorkspaceBaseEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceBaseEnvironment sets the value of the WorkspaceBaseEnvironment field in CreateWorkspaceBaseEnvironmentRequest_SdkV2.
func (m *CreateWorkspaceBaseEnvironmentRequest_SdkV2) SetWorkspaceBaseEnvironment(ctx context.Context, v WorkspaceBaseEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_base_environment"]
	m.WorkspaceBaseEnvironment = types.ListValueMust(t, vs)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto_SdkV2 struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack_trace"] = attrs["stack_trace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksServiceExceptionWithDetailsProto.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":     m.Details,
			"error_code":  m.ErrorCode,
			"message":     m.Message,
			"stack_trace": m.StackTrace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"error_code":  types.StringType,
			"message":     types.StringType,
			"stack_trace": types.StringType,
		},
	}
}

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto_SdkV2) GetDetails(ctx context.Context) ([]types.Object, bool) {
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto_SdkV2.
func (m *DatabricksServiceExceptionWithDetailsProto_SdkV2) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

// A singleton resource representing the default workspace base environment
// configuration. This resource contains the workspace base environments that
// are used as defaults for serverless notebooks and jobs in the workspace, for
// both CPU and GPU compute types.
type DefaultWorkspaceBaseEnvironment_SdkV2 struct {
	// The default workspace base environment for CPU compute. Format:
	// workspace-base-environments/{workspace_base_environment}
	CpuWorkspaceBaseEnvironment types.String `tfsdk:"cpu_workspace_base_environment"`
	// The default workspace base environment for GPU compute. Format:
	// workspace-base-environments/{workspace_base_environment}
	GpuWorkspaceBaseEnvironment types.String `tfsdk:"gpu_workspace_base_environment"`
	// The resource name of this singleton resource. Format:
	// default-workspace-base-environment
	Name types.String `tfsdk:"name"`
}

func (to *DefaultWorkspaceBaseEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultWorkspaceBaseEnvironment_SdkV2) {
}

func (to *DefaultWorkspaceBaseEnvironment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DefaultWorkspaceBaseEnvironment_SdkV2) {
}

func (m DefaultWorkspaceBaseEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cpu_workspace_base_environment"] = attrs["cpu_workspace_base_environment"].SetOptional()
	attrs["gpu_workspace_base_environment"] = attrs["gpu_workspace_base_environment"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultWorkspaceBaseEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultWorkspaceBaseEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultWorkspaceBaseEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (m DefaultWorkspaceBaseEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cpu_workspace_base_environment": m.CpuWorkspaceBaseEnvironment,
			"gpu_workspace_base_environment": m.GpuWorkspaceBaseEnvironment,
			"name":                           m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultWorkspaceBaseEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cpu_workspace_base_environment": types.StringType,
			"gpu_workspace_base_environment": types.StringType,
			"name":                           types.StringType,
		},
	}
}

type DeleteWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// Required. The resource name of the workspace base environment to delete.
	// Format: workspace-base-environments/{workspace_base_environment}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (to *DeleteWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (m DeleteWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// A static resource name of the default workspace base environment. Format:
	// default-workspace-base-environment
	Name types.String `tfsdk:"-"`
}

func (to *GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (to *GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (m GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest_SdkV2 struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest_SdkV2) {
}

func (to *GetOperationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest_SdkV2) {
}

func (m GetOperationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOperationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetOperationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// Required. The resource name of the workspace base environment to
	// retrieve. Format:
	// workspace-base-environments/{workspace_base_environment}
	Name types.String `tfsdk:"-"`
}

func (to *GetWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (to *GetWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (m GetWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListWorkspaceBaseEnvironmentsRequest_SdkV2 struct {
	// The maximum number of environments to return per page. Default is 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token for pagination. Received from a previous
	// ListWorkspaceBaseEnvironments call.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWorkspaceBaseEnvironmentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceBaseEnvironmentsRequest_SdkV2) {
}

func (to *ListWorkspaceBaseEnvironmentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceBaseEnvironmentsRequest_SdkV2) {
}

func (m ListWorkspaceBaseEnvironmentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceBaseEnvironmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceBaseEnvironmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceBaseEnvironmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceBaseEnvironmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceBaseEnvironmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response message for ListWorkspaceBaseEnvironments.
type ListWorkspaceBaseEnvironmentsResponse_SdkV2 struct {
	// Token to retrieve the next page of results. Empty if there are no more
	// results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of workspace base environments.
	WorkspaceBaseEnvironments types.List `tfsdk:"workspace_base_environments"`
}

func (to *ListWorkspaceBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceBaseEnvironmentsResponse_SdkV2) {
	if !from.WorkspaceBaseEnvironments.IsNull() && !from.WorkspaceBaseEnvironments.IsUnknown() && to.WorkspaceBaseEnvironments.IsNull() && len(from.WorkspaceBaseEnvironments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceBaseEnvironments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceBaseEnvironments = from.WorkspaceBaseEnvironments
	}
}

func (to *ListWorkspaceBaseEnvironmentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceBaseEnvironmentsResponse_SdkV2) {
	if !from.WorkspaceBaseEnvironments.IsNull() && !from.WorkspaceBaseEnvironments.IsUnknown() && to.WorkspaceBaseEnvironments.IsNull() && len(from.WorkspaceBaseEnvironments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WorkspaceBaseEnvironments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WorkspaceBaseEnvironments = from.WorkspaceBaseEnvironments
	}
}

func (m ListWorkspaceBaseEnvironmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["workspace_base_environments"] = attrs["workspace_base_environments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceBaseEnvironmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceBaseEnvironmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_base_environments": reflect.TypeOf(WorkspaceBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceBaseEnvironmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceBaseEnvironmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":             m.NextPageToken,
			"workspace_base_environments": m.WorkspaceBaseEnvironments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceBaseEnvironmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"workspace_base_environments": basetypes.ListType{
				ElemType: WorkspaceBaseEnvironment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceBaseEnvironments returns the value of the WorkspaceBaseEnvironments field in ListWorkspaceBaseEnvironmentsResponse_SdkV2 as
// a slice of WorkspaceBaseEnvironment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWorkspaceBaseEnvironmentsResponse_SdkV2) GetWorkspaceBaseEnvironments(ctx context.Context) ([]WorkspaceBaseEnvironment_SdkV2, bool) {
	if m.WorkspaceBaseEnvironments.IsNull() || m.WorkspaceBaseEnvironments.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBaseEnvironment_SdkV2
	d := m.WorkspaceBaseEnvironments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceBaseEnvironments sets the value of the WorkspaceBaseEnvironments field in ListWorkspaceBaseEnvironmentsResponse_SdkV2.
func (m *ListWorkspaceBaseEnvironmentsResponse_SdkV2) SetWorkspaceBaseEnvironments(ctx context.Context, v []WorkspaceBaseEnvironment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_base_environments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceBaseEnvironments = types.ListValueMust(t, vs)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation_SdkV2 struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.List `tfsdk:"error"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata types.Object `tfsdk:"metadata"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name types.String `tfsdk:"name"`
	// The normal, successful response of the operation.
	Response types.Object `tfsdk:"response"`
}

func (to *Operation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *Operation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Operation_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Operation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Operation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation_SdkV2
// only implements ToObjectValue() and Type().
func (m Operation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"done":     m.Done,
			"error":    m.Error,
			"metadata": m.Metadata,
			"name":     m.Name,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done": types.BoolType,
			"error": basetypes.ListType{
				ElemType: DatabricksServiceExceptionWithDetailsProto_SdkV2{}.Type(ctx),
			},
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation_SdkV2 as
// a DatabricksServiceExceptionWithDetailsProto_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation_SdkV2) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto_SdkV2, bool) {
	var e DatabricksServiceExceptionWithDetailsProto_SdkV2
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v []DatabricksServiceExceptionWithDetailsProto_SdkV2
	d := m.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in Operation_SdkV2.
func (m *Operation_SdkV2) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	m.Error = types.ListValueMust(t, vs)
}

// Request message for RefreshWorkspaceBaseEnvironments.
type RefreshWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// Required. The resource name of the workspace base environment to delete.
	// Format: workspace-base-environments/{workspace_base_environment}
	Name types.String `tfsdk:"-"`
}

func (to *RefreshWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RefreshWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (to *RefreshWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RefreshWorkspaceBaseEnvironmentRequest_SdkV2) {
}

func (m RefreshWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RefreshWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RefreshWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RefreshWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RefreshWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	// Required. The default workspace base environment configuration to update.
	DefaultWorkspaceBaseEnvironment types.List `tfsdk:"default_workspace_base_environment"`
	// The resource name of this singleton resource. Format:
	// default-workspace-base-environment
	Name types.String `tfsdk:"-"`
	// Field mask specifying which fields to update. Use comma as the separator
	// for multiple fields (no space). The special value '*' indicates that all
	// fields should be updated (full replacement). Valid field paths:
	// cpu_workspace_base_environment, gpu_workspace_base_environment
	//
	// To unset one or both defaults, include the field path(s) in the mask and
	// omit them from the request body. To unset both, you must list both paths
	// explicitly — the wildcard '*' cannot be used to unset fields.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.DefaultWorkspaceBaseEnvironment.IsNull() && !from.DefaultWorkspaceBaseEnvironment.IsUnknown() {
		if toDefaultWorkspaceBaseEnvironment, ok := to.GetDefaultWorkspaceBaseEnvironment(ctx); ok {
			if fromDefaultWorkspaceBaseEnvironment, ok := from.GetDefaultWorkspaceBaseEnvironment(ctx); ok {
				// Recursively sync the fields of DefaultWorkspaceBaseEnvironment
				toDefaultWorkspaceBaseEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultWorkspaceBaseEnvironment)
				to.SetDefaultWorkspaceBaseEnvironment(ctx, toDefaultWorkspaceBaseEnvironment)
			}
		}
	}
}

func (to *UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.DefaultWorkspaceBaseEnvironment.IsNull() && !from.DefaultWorkspaceBaseEnvironment.IsUnknown() {
		if toDefaultWorkspaceBaseEnvironment, ok := to.GetDefaultWorkspaceBaseEnvironment(ctx); ok {
			if fromDefaultWorkspaceBaseEnvironment, ok := from.GetDefaultWorkspaceBaseEnvironment(ctx); ok {
				toDefaultWorkspaceBaseEnvironment.SyncFieldsDuringRead(ctx, fromDefaultWorkspaceBaseEnvironment)
				to.SetDefaultWorkspaceBaseEnvironment(ctx, toDefaultWorkspaceBaseEnvironment)
			}
		}
	}
}

func (m UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_workspace_base_environment"] = attrs["default_workspace_base_environment"].SetRequired()
	attrs["default_workspace_base_environment"] = attrs["default_workspace_base_environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_workspace_base_environment": reflect.TypeOf(DefaultWorkspaceBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_workspace_base_environment": m.DefaultWorkspaceBaseEnvironment,
			"name":                               m.Name,
			"update_mask":                        m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_workspace_base_environment": basetypes.ListType{
				ElemType: DefaultWorkspaceBaseEnvironment_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDefaultWorkspaceBaseEnvironment returns the value of the DefaultWorkspaceBaseEnvironment field in UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2 as
// a DefaultWorkspaceBaseEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) GetDefaultWorkspaceBaseEnvironment(ctx context.Context) (DefaultWorkspaceBaseEnvironment_SdkV2, bool) {
	var e DefaultWorkspaceBaseEnvironment_SdkV2
	if m.DefaultWorkspaceBaseEnvironment.IsNull() || m.DefaultWorkspaceBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v []DefaultWorkspaceBaseEnvironment_SdkV2
	d := m.DefaultWorkspaceBaseEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultWorkspaceBaseEnvironment sets the value of the DefaultWorkspaceBaseEnvironment field in UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2.
func (m *UpdateDefaultWorkspaceBaseEnvironmentRequest_SdkV2) SetDefaultWorkspaceBaseEnvironment(ctx context.Context, v DefaultWorkspaceBaseEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_workspace_base_environment"]
	m.DefaultWorkspaceBaseEnvironment = types.ListValueMust(t, vs)
}

type UpdateWorkspaceBaseEnvironmentRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
	// Required. The workspace base environment with updated fields. The name
	// field is used to identify the environment to update.
	WorkspaceBaseEnvironment types.List `tfsdk:"workspace_base_environment"`
}

func (to *UpdateWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.WorkspaceBaseEnvironment.IsNull() && !from.WorkspaceBaseEnvironment.IsUnknown() {
		if toWorkspaceBaseEnvironment, ok := to.GetWorkspaceBaseEnvironment(ctx); ok {
			if fromWorkspaceBaseEnvironment, ok := from.GetWorkspaceBaseEnvironment(ctx); ok {
				// Recursively sync the fields of WorkspaceBaseEnvironment
				toWorkspaceBaseEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceBaseEnvironment)
				to.SetWorkspaceBaseEnvironment(ctx, toWorkspaceBaseEnvironment)
			}
		}
	}
}

func (to *UpdateWorkspaceBaseEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceBaseEnvironmentRequest_SdkV2) {
	if !from.WorkspaceBaseEnvironment.IsNull() && !from.WorkspaceBaseEnvironment.IsUnknown() {
		if toWorkspaceBaseEnvironment, ok := to.GetWorkspaceBaseEnvironment(ctx); ok {
			if fromWorkspaceBaseEnvironment, ok := from.GetWorkspaceBaseEnvironment(ctx); ok {
				toWorkspaceBaseEnvironment.SyncFieldsDuringRead(ctx, fromWorkspaceBaseEnvironment)
				to.SetWorkspaceBaseEnvironment(ctx, toWorkspaceBaseEnvironment)
			}
		}
	}
}

func (m UpdateWorkspaceBaseEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_base_environment"] = attrs["workspace_base_environment"].SetRequired()
	attrs["workspace_base_environment"] = attrs["workspace_base_environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBaseEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceBaseEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_base_environment": reflect.TypeOf(WorkspaceBaseEnvironment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBaseEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceBaseEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                       m.Name,
			"workspace_base_environment": m.WorkspaceBaseEnvironment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceBaseEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"workspace_base_environment": basetypes.ListType{
				ElemType: WorkspaceBaseEnvironment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceBaseEnvironment returns the value of the WorkspaceBaseEnvironment field in UpdateWorkspaceBaseEnvironmentRequest_SdkV2 as
// a WorkspaceBaseEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceBaseEnvironmentRequest_SdkV2) GetWorkspaceBaseEnvironment(ctx context.Context) (WorkspaceBaseEnvironment_SdkV2, bool) {
	var e WorkspaceBaseEnvironment_SdkV2
	if m.WorkspaceBaseEnvironment.IsNull() || m.WorkspaceBaseEnvironment.IsUnknown() {
		return e, false
	}
	var v []WorkspaceBaseEnvironment_SdkV2
	d := m.WorkspaceBaseEnvironment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceBaseEnvironment sets the value of the WorkspaceBaseEnvironment field in UpdateWorkspaceBaseEnvironmentRequest_SdkV2.
func (m *UpdateWorkspaceBaseEnvironmentRequest_SdkV2) SetWorkspaceBaseEnvironment(ctx context.Context, v WorkspaceBaseEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_base_environment"]
	m.WorkspaceBaseEnvironment = types.ListValueMust(t, vs)
}

// A WorkspaceBaseEnvironment defines a workspace-level environment
// configuration consisting of an environment version and a list of
// dependencies.
type WorkspaceBaseEnvironment_SdkV2 struct {
	// The type of base environment (CPU or GPU).
	BaseEnvironmentType          types.String `tfsdk:"base_environment_type"`
	EffectiveBaseEnvironmentType types.String `tfsdk:"effective_base_environment_type"`
	// Timestamp when the environment was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// User ID of the creator.
	CreatorUserId types.String `tfsdk:"creator_user_id"`
	// Human-readable display name for the workspace base environment.
	DisplayName types.String `tfsdk:"display_name"`
	// The WSFS or UC Volumes path to the environment YAML file.
	Filepath types.String `tfsdk:"filepath"`
	// Whether this is the default environment for the workspace.
	IsDefault types.Bool `tfsdk:"is_default"`
	// User ID of the last user who updated the environment.
	LastUpdatedUserId types.String `tfsdk:"last_updated_user_id"`
	// Status message providing additional details about the environment status.
	Message types.String `tfsdk:"message"`
	// The resource name of the workspace base environment. Format:
	// workspace-base-environments/{workspace-base-environment}
	Name types.String `tfsdk:"name"`
	// The status of the materialized workspace base environment.
	Status types.String `tfsdk:"status"`
	// Timestamp when the environment was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *WorkspaceBaseEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceBaseEnvironment_SdkV2) {
	to.EffectiveBaseEnvironmentType = to.BaseEnvironmentType
	to.BaseEnvironmentType = from.BaseEnvironmentType
}

func (to *WorkspaceBaseEnvironment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceBaseEnvironment_SdkV2) {
	to.EffectiveBaseEnvironmentType = from.EffectiveBaseEnvironmentType
	if from.EffectiveBaseEnvironmentType.ValueString() == to.BaseEnvironmentType.ValueString() {
		to.BaseEnvironmentType = from.BaseEnvironmentType
	}
}

func (m WorkspaceBaseEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_base_environment_type"] = attrs["effective_base_environment_type"].SetComputed()
	attrs["base_environment_type"] = attrs["base_environment_type"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["filepath"] = attrs["filepath"].SetOptional()
	attrs["is_default"] = attrs["is_default"].SetComputed()
	attrs["last_updated_user_id"] = attrs["last_updated_user_id"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBaseEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceBaseEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBaseEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceBaseEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_environment_type": m.BaseEnvironmentType, "effective_base_environment_type": m.EffectiveBaseEnvironmentType,
			"create_time":          m.CreateTime,
			"creator_user_id":      m.CreatorUserId,
			"display_name":         m.DisplayName,
			"filepath":             m.Filepath,
			"is_default":           m.IsDefault,
			"last_updated_user_id": m.LastUpdatedUserId,
			"message":              m.Message,
			"name":                 m.Name,
			"status":               m.Status,
			"update_time":          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceBaseEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_environment_type":           types.StringType,
			"effective_base_environment_type": types.StringType,
			"create_time":                     timetypes.RFC3339{}.Type(ctx),
			"creator_user_id":                 types.StringType,
			"display_name":                    types.StringType,
			"filepath":                        types.StringType,
			"is_default":                      types.BoolType,
			"last_updated_user_id":            types.StringType,
			"message":                         types.StringType,
			"name":                            types.StringType,
			"status":                          types.StringType,
			"update_time":                     timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// Metadata for the WorkspaceBaseEnvironment long-running operations. This
// message tracks the progress of the workspace base environment long-running
// process.
type WorkspaceBaseEnvironmentOperationMetadata_SdkV2 struct {
}

func (to *WorkspaceBaseEnvironmentOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceBaseEnvironmentOperationMetadata_SdkV2) {
}

func (to *WorkspaceBaseEnvironmentOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceBaseEnvironmentOperationMetadata_SdkV2) {
}

func (m WorkspaceBaseEnvironmentOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBaseEnvironmentOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceBaseEnvironmentOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBaseEnvironmentOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceBaseEnvironmentOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceBaseEnvironmentOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}
