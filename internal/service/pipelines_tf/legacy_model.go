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

type CreatePipeline_SdkV2 struct {
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
	Deployment types.List `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`

	DryRun types.Bool `tfsdk:"dry_run" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition" tf:"optional,object"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Restart window of this pipeline.
	RestartWindow types.List `tfsdk:"restart_window" tf:"optional,object"`
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
	Trigger types.List `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *CreatePipeline_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipeline_SdkV2) {
}

func (newState *CreatePipeline_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreatePipeline_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePipeline.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePipeline_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster_SdkV2{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"trigger":              reflect.TypeOf(PipelineTrigger_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipeline_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePipeline_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreatePipeline_SdkV2) Type(ctx context.Context) attr.Type {
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
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.Type(ctx),
			},
			"development": types.BoolType,
			"dry_run":     types.BoolType,
			"edition":     types.StringType,
			"filters": basetypes.ListType{
				ElemType: Filters{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.Type(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.Type(ctx),
			},
		},
	}
}

// GetClusters returns the value of the Clusters field in CreatePipeline_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in CreatePipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in CreatePipeline_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := o.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	o.Deployment = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in CreatePipeline_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	o.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in CreatePipeline_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := o.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	o.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in CreatePipeline_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := o.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	o.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in CreatePipeline_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in CreatePipeline_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in CreatePipeline_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := o.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	o.RestartWindow = types.ListValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in CreatePipeline_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := o.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	o.Trigger = types.ListValueMust(t, vs)
}

type CreatePipelineResponse_SdkV2 struct {
	// Only returned when dry_run is true.
	EffectiveSettings types.List `tfsdk:"effective_settings" tf:"optional,object"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
}

func (newState *CreatePipelineResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipelineResponse_SdkV2) {
}

func (newState *CreatePipelineResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreatePipelineResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(PipelineSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_settings": o.EffectiveSettings,
			"pipeline_id":        o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_settings": basetypes.ListType{
				ElemType: PipelineSpec{}.Type(ctx),
			},
			"pipeline_id": types.StringType,
		},
	}
}

// GetEffectiveSettings returns the value of the EffectiveSettings field in CreatePipelineResponse_SdkV2 as
// a PipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipelineResponse_SdkV2) GetEffectiveSettings(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if o.EffectiveSettings.IsNull() || o.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := o.EffectiveSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in CreatePipelineResponse_SdkV2.
func (o *CreatePipelineResponse_SdkV2) SetEffectiveSettings(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_settings"]
	o.EffectiveSettings = types.ListValueMust(t, vs)
}

type CronTrigger_SdkV2 struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule" tf:"optional"`

	TimezoneId types.String `tfsdk:"timezone_id" tf:"optional"`
}

func (newState *CronTrigger_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronTrigger_SdkV2) {
}

func (newState *CronTrigger_SdkV2) SyncEffectiveFieldsDuringRead(existingState CronTrigger_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (o CronTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_schedule": o.QuartzCronSchedule,
			"timezone_id":          o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronTrigger_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quartz_cron_schedule": types.StringType,
			"timezone_id":          types.StringType,
		},
	}
}

type DataPlaneId_SdkV2 struct {
	// The instance name of the data plane emitting an event.
	Instance types.String `tfsdk:"instance" tf:"optional"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no" tf:"optional"`
}

func (newState *DataPlaneId_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneId_SdkV2) {
}

func (newState *DataPlaneId_SdkV2) SyncEffectiveFieldsDuringRead(existingState DataPlaneId_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataPlaneId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataPlaneId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneId_SdkV2
// only implements ToObjectValue() and Type().
func (o DataPlaneId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance": o.Instance,
			"seq_no":   o.SeqNo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataPlaneId_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance": types.StringType,
			"seq_no":   types.Int64Type,
		},
	}
}

// Delete a pipeline
type DeletePipelineRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *DeletePipelineRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePipelineRequest_SdkV2) {
}

func (newState *DeletePipelineRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeletePipelineRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePipelineRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePipelineRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePipelineRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type DeletePipelineResponse_SdkV2 struct {
}

func (newState *DeletePipelineResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePipelineResponse_SdkV2) {
}

func (newState *DeletePipelineResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeletePipelineResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EditPipeline_SdkV2 struct {
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
	Deployment types.List `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition" tf:"optional,object"`
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
	RestartWindow types.List `tfsdk:"restart_window" tf:"optional,object"`
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
	Trigger types.List `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *EditPipeline_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipeline_SdkV2) {
}

func (newState *EditPipeline_SdkV2) SyncEffectiveFieldsDuringRead(existingState EditPipeline_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipeline.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPipeline_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster_SdkV2{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"trigger":              reflect.TypeOf(PipelineTrigger_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipeline_SdkV2
// only implements ToObjectValue() and Type().
func (o EditPipeline_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EditPipeline_SdkV2) Type(ctx context.Context) attr.Type {
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
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.Type(ctx),
			},
			"development":            types.BoolType,
			"edition":                types.StringType,
			"expected_last_modified": types.Int64Type,
			"filters": basetypes.ListType{
				ElemType: Filters{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon":      types.BoolType,
			"pipeline_id": types.StringType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.Type(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.Type(ctx),
			},
		},
	}
}

// GetClusters returns the value of the Clusters field in EditPipeline_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in EditPipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in EditPipeline_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := o.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	o.Deployment = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in EditPipeline_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	o.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in EditPipeline_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := o.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	o.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in EditPipeline_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := o.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	o.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in EditPipeline_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in EditPipeline_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in EditPipeline_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := o.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	o.RestartWindow = types.ListValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in EditPipeline_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := o.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	o.Trigger = types.ListValueMust(t, vs)
}

type EditPipelineResponse_SdkV2 struct {
}

func (newState *EditPipelineResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipelineResponse_SdkV2) {
}

func (newState *EditPipelineResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState EditPipelineResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EditPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EditPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EditPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ErrorDetail_SdkV2 struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions types.List `tfsdk:"exceptions" tf:"optional"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal" tf:"optional"`
}

func (newState *ErrorDetail_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ErrorDetail_SdkV2) {
}

func (newState *ErrorDetail_SdkV2) SyncEffectiveFieldsDuringRead(existingState ErrorDetail_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ErrorDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ErrorDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exceptions": reflect.TypeOf(SerializedException_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ErrorDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o ErrorDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exceptions": o.Exceptions,
			"fatal":      o.Fatal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ErrorDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exceptions": basetypes.ListType{
				ElemType: SerializedException{}.Type(ctx),
			},
			"fatal": types.BoolType,
		},
	}
}

