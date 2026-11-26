// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package bundle

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// A request to complete a Version.
type CompleteVersionRequest struct {
	// The reason for completing the version. Must be a terminal reason:
	// VERSION_COMPLETE_SUCCESS, VERSION_COMPLETE_FAILURE, or
	// VERSION_COMPLETE_FORCE_ABORT.
	CompletionReason VersionComplete `json:"completion_reason"`
	// If true, force-completes the version even if the caller is not the
	// original creator. The completion_reason must be
	// VERSION_COMPLETE_FORCE_ABORT when force is true.
	Force bool `json:"force,omitempty"`
	// The name of the version to complete. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CompleteVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CompleteVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDeploymentRequest struct {
	// The deployment to create.
	Deployment Deployment `json:"deployment"`
	// The ID to use for the deployment, which will become the final component
	// of the deployment's resource name (i.e. `deployments/{deployment_id}`).
	DeploymentId string `json:"-" url:"deployment_id"`
}

type CreateOperationRequest struct {
	// The resource operation to create.
	Operation Operation `json:"operation"`
	// The parent version where this operation will be recorded. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Parent string `json:"-" url:"-"`
	// The key identifying the resource this operation applies to. Becomes the
	// final component of the operation's name.
	ResourceKey string `json:"-" url:"resource_key"`
}

type CreateVersionRequest struct {
	// The parent deployment where this version will be created. Format:
	// deployments/{deployment_id}
	Parent string `json:"-" url:"-"`
	// The version to create.
	Version Version `json:"version"`
	// The version ID the caller expects to create. The server validates this
	// equals `last_version_id + 1` on the deployment. If it doesn't match, the
	// server returns `ABORTED`.
	VersionId string `json:"-" url:"version_id"`
}

type DeleteDeploymentRequest struct {
	// Resource name of the deployment to delete. Format:
	// deployments/{deployment_id}
	Name string `json:"-" url:"-"`
}

// A bundle deployment registered with the control plane.
type Deployment struct {
	// When the deployment was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The user who created the deployment (email or principal name).
	CreatedBy string `json:"created_by,omitempty"`
	// Bundle target deployment mode (development or production), derived from
	// the most recent version's mode.
	DeploymentMode DeploymentMode `json:"deployment_mode,omitempty"`
	// When the deployment was destroyed (i.e. `bundle destroy` completed).
	// Unset if the deployment has not been destroyed. Named destroy_time (not
	// delete_time) because this tracks the `databricks bundle destroy` command,
	// not the API-level deletion.
	DestroyTime *time.Time `json:"destroy_time,omitempty"`
	// The user who destroyed the deployment (email or principal name). Unset if
	// the deployment has not been destroyed.
	DestroyedBy string `json:"destroyed_by,omitempty"`
	// Human-readable name for the deployment.
	DisplayName string `json:"display_name,omitempty"`
	// The version_id of the most recent deployment version.
	LastVersionId string `json:"last_version_id,omitempty"`
	// Resource name of the deployment. Format: deployments/{deployment_id}
	Name string `json:"name,omitempty"`
	// Current status of the deployment.
	Status DeploymentStatus `json:"status,omitempty"`
	// The bundle target name associated with this deployment.
	TargetName string `json:"target_name,omitempty"`
	// When the deployment was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Deployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Deployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Bundle target deployment mode. Mirrors the `mode` field on a bundle target in
// `databricks.yml` (see
// https://docs.databricks.com/dev-tools/bundles/deployment-modes).
type DeploymentMode string

const DeploymentModeDeploymentModeDevelopment DeploymentMode = `DEPLOYMENT_MODE_DEVELOPMENT`

const DeploymentModeDeploymentModeProduction DeploymentMode = `DEPLOYMENT_MODE_PRODUCTION`

// String representation for [fmt.Print]
func (f *DeploymentMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentMode) Set(v string) error {
	switch v {
	case `DEPLOYMENT_MODE_DEVELOPMENT`, `DEPLOYMENT_MODE_PRODUCTION`:
		*f = DeploymentMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_MODE_DEVELOPMENT", "DEPLOYMENT_MODE_PRODUCTION"`, v)
	}
}

// Values returns all possible values for DeploymentMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeploymentMode) Values() []DeploymentMode {
	return []DeploymentMode{
		DeploymentModeDeploymentModeDevelopment,
		DeploymentModeDeploymentModeProduction,
	}
}

