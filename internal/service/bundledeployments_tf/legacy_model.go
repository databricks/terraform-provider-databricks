// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package bundledeployments_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// A request to complete a Version.
type CompleteVersionRequest_SdkV2 struct {
	// The reason for completing the version. Must be a terminal reason:
	// VERSION_COMPLETE_SUCCESS, VERSION_COMPLETE_FAILURE, or
	// VERSION_COMPLETE_FORCE_ABORT.
	CompletionReason types.String `tfsdk:"completion_reason"`
	// If true, force-completes the version even if the caller is not the
	// original creator. The completion_reason must be
	// VERSION_COMPLETE_FORCE_ABORT when force is true.
	Force types.Bool `tfsdk:"force"`
	// The name of the version to complete. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name types.String `tfsdk:"-"`
}

func (to *CompleteVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CompleteVersionRequest_SdkV2) {
}

func (to *CompleteVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CompleteVersionRequest_SdkV2) {
}

func (m CompleteVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["completion_reason"] = attrs["completion_reason"].SetRequired()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CompleteVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CompleteVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CompleteVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CompleteVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"completion_reason": m.CompletionReason,
			"force":             m.Force,
			"name":              m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CompleteVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"completion_reason": types.StringType,
			"force":             types.BoolType,
			"name":              types.StringType,
		},
	}
}

type CreateDeploymentRequest_SdkV2 struct {
	// The deployment to create. Caller must set `initial_parent_path`; every
	// other field is populated by the service.
	Deployment types.List `tfsdk:"deployment"`
	// The ID to use for the deployment, which will become the final component
	// of the deployment's resource name (i.e. `deployments/{deployment_id}`).
	DeploymentId types.String `tfsdk:"-"`
}

func (to *CreateDeploymentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDeploymentRequest_SdkV2) {
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				// Recursively sync the fields of Deployment
				toDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
}

func (to *CreateDeploymentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDeploymentRequest_SdkV2) {
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				toDeployment.SyncFieldsDuringRead(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
}

func (m CreateDeploymentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["deployment"] = attrs["deployment"].SetRequired()
	attrs["deployment"] = attrs["deployment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["deployment_id"] = attrs["deployment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDeploymentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"deployment": reflect.TypeOf(Deployment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDeploymentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDeploymentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"deployment":    m.Deployment,
			"deployment_id": m.DeploymentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDeploymentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"deployment": basetypes.ListType{
				ElemType: Deployment_SdkV2{}.Type(ctx),
			},
			"deployment_id": types.StringType,
		},
	}
}

// GetDeployment returns the value of the Deployment field in CreateDeploymentRequest_SdkV2 as
// a Deployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDeploymentRequest_SdkV2) GetDeployment(ctx context.Context) (Deployment_SdkV2, bool) {
	var e Deployment_SdkV2
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v []Deployment_SdkV2
	d := m.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in CreateDeploymentRequest_SdkV2.
func (m *CreateDeploymentRequest_SdkV2) SetDeployment(ctx context.Context, v Deployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	m.Deployment = types.ListValueMust(t, vs)
}

type CreateOperationRequest_SdkV2 struct {
	// The resource operation to create.
	Operation types.List `tfsdk:"operation"`
	// The parent version where this operation will be recorded. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Parent types.String `tfsdk:"-"`
	// The key identifying the resource this operation applies to. Becomes the
	// final component of the operation's name.
	ResourceKey types.String `tfsdk:"-"`
}

func (to *CreateOperationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOperationRequest_SdkV2) {
	if !from.Operation.IsNull() && !from.Operation.IsUnknown() {
		if toOperation, ok := to.GetOperation(ctx); ok {
			if fromOperation, ok := from.GetOperation(ctx); ok {
				// Recursively sync the fields of Operation
				toOperation.SyncFieldsDuringCreateOrUpdate(ctx, fromOperation)
				to.SetOperation(ctx, toOperation)
			}
		}
	}
}

func (to *CreateOperationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateOperationRequest_SdkV2) {
	if !from.Operation.IsNull() && !from.Operation.IsUnknown() {
		if toOperation, ok := to.GetOperation(ctx); ok {
			if fromOperation, ok := from.GetOperation(ctx); ok {
				toOperation.SyncFieldsDuringRead(ctx, fromOperation)
				to.SetOperation(ctx, toOperation)
			}
		}
	}
}

func (m CreateOperationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["operation"] = attrs["operation"].SetRequired()
	attrs["operation"] = attrs["operation"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["resource_key"] = attrs["resource_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateOperationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"operation": reflect.TypeOf(Operation_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOperationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateOperationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operation":    m.Operation,
			"parent":       m.Parent,
			"resource_key": m.ResourceKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOperationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operation": basetypes.ListType{
				ElemType: Operation_SdkV2{}.Type(ctx),
			},
			"parent":       types.StringType,
			"resource_key": types.StringType,
		},
	}
}

// GetOperation returns the value of the Operation field in CreateOperationRequest_SdkV2 as
// a Operation_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateOperationRequest_SdkV2) GetOperation(ctx context.Context) (Operation_SdkV2, bool) {
	var e Operation_SdkV2
	if m.Operation.IsNull() || m.Operation.IsUnknown() {
		return e, false
	}
	var v []Operation_SdkV2
	d := m.Operation.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOperation sets the value of the Operation field in CreateOperationRequest_SdkV2.
