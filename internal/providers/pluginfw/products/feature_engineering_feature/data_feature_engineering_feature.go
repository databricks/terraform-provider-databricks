// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_feature

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

const dataSourceName = "feature_engineering_feature"

var _ datasource.DataSourceWithConfigure = &FeatureDataSource{}

func DataSourceFeature() datasource.DataSource {
	return &FeatureDataSource{}
}

type FeatureDataSource struct {
	Client *autogen.DatabricksClient
}

// FeatureData extends the main model with additional fields.
type FeatureData struct {
	// The description of the feature.
	Description types.String `tfsdk:"description"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"full_name"`
	// The function by which the feature is computed.
	Function types.Object `tfsdk:"function"`
	// The input columns from which the feature is computed.
	Inputs types.List `tfsdk:"inputs"`
	// The data source of the feature.
	Source types.Object `tfsdk:"source"`
	// The time window in which the feature is computed.
	TimeWindow types.Object `tfsdk:"time_window"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FeatureData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FeatureData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function":    reflect.TypeOf(ml_tf.Function{}),
		"inputs":      reflect.TypeOf(types.String{}),
		"source":      reflect.TypeOf(ml_tf.DataSource{}),
		"time_window": reflect.TypeOf(ml_tf.TimeWindow{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureData
// only implements ToObjectValue() and Type().
func (m FeatureData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"full_name":   m.FullName,
			"function":    m.Function,
			"inputs":      m.Inputs,
			"source":      m.Source,
			"time_window": m.TimeWindow,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FeatureData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"description": types.StringType,
			"full_name": types.StringType,
			"function":  ml_tf.Function{}.Type(ctx),
			"inputs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"source":      ml_tf.DataSource{}.Type(ctx),
			"time_window": ml_tf.TimeWindow{}.Type(ctx),
		},
	}
}

func (m FeatureData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["function"] = attrs["function"].SetRequired()
	attrs["inputs"] = attrs["inputs"].SetRequired()
	attrs["source"] = attrs["source"].SetRequired()
	attrs["time_window"] = attrs["time_window"].SetRequired()

	return attrs
}

func (r *FeatureDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *FeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FeatureData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Feature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FeatureDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FeatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config FeatureData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.GetFeature(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get feature_engineering_feature", err.Error())
		return
	}

	var newState FeatureData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
