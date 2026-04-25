// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments_workspace_base_environment

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/environments"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "environments_workspace_base_environment"

var _ resource.ResourceWithConfigure = &WorkspaceBaseEnvironmentResource{}
var _ resource.ResourceWithModifyPlan = &WorkspaceBaseEnvironmentResource{}

func ResourceWorkspaceBaseEnvironment() resource.Resource {
	return &WorkspaceBaseEnvironmentResource{}
}

type WorkspaceBaseEnvironmentResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// WorkspaceBaseEnvironment extends the main model with additional fields.
type WorkspaceBaseEnvironment struct {
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
	// The ID to use for the workspace base environment, which will become the
	// final component of the resource name. This value should be 4-63
	// characters, and valid characters are /[a-z][0-9]-/.
	WorkspaceBaseEnvironmentId types.String `tfsdk:"workspace_base_environment_id"`
	ProviderConfig             types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceBaseEnvironment struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceBaseEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBaseEnvironment
// only implements ToObjectValue() and Type().
func (m WorkspaceBaseEnvironment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"base_environment_type": m.BaseEnvironmentType, "effective_base_environment_type": m.EffectiveBaseEnvironmentType,
			"create_time":                   m.CreateTime,
			"creator_user_id":               m.CreatorUserId,
			"display_name":                  m.DisplayName,
			"filepath":                      m.Filepath,
			"is_default":                    m.IsDefault,
			"last_updated_user_id":          m.LastUpdatedUserId,
			"message":                       m.Message,
			"name":                          m.Name,
			"status":                        m.Status,
			"update_time":                   m.UpdateTime,
			"workspace_base_environment_id": m.WorkspaceBaseEnvironmentId,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceBaseEnvironment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"base_environment_type": types.StringType,
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
			"workspace_base_environment_id":   types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *WorkspaceBaseEnvironment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceBaseEnvironment) {
	to.EffectiveBaseEnvironmentType = to.BaseEnvironmentType
	to.BaseEnvironmentType = from.BaseEnvironmentType
	if !from.WorkspaceBaseEnvironmentId.IsUnknown() {
		to.WorkspaceBaseEnvironmentId = from.WorkspaceBaseEnvironmentId
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *WorkspaceBaseEnvironment) SyncFieldsDuringRead(ctx context.Context, from WorkspaceBaseEnvironment) {
	to.EffectiveBaseEnvironmentType = from.EffectiveBaseEnvironmentType
	if from.EffectiveBaseEnvironmentType.ValueString() == to.BaseEnvironmentType.ValueString() {
		to.BaseEnvironmentType = from.BaseEnvironmentType
	}
	if !from.WorkspaceBaseEnvironmentId.IsUnknown() {
		to.WorkspaceBaseEnvironmentId = from.WorkspaceBaseEnvironmentId
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m WorkspaceBaseEnvironment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["base_environment_type"] = attrs["base_environment_type"].SetOptional()
	attrs["effective_base_environment_type"] = attrs["effective_base_environment_type"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["filepath"] = attrs["filepath"].SetOptional()
	attrs["is_default"] = attrs["is_default"].SetComputed()
	attrs["last_updated_user_id"] = attrs["last_updated_user_id"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_base_environment_id"] = attrs["workspace_base_environment_id"].SetComputed()
	attrs["workspace_base_environment_id"] = attrs["workspace_base_environment_id"].SetOptional()
	attrs["workspace_base_environment_id"] = attrs["workspace_base_environment_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["workspace_base_environment_id"] = attrs["workspace_base_environment_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

func (r *WorkspaceBaseEnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *WorkspaceBaseEnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, WorkspaceBaseEnvironment{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks environments_workspace_base_environment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceBaseEnvironmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *WorkspaceBaseEnvironmentResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *WorkspaceBaseEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceBaseEnvironment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var workspace_base_environment environments.WorkspaceBaseEnvironment

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_base_environment)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := environments.CreateWorkspaceBaseEnvironmentRequest{
		WorkspaceBaseEnvironment:   workspace_base_environment,
		WorkspaceBaseEnvironmentId: plan.WorkspaceBaseEnvironmentId.ValueString(),
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Environments.CreateWorkspaceBaseEnvironment(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create environments_workspace_base_environment", err.Error())
		return
	}

	var newState WorkspaceBaseEnvironment

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for environments_workspace_base_environment to be ready", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *WorkspaceBaseEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState WorkspaceBaseEnvironment
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest environments.GetWorkspaceBaseEnvironmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Environments.GetWorkspaceBaseEnvironment(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get environments_workspace_base_environment", err.Error())
		return
	}

	var newState WorkspaceBaseEnvironment
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *WorkspaceBaseEnvironmentResource) update(ctx context.Context, plan WorkspaceBaseEnvironment, diags *diag.Diagnostics, state *tfsdk.State) {
	var workspace_base_environment environments.WorkspaceBaseEnvironment

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_base_environment)...)
	if diags.HasError() {
		return
	}

	updateRequest := environments.UpdateWorkspaceBaseEnvironmentRequest{
		WorkspaceBaseEnvironment: workspace_base_environment,
		Name:                     plan.Name.ValueString(),
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Environments.UpdateWorkspaceBaseEnvironment(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update environments_workspace_base_environment", err.Error())
		return
	}

	var newState WorkspaceBaseEnvironment

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for environments_workspace_base_environment update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *WorkspaceBaseEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceBaseEnvironment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *WorkspaceBaseEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state WorkspaceBaseEnvironment
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest environments.DeleteWorkspaceBaseEnvironmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Environments.DeleteWorkspaceBaseEnvironment(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete environments_workspace_base_environment", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &WorkspaceBaseEnvironmentResource{}

func (r *WorkspaceBaseEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
