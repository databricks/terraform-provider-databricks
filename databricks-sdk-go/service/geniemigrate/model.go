// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package geniemigrate

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// The request to add workspace files to a migration project.
type AddProjectFilesRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name string `json:"-" url:"-"`
	// The files to add, each identified by `workspace.path` or
	// `workspace.object_id`. Send at most 500 files per request. Duplicate
	// files, and requests that exceed a file-count limit, return
	// `INVALID_PARAMETER_VALUE`. Files already in the project aren't
	// overwritten. An empty list changes nothing.
	Sources []SourceFileRef `json:"sources"`
}

func (s *AddProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lists the files from an `AddProjectFiles` request, including files that were
// already in the project.
type AddProjectFilesResponse struct {
	// The requested files, in request order, without conversion results. Call
	// `ListProjectFiles` or `GetProjectFile` to read their state.
	Files []ProjectFile `json:"files,omitempty"`
}

func (s *AddProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The request message for `CancelOperation` method.
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `json:"-" url:"-"`
}

func (s *CancelOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A conversion of a migration project's files, with its state and the status of
// each file.
type ConversionRun struct {
	// The time the run was accepted, before its background job started.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The time the run ended. Absent while the run is active.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Why the run failed. Set only when `state` is
	// `CONVERSION_RUN_STATE_FAILED`.
	ErrorCode ConversionRunErrorCode `json:"error_code,omitempty"`
	// The files in this run, at most 500, each with its status in this run.
	// Later conversions and reviews can change `ProjectFile.status` but don't
	// change this list.
	Files []ConversionRunFile `json:"files,omitempty"`
	// The resource name of the project being converted, in the format
	// `projects/{project_id}`.
	Project string `json:"project,omitempty"`
	// The skill folders that the run was started with. Empty if the run
	// discovers skills automatically.
	SkillPaths []string `json:"skill_paths,omitempty"`
	// The state of the run.
	State ConversionRunState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConversionRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConversionRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Why a conversion run failed.
type ConversionRunErrorCode string

const ConversionRunErrorCodeConversionRunErrorCodeExpired ConversionRunErrorCode = `CONVERSION_RUN_ERROR_CODE_EXPIRED`

const ConversionRunErrorCodeConversionRunErrorCodeJobFailed ConversionRunErrorCode = `CONVERSION_RUN_ERROR_CODE_JOB_FAILED`

const ConversionRunErrorCodeConversionRunErrorCodeSubmitRejected ConversionRunErrorCode = `CONVERSION_RUN_ERROR_CODE_SUBMIT_REJECTED`

// String representation for [fmt.Print]
func (f *ConversionRunErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConversionRunErrorCode) Set(v string) error {
	switch v {
	case `CONVERSION_RUN_ERROR_CODE_EXPIRED`, `CONVERSION_RUN_ERROR_CODE_JOB_FAILED`, `CONVERSION_RUN_ERROR_CODE_SUBMIT_REJECTED`:
		*f = ConversionRunErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONVERSION_RUN_ERROR_CODE_EXPIRED", "CONVERSION_RUN_ERROR_CODE_JOB_FAILED", "CONVERSION_RUN_ERROR_CODE_SUBMIT_REJECTED"`, v)
	}
}

// Values returns all possible values for ConversionRunErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConversionRunErrorCode) Values() []ConversionRunErrorCode {
	return []ConversionRunErrorCode{
		ConversionRunErrorCodeConversionRunErrorCodeExpired,
		ConversionRunErrorCodeConversionRunErrorCodeJobFailed,
		ConversionRunErrorCodeConversionRunErrorCodeSubmitRejected,
	}
}

// Type always returns ConversionRunErrorCode to satisfy [pflag.Value] interface
func (f *ConversionRunErrorCode) Type() string {
	return "ConversionRunErrorCode"
}

// A file in a conversion run.
type ConversionRunFile struct {
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`.
	Name string `json:"name,omitempty"`
	// The status of the file in this run. `ProjectFile.status` has the file's
	// latest status across all runs.
	Status FileStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConversionRunFile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConversionRunFile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// State of a conversion run.
type ConversionRunState string

const ConversionRunStateConversionRunStateActive ConversionRunState = `CONVERSION_RUN_STATE_ACTIVE`

const ConversionRunStateConversionRunStateCancelled ConversionRunState = `CONVERSION_RUN_STATE_CANCELLED`

const ConversionRunStateConversionRunStateCompleted ConversionRunState = `CONVERSION_RUN_STATE_COMPLETED`

const ConversionRunStateConversionRunStateFailed ConversionRunState = `CONVERSION_RUN_STATE_FAILED`

// String representation for [fmt.Print]
func (f *ConversionRunState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConversionRunState) Set(v string) error {
	switch v {
	case `CONVERSION_RUN_STATE_ACTIVE`, `CONVERSION_RUN_STATE_CANCELLED`, `CONVERSION_RUN_STATE_COMPLETED`, `CONVERSION_RUN_STATE_FAILED`:
		*f = ConversionRunState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONVERSION_RUN_STATE_ACTIVE", "CONVERSION_RUN_STATE_CANCELLED", "CONVERSION_RUN_STATE_COMPLETED", "CONVERSION_RUN_STATE_FAILED"`, v)
	}
}

