// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package pipelines_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names" tf:"optional"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment types.Object `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`

	DryRun types.Bool `tfsdk:"dry_run" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition" tf:"optional,object"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Restart window of this pipeline.
	RestartWindow types.Object `tfsdk:"restart_window" tf:"optional,object"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.Object `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *CreatePipeline) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipeline) {
}

func (newState *CreatePipeline) SyncEffectiveFieldsDuringRead(existingState CreatePipeline) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePipeline.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePipeline) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"trigger":              reflect.TypeOf(PipelineTrigger{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipeline
// only implements ToObjectValue() and Type().
func (o CreatePipeline) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_duplicate_names": o.AllowDuplicateNames,
			"budget_policy_id":      o.BudgetPolicyId,
			"catalog":               o.Catalog,
			"channel":               o.Channel,
			"clusters":              o.Clusters,
			"configuration":         o.Configuration,
			"continuous":            o.Continuous,
			"deployment":            o.Deployment,
			"development":           o.Development,
			"dry_run":               o.DryRun,
			"edition":               o.Edition,
			"filters":               o.Filters,
			"gateway_definition":    o.GatewayDefinition,
			"id":                    o.Id,
			"ingestion_definition":  o.IngestionDefinition,
			"libraries":             o.Libraries,
			"name":                  o.Name,
			"notifications":         o.Notifications,
			"photon":                o.Photon,
			"restart_window":        o.RestartWindow,
			"schema":                o.Schema,
			"serverless":            o.Serverless,
			"storage":               o.Storage,
			"target":                o.Target,
			"trigger":               o.Trigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePipeline) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_duplicate_names": types.BoolType,
			"budget_policy_id":      types.StringType,
			"catalog":               types.StringType,
			"channel":               types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous":           types.BoolType,
			"deployment":           PipelineDeployment{}.Type(ctx),
			"development":          types.BoolType,
			"dry_run":              types.BoolType,
			"edition":              types.StringType,
			"filters":              Filters{}.Type(ctx),
			"gateway_definition":   IngestionGatewayPipelineDefinition{}.Type(ctx),
			"id":                   types.StringType,
			"ingestion_definition": IngestionPipelineDefinition{}.Type(ctx),
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon":         types.BoolType,
			"restart_window": RestartWindow{}.Type(ctx),
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"target":         types.StringType,
			"trigger":        PipelineTrigger{}.Type(ctx),
		},
	}
}

// GetClusters returns the value of the Clusters field in CreatePipeline as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in CreatePipeline.
func (o *CreatePipeline) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in CreatePipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if o.Configuration.IsNull() || o.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in CreatePipeline.
func (o *CreatePipeline) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in CreatePipeline as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDeployment sets the value of the Deployment field in CreatePipeline.
func (o *CreatePipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetFilters returns the value of the Filters field in CreatePipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFilters sets the value of the Filters field in CreatePipeline.
func (o *CreatePipeline) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	o.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in CreatePipeline as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGatewayDefinition sets the value of the GatewayDefinition field in CreatePipeline.
func (o *CreatePipeline) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in CreatePipeline as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIngestionDefinition sets the value of the IngestionDefinition field in CreatePipeline.
func (o *CreatePipeline) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in CreatePipeline as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePipeline.
func (o *CreatePipeline) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in CreatePipeline as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in CreatePipeline.
func (o *CreatePipeline) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in CreatePipeline as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRestartWindow sets the value of the RestartWindow field in CreatePipeline.
func (o *CreatePipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetTrigger returns the value of the Trigger field in CreatePipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTrigger sets the value of the Trigger field in CreatePipeline.
func (o *CreatePipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings types.Object `tfsdk:"effective_settings" tf:"optional,object"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
}

func (newState *CreatePipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipelineResponse) {
}

func (newState *CreatePipelineResponse) SyncEffectiveFieldsDuringRead(existingState CreatePipelineResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(PipelineSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipelineResponse
// only implements ToObjectValue() and Type().
func (o CreatePipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_settings": o.EffectiveSettings,
			"pipeline_id":        o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_settings": PipelineSpec{}.Type(ctx),
			"pipeline_id":        types.StringType,
		},
	}
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in CreatePipelineResponse as
// a PipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipelineResponse) GetEffectiveSettings(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if o.EffectiveSettings.IsNull() || o.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec
	d := o.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectiveSettings sets the value of the EffectiveSettings field in CreatePipelineResponse.
func (o *CreatePipelineResponse) SetEffectiveSettings(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveSettings = vs
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule" tf:"optional"`

	TimezoneId types.String `tfsdk:"timezone_id" tf:"optional"`
}

func (newState *CronTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronTrigger) {
}

func (newState *CronTrigger) SyncEffectiveFieldsDuringRead(existingState CronTrigger) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronTrigger
// only implements ToObjectValue() and Type().
func (o CronTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_schedule": o.QuartzCronSchedule,
			"timezone_id":          o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronTrigger) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quartz_cron_schedule": types.StringType,
			"timezone_id":          types.StringType,
		},
	}
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	Instance types.String `tfsdk:"instance" tf:"optional"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no" tf:"optional"`
}

func (newState *DataPlaneId) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneId) {
}

func (newState *DataPlaneId) SyncEffectiveFieldsDuringRead(existingState DataPlaneId) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataPlaneId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataPlaneId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneId
// only implements ToObjectValue() and Type().
func (o DataPlaneId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance": o.Instance,
			"seq_no":   o.SeqNo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataPlaneId) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance": types.StringType,
			"seq_no":   types.Int64Type,
		},
	}
}

// Delete a pipeline
type DeletePipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *DeletePipelineRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePipelineRequest) {
}

func (newState *DeletePipelineRequest) SyncEffectiveFieldsDuringRead(existingState DeletePipelineRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePipelineRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineRequest
// only implements ToObjectValue() and Type().
func (o DeletePipelineRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePipelineRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type DeletePipelineResponse struct {
}

func (newState *DeletePipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePipelineResponse) {
}

func (newState *DeletePipelineResponse) SyncEffectiveFieldsDuringRead(existingState DeletePipelineResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineResponse
// only implements ToObjectValue() and Type().
func (o DeletePipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names" tf:"optional"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment types.Object `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition" tf:"optional,object"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Unique identifier for this pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// Restart window of this pipeline.
	RestartWindow types.Object `tfsdk:"restart_window" tf:"optional,object"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.Object `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *EditPipeline) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipeline) {
}

func (newState *EditPipeline) SyncEffectiveFieldsDuringRead(existingState EditPipeline) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipeline.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPipeline) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"trigger":              reflect.TypeOf(PipelineTrigger{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipeline
// only implements ToObjectValue() and Type().
func (o EditPipeline) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_duplicate_names":  o.AllowDuplicateNames,
			"budget_policy_id":       o.BudgetPolicyId,
			"catalog":                o.Catalog,
			"channel":                o.Channel,
			"clusters":               o.Clusters,
			"configuration":          o.Configuration,
			"continuous":             o.Continuous,
			"deployment":             o.Deployment,
			"development":            o.Development,
			"edition":                o.Edition,
			"expected_last_modified": o.ExpectedLastModified,
			"filters":                o.Filters,
			"gateway_definition":     o.GatewayDefinition,
			"id":                     o.Id,
			"ingestion_definition":   o.IngestionDefinition,
			"libraries":              o.Libraries,
			"name":                   o.Name,
			"notifications":          o.Notifications,
			"photon":                 o.Photon,
			"pipeline_id":            o.PipelineId,
			"restart_window":         o.RestartWindow,
			"schema":                 o.Schema,
			"serverless":             o.Serverless,
			"storage":                o.Storage,
			"target":                 o.Target,
			"trigger":                o.Trigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EditPipeline) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_duplicate_names": types.BoolType,
			"budget_policy_id":      types.StringType,
			"catalog":               types.StringType,
			"channel":               types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous":             types.BoolType,
			"deployment":             PipelineDeployment{}.Type(ctx),
			"development":            types.BoolType,
			"edition":                types.StringType,
			"expected_last_modified": types.Int64Type,
			"filters":                Filters{}.Type(ctx),
			"gateway_definition":     IngestionGatewayPipelineDefinition{}.Type(ctx),
			"id":                     types.StringType,
			"ingestion_definition":   IngestionPipelineDefinition{}.Type(ctx),
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon":         types.BoolType,
			"pipeline_id":    types.StringType,
			"restart_window": RestartWindow{}.Type(ctx),
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"target":         types.StringType,
			"trigger":        PipelineTrigger{}.Type(ctx),
		},
	}
}

