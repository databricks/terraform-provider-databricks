// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

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
	ActivityType ActivityType `tfsdk:"activity_type"`
	// User-provided comment associated with the activity.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
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
	FromStage Stage `tfsdk:"from_stage"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment types.String `tfsdk:"system_comment"`
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
	ToStage Stage `tfsdk:"to_stage"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
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
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

type ApproveTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity *Activity `tfsdk:"activity"`
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
	AvailableActions types.List `tfsdk:"available_actions"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Comment ID
	Id types.String `tfsdk:"id"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

type CreateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

type CreateCommentResponse struct {
	// Comment details.
	Comment *CommentObject `tfsdk:"comment"`
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation types.String `tfsdk:"artifact_location"`
	// Experiment name.
	Name types.String `tfsdk:"name"`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags types.List `tfsdk:"tags"`
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description"`
	// Register models under this name
	Name types.String `tfsdk:"name"`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags"`
}

type CreateModelResponse struct {
	RegisteredModel *Model `tfsdk:"registered_model"`
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	Description types.String `tfsdk:"description"`
	// Register model under this name
	Name types.String `tfsdk:"name"`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId types.String `tfsdk:"run_id"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink types.String `tfsdk:"run_link"`
	// URI indicating the location of the model artifacts.
	Source types.String `tfsdk:"source"`
	// Additional metadata for model version.
	Tags types.List `tfsdk:"tags"`
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion *ModelVersion `tfsdk:"model_version"`
}

type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
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
	Events types.List `tfsdk:"events"`

	HttpUrlSpec *HttpUrlSpec `tfsdk:"http_url_spec"`

	JobSpec *JobSpec `tfsdk:"job_spec"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `tfsdk:"status"`
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime types.Int64 `tfsdk:"start_time"`
	// Additional metadata for run.
	Tags types.List `tfsdk:"tags"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId types.String `tfsdk:"user_id"`
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `tfsdk:"run"`
}

type CreateTransitionRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

type CreateTransitionRequestResponse struct {
	// Transition request details.
	Request *TransitionRequest `tfsdk:"request"`
}

type CreateWebhookResponse struct {
	Webhook *RegistryWebhook `tfsdk:"webhook"`
}

type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest types.String `tfsdk:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name types.String `tfsdk:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile types.String `tfsdk:"profile"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema types.String `tfsdk:"schema"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	Source types.String `tfsdk:"source"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	SourceType types.String `tfsdk:"source_type"`
}

type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset *Dataset `tfsdk:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags types.List `tfsdk:"tags"`
}

// Delete a comment
type DeleteCommentRequest struct {
	Id types.String `tfsdk:"-" url:"id"`
}

type DeleteCommentResponse struct {
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
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
	RunId types.String `tfsdk:"run_id"`
}

type DeleteRunResponse struct {
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis types.Int64 `tfsdk:"max_timestamp_millis"`
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted types.Int64 `tfsdk:"runs_deleted"`
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key types.String `tfsdk:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id"`
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
	Stage DeleteTransitionRequestStage `tfsdk:"-" url:"stage"`
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
	ArtifactLocation types.String `tfsdk:"artifact_location"`
	// Creation time
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Last update time
	LastUpdateTime types.Int64 `tfsdk:"last_update_time"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage types.String `tfsdk:"lifecycle_stage"`
	// Human readable name that identifies the experiment.
	Name types.String `tfsdk:"name"`
	// Tags: Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags"`
}

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type ExperimentAccessControlResponse struct {
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

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `tfsdk:"permission_level"`
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
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `tfsdk:"permission_level"`
}

type ExperimentPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-" url:"-"`
}

type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

type FileInfo struct {
	// Size in bytes. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir"`
	// Path relative to the root artifact directory run.
	Path types.String `tfsdk:"path"`
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
	PermissionLevels types.List `tfsdk:"permission_levels"`
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
	Experiment *Experiment `tfsdk:"experiment"`
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
	Name types.String `tfsdk:"name"`
	// List of stages.
	Stages types.List `tfsdk:"stages"`
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions"`
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics types.List `tfsdk:"metrics"`
	// Token that can be used to retrieve the next page of metric history
	// results
	NextPageToken types.String `tfsdk:"next_page_token"`
}

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-" url:"name"`
}