// Values returns all possible values for ConversionRunState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConversionRunState) Values() []ConversionRunState {
	return []ConversionRunState{
		ConversionRunStateConversionRunStateActive,
		ConversionRunStateConversionRunStateCancelled,
		ConversionRunStateConversionRunStateCompleted,
		ConversionRunStateConversionRunStateFailed,
	}
}

// Type always returns ConversionRunState to satisfy [pflag.Value] interface
func (f *ConversionRunState) Type() string {
	return "ConversionRunState"
}

// The request to convert files of a migration project.
type ConvertProjectFilesRequest struct {
	// The files to convert, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`. All files must belong to `parent`. At most 500
	// files; a workspace can set a lower limit, and requests over it return
	// `INVALID_PARAMETER_VALUE`.
	Names []string `json:"names,omitempty"`
	// The resource name of the project whose files to convert, in the format
	// `projects/{project_id}`. For example, `projects/123`.
	Parent string `json:"-" url:"-"`
	// Custom skills that the conversion can use, as workspace paths of skill
	// folders. Each folder must contain a `SKILL.md` file and be directly
	// inside a skills folder: `<home>/.assistant/skills`,
	// `/Workspace/.assistant/skills`, or another configured skills folder. At
	// most 10 paths, which together must fit in 1000 bytes when encoded as a
	// JSON array. Larger values return `INVALID_PARAMETER_VALUE`. If empty, the
	// conversion discovers skills automatically.
	SkillPaths []string `json:"skill_paths,omitempty"`
	// Whether to convert each selected file from scratch instead of continuing
	// from its previous results (changes, to-do items, blockers, and notes).
	// Supported only for projects whose target framework is
	// `TARGET_FRAMEWORK_SQL`; other projects return `INVALID_PARAMETER_VALUE`.
	// Defaults to `false`.
	StartOver bool `json:"start_over,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConvertProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConvertProjectFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The result of a completed conversion.
type ConvertProjectFilesResponse struct {
	// The completed conversion run. Each file's outcome is in `run.files`; a
	// completed run can include files that failed to convert.
	Run *ConversionRun `json:"run,omitempty"`
}

func (s *ConvertProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateProjectRequest struct {
	// The project to create. The server sets `name` in the response.
	Project Project `json:"project"`
}

func (s *CreateProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Destination asset kind for a converted file.
type DestinationKind string

const DestinationKindDestinationKindDesignerFile DestinationKind = `DESTINATION_KIND_DESIGNER_FILE`

const DestinationKindDestinationKindFile DestinationKind = `DESTINATION_KIND_FILE`

const DestinationKindDestinationKindNotebook DestinationKind = `DESTINATION_KIND_NOTEBOOK`

const DestinationKindDestinationKindQuery DestinationKind = `DESTINATION_KIND_QUERY`

// String representation for [fmt.Print]
func (f *DestinationKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DestinationKind) Set(v string) error {
	switch v {
	case `DESTINATION_KIND_DESIGNER_FILE`, `DESTINATION_KIND_FILE`, `DESTINATION_KIND_NOTEBOOK`, `DESTINATION_KIND_QUERY`:
		*f = DestinationKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DESTINATION_KIND_DESIGNER_FILE", "DESTINATION_KIND_FILE", "DESTINATION_KIND_NOTEBOOK", "DESTINATION_KIND_QUERY"`, v)
	}
}

