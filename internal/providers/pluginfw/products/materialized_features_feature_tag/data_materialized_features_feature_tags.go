// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package materialized_features_feature_tag

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/ml"
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

const dataSourcesName = "materialized_features_feature_tags"

var _ datasource.DataSourceWithConfigure = &FeatureTagsDataSource{}

func DataSourceFeatureTags() datasource.DataSource {
	return &FeatureTagsDataSource{}
}

// FeatureTagsData extends the main model with additional fields.
type FeatureTagsData struct {
	MaterializedFeatures types.List   `tfsdk:"feature_tags"`
	ProviderConfigData   types.Object `tfsdk:"provider_config"`
}

func (FeatureTagsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags":    reflect.TypeOf(FeatureTagData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m FeatureTagsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tags"] = attrs["feature_tags"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type FeatureTagsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *FeatureTagsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *FeatureTagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FeatureTagsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks FeatureTag",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FeatureTagsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FeatureTagsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config FeatureTagsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest ml.ListFeatureTagsRequest
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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.MaterializedFeatures.ListFeatureTagsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list materialized_features_feature_tags", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var feature_tag FeatureTagData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &feature_tag)...)
		if resp.Diagnostics.HasError() {
			return
		}
		feature_tag.ProviderConfigData = config.ProviderConfigData

		results = append(results, feature_tag.ToObjectValue(ctx))
	}

	var newState FeatureTagsData
	newState.MaterializedFeatures = types.ListValueMust(FeatureTagData{}.Type(ctx), results)
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
