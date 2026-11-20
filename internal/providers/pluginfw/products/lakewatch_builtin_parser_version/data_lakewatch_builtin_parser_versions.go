// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_builtin_parser_version

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

const dataSourcesName = "lakewatch_builtin_parser_versions"

var _ datasource.DataSourceWithConfigure = &BuiltinParserVersionsDataSource{}

func DataSourceBuiltinParserVersions() datasource.DataSource {
	return &BuiltinParserVersionsDataSource{}
}

// BuiltinParserVersionsData extends the main model with additional fields.
type BuiltinParserVersionsData struct {
	Lakewatch types.List `tfsdk:"builtin_parser_versions"`
	// Maximum number of results to return in one page. If unset or 0, at most
	// 100 versions are returned. The maximum allowed value is 1000.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Resource name of the built-in parser whose versions are listed. Format:
	// `builtin-parsers/{parser_id}`.
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (BuiltinParserVersionsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"builtin_parser_versions": reflect.TypeOf(BuiltinParserVersionData{}),
		"provider_config":         reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m BuiltinParserVersionsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["builtin_parser_versions"] = attrs["builtin_parser_versions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type BuiltinParserVersionsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *BuiltinParserVersionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *BuiltinParserVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, BuiltinParserVersionsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks BuiltinParserVersion",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BuiltinParserVersionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *BuiltinParserVersionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config BuiltinParserVersionsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListBuiltinParserVersionsRequest
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

	response, err := client.Lakewatch.ListBuiltinParserVersionsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_builtin_parser_versions", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var builtin_parser_version BuiltinParserVersionData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &builtin_parser_version)...)
		if resp.Diagnostics.HasError() {
			return
		}
		builtin_parser_version.ProviderConfigData = config.ProviderConfigData

		results = append(results, builtin_parser_version.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(BuiltinParserVersionData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