// Values returns all possible values for DestinationKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *DestinationKind) Values() []DestinationKind {
	return []DestinationKind{
		DestinationKindDestinationKindDesignerFile,
		DestinationKindDestinationKindFile,
		DestinationKindDestinationKindNotebook,
		DestinationKindDestinationKindQuery,
	}
}

// Type always returns DestinationKind to satisfy [pflag.Value] interface
func (f *DestinationKind) Type() string {
	return "DestinationKind"
}

// A Databricks job or pipeline created during conversion.
type DestinationResource struct {
	// The ID of the job that conversion created, if any.
	JobId int64 `json:"job_id,omitempty"`
	// The ID (a UUID) of the pipeline that conversion created, if any.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DestinationResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DestinationResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

// Conversion status of a file in a migration project.
type FileStatus string

const FileStatusFileStatusError FileStatus = `FILE_STATUS_ERROR`

const FileStatusFileStatusInProgress FileStatus = `FILE_STATUS_IN_PROGRESS`

const FileStatusFileStatusPartialSuccess FileStatus = `FILE_STATUS_PARTIAL_SUCCESS`

const FileStatusFileStatusQueued FileStatus = `FILE_STATUS_QUEUED`

const FileStatusFileStatusReviewed FileStatus = `FILE_STATUS_REVIEWED`

const FileStatusFileStatusStopped FileStatus = `FILE_STATUS_STOPPED`

const FileStatusFileStatusSuccess FileStatus = `FILE_STATUS_SUCCESS`

// String representation for [fmt.Print]
func (f *FileStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileStatus) Set(v string) error {
	switch v {
	case `FILE_STATUS_ERROR`, `FILE_STATUS_IN_PROGRESS`, `FILE_STATUS_PARTIAL_SUCCESS`, `FILE_STATUS_QUEUED`, `FILE_STATUS_REVIEWED`, `FILE_STATUS_STOPPED`, `FILE_STATUS_SUCCESS`:
		*f = FileStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_STATUS_ERROR", "FILE_STATUS_IN_PROGRESS", "FILE_STATUS_PARTIAL_SUCCESS", "FILE_STATUS_QUEUED", "FILE_STATUS_REVIEWED", "FILE_STATUS_STOPPED", "FILE_STATUS_SUCCESS"`, v)
	}
}

// Values returns all possible values for FileStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileStatus) Values() []FileStatus {
	return []FileStatus{
		FileStatusFileStatusError,
		FileStatusFileStatusInProgress,
		FileStatusFileStatusPartialSuccess,
		FileStatusFileStatusQueued,
		FileStatusFileStatusReviewed,
		FileStatusFileStatusStopped,
		FileStatusFileStatusSuccess,
	}
}

