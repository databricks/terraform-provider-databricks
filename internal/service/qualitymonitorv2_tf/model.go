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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AnomalyDetectionConfig struct {
	// Run id of the last run of the workflow
	LastRunId types.String `tfsdk:"last_run_id"`
	// The status of the last run of the workflow.
	LatestRunStatus types.String `tfsdk:"latest_run_status"`
}

func (newState *AnomalyDetectionConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AnomalyDetectionConfig) {
}

func (newState *AnomalyDetectionConfig) SyncEffectiveFieldsDuringRead(existingState AnomalyDetectionConfig) {
}

func (c AnomalyDetectionConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AnomalyDetectionConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnomalyDetectionConfig
// only implements ToObjectValue() and Type().
func (o AnomalyDetectionConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_run_id":       o.LastRunId,
			"latest_run_status": o.LatestRunStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AnomalyDetectionConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_run_id":       types.StringType,
			"latest_run_status": types.StringType,
		},
	}
}

type CreateQualityMonitorRequest struct {
	QualityMonitor types.Object `tfsdk:"quality_monitor"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o CreateQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quality_monitor": o.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quality_monitor": QualityMonitor{}.Type(ctx),
		},
	}
}

// GetQualityMonitor returns the value of the QualityMonitor field in CreateQualityMonitorRequest as
// a QualityMonitor value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateQualityMonitorRequest) GetQualityMonitor(ctx context.Context) (QualityMonitor, bool) {
	var e QualityMonitor
	if o.QualityMonitor.IsNull() || o.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v []QualityMonitor
	d := o.QualityMonitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQualityMonitor sets the value of the QualityMonitor field in CreateQualityMonitorRequest.
func (o *CreateQualityMonitorRequest) SetQualityMonitor(ctx context.Context, v QualityMonitor) {
	vs := v.ToObjectValue(ctx)
	o.QualityMonitor = vs
}

type DeleteQualityMonitorRequest struct {
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
func (a DeleteQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o DeleteQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type DeleteQualityMonitorResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQualityMonitorResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQualityMonitorResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorResponse
// only implements ToObjectValue() and Type().
func (o DeleteQualityMonitorResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQualityMonitorResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetQualityMonitorRequest struct {
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
func (a GetQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o GetQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type ListQualityMonitorRequest struct {
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
func (a ListQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o ListQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListQualityMonitorResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	QualityMonitors types.List `tfsdk:"quality_monitors"`
}

func (newState *ListQualityMonitorResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQualityMonitorResponse) {
}

func (newState *ListQualityMonitorResponse) SyncEffectiveFieldsDuringRead(existingState ListQualityMonitorResponse) {
}

func (c ListQualityMonitorResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListQualityMonitorResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitors": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorResponse
// only implements ToObjectValue() and Type().
func (o ListQualityMonitorResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":  o.NextPageToken,
			"quality_monitors": o.QualityMonitors,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQualityMonitorResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"quality_monitors": basetypes.ListType{
				ElemType: QualityMonitor{}.Type(ctx),
			},
		},
	}
}

// GetQualityMonitors returns the value of the QualityMonitors field in ListQualityMonitorResponse as
// a slice of QualityMonitor values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQualityMonitorResponse) GetQualityMonitors(ctx context.Context) ([]QualityMonitor, bool) {
	if o.QualityMonitors.IsNull() || o.QualityMonitors.IsUnknown() {
		return nil, false
	}
	var v []QualityMonitor
	d := o.QualityMonitors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQualityMonitors sets the value of the QualityMonitors field in ListQualityMonitorResponse.
func (o *ListQualityMonitorResponse) SetQualityMonitors(ctx context.Context, v []QualityMonitor) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quality_monitors"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.QualityMonitors = types.ListValueMust(t, vs)
}

type QualityMonitor struct {
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *QualityMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan QualityMonitor) {
}

func (newState *QualityMonitor) SyncEffectiveFieldsDuringRead(existingState QualityMonitor) {
}

func (c QualityMonitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetComputed()
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
func (a QualityMonitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(AnomalyDetectionConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitor
// only implements ToObjectValue() and Type().
func (o QualityMonitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anomaly_detection_config": o.AnomalyDetectionConfig,
			"object_id":                o.ObjectId,
			"object_type":              o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QualityMonitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anomaly_detection_config": AnomalyDetectionConfig{}.Type(ctx),
			"object_id":                types.StringType,
			"object_type":              types.StringType,
		},
	}
}

// GetAnomalyDetectionConfig returns the value of the AnomalyDetectionConfig field in QualityMonitor as
// a AnomalyDetectionConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *QualityMonitor) GetAnomalyDetectionConfig(ctx context.Context) (AnomalyDetectionConfig, bool) {
	var e AnomalyDetectionConfig
	if o.AnomalyDetectionConfig.IsNull() || o.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v []AnomalyDetectionConfig
	d := o.AnomalyDetectionConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAnomalyDetectionConfig sets the value of the AnomalyDetectionConfig field in QualityMonitor.
func (o *QualityMonitor) SetAnomalyDetectionConfig(ctx context.Context, v AnomalyDetectionConfig) {
	vs := v.ToObjectValue(ctx)
	o.AnomalyDetectionConfig = vs
}

type UpdateQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`

	QualityMonitor types.Object `tfsdk:"quality_monitor"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o UpdateQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":       o.ObjectId,
			"object_type":     o.ObjectType,
			"quality_monitor": o.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":       types.StringType,
			"object_type":     types.StringType,
			"quality_monitor": QualityMonitor{}.Type(ctx),
		},
	}
}

// GetQualityMonitor returns the value of the QualityMonitor field in UpdateQualityMonitorRequest as
// a QualityMonitor value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateQualityMonitorRequest) GetQualityMonitor(ctx context.Context) (QualityMonitor, bool) {
	var e QualityMonitor
	if o.QualityMonitor.IsNull() || o.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v []QualityMonitor
	d := o.QualityMonitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQualityMonitor sets the value of the QualityMonitor field in UpdateQualityMonitorRequest.
func (o *UpdateQualityMonitorRequest) SetQualityMonitor(ctx context.Context, v QualityMonitor) {
	vs := v.ToObjectValue(ctx)
	o.QualityMonitor = vs
}
