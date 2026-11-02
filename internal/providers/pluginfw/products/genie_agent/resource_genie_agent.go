// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genie_agent

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/genieagents"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/genieagents_tf"
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

const resourceName = "genie_agent"

var _ resource.ResourceWithConfigure = &GenieAgentResource{}
var _ resource.ResourceWithModifyPlan = &GenieAgentResource{}

func ResourceGenieAgent() resource.Resource {
	return &GenieAgentResource{}
}

type GenieAgentResource struct {
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

// GenieAgent extends the main model with additional fields.
type GenieAgent struct {
	// Time when the Genie Agent was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The definition of the Genie Agent: its sample questions, data sources,
	// instructions, and benchmarks. Must set at least one field. Returned on
	// get and on list with `GENIE_AGENT_VIEW_FULL`. Create and update return a
	// submitted definition unchanged; one the service would store differently,
	// or with a column config the workspace can't keep (unknown column, or a
	// display name while those are disabled), is rejected with
	// INVALID_PARAMETER_VALUE.
	Definition types.Object `tfsdk:"definition"`
	// Description of the Genie Agent.
	//
	// The maximum length is 4,096 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	Description types.String `tfsdk:"description"`
	// Display name of the Genie Agent. Must be unique within its folder
	// (effective_parent_path): a create, rename, or move that would collide
	// with an existing Genie Agent's name in the target folder is rejected with
	// ALREADY_EXISTS.
	//
	// The maximum length is 63 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace folder that contains the Genie Agent.
	EffectiveParentPath types.String `tfsdk:"effective_parent_path"`
	// Weak etag of `definition` only; metadata changes do not affect it. Send
	// it on update or delete to detect concurrent definition changes.
	Etag types.String `tfsdk:"etag"`
	// The ID for the new Genie Agent: 32 lowercase hexadecimal characters (a
	// UUID without dashes), not all zeros. Assigned by the server if omitted.
	// Fails with ALREADY_EXISTS if the ID is in use in the workspace, including
	// by a Genie Agent in the trash, or PERMISSION_DENIED if the caller cannot
	// view the Genie Agent using it.
	GenieAgentId types.String `tfsdk:"genie_agent_id"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name types.String `tfsdk:"name"`
	// The workspace folder to create the Genie Agent in. If omitted, the
	// caller's home folder is used. Returned exactly as submitted; see
	// `effective_parent_path` for the current folder. On update, setting it
	// moves the Genie Agent and clearing it does not. At most 4,096 characters.
	ParentPath types.String `tfsdk:"parent_path"`
	// Time when the Genie Agent was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Warehouse associated with the Genie Agent.
	WarehouseId    types.String `tfsdk:"warehouse_id"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// GenieAgent struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m GenieAgent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"definition":      reflect.TypeOf(genieagents_tf.GenieAgentDefinition{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgent
// only implements ToObjectValue() and Type().
func (m GenieAgent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"definition":            m.Definition,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"effective_parent_path": m.EffectiveParentPath,
			"etag":                  m.Etag,
			"genie_agent_id":        m.GenieAgentId,
			"name":                  m.Name,
			"parent_path":           m.ParentPath,
			"update_time":           m.UpdateTime,
			"warehouse_id":          m.WarehouseId,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m GenieAgent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"definition":            genieagents_tf.GenieAgentDefinition{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"effective_parent_path": types.StringType,
			"etag":                  types.StringType,
			"genie_agent_id":        types.StringType,
			"name":                  types.StringType,
			"parent_path":           types.StringType,
			"update_time":           timetypes.RFC3339{}.Type(ctx),
			"warehouse_id":          types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *GenieAgent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgent) {
	if !from.Definition.IsNull() && !from.Definition.IsUnknown() {
		if toDefinition, ok := to.GetDefinition(ctx); ok {
			if fromDefinition, ok := from.GetDefinition(ctx); ok {
				// Recursively sync the fields of Definition
				toDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromDefinition)
				to.SetDefinition(ctx, toDefinition)
			}
		}
	}
	if !from.GenieAgentId.IsUnknown() {
		to.GenieAgentId = from.GenieAgentId
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *GenieAgent) SyncFieldsDuringRead(ctx context.Context, from GenieAgent) {
	if !from.Definition.IsNull() && !from.Definition.IsUnknown() {
		if toDefinition, ok := to.GetDefinition(ctx); ok {
			if fromDefinition, ok := from.GetDefinition(ctx); ok {
				toDefinition.SyncFieldsDuringRead(ctx, fromDefinition)
				to.SetDefinition(ctx, toDefinition)
			}
		}
	}
	if !from.GenieAgentId.IsUnknown() {
		to.GenieAgentId = from.GenieAgentId
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m GenieAgent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["definition"] = attrs["definition"].SetRequired()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["effective_parent_path"] = attrs["effective_parent_path"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()
	attrs["genie_agent_id"] = attrs["genie_agent_id"].SetComputed()
	attrs["genie_agent_id"] = attrs["genie_agent_id"].SetOptional()
	attrs["genie_agent_id"] = attrs["genie_agent_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["genie_agent_id"] = attrs["genie_agent_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetDefinition returns the value of the Definition field in GenieAgent as
// a genieagents_tf.GenieAgentDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgent) GetDefinition(ctx context.Context) (genieagents_tf.GenieAgentDefinition, bool) {
	var e genieagents_tf.GenieAgentDefinition
	if m.Definition.IsNull() || m.Definition.IsUnknown() {
		return e, false
	}
	var v genieagents_tf.GenieAgentDefinition
	d := m.Definition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefinition sets the value of the Definition field in GenieAgent.
func (m *GenieAgent) SetDefinition(ctx context.Context, v genieagents_tf.GenieAgentDefinition) {
	vs := v.ToObjectValue(ctx)
	m.Definition = vs
}

func (r *GenieAgentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *GenieAgentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, GenieAgent{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks genie_agent",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *GenieAgentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *GenieAgentResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *GenieAgentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan GenieAgent
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var genie_agent genieagents.GenieAgent

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &genie_agent)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := genieagents.CreateGenieAgentRequest{
		GenieAgent:   genie_agent,
		GenieAgentId: plan.GenieAgentId.ValueString(),
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

	response, err := client.GenieAgents.CreateGenieAgent(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create genie_agent", err.Error())
		return
	}

	var newState GenieAgent

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

func (r *GenieAgentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState GenieAgent
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest genieagents.GetGenieAgentRequest
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
	response, err := client.GenieAgents.GetGenieAgent(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get genie_agent", err.Error())
		return
	}

	var newState GenieAgent
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

func (r *GenieAgentResource) update(ctx context.Context, plan GenieAgent, diags *diag.Diagnostics, state *tfsdk.State) {
	var genie_agent genieagents.GenieAgent

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &genie_agent)...)
	if diags.HasError() {
		return
	}

	updateRequest := genieagents.UpdateGenieAgentRequest{
		GenieAgent: genie_agent,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("definition,description,display_name,parent_path,warehouse_id", ",")),
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
	response, err := client.GenieAgents.UpdateGenieAgent(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update genie_agent", err.Error())
		return
	}

	var newState GenieAgent

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *GenieAgentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan GenieAgent
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *GenieAgentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state GenieAgent
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest genieagents.DeleteGenieAgentRequest
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

	err := client.GenieAgents.DeleteGenieAgent(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete genie_agent", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &GenieAgentResource{}

func (r *GenieAgentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
