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
	// The type of the last run of the workflow.
	JobType types.String `tfsdk:"job_type"`
	// Run id of the last run of the workflow
	LastRunId types.String `tfsdk:"last_run_id"`
	// The status of the last run of the workflow.
	LatestRunStatus types.String `tfsdk:"latest_run_status"`
}

func (to *AnomalyDetectionConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AnomalyDetectionConfig) {
}

func (to *AnomalyDetectionConfig) SyncFieldsDuringRead(ctx context.Context, from AnomalyDetectionConfig) {
}

func (m AnomalyDetectionConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["job_type"] = attrs["job_type"].SetComputed()
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
func (m AnomalyDetectionConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnomalyDetectionConfig
// only implements ToObjectValue() and Type().
func (m AnomalyDetectionConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_type":          m.JobType,
			"last_run_id":       m.LastRunId,
			"latest_run_status": m.LatestRunStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AnomalyDetectionConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_type":          types.StringType,
			"last_run_id":       types.StringType,
			"latest_run_status": types.StringType,
		},
	}
}

type CreateQualityMonitorRequest struct {
	QualityMonitor types.Object `tfsdk:"quality_monitor"`
}

func (to *CreateQualityMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateQualityMonitorRequest) {
	if !from.QualityMonitor.IsNull() && !from.QualityMonitor.IsUnknown() {
		if toQualityMonitor, ok := to.GetQualityMonitor(ctx); ok {
			if fromQualityMonitor, ok := from.GetQualityMonitor(ctx); ok {
				// Recursively sync the fields of QualityMonitor
				toQualityMonitor.SyncFieldsDuringCreateOrUpdate(ctx, fromQualityMonitor)
				to.SetQualityMonitor(ctx, toQualityMonitor)
			}
		}
	}
}

func (to *CreateQualityMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from CreateQualityMonitorRequest) {
	if !from.QualityMonitor.IsNull() && !from.QualityMonitor.IsUnknown() {
		if toQualityMonitor, ok := to.GetQualityMonitor(ctx); ok {
			if fromQualityMonitor, ok := from.GetQualityMonitor(ctx); ok {
				toQualityMonitor.SyncFieldsDuringRead(ctx, fromQualityMonitor)
				to.SetQualityMonitor(ctx, toQualityMonitor)
			}
		}
	}
}

func (m CreateQualityMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quality_monitor"] = attrs["quality_monitor"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (m CreateQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quality_monitor": m.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quality_monitor": QualityMonitor{}.Type(ctx),
		},
	}
}

// GetQualityMonitor returns the value of the QualityMonitor field in CreateQualityMonitorRequest as
// a QualityMonitor value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateQualityMonitorRequest) GetQualityMonitor(ctx context.Context) (QualityMonitor, bool) {
	var e QualityMonitor
	if m.QualityMonitor.IsNull() || m.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v QualityMonitor
	d := m.QualityMonitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQualityMonitor sets the value of the QualityMonitor field in CreateQualityMonitorRequest.
func (m *CreateQualityMonitorRequest) SetQualityMonitor(ctx context.Context, v QualityMonitor) {
	vs := v.ToObjectValue(ctx)
	m.QualityMonitor = vs
}

type DeleteQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`
}

func (to *DeleteQualityMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteQualityMonitorRequest) {
}

func (to *DeleteQualityMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteQualityMonitorRequest) {
}

func (m DeleteQualityMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (m DeleteQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type GetQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`
}

func (to *GetQualityMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetQualityMonitorRequest) {
}

func (to *GetQualityMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from GetQualityMonitorRequest) {
}

