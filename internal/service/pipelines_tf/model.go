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

	"github.com/databricks/databricks-sdk-go/service/compute"
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

func (newState *CreatePipeline) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipeline) {
}

func (newState *CreatePipeline) SyncEffectiveFieldsDuringRead(existingState CreatePipeline) {
}

func (a CreatePipeline) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.StringType),
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

func (a CreatePipeline) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_duplicate_names": types.BoolType,
			"budget_policy_id":      types.StringType,
			"catalog":               types.StringType,
			"channel":               types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.ToAttrType(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.ToAttrType(ctx),
			},
			"development": types.BoolType,
			"dry_run":     types.BoolType,
			"edition":     types.StringType,
			"filters": basetypes.ListType{
				ElemType: Filters{}.ToAttrType(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.ToAttrType(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.ToAttrType(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.ToAttrType(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.ToAttrType(ctx),
			},
		},
	}
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings types.List `tfsdk:"effective_settings" tf:"optional,object"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
}

func (newState *CreatePipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePipelineResponse) {
}

func (newState *CreatePipelineResponse) SyncEffectiveFieldsDuringRead(existingState CreatePipelineResponse) {
}

func (a CreatePipelineResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_settings": reflect.TypeOf(PipelineSpec{}),
	}
}

func (a CreatePipelineResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_settings": basetypes.ListType{
				ElemType: PipelineSpec{}.ToAttrType(ctx),
			},
			"pipeline_id": types.StringType,
		},
	}
}

type CronTrigger struct {
	QuartzCronSchedule types.String `tfsdk:"quartz_cron_schedule" tf:"optional"`

	TimezoneId types.String `tfsdk:"timezone_id" tf:"optional"`
}

func (newState *CronTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronTrigger) {
}

func (newState *CronTrigger) SyncEffectiveFieldsDuringRead(existingState CronTrigger) {
}

func (a CronTrigger) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CronTrigger) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DataPlaneId) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DataPlaneId) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeletePipelineRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeletePipelineRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeletePipelineResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeletePipelineResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *EditPipeline) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipeline) {
}

func (newState *EditPipeline) SyncEffectiveFieldsDuringRead(existingState EditPipeline) {
}

func (a EditPipeline) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.StringType),
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

func (a EditPipeline) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_duplicate_names": types.BoolType,
			"budget_policy_id":      types.StringType,
			"catalog":               types.StringType,
			"channel":               types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.ToAttrType(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.ToAttrType(ctx),
			},
			"development":            types.BoolType,
			"edition":                types.StringType,
			"expected_last_modified": types.Int64Type,
			"filters": basetypes.ListType{
				ElemType: Filters{}.ToAttrType(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.ToAttrType(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.ToAttrType(ctx),
			},
			"photon":      types.BoolType,
			"pipeline_id": types.StringType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.ToAttrType(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.ToAttrType(ctx),
			},
		},
	}
}

type EditPipelineResponse struct {
}

func (newState *EditPipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditPipelineResponse) {
}

func (newState *EditPipelineResponse) SyncEffectiveFieldsDuringRead(existingState EditPipelineResponse) {
}

func (a EditPipelineResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EditPipelineResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ErrorDetail) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"exceptions": reflect.TypeOf(SerializedException{}),
	}
}

func (a ErrorDetail) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"exceptions": basetypes.ListType{
				ElemType: SerializedException{}.ToAttrType(ctx),
			},
			"fatal": types.BoolType,
		},
	}
}

type FileLibrary struct {
	// The absolute path of the file.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *FileLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileLibrary) {
}

func (newState *FileLibrary) SyncEffectiveFieldsDuringRead(existingState FileLibrary) {
}

func (a FileLibrary) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a FileLibrary) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Filters) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"exclude": reflect.TypeOf(types.StringType),
		"include": reflect.TypeOf(types.StringType),
	}
}

func (a Filters) ToAttrType(ctx context.Context) types.ObjectType {
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

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId types.String `tfsdk:"-"`
}

func (newState *GetPipelinePermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelinePermissionLevelsRequest) {
}

func (newState *GetPipelinePermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetPipelinePermissionLevelsRequest) {
}

func (a GetPipelinePermissionLevelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPipelinePermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetPipelinePermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PipelinePermissionsDescription{}),
	}
}

func (a GetPipelinePermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PipelinePermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
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

func (a GetPipelinePermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPipelinePermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetPipelineRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPipelineRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
	// The pipeline state.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *GetPipelineResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPipelineResponse) {
}

func (newState *GetPipelineResponse) SyncEffectiveFieldsDuringRead(existingState GetPipelineResponse) {
}

func (a GetPipelineResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
		"spec":           reflect.TypeOf(PipelineSpec{}),
	}
}

func (a GetPipelineResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":                      types.StringType,
			"cluster_id":                 types.StringType,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"health":                     types.StringType,
			"last_modified":              types.Int64Type,
			"latest_updates": basetypes.ListType{
				ElemType: UpdateStateInfo{}.ToAttrType(ctx),
			},
			"name":             types.StringType,
			"pipeline_id":      types.StringType,
			"run_as_user_name": types.StringType,
			"spec": basetypes.ListType{
				ElemType: PipelineSpec{}.ToAttrType(ctx),
			},
			"state": types.StringType,
		},
	}
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

func (a GetUpdateRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetUpdateRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pipeline_id": types.StringType,
			"update_id":   types.StringType,
		},
	}
}

type GetUpdateResponse struct {
	// The current update info.
	Update types.List `tfsdk:"update" tf:"optional,object"`
}

func (newState *GetUpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUpdateResponse) {
}

func (newState *GetUpdateResponse) SyncEffectiveFieldsDuringRead(existingState GetUpdateResponse) {
}

func (a GetUpdateResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"update": reflect.TypeOf(UpdateInfo{}),
	}
}

func (a GetUpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"update": basetypes.ListType{
				ElemType: UpdateInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type IngestionConfig struct {
	// Select a specific source report.
	Report types.List `tfsdk:"report" tf:"optional,object"`
	// Select all tables from a specific source schema.
	Schema types.List `tfsdk:"schema" tf:"optional,object"`
	// Select a specific source table.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *IngestionConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionConfig) {
}

func (newState *IngestionConfig) SyncEffectiveFieldsDuringRead(existingState IngestionConfig) {
}

func (a IngestionConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"report": reflect.TypeOf(ReportSpec{}),
		"schema": reflect.TypeOf(SchemaSpec{}),
		"table":  reflect.TypeOf(TableSpec{}),
	}
}

func (a IngestionConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"report": basetypes.ListType{
				ElemType: ReportSpec{}.ToAttrType(ctx),
			},
			"schema": basetypes.ListType{
				ElemType: SchemaSpec{}.ToAttrType(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableSpec{}.ToAttrType(ctx),
			},
		},
	}
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

func (a IngestionGatewayPipelineDefinition) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a IngestionGatewayPipelineDefinition) ToAttrType(ctx context.Context) types.ObjectType {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *IngestionPipelineDefinition) SyncEffectiveFieldsDuringCreateOrUpdate(plan IngestionPipelineDefinition) {
}

func (newState *IngestionPipelineDefinition) SyncEffectiveFieldsDuringRead(existingState IngestionPipelineDefinition) {
}

func (a IngestionPipelineDefinition) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects":             reflect.TypeOf(IngestionConfig{}),
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

func (a IngestionPipelineDefinition) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":      types.StringType,
			"ingestion_gateway_id": types.StringType,
			"objects": basetypes.ListType{
				ElemType: IngestionConfig{}.ToAttrType(ctx),
			},
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.ToAttrType(ctx),
			},
		},
	}
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

func (a ListPipelineEventsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.StringType),
	}
}

func (a ListPipelineEventsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListPipelineEventsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(PipelineEvent{}),
	}
}

func (a ListPipelineEventsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: PipelineEvent{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
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

func (a ListPipelinesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.StringType),
	}
}

func (a ListPipelinesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListPipelinesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"statuses": reflect.TypeOf(PipelineStateInfo{}),
	}
}

func (a ListPipelinesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"statuses": basetypes.ListType{
				ElemType: PipelineStateInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (a ListUpdatesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListUpdatesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListUpdatesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(UpdateInfo{}),
	}
}

func (a ListUpdatesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"updates": basetypes.ListType{
				ElemType: UpdateInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ManualTrigger struct {
}

func (newState *ManualTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan ManualTrigger) {
}

func (newState *ManualTrigger) SyncEffectiveFieldsDuringRead(existingState ManualTrigger) {
}

func (a ManualTrigger) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ManualTrigger) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a NotebookLibrary) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a NotebookLibrary) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Notifications) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"alerts":           reflect.TypeOf(types.StringType),
		"email_recipients": reflect.TypeOf(types.StringType),
	}
}

func (a Notifications) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Origin) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Origin) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PipelineAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PipelineAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PipelineAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PipelinePermission{}),
	}
}

func (a PipelineAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PipelinePermission{}.ToAttrType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale types.List `tfsdk:"autoscale" tf:"optional,object"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes compute.AwsAttributes `tfsdk:"aws_attributes" tf:"optional,object"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes compute.AzureAttributes `tfsdk:"azure_attributes" tf:"optional,object"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf compute.ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional,object"`
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
	GcpAttributes compute.GcpAttributes `tfsdk:"gcp_attributes" tf:"optional,object"`
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

func (a PipelineCluster) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"autoscale":        reflect.TypeOf(PipelineClusterAutoscale{}),
		"aws_attributes":   reflect.TypeOf(compute.AwsAttributes{}),
		"azure_attributes": reflect.TypeOf(compute.AzureAttributes{}),
		"cluster_log_conf": reflect.TypeOf(compute.ClusterLogConf{}),
		"custom_tags":      reflect.TypeOf(types.StringType),
		"gcp_attributes":   reflect.TypeOf(compute.GcpAttributes{}),
		"init_scripts":     reflect.TypeOf(compute.InitScriptInfo{}),
		"spark_conf":       reflect.TypeOf(types.StringType),
		"spark_env_vars":   reflect.TypeOf(types.StringType),
		"ssh_public_keys":  reflect.TypeOf(types.StringType),
	}
}

