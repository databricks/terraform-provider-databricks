// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package feature_engineering_materialized_feature

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "feature_engineering_materialized_feature"

var _ resource.ResourceWithConfigure = &MaterializedFeatureResource{}

func ResourceMaterializedFeature() resource.Resource {
	return &MaterializedFeatureResource{}
}

type MaterializedFeatureResource struct {
	Client *autogen.DatabricksClient
}

// MaterializedFeature extends the main model with additional fields.
type MaterializedFeature struct {
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
// MaterializedFeature struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m MaterializedFeature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"offline_store_config": reflect.TypeOf(ml_tf.OfflineStoreConfig{}),
		"online_store_config":  reflect.TypeOf(ml_tf.OnlineStoreConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MaterializedFeature
// only implements ToObjectValue() and Type().
func (m MaterializedFeature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"feature_name": m.FeatureName,
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
func (m MaterializedFeature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"feature_name": types.StringType,
			"last_materialization_time": types.StringType,
			"materialized_feature_id":   types.StringType,
			"offline_store_config":      ml_tf.OfflineStoreConfig{}.Type(ctx),
			"online_store_config":       ml_tf.OnlineStoreConfig{}.Type(ctx),
			"pipeline_schedule_state":   types.StringType,
			"table_name":                types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *MaterializedFeature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MaterializedFeature) {
	if !from.OfflineStoreConfig.IsNull() && !from.OfflineStoreConfig.IsUnknown() {
		if toOfflineStoreConfig, ok := to.GetOfflineStoreConfig(ctx); ok {
			if fromOfflineStoreConfig, ok := from.GetOfflineStoreConfig(ctx); ok {
				// Recursively sync the fields of OfflineStoreConfig
				toOfflineStoreConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromOfflineStoreConfig)
				to.SetOfflineStoreConfig(ctx, toOfflineStoreConfig)
			}
		}
	}
	if !from.OnlineStoreConfig.IsNull() && !from.OnlineStoreConfig.IsUnknown() {
		if toOnlineStoreConfig, ok := to.GetOnlineStoreConfig(ctx); ok {
			if fromOnlineStoreConfig, ok := from.GetOnlineStoreConfig(ctx); ok {
				// Recursively sync the fields of OnlineStoreConfig
				toOnlineStoreConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromOnlineStoreConfig)
				to.SetOnlineStoreConfig(ctx, toOnlineStoreConfig)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *MaterializedFeature) SyncFieldsDuringRead(ctx context.Context, from MaterializedFeature) {
	if !from.OfflineStoreConfig.IsNull() && !from.OfflineStoreConfig.IsUnknown() {
		if toOfflineStoreConfig, ok := to.GetOfflineStoreConfig(ctx); ok {
			if fromOfflineStoreConfig, ok := from.GetOfflineStoreConfig(ctx); ok {
				toOfflineStoreConfig.SyncFieldsDuringRead(ctx, fromOfflineStoreConfig)
				to.SetOfflineStoreConfig(ctx, toOfflineStoreConfig)
			}
		}
	}
	if !from.OnlineStoreConfig.IsNull() && !from.OnlineStoreConfig.IsUnknown() {
		if toOnlineStoreConfig, ok := to.GetOnlineStoreConfig(ctx); ok {
			if fromOnlineStoreConfig, ok := from.GetOnlineStoreConfig(ctx); ok {
				toOnlineStoreConfig.SyncFieldsDuringRead(ctx, fromOnlineStoreConfig)
				to.SetOnlineStoreConfig(ctx, toOnlineStoreConfig)
			}
		}
	}
}

func (m MaterializedFeature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetRequired()
	attrs["last_materialization_time"] = attrs["last_materialization_time"].SetComputed()
	attrs["materialized_feature_id"] = attrs["materialized_feature_id"].SetComputed()
	attrs["offline_store_config"] = attrs["offline_store_config"].SetOptional()
	attrs["online_store_config"] = attrs["online_store_config"].SetOptional()
	attrs["pipeline_schedule_state"] = attrs["pipeline_schedule_state"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetComputed()

	attrs["materialized_feature_id"] = attrs["materialized_feature_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetOfflineStoreConfig returns the value of the OfflineStoreConfig field in MaterializedFeature as
// a ml_tf.OfflineStoreConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *MaterializedFeature) GetOfflineStoreConfig(ctx context.Context) (ml_tf.OfflineStoreConfig, bool) {
	var e ml_tf.OfflineStoreConfig
	if m.OfflineStoreConfig.IsNull() || m.OfflineStoreConfig.IsUnknown() {
		return e, false
	}
	var v ml_tf.OfflineStoreConfig
	d := m.OfflineStoreConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOfflineStoreConfig sets the value of the OfflineStoreConfig field in MaterializedFeature.
func (m *MaterializedFeature) SetOfflineStoreConfig(ctx context.Context, v ml_tf.OfflineStoreConfig) {
	vs := v.ToObjectValue(ctx)
	m.OfflineStoreConfig = vs
}

// GetOnlineStoreConfig returns the value of the OnlineStoreConfig field in MaterializedFeature as
// a ml_tf.OnlineStoreConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *MaterializedFeature) GetOnlineStoreConfig(ctx context.Context) (ml_tf.OnlineStoreConfig, bool) {
	var e ml_tf.OnlineStoreConfig
	if m.OnlineStoreConfig.IsNull() || m.OnlineStoreConfig.IsUnknown() {
		return e, false
	}
	var v ml_tf.OnlineStoreConfig
	d := m.OnlineStoreConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStoreConfig sets the value of the OnlineStoreConfig field in MaterializedFeature.
func (m *MaterializedFeature) SetOnlineStoreConfig(ctx context.Context, v ml_tf.OnlineStoreConfig) {
	vs := v.ToObjectValue(ctx)
	m.OnlineStoreConfig = vs
}

func (r *MaterializedFeatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *MaterializedFeatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, MaterializedFeature{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks feature_engineering_materialized_feature",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MaterializedFeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *MaterializedFeatureResource) update(ctx context.Context, plan MaterializedFeature, diags *diag.Diagnostics, state *tfsdk.State) {
	var materialized_feature ml.MaterializedFeature

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &materialized_feature)...)
	if diags.HasError() {
		return
	}

	updateRequest := ml.UpdateMaterializedFeatureRequest{
		MaterializedFeature:   materialized_feature,
		MaterializedFeatureId: plan.MaterializedFeatureId.ValueString(),
		UpdateMask:            "feature_name,offline_store_config,online_store_config,pipeline_schedule_state",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.FeatureEngineering.UpdateMaterializedFeature(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update feature_engineering_materialized_feature", err.Error())
		return
	}

	var newState MaterializedFeature

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *MaterializedFeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan MaterializedFeature
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var materialized_feature ml.MaterializedFeature

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &materialized_feature)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := ml.CreateMaterializedFeatureRequest{
		MaterializedFeature: materialized_feature,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureEngineering.CreateMaterializedFeature(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create feature_engineering_materialized_feature", err.Error())
		return
	}

	var newState MaterializedFeature

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

func (r *MaterializedFeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState MaterializedFeature
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetMaterializedFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
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

	var newState MaterializedFeature
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *MaterializedFeatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan MaterializedFeature
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *MaterializedFeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state MaterializedFeature
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest ml.DeleteMaterializedFeatureRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.FeatureEngineering.DeleteMaterializedFeature(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete feature_engineering_materialized_feature", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &MaterializedFeatureResource{}

func (r *MaterializedFeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: materialized_feature_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	materializedFeatureId := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("materialized_feature_id"), materializedFeatureId)...)
}