func (m GetQualityMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (m GetQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetQualityMonitorRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListQualityMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQualityMonitorRequest) {
}

func (to *ListQualityMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from ListQualityMonitorRequest) {
}

func (m ListQualityMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (m ListQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQualityMonitorRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListQualityMonitorResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListQualityMonitorResponse) {
	if !from.QualityMonitors.IsNull() && !from.QualityMonitors.IsUnknown() && to.QualityMonitors.IsNull() && len(from.QualityMonitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QualityMonitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QualityMonitors = from.QualityMonitors
	}
}

func (to *ListQualityMonitorResponse) SyncFieldsDuringRead(ctx context.Context, from ListQualityMonitorResponse) {
	if !from.QualityMonitors.IsNull() && !from.QualityMonitors.IsUnknown() && to.QualityMonitors.IsNull() && len(from.QualityMonitors.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QualityMonitors, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QualityMonitors = from.QualityMonitors
	}
}

func (m ListQualityMonitorResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListQualityMonitorResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitors": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQualityMonitorResponse
// only implements ToObjectValue() and Type().
func (m ListQualityMonitorResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":  m.NextPageToken,
			"quality_monitors": m.QualityMonitors,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListQualityMonitorResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListQualityMonitorResponse) GetQualityMonitors(ctx context.Context) ([]QualityMonitor, bool) {
	if m.QualityMonitors.IsNull() || m.QualityMonitors.IsUnknown() {
		return nil, false
	}
	var v []QualityMonitor
	d := m.QualityMonitors.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQualityMonitors sets the value of the QualityMonitors field in ListQualityMonitorResponse.
func (m *ListQualityMonitorResponse) SetQualityMonitors(ctx context.Context, v []QualityMonitor) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["quality_monitors"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QualityMonitors = types.ListValueMust(t, vs)
}

type QualityMonitor struct {
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"object_type"`
}

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

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QualityMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QualityMonitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(AnomalyDetectionConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitor
// only implements ToObjectValue() and Type().
func (m QualityMonitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anomaly_detection_config": m.AnomalyDetectionConfig,
			"object_id":                m.ObjectId,
			"object_type":              m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QualityMonitor) Type(ctx context.Context) attr.Type {
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
func (m *QualityMonitor) GetAnomalyDetectionConfig(ctx context.Context) (AnomalyDetectionConfig, bool) {
	var e AnomalyDetectionConfig
	if m.AnomalyDetectionConfig.IsNull() || m.AnomalyDetectionConfig.IsUnknown() {
		return e, false
	}
	var v AnomalyDetectionConfig
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
func (m *QualityMonitor) SetAnomalyDetectionConfig(ctx context.Context, v AnomalyDetectionConfig) {
	vs := v.ToObjectValue(ctx)
	m.AnomalyDetectionConfig = vs
}

type UpdateQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType types.String `tfsdk:"-"`

	QualityMonitor types.Object `tfsdk:"quality_monitor"`
}

func (to *UpdateQualityMonitorRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateQualityMonitorRequest) {
	if !from.QualityMonitor.IsNull() && !from.QualityMonitor.IsUnknown() {
		if toQualityMonitor, ok := to.GetQualityMonitor(ctx); ok {
			if fromQualityMonitor, ok := from.GetQualityMonitor(ctx); ok {
				// Recursively sync the fields of QualityMonitor
				toQualityMonitor.SyncFieldsDuringCreateOrUpdate(ctx, fromQualityMonitor)
				to.SetQualityMonitor(ctx, toQualityMonitor)
			}
		}
	}
}

func (to *UpdateQualityMonitorRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateQualityMonitorRequest) {
	if !from.QualityMonitor.IsNull() && !from.QualityMonitor.IsUnknown() {
		if toQualityMonitor, ok := to.GetQualityMonitor(ctx); ok {
			if fromQualityMonitor, ok := from.GetQualityMonitor(ctx); ok {
				toQualityMonitor.SyncFieldsDuringRead(ctx, fromQualityMonitor)
				to.SetQualityMonitor(ctx, toQualityMonitor)
			}
		}
	}
}

func (m UpdateQualityMonitorRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quality_monitor"] = attrs["quality_monitor"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitor": reflect.TypeOf(QualityMonitor{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (m UpdateQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id":       m.ObjectId,
			"object_type":     m.ObjectType,
			"quality_monitor": m.QualityMonitor,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateQualityMonitorRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateQualityMonitorRequest) GetQualityMonitor(ctx context.Context) (QualityMonitor, bool) {
	var e QualityMonitor
	if m.QualityMonitor.IsNull() || m.QualityMonitor.IsUnknown() {
		return e, false
	}
	var v QualityMonitor
	d := m.QualityMonitor.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQualityMonitor sets the value of the QualityMonitor field in UpdateQualityMonitorRequest.
func (m *UpdateQualityMonitorRequest) SetQualityMonitor(ctx context.Context, v QualityMonitor) {
	vs := v.ToObjectValue(ctx)
	m.QualityMonitor = vs
}