func (a PipelineCluster) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apply_policy_default_values": types.BoolType,
			"autoscale": basetypes.ListType{
				ElemType: PipelineClusterAutoscale{}.ToAttrType(ctx),
			},
			"aws_attributes": basetypes.ListType{
				ElemType: compute_tf.AwsAttributes{}.ToAttrType(ctx),
			},
			"azure_attributes": basetypes.ListType{
				ElemType: compute_tf.AzureAttributes{}.ToAttrType(ctx),
			},
			"cluster_log_conf": basetypes.ListType{
				ElemType: compute_tf.ClusterLogConf{}.ToAttrType(ctx),
			},
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"driver_instance_pool_id":      types.StringType,
			"driver_node_type_id":          types.StringType,
			"enable_local_disk_encryption": types.BoolType,
			"gcp_attributes": basetypes.ListType{
				ElemType: compute_tf.GcpAttributes{}.ToAttrType(ctx),
			},
			"init_scripts": basetypes.ListType{
				ElemType: compute_tf.InitScriptInfo{}.ToAttrType(ctx),
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

func (a PipelineClusterAutoscale) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PipelineClusterAutoscale) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PipelineDeployment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PipelineDeployment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kind":               types.StringType,
			"metadata_file_path": types.StringType,
		},
	}
}