type GetModelResponse struct {
	RegisteredModelDatabricks *ModelDatabricks `tfsdk:"registered_model_databricks"`
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
	ArtifactUri types.String `tfsdk:"artifact_uri"`
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-" url:"name"`
	// Model version number
	Version types.String `tfsdk:"-" url:"version"`
}

type GetModelVersionResponse struct {
	ModelVersion *ModelVersion `tfsdk:"model_version"`
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-" url:"-"`
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
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
	Run *Run `tfsdk:"run"`
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization types.String `tfsdk:"authorization"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret types.String `tfsdk:"secret"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url"`
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url"`
}

type InputTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken types.String `tfsdk:"access_token"`
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
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
	Files types.List `tfsdk:"files"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri"`
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
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	RegisteredModels types.List `tfsdk:"registered_models"`
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks"`
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
	Requests types.List `tfsdk:"requests"`
}

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events types.List `tfsdk:"-" url:"events,omitempty"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	ModelName types.String `tfsdk:"-" url:"model_name,omitempty"`
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics types.List `tfsdk:"metrics"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params types.List `tfsdk:"params"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags types.List `tfsdk:"tags"`
}

type LogBatchResponse struct {
}

type LogInputs struct {
	// Dataset inputs
	Datasets types.List `tfsdk:"datasets"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

type LogInputsResponse struct {
}

type LogMetric struct {
	// Name of the metric.
	Key types.String `tfsdk:"key"`
	// ID of the run under which to log the metric. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Step at which to log the metric
	Step types.Int64 `tfsdk:"step"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp types.Int64 `tfsdk:"timestamp"`
	// Double value of the metric being logged.
	Value types.Float64 `tfsdk:"value"`
}

type LogMetricResponse struct {
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

type LogModelResponse struct {
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key types.String `tfsdk:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value types.String `tfsdk:"value"`
}

type LogParamResponse struct {
}

