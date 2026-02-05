// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_kafka_config

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "feature_engineering_kafka_config"

var _ resource.ResourceWithConfigure = &KafkaConfigResource{}

func ResourceKafkaConfig() resource.Resource {
	return &KafkaConfigResource{}
}

type KafkaConfigResource struct {
	Client *autogen.DatabricksClient
}

// KafkaConfig extends the main model with additional fields.
type KafkaConfig struct {
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
// KafkaConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m KafkaConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
// interfere with how the plugin framework retrieves and sets values in state. Thus, KafkaConfig
// only implements ToObjectValue() and Type().
func (m KafkaConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"auth_config": m.AuthConfig,
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
func (m KafkaConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"auth_config": ml_tf.AuthConfig{}.Type(ctx),
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

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *KafkaConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KafkaConfig) {
	if !from.AuthConfig.IsNull() && !from.AuthConfig.IsUnknown() {
		if toAuthConfig, ok := to.GetAuthConfig(ctx); ok {
			if fromAuthConfig, ok := from.GetAuthConfig(ctx); ok {
				// Recursively sync the fields of AuthConfig
				toAuthConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAuthConfig)
				to.SetAuthConfig(ctx, toAuthConfig)
			}
		}
	}
	if !from.BackfillSource.IsNull() && !from.BackfillSource.IsUnknown() {
		if toBackfillSource, ok := to.GetBackfillSource(ctx); ok {
			if fromBackfillSource, ok := from.GetBackfillSource(ctx); ok {
				// Recursively sync the fields of BackfillSource
				toBackfillSource.SyncFieldsDuringCreateOrUpdate(ctx, fromBackfillSource)
				to.SetBackfillSource(ctx, toBackfillSource)
			}
		}
	}
	if !from.KeySchema.IsNull() && !from.KeySchema.IsUnknown() {
		if toKeySchema, ok := to.GetKeySchema(ctx); ok {
			if fromKeySchema, ok := from.GetKeySchema(ctx); ok {
				// Recursively sync the fields of KeySchema
				toKeySchema.SyncFieldsDuringCreateOrUpdate(ctx, fromKeySchema)
				to.SetKeySchema(ctx, toKeySchema)
			}
		}
	}
	if !from.SubscriptionMode.IsNull() && !from.SubscriptionMode.IsUnknown() {
		if toSubscriptionMode, ok := to.GetSubscriptionMode(ctx); ok {
			if fromSubscriptionMode, ok := from.GetSubscriptionMode(ctx); ok {
				// Recursively sync the fields of SubscriptionMode
				toSubscriptionMode.SyncFieldsDuringCreateOrUpdate(ctx, fromSubscriptionMode)
				to.SetSubscriptionMode(ctx, toSubscriptionMode)
			}
		}
	}
	if !from.ValueSchema.IsNull() && !from.ValueSchema.IsUnknown() {
		if toValueSchema, ok := to.GetValueSchema(ctx); ok {
			if fromValueSchema, ok := from.GetValueSchema(ctx); ok {
				// Recursively sync the fields of ValueSchema
				toValueSchema.SyncFieldsDuringCreateOrUpdate(ctx, fromValueSchema)
				to.SetValueSchema(ctx, toValueSchema)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *KafkaConfig) SyncFieldsDuringRead(ctx context.Context, from KafkaConfig) {
	if !from.AuthConfig.IsNull() && !from.AuthConfig.IsUnknown() {
		if toAuthConfig, ok := to.GetAuthConfig(ctx); ok {
			if fromAuthConfig, ok := from.GetAuthConfig(ctx); ok {
				toAuthConfig.SyncFieldsDuringRead(ctx, fromAuthConfig)
				to.SetAuthConfig(ctx, toAuthConfig)
			}
		}
	}
	if !from.BackfillSource.IsNull() && !from.BackfillSource.IsUnknown() {
		if toBackfillSource, ok := to.GetBackfillSource(ctx); ok {
			if fromBackfillSource, ok := from.GetBackfillSource(ctx); ok {
				toBackfillSource.SyncFieldsDuringRead(ctx, fromBackfillSource)
				to.SetBackfillSource(ctx, toBackfillSource)
			}
		}
	}
	if !from.KeySchema.IsNull() && !from.KeySchema.IsUnknown() {
		if toKeySchema, ok := to.GetKeySchema(ctx); ok {
			if fromKeySchema, ok := from.GetKeySchema(ctx); ok {
				toKeySchema.SyncFieldsDuringRead(ctx, fromKeySchema)
				to.SetKeySchema(ctx, toKeySchema)
			}
		}
	}
	if !from.SubscriptionMode.IsNull() && !from.SubscriptionMode.IsUnknown() {
		if toSubscriptionMode, ok := to.GetSubscriptionMode(ctx); ok {
			if fromSubscriptionMode, ok := from.GetSubscriptionMode(ctx); ok {
				toSubscriptionMode.SyncFieldsDuringRead(ctx, fromSubscriptionMode)
				to.SetSubscriptionMode(ctx, toSubscriptionMode)
			}
		}
	}
	if !from.ValueSchema.IsNull() && !from.ValueSchema.IsUnknown() {
		if toValueSchema, ok := to.GetValueSchema(ctx); ok {
			if fromValueSchema, ok := from.GetValueSchema(ctx); ok {
				toValueSchema.SyncFieldsDuringRead(ctx, fromValueSchema)
				to.SetValueSchema(ctx, toValueSchema)
			}
		}
	}
}

func (m KafkaConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auth_config"] = attrs["auth_config"].SetRequired()
	attrs["backfill_source"] = attrs["backfill_source"].SetOptional()
	attrs["bootstrap_servers"] = attrs["bootstrap_servers"].SetRequired()
	attrs["extra_options"] = attrs["extra_options"].SetOptional()
	attrs["key_schema"] = attrs["key_schema"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["subscription_mode"] = attrs["subscription_mode"].SetRequired()
	attrs["value_schema"] = attrs["value_schema"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetAuthConfig returns the value of the AuthConfig field in KafkaConfig as
// a ml_tf.AuthConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetAuthConfig(ctx context.Context) (ml_tf.AuthConfig, bool) {
	var e ml_tf.AuthConfig
	if m.AuthConfig.IsNull() || m.AuthConfig.IsUnknown() {
		return e, false
	}
	var v ml_tf.AuthConfig
	d := m.AuthConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAuthConfig sets the value of the AuthConfig field in KafkaConfig.
func (m *KafkaConfig) SetAuthConfig(ctx context.Context, v ml_tf.AuthConfig) {
	vs := v.ToObjectValue(ctx)
	m.AuthConfig = vs
}

// GetBackfillSource returns the value of the BackfillSource field in KafkaConfig as
// a ml_tf.BackfillSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetBackfillSource(ctx context.Context) (ml_tf.BackfillSource, bool) {
	var e ml_tf.BackfillSource
	if m.BackfillSource.IsNull() || m.BackfillSource.IsUnknown() {
		return e, false
	}
	var v ml_tf.BackfillSource
	d := m.BackfillSource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBackfillSource sets the value of the BackfillSource field in KafkaConfig.
func (m *KafkaConfig) SetBackfillSource(ctx context.Context, v ml_tf.BackfillSource) {
	vs := v.ToObjectValue(ctx)
	m.BackfillSource = vs
}

// GetExtraOptions returns the value of the ExtraOptions field in KafkaConfig as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetExtraOptions(ctx context.Context) (map[string]types.String, bool) {
	if m.ExtraOptions.IsNull() || m.ExtraOptions.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.ExtraOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExtraOptions sets the value of the ExtraOptions field in KafkaConfig.
func (m *KafkaConfig) SetExtraOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_options"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExtraOptions = types.MapValueMust(t, vs)
}

// GetKeySchema returns the value of the KeySchema field in KafkaConfig as
// a ml_tf.SchemaConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetKeySchema(ctx context.Context) (ml_tf.SchemaConfig, bool) {
	var e ml_tf.SchemaConfig
	if m.KeySchema.IsNull() || m.KeySchema.IsUnknown() {
		return e, false
	}
	var v ml_tf.SchemaConfig
	d := m.KeySchema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKeySchema sets the value of the KeySchema field in KafkaConfig.
func (m *KafkaConfig) SetKeySchema(ctx context.Context, v ml_tf.SchemaConfig) {
	vs := v.ToObjectValue(ctx)
	m.KeySchema = vs
}

// GetSubscriptionMode returns the value of the SubscriptionMode field in KafkaConfig as
// a ml_tf.SubscriptionMode value.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetSubscriptionMode(ctx context.Context) (ml_tf.SubscriptionMode, bool) {
	var e ml_tf.SubscriptionMode
	if m.SubscriptionMode.IsNull() || m.SubscriptionMode.IsUnknown() {
		return e, false
	}
	var v ml_tf.SubscriptionMode
	d := m.SubscriptionMode.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptionMode sets the value of the SubscriptionMode field in KafkaConfig.
func (m *KafkaConfig) SetSubscriptionMode(ctx context.Context, v ml_tf.SubscriptionMode) {
	vs := v.ToObjectValue(ctx)
	m.SubscriptionMode = vs
}

// GetValueSchema returns the value of the ValueSchema field in KafkaConfig as
// a ml_tf.SchemaConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *KafkaConfig) GetValueSchema(ctx context.Context) (ml_tf.SchemaConfig, bool) {
	var e ml_tf.SchemaConfig
	if m.ValueSchema.IsNull() || m.ValueSchema.IsUnknown() {
		return e, false
	}
	var v ml_tf.SchemaConfig
	d := m.ValueSchema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValueSchema sets the value of the ValueSchema field in KafkaConfig.
func (m *KafkaConfig) SetValueSchema(ctx context.Context, v ml_tf.SchemaConfig) {
	vs := v.ToObjectValue(ctx)
	m.ValueSchema = vs
}

func (r *KafkaConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *KafkaConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, KafkaConfig{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks feature_engineering_kafka_config",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KafkaConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *KafkaConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KafkaConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var kafka_config ml.KafkaConfig

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &kafka_config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := ml.CreateKafkaConfigRequest{
		KafkaConfig: kafka_config,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.CreateKafkaConfig(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create feature_engineering_kafka_config", err.Error())
		return
	}

	var newState KafkaConfig

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *KafkaConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState KafkaConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetKafkaConfigRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
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
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get feature_engineering_kafka_config", err.Error())
		return
	}

	var newState KafkaConfig
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *KafkaConfigResource) update(ctx context.Context, plan KafkaConfig, diags *diag.Diagnostics, state *tfsdk.State) {
	var kafka_config ml.KafkaConfig

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &kafka_config)...)
	if diags.HasError() {
		return
	}

	updateRequest := ml.UpdateKafkaConfigRequest{
		KafkaConfig: kafka_config,
		Name:        plan.Name.ValueString(),
		UpdateMask:  *fieldmask.New(strings.Split("auth_config,backfill_source,bootstrap_servers,extra_options,key_schema,subscription_mode,value_schema", ",")),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.FeatureEngineering.UpdateKafkaConfig(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update feature_engineering_kafka_config", err.Error())
		return
	}

	var newState KafkaConfig

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *KafkaConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan KafkaConfig
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *KafkaConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state KafkaConfig
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest ml.DeleteKafkaConfigRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.FeatureEngineering.DeleteKafkaConfig(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete feature_engineering_kafka_config", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &KafkaConfigResource{}

func (r *KafkaConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
