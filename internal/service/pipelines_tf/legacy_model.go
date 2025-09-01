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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreatePipeline_SdkV2 struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment types.List `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`

	DryRun types.Bool `tfsdk:"dry_run"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.List `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.List `tfsdk:"event_log"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Restart window of this pipeline.
	RestartWindow types.List `tfsdk:"restart_window"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath types.String `tfsdk:"root_path"`

	RunAs types.List `tfsdk:"run_as"`
	// The default schema (database) where tables are read from or published to.
	Schema types.String `tfsdk:"schema"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags types.Map `tfsdk:"tags"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.List `tfsdk:"trigger"`
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
		"environment":          reflect.TypeOf(PipelinesEnvironment_SdkV2{}),
		"event_log":            reflect.TypeOf(EventLogSpec_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"run_as":               reflect.TypeOf(RunAs_SdkV2{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":           o.Environment,
			"event_log":             o.EventLog,
			"filters":               o.Filters,
			"gateway_definition":    o.GatewayDefinition,
			"id":                    o.Id,
			"ingestion_definition":  o.IngestionDefinition,
			"libraries":             o.Libraries,
			"name":                  o.Name,
			"notifications":         o.Notifications,
			"photon":                o.Photon,
			"restart_window":        o.RestartWindow,
			"root_path":             o.RootPath,
			"run_as":                o.RunAs,
			"schema":                o.Schema,
			"serverless":            o.Serverless,
			"storage":               o.Storage,
			"tags":                  o.Tags,
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
				ElemType: PipelineCluster_SdkV2{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment_SdkV2{}.Type(ctx),
			},
			"development": types.BoolType,
			"dry_run":     types.BoolType,
			"edition":     types.StringType,
			"environment": basetypes.ListType{
				ElemType: PipelinesEnvironment_SdkV2{}.Type(ctx),
			},
			"event_log": basetypes.ListType{
				ElemType: EventLogSpec_SdkV2{}.Type(ctx),
			},
			"filters": basetypes.ListType{
				ElemType: Filters_SdkV2{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications_SdkV2{}.Type(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow_SdkV2{}.Type(ctx),
			},
			"root_path": types.StringType,
			"run_as": basetypes.ListType{
				ElemType: RunAs_SdkV2{}.Type(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target": types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger_SdkV2{}.Type(ctx),
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

// GetEnvironment returns the value of the Environment field in CreatePipeline_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := o.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	o.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in CreatePipeline_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := o.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	o.EventLog = types.ListValueMust(t, vs)
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

// GetRunAs returns the value of the RunAs field in CreatePipeline_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreatePipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreatePipeline_SdkV2.
func (o *CreatePipeline_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
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
	EffectiveSettings types.List `tfsdk:"effective_settings"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (toState *CreatePipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreatePipelineResponse_SdkV2) {
	if !fromPlan.EffectiveSettings.IsNull() && !fromPlan.EffectiveSettings.IsUnknown() {
		if toStateEffectiveSettings, ok := toState.GetEffectiveSettings(ctx); ok {
			if fromPlanEffectiveSettings, ok := fromPlan.GetEffectiveSettings(ctx); ok {
				toStateEffectiveSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEffectiveSettings)
				toState.SetEffectiveSettings(ctx, toStateEffectiveSettings)
			}
		}
	}
}

func (toState *CreatePipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreatePipelineResponse_SdkV2) {
	if !fromState.EffectiveSettings.IsNull() && !fromState.EffectiveSettings.IsUnknown() {
		if toStateEffectiveSettings, ok := toState.GetEffectiveSettings(ctx); ok {
			if fromStateEffectiveSettings, ok := fromState.GetEffectiveSettings(ctx); ok {
				toStateEffectiveSettings.SyncFieldsDuringRead(ctx, fromStateEffectiveSettings)
				toState.SetEffectiveSettings(ctx, toStateEffectiveSettings)
			}
		}
	}
}

func (c CreatePipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_settings"] = attrs["effective_settings"].SetOptional()
	attrs["effective_settings"] = attrs["effective_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()

	return attrs
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
				ElemType: PipelineSpec_SdkV2{}.Type(ctx),
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
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`

	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (toState *CronTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CronTrigger_SdkV2) {
}

func (toState *CronTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CronTrigger_SdkV2) {
}

func (c CronTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quartz_cron_schedule"] = attrs["quartz_cron_schedule"].SetOptional()
	attrs["timezone_id"] = attrs["timezone_id"].SetOptional()

	return attrs
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
	Instance types.String `tfsdk:"instance"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no"`
}

func (toState *DataPlaneId_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DataPlaneId_SdkV2) {
}

func (toState *DataPlaneId_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DataPlaneId_SdkV2) {
}

func (c DataPlaneId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance"] = attrs["instance"].SetOptional()
	attrs["seq_no"] = attrs["seq_no"].SetOptional()

	return attrs
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

type DeletePipelineRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
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

func (toState *DeletePipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeletePipelineResponse_SdkV2) {
}

func (toState *DeletePipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeletePipelineResponse_SdkV2) {
}

func (c DeletePipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	AllowDuplicateNames types.Bool `tfsdk:"allow_duplicate_names"`
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment types.List `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.List `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.List `tfsdk:"event_log"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Unique identifier for this pipeline.
	PipelineId types.String `tfsdk:"-"`
	// Restart window of this pipeline.
	RestartWindow types.List `tfsdk:"restart_window"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath types.String `tfsdk:"root_path"`

	RunAs types.List `tfsdk:"run_as"`
	// The default schema (database) where tables are read from or published to.
	Schema types.String `tfsdk:"schema"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags types.Map `tfsdk:"tags"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.List `tfsdk:"trigger"`
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
		"environment":          reflect.TypeOf(PipelinesEnvironment_SdkV2{}),
		"event_log":            reflect.TypeOf(EventLogSpec_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"run_as":               reflect.TypeOf(RunAs_SdkV2{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":            o.Environment,
			"event_log":              o.EventLog,
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
			"root_path":              o.RootPath,
			"run_as":                 o.RunAs,
			"schema":                 o.Schema,
			"serverless":             o.Serverless,
			"storage":                o.Storage,
			"tags":                   o.Tags,
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
				ElemType: PipelineCluster_SdkV2{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment_SdkV2{}.Type(ctx),
			},
			"development": types.BoolType,
			"edition":     types.StringType,
			"environment": basetypes.ListType{
				ElemType: PipelinesEnvironment_SdkV2{}.Type(ctx),
			},
			"event_log": basetypes.ListType{
				ElemType: EventLogSpec_SdkV2{}.Type(ctx),
			},
			"expected_last_modified": types.Int64Type,
			"filters": basetypes.ListType{
				ElemType: Filters_SdkV2{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications_SdkV2{}.Type(ctx),
			},
			"photon":      types.BoolType,
			"pipeline_id": types.StringType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow_SdkV2{}.Type(ctx),
			},
			"root_path": types.StringType,
			"run_as": basetypes.ListType{
				ElemType: RunAs_SdkV2{}.Type(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target": types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger_SdkV2{}.Type(ctx),
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

// GetEnvironment returns the value of the Environment field in EditPipeline_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := o.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	o.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in EditPipeline_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := o.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	o.EventLog = types.ListValueMust(t, vs)
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

// GetRunAs returns the value of the RunAs field in EditPipeline_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EditPipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EditPipeline_SdkV2.
func (o *EditPipeline_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
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

func (toState *EditPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan EditPipelineResponse_SdkV2) {
}

func (toState *EditPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState EditPipelineResponse_SdkV2) {
}

func (c EditPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Exceptions types.List `tfsdk:"exceptions"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal"`
}

func (toState *ErrorDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ErrorDetail_SdkV2) {
}

func (toState *ErrorDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ErrorDetail_SdkV2) {
}

func (c ErrorDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exceptions"] = attrs["exceptions"].SetOptional()
	attrs["fatal"] = attrs["fatal"].SetOptional()

	return attrs
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
				ElemType: SerializedException_SdkV2{}.Type(ctx),
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

// Configurable event log parameters.
type EventLogSpec_SdkV2 struct {
	// The UC catalog the event log is published under.
	Catalog types.String `tfsdk:"catalog"`
	// The name the event log is published to in UC.
	Name types.String `tfsdk:"name"`
	// The UC schema the event log is published under.
	Schema types.String `tfsdk:"schema"`
}

func (toState *EventLogSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan EventLogSpec_SdkV2) {
}

func (toState *EventLogSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState EventLogSpec_SdkV2) {
}

func (c EventLogSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EventLogSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EventLogSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventLogSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o EventLogSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
			"name":    o.Name,
			"schema":  o.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EventLogSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": types.StringType,
			"name":    types.StringType,
			"schema":  types.StringType,
		},
	}
}

type FileLibrary_SdkV2 struct {
	// The absolute path of the source code.
	Path types.String `tfsdk:"path"`
}

func (toState *FileLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FileLibrary_SdkV2) {
}

func (toState *FileLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FileLibrary_SdkV2) {
}

func (c FileLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
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
	Exclude types.List `tfsdk:"exclude"`
	// Paths to include.
	Include types.List `tfsdk:"include"`
}

func (toState *Filters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Filters_SdkV2) {
}

func (toState *Filters_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Filters_SdkV2) {
}

func (c Filters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exclude"] = attrs["exclude"].SetOptional()
	attrs["include"] = attrs["include"].SetOptional()

	return attrs
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

type GetPipelinePermissionLevelsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
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
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (toState *GetPipelinePermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPipelinePermissionLevelsResponse_SdkV2) {
}

func (toState *GetPipelinePermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetPipelinePermissionLevelsResponse_SdkV2) {
}

func (c GetPipelinePermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
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
				ElemType: PipelinePermissionsDescription_SdkV2{}.Type(ctx),
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

type GetPipelinePermissionsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
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

type GetPipelineRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
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
	Cause types.String `tfsdk:"cause"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health"`
	// The last time the pipeline settings were modified or created.
	LastModified types.Int64 `tfsdk:"last_modified"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates types.List `tfsdk:"latest_updates"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name types.String `tfsdk:"name"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// The user or service principal that the pipeline runs as, if specified in
	// the request. This field indicates the explicit configuration of `run_as`
	// for the pipeline. To find the value in all cases, explicit or implicit,
	// use `run_as_user_name`.
	RunAs types.List `tfsdk:"run_as"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec types.List `tfsdk:"spec"`
	// The pipeline state.
	State types.String `tfsdk:"state"`
}

func (toState *GetPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetPipelineResponse_SdkV2) {
	if !fromPlan.RunAs.IsNull() && !fromPlan.RunAs.IsUnknown() {
		if toStateRunAs, ok := toState.GetRunAs(ctx); ok {
			if fromPlanRunAs, ok := fromPlan.GetRunAs(ctx); ok {
				toStateRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRunAs)
				toState.SetRunAs(ctx, toStateRunAs)
			}
		}
	}
	if !fromPlan.Spec.IsNull() && !fromPlan.Spec.IsUnknown() {
		if toStateSpec, ok := toState.GetSpec(ctx); ok {
			if fromPlanSpec, ok := fromPlan.GetSpec(ctx); ok {
				toStateSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSpec)
				toState.SetSpec(ctx, toStateSpec)
			}
		}
	}
}

func (toState *GetPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetPipelineResponse_SdkV2) {
	if !fromState.RunAs.IsNull() && !fromState.RunAs.IsUnknown() {
		if toStateRunAs, ok := toState.GetRunAs(ctx); ok {
			if fromStateRunAs, ok := fromState.GetRunAs(ctx); ok {
				toStateRunAs.SyncFieldsDuringRead(ctx, fromStateRunAs)
				toState.SetRunAs(ctx, toStateRunAs)
			}
		}
	}
	if !fromState.Spec.IsNull() && !fromState.Spec.IsUnknown() {
		if toStateSpec, ok := toState.GetSpec(ctx); ok {
			if fromStateSpec, ok := fromState.GetSpec(ctx); ok {
				toStateSpec.SyncFieldsDuringRead(ctx, fromStateSpec)
				toState.SetSpec(ctx, toStateSpec)
			}
		}
	}
}

func (c GetPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["last_modified"] = attrs["last_modified"].SetOptional()
	attrs["latest_updates"] = attrs["latest_updates"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as"] = attrs["run_as"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
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
		"run_as":         reflect.TypeOf(RunAs_SdkV2{}),
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
			"run_as":                     o.RunAs,
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
				ElemType: UpdateStateInfo_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"pipeline_id": types.StringType,
			"run_as": basetypes.ListType{
				ElemType: RunAs_SdkV2{}.Type(ctx),
			},
			"run_as_user_name": types.StringType,
			"spec": basetypes.ListType{
				ElemType: PipelineSpec_SdkV2{}.Type(ctx),
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

// GetRunAs returns the value of the RunAs field in GetPipelineResponse_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in GetPipelineResponse_SdkV2.
func (o *GetPipelineResponse_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
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

type GetUpdateRequest_SdkV2 struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
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
	Update types.List `tfsdk:"update"`
}

func (toState *GetUpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetUpdateResponse_SdkV2) {
	if !fromPlan.Update.IsNull() && !fromPlan.Update.IsUnknown() {
		if toStateUpdate, ok := toState.GetUpdate(ctx); ok {
			if fromPlanUpdate, ok := fromPlan.GetUpdate(ctx); ok {
				toStateUpdate.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanUpdate)
				toState.SetUpdate(ctx, toStateUpdate)
			}
		}
	}
}

func (toState *GetUpdateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetUpdateResponse_SdkV2) {
	if !fromState.Update.IsNull() && !fromState.Update.IsUnknown() {
		if toStateUpdate, ok := toState.GetUpdate(ctx); ok {
			if fromStateUpdate, ok := fromState.GetUpdate(ctx); ok {
				toStateUpdate.SyncFieldsDuringRead(ctx, fromStateUpdate)
				toState.SetUpdate(ctx, toStateUpdate)
			}
		}
	}
}

func (c GetUpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["update"] = attrs["update"].SetOptional()
	attrs["update"] = attrs["update"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: UpdateInfo_SdkV2{}.Type(ctx),
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
	Report types.List `tfsdk:"report"`
	// Select all tables from a specific source schema.
	Schema types.List `tfsdk:"schema"`
	// Select a specific source table.
	Table types.List `tfsdk:"table"`
}

func (toState *IngestionConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IngestionConfig_SdkV2) {
	if !fromPlan.Report.IsNull() && !fromPlan.Report.IsUnknown() {
		if toStateReport, ok := toState.GetReport(ctx); ok {
			if fromPlanReport, ok := fromPlan.GetReport(ctx); ok {
				toStateReport.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanReport)
				toState.SetReport(ctx, toStateReport)
			}
		}
	}
	if !fromPlan.Schema.IsNull() && !fromPlan.Schema.IsUnknown() {
		if toStateSchema, ok := toState.GetSchema(ctx); ok {
			if fromPlanSchema, ok := fromPlan.GetSchema(ctx); ok {
				toStateSchema.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSchema)
				toState.SetSchema(ctx, toStateSchema)
			}
		}
	}
	if !fromPlan.Table.IsNull() && !fromPlan.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromPlanTable, ok := fromPlan.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (toState *IngestionConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState IngestionConfig_SdkV2) {
	if !fromState.Report.IsNull() && !fromState.Report.IsUnknown() {
		if toStateReport, ok := toState.GetReport(ctx); ok {
			if fromStateReport, ok := fromState.GetReport(ctx); ok {
				toStateReport.SyncFieldsDuringRead(ctx, fromStateReport)
				toState.SetReport(ctx, toStateReport)
			}
		}
	}
	if !fromState.Schema.IsNull() && !fromState.Schema.IsUnknown() {
		if toStateSchema, ok := toState.GetSchema(ctx); ok {
			if fromStateSchema, ok := fromState.GetSchema(ctx); ok {
				toStateSchema.SyncFieldsDuringRead(ctx, fromStateSchema)
				toState.SetSchema(ctx, toStateSchema)
			}
		}
	}
	if !fromState.Table.IsNull() && !fromState.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromStateTable, ok := fromState.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringRead(ctx, fromStateTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (c IngestionConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["report"] = attrs["report"].SetOptional()
	attrs["report"] = attrs["report"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["schema"] = attrs["schema"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table"] = attrs["table"].SetOptional()
	attrs["table"] = attrs["table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: ReportSpec_SdkV2{}.Type(ctx),
			},
			"schema": basetypes.ListType{
				ElemType: SchemaSpec_SdkV2{}.Type(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableSpec_SdkV2{}.Type(ctx),
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
	ConnectionId types.String `tfsdk:"connection_id"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog types.String `tfsdk:"gateway_storage_catalog"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName types.String `tfsdk:"gateway_storage_name"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema types.String `tfsdk:"gateway_storage_schema"`
}

func (toState *IngestionGatewayPipelineDefinition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IngestionGatewayPipelineDefinition_SdkV2) {
}

func (toState *IngestionGatewayPipelineDefinition_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState IngestionGatewayPipelineDefinition_SdkV2) {
}

func (c IngestionGatewayPipelineDefinition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_id"] = attrs["connection_id"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetRequired()
	attrs["gateway_storage_catalog"] = attrs["gateway_storage_catalog"].SetRequired()
	attrs["gateway_storage_name"] = attrs["gateway_storage_name"].SetOptional()
	attrs["gateway_storage_schema"] = attrs["gateway_storage_schema"].SetRequired()

	return attrs
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
	ConnectionName types.String `tfsdk:"connection_name"`
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	IngestionGatewayId types.String `tfsdk:"ingestion_gateway_id"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects types.List `tfsdk:"objects"`
	// Top-level source configurations
	SourceConfigurations types.List `tfsdk:"source_configurations"`
	// The type of the foreign source. The source type will be inferred from the
	// source connection or ingestion gateway. This field is output only and
	// will be ignored if provided.
	SourceType types.String `tfsdk:"source_type"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration types.List `tfsdk:"table_configuration"`
}

func (toState *IngestionPipelineDefinition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IngestionPipelineDefinition_SdkV2) {
	if !fromPlan.TableConfiguration.IsNull() && !fromPlan.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromPlanTableConfiguration, ok := fromPlan.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (toState *IngestionPipelineDefinition_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState IngestionPipelineDefinition_SdkV2) {
	if !fromState.TableConfiguration.IsNull() && !fromState.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromStateTableConfiguration, ok := fromState.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringRead(ctx, fromStateTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (c IngestionPipelineDefinition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["ingestion_gateway_id"] = attrs["ingestion_gateway_id"].SetOptional()
	attrs["objects"] = attrs["objects"].SetOptional()
	attrs["source_configurations"] = attrs["source_configurations"].SetOptional()
	attrs["source_type"] = attrs["source_type"].SetComputed()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()
	attrs["table_configuration"] = attrs["table_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
		"objects":               reflect.TypeOf(IngestionConfig_SdkV2{}),
		"source_configurations": reflect.TypeOf(SourceConfig_SdkV2{}),
		"table_configuration":   reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name":       o.ConnectionName,
			"ingestion_gateway_id":  o.IngestionGatewayId,
			"objects":               o.Objects,
			"source_configurations": o.SourceConfigurations,
			"source_type":           o.SourceType,
			"table_configuration":   o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":      types.StringType,
			"ingestion_gateway_id": types.StringType,
			"objects": basetypes.ListType{
				ElemType: IngestionConfig_SdkV2{}.Type(ctx),
			},
			"source_configurations": basetypes.ListType{
				ElemType: SourceConfig_SdkV2{}.Type(ctx),
			},
			"source_type": types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig_SdkV2{}.Type(ctx),
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

// GetSourceConfigurations returns the value of the SourceConfigurations field in IngestionPipelineDefinition_SdkV2 as
// a slice of SourceConfig_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition_SdkV2) GetSourceConfigurations(ctx context.Context) ([]SourceConfig_SdkV2, bool) {
	if o.SourceConfigurations.IsNull() || o.SourceConfigurations.IsUnknown() {
		return nil, false
	}
	var v []SourceConfig_SdkV2
	d := o.SourceConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSourceConfigurations sets the value of the SourceConfigurations field in IngestionPipelineDefinition_SdkV2.
func (o *IngestionPipelineDefinition_SdkV2) SetSourceConfigurations(ctx context.Context, v []SourceConfig_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["source_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SourceConfigurations = types.ListValueMust(t, vs)
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

// Configurations that are only applicable for query-based ingestion connectors.
type IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2 struct {
	// The names of the monotonically increasing columns in the source table
	// that are used to enable the table to be read and ingested incrementally
	// through structured streaming. The columns are allowed to have repeated
	// values but have to be non-decreasing. If the source data is merged into
	// the destination (e.g., using SCD Type 1 or Type 2), these columns will
	// implicitly define the `sequence_by` behavior. You can still explicitly
	// set `sequence_by` to override this default.
	CursorColumns types.List `tfsdk:"cursor_columns"`
	// Specifies a SQL WHERE condition that specifies that the source row has
	// been deleted. This is sometimes referred to as "soft-deletes". For
	// example: "Operation = 'DELETE'" or "is_deleted = true". This field is
	// orthogonal to `hard_deletion_sync_interval_in_seconds`, one for
	// soft-deletes and the other for hard-deletes. See also the
	// hard_deletion_sync_min_interval_in_seconds field for handling of "hard
	// deletes" where the source rows are physically removed from the table.
	DeletionCondition types.String `tfsdk:"deletion_condition"`
	// Specifies the minimum interval (in seconds) between snapshots on primary
	// keys for detecting and synchronizing hard deletions—i.e., rows that
	// have been physically removed from the source table. This interval acts as
	// a lower bound. If ingestion runs less frequently than this value, hard
	// deletion synchronization will align with the actual ingestion frequency
	// instead of happening more often. If not set, hard deletion
	// synchronization via snapshots is disabled. This field is mutable and can
	// be updated without triggering a full snapshot.
	HardDeletionSyncMinIntervalInSeconds types.Int64 `tfsdk:"hard_deletion_sync_min_interval_in_seconds"`
}

func (toState *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
}

func (toState *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
}

func (c IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cursor_columns"] = attrs["cursor_columns"].SetOptional()
	attrs["deletion_condition"] = attrs["deletion_condition"].SetOptional()
	attrs["hard_deletion_sync_min_interval_in_seconds"] = attrs["hard_deletion_sync_min_interval_in_seconds"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cursor_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cursor_columns":                             o.CursorColumns,
			"deletion_condition":                         o.DeletionCondition,
			"hard_deletion_sync_min_interval_in_seconds": o.HardDeletionSyncMinIntervalInSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cursor_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"deletion_condition":                         types.StringType,
			"hard_deletion_sync_min_interval_in_seconds": types.Int64Type,
		},
	}
}

// GetCursorColumns returns the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) GetCursorColumns(ctx context.Context) ([]types.String, bool) {
	if o.CursorColumns.IsNull() || o.CursorColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.CursorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCursorColumns sets the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2.
func (o *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SetCursorColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cursor_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CursorColumns = types.ListValueMust(t, vs)
}

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
	// The pipeline to return events for.
	PipelineId types.String `tfsdk:"-"`
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
	Events types.List `tfsdk:"events"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (toState *ListPipelineEventsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListPipelineEventsResponse_SdkV2) {
}

func (toState *ListPipelineEventsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListPipelineEventsResponse_SdkV2) {
}

func (c ListPipelineEventsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["events"] = attrs["events"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["prev_page_token"] = attrs["prev_page_token"].SetOptional()

	return attrs
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
				ElemType: PipelineEvent_SdkV2{}.Type(ctx),
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of events matching the request criteria.
	Statuses types.List `tfsdk:"statuses"`
}

func (toState *ListPipelinesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListPipelinesResponse_SdkV2) {
}

func (toState *ListPipelinesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListPipelinesResponse_SdkV2) {
}

func (c ListPipelinesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["statuses"] = attrs["statuses"].SetOptional()

	return attrs
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
				ElemType: PipelineStateInfo_SdkV2{}.Type(ctx),
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token"`

	Updates types.List `tfsdk:"updates"`
}

func (toState *ListUpdatesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListUpdatesResponse_SdkV2) {
}

func (toState *ListUpdatesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListUpdatesResponse_SdkV2) {
}

func (c ListUpdatesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["prev_page_token"] = attrs["prev_page_token"].SetOptional()
	attrs["updates"] = attrs["updates"].SetOptional()

	return attrs
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
				ElemType: UpdateInfo_SdkV2{}.Type(ctx),
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

func (toState *ManualTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ManualTrigger_SdkV2) {
}

func (toState *ManualTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ManualTrigger_SdkV2) {
}

func (c ManualTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	// The absolute path of the source code.
	Path types.String `tfsdk:"path"`
}

func (toState *NotebookLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NotebookLibrary_SdkV2) {
}

func (toState *NotebookLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState NotebookLibrary_SdkV2) {
}

func (c NotebookLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
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
	Alerts types.List `tfsdk:"alerts"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients types.List `tfsdk:"email_recipients"`
}

func (toState *Notifications_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Notifications_SdkV2) {
}

func (toState *Notifications_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Notifications_SdkV2) {
}

func (c Notifications_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alerts"] = attrs["alerts"].SetOptional()
	attrs["email_recipients"] = attrs["email_recipients"].SetOptional()

	return attrs
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
	BatchId types.Int64 `tfsdk:"batch_id"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud types.String `tfsdk:"cloud"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName types.String `tfsdk:"dataset_name"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId types.String `tfsdk:"flow_id"`
	// The name of the flow. Not unique.
	FlowName types.String `tfsdk:"flow_name"`
	// The optional host name where the event was triggered
	Host types.String `tfsdk:"host"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId types.String `tfsdk:"maintenance_id"`
	// Materialization name.
	MaterializationName types.String `tfsdk:"materialization_name"`
	// The org id of the user. Unique within a cloud.
	OrgId types.Int64 `tfsdk:"org_id"`
	// The id of the pipeline. Globally unique.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// The name of the pipeline. Not unique.
	PipelineName types.String `tfsdk:"pipeline_name"`
	// The cloud region.
	Region types.String `tfsdk:"region"`
	// The id of the request that caused an update.
	RequestId types.String `tfsdk:"request_id"`
	// The id of a (delta) table. Globally unique.
	TableId types.String `tfsdk:"table_id"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId types.String `tfsdk:"uc_resource_id"`
	// The id of an execution. Globally unique.
	UpdateId types.String `tfsdk:"update_id"`
}

func (toState *Origin_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Origin_SdkV2) {
}

func (toState *Origin_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Origin_SdkV2) {
}

func (c Origin_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["batch_id"] = attrs["batch_id"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["dataset_name"] = attrs["dataset_name"].SetOptional()
	attrs["flow_id"] = attrs["flow_id"].SetOptional()
	attrs["flow_name"] = attrs["flow_name"].SetOptional()
	attrs["host"] = attrs["host"].SetOptional()
	attrs["maintenance_id"] = attrs["maintenance_id"].SetOptional()
	attrs["materialization_name"] = attrs["materialization_name"].SetOptional()
	attrs["org_id"] = attrs["org_id"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["pipeline_name"] = attrs["pipeline_name"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["request_id"] = attrs["request_id"].SetOptional()
	attrs["table_id"] = attrs["table_id"].SetOptional()
	attrs["uc_resource_id"] = attrs["uc_resource_id"].SetOptional()
	attrs["update_id"] = attrs["update_id"].SetOptional()

	return attrs
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

type PathPattern_SdkV2 struct {
	// The source code to include for pipelines
	Include types.String `tfsdk:"include"`
}

func (toState *PathPattern_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PathPattern_SdkV2) {
}

func (toState *PathPattern_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PathPattern_SdkV2) {
}

func (c PathPattern_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["include"] = attrs["include"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PathPattern.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PathPattern_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PathPattern_SdkV2
// only implements ToObjectValue() and Type().
func (o PathPattern_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include": o.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PathPattern_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include": types.StringType,
		},
	}
}

type PipelineAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *PipelineAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineAccessControlRequest_SdkV2) {
}

func (toState *PipelineAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineAccessControlRequest_SdkV2) {
}

func (c PipelineAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
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
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (toState *PipelineAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineAccessControlResponse_SdkV2) {
}

func (toState *PipelineAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineAccessControlResponse_SdkV2) {
}

func (c PipelineAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
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
				ElemType: PipelinePermission_SdkV2{}.Type(ctx),
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
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.List `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.List `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.List `tfsdk:"cluster_log_conf"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags types.Map `tfsdk:"custom_tags"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id"`
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes types.List `tfsdk:"gcp_attributes"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts types.List `tfsdk:"init_scripts"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label types.String `tfsdk:"label"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id"`
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
	NumWorkers types.Int64 `tfsdk:"num_workers"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf types.Map `tfsdk:"spark_conf"`
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
	SparkEnvVars types.Map `tfsdk:"spark_env_vars"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys types.List `tfsdk:"ssh_public_keys"`
}

func (toState *PipelineCluster_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineCluster_SdkV2) {
	if !fromPlan.Autoscale.IsNull() && !fromPlan.Autoscale.IsUnknown() {
		if toStateAutoscale, ok := toState.GetAutoscale(ctx); ok {
			if fromPlanAutoscale, ok := fromPlan.GetAutoscale(ctx); ok {
				toStateAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAutoscale)
				toState.SetAutoscale(ctx, toStateAutoscale)
			}
		}
	}
	if !fromPlan.AwsAttributes.IsNull() && !fromPlan.AwsAttributes.IsUnknown() {
		if toStateAwsAttributes, ok := toState.GetAwsAttributes(ctx); ok {
			if fromPlanAwsAttributes, ok := fromPlan.GetAwsAttributes(ctx); ok {
				toStateAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAwsAttributes)
				toState.SetAwsAttributes(ctx, toStateAwsAttributes)
			}
		}
	}
	if !fromPlan.AzureAttributes.IsNull() && !fromPlan.AzureAttributes.IsUnknown() {
		if toStateAzureAttributes, ok := toState.GetAzureAttributes(ctx); ok {
			if fromPlanAzureAttributes, ok := fromPlan.GetAzureAttributes(ctx); ok {
				toStateAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAzureAttributes)
				toState.SetAzureAttributes(ctx, toStateAzureAttributes)
			}
		}
	}
	if !fromPlan.ClusterLogConf.IsNull() && !fromPlan.ClusterLogConf.IsUnknown() {
		if toStateClusterLogConf, ok := toState.GetClusterLogConf(ctx); ok {
			if fromPlanClusterLogConf, ok := fromPlan.GetClusterLogConf(ctx); ok {
				toStateClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanClusterLogConf)
				toState.SetClusterLogConf(ctx, toStateClusterLogConf)
			}
		}
	}
	if !fromPlan.GcpAttributes.IsNull() && !fromPlan.GcpAttributes.IsUnknown() {
		if toStateGcpAttributes, ok := toState.GetGcpAttributes(ctx); ok {
			if fromPlanGcpAttributes, ok := fromPlan.GetGcpAttributes(ctx); ok {
				toStateGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcpAttributes)
				toState.SetGcpAttributes(ctx, toStateGcpAttributes)
			}
		}
	}
}

func (toState *PipelineCluster_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineCluster_SdkV2) {
	if !fromState.Autoscale.IsNull() && !fromState.Autoscale.IsUnknown() {
		if toStateAutoscale, ok := toState.GetAutoscale(ctx); ok {
			if fromStateAutoscale, ok := fromState.GetAutoscale(ctx); ok {
				toStateAutoscale.SyncFieldsDuringRead(ctx, fromStateAutoscale)
				toState.SetAutoscale(ctx, toStateAutoscale)
			}
		}
	}
	if !fromState.AwsAttributes.IsNull() && !fromState.AwsAttributes.IsUnknown() {
		if toStateAwsAttributes, ok := toState.GetAwsAttributes(ctx); ok {
			if fromStateAwsAttributes, ok := fromState.GetAwsAttributes(ctx); ok {
				toStateAwsAttributes.SyncFieldsDuringRead(ctx, fromStateAwsAttributes)
				toState.SetAwsAttributes(ctx, toStateAwsAttributes)
			}
		}
	}
	if !fromState.AzureAttributes.IsNull() && !fromState.AzureAttributes.IsUnknown() {
		if toStateAzureAttributes, ok := toState.GetAzureAttributes(ctx); ok {
			if fromStateAzureAttributes, ok := fromState.GetAzureAttributes(ctx); ok {
				toStateAzureAttributes.SyncFieldsDuringRead(ctx, fromStateAzureAttributes)
				toState.SetAzureAttributes(ctx, toStateAzureAttributes)
			}
		}
	}
	if !fromState.ClusterLogConf.IsNull() && !fromState.ClusterLogConf.IsUnknown() {
		if toStateClusterLogConf, ok := toState.GetClusterLogConf(ctx); ok {
			if fromStateClusterLogConf, ok := fromState.GetClusterLogConf(ctx); ok {
				toStateClusterLogConf.SyncFieldsDuringRead(ctx, fromStateClusterLogConf)
				toState.SetClusterLogConf(ctx, toStateClusterLogConf)
			}
		}
	}
	if !fromState.GcpAttributes.IsNull() && !fromState.GcpAttributes.IsUnknown() {
		if toStateGcpAttributes, ok := toState.GetGcpAttributes(ctx); ok {
			if fromStateGcpAttributes, ok := fromState.GetGcpAttributes(ctx); ok {
				toStateGcpAttributes.SyncFieldsDuringRead(ctx, fromStateGcpAttributes)
				toState.SetGcpAttributes(ctx, toStateGcpAttributes)
			}
		}
	}
}

func (c PipelineCluster_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["init_scripts"] = attrs["init_scripts"].SetOptional()
	attrs["instance_pool_id"] = attrs["instance_pool_id"].SetOptional()
	attrs["label"] = attrs["label"].SetOptional()
	attrs["node_type_id"] = attrs["node_type_id"].SetOptional()
	attrs["num_workers"] = attrs["num_workers"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["spark_conf"] = attrs["spark_conf"].SetOptional()
	attrs["spark_env_vars"] = attrs["spark_env_vars"].SetOptional()
	attrs["ssh_public_keys"] = attrs["ssh_public_keys"].SetOptional()

	return attrs
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
				ElemType: PipelineClusterAutoscale_SdkV2{}.Type(ctx),
			},
			"aws_attributes": basetypes.ListType{
				ElemType: compute_tf.AwsAttributes_SdkV2{}.Type(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: compute_tf.AzureAttributes_SdkV2{}.Type(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: compute_tf.ClusterLogConf_SdkV2{}.Type(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: compute_tf.GcpAttributes_SdkV2{}.Type(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: compute_tf.InitScriptInfo_SdkV2{}.Type(ctx),
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
	MaxWorkers types.Int64 `tfsdk:"max_workers"`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers"`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode types.String `tfsdk:"mode"`
}

func (toState *PipelineClusterAutoscale_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineClusterAutoscale_SdkV2) {
}

func (toState *PipelineClusterAutoscale_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineClusterAutoscale_SdkV2) {
}

func (c PipelineClusterAutoscale_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_workers"] = attrs["max_workers"].SetRequired()
	attrs["min_workers"] = attrs["min_workers"].SetRequired()
	attrs["mode"] = attrs["mode"].SetOptional()

	return attrs
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
	Kind types.String `tfsdk:"kind"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path"`
}

func (toState *PipelineDeployment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineDeployment_SdkV2) {
}

func (toState *PipelineDeployment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineDeployment_SdkV2) {
}

func (c PipelineDeployment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["kind"] = attrs["kind"].SetRequired()
	attrs["metadata_file_path"] = attrs["metadata_file_path"].SetOptional()

	return attrs
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
	Error types.List `tfsdk:"error"`
	// The event type. Should always correspond to the details
	EventType types.String `tfsdk:"event_type"`
	// A time-based, globally unique id.
	Id types.String `tfsdk:"id"`
	// The severity level of the event.
	Level types.String `tfsdk:"level"`
	// Maturity level for event_type.
	MaturityLevel types.String `tfsdk:"maturity_level"`
	// The display message associated with the event.
	Message types.String `tfsdk:"message"`
	// Describes where the event originates from.
	Origin types.List `tfsdk:"origin"`
	// A sequencing object to identify and order events.
	Sequence types.List `tfsdk:"sequence"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (toState *PipelineEvent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineEvent_SdkV2) {
	if !fromPlan.Error.IsNull() && !fromPlan.Error.IsUnknown() {
		if toStateError, ok := toState.GetError(ctx); ok {
			if fromPlanError, ok := fromPlan.GetError(ctx); ok {
				toStateError.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanError)
				toState.SetError(ctx, toStateError)
			}
		}
	}
	if !fromPlan.Origin.IsNull() && !fromPlan.Origin.IsUnknown() {
		if toStateOrigin, ok := toState.GetOrigin(ctx); ok {
			if fromPlanOrigin, ok := fromPlan.GetOrigin(ctx); ok {
				toStateOrigin.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOrigin)
				toState.SetOrigin(ctx, toStateOrigin)
			}
		}
	}
	if !fromPlan.Sequence.IsNull() && !fromPlan.Sequence.IsUnknown() {
		if toStateSequence, ok := toState.GetSequence(ctx); ok {
			if fromPlanSequence, ok := fromPlan.GetSequence(ctx); ok {
				toStateSequence.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSequence)
				toState.SetSequence(ctx, toStateSequence)
			}
		}
	}
}

func (toState *PipelineEvent_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineEvent_SdkV2) {
	if !fromState.Error.IsNull() && !fromState.Error.IsUnknown() {
		if toStateError, ok := toState.GetError(ctx); ok {
			if fromStateError, ok := fromState.GetError(ctx); ok {
				toStateError.SyncFieldsDuringRead(ctx, fromStateError)
				toState.SetError(ctx, toStateError)
			}
		}
	}
	if !fromState.Origin.IsNull() && !fromState.Origin.IsUnknown() {
		if toStateOrigin, ok := toState.GetOrigin(ctx); ok {
			if fromStateOrigin, ok := fromState.GetOrigin(ctx); ok {
				toStateOrigin.SyncFieldsDuringRead(ctx, fromStateOrigin)
				toState.SetOrigin(ctx, toStateOrigin)
			}
		}
	}
	if !fromState.Sequence.IsNull() && !fromState.Sequence.IsUnknown() {
		if toStateSequence, ok := toState.GetSequence(ctx); ok {
			if fromStateSequence, ok := fromState.GetSequence(ctx); ok {
				toStateSequence.SyncFieldsDuringRead(ctx, fromStateSequence)
				toState.SetSequence(ctx, toStateSequence)
			}
		}
	}
}

func (c PipelineEvent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["event_type"] = attrs["event_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["level"] = attrs["level"].SetOptional()
	attrs["maturity_level"] = attrs["maturity_level"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["origin"] = attrs["origin"].SetOptional()
	attrs["origin"] = attrs["origin"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sequence"] = attrs["sequence"].SetOptional()
	attrs["sequence"] = attrs["sequence"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
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
				ElemType: ErrorDetail_SdkV2{}.Type(ctx),
			},
			"event_type":     types.StringType,
			"id":             types.StringType,
			"level":          types.StringType,
			"maturity_level": types.StringType,
			"message":        types.StringType,
			"origin": basetypes.ListType{
				ElemType: Origin_SdkV2{}.Type(ctx),
			},
			"sequence": basetypes.ListType{
				ElemType: Sequencing_SdkV2{}.Type(ctx),
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
	File types.List `tfsdk:"file"`
	// The unified field to include source codes. Each entry can be a notebook
	// path, a file path, or a folder path that ends `/**`. This field cannot be
	// used together with `notebook` or `file`.
	Glob types.List `tfsdk:"glob"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar"`
	// Specification of a maven library to be installed.
	Maven types.List `tfsdk:"maven"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook types.List `tfsdk:"notebook"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl"`
}

func (toState *PipelineLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineLibrary_SdkV2) {
	if !fromPlan.File.IsNull() && !fromPlan.File.IsUnknown() {
		if toStateFile, ok := toState.GetFile(ctx); ok {
			if fromPlanFile, ok := fromPlan.GetFile(ctx); ok {
				toStateFile.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFile)
				toState.SetFile(ctx, toStateFile)
			}
		}
	}
	if !fromPlan.Glob.IsNull() && !fromPlan.Glob.IsUnknown() {
		if toStateGlob, ok := toState.GetGlob(ctx); ok {
			if fromPlanGlob, ok := fromPlan.GetGlob(ctx); ok {
				toStateGlob.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGlob)
				toState.SetGlob(ctx, toStateGlob)
			}
		}
	}
	if !fromPlan.Maven.IsNull() && !fromPlan.Maven.IsUnknown() {
		if toStateMaven, ok := toState.GetMaven(ctx); ok {
			if fromPlanMaven, ok := fromPlan.GetMaven(ctx); ok {
				toStateMaven.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanMaven)
				toState.SetMaven(ctx, toStateMaven)
			}
		}
	}
	if !fromPlan.Notebook.IsNull() && !fromPlan.Notebook.IsUnknown() {
		if toStateNotebook, ok := toState.GetNotebook(ctx); ok {
			if fromPlanNotebook, ok := fromPlan.GetNotebook(ctx); ok {
				toStateNotebook.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanNotebook)
				toState.SetNotebook(ctx, toStateNotebook)
			}
		}
	}
}

func (toState *PipelineLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineLibrary_SdkV2) {
	if !fromState.File.IsNull() && !fromState.File.IsUnknown() {
		if toStateFile, ok := toState.GetFile(ctx); ok {
			if fromStateFile, ok := fromState.GetFile(ctx); ok {
				toStateFile.SyncFieldsDuringRead(ctx, fromStateFile)
				toState.SetFile(ctx, toStateFile)
			}
		}
	}
	if !fromState.Glob.IsNull() && !fromState.Glob.IsUnknown() {
		if toStateGlob, ok := toState.GetGlob(ctx); ok {
			if fromStateGlob, ok := fromState.GetGlob(ctx); ok {
				toStateGlob.SyncFieldsDuringRead(ctx, fromStateGlob)
				toState.SetGlob(ctx, toStateGlob)
			}
		}
	}
	if !fromState.Maven.IsNull() && !fromState.Maven.IsUnknown() {
		if toStateMaven, ok := toState.GetMaven(ctx); ok {
			if fromStateMaven, ok := fromState.GetMaven(ctx); ok {
				toStateMaven.SyncFieldsDuringRead(ctx, fromStateMaven)
				toState.SetMaven(ctx, toStateMaven)
			}
		}
	}
	if !fromState.Notebook.IsNull() && !fromState.Notebook.IsUnknown() {
		if toStateNotebook, ok := toState.GetNotebook(ctx); ok {
			if fromStateNotebook, ok := fromState.GetNotebook(ctx); ok {
				toStateNotebook.SyncFieldsDuringRead(ctx, fromStateNotebook)
				toState.SetNotebook(ctx, toStateNotebook)
			}
		}
	}
}

func (c PipelineLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file"] = attrs["file"].SetOptional()
	attrs["file"] = attrs["file"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["glob"] = attrs["glob"].SetOptional()
	attrs["glob"] = attrs["glob"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["jar"] = attrs["jar"].SetOptional()
	attrs["maven"] = attrs["maven"].SetOptional()
	attrs["maven"] = attrs["maven"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["notebook"] = attrs["notebook"].SetOptional()
	attrs["notebook"] = attrs["notebook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["whl"] = attrs["whl"].SetOptional()

	return attrs
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
		"glob":     reflect.TypeOf(PathPattern_SdkV2{}),
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
			"glob":     o.Glob,
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
				ElemType: FileLibrary_SdkV2{}.Type(ctx),
			},
			"glob": basetypes.ListType{
				ElemType: PathPattern_SdkV2{}.Type(ctx),
			},
			"jar": types.StringType,
			"maven": basetypes.ListType{
				ElemType: compute_tf.MavenLibrary_SdkV2{}.Type(ctx),
			},
			"notebook": basetypes.ListType{
				ElemType: NotebookLibrary_SdkV2{}.Type(ctx),
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

// GetGlob returns the value of the Glob field in PipelineLibrary_SdkV2 as
// a PathPattern_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary_SdkV2) GetGlob(ctx context.Context) (PathPattern_SdkV2, bool) {
	var e PathPattern_SdkV2
	if o.Glob.IsNull() || o.Glob.IsUnknown() {
		return e, false
	}
	var v []PathPattern_SdkV2
	d := o.Glob.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlob sets the value of the Glob field in PipelineLibrary_SdkV2.
func (o *PipelineLibrary_SdkV2) SetGlob(ctx context.Context, v PathPattern_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["glob"]
	o.Glob = types.ListValueMust(t, vs)
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
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PipelinePermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelinePermission_SdkV2) {
}

func (toState *PipelinePermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelinePermission_SdkV2) {
}

func (c PipelinePermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
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
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *PipelinePermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelinePermissions_SdkV2) {
}

func (toState *PipelinePermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelinePermissions_SdkV2) {
}

func (c PipelinePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
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
				ElemType: PipelineAccessControlResponse_SdkV2{}.Type(ctx),
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
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *PipelinePermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelinePermissionsDescription_SdkV2) {
}

func (toState *PipelinePermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelinePermissionsDescription_SdkV2) {
}

func (c PipelinePermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
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
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
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
				ElemType: PipelineAccessControlRequest_SdkV2{}.Type(ctx),
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
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog types.String `tfsdk:"catalog"`
	// DLT Release Channel that specifies which version to use.
	Channel types.String `tfsdk:"channel"`
	// Cluster settings for this pipeline deployment.
	Clusters types.List `tfsdk:"clusters"`
	// String-String configuration for this pipeline execution.
	Configuration types.Map `tfsdk:"configuration"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous types.Bool `tfsdk:"continuous"`
	// Deployment type of this pipeline.
	Deployment types.List `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.List `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.List `tfsdk:"event_log"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.List `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.List `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.List `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Restart window of this pipeline.
	RestartWindow types.List `tfsdk:"restart_window"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath types.String `tfsdk:"root_path"`
	// The default schema (database) where tables are read from or published to.
	Schema types.String `tfsdk:"schema"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless types.Bool `tfsdk:"serverless"`
	// DBFS root directory for storing checkpoints and tables.
	Storage types.String `tfsdk:"storage"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	Tags types.Map `tfsdk:"tags"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target types.String `tfsdk:"target"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger types.List `tfsdk:"trigger"`
}

func (toState *PipelineSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineSpec_SdkV2) {
	if !fromPlan.Deployment.IsNull() && !fromPlan.Deployment.IsUnknown() {
		if toStateDeployment, ok := toState.GetDeployment(ctx); ok {
			if fromPlanDeployment, ok := fromPlan.GetDeployment(ctx); ok {
				toStateDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDeployment)
				toState.SetDeployment(ctx, toStateDeployment)
			}
		}
	}
	if !fromPlan.Environment.IsNull() && !fromPlan.Environment.IsUnknown() {
		if toStateEnvironment, ok := toState.GetEnvironment(ctx); ok {
			if fromPlanEnvironment, ok := fromPlan.GetEnvironment(ctx); ok {
				toStateEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEnvironment)
				toState.SetEnvironment(ctx, toStateEnvironment)
			}
		}
	}
	if !fromPlan.EventLog.IsNull() && !fromPlan.EventLog.IsUnknown() {
		if toStateEventLog, ok := toState.GetEventLog(ctx); ok {
			if fromPlanEventLog, ok := fromPlan.GetEventLog(ctx); ok {
				toStateEventLog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEventLog)
				toState.SetEventLog(ctx, toStateEventLog)
			}
		}
	}
	if !fromPlan.Filters.IsNull() && !fromPlan.Filters.IsUnknown() {
		if toStateFilters, ok := toState.GetFilters(ctx); ok {
			if fromPlanFilters, ok := fromPlan.GetFilters(ctx); ok {
				toStateFilters.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFilters)
				toState.SetFilters(ctx, toStateFilters)
			}
		}
	}
	if !fromPlan.GatewayDefinition.IsNull() && !fromPlan.GatewayDefinition.IsUnknown() {
		if toStateGatewayDefinition, ok := toState.GetGatewayDefinition(ctx); ok {
			if fromPlanGatewayDefinition, ok := fromPlan.GetGatewayDefinition(ctx); ok {
				toStateGatewayDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGatewayDefinition)
				toState.SetGatewayDefinition(ctx, toStateGatewayDefinition)
			}
		}
	}
	if !fromPlan.IngestionDefinition.IsNull() && !fromPlan.IngestionDefinition.IsUnknown() {
		if toStateIngestionDefinition, ok := toState.GetIngestionDefinition(ctx); ok {
			if fromPlanIngestionDefinition, ok := fromPlan.GetIngestionDefinition(ctx); ok {
				toStateIngestionDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanIngestionDefinition)
				toState.SetIngestionDefinition(ctx, toStateIngestionDefinition)
			}
		}
	}
	if !fromPlan.RestartWindow.IsNull() && !fromPlan.RestartWindow.IsUnknown() {
		if toStateRestartWindow, ok := toState.GetRestartWindow(ctx); ok {
			if fromPlanRestartWindow, ok := fromPlan.GetRestartWindow(ctx); ok {
				toStateRestartWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRestartWindow)
				toState.SetRestartWindow(ctx, toStateRestartWindow)
			}
		}
	}
	if !fromPlan.Trigger.IsNull() && !fromPlan.Trigger.IsUnknown() {
		if toStateTrigger, ok := toState.GetTrigger(ctx); ok {
			if fromPlanTrigger, ok := fromPlan.GetTrigger(ctx); ok {
				toStateTrigger.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTrigger)
				toState.SetTrigger(ctx, toStateTrigger)
			}
		}
	}
}

func (toState *PipelineSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineSpec_SdkV2) {
	if !fromState.Deployment.IsNull() && !fromState.Deployment.IsUnknown() {
		if toStateDeployment, ok := toState.GetDeployment(ctx); ok {
			if fromStateDeployment, ok := fromState.GetDeployment(ctx); ok {
				toStateDeployment.SyncFieldsDuringRead(ctx, fromStateDeployment)
				toState.SetDeployment(ctx, toStateDeployment)
			}
		}
	}
	if !fromState.Environment.IsNull() && !fromState.Environment.IsUnknown() {
		if toStateEnvironment, ok := toState.GetEnvironment(ctx); ok {
			if fromStateEnvironment, ok := fromState.GetEnvironment(ctx); ok {
				toStateEnvironment.SyncFieldsDuringRead(ctx, fromStateEnvironment)
				toState.SetEnvironment(ctx, toStateEnvironment)
			}
		}
	}
	if !fromState.EventLog.IsNull() && !fromState.EventLog.IsUnknown() {
		if toStateEventLog, ok := toState.GetEventLog(ctx); ok {
			if fromStateEventLog, ok := fromState.GetEventLog(ctx); ok {
				toStateEventLog.SyncFieldsDuringRead(ctx, fromStateEventLog)
				toState.SetEventLog(ctx, toStateEventLog)
			}
		}
	}
	if !fromState.Filters.IsNull() && !fromState.Filters.IsUnknown() {
		if toStateFilters, ok := toState.GetFilters(ctx); ok {
			if fromStateFilters, ok := fromState.GetFilters(ctx); ok {
				toStateFilters.SyncFieldsDuringRead(ctx, fromStateFilters)
				toState.SetFilters(ctx, toStateFilters)
			}
		}
	}
	if !fromState.GatewayDefinition.IsNull() && !fromState.GatewayDefinition.IsUnknown() {
		if toStateGatewayDefinition, ok := toState.GetGatewayDefinition(ctx); ok {
			if fromStateGatewayDefinition, ok := fromState.GetGatewayDefinition(ctx); ok {
				toStateGatewayDefinition.SyncFieldsDuringRead(ctx, fromStateGatewayDefinition)
				toState.SetGatewayDefinition(ctx, toStateGatewayDefinition)
			}
		}
	}
	if !fromState.IngestionDefinition.IsNull() && !fromState.IngestionDefinition.IsUnknown() {
		if toStateIngestionDefinition, ok := toState.GetIngestionDefinition(ctx); ok {
			if fromStateIngestionDefinition, ok := fromState.GetIngestionDefinition(ctx); ok {
				toStateIngestionDefinition.SyncFieldsDuringRead(ctx, fromStateIngestionDefinition)
				toState.SetIngestionDefinition(ctx, toStateIngestionDefinition)
			}
		}
	}
	if !fromState.RestartWindow.IsNull() && !fromState.RestartWindow.IsUnknown() {
		if toStateRestartWindow, ok := toState.GetRestartWindow(ctx); ok {
			if fromStateRestartWindow, ok := fromState.GetRestartWindow(ctx); ok {
				toStateRestartWindow.SyncFieldsDuringRead(ctx, fromStateRestartWindow)
				toState.SetRestartWindow(ctx, toStateRestartWindow)
			}
		}
	}
	if !fromState.Trigger.IsNull() && !fromState.Trigger.IsUnknown() {
		if toStateTrigger, ok := toState.GetTrigger(ctx); ok {
			if fromStateTrigger, ok := fromState.GetTrigger(ctx); ok {
				toStateTrigger.SyncFieldsDuringRead(ctx, fromStateTrigger)
				toState.SetTrigger(ctx, toStateTrigger)
			}
		}
	}
}

func (c PipelineSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["deployment"] = attrs["deployment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["development"] = attrs["development"].SetOptional()
	attrs["edition"] = attrs["edition"].SetOptional()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["environment"] = attrs["environment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["event_log"] = attrs["event_log"].SetOptional()
	attrs["event_log"] = attrs["event_log"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["filters"] = attrs["filters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["gateway_definition"] = attrs["gateway_definition"].SetOptional()
	attrs["gateway_definition"] = attrs["gateway_definition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["ingestion_definition"] = attrs["ingestion_definition"].SetOptional()
	attrs["ingestion_definition"] = attrs["ingestion_definition"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["photon"] = attrs["photon"].SetOptional()
	attrs["restart_window"] = attrs["restart_window"].SetOptional()
	attrs["restart_window"] = attrs["restart_window"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["root_path"] = attrs["root_path"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()
	attrs["trigger"] = attrs["trigger"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
		"environment":          reflect.TypeOf(PipelinesEnvironment_SdkV2{}),
		"event_log":            reflect.TypeOf(EventLogSpec_SdkV2{}),
		"filters":              reflect.TypeOf(Filters_SdkV2{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition_SdkV2{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition_SdkV2{}),
		"libraries":            reflect.TypeOf(PipelineLibrary_SdkV2{}),
		"notifications":        reflect.TypeOf(Notifications_SdkV2{}),
		"restart_window":       reflect.TypeOf(RestartWindow_SdkV2{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":          o.Environment,
			"event_log":            o.EventLog,
			"filters":              o.Filters,
			"gateway_definition":   o.GatewayDefinition,
			"id":                   o.Id,
			"ingestion_definition": o.IngestionDefinition,
			"libraries":            o.Libraries,
			"name":                 o.Name,
			"notifications":        o.Notifications,
			"photon":               o.Photon,
			"restart_window":       o.RestartWindow,
			"root_path":            o.RootPath,
			"schema":               o.Schema,
			"serverless":           o.Serverless,
			"storage":              o.Storage,
			"tags":                 o.Tags,
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
				ElemType: PipelineCluster_SdkV2{}.Type(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment_SdkV2{}.Type(ctx),
			},
			"development": types.BoolType,
			"edition":     types.StringType,
			"environment": basetypes.ListType{
				ElemType: PipelinesEnvironment_SdkV2{}.Type(ctx),
			},
			"event_log": basetypes.ListType{
				ElemType: EventLogSpec_SdkV2{}.Type(ctx),
			},
			"filters": basetypes.ListType{
				ElemType: Filters_SdkV2{}.Type(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition_SdkV2{}.Type(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications_SdkV2{}.Type(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow_SdkV2{}.Type(ctx),
			},
			"root_path":  types.StringType,
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target": types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger_SdkV2{}.Type(ctx),
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

// GetEnvironment returns the value of the Environment field in PipelineSpec_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := o.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	o.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in PipelineSpec_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := o.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	o.EventLog = types.ListValueMust(t, vs)
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

// GetTags returns the value of the Tags field in PipelineSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in PipelineSpec_SdkV2.
func (o *PipelineSpec_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
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
	ClusterId types.String `tfsdk:"cluster_id"`
	// The username of the pipeline creator.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// The health of a pipeline.
	Health types.String `tfsdk:"health"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates types.List `tfsdk:"latest_updates"`
	// The user-friendly name of the pipeline.
	Name types.String `tfsdk:"name"`
	// The unique identifier of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`

	State types.String `tfsdk:"state"`
}

func (toState *PipelineStateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineStateInfo_SdkV2) {
}

func (toState *PipelineStateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineStateInfo_SdkV2) {
}

func (c PipelineStateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["health"] = attrs["health"].SetOptional()
	attrs["latest_updates"] = attrs["latest_updates"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
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
				ElemType: UpdateStateInfo_SdkV2{}.Type(ctx),
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
	Cron types.List `tfsdk:"cron"`

	Manual types.List `tfsdk:"manual"`
}

func (toState *PipelineTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelineTrigger_SdkV2) {
	if !fromPlan.Cron.IsNull() && !fromPlan.Cron.IsUnknown() {
		if toStateCron, ok := toState.GetCron(ctx); ok {
			if fromPlanCron, ok := fromPlan.GetCron(ctx); ok {
				toStateCron.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCron)
				toState.SetCron(ctx, toStateCron)
			}
		}
	}
	if !fromPlan.Manual.IsNull() && !fromPlan.Manual.IsUnknown() {
		if toStateManual, ok := toState.GetManual(ctx); ok {
			if fromPlanManual, ok := fromPlan.GetManual(ctx); ok {
				toStateManual.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanManual)
				toState.SetManual(ctx, toStateManual)
			}
		}
	}
}

func (toState *PipelineTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelineTrigger_SdkV2) {
	if !fromState.Cron.IsNull() && !fromState.Cron.IsUnknown() {
		if toStateCron, ok := toState.GetCron(ctx); ok {
			if fromStateCron, ok := fromState.GetCron(ctx); ok {
				toStateCron.SyncFieldsDuringRead(ctx, fromStateCron)
				toState.SetCron(ctx, toStateCron)
			}
		}
	}
	if !fromState.Manual.IsNull() && !fromState.Manual.IsUnknown() {
		if toStateManual, ok := toState.GetManual(ctx); ok {
			if fromStateManual, ok := fromState.GetManual(ctx); ok {
				toStateManual.SyncFieldsDuringRead(ctx, fromStateManual)
				toState.SetManual(ctx, toStateManual)
			}
		}
	}
}

func (c PipelineTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cron"] = attrs["cron"].SetOptional()
	attrs["cron"] = attrs["cron"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["manual"] = attrs["manual"].SetOptional()
	attrs["manual"] = attrs["manual"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: CronTrigger_SdkV2{}.Type(ctx),
			},
			"manual": basetypes.ListType{
				ElemType: ManualTrigger_SdkV2{}.Type(ctx),
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

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type PipelinesEnvironment_SdkV2 struct {
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url>
	Dependencies types.List `tfsdk:"dependencies"`
}

func (toState *PipelinesEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PipelinesEnvironment_SdkV2) {
}

func (toState *PipelinesEnvironment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PipelinesEnvironment_SdkV2) {
}

func (c PipelinesEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dependencies"] = attrs["dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelinesEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelinesEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinesEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelinesEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": o.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinesEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in PipelinesEnvironment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinesEnvironment_SdkV2) GetDependencies(ctx context.Context) ([]types.String, bool) {
	if o.Dependencies.IsNull() || o.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in PipelinesEnvironment_SdkV2.
func (o *PipelinesEnvironment_SdkV2) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// PG-specific catalog-level configuration parameters
type PostgresCatalogConfig_SdkV2 struct {
	// Optional. The Postgres slot configuration to use for logical replication
	SlotConfig types.List `tfsdk:"slot_config"`
}

func (toState *PostgresCatalogConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PostgresCatalogConfig_SdkV2) {
	if !fromPlan.SlotConfig.IsNull() && !fromPlan.SlotConfig.IsUnknown() {
		if toStateSlotConfig, ok := toState.GetSlotConfig(ctx); ok {
			if fromPlanSlotConfig, ok := fromPlan.GetSlotConfig(ctx); ok {
				toStateSlotConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSlotConfig)
				toState.SetSlotConfig(ctx, toStateSlotConfig)
			}
		}
	}
}

func (toState *PostgresCatalogConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PostgresCatalogConfig_SdkV2) {
	if !fromState.SlotConfig.IsNull() && !fromState.SlotConfig.IsUnknown() {
		if toStateSlotConfig, ok := toState.GetSlotConfig(ctx); ok {
			if fromStateSlotConfig, ok := fromState.GetSlotConfig(ctx); ok {
				toStateSlotConfig.SyncFieldsDuringRead(ctx, fromStateSlotConfig)
				toState.SetSlotConfig(ctx, toStateSlotConfig)
			}
		}
	}
}

func (c PostgresCatalogConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["slot_config"] = attrs["slot_config"].SetOptional()
	attrs["slot_config"] = attrs["slot_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PostgresCatalogConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PostgresCatalogConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"slot_config": reflect.TypeOf(PostgresSlotConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresCatalogConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o PostgresCatalogConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"slot_config": o.SlotConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PostgresCatalogConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"slot_config": basetypes.ListType{
				ElemType: PostgresSlotConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSlotConfig returns the value of the SlotConfig field in PostgresCatalogConfig_SdkV2 as
// a PostgresSlotConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PostgresCatalogConfig_SdkV2) GetSlotConfig(ctx context.Context) (PostgresSlotConfig_SdkV2, bool) {
	var e PostgresSlotConfig_SdkV2
	if o.SlotConfig.IsNull() || o.SlotConfig.IsUnknown() {
		return e, false
	}
	var v []PostgresSlotConfig_SdkV2
	d := o.SlotConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSlotConfig sets the value of the SlotConfig field in PostgresCatalogConfig_SdkV2.
func (o *PostgresCatalogConfig_SdkV2) SetSlotConfig(ctx context.Context, v PostgresSlotConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slot_config"]
	o.SlotConfig = types.ListValueMust(t, vs)
}

// PostgresSlotConfig contains the configuration for a Postgres logical
// replication slot
type PostgresSlotConfig_SdkV2 struct {
	// The name of the publication to use for the Postgres source
	PublicationName types.String `tfsdk:"publication_name"`
	// The name of the logical replication slot to use for the Postgres source
	SlotName types.String `tfsdk:"slot_name"`
}

func (toState *PostgresSlotConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PostgresSlotConfig_SdkV2) {
}

func (toState *PostgresSlotConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PostgresSlotConfig_SdkV2) {
}

func (c PostgresSlotConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["publication_name"] = attrs["publication_name"].SetOptional()
	attrs["slot_name"] = attrs["slot_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PostgresSlotConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PostgresSlotConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresSlotConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o PostgresSlotConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publication_name": o.PublicationName,
			"slot_name":        o.SlotName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PostgresSlotConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"publication_name": types.StringType,
			"slot_name":        types.StringType,
		},
	}
}

type ReportSpec_SdkV2 struct {
	// Required. Destination catalog to store table.
	DestinationCatalog types.String `tfsdk:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable types.String `tfsdk:"destination_table"`
	// Required. Report URL in the source system.
	SourceUrl types.String `tfsdk:"source_url"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	TableConfiguration types.List `tfsdk:"table_configuration"`
}

func (toState *ReportSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ReportSpec_SdkV2) {
	if !fromPlan.TableConfiguration.IsNull() && !fromPlan.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromPlanTableConfiguration, ok := fromPlan.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (toState *ReportSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ReportSpec_SdkV2) {
	if !fromState.TableConfiguration.IsNull() && !fromState.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromStateTableConfiguration, ok := fromState.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringRead(ctx, fromStateTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (c ReportSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["destination_table"] = attrs["destination_table"].SetOptional()
	attrs["source_url"] = attrs["source_url"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()
	attrs["table_configuration"] = attrs["table_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: TableSpecificConfig_SdkV2{}.Type(ctx),
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
	DaysOfWeek types.List `tfsdk:"days_of_week"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour types.Int64 `tfsdk:"start_hour"`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId types.String `tfsdk:"time_zone_id"`
}

func (toState *RestartWindow_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestartWindow_SdkV2) {
}

func (toState *RestartWindow_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestartWindow_SdkV2) {
}

func (c RestartWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["days_of_week"] = attrs["days_of_week"].SetOptional()
	attrs["start_hour"] = attrs["start_hour"].SetRequired()
	attrs["time_zone_id"] = attrs["time_zone_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestartWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"days_of_week": reflect.TypeOf(types.String{}),
	}
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
			"days_of_week": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_hour":   types.Int64Type,
			"time_zone_id": types.StringType,
		},
	}
}

// GetDaysOfWeek returns the value of the DaysOfWeek field in RestartWindow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RestartWindow_SdkV2) GetDaysOfWeek(ctx context.Context) ([]types.String, bool) {
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

// SetDaysOfWeek sets the value of the DaysOfWeek field in RestartWindow_SdkV2.
func (o *RestartWindow_SdkV2) SetDaysOfWeek(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["days_of_week"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DaysOfWeek = types.ListValueMust(t, vs)
}

// Write-only setting, available only in Create/Update calls. Specifies the user
// or service principal that the pipeline runs as. If not specified, the
// pipeline runs as the user who created the pipeline.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type RunAs_SdkV2 struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName types.String `tfsdk:"user_name"`
}

func (toState *RunAs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RunAs_SdkV2) {
}

func (toState *RunAs_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RunAs_SdkV2) {
}

func (c RunAs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunAs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunAs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunAs_SdkV2
// only implements ToObjectValue() and Type().
func (o RunAs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunAs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type SchemaSpec_SdkV2 struct {
	// Required. Destination catalog to store tables.
	DestinationCatalog types.String `tfsdk:"destination_catalog"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema types.String `tfsdk:"destination_schema"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog types.String `tfsdk:"source_catalog"`
	// Required. Schema name in the source database.
	SourceSchema types.String `tfsdk:"source_schema"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration types.List `tfsdk:"table_configuration"`
}

func (toState *SchemaSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SchemaSpec_SdkV2) {
	if !fromPlan.TableConfiguration.IsNull() && !fromPlan.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromPlanTableConfiguration, ok := fromPlan.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (toState *SchemaSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SchemaSpec_SdkV2) {
	if !fromState.TableConfiguration.IsNull() && !fromState.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromStateTableConfiguration, ok := fromState.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringRead(ctx, fromStateTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (c SchemaSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()
	attrs["source_schema"] = attrs["source_schema"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()
	attrs["table_configuration"] = attrs["table_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: TableSpecificConfig_SdkV2{}.Type(ctx),
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
	// A sequence number, unique and increasing per pipeline.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no"`
	// the ID assigned by the data plane.
	DataPlaneId types.List `tfsdk:"data_plane_id"`
}

func (toState *Sequencing_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Sequencing_SdkV2) {
	if !fromPlan.DataPlaneId.IsNull() && !fromPlan.DataPlaneId.IsUnknown() {
		if toStateDataPlaneId, ok := toState.GetDataPlaneId(ctx); ok {
			if fromPlanDataPlaneId, ok := fromPlan.GetDataPlaneId(ctx); ok {
				toStateDataPlaneId.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDataPlaneId)
				toState.SetDataPlaneId(ctx, toStateDataPlaneId)
			}
		}
	}
}

func (toState *Sequencing_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Sequencing_SdkV2) {
	if !fromState.DataPlaneId.IsNull() && !fromState.DataPlaneId.IsUnknown() {
		if toStateDataPlaneId, ok := toState.GetDataPlaneId(ctx); ok {
			if fromStateDataPlaneId, ok := fromState.GetDataPlaneId(ctx); ok {
				toStateDataPlaneId.SyncFieldsDuringRead(ctx, fromStateDataPlaneId)
				toState.SetDataPlaneId(ctx, toStateDataPlaneId)
			}
		}
	}
}

func (c Sequencing_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["control_plane_seq_no"] = attrs["control_plane_seq_no"].SetOptional()
	attrs["data_plane_id"] = attrs["data_plane_id"].SetOptional()
	attrs["data_plane_id"] = attrs["data_plane_id"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: DataPlaneId_SdkV2{}.Type(ctx),
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
	ClassName types.String `tfsdk:"class_name"`
	// Exception message
	Message types.String `tfsdk:"message"`
	// Stack trace consisting of a list of stack frames
	Stack types.List `tfsdk:"stack"`
}

func (toState *SerializedException_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SerializedException_SdkV2) {
}

func (toState *SerializedException_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SerializedException_SdkV2) {
}

func (c SerializedException_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["class_name"] = attrs["class_name"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack"] = attrs["stack"].SetOptional()

	return attrs
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
				ElemType: StackFrame_SdkV2{}.Type(ctx),
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

// SourceCatalogConfig contains catalog-level custom configuration parameters
// for each source
type SourceCatalogConfig_SdkV2 struct {
	// Postgres-specific catalog-level configuration parameters
	Postgres types.List `tfsdk:"postgres"`
	// Source catalog name
	SourceCatalog types.String `tfsdk:"source_catalog"`
}

func (toState *SourceCatalogConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SourceCatalogConfig_SdkV2) {
	if !fromPlan.Postgres.IsNull() && !fromPlan.Postgres.IsUnknown() {
		if toStatePostgres, ok := toState.GetPostgres(ctx); ok {
			if fromPlanPostgres, ok := fromPlan.GetPostgres(ctx); ok {
				toStatePostgres.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPostgres)
				toState.SetPostgres(ctx, toStatePostgres)
			}
		}
	}
}

func (toState *SourceCatalogConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SourceCatalogConfig_SdkV2) {
	if !fromState.Postgres.IsNull() && !fromState.Postgres.IsUnknown() {
		if toStatePostgres, ok := toState.GetPostgres(ctx); ok {
			if fromStatePostgres, ok := fromState.GetPostgres(ctx); ok {
				toStatePostgres.SyncFieldsDuringRead(ctx, fromStatePostgres)
				toState.SetPostgres(ctx, toStatePostgres)
			}
		}
	}
}

func (c SourceCatalogConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["postgres"] = attrs["postgres"].SetOptional()
	attrs["postgres"] = attrs["postgres"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SourceCatalogConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SourceCatalogConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"postgres": reflect.TypeOf(PostgresCatalogConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceCatalogConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o SourceCatalogConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"postgres":       o.Postgres,
			"source_catalog": o.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SourceCatalogConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"postgres": basetypes.ListType{
				ElemType: PostgresCatalogConfig_SdkV2{}.Type(ctx),
			},
			"source_catalog": types.StringType,
		},
	}
}

// GetPostgres returns the value of the Postgres field in SourceCatalogConfig_SdkV2 as
// a PostgresCatalogConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SourceCatalogConfig_SdkV2) GetPostgres(ctx context.Context) (PostgresCatalogConfig_SdkV2, bool) {
	var e PostgresCatalogConfig_SdkV2
	if o.Postgres.IsNull() || o.Postgres.IsUnknown() {
		return e, false
	}
	var v []PostgresCatalogConfig_SdkV2
	d := o.Postgres.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPostgres sets the value of the Postgres field in SourceCatalogConfig_SdkV2.
func (o *SourceCatalogConfig_SdkV2) SetPostgres(ctx context.Context, v PostgresCatalogConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["postgres"]
	o.Postgres = types.ListValueMust(t, vs)
}

type SourceConfig_SdkV2 struct {
	// Catalog-level source configuration parameters
	Catalog types.List `tfsdk:"catalog"`
}

func (toState *SourceConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SourceConfig_SdkV2) {
	if !fromPlan.Catalog.IsNull() && !fromPlan.Catalog.IsUnknown() {
		if toStateCatalog, ok := toState.GetCatalog(ctx); ok {
			if fromPlanCatalog, ok := fromPlan.GetCatalog(ctx); ok {
				toStateCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCatalog)
				toState.SetCatalog(ctx, toStateCatalog)
			}
		}
	}
}

func (toState *SourceConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SourceConfig_SdkV2) {
	if !fromState.Catalog.IsNull() && !fromState.Catalog.IsUnknown() {
		if toStateCatalog, ok := toState.GetCatalog(ctx); ok {
			if fromStateCatalog, ok := fromState.GetCatalog(ctx); ok {
				toStateCatalog.SyncFieldsDuringRead(ctx, fromStateCatalog)
				toState.SetCatalog(ctx, toStateCatalog)
			}
		}
	}
}

func (c SourceConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["catalog"] = attrs["catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SourceConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SourceConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(SourceCatalogConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o SourceConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SourceConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": basetypes.ListType{
				ElemType: SourceCatalogConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCatalog returns the value of the Catalog field in SourceConfig_SdkV2 as
// a SourceCatalogConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SourceConfig_SdkV2) GetCatalog(ctx context.Context) (SourceCatalogConfig_SdkV2, bool) {
	var e SourceCatalogConfig_SdkV2
	if o.Catalog.IsNull() || o.Catalog.IsUnknown() {
		return e, false
	}
	var v []SourceCatalogConfig_SdkV2
	d := o.Catalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCatalog sets the value of the Catalog field in SourceConfig_SdkV2.
func (o *SourceConfig_SdkV2) SetCatalog(ctx context.Context, v SourceCatalogConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["catalog"]
	o.Catalog = types.ListValueMust(t, vs)
}

type StackFrame_SdkV2 struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name"`
}

func (toState *StackFrame_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StackFrame_SdkV2) {
}

func (toState *StackFrame_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState StackFrame_SdkV2) {
}

func (c StackFrame_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["declaring_class"] = attrs["declaring_class"].SetOptional()
	attrs["file_name"] = attrs["file_name"].SetOptional()
	attrs["line_number"] = attrs["line_number"].SetOptional()
	attrs["method_name"] = attrs["method_name"].SetOptional()

	return attrs
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
	Cause types.String `tfsdk:"cause"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection types.List `tfsdk:"full_refresh_selection"`

	PipelineId types.String `tfsdk:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection types.List `tfsdk:"refresh_selection"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only"`
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
	UpdateId types.String `tfsdk:"update_id"`
}

func (toState *StartUpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StartUpdateResponse_SdkV2) {
}

func (toState *StartUpdateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState StartUpdateResponse_SdkV2) {
}

func (c StartUpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["update_id"] = attrs["update_id"].SetOptional()

	return attrs
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

func (toState *StopPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StopPipelineResponse_SdkV2) {
}

func (toState *StopPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState StopPipelineResponse_SdkV2) {
}

func (c StopPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type StopRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
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
	DestinationCatalog types.String `tfsdk:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema types.String `tfsdk:"destination_schema"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable types.String `tfsdk:"destination_table"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog types.String `tfsdk:"source_catalog"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema types.String `tfsdk:"source_schema"`
	// Required. Table name in the source database.
	SourceTable types.String `tfsdk:"source_table"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration types.List `tfsdk:"table_configuration"`
}

func (toState *TableSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TableSpec_SdkV2) {
	if !fromPlan.TableConfiguration.IsNull() && !fromPlan.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromPlanTableConfiguration, ok := fromPlan.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (toState *TableSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TableSpec_SdkV2) {
	if !fromState.TableConfiguration.IsNull() && !fromState.TableConfiguration.IsUnknown() {
		if toStateTableConfiguration, ok := toState.GetTableConfiguration(ctx); ok {
			if fromStateTableConfiguration, ok := fromState.GetTableConfiguration(ctx); ok {
				toStateTableConfiguration.SyncFieldsDuringRead(ctx, fromStateTableConfiguration)
				toState.SetTableConfiguration(ctx, toStateTableConfiguration)
			}
		}
	}
}

func (c TableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["destination_table"] = attrs["destination_table"].SetOptional()
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()
	attrs["source_schema"] = attrs["source_schema"].SetOptional()
	attrs["source_table"] = attrs["source_table"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()
	attrs["table_configuration"] = attrs["table_configuration"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
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
				ElemType: TableSpecificConfig_SdkV2{}.Type(ctx),
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
	// A list of column names to be excluded for the ingestion. When not
	// specified, include_columns fully controls what columns to be ingested.
	// When specified, all other columns including future ones will be
	// automatically included for ingestion. This field in mutually exclusive
	// with `include_columns`.
	ExcludeColumns types.List `tfsdk:"exclude_columns"`
	// A list of column names to be included for the ingestion. When not
	// specified, all columns except ones in exclude_columns will be included.
	// Future columns will be automatically included. When specified, all other
	// future columns will be automatically excluded from ingestion. This field
	// in mutually exclusive with `exclude_columns`.
	IncludeColumns types.List `tfsdk:"include_columns"`
	// The primary key of the table used to apply changes.
	PrimaryKeys types.List `tfsdk:"primary_keys"`

	QueryBasedConnectorConfig types.List `tfsdk:"query_based_connector_config"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields types.Bool `tfsdk:"salesforce_include_formula_fields"`
	// The SCD type to use to ingest the table.
	ScdType types.String `tfsdk:"scd_type"`
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	SequenceBy types.List `tfsdk:"sequence_by"`
}

func (toState *TableSpecificConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TableSpecificConfig_SdkV2) {
	if !fromPlan.QueryBasedConnectorConfig.IsNull() && !fromPlan.QueryBasedConnectorConfig.IsUnknown() {
		if toStateQueryBasedConnectorConfig, ok := toState.GetQueryBasedConnectorConfig(ctx); ok {
			if fromPlanQueryBasedConnectorConfig, ok := fromPlan.GetQueryBasedConnectorConfig(ctx); ok {
				toStateQueryBasedConnectorConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanQueryBasedConnectorConfig)
				toState.SetQueryBasedConnectorConfig(ctx, toStateQueryBasedConnectorConfig)
			}
		}
	}
}

func (toState *TableSpecificConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TableSpecificConfig_SdkV2) {
	if !fromState.QueryBasedConnectorConfig.IsNull() && !fromState.QueryBasedConnectorConfig.IsUnknown() {
		if toStateQueryBasedConnectorConfig, ok := toState.GetQueryBasedConnectorConfig(ctx); ok {
			if fromStateQueryBasedConnectorConfig, ok := fromState.GetQueryBasedConnectorConfig(ctx); ok {
				toStateQueryBasedConnectorConfig.SyncFieldsDuringRead(ctx, fromStateQueryBasedConnectorConfig)
				toState.SetQueryBasedConnectorConfig(ctx, toStateQueryBasedConnectorConfig)
			}
		}
	}
}

func (c TableSpecificConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exclude_columns"] = attrs["exclude_columns"].SetOptional()
	attrs["include_columns"] = attrs["include_columns"].SetOptional()
	attrs["primary_keys"] = attrs["primary_keys"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["salesforce_include_formula_fields"] = attrs["salesforce_include_formula_fields"].SetOptional()
	attrs["scd_type"] = attrs["scd_type"].SetOptional()
	attrs["sequence_by"] = attrs["sequence_by"].SetOptional()

	return attrs
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
		"exclude_columns":              reflect.TypeOf(types.String{}),
		"include_columns":              reflect.TypeOf(types.String{}),
		"primary_keys":                 reflect.TypeOf(types.String{}),
		"query_based_connector_config": reflect.TypeOf(IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2{}),
		"sequence_by":                  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpecificConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o TableSpecificConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude_columns":                   o.ExcludeColumns,
			"include_columns":                   o.IncludeColumns,
			"primary_keys":                      o.PrimaryKeys,
			"query_based_connector_config":      o.QueryBasedConnectorConfig,
			"salesforce_include_formula_fields": o.SalesforceIncludeFormulaFields,
			"scd_type":                          o.ScdType,
			"sequence_by":                       o.SequenceBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSpecificConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exclude_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"include_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"query_based_connector_config": basetypes.ListType{
				ElemType: IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2{}.Type(ctx),
			},
			"salesforce_include_formula_fields": types.BoolType,
			"scd_type":                          types.StringType,
			"sequence_by": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetExcludeColumns returns the value of the ExcludeColumns field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig_SdkV2) GetExcludeColumns(ctx context.Context) ([]types.String, bool) {
	if o.ExcludeColumns.IsNull() || o.ExcludeColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ExcludeColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExcludeColumns sets the value of the ExcludeColumns field in TableSpecificConfig_SdkV2.
func (o *TableSpecificConfig_SdkV2) SetExcludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExcludeColumns = types.ListValueMust(t, vs)
}

// GetIncludeColumns returns the value of the IncludeColumns field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig_SdkV2) GetIncludeColumns(ctx context.Context) ([]types.String, bool) {
	if o.IncludeColumns.IsNull() || o.IncludeColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IncludeColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIncludeColumns sets the value of the IncludeColumns field in TableSpecificConfig_SdkV2.
func (o *TableSpecificConfig_SdkV2) SetIncludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IncludeColumns = types.ListValueMust(t, vs)
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

// GetQueryBasedConnectorConfig returns the value of the QueryBasedConnectorConfig field in TableSpecificConfig_SdkV2 as
// a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig_SdkV2) GetQueryBasedConnectorConfig(ctx context.Context) (IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2, bool) {
	var e IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
	if o.QueryBasedConnectorConfig.IsNull() || o.QueryBasedConnectorConfig.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
	d := o.QueryBasedConnectorConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryBasedConnectorConfig sets the value of the QueryBasedConnectorConfig field in TableSpecificConfig_SdkV2.
func (o *TableSpecificConfig_SdkV2) SetQueryBasedConnectorConfig(ctx context.Context, v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_based_connector_config"]
	o.QueryBasedConnectorConfig = types.ListValueMust(t, vs)
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
	Cause types.String `tfsdk:"cause"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config types.List `tfsdk:"config"`
	// The time when this update was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// If true, this update will reset all tables before running.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection types.List `tfsdk:"full_refresh_selection"`
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection types.List `tfsdk:"refresh_selection"`
	// The update state.
	State types.String `tfsdk:"state"`
	// The ID of this update.
	UpdateId types.String `tfsdk:"update_id"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly types.Bool `tfsdk:"validate_only"`
}

func (toState *UpdateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateInfo_SdkV2) {
	if !fromPlan.Config.IsNull() && !fromPlan.Config.IsUnknown() {
		if toStateConfig, ok := toState.GetConfig(ctx); ok {
			if fromPlanConfig, ok := fromPlan.GetConfig(ctx); ok {
				toStateConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanConfig)
				toState.SetConfig(ctx, toStateConfig)
			}
		}
	}
}

func (toState *UpdateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateInfo_SdkV2) {
	if !fromState.Config.IsNull() && !fromState.Config.IsUnknown() {
		if toStateConfig, ok := toState.GetConfig(ctx); ok {
			if fromStateConfig, ok := fromState.GetConfig(ctx); ok {
				toStateConfig.SyncFieldsDuringRead(ctx, fromStateConfig)
				toState.SetConfig(ctx, toStateConfig)
			}
		}
	}
}

func (c UpdateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["full_refresh"] = attrs["full_refresh"].SetOptional()
	attrs["full_refresh_selection"] = attrs["full_refresh_selection"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["refresh_selection"] = attrs["refresh_selection"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["update_id"] = attrs["update_id"].SetOptional()
	attrs["validate_only"] = attrs["validate_only"].SetOptional()

	return attrs
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
				ElemType: PipelineSpec_SdkV2{}.Type(ctx),
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
	CreationTime types.String `tfsdk:"creation_time"`

	State types.String `tfsdk:"state"`

	UpdateId types.String `tfsdk:"update_id"`
}

func (toState *UpdateStateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateStateInfo_SdkV2) {
}

func (toState *UpdateStateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateStateInfo_SdkV2) {
}

func (c UpdateStateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["update_id"] = attrs["update_id"].SetOptional()

	return attrs
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
