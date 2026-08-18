// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_gateway_model_service

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

const dataSourcesName = "ai_gateway_model_services"

var _ datasource.DataSourceWithConfigure = &ModelServicesDataSource{}

func DataSourceModelServices() datasource.DataSource {
	return &ModelServicesDataSource{}
}

// ModelServicesData extends the main model with additional fields.
type ModelServicesData struct {
	AiGateway types.List `tfsdk:"model_services"`
	// Maximum number of model services to return. Defaults to 100 when unset or
	// 0; the maximum is 100. Use `next_page_token` to retrieve additional
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

func (ModelServicesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_services":  reflect.TypeOf(ModelServiceData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ModelServicesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	attrs["model_services"] = attrs["model_services"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ModelServicesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ModelServicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ModelServicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ModelServicesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ModelService",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ModelServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ModelServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ModelServicesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListModelServicesRequest
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

	response, err := client.AiGateway.ListModelServicesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list ai_gateway_model_services", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var model_service ModelServiceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &model_service)...)
		if resp.Diagnostics.HasError() {
			return
		}
		model_service.ProviderConfigData = config.ProviderConfigData

		results = append(results, model_service.ToObjectValue(ctx))
	}

	config.AiGateway = types.ListValueMust(ModelServiceData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
