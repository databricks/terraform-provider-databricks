// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_feature

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// FeatureData extends the main model with additional fields.
type FeatureData struct {
	// The description of the feature.
	Description types.String `tfsdk:"description"`
	// The entity columns for the feature, used as aggregation keys and for
	// query-time lookup.
	Entities types.List `tfsdk:"entities"`
	// Deprecated: Use DeltaTableSource.filter_condition or
	// KafkaSource.filter_condition instead. Kept for backwards compatibility.
	// The filter condition applied to the source data before aggregation.
	FilterCondition types.String `tfsdk:"filter_condition"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"full_name"`
	// The function by which the feature is computed.
	Function types.Object `tfsdk:"function"`
	// Deprecated: Use AggregationFunction.inputs instead. Kept for backwards
	// compatibility. The input columns from which the feature is computed.
	Inputs types.List `tfsdk:"inputs"`
	// Lineage context information for this feature. WARNING: This field is
	// primarily intended for internal use by Databricks systems and is
	// automatically populated when features are created through Databricks
	// notebooks or jobs. Users should not manually set this field as incorrect
	// values may lead to inaccurate lineage tracking or unexpected behavior.
	// This field will be set by feature-engineering client and should be left
	// unset by SDK and terraform users.
	LineageContext types.Object `tfsdk:"lineage_context"`
	// The data source of the feature.
	Source types.Object `tfsdk:"source"`
	// Deprecated: Use Function.aggregation_function.time_window instead. Kept
	// for backwards compatibility. The time window in which the feature is
	// computed.
	TimeWindow types.Object `tfsdk:"time_window"`
	// Column recording time, used for point-in-time joins, backfills, and
	// aggregations.
	TimeseriesColumn   types.Object `tfsdk:"timeseries_column"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
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
		"entities":          reflect.TypeOf(ml_tf.EntityColumn{}),
		"function":          reflect.TypeOf(ml_tf.Function{}),
		"inputs":            reflect.TypeOf(types.String{}),
		"lineage_context":   reflect.TypeOf(ml_tf.LineageContext{}),
		"source":            reflect.TypeOf(ml_tf.DataSource{}),
		"time_window":       reflect.TypeOf(ml_tf.TimeWindow{}),
		"timeseries_column": reflect.TypeOf(ml_tf.TimeseriesColumn{}),
		"provider_config":   reflect.TypeOf(ProviderConfigData{}),
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
			"description":       m.Description,
			"entities":          m.Entities,
			"filter_condition":  m.FilterCondition,
			"full_name":         m.FullName,
			"function":          m.Function,
			"inputs":            m.Inputs,
			"lineage_context":   m.LineageContext,
			"source":            m.Source,
			"time_window":       m.TimeWindow,
			"timeseries_column": m.TimeseriesColumn,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FeatureData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"entities": basetypes.ListType{
				ElemType: ml_tf.EntityColumn{}.Type(ctx),
			},
			"filter_condition": types.StringType,
			"full_name":        types.StringType,
			"function":         ml_tf.Function{}.Type(ctx),
			"inputs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"lineage_context":   ml_tf.LineageContext{}.Type(ctx),
			"source":            ml_tf.DataSource{}.Type(ctx),
			"time_window":       ml_tf.TimeWindow{}.Type(ctx),
			"timeseries_column": ml_tf.TimeseriesColumn{}.Type(ctx),

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m FeatureData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetComputed()
	attrs["entities"] = attrs["entities"].SetComputed()
	attrs["filter_condition"] = attrs["filter_condition"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["function"] = attrs["function"].SetComputed()
	attrs["inputs"] = attrs["inputs"].SetComputed()
	attrs["lineage_context"] = attrs["lineage_context"].SetComputed()
	attrs["source"] = attrs["source"].SetComputed()
	attrs["time_window"] = attrs["time_window"].SetComputed()
	attrs["timeseries_column"] = attrs["timeseries_column"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

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

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.GetFeature(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get feature_engineering_feature", err.Error())
		return
	}

	var newState FeatureData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
