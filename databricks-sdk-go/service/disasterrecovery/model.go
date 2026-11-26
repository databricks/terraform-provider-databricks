// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateFailoverGroupRequest struct {
	// The failover group to create.
	FailoverGroup FailoverGroup `json:"failover_group"`
	// Client-provided identifier for the failover group. Used to construct the
	// resource name as {parent}/failover-groups/{failover_group_id}.
	FailoverGroupId string `json:"-" url:"failover_group_id"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`
	// When true, validates the request without creating the failover group.
	ValidateOnly bool `json:"-" url:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStableUrlRequest struct {
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`
	// The stable URL to create.
	StableUrl StableUrl `json:"stable_url"`
	// Client-provided identifier for the stable URL. Used to construct the
	// resource name as {parent}/stable-urls/{stable_url_id}.
	StableUrlId string `json:"-" url:"stable_url_id"`
	// When true, validates the request without creating the stable URL.
	ValidateOnly bool `json:"-" url:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStableUrlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStableUrlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the delete proceeds without an etag check.
	Etag string `json:"-" url:"etag,omitempty"`
	// The fully qualified resource name of the failover group to delete.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"-" url:"-"`
}

func (s *DeleteStableUrlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to failover a failover group to a new primary region.
type FailoverFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the failover proceeds regardless of current
	// state.
	Etag string `json:"etag,omitempty"`
	// The type of failover to perform. FORCED is the only supported type.
	FailoverType FailoverFailoverGroupRequestFailoverType `json:"failover_type"`
	// The fully qualified resource name of the failover group to failover.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
	// The target primary region. Must be one of the participating regions and
	// different from the current effective_primary_region. Serves as an
	// idempotency check.
	TargetPrimaryRegion string `json:"target_primary_region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailoverFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailoverFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of failover to perform.
type FailoverFailoverGroupRequestFailoverType string

const FailoverFailoverGroupRequestFailoverTypeForced FailoverFailoverGroupRequestFailoverType = `FORCED`

// String representation for [fmt.Print]
func (f *FailoverFailoverGroupRequestFailoverType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FailoverFailoverGroupRequestFailoverType) Set(v string) error {
	switch v {
	case `FORCED`:
		*f = FailoverFailoverGroupRequestFailoverType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FORCED"`, v)
	}
}

// Values returns all possible values for FailoverFailoverGroupRequestFailoverType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FailoverFailoverGroupRequestFailoverType) Values() []FailoverFailoverGroupRequestFailoverType {
	return []FailoverFailoverGroupRequestFailoverType{
		FailoverFailoverGroupRequestFailoverTypeForced,
	}
}

// Type always returns FailoverFailoverGroupRequestFailoverType to satisfy [pflag.Value] interface
func (f *FailoverFailoverGroupRequestFailoverType) Type() string {
	return "FailoverFailoverGroupRequestFailoverType"
}

// A failover group manages Databricks Managed Disaster Recovery for a pair of
// workspaces in two regions. It continuously replicates the selected Unity
// Catalog catalogs and workspace assets from the primary region to the
// secondary region and lets you fail over to the secondary region; at least one
// of the two replication types must be enabled. See
// https://docs.databricks.com/aws/en/admin/managed-disaster-recovery for the
// product documentation.
type FailoverGroup struct {
	// Time at which this failover group was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Current effective primary region. Replication flows FROM workspaces in
	// this region. Changes after a successful failover.
	EffectivePrimaryRegion string `json:"effective_primary_region,omitempty"`
	// Opaque version string for optimistic locking. Server-generated and
	// returned in responses.
	Etag string `json:"etag,omitempty"`
	// True when replication errors are preventing the replication point from
	// advancing. Populated only while the failover group is INITIAL_REPLICATION
	// or ACTIVE; the blocking errors are listed in the
	// system.replication.states system table.
	HasReplicationPointBlockingErrors bool `json:"has_replication_point_blocking_errors,omitempty"`
	// Initial primary region. Used only in Create requests to set the starting
	// primary region, and must be one of the regions in `regions`. Not returned
	// in responses.
	InitialPrimaryRegion string `json:"initial_primary_region"`
	// Initial replication progress of all in-scope assets. Set only when the
	// FailoverGroup is in INITIAL_REPLICATION state.
	InitialReplicationProgress *InitialReplicationProgress `json:"initial_replication_progress,omitempty"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"name,omitempty"`
	// The two cloud regions that participate in this failover group, the
	// primary and the secondary. Each workspace set contains one workspace in
	// each region.
	Regions []string `json:"regions"`
	// The most recent point in time up to which all in-scope assets have been
	// replicated to the secondary region. Data written after this point might
	// not exist in the secondary yet and can be lost in an unplanned failover.
	// When both Unity Catalog and workspace assets are replicated, this is the
	// older of the two replication points. Not populated until the first
	// replication point exists, and not populated in ListFailoverGroups
	// responses.
	ReplicationPoint *time.Time `json:"replication_point,omitempty"`
	// Aggregate state of the failover group across Unity Catalog and workspace
	// asset replication. The group is INITIAL_REPLICATION until it has a
	// replication point, ACTIVE while replication is in steady state and
	// failover is available, and FAILING_OVER while a failover is in progress;
	// the `*_FAILED` values mean that the last operation did not complete. Not
	// populated in ListFailoverGroups responses.
	State FailoverGroupState `json:"state,omitempty"`
	// Unity Catalog replication configuration: the catalogs to replicate and
	// the storage location mappings for their data. Omit it to replicate
	// workspace assets only.
	UnityCatalogAssets *UcReplicationConfig `json:"unity_catalog_assets,omitempty"`
	// Time at which this failover group was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Workspace sets, each containing workspaces that replicate to each other.
	// A failover group currently supports exactly one workspace set.
	WorkspaceSets []WorkspaceSet `json:"workspace_sets"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailoverGroup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailoverGroup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The aggregate state of a FailoverGroup.
type FailoverGroupState string

const FailoverGroupStateActive FailoverGroupState = `ACTIVE`

const FailoverGroupStateCreating FailoverGroupState = `CREATING`

const FailoverGroupStateCreationFailed FailoverGroupState = `CREATION_FAILED`

const FailoverGroupStateDeleting FailoverGroupState = `DELETING`

const FailoverGroupStateDeletionFailed FailoverGroupState = `DELETION_FAILED`

const FailoverGroupStateFailingOver FailoverGroupState = `FAILING_OVER`

const FailoverGroupStateFailoverFailed FailoverGroupState = `FAILOVER_FAILED`

const FailoverGroupStateInitialReplication FailoverGroupState = `INITIAL_REPLICATION`

// String representation for [fmt.Print]
func (f *FailoverGroupState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FailoverGroupState) Set(v string) error {
	switch v {
	case `ACTIVE`, `CREATING`, `CREATION_FAILED`, `DELETING`, `DELETION_FAILED`, `FAILING_OVER`, `FAILOVER_FAILED`, `INITIAL_REPLICATION`:
		*f = FailoverGroupState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "CREATING", "CREATION_FAILED", "DELETING", "DELETION_FAILED", "FAILING_OVER", "FAILOVER_FAILED", "INITIAL_REPLICATION"`, v)
	}
}

