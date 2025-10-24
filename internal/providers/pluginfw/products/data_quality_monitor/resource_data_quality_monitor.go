// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_monitor

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/dataquality_tf"
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

const resourceName = "data_quality_monitor"

var _ resource.ResourceWithConfigure = &MonitorResource{}

func ResourceMonitor() resource.Resource {
	return &MonitorResource{}
}

type MonitorResource struct {
	Client *autogen.DatabricksClient
}

// Monitor extends the main model with additional fields.
type Monitor struct {
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
// Monitor struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Monitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(dataquality_tf.AnomalyDetectionConfig{}),
		"data_profiling_config":    reflect.TypeOf(dataquality_tf.DataProfilingConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Monitor
// only implements ToObjectValue() and Type().
func (m Monitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"anomaly_detection_config": m.AnomalyDetectionConfig,
			"data_profiling_config": m.DataProfilingConfig,
			"object_id":             m.ObjectId,
			"object_type":           m.ObjectType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Monitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"anomaly_detection_config": dataquality_tf.AnomalyDetectionConfig{}.Type(ctx),
			"data_profiling_config": dataquality_tf.DataProfilingConfig{}.Type(ctx),
			"object_id":             types.StringType,
			"object_type":           types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Monitor) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Monitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				// Recursively sync the fields of AnomalyDetectionConfig
				toAnomalyDetectionConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
	if !from.DataProfilingConfig.IsNull() && !from.DataProfilingConfig.IsUnknown() {
		if toDataProfilingConfig, ok := to.GetDataProfilingConfig(ctx); ok {
			if fromDataProfilingConfig, ok := from.GetDataProfilingConfig(ctx); ok {
				// Recursively sync the fields of DataProfilingConfig
				toDataProfilingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromDataProfilingConfig)
				to.SetDataProfilingConfig(ctx, toDataProfilingConfig)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Monitor) SyncFieldsDuringRead(ctx context.Context, from Monitor) {
	if !from.AnomalyDetectionConfig.IsNull() && !from.AnomalyDetectionConfig.IsUnknown() {
		if toAnomalyDetectionConfig, ok := to.GetAnomalyDetectionConfig(ctx); ok {
			if fromAnomalyDetectionConfig, ok := from.GetAnomalyDetectionConfig(ctx); ok {
				toAnomalyDetectionConfig.SyncFieldsDuringRead(ctx, fromAnomalyDetectionConfig)
				to.SetAnomalyDetectionConfig(ctx, toAnomalyDetectionConfig)
			}
		}
	}
	if !from.DataProfilingConfig.IsNull() && !from.DataProfilingConfig.IsUnknown() {
		if toDataProfilingConfig, ok := to.GetDataProfilingConfig(ctx); ok {
			if fromDataProfilingConfig, ok := from.GetDataProfilingConfig(ctx); ok {
				toDataProfilingConfig.SyncFieldsDuringRead(ctx, fromDataProfilingConfig)
				to.SetDataProfilingConfig(ctx, toDataProfilingConfig)
			}
		}
	}
}

func (m Monitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetOptional()
	attrs["data_profiling_config"] = attrs["data_profiling_config"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	attrs["object_type"] = attrs["object_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["object_id"] = attrs["object_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in Monitor as
// a dataquality_tf.AnomalyDetectionConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor) GetAnomalyDetectionConfig(ctx context.Context) (dataquality_tf.AnomalyDetectionConfig, bool) {
	var e dataquality_tf.AnomalyDetectionConfig
	if m.AnomalyDetectionConfig.IsNull() || m.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v dataquality_tf.AnomalyDetectionConfig
	d := m.AnomalyDetectionConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in Monitor.
func (m *Monitor) SetAnomalyDetectionConfig(ctx context.Context, v dataquality_tf.AnomalyDetectionConfig) {
	vs := v.ToObjectValue(ctx)
	m.AnomalyDetectionConfig = vs
}

// GetDataProfilingConfig returns the value of the DataProfilingConfig field in Monitor as
// a dataquality_tf.DataProfilingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Monitor) GetDataProfilingConfig(ctx context.Context) (dataquality_tf.DataProfilingConfig, bool) {
	var e dataquality_tf.DataProfilingConfig
	if m.DataProfilingConfig.IsNull() || m.DataProfilingConfig.IsUnknown() {
		return e, false
	}
	var v dataquality_tf.DataProfilingConfig
	d := m.DataProfilingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataProfilingConfig sets the value of the DataProfilingConfig field in Monitor.
func (m *Monitor) SetDataProfilingConfig(ctx context.Context, v dataquality_tf.DataProfilingConfig) {
	vs := v.ToObjectValue(ctx)
	m.DataProfilingConfig = vs
}

func (r *MonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *MonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Monitor{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks data_quality_monitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *MonitorResource) update(ctx context.Context, plan Monitor, diags *diag.Diagnostics, state *tfsdk.State) {
	var monitor dataquality.Monitor

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &monitor)...)
	if diags.HasError() {
		return
	}

	updateRequest := dataquality.UpdateMonitorRequest{
		Monitor:    monitor,
		ObjectId:   plan.ObjectId.ValueString(),
		ObjectType: plan.ObjectType.ValueString(),
		UpdateMask: "anomaly_detection_config,data_profiling_config",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.DataQuality.UpdateMonitor(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update data_quality_monitor", err.Error())
		return
	}

	var newState Monitor

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *MonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Monitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var monitor dataquality.Monitor

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &monitor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := dataquality.CreateMonitorRequest{
		Monitor: monitor,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.CreateMonitor(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create data_quality_monitor", err.Error())
		return
	}

	var newState Monitor

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

func (r *MonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Monitor
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest dataquality.GetMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
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

	var newState Monitor
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *MonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Monitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *MonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Monitor
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest dataquality.DeleteMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.DataQuality.DeleteMonitor(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete data_quality_monitor", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &MonitorResource{}

func (r *MonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