type PipelineEvent struct {
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

func (newState *PipelineEvent) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineEvent) {
}

func (newState *PipelineEvent) SyncEffectiveFieldsDuringRead(existingState PipelineEvent) {
}

func (a PipelineEvent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"error":    reflect.TypeOf(ErrorDetail{}),
		"origin":   reflect.TypeOf(Origin{}),
		"sequence": reflect.TypeOf(Sequencing{}),
	}
}

func (a PipelineEvent) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": basetypes.ListType{
				ElemType: ErrorDetail{}.ToAttrType(ctx),
			},
			"event_type":     types.StringType,
			"id":             types.StringType,
			"level":          types.StringType,
			"maturity_level": types.StringType,
			"message":        types.StringType,
			"origin": basetypes.ListType{
				ElemType: Origin{}.ToAttrType(ctx),
			},
			"sequence": basetypes.ListType{
				ElemType: Sequencing{}.ToAttrType(ctx),
			},
			"timestamp": types.StringType,
		},
	}
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File types.List `tfsdk:"file" tf:"optional,object"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed.
	Maven compute.MavenLibrary `tfsdk:"maven" tf:"optional,object"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook types.List `tfsdk:"notebook" tf:"optional,object"`
	// URI of the whl to be installed.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

func (newState *PipelineLibrary) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineLibrary) {
}

func (newState *PipelineLibrary) SyncEffectiveFieldsDuringRead(existingState PipelineLibrary) {
}

func (a PipelineLibrary) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"file":     reflect.TypeOf(FileLibrary{}),
		"maven":    reflect.TypeOf(compute.MavenLibrary{}),
		"notebook": reflect.TypeOf(NotebookLibrary{}),
	}
}