// GetClusters returns the value of the Clusters field in EditPipeline as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in EditPipeline.
func (o *EditPipeline) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in EditPipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if o.Configuration.IsNull() || o.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in EditPipeline.
func (o *EditPipeline) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in EditPipeline as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDeployment sets the value of the Deployment field in EditPipeline.
func (o *EditPipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetFilters returns the value of the Filters field in EditPipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFilters sets the value of the Filters field in EditPipeline.
func (o *EditPipeline) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	o.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in EditPipeline as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGatewayDefinition sets the value of the GatewayDefinition field in EditPipeline.
func (o *EditPipeline) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in EditPipeline as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIngestionDefinition sets the value of the IngestionDefinition field in EditPipeline.
func (o *EditPipeline) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in EditPipeline as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPipeline.
func (o *EditPipeline) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in EditPipeline as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in EditPipeline.
func (o *EditPipeline) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in EditPipeline as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRestartWindow sets the value of the RestartWindow field in EditPipeline.
func (o *EditPipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetTrigger returns the value of the Trigger field in EditPipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTrigger sets the value of the Trigger field in EditPipeline.
func (o *EditPipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type EditPipelineResponse struct {
}

func (newState *EditPipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipelineResponse) {
}

func (newState *EditPipelineResponse) SyncEffectiveFieldsDuringRead(existingState EditPipelineResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipelineResponse
// only implements ToObjectValue() and Type().
func (o EditPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditPipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions types.List `tfsdk:"exceptions" tf:"optional"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal" tf:"optional"`
}

func (newState *ErrorDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan ErrorDetail) {
}

func (newState *ErrorDetail) SyncEffectiveFieldsDuringRead(existingState ErrorDetail) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ErrorDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ErrorDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exceptions": reflect.TypeOf(SerializedException{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ErrorDetail
// only implements ToObjectValue() and Type().
func (o ErrorDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exceptions": o.Exceptions,
			"fatal":      o.Fatal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ErrorDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exceptions": basetypes.ListType{
				ElemType: SerializedException{}.Type(ctx),
			},
			"fatal": types.BoolType,
		},
	}
}

// GetExceptions returns the value of the Exceptions field in ErrorDetail as
// a slice of SerializedException values.
// If the field is unknown or null, the boolean return value is false.
func (o *ErrorDetail) GetExceptions(ctx context.Context) ([]SerializedException, bool) {
	if o.Exceptions.IsNull() || o.Exceptions.IsUnknown() {
		return nil, false
	}
	var v []SerializedException
	d := o.Exceptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExceptions sets the value of the Exceptions field in ErrorDetail.
func (o *ErrorDetail) SetExceptions(ctx context.Context, v []SerializedException) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exceptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exceptions = types.ListValueMust(t, vs)
}

type FileLibrary struct {
	// The absolute path of the file.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *FileLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileLibrary) {
}

func (newState *FileLibrary) SyncEffectiveFieldsDuringRead(existingState FileLibrary) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileLibrary
// only implements ToObjectValue() and Type().
func (o FileLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type Filters struct {
	// Paths to exclude.
	Exclude types.List `tfsdk:"exclude" tf:"optional"`
	// Paths to include.
	Include types.List `tfsdk:"include" tf:"optional"`
}

func (newState *Filters) SyncEffectiveFieldsDuringCreateOrUpdate(plan Filters) {
}

func (newState *Filters) SyncEffectiveFieldsDuringRead(existingState Filters) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Filters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Filters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude": reflect.TypeOf(types.String{}),
		"include": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filters
// only implements ToObjectValue() and Type().
func (o Filters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude": o.Exclude,
			"include": o.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Filters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exclude": basetypes.ListType{
				ElemType: types.StringType,
			},
			"include": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetExclude returns the value of the Exclude field in Filters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Filters) GetExclude(ctx context.Context) ([]types.String, bool) {
	if o.Exclude.IsNull() || o.Exclude.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Exclude.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExclude sets the value of the Exclude field in Filters.
func (o *Filters) SetExclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exclude = types.ListValueMust(t, vs)
}

// GetInclude returns the value of the Include field in Filters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Filters) GetInclude(ctx context.Context) ([]types.String, bool) {
	if o.Include.IsNull() || o.Include.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Include.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInclude sets the value of the Include field in Filters.
func (o *Filters) SetInclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Include = types.ListValueMust(t, vs)
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelinePermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionLevelsRequest) {
}

func (newState *GetPipelinePermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionLevelsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPipelinePermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionLevelsResponse) {
}

func (newState *GetPipelinePermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PipelinePermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PipelinePermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse as
// a slice of PipelinePermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelinePermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]PipelinePermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse.
func (o *GetPipelinePermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []PipelinePermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get pipeline permissions
type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelinePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionsRequest) {
}

func (newState *GetPipelinePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

// Get a pipeline
type GetPipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelineRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelineRequest) {
}

func (newState *GetPipelineRequest) SyncEffectiveFieldsDuringRead(existingState GetPipelineRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelineRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelineRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineRequest
// only implements ToObjectValue() and Type().
func (o GetPipelineRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelineRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"optional"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health" tf:"optional"`
	// The last time the pipeline settings were modified or created.
	LastModified types.Int64 `tfsdk:"last_modified" tf:"optional"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates types.List `tfsdk:"latest_updates" tf:"optional"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name" tf:"optional"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec types.Object `tfsdk:"spec" tf:"optional,object"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *GetPipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelineResponse) {
}

func (newState *GetPipelineResponse) SyncEffectiveFieldsDuringRead(existingState GetPipelineResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
		"spec":           reflect.TypeOf(PipelineSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineResponse
// only implements ToObjectValue() and Type().
func (o GetPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                      o.Cause,
			"cluster_id":                 o.ClusterId,
			"creator_user_name":          o.CreatorUserName,
			"effective_budget_policy_id": o.EffectiveBudgetPolicyId,
			"health":                     o.Health,
			"last_modified":              o.LastModified,
			"latest_updates":             o.LatestUpdates,
			"name":                       o.Name,
			"pipeline_id":                o.PipelineId,
			"run_as_user_name":           o.RunAsUserName,
			"spec":                       o.Spec,
			"state":                      o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":                      types.StringType,
			"cluster_id":                 types.StringType,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"health":                     types.StringType,
			"last_modified":              types.Int64Type,
			"latest_updates": basetypes.ListType{
				ElemType: UpdateStateInfo{}.Type(ctx),
			},
			"name":             types.StringType,
			"pipeline_id":      types.StringType,
			"run_as_user_name": types.StringType,
			"spec":             PipelineSpec{}.Type(ctx),
			"state":            types.StringType,
		},
	}
}

// GetLatestUpdates returns the value of the LatestUpdates field in GetPipelineResponse as
// a slice of UpdateStateInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo, bool) {
	if o.LatestUpdates.IsNull() || o.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo
	d := o.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in GetPipelineResponse.
func (o *GetPipelineResponse) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestUpdates = types.ListValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in GetPipelineResponse as
// a PipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse) GetSpec(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSpec sets the value of the Spec field in GetPipelineResponse.
func (o *GetPipelineResponse) SetSpec(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

// Get a pipeline update
type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
}

func (newState *GetUpdateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUpdateRequest) {
}

func (newState *GetUpdateRequest) SyncEffectiveFieldsDuringRead(existingState GetUpdateRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateRequest
// only implements ToObjectValue() and Type().
func (o GetUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
			"update_id":   o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetUpdateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
			"update_id":   types.StringType,
		},
	}
}

type GetUpdateResponse struct {
	// The current update info.
	Update types.Object `tfsdk:"update" tf:"optional,object"`
}

func (newState *GetUpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUpdateResponse) {
}

func (newState *GetUpdateResponse) SyncEffectiveFieldsDuringRead(existingState GetUpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"update": reflect.TypeOf(UpdateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateResponse
// only implements ToObjectValue() and Type().
func (o GetUpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update": o.Update,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetUpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update": UpdateInfo{}.Type(ctx),
		},
	}
}

// GetUpdate returns the value of the Update field in GetUpdateResponse as
// a UpdateInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetUpdateResponse) GetUpdate(ctx context.Context) (UpdateInfo, bool) {
	var e UpdateInfo
	if o.Update.IsNull() || o.Update.IsUnknown() {
		return e, false
	}
	var v []UpdateInfo
	d := o.Update.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetUpdate sets the value of the Update field in GetUpdateResponse.
func (o *GetUpdateResponse) SetUpdate(ctx context.Context, v UpdateInfo) {
	vs := v.ToObjectValue(ctx)
	o.Update = vs
}

type IngestionConfig struct {
	// Select a specific source report.
	Report types.Object `tfsdk:"report" tf:"optional,object"`
	// Select all tables from a specific source schema.
	Schema types.Object `tfsdk:"schema" tf:"optional,object"`
	// Select a specific source table.
	Table types.Object `tfsdk:"table" tf:"optional,object"`
}

func (newState *IngestionConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionConfig) {
}

func (newState *IngestionConfig) SyncEffectiveFieldsDuringRead(existingState IngestionConfig) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"report": reflect.TypeOf(ReportSpec{}),
		"schema": reflect.TypeOf(SchemaSpec{}),
		"table":  reflect.TypeOf(TableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionConfig
// only implements ToObjectValue() and Type().
func (o IngestionConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"report": o.Report,
			"schema": o.Schema,
			"table":  o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"report": ReportSpec{}.Type(ctx),
			"schema": SchemaSpec{}.Type(ctx),
			"table":  TableSpec{}.Type(ctx),
		},
	}
}

