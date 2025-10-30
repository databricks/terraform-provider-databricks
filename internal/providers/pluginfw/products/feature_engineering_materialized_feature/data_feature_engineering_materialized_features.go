// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_materialized_feature

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
)

const dataSourcesName = "feature_engineering_materialized_features"

var _ datasource.DataSourceWithConfigure = &MaterializedFeaturesDataSource{}

func DataSourceMaterializedFeatures() datasource.DataSource {
	return &MaterializedFeaturesDataSource{}
}

// MaterializedFeaturesData extends the main model with additional fields.
type MaterializedFeaturesData struct {
	FeatureEngineering types.List `tfsdk:"materialized_features"`
	// Filter by feature name. If specified, only materialized features
	// materialized from this feature will be returned.
	FeatureName types.String `tfsdk:"feature_name"`
	// The maximum number of results to return. Defaults to 100 if not
	// specified. Cannot be greater than 1000.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (MaterializedFeaturesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"materialized_features": reflect.TypeOf(MaterializedFeatureData{}),
	}
}

func (m MaterializedFeaturesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["materialized_features"] = attrs["materialized_features"].SetComputed()
	return attrs
}

type MaterializedFeaturesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *MaterializedFeaturesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *MaterializedFeaturesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MaterializedFeaturesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks MaterializedFeature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MaterializedFeaturesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *MaterializedFeaturesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config MaterializedFeaturesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest ml.ListMaterializedFeaturesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.ListMaterializedFeaturesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list feature_engineering_materialized_features", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var materialized_feature MaterializedFeatureData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &materialized_feature)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, materialized_feature.ToObjectValue(ctx))
	}

	config.FeatureEngineering = types.ListValueMust(MaterializedFeatureData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
