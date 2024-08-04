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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
type ActivityAction string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition.
//
// * `REQUESTED_TRANSITION`: User requested the corresponding stage transition.
//
// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
//
// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
//
// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
//
// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
// archiving existing model versions in a stage.
type ActivityType string

// User applied the corresponding stage transition.
const ActivityTypeAppliedTransition ActivityType = `APPLIED_TRANSITION`

// User approved the corresponding stage transition.
const ActivityTypeApprovedRequest ActivityType = `APPROVED_REQUEST`

// User cancelled an existing transition request.
const ActivityTypeCancelledRequest ActivityType = `CANCELLED_REQUEST`

const ActivityTypeNewComment ActivityType = `NEW_COMMENT`

// User rejected the coressponding stage transition.
const ActivityTypeRejectedRequest ActivityType = `REJECTED_REQUEST`

// User requested the corresponding stage transition.
const ActivityTypeRequestedTransition ActivityType = `REQUESTED_TRANSITION`

// For events performed as a side effect, such as archiving existing model
// versions in a stage.
const ActivityTypeSystemTransition ActivityType = `SYSTEM_TRANSITION`

// String representation for [fmt.Print]
func (f *ActivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityType) Set(v string) error {
	switch v {
	case `APPLIED_TRANSITION`, `APPROVED_REQUEST`, `CANCELLED_REQUEST`, `NEW_COMMENT`, `REJECTED_REQUEST`, `REQUESTED_TRANSITION`, `SYSTEM_TRANSITION`:
		*f = ActivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPLIED_TRANSITION", "APPROVED_REQUEST", "CANCELLED_REQUEST", "NEW_COMMENT", "REJECTED_REQUEST", "REQUESTED_TRANSITION", "SYSTEM_TRANSITION"`, v)
	}
}

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
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

type ApproveTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity *Activity `tfsdk:"activity" tf:"optional"`
}

// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `DELETE_COMMENT`, `EDIT_COMMENT`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE_COMMENT", "EDIT_COMMENT"`, v)
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []types.String `tfsdk:"available_actions" tf:"optional"`
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

type CreateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

type CreateCommentResponse struct {
	// Comment details.
	Comment *CommentObject `tfsdk:"comment" tf:"optional"`
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
	Tags []ExperimentTag `tfsdk:"tags" tf:"optional"`
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Register models under this name
	Name types.String `tfsdk:"name" tf:""`
	// Additional metadata for registered model.
	Tags []ModelTag `tfsdk:"tags" tf:"optional"`
}

type CreateModelResponse struct {
	RegisteredModel *Model `tfsdk:"registered_model" tf:"optional"`
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
	Tags []ModelVersionTag `tfsdk:"tags" tf:"optional"`
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion *ModelVersion `tfsdk:"model_version" tf:"optional"`
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
	Events []types.String `tfsdk:"events" tf:""`

	HttpUrlSpec *HttpUrlSpec `tfsdk:"http_url_spec" tf:"optional"`

	JobSpec *JobSpec `tfsdk:"job_spec" tf:"optional"`
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

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Additional metadata for run.
	Tags []RunTag `tfsdk:"tags" tf:"optional"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `tfsdk:"run" tf:"optional"`
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

type CreateTransitionRequestResponse struct {
	// Transition request details.
	Request *TransitionRequest `tfsdk:"request" tf:"optional"`
}

type CreateWebhookResponse struct {
	Webhook *RegistryWebhook `tfsdk:"webhook" tf:"optional"`
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

type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset *Dataset `tfsdk:"dataset" tf:"optional"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags []InputTag `tfsdk:"tags" tf:"optional"`
}

// Delete a comment
type DeleteCommentRequest struct {
	Id types.String `tfsdk:"-" url:"id"`
}

type DeleteCommentResponse struct {
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

type DeleteExperimentResponse struct {
}

// Delete a model
type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-" url:"name"`
}

type DeleteModelResponse struct {
}

// Delete a model tag
type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-" url:"name"`
}

type DeleteModelTagResponse struct {
}

// Delete a model version.
type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-" url:"name"`
	// Model version number
	Version types.String `tfsdk:"-" url:"version"`
}

type DeleteModelVersionResponse struct {
}