// GetReport returns the value of the Report field in IngestionConfig as
// a ReportSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig) GetReport(ctx context.Context) (ReportSpec, bool) {
	var e ReportSpec
	if o.Report.IsNull() || o.Report.IsUnknown() {
		return e, false
	}
	var v []ReportSpec
	d := o.Report.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetReport sets the value of the Report field in IngestionConfig.
func (o *IngestionConfig) SetReport(ctx context.Context, v ReportSpec) {
	vs := v.ToObjectValue(ctx)
	o.Report = vs
}

// GetSchema returns the value of the Schema field in IngestionConfig as
// a SchemaSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig) GetSchema(ctx context.Context) (SchemaSpec, bool) {
	var e SchemaSpec
	if o.Schema.IsNull() || o.Schema.IsUnknown() {
		return e, false
	}
	var v []SchemaSpec
	d := o.Schema.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSchema sets the value of the Schema field in IngestionConfig.
func (o *IngestionConfig) SetSchema(ctx context.Context, v SchemaSpec) {
	vs := v.ToObjectValue(ctx)
	o.Schema = vs
}

// GetTable returns the value of the Table field in IngestionConfig as
// a TableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig) GetTable(ctx context.Context) (TableSpec, bool) {
	var e TableSpec
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []TableSpec
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTable sets the value of the Table field in IngestionConfig.
func (o *IngestionConfig) SetTable(ctx context.Context, v TableSpec) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	ConnectionId types.String `tfsdk:"connection_id" tf:"optional"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog types.String `tfsdk:"gateway_storage_catalog" tf:"optional"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName types.String `tfsdk:"gateway_storage_name" tf:"optional"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema types.String `tfsdk:"gateway_storage_schema" tf:"optional"`
}

func (newState *IngestionGatewayPipelineDefinition) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionGatewayPipelineDefinition) {
}

func (newState *IngestionGatewayPipelineDefinition) SyncEffectiveFieldsDuringRead(existingState IngestionGatewayPipelineDefinition) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionGatewayPipelineDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionGatewayPipelineDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionGatewayPipelineDefinition
// only implements ToObjectValue() and Type().
func (o IngestionGatewayPipelineDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_id":           o.ConnectionId,
			"connection_name":         o.ConnectionName,
			"gateway_storage_catalog": o.GatewayStorageCatalog,
			"gateway_storage_name":    o.GatewayStorageName,
			"gateway_storage_schema":  o.GatewayStorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionGatewayPipelineDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_id":           types.StringType,
			"connection_name":         types.StringType,
			"gateway_storage_catalog": types.StringType,
			"gateway_storage_name":    types.StringType,
			"gateway_storage_schema":  types.StringType,
		},
	}
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	IngestionGatewayId types.String `tfsdk:"ingestion_gateway_id" tf:"optional"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects types.List `tfsdk:"objects" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration types.Object `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *IngestionPipelineDefinition) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionPipelineDefinition) {
}

func (newState *IngestionPipelineDefinition) SyncEffectiveFieldsDuringRead(existingState IngestionPipelineDefinition) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionPipelineDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionPipelineDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects":             reflect.TypeOf(IngestionConfig{}),
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name":      o.ConnectionName,
			"ingestion_gateway_id": o.IngestionGatewayId,
			"objects":              o.Objects,
			"table_configuration":  o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":      types.StringType,
			"ingestion_gateway_id": types.StringType,
			"objects": basetypes.ListType{
				ElemType: IngestionConfig{}.Type(ctx),
			},
			"table_configuration": TableSpecificConfig{}.Type(ctx),
		},
	}
}

