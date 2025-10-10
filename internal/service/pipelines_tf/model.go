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

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf" // .tmpl
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreatePipeline struct {
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
	Deployment types.Object `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`

	DryRun types.Bool `tfsdk:"dry_run"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.Object `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.Object `tfsdk:"event_log"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Restart window of this pipeline.
	RestartWindow types.Object `tfsdk:"restart_window"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath types.String `tfsdk:"root_path"`

	RunAs types.Object `tfsdk:"run_as"`
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
	Trigger types.Object `tfsdk:"trigger"`
}

func (to *CreatePipeline) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePipeline) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				// Recursively sync the fields of Deployment
				toDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				// Recursively sync the fields of Environment
				toEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				// Recursively sync the fields of EventLog
				toEventLog.SyncFieldsDuringCreateOrUpdate(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				// Recursively sync the fields of Filters
				toFilters.SyncFieldsDuringCreateOrUpdate(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				// Recursively sync the fields of GatewayDefinition
				toGatewayDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				// Recursively sync the fields of IngestionDefinition
				toIngestionDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				// Recursively sync the fields of RestartWindow
				toRestartWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				// Recursively sync the fields of RunAs
				toRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				// Recursively sync the fields of Trigger
				toTrigger.SyncFieldsDuringCreateOrUpdate(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (to *CreatePipeline) SyncFieldsDuringRead(ctx context.Context, from CreatePipeline) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				toDeployment.SyncFieldsDuringRead(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				toEnvironment.SyncFieldsDuringRead(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				toEventLog.SyncFieldsDuringRead(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				toFilters.SyncFieldsDuringRead(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				toGatewayDefinition.SyncFieldsDuringRead(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				toIngestionDefinition.SyncFieldsDuringRead(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				toRestartWindow.SyncFieldsDuringRead(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				toRunAs.SyncFieldsDuringRead(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				toTrigger.SyncFieldsDuringRead(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (c CreatePipeline) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_duplicate_names"] = attrs["allow_duplicate_names"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["development"] = attrs["development"].SetOptional()
	attrs["dry_run"] = attrs["dry_run"].SetOptional()
	attrs["edition"] = attrs["edition"].SetOptional()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["event_log"] = attrs["event_log"].SetOptional()
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["gateway_definition"] = attrs["gateway_definition"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["ingestion_definition"] = attrs["ingestion_definition"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["photon"] = attrs["photon"].SetOptional()
	attrs["restart_window"] = attrs["restart_window"].SetOptional()
	attrs["root_path"] = attrs["root_path"].SetOptional()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()

	return attrs
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
		"environment":          reflect.TypeOf(PipelinesEnvironment{}),
		"event_log":            reflect.TypeOf(EventLogSpec{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"run_as":               reflect.TypeOf(RunAs{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":          PipelinesEnvironment{}.Type(ctx),
			"event_log":            EventLogSpec{}.Type(ctx),
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
			"root_path":      types.StringType,
			"run_as":         RunAs{}.Type(ctx),
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target":  types.StringType,
			"trigger": PipelineTrigger{}.Type(ctx),
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
	var v PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in CreatePipeline.
func (o *CreatePipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in CreatePipeline as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := o.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in CreatePipeline.
func (o *CreatePipeline) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	o.Environment = vs
}

// GetEventLog returns the value of the EventLog field in CreatePipeline as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := o.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in CreatePipeline.
func (o *CreatePipeline) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	o.EventLog = vs
}

// GetFilters returns the value of the Filters field in CreatePipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in CreatePipeline.
func (o *CreatePipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetRunAs returns the value of the RunAs field in CreatePipeline as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := o.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in CreatePipeline.
func (o *CreatePipeline) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	o.RunAs = vs
}

// GetTags returns the value of the Tags field in CreatePipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in CreatePipeline.
func (o *CreatePipeline) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in CreatePipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in CreatePipeline.
func (o *CreatePipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings types.Object `tfsdk:"effective_settings"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (to *CreatePipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePipelineResponse) {
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				// Recursively sync the fields of EffectiveSettings
				toEffectiveSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
}

func (to *CreatePipelineResponse) SyncFieldsDuringRead(ctx context.Context, from CreatePipelineResponse) {
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
}

func (c CreatePipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_settings"] = attrs["effective_settings"].SetOptional()
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
	var v PipelineSpec
	d := o.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in CreatePipelineResponse.
func (o *CreatePipelineResponse) SetEffectiveSettings(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveSettings = vs
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`

	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (to *CronTrigger) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronTrigger) {
}

func (to *CronTrigger) SyncFieldsDuringRead(ctx context.Context, from CronTrigger) {
}

func (c CronTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Instance types.String `tfsdk:"instance"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo types.Int64 `tfsdk:"seq_no"`
}

func (to *DataPlaneId) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataPlaneId) {
}

func (to *DataPlaneId) SyncFieldsDuringRead(ctx context.Context, from DataPlaneId) {
}

func (c DataPlaneId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeletePipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *DeletePipelineRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePipelineRequest) {
}

func (to *DeletePipelineRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePipelineRequest) {
}

func (c DeletePipelineRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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

func (to *DeletePipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePipelineResponse) {
}

func (to *DeletePipelineResponse) SyncFieldsDuringRead(ctx context.Context, from DeletePipelineResponse) {
}

func (c DeletePipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Deployment types.Object `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.Object `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.Object `tfsdk:"event_log"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified types.Int64 `tfsdk:"expected_last_modified"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition"`
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
	RestartWindow types.Object `tfsdk:"restart_window"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	RootPath types.String `tfsdk:"root_path"`

	RunAs types.Object `tfsdk:"run_as"`
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
	Trigger types.Object `tfsdk:"trigger"`
}

func (to *EditPipeline) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPipeline) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				// Recursively sync the fields of Deployment
				toDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				// Recursively sync the fields of Environment
				toEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				// Recursively sync the fields of EventLog
				toEventLog.SyncFieldsDuringCreateOrUpdate(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				// Recursively sync the fields of Filters
				toFilters.SyncFieldsDuringCreateOrUpdate(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				// Recursively sync the fields of GatewayDefinition
				toGatewayDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				// Recursively sync the fields of IngestionDefinition
				toIngestionDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				// Recursively sync the fields of RestartWindow
				toRestartWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				// Recursively sync the fields of RunAs
				toRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				// Recursively sync the fields of Trigger
				toTrigger.SyncFieldsDuringCreateOrUpdate(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (to *EditPipeline) SyncFieldsDuringRead(ctx context.Context, from EditPipeline) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				toDeployment.SyncFieldsDuringRead(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				toEnvironment.SyncFieldsDuringRead(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				toEventLog.SyncFieldsDuringRead(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				toFilters.SyncFieldsDuringRead(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				toGatewayDefinition.SyncFieldsDuringRead(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				toIngestionDefinition.SyncFieldsDuringRead(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				toRestartWindow.SyncFieldsDuringRead(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				toRunAs.SyncFieldsDuringRead(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				toTrigger.SyncFieldsDuringRead(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (c EditPipeline) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_duplicate_names"] = attrs["allow_duplicate_names"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["development"] = attrs["development"].SetOptional()
	attrs["edition"] = attrs["edition"].SetOptional()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["event_log"] = attrs["event_log"].SetOptional()
	attrs["expected_last_modified"] = attrs["expected_last_modified"].SetOptional()
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["gateway_definition"] = attrs["gateway_definition"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["ingestion_definition"] = attrs["ingestion_definition"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["photon"] = attrs["photon"].SetOptional()
	attrs["restart_window"] = attrs["restart_window"].SetOptional()
	attrs["root_path"] = attrs["root_path"].SetOptional()
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
		"environment":          reflect.TypeOf(PipelinesEnvironment{}),
		"event_log":            reflect.TypeOf(EventLogSpec{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"run_as":               reflect.TypeOf(RunAs{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":            PipelinesEnvironment{}.Type(ctx),
			"event_log":              EventLogSpec{}.Type(ctx),
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
			"root_path":      types.StringType,
			"run_as":         RunAs{}.Type(ctx),
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target":  types.StringType,
			"trigger": PipelineTrigger{}.Type(ctx),
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
	var v PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in EditPipeline.
func (o *EditPipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in EditPipeline as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := o.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in EditPipeline.
func (o *EditPipeline) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	o.Environment = vs
}

// GetEventLog returns the value of the EventLog field in EditPipeline as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := o.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in EditPipeline.
func (o *EditPipeline) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	o.EventLog = vs
}

// GetFilters returns the value of the Filters field in EditPipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in EditPipeline.
func (o *EditPipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetRunAs returns the value of the RunAs field in EditPipeline as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := o.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in EditPipeline.
func (o *EditPipeline) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	o.RunAs = vs
}

// GetTags returns the value of the Tags field in EditPipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in EditPipeline.
func (o *EditPipeline) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in EditPipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *EditPipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in EditPipeline.
func (o *EditPipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type EditPipelineResponse struct {
}

func (to *EditPipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPipelineResponse) {
}

func (to *EditPipelineResponse) SyncFieldsDuringRead(ctx context.Context, from EditPipelineResponse) {
}

func (c EditPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	Exceptions types.List `tfsdk:"exceptions"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal types.Bool `tfsdk:"fatal"`
}

func (to *ErrorDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ErrorDetail) {
	if !from.Exceptions.IsNull() && !from.Exceptions.IsUnknown() && to.Exceptions.IsNull() && len(from.Exceptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exceptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exceptions = from.Exceptions
	}
}

func (to *ErrorDetail) SyncFieldsDuringRead(ctx context.Context, from ErrorDetail) {
	if !from.Exceptions.IsNull() && !from.Exceptions.IsUnknown() && to.Exceptions.IsNull() && len(from.Exceptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exceptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exceptions = from.Exceptions
	}
}

func (c ErrorDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

// Configurable event log parameters.
type EventLogSpec struct {
	// The UC catalog the event log is published under.
	Catalog types.String `tfsdk:"catalog"`
	// The name the event log is published to in UC.
	Name types.String `tfsdk:"name"`
	// The UC schema the event log is published under.
	Schema types.String `tfsdk:"schema"`
}

func (to *EventLogSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EventLogSpec) {
}

func (to *EventLogSpec) SyncFieldsDuringRead(ctx context.Context, from EventLogSpec) {
}

func (c EventLogSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EventLogSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventLogSpec
// only implements ToObjectValue() and Type().
func (o EventLogSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
			"name":    o.Name,
			"schema":  o.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EventLogSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": types.StringType,
			"name":    types.StringType,
			"schema":  types.StringType,
		},
	}
}

type FileLibrary struct {
	// The absolute path of the source code.
	Path types.String `tfsdk:"path"`
}

func (to *FileLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileLibrary) {
}

func (to *FileLibrary) SyncFieldsDuringRead(ctx context.Context, from FileLibrary) {
}

func (c FileLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Exclude types.List `tfsdk:"exclude"`
	// Paths to include.
	Include types.List `tfsdk:"include"`
}

func (to *Filters) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Filters) {
	if !from.Exclude.IsNull() && !from.Exclude.IsUnknown() && to.Exclude.IsNull() && len(from.Exclude.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exclude, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exclude = from.Exclude
	}
	if !from.Include.IsNull() && !from.Include.IsUnknown() && to.Include.IsNull() && len(from.Include.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Include, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Include = from.Include
	}
}

func (to *Filters) SyncFieldsDuringRead(ctx context.Context, from Filters) {
	if !from.Exclude.IsNull() && !from.Exclude.IsUnknown() && to.Exclude.IsNull() && len(from.Exclude.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exclude, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exclude = from.Exclude
	}
	if !from.Include.IsNull() && !from.Include.IsUnknown() && to.Include.IsNull() && len(from.Include.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Include, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Include = from.Include
	}
}

func (c Filters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionLevelsRequest) {
}

func (to *GetPipelinePermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionLevelsRequest) {
}

func (c GetPipelinePermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetPipelinePermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetPipelinePermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (c GetPipelinePermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionsRequest) {
}

func (to *GetPipelinePermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionsRequest) {
}

func (c GetPipelinePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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

type GetPipelineRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelineRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelineRequest) {
}

func (to *GetPipelineRequest) SyncFieldsDuringRead(ctx context.Context, from GetPipelineRequest) {
}

func (c GetPipelineRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
	RunAs types.Object `tfsdk:"run_as"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec types.Object `tfsdk:"spec"`
	// The pipeline state.
	State types.String `tfsdk:"state"`
}

func (to *GetPipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelineResponse) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				// Recursively sync the fields of RunAs
				toRunAs.SyncFieldsDuringCreateOrUpdate(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
}

func (to *GetPipelineResponse) SyncFieldsDuringRead(ctx context.Context, from GetPipelineResponse) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
	if !from.RunAs.IsNull() && !from.RunAs.IsUnknown() {
		if toRunAs, ok := to.GetRunAs(ctx); ok {
			if fromRunAs, ok := from.GetRunAs(ctx); ok {
				toRunAs.SyncFieldsDuringRead(ctx, fromRunAs)
				to.SetRunAs(ctx, toRunAs)
			}
		}
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
}

func (c GetPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["run_as_user_name"] = attrs["run_as_user_name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
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
func (a GetPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
		"run_as":         reflect.TypeOf(RunAs{}),
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
			"run_as":                     o.RunAs,
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
			"run_as":           RunAs{}.Type(ctx),
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

// GetRunAs returns the value of the RunAs field in GetPipelineResponse as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := o.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in GetPipelineResponse.
func (o *GetPipelineResponse) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	o.RunAs = vs
}

// GetSpec returns the value of the Spec field in GetPipelineResponse as
// a PipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPipelineResponse) GetSpec(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v PipelineSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in GetPipelineResponse.
func (o *GetPipelineResponse) SetSpec(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
}

func (to *GetUpdateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUpdateRequest) {
}

func (to *GetUpdateRequest) SyncFieldsDuringRead(ctx context.Context, from GetUpdateRequest) {
}

func (c GetUpdateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()
	attrs["update_id"] = attrs["update_id"].SetRequired()

	return attrs
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
	Update types.Object `tfsdk:"update"`
}

func (to *GetUpdateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUpdateResponse) {
	if !from.Update.IsNull() && !from.Update.IsUnknown() {
		if toUpdate, ok := to.GetUpdate(ctx); ok {
			if fromUpdate, ok := from.GetUpdate(ctx); ok {
				// Recursively sync the fields of Update
				toUpdate.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdate)
				to.SetUpdate(ctx, toUpdate)
			}
		}
	}
}

func (to *GetUpdateResponse) SyncFieldsDuringRead(ctx context.Context, from GetUpdateResponse) {
	if !from.Update.IsNull() && !from.Update.IsUnknown() {
		if toUpdate, ok := to.GetUpdate(ctx); ok {
			if fromUpdate, ok := from.GetUpdate(ctx); ok {
				toUpdate.SyncFieldsDuringRead(ctx, fromUpdate)
				to.SetUpdate(ctx, toUpdate)
			}
		}
	}
}

func (c GetUpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["update"] = attrs["update"].SetOptional()

	return attrs
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
	var v UpdateInfo
	d := o.Update.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdate sets the value of the Update field in GetUpdateResponse.
func (o *GetUpdateResponse) SetUpdate(ctx context.Context, v UpdateInfo) {
	vs := v.ToObjectValue(ctx)
	o.Update = vs
}

type IngestionConfig struct {
	// Select a specific source report.
	Report types.Object `tfsdk:"report"`
	// Select all tables from a specific source schema.
	Schema types.Object `tfsdk:"schema"`
	// Select a specific source table.
	Table types.Object `tfsdk:"table"`
}

func (to *IngestionConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionConfig) {
	if !from.Report.IsNull() && !from.Report.IsUnknown() {
		if toReport, ok := to.GetReport(ctx); ok {
			if fromReport, ok := from.GetReport(ctx); ok {
				// Recursively sync the fields of Report
				toReport.SyncFieldsDuringCreateOrUpdate(ctx, fromReport)
				to.SetReport(ctx, toReport)
			}
		}
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				// Recursively sync the fields of Schema
				toSchema.SyncFieldsDuringCreateOrUpdate(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				// Recursively sync the fields of Table
				toTable.SyncFieldsDuringCreateOrUpdate(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (to *IngestionConfig) SyncFieldsDuringRead(ctx context.Context, from IngestionConfig) {
	if !from.Report.IsNull() && !from.Report.IsUnknown() {
		if toReport, ok := to.GetReport(ctx); ok {
			if fromReport, ok := from.GetReport(ctx); ok {
				toReport.SyncFieldsDuringRead(ctx, fromReport)
				to.SetReport(ctx, toReport)
			}
		}
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				toSchema.SyncFieldsDuringRead(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				toTable.SyncFieldsDuringRead(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (c IngestionConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["report"] = attrs["report"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["table"] = attrs["table"].SetOptional()

	return attrs
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
	var v ReportSpec
	d := o.Report.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v SchemaSpec
	d := o.Schema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v TableSpec
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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

func (to *IngestionGatewayPipelineDefinition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionGatewayPipelineDefinition) {
}

func (to *IngestionGatewayPipelineDefinition) SyncFieldsDuringRead(ctx context.Context, from IngestionGatewayPipelineDefinition) {
}

func (c IngestionGatewayPipelineDefinition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	ConnectionName types.String `tfsdk:"connection_name"`
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	IngestionGatewayId types.String `tfsdk:"ingestion_gateway_id"`
	// Netsuite only configuration. When the field is set for a netsuite
	// connector, the jar stored in the field will be validated and added to the
	// classpath of pipeline's cluster.
	NetsuiteJarPath types.String `tfsdk:"netsuite_jar_path"`
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
	TableConfiguration types.Object `tfsdk:"table_configuration"`
}

func (to *IngestionPipelineDefinition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinition) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
	if !from.SourceConfigurations.IsNull() && !from.SourceConfigurations.IsUnknown() && to.SourceConfigurations.IsNull() && len(from.SourceConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SourceConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SourceConfigurations = from.SourceConfigurations
	}
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				// Recursively sync the fields of TableConfiguration
				toTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (to *IngestionPipelineDefinition) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinition) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
	if !from.SourceConfigurations.IsNull() && !from.SourceConfigurations.IsUnknown() && to.SourceConfigurations.IsNull() && len(from.SourceConfigurations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SourceConfigurations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SourceConfigurations = from.SourceConfigurations
	}
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (c IngestionPipelineDefinition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["ingestion_gateway_id"] = attrs["ingestion_gateway_id"].SetOptional()
	attrs["netsuite_jar_path"] = attrs["netsuite_jar_path"].SetOptional()
	attrs["objects"] = attrs["objects"].SetOptional()
	attrs["source_configurations"] = attrs["source_configurations"].SetOptional()
	attrs["source_type"] = attrs["source_type"].SetComputed()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()

	return attrs
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
		"objects":               reflect.TypeOf(IngestionConfig{}),
		"source_configurations": reflect.TypeOf(SourceConfig{}),
		"table_configuration":   reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name":       o.ConnectionName,
			"ingestion_gateway_id":  o.IngestionGatewayId,
			"netsuite_jar_path":     o.NetsuiteJarPath,
			"objects":               o.Objects,
			"source_configurations": o.SourceConfigurations,
			"source_type":           o.SourceType,
			"table_configuration":   o.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":      types.StringType,
			"ingestion_gateway_id": types.StringType,
			"netsuite_jar_path":    types.StringType,
			"objects": basetypes.ListType{
				ElemType: IngestionConfig{}.Type(ctx),
			},
			"source_configurations": basetypes.ListType{
				ElemType: SourceConfig{}.Type(ctx),
			},
			"source_type":         types.StringType,
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

// GetSourceConfigurations returns the value of the SourceConfigurations field in IngestionPipelineDefinition as
// a slice of SourceConfig values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition) GetSourceConfigurations(ctx context.Context) ([]SourceConfig, bool) {
	if o.SourceConfigurations.IsNull() || o.SourceConfigurations.IsUnknown() {
		return nil, false
	}
	var v []SourceConfig
	d := o.SourceConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSourceConfigurations sets the value of the SourceConfigurations field in IngestionPipelineDefinition.
func (o *IngestionPipelineDefinition) SetSourceConfigurations(ctx context.Context, v []SourceConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["source_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SourceConfigurations = types.ListValueMust(t, vs)
}

// GetTableConfiguration returns the value of the TableConfiguration field in IngestionPipelineDefinition as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinition) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if o.TableConfiguration.IsNull() || o.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in IngestionPipelineDefinition.
func (o *IngestionPipelineDefinition) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

// Configurations that are only applicable for query-based ingestion connectors.
type IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig struct {
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

func (to *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) {
	if !from.CursorColumns.IsNull() && !from.CursorColumns.IsUnknown() && to.CursorColumns.IsNull() && len(from.CursorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CursorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CursorColumns = from.CursorColumns
	}
}

func (to *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) {
	if !from.CursorColumns.IsNull() && !from.CursorColumns.IsUnknown() && to.CursorColumns.IsNull() && len(from.CursorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CursorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CursorColumns = from.CursorColumns
	}
}

func (c IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cursor_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cursor_columns":                             o.CursorColumns,
			"deletion_condition":                         o.DeletionCondition,
			"hard_deletion_sync_min_interval_in_seconds": o.HardDeletionSyncMinIntervalInSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) Type(ctx context.Context) attr.Type {
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

// GetCursorColumns returns the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) GetCursorColumns(ctx context.Context) ([]types.String, bool) {
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

// SetCursorColumns sets the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig.
func (o *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) SetCursorColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cursor_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CursorColumns = types.ListValueMust(t, vs)
}

type IngestionPipelineDefinitionWorkdayReportParameters struct {
	// (Optional) Marks the report as incremental. This field is deprecated and
	// should not be used. Use `parameters` instead. The incremental behavior is
	// now controlled by the `parameters` field.
	Incremental types.Bool `tfsdk:"incremental"`
	// Parameters for the Workday report. Each key represents the parameter name
	// (e.g., "start_date", "end_date"), and the corresponding value is a
	// SQL-like expression used to compute the parameter value at runtime.
	// Example: { "start_date": "{ coalesce(current_offset(),
	// date(\"2025-02-01\")) }", "end_date": "{ current_date() - INTERVAL 1 DAY
	// }" }
	Parameters types.Map `tfsdk:"parameters"`
	// (Optional) Additional custom parameters for Workday Report This field is
	// deprecated and should not be used. Use `parameters` instead.
	ReportParameters types.List `tfsdk:"report_parameters"`
}

func (to *IngestionPipelineDefinitionWorkdayReportParameters) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParameters) {
	if !from.ReportParameters.IsNull() && !from.ReportParameters.IsUnknown() && to.ReportParameters.IsNull() && len(from.ReportParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ReportParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ReportParameters = from.ReportParameters
	}
}

func (to *IngestionPipelineDefinitionWorkdayReportParameters) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParameters) {
	if !from.ReportParameters.IsNull() && !from.ReportParameters.IsUnknown() && to.ReportParameters.IsNull() && len(from.ReportParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ReportParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ReportParameters = from.ReportParameters
	}
}

func (c IngestionPipelineDefinitionWorkdayReportParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["incremental"] = attrs["incremental"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["report_parameters"] = attrs["report_parameters"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionPipelineDefinitionWorkdayReportParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionPipelineDefinitionWorkdayReportParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters":        reflect.TypeOf(types.String{}),
		"report_parameters": reflect.TypeOf(IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParameters
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinitionWorkdayReportParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"incremental":       o.Incremental,
			"parameters":        o.Parameters,
			"report_parameters": o.ReportParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinitionWorkdayReportParameters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"incremental": types.BoolType,
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"report_parameters": basetypes.ListType{
				ElemType: IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in IngestionPipelineDefinitionWorkdayReportParameters as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinitionWorkdayReportParameters) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in IngestionPipelineDefinitionWorkdayReportParameters.
func (o *IngestionPipelineDefinitionWorkdayReportParameters) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

// GetReportParameters returns the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters as
// a slice of IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *IngestionPipelineDefinitionWorkdayReportParameters) GetReportParameters(ctx context.Context) ([]IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue, bool) {
	if o.ReportParameters.IsNull() || o.ReportParameters.IsUnknown() {
		return nil, false
	}
	var v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue
	d := o.ReportParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReportParameters sets the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters.
func (o *IngestionPipelineDefinitionWorkdayReportParameters) SetReportParameters(ctx context.Context, v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["report_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ReportParameters = types.ListValueMust(t, vs)
}

type IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue struct {
	// Key for the report parameter, can be a column name or other metadata
	Key types.String `tfsdk:"key"`
	// Value for the report parameter. Possible values it can take are these sql
	// functions: 1. coalesce(current_offset(), date("YYYY-MM-DD")) -> if
	// current_offset() is null, then the passed date, else current_offset() 2.
	// current_date() 3. date_sub(current_date(), x) -> subtract x (some
	// non-negative integer) days from current date
	Value types.String `tfsdk:"value"`
}

func (to *IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) {
}

func (to *IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) {
}

func (c IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue
// only implements ToObjectValue() and Type().
func (o IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

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
	// The pipeline to return events for.
	PipelineId types.String `tfsdk:"-"`
}

func (to *ListPipelineEventsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelineEventsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *ListPipelineEventsRequest) SyncFieldsDuringRead(ctx context.Context, from ListPipelineEventsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (c ListPipelineEventsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
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
	Events types.List `tfsdk:"events"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (to *ListPipelineEventsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelineEventsResponse) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (to *ListPipelineEventsResponse) SyncFieldsDuringRead(ctx context.Context, from ListPipelineEventsResponse) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (c ListPipelineEventsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *ListPipelinesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelinesRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *ListPipelinesRequest) SyncFieldsDuringRead(ctx context.Context, from ListPipelinesRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (c ListPipelinesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of events matching the request criteria.
	Statuses types.List `tfsdk:"statuses"`
}

func (to *ListPipelinesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelinesResponse) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (to *ListPipelinesResponse) SyncFieldsDuringRead(ctx context.Context, from ListPipelinesResponse) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (c ListPipelinesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *ListUpdatesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUpdatesRequest) {
}

func (to *ListUpdatesRequest) SyncFieldsDuringRead(ctx context.Context, from ListUpdatesRequest) {
}

func (c ListUpdatesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["until_update_id"] = attrs["until_update_id"].SetOptional()

	return attrs
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken types.String `tfsdk:"prev_page_token"`

	Updates types.List `tfsdk:"updates"`
}

func (to *ListUpdatesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUpdatesResponse) {
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (to *ListUpdatesResponse) SyncFieldsDuringRead(ctx context.Context, from ListUpdatesResponse) {
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (c ListUpdatesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *ManualTrigger) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManualTrigger) {
}

func (to *ManualTrigger) SyncFieldsDuringRead(ctx context.Context, from ManualTrigger) {
}

func (c ManualTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	// The absolute path of the source code.
	Path types.String `tfsdk:"path"`
}

func (to *NotebookLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotebookLibrary) {
}

func (to *NotebookLibrary) SyncFieldsDuringRead(ctx context.Context, from NotebookLibrary) {
}

func (c NotebookLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Alerts types.List `tfsdk:"alerts"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients types.List `tfsdk:"email_recipients"`
}

func (to *Notifications) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Notifications) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
	if !from.EmailRecipients.IsNull() && !from.EmailRecipients.IsUnknown() && to.EmailRecipients.IsNull() && len(from.EmailRecipients.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailRecipients, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailRecipients = from.EmailRecipients
	}
}

func (to *Notifications) SyncFieldsDuringRead(ctx context.Context, from Notifications) {
	if !from.Alerts.IsNull() && !from.Alerts.IsUnknown() && to.Alerts.IsNull() && len(from.Alerts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Alerts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Alerts = from.Alerts
	}
	if !from.EmailRecipients.IsNull() && !from.EmailRecipients.IsUnknown() && to.EmailRecipients.IsNull() && len(from.EmailRecipients.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmailRecipients, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmailRecipients = from.EmailRecipients
	}
}

func (c Notifications) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *Origin) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Origin) {
}

func (to *Origin) SyncFieldsDuringRead(ctx context.Context, from Origin) {
}

func (c Origin) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type PathPattern struct {
	// The source code to include for pipelines
	Include types.String `tfsdk:"include"`
}

func (to *PathPattern) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PathPattern) {
}

func (to *PathPattern) SyncFieldsDuringRead(ctx context.Context, from PathPattern) {
}

func (c PathPattern) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PathPattern) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PathPattern
// only implements ToObjectValue() and Type().
func (o PathPattern) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include": o.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PathPattern) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include": types.StringType,
		},
	}
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *PipelineAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineAccessControlRequest) {
}

func (to *PipelineAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from PipelineAccessControlRequest) {
}

func (c PipelineAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *PipelineAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *PipelineAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from PipelineAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (c PipelineAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.Object `tfsdk:"autoscale"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes types.Object `tfsdk:"aws_attributes"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes types.Object `tfsdk:"azure_attributes"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf types.Object `tfsdk:"cluster_log_conf"`
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
	GcpAttributes types.Object `tfsdk:"gcp_attributes"`
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

func (to *PipelineCluster) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				// Recursively sync the fields of Autoscale
				toAutoscale.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				// Recursively sync the fields of AwsAttributes
				toAwsAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				// Recursively sync the fields of AzureAttributes
				toAzureAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				// Recursively sync the fields of ClusterLogConf
				toClusterLogConf.SyncFieldsDuringCreateOrUpdate(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				// Recursively sync the fields of GcpAttributes
				toGcpAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
}

func (to *PipelineCluster) SyncFieldsDuringRead(ctx context.Context, from PipelineCluster) {
	if !from.Autoscale.IsNull() && !from.Autoscale.IsUnknown() {
		if toAutoscale, ok := to.GetAutoscale(ctx); ok {
			if fromAutoscale, ok := from.GetAutoscale(ctx); ok {
				toAutoscale.SyncFieldsDuringRead(ctx, fromAutoscale)
				to.SetAutoscale(ctx, toAutoscale)
			}
		}
	}
	if !from.AwsAttributes.IsNull() && !from.AwsAttributes.IsUnknown() {
		if toAwsAttributes, ok := to.GetAwsAttributes(ctx); ok {
			if fromAwsAttributes, ok := from.GetAwsAttributes(ctx); ok {
				toAwsAttributes.SyncFieldsDuringRead(ctx, fromAwsAttributes)
				to.SetAwsAttributes(ctx, toAwsAttributes)
			}
		}
	}
	if !from.AzureAttributes.IsNull() && !from.AzureAttributes.IsUnknown() {
		if toAzureAttributes, ok := to.GetAzureAttributes(ctx); ok {
			if fromAzureAttributes, ok := from.GetAzureAttributes(ctx); ok {
				toAzureAttributes.SyncFieldsDuringRead(ctx, fromAzureAttributes)
				to.SetAzureAttributes(ctx, toAzureAttributes)
			}
		}
	}
	if !from.ClusterLogConf.IsNull() && !from.ClusterLogConf.IsUnknown() {
		if toClusterLogConf, ok := to.GetClusterLogConf(ctx); ok {
			if fromClusterLogConf, ok := from.GetClusterLogConf(ctx); ok {
				toClusterLogConf.SyncFieldsDuringRead(ctx, fromClusterLogConf)
				to.SetClusterLogConf(ctx, toClusterLogConf)
			}
		}
	}
	if !from.GcpAttributes.IsNull() && !from.GcpAttributes.IsUnknown() {
		if toGcpAttributes, ok := to.GetGcpAttributes(ctx); ok {
			if fromGcpAttributes, ok := from.GetGcpAttributes(ctx); ok {
				toGcpAttributes.SyncFieldsDuringRead(ctx, fromGcpAttributes)
				to.SetGcpAttributes(ctx, toGcpAttributes)
			}
		}
	}
	if !from.InitScripts.IsNull() && !from.InitScripts.IsUnknown() && to.InitScripts.IsNull() && len(from.InitScripts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InitScripts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InitScripts = from.InitScripts
	}
	if !from.SshPublicKeys.IsNull() && !from.SshPublicKeys.IsUnknown() && to.SshPublicKeys.IsNull() && len(from.SshPublicKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SshPublicKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SshPublicKeys = from.SshPublicKeys
	}
}

func (c PipelineCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apply_policy_default_values"] = attrs["apply_policy_default_values"].SetOptional()
	attrs["autoscale"] = attrs["autoscale"].SetOptional()
	attrs["aws_attributes"] = attrs["aws_attributes"].SetOptional()
	attrs["azure_attributes"] = attrs["azure_attributes"].SetOptional()
	attrs["cluster_log_conf"] = attrs["cluster_log_conf"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["driver_instance_pool_id"] = attrs["driver_instance_pool_id"].SetOptional()
	attrs["driver_node_type_id"] = attrs["driver_node_type_id"].SetOptional()
	attrs["enable_local_disk_encryption"] = attrs["enable_local_disk_encryption"].SetOptional()
	attrs["gcp_attributes"] = attrs["gcp_attributes"].SetOptional()
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
	var v PipelineClusterAutoscale
	d := o.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v compute_tf.AwsAttributes
	d := o.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v compute_tf.AzureAttributes
	d := o.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v compute_tf.ClusterLogConf
	d := o.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v compute_tf.GcpAttributes
	d := o.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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

func (to *PipelineClusterAutoscale) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineClusterAutoscale) {
}

func (to *PipelineClusterAutoscale) SyncFieldsDuringRead(ctx context.Context, from PipelineClusterAutoscale) {
}

func (c PipelineClusterAutoscale) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Kind types.String `tfsdk:"kind"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath types.String `tfsdk:"metadata_file_path"`
}

func (to *PipelineDeployment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineDeployment) {
}

func (to *PipelineDeployment) SyncFieldsDuringRead(ctx context.Context, from PipelineDeployment) {
}

func (c PipelineDeployment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Error types.Object `tfsdk:"error"`
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
	Origin types.Object `tfsdk:"origin"`
	// A sequencing object to identify and order events.
	Sequence types.Object `tfsdk:"sequence"`
	// The time of the event.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (to *PipelineEvent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineEvent) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
	if !from.Origin.IsNull() && !from.Origin.IsUnknown() {
		if toOrigin, ok := to.GetOrigin(ctx); ok {
			if fromOrigin, ok := from.GetOrigin(ctx); ok {
				// Recursively sync the fields of Origin
				toOrigin.SyncFieldsDuringCreateOrUpdate(ctx, fromOrigin)
				to.SetOrigin(ctx, toOrigin)
			}
		}
	}
	if !from.Sequence.IsNull() && !from.Sequence.IsUnknown() {
		if toSequence, ok := to.GetSequence(ctx); ok {
			if fromSequence, ok := from.GetSequence(ctx); ok {
				// Recursively sync the fields of Sequence
				toSequence.SyncFieldsDuringCreateOrUpdate(ctx, fromSequence)
				to.SetSequence(ctx, toSequence)
			}
		}
	}
}

func (to *PipelineEvent) SyncFieldsDuringRead(ctx context.Context, from PipelineEvent) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
	if !from.Origin.IsNull() && !from.Origin.IsUnknown() {
		if toOrigin, ok := to.GetOrigin(ctx); ok {
			if fromOrigin, ok := from.GetOrigin(ctx); ok {
				toOrigin.SyncFieldsDuringRead(ctx, fromOrigin)
				to.SetOrigin(ctx, toOrigin)
			}
		}
	}
	if !from.Sequence.IsNull() && !from.Sequence.IsUnknown() {
		if toSequence, ok := to.GetSequence(ctx); ok {
			if fromSequence, ok := from.GetSequence(ctx); ok {
				toSequence.SyncFieldsDuringRead(ctx, fromSequence)
				to.SetSequence(ctx, toSequence)
			}
		}
	}
}

func (c PipelineEvent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["event_type"] = attrs["event_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["level"] = attrs["level"].SetOptional()
	attrs["maturity_level"] = attrs["maturity_level"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["origin"] = attrs["origin"].SetOptional()
	attrs["sequence"] = attrs["sequence"].SetOptional()
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
	var v ErrorDetail
	d := o.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v Origin
	d := o.Origin.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v Sequencing
	d := o.Sequence.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSequence sets the value of the Sequence field in PipelineEvent.
func (o *PipelineEvent) SetSequence(ctx context.Context, v Sequencing) {
	vs := v.ToObjectValue(ctx)
	o.Sequence = vs
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File types.Object `tfsdk:"file"`
	// The unified field to include source codes. Each entry can be a notebook
	// path, a file path, or a folder path that ends `/**`. This field cannot be
	// used together with `notebook` or `file`.
	Glob types.Object `tfsdk:"glob"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar"`
	// Specification of a maven library to be installed.
	Maven types.Object `tfsdk:"maven"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook types.Object `tfsdk:"notebook"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl"`
}

func (to *PipelineLibrary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineLibrary) {
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				// Recursively sync the fields of File
				toFile.SyncFieldsDuringCreateOrUpdate(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Glob.IsNull() && !from.Glob.IsUnknown() {
		if toGlob, ok := to.GetGlob(ctx); ok {
			if fromGlob, ok := from.GetGlob(ctx); ok {
				// Recursively sync the fields of Glob
				toGlob.SyncFieldsDuringCreateOrUpdate(ctx, fromGlob)
				to.SetGlob(ctx, toGlob)
			}
		}
	}
	if !from.Maven.IsNull() && !from.Maven.IsUnknown() {
		if toMaven, ok := to.GetMaven(ctx); ok {
			if fromMaven, ok := from.GetMaven(ctx); ok {
				// Recursively sync the fields of Maven
				toMaven.SyncFieldsDuringCreateOrUpdate(ctx, fromMaven)
				to.SetMaven(ctx, toMaven)
			}
		}
	}
	if !from.Notebook.IsNull() && !from.Notebook.IsUnknown() {
		if toNotebook, ok := to.GetNotebook(ctx); ok {
			if fromNotebook, ok := from.GetNotebook(ctx); ok {
				// Recursively sync the fields of Notebook
				toNotebook.SyncFieldsDuringCreateOrUpdate(ctx, fromNotebook)
				to.SetNotebook(ctx, toNotebook)
			}
		}
	}
}

func (to *PipelineLibrary) SyncFieldsDuringRead(ctx context.Context, from PipelineLibrary) {
	if !from.File.IsNull() && !from.File.IsUnknown() {
		if toFile, ok := to.GetFile(ctx); ok {
			if fromFile, ok := from.GetFile(ctx); ok {
				toFile.SyncFieldsDuringRead(ctx, fromFile)
				to.SetFile(ctx, toFile)
			}
		}
	}
	if !from.Glob.IsNull() && !from.Glob.IsUnknown() {
		if toGlob, ok := to.GetGlob(ctx); ok {
			if fromGlob, ok := from.GetGlob(ctx); ok {
				toGlob.SyncFieldsDuringRead(ctx, fromGlob)
				to.SetGlob(ctx, toGlob)
			}
		}
	}
	if !from.Maven.IsNull() && !from.Maven.IsUnknown() {
		if toMaven, ok := to.GetMaven(ctx); ok {
			if fromMaven, ok := from.GetMaven(ctx); ok {
				toMaven.SyncFieldsDuringRead(ctx, fromMaven)
				to.SetMaven(ctx, toMaven)
			}
		}
	}
	if !from.Notebook.IsNull() && !from.Notebook.IsUnknown() {
		if toNotebook, ok := to.GetNotebook(ctx); ok {
			if fromNotebook, ok := from.GetNotebook(ctx); ok {
				toNotebook.SyncFieldsDuringRead(ctx, fromNotebook)
				to.SetNotebook(ctx, toNotebook)
			}
		}
	}
}

func (c PipelineLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file"] = attrs["file"].SetOptional()
	attrs["glob"] = attrs["glob"].SetOptional()
	attrs["jar"] = attrs["jar"].SetOptional()
	attrs["maven"] = attrs["maven"].SetOptional()
	attrs["notebook"] = attrs["notebook"].SetOptional()
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
func (a PipelineLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file":     reflect.TypeOf(FileLibrary{}),
		"glob":     reflect.TypeOf(PathPattern{}),
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
			"glob":     o.Glob,
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
			"glob":     PathPattern{}.Type(ctx),
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
	var v FileLibrary
	d := o.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFile sets the value of the File field in PipelineLibrary.
func (o *PipelineLibrary) SetFile(ctx context.Context, v FileLibrary) {
	vs := v.ToObjectValue(ctx)
	o.File = vs
}

// GetGlob returns the value of the Glob field in PipelineLibrary as
// a PathPattern value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary) GetGlob(ctx context.Context) (PathPattern, bool) {
	var e PathPattern
	if o.Glob.IsNull() || o.Glob.IsUnknown() {
		return e, false
	}
	var v PathPattern
	d := o.Glob.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlob sets the value of the Glob field in PipelineLibrary.
func (o *PipelineLibrary) SetGlob(ctx context.Context, v PathPattern) {
	vs := v.ToObjectValue(ctx)
	o.Glob = vs
}

// GetMaven returns the value of the Maven field in PipelineLibrary as
// a compute_tf.MavenLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineLibrary) GetMaven(ctx context.Context) (compute_tf.MavenLibrary, bool) {
	var e compute_tf.MavenLibrary
	if o.Maven.IsNull() || o.Maven.IsUnknown() {
		return e, false
	}
	var v compute_tf.MavenLibrary
	d := o.Maven.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v NotebookLibrary
	d := o.Notebook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebook sets the value of the Notebook field in PipelineLibrary.
func (o *PipelineLibrary) SetNotebook(ctx context.Context, v NotebookLibrary) {
	vs := v.ToObjectValue(ctx)
	o.Notebook = vs
}

type PipelinePermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *PipelinePermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *PipelinePermission) SyncFieldsDuringRead(ctx context.Context, from PipelinePermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (c PipelinePermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *PipelinePermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *PipelinePermissions) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c PipelinePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *PipelinePermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissionsDescription) {
}

func (to *PipelinePermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissionsDescription) {
}

func (c PipelinePermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *PipelinePermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *PipelinePermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c PipelinePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
	Deployment types.Object `tfsdk:"deployment"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development types.Bool `tfsdk:"development"`
	// Pipeline product edition.
	Edition types.String `tfsdk:"edition"`
	// Environment specification for this pipeline used to install dependencies.
	Environment types.Object `tfsdk:"environment"`
	// Event log configuration for this pipeline
	EventLog types.Object `tfsdk:"event_log"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters types.Object `tfsdk:"filters"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition types.Object `tfsdk:"gateway_definition"`
	// Unique identifier for this pipeline.
	Id types.String `tfsdk:"id"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition types.Object `tfsdk:"ingestion_definition"`
	// Libraries or code needed by this deployment.
	Libraries types.List `tfsdk:"libraries"`
	// Friendly identifier for this pipeline.
	Name types.String `tfsdk:"name"`
	// List of notification settings for this pipeline.
	Notifications types.List `tfsdk:"notifications"`
	// Whether Photon is enabled for this pipeline.
	Photon types.Bool `tfsdk:"photon"`
	// Restart window of this pipeline.
	RestartWindow types.Object `tfsdk:"restart_window"`
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
	Trigger types.Object `tfsdk:"trigger"`
}

func (to *PipelineSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineSpec) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				// Recursively sync the fields of Deployment
				toDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				// Recursively sync the fields of Environment
				toEnvironment.SyncFieldsDuringCreateOrUpdate(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				// Recursively sync the fields of EventLog
				toEventLog.SyncFieldsDuringCreateOrUpdate(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				// Recursively sync the fields of Filters
				toFilters.SyncFieldsDuringCreateOrUpdate(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				// Recursively sync the fields of GatewayDefinition
				toGatewayDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				// Recursively sync the fields of IngestionDefinition
				toIngestionDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				// Recursively sync the fields of RestartWindow
				toRestartWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				// Recursively sync the fields of Trigger
				toTrigger.SyncFieldsDuringCreateOrUpdate(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (to *PipelineSpec) SyncFieldsDuringRead(ctx context.Context, from PipelineSpec) {
	if !from.Clusters.IsNull() && !from.Clusters.IsUnknown() && to.Clusters.IsNull() && len(from.Clusters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Clusters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Clusters = from.Clusters
	}
	if !from.Deployment.IsNull() && !from.Deployment.IsUnknown() {
		if toDeployment, ok := to.GetDeployment(ctx); ok {
			if fromDeployment, ok := from.GetDeployment(ctx); ok {
				toDeployment.SyncFieldsDuringRead(ctx, fromDeployment)
				to.SetDeployment(ctx, toDeployment)
			}
		}
	}
	if !from.Environment.IsNull() && !from.Environment.IsUnknown() {
		if toEnvironment, ok := to.GetEnvironment(ctx); ok {
			if fromEnvironment, ok := from.GetEnvironment(ctx); ok {
				toEnvironment.SyncFieldsDuringRead(ctx, fromEnvironment)
				to.SetEnvironment(ctx, toEnvironment)
			}
		}
	}
	if !from.EventLog.IsNull() && !from.EventLog.IsUnknown() {
		if toEventLog, ok := to.GetEventLog(ctx); ok {
			if fromEventLog, ok := from.GetEventLog(ctx); ok {
				toEventLog.SyncFieldsDuringRead(ctx, fromEventLog)
				to.SetEventLog(ctx, toEventLog)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				toFilters.SyncFieldsDuringRead(ctx, fromFilters)
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.GatewayDefinition.IsNull() && !from.GatewayDefinition.IsUnknown() {
		if toGatewayDefinition, ok := to.GetGatewayDefinition(ctx); ok {
			if fromGatewayDefinition, ok := from.GetGatewayDefinition(ctx); ok {
				toGatewayDefinition.SyncFieldsDuringRead(ctx, fromGatewayDefinition)
				to.SetGatewayDefinition(ctx, toGatewayDefinition)
			}
		}
	}
	if !from.IngestionDefinition.IsNull() && !from.IngestionDefinition.IsUnknown() {
		if toIngestionDefinition, ok := to.GetIngestionDefinition(ctx); ok {
			if fromIngestionDefinition, ok := from.GetIngestionDefinition(ctx); ok {
				toIngestionDefinition.SyncFieldsDuringRead(ctx, fromIngestionDefinition)
				to.SetIngestionDefinition(ctx, toIngestionDefinition)
			}
		}
	}
	if !from.Libraries.IsNull() && !from.Libraries.IsUnknown() && to.Libraries.IsNull() && len(from.Libraries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Libraries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Libraries = from.Libraries
	}
	if !from.Notifications.IsNull() && !from.Notifications.IsUnknown() && to.Notifications.IsNull() && len(from.Notifications.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notifications, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notifications = from.Notifications
	}
	if !from.RestartWindow.IsNull() && !from.RestartWindow.IsUnknown() {
		if toRestartWindow, ok := to.GetRestartWindow(ctx); ok {
			if fromRestartWindow, ok := from.GetRestartWindow(ctx); ok {
				toRestartWindow.SyncFieldsDuringRead(ctx, fromRestartWindow)
				to.SetRestartWindow(ctx, toRestartWindow)
			}
		}
	}
	if !from.Trigger.IsNull() && !from.Trigger.IsUnknown() {
		if toTrigger, ok := to.GetTrigger(ctx); ok {
			if fromTrigger, ok := from.GetTrigger(ctx); ok {
				toTrigger.SyncFieldsDuringRead(ctx, fromTrigger)
				to.SetTrigger(ctx, toTrigger)
			}
		}
	}
}

func (c PipelineSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["development"] = attrs["development"].SetOptional()
	attrs["edition"] = attrs["edition"].SetOptional()
	attrs["environment"] = attrs["environment"].SetOptional()
	attrs["event_log"] = attrs["event_log"].SetOptional()
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["gateway_definition"] = attrs["gateway_definition"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["ingestion_definition"] = attrs["ingestion_definition"].SetOptional()
	attrs["libraries"] = attrs["libraries"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["photon"] = attrs["photon"].SetOptional()
	attrs["restart_window"] = attrs["restart_window"].SetOptional()
	attrs["root_path"] = attrs["root_path"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()

	return attrs
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
		"environment":          reflect.TypeOf(PipelinesEnvironment{}),
		"event_log":            reflect.TypeOf(EventLogSpec{}),
		"filters":              reflect.TypeOf(Filters{}),
		"gateway_definition":   reflect.TypeOf(IngestionGatewayPipelineDefinition{}),
		"ingestion_definition": reflect.TypeOf(IngestionPipelineDefinition{}),
		"libraries":            reflect.TypeOf(PipelineLibrary{}),
		"notifications":        reflect.TypeOf(Notifications{}),
		"restart_window":       reflect.TypeOf(RestartWindow{}),
		"tags":                 reflect.TypeOf(types.String{}),
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
			"environment":          PipelinesEnvironment{}.Type(ctx),
			"event_log":            EventLogSpec{}.Type(ctx),
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
			"root_path":      types.StringType,
			"schema":         types.StringType,
			"serverless":     types.BoolType,
			"storage":        types.StringType,
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"target":  types.StringType,
			"trigger": PipelineTrigger{}.Type(ctx),
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
	var v PipelineDeployment
	d := o.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in PipelineSpec.
func (o *PipelineSpec) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	o.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in PipelineSpec as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if o.Environment.IsNull() || o.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := o.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in PipelineSpec.
func (o *PipelineSpec) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	o.Environment = vs
}

// GetEventLog returns the value of the EventLog field in PipelineSpec as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if o.EventLog.IsNull() || o.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := o.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in PipelineSpec.
func (o *PipelineSpec) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	o.EventLog = vs
}

// GetFilters returns the value of the Filters field in PipelineSpec as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if o.Filters.IsNull() || o.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := o.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionGatewayPipelineDefinition
	d := o.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v IngestionPipelineDefinition
	d := o.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v RestartWindow
	d := o.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in PipelineSpec.
func (o *PipelineSpec) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	o.RestartWindow = vs
}

// GetTags returns the value of the Tags field in PipelineSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in PipelineSpec.
func (o *PipelineSpec) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in PipelineSpec as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelineSpec) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := o.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in PipelineSpec.
func (o *PipelineSpec) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Trigger = vs
}

type PipelineStateInfo struct {
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

func (to *PipelineStateInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineStateInfo) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
}

func (to *PipelineStateInfo) SyncFieldsDuringRead(ctx context.Context, from PipelineStateInfo) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
}

func (c PipelineStateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	Cron types.Object `tfsdk:"cron"`

	Manual types.Object `tfsdk:"manual"`
}

func (to *PipelineTrigger) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineTrigger) {
	if !from.Cron.IsNull() && !from.Cron.IsUnknown() {
		if toCron, ok := to.GetCron(ctx); ok {
			if fromCron, ok := from.GetCron(ctx); ok {
				// Recursively sync the fields of Cron
				toCron.SyncFieldsDuringCreateOrUpdate(ctx, fromCron)
				to.SetCron(ctx, toCron)
			}
		}
	}
	if !from.Manual.IsNull() && !from.Manual.IsUnknown() {
		if toManual, ok := to.GetManual(ctx); ok {
			if fromManual, ok := from.GetManual(ctx); ok {
				// Recursively sync the fields of Manual
				toManual.SyncFieldsDuringCreateOrUpdate(ctx, fromManual)
				to.SetManual(ctx, toManual)
			}
		}
	}
}

func (to *PipelineTrigger) SyncFieldsDuringRead(ctx context.Context, from PipelineTrigger) {
	if !from.Cron.IsNull() && !from.Cron.IsUnknown() {
		if toCron, ok := to.GetCron(ctx); ok {
			if fromCron, ok := from.GetCron(ctx); ok {
				toCron.SyncFieldsDuringRead(ctx, fromCron)
				to.SetCron(ctx, toCron)
			}
		}
	}
	if !from.Manual.IsNull() && !from.Manual.IsUnknown() {
		if toManual, ok := to.GetManual(ctx); ok {
			if fromManual, ok := from.GetManual(ctx); ok {
				toManual.SyncFieldsDuringRead(ctx, fromManual)
				to.SetManual(ctx, toManual)
			}
		}
	}
}

func (c PipelineTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cron"] = attrs["cron"].SetOptional()
	attrs["manual"] = attrs["manual"].SetOptional()

	return attrs
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
	var v CronTrigger
	d := o.Cron.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v ManualTrigger
	d := o.Manual.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManual sets the value of the Manual field in PipelineTrigger.
func (o *PipelineTrigger) SetManual(ctx context.Context, v ManualTrigger) {
	vs := v.ToObjectValue(ctx)
	o.Manual = vs
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type PipelinesEnvironment struct {
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url>
	Dependencies types.List `tfsdk:"dependencies"`
}

func (to *PipelinesEnvironment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinesEnvironment) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (to *PipelinesEnvironment) SyncFieldsDuringRead(ctx context.Context, from PipelinesEnvironment) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (c PipelinesEnvironment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PipelinesEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinesEnvironment
// only implements ToObjectValue() and Type().
func (o PipelinesEnvironment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": o.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelinesEnvironment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in PipelinesEnvironment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PipelinesEnvironment) GetDependencies(ctx context.Context) ([]types.String, bool) {
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

// SetDependencies sets the value of the Dependencies field in PipelinesEnvironment.
func (o *PipelinesEnvironment) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// PG-specific catalog-level configuration parameters
type PostgresCatalogConfig struct {
	// Optional. The Postgres slot configuration to use for logical replication
	SlotConfig types.Object `tfsdk:"slot_config"`
}

func (to *PostgresCatalogConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PostgresCatalogConfig) {
	if !from.SlotConfig.IsNull() && !from.SlotConfig.IsUnknown() {
		if toSlotConfig, ok := to.GetSlotConfig(ctx); ok {
			if fromSlotConfig, ok := from.GetSlotConfig(ctx); ok {
				// Recursively sync the fields of SlotConfig
				toSlotConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromSlotConfig)
				to.SetSlotConfig(ctx, toSlotConfig)
			}
		}
	}
}

func (to *PostgresCatalogConfig) SyncFieldsDuringRead(ctx context.Context, from PostgresCatalogConfig) {
	if !from.SlotConfig.IsNull() && !from.SlotConfig.IsUnknown() {
		if toSlotConfig, ok := to.GetSlotConfig(ctx); ok {
			if fromSlotConfig, ok := from.GetSlotConfig(ctx); ok {
				toSlotConfig.SyncFieldsDuringRead(ctx, fromSlotConfig)
				to.SetSlotConfig(ctx, toSlotConfig)
			}
		}
	}
}

func (c PostgresCatalogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["slot_config"] = attrs["slot_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PostgresCatalogConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PostgresCatalogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"slot_config": reflect.TypeOf(PostgresSlotConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresCatalogConfig
// only implements ToObjectValue() and Type().
func (o PostgresCatalogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"slot_config": o.SlotConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PostgresCatalogConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"slot_config": PostgresSlotConfig{}.Type(ctx),
		},
	}
}

// GetSlotConfig returns the value of the SlotConfig field in PostgresCatalogConfig as
// a PostgresSlotConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PostgresCatalogConfig) GetSlotConfig(ctx context.Context) (PostgresSlotConfig, bool) {
	var e PostgresSlotConfig
	if o.SlotConfig.IsNull() || o.SlotConfig.IsUnknown() {
		return e, false
	}
	var v PostgresSlotConfig
	d := o.SlotConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlotConfig sets the value of the SlotConfig field in PostgresCatalogConfig.
func (o *PostgresCatalogConfig) SetSlotConfig(ctx context.Context, v PostgresSlotConfig) {
	vs := v.ToObjectValue(ctx)
	o.SlotConfig = vs
}

// PostgresSlotConfig contains the configuration for a Postgres logical
// replication slot
type PostgresSlotConfig struct {
	// The name of the publication to use for the Postgres source
	PublicationName types.String `tfsdk:"publication_name"`
	// The name of the logical replication slot to use for the Postgres source
	SlotName types.String `tfsdk:"slot_name"`
}

func (to *PostgresSlotConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PostgresSlotConfig) {
}

func (to *PostgresSlotConfig) SyncFieldsDuringRead(ctx context.Context, from PostgresSlotConfig) {
}

func (c PostgresSlotConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PostgresSlotConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresSlotConfig
// only implements ToObjectValue() and Type().
func (o PostgresSlotConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publication_name": o.PublicationName,
			"slot_name":        o.SlotName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PostgresSlotConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"publication_name": types.StringType,
			"slot_name":        types.StringType,
		},
	}
}

type ReportSpec struct {
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
	TableConfiguration types.Object `tfsdk:"table_configuration"`
}

func (to *ReportSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReportSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				// Recursively sync the fields of TableConfiguration
				toTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (to *ReportSpec) SyncFieldsDuringRead(ctx context.Context, from ReportSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (c ReportSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["destination_table"] = attrs["destination_table"].SetOptional()
	attrs["source_url"] = attrs["source_url"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()

	return attrs
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
	var v TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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

func (to *RestartWindow) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestartWindow) {
	if !from.DaysOfWeek.IsNull() && !from.DaysOfWeek.IsUnknown() && to.DaysOfWeek.IsNull() && len(from.DaysOfWeek.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DaysOfWeek, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DaysOfWeek = from.DaysOfWeek
	}
}

func (to *RestartWindow) SyncFieldsDuringRead(ctx context.Context, from RestartWindow) {
	if !from.DaysOfWeek.IsNull() && !from.DaysOfWeek.IsUnknown() && to.DaysOfWeek.IsNull() && len(from.DaysOfWeek.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DaysOfWeek, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DaysOfWeek = from.DaysOfWeek
	}
}

func (c RestartWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

// Write-only setting, available only in Create/Update calls. Specifies the user
// or service principal that the pipeline runs as. If not specified, the
// pipeline runs as the user who created the pipeline.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type RunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName types.String `tfsdk:"user_name"`
}

func (to *RunAs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunAs) {
}

func (to *RunAs) SyncFieldsDuringRead(ctx context.Context, from RunAs) {
}

func (c RunAs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunAs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunAs
// only implements ToObjectValue() and Type().
func (o RunAs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunAs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type SchemaSpec struct {
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
	TableConfiguration types.Object `tfsdk:"table_configuration"`
}

func (to *SchemaSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SchemaSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				// Recursively sync the fields of TableConfiguration
				toTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (to *SchemaSpec) SyncFieldsDuringRead(ctx context.Context, from SchemaSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (c SchemaSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()
	attrs["source_schema"] = attrs["source_schema"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()

	return attrs
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
	var v TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in SchemaSpec.
func (o *SchemaSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

type Sequencing struct {
	// A sequence number, unique and increasing per pipeline.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no"`
	// the ID assigned by the data plane.
	DataPlaneId types.Object `tfsdk:"data_plane_id"`
}

func (to *Sequencing) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Sequencing) {
	if !from.DataPlaneId.IsNull() && !from.DataPlaneId.IsUnknown() {
		if toDataPlaneId, ok := to.GetDataPlaneId(ctx); ok {
			if fromDataPlaneId, ok := from.GetDataPlaneId(ctx); ok {
				// Recursively sync the fields of DataPlaneId
				toDataPlaneId.SyncFieldsDuringCreateOrUpdate(ctx, fromDataPlaneId)
				to.SetDataPlaneId(ctx, toDataPlaneId)
			}
		}
	}
}

func (to *Sequencing) SyncFieldsDuringRead(ctx context.Context, from Sequencing) {
	if !from.DataPlaneId.IsNull() && !from.DataPlaneId.IsUnknown() {
		if toDataPlaneId, ok := to.GetDataPlaneId(ctx); ok {
			if fromDataPlaneId, ok := from.GetDataPlaneId(ctx); ok {
				toDataPlaneId.SyncFieldsDuringRead(ctx, fromDataPlaneId)
				to.SetDataPlaneId(ctx, toDataPlaneId)
			}
		}
	}
}

func (c Sequencing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["control_plane_seq_no"] = attrs["control_plane_seq_no"].SetOptional()
	attrs["data_plane_id"] = attrs["data_plane_id"].SetOptional()

	return attrs
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
	var v DataPlaneId
	d := o.DataPlaneId.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataPlaneId sets the value of the DataPlaneId field in Sequencing.
func (o *Sequencing) SetDataPlaneId(ctx context.Context, v DataPlaneId) {
	vs := v.ToObjectValue(ctx)
	o.DataPlaneId = vs
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name"`
	// Exception message
	Message types.String `tfsdk:"message"`
	// Stack trace consisting of a list of stack frames
	Stack types.List `tfsdk:"stack"`
}

func (to *SerializedException) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SerializedException) {
	if !from.Stack.IsNull() && !from.Stack.IsUnknown() && to.Stack.IsNull() && len(from.Stack.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stack, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stack = from.Stack
	}
}

func (to *SerializedException) SyncFieldsDuringRead(ctx context.Context, from SerializedException) {
	if !from.Stack.IsNull() && !from.Stack.IsUnknown() && to.Stack.IsNull() && len(from.Stack.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stack, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stack = from.Stack
	}
}

func (c SerializedException) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

// SourceCatalogConfig contains catalog-level custom configuration parameters
// for each source
type SourceCatalogConfig struct {
	// Postgres-specific catalog-level configuration parameters
	Postgres types.Object `tfsdk:"postgres"`
	// Source catalog name
	SourceCatalog types.String `tfsdk:"source_catalog"`
}

func (to *SourceCatalogConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SourceCatalogConfig) {
	if !from.Postgres.IsNull() && !from.Postgres.IsUnknown() {
		if toPostgres, ok := to.GetPostgres(ctx); ok {
			if fromPostgres, ok := from.GetPostgres(ctx); ok {
				// Recursively sync the fields of Postgres
				toPostgres.SyncFieldsDuringCreateOrUpdate(ctx, fromPostgres)
				to.SetPostgres(ctx, toPostgres)
			}
		}
	}
}

func (to *SourceCatalogConfig) SyncFieldsDuringRead(ctx context.Context, from SourceCatalogConfig) {
	if !from.Postgres.IsNull() && !from.Postgres.IsUnknown() {
		if toPostgres, ok := to.GetPostgres(ctx); ok {
			if fromPostgres, ok := from.GetPostgres(ctx); ok {
				toPostgres.SyncFieldsDuringRead(ctx, fromPostgres)
				to.SetPostgres(ctx, toPostgres)
			}
		}
	}
}

func (c SourceCatalogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["postgres"] = attrs["postgres"].SetOptional()
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
func (a SourceCatalogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"postgres": reflect.TypeOf(PostgresCatalogConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceCatalogConfig
// only implements ToObjectValue() and Type().
func (o SourceCatalogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"postgres":       o.Postgres,
			"source_catalog": o.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SourceCatalogConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"postgres":       PostgresCatalogConfig{}.Type(ctx),
			"source_catalog": types.StringType,
		},
	}
}

// GetPostgres returns the value of the Postgres field in SourceCatalogConfig as
// a PostgresCatalogConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *SourceCatalogConfig) GetPostgres(ctx context.Context) (PostgresCatalogConfig, bool) {
	var e PostgresCatalogConfig
	if o.Postgres.IsNull() || o.Postgres.IsUnknown() {
		return e, false
	}
	var v PostgresCatalogConfig
	d := o.Postgres.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPostgres sets the value of the Postgres field in SourceCatalogConfig.
func (o *SourceCatalogConfig) SetPostgres(ctx context.Context, v PostgresCatalogConfig) {
	vs := v.ToObjectValue(ctx)
	o.Postgres = vs
}

type SourceConfig struct {
	// Catalog-level source configuration parameters
	Catalog types.Object `tfsdk:"catalog"`
}

func (to *SourceConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SourceConfig) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				// Recursively sync the fields of Catalog
				toCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (to *SourceConfig) SyncFieldsDuringRead(ctx context.Context, from SourceConfig) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				toCatalog.SyncFieldsDuringRead(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (c SourceConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SourceConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SourceConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(SourceCatalogConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceConfig
// only implements ToObjectValue() and Type().
func (o SourceConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SourceConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": SourceCatalogConfig{}.Type(ctx),
		},
	}
}

// GetCatalog returns the value of the Catalog field in SourceConfig as
// a SourceCatalogConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *SourceConfig) GetCatalog(ctx context.Context) (SourceCatalogConfig, bool) {
	var e SourceCatalogConfig
	if o.Catalog.IsNull() || o.Catalog.IsUnknown() {
		return e, false
	}
	var v SourceCatalogConfig
	d := o.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalog sets the value of the Catalog field in SourceConfig.
func (o *SourceConfig) SetCatalog(ctx context.Context, v SourceCatalogConfig) {
	vs := v.ToObjectValue(ctx)
	o.Catalog = vs
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass types.String `tfsdk:"declaring_class"`
	// File where the method is defined
	FileName types.String `tfsdk:"file_name"`
	// Line from which the method was called
	LineNumber types.Int64 `tfsdk:"line_number"`
	// Name of the method which was called
	MethodName types.String `tfsdk:"method_name"`
}

func (to *StackFrame) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StackFrame) {
}

func (to *StackFrame) SyncFieldsDuringRead(ctx context.Context, from StackFrame) {
}

func (c StackFrame) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *StartUpdate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartUpdate) {
	if !from.FullRefreshSelection.IsNull() && !from.FullRefreshSelection.IsUnknown() && to.FullRefreshSelection.IsNull() && len(from.FullRefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FullRefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FullRefreshSelection = from.FullRefreshSelection
	}
	if !from.RefreshSelection.IsNull() && !from.RefreshSelection.IsUnknown() && to.RefreshSelection.IsNull() && len(from.RefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RefreshSelection = from.RefreshSelection
	}
}

func (to *StartUpdate) SyncFieldsDuringRead(ctx context.Context, from StartUpdate) {
	if !from.FullRefreshSelection.IsNull() && !from.FullRefreshSelection.IsUnknown() && to.FullRefreshSelection.IsNull() && len(from.FullRefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FullRefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FullRefreshSelection = from.FullRefreshSelection
	}
	if !from.RefreshSelection.IsNull() && !from.RefreshSelection.IsUnknown() && to.RefreshSelection.IsNull() && len(from.RefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RefreshSelection = from.RefreshSelection
	}
}

func (c StartUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["full_refresh"] = attrs["full_refresh"].SetOptional()
	attrs["full_refresh_selection"] = attrs["full_refresh_selection"].SetOptional()
	attrs["refresh_selection"] = attrs["refresh_selection"].SetOptional()
	attrs["validate_only"] = attrs["validate_only"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
	UpdateId types.String `tfsdk:"update_id"`
}

func (to *StartUpdateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartUpdateResponse) {
}

func (to *StartUpdateResponse) SyncFieldsDuringRead(ctx context.Context, from StartUpdateResponse) {
}

func (c StartUpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *StopPipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopPipelineResponse) {
}

func (to *StopPipelineResponse) SyncFieldsDuringRead(ctx context.Context, from StopPipelineResponse) {
}

func (c StopPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type StopRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *StopRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopRequest) {
}

func (to *StopRequest) SyncFieldsDuringRead(ctx context.Context, from StopRequest) {
}

func (c StopRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
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
	TableConfiguration types.Object `tfsdk:"table_configuration"`
}

func (to *TableSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				// Recursively sync the fields of TableConfiguration
				toTableConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (to *TableSpec) SyncFieldsDuringRead(ctx context.Context, from TableSpec) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (c TableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_catalog"] = attrs["destination_catalog"].SetRequired()
	attrs["destination_schema"] = attrs["destination_schema"].SetRequired()
	attrs["destination_table"] = attrs["destination_table"].SetOptional()
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()
	attrs["source_schema"] = attrs["source_schema"].SetOptional()
	attrs["source_table"] = attrs["source_table"].SetRequired()
	attrs["table_configuration"] = attrs["table_configuration"].SetOptional()

	return attrs
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
	var v TableSpecificConfig
	d := o.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in TableSpec.
func (o *TableSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	o.TableConfiguration = vs
}

type TableSpecificConfig struct {
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

	QueryBasedConnectorConfig types.Object `tfsdk:"query_based_connector_config"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields types.Bool `tfsdk:"salesforce_include_formula_fields"`
	// The SCD type to use to ingest the table.
	ScdType types.String `tfsdk:"scd_type"`
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	SequenceBy types.List `tfsdk:"sequence_by"`
	// (Optional) Additional custom parameters for Workday Report
	WorkdayReportParameters types.Object `tfsdk:"workday_report_parameters"`
}

func (to *TableSpecificConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableSpecificConfig) {
	if !from.ExcludeColumns.IsNull() && !from.ExcludeColumns.IsUnknown() && to.ExcludeColumns.IsNull() && len(from.ExcludeColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExcludeColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExcludeColumns = from.ExcludeColumns
	}
	if !from.IncludeColumns.IsNull() && !from.IncludeColumns.IsUnknown() && to.IncludeColumns.IsNull() && len(from.IncludeColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IncludeColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IncludeColumns = from.IncludeColumns
	}
	if !from.PrimaryKeys.IsNull() && !from.PrimaryKeys.IsUnknown() && to.PrimaryKeys.IsNull() && len(from.PrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrimaryKeys = from.PrimaryKeys
	}
	if !from.QueryBasedConnectorConfig.IsNull() && !from.QueryBasedConnectorConfig.IsUnknown() {
		if toQueryBasedConnectorConfig, ok := to.GetQueryBasedConnectorConfig(ctx); ok {
			if fromQueryBasedConnectorConfig, ok := from.GetQueryBasedConnectorConfig(ctx); ok {
				// Recursively sync the fields of QueryBasedConnectorConfig
				toQueryBasedConnectorConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromQueryBasedConnectorConfig)
				to.SetQueryBasedConnectorConfig(ctx, toQueryBasedConnectorConfig)
			}
		}
	}
	if !from.SequenceBy.IsNull() && !from.SequenceBy.IsUnknown() && to.SequenceBy.IsNull() && len(from.SequenceBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SequenceBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SequenceBy = from.SequenceBy
	}
	if !from.WorkdayReportParameters.IsNull() && !from.WorkdayReportParameters.IsUnknown() {
		if toWorkdayReportParameters, ok := to.GetWorkdayReportParameters(ctx); ok {
			if fromWorkdayReportParameters, ok := from.GetWorkdayReportParameters(ctx); ok {
				// Recursively sync the fields of WorkdayReportParameters
				toWorkdayReportParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkdayReportParameters)
				to.SetWorkdayReportParameters(ctx, toWorkdayReportParameters)
			}
		}
	}
}

func (to *TableSpecificConfig) SyncFieldsDuringRead(ctx context.Context, from TableSpecificConfig) {
	if !from.ExcludeColumns.IsNull() && !from.ExcludeColumns.IsUnknown() && to.ExcludeColumns.IsNull() && len(from.ExcludeColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExcludeColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExcludeColumns = from.ExcludeColumns
	}
	if !from.IncludeColumns.IsNull() && !from.IncludeColumns.IsUnknown() && to.IncludeColumns.IsNull() && len(from.IncludeColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IncludeColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IncludeColumns = from.IncludeColumns
	}
	if !from.PrimaryKeys.IsNull() && !from.PrimaryKeys.IsUnknown() && to.PrimaryKeys.IsNull() && len(from.PrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrimaryKeys = from.PrimaryKeys
	}
	if !from.QueryBasedConnectorConfig.IsNull() && !from.QueryBasedConnectorConfig.IsUnknown() {
		if toQueryBasedConnectorConfig, ok := to.GetQueryBasedConnectorConfig(ctx); ok {
			if fromQueryBasedConnectorConfig, ok := from.GetQueryBasedConnectorConfig(ctx); ok {
				toQueryBasedConnectorConfig.SyncFieldsDuringRead(ctx, fromQueryBasedConnectorConfig)
				to.SetQueryBasedConnectorConfig(ctx, toQueryBasedConnectorConfig)
			}
		}
	}
	if !from.SequenceBy.IsNull() && !from.SequenceBy.IsUnknown() && to.SequenceBy.IsNull() && len(from.SequenceBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SequenceBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SequenceBy = from.SequenceBy
	}
	if !from.WorkdayReportParameters.IsNull() && !from.WorkdayReportParameters.IsUnknown() {
		if toWorkdayReportParameters, ok := to.GetWorkdayReportParameters(ctx); ok {
			if fromWorkdayReportParameters, ok := from.GetWorkdayReportParameters(ctx); ok {
				toWorkdayReportParameters.SyncFieldsDuringRead(ctx, fromWorkdayReportParameters)
				to.SetWorkdayReportParameters(ctx, toWorkdayReportParameters)
			}
		}
	}
}

func (c TableSpecificConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exclude_columns"] = attrs["exclude_columns"].SetOptional()
	attrs["include_columns"] = attrs["include_columns"].SetOptional()
	attrs["primary_keys"] = attrs["primary_keys"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].SetOptional()
	attrs["salesforce_include_formula_fields"] = attrs["salesforce_include_formula_fields"].SetOptional()
	attrs["scd_type"] = attrs["scd_type"].SetOptional()
	attrs["sequence_by"] = attrs["sequence_by"].SetOptional()
	attrs["workday_report_parameters"] = attrs["workday_report_parameters"].SetOptional()

	return attrs
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
		"exclude_columns":              reflect.TypeOf(types.String{}),
		"include_columns":              reflect.TypeOf(types.String{}),
		"primary_keys":                 reflect.TypeOf(types.String{}),
		"query_based_connector_config": reflect.TypeOf(IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig{}),
		"sequence_by":                  reflect.TypeOf(types.String{}),
		"workday_report_parameters":    reflect.TypeOf(IngestionPipelineDefinitionWorkdayReportParameters{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpecificConfig
// only implements ToObjectValue() and Type().
func (o TableSpecificConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"workday_report_parameters":         o.WorkdayReportParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSpecificConfig) Type(ctx context.Context) attr.Type {
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
			"query_based_connector_config":      IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig{}.Type(ctx),
			"salesforce_include_formula_fields": types.BoolType,
			"scd_type":                          types.StringType,
			"sequence_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"workday_report_parameters": IngestionPipelineDefinitionWorkdayReportParameters{}.Type(ctx),
		},
	}
}

// GetExcludeColumns returns the value of the ExcludeColumns field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetExcludeColumns(ctx context.Context) ([]types.String, bool) {
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

// SetExcludeColumns sets the value of the ExcludeColumns field in TableSpecificConfig.
func (o *TableSpecificConfig) SetExcludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExcludeColumns = types.ListValueMust(t, vs)
}

// GetIncludeColumns returns the value of the IncludeColumns field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetIncludeColumns(ctx context.Context) ([]types.String, bool) {
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

// SetIncludeColumns sets the value of the IncludeColumns field in TableSpecificConfig.
func (o *TableSpecificConfig) SetIncludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IncludeColumns = types.ListValueMust(t, vs)
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

// GetQueryBasedConnectorConfig returns the value of the QueryBasedConnectorConfig field in TableSpecificConfig as
// a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetQueryBasedConnectorConfig(ctx context.Context) (IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig, bool) {
	var e IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
	if o.QueryBasedConnectorConfig.IsNull() || o.QueryBasedConnectorConfig.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
	d := o.QueryBasedConnectorConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryBasedConnectorConfig sets the value of the QueryBasedConnectorConfig field in TableSpecificConfig.
func (o *TableSpecificConfig) SetQueryBasedConnectorConfig(ctx context.Context, v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) {
	vs := v.ToObjectValue(ctx)
	o.QueryBasedConnectorConfig = vs
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

// GetWorkdayReportParameters returns the value of the WorkdayReportParameters field in TableSpecificConfig as
// a IngestionPipelineDefinitionWorkdayReportParameters value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableSpecificConfig) GetWorkdayReportParameters(ctx context.Context) (IngestionPipelineDefinitionWorkdayReportParameters, bool) {
	var e IngestionPipelineDefinitionWorkdayReportParameters
	if o.WorkdayReportParameters.IsNull() || o.WorkdayReportParameters.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinitionWorkdayReportParameters
	d := o.WorkdayReportParameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkdayReportParameters sets the value of the WorkdayReportParameters field in TableSpecificConfig.
func (o *TableSpecificConfig) SetWorkdayReportParameters(ctx context.Context, v IngestionPipelineDefinitionWorkdayReportParameters) {
	vs := v.ToObjectValue(ctx)
	o.WorkdayReportParameters = vs
}

type UpdateInfo struct {
	// What triggered this update.
	Cause types.String `tfsdk:"cause"`
	// The ID of the cluster that the update is running on.
	ClusterId types.String `tfsdk:"cluster_id"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config types.Object `tfsdk:"config"`
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

func (to *UpdateInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInfo) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.FullRefreshSelection.IsNull() && !from.FullRefreshSelection.IsUnknown() && to.FullRefreshSelection.IsNull() && len(from.FullRefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FullRefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FullRefreshSelection = from.FullRefreshSelection
	}
	if !from.RefreshSelection.IsNull() && !from.RefreshSelection.IsUnknown() && to.RefreshSelection.IsNull() && len(from.RefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RefreshSelection = from.RefreshSelection
	}
}

func (to *UpdateInfo) SyncFieldsDuringRead(ctx context.Context, from UpdateInfo) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.FullRefreshSelection.IsNull() && !from.FullRefreshSelection.IsUnknown() && to.FullRefreshSelection.IsNull() && len(from.FullRefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FullRefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FullRefreshSelection = from.FullRefreshSelection
	}
	if !from.RefreshSelection.IsNull() && !from.RefreshSelection.IsUnknown() && to.RefreshSelection.IsNull() && len(from.RefreshSelection.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RefreshSelection, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RefreshSelection = from.RefreshSelection
	}
}

func (c UpdateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
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
	var v PipelineSpec
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	CreationTime types.String `tfsdk:"creation_time"`

	State types.String `tfsdk:"state"`

	UpdateId types.String `tfsdk:"update_id"`
}

func (to *UpdateStateInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateStateInfo) {
}

func (to *UpdateStateInfo) SyncFieldsDuringRead(ctx context.Context, from UpdateStateInfo) {
}

func (c UpdateStateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
