// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package cleanrooms_tf

import (
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	AccessRestricted types.String `tfsdk:"access_restricted" tf:"optional"`

	Comment types.String `tfsdk:"comment" tf:"optional"`
	// When the clean room was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The alias of the collaborator tied to the local clean room.
	LocalCollaboratorAlias types.String `tfsdk:"local_collaborator_alias" tf:"optional"`
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	Name types.String `tfsdk:"name" tf:"optional"`
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	OutputCatalog []CleanRoomOutputCatalog `tfsdk:"output_catalog" tf:"optional,object"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo []CleanRoomRemoteDetail `tfsdk:"remote_detailed_info" tf:"optional,object"`
	// Clean room status.
	Status types.String `tfsdk:"status" tf:"optional"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

func (newState *CleanRoom) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoom) {
}

func (newState *CleanRoom) SyncEffectiveFieldsDuringRead(existingState CleanRoom) {
}

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"optional"`
	// The type of the asset.
	AssetType types.String `tfsdk:"asset_type" tf:"optional"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable []CleanRoomAssetForeignTable `tfsdk:"foreign_table" tf:"optional,object"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails []CleanRoomAssetForeignTableLocalDetails `tfsdk:"foreign_table_local_details" tf:"optional,object"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook []CleanRoomAssetNotebook `tfsdk:"notebook" tf:"optional,object"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias types.String `tfsdk:"owner_collaborator_alias" tf:"optional"`
	// Status of the asset
	Status types.String `tfsdk:"status" tf:"optional"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table []CleanRoomAssetTable `tfsdk:"table" tf:"optional,object"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails []CleanRoomAssetTableLocalDetails `tfsdk:"table_local_details" tf:"optional,object"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View []CleanRoomAssetView `tfsdk:"view" tf:"optional,object"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails []CleanRoomAssetViewLocalDetails `tfsdk:"view_local_details" tf:"optional,object"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails []CleanRoomAssetVolumeLocalDetails `tfsdk:"volume_local_details" tf:"optional,object"`
}

func (newState *CleanRoomAsset) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAsset) {
}

func (newState *CleanRoomAsset) SyncEffectiveFieldsDuringRead(existingState CleanRoomAsset) {
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	Columns catalog.ColumnInfo `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetForeignTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetForeignTable) {
}

func (newState *CleanRoomAssetForeignTable) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetForeignTable) {
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetForeignTableLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetForeignTableLocalDetails) {
}

func (newState *CleanRoomAssetForeignTableLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetForeignTableLocalDetails) {
}

type CleanRoomAssetNotebook struct {
	// Server generated checksum that represents the notebook version.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent types.String `tfsdk:"notebook_content" tf:"optional"`
}

func (newState *CleanRoomAssetNotebook) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetNotebook) {
}

func (newState *CleanRoomAssetNotebook) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetNotebook) {
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	Columns catalog.ColumnInfo `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetTable) {
}

func (newState *CleanRoomAssetTable) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetTable) {
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
	// Partition filtering specification for a shared table.
	Partitions sharing.PartitionSpecificationPartition `tfsdk:"partitions" tf:"optional"`
}

func (newState *CleanRoomAssetTableLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetTableLocalDetails) {
}