// GetObjects returns the value of the Objects field in IngestionPipelineDefinition as
// a slice of IngestionConfig values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition) GetObjects(ctx context.Context) ([]IngestionConfig, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []IngestionConfig
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in IngestionPipelineDefinition.
func (o *IngestionPipelineDefinition) SetObjects(ctx context.Context, v []IngestionConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

// GetTableConfiguration returns the value of the TableConfiguration field in IngestionPipelineDefinition as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTableConfiguration sets the value of the TableConfiguration field in IngestionPipelineDefinition.
func (o *IngestionPipelineDefinition) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

// List pipeline events
type ListPipelineEventsRequest struct {
	// Criteria to select a subset of results, expressed using a SQL-like
	// syntax. The supported filters are: 1. level='INFO' (or WARN or ERROR) 2.
	// level in ('INFO', 'WARN') 3. id='[event-id]' 4. timestamp > 'TIMESTAMP'
	// (or >=,<,<=,=)
	//
	// Composite expressions are supported, for example: level in ('ERROR',
	// 'WARN') AND timestamp> '2021-07-22T06:37:33.083Z'
	Filter types.String `tfsdk:"-"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults types.Int64 `tfsdk:"-"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy types.List `tfsdk:"-"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken types.String `tfsdk:"-"`

	PipelineId types.String `tfsdk:"-"`
}

func (newState *ListPipelineEventsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelineEventsRequest) {
}

func (newState *ListPipelineEventsRequest) SyncEffectiveFieldsDuringRead(existingState ListPipelineEventsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelineEventsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelineEventsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsRequest
// only implements ToObjectValue() and Type().
func (o ListPipelineEventsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      o.Filter,
			"max_results": o.MaxResults,
			"order_by":    o.OrderBy,
			"page_token":  o.PageToken,
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelineEventsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token":  types.StringType,
			"pipeline_id": types.StringType,
		},
	}
}

// GetOrderBy returns the value of the OrderBy field in ListPipelineEventsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelineEventsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in ListPipelineEventsRequest.
func (o *ListPipelineEventsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	Events types.List `tfsdk:"events" tf:"optional"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListPipelineEventsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelineEventsResponse) {
}

func (newState *ListPipelineEventsResponse) SyncEffectiveFieldsDuringRead(existingState ListPipelineEventsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelineEventsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelineEventsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(PipelineEvent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsResponse
// only implements ToObjectValue() and Type().
func (o ListPipelineEventsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          o.Events,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelineEventsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: PipelineEvent{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in ListPipelineEventsResponse as
// a slice of PipelineEvent values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelineEventsResponse) GetEvents(ctx context.Context) ([]PipelineEvent, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []PipelineEvent
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListPipelineEventsResponse.
func (o *ListPipelineEventsResponse) SetEvents(ctx context.Context, v []PipelineEvent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// List pipelines
type ListPipelinesRequest struct {
	// Select a subset of results based on the specified criteria. The supported
	// filters are:
	//
	// * `notebook='<path>'` to select pipelines that reference the provided
	// notebook path. * `name LIKE '[pattern]'` to select pipelines with a name
	// that matches pattern. Wildcards are supported, for example: `name LIKE
	// '%shopping%'`
	//
	// Composite filters are not supported. This field is optional.
	Filter types.String `tfsdk:"-"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults types.Int64 `tfsdk:"-"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy types.List `tfsdk:"-"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListPipelinesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelinesRequest) {
}

func (newState *ListPipelinesRequest) SyncEffectiveFieldsDuringRead(existingState ListPipelinesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelinesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelinesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesRequest
// only implements ToObjectValue() and Type().
func (o ListPipelinesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      o.Filter,
			"max_results": o.MaxResults,
			"order_by":    o.OrderBy,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelinesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token": types.StringType,
		},
	}
}

// GetOrderBy returns the value of the OrderBy field in ListPipelinesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelinesRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in ListPipelinesRequest.
func (o *ListPipelinesRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// The list of events matching the request criteria.
	Statuses types.List `tfsdk:"statuses" tf:"optional"`
}

func (newState *ListPipelinesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelinesResponse) {
}

func (newState *ListPipelinesResponse) SyncEffectiveFieldsDuringRead(existingState ListPipelinesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelinesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelinesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(PipelineStateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesResponse
// only implements ToObjectValue() and Type().
func (o ListPipelinesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"statuses":        o.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelinesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"statuses": basetypes.ListType{
				ElemType: PipelineStateInfo{}.Type(ctx),
			},
		},
	}
}

// GetStatuses returns the value of the Statuses field in ListPipelinesResponse as
// a slice of PipelineStateInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelinesResponse) GetStatuses(ctx context.Context) ([]PipelineStateInfo, bool) {
	if o.Statuses.IsNull() || o.Statuses.IsUnknown() {
		return nil, false
	}
	var v []PipelineStateInfo
	d := o.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListPipelinesResponse.
func (o *ListPipelinesResponse) SetStatuses(ctx context.Context, v []PipelineStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Statuses = types.ListValueMust(t, vs)
}

// List pipeline updates
type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults types.Int64 `tfsdk:"-"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-"`
	// The pipeline to return updates for.
	PipelineId types.String `tfsdk:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId types.String `tfsdk:"-"`
}

func (newState *ListUpdatesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUpdatesRequest) {
}

func (newState *ListUpdatesRequest) SyncEffectiveFieldsDuringRead(existingState ListUpdatesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUpdatesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUpdatesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesRequest
// only implements ToObjectValue() and Type().
func (o ListUpdatesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":     o.MaxResults,
			"page_token":      o.PageToken,
			"pipeline_id":     o.PipelineId,
			"until_update_id": o.UntilUpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListUpdatesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":     types.Int64Type,
			"page_token":      types.StringType,
			"pipeline_id":     types.StringType,
			"until_update_id": types.StringType,
		},
	}
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`

	Updates types.List `tfsdk:"updates" tf:"optional"`
}

func (newState *ListUpdatesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUpdatesResponse) {
}

func (newState *ListUpdatesResponse) SyncEffectiveFieldsDuringRead(existingState ListUpdatesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUpdatesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUpdatesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(UpdateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesResponse
// only implements ToObjectValue() and Type().
func (o ListUpdatesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
			"updates":         o.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListUpdatesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"updates": basetypes.ListType{
				ElemType: UpdateInfo{}.Type(ctx),
			},
		},
	}
}

// GetUpdates returns the value of the Updates field in ListUpdatesResponse as
// a slice of UpdateInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUpdatesResponse) GetUpdates(ctx context.Context) ([]UpdateInfo, bool) {
	if o.Updates.IsNull() || o.Updates.IsUnknown() {
		return nil, false
	}
	var v []UpdateInfo
	d := o.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in ListUpdatesResponse.
func (o *ListUpdatesResponse) SetUpdates(ctx context.Context, v []UpdateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Updates = types.ListValueMust(t, vs)
}

type ManualTrigger struct {
}

func (newState *ManualTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan ManualTrigger) {
}

func (newState *ManualTrigger) SyncEffectiveFieldsDuringRead(existingState ManualTrigger) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManualTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ManualTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManualTrigger
// only implements ToObjectValue() and Type().
func (o ManualTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ManualTrigger) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *NotebookLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookLibrary) {
}

func (newState *NotebookLibrary) SyncEffectiveFieldsDuringRead(existingState NotebookLibrary) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookLibrary
// only implements ToObjectValue() and Type().
func (o NotebookLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	Alerts types.List `tfsdk:"alerts" tf:"optional"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients types.List `tfsdk:"email_recipients" tf:"optional"`
}

func (newState *Notifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan Notifications) {
}

func (newState *Notifications) SyncEffectiveFieldsDuringRead(existingState Notifications) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Notifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Notifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":           reflect.TypeOf(types.String{}),
		"email_recipients": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Notifications
// only implements ToObjectValue() and Type().
func (o Notifications) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":           o.Alerts,
			"email_recipients": o.EmailRecipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Notifications) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alerts": basetypes.ListType{
				ElemType: types.StringType,
			},
			"email_recipients": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAlerts returns the value of the Alerts field in Notifications as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Notifications) GetAlerts(ctx context.Context) ([]types.String, bool) {
	if o.Alerts.IsNull() || o.Alerts.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Alerts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlerts sets the value of the Alerts field in Notifications.
func (o *Notifications) SetAlerts(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Alerts = types.ListValueMust(t, vs)
}

// GetEmailRecipients returns the value of the EmailRecipients field in Notifications as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Notifications) GetEmailRecipients(ctx context.Context) ([]types.String, bool) {
	if o.EmailRecipients.IsNull() || o.EmailRecipients.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.EmailRecipients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailRecipients sets the value of the EmailRecipients field in Notifications.
func (o *Notifications) SetEmailRecipients(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmailRecipients = types.ListValueMust(t, vs)
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId types.Int64 `tfsdk:"batch_id" tf:"optional"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName types.String `tfsdk:"dataset_name" tf:"optional"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId types.String `tfsdk:"flow_id" tf:"optional"`
	// The name of the flow. Not unique.
	FlowName types.String `tfsdk:"flow_name" tf:"optional"`
	// The optional host name where the event was triggered
	Host types.String `tfsdk:"host" tf:"optional"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId types.String `tfsdk:"maintenance_id" tf:"optional"`
	// Materialization name.
	MaterializationName types.String `tfsdk:"materialization_name" tf:"optional"`
	// The org id of the user. Unique within a cloud.
	OrgId types.Int64 `tfsdk:"org_id" tf:"optional"`
	// The id of the pipeline. Globally unique.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// The name of the pipeline. Not unique.
	PipelineName types.String `tfsdk:"pipeline_name" tf:"optional"`
	// The cloud region.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The id of the request that caused an update.
	RequestId types.String `tfsdk:"request_id" tf:"optional"`
	// The id of a (delta) table. Globally unique.
	TableId types.String `tfsdk:"table_id" tf:"optional"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId types.String `tfsdk:"uc_resource_id" tf:"optional"`
	// The id of an execution. Globally unique.
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *Origin) SyncEffectiveFieldsDuringCreateOrUpdate(plan Origin) {
}

func (newState *Origin) SyncEffectiveFieldsDuringRead(existingState Origin) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Origin.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Origin) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Origin
// only implements ToObjectValue() and Type().
func (o Origin) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"batch_id":             o.BatchId,
			"cloud":                o.Cloud,
			"cluster_id":           o.ClusterId,
			"dataset_name":         o.DatasetName,
			"flow_id":              o.FlowId,
			"flow_name":            o.FlowName,
			"host":                 o.Host,
			"maintenance_id":       o.MaintenanceId,
			"materialization_name": o.MaterializationName,
			"org_id":               o.OrgId,
			"pipeline_id":          o.PipelineId,
			"pipeline_name":        o.PipelineName,
			"region":               o.Region,
			"request_id":           o.RequestId,
			"table_id":             o.TableId,
			"uc_resource_id":       o.UcResourceId,
			"update_id":            o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Origin) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"batch_id":             types.Int64Type,
			"cloud":                types.StringType,
			"cluster_id":           types.StringType,
			"dataset_name":         types.StringType,
			"flow_id":              types.StringType,
			"flow_name":            types.StringType,
			"host":                 types.StringType,
			"maintenance_id":       types.StringType,
			"materialization_name": types.StringType,
			"org_id":               types.Int64Type,
			"pipeline_id":          types.StringType,
			"pipeline_name":        types.StringType,
			"region":               types.StringType,
			"request_id":           types.StringType,
			"table_id":             types.StringType,
			"uc_resource_id":       types.StringType,
			"update_id":            types.StringType,
		},
	}
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PipelineAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineAccessControlRequest) {
}

func (newState *PipelineAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState PipelineAccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlRequest
// only implements ToObjectValue() and Type().
func (o PipelineAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PipelineAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PipelineAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineAccessControlResponse) {
}

func (newState *PipelineAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState PipelineAccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PipelinePermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlResponse
// only implements ToObjectValue() and Type().
func (o PipelineAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PipelinePermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in PipelineAccessControlResponse as
// a slice of PipelinePermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineAccessControlResponse) GetAllPermissions(ctx context.Context) ([]PipelinePermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PipelineAccessControlResponse.
func (o *PipelineAccessControlResponse) SetAllPermissions(ctx context.Context, v []PipelinePermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale" tf:"optional,object"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf" tf:"optional,object"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.Object `tfsdk:"gcp_attributes" tf:"optional,object"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label types.String `tfsdk:"label" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf types.Map `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars types.Map `tfsdk:"spark_env_vars" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys" tf:"optional"`
}

func (newState *PipelineCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineCluster) {
}

func (newState *PipelineCluster) SyncEffectiveFieldsDuringRead(existingState PipelineCluster) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(PipelineClusterAutoscale{}),
		"aws_attributes":   reflect.TypeOf(compute_tf.AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(compute_tf.AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(compute_tf.ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"gcp_attributes":   reflect.TypeOf(compute_tf.GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(compute_tf.InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineCluster
// only implements ToObjectValue() and Type().
func (o PipelineCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":  o.ApplyPolicyDefaultValues,
			"autoscale":                    o.Autoscale,
			"aws_attributes":               o.AwsAttributes,
			"azure_attributes":             o.AzureAttributes,
			"cluster_log_conf":             o.ClusterLogConf,
			"custom_tags":                  o.CustomTags,
			"driver_instance_pool_id":      o.DriverInstancePoolId,
			"driver_node_type_id":          o.DriverNodeTypeId,
			"enable_local_disk_encryption": o.EnableLocalDiskEncryption,
			"gcp_attributes":               o.GcpAttributes,
			"init_scripts":                 o.InitScripts,
			"instance_pool_id":             o.InstancePoolId,
			"label":                        o.Label,
			"node_type_id":                 o.NodeTypeId,
			"num_workers":                  o.NumWorkers,
			"policy_id":                    o.PolicyId,
			"spark_conf":                   o.SparkConf,
			"spark_env_vars":               o.SparkEnvVars,
			"ssh_public_keys":              o.SshPublicKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineCluster) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale":                   PipelineClusterAutoscale{}.Type(ctx),
			"aws_attributes":              compute_tf.AwsAttributes{}.Type(ctx),
			"azure_attributes":            compute_tf.AzureAttributes{}.Type(ctx),
			"cluster_log_conf":            compute_tf.ClusterLogConf{}.Type(ctx),
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes":               compute_tf.GcpAttributes{}.Type(ctx),
			"init_scripts": basetypes.ListType{
				ElemType: compute_tf.InitScriptInfo{}.Type(ctx),
			},
			"instance_pool_id": types.StringType,
			"label":            types.StringType,
			"node_type_id":     types.StringType,
			"num_workers":      types.Int64Type,
			"policy_id":        types.StringType,
			"spark_conf": basetypes.MapType{
				ElemType: types.StringType,
			},
			"spark_env_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"ssh_public_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAutoscale returns the value of the Autoscale field in PipelineCluster as
// a PipelineClusterAutoscale value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetAutoscale(ctx context.Context) (PipelineClusterAutoscale, bool) {
	var e PipelineClusterAutoscale
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []PipelineClusterAutoscale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAutoscale sets the value of the Autoscale field in PipelineCluster.
func (o *PipelineCluster) SetAutoscale(ctx context.Context, v PipelineClusterAutoscale) {
	vs := v.ToObjectValue(ctx)
	o.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in PipelineCluster as
// a compute_tf.AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetAwsAttributes(ctx context.Context) (compute_tf.AwsAttributes, bool) {
	var e compute_tf.AwsAttributes
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsAttributes sets the value of the AwsAttributes field in PipelineCluster.
func (o *PipelineCluster) SetAwsAttributes(ctx context.Context, v compute_tf.AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in PipelineCluster as
// a compute_tf.AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetAzureAttributes(ctx context.Context) (compute_tf.AzureAttributes, bool) {
	var e compute_tf.AzureAttributes
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureAttributes sets the value of the AzureAttributes field in PipelineCluster.
func (o *PipelineCluster) SetAzureAttributes(ctx context.Context, v compute_tf.AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	o.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in PipelineCluster as
// a compute_tf.ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetClusterLogConf(ctx context.Context) (compute_tf.ClusterLogConf, bool) {
	var e compute_tf.ClusterLogConf
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetClusterLogConf sets the value of the ClusterLogConf field in PipelineCluster.
func (o *PipelineCluster) SetClusterLogConf(ctx context.Context, v compute_tf.ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	o.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in PipelineCluster.
func (o *PipelineCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in PipelineCluster as
// a compute_tf.GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetGcpAttributes(ctx context.Context) (compute_tf.GcpAttributes, bool) {
	var e compute_tf.GcpAttributes
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGcpAttributes sets the value of the GcpAttributes field in PipelineCluster.
func (o *PipelineCluster) SetGcpAttributes(ctx context.Context, v compute_tf.GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	o.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in PipelineCluster as
// a slice of compute_tf.InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetInitScripts(ctx context.Context) ([]compute_tf.InitScriptInfo, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.InitScriptInfo
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in PipelineCluster.
func (o *PipelineCluster) SetInitScripts(ctx context.Context, v []compute_tf.InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkConf.IsNull() || o.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in PipelineCluster.
func (o *PipelineCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if o.SparkEnvVars.IsNull() || o.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in PipelineCluster.
func (o *PipelineCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in PipelineCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if o.SshPublicKeys.IsNull() || o.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in PipelineCluster.
func (o *PipelineCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers" tf:""`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers" tf:""`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode types.String `tfsdk:"mode" tf:"optional"`
}

func (newState *PipelineClusterAutoscale) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineClusterAutoscale) {
}

func (newState *PipelineClusterAutoscale) SyncEffectiveFieldsDuringRead(existingState PipelineClusterAutoscale) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineClusterAutoscale.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineClusterAutoscale) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineClusterAutoscale
// only implements ToObjectValue() and Type().
func (o PipelineClusterAutoscale) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": o.MaxWorkers,
			"min_workers": o.MinWorkers,
			"mode":        o.Mode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineClusterAutoscale) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_workers": types.Int64Type,
			"min_workers": types.Int64Type,
			"mode":        types.StringType,
		},
	}
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	Kind types.String `tfsdk:"kind" tf:"optional"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path" tf:"optional"`
}

func (newState *PipelineDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineDeployment) {
}

func (newState *PipelineDeployment) SyncEffectiveFieldsDuringRead(existingState PipelineDeployment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineDeployment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineDeployment
// only implements ToObjectValue() and Type().
func (o PipelineDeployment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kind":               o.Kind,
			"metadata_file_path": o.MetadataFilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineDeployment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kind":               types.StringType,
			"metadata_file_path": types.StringType,
		},
	}
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	Error types.Object `tfsdk:"error" tf:"optional,object"`
	// The event type. Should always correspond to the details
	EventType types.String `tfsdk:"event_type" tf:"optional"`
	// A time-based, globally unique id.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The severity level of the event.
	Level types.String `tfsdk:"level" tf:"optional"`
	// Maturity level for event_type.
	MaturityLevel types.String `tfsdk:"maturity_level" tf:"optional"`
	// The display message associated with the event.
	Message types.String `tfsdk:"message" tf:"optional"`
	// Describes where the event originates from.
	Origin types.Object `tfsdk:"origin" tf:"optional,object"`
	// A sequencing object to identify and order events.
	Sequence types.Object `tfsdk:"sequence" tf:"optional,object"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *PipelineEvent) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineEvent) {
}

func (newState *PipelineEvent) SyncEffectiveFieldsDuringRead(existingState PipelineEvent) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineEvent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineEvent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error":    reflect.TypeOf(ErrorDetail{}),
		"origin":   reflect.TypeOf(Origin{}),
		"sequence": reflect.TypeOf(Sequencing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineEvent
// only implements ToObjectValue() and Type().
func (o PipelineEvent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":          o.Error,
			"event_type":     o.EventType,
			"id":             o.Id,
			"level":          o.Level,
			"maturity_level": o.MaturityLevel,
			"message":        o.Message,
			"origin":         o.Origin,
			"sequence":       o.Sequence,
			"timestamp":      o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineEvent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error":          ErrorDetail{}.Type(ctx),
			"event_type":     types.StringType,
			"id":             types.StringType,
			"level":          types.StringType,
			"maturity_level": types.StringType,
			"message":        types.StringType,
			"origin":         Origin{}.Type(ctx),
			"sequence":       Sequencing{}.Type(ctx),
			"timestamp":      types.StringType,
		},
	}
}

// GetError returns the value of the Error field in PipelineEvent as
// a ErrorDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent) GetError(ctx context.Context) (ErrorDetail, bool) {
	var e ErrorDetail
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []ErrorDetail
	d := o.Error.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetError sets the value of the Error field in PipelineEvent.
func (o *PipelineEvent) SetError(ctx context.Context, v ErrorDetail) {
	vs := v.ToObjectValue(ctx)
	o.Error = vs
}

// GetOrigin returns the value of the Origin field in PipelineEvent as
// a Origin value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent) GetOrigin(ctx context.Context) (Origin, bool) {
	var e Origin
	if o.Origin.IsNull() || o.Origin.IsUnknown() {
		return e, false
	}
	var v []Origin
	d := o.Origin.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetOrigin sets the value of the Origin field in PipelineEvent.
func (o *PipelineEvent) SetOrigin(ctx context.Context, v Origin) {
	vs := v.ToObjectValue(ctx)
	o.Origin = vs
}

// GetSequence returns the value of the Sequence field in PipelineEvent as
// a Sequencing value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent) GetSequence(ctx context.Context) (Sequencing, bool) {
	var e Sequencing
	if o.Sequence.IsNull() || o.Sequence.IsUnknown() {
		return e, false
	}
	var v []Sequencing
	d := o.Sequence.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSequence sets the value of the Sequence field in PipelineEvent.
func (o *PipelineEvent) SetSequence(ctx context.Context, v Sequencing) {
	vs := v.ToObjectValue(ctx)
	o.Sequence = vs
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File types.Object `tfsdk:"file" tf:"optional,object"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed.
	Maven types.Object `tfsdk:"maven" tf:"optional,object"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook types.Object `tfsdk:"notebook" tf:"optional,object"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

func (newState *PipelineLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineLibrary) {
}

func (newState *PipelineLibrary) SyncEffectiveFieldsDuringRead(existingState PipelineLibrary) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file":     reflect.TypeOf(FileLibrary{}),
		"maven":    reflect.TypeOf(compute_tf.MavenLibrary{}),
		"notebook": reflect.TypeOf(NotebookLibrary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineLibrary
// only implements ToObjectValue() and Type().
func (o PipelineLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file":     o.File,
			"jar":      o.Jar,
			"maven":    o.Maven,
			"notebook": o.Notebook,
			"whl":      o.Whl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineLibrary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file":     FileLibrary{}.Type(ctx),
			"jar":      types.StringType,
			"maven":    compute_tf.MavenLibrary{}.Type(ctx),
			"notebook": NotebookLibrary{}.Type(ctx),
			"whl":      types.StringType,
		},
	}
}

// GetFile returns the value of the File field in PipelineLibrary as
// a FileLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary) GetFile(ctx context.Context) (FileLibrary, bool) {
	var e FileLibrary
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []FileLibrary
	d := o.File.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFile sets the value of the File field in PipelineLibrary.
func (o *PipelineLibrary) SetFile(ctx context.Context, v FileLibrary) {
	vs := v.ToObjectValue(ctx)
	o.File = vs
}

// GetMaven returns the value of the Maven field in PipelineLibrary as
// a compute_tf.MavenLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary) GetMaven(ctx context.Context) (compute_tf.MavenLibrary, bool) {
	var e compute_tf.MavenLibrary
	if o.Maven.IsNull() || o.Maven.IsUnknown() {
		return e, false
	}
	var v []compute_tf.MavenLibrary
	d := o.Maven.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMaven sets the value of the Maven field in PipelineLibrary.
func (o *PipelineLibrary) SetMaven(ctx context.Context, v compute_tf.MavenLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Maven = vs
}

// GetNotebook returns the value of the Notebook field in PipelineLibrary as
// a NotebookLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary) GetNotebook(ctx context.Context) (NotebookLibrary, bool) {
	var e NotebookLibrary
	if o.Notebook.IsNull() || o.Notebook.IsUnknown() {
		return e, false
	}
	var v []NotebookLibrary
	d := o.Notebook.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNotebook sets the value of the Notebook field in PipelineLibrary.
func (o *PipelineLibrary) SetNotebook(ctx context.Context, v NotebookLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Notebook = vs
}

type PipelinePermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PipelinePermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermission) {
}

func (newState *PipelinePermission) SyncEffectiveFieldsDuringRead(existingState PipelinePermission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermission
// only implements ToObjectValue() and Type().
func (o PipelinePermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in PipelinePermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in PipelinePermission.
func (o *PipelinePermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type PipelinePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *PipelinePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissions) {
}

func (newState *PipelinePermissions) SyncEffectiveFieldsDuringRead(existingState PipelinePermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissions
// only implements ToObjectValue() and Type().
func (o PipelinePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PipelineAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PipelinePermissions as
// a slice of PipelineAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermissions) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissions.
func (o *PipelinePermissions) SetAccessControlList(ctx context.Context, v []PipelineAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PipelinePermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PipelinePermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissionsDescription) {
}

func (newState *PipelinePermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PipelinePermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsDescription
// only implements ToObjectValue() and Type().
func (o PipelinePermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PipelinePermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *PipelinePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissionsRequest) {
}

func (newState *PipelinePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PipelinePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsRequest
// only implements ToObjectValue() and Type().
func (o PipelinePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"pipeline_id":         o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PipelineAccessControlRequest{}.Type(ctx),
			},
			"pipeline_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PipelinePermissionsRequest as
// a slice of PipelineAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermissionsRequest) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissionsRequest.
func (o *PipelinePermissionsRequest) SetAccessControlList(ctx context.Context, v []PipelineAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel" tf:"optional"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters" tf:"optional"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration" tf:"optional"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous" tf:"optional"`
	// Deployment type of this pipeline.
	Deployment types.Object `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition" tf:"optional,object"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Restart window of this pipeline.
	RestartWindow types.Object `tfsdk:"restart_window" tf:"optional,object"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless" tf:"optional"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage" tf:"optional"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target types.String `tfsdk:"target" tf:"optional"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.Object `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *PipelineSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineSpec) {
}

func (newState *PipelineSpec) SyncEffectiveFieldsDuringRead(existingState PipelineSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"trigger":              reflect.TypeOf(PipelineTrigger{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineSpec
// only implements ToObjectValue() and Type().
func (o PipelineSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":     o.BudgetPolicyId,
			"catalog":              o.Catalog,
			"channel":              o.Channel,
			"clusters":             o.Clusters,
			"configuration":        o.Configuration,
			"continuous":           o.Continuous,
			"deployment":           o.Deployment,
			"development":          o.Development,
			"edition":              o.Edition,
			"filters":              o.Filters,
			"gateway_definition":   o.GatewayDefinition,
			"id":                   o.Id,
			"ingestion_definition": o.IngestionDefinition,
			"libraries":            o.Libraries,
			"name":                 o.Name,
			"notifications":        o.Notifications,
			"photon":               o.Photon,
			"restart_window":       o.RestartWindow,
			"schema":               o.Schema,
			"serverless":           o.Serverless,
			"storage":              o.Storage,
			"target":               o.Target,
			"trigger":              o.Trigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"catalog":          types.StringType,
			"channel":          types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous":           types.BoolType,
			"deployment":           PipelineDeployment{}.Type(ctx),
			"development":          types.BoolType,
			"edition":              types.StringType,
			"filters":              Filters{}.Type(ctx),
			"gateway_definition":   IngestionGatewayPipelineDefinition{}.Type(ctx),
			"id":                   types.StringType,
			"ingestion_definition": IngestionPipelineDefinition{}.Type(ctx),
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon":         types.BoolType,
			"restart_window": RestartWindow{}.Type(ctx),
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"target":         types.StringType,
			"trigger":        PipelineTrigger{}.Type(ctx),
		},
	}
}

// GetClusters returns the value of the Clusters field in PipelineSpec as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in PipelineSpec.
func (o *PipelineSpec) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in PipelineSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if o.Configuration.IsNull() || o.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in PipelineSpec.
func (o *PipelineSpec) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in PipelineSpec as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDeployment sets the value of the Deployment field in PipelineSpec.
func (o *PipelineSpec) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetFilters returns the value of the Filters field in PipelineSpec as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFilters sets the value of the Filters field in PipelineSpec.
func (o *PipelineSpec) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	o.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in PipelineSpec as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGatewayDefinition sets the value of the GatewayDefinition field in PipelineSpec.
func (o *PipelineSpec) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in PipelineSpec as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIngestionDefinition sets the value of the IngestionDefinition field in PipelineSpec.
func (o *PipelineSpec) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	o.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in PipelineSpec as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in PipelineSpec.
func (o *PipelineSpec) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in PipelineSpec as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in PipelineSpec.
func (o *PipelineSpec) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in PipelineSpec as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRestartWindow sets the value of the RestartWindow field in PipelineSpec.
func (o *PipelineSpec) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetTrigger returns the value of the Trigger field in PipelineSpec as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTrigger sets the value of the Trigger field in PipelineSpec.
func (o *PipelineSpec) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health" tf:"optional"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates types.List `tfsdk:"latest_updates" tf:"optional"`
	// The user-friendly name of the pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The unique identifier of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName types.String `tfsdk:"run_as_user_name" tf:"optional"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *PipelineStateInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineStateInfo) {
}

func (newState *PipelineStateInfo) SyncEffectiveFieldsDuringRead(existingState PipelineStateInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineStateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineStateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineStateInfo
// only implements ToObjectValue() and Type().
func (o PipelineStateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":        o.ClusterId,
			"creator_user_name": o.CreatorUserName,
			"health":            o.Health,
			"latest_updates":    o.LatestUpdates,
			"name":              o.Name,
			"pipeline_id":       o.PipelineId,
			"run_as_user_name":  o.RunAsUserName,
			"state":             o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineStateInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":        types.StringType,
			"creator_user_name": types.StringType,
			"health":            types.StringType,
			"latest_updates": basetypes.ListType{
				ElemType: UpdateStateInfo{}.Type(ctx),
			},
			"name":             types.StringType,
			"pipeline_id":      types.StringType,
			"run_as_user_name": types.StringType,
			"state":            types.StringType,
		},
	}
}

// GetLatestUpdates returns the value of the LatestUpdates field in PipelineStateInfo as
// a slice of UpdateStateInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineStateInfo) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo, bool) {
	if o.LatestUpdates.IsNull() || o.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo
	d := o.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in PipelineStateInfo.
func (o *PipelineStateInfo) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestUpdates = types.ListValueMust(t, vs)
}

type PipelineTrigger struct {
	Cron types.Object `tfsdk:"cron" tf:"optional,object"`

	Manual types.Object `tfsdk:"manual" tf:"optional,object"`
}

func (newState *PipelineTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineTrigger) {
}

func (newState *PipelineTrigger) SyncEffectiveFieldsDuringRead(existingState PipelineTrigger) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron":   reflect.TypeOf(CronTrigger{}),
		"manual": reflect.TypeOf(ManualTrigger{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineTrigger
// only implements ToObjectValue() and Type().
func (o PipelineTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cron":   o.Cron,
			"manual": o.Manual,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineTrigger) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cron":   CronTrigger{}.Type(ctx),
			"manual": ManualTrigger{}.Type(ctx),
		},
	}
}

// GetCron returns the value of the Cron field in PipelineTrigger as
// a CronTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineTrigger) GetCron(ctx context.Context) (CronTrigger, bool) {
	var e CronTrigger
	if o.Cron.IsNull() || o.Cron.IsUnknown() {
		return e, false
	}
	var v []CronTrigger
	d := o.Cron.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCron sets the value of the Cron field in PipelineTrigger.
func (o *PipelineTrigger) SetCron(ctx context.Context, v CronTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Cron = vs
}

// GetManual returns the value of the Manual field in PipelineTrigger as
// a ManualTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineTrigger) GetManual(ctx context.Context) (ManualTrigger, bool) {
	var e ManualTrigger
	if o.Manual.IsNull() || o.Manual.IsUnknown() {
		return e, false
	}
	var v []ManualTrigger
	d := o.Manual.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetManual sets the value of the Manual field in PipelineTrigger.
func (o *PipelineTrigger) SetManual(ctx context.Context, v ManualTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Manual = vs
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog types.String `tfsdk:"destination_catalog" tf:"optional"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema" tf:"optional"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable types.String `tfsdk:"destination_table" tf:"optional"`
	// Required. Report URL in the source system.
	SourceUrl types.String `tfsdk:"source_url" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	TableConfiguration types.Object `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *ReportSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReportSpec) {
}

func (newState *ReportSpec) SyncEffectiveFieldsDuringRead(existingState ReportSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReportSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReportSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReportSpec
// only implements ToObjectValue() and Type().
func (o ReportSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": o.DestinationCatalog,
			"destination_schema":  o.DestinationSchema,
			"destination_table":   o.DestinationTable,
			"source_url":          o.SourceUrl,
			"table_configuration": o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReportSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_url":          types.StringType,
			"table_configuration": TableSpecificConfig{}.Type(ctx),
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in ReportSpec as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ReportSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTableConfiguration sets the value of the TableConfiguration field in ReportSpec.
func (o *ReportSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	DaysOfWeek types.List `tfsdk:"days_of_week" tf:"optional"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour types.Int64 `tfsdk:"start_hour" tf:""`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId types.String `tfsdk:"time_zone_id" tf:"optional"`
}

func (newState *RestartWindow) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestartWindow) {
}

func (newState *RestartWindow) SyncEffectiveFieldsDuringRead(existingState RestartWindow) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"days_of_week": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartWindow
// only implements ToObjectValue() and Type().
func (o RestartWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"days_of_week": o.DaysOfWeek,
			"start_hour":   o.StartHour,
			"time_zone_id": o.TimeZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestartWindow) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"days_of_week": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_hour":   types.Int64Type,
			"time_zone_id": types.StringType,
		},
	}
}

// GetDaysOfWeek returns the value of the DaysOfWeek field in RestartWindow as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RestartWindow) GetDaysOfWeek(ctx context.Context) ([]types.String, bool) {
	if o.DaysOfWeek.IsNull() || o.DaysOfWeek.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DaysOfWeek.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDaysOfWeek sets the value of the DaysOfWeek field in RestartWindow.
func (o *RestartWindow) SetDaysOfWeek(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["days_of_week"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DaysOfWeek = types.ListValueMust(t, vs)
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	DestinationCatalog types.String `tfsdk:"destination_catalog" tf:"optional"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema types.String `tfsdk:"destination_schema" tf:"optional"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog types.String `tfsdk:"source_catalog" tf:"optional"`
	// Required. Schema name in the source database.
	SourceSchema types.String `tfsdk:"source_schema" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration types.Object `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *SchemaSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaSpec) {
}

func (newState *SchemaSpec) SyncEffectiveFieldsDuringRead(existingState SchemaSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SchemaSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SchemaSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaSpec
// only implements ToObjectValue() and Type().
func (o SchemaSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": o.DestinationCatalog,
			"destination_schema":  o.DestinationSchema,
			"source_catalog":      o.SourceCatalog,
			"source_schema":       o.SourceSchema,
			"table_configuration": o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SchemaSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"table_configuration": TableSpecificConfig{}.Type(ctx),
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in SchemaSpec as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTableConfiguration sets the value of the TableConfiguration field in SchemaSpec.
func (o *SchemaSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no" tf:"optional"`
	// the ID assigned by the data plane.
	DataPlaneId types.Object `tfsdk:"data_plane_id" tf:"optional,object"`
}

func (newState *Sequencing) SyncEffectiveFieldsDuringCreateOrUpdate(plan Sequencing) {
}

func (newState *Sequencing) SyncEffectiveFieldsDuringRead(existingState Sequencing) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Sequencing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Sequencing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_id": reflect.TypeOf(DataPlaneId{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sequencing
// only implements ToObjectValue() and Type().
func (o Sequencing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"control_plane_seq_no": o.ControlPlaneSeqNo,
			"data_plane_id":        o.DataPlaneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Sequencing) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"control_plane_seq_no": types.Int64Type,
			"data_plane_id":        DataPlaneId{}.Type(ctx),
		},
	}
}

// GetDataPlaneId returns the value of the DataPlaneId field in Sequencing as
// a DataPlaneId value.
// If the field is unknown or null, the boolean return value is false.
func (o *Sequencing) GetDataPlaneId(ctx context.Context) (DataPlaneId, bool) {
	var e DataPlaneId
	if o.DataPlaneId.IsNull() || o.DataPlaneId.IsUnknown() {
		return e, false
	}
	var v []DataPlaneId
	d := o.DataPlaneId.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDataPlaneId sets the value of the DataPlaneId field in Sequencing.
func (o *Sequencing) SetDataPlaneId(ctx context.Context, v DataPlaneId) {
	vs := v.ToObjectValue(ctx)
	o.DataPlaneId = vs
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name" tf:"optional"`
	// Exception message
	Message types.String `tfsdk:"message" tf:"optional"`
	// Stack trace consisting of a list of stack frames
	Stack types.List `tfsdk:"stack" tf:"optional"`
}

func (newState *SerializedException) SyncEffectiveFieldsDuringCreateOrUpdate(plan SerializedException) {
}

func (newState *SerializedException) SyncEffectiveFieldsDuringRead(existingState SerializedException) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SerializedException.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SerializedException) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stack": reflect.TypeOf(StackFrame{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SerializedException
// only implements ToObjectValue() and Type().
func (o SerializedException) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"class_name": o.ClassName,
			"message":    o.Message,
			"stack":      o.Stack,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SerializedException) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"class_name": types.StringType,
			"message":    types.StringType,
			"stack": basetypes.ListType{
				ElemType: StackFrame{}.Type(ctx),
			},
		},
	}
}

// GetStack returns the value of the Stack field in SerializedException as
// a slice of StackFrame values.
// If the field is unknown or null, the boolean return value is false.
func (o *SerializedException) GetStack(ctx context.Context) ([]StackFrame, bool) {
	if o.Stack.IsNull() || o.Stack.IsUnknown() {
		return nil, false
	}
	var v []StackFrame
	d := o.Stack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStack sets the value of the Stack field in SerializedException.
func (o *SerializedException) SetStack(ctx context.Context, v []StackFrame) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stack"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stack = types.ListValueMust(t, vs)
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class" tf:"optional"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name" tf:"optional"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number" tf:"optional"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name" tf:"optional"`
}

func (newState *StackFrame) SyncEffectiveFieldsDuringCreateOrUpdate(plan StackFrame) {
}

func (newState *StackFrame) SyncEffectiveFieldsDuringRead(existingState StackFrame) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StackFrame.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StackFrame) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StackFrame
// only implements ToObjectValue() and Type().
func (o StackFrame) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"declaring_class": o.DeclaringClass,
			"file_name":       o.FileName,
			"line_number":     o.LineNumber,
			"method_name":     o.MethodName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StackFrame) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"declaring_class": types.StringType,
			"file_name":       types.StringType,
			"line_number":     types.Int64Type,
			"method_name":     types.StringType,
		},
	}
}

type StartUpdate struct {
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection types.List `tfsdk:"full_refresh_selection" tf:"optional"`

	PipelineId types.String `tfsdk:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection types.List `tfsdk:"refresh_selection" tf:"optional"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

func (newState *StartUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartUpdate) {
}

func (newState *StartUpdate) SyncEffectiveFieldsDuringRead(existingState StartUpdate) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdate
// only implements ToObjectValue() and Type().
func (o StartUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                  o.Cause,
			"full_refresh":           o.FullRefresh,
			"full_refresh_selection": o.FullRefreshSelection,
			"pipeline_id":            o.PipelineId,
			"refresh_selection":      o.RefreshSelection,
			"validate_only":          o.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":        types.StringType,
			"full_refresh": types.BoolType,
			"full_refresh_selection": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pipeline_id": types.StringType,
			"refresh_selection": basetypes.ListType{
				ElemType: types.StringType,
			},
			"validate_only": types.BoolType,
		},
	}
}

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in StartUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *StartUpdate) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if o.FullRefreshSelection.IsNull() || o.FullRefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FullRefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in StartUpdate.
func (o *StartUpdate) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in StartUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *StartUpdate) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if o.RefreshSelection.IsNull() || o.RefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshSelection sets the value of the RefreshSelection field in StartUpdate.
func (o *StartUpdate) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RefreshSelection = types.ListValueMust(t, vs)
}

type StartUpdateResponse struct {
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *StartUpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartUpdateResponse) {
}

func (newState *StartUpdateResponse) SyncEffectiveFieldsDuringRead(existingState StartUpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartUpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartUpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdateResponse
// only implements ToObjectValue() and Type().
func (o StartUpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_id": o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartUpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update_id": types.StringType,
		},
	}
}

type StopPipelineResponse struct {
}

func (newState *StopPipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopPipelineResponse) {
}

func (newState *StopPipelineResponse) SyncEffectiveFieldsDuringRead(existingState StopPipelineResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopPipelineResponse
// only implements ToObjectValue() and Type().
func (o StopPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StopPipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Stop a pipeline
type StopRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *StopRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopRequest) {
}

func (newState *StopRequest) SyncEffectiveFieldsDuringRead(existingState StopRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest
// only implements ToObjectValue() and Type().
func (o StopRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog types.String `tfsdk:"destination_catalog" tf:"optional"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema" tf:"optional"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable types.String `tfsdk:"destination_table" tf:"optional"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog types.String `tfsdk:"source_catalog" tf:"optional"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema types.String `tfsdk:"source_schema" tf:"optional"`
	// Required. Table name in the source database.
	SourceTable types.String `tfsdk:"source_table" tf:"optional"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration types.Object `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *TableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSpec) {
}

func (newState *TableSpec) SyncEffectiveFieldsDuringRead(existingState TableSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpec
// only implements ToObjectValue() and Type().
func (o TableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": o.DestinationCatalog,
			"destination_schema":  o.DestinationSchema,
			"destination_table":   o.DestinationTable,
			"source_catalog":      o.SourceCatalog,
			"source_schema":       o.SourceSchema,
			"source_table":        o.SourceTable,
			"table_configuration": o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"source_table":        types.StringType,
			"table_configuration": TableSpecificConfig{}.Type(ctx),
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in TableSpec as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTableConfiguration sets the value of the TableConfiguration field in TableSpec.
func (o *TableSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

type TableSpecificConfig struct {
	// The primary key of the table used to apply changes.
	PrimaryKeys types.List `tfsdk:"primary_keys" tf:"optional"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields types.Bool `tfsdk:"salesforce_include_formula_fields" tf:"optional"`
	// The SCD type to use to ingest the table.
	ScdType types.String `tfsdk:"scd_type" tf:"optional"`
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	SequenceBy types.List `tfsdk:"sequence_by" tf:"optional"`
}

func (newState *TableSpecificConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSpecificConfig) {
}

func (newState *TableSpecificConfig) SyncEffectiveFieldsDuringRead(existingState TableSpecificConfig) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSpecificConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSpecificConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
		"sequence_by":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpecificConfig
// only implements ToObjectValue() and Type().
func (o TableSpecificConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"primary_keys":                      o.PrimaryKeys,
			"salesforce_include_formula_fields": o.SalesforceIncludeFormulaFields,
			"scd_type":                          o.ScdType,
			"sequence_by":                       o.SequenceBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSpecificConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"salesforce_include_formula_fields": types.BoolType,
			"scd_type":                          types.StringType,
			"sequence_by": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if o.PrimaryKeys.IsNull() || o.PrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeys sets the value of the PrimaryKeys field in TableSpecificConfig.
func (o *TableSpecificConfig) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeys = types.ListValueMust(t, vs)
}

// GetSequenceBy returns the value of the SequenceBy field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetSequenceBy(ctx context.Context) ([]types.String, bool) {
	if o.SequenceBy.IsNull() || o.SequenceBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SequenceBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSequenceBy sets the value of the SequenceBy field in TableSpecificConfig.
func (o *TableSpecificConfig) SetSequenceBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SequenceBy = types.ListValueMust(t, vs)
}

type UpdateInfo struct {
	// What triggered this update.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config types.Object `tfsdk:"config" tf:"optional,object"`
	// The time when this update was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection types.List `tfsdk:"full_refresh_selection" tf:"optional"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection types.List `tfsdk:"refresh_selection" tf:"optional"`
	// The update state.
	State types.String `tfsdk:"state" tf:"optional"`
	// The ID of this update.
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

func (newState *UpdateInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInfo) {
}

func (newState *UpdateInfo) SyncEffectiveFieldsDuringRead(existingState UpdateInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":                 reflect.TypeOf(PipelineSpec{}),
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInfo
// only implements ToObjectValue() and Type().
func (o UpdateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                  o.Cause,
			"cluster_id":             o.ClusterId,
			"config":                 o.Config,
			"creation_time":          o.CreationTime,
			"full_refresh":           o.FullRefresh,
			"full_refresh_selection": o.FullRefreshSelection,
			"pipeline_id":            o.PipelineId,
			"refresh_selection":      o.RefreshSelection,
			"state":                  o.State,
			"update_id":              o.UpdateId,
			"validate_only":          o.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":         types.StringType,
			"cluster_id":    types.StringType,
			"config":        PipelineSpec{}.Type(ctx),
			"creation_time": types.Int64Type,
			"full_refresh":  types.BoolType,
			"full_refresh_selection": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pipeline_id": types.StringType,
			"refresh_selection": basetypes.ListType{
				ElemType: types.StringType,
			},
			"state":         types.StringType,
			"update_id":     types.StringType,
			"validate_only": types.BoolType,
		},
	}
}

// GetConfig returns the value of the Config field in UpdateInfo as
// a PipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo) GetConfig(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetConfig sets the value of the Config field in UpdateInfo.
func (o *UpdateInfo) SetConfig(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in UpdateInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if o.FullRefreshSelection.IsNull() || o.FullRefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FullRefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in UpdateInfo.
func (o *UpdateInfo) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in UpdateInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if o.RefreshSelection.IsNull() || o.RefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshSelection sets the value of the RefreshSelection field in UpdateInfo.
func (o *UpdateInfo) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RefreshSelection = types.ListValueMust(t, vs)
}

type UpdateStateInfo struct {
	CreationTime types.String `tfsdk:"creation_time" tf:"optional"`

	State types.String `tfsdk:"state" tf:"optional"`

	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *UpdateStateInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStateInfo) {
}

func (newState *UpdateStateInfo) SyncEffectiveFieldsDuringRead(existingState UpdateStateInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateStateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateStateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStateInfo
// only implements ToObjectValue() and Type().
func (o UpdateStateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time": o.CreationTime,
			"state":         o.State,
			"update_id":     o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateStateInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.StringType,
			"state":         types.StringType,
			"update_id":     types.StringType,
		},
	}
}
