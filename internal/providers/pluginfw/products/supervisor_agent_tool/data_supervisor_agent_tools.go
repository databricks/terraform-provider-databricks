// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisor_agent_tool

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/supervisoragents"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourcesName = "supervisor_agent_tools"

var _ datasource.DataSourceWithConfigure = &ToolsDataSource{}

func DataSourceTools() datasource.DataSource {
	return &ToolsDataSource{}
}

// ToolsData extends the main model with additional fields.
type ToolsData struct {
	SupervisorAgents types.List `tfsdk:"tools"`

	PageSize types.Int64 `tfsdk:"page_size"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ToolsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tools":           reflect.TypeOf(ToolData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ToolsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["tools"] = attrs["tools"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ToolsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ToolsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ToolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ToolsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Tool",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ToolsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ToolsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ToolsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest supervisoragents.ListToolsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
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

	response, err := client.SupervisorAgents.ListToolsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list supervisor_agent_tools", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var tool ToolData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &tool)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tool.ProviderConfigData = config.ProviderConfigData

		results = append(results, tool.ToObjectValue(ctx))
	}

	config.SupervisorAgents = types.ListValueMust(ToolData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
