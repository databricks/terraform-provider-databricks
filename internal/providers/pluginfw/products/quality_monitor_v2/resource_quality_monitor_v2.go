// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package quality_monitor_v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/qualitymonitorv2_tf"
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

const resourceName = "quality_monitor_v2"

var _ resource.ResourceWithConfigure = &QualityMonitorResource{}

func ResourceQualityMonitor() resource.Resource {
	return &QualityMonitorResource{}
}

type QualityMonitorResource struct {
	Client *autogen.DatabricksClient
}

// QualityMonitor extends the main model with additional fields.
type QualityMonitor struct {
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"object_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// QualityMonitor struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m QualityMonitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(qualitymonitorv2_tf.AnomalyDetectionConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitor
// only implements ToObjectValue() and Type().
func (m QualityMonitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"anomaly_detection_config": m.AnomalyDetectionConfig,
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m QualityMonitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"anomaly_detection_config": qualitymonitorv2_tf.AnomalyDetectionConfig{}.Type(ctx),
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *QualityMonitor) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QualityMonitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				// Recursively sync the fields of AnomalyDetectionConfig
				toAnomalyDetectionConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *QualityMonitor) SyncFieldsDuringRead(ctx context.Context, from QualityMonitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				toAnomalyDetectionConfig.SyncFieldsDuringRead(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
}

func (m QualityMonitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	attrs["object_type"] = attrs["object_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["object_id"] = attrs["object_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in QualityMonitor as
// a qualitymonitorv2_tf.AnomalyDetectionConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *QualityMonitor) GetAnomalyDetectionConfig(ctx context.Context) (qualitymonitorv2_tf.AnomalyDetectionConfig, bool) {
	var e qualitymonitorv2_tf.AnomalyDetectionConfig
	if m.AnomalyDetectionConfig.IsNull() || m.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v qualitymonitorv2_tf.AnomalyDetectionConfig
	d := m.AnomalyDetectionConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in QualityMonitor.
func (m *QualityMonitor) SetAnomalyDetectionConfig(ctx context.Context, v qualitymonitorv2_tf.AnomalyDetectionConfig) {
	vs := v.ToObjectValue(ctx)
	m.AnomalyDetectionConfig = vs
}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, QualityMonitor{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks quality_monitor_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *QualityMonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *QualityMonitorResource) update(ctx context.Context, plan QualityMonitor, diags *diag.Diagnostics, state *tfsdk.State) {
	var quality_monitor qualitymonitorv2.QualityMonitor

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &quality_monitor)...)
	if diags.HasError() {
		return
	}

	updateRequest := qualitymonitorv2.UpdateQualityMonitorRequest{
		QualityMonitor: quality_monitor,
		ObjectId:       plan.ObjectId.ValueString(),
		ObjectType:     plan.ObjectType.ValueString(),
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.QualityMonitorV2.UpdateQualityMonitor(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update quality_monitor_v2", err.Error())
		return
	}

	var newState QualityMonitor

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *QualityMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan QualityMonitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var quality_monitor qualitymonitorv2.QualityMonitor

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &quality_monitor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := qualitymonitorv2.CreateQualityMonitorRequest{
		QualityMonitor: quality_monitor,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.QualityMonitorV2.CreateQualityMonitor(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create quality_monitor_v2", err.Error())
		return
	}

	var newState QualityMonitor

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

func (r *QualityMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState QualityMonitor
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest qualitymonitorv2.GetQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.QualityMonitorV2.GetQualityMonitor(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get quality_monitor_v2", err.Error())
		return
	}

	var newState QualityMonitor
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *QualityMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan QualityMonitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state QualityMonitor
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest qualitymonitorv2.DeleteQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.QualityMonitorV2.DeleteQualityMonitor(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete quality_monitor_v2", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &QualityMonitorResource{}

func (r *QualityMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: object_type,object_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	objectType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_type"), objectType)...)
	objectId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_id"), objectId)...)
}
