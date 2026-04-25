// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisor_agent

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

const dataSourcesName = "supervisor_agents"

var _ datasource.DataSourceWithConfigure = &SupervisorAgentsDataSource{}

func DataSourceSupervisorAgents() datasource.DataSource {
	return &SupervisorAgentsDataSource{}
}

// SupervisorAgentsData extends the main model with additional fields.
type SupervisorAgentsData struct {
	SupervisorAgents types.List `tfsdk:"supervisor_agents"`
	// The maximum number of supervisor agents to return. If unspecified, at
	// most 100 supervisor agents will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SupervisorAgentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agents": reflect.TypeOf(SupervisorAgentData{}),
		"provider_config":   reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SupervisorAgentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["supervisor_agents"] = attrs["supervisor_agents"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SupervisorAgentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SupervisorAgentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SupervisorAgentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SupervisorAgentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks SupervisorAgent",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SupervisorAgentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SupervisorAgentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SupervisorAgentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest supervisoragents.ListSupervisorAgentsRequest
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

	response, err := client.SupervisorAgents.ListSupervisorAgentsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list supervisor_agents", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var supervisor_agent SupervisorAgentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &supervisor_agent)...)
		if resp.Diagnostics.HasError() {
			return
		}
		supervisor_agent.ProviderConfigData = config.ProviderConfigData

		results = append(results, supervisor_agent.ToObjectValue(ctx))
	}

	config.SupervisorAgents = types.ListValueMust(SupervisorAgentData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