// Type always returns DeploymentMode to satisfy [pflag.Value] interface
func (f *DeploymentMode) Type() string {
	return "DeploymentMode"
}

// Type of a deployment resource.
type DeploymentResourceType string

const DeploymentResourceTypeDeploymentResourceTypeAlert DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_ALERT`

const DeploymentResourceTypeDeploymentResourceTypeApp DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_APP`

const DeploymentResourceTypeDeploymentResourceTypeCatalog DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_CATALOG`

const DeploymentResourceTypeDeploymentResourceTypeCluster DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_CLUSTER`

const DeploymentResourceTypeDeploymentResourceTypeDashboard DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DASHBOARD`

const DeploymentResourceTypeDeploymentResourceTypeDatabaseCatalog DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG`

const DeploymentResourceTypeDeploymentResourceTypeDatabaseInstance DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE`

const DeploymentResourceTypeDeploymentResourceTypeExperiment DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT`

const DeploymentResourceTypeDeploymentResourceTypeExternalLocation DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION`

const DeploymentResourceTypeDeploymentResourceTypeJob DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_JOB`

const DeploymentResourceTypeDeploymentResourceTypeModel DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_MODEL`

const DeploymentResourceTypeDeploymentResourceTypeModelServingEndpoint DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT`

const DeploymentResourceTypeDeploymentResourceTypePipeline DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_PIPELINE`

const DeploymentResourceTypeDeploymentResourceTypePostgresBranch DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH`

const DeploymentResourceTypeDeploymentResourceTypePostgresEndpoint DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT`

const DeploymentResourceTypeDeploymentResourceTypePostgresProject DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT`

const DeploymentResourceTypeDeploymentResourceTypeQualityMonitor DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR`

const DeploymentResourceTypeDeploymentResourceTypeRegisteredModel DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL`

const DeploymentResourceTypeDeploymentResourceTypeSchema DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SCHEMA`

const DeploymentResourceTypeDeploymentResourceTypeSecretScope DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE`

const DeploymentResourceTypeDeploymentResourceTypeSqlWarehouse DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE`

const DeploymentResourceTypeDeploymentResourceTypeSyncedDatabaseTable DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE`