// Delete a model version tag
type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-" url:"name"`
	// Model version number that the tag was logged under.
	Version types.String `tfsdk:"-" url:"version"`
}

type DeleteModelVersionTagResponse struct {
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.String `tfsdk:"run_id" tf:""`
}

type DeleteRunResponse struct {
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

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted types.Int64 `tfsdk:"runs_deleted" tf:"optional"`
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:""`
}

type DeleteTagResponse struct {
}

// Delete a transition request
type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"-" url:"comment,omitempty"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator types.String `tfsdk:"-" url:"creator"`
	// Name of the model.
	Name types.String `tfsdk:"-" url:"name"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"-" url:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"-" url:"version"`
}

type DeleteTransitionRequestResponse struct {
}

type DeleteTransitionRequestStage string

const DeleteTransitionRequestStageArchived DeleteTransitionRequestStage = `Archived`

const DeleteTransitionRequestStageNone DeleteTransitionRequestStage = `None`

const DeleteTransitionRequestStageProduction DeleteTransitionRequestStage = `Production`

const DeleteTransitionRequestStageStaging DeleteTransitionRequestStage = `Staging`

// String representation for [fmt.Print]
func (f *DeleteTransitionRequestStage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeleteTransitionRequestStage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = DeleteTransitionRequestStage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns DeleteTransitionRequestStage to satisfy [pflag.Value] interface
func (f *DeleteTransitionRequestStage) Type() string {
	return "DeleteTransitionRequestStage"
}

// Delete a webhook
type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-" url:"id,omitempty"`
}

type DeleteWebhookResponse struct {
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
	Tags []ExperimentTag `tfsdk:"tags" tf:"optional"`
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

type ExperimentAccessControlResponse struct {
	// All permissions.
	AllPermissions []ExperimentPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type ExperimentPermissionLevel string

const ExperimentPermissionLevelCanEdit ExperimentPermissionLevel = `CAN_EDIT`

const ExperimentPermissionLevelCanManage ExperimentPermissionLevel = `CAN_MANAGE`

const ExperimentPermissionLevelCanRead ExperimentPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *ExperimentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExperimentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = ExperimentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

type ExperimentPermissions struct {
	AccessControlList []ExperimentAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type ExperimentPermissionsRequest struct {
	AccessControlList []ExperimentAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-" url:"-"`
}

type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

type FileInfo struct {
	// Size in bytes. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size" tf:"optional"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir" tf:"optional"`
	// Path relative to the root artifact directory run.
	Path types.String `tfsdk:"path" tf:"optional"`
}

// Get metadata
type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-" url:"experiment_name"`
}

// Get experiment permission levels
type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-" url:"-"`
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ExperimentPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get experiment permissions
type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-" url:"-"`
}

// Get an experiment
type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"-" url:"experiment_id"`
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment *Experiment `tfsdk:"experiment" tf:"optional"`
}

// Get history of a given metric within a run
type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// Name of the metric.
	MetricKey types.String `tfsdk:"-" url:"metric_key"`
	// Token indicating the page of metric histories to fetch.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId types.String `tfsdk:"-" url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run from which to fetch metric
	// values. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-" url:"run_uuid,omitempty"`
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
	// List of stages.
	Stages []types.String `tfsdk:"stages" tf:"optional"`
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions []ModelVersion `tfsdk:"model_versions" tf:"optional"`
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics []Metric `tfsdk:"metrics" tf:"optional"`
	// Token that can be used to retrieve the next page of metric history
	// results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-" url:"name"`
}

type GetModelResponse struct {
	RegisteredModelDatabricks *ModelDatabricks `tfsdk:"registered_model_databricks" tf:"optional"`
}

// Get a model version URI
type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-" url:"name"`
	// Model version number
	Version types.String `tfsdk:"-" url:"version"`
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri types.String `tfsdk:"artifact_uri" tf:"optional"`
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-" url:"name"`
	// Model version number
	Version types.String `tfsdk:"-" url:"version"`
}

type GetModelVersionResponse struct {
	ModelVersion *ModelVersion `tfsdk:"model_version" tf:"optional"`
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RegisteredModelPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get registered model permissions
type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-" url:"-"`
}