// Values returns all possible values for FailoverGroupState.
//
// There is no guarantee on the order of the values in the slice.
func (f *FailoverGroupState) Values() []FailoverGroupState {
	return []FailoverGroupState{
		FailoverGroupStateActive,
		FailoverGroupStateCreating,
		FailoverGroupStateCreationFailed,
		FailoverGroupStateDeleting,
		FailoverGroupStateDeletionFailed,
		FailoverGroupStateFailingOver,
		FailoverGroupStateFailoverFailed,
		FailoverGroupStateInitialReplication,
	}
}

// Type always returns FailoverGroupState to satisfy [pflag.Value] interface
func (f *FailoverGroupState) Type() string {
	return "FailoverGroupState"
}

type GetFailoverGroupRequest struct {
	// The fully qualified resource name of the failover group. Format:
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
}

func (s *GetFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"-" url:"-"`
}

func (s *GetStableUrlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress of the initial replication of all in-scope assets (Unity Catalog and
// workspace assets) of a FailoverGroup. Populated only while the FailoverGroup
// is in INITIAL_REPLICATION.
type InitialReplicationProgress struct {
	// Percentage of in-scope assets replicated so far, from 0 to 100. Stays
	// below 100 while errors block the first replication point.
	Completed int `json:"completed,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InitialReplicationProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InitialReplicationProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFailoverGroupsRequest struct {
	// Maximum number of failover groups to return per page: - when set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - when set to 0 or unset, the page length is set
	// to a server configured value (recommended); - when set to a value less
	// than 0, an invalid parameter error is returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token received from a previous ListFailoverGroups call. Provide this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFailoverGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFailoverGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing failover groups.
type ListFailoverGroupsResponse struct {
	// The failover groups for this account.
	FailoverGroups []FailoverGroup `json:"failover_groups,omitempty"`
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFailoverGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFailoverGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStableUrlsRequest struct {
	// Maximum number of stable URLs to return per page: - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0 or unset, the page length is set to a
	// server configured value (recommended); - when set to a value less than 0,
	// an invalid parameter error is returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token received from a previous ListStableUrls call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStableUrlsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStableUrlsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing stable URLs.
type ListStableUrlsResponse struct {
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The stable URLs for this account.
	StableUrls []StableUrl `json:"stable_urls,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStableUrlsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStableUrlsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A location mapping identified by a name, with URIs per region. The system
// derives replication direction from effective_primary_region.
type LocationMapping struct {
	// Identifier of this location mapping, unique within the failover group.
	Name string `json:"name"`
	// The storage URI of this location in each region, with exactly one entry
	// per region of the failover group.
	UriByRegion []LocationMappingEntry `json:"uri_by_region"`
}

func (s *LocationMapping) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A single entry in a location mapping, mapping a region to a storage URI. Used
// instead of map<string, string> for proto2 compatibility.
type LocationMappingEntry struct {
	// Cloud region this URI applies to. Must be one of the failover group's
	// regions.
	Region string `json:"region"`
	// Cloud storage URI of the location in this region. End the URI with `/*`
	// to map every path under it as a prefix; within one mapping, either all
	// URIs or none carry the `/*` suffix.
	Uri string `json:"uri"`
}

func (s *LocationMappingEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A stable URL provides a failover-aware endpoint for accessing a workspace.
// Its lifecycle is independent of any failover group.
type StableUrl struct {
	// The workspace this stable URL currently routes to. Set to
	// `initial_workspace_id` at creation, advanced to the failover group's
	// primary while attached (including across a failover), and preserved when
	// the stable URL is detached from its failover group. Read this to see
	// where an unattached stable URL points: after a failover followed by a
	// detach it reflects the post-failover primary, not `initial_workspace_id`.
	EffectiveWorkspaceId string `json:"effective_workspace_id,omitempty"`
	// Fully qualified resource name of the FailoverGroup this stable URL is
	// currently linked to, in the format
	// `accounts/{account_id}/failover-groups/{failover_group_id}`. Empty when
	// the stable URL is not attached to any failover group.
	FailoverGroupName string `json:"failover_group_name,omitempty"`
	// The workspace this stable URL is initially bound to. Used only in Create
	// requests to associate the stable URL with a workspace. Not returned in
	// responses.
	InitialWorkspaceId string `json:"initial_workspace_id"`
	// Fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"name,omitempty"`
	// The stable workspace ID for this stable URL. Generated on creation and
	// immutable thereafter; identifies the URL across failovers and is the same
	// value embedded in the `url` (as the `w=` query parameter for SPOG URLs,
	// or in the `conn-<id>` hostname for Private-Link URLs).
	StableWorkspaceId string `json:"stable_workspace_id,omitempty"`
	// The stable URL endpoint. Generated on creation and immutable thereafter.
	// For non-Private-Link workspaces this is
	// `https://<spog_host>/?w=<connection_id>`. For Private-Link workspaces
	// this is the per-connection hostname.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StableUrl) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StableUrl) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog catalog to replicate.
type UcCatalog struct {
	// Name of the Unity Catalog catalog to replicate.
	Name string `json:"name"`
}

func (s *UcCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Unity Catalog replication configuration (top-level, not per-set).
type UcReplicationConfig struct {
	// Unity Catalog catalogs to replicate, including their metadata and managed
	// table data. Adding a catalog to an existing failover group returns the
	// group to INITIAL_REPLICATION until the catalog has been replicated.
	Catalogs []UcCatalog `json:"catalogs"`
	// Name of the workspace set whose workspaces run the data replication for
	// the catalogs' underlying storage. Must match the name of a workspace set
	// in this failover group.
	DataReplicationWorkspaceSet string `json:"data_replication_workspace_set"`
	// Storage location mappings. Each mapping pairs a storage path used by the
	// replicated catalogs in one region with the corresponding path in the
	// other region, so that replicated data is written to the matching storage
	// in the secondary region.
	LocationMappings []LocationMapping `json:"location_mappings,omitempty"`
}

func (s *UcReplicationConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateFailoverGroupRequest struct {
	// Optional opaque version string for optimistic locking, obtained from a
	// prior read of the failover group. If provided, the update is rejected
	// unless it matches the failover group's current etag. If omitted, the
	// update proceeds without an optimistic-lock check.
	Etag string `json:"-" url:"etag,omitempty"`
	// The failover group with updated fields. The name field identifies the
	// resource and is populated from the URL path.
	FailoverGroup FailoverGroup `json:"failover_group"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
	// Comma-separated list of the fields to update. The supported paths are
	// `workspace_sets`, `unity_catalog_assets`, `unity_catalog_assets.catalogs`
	// and `unity_catalog_assets.location_mappings`; any other path is rejected.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A set of workspaces that replicate to each other across regions.
type WorkspaceSet struct {
	// Identifier of the workspace set, unique within the failover group and
	// referenced by unity_catalog_assets.data_replication_workspace_set. Cannot
	// be changed after creation.
	Name string `json:"name"`
	// Whether to replicate workspace assets (notebooks, jobs, clusters, SQL
	// warehouses and so on) and their ACLs for this set. Defaults to false. Can
	// be turned on or off with UpdateFailoverGroup.
	ReplicateWorkspaceAssets bool `json:"replicate_workspace_assets,omitempty"`
	// Resource names of stable URLs associated with this workspace set. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}. The referenced stable
	// URLs must already exist (via CreateStableUrl), and the list can only be
	// set when replicate_workspace_assets is true. While workspace asset
	// replication stays enabled, an update can attach or detach one stable URL
	// at a time.
	StableUrlNames []string `json:"stable_url_names,omitempty"`
	// IDs of the two workspaces in this set, one in each of the failover
	// group's regions. The system derives and validates each workspace's
	// region. Both workspaces must belong to the account and have the Mission
	// Critical add-on enabled, and they cannot be changed after creation.
	WorkspaceIds []string `json:"workspace_ids"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceSet) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceSet) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
