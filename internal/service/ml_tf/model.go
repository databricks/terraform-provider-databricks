// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package ml_tf

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Activity recorded for the action.
type Activity struct {
	// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied
	// the corresponding stage transition.
	//
	// * `REQUESTED_TRANSITION`: User requested the corresponding stage
	// transition.
	//
	// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
	//
	// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
	//
	// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
	//
	// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
	// archiving existing model versions in a stage.
	ActivityType types.String `tfsdk:"activity_type" tf:"optional"`
	// User-provided comment associated with the activity.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	FromStage types.String `tfsdk:"from_stage" tf:"optional"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment types.String `tfsdk:"system_comment" tf:"optional"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage types.String `tfsdk:"to_stage" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *Activity) SyncEffectiveFieldsDuringCreateOrUpdate(plan Activity) {
}

func (newState *Activity) SyncEffectiveFieldsDuringRead(existingState Activity) {
}

func (a Activity) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Activity) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ActivityType":         types.StringType,
			"Comment":              types.StringType,
			"CreationTimestamp":    types.Int64Type,
			"FromStage":            types.StringType,
			"Id":                   types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"SystemComment":        types.StringType,
			"ToStage":              types.StringType,
			"UserId":               types.StringType,
		},
	}
}

type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions" tf:""`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *ApproveTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApproveTransitionRequest) {
}

func (newState *ApproveTransitionRequest) SyncEffectiveFieldsDuringRead(existingState ApproveTransitionRequest) {
}

func (a ApproveTransitionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ApproveTransitionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArchiveExistingVersions": types.BoolType,
			"Comment":                 types.StringType,
			"Name":                    types.StringType,
			"Stage":                   types.StringType,
			"Version":                 types.StringType,
		},
	}
}

type ApproveTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity types.List `tfsdk:"activity" tf:"optional,object"`
}

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApproveTransitionRequestResponse) {
}

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState ApproveTransitionRequestResponse) {
}

func (a ApproveTransitionRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Activity": reflect.TypeOf(Activity{}),
	}
}

func (a ApproveTransitionRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Activity": Activity{}.ToAttrType(ctx),
		},
	}
}

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions" tf:"optional"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Comment ID
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *CommentObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommentObject) {
}

func (newState *CommentObject) SyncEffectiveFieldsDuringRead(existingState CommentObject) {
}

func (a CommentObject) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AvailableActions": reflect.TypeOf(types.StringType),
	}
}

func (a CommentObject) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AvailableActions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Comment":              types.StringType,
			"CreationTimestamp":    types.Int64Type,
			"Id":                   types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"UserId":               types.StringType,
		},
	}
}

type CreateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *CreateComment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateComment) {
}

func (newState *CreateComment) SyncEffectiveFieldsDuringRead(existingState CreateComment) {
}

func (a CreateComment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateComment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": types.StringType,
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type CreateCommentResponse struct {
	// Comment details.
	Comment types.List `tfsdk:"comment" tf:"optional,object"`
}

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCommentResponse) {
}

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringRead(existingState CreateCommentResponse) {
}

func (a CreateCommentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Comment": reflect.TypeOf(CommentObject{}),
	}
}

func (a CreateCommentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": CommentObject{}.ToAttrType(ctx),
		},
	}
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation types.String `tfsdk:"artifact_location" tf:"optional"`
	// Experiment name.
	Name types.String `tfsdk:"name" tf:""`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExperiment) {
}

func (newState *CreateExperiment) SyncEffectiveFieldsDuringRead(existingState CreateExperiment) {
}

func (a CreateExperiment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ExperimentTag{}),
	}
}

func (a CreateExperiment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArtifactLocation": types.StringType,
			"Name":             types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ExperimentTag{}.ToAttrType(ctx),
			},
		},
	}
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExperimentResponse) {
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringRead(existingState CreateExperimentResponse) {
}

func (a CreateExperimentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateExperimentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Register models under this name
	Name types.String `tfsdk:"name" tf:""`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelRequest) {
}

func (newState *CreateModelRequest) SyncEffectiveFieldsDuringRead(existingState CreateModelRequest) {
}

func (a CreateModelRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ModelTag{}),
	}
}

func (a CreateModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Name":        types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelTag{}.ToAttrType(ctx),
			},
		},
	}
}

type CreateModelResponse struct {
	RegisteredModel types.List `tfsdk:"registered_model" tf:"optional,object"`
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelResponse) {
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelResponse) {
}

func (a CreateModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RegisteredModel": reflect.TypeOf(Model{}),
	}
}

func (a CreateModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RegisteredModel": Model{}.ToAttrType(ctx),
		},
	}
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Register model under this name
	Name types.String `tfsdk:"name" tf:""`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI indicating the location of the model artifacts.
	Source types.String `tfsdk:"source" tf:""`
	// Additional metadata for model version.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelVersionRequest) {
}

func (newState *CreateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState CreateModelVersionRequest) {
}

func (a CreateModelVersionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

func (a CreateModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Name":        types.StringType,
			"RunId":       types.StringType,
			"RunLink":     types.StringType,
			"Source":      types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelVersionTag{}.ToAttrType(ctx),
			},
		},
	}
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelVersionResponse) {
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelVersionResponse) {
}

func (a CreateModelVersionResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ModelVersion": reflect.TypeOf(ModelVersion{}),
	}
}

func (a CreateModelVersionResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelVersion": ModelVersion{}.ToAttrType(ctx),
		},
	}
}

type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events" tf:""`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *CreateRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegistryWebhook) {
}

func (newState *CreateRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState CreateRegistryWebhook) {
}

func (a CreateRegistryWebhook) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Events":      reflect.TypeOf(types.StringType),
		"HttpUrlSpec": reflect.TypeOf(HttpUrlSpec{}),
		"JobSpec":     reflect.TypeOf(JobSpec{}),
	}
}

func (a CreateRegistryWebhook) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"HttpUrlSpec": HttpUrlSpec{}.ToAttrType(ctx),
			"JobSpec":     JobSpec{}.ToAttrType(ctx),
			"ModelName":   types.StringType,
			"Status":      types.StringType,
		},
	}
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Additional metadata for run.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *CreateRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRun) {
}

func (newState *CreateRun) SyncEffectiveFieldsDuringRead(existingState CreateRun) {
}

func (a CreateRun) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(RunTag{}),
	}
}

func (a CreateRun) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
			"StartTime":    types.Int64Type,
			"Tags": basetypes.ListType{
				ElemType: RunTag{}.ToAttrType(ctx),
			},
			"UserId": types.StringType,
		},
	}
}

type CreateRunResponse struct {
	// The newly created run.
	Run types.List `tfsdk:"run" tf:"optional,object"`
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRunResponse) {
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringRead(existingState CreateRunResponse) {
}

func (a CreateRunResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Run": reflect.TypeOf(Run{}),
	}
}

func (a CreateRunResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Run": Run{}.ToAttrType(ctx),
		},
	}
}

type CreateTransitionRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *CreateTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTransitionRequest) {
}

func (newState *CreateTransitionRequest) SyncEffectiveFieldsDuringRead(existingState CreateTransitionRequest) {
}

func (a CreateTransitionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateTransitionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": types.StringType,
			"Name":    types.StringType,
			"Stage":   types.StringType,
			"Version": types.StringType,
		},
	}
}

type CreateTransitionRequestResponse struct {
	// Transition request details.
	Request types.List `tfsdk:"request" tf:"optional,object"`
}

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTransitionRequestResponse) {
}

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState CreateTransitionRequestResponse) {
}

func (a CreateTransitionRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Request": reflect.TypeOf(TransitionRequest{}),
	}
}