func (newState *CleanRoomAssetTableLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetTableLocalDetails) {
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	Columns catalog.ColumnInfo `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetView) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetView) {
}

func (newState *CleanRoomAssetView) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetView) {
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetViewLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetViewLocalDetails) {
}

func (newState *CleanRoomAssetViewLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetViewLocalDetails) {
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetVolumeLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetVolumeLocalDetails) {
}

func (newState *CleanRoomAssetVolumeLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetVolumeLocalDetails) {
}

// Publicly visible clean room collaborator.
type CleanRoomCollaborator struct {
	// Collaborator alias specified by the clean room creator. It is unique
	// across all collaborators of this clean room, and used to derive multiple
	// values internally such as catalog alias and clean room name for single
	// metastore clean rooms. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CollaboratorAlias types.String `tfsdk:"collaborator_alias" tf:"optional"`
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	InviteRecipientEmail types.String `tfsdk:"invite_recipient_email" tf:"optional"`
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	InviteRecipientWorkspaceId types.Int64 `tfsdk:"invite_recipient_workspace_id" tf:"optional"`
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	OrganizationName types.String `tfsdk:"organization_name" tf:"optional"`
}

func (newState *CleanRoomCollaborator) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomCollaborator) {
}

func (newState *CleanRoomCollaborator) SyncEffectiveFieldsDuringRead(existingState CleanRoomCollaborator) {
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo []CollaboratorJobRunInfo `tfsdk:"collaborator_job_run_info" tf:"optional,object"`
	// State of the task run.
	NotebookJobRunState jobs.CleanRoomTaskRunState `tfsdk:"notebook_job_run_state" tf:"optional,object"`
	// Asset name of the notebook executed in this task run.
	NotebookName types.String `tfsdk:"notebook_name" tf:"optional"`
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime types.Int64 `tfsdk:"output_schema_expiration_time" tf:"optional"`
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:"optional"`
	// Duration of the task run, in milliseconds.
	RunDuration types.Int64 `tfsdk:"run_duration" tf:"optional"`
	// When the task run started, in epoch milliseconds.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
}

func (newState *CleanRoomNotebookTaskRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomNotebookTaskRun) {
}

func (newState *CleanRoomNotebookTaskRun) SyncEffectiveFieldsDuringRead(existingState CleanRoomNotebookTaskRun) {
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *CleanRoomOutputCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomOutputCatalog) {
}

func (newState *CleanRoomOutputCatalog) SyncEffectiveFieldsDuringRead(existingState CleanRoomOutputCatalog) {
}

// Publicly visible central clean room details.
type CleanRoomRemoteDetail struct {
	// Central clean room ID.
	CentralCleanRoomId types.String `tfsdk:"central_clean_room_id" tf:"optional"`
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	CloudVendor types.String `tfsdk:"cloud_vendor" tf:"optional"`
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	Collaborators []CleanRoomCollaborator `tfsdk:"collaborators" tf:"optional"`
	// The compliance security profile used to process regulated data following
	// compliance standards.
	ComplianceSecurityProfile []ComplianceSecurityProfile `tfsdk:"compliance_security_profile" tf:"optional,object"`
	// Collaborator who creates the clean room.
	Creator []CleanRoomCollaborator `tfsdk:"creator" tf:"optional,object"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy settings.EgressNetworkPolicy `tfsdk:"egress_network_policy" tf:"optional,object"`
	// Region of the central clean room.
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *CleanRoomRemoteDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomRemoteDetail) {
}

func (newState *CleanRoomRemoteDetail) SyncEffectiveFieldsDuringRead(existingState CleanRoomRemoteDetail) {
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	CollaboratorAlias types.String `tfsdk:"collaborator_alias" tf:"optional"`
	// Job ID of the task run in the collaborator's workspace.
	CollaboratorJobId types.Int64 `tfsdk:"collaborator_job_id" tf:"optional"`
	// Job run ID of the task run in the collaborator's workspace.
	CollaboratorJobRunId types.Int64 `tfsdk:"collaborator_job_run_id" tf:"optional"`
	// Task run ID of the task run in the collaborator's workspace.
	CollaboratorTaskRunId types.Int64 `tfsdk:"collaborator_task_run_id" tf:"optional"`
	// ID of the collaborator's workspace that triggered the task run.
	CollaboratorWorkspaceId types.Int64 `tfsdk:"collaborator_workspace_id" tf:"optional"`
}

func (newState *CollaboratorJobRunInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CollaboratorJobRunInfo) {
}

func (newState *CollaboratorJobRunInfo) SyncEffectiveFieldsDuringRead(existingState CollaboratorJobRunInfo) {
}

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards settings.ComplianceStandard `tfsdk:"compliance_standards" tf:"optional"`
	// Whether the compliance security profile is enabled.
	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfile) {
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfile) {
}

// Create an asset
type CreateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset []CleanRoomAsset `tfsdk:"asset" tf:"optional,object"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *CreateCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomAssetRequest) {
}

