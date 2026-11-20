// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_builtin_parser

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

const dataSourcesName = "lakewatch_builtin_parsers"

var _ datasource.DataSourceWithConfigure = &BuiltinParsersDataSource{}

func DataSourceBuiltinParsers() datasource.DataSource {
	return &BuiltinParsersDataSource{}
}

// BuiltinParsersData extends the main model with additional fields.
type BuiltinParsersData struct {
	Lakewatch types.List `tfsdk:"builtin_parsers"`
	// AIP-160 filter expression. Built-ins are filterable by their `tags`; the
	// exact tag-filter grammar is defined by the server.
	Filter types.String `tfsdk:"filter"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (BuiltinParsersData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"builtin_parsers": reflect.TypeOf(BuiltinParserData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m BuiltinParsersData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	attrs["builtin_parsers"] = attrs["builtin_parsers"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type BuiltinParsersDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *BuiltinParsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *BuiltinParsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, BuiltinParsersData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks BuiltinParser",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BuiltinParsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *BuiltinParsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config BuiltinParsersData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListBuiltinParsersRequest
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

	response, err := client.Lakewatch.ListBuiltinParsersAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_builtin_parsers", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var builtin_parser BuiltinParserData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &builtin_parser)...)
		if resp.Diagnostics.HasError() {
			return
		}
		builtin_parser.ProviderConfigData = config.ProviderConfigData

		results = append(results, builtin_parser.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(BuiltinParserData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