func (a CreateTransitionRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Request": TransitionRequest{}.ToAttrType(ctx),
		},
	}
}

type CreateWebhookResponse struct {
	Webhook types.List `tfsdk:"webhook" tf:"optional,object"`
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWebhookResponse) {
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState CreateWebhookResponse) {
}

func (a CreateWebhookResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Webhook": reflect.TypeOf(RegistryWebhook{}),
	}
}

func (a CreateWebhookResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Webhook": RegistryWebhook{}.ToAttrType(ctx),
		},
	}
}

type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest types.String `tfsdk:"digest" tf:"optional"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name types.String `tfsdk:"name" tf:"optional"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile types.String `tfsdk:"profile" tf:"optional"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	Source types.String `tfsdk:"source" tf:"optional"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	SourceType types.String `tfsdk:"source_type" tf:"optional"`
}

func (newState *Dataset) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dataset) {
}

func (newState *Dataset) SyncEffectiveFieldsDuringRead(existingState Dataset) {
}

func (a Dataset) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Dataset) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Digest":     types.StringType,
			"Name":       types.StringType,
			"Profile":    types.StringType,
			"Schema":     types.StringType,
			"Source":     types.StringType,
			"SourceType": types.StringType,
		},
	}
}

type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset types.List `tfsdk:"dataset" tf:"optional,object"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DatasetInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatasetInput) {
}

func (newState *DatasetInput) SyncEffectiveFieldsDuringRead(existingState DatasetInput) {
}

func (a DatasetInput) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Dataset": reflect.TypeOf(Dataset{}),
		"Tags":    reflect.TypeOf(InputTag{}),
	}
}

func (a DatasetInput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Dataset": Dataset{}.ToAttrType(ctx),
			"Tags": basetypes.ListType{
				ElemType: InputTag{}.ToAttrType(ctx),
			},
		},
	}
}

// Delete a comment
type DeleteCommentRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteCommentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCommentRequest) {
}

func (newState *DeleteCommentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCommentRequest) {
}

func (a DeleteCommentRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteCommentRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteCommentResponse struct {
}

func (newState *DeleteCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCommentResponse) {
}

func (newState *DeleteCommentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCommentResponse) {
}

func (a DeleteCommentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteCommentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

func (newState *DeleteExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExperiment) {
}

func (newState *DeleteExperiment) SyncEffectiveFieldsDuringRead(existingState DeleteExperiment) {
}

func (a DeleteExperiment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExperiment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

type DeleteExperimentResponse struct {
}

func (newState *DeleteExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExperimentResponse) {
}

func (newState *DeleteExperimentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExperimentResponse) {
}

func (a DeleteExperimentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteExperimentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model
type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelRequest) {
}

func (newState *DeleteModelRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelRequest) {
}

func (a DeleteModelRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name": types.StringType,
		},
	}
}

type DeleteModelResponse struct {
}

func (newState *DeleteModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelResponse) {
}

func (newState *DeleteModelResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelResponse) {
}

func (a DeleteModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model tag
type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteModelTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelTagRequest) {
}

func (newState *DeleteModelTagRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelTagRequest) {
}

func (a DeleteModelTagRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelTagRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":  types.StringType,
			"Name": types.StringType,
		},
	}
}

type DeleteModelTagResponse struct {
}

func (newState *DeleteModelTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelTagResponse) {
}

func (newState *DeleteModelTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelTagResponse) {
}

func (a DeleteModelTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model version.
type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionRequest) {
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionRequest) {
}

func (a DeleteModelVersionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type DeleteModelVersionResponse struct {
}

func (newState *DeleteModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionResponse) {
}

func (newState *DeleteModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionResponse) {
}

func (a DeleteModelVersionResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelVersionResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model version tag
type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
	// Model version number that the tag was logged under.
	Version types.String `tfsdk:"-"`
}

func (newState *DeleteModelVersionTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionTagRequest) {
}

func (newState *DeleteModelVersionTagRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionTagRequest) {
}

func (a DeleteModelVersionTagRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelVersionTagRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":     types.StringType,
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type DeleteModelVersionTagResponse struct {
}

func (newState *DeleteModelVersionTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionTagResponse) {
}

func (newState *DeleteModelVersionTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionTagResponse) {
}

func (a DeleteModelVersionTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteModelVersionTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRun) {
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringRead(existingState DeleteRun) {
}

func (a DeleteRun) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRun) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunId": types.StringType,
		},
	}
}

type DeleteRunResponse struct {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunResponse) {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunResponse) {
}

func (a DeleteRunResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRunResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs" tf:"optional"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis types.Int64 `tfsdk:"max_timestamp_millis" tf:""`
}

func (newState *DeleteRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRuns) {
}

func (newState *DeleteRuns) SyncEffectiveFieldsDuringRead(existingState DeleteRuns) {
}

func (a DeleteRuns) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRuns) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId":       types.StringType,
			"MaxRuns":            types.Int64Type,
			"MaxTimestampMillis": types.Int64Type,
		},
	}
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted types.Int64 `tfsdk:"runs_deleted" tf:"optional"`
}

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunsResponse) {
}

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunsResponse) {
}

func (a DeleteRunsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRunsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunsDeleted": types.Int64Type,
		},
	}
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *DeleteTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTag) {
}

func (newState *DeleteTag) SyncEffectiveFieldsDuringRead(existingState DeleteTag) {
}

func (a DeleteTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"RunId": types.StringType,
		},
	}
}

type DeleteTagResponse struct {
}

func (newState *DeleteTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTagResponse) {
}

func (newState *DeleteTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteTagResponse) {
}

func (a DeleteTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a transition request
type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"-"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator types.String `tfsdk:"-"`
	// Name of the model.
	Name types.String `tfsdk:"-"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (newState *DeleteTransitionRequestRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTransitionRequestRequest) {
}

func (newState *DeleteTransitionRequestRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTransitionRequestRequest) {
}

func (a DeleteTransitionRequestRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteTransitionRequestRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": types.StringType,
			"Creator": types.StringType,
			"Name":    types.StringType,
			"Stage":   types.StringType,
			"Version": types.StringType,
		},
	}
}

type DeleteTransitionRequestResponse struct {
}

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTransitionRequestResponse) {
}

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState DeleteTransitionRequestResponse) {
}

func (a DeleteTransitionRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteTransitionRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a webhook
type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteWebhookRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWebhookRequest) {
}

func (newState *DeleteWebhookRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWebhookRequest) {
}

func (a DeleteWebhookRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteWebhookRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Id": types.StringType,
		},
	}
}

type DeleteWebhookResponse struct {
}

func (newState *DeleteWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWebhookResponse) {
}

func (newState *DeleteWebhookResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWebhookResponse) {
}

