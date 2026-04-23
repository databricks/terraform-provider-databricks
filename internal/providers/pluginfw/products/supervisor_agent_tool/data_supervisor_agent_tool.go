// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisor_agent_tool

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/supervisoragents"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/supervisoragents_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "supervisor_agent_tool"

var _ datasource.DataSourceWithConfigure = &ToolDataSource{}

func DataSourceTool() datasource.DataSource {
	return &ToolDataSource{}
}

type ToolDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// ToolData extends the main model with additional fields.
type ToolData struct {
	App types.Object `tfsdk:"app"`

	UcConnection types.Object `tfsdk:"uc_connection"`
	// Description of what this tool does (user-facing).
	Description types.String `tfsdk:"description"`

	GenieSpace types.Object `tfsdk:"genie_space"`
	// Deprecated: Use tool_id instead.
	Id types.String `tfsdk:"id"`

	KnowledgeAssistant types.Object `tfsdk:"knowledge_assistant"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"name"`
	// User specified id of the Tool.
	ToolId types.String `tfsdk:"tool_id"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "connection", "app", "volume", "lakeview_dashboard",
	// "serving_endpoint", "uc_table", "vector_search_index".
	ToolType types.String `tfsdk:"tool_type"`

	UcFunction types.Object `tfsdk:"uc_function"`

	Volume             types.Object `tfsdk:"volume"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ToolData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ToolData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":                 reflect.TypeOf(supervisoragents_tf.App{}),
		"uc_connection":       reflect.TypeOf(supervisoragents_tf.Connection{}),
		"genie_space":         reflect.TypeOf(supervisoragents_tf.GenieSpace{}),
		"knowledge_assistant": reflect.TypeOf(supervisoragents_tf.KnowledgeAssistant{}),
		"uc_function":         reflect.TypeOf(supervisoragents_tf.UcFunction{}),
		"volume":              reflect.TypeOf(supervisoragents_tf.Volume{}),
		"provider_config":     reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ToolData
// only implements ToObjectValue() and Type().
func (m ToolData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":                 m.App,
			"uc_connection":       m.UcConnection,
			"description":         m.Description,
			"genie_space":         m.GenieSpace,
			"id":                  m.Id,
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"tool_id":             m.ToolId,
			"tool_type":           m.ToolType,
			"uc_function":         m.UcFunction,
			"volume":              m.Volume,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ToolData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":                 supervisoragents_tf.App{}.Type(ctx),
			"uc_connection":       supervisoragents_tf.Connection{}.Type(ctx),
			"description":         types.StringType,
			"genie_space":         supervisoragents_tf.GenieSpace{}.Type(ctx),
			"id":                  types.StringType,
			"knowledge_assistant": supervisoragents_tf.KnowledgeAssistant{}.Type(ctx),
			"name":                types.StringType,
			"tool_id":             types.StringType,
			"tool_type":           types.StringType,
			"uc_function":         supervisoragents_tf.UcFunction{}.Type(ctx),
			"volume":              supervisoragents_tf.Volume{}.Type(ctx),

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m ToolData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetComputed()
	attrs["uc_connection"] = attrs["uc_connection"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["genie_space"] = attrs["genie_space"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["tool_id"] = attrs["tool_id"].SetComputed()
	attrs["tool_type"] = attrs["tool_type"].SetComputed()
	attrs["uc_function"] = attrs["uc_function"].SetComputed()
	attrs["volume"] = attrs["volume"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *ToolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ToolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ToolData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Tool",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ToolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ToolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ToolData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest supervisoragents.GetToolRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
		resp.Diagnostics.AddError("failed to get supervisor_agent_tool", err.Error())
		return
	}

	var newState ToolData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
