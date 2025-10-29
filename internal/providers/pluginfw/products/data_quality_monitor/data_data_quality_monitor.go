// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_monitor

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/dataquality_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "data_quality_monitor"

var _ datasource.DataSourceWithConfigure = &MonitorDataSource{}

func DataSourceMonitor() datasource.DataSource {
	return &MonitorDataSource{}
}

type MonitorDataSource struct {
	Client *autogen.DatabricksClient
}

// MonitorData extends the main model with additional fields.
type MonitorData struct {
	// Anomaly Detection Configuration, applicable to `schema` object types.
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// Data Profiling Configuration, applicable to `table` object types. Exactly
	// one `Analysis Configuration` must be present.
	DataProfilingConfig types.Object `tfsdk:"data_profiling_config"`
	// The UUID of the request object. It is `schema_id` for `schema`, and
	// `table_id` for `table`.
	//
	// Find the `schema_id` from either: 1. The [schema_id] of the `Schemas`
	// resource. 2. In [Catalog Explorer] > select the `schema` > go to the
	// `Details` tab > the `Schema ID` field.
	//
	// Find the `table_id` from either: 1. The [table_id] of the `Tables`
	// resource. 2. In [Catalog Explorer] > select the `table` > go to the
	// `Details` tab > the `Table ID` field.
	//
	// [Catalog Explorer]: https://docs.databricks.com/aws/en/catalog-explorer/
	// [schema_id]: https://docs.databricks.com/api/workspace/schemas/get#schema_id
	// [table_id]: https://docs.databricks.com/api/workspace/tables/get#table_id
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType types.String `tfsdk:"object_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// MonitorData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m MonitorData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(dataquality_tf.AnomalyDetectionConfig{}),
		"data_profiling_config":    reflect.TypeOf(dataquality_tf.DataProfilingConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorData
// only implements ToObjectValue() and Type().
func (m MonitorData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anomaly_detection_config": m.AnomalyDetectionConfig,
			"data_profiling_config":    m.DataProfilingConfig,
			"object_id":                m.ObjectId,
			"object_type":              m.ObjectType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m MonitorData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anomaly_detection_config": dataquality_tf.AnomalyDetectionConfig{}.Type(ctx),
			"data_profiling_config":    dataquality_tf.DataProfilingConfig{}.Type(ctx),
			"object_id":                types.StringType,
			"object_type":              types.StringType,
		},
	}
}

func (m MonitorData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetComputed()
	attrs["data_profiling_config"] = attrs["data_profiling_config"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	return attrs
}

func (r *MonitorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *MonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MonitorData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Monitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MonitorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *MonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config MonitorData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest dataquality.GetMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.GetMonitor(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get data_quality_monitor", err.Error())
		return
	}

	var newState MonitorData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
