// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The delta sharing authentication type.
type AuthenticationType string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

// String representation for [fmt.Print]
func (f *AuthenticationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `TOKEN`:
		*f = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "TOKEN"`, v)
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (f *AuthenticationType) Type() string {
	return "AuthenticationType"
}

type CentralCleanRoomInfo struct {
	// All assets from all collaborators that are available in the clean room.
	// Only one of table_info or notebook_info will be filled in.
	CleanRoomAssets []CleanRoomAssetInfo `tfsdk:"clean_room_assets"`
	// All collaborators who are in the clean room.
	Collaborators []CleanRoomCollaboratorInfo `tfsdk:"collaborators"`
	// The collaborator who created the clean room.
	Creator *CleanRoomCollaboratorInfo `tfsdk:"creator"`
	// The cloud where clean room tasks will be run.
	StationCloud types.String `tfsdk:"station_cloud"`
	// The region where clean room tasks will be run.
	StationRegion types.String `tfsdk:"station_region"`
}

type CleanRoomAssetInfo struct {
	// Time at which this asset was added, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// Details about the notebook asset.
	NotebookInfo *CleanRoomNotebookInfo `tfsdk:"notebook_info"`
	// The collaborator who owns the asset.
	Owner *CleanRoomCollaboratorInfo `tfsdk:"owner"`
	// Details about the table asset.
	TableInfo *CleanRoomTableInfo `tfsdk:"table_info"`
	// Time at which this asset was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

type CleanRoomCatalog struct {
	// Name of the catalog in the clean room station. Empty for notebooks.
	CatalogName types.String `tfsdk:"catalog_name"`
	// The details of the shared notebook files.
	NotebookFiles []SharedDataObject `tfsdk:"notebook_files"`
	// The details of the shared tables.
	Tables []SharedDataObject `tfsdk:"tables"`
}

type CleanRoomCatalogUpdate struct {
	// The name of the catalog to update assets.
	CatalogName types.String `tfsdk:"catalog_name"`
	// The updates to the assets in the catalog.
	Updates *SharedDataObjectUpdate `tfsdk:"updates"`
}

type CleanRoomCollaboratorInfo struct {
	// The global Unity Catalog metastore id of the collaborator. Also known as
	// the sharing identifier. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id"`
	// The organization name of the collaborator. This is configured in the
	// metastore for Delta Sharing and is used to identify the organization to
	// other collaborators.
	OrganizationName types.String `tfsdk:"organization_name"`
}

type CleanRoomInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this clean room was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of clean room creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Catalog aliases shared by the current collaborator with asset details.
	LocalCatalogs []CleanRoomCatalog `tfsdk:"local_catalogs"`
	// Name of the clean room.
	Name types.String `tfsdk:"name"`
	// Username of current owner of clean room.
	Owner types.String `tfsdk:"owner"`
	// Central clean room details.
	RemoteDetailedInfo *CentralCleanRoomInfo `tfsdk:"remote_detailed_info"`
	// Time at which this clean room was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of clean room updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

type CleanRoomNotebookInfo struct {
	// The base64 representation of the notebook content in HTML.
	NotebookContent types.String `tfsdk:"notebook_content"`
	// The name of the notebook.
	NotebookName types.String `tfsdk:"notebook_name"`
}

type CleanRoomTableInfo struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `tfsdk:"columns"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName types.String `tfsdk:"full_name"`
	// Name of table, relative to parent schema.
	Name types.String `tfsdk:"name"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name"`
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`

	Mask *ColumnMask `tfsdk:"mask"`
	// Name of Column.
	Name types.String `tfsdk:"name"`
	// Whether field may be Null (default: true).
	Nullable types.Bool `tfsdk:"nullable"`
	// Partition index for column.
	PartitionIndex types.Int64 `tfsdk:"partition_index"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// Full data type specification, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName ColumnTypeName `tfsdk:"type_name"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// Full data type specification as SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text"`
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName types.String `tfsdk:"function_name"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []types.String `tfsdk:"using_column_names"`
}

// Name of type (INT, STRUCT, MAP, etc.).
type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE"`, v)
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

type CreateCleanRoom struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the clean room.
	Name types.String `tfsdk:"name"`
	// Central clean room details.
	RemoteDetailedInfo CentralCleanRoomInfo `tfsdk:"remote_detailed_info"`
}

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code"`
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the share.
	Name types.String `tfsdk:"name"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// The name of the clean room.
	Name types.String `tfsdk:"-" url:"-"`
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-" url:"-"`
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-" url:"-"`
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-" url:"-"`
}

