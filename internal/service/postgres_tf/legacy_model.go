// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package postgres_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Branch_SdkV2 struct {
	// A timestamp indicating when the branch was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState types.String `tfsdk:"current_state"`
	// Whether the branch is the project's default branch. This field is only
	// returned on create/update responses. See effective_default for the value
	// that is actually applied to the branch.
	Default types.Bool `tfsdk:"default"`
	// Whether the branch is the project's default branch.
	EffectiveDefault types.Bool `tfsdk:"effective_default"`
	// Whether the branch is protected.
	EffectiveIsProtected types.Bool `tfsdk:"effective_is_protected"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	EffectiveSourceBranch types.String `tfsdk:"effective_source_branch"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	EffectiveSourceBranchLsn types.String `tfsdk:"effective_source_branch_lsn"`
	// The point in time on the source branch from which this branch was
	// created.
	EffectiveSourceBranchTime timetypes.RFC3339 `tfsdk:"effective_source_branch_time"`
	// Whether the branch is protected.
	IsProtected types.Bool `tfsdk:"is_protected"`
	// The logical size of the branch.
	LogicalSizeBytes types.Int64 `tfsdk:"logical_size_bytes"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"name"`
	// The project containing this branch. Format: projects/{project_id}
	Parent types.String `tfsdk:"parent"`
	// The pending state of the branch, if a state transition is in progress.
	PendingState types.String `tfsdk:"pending_state"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch types.String `tfsdk:"source_branch"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn types.String `tfsdk:"source_branch_lsn"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime timetypes.RFC3339 `tfsdk:"source_branch_time"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime timetypes.RFC3339 `tfsdk:"state_change_time"`
	// System generated unique ID for the branch.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Branch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Branch_SdkV2) {
	if !from.Default.IsUnknown() && !from.Default.IsNull() {
		// Default is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Default = from.Default
	}
	if !from.IsProtected.IsUnknown() && !from.IsProtected.IsNull() {
		// IsProtected is an input only field and not returned by the service, so we keep the value from the prior state.
		to.IsProtected = from.IsProtected
	}
	if !from.SourceBranch.IsUnknown() && !from.SourceBranch.IsNull() {
		// SourceBranch is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranch = from.SourceBranch
	}
	if !from.SourceBranchLsn.IsUnknown() && !from.SourceBranchLsn.IsNull() {
		// SourceBranchLsn is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranchLsn = from.SourceBranchLsn
	}
	if !from.SourceBranchTime.IsUnknown() && !from.SourceBranchTime.IsNull() {
		// SourceBranchTime is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranchTime = from.SourceBranchTime
	}
}

func (to *Branch_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Branch_SdkV2) {
	if !from.Default.IsUnknown() && !from.Default.IsNull() {
		// Default is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Default = from.Default
	}
	if !from.IsProtected.IsUnknown() && !from.IsProtected.IsNull() {
		// IsProtected is an input only field and not returned by the service, so we keep the value from the prior state.
		to.IsProtected = from.IsProtected
	}
	if !from.SourceBranch.IsUnknown() && !from.SourceBranch.IsNull() {
		// SourceBranch is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranch = from.SourceBranch
	}
	if !from.SourceBranchLsn.IsUnknown() && !from.SourceBranchLsn.IsNull() {
		// SourceBranchLsn is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranchLsn = from.SourceBranchLsn
	}
	if !from.SourceBranchTime.IsUnknown() && !from.SourceBranchTime.IsNull() {
		// SourceBranchTime is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SourceBranchTime = from.SourceBranchTime
	}
}