const DeploymentResourceTypeDeploymentResourceTypeVolume DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_VOLUME`

// String representation for [fmt.Print]
func (f *DeploymentResourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentResourceType) Set(v string) error {
	switch v {
	case `DEPLOYMENT_RESOURCE_TYPE_ALERT`, `DEPLOYMENT_RESOURCE_TYPE_APP`, `DEPLOYMENT_RESOURCE_TYPE_CATALOG`, `DEPLOYMENT_RESOURCE_TYPE_CLUSTER`, `DEPLOYMENT_RESOURCE_TYPE_DASHBOARD`, `DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG`, `DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE`, `DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT`, `DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION`, `DEPLOYMENT_RESOURCE_TYPE_JOB`, `DEPLOYMENT_RESOURCE_TYPE_MODEL`, `DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT`, `DEPLOYMENT_RESOURCE_TYPE_PIPELINE`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT`, `DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR`, `DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL`, `DEPLOYMENT_RESOURCE_TYPE_SCHEMA`, `DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE`, `DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE`, `DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE`, `DEPLOYMENT_RESOURCE_TYPE_VOLUME`:
		*f = DeploymentResourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_RESOURCE_TYPE_ALERT", "DEPLOYMENT_RESOURCE_TYPE_APP", "DEPLOYMENT_RESOURCE_TYPE_CATALOG", "DEPLOYMENT_RESOURCE_TYPE_CLUSTER", "DEPLOYMENT_RESOURCE_TYPE_DASHBOARD", "DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG", "DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE", "DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT", "DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION", "DEPLOYMENT_RESOURCE_TYPE_JOB", "DEPLOYMENT_RESOURCE_TYPE_MODEL", "DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT", "DEPLOYMENT_RESOURCE_TYPE_PIPELINE", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT", "DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR", "DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL", "DEPLOYMENT_RESOURCE_TYPE_SCHEMA", "DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE", "DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE", "DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE", "DEPLOYMENT_RESOURCE_TYPE_VOLUME"`, v)
	}
}

// Values returns all possible values for DeploymentResourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeploymentResourceType) Values() []DeploymentResourceType {
	return []DeploymentResourceType{
		DeploymentResourceTypeDeploymentResourceTypeAlert,
		DeploymentResourceTypeDeploymentResourceTypeApp,
		DeploymentResourceTypeDeploymentResourceTypeCatalog,
		DeploymentResourceTypeDeploymentResourceTypeCluster,
		DeploymentResourceTypeDeploymentResourceTypeDashboard,
		DeploymentResourceTypeDeploymentResourceTypeDatabaseCatalog,
		DeploymentResourceTypeDeploymentResourceTypeDatabaseInstance,
		DeploymentResourceTypeDeploymentResourceTypeExperiment,
		DeploymentResourceTypeDeploymentResourceTypeExternalLocation,
		DeploymentResourceTypeDeploymentResourceTypeJob,
		DeploymentResourceTypeDeploymentResourceTypeModel,
		DeploymentResourceTypeDeploymentResourceTypeModelServingEndpoint,
		DeploymentResourceTypeDeploymentResourceTypePipeline,
		DeploymentResourceTypeDeploymentResourceTypePostgresBranch,
		DeploymentResourceTypeDeploymentResourceTypePostgresEndpoint,
		DeploymentResourceTypeDeploymentResourceTypePostgresProject,
		DeploymentResourceTypeDeploymentResourceTypeQualityMonitor,
		DeploymentResourceTypeDeploymentResourceTypeRegisteredModel,
		DeploymentResourceTypeDeploymentResourceTypeSchema,
		DeploymentResourceTypeDeploymentResourceTypeSecretScope,
		DeploymentResourceTypeDeploymentResourceTypeSqlWarehouse,
		DeploymentResourceTypeDeploymentResourceTypeSyncedDatabaseTable,
		DeploymentResourceTypeDeploymentResourceTypeVolume,
	}
}

// Type always returns DeploymentResourceType to satisfy [pflag.Value] interface
func (f *DeploymentResourceType) Type() string {
	return "DeploymentResourceType"
}

// Status of a deployment.
type DeploymentStatus string

const DeploymentStatusDeploymentStatusActive DeploymentStatus = `DEPLOYMENT_STATUS_ACTIVE`

const DeploymentStatusDeploymentStatusDeleted DeploymentStatus = `DEPLOYMENT_STATUS_DELETED`

const DeploymentStatusDeploymentStatusFailed DeploymentStatus = `DEPLOYMENT_STATUS_FAILED`

const DeploymentStatusDeploymentStatusInProgress DeploymentStatus = `DEPLOYMENT_STATUS_IN_PROGRESS`

// String representation for [fmt.Print]
func (f *DeploymentStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentStatus) Set(v string) error {
	switch v {
	case `DEPLOYMENT_STATUS_ACTIVE`, `DEPLOYMENT_STATUS_DELETED`, `DEPLOYMENT_STATUS_FAILED`, `DEPLOYMENT_STATUS_IN_PROGRESS`:
		*f = DeploymentStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_STATUS_ACTIVE", "DEPLOYMENT_STATUS_DELETED", "DEPLOYMENT_STATUS_FAILED", "DEPLOYMENT_STATUS_IN_PROGRESS"`, v)
	}
}

// Values returns all possible values for DeploymentStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		DeploymentStatusDeploymentStatusActive,
		DeploymentStatusDeploymentStatusDeleted,
		DeploymentStatusDeploymentStatusFailed,
		DeploymentStatusDeploymentStatusInProgress,
	}
}

// Type always returns DeploymentStatus to satisfy [pflag.Value] interface
func (f *DeploymentStatus) Type() string {
	return "DeploymentStatus"
}

type GetDeploymentRequest struct {
	// Resource name of the deployment to retrieve. Format:
	// deployments/{deployment_id}
	Name string `json:"-" url:"-"`
}

type GetOperationRequest struct {
	// The name of the resource operation to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name string `json:"-" url:"-"`
}

type GetResourceRequest struct {
	// The name of the resource to retrieve. Format:
	// deployments/{deployment_id}/resources/{resource_key}
	Name string `json:"-" url:"-"`
}

type GetVersionRequest struct {
	// The name of the version to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"-" url:"-"`
}

// A request to send a heartbeat for a Version.
type HeartbeatRequest struct {
	// The version whose lock to renew. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"-" url:"-"`
}

