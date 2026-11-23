// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_parser_version

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/lakewatch"
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

const dataSourcesName = "lakewatch_parser_versions"

var _ datasource.DataSourceWithConfigure = &ParserVersionsDataSource{}

func DataSourceParserVersions() datasource.DataSource {
	return &ParserVersionsDataSource{}
}

// ParserVersionsData extends the main model with additional fields.
type ParserVersionsData struct {
	Lakewatch types.List `tfsdk:"parser_versions"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Resource name of the parser whose versions are listed. Format:
	// `parsers/{parser_id}`.
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ParserVersionsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parser_versions": reflect.TypeOf(ParserVersionData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ParserVersionsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["parser_versions"] = attrs["parser_versions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ParserVersionsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ParserVersionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ParserVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ParserVersionsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ParserVersion",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ParserVersionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ParserVersionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ParserVersionsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListParserVersionsRequest
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

	response, err := client.Lakewatch.ListParserVersionsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_parser_versions", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var parser_version ParserVersionData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &parser_version)...)
		if resp.Diagnostics.HasError() {
			return
		}
		parser_version.ProviderConfigData = config.ProviderConfigData

		results = append(results, parser_version.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(ParserVersionData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
