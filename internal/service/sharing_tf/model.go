// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sharing_tf

import (
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CentralCleanRoomInfo struct {
	// All assets from all collaborators that are available in the clean room.
	// Only one of table_info or notebook_info will be filled in.
	CleanRoomAssets []CleanRoomAssetInfo `tfsdk:"clean_room_assets" tf:"optional"`
	// All collaborators who are in the clean room.
	Collaborators []CleanRoomCollaboratorInfo `tfsdk:"collaborators" tf:"optional"`
	// The collaborator who created the clean room.
	Creator []CleanRoomCollaboratorInfo `tfsdk:"creator" tf:"optional"`
	// The cloud where clean room tasks will be run.
	StationCloud types.String `tfsdk:"station_cloud" tf:"optional"`
	// The region where clean room tasks will be run.
	StationRegion types.String `tfsdk:"station_region" tf:"optional"`
}

type CleanRoomAssetInfo struct {
	// Time at which this asset was added, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"optional"`
	// Details about the notebook asset.
	NotebookInfo []CleanRoomNotebookInfo `tfsdk:"notebook_info" tf:"optional"`
	// The collaborator who owns the asset.
	Owner []CleanRoomCollaboratorInfo `tfsdk:"owner" tf:"optional"`
	// Details about the table asset.
	TableInfo []CleanRoomTableInfo `tfsdk:"table_info" tf:"optional"`
	// Time at which this asset was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

type CleanRoomCatalog struct {
	// Name of the catalog in the clean room station. Empty for notebooks.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The details of the shared notebook files.
	NotebookFiles []SharedDataObject `tfsdk:"notebook_files" tf:"optional"`
	// The details of the shared tables.
	Tables []SharedDataObject `tfsdk:"tables" tf:"optional"`
}

type CleanRoomCatalogUpdate struct {
	// The name of the catalog to update assets.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The updates to the assets in the catalog.
	Updates []SharedDataObjectUpdate `tfsdk:"updates" tf:"optional"`
}

type CleanRoomCollaboratorInfo struct {
	// The global Unity Catalog metastore id of the collaborator. Also known as
	// the sharing identifier. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// The organization name of the collaborator. This is configured in the
	// metastore for Delta Sharing and is used to identify the organization to
	// other collaborators.
	OrganizationName types.String `tfsdk:"organization_name" tf:"optional"`
}

type CleanRoomInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this clean room was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of clean room creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Catalog aliases shared by the current collaborator with asset details.
	LocalCatalogs []CleanRoomCatalog `tfsdk:"local_catalogs" tf:"optional"`
	// Name of the clean room.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of clean room.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Central clean room details.
	RemoteDetailedInfo []CentralCleanRoomInfo `tfsdk:"remote_detailed_info" tf:"optional"`
	// Time at which this clean room was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of clean room updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type CleanRoomNotebookInfo struct {
	// The base64 representation of the notebook content in HTML.
	NotebookContent types.String `tfsdk:"notebook_content" tf:"optional"`
	// The name of the notebook.
	NotebookName types.String `tfsdk:"notebook_name" tf:"optional"`
}

type CleanRoomTableInfo struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `tfsdk:"columns" tf:"optional"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Name of table, relative to parent schema.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`

	Mask []ColumnMask `tfsdk:"mask" tf:"optional"`
	// Name of Column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether field may be Null (default: true).
	Nullable types.Bool `tfsdk:"nullable" tf:"optional"`
	// Partition index for column.
	PartitionIndex types.Int64 `tfsdk:"partition_index" tf:"optional"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// Full data type specification, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json" tf:"optional"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName types.String `tfsdk:"type_name" tf:"optional"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type specification as SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:"optional"`
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:"optional"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []types.String `tfsdk:"using_column_names" tf:"optional"`
}

type CreateCleanRoom struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the clean room.
	Name types.String `tfsdk:"name" tf:""`
	// Central clean room details.
	RemoteDetailedInfo []CentralCleanRoomInfo `tfsdk:"remote_detailed_info" tf:""`
}

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:""`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:""`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:""`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// The name of the clean room.
	Name types.String `tfsdk:"-"`
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

type DeleteResponse struct {
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

type GetActivationUrlInfoResponse struct {
}

// Get a clean room
type GetCleanRoomRequest struct {
	// Whether to include remote details (central) on the clean room.
	IncludeRemoteDetails types.Bool `tfsdk:"-"`
	// The name of the clean room.
	Name types.String `tfsdk:"-"`
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `tfsdk:"permissions_out" tf:"optional"`
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []types.String `tfsdk:"allowed_ip_addresses" tf:"optional"`
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return. If not set, all the clean rooms
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListCleanRoomsResponse struct {
	// An array of clean rooms. Remote details (central) are not included.
	CleanRooms []CleanRoomInfo `tfsdk:"clean_rooms" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider shares.
	Shares []ProviderShare `tfsdk:"shares" tf:"optional"`
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider information objects.
	Providers []ProviderInfo `tfsdk:"providers" tf:"optional"`
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of recipient information objects.
	Recipients []RecipientInfo `tfsdk:"recipients" tf:"optional"`
}

// List shares by Provider
type ListSharesRequest struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the provider in which to list shares.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share information objects.
	Shares []ShareInfo `tfsdk:"shares" tf:"optional"`
}

type Partition struct {
	// An array of partition values.
	Values []PartitionValue `tfsdk:"values" tf:"optional"`
}

type PartitionValue struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op" tf:"optional"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key" tf:"optional"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value" tf:"optional"`
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges []types.String `tfsdk:"privileges" tf:"optional"`
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id" tf:"optional"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile []RecipientProfile `tfsdk:"recipient_profile" tf:"optional"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name" tf:"optional"`
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated" tf:"optional"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `tfsdk:"tokens" tf:"optional"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version" tf:"optional"`
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds" tf:""`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:""`
}

type ShareInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `tfsdk:"objects" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the Recipient.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments" tf:"optional"`
	// The share name.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"optional"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by" tf:"optional"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled types.Bool `tfsdk:"cdf_enabled" tf:"optional"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The type of the data object.
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus types.String `tfsdk:"history_data_sharing_status" tf:"optional"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name types.String `tfsdk:"name" tf:""`
	// Array of partitions for the shared data.
	Partitions []Partition `tfsdk:"partitions" tf:"optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs types.String `tfsdk:"shared_as" tf:"optional"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion types.Int64 `tfsdk:"start_version" tf:"optional"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status" tf:"optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as" tf:"optional"`
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action" tf:"optional"`
	// The data object that is being added, removed, or updated.
	DataObject []SharedDataObject `tfsdk:"data_object" tf:"optional"`
}

type UpdateCleanRoom struct {
	// Array of shared data object updates.
	CatalogUpdates []CleanRoomCatalogUpdate `tfsdk:"catalog_updates" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the clean room.
	Name types.String `tfsdk:"-"`
	// Username of current owner of clean room.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

type UpdatePermissionsResponse struct {
}

type UpdateProvider struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList []IpAccessList `tfsdk:"ip_access_list" tf:"optional"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs []SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs" tf:"optional"`
}

type UpdateResponse struct {
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `tfsdk:"updates" tf:"optional"`
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes catalog.PermissionsChange `tfsdk:"changes" tf:"optional"`
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}