func (a DeleteWebhookResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteWebhookResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Experiment struct {
	// Location where artifacts for the experiment are stored.
	ArtifactLocation types.String `tfsdk:"artifact_location" tf:"optional"`
	// Creation time
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Last update time
	LastUpdateTime types.Int64 `tfsdk:"last_update_time" tf:"optional"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage types.String `tfsdk:"lifecycle_stage" tf:"optional"`
	// Human readable name that identifies the experiment.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Tags: Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *Experiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan Experiment) {
}

func (newState *Experiment) SyncEffectiveFieldsDuringRead(existingState Experiment) {
}

func (a Experiment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ExperimentTag{}),
	}
}

func (a Experiment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArtifactLocation": types.StringType,
			"CreationTime":     types.Int64Type,
			"ExperimentId":     types.StringType,
			"LastUpdateTime":   types.Int64Type,
			"LifecycleStage":   types.StringType,
			"Name":             types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ExperimentTag{}.ToAttrType(ctx),
			},
		},
	}
}

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlRequest) {
}

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlRequest) {
}

func (a ExperimentAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExperimentAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"GroupName":            types.StringType,
			"PermissionLevel":      types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
}

type ExperimentAccessControlResponse struct {
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

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlResponse) {
}

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlResponse) {
}

func (a ExperimentAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AllPermissions": reflect.TypeOf(ExperimentPermission{}),
	}
}

func (a ExperimentAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllPermissions": basetypes.ListType{
				ElemType: ExperimentPermission{}.ToAttrType(ctx),
			},
			"DisplayName":          types.StringType,
			"GroupName":            types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
}

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermission) {
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringRead(existingState ExperimentPermission) {
}

func (a ExperimentPermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"InheritedFromObject": reflect.TypeOf(types.StringType),
	}
}

func (a ExperimentPermission) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Inherited": types.BoolType,
			"InheritedFromObject": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PermissionLevel": types.StringType,
		},
	}
}

type ExperimentPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissions) {
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissions) {
}

func (a ExperimentPermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(ExperimentAccessControlResponse{}),
	}
}

func (a ExperimentPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: ExperimentAccessControlResponse{}.ToAttrType(ctx),
			},
			"ObjectId":   types.StringType,
			"ObjectType": types.StringType,
		},
	}
}

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissionsDescription) {
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissionsDescription) {
}

func (a ExperimentPermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExperimentPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description":     types.StringType,
			"PermissionLevel": types.StringType,
		},
	}
}

type ExperimentPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *ExperimentPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissionsRequest) {
}

func (newState *ExperimentPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissionsRequest) {
}

func (a ExperimentPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(ExperimentAccessControlRequest{}),
	}
}

func (a ExperimentPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: ExperimentAccessControlRequest{}.ToAttrType(ctx),
			},
			"ExperimentId": types.StringType,
		},
	}
}

type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentTag) {
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringRead(existingState ExperimentTag) {
}

func (a ExperimentTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExperimentTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type FileInfo struct {
	// Size in bytes. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size" tf:"optional"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir" tf:"optional"`
	// Path relative to the root artifact directory run.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *FileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncEffectiveFieldsDuringRead(existingState FileInfo) {
}

func (a FileInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a FileInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FileSize": types.Int64Type,
			"IsDir":    types.BoolType,
			"Path":     types.StringType,
		},
	}
}

// Get metadata
type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

func (newState *GetByNameRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetByNameRequest) {
}

func (newState *GetByNameRequest) SyncEffectiveFieldsDuringRead(existingState GetByNameRequest) {
}

func (a GetByNameRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetByNameRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentName": types.StringType,
		},
	}
}

// Get experiment permission levels
type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionLevelsRequest) {
}

func (newState *GetExperimentPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionLevelsRequest) {
}

func (a GetExperimentPermissionLevelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetExperimentPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionLevelsResponse) {
}

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionLevelsResponse) {
}

func (a GetExperimentPermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PermissionLevels": reflect.TypeOf(ExperimentPermissionsDescription{}),
	}
}

func (a GetExperimentPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PermissionLevels": basetypes.ListType{
				ElemType: ExperimentPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Get experiment permissions
type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionsRequest) {
}

func (newState *GetExperimentPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionsRequest) {
}

func (a GetExperimentPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetExperimentPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

// Get an experiment
type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentRequest) {
}

func (newState *GetExperimentRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentRequest) {
}

func (a GetExperimentRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetExperimentRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment types.List `tfsdk:"experiment" tf:"optional,object"`
}

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentResponse) {
}

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentResponse) {
}

func (a GetExperimentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Experiment": reflect.TypeOf(Experiment{}),
	}
}

func (a GetExperimentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Experiment": Experiment{}.ToAttrType(ctx),
		},
	}
}

// Get history of a given metric within a run
type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the metric.
	MetricKey types.String `tfsdk:"-"`
	// Token indicating the page of metric histories to fetch.
	PageToken types.String `tfsdk:"-"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use run_id instead] ID of the run from which to fetch metric
	// values. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *GetHistoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetHistoryRequest) {
}

func (newState *GetHistoryRequest) SyncEffectiveFieldsDuringRead(existingState GetHistoryRequest) {
}

func (a GetHistoryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetHistoryRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"MaxResults": types.Int64Type,
			"MetricKey":  types.StringType,
			"PageToken":  types.StringType,
			"RunId":      types.StringType,
			"RunUuid":    types.StringType,
		},
	}
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
	// List of stages.
	Stages types.List `tfsdk:"stages" tf:"optional"`
}

func (newState *GetLatestVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionsRequest) {
}

func (newState *GetLatestVersionsRequest) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionsRequest) {
}

func (a GetLatestVersionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Stages": reflect.TypeOf(types.StringType),
	}
}

func (a GetLatestVersionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name": types.StringType,
			"Stages": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionsResponse) {
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionsResponse) {
}

func (a GetLatestVersionsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ModelVersions": reflect.TypeOf(ModelVersion{}),
	}
}

func (a GetLatestVersionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelVersions": basetypes.ListType{
				ElemType: ModelVersion{}.ToAttrType(ctx),
			},
		},
	}
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Token that can be used to retrieve the next page of metric history
	// results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetricHistoryResponse) {
}

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetricHistoryResponse) {
}

func (a GetMetricHistoryResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Metrics": reflect.TypeOf(Metric{}),
	}
}

func (a GetMetricHistoryResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Metrics": basetypes.ListType{
				ElemType: Metric{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (newState *GetModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelRequest) {
}

func (newState *GetModelRequest) SyncEffectiveFieldsDuringRead(existingState GetModelRequest) {
}

func (a GetModelRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name": types.StringType,
		},
	}
}

type GetModelResponse struct {
	RegisteredModelDatabricks types.List `tfsdk:"registered_model_databricks" tf:"optional,object"`
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelResponse) {
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringRead(existingState GetModelResponse) {
}

func (a GetModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RegisteredModelDatabricks": reflect.TypeOf(ModelDatabricks{}),
	}
}

func (a GetModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RegisteredModelDatabricks": ModelDatabricks{}.ToAttrType(ctx),
		},
	}
}

// Get a model version URI
type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *GetModelVersionDownloadUriRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionDownloadUriRequest) {
}

func (newState *GetModelVersionDownloadUriRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionDownloadUriRequest) {
}

func (a GetModelVersionDownloadUriRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetModelVersionDownloadUriRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri types.String `tfsdk:"artifact_uri" tf:"optional"`
}

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionDownloadUriResponse) {
}

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionDownloadUriResponse) {
}

func (a GetModelVersionDownloadUriResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetModelVersionDownloadUriResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArtifactUri": types.StringType,
		},
	}
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionRequest) {
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionRequest) {
}

func (a GetModelVersionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type GetModelVersionResponse struct {
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionResponse) {
}

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionResponse) {
}

func (a GetModelVersionResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ModelVersion": reflect.TypeOf(ModelVersion{}),
	}
}

func (a GetModelVersionResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelVersion": ModelVersion{}.ToAttrType(ctx),
		},
	}
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *GetRegisteredModelPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionLevelsRequest) {
}

func (newState *GetRegisteredModelPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionLevelsRequest) {
}

func (a GetRegisteredModelPermissionLevelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRegisteredModelPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RegisteredModelId": types.StringType,
		},
	}
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionLevelsResponse) {
}

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionLevelsResponse) {
}

func (a GetRegisteredModelPermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PermissionLevels": reflect.TypeOf(RegisteredModelPermissionsDescription{}),
	}
}

func (a GetRegisteredModelPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PermissionLevels": basetypes.ListType{
				ElemType: RegisteredModelPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Get registered model permissions
type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *GetRegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionsRequest) {
}

func (newState *GetRegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionsRequest) {
}

func (a GetRegisteredModelPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRegisteredModelPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RegisteredModelId": types.StringType,
		},
	}
}

// Get a run
type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use run_id instead] ID of the run to fetch. This field will
	// be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunRequest) {
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringRead(existingState GetRunRequest) {
}

func (a GetRunRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRunRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunId":   types.StringType,
			"RunUuid": types.StringType,
		},
	}
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run types.List `tfsdk:"run" tf:"optional,object"`
}

func (newState *GetRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunResponse) {
}

func (newState *GetRunResponse) SyncEffectiveFieldsDuringRead(existingState GetRunResponse) {
}

func (a GetRunResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Run": reflect.TypeOf(Run{}),
	}
}

func (a GetRunResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Run": Run{}.ToAttrType(ctx),
		},
	}
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization types.String `tfsdk:"authorization" tf:"optional"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification" tf:"optional"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret types.String `tfsdk:"secret" tf:"optional"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url" tf:""`
}

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpec) {
}

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpec) {
}

func (a HttpUrlSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a HttpUrlSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Authorization":         types.StringType,
			"EnableSslVerification": types.BoolType,
			"Secret":                types.StringType,
			"Url":                   types.StringType,
		},
	}
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification" tf:"optional"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpecWithoutSecret) {
}

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpecWithoutSecret) {
}

func (a HttpUrlSpecWithoutSecret) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a HttpUrlSpecWithoutSecret) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EnableSslVerification": types.BoolType,
			"Url":                   types.StringType,
		},
	}
}

type InputTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *InputTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan InputTag) {
}

func (newState *InputTag) SyncEffectiveFieldsDuringRead(existingState InputTag) {
}

func (a InputTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a InputTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken types.String `tfsdk:"access_token" tf:""`
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id" tf:""`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url" tf:"optional"`
}