func (m *CreateOperationRequest_SdkV2) SetOperation(ctx context.Context, v Operation_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["operation"]
	m.Operation = types.ListValueMust(t, vs)
}

type CreateVersionRequest_SdkV2 struct {
	// The parent deployment where this version will be created. Format:
	// deployments/{deployment_id}
	Parent types.String `tfsdk:"-"`
	// The version to create.
	Version types.List `tfsdk:"version"`
	// The ID to use for the version, which becomes the final component of the
	// version's resource name. A numeric string (base-10, fits in a signed
	// 64-bit integer) chosen by the caller; must be greater than or equal to 1.
	// Must be numerically greater than the deployment's most recent version
	// (see `version.previous_version_id`); it does not need to start at 1 or
	// increase by exactly 1. If the value is not numerically greater, the
	// server returns `INVALID_PARAMETER_VALUE`.
	VersionId types.String `tfsdk:"-"`
}

func (to *CreateVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVersionRequest_SdkV2) {
	if !from.Version.IsNull() && !from.Version.IsUnknown() {
		if toVersion, ok := to.GetVersion(ctx); ok {
			if fromVersion, ok := from.GetVersion(ctx); ok {
				// Recursively sync the fields of Version
				toVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromVersion)
				to.SetVersion(ctx, toVersion)
			}
		}
	}
}

func (to *CreateVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateVersionRequest_SdkV2) {
	if !from.Version.IsNull() && !from.Version.IsUnknown() {
		if toVersion, ok := to.GetVersion(ctx); ok {
			if fromVersion, ok := from.GetVersion(ctx); ok {
				toVersion.SyncFieldsDuringRead(ctx, fromVersion)
				to.SetVersion(ctx, toVersion)
			}
		}
	}
}

func (m CreateVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["version"] = attrs["version"].SetRequired()
	attrs["version"] = attrs["version"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["version_id"] = attrs["version_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"version": reflect.TypeOf(Version_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":     m.Parent,
			"version":    m.Version,
			"version_id": m.VersionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"version": basetypes.ListType{
				ElemType: Version_SdkV2{}.Type(ctx),
			},
			"version_id": types.StringType,
		},
	}
}

// GetVersion returns the value of the Version field in CreateVersionRequest_SdkV2 as
// a Version_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateVersionRequest_SdkV2) GetVersion(ctx context.Context) (Version_SdkV2, bool) {
	var e Version_SdkV2
	if m.Version.IsNull() || m.Version.IsUnknown() {
		return e, false
	}
	var v []Version_SdkV2
	d := m.Version.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVersion sets the value of the Version field in CreateVersionRequest_SdkV2.
func (m *CreateVersionRequest_SdkV2) SetVersion(ctx context.Context, v Version_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["version"]
	m.Version = types.ListValueMust(t, vs)
}

type DeleteDeploymentRequest_SdkV2 struct {
	// Resource name of the deployment to delete. Format:
	// deployments/{deployment_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDeploymentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDeploymentRequest_SdkV2) {
}

func (to *DeleteDeploymentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDeploymentRequest_SdkV2) {
}

func (m DeleteDeploymentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDeploymentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDeploymentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDeploymentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDeploymentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A bundle deployment registered with the control plane.
type Deployment_SdkV2 struct {
	// When the deployment was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The user who created the deployment (email or principal name).
	CreatedBy types.String `tfsdk:"created_by"`
	// Bundle target deployment mode (development or production), derived from
	// the most recent version's mode.
	DeploymentMode types.String `tfsdk:"deployment_mode"`
	// When the deployment was destroyed (i.e. `bundle destroy` completed).
	// Unset if the deployment has not been destroyed. Named destroy_time (not
	// delete_time) because this tracks the `databricks bundle destroy` command,
	// not the API-level deletion.
	DestroyTime timetypes.RFC3339 `tfsdk:"destroy_time"`
	// The user who destroyed the deployment (email or principal name). Unset if
	// the deployment has not been destroyed.
	DestroyedBy types.String `tfsdk:"destroyed_by"`
	// Human-readable name for the deployment. Output only: it is denormalized
	// from the latest version, not set directly on the deployment.
	DisplayName types.String `tfsdk:"display_name"`
	// Git provenance of the deployment's source, derived from the latest
	// version.
	GitInfo types.List `tfsdk:"git_info"`
	// The workspace path of the folder where the deployment is initially
	// created. Includes a leading slash and no trailing slash. On create, the
	// deployment is registered as a typed BUNDLE_DEPLOYMENT tree node under
	// this folder, which must already exist. This field is input only and is
	// not returned in create, get, or list responses. The service rejects
	// create requests that omit it.
	InitialParentPath types.String `tfsdk:"initial_parent_path"`
	// The version_id of the most recent deployment version.
	LastVersionId types.String `tfsdk:"last_version_id"`
	// Resource name of the deployment. Format: deployments/{deployment_id}
	Name types.String `tfsdk:"name"`
	// Current status of the deployment.
	Status types.String `tfsdk:"status"`
	// The bundle target name associated with this deployment. Output only: it
	// is denormalized from the latest version, not set directly on the
	// deployment.
	TargetName types.String `tfsdk:"target_name"`
	// When the deployment was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Workspace location of the deployment, derived from the latest version.
	WorkspaceInfo types.List `tfsdk:"workspace_info"`
}

func (to *Deployment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Deployment_SdkV2) {
	if !from.GitInfo.IsNull() && !from.GitInfo.IsUnknown() {
		if toGitInfo, ok := to.GetGitInfo(ctx); ok {
			if fromGitInfo, ok := from.GetGitInfo(ctx); ok {
				// Recursively sync the fields of GitInfo
				toGitInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGitInfo)
				to.SetGitInfo(ctx, toGitInfo)
			}
		}
	}
	if !from.InitialParentPath.IsUnknown() && !from.InitialParentPath.IsNull() {
		// InitialParentPath is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialParentPath = from.InitialParentPath
	}
	if !from.WorkspaceInfo.IsNull() && !from.WorkspaceInfo.IsUnknown() {
		if toWorkspaceInfo, ok := to.GetWorkspaceInfo(ctx); ok {
			if fromWorkspaceInfo, ok := from.GetWorkspaceInfo(ctx); ok {
				// Recursively sync the fields of WorkspaceInfo
				toWorkspaceInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceInfo)
				to.SetWorkspaceInfo(ctx, toWorkspaceInfo)
			}
		}
	}
}

