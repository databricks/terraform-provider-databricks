// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledge_assistant

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
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

const resourceName = "knowledge_assistant"

var _ resource.ResourceWithConfigure = &KnowledgeAssistantResource{}
var _ resource.ResourceWithModifyPlan = &KnowledgeAssistantResource{}

func ResourceKnowledgeAssistant() resource.Resource {
	return &KnowledgeAssistantResource{}
}

type KnowledgeAssistantResource struct {
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

// KnowledgeAssistant extends the main model with additional fields.
type KnowledgeAssistant struct {
	// Creation timestamp.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The creator of the Knowledge Assistant.
	Creator types.String `tfsdk:"creator"`
	// Description of what this agent can do (user-facing). Required when
	// creating a Knowledge Assistant. When updating a Knowledge Assistant,
	// optional unless included in update_mask.
	Description types.String `tfsdk:"description"`
	// The display name of the Knowledge Assistant, unique at workspace level.
	// Required when creating a Knowledge Assistant. When updating a Knowledge
	// Assistant, optional unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`
	// The name of the knowledge assistant agent endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Error details when the Knowledge Assistant is in FAILED state.
	ErrorInfo types.String `tfsdk:"error_info"`
	// The MLflow experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The universally unique identifier (UUID) of the Knowledge Assistant.
	Id types.String `tfsdk:"id"`
	// Additional global instructions on how the agent should generate answers.
	// Optional on create and update. When updating a Knowledge Assistant,
	// include this field in update_mask to modify it.
	Instructions types.String `tfsdk:"instructions"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"name"`
	// State of the Knowledge Assistant. Not returned in List responses.
	State          types.String `tfsdk:"state"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// KnowledgeAssistant struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m KnowledgeAssistant) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistant
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistant) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"creator":       m.Creator,
			"description":   m.Description,
			"display_name":  m.DisplayName,
			"endpoint_name": m.EndpointName,
			"error_info":    m.ErrorInfo,
			"experiment_id": m.ExperimentId,
			"id":            m.Id,
			"instructions":  m.Instructions,
			"name":          m.Name,
			"state":         m.State,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m KnowledgeAssistant) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"creator":       types.StringType,
			"description":   types.StringType,
			"display_name":  types.StringType,
			"endpoint_name": types.StringType,
			"error_info":    types.StringType,
			"experiment_id": types.StringType,
			"id":            types.StringType,
			"instructions":  types.StringType,
			"name":          types.StringType,
			"state":         types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *KnowledgeAssistant) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistant) {
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *KnowledgeAssistant) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistant) {
	to.ProviderConfig = from.ProviderConfig

}

func (m KnowledgeAssistant) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["error_info"] = attrs["error_info"].SetComputed()
	attrs["experiment_id"] = attrs["experiment_id"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

func (r *KnowledgeAssistantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *KnowledgeAssistantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, KnowledgeAssistant{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks knowledge_assistant",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KnowledgeAssistantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *KnowledgeAssistantResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *KnowledgeAssistantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KnowledgeAssistant
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var knowledge_assistant knowledgeassistants.KnowledgeAssistant

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &knowledge_assistant)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := knowledgeassistants.CreateKnowledgeAssistantRequest{
		KnowledgeAssistant: knowledge_assistant,
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

	response, err := client.KnowledgeAssistants.CreateKnowledgeAssistant(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create knowledge_assistant", err.Error())
		return
	}

	var newState KnowledgeAssistant

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

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

func (r *KnowledgeAssistantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState KnowledgeAssistant
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest knowledgeassistants.GetKnowledgeAssistantRequest
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
	response, err := client.KnowledgeAssistants.GetKnowledgeAssistant(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get knowledge_assistant", err.Error())
		return
	}

	var newState KnowledgeAssistant
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

func (r *KnowledgeAssistantResource) update(ctx context.Context, plan KnowledgeAssistant, diags *diag.Diagnostics, state *tfsdk.State) {
	var knowledge_assistant knowledgeassistants.KnowledgeAssistant

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &knowledge_assistant)...)
	if diags.HasError() {
		return
	}

	updateRequest := knowledgeassistants.UpdateKnowledgeAssistantRequest{
		KnowledgeAssistant: knowledge_assistant,
		Name:               plan.Name.ValueString(),
		UpdateMask:         *fieldmask.New(strings.Split("description,display_name,instructions", ",")),
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
	response, err := client.KnowledgeAssistants.UpdateKnowledgeAssistant(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update knowledge_assistant", err.Error())
		return
	}

	var newState KnowledgeAssistant

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *KnowledgeAssistantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KnowledgeAssistant
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *KnowledgeAssistantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state KnowledgeAssistant
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest knowledgeassistants.DeleteKnowledgeAssistantRequest
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

	err := client.KnowledgeAssistants.DeleteKnowledgeAssistant(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete knowledge_assistant", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &KnowledgeAssistantResource{}

func (r *KnowledgeAssistantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