func (newState *JobSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpec) {
}

func (newState *JobSpec) SyncEffectiveFieldsDuringRead(existingState JobSpec) {
}

func (a JobSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a JobSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessToken":  types.StringType,
			"JobId":        types.StringType,
			"WorkspaceUrl": types.StringType,
		},
	}
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id" tf:"optional"`
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	WorkspaceUrl types.String `tfsdk:"workspace_url" tf:"optional"`
}

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpecWithoutSecret) {
}

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState JobSpecWithoutSecret) {
}

func (a JobSpecWithoutSecret) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a JobSpecWithoutSecret) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"JobId":        types.StringType,
			"WorkspaceUrl": types.StringType,
		},
	}
}

// Get all artifacts
type ListArtifactsRequest struct {
	// Token indicating the page of artifact results to fetch. `page_token` is
	// not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken types.String `tfsdk:"-"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path types.String `tfsdk:"-"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use run_id instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *ListArtifactsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListArtifactsRequest) {
}

func (newState *ListArtifactsRequest) SyncEffectiveFieldsDuringRead(existingState ListArtifactsRequest) {
}

func (a ListArtifactsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListArtifactsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageToken": types.StringType,
			"Path":      types.StringType,
			"RunId":     types.StringType,
			"RunUuid":   types.StringType,
		},
	}
}

type ListArtifactsResponse struct {
	// File location and metadata for artifacts.
	Files types.List `tfsdk:"files" tf:"optional"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri" tf:"optional"`
}

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListArtifactsResponse) {
}

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringRead(existingState ListArtifactsResponse) {
}

func (a ListArtifactsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Files": reflect.TypeOf(FileInfo{}),
	}
}

func (a ListArtifactsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Files": basetypes.ListType{
				ElemType: FileInfo{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
			"RootUri":       types.StringType,
		},
	}
}

// List experiments
type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults types.Int64 `tfsdk:"-"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"-"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"-"`
}

func (newState *ListExperimentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExperimentsRequest) {
}

func (newState *ListExperimentsRequest) SyncEffectiveFieldsDuringRead(existingState ListExperimentsRequest) {
}

func (a ListExperimentsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListExperimentsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"MaxResults": types.Int64Type,
			"PageToken":  types.StringType,
			"ViewType":   types.StringType,
		},
	}
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments types.List `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExperimentsResponse) {
}

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState ListExperimentsResponse) {
}

func (a ListExperimentsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Experiments": reflect.TypeOf(Experiment{}),
	}
}

func (a ListExperimentsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Experiments": basetypes.ListType{
				ElemType: Experiment{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// List models
type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelsRequest) {
}

func (newState *ListModelsRequest) SyncEffectiveFieldsDuringRead(existingState ListModelsRequest) {
}

func (a ListModelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListModelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"MaxResults": types.Int64Type,
			"PageToken":  types.StringType,
		},
	}
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelsResponse) {
}

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelsResponse) {
}

func (a ListModelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RegisteredModels": reflect.TypeOf(Model{}),
	}
}

func (a ListModelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"RegisteredModels": basetypes.ListType{
				ElemType: Model{}.ToAttrType(ctx),
			},
		},
	}
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks" tf:"optional"`
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegistryWebhooks) {
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringRead(existingState ListRegistryWebhooks) {
}

func (a ListRegistryWebhooks) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Webhooks": reflect.TypeOf(RegistryWebhook{}),
	}
}

func (a ListRegistryWebhooks) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"Webhooks": basetypes.ListType{
				ElemType: RegistryWebhook{}.ToAttrType(ctx),
			},
		},
	}
}

// List transition requests
type ListTransitionRequestsRequest struct {
	// Name of the model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (newState *ListTransitionRequestsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTransitionRequestsRequest) {
}

func (newState *ListTransitionRequestsRequest) SyncEffectiveFieldsDuringRead(existingState ListTransitionRequestsRequest) {
}

func (a ListTransitionRequestsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListTransitionRequestsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":    types.StringType,
			"Version": types.StringType,
		},
	}
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests types.List `tfsdk:"requests" tf:"optional"`
}

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTransitionRequestsResponse) {
}

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringRead(existingState ListTransitionRequestsResponse) {
}

func (a ListTransitionRequestsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Requests": reflect.TypeOf(Activity{}),
	}
}

func (a ListTransitionRequestsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Requests": basetypes.ListType{
				ElemType: Activity{}.ToAttrType(ctx),
			},
		},
	}
}

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events types.List `tfsdk:"-"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	ModelName types.String `tfsdk:"-"`
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListWebhooksRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWebhooksRequest) {
}

func (newState *ListWebhooksRequest) SyncEffectiveFieldsDuringRead(existingState ListWebhooksRequest) {
}

func (a ListWebhooksRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Events": reflect.TypeOf(types.StringType),
	}
}

func (a ListWebhooksRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"ModelName": types.StringType,
			"PageToken": types.StringType,
		},
	}
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params types.List `tfsdk:"params" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *LogBatch) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogBatch) {
}

func (newState *LogBatch) SyncEffectiveFieldsDuringRead(existingState LogBatch) {
}

func (a LogBatch) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Metrics": reflect.TypeOf(Metric{}),
		"Params":  reflect.TypeOf(Param{}),
		"Tags":    reflect.TypeOf(RunTag{}),
	}
}