// Response for Heartbeat.
type HeartbeatResponse struct {
	// The new lock expiry time after renewal.
	ExpireTime *time.Time `json:"expire_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *HeartbeatResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s HeartbeatResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDeploymentsRequest struct {
	// The maximum number of deployments to return. The service may return fewer
	// than this value. If unspecified, at most 50 deployments will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDeployments` call. Provide
	// this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDeploymentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDeploymentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListDeployments.
type ListDeploymentsResponse struct {
	// The deployments from the queried workspace.
	Deployments []Deployment `json:"deployments,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDeploymentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDeploymentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListOperationsRequest struct {
	// The maximum number of operations to return. The service may return fewer
	// than this value. If unspecified, at most 50 operations will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListOperations` call. Provide
	// this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent version. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOperationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOperationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListOperations.
type ListOperationsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The resource operations under the specified version.
	Operations []Operation `json:"operations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOperationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOperationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResourcesRequest struct {
	// The maximum number of resources to return. The service may return fewer
	// than this value. If unspecified, at most 50 resources will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListResources` call. Provide this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent deployment. Format: deployments/{deployment_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListResourcesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListResourcesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListResources.
type ListResourcesResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The resources under the specified deployment.
	Resources []Resource `json:"resources,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListResourcesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListResourcesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVersionsRequest struct {
	// The maximum number of versions to return. The service may return fewer
	// than this value. If unspecified, at most 50 versions will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListVersions` call. Provide this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent deployment. Format: deployments/{deployment_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListVersions.
type ListVersionsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The versions under the specified deployment.
	Versions []Version `json:"versions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An operation on a single resource performed during a version. Operations are
// append-only and record the result of applying a resource change to the
// workspace.
type Operation struct {
	// The type of operation performed on this resource.
	ActionType OperationActionType `json:"action_type"`
	// When the operation was recorded.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Error message if the operation failed. Set when status is
	// OPERATION_STATUS_FAILED. Captures the error encountered while applying
	// the resource to the workspace.
	ErrorMessage string `json:"error_message,omitempty"`
	// Resource name of the operation. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name string `json:"name,omitempty"`
	// ID reference for the actual resource in the workspace (e.g. the job ID,
	// pipeline ID).
	ResourceId string `json:"resource_id"`
	// Resource identifier within the bundle (e.g. "jobs.foo", "pipelines.bar",
	// "jobs.foo.permissions", "files.<rel-path>"). Can be an arbitrary UTF-8
	// encoded string key. This key links the operation to the corresponding
	// deployment-level Resource.
	ResourceKey string `json:"resource_key,omitempty"`
	// The type of the deployment resource this operation applies to. Derived
	// from the `resource_key` prefix (e.g. "jobs" → JOB); the caller does not
	// set this field.
	ResourceType DeploymentResourceType `json:"resource_type,omitempty"`
	// Serialized local config state after the operation. Should be unset for
	// delete operations.
	State *json.RawMessage `json:"state,omitempty"`
	// Whether the operation succeeded or failed.
	Status OperationStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of action performed on a resource during a deployment.
type OperationActionType string

const OperationActionTypeOperationActionTypeBind OperationActionType = `OPERATION_ACTION_TYPE_BIND`

const OperationActionTypeOperationActionTypeBindAndUpdate OperationActionType = `OPERATION_ACTION_TYPE_BIND_AND_UPDATE`

const OperationActionTypeOperationActionTypeCreate OperationActionType = `OPERATION_ACTION_TYPE_CREATE`

const OperationActionTypeOperationActionTypeDelete OperationActionType = `OPERATION_ACTION_TYPE_DELETE`

const OperationActionTypeOperationActionTypeInitialRegister OperationActionType = `OPERATION_ACTION_TYPE_INITIAL_REGISTER`

const OperationActionTypeOperationActionTypeRecreate OperationActionType = `OPERATION_ACTION_TYPE_RECREATE`

const OperationActionTypeOperationActionTypeResize OperationActionType = `OPERATION_ACTION_TYPE_RESIZE`

const OperationActionTypeOperationActionTypeUpdate OperationActionType = `OPERATION_ACTION_TYPE_UPDATE`

const OperationActionTypeOperationActionTypeUpdateWithId OperationActionType = `OPERATION_ACTION_TYPE_UPDATE_WITH_ID`

// String representation for [fmt.Print]
func (f *OperationActionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OperationActionType) Set(v string) error {
	switch v {
	case `OPERATION_ACTION_TYPE_BIND`, `OPERATION_ACTION_TYPE_BIND_AND_UPDATE`, `OPERATION_ACTION_TYPE_CREATE`, `OPERATION_ACTION_TYPE_DELETE`, `OPERATION_ACTION_TYPE_INITIAL_REGISTER`, `OPERATION_ACTION_TYPE_RECREATE`, `OPERATION_ACTION_TYPE_RESIZE`, `OPERATION_ACTION_TYPE_UPDATE`, `OPERATION_ACTION_TYPE_UPDATE_WITH_ID`:
		*f = OperationActionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPERATION_ACTION_TYPE_BIND", "OPERATION_ACTION_TYPE_BIND_AND_UPDATE", "OPERATION_ACTION_TYPE_CREATE", "OPERATION_ACTION_TYPE_DELETE", "OPERATION_ACTION_TYPE_INITIAL_REGISTER", "OPERATION_ACTION_TYPE_RECREATE", "OPERATION_ACTION_TYPE_RESIZE", "OPERATION_ACTION_TYPE_UPDATE", "OPERATION_ACTION_TYPE_UPDATE_WITH_ID"`, v)
	}
}

// Values returns all possible values for OperationActionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OperationActionType) Values() []OperationActionType {
	return []OperationActionType{
		OperationActionTypeOperationActionTypeBind,
		OperationActionTypeOperationActionTypeBindAndUpdate,
		OperationActionTypeOperationActionTypeCreate,
		OperationActionTypeOperationActionTypeDelete,
		OperationActionTypeOperationActionTypeInitialRegister,
		OperationActionTypeOperationActionTypeRecreate,
		OperationActionTypeOperationActionTypeResize,
		OperationActionTypeOperationActionTypeUpdate,
		OperationActionTypeOperationActionTypeUpdateWithId,
	}
}