type GetActivationUrlInfoResponse struct {
}

// Get a clean room
type GetCleanRoomRequest struct {
	// Whether to include remote details (central) on the clean room.
	IncludeRemoteDetails types.Bool `tfsdk:"-" url:"include_remote_details,omitempty"`
	// The name of the clean room.
	Name types.String `tfsdk:"-" url:"-"`
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-" url:"-"`
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-" url:"-"`
}

type GetRecipientSharePermissionsResponse struct {
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `tfsdk:"permissions_out"`
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name types.String `tfsdk:"-" url:"-"`
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []types.String `tfsdk:"allowed_ip_addresses"`
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return. If not set, all the clean rooms
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListCleanRoomsResponse struct {
	// An array of clean rooms. Remote details (central) are not included.
	CleanRooms []CleanRoomInfo `tfsdk:"clean_rooms"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

type ListProviderSharesResponse struct {
	// An array of provider shares.
	Shares []ProviderShare `tfsdk:"shares"`
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId types.String `tfsdk:"-" url:"data_provider_global_metastore_id,omitempty"`
}

type ListProvidersResponse struct {
	// An array of provider information objects.
	Providers []ProviderInfo `tfsdk:"providers"`
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"-" url:"data_recipient_global_metastore_id,omitempty"`
}

type ListRecipientsResponse struct {
	// An array of recipient information objects.
	Recipients []RecipientInfo `tfsdk:"recipients"`
}

// List shares by Provider
type ListSharesRequest struct {
	// Name of the provider in which to list shares.
	Name types.String `tfsdk:"-" url:"-"`
}

type ListSharesResponse struct {
	// An array of data share information objects.
	Shares []ShareInfo `tfsdk:"shares"`
}

type Partition struct {
	// An array of partition values.
	Values []PartitionValue `tfsdk:"values"`
}

type PartitionValue struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name"`
	// The operator to apply for the value.
	Op PartitionValueOp `tfsdk:"op"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value"`
}

// The operator to apply for the value.
type PartitionValueOp string

const PartitionValueOpEqual PartitionValueOp = `EQUAL`

const PartitionValueOpLike PartitionValueOp = `LIKE`

// String representation for [fmt.Print]
func (f *PartitionValueOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PartitionValueOp) Set(v string) error {
	switch v {
	case `EQUAL`, `LIKE`:
		*f = PartitionValueOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "LIKE"`, v)
	}
}

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (f *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

type Privilege string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel Privilege = `CREATE_MODEL`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential Privilege = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeCreateVolume Privilege = `CREATE_VOLUME`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeReadVolume Privilege = `READ_VOLUME`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseConnection Privilege = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets Privilege = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume Privilege = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE`, `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `MANAGE_ALLOWLIST`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "CREATE", "CREATE_CATALOG", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "MANAGE_ALLOWLIST", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal"`
	// The privileges assigned to the principal.
	Privileges []Privilege `tfsdk:"privileges"`
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile *RecipientProfile `tfsdk:"recipient_profile"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name"`
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud types.String `tfsdk:"cloud"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `tfsdk:"tokens"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-" url:"-"`
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion"`
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds"`
	// The name of the recipient.
	Name types.String `tfsdk:"-" url:"-"`
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[types.String]types.String `tfsdk:"properties"`
}

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[types.String]types.String

type ShareInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Name of the share.
	Name types.String `tfsdk:"name"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `tfsdk:"objects"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// The name of the Recipient.
	Name types.String `tfsdk:"-" url:"-"`
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments"`
	// The share name.
	ShareName types.String `tfsdk:"share_name"`
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled types.Bool `tfsdk:"cdf_enabled"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment types.String `tfsdk:"comment"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content"`
	// The type of the data object.
	DataObjectType SharedDataObjectDataObjectType `tfsdk:"data_object_type"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `tfsdk:"history_data_sharing_status"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name types.String `tfsdk:"name"`
	// Array of partitions for the shared data.
	Partitions []Partition `tfsdk:"partitions"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs types.String `tfsdk:"shared_as"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion types.Int64 `tfsdk:"start_version"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus `tfsdk:"status"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as"`
}

// The type of the data object.
type SharedDataObjectDataObjectType string

const SharedDataObjectDataObjectTypeMaterializedView SharedDataObjectDataObjectType = `MATERIALIZED_VIEW`