// Get a run
type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId types.String `tfsdk:"-" url:"run_id"`
	// [Deprecated, use run_id instead] ID of the run to fetch. This field will
	// be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-" url:"run_uuid,omitempty"`
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run *Run `tfsdk:"run" tf:"optional"`
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

type InputTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
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

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id" tf:"optional"`
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	WorkspaceUrl types.String `tfsdk:"workspace_url" tf:"optional"`
}

// Get all artifacts
type ListArtifactsRequest struct {
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path types.String `tfsdk:"-" url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId types.String `tfsdk:"-" url:"run_id,omitempty"`
	// [Deprecated, use run_id instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-" url:"run_uuid,omitempty"`
}

type ListArtifactsResponse struct {
	// File location and metadata for artifacts.
	Files []FileInfo `tfsdk:"files" tf:"optional"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri" tf:"optional"`
}

// List experiments
type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"-" url:"view_type,omitempty"`
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []Experiment `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List models
type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels []Model `tfsdk:"registered_models" tf:"optional"`
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Array of registry webhooks.
	Webhooks []RegistryWebhook `tfsdk:"webhooks" tf:"optional"`
}

// List transition requests
type ListTransitionRequestsRequest struct {
	// Name of the model.
	Name types.String `tfsdk:"-" url:"name"`
	// Version of the model.
	Version types.String `tfsdk:"-" url:"version"`
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests []Activity `tfsdk:"requests" tf:"optional"`
}

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []types.String `tfsdk:"-" url:"events,omitempty"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	ModelName types.String `tfsdk:"-" url:"model_name,omitempty"`
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics []Metric `tfsdk:"metrics" tf:"optional"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params []Param `tfsdk:"params" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags []RunTag `tfsdk:"tags" tf:"optional"`
}

type LogBatchResponse struct {
}

type LogInputs struct {
	// Dataset inputs
	Datasets []DatasetInput `tfsdk:"datasets" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

type LogInputsResponse struct {
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

type LogMetricResponse struct {
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

type LogModelResponse struct {
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

type LogParamResponse struct {
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
	LatestVersions []ModelVersion `tfsdk:"latest_versions" tf:"optional"`
	// Unique name for the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags []ModelTag `tfsdk:"tags" tf:"optional"`
	// User that created this `registered_model`
	UserId types.String `tfsdk:"user_id" tf:"optional"`
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
	LatestVersions []ModelVersion `tfsdk:"latest_versions" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Array of tags associated with the model.
	Tags []ModelTag `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

type ModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
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
	Tags []ModelVersionTag `tfsdk:"tags" tf:"optional"`
	// User that created this `model_version`.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Model's version number.
	Version types.String `tfsdk:"version" tf:"optional"`
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
	Tags []ModelVersionTag `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:"optional"`
}

// Current status of `model_version`
type ModelVersionStatus string

const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

const ModelVersionStatusReady ModelVersionStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

type ModelVersionTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

type Param struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Value associated with this param.
	Value types.String `tfsdk:"value" tf:"optional"`
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
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

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	AllPermissions []RegisteredModelPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type RegisteredModelPermissionLevel string

const RegisteredModelPermissionLevelCanEdit RegisteredModelPermissionLevel = `CAN_EDIT`

const RegisteredModelPermissionLevelCanManage RegisteredModelPermissionLevel = `CAN_MANAGE`

const RegisteredModelPermissionLevelCanManageProductionVersions RegisteredModelPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const RegisteredModelPermissionLevelCanManageStagingVersions RegisteredModelPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const RegisteredModelPermissionLevelCanRead RegisteredModelPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *RegisteredModelPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegisteredModelPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = RegisteredModelPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

type RegisteredModelPermissions struct {
	AccessControlList []RegisteredModelAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList []RegisteredModelAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-" url:"-"`
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
	Events []types.String `tfsdk:"events" tf:"optional"`

	HttpUrlSpec *HttpUrlSpecWithoutSecret `tfsdk:"http_url_spec" tf:"optional"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:"optional"`

	JobSpec *JobSpecWithoutSecret `tfsdk:"job_spec" tf:"optional"`
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

type RegistryWebhookEvent string

const RegistryWebhookEventCommentCreated RegistryWebhookEvent = `COMMENT_CREATED`

const RegistryWebhookEventModelVersionCreated RegistryWebhookEvent = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEvent = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEvent = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEvent = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_STAGING_CREATED`

// String representation for [fmt.Print]
func (f *RegistryWebhookEvent) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookEvent) Set(v string) error {
	switch v {
	case `COMMENT_CREATED`, `MODEL_VERSION_CREATED`, `MODEL_VERSION_TAG_SET`, `MODEL_VERSION_TRANSITIONED_STAGE`, `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`, `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`, `MODEL_VERSION_TRANSITIONED_TO_STAGING`, `REGISTERED_MODEL_CREATED`, `TRANSITION_REQUEST_CREATED`, `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`, `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`, `TRANSITION_REQUEST_TO_STAGING_CREATED`:
		*f = RegistryWebhookEvent(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMENT_CREATED", "MODEL_VERSION_CREATED", "MODEL_VERSION_TAG_SET", "MODEL_VERSION_TRANSITIONED_STAGE", "MODEL_VERSION_TRANSITIONED_TO_ARCHIVED", "MODEL_VERSION_TRANSITIONED_TO_PRODUCTION", "MODEL_VERSION_TRANSITIONED_TO_STAGING", "REGISTERED_MODEL_CREATED", "TRANSITION_REQUEST_CREATED", "TRANSITION_REQUEST_TO_ARCHIVED_CREATED", "TRANSITION_REQUEST_TO_PRODUCTION_CREATED", "TRANSITION_REQUEST_TO_STAGING_CREATED"`, v)
	}
}

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens.
//
// * `DISABLED`: Webhook is not triggered.
//
// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is not
// triggered on a real event.
type RegistryWebhookStatus string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

// String representation for [fmt.Print]
func (f *RegistryWebhookStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `DISABLED`, `TEST_MODE`:
		*f = RegistryWebhookStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DISABLED", "TEST_MODE"`, v)
	}
}

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
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

type RejectTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity *Activity `tfsdk:"activity" tf:"optional"`
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

type RenameModelResponse struct {
	RegisteredModel *Model `tfsdk:"registered_model" tf:"optional"`
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

type RestoreExperimentResponse struct {
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id" tf:""`
}

type RestoreRunResponse struct {
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

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored types.Int64 `tfsdk:"runs_restored" tf:"optional"`
}

type Run struct {
	// Run data.
	Data *RunData `tfsdk:"data" tf:"optional"`
	// Run metadata.
	Info *RunInfo `tfsdk:"info" tf:"optional"`
	// Run inputs.
	Inputs *RunInputs `tfsdk:"inputs" tf:"optional"`
}

type RunData struct {
	// Run metrics.
	Metrics []Metric `tfsdk:"metrics" tf:"optional"`
	// Run parameters.
	Params []Param `tfsdk:"params" tf:"optional"`
	// Additional metadata key-value pairs.
	Tags []RunTag `tfsdk:"tags" tf:"optional"`
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

// Current status of the run.
type RunInfoStatus string

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *RunInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunInfoStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = RunInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

type RunInputs struct {
	// Run metrics.
	DatasetInputs []DatasetInput `tfsdk:"dataset_inputs" tf:"optional"`
}

type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
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
	OrderBy []types.String `tfsdk:"order_by" tf:"optional"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"view_type" tf:"optional"`
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments []Experiment `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// Qualifier for type of experiments to be returned. If unspecified, return only
// active experiments.
type SearchExperimentsViewType string

const SearchExperimentsViewTypeActiveOnly SearchExperimentsViewType = `ACTIVE_ONLY`

const SearchExperimentsViewTypeAll SearchExperimentsViewType = `ALL`

const SearchExperimentsViewTypeDeletedOnly SearchExperimentsViewType = `DELETED_ONLY`

// String representation for [fmt.Print]
func (f *SearchExperimentsViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SearchExperimentsViewType) Set(v string) error {
	switch v {
	case `ACTIVE_ONLY`, `ALL`, `DELETED_ONLY`:
		*f = SearchExperimentsViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_ONLY", "ALL", "DELETED_ONLY"`, v)
	}
}

// Type always returns SearchExperimentsViewType to satisfy [pflag.Value] interface
func (f *SearchExperimentsViewType) Type() string {
	return "SearchExperimentsViewType"
}

