// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_feature

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "feature_engineering_feature"

var _ resource.ResourceWithConfigure = &FeatureResource{}

func ResourceFeature() resource.Resource {
	return &FeatureResource{}
}

type FeatureResource struct {
	Client *autogen.DatabricksClient
}

type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

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
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"full_name"`
	// The function by which the feature is computed.
	Function types.Object `tfsdk:"function"`
	// The input columns from which the feature is computed.
	Inputs types.List `tfsdk:"inputs"`
	// The data source of the feature.
	Source types.Object `tfsdk:"source"`
	// The time window in which the feature is computed.
	TimeWindow     types.Object `tfsdk:"time_window"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
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
		"function":        reflect.TypeOf(ml_tf.Function{}),
		"inputs":          reflect.TypeOf(types.String{}),
		"source":          reflect.TypeOf(ml_tf.DataSource{}),
		"time_window":     reflect.TypeOf(ml_tf.TimeWindow{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
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
			"full_name":   m.FullName,
			"function":    m.Function,
			"inputs":      m.Inputs,
			"source":      m.Source,
			"time_window": m.TimeWindow,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Feature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"description": types.StringType,
			"full_name": types.StringType,
			"function":  ml_tf.Function{}.Type(ctx),
			"inputs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"source":      ml_tf.DataSource{}.Type(ctx),
			"time_window": ml_tf.TimeWindow{}.Type(ctx),

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Feature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Feature) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				// Recursively sync the fields of Function
				toFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
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
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Feature) SyncFieldsDuringRead(ctx context.Context, from Feature) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				toFunction.SyncFieldsDuringRead(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
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
	to.ProviderConfig = from.ProviderConfig

}

func (m Feature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["function"] = attrs["function"].SetRequired()
	attrs["function"] = attrs["function"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["inputs"] = attrs["inputs"].SetRequired()
	attrs["inputs"] = attrs["inputs"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].SetRequired()
	attrs["source"] = attrs["source"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["time_window"] = attrs["time_window"].SetRequired()
	attrs["time_window"] = attrs["time_window"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
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

func (r *FeatureResource) update(ctx context.Context, plan Feature, diags *diag.Diagnostics, state *tfsdk.State) {
	var feature ml.Feature

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &feature)...)
	if diags.HasError() {
		return
	}

	updateRequest := ml.UpdateFeatureRequest{
		Feature:    feature,
		FullName:   plan.FullName.ValueString(),
		UpdateMask: "description",
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

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
