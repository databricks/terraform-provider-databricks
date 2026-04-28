// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisor_agent_tool

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/supervisoragents"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/supervisoragents_tf"
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

const resourceName = "supervisor_agent_tool"

var _ resource.ResourceWithConfigure = &ToolResource{}
var _ resource.ResourceWithModifyPlan = &ToolResource{}

func ResourceTool() resource.Resource {
	return &ToolResource{}
}

type ToolResource struct {
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

// Tool extends the main model with additional fields.
type Tool struct {
	App types.Object `tfsdk:"app"`
	// Description of what this tool does (user-facing).
	Description types.String `tfsdk:"description"`

	GenieSpace types.Object `tfsdk:"genie_space"`
	// Deprecated: Use tool_id instead.
	Id types.String `tfsdk:"id"`

	KnowledgeAssistant types.Object `tfsdk:"knowledge_assistant"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"name"`
	// Parent resource where this tool will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"parent"`
	// User specified id of the Tool.
	ToolId types.String `tfsdk:"tool_id"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "uc_connection", "app", "volume", "lakeview_dashboard",
	// "serving_endpoint", "uc_table", "vector_search_index".
	ToolType types.String `tfsdk:"tool_type"`

	UcConnection types.Object `tfsdk:"uc_connection"`

	UcFunction types.Object `tfsdk:"uc_function"`

	Volume         types.Object `tfsdk:"volume"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Tool struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Tool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":                 reflect.TypeOf(supervisoragents_tf.App{}),
		"genie_space":         reflect.TypeOf(supervisoragents_tf.GenieSpace{}),
		"knowledge_assistant": reflect.TypeOf(supervisoragents_tf.KnowledgeAssistant{}),
		"uc_connection":       reflect.TypeOf(supervisoragents_tf.UcConnection{}),
		"uc_function":         reflect.TypeOf(supervisoragents_tf.UcFunction{}),
		"volume":              reflect.TypeOf(supervisoragents_tf.Volume{}),
		"provider_config":     reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Tool
// only implements ToObjectValue() and Type().
func (m Tool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"app": m.App,
			"description":         m.Description,
			"genie_space":         m.GenieSpace,
			"id":                  m.Id,
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"parent":              m.Parent,
			"tool_id":             m.ToolId,
			"tool_type":           m.ToolType,
			"uc_connection":       m.UcConnection,
			"uc_function":         m.UcFunction,
			"volume":              m.Volume,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Tool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"app": supervisoragents_tf.App{}.Type(ctx),
			"description":         types.StringType,
			"genie_space":         supervisoragents_tf.GenieSpace{}.Type(ctx),
			"id":                  types.StringType,
			"knowledge_assistant": supervisoragents_tf.KnowledgeAssistant{}.Type(ctx),
			"name":                types.StringType,
			"parent":              types.StringType,
			"tool_id":             types.StringType,
			"tool_type":           types.StringType,
			"uc_connection":       supervisoragents_tf.UcConnection{}.Type(ctx),
			"uc_function":         supervisoragents_tf.UcFunction{}.Type(ctx),
			"volume":              supervisoragents_tf.Volume{}.Type(ctx),

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Tool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Tool) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				// Recursively sync the fields of GenieSpace
				toGenieSpace.SyncFieldsDuringCreateOrUpdate(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.UcConnection.IsNull() && !from.UcConnection.IsUnknown() {
		if toUcConnection, ok := to.GetUcConnection(ctx); ok {
			if fromUcConnection, ok := from.GetUcConnection(ctx); ok {
				// Recursively sync the fields of UcConnection
				toUcConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromUcConnection)
				to.SetUcConnection(ctx, toUcConnection)
			}
		}
	}
	if !from.UcFunction.IsNull() && !from.UcFunction.IsUnknown() {
		if toUcFunction, ok := to.GetUcFunction(ctx); ok {
			if fromUcFunction, ok := from.GetUcFunction(ctx); ok {
				// Recursively sync the fields of UcFunction
				toUcFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromUcFunction)
				to.SetUcFunction(ctx, toUcFunction)
			}
		}
	}
	if !from.Volume.IsNull() && !from.Volume.IsUnknown() {
		if toVolume, ok := to.GetVolume(ctx); ok {
			if fromVolume, ok := from.GetVolume(ctx); ok {
				// Recursively sync the fields of Volume
				toVolume.SyncFieldsDuringCreateOrUpdate(ctx, fromVolume)
				to.SetVolume(ctx, toVolume)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Tool) SyncFieldsDuringRead(ctx context.Context, from Tool) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				toGenieSpace.SyncFieldsDuringRead(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.UcConnection.IsNull() && !from.UcConnection.IsUnknown() {
		if toUcConnection, ok := to.GetUcConnection(ctx); ok {
			if fromUcConnection, ok := from.GetUcConnection(ctx); ok {
				toUcConnection.SyncFieldsDuringRead(ctx, fromUcConnection)
				to.SetUcConnection(ctx, toUcConnection)
			}
		}
	}
	if !from.UcFunction.IsNull() && !from.UcFunction.IsUnknown() {
		if toUcFunction, ok := to.GetUcFunction(ctx); ok {
			if fromUcFunction, ok := from.GetUcFunction(ctx); ok {
				toUcFunction.SyncFieldsDuringRead(ctx, fromUcFunction)
				to.SetUcFunction(ctx, toUcFunction)
			}
		}
	}
	if !from.Volume.IsNull() && !from.Volume.IsUnknown() {
		if toVolume, ok := to.GetVolume(ctx); ok {
			if fromVolume, ok := from.GetVolume(ctx); ok {
				toVolume.SyncFieldsDuringRead(ctx, fromVolume)
				to.SetVolume(ctx, toVolume)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Tool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetOptional()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["genie_space"] = attrs["genie_space"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["tool_id"] = attrs["tool_id"].SetRequired()
	attrs["tool_id"] = attrs["tool_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["tool_id"] = attrs["tool_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["tool_type"] = attrs["tool_type"].SetRequired()
	attrs["uc_connection"] = attrs["uc_connection"].SetOptional()
	attrs["uc_function"] = attrs["uc_function"].SetOptional()
	attrs["volume"] = attrs["volume"].SetOptional()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetApp returns the value of the App field in Tool as
// a supervisoragents_tf.App value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetApp(ctx context.Context) (supervisoragents_tf.App, bool) {
	var e supervisoragents_tf.App
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.App
	d := m.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApp sets the value of the App field in Tool.
func (m *Tool) SetApp(ctx context.Context, v supervisoragents_tf.App) {
	vs := v.ToObjectValue(ctx)
	m.App = vs
}

// GetGenieSpace returns the value of the GenieSpace field in Tool as
// a supervisoragents_tf.GenieSpace value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetGenieSpace(ctx context.Context) (supervisoragents_tf.GenieSpace, bool) {
	var e supervisoragents_tf.GenieSpace
	if m.GenieSpace.IsNull() || m.GenieSpace.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.GenieSpace
	d := m.GenieSpace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieSpace sets the value of the GenieSpace field in Tool.
func (m *Tool) SetGenieSpace(ctx context.Context, v supervisoragents_tf.GenieSpace) {
	vs := v.ToObjectValue(ctx)
	m.GenieSpace = vs
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in Tool as
// a supervisoragents_tf.KnowledgeAssistant value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetKnowledgeAssistant(ctx context.Context) (supervisoragents_tf.KnowledgeAssistant, bool) {
	var e supervisoragents_tf.KnowledgeAssistant
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.KnowledgeAssistant
	d := m.KnowledgeAssistant.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in Tool.
func (m *Tool) SetKnowledgeAssistant(ctx context.Context, v supervisoragents_tf.KnowledgeAssistant) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeAssistant = vs
}

// GetUcConnection returns the value of the UcConnection field in Tool as
// a supervisoragents_tf.UcConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetUcConnection(ctx context.Context) (supervisoragents_tf.UcConnection, bool) {
	var e supervisoragents_tf.UcConnection
	if m.UcConnection.IsNull() || m.UcConnection.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.UcConnection
	d := m.UcConnection.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcConnection sets the value of the UcConnection field in Tool.
func (m *Tool) SetUcConnection(ctx context.Context, v supervisoragents_tf.UcConnection) {
	vs := v.ToObjectValue(ctx)
	m.UcConnection = vs
}

// GetUcFunction returns the value of the UcFunction field in Tool as
// a supervisoragents_tf.UcFunction value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetUcFunction(ctx context.Context) (supervisoragents_tf.UcFunction, bool) {
	var e supervisoragents_tf.UcFunction
	if m.UcFunction.IsNull() || m.UcFunction.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.UcFunction
	d := m.UcFunction.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcFunction sets the value of the UcFunction field in Tool.
func (m *Tool) SetUcFunction(ctx context.Context, v supervisoragents_tf.UcFunction) {
	vs := v.ToObjectValue(ctx)
	m.UcFunction = vs
}

// GetVolume returns the value of the Volume field in Tool as
// a supervisoragents_tf.Volume value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetVolume(ctx context.Context) (supervisoragents_tf.Volume, bool) {
	var e supervisoragents_tf.Volume
	if m.Volume.IsNull() || m.Volume.IsUnknown() {
		return e, false
	}
	var v supervisoragents_tf.Volume
	d := m.Volume.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolume sets the value of the Volume field in Tool.
func (m *Tool) SetVolume(ctx context.Context, v supervisoragents_tf.Volume) {
	vs := v.ToObjectValue(ctx)
	m.Volume = vs
}

func (r *ToolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ToolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Tool{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks supervisor_agent_tool",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ToolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ToolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

func (r *ToolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Tool
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var tool supervisoragents.Tool

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &tool)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := supervisoragents.CreateToolRequest{
		Tool:   tool,
		Parent: plan.Parent.ValueString(),
		ToolId: plan.ToolId.ValueString(),
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

	response, err := client.SupervisorAgents.CreateTool(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create supervisor_agent_tool", err.Error())
		return
	}

	var newState Tool

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

func (r *ToolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Tool
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest supervisoragents.GetToolRequest
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
	response, err := client.SupervisorAgents.GetTool(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get supervisor_agent_tool", err.Error())
		return
	}

	var newState Tool
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

func (r *ToolResource) update(ctx context.Context, plan Tool, diags *diag.Diagnostics, state *tfsdk.State) {
	var tool supervisoragents.Tool

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &tool)...)
	if diags.HasError() {
		return
	}

	updateRequest := supervisoragents.UpdateToolRequest{
		Tool:       tool,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("app,description,genie_space,knowledge_assistant,tool_type,uc_connection,uc_function,volume", ",")),
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
	response, err := client.SupervisorAgents.UpdateTool(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update supervisor_agent_tool", err.Error())
		return
	}

	var newState Tool

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ToolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Tool
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ToolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Tool
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest supervisoragents.DeleteToolRequest
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

	err := client.SupervisorAgents.DeleteTool(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete supervisor_agent_tool", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ToolResource{}

func (r *ToolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
