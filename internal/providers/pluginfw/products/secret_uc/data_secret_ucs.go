// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secret_uc

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

const dataSourcesName = "secret_ucs"

var _ datasource.DataSourceWithConfigure = &SecretsDataSource{}

func DataSourceSecrets() datasource.DataSource {
	return &SecretsDataSource{}
}

// SecretsData extends the main model with additional fields.
type SecretsData struct {
	SecretsUc types.List `tfsdk:"secrets"`
	// The name of the catalog under which to list secrets. Both
	// **catalog_name** and **schema_name** must be specified together.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Whether to include secrets in the response for which you only have the
	// **BROWSE** privilege, which limits access to metadata.
	IncludeBrowse types.Bool `tfsdk:"include_browse"`
	// Maximum number of secrets to return.
	//
	// - If not specified, at most 10000 secrets are returned. - If set to a
	// value greater than 0, the page length is the minimum of this value and
	// 10000. - If set to 0, the page length is set to 10000. - If set to a
	// value less than 0, an invalid parameter error is returned.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The name of the schema under which to list secrets. Both **catalog_name**
	// and **schema_name** must be specified together.
	SchemaName         types.String `tfsdk:"schema_name"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SecretsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets":         reflect.TypeOf(SecretData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SecretsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["include_browse"] = attrs["include_browse"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["secrets"] = attrs["secrets"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SecretsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SecretsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SecretsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SecretsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Secret",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SecretsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SecretsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SecretsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListSecretsRequest
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

	response, err := client.SecretsUc.ListSecretsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list secret_ucs", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var secret SecretData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &secret)...)
		if resp.Diagnostics.HasError() {
			return
		}
		secret.ProviderConfigData = config.ProviderConfigData

		results = append(results, secret.ToObjectValue(ctx))
	}

	config.SecretsUc = types.ListValueMust(SecretData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