// Type always returns OperationActionType to satisfy [pflag.Value] interface
func (f *OperationActionType) Type() string {
	return "OperationActionType"
}

// Status of a resource operation.
type OperationStatus string

const OperationStatusOperationStatusFailed OperationStatus = `OPERATION_STATUS_FAILED`

const OperationStatusOperationStatusSucceeded OperationStatus = `OPERATION_STATUS_SUCCEEDED`

// String representation for [fmt.Print]
func (f *OperationStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OperationStatus) Set(v string) error {
	switch v {
	case `OPERATION_STATUS_FAILED`, `OPERATION_STATUS_SUCCEEDED`:
		*f = OperationStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPERATION_STATUS_FAILED", "OPERATION_STATUS_SUCCEEDED"`, v)
	}
}

// Values returns all possible values for OperationStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		OperationStatusOperationStatusFailed,
		OperationStatusOperationStatusSucceeded,
	}
}

// Type always returns OperationStatus to satisfy [pflag.Value] interface
func (f *OperationStatus) Type() string {
	return "OperationStatus"
}

// A resource managed by a deployment. Resources are implicitly created,
// updated, or deleted when operations are recorded on a version.
type Resource struct {
	// The action performed on this resource during the last version.
	LastActionType OperationActionType `json:"last_action_type,omitempty"`
	// The version_id of the last version where this resource was updated.
	LastVersionId string `json:"last_version_id,omitempty"`
	// Resource name. Format:
	// deployments/{deployment_id}/resources/{resource_key}
	Name string `json:"name,omitempty"`
	// ID that references the actual resource in the workspace (e.g. the job ID,
	// pipeline ID).
	ResourceId string `json:"resource_id,omitempty"`
	// Resource identifier within the bundle (e.g. "jobs.foo", "pipelines.bar",
	// "jobs.foo.permissions").
	ResourceKey string `json:"resource_key,omitempty"`
	// The type of the deployment resource.
	ResourceType DeploymentResourceType `json:"resource_type"`
	// Serialized local config state (what the CLI deployed).
	State *json.RawMessage `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Resource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Resource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single invocation of a deploy or destroy command against a deployment.
// Creating a version acquires an exclusive lock on the parent deployment.
type Version struct {
	// CLI version used to initiate the version.
	CliVersion string `json:"cli_version"`
	// When the version completed. Unset while the version is in progress.
	CompleteTime *time.Time `json:"complete_time,omitempty"`
	// The user who completed the version (email or principal name). May differ
	// from `created_by` when another user force-completes the version.
	CompletedBy string `json:"completed_by,omitempty"`
	// Why the version was completed. Unset while in progress. Set when status
	// transitions to COMPLETED.
	CompletionReason VersionComplete `json:"completion_reason,omitempty"`
	// When the version was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The user who created the version (email or principal name).
	CreatedBy string `json:"created_by,omitempty"`
	// Bundle target deployment mode (development or production), captured at
	// the time of this version.
	DeploymentMode DeploymentMode `json:"deployment_mode,omitempty"`
	// Display name for the deployment, captured at the time of this version.
	DisplayName string `json:"display_name,omitempty"`
	// Resource name of the version. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"name,omitempty"`
	// Status of the version: IN_PROGRESS or COMPLETED.
	Status VersionStatus `json:"status,omitempty"`
	// Target name of the deployment, captured at the time of this version.
	TargetName string `json:"target_name,omitempty"`
	// Monotonically increasing version identifier within the parent deployment.
	// Assigned by the client on creation.
	VersionId string `json:"version_id,omitempty"`
	// Type of version (deploy or destroy).
	VersionType VersionType `json:"version_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Version) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Version) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reason why a version was completed.
type VersionComplete string

const VersionCompleteVersionCompleteFailure VersionComplete = `VERSION_COMPLETE_FAILURE`

const VersionCompleteVersionCompleteForceAbort VersionComplete = `VERSION_COMPLETE_FORCE_ABORT`

const VersionCompleteVersionCompleteLeaseExpired VersionComplete = `VERSION_COMPLETE_LEASE_EXPIRED`

const VersionCompleteVersionCompleteSuccess VersionComplete = `VERSION_COMPLETE_SUCCESS`

// String representation for [fmt.Print]
func (f *VersionComplete) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VersionComplete) Set(v string) error {
	switch v {
	case `VERSION_COMPLETE_FAILURE`, `VERSION_COMPLETE_FORCE_ABORT`, `VERSION_COMPLETE_LEASE_EXPIRED`, `VERSION_COMPLETE_SUCCESS`:
		*f = VersionComplete(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "VERSION_COMPLETE_FAILURE", "VERSION_COMPLETE_FORCE_ABORT", "VERSION_COMPLETE_LEASE_EXPIRED", "VERSION_COMPLETE_SUCCESS"`, v)
	}
}