func (a LogBatch) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Metrics": basetypes.ListType{
				ElemType: Metric{}.ToAttrType(ctx),
			},
			"Params": basetypes.ListType{
				ElemType: Param{}.ToAttrType(ctx),
			},
			"RunId": types.StringType,
			"Tags": basetypes.ListType{
				ElemType: RunTag{}.ToAttrType(ctx),
			},
		},
	}
}

type LogBatchResponse struct {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogBatchResponse) {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringRead(existingState LogBatchResponse) {
}

func (a LogBatchResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogBatchResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogInputs struct {
	// Dataset inputs
	Datasets types.List `tfsdk:"datasets" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

func (newState *LogInputs) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogInputs) {
}

func (newState *LogInputs) SyncEffectiveFieldsDuringRead(existingState LogInputs) {
}

func (a LogInputs) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Datasets": reflect.TypeOf(DatasetInput{}),
	}
}

func (a LogInputs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Datasets": basetypes.ListType{
				ElemType: DatasetInput{}.ToAttrType(ctx),
			},
			"RunId": types.StringType,
		},
	}
}

type LogInputsResponse struct {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogInputsResponse) {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringRead(existingState LogInputsResponse) {
}

func (a LogInputsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogInputsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogMetric struct {
	// Name of the metric.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the metric. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Step at which to log the metric
	Step types.Int64 `tfsdk:"step" tf:"optional"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:""`
	// Double value of the metric being logged.
	Value types.Float64 `tfsdk:"value" tf:""`
}

func (newState *LogMetric) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogMetric) {
}

func (newState *LogMetric) SyncEffectiveFieldsDuringRead(existingState LogMetric) {
}

func (a LogMetric) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogMetric) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":       types.StringType,
			"RunId":     types.StringType,
			"RunUuid":   types.StringType,
			"Step":      types.Int64Type,
			"Timestamp": types.Int64Type,
			"Value":     types.Float64Type,
		},
	}
}

type LogMetricResponse struct {
}

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogMetricResponse) {
}

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringRead(existingState LogMetricResponse) {
}

func (a LogMetricResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogMetricResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

func (newState *LogModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogModel) {
}

func (newState *LogModel) SyncEffectiveFieldsDuringRead(existingState LogModel) {
}

func (a LogModel) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogModel) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelJson": types.StringType,
			"RunId":     types.StringType,
		},
	}
}

type LogModelResponse struct {
}

func (newState *LogModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogModelResponse) {
}

func (newState *LogModelResponse) SyncEffectiveFieldsDuringRead(existingState LogModelResponse) {
}

func (a LogModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the param. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *LogParam) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogParam) {
}

func (newState *LogParam) SyncEffectiveFieldsDuringRead(existingState LogParam) {
}

func (a LogParam) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogParam) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":     types.StringType,
			"RunId":   types.StringType,
			"RunUuid": types.StringType,
			"Value":   types.StringType,
		},
	}
}

type LogParamResponse struct {
}

func (newState *LogParamResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogParamResponse) {
}

func (newState *LogParamResponse) SyncEffectiveFieldsDuringRead(existingState LogParamResponse) {
}

func (a LogParamResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a LogParamResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Metric struct {
	// Key identifying this metric.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Step at which to log the metric.
	Step types.Int64 `tfsdk:"step" tf:"optional"`
	// The timestamp at which this metric was recorded.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`
	// Value associated with this metric.
	Value types.Float64 `tfsdk:"value" tf:"optional"`
}

func (newState *Metric) SyncEffectiveFieldsDuringCreateOrUpdate(plan Metric) {
}

func (newState *Metric) SyncEffectiveFieldsDuringRead(existingState Metric) {
}

func (a Metric) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Metric) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":       types.StringType,
			"Step":      types.Int64Type,
			"Timestamp": types.Int64Type,
			"Value":     types.Float64Type,
		},
	}
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Description of this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions types.List `tfsdk:"latest_versions" tf:"optional"`
	// Unique name for the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// User that created this `registered_model`
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *Model) SyncEffectiveFieldsDuringCreateOrUpdate(plan Model) {
}

func (newState *Model) SyncEffectiveFieldsDuringRead(existingState Model) {
}

func (a Model) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"LatestVersions": reflect.TypeOf(ModelVersion{}),
		"Tags":           reflect.TypeOf(ModelTag{}),
	}
}

func (a Model) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp":    types.Int64Type,
			"Description":          types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"LatestVersions": basetypes.ListType{
				ElemType: ModelVersion{}.ToAttrType(ctx),
			},
			"Name": types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelTag{}.ToAttrType(ctx),
			},
			"UserId": types.StringType,
		},
	}
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions types.List `tfsdk:"latest_versions" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Array of tags associated with the model.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDatabricks) {
}

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelDatabricks) {
}

func (a ModelDatabricks) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"LatestVersions": reflect.TypeOf(ModelVersion{}),
		"Tags":           reflect.TypeOf(ModelTag{}),
	}
}

func (a ModelDatabricks) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp":    types.Int64Type,
			"Description":          types.StringType,
			"Id":                   types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"LatestVersions": basetypes.ListType{
				ElemType: ModelVersion{}.ToAttrType(ctx),
			},
			"Name":            types.StringType,
			"PermissionLevel": types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelTag{}.ToAttrType(ctx),
			},
			"UserId": types.StringType,
		},
	}
}

type ModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ModelTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelTag) {
}

func (newState *ModelTag) SyncEffectiveFieldsDuringRead(existingState ModelTag) {
}

func (a ModelTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ModelTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Current stage for this `model_version`.
	CurrentStage types.String `tfsdk:"current_stage" tf:"optional"`
	// Description of this `model_version`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Unique name of the model
	Name types.String `tfsdk:"name" tf:"optional"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// Run Link: Direct link to the run that generated this version
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source types.String `tfsdk:"source" tf:"optional"`
	// Current status of `model_version`
	Status types.String `tfsdk:"status" tf:"optional"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// User that created this `model_version`.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Model's version number.
	Version types.String `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersion) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersion) {
}

func (newState *ModelVersion) SyncEffectiveFieldsDuringRead(existingState ModelVersion) {
}

func (a ModelVersion) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

