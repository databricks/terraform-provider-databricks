// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// If changed, also update
// estore/namespaces/defaultbaseenvironments/latest.proto
type BaseEnvironmentType string

const BaseEnvironmentTypeCpu BaseEnvironmentType = `CPU`

const BaseEnvironmentTypeGpu BaseEnvironmentType = `GPU`

// String representation for [fmt.Print]
func (f *BaseEnvironmentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BaseEnvironmentType) Set(v string) error {
	switch v {
	case `CPU`, `GPU`:
		*f = BaseEnvironmentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "GPU"`, v)
	}
}

// Values returns all possible values for BaseEnvironmentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *BaseEnvironmentType) Values() []BaseEnvironmentType {
	return []BaseEnvironmentType{
		BaseEnvironmentTypeCpu,
		BaseEnvironmentTypeGpu,
	}
}

// Type always returns BaseEnvironmentType to satisfy [pflag.Value] interface
func (f *BaseEnvironmentType) Type() string {
	return "BaseEnvironmentType"
}

type CreateWorkspaceBaseEnvironmentRequest struct {
	// A unique identifier for this request. A random UUID is recommended. This
	// request is only idempotent if a request_id is provided.
	RequestId string `json:"-" url:"request_id,omitempty"`
	// Required. The workspace base environment to create.
	WorkspaceBaseEnvironment WorkspaceBaseEnvironment `json:"workspace_base_environment"`
	// The ID to use for the workspace base environment, which will become the
	// final component of the resource name. This value should be 4-63
	// characters, and valid characters are /[a-z][0-9]-/.
	WorkspaceBaseEnvironmentId string `json:"-" url:"workspace_base_environment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWorkspaceBaseEnvironmentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// A singleton resource representing the default workspace base environment
// configuration. This resource contains the workspace base environments that
// are used as defaults for serverless notebooks and jobs in the workspace, for
// both CPU and GPU compute types. Each default can reference a
// Databricks-provided base environment or a base environment created in the
// workspace.
//
// When cpu_workspace_base_environment or gpu_workspace_base_environment is
// unset, new serverless notebooks for that compute type use the Databricks
// Standard base environment: the latest stable standard environment version
// provided by Databricks, with no workspace base environment applied. The
// workspace settings page shows this as Standard Latest.
type DefaultWorkspaceBaseEnvironment struct {
	// The default workspace base environment for CPU compute. If unset, CPU
	// compute defaults to the Databricks Standard base environment (Standard
	// Latest). Format: workspace-base-environments/{workspace_base_environment}
	CpuWorkspaceBaseEnvironment string `json:"cpu_workspace_base_environment,omitempty"`
	// The default workspace base environment for GPU compute. If unset, GPU
	// compute defaults to the Databricks Standard base environment (Standard
	// Latest). Format: workspace-base-environments/{workspace_base_environment}
	GpuWorkspaceBaseEnvironment string `json:"gpu_workspace_base_environment,omitempty"`
	// The resource name of this singleton resource. Format:
	// default-workspace-base-environment
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultWorkspaceBaseEnvironment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultWorkspaceBaseEnvironment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteWorkspaceBaseEnvironmentRequest struct {
	// Required. The resource name of the workspace base environment to delete.
	// Format: workspace-base-environments/{workspace_base_environment}
	Name string `json:"-" url:"-"`
}

func (s *DeleteWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Environment specification for a WorkspaceBaseEnvironment. Contains the
// environment version and dependencies configuration.
type EnvironmentSpec struct {
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a valid pip requirements file line per
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/.
	// Allowed dependencies include a requirement specifier, an archive URL, a
	// local project path (such as WSFS or UC Volumes in Databricks), or a VCS
	// project URL.
	Dependencies []string `json:"dependencies,omitempty"`
	// Environment version used by the environment. Each version comes with a
	// specific Python version and a set of Python packages. The version is a
	// string, consisting of an integer.
	EnvironmentVersion string `json:"environment_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnvironmentSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvironmentSpec) MarshalJSON() ([]byte, error) {
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

type GetDefaultWorkspaceBaseEnvironmentRequest struct {
	// A static resource name of the default workspace base environment. Format:
	// default-workspace-base-environment
	Name string `json:"-" url:"-"`
}

func (s *GetDefaultWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceBaseEnvironmentRequest struct {
	// Required. The resource name of the workspace base environment to
	// retrieve. Format:
	// workspace-base-environments/{workspace_base_environment}
	Name string `json:"-" url:"-"`
}

func (s *GetWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListWorkspaceBaseEnvironmentsRequest struct {
	// The maximum number of environments to return per page. Default is 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token for pagination. Received from a previous
	// ListWorkspaceBaseEnvironments call.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceBaseEnvironmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceBaseEnvironmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for ListWorkspaceBaseEnvironments.
type ListWorkspaceBaseEnvironmentsResponse struct {
	// Token to retrieve the next page of results. Empty if there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of workspace base environments.
	WorkspaceBaseEnvironments []WorkspaceBaseEnvironment `json:"workspace_base_environments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceBaseEnvironmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceBaseEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Request message for RefreshWorkspaceBaseEnvironments.
type RefreshWorkspaceBaseEnvironmentRequest struct {
	// Required. The resource name of the workspace base environment to refresh.
	// Format: workspace-base-environments/{workspace_base_environment}
	Name string `json:"-" url:"-"`
}

func (s *RefreshWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDefaultWorkspaceBaseEnvironmentRequest struct {
	// Required. The default workspace base environment configuration to update.
	DefaultWorkspaceBaseEnvironment DefaultWorkspaceBaseEnvironment `json:"default_workspace_base_environment"`
	// The resource name of this singleton resource. Format:
	// default-workspace-base-environment
	Name string `json:"-" url:"-"`
	// Field mask specifying which fields to update. Use comma as the separator
	// for multiple fields (no space). The special value '*' indicates that all
	// fields should be updated (full replacement). Valid field paths:
	// cpu_workspace_base_environment, gpu_workspace_base_environment
	//
	// To unset one or both defaults, include the field path(s) in the mask and
	// omit them from the request body. To unset both, you must list both paths
	// explicitly — the wildcard '*' cannot be used to unset fields. Unsetting
	// a default makes that compute type use the Databricks Standard base
	// environment. The Standard base environment has no resource name, so
	// unsetting the field is the only way to select it.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateDefaultWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceBaseEnvironmentRequest struct {
	// The resource name of the workspace base environment to update. Format:
	// workspace-base-environments/{workspace_base_environment}
	Name string `json:"-" url:"-"`
	// The workspace base environment with its new configuration. The `name`
	// path parameter identifies the environment to update. Set `display_name`
	// and `filepath` on every update, even when they do not change. An update
	// keeps the stored `base_environment_type`, but the YAML file is validated
	// against the type in the request, which defaults to `CPU`, so also set it
	// for a GPU environment.
	WorkspaceBaseEnvironment WorkspaceBaseEnvironment `json:"workspace_base_environment"`
}

func (s *UpdateWorkspaceBaseEnvironmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A WorkspaceBaseEnvironment defines a workspace-level environment
// configuration consisting of an environment version and a list of
// dependencies.
type WorkspaceBaseEnvironment struct {
	// The type of base environment (CPU or GPU). The default is `CPU`.
	BaseEnvironmentType BaseEnvironmentType `json:"base_environment_type,omitempty"`
	// Timestamp when the environment was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// User ID of the creator.
	CreatorUserId string `json:"creator_user_id,omitempty"`
	// Human-readable display name for the workspace base environment.
	DisplayName string `json:"display_name"`
	// The workspace file or Unity Catalog volume path to the YAML file that
	// defines the environment version and dependencies. The file name must end
	// in `.yaml` or `.yml`. Create and update requests must set this field.
	Filepath string `json:"filepath,omitempty"`
	// Whether this is the default environment for the workspace.
	IsDefault bool `json:"is_default,omitempty"`
	// User ID of the last user who updated the environment.
	LastUpdatedUserId string `json:"last_updated_user_id,omitempty"`
	// Status message providing additional details about the environment status.
	// When the status is `FAILED`, the message contains the build error.
	Message string `json:"message,omitempty"`
	// The resource name of the workspace base environment. Format:
	// workspace-base-environments/{workspace-base-environment} The server
	// ignores this field on create and builds the name from
	// `workspace_base_environment_id`, or from `display_name` plus a random
	// suffix when no ID is given.
	Name string `json:"name,omitempty"`
	// The environment specification containing version and dependencies, read
	// from the YAML file at `filepath`.
	Spec *EnvironmentSpec `json:"spec,omitempty"`
	// The status of the materialized workspace base environment, which the
	// server builds from the YAML file. The status is `CREATED` once the build
	// has succeeded.
	Status WorkspaceBaseEnvironmentCacheStatus `json:"status,omitempty"`
	// Timestamp when the environment was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceBaseEnvironment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceBaseEnvironment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of the environment materialization.
type WorkspaceBaseEnvironmentCacheStatus string

const WorkspaceBaseEnvironmentCacheStatusCreated WorkspaceBaseEnvironmentCacheStatus = `CREATED`

const WorkspaceBaseEnvironmentCacheStatusExpired WorkspaceBaseEnvironmentCacheStatus = `EXPIRED`

const WorkspaceBaseEnvironmentCacheStatusFailed WorkspaceBaseEnvironmentCacheStatus = `FAILED`

const WorkspaceBaseEnvironmentCacheStatusInvalid WorkspaceBaseEnvironmentCacheStatus = `INVALID`

const WorkspaceBaseEnvironmentCacheStatusPending WorkspaceBaseEnvironmentCacheStatus = `PENDING`

const WorkspaceBaseEnvironmentCacheStatusRefreshing WorkspaceBaseEnvironmentCacheStatus = `REFRESHING`

// String representation for [fmt.Print]
func (f *WorkspaceBaseEnvironmentCacheStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceBaseEnvironmentCacheStatus) Set(v string) error {
	switch v {
	case `CREATED`, `EXPIRED`, `FAILED`, `INVALID`, `PENDING`, `REFRESHING`:
		*f = WorkspaceBaseEnvironmentCacheStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATED", "EXPIRED", "FAILED", "INVALID", "PENDING", "REFRESHING"`, v)
	}
}

// Values returns all possible values for WorkspaceBaseEnvironmentCacheStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceBaseEnvironmentCacheStatus) Values() []WorkspaceBaseEnvironmentCacheStatus {
	return []WorkspaceBaseEnvironmentCacheStatus{
		WorkspaceBaseEnvironmentCacheStatusCreated,
		WorkspaceBaseEnvironmentCacheStatusExpired,
		WorkspaceBaseEnvironmentCacheStatusFailed,
		WorkspaceBaseEnvironmentCacheStatusInvalid,
		WorkspaceBaseEnvironmentCacheStatusPending,
		WorkspaceBaseEnvironmentCacheStatusRefreshing,
	}
}

// Type always returns WorkspaceBaseEnvironmentCacheStatus to satisfy [pflag.Value] interface
func (f *WorkspaceBaseEnvironmentCacheStatus) Type() string {
	return "WorkspaceBaseEnvironmentCacheStatus"
}

// Metadata for the WorkspaceBaseEnvironment long-running operations. This
// message tracks the progress of the workspace base environment long-running
// process.
type WorkspaceBaseEnvironmentOperationMetadata struct {
}

func (s *WorkspaceBaseEnvironmentOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
