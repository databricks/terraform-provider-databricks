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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ApplyEnvironmentRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *ApplyEnvironmentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApplyEnvironmentRequest_SdkV2) {
}

func (to *ApplyEnvironmentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ApplyEnvironmentRequest_SdkV2) {
}

func (m ApplyEnvironmentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ApplyEnvironmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplyEnvironmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ApplyEnvironmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApplyEnvironmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type ApplyEnvironmentRequestResponse_SdkV2 struct {
}

func (to *ApplyEnvironmentRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApplyEnvironmentRequestResponse_SdkV2) {
}

func (to *ApplyEnvironmentRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ApplyEnvironmentRequestResponse_SdkV2) {
}

func (m ApplyEnvironmentRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApplyEnvironmentRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApplyEnvironmentRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplyEnvironmentRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ApplyEnvironmentRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ApplyEnvironmentRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ConnectionParameters_SdkV2 struct {
	// Source catalog for initial connection. This is necessary for schema
	// exploration in some database systems like Oracle, and optional but
	// nice-to-have in some other database systems like Postgres. For Oracle
	// databases, this maps to a service name.
	SourceCatalog types.String `tfsdk:"source_catalog"`
}

func (to *ConnectionParameters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConnectionParameters_SdkV2) {
}

func (to *ConnectionParameters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ConnectionParameters_SdkV2) {
}

func (m ConnectionParameters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ConnectionParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConnectionParameters_SdkV2
// only implements ToObjectValue() and Type().
func (m ConnectionParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_catalog": m.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConnectionParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_catalog": types.StringType,
		},
	}
}

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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *CreatePipeline_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePipeline_SdkV2) {
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

func (to *CreatePipeline_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePipeline_SdkV2) {
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

func (m CreatePipeline_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_duplicate_names"] = attrs["allow_duplicate_names"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["channel"] = attrs["channel"].SetOptional()
	attrs["clusters"] = attrs["clusters"].SetOptional()
	attrs["configuration"] = attrs["configuration"].SetOptional()
	attrs["continuous"] = attrs["continuous"].SetOptional()
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["deployment"] = attrs["deployment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["development"] = attrs["development"].SetOptional()
	attrs["dry_run"] = attrs["dry_run"].SetOptional()
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
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as"] = attrs["run_as"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()
	attrs["trigger"] = attrs["trigger"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreatePipeline_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CreatePipeline_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreatePipeline_SdkV2) Type(ctx context.Context) attr.Type {
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
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in CreatePipeline_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in CreatePipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in CreatePipeline_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := m.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	m.Deployment = types.ListValueMust(t, vs)
}

// GetEnvironment returns the value of the Environment field in CreatePipeline_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := m.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	m.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in CreatePipeline_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := m.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	m.EventLog = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in CreatePipeline_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	m.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in CreatePipeline_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := m.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	m.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in CreatePipeline_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := m.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	m.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in CreatePipeline_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in CreatePipeline_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in CreatePipeline_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := m.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	m.RestartWindow = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in CreatePipeline_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := m.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	m.RunAs = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreatePipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in CreatePipeline_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePipeline_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := m.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in CreatePipeline_SdkV2.
func (m *CreatePipeline_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	m.Trigger = types.ListValueMust(t, vs)
}

type CreatePipelineResponse_SdkV2 struct {
	// Only returned when dry_run is true.
	EffectiveSettings types.List `tfsdk:"effective_settings"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (to *CreatePipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePipelineResponse_SdkV2) {
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

func (to *CreatePipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePipelineResponse_SdkV2) {
	if !from.EffectiveSettings.IsNull() && !from.EffectiveSettings.IsUnknown() {
		if toEffectiveSettings, ok := to.GetEffectiveSettings(ctx); ok {
			if fromEffectiveSettings, ok := from.GetEffectiveSettings(ctx); ok {
				toEffectiveSettings.SyncFieldsDuringRead(ctx, fromEffectiveSettings)
				to.SetEffectiveSettings(ctx, toEffectiveSettings)
			}
		}
	}
}

func (m CreatePipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(PipelineSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_settings": m.EffectiveSettings,
			"pipeline_id":        m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreatePipelineResponse_SdkV2) GetEffectiveSettings(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if m.EffectiveSettings.IsNull() || m.EffectiveSettings.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := m.EffectiveSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveSettings sets the value of the EffectiveSettings field in CreatePipelineResponse_SdkV2.
func (m *CreatePipelineResponse_SdkV2) SetEffectiveSettings(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_settings"]
	m.EffectiveSettings = types.ListValueMust(t, vs)
}

type CronTrigger_SdkV2 struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule"`

	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (to *CronTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CronTrigger_SdkV2) {
}

func (to *CronTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CronTrigger_SdkV2) {
}

func (m CronTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CronTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (m CronTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quartz_cron_schedule": m.QuartzCronSchedule,
			"timezone_id":          m.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CronTrigger_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DataPlaneId_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataPlaneId_SdkV2) {
}

func (to *DataPlaneId_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataPlaneId_SdkV2) {
}

func (m DataPlaneId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataPlaneId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneId_SdkV2
// only implements ToObjectValue() and Type().
func (m DataPlaneId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance": m.Instance,
			"seq_no":   m.SeqNo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataPlaneId_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeletePipelineRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePipelineRequest_SdkV2) {
}

func (to *DeletePipelineRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePipelineRequest_SdkV2) {
}

func (m DeletePipelineRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePipelineRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePipelineRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePipelineRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type DeletePipelineResponse_SdkV2 struct {
}

func (to *DeletePipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePipelineResponse_SdkV2) {
}

func (to *DeletePipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePipelineResponse_SdkV2) {
}

func (m DeletePipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *EditPipeline_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPipeline_SdkV2) {
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

func (to *EditPipeline_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EditPipeline_SdkV2) {
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

func (m EditPipeline_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_duplicate_names"] = attrs["allow_duplicate_names"].SetOptional()
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
	attrs["expected_last_modified"] = attrs["expected_last_modified"].SetOptional()
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
	attrs["run_as"] = attrs["run_as"].SetOptional()
	attrs["run_as"] = attrs["run_as"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["serverless"] = attrs["serverless"].SetOptional()
	attrs["storage"] = attrs["storage"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["target"] = attrs["target"].SetOptional()
	attrs["trigger"] = attrs["trigger"].SetOptional()
	attrs["trigger"] = attrs["trigger"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EditPipeline_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EditPipeline_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EditPipeline_SdkV2) Type(ctx context.Context) attr.Type {
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
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in EditPipeline_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in EditPipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in EditPipeline_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := m.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	m.Deployment = types.ListValueMust(t, vs)
}

// GetEnvironment returns the value of the Environment field in EditPipeline_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := m.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	m.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in EditPipeline_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := m.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	m.EventLog = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in EditPipeline_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	m.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in EditPipeline_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := m.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	m.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in EditPipeline_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := m.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	m.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in EditPipeline_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in EditPipeline_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in EditPipeline_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := m.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	m.RestartWindow = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in EditPipeline_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := m.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	m.RunAs = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in EditPipeline_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in EditPipeline_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EditPipeline_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := m.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in EditPipeline_SdkV2.
func (m *EditPipeline_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	m.Trigger = types.ListValueMust(t, vs)
}

type EditPipelineResponse_SdkV2 struct {
}

func (to *EditPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EditPipelineResponse_SdkV2) {
}

func (to *EditPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EditPipelineResponse_SdkV2) {
}

func (m EditPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EditPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EditPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EditPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m EditPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EditPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ErrorDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ErrorDetail_SdkV2) {
	if !from.Exceptions.IsNull() && !from.Exceptions.IsUnknown() && to.Exceptions.IsNull() && len(from.Exceptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exceptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exceptions = from.Exceptions
	}
}

func (to *ErrorDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ErrorDetail_SdkV2) {
	if !from.Exceptions.IsNull() && !from.Exceptions.IsUnknown() && to.Exceptions.IsNull() && len(from.Exceptions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Exceptions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Exceptions = from.Exceptions
	}
}

func (m ErrorDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ErrorDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exceptions": reflect.TypeOf(SerializedException_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ErrorDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m ErrorDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exceptions": m.Exceptions,
			"fatal":      m.Fatal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ErrorDetail_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ErrorDetail_SdkV2) GetExceptions(ctx context.Context) ([]SerializedException_SdkV2, bool) {
	if m.Exceptions.IsNull() || m.Exceptions.IsUnknown() {
		return nil, false
	}
	var v []SerializedException_SdkV2
	d := m.Exceptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExceptions sets the value of the Exceptions field in ErrorDetail_SdkV2.
func (m *ErrorDetail_SdkV2) SetExceptions(ctx context.Context, v []SerializedException_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exceptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exceptions = types.ListValueMust(t, vs)
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

func (to *EventLogSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EventLogSpec_SdkV2) {
}

func (to *EventLogSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EventLogSpec_SdkV2) {
}

func (m EventLogSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EventLogSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EventLogSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m EventLogSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
			"name":    m.Name,
			"schema":  m.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EventLogSpec_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *FileLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileLibrary_SdkV2) {
}

func (to *FileLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileLibrary_SdkV2) {
}

func (m FileLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FileLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (m FileLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileLibrary_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Filters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Filters_SdkV2) {
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

func (to *Filters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Filters_SdkV2) {
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

func (m Filters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Filters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude": reflect.TypeOf(types.String{}),
		"include": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Filters_SdkV2
// only implements ToObjectValue() and Type().
func (m Filters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"exclude": m.Exclude,
			"include": m.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Filters_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Filters_SdkV2) GetExclude(ctx context.Context) ([]types.String, bool) {
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

// SetExclude sets the value of the Exclude field in Filters_SdkV2.
func (m *Filters_SdkV2) SetExclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Exclude = types.ListValueMust(t, vs)
}

// GetInclude returns the value of the Include field in Filters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Filters_SdkV2) GetInclude(ctx context.Context) ([]types.String, bool) {
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

// SetInclude sets the value of the Include field in Filters_SdkV2.
func (m *Filters_SdkV2) SetInclude(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Include = types.ListValueMust(t, vs)
}

type GetPipelinePermissionLevelsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionLevelsRequest_SdkV2) {
}

func (to *GetPipelinePermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionLevelsRequest_SdkV2) {
}

func (m GetPipelinePermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetPipelinePermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetPipelinePermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetPipelinePermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PipelinePermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetPipelinePermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]PipelinePermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetPipelinePermissionLevelsResponse_SdkV2.
func (m *GetPipelinePermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []PipelinePermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetPipelinePermissionsRequest_SdkV2 struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelinePermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelinePermissionsRequest_SdkV2) {
}

func (to *GetPipelinePermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPipelinePermissionsRequest_SdkV2) {
}

func (m GetPipelinePermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelinePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelinePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPipelinePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelinePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type GetPipelineRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *GetPipelineRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelineRequest_SdkV2) {
}

func (to *GetPipelineRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPipelineRequest_SdkV2) {
}

func (m GetPipelineRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelineRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPipelineRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPipelineRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	RunAs types.List `tfsdk:"run_as"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec types.List `tfsdk:"spec"`
	// The pipeline state.
	State types.String `tfsdk:"state"`
}

func (to *GetPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPipelineResponse_SdkV2) {
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

func (to *GetPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPipelineResponse_SdkV2) {
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

func (m GetPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo_SdkV2{}),
		"run_as":         reflect.TypeOf(RunAs_SdkV2{}),
		"spec":           reflect.TypeOf(PipelineSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetPipelineResponse_SdkV2) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo_SdkV2, bool) {
	if m.LatestUpdates.IsNull() || m.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo_SdkV2
	d := m.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in GetPipelineResponse_SdkV2.
func (m *GetPipelineResponse_SdkV2) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestUpdates = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in GetPipelineResponse_SdkV2 as
// a RunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPipelineResponse_SdkV2) GetRunAs(ctx context.Context) (RunAs_SdkV2, bool) {
	var e RunAs_SdkV2
	if m.RunAs.IsNull() || m.RunAs.IsUnknown() {
		return e, false
	}
	var v []RunAs_SdkV2
	d := m.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in GetPipelineResponse_SdkV2.
func (m *GetPipelineResponse_SdkV2) SetRunAs(ctx context.Context, v RunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	m.RunAs = types.ListValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in GetPipelineResponse_SdkV2 as
// a PipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPipelineResponse_SdkV2) GetSpec(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in GetPipelineResponse_SdkV2.
func (m *GetPipelineResponse_SdkV2) SetSpec(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

type GetUpdateRequest_SdkV2 struct {
	// The ID of the pipeline.
	PipelineId types.String `tfsdk:"-"`
	// The ID of the update.
	UpdateId types.String `tfsdk:"-"`
}

func (to *GetUpdateRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUpdateRequest_SdkV2) {
}

func (to *GetUpdateRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUpdateRequest_SdkV2) {
}

func (m GetUpdateRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetUpdateRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetUpdateRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
			"update_id":   m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUpdateRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetUpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetUpdateResponse_SdkV2) {
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

func (to *GetUpdateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetUpdateResponse_SdkV2) {
	if !from.Update.IsNull() && !from.Update.IsUnknown() {
		if toUpdate, ok := to.GetUpdate(ctx); ok {
			if fromUpdate, ok := from.GetUpdate(ctx); ok {
				toUpdate.SyncFieldsDuringRead(ctx, fromUpdate)
				to.SetUpdate(ctx, toUpdate)
			}
		}
	}
}

func (m GetUpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetUpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"update": reflect.TypeOf(UpdateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetUpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetUpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update": m.Update,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetUpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *GetUpdateResponse_SdkV2) GetUpdate(ctx context.Context) (UpdateInfo_SdkV2, bool) {
	var e UpdateInfo_SdkV2
	if m.Update.IsNull() || m.Update.IsUnknown() {
		return e, false
	}
	var v []UpdateInfo_SdkV2
	d := m.Update.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUpdate sets the value of the Update field in GetUpdateResponse_SdkV2.
func (m *GetUpdateResponse_SdkV2) SetUpdate(ctx context.Context, v UpdateInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["update"]
	m.Update = types.ListValueMust(t, vs)
}

type IngestionConfig_SdkV2 struct {
	// Select a specific source report.
	Report types.List `tfsdk:"report"`
	// Select all tables from a specific source schema.
	Schema types.List `tfsdk:"schema"`
	// Select a specific source table.
	Table types.List `tfsdk:"table"`
}

func (to *IngestionConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionConfig_SdkV2) {
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

func (to *IngestionConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionConfig_SdkV2) {
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

func (m IngestionConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"report": reflect.TypeOf(ReportSpec_SdkV2{}),
		"schema": reflect.TypeOf(SchemaSpec_SdkV2{}),
		"table":  reflect.TypeOf(TableSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"report": m.Report,
			"schema": m.Schema,
			"table":  m.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *IngestionConfig_SdkV2) GetReport(ctx context.Context) (ReportSpec_SdkV2, bool) {
	var e ReportSpec_SdkV2
	if m.Report.IsNull() || m.Report.IsUnknown() {
		return e, false
	}
	var v []ReportSpec_SdkV2
	d := m.Report.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReport sets the value of the Report field in IngestionConfig_SdkV2.
func (m *IngestionConfig_SdkV2) SetReport(ctx context.Context, v ReportSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["report"]
	m.Report = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in IngestionConfig_SdkV2 as
// a SchemaSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionConfig_SdkV2) GetSchema(ctx context.Context) (SchemaSpec_SdkV2, bool) {
	var e SchemaSpec_SdkV2
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return e, false
	}
	var v []SchemaSpec_SdkV2
	d := m.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in IngestionConfig_SdkV2.
func (m *IngestionConfig_SdkV2) SetSchema(ctx context.Context, v SchemaSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	m.Schema = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in IngestionConfig_SdkV2 as
// a TableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionConfig_SdkV2) GetTable(ctx context.Context) (TableSpec_SdkV2, bool) {
	var e TableSpec_SdkV2
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v []TableSpec_SdkV2
	d := m.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in IngestionConfig_SdkV2.
func (m *IngestionConfig_SdkV2) SetTable(ctx context.Context, v TableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	m.Table = types.ListValueMust(t, vs)
}

type IngestionGatewayPipelineDefinition_SdkV2 struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	ConnectionId types.String `tfsdk:"connection_id"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Optional, Internal. Parameters required to establish an initial
	// connection with the source.
	ConnectionParameters types.List `tfsdk:"connection_parameters"`
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

func (to *IngestionGatewayPipelineDefinition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionGatewayPipelineDefinition_SdkV2) {
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

func (to *IngestionGatewayPipelineDefinition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionGatewayPipelineDefinition_SdkV2) {
	if !from.ConnectionParameters.IsNull() && !from.ConnectionParameters.IsUnknown() {
		if toConnectionParameters, ok := to.GetConnectionParameters(ctx); ok {
			if fromConnectionParameters, ok := from.GetConnectionParameters(ctx); ok {
				toConnectionParameters.SyncFieldsDuringRead(ctx, fromConnectionParameters)
				to.SetConnectionParameters(ctx, toConnectionParameters)
			}
		}
	}
}

func (m IngestionGatewayPipelineDefinition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_id"] = attrs["connection_id"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetRequired()
	attrs["connection_parameters"] = attrs["connection_parameters"].SetOptional()
	attrs["connection_parameters"] = attrs["connection_parameters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m IngestionGatewayPipelineDefinition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"connection_parameters": reflect.TypeOf(ConnectionParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionGatewayPipelineDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionGatewayPipelineDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m IngestionGatewayPipelineDefinition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_id":   types.StringType,
			"connection_name": types.StringType,
			"connection_parameters": basetypes.ListType{
				ElemType: ConnectionParameters_SdkV2{}.Type(ctx),
			},
			"gateway_storage_catalog": types.StringType,
			"gateway_storage_name":    types.StringType,
			"gateway_storage_schema":  types.StringType,
		},
	}
}

// GetConnectionParameters returns the value of the ConnectionParameters field in IngestionGatewayPipelineDefinition_SdkV2 as
// a ConnectionParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionGatewayPipelineDefinition_SdkV2) GetConnectionParameters(ctx context.Context) (ConnectionParameters_SdkV2, bool) {
	var e ConnectionParameters_SdkV2
	if m.ConnectionParameters.IsNull() || m.ConnectionParameters.IsUnknown() {
		return e, false
	}
	var v []ConnectionParameters_SdkV2
	d := m.ConnectionParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConnectionParameters sets the value of the ConnectionParameters field in IngestionGatewayPipelineDefinition_SdkV2.
func (m *IngestionGatewayPipelineDefinition_SdkV2) SetConnectionParameters(ctx context.Context, v ConnectionParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["connection_parameters"]
	m.ConnectionParameters = types.ListValueMust(t, vs)
}

type IngestionPipelineDefinition_SdkV2 struct {
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
	TableConfiguration types.List `tfsdk:"table_configuration"`
}

func (to *IngestionPipelineDefinition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinition_SdkV2) {
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

func (to *IngestionPipelineDefinition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinition_SdkV2) {
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

func (m IngestionPipelineDefinition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["ingest_from_uc_foreign_catalog"] = attrs["ingest_from_uc_foreign_catalog"].SetOptional()
	attrs["ingestion_gateway_id"] = attrs["ingestion_gateway_id"].SetOptional()
	attrs["netsuite_jar_path"] = attrs["netsuite_jar_path"].SetOptional()
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
func (m IngestionPipelineDefinition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects":               reflect.TypeOf(IngestionConfig_SdkV2{}),
		"source_configurations": reflect.TypeOf(SourceConfig_SdkV2{}),
		"table_configuration":   reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m IngestionPipelineDefinition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":                types.StringType,
			"ingest_from_uc_foreign_catalog": types.BoolType,
			"ingestion_gateway_id":           types.StringType,
			"netsuite_jar_path":              types.StringType,
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
func (m *IngestionPipelineDefinition_SdkV2) GetObjects(ctx context.Context) ([]IngestionConfig_SdkV2, bool) {
	if m.Objects.IsNull() || m.Objects.IsUnknown() {
		return nil, false
	}
	var v []IngestionConfig_SdkV2
	d := m.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in IngestionPipelineDefinition_SdkV2.
func (m *IngestionPipelineDefinition_SdkV2) SetObjects(ctx context.Context, v []IngestionConfig_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Objects = types.ListValueMust(t, vs)
}

// GetSourceConfigurations returns the value of the SourceConfigurations field in IngestionPipelineDefinition_SdkV2 as
// a slice of SourceConfig_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinition_SdkV2) GetSourceConfigurations(ctx context.Context) ([]SourceConfig_SdkV2, bool) {
	if m.SourceConfigurations.IsNull() || m.SourceConfigurations.IsUnknown() {
		return nil, false
	}
	var v []SourceConfig_SdkV2
	d := m.SourceConfigurations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSourceConfigurations sets the value of the SourceConfigurations field in IngestionPipelineDefinition_SdkV2.
func (m *IngestionPipelineDefinition_SdkV2) SetSourceConfigurations(ctx context.Context, v []SourceConfig_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["source_configurations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SourceConfigurations = types.ListValueMust(t, vs)
}

// GetTableConfiguration returns the value of the TableConfiguration field in IngestionPipelineDefinition_SdkV2 as
// a TableSpecificConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinition_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := m.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in IngestionPipelineDefinition_SdkV2.
func (m *IngestionPipelineDefinition_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	m.TableConfiguration = types.ListValueMust(t, vs)
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

func (to *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
	if !from.CursorColumns.IsNull() && !from.CursorColumns.IsUnknown() && to.CursorColumns.IsNull() && len(from.CursorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CursorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CursorColumns = from.CursorColumns
	}
}

func (to *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
	if !from.CursorColumns.IsNull() && !from.CursorColumns.IsUnknown() && to.CursorColumns.IsNull() && len(from.CursorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CursorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CursorColumns = from.CursorColumns
	}
}

func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cursor_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cursor_columns":                             m.CursorColumns,
			"deletion_condition":                         m.DeletionCondition,
			"hard_deletion_sync_min_interval_in_seconds": m.HardDeletionSyncMinIntervalInSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) GetCursorColumns(ctx context.Context) ([]types.String, bool) {
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

// SetCursorColumns sets the value of the CursorColumns field in IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2.
func (m *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) SetCursorColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cursor_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CursorColumns = types.ListValueMust(t, vs)
}

type IngestionPipelineDefinitionWorkdayReportParameters_SdkV2 struct {
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

func (to *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) {
	if !from.ReportParameters.IsNull() && !from.ReportParameters.IsUnknown() && to.ReportParameters.IsNull() && len(from.ReportParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ReportParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ReportParameters = from.ReportParameters
	}
}

func (to *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) {
	if !from.ReportParameters.IsNull() && !from.ReportParameters.IsUnknown() && to.ReportParameters.IsNull() && len(from.ReportParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ReportParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ReportParameters = from.ReportParameters
	}
}

func (m IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters":        reflect.TypeOf(types.String{}),
		"report_parameters": reflect.TypeOf(IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParameters_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"incremental":       m.Incremental,
			"parameters":        m.Parameters,
			"report_parameters": m.ReportParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"incremental": types.BoolType,
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"report_parameters": basetypes.ListType{
				ElemType: IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in IngestionPipelineDefinitionWorkdayReportParameters_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
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

// SetParameters sets the value of the Parameters field in IngestionPipelineDefinitionWorkdayReportParameters_SdkV2.
func (m *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.MapValueMust(t, vs)
}

// GetReportParameters returns the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters_SdkV2 as
// a slice of IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) GetReportParameters(ctx context.Context) ([]IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2, bool) {
	if m.ReportParameters.IsNull() || m.ReportParameters.IsUnknown() {
		return nil, false
	}
	var v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2
	d := m.ReportParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReportParameters sets the value of the ReportParameters field in IngestionPipelineDefinitionWorkdayReportParameters_SdkV2.
func (m *IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) SetReportParameters(ctx context.Context, v []IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["report_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ReportParameters = types.ListValueMust(t, vs)
}

type IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2 struct {
	// Key for the report parameter, can be a column name or other metadata
	Key types.String `tfsdk:"key"`
	// Value for the report parameter. Possible values it can take are these sql
	// functions: 1. coalesce(current_offset(), date("YYYY-MM-DD")) -> if
	// current_offset() is null, then the passed date, else current_offset() 2.
	// current_date() 3. date_sub(current_date(), x) -> subtract x (some
	// non-negative integer) days from current date
	Value types.String `tfsdk:"value"`
}

func (to *IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) {
}

func (to *IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) {
}

func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2
// only implements ToObjectValue() and Type().
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IngestionPipelineDefinitionWorkdayReportParametersQueryKeyValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
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

func (to *ListPipelineEventsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelineEventsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *ListPipelineEventsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPipelineEventsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m ListPipelineEventsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelineEventsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPipelineEventsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListPipelineEventsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelineEventsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in ListPipelineEventsRequest_SdkV2.
func (m *ListPipelineEventsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelineEventsResponse_SdkV2 struct {
	// The list of events matching the request criteria.
	Events types.List `tfsdk:"events"`
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

func (to *ListPipelineEventsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelineEventsResponse_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (to *ListPipelineEventsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPipelineEventsResponse_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (m ListPipelineEventsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelineEventsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(PipelineEvent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelineEventsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPipelineEventsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":          m.Events,
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelineEventsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelineEventsResponse_SdkV2) GetEvents(ctx context.Context) ([]PipelineEvent_SdkV2, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []PipelineEvent_SdkV2
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListPipelineEventsResponse_SdkV2.
func (m *ListPipelineEventsResponse_SdkV2) SetEvents(ctx context.Context, v []PipelineEvent_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
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

func (to *ListPipelinesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelinesRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *ListPipelinesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPipelinesRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m ListPipelinesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelinesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPipelinesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListPipelinesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelinesRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in ListPipelinesRequest_SdkV2.
func (m *ListPipelinesRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type ListPipelinesResponse_SdkV2 struct {
	// If present, a token to fetch the next page of events.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The list of events matching the request criteria.
	Statuses types.List `tfsdk:"statuses"`
}

func (to *ListPipelinesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPipelinesResponse_SdkV2) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (to *ListPipelinesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPipelinesResponse_SdkV2) {
	if !from.Statuses.IsNull() && !from.Statuses.IsUnknown() && to.Statuses.IsNull() && len(from.Statuses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Statuses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Statuses = from.Statuses
	}
}

func (m ListPipelinesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPipelinesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(PipelineStateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPipelinesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPipelinesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"statuses":        m.Statuses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPipelinesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListPipelinesResponse_SdkV2) GetStatuses(ctx context.Context) ([]PipelineStateInfo_SdkV2, bool) {
	if m.Statuses.IsNull() || m.Statuses.IsUnknown() {
		return nil, false
	}
	var v []PipelineStateInfo_SdkV2
	d := m.Statuses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatuses sets the value of the Statuses field in ListPipelinesResponse_SdkV2.
func (m *ListPipelinesResponse_SdkV2) SetStatuses(ctx context.Context, v []PipelineStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["statuses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Statuses = types.ListValueMust(t, vs)
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

func (to *ListUpdatesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUpdatesRequest_SdkV2) {
}

func (to *ListUpdatesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUpdatesRequest_SdkV2) {
}

func (m ListUpdatesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListUpdatesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListUpdatesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListUpdatesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListUpdatesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListUpdatesResponse_SdkV2) {
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (to *ListUpdatesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListUpdatesResponse_SdkV2) {
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (m ListUpdatesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListUpdatesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(UpdateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListUpdatesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListUpdatesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"prev_page_token": m.PrevPageToken,
			"updates":         m.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListUpdatesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListUpdatesResponse_SdkV2) GetUpdates(ctx context.Context) ([]UpdateInfo_SdkV2, bool) {
	if m.Updates.IsNull() || m.Updates.IsUnknown() {
		return nil, false
	}
	var v []UpdateInfo_SdkV2
	d := m.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in ListUpdatesResponse_SdkV2.
func (m *ListUpdatesResponse_SdkV2) SetUpdates(ctx context.Context, v []UpdateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Updates = types.ListValueMust(t, vs)
}

type ManualTrigger_SdkV2 struct {
}

func (to *ManualTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManualTrigger_SdkV2) {
}

func (to *ManualTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ManualTrigger_SdkV2) {
}

func (m ManualTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManualTrigger.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ManualTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManualTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (m ManualTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ManualTrigger_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type NotebookLibrary_SdkV2 struct {
	// The absolute path of the source code.
	Path types.String `tfsdk:"path"`
}

func (to *NotebookLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotebookLibrary_SdkV2) {
}

func (to *NotebookLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotebookLibrary_SdkV2) {
}

func (m NotebookLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NotebookLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookLibrary_SdkV2
// only implements ToObjectValue() and Type().
func (m NotebookLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotebookLibrary_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *Notifications_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Notifications_SdkV2) {
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

func (to *Notifications_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Notifications_SdkV2) {
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

func (m Notifications_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Notifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":           reflect.TypeOf(types.String{}),
		"email_recipients": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Notifications_SdkV2
// only implements ToObjectValue() and Type().
func (m Notifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alerts":           m.Alerts,
			"email_recipients": m.EmailRecipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Notifications_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Notifications_SdkV2) GetAlerts(ctx context.Context) ([]types.String, bool) {
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

// SetAlerts sets the value of the Alerts field in Notifications_SdkV2.
func (m *Notifications_SdkV2) SetAlerts(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["alerts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Alerts = types.ListValueMust(t, vs)
}

// GetEmailRecipients returns the value of the EmailRecipients field in Notifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Notifications_SdkV2) GetEmailRecipients(ctx context.Context) ([]types.String, bool) {
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

// SetEmailRecipients sets the value of the EmailRecipients field in Notifications_SdkV2.
func (m *Notifications_SdkV2) SetEmailRecipients(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmailRecipients = types.ListValueMust(t, vs)
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

func (to *Origin_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Origin_SdkV2) {
}

func (to *Origin_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Origin_SdkV2) {
}

func (m Origin_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Origin_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Origin_SdkV2
// only implements ToObjectValue() and Type().
func (m Origin_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Origin_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PathPattern_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PathPattern_SdkV2) {
}

func (to *PathPattern_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PathPattern_SdkV2) {
}

func (m PathPattern_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PathPattern_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PathPattern_SdkV2
// only implements ToObjectValue() and Type().
func (m PathPattern_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include": m.Include,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PathPattern_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PipelineAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineAccessControlRequest_SdkV2) {
}

func (to *PipelineAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineAccessControlRequest_SdkV2) {
}

func (m PipelineAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PipelineAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *PipelineAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m PipelineAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PipelinePermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]PipelinePermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []PipelinePermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in PipelineAccessControlResponse_SdkV2.
func (m *PipelineAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []PipelinePermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
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

func (to *PipelineCluster_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineCluster_SdkV2) {
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

func (to *PipelineCluster_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineCluster_SdkV2) {
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

func (m PipelineCluster_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineCluster_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineCluster_SdkV2) GetAutoscale(ctx context.Context) (PipelineClusterAutoscale_SdkV2, bool) {
	var e PipelineClusterAutoscale_SdkV2
	if m.Autoscale.IsNull() || m.Autoscale.IsUnknown() {
		return e, false
	}
	var v []PipelineClusterAutoscale_SdkV2
	d := m.Autoscale.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoscale sets the value of the Autoscale field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetAutoscale(ctx context.Context, v PipelineClusterAutoscale_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["autoscale"]
	m.Autoscale = types.ListValueMust(t, vs)
}

// GetAwsAttributes returns the value of the AwsAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.AwsAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetAwsAttributes(ctx context.Context) (compute_tf.AwsAttributes_SdkV2, bool) {
	var e compute_tf.AwsAttributes_SdkV2
	if m.AwsAttributes.IsNull() || m.AwsAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AwsAttributes_SdkV2
	d := m.AwsAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsAttributes sets the value of the AwsAttributes field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetAwsAttributes(ctx context.Context, v compute_tf.AwsAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_attributes"]
	m.AwsAttributes = types.ListValueMust(t, vs)
}

// GetAzureAttributes returns the value of the AzureAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.AzureAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetAzureAttributes(ctx context.Context) (compute_tf.AzureAttributes_SdkV2, bool) {
	var e compute_tf.AzureAttributes_SdkV2
	if m.AzureAttributes.IsNull() || m.AzureAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.AzureAttributes_SdkV2
	d := m.AzureAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAttributes sets the value of the AzureAttributes field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetAzureAttributes(ctx context.Context, v compute_tf.AzureAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_attributes"]
	m.AzureAttributes = types.ListValueMust(t, vs)
}

// GetClusterLogConf returns the value of the ClusterLogConf field in PipelineCluster_SdkV2 as
// a compute_tf.ClusterLogConf_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetClusterLogConf(ctx context.Context) (compute_tf.ClusterLogConf_SdkV2, bool) {
	var e compute_tf.ClusterLogConf_SdkV2
	if m.ClusterLogConf.IsNull() || m.ClusterLogConf.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterLogConf_SdkV2
	d := m.ClusterLogConf.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterLogConf sets the value of the ClusterLogConf field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetClusterLogConf(ctx context.Context, v compute_tf.ClusterLogConf_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_log_conf"]
	m.ClusterLogConf = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpAttributes returns the value of the GcpAttributes field in PipelineCluster_SdkV2 as
// a compute_tf.GcpAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetGcpAttributes(ctx context.Context) (compute_tf.GcpAttributes_SdkV2, bool) {
	var e compute_tf.GcpAttributes_SdkV2
	if m.GcpAttributes.IsNull() || m.GcpAttributes.IsUnknown() {
		return e, false
	}
	var v []compute_tf.GcpAttributes_SdkV2
	d := m.GcpAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpAttributes sets the value of the GcpAttributes field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetGcpAttributes(ctx context.Context, v compute_tf.GcpAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_attributes"]
	m.GcpAttributes = types.ListValueMust(t, vs)
}

// GetInitScripts returns the value of the InitScripts field in PipelineCluster_SdkV2 as
// a slice of compute_tf.InitScriptInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetInitScripts(ctx context.Context) ([]compute_tf.InitScriptInfo_SdkV2, bool) {
	if m.InitScripts.IsNull() || m.InitScripts.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.InitScriptInfo_SdkV2
	d := m.InitScripts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitScripts sets the value of the InitScripts field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetInitScripts(ctx context.Context, v []compute_tf.InitScriptInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["init_scripts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InitScripts = types.ListValueMust(t, vs)
}

// GetSparkConf returns the value of the SparkConf field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetSparkConf(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkConf sets the value of the SparkConf field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetSparkConf(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_conf"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkConf = types.MapValueMust(t, vs)
}

// GetSparkEnvVars returns the value of the SparkEnvVars field in PipelineCluster_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetSparkEnvVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetSparkEnvVars sets the value of the SparkEnvVars field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetSparkEnvVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SparkEnvVars = types.MapValueMust(t, vs)
}

// GetSshPublicKeys returns the value of the SshPublicKeys field in PipelineCluster_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineCluster_SdkV2) GetSshPublicKeys(ctx context.Context) ([]types.String, bool) {
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

// SetSshPublicKeys sets the value of the SshPublicKeys field in PipelineCluster_SdkV2.
func (m *PipelineCluster_SdkV2) SetSshPublicKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ssh_public_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SshPublicKeys = types.ListValueMust(t, vs)
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

func (to *PipelineClusterAutoscale_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineClusterAutoscale_SdkV2) {
}

func (to *PipelineClusterAutoscale_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineClusterAutoscale_SdkV2) {
}

func (m PipelineClusterAutoscale_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineClusterAutoscale_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineClusterAutoscale_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineClusterAutoscale_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_workers": m.MaxWorkers,
			"min_workers": m.MinWorkers,
			"mode":        m.Mode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineClusterAutoscale_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PipelineDeployment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineDeployment_SdkV2) {
}

func (to *PipelineDeployment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineDeployment_SdkV2) {
}

func (m PipelineDeployment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineDeployment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineDeployment_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineDeployment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kind":               m.Kind,
			"metadata_file_path": m.MetadataFilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineDeployment_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PipelineEvent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineEvent_SdkV2) {
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

func (to *PipelineEvent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineEvent_SdkV2) {
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

func (m PipelineEvent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineEvent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error":    reflect.TypeOf(ErrorDetail_SdkV2{}),
		"origin":   reflect.TypeOf(Origin_SdkV2{}),
		"sequence": reflect.TypeOf(Sequencing_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineEvent_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineEvent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineEvent_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineEvent_SdkV2) GetError(ctx context.Context) (ErrorDetail_SdkV2, bool) {
	var e ErrorDetail_SdkV2
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v []ErrorDetail_SdkV2
	d := m.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in PipelineEvent_SdkV2.
func (m *PipelineEvent_SdkV2) SetError(ctx context.Context, v ErrorDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	m.Error = types.ListValueMust(t, vs)
}

// GetOrigin returns the value of the Origin field in PipelineEvent_SdkV2 as
// a Origin_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineEvent_SdkV2) GetOrigin(ctx context.Context) (Origin_SdkV2, bool) {
	var e Origin_SdkV2
	if m.Origin.IsNull() || m.Origin.IsUnknown() {
		return e, false
	}
	var v []Origin_SdkV2
	d := m.Origin.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOrigin sets the value of the Origin field in PipelineEvent_SdkV2.
func (m *PipelineEvent_SdkV2) SetOrigin(ctx context.Context, v Origin_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["origin"]
	m.Origin = types.ListValueMust(t, vs)
}

// GetSequence returns the value of the Sequence field in PipelineEvent_SdkV2 as
// a Sequencing_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineEvent_SdkV2) GetSequence(ctx context.Context) (Sequencing_SdkV2, bool) {
	var e Sequencing_SdkV2
	if m.Sequence.IsNull() || m.Sequence.IsUnknown() {
		return e, false
	}
	var v []Sequencing_SdkV2
	d := m.Sequence.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSequence sets the value of the Sequence field in PipelineEvent_SdkV2.
func (m *PipelineEvent_SdkV2) SetSequence(ctx context.Context, v Sequencing_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence"]
	m.Sequence = types.ListValueMust(t, vs)
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

func (to *PipelineLibrary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineLibrary_SdkV2) {
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

func (to *PipelineLibrary_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineLibrary_SdkV2) {
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

func (m PipelineLibrary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineLibrary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineLibrary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineLibrary_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineLibrary_SdkV2) GetFile(ctx context.Context) (FileLibrary_SdkV2, bool) {
	var e FileLibrary_SdkV2
	if m.File.IsNull() || m.File.IsUnknown() {
		return e, false
	}
	var v []FileLibrary_SdkV2
	d := m.File.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in PipelineLibrary_SdkV2.
func (m *PipelineLibrary_SdkV2) SetFile(ctx context.Context, v FileLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file"]
	m.File = types.ListValueMust(t, vs)
}

// GetGlob returns the value of the Glob field in PipelineLibrary_SdkV2 as
// a PathPattern_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary_SdkV2) GetGlob(ctx context.Context) (PathPattern_SdkV2, bool) {
	var e PathPattern_SdkV2
	if m.Glob.IsNull() || m.Glob.IsUnknown() {
		return e, false
	}
	var v []PathPattern_SdkV2
	d := m.Glob.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGlob sets the value of the Glob field in PipelineLibrary_SdkV2.
func (m *PipelineLibrary_SdkV2) SetGlob(ctx context.Context, v PathPattern_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["glob"]
	m.Glob = types.ListValueMust(t, vs)
}

// GetMaven returns the value of the Maven field in PipelineLibrary_SdkV2 as
// a compute_tf.MavenLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary_SdkV2) GetMaven(ctx context.Context) (compute_tf.MavenLibrary_SdkV2, bool) {
	var e compute_tf.MavenLibrary_SdkV2
	if m.Maven.IsNull() || m.Maven.IsUnknown() {
		return e, false
	}
	var v []compute_tf.MavenLibrary_SdkV2
	d := m.Maven.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaven sets the value of the Maven field in PipelineLibrary_SdkV2.
func (m *PipelineLibrary_SdkV2) SetMaven(ctx context.Context, v compute_tf.MavenLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["maven"]
	m.Maven = types.ListValueMust(t, vs)
}

// GetNotebook returns the value of the Notebook field in PipelineLibrary_SdkV2 as
// a NotebookLibrary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineLibrary_SdkV2) GetNotebook(ctx context.Context) (NotebookLibrary_SdkV2, bool) {
	var e NotebookLibrary_SdkV2
	if m.Notebook.IsNull() || m.Notebook.IsUnknown() {
		return e, false
	}
	var v []NotebookLibrary_SdkV2
	d := m.Notebook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebook sets the value of the Notebook field in PipelineLibrary_SdkV2.
func (m *PipelineLibrary_SdkV2) SetNotebook(ctx context.Context, v NotebookLibrary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook"]
	m.Notebook = types.ListValueMust(t, vs)
}

type PipelinePermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *PipelinePermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *PipelinePermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelinePermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m PipelinePermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermission_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelinePermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in PipelinePermission_SdkV2.
func (m *PipelinePermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type PipelinePermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *PipelinePermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *PipelinePermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m PipelinePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelinePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissions_SdkV2.
func (m *PipelinePermissions_SdkV2) SetAccessControlList(ctx context.Context, v []PipelineAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type PipelinePermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *PipelinePermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissionsDescription_SdkV2) {
}

func (to *PipelinePermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissionsDescription_SdkV2) {
}

func (m PipelinePermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelinePermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *PipelinePermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinePermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *PipelinePermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelinePermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m PipelinePermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelinePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"pipeline_id":         m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelinePermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]PipelineAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []PipelineAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in PipelinePermissionsRequest_SdkV2.
func (m *PipelinePermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []PipelineAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
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
	// Usage policy of this pipeline.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *PipelineSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineSpec_SdkV2) {
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

func (to *PipelineSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineSpec_SdkV2) {
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

func (m PipelineSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m PipelineSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
			"usage_policy_id": types.StringType,
		},
	}
}

// GetClusters returns the value of the Clusters field in PipelineSpec_SdkV2 as
// a slice of PipelineCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetClusters(ctx context.Context) ([]PipelineCluster_SdkV2, bool) {
	if m.Clusters.IsNull() || m.Clusters.IsUnknown() {
		return nil, false
	}
	var v []PipelineCluster_SdkV2
	d := m.Clusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClusters sets the value of the Clusters field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetClusters(ctx context.Context, v []PipelineCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Clusters = types.ListValueMust(t, vs)
}

// GetConfiguration returns the value of the Configuration field in PipelineSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetConfiguration(ctx context.Context) (map[string]types.String, bool) {
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

// SetConfiguration sets the value of the Configuration field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetConfiguration(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["configuration"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Configuration = types.MapValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in PipelineSpec_SdkV2 as
// a PipelineDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetDeployment(ctx context.Context) (PipelineDeployment_SdkV2, bool) {
	var e PipelineDeployment_SdkV2
	if m.Deployment.IsNull() || m.Deployment.IsUnknown() {
		return e, false
	}
	var v []PipelineDeployment_SdkV2
	d := m.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetDeployment(ctx context.Context, v PipelineDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	m.Deployment = types.ListValueMust(t, vs)
}

// GetEnvironment returns the value of the Environment field in PipelineSpec_SdkV2 as
// a PipelinesEnvironment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetEnvironment(ctx context.Context) (PipelinesEnvironment_SdkV2, bool) {
	var e PipelinesEnvironment_SdkV2
	if m.Environment.IsNull() || m.Environment.IsUnknown() {
		return e, false
	}
	var v []PipelinesEnvironment_SdkV2
	d := m.Environment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnvironment sets the value of the Environment field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetEnvironment(ctx context.Context, v PipelinesEnvironment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	m.Environment = types.ListValueMust(t, vs)
}

// GetEventLog returns the value of the EventLog field in PipelineSpec_SdkV2 as
// a EventLogSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetEventLog(ctx context.Context) (EventLogSpec_SdkV2, bool) {
	var e EventLogSpec_SdkV2
	if m.EventLog.IsNull() || m.EventLog.IsUnknown() {
		return e, false
	}
	var v []EventLogSpec_SdkV2
	d := m.EventLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEventLog sets the value of the EventLog field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetEventLog(ctx context.Context, v EventLogSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["event_log"]
	m.EventLog = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in PipelineSpec_SdkV2 as
// a Filters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetFilters(ctx context.Context) (Filters_SdkV2, bool) {
	var e Filters_SdkV2
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return e, false
	}
	var v []Filters_SdkV2
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFilters sets the value of the Filters field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetFilters(ctx context.Context, v Filters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	m.Filters = types.ListValueMust(t, vs)
}

// GetGatewayDefinition returns the value of the GatewayDefinition field in PipelineSpec_SdkV2 as
// a IngestionGatewayPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetGatewayDefinition(ctx context.Context) (IngestionGatewayPipelineDefinition_SdkV2, bool) {
	var e IngestionGatewayPipelineDefinition_SdkV2
	if m.GatewayDefinition.IsNull() || m.GatewayDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionGatewayPipelineDefinition_SdkV2
	d := m.GatewayDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGatewayDefinition sets the value of the GatewayDefinition field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetGatewayDefinition(ctx context.Context, v IngestionGatewayPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_definition"]
	m.GatewayDefinition = types.ListValueMust(t, vs)
}

// GetIngestionDefinition returns the value of the IngestionDefinition field in PipelineSpec_SdkV2 as
// a IngestionPipelineDefinition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetIngestionDefinition(ctx context.Context) (IngestionPipelineDefinition_SdkV2, bool) {
	var e IngestionPipelineDefinition_SdkV2
	if m.IngestionDefinition.IsNull() || m.IngestionDefinition.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinition_SdkV2
	d := m.IngestionDefinition.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIngestionDefinition sets the value of the IngestionDefinition field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetIngestionDefinition(ctx context.Context, v IngestionPipelineDefinition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ingestion_definition"]
	m.IngestionDefinition = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in PipelineSpec_SdkV2 as
// a slice of PipelineLibrary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetLibraries(ctx context.Context) ([]PipelineLibrary_SdkV2, bool) {
	if m.Libraries.IsNull() || m.Libraries.IsUnknown() {
		return nil, false
	}
	var v []PipelineLibrary_SdkV2
	d := m.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetLibraries(ctx context.Context, v []PipelineLibrary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["libraries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Libraries = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in PipelineSpec_SdkV2 as
// a slice of Notifications_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetNotifications(ctx context.Context) ([]Notifications_SdkV2, bool) {
	if m.Notifications.IsNull() || m.Notifications.IsUnknown() {
		return nil, false
	}
	var v []Notifications_SdkV2
	d := m.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotifications sets the value of the Notifications field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetNotifications(ctx context.Context, v []Notifications_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notifications = types.ListValueMust(t, vs)
}

// GetRestartWindow returns the value of the RestartWindow field in PipelineSpec_SdkV2 as
// a RestartWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetRestartWindow(ctx context.Context) (RestartWindow_SdkV2, bool) {
	var e RestartWindow_SdkV2
	if m.RestartWindow.IsNull() || m.RestartWindow.IsUnknown() {
		return e, false
	}
	var v []RestartWindow_SdkV2
	d := m.RestartWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestartWindow sets the value of the RestartWindow field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetRestartWindow(ctx context.Context, v RestartWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["restart_window"]
	m.RestartWindow = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in PipelineSpec_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetTags sets the value of the Tags field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.MapValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in PipelineSpec_SdkV2 as
// a PipelineTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineSpec_SdkV2) GetTrigger(ctx context.Context) (PipelineTrigger_SdkV2, bool) {
	var e PipelineTrigger_SdkV2
	if m.Trigger.IsNull() || m.Trigger.IsUnknown() {
		return e, false
	}
	var v []PipelineTrigger_SdkV2
	d := m.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in PipelineSpec_SdkV2.
func (m *PipelineSpec_SdkV2) SetTrigger(ctx context.Context, v PipelineTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	m.Trigger = types.ListValueMust(t, vs)
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

func (to *PipelineStateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineStateInfo_SdkV2) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
}

func (to *PipelineStateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineStateInfo_SdkV2) {
	if !from.LatestUpdates.IsNull() && !from.LatestUpdates.IsUnknown() && to.LatestUpdates.IsNull() && len(from.LatestUpdates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestUpdates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestUpdates = from.LatestUpdates
	}
}

func (m PipelineStateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineStateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineStateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineStateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PipelineStateInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineStateInfo_SdkV2) GetLatestUpdates(ctx context.Context) ([]UpdateStateInfo_SdkV2, bool) {
	if m.LatestUpdates.IsNull() || m.LatestUpdates.IsUnknown() {
		return nil, false
	}
	var v []UpdateStateInfo_SdkV2
	d := m.LatestUpdates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestUpdates sets the value of the LatestUpdates field in PipelineStateInfo_SdkV2.
func (m *PipelineStateInfo_SdkV2) SetLatestUpdates(ctx context.Context, v []UpdateStateInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestUpdates = types.ListValueMust(t, vs)
}

type PipelineTrigger_SdkV2 struct {
	Cron types.List `tfsdk:"cron"`

	Manual types.List `tfsdk:"manual"`
}

func (to *PipelineTrigger_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelineTrigger_SdkV2) {
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

func (to *PipelineTrigger_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelineTrigger_SdkV2) {
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

func (m PipelineTrigger_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelineTrigger_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron":   reflect.TypeOf(CronTrigger_SdkV2{}),
		"manual": reflect.TypeOf(ManualTrigger_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineTrigger_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelineTrigger_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cron":   m.Cron,
			"manual": m.Manual,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelineTrigger_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelineTrigger_SdkV2) GetCron(ctx context.Context) (CronTrigger_SdkV2, bool) {
	var e CronTrigger_SdkV2
	if m.Cron.IsNull() || m.Cron.IsUnknown() {
		return e, false
	}
	var v []CronTrigger_SdkV2
	d := m.Cron.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCron sets the value of the Cron field in PipelineTrigger_SdkV2.
func (m *PipelineTrigger_SdkV2) SetCron(ctx context.Context, v CronTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cron"]
	m.Cron = types.ListValueMust(t, vs)
}

// GetManual returns the value of the Manual field in PipelineTrigger_SdkV2 as
// a ManualTrigger_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PipelineTrigger_SdkV2) GetManual(ctx context.Context) (ManualTrigger_SdkV2, bool) {
	var e ManualTrigger_SdkV2
	if m.Manual.IsNull() || m.Manual.IsUnknown() {
		return e, false
	}
	var v []ManualTrigger_SdkV2
	d := m.Manual.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManual sets the value of the Manual field in PipelineTrigger_SdkV2.
func (m *PipelineTrigger_SdkV2) SetManual(ctx context.Context, v ManualTrigger_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["manual"]
	m.Manual = types.ListValueMust(t, vs)
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

func (to *PipelinesEnvironment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PipelinesEnvironment_SdkV2) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (to *PipelinesEnvironment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PipelinesEnvironment_SdkV2) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (m PipelinesEnvironment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PipelinesEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelinesEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (m PipelinesEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": m.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PipelinesEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PipelinesEnvironment_SdkV2) GetDependencies(ctx context.Context) ([]types.String, bool) {
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

// SetDependencies sets the value of the Dependencies field in PipelinesEnvironment_SdkV2.
func (m *PipelinesEnvironment_SdkV2) SetDependencies(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Dependencies = types.ListValueMust(t, vs)
}

// PG-specific catalog-level configuration parameters
type PostgresCatalogConfig_SdkV2 struct {
	// Optional. The Postgres slot configuration to use for logical replication
	SlotConfig types.List `tfsdk:"slot_config"`
}

func (to *PostgresCatalogConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PostgresCatalogConfig_SdkV2) {
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

func (to *PostgresCatalogConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PostgresCatalogConfig_SdkV2) {
	if !from.SlotConfig.IsNull() && !from.SlotConfig.IsUnknown() {
		if toSlotConfig, ok := to.GetSlotConfig(ctx); ok {
			if fromSlotConfig, ok := from.GetSlotConfig(ctx); ok {
				toSlotConfig.SyncFieldsDuringRead(ctx, fromSlotConfig)
				to.SetSlotConfig(ctx, toSlotConfig)
			}
		}
	}
}

func (m PostgresCatalogConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PostgresCatalogConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"slot_config": reflect.TypeOf(PostgresSlotConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresCatalogConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m PostgresCatalogConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"slot_config": m.SlotConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PostgresCatalogConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *PostgresCatalogConfig_SdkV2) GetSlotConfig(ctx context.Context) (PostgresSlotConfig_SdkV2, bool) {
	var e PostgresSlotConfig_SdkV2
	if m.SlotConfig.IsNull() || m.SlotConfig.IsUnknown() {
		return e, false
	}
	var v []PostgresSlotConfig_SdkV2
	d := m.SlotConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSlotConfig sets the value of the SlotConfig field in PostgresCatalogConfig_SdkV2.
func (m *PostgresCatalogConfig_SdkV2) SetSlotConfig(ctx context.Context, v PostgresSlotConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["slot_config"]
	m.SlotConfig = types.ListValueMust(t, vs)
}

// PostgresSlotConfig contains the configuration for a Postgres logical
// replication slot
type PostgresSlotConfig_SdkV2 struct {
	// The name of the publication to use for the Postgres source
	PublicationName types.String `tfsdk:"publication_name"`
	// The name of the logical replication slot to use for the Postgres source
	SlotName types.String `tfsdk:"slot_name"`
}

func (to *PostgresSlotConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PostgresSlotConfig_SdkV2) {
}

func (to *PostgresSlotConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PostgresSlotConfig_SdkV2) {
}

func (m PostgresSlotConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PostgresSlotConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PostgresSlotConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m PostgresSlotConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publication_name": m.PublicationName,
			"slot_name":        m.SlotName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PostgresSlotConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ReportSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReportSpec_SdkV2) {
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

func (to *ReportSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ReportSpec_SdkV2) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (m ReportSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ReportSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReportSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m ReportSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ReportSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ReportSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := m.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in ReportSpec_SdkV2.
func (m *ReportSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	m.TableConfiguration = types.ListValueMust(t, vs)
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

func (to *RestartWindow_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestartWindow_SdkV2) {
	if !from.DaysOfWeek.IsNull() && !from.DaysOfWeek.IsUnknown() && to.DaysOfWeek.IsNull() && len(from.DaysOfWeek.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DaysOfWeek, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DaysOfWeek = from.DaysOfWeek
	}
}

func (to *RestartWindow_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestartWindow_SdkV2) {
	if !from.DaysOfWeek.IsNull() && !from.DaysOfWeek.IsUnknown() && to.DaysOfWeek.IsNull() && len(from.DaysOfWeek.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DaysOfWeek, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DaysOfWeek = from.DaysOfWeek
	}
}

func (m RestartWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestartWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"days_of_week": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartWindow_SdkV2
// only implements ToObjectValue() and Type().
func (m RestartWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"days_of_week": m.DaysOfWeek,
			"start_hour":   m.StartHour,
			"time_zone_id": m.TimeZoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestartWindow_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *RestartWindow_SdkV2) GetDaysOfWeek(ctx context.Context) ([]types.String, bool) {
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

// SetDaysOfWeek sets the value of the DaysOfWeek field in RestartWindow_SdkV2.
func (m *RestartWindow_SdkV2) SetDaysOfWeek(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["days_of_week"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DaysOfWeek = types.ListValueMust(t, vs)
}

type RestorePipelineRequest_SdkV2 struct {
	// The ID of the pipeline to restore
	PipelineId types.String `tfsdk:"-"`
}

func (to *RestorePipelineRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestorePipelineRequest_SdkV2) {
}

func (to *RestorePipelineRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestorePipelineRequest_SdkV2) {
}

func (m RestorePipelineRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestorePipelineRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestorePipelineRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RestorePipelineRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestorePipelineRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
		},
	}
}

type RestorePipelineRequestResponse_SdkV2 struct {
}

func (to *RestorePipelineRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestorePipelineRequestResponse_SdkV2) {
}

func (to *RestorePipelineRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestorePipelineRequestResponse_SdkV2) {
}

func (m RestorePipelineRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestorePipelineRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestorePipelineRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestorePipelineRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RestorePipelineRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestorePipelineRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
type RunAs_SdkV2 struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName types.String `tfsdk:"user_name"`
}

func (to *RunAs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunAs_SdkV2) {
}

func (to *RunAs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RunAs_SdkV2) {
}

func (m RunAs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunAs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunAs_SdkV2
// only implements ToObjectValue() and Type().
func (m RunAs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunAs_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *SchemaSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SchemaSpec_SdkV2) {
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

func (to *SchemaSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SchemaSpec_SdkV2) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (m SchemaSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SchemaSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m SchemaSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SchemaSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *SchemaSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := m.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in SchemaSpec_SdkV2.
func (m *SchemaSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	m.TableConfiguration = types.ListValueMust(t, vs)
}

type Sequencing_SdkV2 struct {
	// A sequence number, unique and increasing per pipeline.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no"`
	// the ID assigned by the data plane.
	DataPlaneId types.List `tfsdk:"data_plane_id"`
}

func (to *Sequencing_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Sequencing_SdkV2) {
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

func (to *Sequencing_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Sequencing_SdkV2) {
	if !from.DataPlaneId.IsNull() && !from.DataPlaneId.IsUnknown() {
		if toDataPlaneId, ok := to.GetDataPlaneId(ctx); ok {
			if fromDataPlaneId, ok := from.GetDataPlaneId(ctx); ok {
				toDataPlaneId.SyncFieldsDuringRead(ctx, fromDataPlaneId)
				to.SetDataPlaneId(ctx, toDataPlaneId)
			}
		}
	}
}

func (m Sequencing_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Sequencing_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_id": reflect.TypeOf(DataPlaneId_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sequencing_SdkV2
// only implements ToObjectValue() and Type().
func (m Sequencing_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"control_plane_seq_no": m.ControlPlaneSeqNo,
			"data_plane_id":        m.DataPlaneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Sequencing_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *Sequencing_SdkV2) GetDataPlaneId(ctx context.Context) (DataPlaneId_SdkV2, bool) {
	var e DataPlaneId_SdkV2
	if m.DataPlaneId.IsNull() || m.DataPlaneId.IsUnknown() {
		return e, false
	}
	var v []DataPlaneId_SdkV2
	d := m.DataPlaneId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneId sets the value of the DataPlaneId field in Sequencing_SdkV2.
func (m *Sequencing_SdkV2) SetDataPlaneId(ctx context.Context, v DataPlaneId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_plane_id"]
	m.DataPlaneId = types.ListValueMust(t, vs)
}

type SerializedException_SdkV2 struct {
	// Runtime class of the exception
	ClassName types.String `tfsdk:"class_name"`
	// Exception message
	Message types.String `tfsdk:"message"`
	// Stack trace consisting of a list of stack frames
	Stack types.List `tfsdk:"stack"`
}

func (to *SerializedException_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SerializedException_SdkV2) {
	if !from.Stack.IsNull() && !from.Stack.IsUnknown() && to.Stack.IsNull() && len(from.Stack.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stack, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stack = from.Stack
	}
}

func (to *SerializedException_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SerializedException_SdkV2) {
	if !from.Stack.IsNull() && !from.Stack.IsUnknown() && to.Stack.IsNull() && len(from.Stack.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stack, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stack = from.Stack
	}
}

func (m SerializedException_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SerializedException_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stack": reflect.TypeOf(StackFrame_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SerializedException_SdkV2
// only implements ToObjectValue() and Type().
func (m SerializedException_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"class_name": m.ClassName,
			"message":    m.Message,
			"stack":      m.Stack,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SerializedException_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *SerializedException_SdkV2) GetStack(ctx context.Context) ([]StackFrame_SdkV2, bool) {
	if m.Stack.IsNull() || m.Stack.IsUnknown() {
		return nil, false
	}
	var v []StackFrame_SdkV2
	d := m.Stack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStack sets the value of the Stack field in SerializedException_SdkV2.
func (m *SerializedException_SdkV2) SetStack(ctx context.Context, v []StackFrame_SdkV2) {
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
type SourceCatalogConfig_SdkV2 struct {
	// Postgres-specific catalog-level configuration parameters
	Postgres types.List `tfsdk:"postgres"`
	// Source catalog name
	SourceCatalog types.String `tfsdk:"source_catalog"`
}

func (to *SourceCatalogConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SourceCatalogConfig_SdkV2) {
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

func (to *SourceCatalogConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SourceCatalogConfig_SdkV2) {
	if !from.Postgres.IsNull() && !from.Postgres.IsUnknown() {
		if toPostgres, ok := to.GetPostgres(ctx); ok {
			if fromPostgres, ok := from.GetPostgres(ctx); ok {
				toPostgres.SyncFieldsDuringRead(ctx, fromPostgres)
				to.SetPostgres(ctx, toPostgres)
			}
		}
	}
}

func (m SourceCatalogConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SourceCatalogConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"postgres": reflect.TypeOf(PostgresCatalogConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceCatalogConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m SourceCatalogConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"postgres":       m.Postgres,
			"source_catalog": m.SourceCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SourceCatalogConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *SourceCatalogConfig_SdkV2) GetPostgres(ctx context.Context) (PostgresCatalogConfig_SdkV2, bool) {
	var e PostgresCatalogConfig_SdkV2
	if m.Postgres.IsNull() || m.Postgres.IsUnknown() {
		return e, false
	}
	var v []PostgresCatalogConfig_SdkV2
	d := m.Postgres.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPostgres sets the value of the Postgres field in SourceCatalogConfig_SdkV2.
func (m *SourceCatalogConfig_SdkV2) SetPostgres(ctx context.Context, v PostgresCatalogConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["postgres"]
	m.Postgres = types.ListValueMust(t, vs)
}

type SourceConfig_SdkV2 struct {
	// Catalog-level source configuration parameters
	Catalog types.List `tfsdk:"catalog"`
}

func (to *SourceConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SourceConfig_SdkV2) {
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

func (to *SourceConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SourceConfig_SdkV2) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				toCatalog.SyncFieldsDuringRead(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (m SourceConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SourceConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(SourceCatalogConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m SourceConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SourceConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *SourceConfig_SdkV2) GetCatalog(ctx context.Context) (SourceCatalogConfig_SdkV2, bool) {
	var e SourceCatalogConfig_SdkV2
	if m.Catalog.IsNull() || m.Catalog.IsUnknown() {
		return e, false
	}
	var v []SourceCatalogConfig_SdkV2
	d := m.Catalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCatalog sets the value of the Catalog field in SourceConfig_SdkV2.
func (m *SourceConfig_SdkV2) SetCatalog(ctx context.Context, v SourceCatalogConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["catalog"]
	m.Catalog = types.ListValueMust(t, vs)
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

func (to *StackFrame_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StackFrame_SdkV2) {
}

func (to *StackFrame_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StackFrame_SdkV2) {
}

func (m StackFrame_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StackFrame_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StackFrame_SdkV2
// only implements ToObjectValue() and Type().
func (m StackFrame_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m StackFrame_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *StartUpdate_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartUpdate_SdkV2) {
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

func (to *StartUpdate_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StartUpdate_SdkV2) {
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

func (m StartUpdate_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (m StartUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m StartUpdate_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *StartUpdate_SdkV2) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in StartUpdate_SdkV2.
func (m *StartUpdate_SdkV2) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in StartUpdate_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *StartUpdate_SdkV2) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetRefreshSelection sets the value of the RefreshSelection field in StartUpdate_SdkV2.
func (m *StartUpdate_SdkV2) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RefreshSelection = types.ListValueMust(t, vs)
}

type StartUpdateResponse_SdkV2 struct {
	UpdateId types.String `tfsdk:"update_id"`
}

func (to *StartUpdateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartUpdateResponse_SdkV2) {
}

func (to *StartUpdateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StartUpdateResponse_SdkV2) {
}

func (m StartUpdateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartUpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartUpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StartUpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"update_id": m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartUpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update_id": types.StringType,
		},
	}
}

type StopPipelineResponse_SdkV2 struct {
}

func (to *StopPipelineResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopPipelineResponse_SdkV2) {
}

func (to *StopPipelineResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StopPipelineResponse_SdkV2) {
}

func (m StopPipelineResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopPipelineResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StopPipelineResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopPipelineResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StopPipelineResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m StopPipelineResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type StopRequest_SdkV2 struct {
	PipelineId types.String `tfsdk:"-"`
}

func (to *StopRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopRequest_SdkV2) {
}

func (to *StopRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StopRequest_SdkV2) {
}

func (m StopRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StopRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m StopRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *TableSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableSpec_SdkV2) {
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

func (to *TableSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TableSpec_SdkV2) {
	if !from.TableConfiguration.IsNull() && !from.TableConfiguration.IsUnknown() {
		if toTableConfiguration, ok := to.GetTableConfiguration(ctx); ok {
			if fromTableConfiguration, ok := from.GetTableConfiguration(ctx); ok {
				toTableConfiguration.SyncFieldsDuringRead(ctx, fromTableConfiguration)
				to.SetTableConfiguration(ctx, toTableConfiguration)
			}
		}
	}
}

func (m TableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m TableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TableSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *TableSpec_SdkV2) GetTableConfiguration(ctx context.Context) (TableSpecificConfig_SdkV2, bool) {
	var e TableSpecificConfig_SdkV2
	if m.TableConfiguration.IsNull() || m.TableConfiguration.IsUnknown() {
		return e, false
	}
	var v []TableSpecificConfig_SdkV2
	d := m.TableConfiguration.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableConfiguration sets the value of the TableConfiguration field in TableSpec_SdkV2.
func (m *TableSpec_SdkV2) SetTableConfiguration(ctx context.Context, v TableSpecificConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table_configuration"]
	m.TableConfiguration = types.ListValueMust(t, vs)
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
	WorkdayReportParameters types.List `tfsdk:"workday_report_parameters"`
}

func (to *TableSpecificConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableSpecificConfig_SdkV2) {
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

func (to *TableSpecificConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TableSpecificConfig_SdkV2) {
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

func (m TableSpecificConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["exclude_columns"] = attrs["exclude_columns"].SetOptional()
	attrs["include_columns"] = attrs["include_columns"].SetOptional()
	attrs["primary_keys"] = attrs["primary_keys"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].SetOptional()
	attrs["query_based_connector_config"] = attrs["query_based_connector_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["row_filter"] = attrs["row_filter"].SetOptional()
	attrs["salesforce_include_formula_fields"] = attrs["salesforce_include_formula_fields"].SetOptional()
	attrs["scd_type"] = attrs["scd_type"].SetOptional()
	attrs["sequence_by"] = attrs["sequence_by"].SetOptional()
	attrs["workday_report_parameters"] = attrs["workday_report_parameters"].SetOptional()
	attrs["workday_report_parameters"] = attrs["workday_report_parameters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSpecificConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TableSpecificConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude_columns":              reflect.TypeOf(types.String{}),
		"include_columns":              reflect.TypeOf(types.String{}),
		"primary_keys":                 reflect.TypeOf(types.String{}),
		"query_based_connector_config": reflect.TypeOf(IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2{}),
		"sequence_by":                  reflect.TypeOf(types.String{}),
		"workday_report_parameters":    reflect.TypeOf(IngestionPipelineDefinitionWorkdayReportParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSpecificConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m TableSpecificConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TableSpecificConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
			"row_filter":                        types.StringType,
			"salesforce_include_formula_fields": types.BoolType,
			"scd_type":                          types.StringType,
			"sequence_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"workday_report_parameters": basetypes.ListType{
				ElemType: IngestionPipelineDefinitionWorkdayReportParameters_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetExcludeColumns returns the value of the ExcludeColumns field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetExcludeColumns(ctx context.Context) ([]types.String, bool) {
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

// SetExcludeColumns sets the value of the ExcludeColumns field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetExcludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["exclude_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExcludeColumns = types.ListValueMust(t, vs)
}

// GetIncludeColumns returns the value of the IncludeColumns field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetIncludeColumns(ctx context.Context) ([]types.String, bool) {
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

// SetIncludeColumns sets the value of the IncludeColumns field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetIncludeColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IncludeColumns = types.ListValueMust(t, vs)
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeys sets the value of the PrimaryKeys field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeys = types.ListValueMust(t, vs)
}

// GetQueryBasedConnectorConfig returns the value of the QueryBasedConnectorConfig field in TableSpecificConfig_SdkV2 as
// a IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetQueryBasedConnectorConfig(ctx context.Context) (IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2, bool) {
	var e IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
	if m.QueryBasedConnectorConfig.IsNull() || m.QueryBasedConnectorConfig.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2
	d := m.QueryBasedConnectorConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryBasedConnectorConfig sets the value of the QueryBasedConnectorConfig field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetQueryBasedConnectorConfig(ctx context.Context, v IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_based_connector_config"]
	m.QueryBasedConnectorConfig = types.ListValueMust(t, vs)
}

// GetSequenceBy returns the value of the SequenceBy field in TableSpecificConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetSequenceBy(ctx context.Context) ([]types.String, bool) {
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

// SetSequenceBy sets the value of the SequenceBy field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetSequenceBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sequence_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SequenceBy = types.ListValueMust(t, vs)
}

// GetWorkdayReportParameters returns the value of the WorkdayReportParameters field in TableSpecificConfig_SdkV2 as
// a IngestionPipelineDefinitionWorkdayReportParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *TableSpecificConfig_SdkV2) GetWorkdayReportParameters(ctx context.Context) (IngestionPipelineDefinitionWorkdayReportParameters_SdkV2, bool) {
	var e IngestionPipelineDefinitionWorkdayReportParameters_SdkV2
	if m.WorkdayReportParameters.IsNull() || m.WorkdayReportParameters.IsUnknown() {
		return e, false
	}
	var v []IngestionPipelineDefinitionWorkdayReportParameters_SdkV2
	d := m.WorkdayReportParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkdayReportParameters sets the value of the WorkdayReportParameters field in TableSpecificConfig_SdkV2.
func (m *TableSpecificConfig_SdkV2) SetWorkdayReportParameters(ctx context.Context, v IngestionPipelineDefinitionWorkdayReportParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workday_report_parameters"]
	m.WorkdayReportParameters = types.ListValueMust(t, vs)
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

func (to *UpdateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInfo_SdkV2) {
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

func (to *UpdateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateInfo_SdkV2) {
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

func (m UpdateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cause"] = attrs["cause"].SetOptional()
	attrs["cluster_id"] = attrs["cluster_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":                 reflect.TypeOf(PipelineSpec_SdkV2{}),
		"full_refresh_selection": reflect.TypeOf(types.String{}),
		"refresh_selection":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetConfig returns the value of the Config field in UpdateInfo_SdkV2 as
// a PipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInfo_SdkV2) GetConfig(ctx context.Context) (PipelineSpec_SdkV2, bool) {
	var e PipelineSpec_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []PipelineSpec_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in UpdateInfo_SdkV2.
func (m *UpdateInfo_SdkV2) SetConfig(ctx context.Context, v PipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// GetFullRefreshSelection returns the value of the FullRefreshSelection field in UpdateInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInfo_SdkV2) GetFullRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetFullRefreshSelection sets the value of the FullRefreshSelection field in UpdateInfo_SdkV2.
func (m *UpdateInfo_SdkV2) SetFullRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["full_refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FullRefreshSelection = types.ListValueMust(t, vs)
}

// GetRefreshSelection returns the value of the RefreshSelection field in UpdateInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInfo_SdkV2) GetRefreshSelection(ctx context.Context) ([]types.String, bool) {
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

// SetRefreshSelection sets the value of the RefreshSelection field in UpdateInfo_SdkV2.
func (m *UpdateInfo_SdkV2) SetRefreshSelection(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["refresh_selection"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RefreshSelection = types.ListValueMust(t, vs)
}

type UpdateStateInfo_SdkV2 struct {
	CreationTime types.String `tfsdk:"creation_time"`

	State types.String `tfsdk:"state"`

	UpdateId types.String `tfsdk:"update_id"`
}

func (to *UpdateStateInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateStateInfo_SdkV2) {
}

func (to *UpdateStateInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateStateInfo_SdkV2) {
}

func (m UpdateStateInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateStateInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStateInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateStateInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time": m.CreationTime,
			"state":         m.State,
			"update_id":     m.UpdateId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateStateInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.StringType,
			"state":         types.StringType,
			"update_id":     types.StringType,
		},
	}
}