func (a ModelVersion) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp":    types.Int64Type,
			"CurrentStage":         types.StringType,
			"Description":          types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"Name":                 types.StringType,
			"RunId":                types.StringType,
			"RunLink":              types.StringType,
			"Source":               types.StringType,
			"Status":               types.StringType,
			"StatusMessage":        types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelVersionTag{}.ToAttrType(ctx),
			},
			"UserId":  types.StringType,
			"Version": types.StringType,
		},
	}
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	CurrentStage types.String `tfsdk:"current_stage" tf:"optional"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source types.String `tfsdk:"source" tf:"optional"`
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	Status types.String `tfsdk:"status" tf:"optional"`
	// Details on the current status, for example why registration failed.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
	// Array of tags that are associated with the model version.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionDatabricks) {
}

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelVersionDatabricks) {
}

func (a ModelVersionDatabricks) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

func (a ModelVersionDatabricks) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp":    types.Int64Type,
			"CurrentStage":         types.StringType,
			"Description":          types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"Name":                 types.StringType,
			"PermissionLevel":      types.StringType,
			"RunId":                types.StringType,
			"RunLink":              types.StringType,
			"Source":               types.StringType,
			"Status":               types.StringType,
			"StatusMessage":        types.StringType,
			"Tags": basetypes.ListType{
				ElemType: ModelVersionTag{}.ToAttrType(ctx),
			},
			"UserId":  types.StringType,
			"Version": types.StringType,
		},
	}
}

type ModelVersionTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionTag) {
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringRead(existingState ModelVersionTag) {
}

func (a ModelVersionTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ModelVersionTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type Param struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Value associated with this param.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *Param) SyncEffectiveFieldsDuringCreateOrUpdate(plan Param) {
}

func (newState *Param) SyncEffectiveFieldsDuringRead(existingState Param) {
}

func (a Param) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Param) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlRequest) {
}

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlRequest) {
}

func (a RegisteredModelAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RegisteredModelAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"GroupName":            types.StringType,
			"PermissionLevel":      types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
}

type RegisteredModelAccessControlResponse struct {
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

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlResponse) {
}

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlResponse) {
}

func (a RegisteredModelAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AllPermissions": reflect.TypeOf(RegisteredModelPermission{}),
	}
}

func (a RegisteredModelAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AllPermissions": basetypes.ListType{
				ElemType: RegisteredModelPermission{}.ToAttrType(ctx),
			},
			"DisplayName":          types.StringType,
			"GroupName":            types.StringType,
			"ServicePrincipalName": types.StringType,
			"UserName":             types.StringType,
		},
	}
}

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermission) {
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermission) {
}

func (a RegisteredModelPermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"InheritedFromObject": reflect.TypeOf(types.StringType),
	}
}

func (a RegisteredModelPermission) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Inherited": types.BoolType,
			"InheritedFromObject": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PermissionLevel": types.StringType,
		},
	}
}

type RegisteredModelPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissions) {
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissions) {
}

func (a RegisteredModelPermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(RegisteredModelAccessControlResponse{}),
	}
}

func (a RegisteredModelPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: RegisteredModelAccessControlResponse{}.ToAttrType(ctx),
			},
			"ObjectId":   types.StringType,
			"ObjectType": types.StringType,
		},
	}
}

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissionsDescription) {
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissionsDescription) {
}

func (a RegisteredModelPermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RegisteredModelPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description":     types.StringType,
			"PermissionLevel": types.StringType,
		},
	}
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *RegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissionsRequest) {
}

func (newState *RegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissionsRequest) {
}

func (a RegisteredModelPermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(RegisteredModelAccessControlRequest{}),
	}
}

func (a RegisteredModelPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AccessControlList": basetypes.ListType{
				ElemType: RegisteredModelAccessControlRequest{}.ToAttrType(ctx),
			},
			"RegisteredModelId": types.StringType,
		},
	}
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events" tf:"optional"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:"optional"`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegistryWebhook) {
}

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringRead(existingState RegistryWebhook) {
}

func (a RegistryWebhook) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Events":      reflect.TypeOf(types.StringType),
		"HttpUrlSpec": reflect.TypeOf(HttpUrlSpecWithoutSecret{}),
		"JobSpec":     reflect.TypeOf(JobSpecWithoutSecret{}),
	}
}

func (a RegistryWebhook) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp": types.Int64Type,
			"Description":       types.StringType,
			"Events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"HttpUrlSpec":          HttpUrlSpecWithoutSecret{}.ToAttrType(ctx),
			"Id":                   types.StringType,
			"JobSpec":              JobSpecWithoutSecret{}.ToAttrType(ctx),
			"LastUpdatedTimestamp": types.Int64Type,
			"ModelName":            types.StringType,
			"Status":               types.StringType,
		},
	}
}

type RejectTransitionRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *RejectTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RejectTransitionRequest) {
}

func (newState *RejectTransitionRequest) SyncEffectiveFieldsDuringRead(existingState RejectTransitionRequest) {
}

func (a RejectTransitionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RejectTransitionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": types.StringType,
			"Name":    types.StringType,
			"Stage":   types.StringType,
			"Version": types.StringType,
		},
	}
}

type RejectTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity types.List `tfsdk:"activity" tf:"optional,object"`
}

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RejectTransitionRequestResponse) {
}

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState RejectTransitionRequestResponse) {
}

func (a RejectTransitionRequestResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Activity": reflect.TypeOf(Activity{}),
	}
}

func (a RejectTransitionRequestResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Activity": Activity{}.ToAttrType(ctx),
		},
	}
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

func (newState *RenameModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RenameModelRequest) {
}

func (newState *RenameModelRequest) SyncEffectiveFieldsDuringRead(existingState RenameModelRequest) {
}

func (a RenameModelRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RenameModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name":    types.StringType,
			"NewName": types.StringType,
		},
	}
}

type RenameModelResponse struct {
	RegisteredModel types.List `tfsdk:"registered_model" tf:"optional,object"`
}

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RenameModelResponse) {
}

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringRead(existingState RenameModelResponse) {
}

func (a RenameModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RegisteredModel": reflect.TypeOf(Model{}),
	}
}

func (a RenameModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RegisteredModel": Model{}.ToAttrType(ctx),
		},
	}
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

func (newState *RestoreExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreExperiment) {
}

func (newState *RestoreExperiment) SyncEffectiveFieldsDuringRead(existingState RestoreExperiment) {
}

func (a RestoreExperiment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreExperiment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
		},
	}
}

type RestoreExperimentResponse struct {
}

func (newState *RestoreExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreExperimentResponse) {
}

func (newState *RestoreExperimentResponse) SyncEffectiveFieldsDuringRead(existingState RestoreExperimentResponse) {
}

func (a RestoreExperimentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreExperimentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *RestoreRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRun) {
}

func (newState *RestoreRun) SyncEffectiveFieldsDuringRead(existingState RestoreRun) {
}

func (a RestoreRun) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreRun) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunId": types.StringType,
		},
	}
}

type RestoreRunResponse struct {
}

func (newState *RestoreRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRunResponse) {
}

func (newState *RestoreRunResponse) SyncEffectiveFieldsDuringRead(existingState RestoreRunResponse) {
}

func (a RestoreRunResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreRunResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs" tf:"optional"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis types.Int64 `tfsdk:"min_timestamp_millis" tf:""`
}

func (newState *RestoreRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRuns) {
}

func (newState *RestoreRuns) SyncEffectiveFieldsDuringRead(existingState RestoreRuns) {
}

func (a RestoreRuns) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreRuns) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId":       types.StringType,
			"MaxRuns":            types.Int64Type,
			"MinTimestampMillis": types.Int64Type,
		},
	}
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored types.Int64 `tfsdk:"runs_restored" tf:"optional"`
}

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRunsResponse) {
}

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringRead(existingState RestoreRunsResponse) {
}

func (a RestoreRunsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RestoreRunsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunsRestored": types.Int64Type,
		},
	}
}

type Run struct {
	// Run data.
	Data types.List `tfsdk:"data" tf:"optional,object"`
	// Run metadata.
	Info types.List `tfsdk:"info" tf:"optional,object"`
	// Run inputs.
	Inputs types.List `tfsdk:"inputs" tf:"optional,object"`
}

func (newState *Run) SyncEffectiveFieldsDuringCreateOrUpdate(plan Run) {
}

func (newState *Run) SyncEffectiveFieldsDuringRead(existingState Run) {
}

func (a Run) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Data":   reflect.TypeOf(RunData{}),
		"Info":   reflect.TypeOf(RunInfo{}),
		"Inputs": reflect.TypeOf(RunInputs{}),
	}
}

func (a Run) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Data":   RunData{}.ToAttrType(ctx),
			"Info":   RunInfo{}.ToAttrType(ctx),
			"Inputs": RunInputs{}.ToAttrType(ctx),
		},
	}
}

type RunData struct {
	// Run metrics.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Run parameters.
	Params types.List `tfsdk:"params" tf:"optional"`
	// Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *RunData) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunData) {
}

