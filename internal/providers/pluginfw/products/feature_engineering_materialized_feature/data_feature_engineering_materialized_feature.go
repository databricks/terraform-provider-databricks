// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_materialized_feature

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "feature_engineering_materialized_feature"

var _ datasource.DataSourceWithConfigure = &MaterializedFeatureDataSource{}

func DataSourceMaterializedFeature() datasource.DataSource {
	return &MaterializedFeatureDataSource{}
}

type MaterializedFeatureDataSource struct {
	Client *autogen.DatabricksClient
}

// MaterializedFeatureData extends the main model with additional fields.
type MaterializedFeatureData struct {
	// The full name of the feature in Unity Catalog.
	FeatureName types.String `tfsdk:"feature_name"`
	// The timestamp when the pipeline last ran and updated the materialized
	// feature values. If the pipeline has not run yet, this field will be null.
	LastMaterializationTime types.String `tfsdk:"last_materialization_time"`
	// Unique identifier for the materialized feature.
	MaterializedFeatureId types.String `tfsdk:"materialized_feature_id"`

	OfflineStoreConfig types.Object `tfsdk:"offline_store_config"`

	OnlineStoreConfig types.Object `tfsdk:"online_store_config"`
	// The schedule state of the materialization pipeline.
	PipelineScheduleState types.String `tfsdk:"pipeline_schedule_state"`
	// The fully qualified Unity Catalog path to the table containing the
	// materialized feature (Delta table or Lakebase table). Output only.
	TableName types.String `tfsdk:"table_name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// MaterializedFeatureData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m MaterializedFeatureData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"offline_store_config": reflect.TypeOf(ml_tf.OfflineStoreConfig{}),
		"online_store_config":  reflect.TypeOf(ml_tf.OnlineStoreConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MaterializedFeatureData
// only implements ToObjectValue() and Type().
func (m MaterializedFeatureData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name":              m.FeatureName,
			"last_materialization_time": m.LastMaterializationTime,
			"materialized_feature_id":   m.MaterializedFeatureId,
			"offline_store_config":      m.OfflineStoreConfig,
			"online_store_config":       m.OnlineStoreConfig,
			"pipeline_schedule_state":   m.PipelineScheduleState,
			"table_name":                m.TableName,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m MaterializedFeatureData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name":              types.StringType,
			"last_materialization_time": types.StringType,
			"materialized_feature_id":   types.StringType,
			"offline_store_config":      ml_tf.OfflineStoreConfig{}.Type(ctx),
			"online_store_config":       ml_tf.OnlineStoreConfig{}.Type(ctx),
			"pipeline_schedule_state":   types.StringType,
			"table_name":                types.StringType,
		},
	}
}

func (m MaterializedFeatureData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetComputed()
	attrs["last_materialization_time"] = attrs["last_materialization_time"].SetComputed()
	attrs["materialized_feature_id"] = attrs["materialized_feature_id"].SetRequired()
	attrs["offline_store_config"] = attrs["offline_store_config"].SetComputed()
	attrs["online_store_config"] = attrs["online_store_config"].SetComputed()
	attrs["pipeline_schedule_state"] = attrs["pipeline_schedule_state"].SetComputed()
	attrs["table_name"] = attrs["table_name"].SetComputed()

	return attrs
}

func (r *MaterializedFeatureDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *MaterializedFeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MaterializedFeatureData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks MaterializedFeature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MaterializedFeatureDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *MaterializedFeatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config MaterializedFeatureData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetMaterializedFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.GetMaterializedFeature(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get feature_engineering_materialized_feature", err.Error())
		return
	}

	var newState MaterializedFeatureData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