func (newState *CreateCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomAssetRequest) {
}

// Create an output catalog
type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog []CleanRoomOutputCatalog `tfsdk:"output_catalog" tf:"optional,object"`
}

func (newState *CreateCleanRoomOutputCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomOutputCatalogRequest) {
}

func (newState *CreateCleanRoomOutputCatalogRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomOutputCatalogRequest) {
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog []CleanRoomOutputCatalog `tfsdk:"output_catalog" tf:"optional,object"`
}

func (newState *CreateCleanRoomOutputCatalogResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomOutputCatalogResponse) {
}

func (newState *CreateCleanRoomOutputCatalogResponse) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomOutputCatalogResponse) {
}

// Create a clean room
type CreateCleanRoomRequest struct {
	CleanRoom []CleanRoom `tfsdk:"clean_room" tf:"optional,object"`
}

func (newState *CreateCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomRequest) {
}

func (newState *CreateCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomRequest) {
}

// Delete an asset
type DeleteCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName types.String `tfsdk:"-"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *DeleteCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomAssetRequest) {
}

func (newState *DeleteCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomAssetRequest) {
}

// Response for delete clean room request. Using an empty message since the
// generic Empty proto does not externd UnshadedMessageMarker.
type DeleteCleanRoomAssetResponse struct {
}

func (newState *DeleteCleanRoomAssetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomAssetResponse) {
}

func (newState *DeleteCleanRoomAssetResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomAssetResponse) {
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomRequest) {
}

func (newState *DeleteCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// Get an asset
type GetCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName types.String `tfsdk:"-"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *GetCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCleanRoomAssetRequest) {
}

func (newState *GetCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState GetCleanRoomAssetRequest) {
}

// Get a clean room
type GetCleanRoomRequest struct {
	Name types.String `tfsdk:"-"`
}

func (newState *GetCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCleanRoomRequest) {
}

func (newState *GetCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState GetCleanRoomRequest) {
}

// List assets
type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomAssetsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomAssetsRequest) {
}

func (newState *ListCleanRoomAssetsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomAssetsRequest) {
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	Assets []CleanRoomAsset `tfsdk:"assets" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCleanRoomAssetsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomAssetsResponse) {
}

func (newState *ListCleanRoomAssetsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomAssetsResponse) {
}

// List notebook task runs
type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Notebook name
	NotebookName types.String `tfsdk:"-"`
	// The maximum number of task runs to return
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomNotebookTaskRunsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomNotebookTaskRunsRequest) {
}

func (newState *ListCleanRoomNotebookTaskRunsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomNotebookTaskRunsRequest) {
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Name of the clean room.
	Runs []CleanRoomNotebookTaskRun `tfsdk:"runs" tf:"optional"`
}

func (newState *ListCleanRoomNotebookTaskRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomNotebookTaskRunsResponse) {
}

func (newState *ListCleanRoomNotebookTaskRunsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomNotebookTaskRunsResponse) {
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomsRequest) {
}

func (newState *ListCleanRoomsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomsRequest) {
}

type ListCleanRoomsResponse struct {
	CleanRooms []CleanRoom `tfsdk:"clean_rooms" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCleanRoomsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomsResponse) {
}

func (newState *ListCleanRoomsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomsResponse) {
}

// Update an asset
type UpdateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset []CleanRoomAsset `tfsdk:"asset" tf:"optional,object"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCleanRoomAssetRequest) {
}

func (newState *UpdateCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCleanRoomAssetRequest) {
}

type UpdateCleanRoomRequest struct {
	CleanRoom []CleanRoom `tfsdk:"clean_room" tf:"optional,object"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCleanRoomRequest) {
}

func (newState *UpdateCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCleanRoomRequest) {
}