// GetExceptions returns the value of the Exceptions field in ErrorDetail_SdkV2 as
// a slice of SerializedException_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ErrorDetail_SdkV2) GetExceptions(ctx context.Context) ([]SerializedException_SdkV2, bool) {
	if o.Exceptions.IsNull() || o.Exceptions.IsUnknown() {
		return nil, false
	}
	var v []SerializedException_SdkV2
	d := o.Exceptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExceptions sets the value of the Exceptions field in ErrorDetail_SdkV2.
func (o *ErrorDetail_SdkV2) SetExceptions(ctx context.Context, v []SerializedException_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exceptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exceptions = types.ListValueMust(t, vs)
}

type FileLibrary_SdkV2 struct {
	// The absolute path of the file.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *FileLibrary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileLibrary_SdkV2) {
}

func (newState *FileLibrary_SdkV2) SyncEffectiveFieldsDuringRead(existingState FileLibrary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o FileLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileLibrary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type Filters_SdkV2 struct {
	// Paths to exclude.
	Exclude types.List `tfsdk:"exclude" tf:"optional"`
	// Paths to include.
	Include types.List `tfsdk:"include" tf:"optional"`
}

func (newState *Filters_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Filters_SdkV2) {
}

func (newState *Filters_SdkV2) SyncEffectiveFieldsDuringRead(existingState Filters_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Filters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Filters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude": reflect.TypeOf(types.String{}),
		"include": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filters_SdkV2
// only implements ToObjectValue() and Type().
func (o Filters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude": o.Exclude,
			"include": o.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Filters_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetExclude returns the value of the Exclude field in Filters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Filters_SdkV2) GetExclude(ctx context.Context) ([]types.String, bool) {
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

// SetExclude sets the value of the Exclude field in Filters_SdkV2.
func (o *Filters_SdkV2) SetExclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Exclude = types.ListValueMust(t, vs)
}

// GetInclude returns the value of the Include field in Filters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Filters_SdkV2) GetInclude(ctx context.Context) ([]types.String, bool) {
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

// SetInclude sets the value of the Include field in Filters_SdkV2.
func (o *Filters_SdkV2) SetInclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Include = types.ListValueMust(t, vs)
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelinePermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionLevelsRequest_SdkV2) {
}

func (newState *GetPipelinePermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionLevelsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type GetPipelinePermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPipelinePermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionLevelsResponse_SdkV2) {
}

func (newState *GetPipelinePermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionLevelsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PipelinePermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PipelinePermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse_SdkV2 as
// a slice of PipelinePermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelinePermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]PipelinePermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse_SdkV2.
func (o *GetPipelinePermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []PipelinePermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get pipeline permissions
type GetPipelinePermissionsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelinePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionsRequest_SdkV2) {
}

func (newState *GetPipelinePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelinePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelinePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPipelinePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelinePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

// Get a pipeline
type GetPipelineRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelineRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelineRequest_SdkV2) {
}

func (newState *GetPipelineRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPipelineRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelineRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelineRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPipelineRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPipelineRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type GetPipelineResponse_SdkV2 struct {
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
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *GetPipelineResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelineResponse_SdkV2) {
}

func (newState *GetPipelineResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPipelineResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo_SdkV2{}),
		"spec":           reflect.TypeOf(PipelineSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
			"spec": basetypes.ListType{
				ElemType: PipelineSpec{}.Type(ctx),
			},
			"state": types.StringType,
		},
	}
}

// GetLatestUpdates returns the value of the LatestUpdates field in GetPipelineResponse_SdkV2 as
// a slice of UpdateStateInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse_SdkV2) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo_SdkV2, bool) {
	if o.LatestUpdates.IsNull() || o.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo_SdkV2
	d := o.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in GetPipelineResponse_SdkV2.
func (o *GetPipelineResponse_SdkV2) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestUpdates = types.ListValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in GetPipelineResponse_SdkV2 as
// a PipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse_SdkV2) GetSpec(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := o.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in GetPipelineResponse_SdkV2.
func (o *GetPipelineResponse_SdkV2) SetSpec(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	o.Spec = types.ListValueMust(t, vs)
}

// Get a pipeline update
type GetUpdateRequest_SdkV2 struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
}

func (newState *GetUpdateRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUpdateRequest_SdkV2) {
}

func (newState *GetUpdateRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetUpdateRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUpdateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetUpdateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
			"update_id":   o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetUpdateRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
			"update_id":   types.StringType,
		},
	}
}

type GetUpdateResponse_SdkV2 struct {
	// The current update info.
	Update types.List `tfsdk:"update" tf:"optional,object"`
}

func (newState *GetUpdateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUpdateResponse_SdkV2) {
}

func (newState *GetUpdateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetUpdateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"update": reflect.TypeOf(UpdateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetUpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update": o.Update,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetUpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update": basetypes.ListType{
				ElemType: UpdateInfo{}.Type(ctx),
			},
		},
	}
}

// GetUpdate returns the value of the Update field in GetUpdateResponse_SdkV2 as
// a UpdateInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetUpdateResponse_SdkV2) GetUpdate(ctx context.Context) (UpdateInfo_SdkV2, bool) {
	var e UpdateInfo_SdkV2
	if o.Update.IsNull() || o.Update.IsUnknown() {
		return e, false
	}
	var v []UpdateInfo_SdkV2
	d := o.Update.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdate sets the value of the Update field in GetUpdateResponse_SdkV2.
func (o *GetUpdateResponse_SdkV2) SetUpdate(ctx context.Context, v UpdateInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["update"]
	o.Update = types.ListValueMust(t, vs)
}

type IngestionConfig_SdkV2 struct {
	// Select a specific source report.
	Report types.List `tfsdk:"report" tf:"optional,object"`
	// Select all tables from a specific source schema.
	Schema types.List `tfsdk:"schema" tf:"optional,object"`
	// Select a specific source table.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *IngestionConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionConfig_SdkV2) {
}