func (to *Deployment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Deployment_SdkV2) {
	if !from.GitInfo.IsNull() && !from.GitInfo.IsUnknown() {
		if toGitInfo, ok := to.GetGitInfo(ctx); ok {
			if fromGitInfo, ok := from.GetGitInfo(ctx); ok {
				toGitInfo.SyncFieldsDuringRead(ctx, fromGitInfo)
				to.SetGitInfo(ctx, toGitInfo)
			}
		}
	}
	if !from.InitialParentPath.IsUnknown() && !from.InitialParentPath.IsNull() {
		// InitialParentPath is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialParentPath = from.InitialParentPath
	}
	if !from.WorkspaceInfo.IsNull() && !from.WorkspaceInfo.IsUnknown() {
		if toWorkspaceInfo, ok := to.GetWorkspaceInfo(ctx); ok {
			if fromWorkspaceInfo, ok := from.GetWorkspaceInfo(ctx); ok {
				toWorkspaceInfo.SyncFieldsDuringRead(ctx, fromWorkspaceInfo)
				to.SetWorkspaceInfo(ctx, toWorkspaceInfo)
			}
		}
	}
}

func (m Deployment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["deployment_mode"] = attrs["deployment_mode"].SetComputed()
	attrs["destroy_time"] = attrs["destroy_time"].SetComputed()
	attrs["destroyed_by"] = attrs["destroyed_by"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["git_info"] = attrs["git_info"].SetComputed()
	attrs["git_info"] = attrs["git_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["initial_parent_path"] = attrs["initial_parent_path"].SetOptional()
	attrs["initial_parent_path"] = attrs["initial_parent_path"].SetComputed()
	attrs["initial_parent_path"] = attrs["initial_parent_path"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["last_version_id"] = attrs["last_version_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["target_name"] = attrs["target_name"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_info"] = attrs["workspace_info"].SetComputed()
	attrs["workspace_info"] = attrs["workspace_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Deployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Deployment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_info":       reflect.TypeOf(GitInfo_SdkV2{}),
		"workspace_info": reflect.TypeOf(WorkspaceInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Deployment_SdkV2
// only implements ToObjectValue() and Type().
func (m Deployment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":         m.CreateTime,
			"created_by":          m.CreatedBy,
			"deployment_mode":     m.DeploymentMode,
			"destroy_time":        m.DestroyTime,
			"destroyed_by":        m.DestroyedBy,
			"display_name":        m.DisplayName,
			"git_info":            m.GitInfo,
			"initial_parent_path": m.InitialParentPath,
			"last_version_id":     m.LastVersionId,
			"name":                m.Name,
			"status":              m.Status,
			"target_name":         m.TargetName,
			"update_time":         m.UpdateTime,
			"workspace_info":      m.WorkspaceInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Deployment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"created_by":      types.StringType,
			"deployment_mode": types.StringType,
			"destroy_time":    timetypes.RFC3339{}.Type(ctx),
			"destroyed_by":    types.StringType,
			"display_name":    types.StringType,
			"git_info": basetypes.ListType{
				ElemType: GitInfo_SdkV2{}.Type(ctx),
			},
			"initial_parent_path": types.StringType,
			"last_version_id":     types.StringType,
			"name":                types.StringType,
			"status":              types.StringType,
			"target_name":         types.StringType,
			"update_time":         timetypes.RFC3339{}.Type(ctx),
			"workspace_info": basetypes.ListType{
				ElemType: WorkspaceInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGitInfo returns the value of the GitInfo field in Deployment_SdkV2 as
// a GitInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Deployment_SdkV2) GetGitInfo(ctx context.Context) (GitInfo_SdkV2, bool) {
	var e GitInfo_SdkV2
	if m.GitInfo.IsNull() || m.GitInfo.IsUnknown() {
		return e, false
	}
	var v []GitInfo_SdkV2
	d := m.GitInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitInfo sets the value of the GitInfo field in Deployment_SdkV2.
func (m *Deployment_SdkV2) SetGitInfo(ctx context.Context, v GitInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["git_info"]
	m.GitInfo = types.ListValueMust(t, vs)
}

// GetWorkspaceInfo returns the value of the WorkspaceInfo field in Deployment_SdkV2 as
// a WorkspaceInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Deployment_SdkV2) GetWorkspaceInfo(ctx context.Context) (WorkspaceInfo_SdkV2, bool) {
	var e WorkspaceInfo_SdkV2
	if m.WorkspaceInfo.IsNull() || m.WorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v []WorkspaceInfo_SdkV2
	d := m.WorkspaceInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceInfo sets the value of the WorkspaceInfo field in Deployment_SdkV2.
func (m *Deployment_SdkV2) SetWorkspaceInfo(ctx context.Context, v WorkspaceInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_info"]
	m.WorkspaceInfo = types.ListValueMust(t, vs)
}

type GetDeploymentRequest_SdkV2 struct {
	// Resource name of the deployment to retrieve. Format:
	// deployments/{deployment_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetDeploymentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDeploymentRequest_SdkV2) {
}

func (to *GetDeploymentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDeploymentRequest_SdkV2) {
}

func (m GetDeploymentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDeploymentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDeploymentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDeploymentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDeploymentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest_SdkV2 struct {
	// The name of the resource operation to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
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

type GetResourceRequest_SdkV2 struct {
	// The name of the resource to retrieve. Format:
	// deployments/{deployment_id}/resources/{resource_key}
	Name types.String `tfsdk:"-"`
}

func (to *GetResourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetResourceRequest_SdkV2) {
}

func (to *GetResourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetResourceRequest_SdkV2) {
}

func (m GetResourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetResourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetResourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetResourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetResourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetVersionRequest_SdkV2 struct {
	// The name of the version to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVersionRequest_SdkV2) {
}

func (to *GetVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetVersionRequest_SdkV2) {
}

func (m GetVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Git provenance of a bundle's source, captured at deploy time. Lets consumers
// link a deployed resource back to its source in version control.
type GitInfo_SdkV2 struct {
	// Branch the source was deployed from.
	Branch types.String `tfsdk:"branch"`
	// Commit SHA of the deployed source.
	Commit types.String `tfsdk:"commit"`
	// URL of the git remote the source was deployed from.
	OriginUrl types.String `tfsdk:"origin_url"`
}

func (to *GitInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GitInfo_SdkV2) {
}

func (to *GitInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GitInfo_SdkV2) {
}

func (m GitInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["commit"] = attrs["commit"].SetOptional()
	attrs["origin_url"] = attrs["origin_url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GitInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GitInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m GitInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":     m.Branch,
			"commit":     m.Commit,
			"origin_url": m.OriginUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GitInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":     types.StringType,
			"commit":     types.StringType,
			"origin_url": types.StringType,
		},
	}
}

// A request to send a heartbeat for a Version.
type HeartbeatRequest_SdkV2 struct {
	// The version whose lock to renew. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name types.String `tfsdk:"-"`
}

func (to *HeartbeatRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HeartbeatRequest_SdkV2) {
}

func (to *HeartbeatRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from HeartbeatRequest_SdkV2) {
}

func (m HeartbeatRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in HeartbeatRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m HeartbeatRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HeartbeatRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m HeartbeatRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HeartbeatRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Response for Heartbeat.
type HeartbeatResponse_SdkV2 struct {
	// The new lock expiry time after renewal.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`
}

func (to *HeartbeatResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HeartbeatResponse_SdkV2) {
}

func (to *HeartbeatResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from HeartbeatResponse_SdkV2) {
}

func (m HeartbeatResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["expire_time"] = attrs["expire_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in HeartbeatResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m HeartbeatResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HeartbeatResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m HeartbeatResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expire_time": m.ExpireTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HeartbeatResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expire_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

type ListDeploymentsRequest_SdkV2 struct {
	// The maximum number of deployments to return. The service may return fewer
	// than this value. If unspecified, at most 50 deployments will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListDeployments` call. Provide
	// this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDeploymentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDeploymentsRequest_SdkV2) {
}

func (to *ListDeploymentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDeploymentsRequest_SdkV2) {
}

func (m ListDeploymentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDeploymentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDeploymentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDeploymentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDeploymentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDeploymentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response for ListDeployments.
type ListDeploymentsResponse_SdkV2 struct {
	// The deployments from the queried workspace.
	Deployments types.List `tfsdk:"deployments"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDeploymentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDeploymentsResponse_SdkV2) {
	if !from.Deployments.IsNull() && !from.Deployments.IsUnknown() && to.Deployments.IsNull() && len(from.Deployments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Deployments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Deployments = from.Deployments
	}
}

func (to *ListDeploymentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDeploymentsResponse_SdkV2) {
	if !from.Deployments.IsNull() && !from.Deployments.IsUnknown() && to.Deployments.IsNull() && len(from.Deployments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Deployments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Deployments = from.Deployments
	}
}

func (m ListDeploymentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["deployments"] = attrs["deployments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDeploymentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDeploymentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"deployments": reflect.TypeOf(Deployment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDeploymentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDeploymentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"deployments":     m.Deployments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDeploymentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"deployments": basetypes.ListType{
				ElemType: Deployment_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDeployments returns the value of the Deployments field in ListDeploymentsResponse_SdkV2 as
// a slice of Deployment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDeploymentsResponse_SdkV2) GetDeployments(ctx context.Context) ([]Deployment_SdkV2, bool) {
	if m.Deployments.IsNull() || m.Deployments.IsUnknown() {
		return nil, false
	}
	var v []Deployment_SdkV2
	d := m.Deployments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployments sets the value of the Deployments field in ListDeploymentsResponse_SdkV2.
func (m *ListDeploymentsResponse_SdkV2) SetDeployments(ctx context.Context, v []Deployment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["deployments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Deployments = types.ListValueMust(t, vs)
}

type ListOperationsRequest_SdkV2 struct {
	// The maximum number of operations to return. The service may return fewer
	// than this value. If unspecified, at most 50 operations will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListOperations` call. Provide
	// this to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The parent version. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListOperationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOperationsRequest_SdkV2) {
}

func (to *ListOperationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListOperationsRequest_SdkV2) {
}

func (m ListOperationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOperationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListOperationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOperationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListOperationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOperationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for ListOperations.
type ListOperationsResponse_SdkV2 struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The resource operations under the specified version.
	Operations types.List `tfsdk:"operations"`
}

func (to *ListOperationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOperationsResponse_SdkV2) {
	if !from.Operations.IsNull() && !from.Operations.IsUnknown() && to.Operations.IsNull() && len(from.Operations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Operations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Operations = from.Operations
	}
}

func (to *ListOperationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListOperationsResponse_SdkV2) {
	if !from.Operations.IsNull() && !from.Operations.IsUnknown() && to.Operations.IsNull() && len(from.Operations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Operations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Operations = from.Operations
	}
}

func (m ListOperationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["operations"] = attrs["operations"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOperationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListOperationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"operations": reflect.TypeOf(Operation_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOperationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListOperationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"operations":      m.Operations,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOperationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"operations": basetypes.ListType{
				ElemType: Operation_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOperations returns the value of the Operations field in ListOperationsResponse_SdkV2 as
// a slice of Operation_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListOperationsResponse_SdkV2) GetOperations(ctx context.Context) ([]Operation_SdkV2, bool) {
	if m.Operations.IsNull() || m.Operations.IsUnknown() {
		return nil, false
	}
	var v []Operation_SdkV2
	d := m.Operations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOperations sets the value of the Operations field in ListOperationsResponse_SdkV2.
func (m *ListOperationsResponse_SdkV2) SetOperations(ctx context.Context, v []Operation_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["operations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Operations = types.ListValueMust(t, vs)
}

type ListResourcesRequest_SdkV2 struct {
	// The maximum number of resources to return. The service may return fewer
	// than this value. If unspecified, at most 50 resources will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListResources` call. Provide this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The parent deployment. Format: deployments/{deployment_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListResourcesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResourcesRequest_SdkV2) {
}

func (to *ListResourcesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListResourcesRequest_SdkV2) {
}

func (m ListResourcesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListResourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListResourcesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResourcesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListResourcesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListResourcesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for ListResources.
type ListResourcesResponse_SdkV2 struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The resources under the specified deployment.
	Resources types.List `tfsdk:"resources"`
}

func (to *ListResourcesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResourcesResponse_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (to *ListResourcesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListResourcesResponse_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (m ListResourcesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["resources"] = attrs["resources"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListResourcesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListResourcesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(Resource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResourcesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListResourcesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"resources":       m.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListResourcesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"resources": basetypes.ListType{
				ElemType: Resource_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResources returns the value of the Resources field in ListResourcesResponse_SdkV2 as
// a slice of Resource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListResourcesResponse_SdkV2) GetResources(ctx context.Context) ([]Resource_SdkV2, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []Resource_SdkV2
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in ListResourcesResponse_SdkV2.
func (m *ListResourcesResponse_SdkV2) SetResources(ctx context.Context, v []Resource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

type ListVersionsRequest_SdkV2 struct {
	// The maximum number of versions to return. The service may return fewer
	// than this value. If unspecified, at most 50 versions will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListVersions` call. Provide this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The parent deployment. Format: deployments/{deployment_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVersionsRequest_SdkV2) {
}

func (to *ListVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVersionsRequest_SdkV2) {
}

func (m ListVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for ListVersions.
type ListVersionsResponse_SdkV2 struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The versions under the specified deployment.
	Versions types.List `tfsdk:"versions"`
}

func (to *ListVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVersionsResponse_SdkV2) {
	if !from.Versions.IsNull() && !from.Versions.IsUnknown() && to.Versions.IsNull() && len(from.Versions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Versions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Versions = from.Versions
	}
}

func (to *ListVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVersionsResponse_SdkV2) {
	if !from.Versions.IsNull() && !from.Versions.IsUnknown() && to.Versions.IsNull() && len(from.Versions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Versions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Versions = from.Versions
	}
}

func (m ListVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["versions"] = attrs["versions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"versions": reflect.TypeOf(Version_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"versions":        m.Versions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"versions": basetypes.ListType{
				ElemType: Version_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVersions returns the value of the Versions field in ListVersionsResponse_SdkV2 as
// a slice of Version_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListVersionsResponse_SdkV2) GetVersions(ctx context.Context) ([]Version_SdkV2, bool) {
	if m.Versions.IsNull() || m.Versions.IsUnknown() {
		return nil, false
	}
	var v []Version_SdkV2
	d := m.Versions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVersions sets the value of the Versions field in ListVersionsResponse_SdkV2.
func (m *ListVersionsResponse_SdkV2) SetVersions(ctx context.Context, v []Version_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Versions = types.ListValueMust(t, vs)
}

// An operation on a single resource performed during a version. Operations
// record the result of applying a resource change to the workspace. Most fields
// are immutable once recorded; `state`, `error_message`, `resource_id`, and
// `status` may be updated afterwards (via UpdateOperation), guarded by
// `sequence_id` for optimistic concurrency control.
type Operation_SdkV2 struct {
	// The type of operation performed on this resource.
	ActionType types.String `tfsdk:"action_type"`
	// When the operation was recorded.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Error message if the operation failed. Set when status is
	// OPERATION_STATUS_FAILED. Captures the error encountered while applying
	// the resource to the workspace. Mutable: may be updated after creation via
	// UpdateOperation; setting it to an empty string clears it. After an update
	// is applied, an operation whose status is OPERATION_STATUS_SUCCEEDED
	// cannot carry an error_message.
	ErrorMessage types.String `tfsdk:"error_message"`
	// Resource name of the operation. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name types.String `tfsdk:"name"`
	// ID of the actual resource in the workspace (e.g. the job ID, pipeline
	// ID). Optional at creation: CREATE and RECREATE operations produce a new
	// resource whose ID is not yet known when the operation is recorded.
	// Mutable: may be filled in (or corrected) later via UpdateOperation once
	// the ID is known.
	ResourceId types.String `tfsdk:"resource_id"`
	// Resource identifier within the bundle (e.g. "jobs.foo", "pipelines.bar",
	// "jobs.foo.permissions", "files.<rel-path>"). Can be an arbitrary UTF-8
	// encoded string key. This key links the operation to the corresponding
	// deployment-level Resource.
	ResourceKey types.String `tfsdk:"resource_key"`
	// The type of the deployment resource this operation applies to. Derived
	// from the `resource_key` prefix (e.g. "jobs" → JOB); the caller does not
	// set this field.
	ResourceType types.String `tfsdk:"resource_type"`
	// Serialized local config state after the operation. Should be unset for
	// delete operations. Mutable: may be updated after creation via
	// UpdateOperation. When updating, the caller must echo the last-observed
	// `sequence_id` as a concurrency precondition.
	State jsontypes.Normalized `tfsdk:"state"`
	// Whether the operation succeeded or failed. Mutable: may be updated after
	// creation via UpdateOperation, e.g. when an operation recorded as failed
	// is retried and eventually succeeds. A succeeded operation cannot carry an
	// `error_message`.
	Status types.String `tfsdk:"status"`
}

func (to *Operation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation_SdkV2) {
}

func (to *Operation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Operation_SdkV2) {
}

func (m Operation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action_type"] = attrs["action_type"].SetRequired()
	attrs["action_type"] = attrs["action_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["resource_key"] = attrs["resource_key"].SetOptional()
	attrs["resource_key"] = attrs["resource_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["resource_type"] = attrs["resource_type"].SetComputed()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["status"] = attrs["status"].SetRequired()

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
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation_SdkV2
// only implements ToObjectValue() and Type().
func (m Operation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action_type":   m.ActionType,
			"create_time":   m.CreateTime,
			"error_message": m.ErrorMessage,
			"name":          m.Name,
			"resource_id":   m.ResourceId,
			"resource_key":  m.ResourceKey,
			"resource_type": m.ResourceType,
			"state":         m.State,
			"status":        m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action_type":   types.StringType,
			"create_time":   timetypes.RFC3339{}.Type(ctx),
			"error_message": types.StringType,
			"name":          types.StringType,
			"resource_id":   types.StringType,
			"resource_key":  types.StringType,
			"resource_type": types.StringType,
			"state":         jsontypes.NormalizedType{},
			"status":        types.StringType,
		},
	}
}

// A resource managed by a deployment. Resources are implicitly created,
// updated, or deleted when operations are recorded on a version.
type Resource_SdkV2 struct {
	// The action performed on this resource during the last version.
	LastActionType types.String `tfsdk:"last_action_type"`
	// The version_id of the last version where this resource was updated.
	LastVersionId types.String `tfsdk:"last_version_id"`
	// Resource name. Format:
	// deployments/{deployment_id}/resources/{resource_key}
	Name types.String `tfsdk:"name"`
	// ID that references the actual resource in the workspace (e.g. the job ID,
	// pipeline ID).
	ResourceId types.String `tfsdk:"resource_id"`
	// Resource identifier within the bundle (e.g. "jobs.foo", "pipelines.bar",
	// "jobs.foo.permissions").
	ResourceKey types.String `tfsdk:"resource_key"`
	// The type of the deployment resource.
	ResourceType types.String `tfsdk:"resource_type"`
	// Serialized local config state (what the CLI deployed).
	State jsontypes.Normalized `tfsdk:"state"`
	// When the last operation that updated this resource's recorded state was
	// applied. Pairs with last_action_type and last_version_id (all three
	// advance together on that write).
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Resource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Resource_SdkV2) {
}

func (to *Resource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Resource_SdkV2) {
}

func (m Resource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_action_type"] = attrs["last_action_type"].SetComputed()
	attrs["last_version_id"] = attrs["last_version_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["resource_key"] = attrs["resource_key"].SetOptional()
	attrs["resource_key"] = attrs["resource_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["resource_type"] = attrs["resource_type"].SetRequired()
	attrs["resource_type"] = attrs["resource_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Resource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Resource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Resource_SdkV2
// only implements ToObjectValue() and Type().
func (m Resource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_action_type": m.LastActionType,
			"last_version_id":  m.LastVersionId,
			"name":             m.Name,
			"resource_id":      m.ResourceId,
			"resource_key":     m.ResourceKey,
			"resource_type":    m.ResourceType,
			"state":            m.State,
			"update_time":      m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Resource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_action_type": types.StringType,
			"last_version_id":  types.StringType,
			"name":             types.StringType,
			"resource_id":      types.StringType,
			"resource_key":     types.StringType,
			"resource_type":    types.StringType,
			"state":            jsontypes.NormalizedType{},
			"update_time":      timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// A single invocation of a deploy or destroy command against a deployment.
// Creating a version acquires an exclusive lock on the parent deployment.
type Version_SdkV2 struct {
	// CLI version used to initiate the version.
	CliVersion types.String `tfsdk:"cli_version"`
	// When the version completed. Unset while the version is in progress.
	CompleteTime timetypes.RFC3339 `tfsdk:"complete_time"`
	// The user who completed the version (email or principal name). May differ
	// from `created_by` when another user force-completes the version.
	CompletedBy types.String `tfsdk:"completed_by"`
	// Why the version was completed. Unset while in progress. Set when status
	// transitions to COMPLETED.
	CompletionReason types.String `tfsdk:"completion_reason"`
	// When the version was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The user who created the version (email or principal name).
	CreatedBy types.String `tfsdk:"created_by"`
	// Bundle target deployment mode (development or production), captured at
	// the time of this version.
	DeploymentMode types.String `tfsdk:"deployment_mode"`
	// Display name for the deployment, captured at the time of this version.
	DisplayName types.String `tfsdk:"display_name"`
	// Git provenance of the source, captured at the time of this version.
	GitInfo types.List `tfsdk:"git_info"`
	// Resource name of the version. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name types.String `tfsdk:"name"`
	// Status of the version: IN_PROGRESS or COMPLETED.
	Status types.String `tfsdk:"status"`
	// Target name of the deployment, captured at the time of this version.
	TargetName types.String `tfsdk:"target_name"`
	// Version identifier within the parent deployment, assigned by the client
	// on creation. A numeric string (base-10, fits in a signed 64-bit integer)
	// that is greater than or equal to 1. Version IDs are strictly increasing
	// within a deployment but are not required to start at 1 or to be
	// contiguous.
	VersionId types.String `tfsdk:"version_id"`
	// Type of version (deploy or destroy).
	VersionType types.String `tfsdk:"version_type"`
	// Workspace location of the deployment, captured at the time of this
	// version.
	WorkspaceInfo types.List `tfsdk:"workspace_info"`
}

func (to *Version_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Version_SdkV2) {
	if !from.GitInfo.IsNull() && !from.GitInfo.IsUnknown() {
		if toGitInfo, ok := to.GetGitInfo(ctx); ok {
			if fromGitInfo, ok := from.GetGitInfo(ctx); ok {
				// Recursively sync the fields of GitInfo
				toGitInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGitInfo)
				to.SetGitInfo(ctx, toGitInfo)
			}
		}
	}
	if !from.WorkspaceInfo.IsNull() && !from.WorkspaceInfo.IsUnknown() {
		if toWorkspaceInfo, ok := to.GetWorkspaceInfo(ctx); ok {
			if fromWorkspaceInfo, ok := from.GetWorkspaceInfo(ctx); ok {
				// Recursively sync the fields of WorkspaceInfo
				toWorkspaceInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceInfo)
				to.SetWorkspaceInfo(ctx, toWorkspaceInfo)
			}
		}
	}
}

func (to *Version_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Version_SdkV2) {
	if !from.GitInfo.IsNull() && !from.GitInfo.IsUnknown() {
		if toGitInfo, ok := to.GetGitInfo(ctx); ok {
			if fromGitInfo, ok := from.GetGitInfo(ctx); ok {
				toGitInfo.SyncFieldsDuringRead(ctx, fromGitInfo)
				to.SetGitInfo(ctx, toGitInfo)
			}
		}
	}
	if !from.WorkspaceInfo.IsNull() && !from.WorkspaceInfo.IsUnknown() {
		if toWorkspaceInfo, ok := to.GetWorkspaceInfo(ctx); ok {
			if fromWorkspaceInfo, ok := from.GetWorkspaceInfo(ctx); ok {
				toWorkspaceInfo.SyncFieldsDuringRead(ctx, fromWorkspaceInfo)
				to.SetWorkspaceInfo(ctx, toWorkspaceInfo)
			}
		}
	}
}

func (m Version_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cli_version"] = attrs["cli_version"].SetRequired()
	attrs["cli_version"] = attrs["cli_version"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["complete_time"] = attrs["complete_time"].SetComputed()
	attrs["completed_by"] = attrs["completed_by"].SetComputed()
	attrs["completion_reason"] = attrs["completion_reason"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["deployment_mode"] = attrs["deployment_mode"].SetOptional()
	attrs["deployment_mode"] = attrs["deployment_mode"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["git_info"] = attrs["git_info"].SetOptional()
	attrs["git_info"] = attrs["git_info"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["git_info"] = attrs["git_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["target_name"] = attrs["target_name"].SetOptional()
	attrs["target_name"] = attrs["target_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["version_id"] = attrs["version_id"].SetComputed()
	attrs["version_type"] = attrs["version_type"].SetRequired()
	attrs["version_type"] = attrs["version_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["workspace_info"] = attrs["workspace_info"].SetOptional()
	attrs["workspace_info"] = attrs["workspace_info"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["workspace_info"] = attrs["workspace_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Version.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Version_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_info":       reflect.TypeOf(GitInfo_SdkV2{}),
		"workspace_info": reflect.TypeOf(WorkspaceInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Version_SdkV2
// only implements ToObjectValue() and Type().
func (m Version_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cli_version":       m.CliVersion,
			"complete_time":     m.CompleteTime,
			"completed_by":      m.CompletedBy,
			"completion_reason": m.CompletionReason,
			"create_time":       m.CreateTime,
			"created_by":        m.CreatedBy,
			"deployment_mode":   m.DeploymentMode,
			"display_name":      m.DisplayName,
			"git_info":          m.GitInfo,
			"name":              m.Name,
			"status":            m.Status,
			"target_name":       m.TargetName,
			"version_id":        m.VersionId,
			"version_type":      m.VersionType,
			"workspace_info":    m.WorkspaceInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Version_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cli_version":       types.StringType,
			"complete_time":     timetypes.RFC3339{}.Type(ctx),
			"completed_by":      types.StringType,
			"completion_reason": types.StringType,
			"create_time":       timetypes.RFC3339{}.Type(ctx),
			"created_by":        types.StringType,
			"deployment_mode":   types.StringType,
			"display_name":      types.StringType,
			"git_info": basetypes.ListType{
				ElemType: GitInfo_SdkV2{}.Type(ctx),
			},
			"name":         types.StringType,
			"status":       types.StringType,
			"target_name":  types.StringType,
			"version_id":   types.StringType,
			"version_type": types.StringType,
			"workspace_info": basetypes.ListType{
				ElemType: WorkspaceInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGitInfo returns the value of the GitInfo field in Version_SdkV2 as
// a GitInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Version_SdkV2) GetGitInfo(ctx context.Context) (GitInfo_SdkV2, bool) {
	var e GitInfo_SdkV2
	if m.GitInfo.IsNull() || m.GitInfo.IsUnknown() {
		return e, false
	}
	var v []GitInfo_SdkV2
	d := m.GitInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitInfo sets the value of the GitInfo field in Version_SdkV2.
func (m *Version_SdkV2) SetGitInfo(ctx context.Context, v GitInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["git_info"]
	m.GitInfo = types.ListValueMust(t, vs)
}

// GetWorkspaceInfo returns the value of the WorkspaceInfo field in Version_SdkV2 as
// a WorkspaceInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Version_SdkV2) GetWorkspaceInfo(ctx context.Context) (WorkspaceInfo_SdkV2, bool) {
	var e WorkspaceInfo_SdkV2
	if m.WorkspaceInfo.IsNull() || m.WorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v []WorkspaceInfo_SdkV2
	d := m.WorkspaceInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceInfo sets the value of the WorkspaceInfo field in Version_SdkV2.
func (m *Version_SdkV2) SetWorkspaceInfo(ctx context.Context, v WorkspaceInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_info"]
	m.WorkspaceInfo = types.ListValueMust(t, vs)
}

// Workspace location of a bundle deployment, captured at deploy time.
type WorkspaceInfo_SdkV2 struct {
	// Path of the bundle root (the directory containing databricks.yml)
	// relative to git_folder_path. Empty when the deployment is not from a
	// Databricks Git folder.
	BundleRootPath types.String `tfsdk:"bundle_root_path"`
	// Absolute workspace path where the deployed bundle files live. Mirrors the
	// workspace.file_path field in DABs bundle config.
	FilePath types.String `tfsdk:"file_path"`
	// When deployed from a Databricks Git folder, the absolute workspace path
	// of that folder; empty for local deploys.
	GitFolderPath types.String `tfsdk:"git_folder_path"`
	// Absolute workspace path of the deployment root — the base path the
	// deployed files live under. Mirrors workspace.root_path in the DABs bundle
	// config; file_path is its files subdirectory.
	RootPath types.String `tfsdk:"root_path"`
	// Whether files are served directly from the source sync root instead of
	// being copied into file_path.
	SourceLinked types.Bool `tfsdk:"source_linked"`
}

func (to *WorkspaceInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceInfo_SdkV2) {
}

func (to *WorkspaceInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceInfo_SdkV2) {
}

func (m WorkspaceInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bundle_root_path"] = attrs["bundle_root_path"].SetOptional()
	attrs["file_path"] = attrs["file_path"].SetOptional()
	attrs["git_folder_path"] = attrs["git_folder_path"].SetOptional()
	attrs["root_path"] = attrs["root_path"].SetOptional()
	attrs["source_linked"] = attrs["source_linked"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bundle_root_path": m.BundleRootPath,
			"file_path":        m.FilePath,
			"git_folder_path":  m.GitFolderPath,
			"root_path":        m.RootPath,
			"source_linked":    m.SourceLinked,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bundle_root_path": types.StringType,
			"file_path":        types.StringType,
			"git_folder_path":  types.StringType,
			"root_path":        types.StringType,
			"source_linked":    types.BoolType,
		},
	}
}
