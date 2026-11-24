// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package bundledeployments

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// A request to complete a Version.
type CompleteVersionRequest struct {
	// The reason for completing the version. Must be VERSION_COMPLETE_SUCCESS
	// or VERSION_COMPLETE_FAILURE.
	CompletionReason VersionComplete `json:"completion_reason"`
	// If true, force-completes the version even if the caller is not the
	// original creator.
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
	// The deployment to create. `initial_parent_path` is required.
	// `display_name`, `target_name`, `deployment_mode`, and `workspace_info`
	// may be set; every other field is assigned by the service and ignored on
	// input.
	Deployment Deployment `json:"deployment"`
}

func (s *CreateDeploymentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateVersionRequest struct {
	// The parent deployment where this version will be created. Format:
	// deployments/{deployment_id}
	Parent string `json:"-" url:"-"`
	// The version to create.
	Version Version `json:"version"`
	// The ID to use for the version, which becomes the final component of the
	// version's resource name. A numeric string (base-10, fits in a signed
	// 64-bit integer) chosen by the caller; must be greater than or equal to 1.
	// Must be numerically greater than the deployment's most recent version
	// (see `version.previous_version_id`); it does not need to start at 1 or
	// increase by exactly 1. If the value is not numerically greater, the
	// server returns `INVALID_PARAMETER_VALUE`.
	VersionId string `json:"-" url:"version_id"`
}

func (s *CreateVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Dashboard-specific per-resource metadata. Set only for dashboard resources.
type DashboardMetadata struct {
	// Path of the file that declares this dashboard, relative to the bundle's
	// workspace.file_path (Deployment.workspace_info.file_path) — join the
	// two to get the file's absolute workspace path.
	//
	// For now this lives only on the dashboard metadata, and is a single string
	// because it was a single string (`relative_path`) in the legacy bundle
	// metadata.json. We may generalize it in the future: lifting it to a
	// top-level field on Resource/Operation (every resource type has a
	// definition location) and converting it to a repeated field, since a
	// resource can be declared across multiple files/locations.
	DefinitionPath string `json:"definition_path,omitempty"`
	// Path of the dashboard's source artifact (its `.lvdash.json`), relative to
	// the deployment root.
	SourcePath string `json:"source_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDeploymentRequest struct {
	// Resource name of the deployment to delete. Format:
	// deployments/{deployment_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteDeploymentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A bundle deployment registered with the control plane.
type Deployment struct {
	// When the deployment was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The user who created the deployment (email or principal name). Empty if
	// authoritative deployment metadata does not identify a creator or the
	// principal cannot be resolved.
	CreatedBy string `json:"created_by,omitempty"`
	// Bundle target deployment mode (development or production).
	DeploymentMode DeploymentMode `json:"deployment_mode,omitempty"`
	// When deletion was recorded. Unset if deletion has not been recorded. This
	// response metadata does not determine the deployment's lifecycle status.
	DestroyTime *time.Time `json:"destroy_time,omitempty"`
	// The user who destroyed the deployment (email or principal name). Unset if
	// the deployment has not been destroyed.
	DestroyedBy string `json:"destroyed_by,omitempty"`
	// Human-readable name for the deployment, up to 256 characters.
	DisplayName string `json:"display_name,omitempty"`
	// Git provenance of the deployment's source, derived from the latest
	// version.
	GitInfo *GitInfo `json:"git_info,omitempty"`
	// The workspace path of the existing folder where the deployment is
	// initially created. Must be absolute and canonical, with single
	// separators, no `.` or `..` segments, and no trailing slash unless the
	// path is `/`. It may contain at most 24 path segments, excluding an
	// optional leading `/Workspace` segment. The complete path may contain up
	// to 1,024 characters, and each segment may contain up to 511 characters.
	// This field is input only and is not returned in create, get, or list
	// responses.
	InitialParentPath string `json:"initial_parent_path,omitempty"`
	// The version_id of the most recent version that completed successfully.
	// Unset until a version has completed successfully. Unlike last_version_id,
	// it is not advanced when a version fails, so it always points at the last
	// known-good deployment state (or is unset if there has never been one).
	LastSuccessfulVersionId string `json:"last_successful_version_id,omitempty"`
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
	// The user who most recently updated the deployment (email or principal
	// name). Empty if authoritative deployment metadata does not identify a
	// modifier or the principal cannot be resolved.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Workspace location of the deployment.
	WorkspaceInfo *WorkspaceInfo `json:"workspace_info,omitempty"`

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

const DeploymentResourceTypeDeploymentResourceTypeClusterPolicy DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_CLUSTER_POLICY`

const DeploymentResourceTypeDeploymentResourceTypeDashboard DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DASHBOARD`

const DeploymentResourceTypeDeploymentResourceTypeDatabaseCatalog DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG`

const DeploymentResourceTypeDeploymentResourceTypeDatabaseInstance DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE`

const DeploymentResourceTypeDeploymentResourceTypeExperiment DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT`

const DeploymentResourceTypeDeploymentResourceTypeExternalLocation DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION`

const DeploymentResourceTypeDeploymentResourceTypeGenieSpace DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_GENIE_SPACE`

const DeploymentResourceTypeDeploymentResourceTypeInstancePool DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_INSTANCE_POOL`

const DeploymentResourceTypeDeploymentResourceTypeJob DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_JOB`

const DeploymentResourceTypeDeploymentResourceTypeJobRun DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_JOB_RUN`

const DeploymentResourceTypeDeploymentResourceTypeModel DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_MODEL`

const DeploymentResourceTypeDeploymentResourceTypeModelServingEndpoint DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT`

const DeploymentResourceTypeDeploymentResourceTypePipeline DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_PIPELINE`

const DeploymentResourceTypeDeploymentResourceTypePostgresBranch DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH`

const DeploymentResourceTypeDeploymentResourceTypePostgresCatalog DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_CATALOG`

const DeploymentResourceTypeDeploymentResourceTypePostgresDatabase DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_DATABASE`

const DeploymentResourceTypeDeploymentResourceTypePostgresEndpoint DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT`

const DeploymentResourceTypeDeploymentResourceTypePostgresProject DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT`

const DeploymentResourceTypeDeploymentResourceTypePostgresRole DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ROLE`

const DeploymentResourceTypeDeploymentResourceTypePostgresSyncedTable DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_SYNCED_TABLE`

const DeploymentResourceTypeDeploymentResourceTypeQualityMonitor DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR`

const DeploymentResourceTypeDeploymentResourceTypeRegisteredModel DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL`

const DeploymentResourceTypeDeploymentResourceTypeSchema DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SCHEMA`

const DeploymentResourceTypeDeploymentResourceTypeSecret DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SECRET`

const DeploymentResourceTypeDeploymentResourceTypeSecretScope DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE`

const DeploymentResourceTypeDeploymentResourceTypeSqlWarehouse DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE`

const DeploymentResourceTypeDeploymentResourceTypeSyncedDatabaseTable DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE`

const DeploymentResourceTypeDeploymentResourceTypeVectorSearchEndpoint DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_ENDPOINT`

const DeploymentResourceTypeDeploymentResourceTypeVectorSearchIndex DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_INDEX`

const DeploymentResourceTypeDeploymentResourceTypeVolume DeploymentResourceType = `DEPLOYMENT_RESOURCE_TYPE_VOLUME`

// String representation for [fmt.Print]
func (f *DeploymentResourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentResourceType) Set(v string) error {
	switch v {
	case `DEPLOYMENT_RESOURCE_TYPE_ALERT`, `DEPLOYMENT_RESOURCE_TYPE_APP`, `DEPLOYMENT_RESOURCE_TYPE_CATALOG`, `DEPLOYMENT_RESOURCE_TYPE_CLUSTER`, `DEPLOYMENT_RESOURCE_TYPE_CLUSTER_POLICY`, `DEPLOYMENT_RESOURCE_TYPE_DASHBOARD`, `DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG`, `DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE`, `DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT`, `DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION`, `DEPLOYMENT_RESOURCE_TYPE_GENIE_SPACE`, `DEPLOYMENT_RESOURCE_TYPE_INSTANCE_POOL`, `DEPLOYMENT_RESOURCE_TYPE_JOB`, `DEPLOYMENT_RESOURCE_TYPE_JOB_RUN`, `DEPLOYMENT_RESOURCE_TYPE_MODEL`, `DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT`, `DEPLOYMENT_RESOURCE_TYPE_PIPELINE`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_CATALOG`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_DATABASE`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ROLE`, `DEPLOYMENT_RESOURCE_TYPE_POSTGRES_SYNCED_TABLE`, `DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR`, `DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL`, `DEPLOYMENT_RESOURCE_TYPE_SCHEMA`, `DEPLOYMENT_RESOURCE_TYPE_SECRET`, `DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE`, `DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE`, `DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE`, `DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_ENDPOINT`, `DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_INDEX`, `DEPLOYMENT_RESOURCE_TYPE_VOLUME`:
		*f = DeploymentResourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_RESOURCE_TYPE_ALERT", "DEPLOYMENT_RESOURCE_TYPE_APP", "DEPLOYMENT_RESOURCE_TYPE_CATALOG", "DEPLOYMENT_RESOURCE_TYPE_CLUSTER", "DEPLOYMENT_RESOURCE_TYPE_CLUSTER_POLICY", "DEPLOYMENT_RESOURCE_TYPE_DASHBOARD", "DEPLOYMENT_RESOURCE_TYPE_DATABASE_CATALOG", "DEPLOYMENT_RESOURCE_TYPE_DATABASE_INSTANCE", "DEPLOYMENT_RESOURCE_TYPE_EXPERIMENT", "DEPLOYMENT_RESOURCE_TYPE_EXTERNAL_LOCATION", "DEPLOYMENT_RESOURCE_TYPE_GENIE_SPACE", "DEPLOYMENT_RESOURCE_TYPE_INSTANCE_POOL", "DEPLOYMENT_RESOURCE_TYPE_JOB", "DEPLOYMENT_RESOURCE_TYPE_JOB_RUN", "DEPLOYMENT_RESOURCE_TYPE_MODEL", "DEPLOYMENT_RESOURCE_TYPE_MODEL_SERVING_ENDPOINT", "DEPLOYMENT_RESOURCE_TYPE_PIPELINE", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_BRANCH", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_CATALOG", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_DATABASE", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ENDPOINT", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_PROJECT", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_ROLE", "DEPLOYMENT_RESOURCE_TYPE_POSTGRES_SYNCED_TABLE", "DEPLOYMENT_RESOURCE_TYPE_QUALITY_MONITOR", "DEPLOYMENT_RESOURCE_TYPE_REGISTERED_MODEL", "DEPLOYMENT_RESOURCE_TYPE_SCHEMA", "DEPLOYMENT_RESOURCE_TYPE_SECRET", "DEPLOYMENT_RESOURCE_TYPE_SECRET_SCOPE", "DEPLOYMENT_RESOURCE_TYPE_SQL_WAREHOUSE", "DEPLOYMENT_RESOURCE_TYPE_SYNCED_DATABASE_TABLE", "DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_ENDPOINT", "DEPLOYMENT_RESOURCE_TYPE_VECTOR_SEARCH_INDEX", "DEPLOYMENT_RESOURCE_TYPE_VOLUME"`, v)
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
		DeploymentResourceTypeDeploymentResourceTypeClusterPolicy,
		DeploymentResourceTypeDeploymentResourceTypeDashboard,
		DeploymentResourceTypeDeploymentResourceTypeDatabaseCatalog,
		DeploymentResourceTypeDeploymentResourceTypeDatabaseInstance,
		DeploymentResourceTypeDeploymentResourceTypeExperiment,
		DeploymentResourceTypeDeploymentResourceTypeExternalLocation,
		DeploymentResourceTypeDeploymentResourceTypeGenieSpace,
		DeploymentResourceTypeDeploymentResourceTypeInstancePool,
		DeploymentResourceTypeDeploymentResourceTypeJob,
		DeploymentResourceTypeDeploymentResourceTypeJobRun,
		DeploymentResourceTypeDeploymentResourceTypeModel,
		DeploymentResourceTypeDeploymentResourceTypeModelServingEndpoint,
		DeploymentResourceTypeDeploymentResourceTypePipeline,
		DeploymentResourceTypeDeploymentResourceTypePostgresBranch,
		DeploymentResourceTypeDeploymentResourceTypePostgresCatalog,
		DeploymentResourceTypeDeploymentResourceTypePostgresDatabase,
		DeploymentResourceTypeDeploymentResourceTypePostgresEndpoint,
		DeploymentResourceTypeDeploymentResourceTypePostgresProject,
		DeploymentResourceTypeDeploymentResourceTypePostgresRole,
		DeploymentResourceTypeDeploymentResourceTypePostgresSyncedTable,
		DeploymentResourceTypeDeploymentResourceTypeQualityMonitor,
		DeploymentResourceTypeDeploymentResourceTypeRegisteredModel,
		DeploymentResourceTypeDeploymentResourceTypeSchema,
		DeploymentResourceTypeDeploymentResourceTypeSecret,
		DeploymentResourceTypeDeploymentResourceTypeSecretScope,
		DeploymentResourceTypeDeploymentResourceTypeSqlWarehouse,
		DeploymentResourceTypeDeploymentResourceTypeSyncedDatabaseTable,
		DeploymentResourceTypeDeploymentResourceTypeVectorSearchEndpoint,
		DeploymentResourceTypeDeploymentResourceTypeVectorSearchIndex,
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

func (s *GetDeploymentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOperationRequest struct {
	// The name of the resource operation to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetResourceRequest struct {
	// The name of the resource to retrieve. Format:
	// deployments/{deployment_id}/resources/{resource_key}
	Name string `json:"-" url:"-"`
}

func (s *GetResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetVersionRequest struct {
	// The name of the version to retrieve. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"-" url:"-"`
}

func (s *GetVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Git provenance of a bundle's source, captured at deploy time. Lets consumers
// link a deployed resource back to its source in version control.
type GitInfo struct {
	// Branch the source was deployed from.
	Branch string `json:"branch,omitempty"`
	// Commit SHA of the deployed source.
	Commit string `json:"commit,omitempty"`
	// URL of the git remote the source was deployed from.
	OriginUrl string `json:"origin_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GitInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A request to send a heartbeat for a Version.
type HeartbeatRequest struct {
	// The version whose lock to renew. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"-" url:"-"`
}

func (s *HeartbeatRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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
	// than this value. If unspecified, at most 20 deployments will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDeployments` call. Provide
	// this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Whether to include soft-deleted deployments in the response. Defaults to
	// false. When true, returns both live and soft-deleted deployments visible
	// to the caller, subject to the filter. Soft-deleted deployments have
	// status `DEPLOYMENT_STATUS_DELETED`.
	ShowDeleted bool `json:"-" url:"show_deleted,omitempty"`

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
	// than this value. If unspecified, at most 20 versions will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
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

// An operation on a single resource performed during a version. The full set of
// operations for a version is recorded when the version is created: each
// carries its `resource_key` and `action_type` and starts in
// `OPERATION_STATUS_PENDING`. As each resource is applied, its operation is
// updated (via UpdateOperation) to record the result of applying the change to
// the workspace. `state`, `error_message`, `resource_id`, `status`, and
// `dashboard_metadata` may be updated afterwards, guarded by `sequence_id` for
// optimistic concurrency control; all other fields are immutable once recorded.
type Operation struct {
	// The type of operation performed on this resource. Set when the version is
	// created and immutable thereafter.
	ActionType OperationActionType `json:"action_type,omitempty"`
	// When the operation was recorded.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Dashboard-specific metadata; set only for dashboard resources. Mutable:
	// may be set or updated via UpdateOperation as the resource is applied, and
	// is mirrored onto the corresponding deployment-level resource.
	DashboardMetadata *DashboardMetadata `json:"dashboard_metadata,omitempty"`
	// Error message if the operation failed. Set when status is
	// OPERATION_STATUS_FAILED. Captures the error encountered while applying
	// the resource to the workspace. Mutable: may be updated after creation via
	// UpdateOperation; setting it to an empty string clears it. After an update
	// is applied, an operation whose status is OPERATION_STATUS_SUCCEEDED
	// cannot carry an error_message.
	ErrorMessage string `json:"error_message,omitempty"`
	// Resource name of the operation. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name string `json:"name,omitempty"`
	// ID of the actual resource in the workspace (e.g. the job ID, pipeline
	// ID). Required whenever `state` is set, because state records a resource
	// that exists. A CREATE or RECREATE that has not produced its resource yet
	// records neither. Mutable: may be filled in (or corrected) later via
	// UpdateOperation once the ID is known.
	ResourceId string `json:"resource_id,omitempty"`
	// Resource identifier within the bundle (e.g. "jobs.foo", "pipelines.bar",
	// "jobs.foo.permissions", "files.<rel-path>"). Can be an arbitrary UTF-8
	// encoded string key. This key links the operation to the corresponding
	// deployment-level Resource. Set when the version is created and immutable
	// thereafter.
	ResourceKey string `json:"resource_key,omitempty"`
	// The type of the deployment resource this operation applies to. Derived
	// from the `resource_key` prefix (e.g. "jobs" → JOB); the caller does not
	// set this field.
	ResourceType DeploymentResourceType `json:"resource_type,omitempty"`
	// Monotonically increasing revision used for optimistic concurrency control
	// (the AIP-154 concurrency token for this resource, realized as a sequence
	// number rather than an opaque etag). The server assigns 0 when the
	// operation is created and increments it on every successful
	// UpdateOperation, so a never-updated operation is at 0 and the first
	// successful update makes it 1. It is OPTIONAL rather than OUTPUT_ONLY
	// because it is dual-purpose: GetOperation returns the current value, and
	// UpdateOperation reads the caller-supplied value as a precondition. The
	// caller must echo the value it last observed; if it no longer matches the
	// server's value, the update is rejected with ABORTED so the caller can
	// re-read and retry.
	SequenceId int64 `json:"sequence_id,omitempty"`
	// Serialized local config state after the operation. Its presence records
	// whether the resource is still tracked by the deployment. Explicitly
	// clearing state removes the resource from the deployment; omitting it from
	// the update mask leaves resource state unchanged. State is stored only on
	// Resources. UpdateOperation accepts state but omits it from its response,
	// as do subsequent GetOperation and ListOperations.
	//
	// Mutable: may be updated after creation via UpdateOperation. When
	// updating, the caller must echo the last-observed `sequence_id` as a
	// concurrency precondition.
	//
	// The string is not parsed; it is projected onto the corresponding
	// Resource's state. It is deliberately not google.protobuf.Value, whose
	// only numeric case is `double number_value`, so parsing the client's JSON
	// into it rewrites every integer as a double - `1` reads back as `1.0`,
	// which no longer deserializes into an integer field - and silently loses
	// precision above 2^53, which is within range for IDs the client records.
	//
	// A string rather than bytes: the payload is always UTF-8 JSON, and proto3
	// JSON maps bytes to base64, which inflates every request and response by a
	// third and makes state unreadable in logs and API responses. Both generate
	// the same OpenAPI schema ("type": "string"), so the SDKs are identical
	// either way.
	State string `json:"state,omitempty"`
	// Status of the operation. Starts as OPERATION_STATUS_PENDING when the
	// version is created and moves to a terminal status once the resource is
	// applied. Mutable: updated via UpdateOperation, e.g. when an operation
	// recorded as failed is retried and eventually succeeds.
	Status OperationStatus `json:"status,omitempty"`
	// When the operation was last updated. Set to `create_time` when the
	// operation is created and to the server timestamp on each successful
	// UpdateOperation.
	UpdateTime *time.Time `json:"update_time,omitempty"`

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

const OperationActionTypeOperationActionTypeUnbind OperationActionType = `OPERATION_ACTION_TYPE_UNBIND`

const OperationActionTypeOperationActionTypeUpdate OperationActionType = `OPERATION_ACTION_TYPE_UPDATE`

const OperationActionTypeOperationActionTypeUpdateWithId OperationActionType = `OPERATION_ACTION_TYPE_UPDATE_WITH_ID`

// String representation for [fmt.Print]
func (f *OperationActionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OperationActionType) Set(v string) error {
	switch v {
	case `OPERATION_ACTION_TYPE_BIND`, `OPERATION_ACTION_TYPE_BIND_AND_UPDATE`, `OPERATION_ACTION_TYPE_CREATE`, `OPERATION_ACTION_TYPE_DELETE`, `OPERATION_ACTION_TYPE_INITIAL_REGISTER`, `OPERATION_ACTION_TYPE_RECREATE`, `OPERATION_ACTION_TYPE_RESIZE`, `OPERATION_ACTION_TYPE_UNBIND`, `OPERATION_ACTION_TYPE_UPDATE`, `OPERATION_ACTION_TYPE_UPDATE_WITH_ID`:
		*f = OperationActionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPERATION_ACTION_TYPE_BIND", "OPERATION_ACTION_TYPE_BIND_AND_UPDATE", "OPERATION_ACTION_TYPE_CREATE", "OPERATION_ACTION_TYPE_DELETE", "OPERATION_ACTION_TYPE_INITIAL_REGISTER", "OPERATION_ACTION_TYPE_RECREATE", "OPERATION_ACTION_TYPE_RESIZE", "OPERATION_ACTION_TYPE_UNBIND", "OPERATION_ACTION_TYPE_UPDATE", "OPERATION_ACTION_TYPE_UPDATE_WITH_ID"`, v)
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
		OperationActionTypeOperationActionTypeUnbind,
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

const OperationStatusOperationStatusPending OperationStatus = `OPERATION_STATUS_PENDING`

const OperationStatusOperationStatusSucceeded OperationStatus = `OPERATION_STATUS_SUCCEEDED`

// String representation for [fmt.Print]
func (f *OperationStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OperationStatus) Set(v string) error {
	switch v {
	case `OPERATION_STATUS_FAILED`, `OPERATION_STATUS_PENDING`, `OPERATION_STATUS_SUCCEEDED`:
		*f = OperationStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPERATION_STATUS_FAILED", "OPERATION_STATUS_PENDING", "OPERATION_STATUS_SUCCEEDED"`, v)
	}
}

// Values returns all possible values for OperationStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		OperationStatusOperationStatusFailed,
		OperationStatusOperationStatusPending,
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
	// Dashboard-specific metadata; set only for dashboard resources.
	DashboardMetadata *DashboardMetadata `json:"dashboard_metadata,omitempty"`
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
	// Serialized local config state (what the CLI deployed). Opaque to this
	// service; see Operation.state for why this is a string and not
	// google.protobuf.Value.
	State string `json:"state,omitempty"`
	// When the last operation that updated this resource's recorded state was
	// applied. Pairs with last_action_type and last_version_id (all three
	// advance together on that write).
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Resource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Resource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A resource operation to record when a version is created. Each staged
// operation identifies the resource it applies to and the action planned for
// it; the server records the operation in `OPERATION_STATUS_PENDING`, and its
// outcome is filled in later via UpdateOperation.
type StagedOperation struct {
	// The type of operation planned for this resource.
	ActionType OperationActionType `json:"action_type"`
	// The key identifying the resource this operation applies to (e.g.
	// "jobs.foo", "pipelines.bar"). Becomes the final component of the
	// operation's name and must be unique among the operations in the version.
	ResourceKey string `json:"resource_key"`
}

func (s *StagedOperation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDeploymentRequest struct {
	// The deployment to update. Its `name` selects the deployment; the fields
	// named in `update_mask` carry the new values. All other fields are
	// ignored.
	Deployment Deployment `json:"deployment"`
	// Resource name of the deployment. Format: deployments/{deployment_id}
	Name string `json:"-" url:"-"`
	// The fields to update; supported paths are `display_name`,
	// `deployment_mode`, `target_name`, and `workspace_info`. An empty mask or
	// any other path returns INVALID_PARAMETER_VALUE.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateDeploymentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateOperationRequest struct {
	// Resource name of the operation. Format:
	// deployments/{deployment_id}/versions/{version_id}/operations/{resource_key}
	Name string `json:"-" url:"-"`
	// The operation to update. Its `name` selects the operation; the fields
	// named in `update_mask` carry the new values; and `sequence_id` carries
	// the optimistic-concurrency precondition (see the field docs on
	// Operation). All other fields are ignored.
	Operation Operation `json:"operation"`
	// The set of fields to update. Required; supported paths are `state`,
	// `error_message`, `resource_id`, `status`, and `dashboard_metadata`. An
	// empty mask or any other path is rejected with INVALID_PARAMETER_VALUE.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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
	// Display name for the deployment, captured at the time of this version. Up
	// to 256 characters. When present, creating the version updates the
	// deployment display name. An empty value clears it; an absent value leaves
	// the current deployment display name unchanged.
	DisplayName string `json:"display_name,omitempty"`
	// Git provenance of the source, captured at the time of this version.
	GitInfo *GitInfo `json:"git_info,omitempty"`
	// Resource name of the version. Format:
	// deployments/{deployment_id}/versions/{version_id}
	Name string `json:"name,omitempty"`
	// The full operation plan for this version: one PENDING operation per
	// entry, recorded in the same transaction. Input only -- supplied on create
	// and never returned by create/get/list; read the recorded operations via
	// ListOperations.
	Operations []StagedOperation `json:"operations,omitempty"`
	// The version_id this version was created on top of — the deployment's
	// most recent version at creation time. Leave unset when creating the first
	// version (the deployment has no prior versions). Set by the client on
	// creation and immutable thereafter.
	//
	// Acts as an optimistic-concurrency precondition: the server requires it to
	// equal the deployment's current most-recent version (and to be unset when
	// the deployment has no versions) and returns `INVALID_PARAMETER_VALUE` on
	// mismatch, so a deploy racing against a concurrent deploy is rejected
	// rather than silently overwriting it.
	PreviousVersionId string `json:"previous_version_id,omitempty"`
	// Status of the version: IN_PROGRESS or COMPLETED.
	Status VersionStatus `json:"status,omitempty"`
	// Target name of the deployment, captured at the time of this version.
	TargetName string `json:"target_name,omitempty"`
	// Version identifier within the parent deployment, assigned by the client
	// on creation. A numeric string (base-10, fits in a signed 64-bit integer)
	// that is greater than or equal to 1. Version IDs are strictly increasing
	// within a deployment but are not required to start at 1 or to be
	// contiguous.
	VersionId string `json:"version_id,omitempty"`
	// Type of version (deploy or destroy).
	VersionType VersionType `json:"version_type"`
	// Workspace location of the deployment, captured at the time of this
	// version.
	WorkspaceInfo *WorkspaceInfo `json:"workspace_info,omitempty"`

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

// Workspace location of a bundle deployment, captured at deploy time.
type WorkspaceInfo struct {
	// Path of the bundle root (the directory containing databricks.yml)
	// relative to git_folder_path. Empty when the deployment is not from a
	// Databricks Git folder.
	BundleRootPath string `json:"bundle_root_path,omitempty"`
	// Absolute workspace path where the deployed bundle files live. Mirrors the
	// workspace.file_path field in DABs bundle config.
	FilePath string `json:"file_path,omitempty"`
	// When deployed from a Databricks Git folder, the absolute workspace path
	// of that folder; empty for local deploys.
	GitFolderPath string `json:"git_folder_path,omitempty"`
	// Absolute workspace path of the deployment root — the base path the
	// deployed files live under. Mirrors workspace.root_path in the DABs bundle
	// config; file_path is its files subdirectory.
	RootPath string `json:"root_path,omitempty"`
	// Whether files are served directly from the source sync root instead of
	// being copied into file_path.
	SourceLinked bool `json:"source_linked,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