func (newState *IngestionConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState IngestionConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"report": reflect.TypeOf(ReportSpec_SdkV2{}),
		"schema": reflect.TypeOf(SchemaSpec_SdkV2{}),
		"table":  reflect.TypeOf(TableSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o IngestionConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"report": o.Report,
			"schema": o.Schema,
			"table":  o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"report": basetypes.ListType{
				ElemType: ReportSpec{}.Type(ctx),
			},
			"schema": basetypes.ListType{
				ElemType: SchemaSpec{}.Type(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableSpec{}.Type(ctx),
			},
		},
	}
}

// GetReport returns the value of the Report field in IngestionConfig_SdkV2 as
// a ReportSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig_SdkV2) GetReport(ctx context.Context) (ReportSpec_SdkV2, bool) {
	var e ReportSpec_SdkV2
	if o.Report.IsNull() || o.Report.IsUnknown() {
		return e, false
	}
	var v []ReportSpec_SdkV2
	d := o.Report.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReport sets the value of the Report field in IngestionConfig_SdkV2.
func (o *IngestionConfig_SdkV2) SetReport(ctx context.Context, v ReportSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["report"]
	o.Report = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in IngestionConfig_SdkV2 as
// a SchemaSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig_SdkV2) GetSchema(ctx context.Context) (SchemaSpec_SdkV2, bool) {
	var e SchemaSpec_SdkV2
	if o.Schema.IsNull() || o.Schema.IsUnknown() {
		return e, false
	}
	var v []SchemaSpec_SdkV2
	d := o.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in IngestionConfig_SdkV2.
func (o *IngestionConfig_SdkV2) SetSchema(ctx context.Context, v SchemaSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	o.Schema = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in IngestionConfig_SdkV2 as
// a TableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionConfig_SdkV2) GetTable(ctx context.Context) (TableSpec_SdkV2, bool) {
	var e TableSpec_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []TableSpec_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in IngestionConfig_SdkV2.
func (o *IngestionConfig_SdkV2) SetTable(ctx context.Context, v TableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

type IngestionGatewayPipelineDefinition_SdkV2 struct {
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

func (newState *IngestionGatewayPipelineDefinition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionGatewayPipelineDefinition_SdkV2) {
}

func (newState *IngestionGatewayPipelineDefinition_SdkV2) SyncEffectiveFieldsDuringRead(existingState IngestionGatewayPipelineDefinition_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionGatewayPipelineDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionGatewayPipelineDefinition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionGatewayPipelineDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (o IngestionGatewayPipelineDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o IngestionGatewayPipelineDefinition_SdkV2) Type(ctx context.Context) attr.Type {
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

type IngestionPipelineDefinition_SdkV2 struct {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *IngestionPipelineDefinition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionPipelineDefinition_SdkV2) {
}

func (newState *IngestionPipelineDefinition_SdkV2) SyncEffectiveFieldsDuringRead(existingState IngestionPipelineDefinition_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionPipelineDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionPipelineDefinition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects":             reflect.TypeOf(IngestionConfig_SdkV2{}),
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o IngestionPipelineDefinition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":      types.StringType,
			"ingestion_gateway_id": types.StringType,
			"objects": basetypes.ListType{
				ElemType: IngestionConfig{}.Type(ctx),
			},
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.Type(ctx),
			},
		},
	}
}

// GetObjects returns the value of the Objects field in IngestionPipelineDefinition_SdkV2 as
// a slice of IngestionConfig_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition_SdkV2) GetObjects(ctx context.Context) ([]IngestionConfig_SdkV2, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []IngestionConfig_SdkV2
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in IngestionPipelineDefinition_SdkV2.
func (o *IngestionPipelineDefinition_SdkV2) SetObjects(ctx context.Context, v []IngestionConfig_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

// GetTableConfiguration returns the value of the TableConfiguration field in IngestionPipelineDefinition_SdkV2 as
// a TableSpecificConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := o.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in IngestionPipelineDefinition_SdkV2.
func (o *IngestionPipelineDefinition_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	o.TableConfiguration = types.ListValueMust(t, vs)
}

// List pipeline events
type ListPipelineEventsRequest_SdkV2 struct {
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

func (newState *ListPipelineEventsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelineEventsRequest_SdkV2) {
}

func (newState *ListPipelineEventsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListPipelineEventsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelineEventsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelineEventsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPipelineEventsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListPipelineEventsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOrderBy returns the value of the OrderBy field in ListPipelineEventsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelineEventsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in ListPipelineEventsRequest_SdkV2.
func (o *ListPipelineEventsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelineEventsResponse_SdkV2 struct {
	// The list of events matching the request criteria.
	Events types.List `tfsdk:"events" tf:"optional"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListPipelineEventsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelineEventsResponse_SdkV2) {
}

func (newState *ListPipelineEventsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListPipelineEventsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelineEventsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelineEventsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(PipelineEvent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPipelineEventsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          o.Events,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelineEventsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetEvents returns the value of the Events field in ListPipelineEventsResponse_SdkV2 as
// a slice of PipelineEvent_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelineEventsResponse_SdkV2) GetEvents(ctx context.Context) ([]PipelineEvent_SdkV2, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []PipelineEvent_SdkV2
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListPipelineEventsResponse_SdkV2.
func (o *ListPipelineEventsResponse_SdkV2) SetEvents(ctx context.Context, v []PipelineEvent_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// List pipelines
type ListPipelinesRequest_SdkV2 struct {
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

func (newState *ListPipelinesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelinesRequest_SdkV2) {
}

func (newState *ListPipelinesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListPipelinesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelinesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelinesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPipelinesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListPipelinesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOrderBy returns the value of the OrderBy field in ListPipelinesRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelinesRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in ListPipelinesRequest_SdkV2.
func (o *ListPipelinesRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelinesResponse_SdkV2 struct {
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// The list of events matching the request criteria.
	Statuses types.List `tfsdk:"statuses" tf:"optional"`
}

func (newState *ListPipelinesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPipelinesResponse_SdkV2) {
}

func (newState *ListPipelinesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListPipelinesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPipelinesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPipelinesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(PipelineStateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPipelinesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"statuses":        o.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPipelinesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"statuses": basetypes.ListType{
				ElemType: PipelineStateInfo{}.Type(ctx),
			},
		},
	}
}

// GetStatuses returns the value of the Statuses field in ListPipelinesResponse_SdkV2 as
// a slice of PipelineStateInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListPipelinesResponse_SdkV2) GetStatuses(ctx context.Context) ([]PipelineStateInfo_SdkV2, bool) {
	if o.Statuses.IsNull() || o.Statuses.IsUnknown() {
		return nil, false
	}
	var v []PipelineStateInfo_SdkV2
	d := o.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListPipelinesResponse_SdkV2.
func (o *ListPipelinesResponse_SdkV2) SetStatuses(ctx context.Context, v []PipelineStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Statuses = types.ListValueMust(t, vs)
}

// List pipeline updates
type ListUpdatesRequest_SdkV2 struct {
	// Max number of entries to return in a single page.
	MaxResults types.Int64 `tfsdk:"-"`
	// Page token returned by previous call
	PageToken types.String `tfsdk:"-"`
	// The pipeline to return updates for.
	PipelineId types.String `tfsdk:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId types.String `tfsdk:"-"`
}

func (newState *ListUpdatesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUpdatesRequest_SdkV2) {
}

func (newState *ListUpdatesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListUpdatesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUpdatesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUpdatesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListUpdatesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListUpdatesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":     types.Int64Type,
			"page_token":      types.StringType,
			"pipeline_id":     types.StringType,
			"until_update_id": types.StringType,
		},
	}
}

type ListUpdatesResponse_SdkV2 struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`

	Updates types.List `tfsdk:"updates" tf:"optional"`
}

func (newState *ListUpdatesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUpdatesResponse_SdkV2) {
}

func (newState *ListUpdatesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListUpdatesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUpdatesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUpdatesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(UpdateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListUpdatesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
			"updates":         o.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListUpdatesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetUpdates returns the value of the Updates field in ListUpdatesResponse_SdkV2 as
// a slice of UpdateInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListUpdatesResponse_SdkV2) GetUpdates(ctx context.Context) ([]UpdateInfo_SdkV2, bool) {
	if o.Updates.IsNull() || o.Updates.IsUnknown() {
		return nil, false
	}
	var v []UpdateInfo_SdkV2
	d := o.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in ListUpdatesResponse_SdkV2.
func (o *ListUpdatesResponse_SdkV2) SetUpdates(ctx context.Context, v []UpdateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Updates = types.ListValueMust(t, vs)
}

type ManualTrigger_SdkV2 struct {
}

func (newState *ManualTrigger_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ManualTrigger_SdkV2) {
}

func (newState *ManualTrigger_SdkV2) SyncEffectiveFieldsDuringRead(existingState ManualTrigger_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManualTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ManualTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManualTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (o ManualTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ManualTrigger_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type NotebookLibrary_SdkV2 struct {
	// The absolute path of the notebook.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *NotebookLibrary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookLibrary_SdkV2) {
}

func (newState *NotebookLibrary_SdkV2) SyncEffectiveFieldsDuringRead(existingState NotebookLibrary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o NotebookLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookLibrary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type Notifications_SdkV2 struct {
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

func (newState *Notifications_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Notifications_SdkV2) {
}

func (newState *Notifications_SdkV2) SyncEffectiveFieldsDuringRead(existingState Notifications_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Notifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Notifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":           reflect.TypeOf(types.String{}),
		"email_recipients": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Notifications_SdkV2
// only implements ToObjectValue() and Type().
func (o Notifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":           o.Alerts,
			"email_recipients": o.EmailRecipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Notifications_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAlerts returns the value of the Alerts field in Notifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Notifications_SdkV2) GetAlerts(ctx context.Context) ([]types.String, bool) {
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

// SetAlerts sets the value of the Alerts field in Notifications_SdkV2.
func (o *Notifications_SdkV2) SetAlerts(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Alerts = types.ListValueMust(t, vs)
}

// GetEmailRecipients returns the value of the EmailRecipients field in Notifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Notifications_SdkV2) GetEmailRecipients(ctx context.Context) ([]types.String, bool) {
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

// SetEmailRecipients sets the value of the EmailRecipients field in Notifications_SdkV2.
func (o *Notifications_SdkV2) SetEmailRecipients(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmailRecipients = types.ListValueMust(t, vs)
}

type Origin_SdkV2 struct {
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

func (newState *Origin_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Origin_SdkV2) {
}

func (newState *Origin_SdkV2) SyncEffectiveFieldsDuringRead(existingState Origin_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Origin.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Origin_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Origin_SdkV2
// only implements ToObjectValue() and Type().
func (o Origin_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Origin_SdkV2) Type(ctx context.Context) attr.Type {
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

type PipelineAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PipelineAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineAccessControlRequest_SdkV2) {
}

func (newState *PipelineAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineAccessControlRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PipelineAccessControlResponse_SdkV2 struct {
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

func (newState *PipelineAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineAccessControlResponse_SdkV2) {
}

func (newState *PipelineAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineAccessControlResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PipelinePermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAllPermissions returns the value of the AllPermissions field in PipelineAccessControlResponse_SdkV2 as
// a slice of PipelinePermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]PipelinePermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PipelineAccessControlResponse_SdkV2.
func (o *PipelineAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []PipelinePermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type PipelineCluster_SdkV2 struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf" tf:"optional,object"`
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
	GcpAttributes types.List `tfsdk:"gcp_attributes" tf:"optional,object"`
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

func (newState *PipelineCluster_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineCluster_SdkV2) {
}

func (newState *PipelineCluster_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineCluster_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(PipelineClusterAutoscale_SdkV2{}),
		"aws_attributes":   reflect.TypeOf(compute_tf.AwsAttributes_SdkV2{}),
		"azure_attributes": reflect.TypeOf(compute_tf.AzureAttributes_SdkV2{}),
		"cluster_log_conf": reflect.TypeOf(compute_tf.ClusterLogConf_SdkV2{}),
		"custom_tags":      reflect.TypeOf(types.String{}),
		"gcp_attributes":   reflect.TypeOf(compute_tf.GcpAttributes_SdkV2{}),
		"init_scripts":     reflect.TypeOf(compute_tf.InitScriptInfo_SdkV2{}),
		"spark_conf":       reflect.TypeOf(types.String{}),
		"spark_env_vars":   reflect.TypeOf(types.String{}),
		"ssh_public_keys":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: PipelineClusterAutoscale{}.Type(ctx),
			},
			"aws_attributes": basetypes.ListType{
				ElemType: compute_tf.AwsAttributes{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: compute_tf.AzureAttributes{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: compute_tf.ClusterLogConf{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: compute_tf.GcpAttributes{}.Type(ctx),
			},
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

// GetAutoscale returns the value of the Autoscale field in PipelineCluster_SdkV2 as
// a PipelineClusterAutoscale_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetAutoscale(ctx context.Context) (PipelineClusterAutoscale_SdkV2, bool) {
	var e PipelineClusterAutoscale_SdkV2
	if o.Autoscale.IsNull() || o.Autoscale.IsUnknown() {
		return e, false
	}
	var v []PipelineClusterAutoscale_SdkV2
	d := o.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetAutoscale(ctx context.Context, v PipelineClusterAutoscale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	o.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetAwsAttributes(ctx context.Context) (compute_tf.AwsAttributes_SdkV2, bool) {
	var e compute_tf.AwsAttributes_SdkV2
	if o.AwsAttributes.IsNull() || o.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AwsAttributes_SdkV2
	d := o.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetAwsAttributes(ctx context.Context, v compute_tf.AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	o.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetAzureAttributes(ctx context.Context) (compute_tf.AzureAttributes_SdkV2, bool) {
	var e compute_tf.AzureAttributes_SdkV2
	if o.AzureAttributes.IsNull() || o.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AzureAttributes_SdkV2
	d := o.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetAzureAttributes(ctx context.Context, v compute_tf.AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	o.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in PipelineCluster_SdkV2 as
// a compute_tf.ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetClusterLogConf(ctx context.Context) (compute_tf.ClusterLogConf_SdkV2, bool) {
	var e compute_tf.ClusterLogConf_SdkV2
	if o.ClusterLogConf.IsNull() || o.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterLogConf_SdkV2
	d := o.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetClusterLogConf(ctx context.Context, v compute_tf.ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	o.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetGcpAttributes(ctx context.Context) (compute_tf.GcpAttributes_SdkV2, bool) {
	var e compute_tf.GcpAttributes_SdkV2
	if o.GcpAttributes.IsNull() || o.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.GcpAttributes_SdkV2
	d := o.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetGcpAttributes(ctx context.Context, v compute_tf.GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	o.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in PipelineCluster_SdkV2 as
// a slice of compute_tf.InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetInitScripts(ctx context.Context) ([]compute_tf.InitScriptInfo_SdkV2, bool) {
	if o.InitScripts.IsNull() || o.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.InitScriptInfo_SdkV2
	d := o.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetInitScripts(ctx context.Context, v []compute_tf.InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in PipelineCluster_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineCluster_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in PipelineCluster_SdkV2.
func (o *PipelineCluster_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SshPublicKeys = types.ListValueMust(t, vs)
}

type PipelineClusterAutoscale_SdkV2 struct {
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

func (newState *PipelineClusterAutoscale_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineClusterAutoscale_SdkV2) {
}

func (newState *PipelineClusterAutoscale_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineClusterAutoscale_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineClusterAutoscale.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineClusterAutoscale_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineClusterAutoscale_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineClusterAutoscale_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": o.MaxWorkers,
			"min_workers": o.MinWorkers,
			"mode":        o.Mode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineClusterAutoscale_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_workers": types.Int64Type,
			"min_workers": types.Int64Type,
			"mode":        types.StringType,
		},
	}
}

type PipelineDeployment_SdkV2 struct {
	// The deployment method that manages the pipeline.
	Kind types.String `tfsdk:"kind" tf:"optional"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path" tf:"optional"`
}

func (newState *PipelineDeployment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineDeployment_SdkV2) {
}

func (newState *PipelineDeployment_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineDeployment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineDeployment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineDeployment_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineDeployment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kind":               o.Kind,
			"metadata_file_path": o.MetadataFilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineDeployment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kind":               types.StringType,
			"metadata_file_path": types.StringType,
		},
	}
}

type PipelineEvent_SdkV2 struct {
	// Information about an error captured by the event.
	Error types.List `tfsdk:"error" tf:"optional,object"`
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
	Origin types.List `tfsdk:"origin" tf:"optional,object"`
	// A sequencing object to identify and order events.
	Sequence types.List `tfsdk:"sequence" tf:"optional,object"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *PipelineEvent_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineEvent_SdkV2) {
}

func (newState *PipelineEvent_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineEvent_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineEvent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineEvent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error":    reflect.TypeOf(ErrorDetail_SdkV2{}),
		"origin":   reflect.TypeOf(Origin_SdkV2{}),
		"sequence": reflect.TypeOf(Sequencing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineEvent_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineEvent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineEvent_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": basetypes.ListType{
				ElemType: ErrorDetail{}.Type(ctx),
			},
			"event_type":     types.StringType,
			"id":             types.StringType,
			"level":          types.StringType,
			"maturity_level": types.StringType,
			"message":        types.StringType,
			"origin": basetypes.ListType{
				ElemType: Origin{}.Type(ctx),
			},
			"sequence": basetypes.ListType{
				ElemType: Sequencing{}.Type(ctx),
			},
			"timestamp": types.StringType,
		},
	}
}

// GetError returns the value of the Error field in PipelineEvent_SdkV2 as
// a ErrorDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent_SdkV2) GetError(ctx context.Context) (ErrorDetail_SdkV2, bool) {
	var e ErrorDetail_SdkV2
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []ErrorDetail_SdkV2
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in PipelineEvent_SdkV2.
func (o *PipelineEvent_SdkV2) SetError(ctx context.Context, v ErrorDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

// GetOrigin returns the value of the Origin field in PipelineEvent_SdkV2 as
// a Origin_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent_SdkV2) GetOrigin(ctx context.Context) (Origin_SdkV2, bool) {
	var e Origin_SdkV2
	if o.Origin.IsNull() || o.Origin.IsUnknown() {
		return e, false
	}
	var v []Origin_SdkV2
	d := o.Origin.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOrigin sets the value of the Origin field in PipelineEvent_SdkV2.
func (o *PipelineEvent_SdkV2) SetOrigin(ctx context.Context, v Origin_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["origin"]
	o.Origin = types.ListValueMust(t, vs)
}

// GetSequence returns the value of the Sequence field in PipelineEvent_SdkV2 as
// a Sequencing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineEvent_SdkV2) GetSequence(ctx context.Context) (Sequencing_SdkV2, bool) {
	var e Sequencing_SdkV2
	if o.Sequence.IsNull() || o.Sequence.IsUnknown() {
		return e, false
	}
	var v []Sequencing_SdkV2
	d := o.Sequence.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSequence sets the value of the Sequence field in PipelineEvent_SdkV2.
func (o *PipelineEvent_SdkV2) SetSequence(ctx context.Context, v Sequencing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence"]
	o.Sequence = types.ListValueMust(t, vs)
}

type PipelineLibrary_SdkV2 struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File types.List `tfsdk:"file" tf:"optional,object"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed.
	Maven types.List `tfsdk:"maven" tf:"optional,object"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook types.List `tfsdk:"notebook" tf:"optional,object"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

func (newState *PipelineLibrary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineLibrary_SdkV2) {
}

func (newState *PipelineLibrary_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineLibrary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineLibrary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file":     reflect.TypeOf(FileLibrary_SdkV2{}),
		"maven":    reflect.TypeOf(compute_tf.MavenLibrary_SdkV2{}),
		"notebook": reflect.TypeOf(NotebookLibrary_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineLibrary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file": basetypes.ListType{
				ElemType: FileLibrary{}.Type(ctx),
			},
			"jar": types.StringType,
			"maven": basetypes.ListType{
				ElemType: compute_tf.MavenLibrary{}.Type(ctx),
			},
			"notebook": basetypes.ListType{
				ElemType: NotebookLibrary{}.Type(ctx),
			},
			"whl": types.StringType,
		},
	}
}

// GetFile returns the value of the File field in PipelineLibrary_SdkV2 as
// a FileLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary_SdkV2) GetFile(ctx context.Context) (FileLibrary_SdkV2, bool) {
	var e FileLibrary_SdkV2
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []FileLibrary_SdkV2
	d := o.File.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in PipelineLibrary_SdkV2.
func (o *PipelineLibrary_SdkV2) SetFile(ctx context.Context, v FileLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file"]
	o.File = types.ListValueMust(t, vs)
}

// GetMaven returns the value of the Maven field in PipelineLibrary_SdkV2 as
// a compute_tf.MavenLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary_SdkV2) GetMaven(ctx context.Context) (compute_tf.MavenLibrary_SdkV2, bool) {
	var e compute_tf.MavenLibrary_SdkV2
	if o.Maven.IsNull() || o.Maven.IsUnknown() {
		return e, false
	}
	var v []compute_tf.MavenLibrary_SdkV2
	d := o.Maven.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaven sets the value of the Maven field in PipelineLibrary_SdkV2.
func (o *PipelineLibrary_SdkV2) SetMaven(ctx context.Context, v compute_tf.MavenLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["maven"]
	o.Maven = types.ListValueMust(t, vs)
}

// GetNotebook returns the value of the Notebook field in PipelineLibrary_SdkV2 as
// a NotebookLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary_SdkV2) GetNotebook(ctx context.Context) (NotebookLibrary_SdkV2, bool) {
	var e NotebookLibrary_SdkV2
	if o.Notebook.IsNull() || o.Notebook.IsUnknown() {
		return e, false
	}
	var v []NotebookLibrary_SdkV2
	d := o.Notebook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebook sets the value of the Notebook field in PipelineLibrary_SdkV2.
func (o *PipelineLibrary_SdkV2) SetNotebook(ctx context.Context, v NotebookLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook"]
	o.Notebook = types.ListValueMust(t, vs)
}

type PipelinePermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PipelinePermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermission_SdkV2) {
}

func (newState *PipelinePermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelinePermission_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermission_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelinePermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in PipelinePermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in PipelinePermission_SdkV2.
func (o *PipelinePermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type PipelinePermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *PipelinePermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissions_SdkV2) {
}

func (newState *PipelinePermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelinePermissions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelinePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissions_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAccessControlList returns the value of the AccessControlList field in PipelinePermissions_SdkV2 as
// a slice of PipelineAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissions_SdkV2.
func (o *PipelinePermissions_SdkV2) SetAccessControlList(ctx context.Context, v []PipelineAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PipelinePermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PipelinePermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissionsDescription_SdkV2) {
}

func (newState *PipelinePermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelinePermissionsDescription_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelinePermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PipelinePermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *PipelinePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissionsRequest_SdkV2) {
}

func (newState *PipelinePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelinePermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelinePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"pipeline_id":         o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PipelineAccessControlRequest{}.Type(ctx),
			},
			"pipeline_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in PipelinePermissionsRequest_SdkV2 as
// a slice of PipelineAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinePermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissionsRequest_SdkV2.
func (o *PipelinePermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []PipelineAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type PipelineSpec_SdkV2 struct {
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
	Deployment types.List `tfsdk:"deployment" tf:"optional,object"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development" tf:"optional"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition" tf:"optional"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters" tf:"optional,object"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition" tf:"optional,object"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition" tf:"optional,object"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries" tf:"optional"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications" tf:"optional"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon" tf:"optional"`
	// Restart window of this pipeline.
	RestartWindow types.List `tfsdk:"restart_window" tf:"optional,object"`
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
	Trigger types.List `tfsdk:"trigger" tf:"optional,object"`
}

func (newState *PipelineSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineSpec_SdkV2) {
}

func (newState *PipelineSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster_SdkV2{}),
		"configuration":        reflect.TypeOf(types.String{}),
		"deployment":           reflect.TypeOf(PipelineDeployment_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"trigger":              reflect.TypeOf(PipelineTrigger_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.Type(ctx),
			},
			"development": types.BoolType,
			"edition":     types.StringType,
			"filters": basetypes.ListType{
				ElemType: Filters{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.Type(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.Type(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.Type(ctx),
			},
		},
	}
}

// GetClusters returns the value of the Clusters field in PipelineSpec_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if o.Clusters.IsNull() || o.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := o.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in PipelineSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in PipelineSpec_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := o.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	o.Deployment = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in PipelineSpec_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := o.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	o.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in PipelineSpec_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if o.GatewayDefinition.IsNull() || o.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := o.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	o.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in PipelineSpec_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if o.IngestionDefinition.IsNull() || o.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := o.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	o.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in PipelineSpec_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in PipelineSpec_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in PipelineSpec_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if o.RestartWindow.IsNull() || o.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := o.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	o.RestartWindow = types.ListValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in PipelineSpec_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := o.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	o.Trigger = types.ListValueMust(t, vs)
}

type PipelineStateInfo_SdkV2 struct {
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

func (newState *PipelineStateInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineStateInfo_SdkV2) {
}

func (newState *PipelineStateInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineStateInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineStateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineStateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineStateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineStateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineStateInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetLatestUpdates returns the value of the LatestUpdates field in PipelineStateInfo_SdkV2 as
// a slice of UpdateStateInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineStateInfo_SdkV2) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo_SdkV2, bool) {
	if o.LatestUpdates.IsNull() || o.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo_SdkV2
	d := o.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in PipelineStateInfo_SdkV2.
func (o *PipelineStateInfo_SdkV2) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestUpdates = types.ListValueMust(t, vs)
}

type PipelineTrigger_SdkV2 struct {
	Cron types.List `tfsdk:"cron" tf:"optional,object"`

	Manual types.List `tfsdk:"manual" tf:"optional,object"`
}

func (newState *PipelineTrigger_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineTrigger_SdkV2) {
}

func (newState *PipelineTrigger_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineTrigger_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron":   reflect.TypeOf(CronTrigger_SdkV2{}),
		"manual": reflect.TypeOf(ManualTrigger_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cron":   o.Cron,
			"manual": o.Manual,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineTrigger_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cron": basetypes.ListType{
				ElemType: CronTrigger{}.Type(ctx),
			},
			"manual": basetypes.ListType{
				ElemType: ManualTrigger{}.Type(ctx),
			},
		},
	}
}

// GetCron returns the value of the Cron field in PipelineTrigger_SdkV2 as
// a CronTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineTrigger_SdkV2) GetCron(ctx context.Context) (CronTrigger_SdkV2, bool) {
	var e CronTrigger_SdkV2
	if o.Cron.IsNull() || o.Cron.IsUnknown() {
		return e, false
	}
	var v []CronTrigger_SdkV2
	d := o.Cron.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCron sets the value of the Cron field in PipelineTrigger_SdkV2.
func (o *PipelineTrigger_SdkV2) SetCron(ctx context.Context, v CronTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cron"]
	o.Cron = types.ListValueMust(t, vs)
}

// GetManual returns the value of the Manual field in PipelineTrigger_SdkV2 as
// a ManualTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineTrigger_SdkV2) GetManual(ctx context.Context) (ManualTrigger_SdkV2, bool) {
	var e ManualTrigger_SdkV2
	if o.Manual.IsNull() || o.Manual.IsUnknown() {
		return e, false
	}
	var v []ManualTrigger_SdkV2
	d := o.Manual.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManual sets the value of the Manual field in PipelineTrigger_SdkV2.
func (o *PipelineTrigger_SdkV2) SetManual(ctx context.Context, v ManualTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["manual"]
	o.Manual = types.ListValueMust(t, vs)
}

type ReportSpec_SdkV2 struct {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *ReportSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReportSpec_SdkV2) {
}

func (newState *ReportSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState ReportSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReportSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReportSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReportSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o ReportSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ReportSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_url":          types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.Type(ctx),
			},
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in ReportSpec_SdkV2 as
// a TableSpecificConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ReportSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := o.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in ReportSpec_SdkV2.
func (o *ReportSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	o.TableConfiguration = types.ListValueMust(t, vs)
}

type RestartWindow_SdkV2 struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	DaysOfWeek types.String `tfsdk:"days_of_week" tf:"optional"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour types.Int64 `tfsdk:"start_hour" tf:""`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId types.String `tfsdk:"time_zone_id" tf:"optional"`
}

func (newState *RestartWindow_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestartWindow_SdkV2) {
}

func (newState *RestartWindow_SdkV2) SyncEffectiveFieldsDuringRead(existingState RestartWindow_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartWindow_SdkV2
// only implements ToObjectValue() and Type().
func (o RestartWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"days_of_week": o.DaysOfWeek,
			"start_hour":   o.StartHour,
			"time_zone_id": o.TimeZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestartWindow_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"days_of_week": types.StringType,
			"start_hour":   types.Int64Type,
			"time_zone_id": types.StringType,
		},
	}
}

type SchemaSpec_SdkV2 struct {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *SchemaSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaSpec_SdkV2) {
}

func (newState *SchemaSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState SchemaSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SchemaSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SchemaSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o SchemaSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SchemaSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.Type(ctx),
			},
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in SchemaSpec_SdkV2 as
// a TableSpecificConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := o.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in SchemaSpec_SdkV2.
func (o *SchemaSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	o.TableConfiguration = types.ListValueMust(t, vs)
}

type Sequencing_SdkV2 struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no" tf:"optional"`
	// the ID assigned by the data plane.
	DataPlaneId types.List `tfsdk:"data_plane_id" tf:"optional,object"`
}

func (newState *Sequencing_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Sequencing_SdkV2) {
}

func (newState *Sequencing_SdkV2) SyncEffectiveFieldsDuringRead(existingState Sequencing_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Sequencing.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Sequencing_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_id": reflect.TypeOf(DataPlaneId_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sequencing_SdkV2
// only implements ToObjectValue() and Type().
func (o Sequencing_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"control_plane_seq_no": o.ControlPlaneSeqNo,
			"data_plane_id":        o.DataPlaneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Sequencing_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"control_plane_seq_no": types.Int64Type,
			"data_plane_id": basetypes.ListType{
				ElemType: DataPlaneId{}.Type(ctx),
			},
		},
	}
}

// GetDataPlaneId returns the value of the DataPlaneId field in Sequencing_SdkV2 as
// a DataPlaneId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Sequencing_SdkV2) GetDataPlaneId(ctx context.Context) (DataPlaneId_SdkV2, bool) {
	var e DataPlaneId_SdkV2
	if o.DataPlaneId.IsNull() || o.DataPlaneId.IsUnknown() {
		return e, false
	}
	var v []DataPlaneId_SdkV2
	d := o.DataPlaneId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneId sets the value of the DataPlaneId field in Sequencing_SdkV2.
func (o *Sequencing_SdkV2) SetDataPlaneId(ctx context.Context, v DataPlaneId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_plane_id"]
	o.DataPlaneId = types.ListValueMust(t, vs)
}

type SerializedException_SdkV2 struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name" tf:"optional"`
	// Exception message
	Message types.String `tfsdk:"message" tf:"optional"`
	// Stack trace consisting of a list of stack frames
	Stack types.List `tfsdk:"stack" tf:"optional"`
}

func (newState *SerializedException_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SerializedException_SdkV2) {
}

func (newState *SerializedException_SdkV2) SyncEffectiveFieldsDuringRead(existingState SerializedException_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SerializedException.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SerializedException_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stack": reflect.TypeOf(StackFrame_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SerializedException_SdkV2
// only implements ToObjectValue() and Type().
func (o SerializedException_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"class_name": o.ClassName,
			"message":    o.Message,
			"stack":      o.Stack,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SerializedException_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetStack returns the value of the Stack field in SerializedException_SdkV2 as
// a slice of StackFrame_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SerializedException_SdkV2) GetStack(ctx context.Context) ([]StackFrame_SdkV2, bool) {
	if o.Stack.IsNull() || o.Stack.IsUnknown() {
		return nil, false
	}
	var v []StackFrame_SdkV2
	d := o.Stack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStack sets the value of the Stack field in SerializedException_SdkV2.
func (o *SerializedException_SdkV2) SetStack(ctx context.Context, v []StackFrame_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stack"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stack = types.ListValueMust(t, vs)
}

type StackFrame_SdkV2 struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class" tf:"optional"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name" tf:"optional"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number" tf:"optional"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name" tf:"optional"`
}

func (newState *StackFrame_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StackFrame_SdkV2) {
}

func (newState *StackFrame_SdkV2) SyncEffectiveFieldsDuringRead(existingState StackFrame_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StackFrame.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StackFrame_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StackFrame_SdkV2
// only implements ToObjectValue() and Type().
func (o StackFrame_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StackFrame_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"declaring_class": types.StringType,
			"file_name":       types.StringType,
			"line_number":     types.Int64Type,
			"method_name":     types.StringType,
		},
	}
}

type StartUpdate_SdkV2 struct {
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

func (newState *StartUpdate_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartUpdate_SdkV2) {
}

func (newState *StartUpdate_SdkV2) SyncEffectiveFieldsDuringRead(existingState StartUpdate_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (o StartUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StartUpdate_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in StartUpdate_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *StartUpdate_SdkV2) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in StartUpdate_SdkV2.
func (o *StartUpdate_SdkV2) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in StartUpdate_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *StartUpdate_SdkV2) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetRefreshSelection sets the value of the RefreshSelection field in StartUpdate_SdkV2.
func (o *StartUpdate_SdkV2) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RefreshSelection = types.ListValueMust(t, vs)
}

type StartUpdateResponse_SdkV2 struct {
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *StartUpdateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartUpdateResponse_SdkV2) {
}

func (newState *StartUpdateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState StartUpdateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartUpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartUpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StartUpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_id": o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartUpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update_id": types.StringType,
		},
	}
}

type StopPipelineResponse_SdkV2 struct {
}

func (newState *StopPipelineResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopPipelineResponse_SdkV2) {
}

func (newState *StopPipelineResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState StopPipelineResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o StopPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o StopPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Stop a pipeline
type StopRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (newState *StopRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopRequest_SdkV2) {
}

func (newState *StopRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState StopRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StopRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StopRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StopRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type TableSpec_SdkV2 struct {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *TableSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSpec_SdkV2) {
}

func (newState *TableSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o TableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"source_table":        types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.Type(ctx),
			},
		},
	}
}

// GetTableConfiguration returns the value of the TableConfiguration field in TableSpec_SdkV2 as
// a TableSpecificConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := o.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in TableSpec_SdkV2.
func (o *TableSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	o.TableConfiguration = types.ListValueMust(t, vs)
}

type TableSpecificConfig_SdkV2 struct {
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

func (newState *TableSpecificConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSpecificConfig_SdkV2) {
}

func (newState *TableSpecificConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableSpecificConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSpecificConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSpecificConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
		"sequence_by":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpecificConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o TableSpecificConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TableSpecificConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetPrimaryKeys returns the value of the PrimaryKeys field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig_SdkV2) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeys sets the value of the PrimaryKeys field in TableSpecificConfig_SdkV2.
func (o *TableSpecificConfig_SdkV2) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeys = types.ListValueMust(t, vs)
}

// GetSequenceBy returns the value of the SequenceBy field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig_SdkV2) GetSequenceBy(ctx context.Context) ([]types.String, bool) {
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

// SetSequenceBy sets the value of the SequenceBy field in TableSpecificConfig_SdkV2.
func (o *TableSpecificConfig_SdkV2) SetSequenceBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SequenceBy = types.ListValueMust(t, vs)
}

type UpdateInfo_SdkV2 struct {
	// What triggered this update.
	Cause types.String `tfsdk:"cause" tf:"optional"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config types.List `tfsdk:"config" tf:"optional,object"`
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

func (newState *UpdateInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInfo_SdkV2) {
}

func (newState *UpdateInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":                 reflect.TypeOf(PipelineSpec_SdkV2{}),
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":      types.StringType,
			"cluster_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: PipelineSpec{}.Type(ctx),
			},
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

// GetConfig returns the value of the Config field in UpdateInfo_SdkV2 as
// a PipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo_SdkV2) GetConfig(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in UpdateInfo_SdkV2.
func (o *UpdateInfo_SdkV2) SetConfig(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
}

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in UpdateInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo_SdkV2) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in UpdateInfo_SdkV2.
func (o *UpdateInfo_SdkV2) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in UpdateInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInfo_SdkV2) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetRefreshSelection sets the value of the RefreshSelection field in UpdateInfo_SdkV2.
func (o *UpdateInfo_SdkV2) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RefreshSelection = types.ListValueMust(t, vs)
}

type UpdateStateInfo_SdkV2 struct {
	CreationTime types.String `tfsdk:"creation_time" tf:"optional"`

	State types.String `tfsdk:"state" tf:"optional"`

	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *UpdateStateInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStateInfo_SdkV2) {
}

func (newState *UpdateStateInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateStateInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateStateInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateStateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateStateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time": o.CreationTime,
			"state":         o.State,
			"update_id":     o.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateStateInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.StringType,
			"state":         types.StringType,
			"update_id":     types.StringType,
		},
	}
}