type Metric struct {
	// Key identifying this metric.
	Key types.String `tfsdk:"key"`
	// Step at which to log the metric.
	Step types.Int64 `tfsdk:"step"`
	// The timestamp at which this metric was recorded.
	Timestamp types.Int64 `tfsdk:"timestamp"`
	// Value associated with this metric.
	Value types.Float64 `tfsdk:"value"`
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Description of this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions types.List `tfsdk:"latest_versions"`
	// Unique name for the model.
	Name types.String `tfsdk:"name"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags types.List `tfsdk:"tags"`
	// User that created this `registered_model`
	UserId types.String `tfsdk:"user_id"`
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions types.List `tfsdk:"latest_versions"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`
	// Array of tags associated with the model.
	Tags types.List `tfsdk:"tags"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

type ModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Current stage for this `model_version`.
	CurrentStage types.String `tfsdk:"current_stage"`
	// Description of this `model_version`.
	Description types.String `tfsdk:"description"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Unique name of the model
	Name types.String `tfsdk:"name"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId types.String `tfsdk:"run_id"`
	// Run Link: Direct link to the run that generated this version
	RunLink types.String `tfsdk:"run_link"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source types.String `tfsdk:"source"`
	// Current status of `model_version`
	Status ModelVersionStatus `tfsdk:"status"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage types.String `tfsdk:"status_message"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags types.List `tfsdk:"tags"`
	// User that created this `model_version`.
	UserId types.String `tfsdk:"user_id"`
	// Model's version number.
	Version types.String `tfsdk:"version"`
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	CurrentStage Stage `tfsdk:"current_stage"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId types.String `tfsdk:"run_id"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink types.String `tfsdk:"run_link"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source types.String `tfsdk:"source"`
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	Status Status `tfsdk:"status"`
	// Details on the current status, for example why registration failed.
	StatusMessage types.String `tfsdk:"status_message"`
	// Array of tags that are associated with the model version.
	Tags types.List `tfsdk:"tags"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
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
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

type Param struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key"`
	// Value associated with this param.
	Value types.String `tfsdk:"value"`
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
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type RegisteredModelAccessControlResponse struct {
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

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `tfsdk:"permission_level"`
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
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `tfsdk:"permission_level"`
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-" url:"-"`
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
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
	Events types.List `tfsdk:"events"`

	HttpUrlSpec *HttpUrlSpecWithoutSecret `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec *JobSpecWithoutSecret `tfsdk:"job_spec"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `tfsdk:"status"`
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
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

type RejectTransitionRequestResponse struct {
	// Activity recorded for the action.
	Activity *Activity `tfsdk:"activity"`
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name"`
}

type RenameModelResponse struct {
	RegisteredModel *Model `tfsdk:"registered_model"`
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

type RestoreExperimentResponse struct {
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id"`
}

type RestoreRunResponse struct {
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis types.Int64 `tfsdk:"min_timestamp_millis"`
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored types.Int64 `tfsdk:"runs_restored"`
}

type Run struct {
	// Run data.
	Data *RunData `tfsdk:"data"`
	// Run metadata.
	Info *RunInfo `tfsdk:"info"`
	// Run inputs.
	Inputs *RunInputs `tfsdk:"inputs"`
}

type RunData struct {
	// Run metrics.
	Metrics types.List `tfsdk:"metrics"`
	// Run parameters.
	Params types.List `tfsdk:"params"`
	// Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags"`
}

type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like `s3://bucket/directory` or `dbfs:/my/directory`. If not set, the
	// local `./mlruns` directory is chosen.
	ArtifactUri types.String `tfsdk:"artifact_uri"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage types.String `tfsdk:"lifecycle_stage"`
	// Unique identifier for the run.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime types.Int64 `tfsdk:"start_time"`
	// Current status of the run.
	Status RunInfoStatus `tfsdk:"status"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId types.String `tfsdk:"user_id"`
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
	DatasetInputs types.List `tfsdk:"dataset_inputs"`
}

type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	Filter types.String `tfsdk:"filter"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults types.Int64 `tfsdk:"max_results"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy types.List `tfsdk:"order_by"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"page_token"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType SearchExperimentsViewType `tfsdk:"view_type"`
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
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
	OrderBy types.List `tfsdk:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token"`
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
	OrderBy types.List `tfsdk:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models"`
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds types.List `tfsdk:"experiment_ids"`
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
	Filter types.String `tfsdk:"filter"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults types.Int64 `tfsdk:"max_results"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional "DESC" or "ASC" annotation, where "ASC" is the
	// default. Example: ["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"] Tiebreaks are done by start_time DESC followed by run_id
	// for runs with the same start time (and this is the default ordering
	// criterion if order_by is not provided).
	OrderBy types.List `tfsdk:"order_by"`
	// Token for the current page of runs.
	PageToken types.String `tfsdk:"page_token"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType SearchRunsRunViewType `tfsdk:"run_view_type"`
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs"`
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
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
}

type SetExperimentTagResponse struct {
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// Unique name of the model.
	Name types.String `tfsdk:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
}

type SetModelTagResponse struct {
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// Unique name of the model.
	Name types.String `tfsdk:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
	// Model version number.
	Version types.String `tfsdk:"version"`
}

type SetModelVersionTagResponse struct {
}

type SetTag struct {
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
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
	Body types.String `tfsdk:"body"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code"`
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `tfsdk:"event"`
	// Webhook ID
	Id types.String `tfsdk:"id"`
}

type TestRegistryWebhookResponse struct {
	// Test webhook response object.
	Webhook *TestRegistryWebhook `tfsdk:"webhook"`
}

type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions"`
	// User-provided comment associated with the transition request.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
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
	ToStage Stage `tfsdk:"to_stage"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

type TransitionStageResponse struct {
	ModelVersion *ModelVersionDatabricks `tfsdk:"model_version"`
}

type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id"`
}

type UpdateCommentResponse struct {
	// Comment details.
	Comment *CommentObject `tfsdk:"comment"`
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name"`
}

type UpdateExperimentResponse struct {
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
}

type UpdateModelResponse struct {
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Name of the registered model
	Name types.String `tfsdk:"name"`
	// Model version number
	Version types.String `tfsdk:"version"`
}

type UpdateModelVersionResponse struct {
}

type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
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
	Events types.List `tfsdk:"events"`

	HttpUrlSpec *HttpUrlSpec `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec *JobSpec `tfsdk:"job_spec"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `tfsdk:"status"`
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime types.Int64 `tfsdk:"end_time"`
	// ID of the run to update. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use run_id instead] ID of the run to update.. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Updated status of the run.
	Status UpdateRunStatus `tfsdk:"status"`
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `tfsdk:"run_info"`
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
