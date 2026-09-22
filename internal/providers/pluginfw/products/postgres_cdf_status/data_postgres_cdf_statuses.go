// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_cdf_status

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

const dataSourcesName = "postgres_cdf_statuses"

var _ datasource.DataSourceWithConfigure = &CdfStatusDataSource{}

func DataSourceCdfStatus() datasource.DataSource {
	return &CdfStatusDataSource{}
}

// CdfStatusesData extends the main model with additional fields.
type CdfStatusesData struct {
	Postgres types.List `tfsdk:"cdf_statuses"`
	// Maximum number of CdfStatuses to return.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The parent CdfConfig to list CdfStatuses for. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (CdfStatusesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cdf_statuses":    reflect.TypeOf(CdfStatusData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m CdfStatusesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["cdf_statuses"] = attrs["cdf_statuses"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type CdfStatusDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CdfStatusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CdfStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CdfStatusesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CdfStatus",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CdfStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CdfStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config CdfStatusesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest postgres.ListCdfStatusesRequest
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

	response, err := client.Postgres.ListCdfStatusesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list postgres_cdf_statuses", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var cdf_status CdfStatusData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &cdf_status)...)
		if resp.Diagnostics.HasError() {
			return
		}
		cdf_status.ProviderConfigData = config.ProviderConfigData

		results = append(results, cdf_status.ToObjectValue(ctx))
	}

	config.Postgres = types.ListValueMust(CdfStatusData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
