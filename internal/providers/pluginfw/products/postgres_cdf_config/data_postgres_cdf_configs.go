// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_cdf_config

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/postgres"
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

const dataSourcesName = "postgres_cdf_configs"

var _ datasource.DataSourceWithConfigure = &CdfConfigsDataSource{}

func DataSourceCdfConfigs() datasource.DataSource {
	return &CdfConfigsDataSource{}
}

// CdfConfigsData extends the main model with additional fields.
type CdfConfigsData struct {
	Postgres types.List `tfsdk:"cdf_configs"`
	// Maximum number of CdfConfigs to return.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The parent database to list CdfConfigs for. Format:
	// projects/{project}/branches/{branch}/databases/{database}
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (CdfConfigsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cdf_configs":     reflect.TypeOf(CdfConfigData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m CdfConfigsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["cdf_configs"] = attrs["cdf_configs"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type CdfConfigsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CdfConfigsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CdfConfigsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CdfConfigsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CdfConfig",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CdfConfigsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CdfConfigsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config CdfConfigsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest postgres.ListCdfConfigsRequest
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

	response, err := client.Postgres.ListCdfConfigsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list postgres_cdf_configs", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var cdf_config CdfConfigData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &cdf_config)...)
		if resp.Diagnostics.HasError() {
			return
		}
		cdf_config.ProviderConfigData = config.ProviderConfigData

		results = append(results, cdf_config.ToObjectValue(ctx))
	}

	config.Postgres = types.ListValueMust(CdfConfigData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
