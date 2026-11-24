// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_parser

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

const dataSourcesName = "lakewatch_parsers"

var _ datasource.DataSourceWithConfigure = &ParsersDataSource{}

func DataSourceParsers() datasource.DataSource {
	return &ParsersDataSource{}
}

// ParsersData extends the main model with additional fields.
type ParsersData struct {
	Lakewatch types.List `tfsdk:"parsers"`
	// Name of the parent UC catalog to list parsers from. Required: List is
	// scoped to a single catalog, mirroring the Unity Catalog tables/functions
	// APIs, which take catalog_name as a request field (bound to a query
	// parameter) rather than a catalog path segment.
	CatalogName types.String `tfsdk:"catalog_name"`
	// AIP-160 filter expression. Parsers are filterable by `parser_name`.
	Filter types.String `tfsdk:"filter"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Name of the parent UC schema. Optional: when set, results are narrowed to
	// this schema; when omitted, parsers across all schemas in `catalog_name`
	// are returned.
	SchemaName         types.String `tfsdk:"schema_name"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ParsersData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parsers":         reflect.TypeOf(ParserData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ParsersData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()

	attrs["parsers"] = attrs["parsers"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ParsersDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ParsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ParsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ParsersData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Parser",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ParsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ParsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ParsersData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListParsersRequest
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

	response, err := client.Lakewatch.ListParsersAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_parsers", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var parser ParserData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &parser)...)
		if resp.Diagnostics.HasError() {
			return
		}
		parser.ProviderConfigData = config.ProviderConfigData

		results = append(results, parser.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(ParserData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
