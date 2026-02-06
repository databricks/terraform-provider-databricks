// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_kafka_config

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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "feature_engineering_kafka_config"

var _ datasource.DataSourceWithConfigure = &KafkaConfigDataSource{}

func DataSourceKafkaConfig() datasource.DataSource {
	return &KafkaConfigDataSource{}
}

type KafkaConfigDataSource struct {
	Client *autogen.DatabricksClient
}

// KafkaConfigData extends the main model with additional fields.
type KafkaConfigData struct {
	// Authentication configuration for connection to topics.
	AuthConfig types.Object `tfsdk:"auth_config"`
	// A user-provided and managed source for backfilling data. Historical data
	// is used when creating a training set from streaming features linked to
	// this Kafka config. In the future, a separate table will be maintained by
	// Databricks for forward filling data. The schema for this source must
	// match exactly that of the key and value schemas specified for this Kafka
	// config.
	BackfillSource types.Object `tfsdk:"backfill_source"`
	// A comma-separated list of host/port pairs pointing to Kafka cluster.
	BootstrapServers types.String `tfsdk:"bootstrap_servers"`
	// Catch-all for miscellaneous options. Keys should be source options or
	// Kafka consumer options (kafka.*)
	ExtraOptions types.Map `tfsdk:"extra_options"`
	// Schema configuration for extracting message keys from topics. At least
	// one of key_schema and value_schema must be provided.
	KeySchema types.Object `tfsdk:"key_schema"`
	// Name that uniquely identifies this Kafka config within the metastore.
	// This will be the identifier used from the Feature object to reference
	// these configs for a feature. Can be distinct from topic name.
	Name types.String `tfsdk:"name"`
	// Options to configure which Kafka topics to pull data from.
	SubscriptionMode types.Object `tfsdk:"subscription_mode"`
	// Schema configuration for extracting message values from topics. At least
	// one of key_schema and value_schema must be provided.
	ValueSchema types.Object `tfsdk:"value_schema"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// KafkaConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m KafkaConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auth_config":       reflect.TypeOf(ml_tf.AuthConfig{}),
		"backfill_source":   reflect.TypeOf(ml_tf.BackfillSource{}),
		"extra_options":     reflect.TypeOf(types.String{}),
		"key_schema":        reflect.TypeOf(ml_tf.SchemaConfig{}),
		"subscription_mode": reflect.TypeOf(ml_tf.SubscriptionMode{}),
		"value_schema":      reflect.TypeOf(ml_tf.SchemaConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KafkaConfigData
// only implements ToObjectValue() and Type().
func (m KafkaConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auth_config":       m.AuthConfig,
			"backfill_source":   m.BackfillSource,
			"bootstrap_servers": m.BootstrapServers,
			"extra_options":     m.ExtraOptions,
			"key_schema":        m.KeySchema,
			"name":              m.Name,
			"subscription_mode": m.SubscriptionMode,
			"value_schema":      m.ValueSchema,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m KafkaConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_config":       ml_tf.AuthConfig{}.Type(ctx),
			"backfill_source":   ml_tf.BackfillSource{}.Type(ctx),
			"bootstrap_servers": types.StringType,
			"extra_options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"key_schema":        ml_tf.SchemaConfig{}.Type(ctx),
			"name":              types.StringType,
			"subscription_mode": ml_tf.SubscriptionMode{}.Type(ctx),
			"value_schema":      ml_tf.SchemaConfig{}.Type(ctx),
		},
	}
}

func (m KafkaConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auth_config"] = attrs["auth_config"].SetComputed()
	attrs["backfill_source"] = attrs["backfill_source"].SetComputed()
	attrs["bootstrap_servers"] = attrs["bootstrap_servers"].SetComputed()
	attrs["extra_options"] = attrs["extra_options"].SetComputed()
	attrs["key_schema"] = attrs["key_schema"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["subscription_mode"] = attrs["subscription_mode"].SetComputed()
	attrs["value_schema"] = attrs["value_schema"].SetComputed()

	return attrs
}

func (r *KafkaConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *KafkaConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, KafkaConfigData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks KafkaConfig",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KafkaConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *KafkaConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config KafkaConfigData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetKafkaConfigRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.GetKafkaConfig(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get feature_engineering_kafka_config", err.Error())
		return
	}

	var newState KafkaConfigData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
