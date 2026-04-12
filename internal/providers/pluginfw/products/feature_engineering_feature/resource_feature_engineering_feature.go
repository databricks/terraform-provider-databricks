// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_feature

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "feature_engineering_feature"

var _ resource.ResourceWithConfigure = &FeatureResource{}
var _ resource.ResourceWithModifyPlan = &FeatureResource{}

func ResourceFeature() resource.Resource {
	return &FeatureResource{}
}

type FeatureResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// Feature extends the main model with additional fields.
type Feature struct {
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
	TimeseriesColumn types.Object `tfsdk:"timeseries_column"`
	ProviderConfig   types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Feature struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Feature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entities":          reflect.TypeOf(ml_tf.EntityColumn{}),
		"function":          reflect.TypeOf(ml_tf.Function{}),
		"inputs":            reflect.TypeOf(types.String{}),
		"lineage_context":   reflect.TypeOf(ml_tf.LineageContext{}),
		"source":            reflect.TypeOf(ml_tf.DataSource{}),
		"time_window":       reflect.TypeOf(ml_tf.TimeWindow{}),
		"timeseries_column": reflect.TypeOf(ml_tf.TimeseriesColumn{}),
		"provider_config":   reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Feature
// only implements ToObjectValue() and Type().
func (m Feature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"description": m.Description,
			"entities":          m.Entities,
			"filter_condition":  m.FilterCondition,
			"full_name":         m.FullName,
			"function":          m.Function,
			"inputs":            m.Inputs,
			"lineage_context":   m.LineageContext,
			"source":            m.Source,
			"time_window":       m.TimeWindow,
			"timeseries_column": m.TimeseriesColumn,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Feature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"description": types.StringType,
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

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Feature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Feature) {
	if !from.Entities.IsNull() && !from.Entities.IsUnknown() && to.Entities.IsNull() && len(from.Entities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entities = from.Entities
	}
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				// Recursively sync the fields of Function
				toFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Inputs.IsNull() && !from.Inputs.IsUnknown() && to.Inputs.IsNull() && len(from.Inputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Inputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Inputs = from.Inputs
	}
	if !from.LineageContext.IsNull() && !from.LineageContext.IsUnknown() {
		if toLineageContext, ok := to.GetLineageContext(ctx); ok {
			if fromLineageContext, ok := from.GetLineageContext(ctx); ok {
				// Recursively sync the fields of LineageContext
				toLineageContext.SyncFieldsDuringCreateOrUpdate(ctx, fromLineageContext)
				to.SetLineageContext(ctx, toLineageContext)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				// Recursively sync the fields of Source
				toSource.SyncFieldsDuringCreateOrUpdate(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.TimeWindow.IsNull() && !from.TimeWindow.IsUnknown() {
		if toTimeWindow, ok := to.GetTimeWindow(ctx); ok {
			if fromTimeWindow, ok := from.GetTimeWindow(ctx); ok {
				// Recursively sync the fields of TimeWindow
				toTimeWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromTimeWindow)
				to.SetTimeWindow(ctx, toTimeWindow)
			}
		}
	}
	if !from.TimeseriesColumn.IsNull() && !from.TimeseriesColumn.IsUnknown() {
		if toTimeseriesColumn, ok := to.GetTimeseriesColumn(ctx); ok {
			if fromTimeseriesColumn, ok := from.GetTimeseriesColumn(ctx); ok {
				// Recursively sync the fields of TimeseriesColumn
				toTimeseriesColumn.SyncFieldsDuringCreateOrUpdate(ctx, fromTimeseriesColumn)
				to.SetTimeseriesColumn(ctx, toTimeseriesColumn)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Feature) SyncFieldsDuringRead(ctx context.Context, from Feature) {
	if !from.Entities.IsNull() && !from.Entities.IsUnknown() && to.Entities.IsNull() && len(from.Entities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entities = from.Entities
	}
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				toFunction.SyncFieldsDuringRead(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Inputs.IsNull() && !from.Inputs.IsUnknown() && to.Inputs.IsNull() && len(from.Inputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Inputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Inputs = from.Inputs
	}
	if !from.LineageContext.IsNull() && !from.LineageContext.IsUnknown() {
		if toLineageContext, ok := to.GetLineageContext(ctx); ok {
			if fromLineageContext, ok := from.GetLineageContext(ctx); ok {
				toLineageContext.SyncFieldsDuringRead(ctx, fromLineageContext)
				to.SetLineageContext(ctx, toLineageContext)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				toSource.SyncFieldsDuringRead(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.TimeWindow.IsNull() && !from.TimeWindow.IsUnknown() {
		if toTimeWindow, ok := to.GetTimeWindow(ctx); ok {
			if fromTimeWindow, ok := from.GetTimeWindow(ctx); ok {
				toTimeWindow.SyncFieldsDuringRead(ctx, fromTimeWindow)
				to.SetTimeWindow(ctx, toTimeWindow)
			}
		}
	}
	if !from.TimeseriesColumn.IsNull() && !from.TimeseriesColumn.IsUnknown() {
		if toTimeseriesColumn, ok := to.GetTimeseriesColumn(ctx); ok {
			if fromTimeseriesColumn, ok := from.GetTimeseriesColumn(ctx); ok {
				toTimeseriesColumn.SyncFieldsDuringRead(ctx, fromTimeseriesColumn)
				to.SetTimeseriesColumn(ctx, toTimeseriesColumn)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Feature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["entities"] = attrs["entities"].SetOptional()
	attrs["filter_condition"] = attrs["filter_condition"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["function"] = attrs["function"].SetRequired()
	attrs["function"] = attrs["function"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["inputs"] = attrs["inputs"].SetOptional()
	attrs["inputs"] = attrs["inputs"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["lineage_context"] = attrs["lineage_context"].SetOptional()
	attrs["source"] = attrs["source"].SetRequired()
	attrs["source"] = attrs["source"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["time_window"] = attrs["time_window"].SetOptional()
	attrs["time_window"] = attrs["time_window"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["timeseries_column"] = attrs["timeseries_column"].SetOptional()

	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

// GetEntities returns the value of the Entities field in Feature as
// a slice of ml_tf.EntityColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetEntities(ctx context.Context) ([]ml_tf.EntityColumn, bool) {
	if m.Entities.IsNull() || m.Entities.IsUnknown() {
		return nil, false
	}
	var v []ml_tf.EntityColumn
	d := m.Entities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntities sets the value of the Entities field in Feature.
func (m *Feature) SetEntities(ctx context.Context, v []ml_tf.EntityColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entities = types.ListValueMust(t, vs)
}

// GetFunction returns the value of the Function field in Feature as
// a ml_tf.Function value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetFunction(ctx context.Context) (ml_tf.Function, bool) {
	var e ml_tf.Function
	if m.Function.IsNull() || m.Function.IsUnknown() {
		return e, false
	}
	var v ml_tf.Function
	d := m.Function.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunction sets the value of the Function field in Feature.
func (m *Feature) SetFunction(ctx context.Context, v ml_tf.Function) {
	vs := v.ToObjectValue(ctx)
	m.Function = vs
}

// GetInputs returns the value of the Inputs field in Feature as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetInputs(ctx context.Context) ([]types.String, bool) {
	if m.Inputs.IsNull() || m.Inputs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Inputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputs sets the value of the Inputs field in Feature.
func (m *Feature) SetInputs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Inputs = types.ListValueMust(t, vs)
}

// GetLineageContext returns the value of the LineageContext field in Feature as
// a ml_tf.LineageContext value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetLineageContext(ctx context.Context) (ml_tf.LineageContext, bool) {
	var e ml_tf.LineageContext
	if m.LineageContext.IsNull() || m.LineageContext.IsUnknown() {
		return e, false
	}
	var v ml_tf.LineageContext
	d := m.LineageContext.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLineageContext sets the value of the LineageContext field in Feature.
func (m *Feature) SetLineageContext(ctx context.Context, v ml_tf.LineageContext) {
	vs := v.ToObjectValue(ctx)
	m.LineageContext = vs
}

// GetSource returns the value of the Source field in Feature as
// a ml_tf.DataSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetSource(ctx context.Context) (ml_tf.DataSource, bool) {
	var e ml_tf.DataSource
	if m.Source.IsNull() || m.Source.IsUnknown() {
		return e, false
	}
	var v ml_tf.DataSource
	d := m.Source.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSource sets the value of the Source field in Feature.
func (m *Feature) SetSource(ctx context.Context, v ml_tf.DataSource) {
	vs := v.ToObjectValue(ctx)
	m.Source = vs
}

// GetTimeWindow returns the value of the TimeWindow field in Feature as
// a ml_tf.TimeWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetTimeWindow(ctx context.Context) (ml_tf.TimeWindow, bool) {
	var e ml_tf.TimeWindow
	if m.TimeWindow.IsNull() || m.TimeWindow.IsUnknown() {
		return e, false
	}
	var v ml_tf.TimeWindow
	d := m.TimeWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeWindow sets the value of the TimeWindow field in Feature.
func (m *Feature) SetTimeWindow(ctx context.Context, v ml_tf.TimeWindow) {
	vs := v.ToObjectValue(ctx)
	m.TimeWindow = vs
}

// GetTimeseriesColumn returns the value of the TimeseriesColumn field in Feature as
// a ml_tf.TimeseriesColumn value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetTimeseriesColumn(ctx context.Context) (ml_tf.TimeseriesColumn, bool) {
	var e ml_tf.TimeseriesColumn
	if m.TimeseriesColumn.IsNull() || m.TimeseriesColumn.IsUnknown() {
		return e, false
	}
	var v ml_tf.TimeseriesColumn
	d := m.TimeseriesColumn.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeseriesColumn sets the value of the TimeseriesColumn field in Feature.
func (m *Feature) SetTimeseriesColumn(ctx context.Context, v ml_tf.TimeseriesColumn) {
	vs := v.ToObjectValue(ctx)
	m.TimeseriesColumn = vs
}

func (r *FeatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *FeatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Feature{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks feature_engineering_feature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *FeatureResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip validation on destroy plans (plan is null).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	var plan Feature
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, validateDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())
	resp.Diagnostics.Append(validateDiags...)
}

func (r *FeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Feature
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var feature ml.Feature

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &feature)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := ml.CreateFeatureRequest{
		Feature: feature,
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.FeatureEngineering.CreateFeature(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create feature_engineering_feature", err.Error())
		return
	}

	var newState Feature

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

func (r *FeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Feature
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get feature_engineering_feature", err.Error())
		return
	}

	var newState Feature
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *FeatureResource) update(ctx context.Context, plan Feature, diags *diag.Diagnostics, state *tfsdk.State) {
	var feature ml.Feature

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &feature)...)
	if diags.HasError() {
		return
	}

	updateRequest := ml.UpdateFeatureRequest{
		Feature:    feature,
		FullName:   plan.FullName.ValueString(),
		UpdateMask: "description,entities,filter_condition,lineage_context,timeseries_column",
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.FeatureEngineering.UpdateFeature(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update feature_engineering_feature", err.Error())
		return
	}

	var newState Feature

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *FeatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Feature
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *FeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Feature
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest ml.DeleteFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	err := client.FeatureEngineering.DeleteFeature(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete feature_engineering_feature", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &FeatureResource{}

func (r *FeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: full_name. Got: %q",
				req.ID,
			),
		)
		return
	}

	fullName := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("full_name"), fullName)...)
}