// Values returns all possible values for VersionComplete.
//
// There is no guarantee on the order of the values in the slice.
func (f *VersionComplete) Values() []VersionComplete {
	return []VersionComplete{
		VersionCompleteVersionCompleteFailure,
		VersionCompleteVersionCompleteForceAbort,
		VersionCompleteVersionCompleteLeaseExpired,
		VersionCompleteVersionCompleteSuccess,
	}
}

// Type always returns VersionComplete to satisfy [pflag.Value] interface
func (f *VersionComplete) Type() string {
	return "VersionComplete"
}

// Status of a version.
type VersionStatus string

const VersionStatusVersionStatusCompleted VersionStatus = `VERSION_STATUS_COMPLETED`

const VersionStatusVersionStatusInProgress VersionStatus = `VERSION_STATUS_IN_PROGRESS`

// String representation for [fmt.Print]
func (f *VersionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VersionStatus) Set(v string) error {
	switch v {
	case `VERSION_STATUS_COMPLETED`, `VERSION_STATUS_IN_PROGRESS`:
		*f = VersionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "VERSION_STATUS_COMPLETED", "VERSION_STATUS_IN_PROGRESS"`, v)
	}
}

// Values returns all possible values for VersionStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *VersionStatus) Values() []VersionStatus {
	return []VersionStatus{
		VersionStatusVersionStatusCompleted,
		VersionStatusVersionStatusInProgress,
	}
}

// Type always returns VersionStatus to satisfy [pflag.Value] interface
func (f *VersionStatus) Type() string {
	return "VersionStatus"
}

// Type of version.
type VersionType string

const VersionTypeVersionTypeDeploy VersionType = `VERSION_TYPE_DEPLOY`

const VersionTypeVersionTypeDestroy VersionType = `VERSION_TYPE_DESTROY`

// String representation for [fmt.Print]
func (f *VersionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VersionType) Set(v string) error {
	switch v {
	case `VERSION_TYPE_DEPLOY`, `VERSION_TYPE_DESTROY`:
		*f = VersionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "VERSION_TYPE_DEPLOY", "VERSION_TYPE_DESTROY"`, v)
	}
}

// Values returns all possible values for VersionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *VersionType) Values() []VersionType {
	return []VersionType{
		VersionTypeVersionTypeDeploy,
		VersionTypeVersionTypeDestroy,
	}
}

// Type always returns VersionType to satisfy [pflag.Value] interface
func (f *VersionType) Type() string {
	return "VersionType"
}
