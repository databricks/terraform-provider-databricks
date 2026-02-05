// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_setting_user_preference_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settingsv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "account_setting_user_preference_v2"

var _ datasource.DataSourceWithConfigure = &UserPreferenceDataSource{}

func DataSourceUserPreference() datasource.DataSource {
	return &UserPreferenceDataSource{}
}

type UserPreferenceDataSource struct {
	Client *autogen.DatabricksClient
}

// UserPreferenceData extends the main model with additional fields.
type UserPreferenceData struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	StringVal types.Object `tfsdk:"string_val"`
	// User ID of the user.
	UserId types.String `tfsdk:"user_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// UserPreferenceData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m UserPreferenceData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val":           reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_boolean_val": reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_string_val":  reflect.TypeOf(settingsv2_tf.StringMessage{}),
		"string_val":            reflect.TypeOf(settingsv2_tf.StringMessage{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserPreferenceData
// only implements ToObjectValue() and Type().
func (m UserPreferenceData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":           m.BooleanVal,
			"effective_boolean_val": m.EffectiveBooleanVal,
			"effective_string_val":  m.EffectiveStringVal,
			"name":                  m.Name,
			"string_val":            m.StringVal,
			"user_id":               m.UserId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m UserPreferenceData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":           settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_boolean_val": settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_string_val":  settingsv2_tf.StringMessage{}.Type(ctx),
			"name":                  types.StringType,
			"string_val":            settingsv2_tf.StringMessage{}.Type(ctx),
			"user_id":               types.StringType,
		},
	}
}

func (m UserPreferenceData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetComputed()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["string_val"] = attrs["string_val"].SetComputed()
	attrs["user_id"] = attrs["user_id"].SetRequired()

	return attrs
}

func (r *UserPreferenceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *UserPreferenceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, UserPreferenceData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks UserPreference",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *UserPreferenceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *UserPreferenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config UserPreferenceData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settingsv2.GetPublicAccountUserPreferenceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.SettingsV2.GetPublicAccountUserPreference(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get account_setting_user_preference_v2", err.Error())
		return
	}

	var newState UserPreferenceData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
