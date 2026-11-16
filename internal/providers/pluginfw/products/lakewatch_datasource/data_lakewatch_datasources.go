// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_datasource

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

const dataSourcesName = "lakewatch_datasources"

var _ datasource.DataSourceWithConfigure = &DatasourcesDataSource{}

func DataSourceDatasources() datasource.DataSource {
	return &DatasourcesDataSource{}
}

// DatasourcesData extends the main model with additional fields.
type DatasourcesData struct {
	Lakewatch types.List `tfsdk:"datasources"`
	// Name of the Unity Catalog catalog to list datasources from. Required;
	// each request lists the datasources of a single catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Filter expression that restricts the results to datasources with a given
	// name, for example `datasource_name = "my_source"`. Only a single
	// exact-match comparison on `datasource_name` is supported, and it matches
	// the datasource's current Unity Catalog name, which reads report in
	// `effective_datasource_name`. Filtered results are returned in a single
	// page.
	Filter types.String `tfsdk:"filter"`
	// Maximum number of results to return in one page. This is a best-effort
	// upper bound, not a guarantee: the server may return fewer items than
	// requested (even zero) while more results still exist. If unset or 0, the
	// server picks a default.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Name of a Unity Catalog schema in `catalog_name`. Optional; if set, only
	// the datasources in this schema are returned, and if omitted, the
	// datasources in all schemas of the catalog are returned.
	SchemaName types.String `tfsdk:"schema_name"`
	// Which fields to return for each datasource. Defaults to
	// `DATASOURCE_VIEW_FULL`. `DATASOURCE_VIEW_BASIC` returns only the identity
	// and Unity Catalog location fields.
	View               types.String `tfsdk:"view"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (DatasourcesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasources":     reflect.TypeOf(DatasourceData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m DatasourcesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	attrs["datasources"] = attrs["datasources"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type DatasourcesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DatasourcesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DatasourcesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatasourcesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Datasource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatasourcesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DatasourcesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DatasourcesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListDatasourcesRequest
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

	response, err := client.Lakewatch.ListDatasourcesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_datasources", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var datasource DatasourceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &datasource)...)
		if resp.Diagnostics.HasError() {
			return
		}
		datasource.ProviderConfigData = config.ProviderConfigData

		results = append(results, datasource.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(DatasourceData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