// Searches model versions
type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy []types.String `tfsdk:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions []ModelVersion `tfsdk:"model_versions" tf:"optional"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// Search models
type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy []types.String `tfsdk:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Registered Models that match the search criteria.
	RegisteredModels []Model `tfsdk:"registered_models" tf:"optional"`
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds []types.String `tfsdk:"experiment_ids" tf:"optional"`
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
	OrderBy []types.String `tfsdk:"order_by" tf:"optional"`
	// Token for the current page of runs.
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType types.String `tfsdk:"run_view_type" tf:"optional"`
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Runs that match the search criteria.
	Runs []Run `tfsdk:"runs" tf:"optional"`
}

// Whether to display only active, only deleted, or all runs. Defaults to only
// active runs.
type SearchRunsRunViewType string

const SearchRunsRunViewTypeActiveOnly SearchRunsRunViewType = `ACTIVE_ONLY`

const SearchRunsRunViewTypeAll SearchRunsRunViewType = `ALL`

const SearchRunsRunViewTypeDeletedOnly SearchRunsRunViewType = `DELETED_ONLY`

// String representation for [fmt.Print]
func (f *SearchRunsRunViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SearchRunsRunViewType) Set(v string) error {
	switch v {
	case `ACTIVE_ONLY`, `ALL`, `DELETED_ONLY`:
		*f = SearchRunsRunViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_ONLY", "ALL", "DELETED_ONLY"`, v)
	}
}

// Type always returns SearchRunsRunViewType to satisfy [pflag.Value] interface
func (f *SearchRunsRunViewType) Type() string {
	return "SearchRunsRunViewType"
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

type SetExperimentTagResponse struct {
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

type SetModelTagResponse struct {
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

type SetModelVersionTagResponse struct {
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

type SetTagResponse struct {
}

// Stage of the model version. Valid values are:
//
// * `None`: The initial stage of a model version.
//
// * `Staging`: Staging or pre-production stage.
//
// * `Production`: Production stage.
//
// * `Archived`: Archived stage.
type Stage string

// Archived stage.
const StageArchived Stage = `Archived`

// The initial stage of a model version.
const StageNone Stage = `None`

// Production stage.
const StageProduction Stage = `Production`

// Staging or pre-production stage.
const StageStaging Stage = `Staging`

// String representation for [fmt.Print]
func (f *Stage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Stage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = Stage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns Stage to satisfy [pflag.Value] interface
func (f *Stage) Type() string {
	return "Stage"
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string

// Request to register a new model version has failed.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady Status = `READY`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	Body types.String `tfsdk:"body" tf:"optional"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code" tf:"optional"`
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event types.String `tfsdk:"event" tf:"optional"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`
}

type TestRegistryWebhookResponse struct {
	// Test webhook response object.
	Webhook *TestRegistryWebhook `tfsdk:"webhook" tf:"optional"`
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

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []types.String `tfsdk:"available_actions" tf:"optional"`
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

type TransitionStageResponse struct {
	ModelVersion *ModelVersionDatabricks `tfsdk:"model_version" tf:"optional"`
}

type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id" tf:""`
}

type UpdateCommentResponse struct {
	// Comment details.
	Comment *CommentObject `tfsdk:"comment" tf:"optional"`
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

type UpdateExperimentResponse struct {
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
}

type UpdateModelResponse struct {
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Name of the registered model
	Name types.String `tfsdk:"name" tf:""`
	// Model version number
	Version types.String `tfsdk:"version" tf:""`
}

type UpdateModelVersionResponse struct {
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
	Events []types.String `tfsdk:"events" tf:"optional"`

	HttpUrlSpec *HttpUrlSpec `tfsdk:"http_url_spec" tf:"optional"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`

	JobSpec *JobSpec `tfsdk:"job_spec" tf:"optional"`
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

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `tfsdk:"run_info" tf:"optional"`
}

// Updated status of the run.
type UpdateRunStatus string

const UpdateRunStatusFailed UpdateRunStatus = `FAILED`

const UpdateRunStatusFinished UpdateRunStatus = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatus = `KILLED`

const UpdateRunStatusRunning UpdateRunStatus = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *UpdateRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateRunStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = UpdateRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

type UpdateWebhookResponse struct {
}
