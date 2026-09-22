// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_search_index

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/aisearch"
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

const dataSourcesName = "ai_search_indexes"

var _ datasource.DataSourceWithConfigure = &IndexesDataSource{}

func DataSourceIndexes() datasource.DataSource {
	return &IndexesDataSource{}
}

// IndexesData extends the main model with additional fields.
type IndexesData struct {
	AiSearch types.List `tfsdk:"indexes"`
	// Best-effort upper bound on the number of results to return. Honored as an
	// upper bound by the shim: `page_size` only narrows the legacy backend's
	// response, never widens it, so the practical cap is `min(page_size,
	// legacy_fixed_page_size)`.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The Endpoint that owns this collection of indexes. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (IndexesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"indexes":         reflect.TypeOf(IndexData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m IndexesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["indexes"] = attrs["indexes"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type IndexesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *IndexesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *IndexesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, IndexesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Index",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *IndexesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *IndexesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config IndexesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest aisearch.ListIndexesRequest
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

	response, err := client.AiSearch.ListIndexesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list ai_search_indexes", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var index IndexData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &index)...)
		if resp.Diagnostics.HasError() {
			return
		}
		index.ProviderConfigData = config.ProviderConfigData

		results = append(results, index.ToObjectValue(ctx))
	}

	config.AiSearch = types.ListValueMust(IndexData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
