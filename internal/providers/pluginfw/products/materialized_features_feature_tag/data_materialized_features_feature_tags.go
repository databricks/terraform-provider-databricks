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
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "materialized_features_feature_tags"

var _ datasource.DataSourceWithConfigure = &FeatureTagsDataSource{}

func DataSourceFeatureTags() datasource.DataSource {
	return &FeatureTagsDataSource{}
}

// FeatureTagsData extends the main model with additional fields.
type FeatureTagsData struct {
	MaterializedFeatures types.List   `tfsdk:"feature_tags"`
	WorkspaceID          types.String `tfsdk:"workspace_id"`
}

func (FeatureTagsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags": reflect.TypeOf(ml_tf.FeatureTag{}),
	}
}

type FeatureTagsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *FeatureTagsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *FeatureTagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FeatureTagsData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("feature_tags")
		c.SetOptional("workspace_id")
		return c
	})
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

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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

	response, err := client.MaterializedFeatures.ListFeatureTagsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list materialized_features_feature_tags", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var feature_tag ml_tf.FeatureTag
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &feature_tag)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, feature_tag.ToObjectValue(ctx))
	}

	var newState FeatureTagsData
	newState.MaterializedFeatures = types.ListValueMust(ml_tf.FeatureTag{}.Type(ctx), results)
	newState.WorkspaceID = config.WorkspaceID
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