func (a PipelineLibrary) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file": basetypes.ListType{
				ElemType: FileLibrary{}.ToAttrType(ctx),
			},
			"jar": types.StringType,
			"maven": basetypes.ListType{
				ElemType: compute_tf.MavenLibrary{}.ToAttrType(ctx),
			},
			"notebook": basetypes.ListType{
				ElemType: NotebookLibrary{}.ToAttrType(ctx),
			},
			"whl": types.StringType,
		},
	}
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

func (a PipelinePermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.StringType),
	}
}

func (a PipelinePermission) ToAttrType(ctx context.Context) types.ObjectType {
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

type PipelinePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *PipelinePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelinePermissions) {
}

func (newState *PipelinePermissions) SyncEffectiveFieldsDuringRead(existingState PipelinePermissions) {
}

func (a PipelinePermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlResponse{}),
	}
}

func (a PipelinePermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PipelineAccessControlResponse{}.ToAttrType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
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

func (a PipelinePermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PipelinePermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PipelinePermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PipelineAccessControlRequest{}),
	}
}

func (a PipelinePermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PipelineAccessControlRequest{}.ToAttrType(ctx),
			},
			"pipeline_id": types.StringType,
		},
	}
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

func (newState *PipelineSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineSpec) {
}

func (newState *PipelineSpec) SyncEffectiveFieldsDuringRead(existingState PipelineSpec) {
}

func (a PipelineSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"clusters":             reflect.TypeOf(PipelineCluster{}),
		"configuration":        reflect.TypeOf(types.StringType),
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

func (a PipelineSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"catalog":          types.StringType,
			"channel":          types.StringType,
			"clusters": basetypes.ListType{
				ElemType: PipelineCluster{}.ToAttrType(ctx),
			},
			"configuration": basetypes.MapType{
				ElemType: types.StringType,
			},
			"continuous": types.BoolType,
			"deployment": basetypes.ListType{
				ElemType: PipelineDeployment{}.ToAttrType(ctx),
			},
			"development": types.BoolType,
			"edition":     types.StringType,
			"filters": basetypes.ListType{
				ElemType: Filters{}.ToAttrType(ctx),
			},
			"gateway_definition": basetypes.ListType{
				ElemType: IngestionGatewayPipelineDefinition{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"ingestion_definition": basetypes.ListType{
				ElemType: IngestionPipelineDefinition{}.ToAttrType(ctx),
			},
			"libraries": basetypes.ListType{
				ElemType: PipelineLibrary{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"notifications": basetypes.ListType{
				ElemType: Notifications{}.ToAttrType(ctx),
			},
			"photon": types.BoolType,
			"restart_window": basetypes.ListType{
				ElemType: RestartWindow{}.ToAttrType(ctx),
			},
			"schema":     types.StringType,
			"serverless": types.BoolType,
			"storage":    types.StringType,
			"target":     types.StringType,
			"trigger": basetypes.ListType{
				ElemType: PipelineTrigger{}.ToAttrType(ctx),
			},
		},
	}
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

func (a PipelineStateInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_updates": reflect.TypeOf(UpdateStateInfo{}),
	}
}

func (a PipelineStateInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":        types.StringType,
			"creator_user_name": types.StringType,
			"health":            types.StringType,
			"latest_updates": basetypes.ListType{
				ElemType: UpdateStateInfo{}.ToAttrType(ctx),
			},
			"name":             types.StringType,
			"pipeline_id":      types.StringType,
			"run_as_user_name": types.StringType,
			"state":            types.StringType,
		},
	}
}

type PipelineTrigger struct {
	Cron types.List `tfsdk:"cron" tf:"optional,object"`

	Manual []ManualTrigger `tfsdk:"manual" tf:"optional,object"`
}

func (newState *PipelineTrigger) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineTrigger) {
}

func (newState *PipelineTrigger) SyncEffectiveFieldsDuringRead(existingState PipelineTrigger) {
}

func (a PipelineTrigger) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron":   reflect.TypeOf(CronTrigger{}),
		"manual": reflect.TypeOf(ManualTrigger{}),
	}
}

func (a PipelineTrigger) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cron": basetypes.ListType{
				ElemType: CronTrigger{}.ToAttrType(ctx),
			},
			"manual": basetypes.ListType{
				ElemType: ManualTrigger{}.ToAttrType(ctx),
			},
		},
	}
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *ReportSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReportSpec) {
}

func (newState *ReportSpec) SyncEffectiveFieldsDuringRead(existingState ReportSpec) {
}

