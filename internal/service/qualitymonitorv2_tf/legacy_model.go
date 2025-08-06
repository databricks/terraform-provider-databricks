// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package qualitymonitorv2_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AnomalyDetectionConfig_SdkV2 struct {
	// Run id of the last run of the workflow
	LastRunId types.String `tfsdk:"last_run_id"`
	// The status of the last run of the workflow.
	LatestRunStatus types.String `tfsdk:"latest_run_status"`
}

func (toState *AnomalyDetectionConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AnomalyDetectionConfig_SdkV2) {
}

func (toState *AnomalyDetectionConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AnomalyDetectionConfig_SdkV2) {
}

func (c AnomalyDetectionConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_run_id"] = attrs["last_run_id"].SetComputed()
	attrs["latest_run_status"] = attrs["latest_run_status"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AnomalyDetectionConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AnomalyDetectionConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnomalyDetectionConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AnomalyDetectionConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_run_id":       o.LastRunId,
			"latest_run_status": o.LatestRunStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AnomalyDetectionConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_run_id":       types.StringType,
			"latest_run_status": types.StringType,
		},
	}
}

type CreateQualityMonitorRequest_SdkV2 struct {
	QualityMonitor types.List `tfsdk:"quality_monitor"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quality_monitor": o.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quality_monitor": basetypes.ListType{
				ElemType: QualityMonitor_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQualityMonitor returns the value of the QualityMonitor field in CreateQualityMonitorRequest_SdkV2 as
// a QualityMonitor_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQualityMonitorRequest_SdkV2) GetQualityMonitor(ctx context.Context) (QualityMonitor_SdkV2, bool) {
	var e QualityMonitor_SdkV2
	if o.QualityMonitor.IsNull() || o.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v []QualityMonitor_SdkV2
	d := o.QualityMonitor.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQualityMonitor sets the value of the QualityMonitor field in CreateQualityMonitorRequest_SdkV2.
func (o *CreateQualityMonitorRequest_SdkV2) SetQualityMonitor(ctx context.Context, v QualityMonitor_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quality_monitor"]
	o.QualityMonitor = types.ListValueMust(t, vs)
}

type DeleteQualityMonitorRequest_SdkV2 struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type GetQualityMonitorRequest_SdkV2 struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type ListQualityMonitorRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListQualityMonitorResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	QualityMonitors types.List `tfsdk:"quality_monitors"`
}

func (toState *ListQualityMonitorResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListQualityMonitorResponse_SdkV2) {
}

func (toState *ListQualityMonitorResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListQualityMonitorResponse_SdkV2) {
}

func (c ListQualityMonitorResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["quality_monitors"] = attrs["quality_monitors"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQualityMonitorResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQualityMonitorResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitors": reflect.TypeOf(QualityMonitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQualityMonitorResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":  o.NextPageToken,
			"quality_monitors": o.QualityMonitors,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQualityMonitorResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"quality_monitors": basetypes.ListType{
				ElemType: QualityMonitor_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQualityMonitors returns the value of the QualityMonitors field in ListQualityMonitorResponse_SdkV2 as
// a slice of QualityMonitor_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQualityMonitorResponse_SdkV2) GetQualityMonitors(ctx context.Context) ([]QualityMonitor_SdkV2, bool) {
	if o.QualityMonitors.IsNull() || o.QualityMonitors.IsUnknown() {
		return nil, false
	}
	var v []QualityMonitor_SdkV2
	d := o.QualityMonitors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQualityMonitors sets the value of the QualityMonitors field in ListQualityMonitorResponse_SdkV2.
func (o *ListQualityMonitorResponse_SdkV2) SetQualityMonitors(ctx context.Context, v []QualityMonitor_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quality_monitors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.QualityMonitors = types.ListValueMust(t, vs)
}

type QualityMonitor_SdkV2 struct {
	AnomalyDetectionConfig types.List `tfsdk:"anomaly_detection_config"`
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *QualityMonitor_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan QualityMonitor_SdkV2) {
	if !fromPlan.AnomalyDetectionConfig.IsNull() && !fromPlan.AnomalyDetectionConfig.IsUnknown() {
		if toStateAnomalyDetectionConfig, ok := toState.GetAnomalyDetectionConfig(ctx); ok {
			if fromPlanAnomalyDetectionConfig, ok := fromPlan.GetAnomalyDetectionConfig(ctx); ok {
				toStateAnomalyDetectionConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAnomalyDetectionConfig)
				toState.SetAnomalyDetectionConfig(ctx, toStateAnomalyDetectionConfig)
			}
		}
	}
}

func (toState *QualityMonitor_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState QualityMonitor_SdkV2) {
	if !fromState.AnomalyDetectionConfig.IsNull() && !fromState.AnomalyDetectionConfig.IsUnknown() {
		if toStateAnomalyDetectionConfig, ok := toState.GetAnomalyDetectionConfig(ctx); ok {
			if fromStateAnomalyDetectionConfig, ok := fromState.GetAnomalyDetectionConfig(ctx); ok {
				toStateAnomalyDetectionConfig.SyncFieldsDuringRead(ctx, fromStateAnomalyDetectionConfig)
				toState.SetAnomalyDetectionConfig(ctx, toStateAnomalyDetectionConfig)
			}
		}
	}
}

func (c QualityMonitor_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetComputed()
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QualityMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QualityMonitor_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(AnomalyDetectionConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitor_SdkV2
// only implements ToObjectValue() and Type().
func (o QualityMonitor_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anomaly_detection_config": o.AnomalyDetectionConfig,
			"object_id":                o.ObjectId,
			"object_type":              o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QualityMonitor_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anomaly_detection_config": basetypes.ListType{
				ElemType: AnomalyDetectionConfig_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in QualityMonitor_SdkV2 as
// a AnomalyDetectionConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QualityMonitor_SdkV2) GetAnomalyDetectionConfig(ctx context.Context) (AnomalyDetectionConfig_SdkV2, bool) {
	var e AnomalyDetectionConfig_SdkV2
	if o.AnomalyDetectionConfig.IsNull() || o.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v []AnomalyDetectionConfig_SdkV2
	d := o.AnomalyDetectionConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in QualityMonitor_SdkV2.
func (o *QualityMonitor_SdkV2) SetAnomalyDetectionConfig(ctx context.Context, v AnomalyDetectionConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["anomaly_detection_config"]
	o.AnomalyDetectionConfig = types.ListValueMust(t, vs)
}

type UpdateQualityMonitorRequest_SdkV2 struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`

	QualityMonitor types.List `tfsdk:"quality_monitor"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":       o.ObjectId,
			"object_type":     o.ObjectType,
			"quality_monitor": o.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
			"quality_monitor": basetypes.ListType{
				ElemType: QualityMonitor_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQualityMonitor returns the value of the QualityMonitor field in UpdateQualityMonitorRequest_SdkV2 as
// a QualityMonitor_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQualityMonitorRequest_SdkV2) GetQualityMonitor(ctx context.Context) (QualityMonitor_SdkV2, bool) {
	var e QualityMonitor_SdkV2
	if o.QualityMonitor.IsNull() || o.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v []QualityMonitor_SdkV2
	d := o.QualityMonitor.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQualityMonitor sets the value of the QualityMonitor field in UpdateQualityMonitorRequest_SdkV2.
func (o *UpdateQualityMonitorRequest_SdkV2) SetQualityMonitor(ctx context.Context, v QualityMonitor_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quality_monitor"]
	o.QualityMonitor = types.ListValueMust(t, vs)
}
