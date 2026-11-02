// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genie_agent

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/genieagents"
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

const dataSourcesName = "genie_agents"

var _ datasource.DataSourceWithConfigure = &GenieAgentsDataSource{}

func DataSourceGenieAgents() datasource.DataSource {
	return &GenieAgentsDataSource{}
}

// GenieAgentsData extends the main model with additional fields.
type GenieAgentsData struct {
	GenieAgents types.List `tfsdk:"genie_agents"`
	// Maximum number of Genie Agents to return per page. Defaults to 20 when
	// unset or 0; values above 100 are coerced to 100. Negative values are
	// rejected with INVALID_PARAMETER_VALUE. With `GENIE_AGENT_VIEW_FULL`, a
	// smaller per-page maximum applies.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The fields to return for each Genie Agent. Defaults to
	// `GENIE_AGENT_VIEW_BASIC`. Use `GENIE_AGENT_VIEW_FULL` only when the agent
	// bodies are needed, since it returns smaller pages.
	View               types.String `tfsdk:"view"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (GenieAgentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"genie_agents":    reflect.TypeOf(GenieAgentData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m GenieAgentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	attrs["genie_agents"] = attrs["genie_agents"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type GenieAgentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *GenieAgentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *GenieAgentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GenieAgentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks GenieAgent",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *GenieAgentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *GenieAgentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config GenieAgentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest genieagents.ListGenieAgentsRequest
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

	response, err := client.GenieAgents.ListGenieAgentsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list genie_agents", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var genie_agent GenieAgentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &genie_agent)...)
		if resp.Diagnostics.HasError() {
			return
		}
		genie_agent.ProviderConfigData = config.ProviderConfigData

		results = append(results, genie_agent.ToObjectValue(ctx))
	}

	config.GenieAgents = types.ListValueMust(GenieAgentData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