func (a ReportSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

func (a ReportSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_url":          types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type RestartWindow struct {
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

func (newState *RestartWindow) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestartWindow) {
}

func (newState *RestartWindow) SyncEffectiveFieldsDuringRead(existingState RestartWindow) {
}

func (a RestartWindow) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestartWindow) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"days_of_week": types.StringType,
			"start_hour":   types.Int64Type,
			"time_zone_id": types.StringType,
		},
	}
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *SchemaSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaSpec) {
}

func (newState *SchemaSpec) SyncEffectiveFieldsDuringRead(existingState SchemaSpec) {
}

func (a SchemaSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

func (a SchemaSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo types.Int64 `tfsdk:"control_plane_seq_no" tf:"optional"`
	// the ID assigned by the data plane.
	DataPlaneId types.List `tfsdk:"data_plane_id" tf:"optional,object"`
}

func (newState *Sequencing) SyncEffectiveFieldsDuringCreateOrUpdate(plan Sequencing) {
}

func (newState *Sequencing) SyncEffectiveFieldsDuringRead(existingState Sequencing) {
}

func (a Sequencing) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_plane_id": reflect.TypeOf(DataPlaneId{}),
	}
}

func (a Sequencing) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"control_plane_seq_no": types.Int64Type,
			"data_plane_id": basetypes.ListType{
				ElemType: DataPlaneId{}.ToAttrType(ctx),
			},
		},
	}
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

func (a SerializedException) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"stack": reflect.TypeOf(StackFrame{}),
	}
}

func (a SerializedException) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"class_name": types.StringType,
			"message":    types.StringType,
			"stack": basetypes.ListType{
				ElemType: StackFrame{}.ToAttrType(ctx),
			},
		},
	}
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

func (a StackFrame) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a StackFrame) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a StartUpdate) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_refresh_selection": reflect.TypeOf(types.StringType),
		"refresh_selection":      reflect.TypeOf(types.StringType),
	}
}

func (a StartUpdate) ToAttrType(ctx context.Context) types.ObjectType {
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

type StartUpdateResponse struct {
	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *StartUpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartUpdateResponse) {
}

func (newState *StartUpdateResponse) SyncEffectiveFieldsDuringRead(existingState StartUpdateResponse) {
}

func (a StartUpdateResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a StartUpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a StopPipelineResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a StopPipelineResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a StopRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a StopRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	TableConfiguration types.List `tfsdk:"table_configuration" tf:"optional,object"`
}

func (newState *TableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSpec) {
}

func (newState *TableSpec) SyncEffectiveFieldsDuringRead(existingState TableSpec) {
}

func (a TableSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_configuration": reflect.TypeOf(TableSpecificConfig{}),
	}
}

func (a TableSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_catalog": types.StringType,
			"destination_schema":  types.StringType,
			"destination_table":   types.StringType,
			"source_catalog":      types.StringType,
			"source_schema":       types.StringType,
			"source_table":        types.StringType,
			"table_configuration": basetypes.ListType{
				ElemType: TableSpecificConfig{}.ToAttrType(ctx),
			},
		},
	}
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

func (a TableSpecificConfig) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.StringType),
		"sequence_by":  reflect.TypeOf(types.StringType),
	}
}

func (a TableSpecificConfig) ToAttrType(ctx context.Context) types.ObjectType {
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

type UpdateInfo struct {
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

func (newState *UpdateInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateInfo) {
}

func (newState *UpdateInfo) SyncEffectiveFieldsDuringRead(existingState UpdateInfo) {
}

func (a UpdateInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":                 reflect.TypeOf(PipelineSpec{}),
		"full_refresh_selection": reflect.TypeOf(types.StringType),
		"refresh_selection":      reflect.TypeOf(types.StringType),
	}
}

func (a UpdateInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cause":      types.StringType,
			"cluster_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: PipelineSpec{}.ToAttrType(ctx),
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

type UpdateStateInfo struct {
	CreationTime types.String `tfsdk:"creation_time" tf:"optional"`

	State types.String `tfsdk:"state" tf:"optional"`

	UpdateId types.String `tfsdk:"update_id" tf:"optional"`
}

func (newState *UpdateStateInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStateInfo) {
}

func (newState *UpdateStateInfo) SyncEffectiveFieldsDuringRead(existingState UpdateStateInfo) {
}

func (a UpdateStateInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateStateInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.StringType,
			"state":         types.StringType,
			"update_id":     types.StringType,
		},
	}
}
