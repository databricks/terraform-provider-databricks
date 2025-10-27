// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_feature

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

const dataSourcesName = "feature_engineering_features"

var _ datasource.DataSourceWithConfigure = &FeaturesDataSource{}

func DataSourceFeatures() datasource.DataSource {
	return &FeaturesDataSource{}
}

// FeaturesData extends the main model with additional fields.
type FeaturesData struct {
	FeatureEngineering types.List `tfsdk:"features"`
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (FeaturesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(FeatureData{}),
	}
}

func (m FeaturesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["features"] = attrs["features"].SetComputed()
	return attrs
}

type FeaturesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *FeaturesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *FeaturesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FeaturesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Feature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FeaturesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FeaturesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config FeaturesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest ml.ListFeaturesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.ListFeaturesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list feature_engineering_features", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var feature FeatureData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &feature)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, feature.ToObjectValue(ctx))
	}

	config.FeatureEngineering = types.ListValueMust(FeatureData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
