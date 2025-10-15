// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_metadata

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

const dataSourcesName = "external_metadatas"

var _ datasource.DataSourceWithConfigure = &ExternalMetadatasDataSource{}

func DataSourceExternalMetadatas() datasource.DataSource {
	return &ExternalMetadatasDataSource{}
}

// ExternalMetadatasData extends the main model with additional fields.
type ExternalMetadatasData struct {
	ExternalMetadata   types.List   `tfsdk:"external_metadata"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ExternalMetadatasData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_metadata": reflect.TypeOf(ExternalMetadataData{}),
		"provider_config":   reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ExternalMetadatasData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_metadata"] = attrs["external_metadata"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ExternalMetadatasDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ExternalMetadatasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ExternalMetadatasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ExternalMetadatasData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ExternalMetadata",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalMetadatasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ExternalMetadatasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ExternalMetadatasData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListExternalMetadataRequest
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

	response, err := client.ExternalMetadata.ListExternalMetadataAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list external_metadatas", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var external_metadata ExternalMetadataData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &external_metadata)...)
		if resp.Diagnostics.HasError() {
			return
		}
		external_metadata.ProviderConfigData = config.ProviderConfigData

		results = append(results, external_metadata.ToObjectValue(ctx))
	}

	var newState ExternalMetadatasData
	newState.ExternalMetadata = types.ListValueMust(ExternalMetadataData{}.Type(ctx), results)
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
