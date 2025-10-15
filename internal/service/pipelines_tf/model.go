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

type ApplyEnvironmentRequest struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *ApplyEnvironmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApplyEnvironmentRequest) {
}

func (to *ApplyEnvironmentRequest) SyncFieldsDuringRead(ctx context.Context, from ApplyEnvironmentRequest) {
}

func (m ApplyEnvironmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApplyEnvironmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApplyEnvironmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplyEnvironmentRequest
// only implements ToObjectValue() and Type().
func (m ApplyEnvironmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApplyEnvironmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type ApplyEnvironmentRequestResponse struct {
}

func (to *ApplyEnvironmentRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApplyEnvironmentRequestResponse) {
}

func (to *ApplyEnvironmentRequestResponse) SyncFieldsDuringRead(ctx context.Context, from ApplyEnvironmentRequestResponse) {
}

func (m ApplyEnvironmentRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApplyEnvironmentRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApplyEnvironmentRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplyEnvironmentRequestResponse
// only implements ToObjectValue() and Type().
func (m ApplyEnvironmentRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ApplyEnvironmentRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ConnectionParameters struct {
	// Source catalog for initial connection. This is necessary for schema
	// exploration in some database systems like Oracle, and optional but
	// nice-to-have in some other database systems like Postgres. For Oracle
	// databases, this maps to a service name.
	SourceCatalog types.String `tfsdk:"source_catalog"`
}

func (to *ConnectionParameters) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConnectionParameters) {
}

func (to *ConnectionParameters) SyncFieldsDuringRead(ctx context.Context, from ConnectionParameters) {
}

func (m ConnectionParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["source_catalog"] = attrs["source_catalog"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConnectionParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ConnectionParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConnectionParameters
// only implements ToObjectValue() and Type().
func (m ConnectionParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_catalog": m.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConnectionParameters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_catalog": types.StringType,
		},
	}
}

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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
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

func (m CreatePipeline) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePipeline.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePipeline) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CreatePipeline) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_duplicate_names": m.AllowDuplicateNames,
			"budget_policy_id":      m.BudgetPolicyId,
			"catalog":               m.Catalog,
			"channel":               m.Channel,
			"clusters":              m.Clusters,
			"configuration":         m.Configuration,
			"continuous":            m.Continuous,
			"deployment":            m.Deployment,
			"development":           m.Development,
			"dry_run":               m.DryRun,
			"edition":               m.Edition,
			"environment":           m.Environment,
			"event_log":             m.EventLog,
			"filters":               m.Filters,
			"gateway_definition":    m.GatewayDefinition,
			"id":                    m.Id,
			"ingestion_definition":  m.IngestionDefinition,
			"libraries":             m.Libraries,
			"name":                  m.Name,
			"notifications":         m.Notifications,
			"photon":                m.Photon,
			"restart_window":        m.RestartWindow,
			"root_path":             m.RootPath,
			"run_as":                m.RunAs,
			"schema":                m.Schema,
			"serverless":            m.Serverless,
			"storage":               m.Storage,
			"tags":                  m.Tags,
			"target":                m.Target,
			"trigger":               m.Trigger,
			"usage_policy_id":       m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePipeline) Type(ctx context.Context) attr.Type {
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
			"target":          types.StringType,
			"trigger":         PipelineTrigger{}.Type(ctx),
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in CreatePipeline as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in CreatePipeline.
func (m *CreatePipeline) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in CreatePipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if m.Configuration.IsNull() || m.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in CreatePipeline.
func (m *CreatePipeline) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in CreatePipeline as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v PipelineDeployment
	d := m.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in CreatePipeline.
func (m *CreatePipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	m.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in CreatePipeline as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := m.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in CreatePipeline.
func (m *CreatePipeline) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.Environment = vs
}

// GetEventLog returns the value of the EventLog field in CreatePipeline as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := m.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in CreatePipeline.
func (m *CreatePipeline) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	m.EventLog = vs
}

// GetFilters returns the value of the Filters field in CreatePipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := m.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in CreatePipeline.
func (m *CreatePipeline) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	m.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in CreatePipeline as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionGatewayPipelineDefinition
	d := m.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in CreatePipeline.
func (m *CreatePipeline) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in CreatePipeline as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinition
	d := m.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in CreatePipeline.
func (m *CreatePipeline) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in CreatePipeline as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePipeline.
func (m *CreatePipeline) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in CreatePipeline as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in CreatePipeline.
func (m *CreatePipeline) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in CreatePipeline as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v RestartWindow
	d := m.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in CreatePipeline.
func (m *CreatePipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	m.RestartWindow = vs
}

// GetRunAs returns the value of the RunAs field in CreatePipeline as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := m.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in CreatePipeline.
func (m *CreatePipeline) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	m.RunAs = vs
}