func (newState *RunData) SyncEffectiveFieldsDuringRead(existingState RunData) {
}

func (a RunData) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Metrics": reflect.TypeOf(Metric{}),
		"Params":  reflect.TypeOf(Param{}),
		"Tags":    reflect.TypeOf(RunTag{}),
	}
}

func (a RunData) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Metrics": basetypes.ListType{
				ElemType: Metric{}.ToAttrType(ctx),
			},
			"Params": basetypes.ListType{
				ElemType: Param{}.ToAttrType(ctx),
			},
			"Tags": basetypes.ListType{
				ElemType: RunTag{}.ToAttrType(ctx),
			},
		},
	}
}

type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like `s3://bucket/directory` or `dbfs:/my/directory`. If not set, the
	// local `./mlruns` directory is chosen.
	ArtifactUri types.String `tfsdk:"artifact_uri" tf:"optional"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage types.String `tfsdk:"lifecycle_stage" tf:"optional"`
	// Unique identifier for the run.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Current status of the run.
	Status types.String `tfsdk:"status" tf:"optional"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *RunInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInfo) {
}

func (newState *RunInfo) SyncEffectiveFieldsDuringRead(existingState RunInfo) {
}

func (a RunInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RunInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArtifactUri":    types.StringType,
			"EndTime":        types.Int64Type,
			"ExperimentId":   types.StringType,
			"LifecycleStage": types.StringType,
			"RunId":          types.StringType,
			"RunUuid":        types.StringType,
			"StartTime":      types.Int64Type,
			"Status":         types.StringType,
			"UserId":         types.StringType,
		},
	}
}

type RunInputs struct {
	// Run metrics.
	DatasetInputs types.List `tfsdk:"dataset_inputs" tf:"optional"`
}

func (newState *RunInputs) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInputs) {
}

func (newState *RunInputs) SyncEffectiveFieldsDuringRead(existingState RunInputs) {
}

func (a RunInputs) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DatasetInputs": reflect.TypeOf(DatasetInput{}),
	}
}

func (a RunInputs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DatasetInputs": basetypes.ListType{
				ElemType: DatasetInput{}.ToAttrType(ctx),
			},
		},
	}
}

type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *RunTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunTag) {
}

func (newState *RunTag) SyncEffectiveFieldsDuringRead(existingState RunTag) {
}

func (a RunTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RunTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": types.StringType,
		},
	}
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	Filter types.String `tfsdk:"filter" tf:"optional"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults types.Int64 `tfsdk:"max_results" tf:"optional"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy types.List `tfsdk:"order_by" tf:"optional"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"view_type" tf:"optional"`
}

func (newState *SearchExperiments) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchExperiments) {
}

func (newState *SearchExperiments) SyncEffectiveFieldsDuringRead(existingState SearchExperiments) {
}

func (a SearchExperiments) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"OrderBy": reflect.TypeOf(types.StringType),
	}
}

func (a SearchExperiments) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter":     types.StringType,
			"MaxResults": types.Int64Type,
			"OrderBy": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PageToken": types.StringType,
			"ViewType":  types.StringType,
		},
	}
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchExperimentsResponse) {
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState SearchExperimentsResponse) {
}

func (a SearchExperimentsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Experiments": reflect.TypeOf(Experiment{}),
	}
}

func (a SearchExperimentsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Experiments": basetypes.ListType{
				ElemType: Experiment{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// Searches model versions
type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults types.Int64 `tfsdk:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy types.List `tfsdk:"-"`
	// Pagination token to go to next page based on previous search query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *SearchModelVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelVersionsRequest) {
}

func (newState *SearchModelVersionsRequest) SyncEffectiveFieldsDuringRead(existingState SearchModelVersionsRequest) {
}

func (a SearchModelVersionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"OrderBy": reflect.TypeOf(types.StringType),
	}
}

func (a SearchModelVersionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter":     types.StringType,
			"MaxResults": types.Int64Type,
			"OrderBy": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PageToken": types.StringType,
		},
	}
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelVersionsResponse) {
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelVersionsResponse) {
}

func (a SearchModelVersionsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ModelVersions": reflect.TypeOf(ModelVersion{}),
	}
}

func (a SearchModelVersionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelVersions": basetypes.ListType{
				ElemType: ModelVersion{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
		},
	}
}

// Search models
type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy types.List `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *SearchModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelsRequest) {
}

func (newState *SearchModelsRequest) SyncEffectiveFieldsDuringRead(existingState SearchModelsRequest) {
}

func (a SearchModelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"OrderBy": reflect.TypeOf(types.StringType),
	}
}

func (a SearchModelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Filter":     types.StringType,
			"MaxResults": types.Int64Type,
			"OrderBy": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PageToken": types.StringType,
		},
	}
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelsResponse) {
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelsResponse) {
}

func (a SearchModelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RegisteredModels": reflect.TypeOf(Model{}),
	}
}

func (a SearchModelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"RegisteredModels": basetypes.ListType{
				ElemType: Model{}.ToAttrType(ctx),
			},
		},
	}
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds types.List `tfsdk:"experiment_ids" tf:"optional"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	Filter types.String `tfsdk:"filter" tf:"optional"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults types.Int64 `tfsdk:"max_results" tf:"optional"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional "DESC" or "ASC" annotation, where "ASC" is the
	// default. Example: ["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"] Tiebreaks are done by start_time DESC followed by run_id
	// for runs with the same start time (and this is the default ordering
	// criterion if order_by is not provided).
	OrderBy types.List `tfsdk:"order_by" tf:"optional"`
	// Token for the current page of runs.
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType types.String `tfsdk:"run_view_type" tf:"optional"`
}

func (newState *SearchRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchRuns) {
}

func (newState *SearchRuns) SyncEffectiveFieldsDuringRead(existingState SearchRuns) {
}

func (a SearchRuns) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ExperimentIds": reflect.TypeOf(types.StringType),
		"OrderBy":       reflect.TypeOf(types.StringType),
	}
}

func (a SearchRuns) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentIds": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Filter":     types.StringType,
			"MaxResults": types.Int64Type,
			"OrderBy": basetypes.ListType{
				ElemType: types.StringType,
			},
			"PageToken":   types.StringType,
			"RunViewType": types.StringType,
		},
	}
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs" tf:"optional"`
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchRunsResponse) {
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringRead(existingState SearchRunsResponse) {
}

func (a SearchRunsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Runs": reflect.TypeOf(Run{}),
	}
}

func (a SearchRunsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"Runs": basetypes.ListType{
				ElemType: Run{}.ToAttrType(ctx),
			},
		},
	}
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetExperimentTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetExperimentTag) {
}

func (newState *SetExperimentTag) SyncEffectiveFieldsDuringRead(existingState SetExperimentTag) {
}

func (a SetExperimentTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetExperimentTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
			"Key":          types.StringType,
			"Value":        types.StringType,
		},
	}
}

type SetExperimentTagResponse struct {
}

func (newState *SetExperimentTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetExperimentTagResponse) {
}

func (newState *SetExperimentTagResponse) SyncEffectiveFieldsDuringRead(existingState SetExperimentTagResponse) {
}

func (a SetExperimentTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetExperimentTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// Unique name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetModelTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelTagRequest) {
}

func (newState *SetModelTagRequest) SyncEffectiveFieldsDuringRead(existingState SetModelTagRequest) {
}

func (a SetModelTagRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetModelTagRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Name":  types.StringType,
			"Value": types.StringType,
		},
	}
}

type SetModelTagResponse struct {
}

func (newState *SetModelTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelTagResponse) {
}

func (newState *SetModelTagResponse) SyncEffectiveFieldsDuringRead(existingState SetModelTagResponse) {
}