// Type always returns FileStatus to satisfy [pflag.Value] interface
func (f *FileStatus) Type() string {
	return "FileStatus"
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProjectFileRequest struct {
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Name string `json:"-" url:"-"`
}

func (s *GetProjectFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name string `json:"-" url:"-"`
}

func (s *GetProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListProjectFilesRequest struct {
	// The maximum number of files to return, at most 500. If omitted, zero, or
	// negative, the server chooses. The server can return a different number of
	// files than requested, so keep requesting pages until `next_page_token` is
	// empty.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response. Omit it for the first
	// page. Use it only with the same `parent`.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The response to `ListProjectFiles`.
type ListProjectFilesResponse struct {
	// The token for the next page. Empty when there are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The files in this page.
	ProjectFiles []ProjectFile `json:"project_files,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectFilesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsRequest struct {
	// The folder to list, as `folder = "folders/{folder_id}"`. For example,
	// `folder = "folders/456"`. No other filter is supported. Defaults to the
	// caller's home folder.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of projects to return, at most 500. Defaults to 100 if
	// omitted, zero, or negative. A page can have fewer projects than
	// requested, or none, even when more pages follow.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response. Omit it for the first
	// page. Use the same `filter` for every page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The response to `ListProjects`.
type ListProjectsResponse struct {
	// The token for the next page. Empty when there are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The projects in this page, without `project_config` or `changes_locked`.
	// Call `GetProject` for the full project.
	Projects []Project `json:"projects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The location of a file or folder. Exactly one location type is set. More
// location types can be added later, so clients should ignore types they don't
// recognize.
type Location struct {
	// A file or folder in the workspace.
	Workspace *WorkspaceLocation `json:"workspace,omitempty"`
}

func (s *Location) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A migration project: the conversion settings and source files of one
// migration.
type Project struct {
	// Whether the source dialect and target framework are locked. They lock
	// after a conversion or analysis runs on the project. Other settings stay
	// editable.
	ChangesLocked bool `json:"changes_locked,omitempty"`
	// The time the project was created, if available. Can be absent from
	// `CreateProject` and `UpdateProject` responses.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The display name of the project. Required on create, and can't be blank.
	// Changing it renames the project.
	DisplayName string `json:"display_name"`
	// The workspace folder that contains the project, in the format
	// `folders/{folder_id}`. For example, `folders/456`. If omitted on create,
	// defaults to the caller's home folder. `UpdateProject` can't change it,
	// but moving the project in the workspace does. Returned when available.
	// Converted files are written to `project_config.destination_location`,
	// which can be a different folder.
	Folder string `json:"folder,omitempty"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. The server assigns it when the project is
	// created.
	Name string `json:"name,omitempty"`
	// The conversion settings of the project. Required on create. Returned by
	// `CreateProject`, `GetProject`, and `UpdateProject`, but not by
	// `ListProjects`.
	ProjectConfig ProjectConfig `json:"project_config"`
	// The time the project was last modified, if available. Can be absent from
	// `CreateProject` and `UpdateProject` responses.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Project) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Project) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Conversion settings of a migration project.
type ProjectConfig struct {
	// The workspace folder where converted files are written. Required on
	// create; set `workspace.path` or `workspace.object_id`. `UpdateProject`
	// can't change it.
	DestinationLocation *Location `json:"destination_location,omitempty"`
	// The workspace JSON file that holds the project's lineage report.
	// Optional. To clear it, include `project_config.lineage_location` in
	// `update_mask` and leave this field unset.
	LineageLocation *Location `json:"lineage_location,omitempty"`
	// The dialect of the project's source files. Can't be changed when
	// `Project.changes_locked` is `true`.
	SourceDialect SourceDialect `json:"source_dialect,omitempty"`
	// The framework that source files are converted to. Required on create, and
	// can't be `TARGET_FRAMEWORK_UNSPECIFIED`. Can't be changed when
	// `Project.changes_locked` is `true`.
	TargetFramework TargetFramework `json:"target_framework"`
}

func (s *ProjectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A source file in a migration project, with its latest conversion results. The
// service sets the conversion results. To accept the converted output, call
// `ReviewProjectFiles`; to undo that, call `UnreviewProjectFiles`. Neither
// method changes the conversion results.
type ProjectFile struct {
	// Issues that block the conversion. Returned only by `GetProjectFile`,
	// which reads them from the converted file.
	Blockers []string `json:"blockers,omitempty"`
	// The changes made during conversion. Returned only by `GetProjectFile`,
	// which reads them from the converted file.
	Changes []string `json:"changes,omitempty"`
	// The location of the asset that this file was derived from. Absent if the
	// file wasn't derived from another asset.
	DerivedFrom *Location `json:"derived_from,omitempty"`
	// The location of the converted file. Absent until a conversion writes it.
	Destination *Location `json:"destination,omitempty"`
	// The kind of the converted file, when recorded.
	DestinationKind DestinationKind `json:"destination_kind,omitempty"`
	// The Databricks job or pipeline that conversion created for this file, if
	// any.
	DestinationResource *DestinationResource `json:"destination_resource,omitempty"`
	// The error reported for this file, if any. `GetProjectFile` returns the
	// error stored with the converted file when available, and otherwise the
	// latest error recorded for the project file.
	ErrorMessage string `json:"error_message,omitempty"`
	// Whether the file has conversion notes. Returned by `ListProjectFiles`;
	// call `GetProjectFile` to read the notes.
	HasNotes bool `json:"has_notes,omitempty"`
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`. The server assigns `file_id`, an opaque value
	// of at most 63 characters.
	Name string `json:"name,omitempty"`
	// Informational notes about the conversion that need no decision. Returned
	// only by `GetProjectFile`, which reads them from the converted file.
	Notes []string `json:"notes,omitempty"`
	// The location of the source file.
	Source *Location `json:"source,omitempty"`
	// The file's latest conversion status across all runs. For its status in a
	// specific run, read `ConversionRun.files` in that run's operation
	// metadata.
	Status FileStatus `json:"status,omitempty"`
	// Follow-up work (to-do items) identified during conversion. Returned only
	// by `GetProjectFile`, which reads them from the converted file.
	Todos []string `json:"todos,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectFile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectFile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The request to remove files from a migration project.
type RemoveProjectFilesRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name string `json:"-" url:"-"`
	// The files to remove, by `ProjectFile.name`. For example,
	// `projects/123/files/789`. Every file must belong to the project in
	// `name`. An empty list changes nothing. Send at most 10 files per request,
	// the default limit. Requests over a project's limit return
	// `INVALID_PARAMETER_VALUE`.
	Names []string `json:"names,omitempty"`
}

func (s *RemoveProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lists the files from a `RemoveProjectFiles` request.
type RemoveProjectFilesResponse struct {
	// The requested files, in request order. Each entry identifies a source
	// file and has no conversion results.
	Files []ProjectFile `json:"files,omitempty"`
}

func (s *RemoveProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The request to mark files of a migration project as reviewed.
type ReviewProjectFilesRequest struct {
	// The files to review: 1 to 100 distinct file resource names, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Names []string `json:"names"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. Every file in `names` must belong to this
	// project.
	Parent string `json:"-" url:"-"`
}

func (s *ReviewProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The files reviewed by `ReviewProjectFiles`, in request order.
type ReviewProjectFilesResponse struct {
	// The reviewed files, in request order. To read a file's changes, to-do
	// items, blockers, and notes, call `GetProjectFile`.
	ProjectFiles []ProjectFile `json:"project_files,omitempty"`
}

func (s *ReviewProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Source dialect of the files in a migration project.
type SourceDialect string

const SourceDialectSourceDialectAlteryx SourceDialect = `SOURCE_DIALECT_ALTERYX`

const SourceDialectSourceDialectBigquery SourceDialect = `SOURCE_DIALECT_BIGQUERY`

const SourceDialectSourceDialectInformatica SourceDialect = `SOURCE_DIALECT_INFORMATICA`

const SourceDialectSourceDialectOracle SourceDialect = `SOURCE_DIALECT_ORACLE`

const SourceDialectSourceDialectRedshift SourceDialect = `SOURCE_DIALECT_REDSHIFT`

const SourceDialectSourceDialectSnowflake SourceDialect = `SOURCE_DIALECT_SNOWFLAKE`

const SourceDialectSourceDialectSsis SourceDialect = `SOURCE_DIALECT_SSIS`

const SourceDialectSourceDialectTeradata SourceDialect = `SOURCE_DIALECT_TERADATA`

const SourceDialectSourceDialectTSql SourceDialect = `SOURCE_DIALECT_T_SQL`

// String representation for [fmt.Print]
func (f *SourceDialect) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SourceDialect) Set(v string) error {
	switch v {
	case `SOURCE_DIALECT_ALTERYX`, `SOURCE_DIALECT_BIGQUERY`, `SOURCE_DIALECT_INFORMATICA`, `SOURCE_DIALECT_ORACLE`, `SOURCE_DIALECT_REDSHIFT`, `SOURCE_DIALECT_SNOWFLAKE`, `SOURCE_DIALECT_SSIS`, `SOURCE_DIALECT_TERADATA`, `SOURCE_DIALECT_T_SQL`:
		*f = SourceDialect(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SOURCE_DIALECT_ALTERYX", "SOURCE_DIALECT_BIGQUERY", "SOURCE_DIALECT_INFORMATICA", "SOURCE_DIALECT_ORACLE", "SOURCE_DIALECT_REDSHIFT", "SOURCE_DIALECT_SNOWFLAKE", "SOURCE_DIALECT_SSIS", "SOURCE_DIALECT_TERADATA", "SOURCE_DIALECT_T_SQL"`, v)
	}
}

// Values returns all possible values for SourceDialect.
//
// There is no guarantee on the order of the values in the slice.
func (f *SourceDialect) Values() []SourceDialect {
	return []SourceDialect{
		SourceDialectSourceDialectAlteryx,
		SourceDialectSourceDialectBigquery,
		SourceDialectSourceDialectInformatica,
		SourceDialectSourceDialectOracle,
		SourceDialectSourceDialectRedshift,
		SourceDialectSourceDialectSnowflake,
		SourceDialectSourceDialectSsis,
		SourceDialectSourceDialectTeradata,
		SourceDialectSourceDialectTSql,
	}
}

// Type always returns SourceDialect to satisfy [pflag.Value] interface
func (f *SourceDialect) Type() string {
	return "SourceDialect"
}

// A reference to a source file in the workspace. Set `workspace.path` or
// `workspace.object_id`. For example, `{"workspace": {"path":
// "/Workspace/Shared/migration/source.sql"}}`.
type SourceFileRef struct {
	// The location of the asset that this file was derived from, if any. Used
	// only by `AddProjectFiles`, and can't be changed after the file is added.
	DerivedFrom *Location `json:"derived_from,omitempty"`
	// The source file in the workspace. Must identify an existing file the
	// caller can read.
	Workspace *WorkspaceLocation `json:"workspace,omitempty"`
}

func (s *SourceFileRef) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Output framework of the converted files in a migration project.
type TargetFramework string

const TargetFrameworkTargetFrameworkDesigner TargetFramework = `TARGET_FRAMEWORK_DESIGNER`

const TargetFrameworkTargetFrameworkPyspark TargetFramework = `TARGET_FRAMEWORK_PYSPARK`

const TargetFrameworkTargetFrameworkSparkDeclarativePipelines TargetFramework = `TARGET_FRAMEWORK_SPARK_DECLARATIVE_PIPELINES`

const TargetFrameworkTargetFrameworkSql TargetFramework = `TARGET_FRAMEWORK_SQL`

// String representation for [fmt.Print]
func (f *TargetFramework) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TargetFramework) Set(v string) error {
	switch v {
	case `TARGET_FRAMEWORK_DESIGNER`, `TARGET_FRAMEWORK_PYSPARK`, `TARGET_FRAMEWORK_SPARK_DECLARATIVE_PIPELINES`, `TARGET_FRAMEWORK_SQL`:
		*f = TargetFramework(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TARGET_FRAMEWORK_DESIGNER", "TARGET_FRAMEWORK_PYSPARK", "TARGET_FRAMEWORK_SPARK_DECLARATIVE_PIPELINES", "TARGET_FRAMEWORK_SQL"`, v)
	}
}

// Values returns all possible values for TargetFramework.
//
// There is no guarantee on the order of the values in the slice.
func (f *TargetFramework) Values() []TargetFramework {
	return []TargetFramework{
		TargetFrameworkTargetFrameworkDesigner,
		TargetFrameworkTargetFrameworkPyspark,
		TargetFrameworkTargetFrameworkSparkDeclarativePipelines,
		TargetFrameworkTargetFrameworkSql,
	}
}

// Type always returns TargetFramework to satisfy [pflag.Value] interface
func (f *TargetFramework) Type() string {
	return "TargetFramework"
}

// The request to remove the reviewed status from files of a migration project.
type UnreviewProjectFilesRequest struct {
	// The files to unreview: 1 to 100 distinct file resource names, in the
	// format `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Names []string `json:"names"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. Every file in `names` must belong to this
	// project.
	Parent string `json:"-" url:"-"`
}

func (s *UnreviewProjectFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The files unreviewed by `UnreviewProjectFiles`, in request order.
type UnreviewProjectFilesResponse struct {
	// The unreviewed files, in request order. To read a file's changes, to-do
	// items, blockers, and notes, call `GetProjectFile`.
	ProjectFiles []ProjectFile `json:"project_files,omitempty"`
}

func (s *UnreviewProjectFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. The server assigns it when the project is
	// created.
	Name string `json:"-" url:"-"`
	// The project to update, identified by `name`.
	Project Project `json:"project"`
	// The fields to update. Supported fields are `display_name`,
	// `project_config.source_dialect`, `project_config.target_framework`, and
	// `project_config.lineage_location`. `project_config` and `*` update all
	// three `project_config` fields and can't be combined with `display_name`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A file or folder in the workspace. On input, set `path` or `object_id`; if
// both are set, `object_id` is used. Responses set both when known.
type WorkspaceLocation struct {
	// The workspace object ID, as returned in `object_id` by the Workspace API.
	// For example, `12345`.
	ObjectId int64 `json:"object_id,omitempty"`
	// The absolute workspace path. For example,
	// `/Workspace/Shared/migration/source.sql`.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