func (m Branch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["default"] = attrs["default"].SetOptional()
	attrs["default"] = attrs["default"].SetComputed()
	attrs["default"] = attrs["default"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_default"] = attrs["effective_default"].SetComputed()
	attrs["effective_is_protected"] = attrs["effective_is_protected"].SetComputed()
	attrs["effective_source_branch"] = attrs["effective_source_branch"].SetComputed()
	attrs["effective_source_branch_lsn"] = attrs["effective_source_branch_lsn"].SetComputed()
	attrs["effective_source_branch_time"] = attrs["effective_source_branch_time"].SetComputed()
	attrs["is_protected"] = attrs["is_protected"].SetOptional()
	attrs["is_protected"] = attrs["is_protected"].SetComputed()
	attrs["is_protected"] = attrs["is_protected"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["logical_size_bytes"] = attrs["logical_size_bytes"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["source_branch"] = attrs["source_branch"].SetOptional()
	attrs["source_branch"] = attrs["source_branch"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_branch"] = attrs["source_branch"].SetComputed()
	attrs["source_branch"] = attrs["source_branch"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].SetOptional()
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].SetComputed()
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["source_branch_time"] = attrs["source_branch_time"].SetOptional()
	attrs["source_branch_time"] = attrs["source_branch_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_branch_time"] = attrs["source_branch_time"].SetComputed()
	attrs["source_branch_time"] = attrs["source_branch_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["state_change_time"] = attrs["state_change_time"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Branch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Branch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Branch_SdkV2
// only implements ToObjectValue() and Type().
func (m Branch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":                  m.CreateTime,
			"current_state":                m.CurrentState,
			"default":                      m.Default,
			"effective_default":            m.EffectiveDefault,
			"effective_is_protected":       m.EffectiveIsProtected,
			"effective_source_branch":      m.EffectiveSourceBranch,
			"effective_source_branch_lsn":  m.EffectiveSourceBranchLsn,
			"effective_source_branch_time": m.EffectiveSourceBranchTime,
			"is_protected":                 m.IsProtected,
			"logical_size_bytes":           m.LogicalSizeBytes,
			"name":                         m.Name,
			"parent":                       m.Parent,
			"pending_state":                m.PendingState,
			"source_branch":                m.SourceBranch,
			"source_branch_lsn":            m.SourceBranchLsn,
			"source_branch_time":           m.SourceBranchTime,
			"state_change_time":            m.StateChangeTime,
			"uid":                          m.Uid,
			"update_time":                  m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Branch_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":                  timetypes.RFC3339{}.Type(ctx),
			"current_state":                types.StringType,
			"default":                      types.BoolType,
			"effective_default":            types.BoolType,
			"effective_is_protected":       types.BoolType,
			"effective_source_branch":      types.StringType,
			"effective_source_branch_lsn":  types.StringType,
			"effective_source_branch_time": timetypes.RFC3339{}.Type(ctx),
			"is_protected":                 types.BoolType,
			"logical_size_bytes":           types.Int64Type,
			"name":                         types.StringType,
			"parent":                       types.StringType,
			"pending_state":                types.StringType,
			"source_branch":                types.StringType,
			"source_branch_lsn":            types.StringType,
			"source_branch_time":           timetypes.RFC3339{}.Type(ctx),
			"state_change_time":            timetypes.RFC3339{}.Type(ctx),
			"uid":                          types.StringType,
			"update_time":                  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

type BranchOperationMetadata_SdkV2 struct {
}

func (to *BranchOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchOperationMetadata_SdkV2) {
}

func (to *BranchOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BranchOperationMetadata_SdkV2) {
}

func (m BranchOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m BranchOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m BranchOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateBranchRequest_SdkV2 struct {
	// The Branch to create.
	Branch types.List `tfsdk:"branch"`
	// The ID to use for the Branch, which will become the final component of
	// the branch's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	BranchId types.String `tfsdk:"-"`
	// The Project where this Branch will be created. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				// Recursively sync the fields of Branch
				toBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (to *CreateBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m CreateBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
	attrs["branch"] = attrs["branch"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":    m.Branch,
			"branch_id": m.BranchId,
			"parent":    m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"branch_id": types.StringType,
			"parent":    types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in CreateBranchRequest_SdkV2 as
// a Branch_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBranchRequest_SdkV2) GetBranch(ctx context.Context) (Branch_SdkV2, bool) {
	var e Branch_SdkV2
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v []Branch_SdkV2
	d := m.Branch.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBranch sets the value of the Branch field in CreateBranchRequest_SdkV2.
func (m *CreateBranchRequest_SdkV2) SetBranch(ctx context.Context, v Branch_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branch"]
	m.Branch = types.ListValueMust(t, vs)
}

type CreateEndpointRequest_SdkV2 struct {
	// The Endpoint to create.
	Endpoint types.List `tfsdk:"endpoint"`
	// The ID to use for the Endpoint, which will become the final component of
	// the endpoint's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	EndpointId types.String `tfsdk:"-"`
	// The Branch where this Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"endpoint_id": m.EndpointId,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"endpoint_id": types.StringType,
			"parent":      types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest_SdkV2.
func (m *CreateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

type CreateProjectRequest_SdkV2 struct {
	// The Project to create.
	Project types.List `tfsdk:"project"`
	// The ID to use for the Project, which will become the final component of
	// the project's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *CreateProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m CreateProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
	attrs["project"] = attrs["project"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["project_id"] = attrs["project_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project":    m.Project,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
			"project_id": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in CreateProjectRequest_SdkV2 as
// a Project_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProjectRequest_SdkV2) GetProject(ctx context.Context) (Project_SdkV2, bool) {
	var e Project_SdkV2
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v []Project_SdkV2
	d := m.Project.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProject sets the value of the Project field in CreateProjectRequest_SdkV2.
func (m *CreateProjectRequest_SdkV2) SetProject(ctx context.Context, v Project_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project"]
	m.Project = types.ListValueMust(t, vs)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto_SdkV2 struct {
	// @pbjson-skip
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

type DeleteBranchRequest_SdkV2 struct {
	// The name of the Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBranchRequest_SdkV2) {
}

func (to *DeleteBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteBranchRequest_SdkV2) {
}

func (m DeleteBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteEndpointRequest_SdkV2 struct {
	// The name of the Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (m DeleteEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteProjectRequest_SdkV2 struct {
	// The name of the Project to delete. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProjectRequest_SdkV2) {
}

func (to *DeleteProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteProjectRequest_SdkV2) {
}

func (m DeleteProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type Endpoint_SdkV2 struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// A timestamp indicating when the compute endpoint was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	CurrentState types.String `tfsdk:"current_state"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled types.Bool `tfsdk:"disabled"`
	// The maximum number of Compute Units.
	EffectiveAutoscalingLimitMaxCu types.Float64 `tfsdk:"effective_autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	EffectiveAutoscalingLimitMinCu types.Float64 `tfsdk:"effective_autoscaling_limit_min_cu"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	EffectiveDisabled types.Bool `tfsdk:"effective_disabled"`

	EffectivePoolerMode types.String `tfsdk:"effective_pooler_mode"`

	EffectiveSettings types.List `tfsdk:"effective_settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	EffectiveSuspendTimeoutDuration timetypes.GoDuration `tfsdk:"effective_suspend_timeout_duration"`
	// The endpoint type. There could be only one READ_WRITE endpoint per
	// branch.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host types.String `tfsdk:"host"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime timetypes.RFC3339 `tfsdk:"last_active_time"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"name"`
	// The branch containing this endpoint. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"parent"`

	PendingState types.String `tfsdk:"pending_state"`

	PoolerMode types.String `tfsdk:"pooler_mode"`

	Settings types.List `tfsdk:"settings"`
	// A timestamp indicating when the compute endpoint was last started.
	StartTime timetypes.RFC3339 `tfsdk:"start_time"`
	// A timestamp indicating when the compute endpoint was last suspended.
	SuspendTime timetypes.RFC3339 `tfsdk:"suspend_time"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
	// System generated unique ID for the endpoint.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Endpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint_SdkV2) {
	if !from.AutoscalingLimitMaxCu.IsUnknown() && !from.AutoscalingLimitMaxCu.IsNull() {
		// AutoscalingLimitMaxCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMaxCu = from.AutoscalingLimitMaxCu
	}
	if !from.AutoscalingLimitMinCu.IsUnknown() && !from.AutoscalingLimitMinCu.IsNull() {
		// AutoscalingLimitMinCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMinCu = from.AutoscalingLimitMinCu
	}
	if !from.Disabled.IsUnknown() && !from.Disabled.IsNull() {
		// Disabled is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Disabled = from.Disabled
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				// Recursively sync the fields of EffectiveSettings
				toEffectiveSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	if !from.PoolerMode.IsUnknown() && !from.PoolerMode.IsNull() {
		// PoolerMode is an input only field and not returned by the service, so we keep the value from the prior state.
		to.PoolerMode = from.PoolerMode
	}
	if !from.Settings.IsUnknown() && !from.Settings.IsNull() {
		// Settings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Settings = from.Settings
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
	if !from.SuspendTimeoutDuration.IsUnknown() && !from.SuspendTimeoutDuration.IsNull() {
		// SuspendTimeoutDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SuspendTimeoutDuration = from.SuspendTimeoutDuration
	}
}

func (to *Endpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Endpoint_SdkV2) {
	if !from.AutoscalingLimitMaxCu.IsUnknown() && !from.AutoscalingLimitMaxCu.IsNull() {
		// AutoscalingLimitMaxCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMaxCu = from.AutoscalingLimitMaxCu
	}
	if !from.AutoscalingLimitMinCu.IsUnknown() && !from.AutoscalingLimitMinCu.IsNull() {
		// AutoscalingLimitMinCu is an input only field and not returned by the service, so we keep the value from the prior state.
		to.AutoscalingLimitMinCu = from.AutoscalingLimitMinCu
	}
	if !from.Disabled.IsUnknown() && !from.Disabled.IsNull() {
		// Disabled is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Disabled = from.Disabled
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	if !from.PoolerMode.IsUnknown() && !from.PoolerMode.IsNull() {
		// PoolerMode is an input only field and not returned by the service, so we keep the value from the prior state.
		to.PoolerMode = from.PoolerMode
	}
	if !from.Settings.IsUnknown() && !from.Settings.IsNull() {
		// Settings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Settings = from.Settings
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
	if !from.SuspendTimeoutDuration.IsUnknown() && !from.SuspendTimeoutDuration.IsNull() {
		// SuspendTimeoutDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.SuspendTimeoutDuration = from.SuspendTimeoutDuration
	}
}

func (m Endpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetComputed()
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].(tfschema.Float64AttributeBuilder).AddPlanModifier(float64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetComputed()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].(tfschema.Float64AttributeBuilder).AddPlanModifier(float64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetOptional()
	attrs["disabled"] = attrs["disabled"].SetComputed()
	attrs["disabled"] = attrs["disabled"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_autoscaling_limit_max_cu"] = attrs["effective_autoscaling_limit_max_cu"].SetComputed()
	attrs["effective_autoscaling_limit_min_cu"] = attrs["effective_autoscaling_limit_min_cu"].SetComputed()
	attrs["effective_disabled"] = attrs["effective_disabled"].SetComputed()
	attrs["effective_pooler_mode"] = attrs["effective_pooler_mode"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_suspend_timeout_duration"] = attrs["effective_suspend_timeout_duration"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["host"] = attrs["host"].SetComputed()
	attrs["last_active_time"] = attrs["last_active_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetOptional()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["settings"] = attrs["settings"].SetComputed()
	attrs["settings"] = attrs["settings"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["start_time"] = attrs["start_time"].SetComputed()
	attrs["suspend_time"] = attrs["suspend_time"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Endpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Endpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(EndpointSettings_SdkV2{}),
		"settings":           reflect.TypeOf(EndpointSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint_SdkV2
// only implements ToObjectValue() and Type().
func (m Endpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu":           m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu":           m.AutoscalingLimitMinCu,
			"create_time":                        m.CreateTime,
			"current_state":                      m.CurrentState,
			"disabled":                           m.Disabled,
			"effective_autoscaling_limit_max_cu": m.EffectiveAutoscalingLimitMaxCu,
			"effective_autoscaling_limit_min_cu": m.EffectiveAutoscalingLimitMinCu,
			"effective_disabled":                 m.EffectiveDisabled,
			"effective_pooler_mode":              m.EffectivePoolerMode,
			"effective_settings":                 m.EffectiveSettings,
			"effective_suspend_timeout_duration": m.EffectiveSuspendTimeoutDuration,
			"endpoint_type":                      m.EndpointType,
			"host":                               m.Host,
			"last_active_time":                   m.LastActiveTime,
			"name":                               m.Name,
			"parent":                             m.Parent,
			"pending_state":                      m.PendingState,
			"pooler_mode":                        m.PoolerMode,
			"settings":                           m.Settings,
			"start_time":                         m.StartTime,
			"suspend_time":                       m.SuspendTime,
			"suspend_timeout_duration":           m.SuspendTimeoutDuration,
			"uid":                                m.Uid,
			"update_time":                        m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Endpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu":           types.Float64Type,
			"autoscaling_limit_min_cu":           types.Float64Type,
			"create_time":                        timetypes.RFC3339{}.Type(ctx),
			"current_state":                      types.StringType,
			"disabled":                           types.BoolType,
			"effective_autoscaling_limit_max_cu": types.Float64Type,
			"effective_autoscaling_limit_min_cu": types.Float64Type,
			"effective_disabled":                 types.BoolType,
			"effective_pooler_mode":              types.StringType,
			"effective_settings": basetypes.ListType{
				ElemType: EndpointSettings_SdkV2{}.Type(ctx),
			},
			"effective_suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
			"endpoint_type":                      types.StringType,
			"host":                               types.StringType,
			"last_active_time":                   timetypes.RFC3339{}.Type(ctx),
			"name":                               types.StringType,
			"parent":                             types.StringType,
			"pending_state":                      types.StringType,
			"pooler_mode":                        types.StringType,
			"settings": basetypes.ListType{
				ElemType: EndpointSettings_SdkV2{}.Type(ctx),
			},
			"start_time":               timetypes.RFC3339{}.Type(ctx),
			"suspend_time":             timetypes.RFC3339{}.Type(ctx),
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
			"uid":                      types.StringType,
			"update_time":              timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Endpoint_SdkV2 as
// a EndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetEffectiveSettings(ctx context.Context) (EndpointSettings_SdkV2, bool) {
	var e EndpointSettings_SdkV2
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v []EndpointSettings_SdkV2
	d := m.EffectiveSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetEffectiveSettings(ctx context.Context, v EndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_settings"]
	m.EffectiveSettings = types.ListValueMust(t, vs)
}

// GetSettings returns the value of the Settings field in Endpoint_SdkV2 as
// a EndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetSettings(ctx context.Context) (EndpointSettings_SdkV2, bool) {
	var e EndpointSettings_SdkV2
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v []EndpointSettings_SdkV2
	d := m.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetSettings(ctx context.Context, v EndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	m.Settings = types.ListValueMust(t, vs)
}

type EndpointOperationMetadata_SdkV2 struct {
}

func (to *EndpointOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointOperationMetadata_SdkV2) {
}

func (to *EndpointOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointOperationMetadata_SdkV2) {
}

func (m EndpointOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A collection of settings for a compute endpoint.
type EndpointSettings_SdkV2 struct {
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings types.Map `tfsdk:"pgbouncer_settings"`
}

func (to *EndpointSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSettings_SdkV2) {
}

func (to *EndpointSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointSettings_SdkV2) {
}

func (m EndpointSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()
	attrs["pgbouncer_settings"] = attrs["pgbouncer_settings"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pg_settings":        m.PgSettings,
			"pgbouncer_settings": m.PgbouncerSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pgbouncer_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in EndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings_SdkV2) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in EndpointSettings_SdkV2.
func (m *EndpointSettings_SdkV2) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in EndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings_SdkV2) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgbouncerSettings.IsNull() || m.PgbouncerSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgbouncerSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in EndpointSettings_SdkV2.
func (m *EndpointSettings_SdkV2) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
}

type GetBranchRequest_SdkV2 struct {
	// The name of the Branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBranchRequest_SdkV2) {
}

func (to *GetBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBranchRequest_SdkV2) {
}

func (m GetBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetEndpointRequest_SdkV2 struct {
	// The name of the Endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (m GetEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type GetProjectRequest_SdkV2 struct {
	// The name of the Project to retrieve. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectRequest_SdkV2) {
}

func (to *GetProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetProjectRequest_SdkV2) {
}

func (m GetProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListBranchesRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Branches. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListBranchesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesRequest_SdkV2) {
}

func (to *ListBranchesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBranchesRequest_SdkV2) {
}

func (m ListBranchesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBranchesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListBranchesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBranchesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListBranchesResponse_SdkV2 struct {
	// List of branches.
	Branches types.List `tfsdk:"branches"`
	// Pagination token to request the next page of branches.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListBranchesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesResponse_SdkV2) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (to *ListBranchesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBranchesResponse_SdkV2) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (m ListBranchesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branches"] = attrs["branches"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBranchesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListBranchesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branches": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBranchesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branches":        m.Branches,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branches": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBranches returns the value of the Branches field in ListBranchesResponse_SdkV2 as
// a slice of Branch_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBranchesResponse_SdkV2) GetBranches(ctx context.Context) ([]Branch_SdkV2, bool) {
	if m.Branches.IsNull() || m.Branches.IsUnknown() {
		return nil, false
	}
	var v []Branch_SdkV2
	d := m.Branches.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranches sets the value of the Branches field in ListBranchesResponse_SdkV2.
func (m *ListBranchesResponse_SdkV2) SetBranches(ctx context.Context, v []Branch_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branches"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Branches = types.ListValueMust(t, vs)
}

type ListEndpointsRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Endpoints. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (m ListEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListEndpointsResponse_SdkV2 struct {
	// List of endpoints.
	Endpoints types.List `tfsdk:"endpoints"`
	// Pagination token to request the next page of endpoints.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoints"] = attrs["endpoints"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse_SdkV2 as
// a slice of Endpoint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse_SdkV2) GetEndpoints(ctx context.Context) ([]Endpoint_SdkV2, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse_SdkV2.
func (m *ListEndpointsResponse_SdkV2) SetEndpoints(ctx context.Context, v []Endpoint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListProjectsRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Projects. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProjectsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsRequest_SdkV2) {
}

func (to *ListProjectsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProjectsRequest_SdkV2) {
}

func (m ListProjectsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProjectsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProjectsResponse_SdkV2 struct {
	// Pagination token to request the next page of projects.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of projects.
	Projects types.List `tfsdk:"projects"`
}

func (to *ListProjectsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsResponse_SdkV2) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (to *ListProjectsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProjectsResponse_SdkV2) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (m ListProjectsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["projects"] = attrs["projects"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProjectsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"projects":        m.Projects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"projects": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetProjects returns the value of the Projects field in ListProjectsResponse_SdkV2 as
// a slice of Project_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectsResponse_SdkV2) GetProjects(ctx context.Context) ([]Project_SdkV2, bool) {
	if m.Projects.IsNull() || m.Projects.IsUnknown() {
		return nil, false
	}
	var v []Project_SdkV2
	d := m.Projects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjects sets the value of the Projects field in ListProjectsResponse_SdkV2.
func (m *ListProjectsResponse_SdkV2) SetProjects(ctx context.Context, v []Project_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["projects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Projects = types.ListValueMust(t, vs)
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

type Project_SdkV2 struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime timetypes.RFC3339 `tfsdk:"compute_last_active_time"`
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	DefaultEndpointSettings types.List `tfsdk:"default_endpoint_settings"`
	// Human-readable project name.
	DisplayName types.String `tfsdk:"display_name"`

	EffectiveDefaultEndpointSettings types.List `tfsdk:"effective_default_endpoint_settings"`

	EffectiveDisplayName types.String `tfsdk:"effective_display_name"`

	EffectiveHistoryRetentionDuration timetypes.GoDuration `tfsdk:"effective_history_retention_duration"`

	EffectivePgVersion types.Int64 `tfsdk:"effective_pg_version"`

	EffectiveSettings types.List `tfsdk:"effective_settings"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration timetypes.GoDuration `tfsdk:"history_retention_duration"`
	// The resource name of the project. Format: projects/{project_id}
	Name types.String `tfsdk:"name"`
	// The major Postgres version number.
	PgVersion types.Int64 `tfsdk:"pg_version"`

	Settings types.List `tfsdk:"settings"`
	// The current space occupied by the project in storage. Synthetic storage
	// size combines the logical data size and Write-Ahead Log (WAL) size for
	// all branches in a project.
	SyntheticStorageSizeBytes types.Int64 `tfsdk:"synthetic_storage_size_bytes"`
	// System generated unique ID for the project.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the project was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Project_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project_SdkV2) {
	if !from.DefaultEndpointSettings.IsUnknown() && !from.DefaultEndpointSettings.IsNull() {
		// DefaultEndpointSettings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DefaultEndpointSettings = from.DefaultEndpointSettings
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of DefaultEndpointSettings
				toDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.DisplayName.IsUnknown() && !from.DisplayName.IsNull() {
		// DisplayName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DisplayName = from.DisplayName
	}
	if !from.EffectiveDefaultEndpointSettings.IsNull() && !from.EffectiveDefaultEndpointSettings.IsUnknown() {
		if toEffectiveDefaultEndpointSettings, ok := to.GetEffectiveDefaultEndpointSettings(ctx); ok {
			if fromEffectiveDefaultEndpointSettings, ok := from.GetEffectiveDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of EffectiveDefaultEndpointSettings
				toEffectiveDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveDefaultEndpointSettings)
				to.SetEffectiveDefaultEndpointSettings(ctx, toEffectiveDefaultEndpointSettings)
			}
		}
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				// Recursively sync the fields of EffectiveSettings
				toEffectiveSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	if !from.HistoryRetentionDuration.IsUnknown() && !from.HistoryRetentionDuration.IsNull() {
		// HistoryRetentionDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.HistoryRetentionDuration = from.HistoryRetentionDuration
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (to *Project_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Project_SdkV2) {
	if !from.DefaultEndpointSettings.IsUnknown() && !from.DefaultEndpointSettings.IsNull() {
		// DefaultEndpointSettings is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DefaultEndpointSettings = from.DefaultEndpointSettings
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.DisplayName.IsUnknown() && !from.DisplayName.IsNull() {
		// DisplayName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DisplayName = from.DisplayName
	}
	if !from.EffectiveDefaultEndpointSettings.IsNull() && !from.EffectiveDefaultEndpointSettings.IsUnknown() {
		if toEffectiveDefaultEndpointSettings, ok := to.GetEffectiveDefaultEndpointSettings(ctx); ok {
			if fromEffectiveDefaultEndpointSettings, ok := from.GetEffectiveDefaultEndpointSettings(ctx); ok {
				toEffectiveDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromEffectiveDefaultEndpointSettings)
				to.SetEffectiveDefaultEndpointSettings(ctx, toEffectiveDefaultEndpointSettings)
			}
		}
	}
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
	if !from.HistoryRetentionDuration.IsUnknown() && !from.HistoryRetentionDuration.IsNull() {
		// HistoryRetentionDuration is an input only field and not returned by the service, so we keep the value from the prior state.
		to.HistoryRetentionDuration = from.HistoryRetentionDuration
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m Project_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["compute_last_active_time"] = attrs["compute_last_active_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_default_endpoint_settings"] = attrs["effective_default_endpoint_settings"].SetComputed()
	attrs["effective_default_endpoint_settings"] = attrs["effective_default_endpoint_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_display_name"] = attrs["effective_display_name"].SetComputed()
	attrs["effective_history_retention_duration"] = attrs["effective_history_retention_duration"].SetComputed()
	attrs["effective_pg_version"] = attrs["effective_pg_version"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetOptional()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["settings"] = attrs["settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["synthetic_storage_size_bytes"] = attrs["synthetic_storage_size_bytes"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Project.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Project_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings":           reflect.TypeOf(ProjectDefaultEndpointSettings_SdkV2{}),
		"effective_default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings_SdkV2{}),
		"effective_settings":                  reflect.TypeOf(ProjectSettings_SdkV2{}),
		"settings":                            reflect.TypeOf(ProjectSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project_SdkV2
// only implements ToObjectValue() and Type().
func (m Project_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_logical_size_limit_bytes":      m.BranchLogicalSizeLimitBytes,
			"compute_last_active_time":             m.ComputeLastActiveTime,
			"create_time":                          m.CreateTime,
			"default_endpoint_settings":            m.DefaultEndpointSettings,
			"display_name":                         m.DisplayName,
			"effective_default_endpoint_settings":  m.EffectiveDefaultEndpointSettings,
			"effective_display_name":               m.EffectiveDisplayName,
			"effective_history_retention_duration": m.EffectiveHistoryRetentionDuration,
			"effective_pg_version":                 m.EffectivePgVersion,
			"effective_settings":                   m.EffectiveSettings,
			"history_retention_duration":           m.HistoryRetentionDuration,
			"name":                                 m.Name,
			"pg_version":                           m.PgVersion,
			"settings":                             m.Settings,
			"synthetic_storage_size_bytes":         m.SyntheticStorageSizeBytes,
			"uid":                                  m.Uid,
			"update_time":                          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Project_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes": types.Int64Type,
			"compute_last_active_time":        timetypes.RFC3339{}.Type(ctx),
			"create_time":                     timetypes.RFC3339{}.Type(ctx),
			"default_endpoint_settings": basetypes.ListType{
				ElemType: ProjectDefaultEndpointSettings_SdkV2{}.Type(ctx),
			},
			"display_name": types.StringType,
			"effective_default_endpoint_settings": basetypes.ListType{
				ElemType: ProjectDefaultEndpointSettings_SdkV2{}.Type(ctx),
			},
			"effective_display_name":               types.StringType,
			"effective_history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"effective_pg_version":                 types.Int64Type,
			"effective_settings": basetypes.ListType{
				ElemType: ProjectSettings_SdkV2{}.Type(ctx),
			},
			"history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"name":                       types.StringType,
			"pg_version":                 types.Int64Type,
			"settings": basetypes.ListType{
				ElemType: ProjectSettings_SdkV2{}.Type(ctx),
			},
			"synthetic_storage_size_bytes": types.Int64Type,
			"uid":                          types.StringType,
			"update_time":                  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in Project_SdkV2 as
// a ProjectDefaultEndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings_SdkV2, bool) {
	var e ProjectDefaultEndpointSettings_SdkV2
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v []ProjectDefaultEndpointSettings_SdkV2
	d := m.DefaultEndpointSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in Project_SdkV2.
func (m *Project_SdkV2) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_endpoint_settings"]
	m.DefaultEndpointSettings = types.ListValueMust(t, vs)
}

// GetEffectiveDefaultEndpointSettings returns the value of the EffectiveDefaultEndpointSettings field in Project_SdkV2 as
// a ProjectDefaultEndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetEffectiveDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings_SdkV2, bool) {
	var e ProjectDefaultEndpointSettings_SdkV2
	if m.EffectiveDefaultEndpointSettings.IsNull() || m.EffectiveDefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v []ProjectDefaultEndpointSettings_SdkV2
	d := m.EffectiveDefaultEndpointSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveDefaultEndpointSettings sets the value of the EffectiveDefaultEndpointSettings field in Project_SdkV2.
func (m *Project_SdkV2) SetEffectiveDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_default_endpoint_settings"]
	m.EffectiveDefaultEndpointSettings = types.ListValueMust(t, vs)
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Project_SdkV2 as
// a ProjectSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetEffectiveSettings(ctx context.Context) (ProjectSettings_SdkV2, bool) {
	var e ProjectSettings_SdkV2
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v []ProjectSettings_SdkV2
	d := m.EffectiveSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Project_SdkV2.
func (m *Project_SdkV2) SetEffectiveSettings(ctx context.Context, v ProjectSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_settings"]
	m.EffectiveSettings = types.ListValueMust(t, vs)
}

// GetSettings returns the value of the Settings field in Project_SdkV2 as
// a ProjectSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetSettings(ctx context.Context) (ProjectSettings_SdkV2, bool) {
	var e ProjectSettings_SdkV2
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v []ProjectSettings_SdkV2
	d := m.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in Project_SdkV2.
func (m *Project_SdkV2) SetSettings(ctx context.Context, v ProjectSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	m.Settings = types.ListValueMust(t, vs)
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings_SdkV2 struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings types.Map `tfsdk:"pgbouncer_settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *ProjectDefaultEndpointSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectDefaultEndpointSettings_SdkV2) {
}

func (to *ProjectDefaultEndpointSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectDefaultEndpointSettings_SdkV2) {
}

func (m ProjectDefaultEndpointSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()
	attrs["pgbouncer_settings"] = attrs["pgbouncer_settings"].SetOptional()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectDefaultEndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectDefaultEndpointSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectDefaultEndpointSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectDefaultEndpointSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"pg_settings":              m.PgSettings,
			"pgbouncer_settings":       m.PgbouncerSettings,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectDefaultEndpointSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pgbouncer_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in ProjectDefaultEndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings_SdkV2) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in ProjectDefaultEndpointSettings_SdkV2.
func (m *ProjectDefaultEndpointSettings_SdkV2) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in ProjectDefaultEndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings_SdkV2) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgbouncerSettings.IsNull() || m.PgbouncerSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgbouncerSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in ProjectDefaultEndpointSettings_SdkV2.
func (m *ProjectDefaultEndpointSettings_SdkV2) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
}

type ProjectOperationMetadata_SdkV2 struct {
}

func (to *ProjectOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectOperationMetadata_SdkV2) {
}

func (to *ProjectOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectOperationMetadata_SdkV2) {
}

func (m ProjectOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ProjectSettings_SdkV2 struct {
	// Sets wal_level=logical for all compute endpoints in this project. All
	// active endpoints will be suspended. Once enabled, logical replication
	// cannot be disabled.
	EnableLogicalReplication types.Bool `tfsdk:"enable_logical_replication"`
}

func (to *ProjectSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectSettings_SdkV2) {
}

func (to *ProjectSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectSettings_SdkV2) {
}

func (m ProjectSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_logical_replication"] = attrs["enable_logical_replication"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_logical_replication": m.EnableLogicalReplication,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_logical_replication": types.BoolType,
		},
	}
}

type UpdateBranchRequest_SdkV2 struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch types.List `tfsdk:"branch"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				// Recursively sync the fields of Branch
				toBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (to *UpdateBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m UpdateBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
	attrs["branch"] = attrs["branch"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":      m.Branch,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in UpdateBranchRequest_SdkV2 as
// a Branch_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBranchRequest_SdkV2) GetBranch(ctx context.Context) (Branch_SdkV2, bool) {
	var e Branch_SdkV2
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v []Branch_SdkV2
	d := m.Branch.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBranch sets the value of the Branch field in UpdateBranchRequest_SdkV2.
func (m *UpdateBranchRequest_SdkV2) SetBranch(ctx context.Context, v Branch_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branch"]
	m.Branch = types.ListValueMust(t, vs)
}

type UpdateEndpointRequest_SdkV2 struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.List `tfsdk:"endpoint"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m UpdateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in UpdateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in UpdateEndpointRequest_SdkV2.
func (m *UpdateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

type UpdateProjectRequest_SdkV2 struct {
	// The resource name of the project. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project types.List `tfsdk:"project"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *UpdateProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m UpdateProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
	attrs["project"] = attrs["project"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"project":     m.Project,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"project": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in UpdateProjectRequest_SdkV2 as
// a Project_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProjectRequest_SdkV2) GetProject(ctx context.Context) (Project_SdkV2, bool) {
	var e Project_SdkV2
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v []Project_SdkV2
	d := m.Project.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProject sets the value of the Project field in UpdateProjectRequest_SdkV2.
func (m *UpdateProjectRequest_SdkV2) SetProject(ctx context.Context, v Project_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project"]
	m.Project = types.ListValueMust(t, vs)
}