func (a SetModelTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetModelTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// Unique name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
	// Model version number.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *SetModelVersionTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelVersionTagRequest) {
}

func (newState *SetModelVersionTagRequest) SyncEffectiveFieldsDuringRead(existingState SetModelVersionTagRequest) {
}

func (a SetModelVersionTagRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetModelVersionTagRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":     types.StringType,
			"Name":    types.StringType,
			"Value":   types.StringType,
			"Version": types.StringType,
		},
	}
}

type SetModelVersionTagResponse struct {
}

func (newState *SetModelVersionTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelVersionTagResponse) {
}

func (newState *SetModelVersionTagResponse) SyncEffectiveFieldsDuringRead(existingState SetModelVersionTagResponse) {
}

func (a SetModelVersionTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetModelVersionTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetTag struct {
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the tag. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetTag) {
}

func (newState *SetTag) SyncEffectiveFieldsDuringRead(existingState SetTag) {
}

func (a SetTag) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":     types.StringType,
			"RunId":   types.StringType,
			"RunUuid": types.StringType,
			"Value":   types.StringType,
		},
	}
}

type SetTagResponse struct {
}

func (newState *SetTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetTagResponse) {
}

func (newState *SetTagResponse) SyncEffectiveFieldsDuringRead(existingState SetTagResponse) {
}

func (a SetTagResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SetTagResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	Body types.String `tfsdk:"body" tf:"optional"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code" tf:"optional"`
}

func (newState *TestRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhook) {
}

func (newState *TestRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhook) {
}

func (a TestRegistryWebhook) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TestRegistryWebhook) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Body":       types.StringType,
			"StatusCode": types.Int64Type,
		},
	}
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event types.String `tfsdk:"event" tf:"optional"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *TestRegistryWebhookRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhookRequest) {
}

func (newState *TestRegistryWebhookRequest) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhookRequest) {
}

func (a TestRegistryWebhookRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TestRegistryWebhookRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Event": types.StringType,
			"Id":    types.StringType,
		},
	}
}

type TestRegistryWebhookResponse struct {
	// Test webhook response object.
	Webhook types.List `tfsdk:"webhook" tf:"optional,object"`
}

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhookResponse) {
}

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhookResponse) {
}

func (a TestRegistryWebhookResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Webhook": reflect.TypeOf(TestRegistryWebhook{}),
	}
}

func (a TestRegistryWebhookResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Webhook": TestRegistryWebhook{}.ToAttrType(ctx),
		},
	}
}

type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions" tf:""`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *TransitionModelVersionStageDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionModelVersionStageDatabricks) {
}

func (newState *TransitionModelVersionStageDatabricks) SyncEffectiveFieldsDuringRead(existingState TransitionModelVersionStageDatabricks) {
}

func (a TransitionModelVersionStageDatabricks) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a TransitionModelVersionStageDatabricks) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ArchiveExistingVersions": types.BoolType,
			"Comment":                 types.StringType,
			"Name":                    types.StringType,
			"Stage":                   types.StringType,
			"Version":                 types.StringType,
		},
	}
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions" tf:"optional"`
	// User-provided comment associated with the transition request.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage types.String `tfsdk:"to_stage" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *TransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionRequest) {
}

func (newState *TransitionRequest) SyncEffectiveFieldsDuringRead(existingState TransitionRequest) {
}

func (a TransitionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AvailableActions": reflect.TypeOf(types.StringType),
	}
}

func (a TransitionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"AvailableActions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"Comment":           types.StringType,
			"CreationTimestamp": types.Int64Type,
			"ToStage":           types.StringType,
			"UserId":            types.StringType,
		},
	}
}

type TransitionStageResponse struct {
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionStageResponse) {
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringRead(existingState TransitionStageResponse) {
}

func (a TransitionStageResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ModelVersion": reflect.TypeOf(ModelVersionDatabricks{}),
	}
}

func (a TransitionStageResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ModelVersion": ModelVersionDatabricks{}.ToAttrType(ctx),
		},
	}
}

type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *UpdateComment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateComment) {
}

func (newState *UpdateComment) SyncEffectiveFieldsDuringRead(existingState UpdateComment) {
}

func (a UpdateComment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateComment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": types.StringType,
			"Id":      types.StringType,
		},
	}
}

type UpdateCommentResponse struct {
	// Comment details.
	Comment types.List `tfsdk:"comment" tf:"optional,object"`
}

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCommentResponse) {
}

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateCommentResponse) {
}

func (a UpdateCommentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Comment": reflect.TypeOf(CommentObject{}),
	}
}

func (a UpdateCommentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Comment": CommentObject{}.ToAttrType(ctx),
		},
	}
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

func (newState *UpdateExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExperiment) {
}

func (newState *UpdateExperiment) SyncEffectiveFieldsDuringRead(existingState UpdateExperiment) {
}

func (a UpdateExperiment) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateExperiment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ExperimentId": types.StringType,
			"NewName":      types.StringType,
		},
	}
}

type UpdateExperimentResponse struct {
}

func (newState *UpdateExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExperimentResponse) {
}

func (newState *UpdateExperimentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateExperimentResponse) {
}

func (a UpdateExperimentResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateExperimentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *UpdateModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelRequest) {
}

func (newState *UpdateModelRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelRequest) {
}

func (a UpdateModelRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Name":        types.StringType,
		},
	}
}

type UpdateModelResponse struct {
}

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelResponse) {
}

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelResponse) {
}

func (a UpdateModelResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateModelResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Name of the registered model
	Name types.String `tfsdk:"name" tf:""`
	// Model version number
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest) {
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest) {
}

func (a UpdateModelVersionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Name":        types.StringType,
			"Version":     types.StringType,
		},
	}
}

type UpdateModelVersionResponse struct {
}

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionResponse) {
}

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionResponse) {
}

func (a UpdateModelVersionResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateModelVersionResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events" tf:"optional"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *UpdateRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegistryWebhook) {
}

func (newState *UpdateRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState UpdateRegistryWebhook) {
}

func (a UpdateRegistryWebhook) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Events":      reflect.TypeOf(types.StringType),
		"HttpUrlSpec": reflect.TypeOf(HttpUrlSpec{}),
		"JobSpec":     reflect.TypeOf(JobSpec{}),
	}
}

func (a UpdateRegistryWebhook) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Description": types.StringType,
			"Events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"HttpUrlSpec": HttpUrlSpec{}.ToAttrType(ctx),
			"Id":          types.StringType,
			"JobSpec":     JobSpec{}.ToAttrType(ctx),
			"Status":      types.StringType,
		},
	}
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// ID of the run to update. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run to update.. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Updated status of the run.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *UpdateRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRun) {
}

func (newState *UpdateRun) SyncEffectiveFieldsDuringRead(existingState UpdateRun) {
}

func (a UpdateRun) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateRun) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndTime": types.Int64Type,
			"RunId":   types.StringType,
			"RunUuid": types.StringType,
			"Status":  types.StringType,
		},
	}
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo types.List `tfsdk:"run_info" tf:"optional,object"`
}

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRunResponse) {
}

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringRead(existingState UpdateRunResponse) {
}

func (a UpdateRunResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"RunInfo": reflect.TypeOf(RunInfo{}),
	}
}

func (a UpdateRunResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"RunInfo": RunInfo{}.ToAttrType(ctx),
		},
	}
}

type UpdateWebhookResponse struct {
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWebhookResponse) {
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState UpdateWebhookResponse) {
}

func (a UpdateWebhookResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateWebhookResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}