// GetTags returns the value of the Tags field in CreatePipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreatePipeline.
func (m *CreatePipeline) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in CreatePipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := m.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in CreatePipeline.
func (m *CreatePipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	m.Trigger = vs
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

func (m CreatePipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(PipelineSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipelineResponse
// only implements ToObjectValue() and Type().
func (m CreatePipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_settings": m.EffectiveSettings,
			"pipeline_id":        m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePipelineResponse) Type(ctx context.Context) attr.Type {
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
func (m *CreatePipelineResponse) GetEffectiveSettings(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v PipelineSpec
	d := m.EffectiveSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in CreatePipelineResponse.
func (m *CreatePipelineResponse) SetEffectiveSettings(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveSettings = vs
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`

	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (to *CronTrigger) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronTrigger) {
}

func (to *CronTrigger) SyncFieldsDuringRead(ctx context.Context, from CronTrigger) {
}

func (m CronTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CronTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronTrigger
// only implements ToObjectValue() and Type().
func (m CronTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_schedule": m.QuartzCronSchedule,
			"timezone_id":          m.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CronTrigger) Type(ctx context.Context) attr.Type {
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

func (m DataPlaneId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataPlaneId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneId
// only implements ToObjectValue() and Type().
func (m DataPlaneId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance": m.Instance,
			"seq_no":   m.SeqNo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataPlaneId) Type(ctx context.Context) attr.Type {
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

func (m DeletePipelineRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePipelineRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineRequest
// only implements ToObjectValue() and Type().
func (m DeletePipelineRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePipelineRequest) Type(ctx context.Context) attr.Type {
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

func (m DeletePipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineResponse
// only implements ToObjectValue() and Type().
func (m DeletePipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePipelineResponse) Type(ctx context.Context) attr.Type {
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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
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

func (m EditPipeline) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
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
func (m EditPipeline) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EditPipeline) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_duplicate_names":  m.AllowDuplicateNames,
			"budget_policy_id":       m.BudgetPolicyId,
			"catalog":                m.Catalog,
			"channel":                m.Channel,
			"clusters":               m.Clusters,
			"configuration":          m.Configuration,
			"continuous":             m.Continuous,
			"deployment":             m.Deployment,
			"development":            m.Development,
			"edition":                m.Edition,
			"environment":            m.Environment,
			"event_log":              m.EventLog,
			"expected_last_modified": m.ExpectedLastModified,
			"filters":                m.Filters,
			"gateway_definition":     m.GatewayDefinition,
			"id":                     m.Id,
			"ingestion_definition":   m.IngestionDefinition,
			"libraries":              m.Libraries,
			"name":                   m.Name,
			"notifications":          m.Notifications,
			"photon":                 m.Photon,
			"pipeline_id":            m.PipelineId,
			"restart_window":         m.RestartWindow,
			"root_path":              m.RootPath,
			"run_as":                 m.RunAs,
			"schema":                 m.Schema,
			"serverless":             m.Serverless,
			"storage":                m.Storage,
			"tags":                   m.Tags,
			"target":                 m.Target,
			"trigger":                m.Trigger,
			"usage_policy_id":        m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EditPipeline) Type(ctx context.Context) attr.Type {
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
			"target":          types.StringType,
			"trigger":         PipelineTrigger{}.Type(ctx),
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in EditPipeline as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in EditPipeline.
func (m *EditPipeline) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in EditPipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if m.Configuration.IsNull() || m.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in EditPipeline.
func (m *EditPipeline) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in EditPipeline as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v PipelineDeployment
	d := m.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in EditPipeline.
func (m *EditPipeline) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	m.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in EditPipeline as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := m.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in EditPipeline.
func (m *EditPipeline) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.Environment = vs
}

// GetEventLog returns the value of the EventLog field in EditPipeline as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := m.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in EditPipeline.
func (m *EditPipeline) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	m.EventLog = vs
}

// GetFilters returns the value of the Filters field in EditPipeline as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := m.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in EditPipeline.
func (m *EditPipeline) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	m.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in EditPipeline as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionGatewayPipelineDefinition
	d := m.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in EditPipeline.
func (m *EditPipeline) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in EditPipeline as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinition
	d := m.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in EditPipeline.
func (m *EditPipeline) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in EditPipeline as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPipeline.
func (m *EditPipeline) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in EditPipeline as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in EditPipeline.
func (m *EditPipeline) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in EditPipeline as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v RestartWindow
	d := m.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in EditPipeline.
func (m *EditPipeline) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	m.RestartWindow = vs
}

// GetRunAs returns the value of the RunAs field in EditPipeline as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := m.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in EditPipeline.
func (m *EditPipeline) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	m.RunAs = vs
}

// GetTags returns the value of the Tags field in EditPipeline as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EditPipeline.
func (m *EditPipeline) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in EditPipeline as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := m.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in EditPipeline.
func (m *EditPipeline) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	m.Trigger = vs
}

type EditPipelineResponse struct {
}

func (to *EditPipelineResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPipelineResponse) {
}

func (to *EditPipelineResponse) SyncFieldsDuringRead(ctx context.Context, from EditPipelineResponse) {
}

func (m EditPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipelineResponse
// only implements ToObjectValue() and Type().
func (m EditPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditPipelineResponse) Type(ctx context.Context) attr.Type {
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

func (m ErrorDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ErrorDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exceptions": reflect.TypeOf(SerializedException{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ErrorDetail
// only implements ToObjectValue() and Type().
func (m ErrorDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exceptions": m.Exceptions,
			"fatal":      m.Fatal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ErrorDetail) Type(ctx context.Context) attr.Type {
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
func (m *ErrorDetail) GetExceptions(ctx context.Context) ([]SerializedException, bool) {
	if m.Exceptions.IsNull() || m.Exceptions.IsUnknown() {
		return nil, false
	}
	var v []SerializedException
	d := m.Exceptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExceptions sets the value of the Exceptions field in ErrorDetail.
func (m *ErrorDetail) SetExceptions(ctx context.Context, v []SerializedException) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exceptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exceptions = types.ListValueMust(t, vs)
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

func (m EventLogSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EventLogSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventLogSpec
// only implements ToObjectValue() and Type().
func (m EventLogSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
			"name":    m.Name,
			"schema":  m.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EventLogSpec) Type(ctx context.Context) attr.Type {
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

func (m FileLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FileLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileLibrary
// only implements ToObjectValue() and Type().
func (m FileLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileLibrary) Type(ctx context.Context) attr.Type {
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

func (m Filters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Filters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude": reflect.TypeOf(types.String{}),
		"include": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filters
// only implements ToObjectValue() and Type().
func (m Filters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude": m.Exclude,
			"include": m.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Filters) Type(ctx context.Context) attr.Type {
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
func (m *Filters) GetExclude(ctx context.Context) ([]types.String, bool) {
	if m.Exclude.IsNull() || m.Exclude.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Exclude.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExclude sets the value of the Exclude field in Filters.
func (m *Filters) SetExclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exclude = types.ListValueMust(t, vs)
}

// GetInclude returns the value of the Include field in Filters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Filters) GetInclude(ctx context.Context) ([]types.String, bool) {
	if m.Include.IsNull() || m.Include.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Include.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInclude sets the value of the Include field in Filters.
func (m *Filters) SetInclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Include = types.ListValueMust(t, vs)
}

type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionLevelsRequest) {
}

func (to *GetPipelinePermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionLevelsRequest) {
}

func (m GetPipelinePermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (m GetPipelinePermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PipelinePermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetPipelinePermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]PipelinePermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse.
func (m *GetPipelinePermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []PipelinePermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionsRequest) {
}

func (to *GetPipelinePermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionsRequest) {
}

func (m GetPipelinePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionsRequest) Type(ctx context.Context) attr.Type {
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

func (m GetPipelineRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelineRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineRequest
// only implements ToObjectValue() and Type().
func (m GetPipelineRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelineRequest) Type(ctx context.Context) attr.Type {
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
	// Serverless usage policy ID of the pipeline.
	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
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

func (m GetPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["creator_user_name"] = attrs["creator_user_name"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetOptional()
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
func (m GetPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
		"run_as":         reflect.TypeOf(RunAs{}),
		"spec":           reflect.TypeOf(PipelineSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineResponse
// only implements ToObjectValue() and Type().
func (m GetPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                      m.Cause,
			"cluster_id":                 m.ClusterId,
			"creator_user_name":          m.CreatorUserName,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"effective_usage_policy_id":  m.EffectiveUsagePolicyId,
			"health":                     m.Health,
			"last_modified":              m.LastModified,
			"latest_updates":             m.LatestUpdates,
			"name":                       m.Name,
			"pipeline_id":                m.PipelineId,
			"run_as":                     m.RunAs,
			"run_as_user_name":           m.RunAsUserName,
			"spec":                       m.Spec,
			"state":                      m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelineResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":                      types.StringType,
			"cluster_id":                 types.StringType,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"effective_usage_policy_id":  types.StringType,
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
func (m *GetPipelineResponse) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo, bool) {
	if m.LatestUpdates.IsNull() || m.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo
	d := m.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in GetPipelineResponse.
func (m *GetPipelineResponse) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestUpdates = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in GetPipelineResponse as
// a RunAs value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPipelineResponse) GetRunAs(ctx context.Context) (RunAs, bool) {
	var e RunAs
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v RunAs
	d := m.RunAs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunAs sets the value of the RunAs field in GetPipelineResponse.
func (m *GetPipelineResponse) SetRunAs(ctx context.Context, v RunAs) {
	vs := v.ToObjectValue(ctx)
	m.RunAs = vs
}

// GetSpec returns the value of the Spec field in GetPipelineResponse as
// a PipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPipelineResponse) GetSpec(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v PipelineSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in GetPipelineResponse.
func (m *GetPipelineResponse) SetSpec(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
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

func (m GetUpdateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateRequest
// only implements ToObjectValue() and Type().
func (m GetUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
			"update_id":   m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUpdateRequest) Type(ctx context.Context) attr.Type {
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

func (m GetUpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetUpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"update": reflect.TypeOf(UpdateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateResponse
// only implements ToObjectValue() and Type().
func (m GetUpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update": m.Update,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update": UpdateInfo{}.Type(ctx),
		},
	}
}

// GetUpdate returns the value of the Update field in GetUpdateResponse as
// a UpdateInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetUpdateResponse) GetUpdate(ctx context.Context) (UpdateInfo, bool) {
	var e UpdateInfo
	if m.Update.IsNull() || m.Update.IsUnknown() {
		return e, false
	}
	var v UpdateInfo
	d := m.Update.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdate sets the value of the Update field in GetUpdateResponse.
func (m *GetUpdateResponse) SetUpdate(ctx context.Context, v UpdateInfo) {
	vs := v.ToObjectValue(ctx)
	m.Update = vs
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

func (m IngestionConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"report": reflect.TypeOf(ReportSpec{}),
		"schema": reflect.TypeOf(SchemaSpec{}),
		"table":  reflect.TypeOf(TableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionConfig
// only implements ToObjectValue() and Type().
func (m IngestionConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"report": m.Report,
			"schema": m.Schema,
			"table":  m.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionConfig) Type(ctx context.Context) attr.Type {
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
func (m *IngestionConfig) GetReport(ctx context.Context) (ReportSpec, bool) {
	var e ReportSpec
	if m.Report.IsNull() || m.Report.IsUnknown() {
		return e, false
	}
	var v ReportSpec
	d := m.Report.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReport sets the value of the Report field in IngestionConfig.
func (m *IngestionConfig) SetReport(ctx context.Context, v ReportSpec) {
	vs := v.ToObjectValue(ctx)
	m.Report = vs
}

// GetSchema returns the value of the Schema field in IngestionConfig as
// a SchemaSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionConfig) GetSchema(ctx context.Context) (SchemaSpec, bool) {
	var e SchemaSpec
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return e, false
	}
	var v SchemaSpec
	d := m.Schema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchema sets the value of the Schema field in IngestionConfig.
func (m *IngestionConfig) SetSchema(ctx context.Context, v SchemaSpec) {
	vs := v.ToObjectValue(ctx)
	m.Schema = vs
}

// GetTable returns the value of the Table field in IngestionConfig as
// a TableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionConfig) GetTable(ctx context.Context) (TableSpec, bool) {
	var e TableSpec
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v TableSpec
	d := m.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in IngestionConfig.
func (m *IngestionConfig) SetTable(ctx context.Context, v TableSpec) {
	vs := v.ToObjectValue(ctx)
	m.Table = vs
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	ConnectionId types.String `tfsdk:"connection_id"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Optional, Internal. Parameters required to establish an initial
	// connection with the source.
	ConnectionParameters types.Object `tfsdk:"connection_parameters"`
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
	if !from.ConnectionParameters.IsNull() && !from.ConnectionParameters.IsUnknown() {
		if toConnectionParameters, ok := to.GetConnectionParameters(ctx); ok {
			if fromConnectionParameters, ok := from.GetConnectionParameters(ctx); ok {
				// Recursively sync the fields of ConnectionParameters
				toConnectionParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromConnectionParameters)
				to.SetConnectionParameters(ctx, toConnectionParameters)
			}
		}
	}
}

func (to *IngestionGatewayPipelineDefinition) SyncFieldsDuringRead(ctx context.Context, from IngestionGatewayPipelineDefinition) {
	if !from.ConnectionParameters.IsNull() && !from.ConnectionParameters.IsUnknown() {
		if toConnectionParameters, ok := to.GetConnectionParameters(ctx); ok {
			if fromConnectionParameters, ok := from.GetConnectionParameters(ctx); ok {
				toConnectionParameters.SyncFieldsDuringRead(ctx, fromConnectionParameters)
				to.SetConnectionParameters(ctx, toConnectionParameters)
			}
		}
	}
}

func (m IngestionGatewayPipelineDefinition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_id"] = attrs["connection_id"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetRequired()
	attrs["connection_parameters"] = attrs["connection_parameters"].SetOptional()
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
func (m IngestionGatewayPipelineDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"connection_parameters": reflect.TypeOf(ConnectionParameters{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionGatewayPipelineDefinition
// only implements ToObjectValue() and Type().
func (m IngestionGatewayPipelineDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_id":           m.ConnectionId,
			"connection_name":         m.ConnectionName,
			"connection_parameters":   m.ConnectionParameters,
			"gateway_storage_catalog": m.GatewayStorageCatalog,
			"gateway_storage_name":    m.GatewayStorageName,
			"gateway_storage_schema":  m.GatewayStorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionGatewayPipelineDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_id":           types.StringType,
			"connection_name":         types.StringType,
			"connection_parameters":   ConnectionParameters{}.Type(ctx),
			"gateway_storage_catalog": types.StringType,
			"gateway_storage_name":    types.StringType,
			"gateway_storage_schema":  types.StringType,
		},
	}
}

// GetConnectionParameters returns the value of the ConnectionParameters field in IngestionGatewayPipelineDefinition as
// a ConnectionParameters value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionGatewayPipelineDefinition) GetConnectionParameters(ctx context.Context) (ConnectionParameters, bool) {
	var e ConnectionParameters
	if m.ConnectionParameters.IsNull() || m.ConnectionParameters.IsUnknown() {
		return e, false
	}
	var v ConnectionParameters
	d := m.ConnectionParameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConnectionParameters sets the value of the ConnectionParameters field in IngestionGatewayPipelineDefinition.
func (m *IngestionGatewayPipelineDefinition) SetConnectionParameters(ctx context.Context, v ConnectionParameters) {
	vs := v.ToObjectValue(ctx)
	m.ConnectionParameters = vs
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Immutable. If set to true, the pipeline will ingest tables from the UC
	// foreign catalogs directly without the need to specify a UC connection or
	// ingestion gateway. The `source_catalog` fields in objects of
	// IngestionConfig are interpreted as the UC foreign catalogs to ingest
	// from.
	IngestFromUcForeignCatalog types.Bool `tfsdk:"ingest_from_uc_foreign_catalog"`
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

func (m IngestionPipelineDefinition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["ingest_from_uc_foreign_catalog"] = attrs["ingest_from_uc_foreign_catalog"].SetOptional()
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
func (m IngestionPipelineDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects":               reflect.TypeOf(IngestionConfig{}),
		"source_configurations": reflect.TypeOf(SourceConfig{}),
		"table_configuration":   reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name":                m.ConnectionName,
			"ingest_from_uc_foreign_catalog": m.IngestFromUcForeignCatalog,
			"ingestion_gateway_id":           m.IngestionGatewayId,
			"netsuite_jar_path":              m.NetsuiteJarPath,
			"objects":                        m.Objects,
			"source_configurations":          m.SourceConfigurations,
			"source_type":                    m.SourceType,
			"table_configuration":            m.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":                types.StringType,
			"ingest_from_uc_foreign_catalog": types.BoolType,
			"ingestion_gateway_id":           types.StringType,
			"netsuite_jar_path":              types.StringType,
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
func (m *IngestionPipelineDefinition) GetObjects(ctx context.Context) ([]IngestionConfig, bool) {
	if m.Objects.IsNull() || m.Objects.IsUnknown() {
		return nil, false
	}
	var v []IngestionConfig
	d := m.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in IngestionPipelineDefinition.
func (m *IngestionPipelineDefinition) SetObjects(ctx context.Context, v []IngestionConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Objects = types.ListValueMust(t, vs)
}

// GetSourceConfigurations returns the value of the SourceConfigurations field in IngestionPipelineDefinition as
// a slice of SourceConfig values.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinition) GetSourceConfigurations(ctx context.Context) ([]SourceConfig, bool) {
	if m.SourceConfigurations.IsNull() || m.SourceConfigurations.IsUnknown() {
		return nil, false
	}
	var v []SourceConfig
	d := m.SourceConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSourceConfigurations sets the value of the SourceConfigurations field in IngestionPipelineDefinition.
func (m *IngestionPipelineDefinition) SetSourceConfigurations(ctx context.Context, v []SourceConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["source_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SourceConfigurations = types.ListValueMust(t, vs)
}

// GetTableConfiguration returns the value of the TableConfiguration field in IngestionPipelineDefinition as
// a TableSpecificConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinition) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v TableSpecificConfig
	d := m.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in IngestionPipelineDefinition.
func (m *IngestionPipelineDefinition) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	m.TableConfiguration = vs
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

func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cursor_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cursor_columns":                             m.CursorColumns,
			"deletion_condition":                         m.DeletionCondition,
			"hard_deletion_sync_min_interval_in_seconds": m.HardDeletionSyncMinIntervalInSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) Type(ctx context.Context) attr.Type {
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
func (m *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) GetCursorColumns(ctx context.Context) ([]types.String, bool) {
	if m.CursorColumns.IsNull() || m.CursorColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.CursorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCursorColumns sets the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig.
func (m *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) SetCursorColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cursor_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CursorColumns = types.ListValueMust(t, vs)
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

func (m IngestionPipelineDefinitionWorkdayReportParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionWorkdayReportParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters":        reflect.TypeOf(types.String{}),
		"report_parameters": reflect.TypeOf(IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParameters
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionWorkdayReportParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"incremental":       m.Incremental,
			"parameters":        m.Parameters,
			"report_parameters": m.ReportParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionWorkdayReportParameters) Type(ctx context.Context) attr.Type {
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
func (m *IngestionPipelineDefinitionWorkdayReportParameters) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in IngestionPipelineDefinitionWorkdayReportParameters.
func (m *IngestionPipelineDefinitionWorkdayReportParameters) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.MapValueMust(t, vs)
}

// GetReportParameters returns the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters as
// a slice of IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinitionWorkdayReportParameters) GetReportParameters(ctx context.Context) ([]IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue, bool) {
	if m.ReportParameters.IsNull() || m.ReportParameters.IsUnknown() {
		return nil, false
	}
	var v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue
	d := m.ReportParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReportParameters sets the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters.
func (m *IngestionPipelineDefinitionWorkdayReportParameters) SetReportParameters(ctx context.Context, v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["report_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ReportParameters = types.ListValueMust(t, vs)
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

func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue) Type(ctx context.Context) attr.Type {
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

func (m ListPipelineEventsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelineEventsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsRequest
// only implements ToObjectValue() and Type().
func (m ListPipelineEventsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      m.Filter,
			"max_results": m.MaxResults,
			"order_by":    m.OrderBy,
			"page_token":  m.PageToken,
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelineEventsRequest) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelineEventsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in ListPipelineEventsRequest.
func (m *ListPipelineEventsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
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

func (m ListPipelineEventsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelineEventsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(PipelineEvent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsResponse
// only implements ToObjectValue() and Type().
func (m ListPipelineEventsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          m.Events,
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelineEventsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelineEventsResponse) GetEvents(ctx context.Context) ([]PipelineEvent, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []PipelineEvent
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListPipelineEventsResponse.
func (m *ListPipelineEventsResponse) SetEvents(ctx context.Context, v []PipelineEvent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
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

func (m ListPipelinesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelinesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesRequest
// only implements ToObjectValue() and Type().
func (m ListPipelinesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      m.Filter,
			"max_results": m.MaxResults,
			"order_by":    m.OrderBy,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelinesRequest) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelinesRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in ListPipelinesRequest.
func (m *ListPipelinesRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
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

func (m ListPipelinesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelinesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(PipelineStateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesResponse
// only implements ToObjectValue() and Type().
func (m ListPipelinesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"statuses":        m.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelinesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelinesResponse) GetStatuses(ctx context.Context) ([]PipelineStateInfo, bool) {
	if m.Statuses.IsNull() || m.Statuses.IsUnknown() {
		return nil, false
	}
	var v []PipelineStateInfo
	d := m.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListPipelinesResponse.
func (m *ListPipelinesResponse) SetStatuses(ctx context.Context, v []PipelineStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Statuses = types.ListValueMust(t, vs)
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

func (m ListUpdatesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListUpdatesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesRequest
// only implements ToObjectValue() and Type().
func (m ListUpdatesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":     m.MaxResults,
			"page_token":      m.PageToken,
			"pipeline_id":     m.PipelineId,
			"until_update_id": m.UntilUpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUpdatesRequest) Type(ctx context.Context) attr.Type {
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

func (m ListUpdatesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListUpdatesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(UpdateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesResponse
// only implements ToObjectValue() and Type().
func (m ListUpdatesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
			"updates":         m.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUpdatesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListUpdatesResponse) GetUpdates(ctx context.Context) ([]UpdateInfo, bool) {
	if m.Updates.IsNull() || m.Updates.IsUnknown() {
		return nil, false
	}
	var v []UpdateInfo
	d := m.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in ListUpdatesResponse.
func (m *ListUpdatesResponse) SetUpdates(ctx context.Context, v []UpdateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Updates = types.ListValueMust(t, vs)
}

type ManualTrigger struct {
}

func (to *ManualTrigger) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManualTrigger) {
}

func (to *ManualTrigger) SyncFieldsDuringRead(ctx context.Context, from ManualTrigger) {
}

func (m ManualTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManualTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ManualTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManualTrigger
// only implements ToObjectValue() and Type().
func (m ManualTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ManualTrigger) Type(ctx context.Context) attr.Type {
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

func (m NotebookLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NotebookLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookLibrary
// only implements ToObjectValue() and Type().
func (m NotebookLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotebookLibrary) Type(ctx context.Context) attr.Type {
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

func (m Notifications) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Notifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":           reflect.TypeOf(types.String{}),
		"email_recipients": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Notifications
// only implements ToObjectValue() and Type().
func (m Notifications) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":           m.Alerts,
			"email_recipients": m.EmailRecipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Notifications) Type(ctx context.Context) attr.Type {
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
func (m *Notifications) GetAlerts(ctx context.Context) ([]types.String, bool) {
	if m.Alerts.IsNull() || m.Alerts.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Alerts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAlerts sets the value of the Alerts field in Notifications.
func (m *Notifications) SetAlerts(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Alerts = types.ListValueMust(t, vs)
}

// GetEmailRecipients returns the value of the EmailRecipients field in Notifications as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Notifications) GetEmailRecipients(ctx context.Context) ([]types.String, bool) {
	if m.EmailRecipients.IsNull() || m.EmailRecipients.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EmailRecipients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailRecipients sets the value of the EmailRecipients field in Notifications.
func (m *Notifications) SetEmailRecipients(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmailRecipients = types.ListValueMust(t, vs)
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

func (m Origin) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Origin) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Origin
// only implements ToObjectValue() and Type().
func (m Origin) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"batch_id":             m.BatchId,
			"cloud":                m.Cloud,
			"cluster_id":           m.ClusterId,
			"dataset_name":         m.DatasetName,
			"flow_id":              m.FlowId,
			"flow_name":            m.FlowName,
			"host":                 m.Host,
			"maintenance_id":       m.MaintenanceId,
			"materialization_name": m.MaterializationName,
			"org_id":               m.OrgId,
			"pipeline_id":          m.PipelineId,
			"pipeline_name":        m.PipelineName,
			"region":               m.Region,
			"request_id":           m.RequestId,
			"table_id":             m.TableId,
			"uc_resource_id":       m.UcResourceId,
			"update_id":            m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Origin) Type(ctx context.Context) attr.Type {
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

func (m PathPattern) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PathPattern) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PathPattern
// only implements ToObjectValue() and Type().
func (m PathPattern) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include": m.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PathPattern) Type(ctx context.Context) attr.Type {
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

func (m PipelineAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlRequest
// only implements ToObjectValue() and Type().
func (m PipelineAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (m PipelineAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PipelinePermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlResponse
// only implements ToObjectValue() and Type().
func (m PipelineAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (m *PipelineAccessControlResponse) GetAllPermissions(ctx context.Context) ([]PipelinePermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PipelineAccessControlResponse.
func (m *PipelineAccessControlResponse) SetAllPermissions(ctx context.Context, v []PipelinePermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
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

func (m PipelineCluster) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineCluster) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apply_policy_default_values":  m.ApplyPolicyDefaultValues,
			"autoscale":                    m.Autoscale,
			"aws_attributes":               m.AwsAttributes,
			"azure_attributes":             m.AzureAttributes,
			"cluster_log_conf":             m.ClusterLogConf,
			"custom_tags":                  m.CustomTags,
			"driver_instance_pool_id":      m.DriverInstancePoolId,
			"driver_node_type_id":          m.DriverNodeTypeId,
			"enable_local_disk_encryption": m.EnableLocalDiskEncryption,
			"gcp_attributes":               m.GcpAttributes,
			"init_scripts":                 m.InitScripts,
			"instance_pool_id":             m.InstancePoolId,
			"label":                        m.Label,
			"node_type_id":                 m.NodeTypeId,
			"num_workers":                  m.NumWorkers,
			"policy_id":                    m.PolicyId,
			"spark_conf":                   m.SparkConf,
			"spark_env_vars":               m.SparkEnvVars,
			"ssh_public_keys":              m.SshPublicKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineCluster) Type(ctx context.Context) attr.Type {
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
func (m *PipelineCluster) GetAutoscale(ctx context.Context) (PipelineClusterAutoscale, bool) {
	var e PipelineClusterAutoscale
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v PipelineClusterAutoscale
	d := m.Autoscale.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoscale sets the value of the Autoscale field in PipelineCluster.
func (m *PipelineCluster) SetAutoscale(ctx context.Context, v PipelineClusterAutoscale) {
	vs := v.ToObjectValue(ctx)
	m.Autoscale = vs
}

// GetAwsAttributes returns the value of the AwsAttributes field in PipelineCluster as
// a compute_tf.AwsAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetAwsAttributes(ctx context.Context) (compute_tf.AwsAttributes, bool) {
	var e compute_tf.AwsAttributes
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v compute_tf.AwsAttributes
	d := m.AwsAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsAttributes sets the value of the AwsAttributes field in PipelineCluster.
func (m *PipelineCluster) SetAwsAttributes(ctx context.Context, v compute_tf.AwsAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AwsAttributes = vs
}

// GetAzureAttributes returns the value of the AzureAttributes field in PipelineCluster as
// a compute_tf.AzureAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetAzureAttributes(ctx context.Context) (compute_tf.AzureAttributes, bool) {
	var e compute_tf.AzureAttributes
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v compute_tf.AzureAttributes
	d := m.AzureAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureAttributes sets the value of the AzureAttributes field in PipelineCluster.
func (m *PipelineCluster) SetAzureAttributes(ctx context.Context, v compute_tf.AzureAttributes) {
	vs := v.ToObjectValue(ctx)
	m.AzureAttributes = vs
}

// GetClusterLogConf returns the value of the ClusterLogConf field in PipelineCluster as
// a compute_tf.ClusterLogConf value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetClusterLogConf(ctx context.Context) (compute_tf.ClusterLogConf, bool) {
	var e compute_tf.ClusterLogConf
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v compute_tf.ClusterLogConf
	d := m.ClusterLogConf.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in PipelineCluster.
func (m *PipelineCluster) SetClusterLogConf(ctx context.Context, v compute_tf.ClusterLogConf) {
	vs := v.ToObjectValue(ctx)
	m.ClusterLogConf = vs
}

// GetCustomTags returns the value of the CustomTags field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in PipelineCluster.
func (m *PipelineCluster) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in PipelineCluster as
// a compute_tf.GcpAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetGcpAttributes(ctx context.Context) (compute_tf.GcpAttributes, bool) {
	var e compute_tf.GcpAttributes
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v compute_tf.GcpAttributes
	d := m.GcpAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpAttributes sets the value of the GcpAttributes field in PipelineCluster.
func (m *PipelineCluster) SetGcpAttributes(ctx context.Context, v compute_tf.GcpAttributes) {
	vs := v.ToObjectValue(ctx)
	m.GcpAttributes = vs
}

// GetInitScripts returns the value of the InitScripts field in PipelineCluster as
// a slice of compute_tf.InitScriptInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetInitScripts(ctx context.Context) ([]compute_tf.InitScriptInfo, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.InitScriptInfo
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in PipelineCluster.
func (m *PipelineCluster) SetInitScripts(ctx context.Context, v []compute_tf.InitScriptInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkConf.IsNull() || m.SparkConf.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkConf sets the value of the SparkConf field in PipelineCluster.
func (m *PipelineCluster) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in PipelineCluster as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
	if m.SparkEnvVars.IsNull() || m.SparkEnvVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.SparkEnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkEnvVars sets the value of the SparkEnvVars field in PipelineCluster.
func (m *PipelineCluster) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in PipelineCluster as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
	if m.SshPublicKeys.IsNull() || m.SshPublicKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SshPublicKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSshPublicKeys sets the value of the SshPublicKeys field in PipelineCluster.
func (m *PipelineCluster) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
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

func (m PipelineClusterAutoscale) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineClusterAutoscale) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineClusterAutoscale
// only implements ToObjectValue() and Type().
func (m PipelineClusterAutoscale) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": m.MaxWorkers,
			"min_workers": m.MinWorkers,
			"mode":        m.Mode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineClusterAutoscale) Type(ctx context.Context) attr.Type {
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

func (m PipelineDeployment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineDeployment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineDeployment
// only implements ToObjectValue() and Type().
func (m PipelineDeployment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kind":               m.Kind,
			"metadata_file_path": m.MetadataFilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineDeployment) Type(ctx context.Context) attr.Type {
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

func (m PipelineEvent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineEvent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error":    reflect.TypeOf(ErrorDetail{}),
		"origin":   reflect.TypeOf(Origin{}),
		"sequence": reflect.TypeOf(Sequencing{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineEvent
// only implements ToObjectValue() and Type().
func (m PipelineEvent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":          m.Error,
			"event_type":     m.EventType,
			"id":             m.Id,
			"level":          m.Level,
			"maturity_level": m.MaturityLevel,
			"message":        m.Message,
			"origin":         m.Origin,
			"sequence":       m.Sequence,
			"timestamp":      m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineEvent) Type(ctx context.Context) attr.Type {
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
func (m *PipelineEvent) GetError(ctx context.Context) (ErrorDetail, bool) {
	var e ErrorDetail
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v ErrorDetail
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in PipelineEvent.
func (m *PipelineEvent) SetError(ctx context.Context, v ErrorDetail) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

// GetOrigin returns the value of the Origin field in PipelineEvent as
// a Origin value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineEvent) GetOrigin(ctx context.Context) (Origin, bool) {
	var e Origin
	if m.Origin.IsNull() || m.Origin.IsUnknown() {
		return e, false
	}
	var v Origin
	d := m.Origin.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrigin sets the value of the Origin field in PipelineEvent.
func (m *PipelineEvent) SetOrigin(ctx context.Context, v Origin) {
	vs := v.ToObjectValue(ctx)
	m.Origin = vs
}

// GetSequence returns the value of the Sequence field in PipelineEvent as
// a Sequencing value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineEvent) GetSequence(ctx context.Context) (Sequencing, bool) {
	var e Sequencing
	if m.Sequence.IsNull() || m.Sequence.IsUnknown() {
		return e, false
	}
	var v Sequencing
	d := m.Sequence.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSequence sets the value of the Sequence field in PipelineEvent.
func (m *PipelineEvent) SetSequence(ctx context.Context, v Sequencing) {
	vs := v.ToObjectValue(ctx)
	m.Sequence = vs
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

func (m PipelineLibrary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineLibrary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineLibrary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file":     m.File,
			"glob":     m.Glob,
			"jar":      m.Jar,
			"maven":    m.Maven,
			"notebook": m.Notebook,
			"whl":      m.Whl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineLibrary) Type(ctx context.Context) attr.Type {
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
func (m *PipelineLibrary) GetFile(ctx context.Context) (FileLibrary, bool) {
	var e FileLibrary
	if m.File.IsNull() || m.File.IsUnknown() {
		return e, false
	}
	var v FileLibrary
	d := m.File.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFile sets the value of the File field in PipelineLibrary.
func (m *PipelineLibrary) SetFile(ctx context.Context, v FileLibrary) {
	vs := v.ToObjectValue(ctx)
	m.File = vs
}

// GetGlob returns the value of the Glob field in PipelineLibrary as
// a PathPattern value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary) GetGlob(ctx context.Context) (PathPattern, bool) {
	var e PathPattern
	if m.Glob.IsNull() || m.Glob.IsUnknown() {
		return e, false
	}
	var v PathPattern
	d := m.Glob.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGlob sets the value of the Glob field in PipelineLibrary.
func (m *PipelineLibrary) SetGlob(ctx context.Context, v PathPattern) {
	vs := v.ToObjectValue(ctx)
	m.Glob = vs
}

// GetMaven returns the value of the Maven field in PipelineLibrary as
// a compute_tf.MavenLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary) GetMaven(ctx context.Context) (compute_tf.MavenLibrary, bool) {
	var e compute_tf.MavenLibrary
	if m.Maven.IsNull() || m.Maven.IsUnknown() {
		return e, false
	}
	var v compute_tf.MavenLibrary
	d := m.Maven.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMaven sets the value of the Maven field in PipelineLibrary.
func (m *PipelineLibrary) SetMaven(ctx context.Context, v compute_tf.MavenLibrary) {
	vs := v.ToObjectValue(ctx)
	m.Maven = vs
}

// GetNotebook returns the value of the Notebook field in PipelineLibrary as
// a NotebookLibrary value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary) GetNotebook(ctx context.Context) (NotebookLibrary, bool) {
	var e NotebookLibrary
	if m.Notebook.IsNull() || m.Notebook.IsUnknown() {
		return e, false
	}
	var v NotebookLibrary
	d := m.Notebook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebook sets the value of the Notebook field in PipelineLibrary.
func (m *PipelineLibrary) SetNotebook(ctx context.Context, v NotebookLibrary) {
	vs := v.ToObjectValue(ctx)
	m.Notebook = vs
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

func (m PipelinePermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermission
// only implements ToObjectValue() and Type().
func (m PipelinePermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermission) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in PipelinePermission.
func (m *PipelinePermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
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

func (m PipelinePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissions
// only implements ToObjectValue() and Type().
func (m PipelinePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissions) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermissions) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissions.
func (m *PipelinePermissions) SetAccessControlList(ctx context.Context, v []PipelineAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type PipelinePermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *PipelinePermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissionsDescription) {
}

func (to *PipelinePermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissionsDescription) {
}

func (m PipelinePermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsDescription
// only implements ToObjectValue() and Type().
func (m PipelinePermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissionsDescription) Type(ctx context.Context) attr.Type {
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

func (m PipelinePermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsRequest
// only implements ToObjectValue() and Type().
func (m PipelinePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"pipeline_id":         m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermissionsRequest) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissionsRequest.
func (m *PipelinePermissionsRequest) SetAccessControlList(ctx context.Context, v []PipelineAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
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

func (m PipelineSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PipelineSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":     m.BudgetPolicyId,
			"catalog":              m.Catalog,
			"channel":              m.Channel,
			"clusters":             m.Clusters,
			"configuration":        m.Configuration,
			"continuous":           m.Continuous,
			"deployment":           m.Deployment,
			"development":          m.Development,
			"edition":              m.Edition,
			"environment":          m.Environment,
			"event_log":            m.EventLog,
			"filters":              m.Filters,
			"gateway_definition":   m.GatewayDefinition,
			"id":                   m.Id,
			"ingestion_definition": m.IngestionDefinition,
			"libraries":            m.Libraries,
			"name":                 m.Name,
			"notifications":        m.Notifications,
			"photon":               m.Photon,
			"restart_window":       m.RestartWindow,
			"root_path":            m.RootPath,
			"schema":               m.Schema,
			"serverless":           m.Serverless,
			"storage":              m.Storage,
			"tags":                 m.Tags,
			"target":               m.Target,
			"trigger":              m.Trigger,
			"usage_policy_id":      m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineSpec) Type(ctx context.Context) attr.Type {
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
			"target":          types.StringType,
			"trigger":         PipelineTrigger{}.Type(ctx),
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in PipelineSpec as
// a slice of PipelineCluster values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetClusters(ctx context.Context) ([]PipelineCluster, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in PipelineSpec.
func (m *PipelineSpec) SetClusters(ctx context.Context, v []PipelineCluster) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in PipelineSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
	if m.Configuration.IsNull() || m.Configuration.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Configuration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfiguration sets the value of the Configuration field in PipelineSpec.
func (m *PipelineSpec) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in PipelineSpec as
// a PipelineDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetDeployment(ctx context.Context) (PipelineDeployment, bool) {
	var e PipelineDeployment
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v PipelineDeployment
	d := m.Deployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeployment sets the value of the Deployment field in PipelineSpec.
func (m *PipelineSpec) SetDeployment(ctx context.Context, v PipelineDeployment) {
	vs := v.ToObjectValue(ctx)
	m.Deployment = vs
}

// GetEnvironment returns the value of the Environment field in PipelineSpec as
// a PipelinesEnvironment value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetEnvironment(ctx context.Context) (PipelinesEnvironment, bool) {
	var e PipelinesEnvironment
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v PipelinesEnvironment
	d := m.Environment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironment sets the value of the Environment field in PipelineSpec.
func (m *PipelineSpec) SetEnvironment(ctx context.Context, v PipelinesEnvironment) {
	vs := v.ToObjectValue(ctx)
	m.Environment = vs
}

// GetEventLog returns the value of the EventLog field in PipelineSpec as
// a EventLogSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetEventLog(ctx context.Context) (EventLogSpec, bool) {
	var e EventLogSpec
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v EventLogSpec
	d := m.EventLog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEventLog sets the value of the EventLog field in PipelineSpec.
func (m *PipelineSpec) SetEventLog(ctx context.Context, v EventLogSpec) {
	vs := v.ToObjectValue(ctx)
	m.EventLog = vs
}

// GetFilters returns the value of the Filters field in PipelineSpec as
// a Filters value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetFilters(ctx context.Context) (Filters, bool) {
	var e Filters
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v Filters
	d := m.Filters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in PipelineSpec.
func (m *PipelineSpec) SetFilters(ctx context.Context, v Filters) {
	vs := v.ToObjectValue(ctx)
	m.Filters = vs
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in PipelineSpec as
// a IngestionGatewayPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition, bool) {
	var e IngestionGatewayPipelineDefinition
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionGatewayPipelineDefinition
	d := m.GatewayDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in PipelineSpec.
func (m *PipelineSpec) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.GatewayDefinition = vs
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in PipelineSpec as
// a IngestionPipelineDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition, bool) {
	var e IngestionPipelineDefinition
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinition
	d := m.IngestionDefinition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in PipelineSpec.
func (m *PipelineSpec) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition) {
	vs := v.ToObjectValue(ctx)
	m.IngestionDefinition = vs
}

// GetLibraries returns the value of the Libraries field in PipelineSpec as
// a slice of PipelineLibrary values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetLibraries(ctx context.Context) ([]PipelineLibrary, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in PipelineSpec.
func (m *PipelineSpec) SetLibraries(ctx context.Context, v []PipelineLibrary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in PipelineSpec as
// a slice of Notifications values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetNotifications(ctx context.Context) ([]Notifications, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in PipelineSpec.
func (m *PipelineSpec) SetNotifications(ctx context.Context, v []Notifications) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in PipelineSpec as
// a RestartWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetRestartWindow(ctx context.Context) (RestartWindow, bool) {
	var e RestartWindow
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v RestartWindow
	d := m.RestartWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestartWindow sets the value of the RestartWindow field in PipelineSpec.
func (m *PipelineSpec) SetRestartWindow(ctx context.Context, v RestartWindow) {
	vs := v.ToObjectValue(ctx)
	m.RestartWindow = vs
}

// GetTags returns the value of the Tags field in PipelineSpec as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in PipelineSpec.
func (m *PipelineSpec) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in PipelineSpec as
// a PipelineTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec) GetTrigger(ctx context.Context) (PipelineTrigger, bool) {
	var e PipelineTrigger
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v PipelineTrigger
	d := m.Trigger.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrigger sets the value of the Trigger field in PipelineSpec.
func (m *PipelineSpec) SetTrigger(ctx context.Context, v PipelineTrigger) {
	vs := v.ToObjectValue(ctx)
	m.Trigger = vs
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

func (m PipelineStateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineStateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineStateInfo
// only implements ToObjectValue() and Type().
func (m PipelineStateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":        m.ClusterId,
			"creator_user_name": m.CreatorUserName,
			"health":            m.Health,
			"latest_updates":    m.LatestUpdates,
			"name":              m.Name,
			"pipeline_id":       m.PipelineId,
			"run_as_user_name":  m.RunAsUserName,
			"state":             m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineStateInfo) Type(ctx context.Context) attr.Type {
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
func (m *PipelineStateInfo) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo, bool) {
	if m.LatestUpdates.IsNull() || m.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo
	d := m.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in PipelineStateInfo.
func (m *PipelineStateInfo) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestUpdates = types.ListValueMust(t, vs)
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

func (m PipelineTrigger) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineTrigger) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron":   reflect.TypeOf(CronTrigger{}),
		"manual": reflect.TypeOf(ManualTrigger{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineTrigger
// only implements ToObjectValue() and Type().
func (m PipelineTrigger) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cron":   m.Cron,
			"manual": m.Manual,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineTrigger) Type(ctx context.Context) attr.Type {
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
func (m *PipelineTrigger) GetCron(ctx context.Context) (CronTrigger, bool) {
	var e CronTrigger
	if m.Cron.IsNull() || m.Cron.IsUnknown() {
		return e, false
	}
	var v CronTrigger
	d := m.Cron.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCron sets the value of the Cron field in PipelineTrigger.
func (m *PipelineTrigger) SetCron(ctx context.Context, v CronTrigger) {
	vs := v.ToObjectValue(ctx)
	m.Cron = vs
}

// GetManual returns the value of the Manual field in PipelineTrigger as
// a ManualTrigger value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineTrigger) GetManual(ctx context.Context) (ManualTrigger, bool) {
	var e ManualTrigger
	if m.Manual.IsNull() || m.Manual.IsUnknown() {
		return e, false
	}
	var v ManualTrigger
	d := m.Manual.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManual sets the value of the Manual field in PipelineTrigger.
func (m *PipelineTrigger) SetManual(ctx context.Context, v ManualTrigger) {
	vs := v.ToObjectValue(ctx)
	m.Manual = vs
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

func (m PipelinesEnvironment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinesEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinesEnvironment
// only implements ToObjectValue() and Type().
func (m PipelinesEnvironment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": m.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinesEnvironment) Type(ctx context.Context) attr.Type {
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
func (m *PipelinesEnvironment) GetDependencies(ctx context.Context) ([]types.String, bool) {
	if m.Dependencies.IsNull() || m.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in PipelinesEnvironment.
func (m *PipelinesEnvironment) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Dependencies = types.ListValueMust(t, vs)
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

func (m PostgresCatalogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PostgresCatalogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"slot_config": reflect.TypeOf(PostgresSlotConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresCatalogConfig
// only implements ToObjectValue() and Type().
func (m PostgresCatalogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"slot_config": m.SlotConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PostgresCatalogConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"slot_config": PostgresSlotConfig{}.Type(ctx),
		},
	}
}

// GetSlotConfig returns the value of the SlotConfig field in PostgresCatalogConfig as
// a PostgresSlotConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *PostgresCatalogConfig) GetSlotConfig(ctx context.Context) (PostgresSlotConfig, bool) {
	var e PostgresSlotConfig
	if m.SlotConfig.IsNull() || m.SlotConfig.IsUnknown() {
		return e, false
	}
	var v PostgresSlotConfig
	d := m.SlotConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlotConfig sets the value of the SlotConfig field in PostgresCatalogConfig.
func (m *PostgresCatalogConfig) SetSlotConfig(ctx context.Context, v PostgresSlotConfig) {
	vs := v.ToObjectValue(ctx)
	m.SlotConfig = vs
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

func (m PostgresSlotConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PostgresSlotConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresSlotConfig
// only implements ToObjectValue() and Type().
func (m PostgresSlotConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publication_name": m.PublicationName,
			"slot_name":        m.SlotName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PostgresSlotConfig) Type(ctx context.Context) attr.Type {
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

func (m ReportSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ReportSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReportSpec
// only implements ToObjectValue() and Type().
func (m ReportSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": m.DestinationCatalog,
			"destination_schema":  m.DestinationSchema,
			"destination_table":   m.DestinationTable,
			"source_url":          m.SourceUrl,
			"table_configuration": m.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReportSpec) Type(ctx context.Context) attr.Type {
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
func (m *ReportSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v TableSpecificConfig
	d := m.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in ReportSpec.
func (m *ReportSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	m.TableConfiguration = vs
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

func (m RestartWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestartWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"days_of_week": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartWindow
// only implements ToObjectValue() and Type().
func (m RestartWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"days_of_week": m.DaysOfWeek,
			"start_hour":   m.StartHour,
			"time_zone_id": m.TimeZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestartWindow) Type(ctx context.Context) attr.Type {
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
func (m *RestartWindow) GetDaysOfWeek(ctx context.Context) ([]types.String, bool) {
	if m.DaysOfWeek.IsNull() || m.DaysOfWeek.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DaysOfWeek.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDaysOfWeek sets the value of the DaysOfWeek field in RestartWindow.
func (m *RestartWindow) SetDaysOfWeek(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["days_of_week"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DaysOfWeek = types.ListValueMust(t, vs)
}

type RestorePipelineRequest struct {
	// The ID of the pipeline to restore
	PipelineId types.String `tfsdk:"-"`
}

func (to *RestorePipelineRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestorePipelineRequest) {
}

func (to *RestorePipelineRequest) SyncFieldsDuringRead(ctx context.Context, from RestorePipelineRequest) {
}

func (m RestorePipelineRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pipeline_id"] = attrs["pipeline_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestorePipelineRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestorePipelineRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestorePipelineRequest
// only implements ToObjectValue() and Type().
func (m RestorePipelineRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestorePipelineRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type RestorePipelineRequestResponse struct {
}

func (to *RestorePipelineRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestorePipelineRequestResponse) {
}

func (to *RestorePipelineRequestResponse) SyncFieldsDuringRead(ctx context.Context, from RestorePipelineRequestResponse) {
}

func (m RestorePipelineRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestorePipelineRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestorePipelineRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestorePipelineRequestResponse
// only implements ToObjectValue() and Type().
func (m RestorePipelineRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestorePipelineRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (m RunAs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunAs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunAs
// only implements ToObjectValue() and Type().
func (m RunAs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunAs) Type(ctx context.Context) attr.Type {
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

func (m SchemaSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SchemaSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaSpec
// only implements ToObjectValue() and Type().
func (m SchemaSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": m.DestinationCatalog,
			"destination_schema":  m.DestinationSchema,
			"source_catalog":      m.SourceCatalog,
			"source_schema":       m.SourceSchema,
			"table_configuration": m.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SchemaSpec) Type(ctx context.Context) attr.Type {
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
func (m *SchemaSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v TableSpecificConfig
	d := m.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in SchemaSpec.
func (m *SchemaSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	m.TableConfiguration = vs
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

func (m Sequencing) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Sequencing) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_id": reflect.TypeOf(DataPlaneId{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sequencing
// only implements ToObjectValue() and Type().
func (m Sequencing) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"control_plane_seq_no": m.ControlPlaneSeqNo,
			"data_plane_id":        m.DataPlaneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Sequencing) Type(ctx context.Context) attr.Type {
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
func (m *Sequencing) GetDataPlaneId(ctx context.Context) (DataPlaneId, bool) {
	var e DataPlaneId
	if m.DataPlaneId.IsNull() || m.DataPlaneId.IsUnknown() {
		return e, false
	}
	var v DataPlaneId
	d := m.DataPlaneId.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataPlaneId sets the value of the DataPlaneId field in Sequencing.
func (m *Sequencing) SetDataPlaneId(ctx context.Context, v DataPlaneId) {
	vs := v.ToObjectValue(ctx)
	m.DataPlaneId = vs
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

func (m SerializedException) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SerializedException) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stack": reflect.TypeOf(StackFrame{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SerializedException
// only implements ToObjectValue() and Type().
func (m SerializedException) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"class_name": m.ClassName,
			"message":    m.Message,
			"stack":      m.Stack,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SerializedException) Type(ctx context.Context) attr.Type {
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
func (m *SerializedException) GetStack(ctx context.Context) ([]StackFrame, bool) {
	if m.Stack.IsNull() || m.Stack.IsUnknown() {
		return nil, false
	}
	var v []StackFrame
	d := m.Stack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStack sets the value of the Stack field in SerializedException.
func (m *SerializedException) SetStack(ctx context.Context, v []StackFrame) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stack"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Stack = types.ListValueMust(t, vs)
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

func (m SourceCatalogConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SourceCatalogConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"postgres": reflect.TypeOf(PostgresCatalogConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceCatalogConfig
// only implements ToObjectValue() and Type().
func (m SourceCatalogConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"postgres":       m.Postgres,
			"source_catalog": m.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SourceCatalogConfig) Type(ctx context.Context) attr.Type {
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
func (m *SourceCatalogConfig) GetPostgres(ctx context.Context) (PostgresCatalogConfig, bool) {
	var e PostgresCatalogConfig
	if m.Postgres.IsNull() || m.Postgres.IsUnknown() {
		return e, false
	}
	var v PostgresCatalogConfig
	d := m.Postgres.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPostgres sets the value of the Postgres field in SourceCatalogConfig.
func (m *SourceCatalogConfig) SetPostgres(ctx context.Context, v PostgresCatalogConfig) {
	vs := v.ToObjectValue(ctx)
	m.Postgres = vs
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

func (m SourceConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SourceConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(SourceCatalogConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceConfig
// only implements ToObjectValue() and Type().
func (m SourceConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SourceConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": SourceCatalogConfig{}.Type(ctx),
		},
	}
}

// GetCatalog returns the value of the Catalog field in SourceConfig as
// a SourceCatalogConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *SourceConfig) GetCatalog(ctx context.Context) (SourceCatalogConfig, bool) {
	var e SourceCatalogConfig
	if m.Catalog.IsNull() || m.Catalog.IsUnknown() {
		return e, false
	}
	var v SourceCatalogConfig
	d := m.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalog sets the value of the Catalog field in SourceConfig.
func (m *SourceConfig) SetCatalog(ctx context.Context, v SourceCatalogConfig) {
	vs := v.ToObjectValue(ctx)
	m.Catalog = vs
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

func (m StackFrame) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StackFrame) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StackFrame
// only implements ToObjectValue() and Type().
func (m StackFrame) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"declaring_class": m.DeclaringClass,
			"file_name":       m.FileName,
			"line_number":     m.LineNumber,
			"method_name":     m.MethodName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StackFrame) Type(ctx context.Context) attr.Type {
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

func (m StartUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdate
// only implements ToObjectValue() and Type().
func (m StartUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                  m.Cause,
			"full_refresh":           m.FullRefresh,
			"full_refresh_selection": m.FullRefreshSelection,
			"pipeline_id":            m.PipelineId,
			"refresh_selection":      m.RefreshSelection,
			"validate_only":          m.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartUpdate) Type(ctx context.Context) attr.Type {
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
func (m *StartUpdate) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if m.FullRefreshSelection.IsNull() || m.FullRefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FullRefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in StartUpdate.
func (m *StartUpdate) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in StartUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *StartUpdate) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if m.RefreshSelection.IsNull() || m.RefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshSelection sets the value of the RefreshSelection field in StartUpdate.
func (m *StartUpdate) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RefreshSelection = types.ListValueMust(t, vs)
}

type StartUpdateResponse struct {
	UpdateId types.String `tfsdk:"update_id"`
}

func (to *StartUpdateResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartUpdateResponse) {
}

func (to *StartUpdateResponse) SyncFieldsDuringRead(ctx context.Context, from StartUpdateResponse) {
}

func (m StartUpdateResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartUpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdateResponse
// only implements ToObjectValue() and Type().
func (m StartUpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_id": m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartUpdateResponse) Type(ctx context.Context) attr.Type {
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

func (m StopPipelineResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StopPipelineResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopPipelineResponse
// only implements ToObjectValue() and Type().
func (m StopPipelineResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m StopPipelineResponse) Type(ctx context.Context) attr.Type {
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

func (m StopRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StopRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest
// only implements ToObjectValue() and Type().
func (m StopRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopRequest) Type(ctx context.Context) attr.Type {
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

func (m TableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpec
// only implements ToObjectValue() and Type().
func (m TableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_catalog": m.DestinationCatalog,
			"destination_schema":  m.DestinationSchema,
			"destination_table":   m.DestinationTable,
			"source_catalog":      m.SourceCatalog,
			"source_schema":       m.SourceSchema,
			"source_table":        m.SourceTable,
			"table_configuration": m.TableConfiguration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TableSpec) Type(ctx context.Context) attr.Type {
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
func (m *TableSpec) GetTableConfiguration(ctx context.Context) (TableSpecificConfig, bool) {
	var e TableSpecificConfig
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v TableSpecificConfig
	d := m.TableConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConfiguration sets the value of the TableConfiguration field in TableSpec.
func (m *TableSpec) SetTableConfiguration(ctx context.Context, v TableSpecificConfig) {
	vs := v.ToObjectValue(ctx)
	m.TableConfiguration = vs
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
	// (Optional, Immutable) The row filter condition to be applied to the
	// table. It must not contain the WHERE keyword, only the actual filter
	// condition. It must be in DBSQL format.
	RowFilter types.String `tfsdk:"row_filter"`
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

func (m TableSpecificConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exclude_columns"] = attrs["exclude_columns"].SetOptional()
	attrs["include_columns"] = attrs["include_columns"].SetOptional()
	attrs["primary_keys"] = attrs["primary_keys"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].SetOptional()
	attrs["row_filter"] = attrs["row_filter"].SetOptional()
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
func (m TableSpecificConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m TableSpecificConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude_columns":                   m.ExcludeColumns,
			"include_columns":                   m.IncludeColumns,
			"primary_keys":                      m.PrimaryKeys,
			"query_based_connector_config":      m.QueryBasedConnectorConfig,
			"row_filter":                        m.RowFilter,
			"salesforce_include_formula_fields": m.SalesforceIncludeFormulaFields,
			"scd_type":                          m.ScdType,
			"sequence_by":                       m.SequenceBy,
			"workday_report_parameters":         m.WorkdayReportParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TableSpecificConfig) Type(ctx context.Context) attr.Type {
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
			"row_filter":                        types.StringType,
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
func (m *TableSpecificConfig) GetExcludeColumns(ctx context.Context) ([]types.String, bool) {
	if m.ExcludeColumns.IsNull() || m.ExcludeColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ExcludeColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExcludeColumns sets the value of the ExcludeColumns field in TableSpecificConfig.
func (m *TableSpecificConfig) SetExcludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExcludeColumns = types.ListValueMust(t, vs)
}

// GetIncludeColumns returns the value of the IncludeColumns field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig) GetIncludeColumns(ctx context.Context) ([]types.String, bool) {
	if m.IncludeColumns.IsNull() || m.IncludeColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IncludeColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIncludeColumns sets the value of the IncludeColumns field in TableSpecificConfig.
func (m *TableSpecificConfig) SetIncludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IncludeColumns = types.ListValueMust(t, vs)
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.PrimaryKeys.IsNull() || m.PrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeys sets the value of the PrimaryKeys field in TableSpecificConfig.
func (m *TableSpecificConfig) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeys = types.ListValueMust(t, vs)
}

// GetQueryBasedConnectorConfig returns the value of the QueryBasedConnectorConfig field in TableSpecificConfig as
// a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig) GetQueryBasedConnectorConfig(ctx context.Context) (IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig, bool) {
	var e IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
	if m.QueryBasedConnectorConfig.IsNull() || m.QueryBasedConnectorConfig.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig
	d := m.QueryBasedConnectorConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryBasedConnectorConfig sets the value of the QueryBasedConnectorConfig field in TableSpecificConfig.
func (m *TableSpecificConfig) SetQueryBasedConnectorConfig(ctx context.Context, v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) {
	vs := v.ToObjectValue(ctx)
	m.QueryBasedConnectorConfig = vs
}

// GetSequenceBy returns the value of the SequenceBy field in TableSpecificConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig) GetSequenceBy(ctx context.Context) ([]types.String, bool) {
	if m.SequenceBy.IsNull() || m.SequenceBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SequenceBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSequenceBy sets the value of the SequenceBy field in TableSpecificConfig.
func (m *TableSpecificConfig) SetSequenceBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SequenceBy = types.ListValueMust(t, vs)
}

// GetWorkdayReportParameters returns the value of the WorkdayReportParameters field in TableSpecificConfig as
// a IngestionPipelineDefinitionWorkdayReportParameters value.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig) GetWorkdayReportParameters(ctx context.Context) (IngestionPipelineDefinitionWorkdayReportParameters, bool) {
	var e IngestionPipelineDefinitionWorkdayReportParameters
	if m.WorkdayReportParameters.IsNull() || m.WorkdayReportParameters.IsUnknown() {
		return e, false
	}
	var v IngestionPipelineDefinitionWorkdayReportParameters
	d := m.WorkdayReportParameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkdayReportParameters sets the value of the WorkdayReportParameters field in TableSpecificConfig.
func (m *TableSpecificConfig) SetWorkdayReportParameters(ctx context.Context, v IngestionPipelineDefinitionWorkdayReportParameters) {
	vs := v.ToObjectValue(ctx)
	m.WorkdayReportParameters = vs
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
	// Indicates whether the update is either part of a continuous job run, or
	// running in legacy continuous pipeline mode.
	Mode types.String `tfsdk:"mode"`
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

func (m UpdateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["full_refresh"] = attrs["full_refresh"].SetOptional()
	attrs["full_refresh_selection"] = attrs["full_refresh_selection"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
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
func (m UpdateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":                 reflect.TypeOf(PipelineSpec{}),
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInfo
// only implements ToObjectValue() and Type().
func (m UpdateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cause":                  m.Cause,
			"cluster_id":             m.ClusterId,
			"config":                 m.Config,
			"creation_time":          m.CreationTime,
			"full_refresh":           m.FullRefresh,
			"full_refresh_selection": m.FullRefreshSelection,
			"mode":                   m.Mode,
			"pipeline_id":            m.PipelineId,
			"refresh_selection":      m.RefreshSelection,
			"state":                  m.State,
			"update_id":              m.UpdateId,
			"validate_only":          m.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateInfo) Type(ctx context.Context) attr.Type {
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
			"mode":        types.StringType,
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
func (m *UpdateInfo) GetConfig(ctx context.Context) (PipelineSpec, bool) {
	var e PipelineSpec
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v PipelineSpec
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in UpdateInfo.
func (m *UpdateInfo) SetConfig(ctx context.Context, v PipelineSpec) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in UpdateInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInfo) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if m.FullRefreshSelection.IsNull() || m.FullRefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FullRefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in UpdateInfo.
func (m *UpdateInfo) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in UpdateInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInfo) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
	if m.RefreshSelection.IsNull() || m.RefreshSelection.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RefreshSelection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshSelection sets the value of the RefreshSelection field in UpdateInfo.
func (m *UpdateInfo) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RefreshSelection = types.ListValueMust(t, vs)
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

func (m UpdateStateInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateStateInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStateInfo
// only implements ToObjectValue() and Type().
func (m UpdateStateInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time": m.CreationTime,
			"state":         m.State,
			"update_id":     m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateStateInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.StringType,
			"state":         types.StringType,
			"update_id":     types.StringType,
		},
	}
}