const SharedDataObjectDataObjectTypeModel SharedDataObjectDataObjectType = `MODEL`

const SharedDataObjectDataObjectTypeNotebookFile SharedDataObjectDataObjectType = `NOTEBOOK_FILE`

const SharedDataObjectDataObjectTypeSchema SharedDataObjectDataObjectType = `SCHEMA`

const SharedDataObjectDataObjectTypeStreamingTable SharedDataObjectDataObjectType = `STREAMING_TABLE`

const SharedDataObjectDataObjectTypeTable SharedDataObjectDataObjectType = `TABLE`

const SharedDataObjectDataObjectTypeView SharedDataObjectDataObjectType = `VIEW`

// String representation for [fmt.Print]
func (f *SharedDataObjectDataObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectDataObjectType) Set(v string) error {
	switch v {
	case `MATERIALIZED_VIEW`, `MODEL`, `NOTEBOOK_FILE`, `SCHEMA`, `STREAMING_TABLE`, `TABLE`, `VIEW`:
		*f = SharedDataObjectDataObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MATERIALIZED_VIEW", "MODEL", "NOTEBOOK_FILE", "SCHEMA", "STREAMING_TABLE", "TABLE", "VIEW"`, v)
	}
}

// Type always returns SharedDataObjectDataObjectType to satisfy [pflag.Value] interface
func (f *SharedDataObjectDataObjectType) Type() string {
	return "SharedDataObjectDataObjectType"
}

// Whether to enable or disable sharing of data history. If not specified, the
// default is **DISABLED**.
type SharedDataObjectHistoryDataSharingStatus string

const SharedDataObjectHistoryDataSharingStatusDisabled SharedDataObjectHistoryDataSharingStatus = `DISABLED`

const SharedDataObjectHistoryDataSharingStatusEnabled SharedDataObjectHistoryDataSharingStatus = `ENABLED`

// String representation for [fmt.Print]
func (f *SharedDataObjectHistoryDataSharingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectHistoryDataSharingStatus) Set(v string) error {
	switch v {
	case `DISABLED`, `ENABLED`:
		*f = SharedDataObjectHistoryDataSharingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLED", "ENABLED"`, v)
	}
}

// Type always returns SharedDataObjectHistoryDataSharingStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectHistoryDataSharingStatus) Type() string {
	return "SharedDataObjectHistoryDataSharingStatus"
}

// One of: **ACTIVE**, **PERMISSION_DENIED**.
type SharedDataObjectStatus string

const SharedDataObjectStatusActive SharedDataObjectStatus = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatus = `PERMISSION_DENIED`

// String representation for [fmt.Print]
func (f *SharedDataObjectStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `PERMISSION_DENIED`:
		*f = SharedDataObjectStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PERMISSION_DENIED"`, v)
	}
}

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectStatus) Type() string {
	return "SharedDataObjectStatus"
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action SharedDataObjectUpdateAction `tfsdk:"action"`
	// The data object that is being added, removed, or updated.
	DataObject *SharedDataObject `tfsdk:"data_object"`
}

// One of: **ADD**, **REMOVE**, **UPDATE**.
type SharedDataObjectUpdateAction string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

const SharedDataObjectUpdateActionUpdate SharedDataObjectUpdateAction = `UPDATE`

// String representation for [fmt.Print]
func (f *SharedDataObjectUpdateAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectUpdateAction) Set(v string) error {
	switch v {
	case `ADD`, `REMOVE`, `UPDATE`:
		*f = SharedDataObjectUpdateAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD", "REMOVE", "UPDATE"`, v)
	}
}

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (f *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

type UpdateCleanRoom struct {
	// Array of shared data object updates.
	CatalogUpdates []CleanRoomCatalogUpdate `tfsdk:"catalog_updates"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the clean room.
	Name types.String `tfsdk:"-" url:"-"`
	// Username of current owner of clean room.
	Owner types.String `tfsdk:"owner"`
}

type UpdatePermissionsResponse struct {
}

type UpdateProvider struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// Name of the provider.
	Name types.String `tfsdk:"-" url:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Name of the recipient.
	Name types.String `tfsdk:"-" url:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`
}

type UpdateResponse struct {
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the share.
	Name types.String `tfsdk:"-" url:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `tfsdk:"updates"`
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes []catalog.PermissionsChange `tfsdk:"changes"`
	// The name of the share.
	Name types.String `tfsdk:"-" url:"-"`
}
