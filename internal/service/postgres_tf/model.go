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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Branch struct {
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

func (to *Branch) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Branch) {
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

func (to *Branch) SyncFieldsDuringRead(ctx context.Context, from Branch) {
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

func (m Branch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Branch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Branch
// only implements ToObjectValue() and Type().
func (m Branch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Branch) Type(ctx context.Context) attr.Type {
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

type BranchOperationMetadata struct {
}

func (to *BranchOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchOperationMetadata) {
}

func (to *BranchOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from BranchOperationMetadata) {
}

func (m BranchOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchOperationMetadata
// only implements ToObjectValue() and Type().
func (m BranchOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m BranchOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateBranchRequest struct {
	// The Branch to create.
	Branch types.Object `tfsdk:"branch"`
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

func (to *CreateBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBranchRequest) {
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

func (to *CreateBranchRequest) SyncFieldsDuringRead(ctx context.Context, from CreateBranchRequest) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m CreateBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
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
func (m CreateBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBranchRequest
// only implements ToObjectValue() and Type().
func (m CreateBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":    m.Branch,
			"branch_id": m.BranchId,
			"parent":    m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":    Branch{}.Type(ctx),
			"branch_id": types.StringType,
			"parent":    types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in CreateBranchRequest as
// a Branch value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBranchRequest) GetBranch(ctx context.Context) (Branch, bool) {
	var e Branch
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v Branch
	d := m.Branch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranch sets the value of the Branch field in CreateBranchRequest.
func (m *CreateBranchRequest) SetBranch(ctx context.Context, v Branch) {
	vs := v.ToObjectValue(ctx)
	m.Branch = vs
}

type CreateEndpointRequest struct {
	// The Endpoint to create.
	Endpoint types.Object `tfsdk:"endpoint"`
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

func (to *CreateEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest) {
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

func (to *CreateEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
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
func (m CreateEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"endpoint_id": m.EndpointId,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint":    Endpoint{}.Type(ctx),
			"endpoint_id": types.StringType,
			"parent":      types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest as
// a Endpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest) GetEndpoint(ctx context.Context) (Endpoint, bool) {
	var e Endpoint
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v Endpoint
	d := m.Endpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest.
func (m *CreateEndpointRequest) SetEndpoint(ctx context.Context, v Endpoint) {
	vs := v.ToObjectValue(ctx)
	m.Endpoint = vs
}

type CreateProjectRequest struct {
	// The Project to create.
	Project types.Object `tfsdk:"project"`
	// The ID to use for the Project, which will become the final component of
	// the project's resource name.
	//
	// This value should be 4-63 characters, and valid characters are
	// /[a-z][0-9]-/.
	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProjectRequest) {
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

func (to *CreateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from CreateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m CreateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
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
func (m CreateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProjectRequest
// only implements ToObjectValue() and Type().
func (m CreateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project":    m.Project,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project":    Project{}.Type(ctx),
			"project_id": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in CreateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in CreateProjectRequest.
func (m *CreateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	// @pbjson-skip
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabricksServiceExceptionWithDetailsProto) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DatabricksServiceExceptionWithDetailsProto) Type(ctx context.Context) attr.Type {
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

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto) GetDetails(ctx context.Context) ([]types.Object, bool) {
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

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto.
func (m *DatabricksServiceExceptionWithDetailsProto) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteBranchRequest struct {
	// The name of the Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBranchRequest) {
}

func (to *DeleteBranchRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteBranchRequest) {
}

func (m DeleteBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBranchRequest
// only implements ToObjectValue() and Type().
func (m DeleteBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteEndpointRequest struct {
	// The name of the Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest) {
}

func (to *DeleteEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest) {
}

func (m DeleteEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteProjectRequest struct {
	// The name of the Project to delete. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProjectRequest) {
}

func (to *DeleteProjectRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteProjectRequest) {
}

func (m DeleteProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProjectRequest
// only implements ToObjectValue() and Type().
func (m DeleteProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type Endpoint struct {
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

	EffectiveSettings types.Object `tfsdk:"effective_settings"`
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

	Settings types.Object `tfsdk:"settings"`
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

func (to *Endpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint) {
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

func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
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

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["settings"] = attrs["settings"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
func (m Endpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(EndpointSettings{}),
		"settings":           reflect.TypeOf(EndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint
// only implements ToObjectValue() and Type().
func (m Endpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Endpoint) Type(ctx context.Context) attr.Type {
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
			"effective_settings":                 EndpointSettings{}.Type(ctx),
			"effective_suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
			"endpoint_type":                      types.StringType,
			"host":                               types.StringType,
			"last_active_time":                   timetypes.RFC3339{}.Type(ctx),
			"name":                               types.StringType,
			"parent":                             types.StringType,
			"pending_state":                      types.StringType,
			"pooler_mode":                        types.StringType,
			"settings":                           EndpointSettings{}.Type(ctx),
			"start_time":                         timetypes.RFC3339{}.Type(ctx),
			"suspend_time":                       timetypes.RFC3339{}.Type(ctx),
			"suspend_timeout_duration":           timetypes.GoDuration{}.Type(ctx),
			"uid":                                types.StringType,
			"update_time":                        timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Endpoint as
// a EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetEffectiveSettings(ctx context.Context) (EndpointSettings, bool) {
	var e EndpointSettings
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v EndpointSettings
	d := m.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Endpoint.
func (m *Endpoint) SetEffectiveSettings(ctx context.Context, v EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveSettings = vs
}

// GetSettings returns the value of the Settings field in Endpoint as
// a EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetSettings(ctx context.Context) (EndpointSettings, bool) {
	var e EndpointSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v EndpointSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in Endpoint.
func (m *Endpoint) SetSettings(ctx context.Context, v EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

type EndpointOperationMetadata struct {
}

func (to *EndpointOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointOperationMetadata) {
}

func (to *EndpointOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from EndpointOperationMetadata) {
}

func (m EndpointOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointOperationMetadata
// only implements ToObjectValue() and Type().
func (m EndpointOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A collection of settings for a compute endpoint.
type EndpointSettings struct {
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings types.Map `tfsdk:"pgbouncer_settings"`
}

func (to *EndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSettings) {
}

func (to *EndpointSettings) SyncFieldsDuringRead(ctx context.Context, from EndpointSettings) {
}

func (m EndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSettings
// only implements ToObjectValue() and Type().
func (m EndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pg_settings":        m.PgSettings,
			"pgbouncer_settings": m.PgbouncerSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointSettings) Type(ctx context.Context) attr.Type {
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

// GetPgSettings returns the value of the PgSettings field in EndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgSettings sets the value of the PgSettings field in EndpointSettings.
func (m *EndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in EndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in EndpointSettings.
func (m *EndpointSettings) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
}

type GetBranchRequest struct {
	// The name of the Branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBranchRequest) {
}

func (to *GetBranchRequest) SyncFieldsDuringRead(ctx context.Context, from GetBranchRequest) {
}

func (m GetBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBranchRequest
// only implements ToObjectValue() and Type().
func (m GetBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetEndpointRequest struct {
	// The name of the Endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest) {
}

func (to *GetEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest) {
}

func (m GetEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest) {
}

func (to *GetOperationRequest) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest) {
}

func (m GetOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest
// only implements ToObjectValue() and Type().
func (m GetOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetProjectRequest struct {
	// The name of the Project to retrieve. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectRequest) {
}

func (to *GetProjectRequest) SyncFieldsDuringRead(ctx context.Context, from GetProjectRequest) {
}

func (m GetProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectRequest
// only implements ToObjectValue() and Type().
func (m GetProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListBranchesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Branches. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListBranchesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesRequest) {
}

func (to *ListBranchesRequest) SyncFieldsDuringRead(ctx context.Context, from ListBranchesRequest) {
}

func (m ListBranchesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBranchesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesRequest
// only implements ToObjectValue() and Type().
func (m ListBranchesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListBranchesResponse struct {
	// List of branches.
	Branches types.List `tfsdk:"branches"`
	// Pagination token to request the next page of branches.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListBranchesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesResponse) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (to *ListBranchesResponse) SyncFieldsDuringRead(ctx context.Context, from ListBranchesResponse) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (m ListBranchesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBranchesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branches": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesResponse
// only implements ToObjectValue() and Type().
func (m ListBranchesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branches":        m.Branches,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branches": basetypes.ListType{
				ElemType: Branch{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBranches returns the value of the Branches field in ListBranchesResponse as
// a slice of Branch values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBranchesResponse) GetBranches(ctx context.Context) ([]Branch, bool) {
	if m.Branches.IsNull() || m.Branches.IsUnknown() {
		return nil, false
	}
	var v []Branch
	d := m.Branches.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranches sets the value of the Branches field in ListBranchesResponse.
func (m *ListBranchesResponse) SetBranches(ctx context.Context, v []Branch) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branches"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Branches = types.ListValueMust(t, vs)
}

type ListEndpointsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Endpoints. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest) {
}

func (to *ListEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest) {
}

func (m ListEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListEndpointsResponse struct {
	// List of endpoints.
	Endpoints types.List `tfsdk:"endpoints"`
	// Pagination token to request the next page of endpoints.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListEndpointsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: Endpoint{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse as
// a slice of Endpoint values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse) GetEndpoints(ctx context.Context) ([]Endpoint, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []Endpoint
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse.
func (m *ListEndpointsResponse) SetEndpoints(ctx context.Context, v []Endpoint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListProjectsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Projects. Requests first page
	// if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProjectsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsRequest) {
}

func (to *ListProjectsRequest) SyncFieldsDuringRead(ctx context.Context, from ListProjectsRequest) {
}

func (m ListProjectsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProjectsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsRequest
// only implements ToObjectValue() and Type().
func (m ListProjectsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProjectsResponse struct {
	// Pagination token to request the next page of projects.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of projects.
	Projects types.List `tfsdk:"projects"`
}

func (to *ListProjectsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (to *ListProjectsResponse) SyncFieldsDuringRead(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (m ListProjectsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProjectsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsResponse
// only implements ToObjectValue() and Type().
func (m ListProjectsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"projects":        m.Projects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"projects": basetypes.ListType{
				ElemType: Project{}.Type(ctx),
			},
		},
	}
}

// GetProjects returns the value of the Projects field in ListProjectsResponse as
// a slice of Project values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectsResponse) GetProjects(ctx context.Context) ([]Project, bool) {
	if m.Projects.IsNull() || m.Projects.IsUnknown() {
		return nil, false
	}
	var v []Project
	d := m.Projects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjects sets the value of the Projects field in ListProjectsResponse.
func (m *ListProjectsResponse) SetProjects(ctx context.Context, v []Project) {
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
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.Object `tfsdk:"error"`
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

func (to *Operation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation) {
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

func (to *Operation) SyncFieldsDuringRead(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
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
func (m Operation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation
// only implements ToObjectValue() and Type().
func (m Operation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Operation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done":     types.BoolType,
			"error":    DatabricksServiceExceptionWithDetailsProto{}.Type(ctx),
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation as
// a DatabricksServiceExceptionWithDetailsProto value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto, bool) {
	var e DatabricksServiceExceptionWithDetailsProto
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v DatabricksServiceExceptionWithDetailsProto
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in Operation.
func (m *Operation) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

type Project struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime timetypes.RFC3339 `tfsdk:"compute_last_active_time"`
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	DefaultEndpointSettings types.Object `tfsdk:"default_endpoint_settings"`
	// Human-readable project name.
	DisplayName types.String `tfsdk:"display_name"`

	EffectiveDefaultEndpointSettings types.Object `tfsdk:"effective_default_endpoint_settings"`

	EffectiveDisplayName types.String `tfsdk:"effective_display_name"`

	EffectiveHistoryRetentionDuration timetypes.GoDuration `tfsdk:"effective_history_retention_duration"`

	EffectivePgVersion types.Int64 `tfsdk:"effective_pg_version"`

	EffectiveSettings types.Object `tfsdk:"effective_settings"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration timetypes.GoDuration `tfsdk:"history_retention_duration"`
	// The resource name of the project. Format: projects/{project_id}
	Name types.String `tfsdk:"name"`
	// The major Postgres version number.
	PgVersion types.Int64 `tfsdk:"pg_version"`

	Settings types.Object `tfsdk:"settings"`
	// The current space occupied by the project in storage. Synthetic storage
	// size combines the logical data size and Write-Ahead Log (WAL) size for
	// all branches in a project.
	SyntheticStorageSizeBytes types.Int64 `tfsdk:"synthetic_storage_size_bytes"`
	// System generated unique ID for the project.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the project was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Project) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project) {
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

func (to *Project) SyncFieldsDuringRead(ctx context.Context, from Project) {
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

func (m Project) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["compute_last_active_time"] = attrs["compute_last_active_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_default_endpoint_settings"] = attrs["effective_default_endpoint_settings"].SetComputed()
	attrs["effective_display_name"] = attrs["effective_display_name"].SetComputed()
	attrs["effective_history_retention_duration"] = attrs["effective_history_retention_duration"].SetComputed()
	attrs["effective_pg_version"] = attrs["effective_pg_version"].SetComputed()
	attrs["effective_settings"] = attrs["effective_settings"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetOptional()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
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
func (m Project) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings":           reflect.TypeOf(ProjectDefaultEndpointSettings{}),
		"effective_default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings{}),
		"effective_settings":                  reflect.TypeOf(ProjectSettings{}),
		"settings":                            reflect.TypeOf(ProjectSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project
// only implements ToObjectValue() and Type().
func (m Project) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Project) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes":      types.Int64Type,
			"compute_last_active_time":             timetypes.RFC3339{}.Type(ctx),
			"create_time":                          timetypes.RFC3339{}.Type(ctx),
			"default_endpoint_settings":            ProjectDefaultEndpointSettings{}.Type(ctx),
			"display_name":                         types.StringType,
			"effective_default_endpoint_settings":  ProjectDefaultEndpointSettings{}.Type(ctx),
			"effective_display_name":               types.StringType,
			"effective_history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"effective_pg_version":                 types.Int64Type,
			"effective_settings":                   ProjectSettings{}.Type(ctx),
			"history_retention_duration":           timetypes.GoDuration{}.Type(ctx),
			"name":                                 types.StringType,
			"pg_version":                           types.Int64Type,
			"settings":                             ProjectSettings{}.Type(ctx),
			"synthetic_storage_size_bytes":         types.Int64Type,
			"uid":                                  types.StringType,
			"update_time":                          timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in Project as
// a ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings, bool) {
	var e ProjectDefaultEndpointSettings
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v ProjectDefaultEndpointSettings
	d := m.DefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in Project.
func (m *Project) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.DefaultEndpointSettings = vs
}

// GetEffectiveDefaultEndpointSettings returns the value of the EffectiveDefaultEndpointSettings field in Project as
// a ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetEffectiveDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings, bool) {
	var e ProjectDefaultEndpointSettings
	if m.EffectiveDefaultEndpointSettings.IsNull() || m.EffectiveDefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v ProjectDefaultEndpointSettings
	d := m.EffectiveDefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveDefaultEndpointSettings sets the value of the EffectiveDefaultEndpointSettings field in Project.
func (m *Project) SetEffectiveDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveDefaultEndpointSettings = vs
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in Project as
// a ProjectSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetEffectiveSettings(ctx context.Context) (ProjectSettings, bool) {
	var e ProjectSettings
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v ProjectSettings
	d := m.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in Project.
func (m *Project) SetEffectiveSettings(ctx context.Context, v ProjectSettings) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveSettings = vs
}

// GetSettings returns the value of the Settings field in Project as
// a ProjectSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetSettings(ctx context.Context) (ProjectSettings, bool) {
	var e ProjectSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v ProjectSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in Project.
func (m *Project) SetSettings(ctx context.Context, v ProjectSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings struct {
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

func (to *ProjectDefaultEndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectDefaultEndpointSettings) {
}

func (to *ProjectDefaultEndpointSettings) SyncFieldsDuringRead(ctx context.Context, from ProjectDefaultEndpointSettings) {
}

func (m ProjectDefaultEndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ProjectDefaultEndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectDefaultEndpointSettings
// only implements ToObjectValue() and Type().
func (m ProjectDefaultEndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ProjectDefaultEndpointSettings) Type(ctx context.Context) attr.Type {
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

// GetPgSettings returns the value of the PgSettings field in ProjectDefaultEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgSettings sets the value of the PgSettings field in ProjectDefaultEndpointSettings.
func (m *ProjectDefaultEndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in ProjectDefaultEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in ProjectDefaultEndpointSettings.
func (m *ProjectDefaultEndpointSettings) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
}

type ProjectOperationMetadata struct {
}

func (to *ProjectOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectOperationMetadata) {
}

func (to *ProjectOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from ProjectOperationMetadata) {
}

func (m ProjectOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectOperationMetadata
// only implements ToObjectValue() and Type().
func (m ProjectOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ProjectSettings struct {
	// Sets wal_level=logical for all compute endpoints in this project. All
	// active endpoints will be suspended. Once enabled, logical replication
	// cannot be disabled.
	EnableLogicalReplication types.Bool `tfsdk:"enable_logical_replication"`
}

func (to *ProjectSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectSettings) {
}

func (to *ProjectSettings) SyncFieldsDuringRead(ctx context.Context, from ProjectSettings) {
}

func (m ProjectSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ProjectSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectSettings
// only implements ToObjectValue() and Type().
func (m ProjectSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_logical_replication": m.EnableLogicalReplication,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_logical_replication": types.BoolType,
		},
	}
}

type UpdateBranchRequest struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch types.Object `tfsdk:"branch"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBranchRequest) {
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

func (to *UpdateBranchRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateBranchRequest) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m UpdateBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
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
func (m UpdateBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBranchRequest
// only implements ToObjectValue() and Type().
func (m UpdateBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":      m.Branch,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":      Branch{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in UpdateBranchRequest as
// a Branch value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBranchRequest) GetBranch(ctx context.Context) (Branch, bool) {
	var e Branch
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v Branch
	d := m.Branch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranch sets the value of the Branch field in UpdateBranchRequest.
func (m *UpdateBranchRequest) SetBranch(ctx context.Context, v Branch) {
	vs := v.ToObjectValue(ctx)
	m.Branch = vs
}

type UpdateEndpointRequest struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.Object `tfsdk:"endpoint"`
	// The resource name of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointRequest) {
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

func (to *UpdateEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m UpdateEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
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
func (m UpdateEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointRequest
// only implements ToObjectValue() and Type().
func (m UpdateEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint":    Endpoint{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in UpdateEndpointRequest as
// a Endpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointRequest) GetEndpoint(ctx context.Context) (Endpoint, bool) {
	var e Endpoint
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v Endpoint
	d := m.Endpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoint sets the value of the Endpoint field in UpdateEndpointRequest.
func (m *UpdateEndpointRequest) SetEndpoint(ctx context.Context, v Endpoint) {
	vs := v.ToObjectValue(ctx)
	m.Endpoint = vs
}

type UpdateProjectRequest struct {
	// The resource name of the project. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project types.Object `tfsdk:"project"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProjectRequest) {
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

func (to *UpdateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m UpdateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
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
func (m UpdateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProjectRequest
// only implements ToObjectValue() and Type().
func (m UpdateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"project":     m.Project,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"project":     Project{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in UpdateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in UpdateProjectRequest.
func (m *UpdateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}
