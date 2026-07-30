// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_gateway_mcp_service

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
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

const dataSourcesName = "ai_gateway_mcp_services"

var _ datasource.DataSourceWithConfigure = &McpServicesDataSource{}

func DataSourceMcpServices() datasource.DataSource {
	return &McpServicesDataSource{}
}

// McpServicesData extends the main model with additional fields.
type McpServicesData struct {
	AiGateway types.List `tfsdk:"mcp_services"`
	// Whether to include MCP services for which the principal can only access
	// selective metadata.
	IncludeBrowse types.Bool `tfsdk:"include_browse"`
	// Maximum number of MCP services to return. Defaults to 100 when unset or
	// 0; the maximum is 1000. Use `next_page_token` to retrieve additional
	// pages.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Resource name of the parent schema to list within, as
	// `schemas/{catalog}.{schema}`. Each `{...}` component is capped at 255
	// characters individually.
	Parent types.String `tfsdk:"parent"`
	// View selector controlling which fields are populated per row.
	View               types.String `tfsdk:"view"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (McpServicesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"mcp_services":    reflect.TypeOf(McpServiceData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m McpServicesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["include_browse"] = attrs["include_browse"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	attrs["mcp_services"] = attrs["mcp_services"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type McpServicesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *McpServicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *McpServicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, McpServicesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks McpService",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *McpServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *McpServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config McpServicesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListMcpServicesRequest
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

	response, err := client.AiGateway.ListMcpServicesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list ai_gateway_mcp_services", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var mcp_service McpServiceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &mcp_service)...)
		if resp.Diagnostics.HasError() {
			return
		}
		mcp_service.ProviderConfigData = config.ProviderConfigData

		results = append(results, mcp_service.ToObjectValue(ctx))
	}

	config.AiGateway = types.ListValueMust(McpServiceData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
