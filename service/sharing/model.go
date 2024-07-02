// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
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
	StationCloud string `tfsdk:"station_cloud"`
	// The region where clean room tasks will be run.
	StationRegion string `tfsdk:"station_region"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CentralCleanRoomInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CentralCleanRoomInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomAssetInfo struct {
	// Time at which this asset was added, in epoch milliseconds.
	AddedAt int64 `tfsdk:"added_at"`
	// Details about the notebook asset.
	NotebookInfo *CleanRoomNotebookInfo `tfsdk:"notebook_info"`
	// The collaborator who owns the asset.
	Owner *CleanRoomCollaboratorInfo `tfsdk:"owner"`
	// Details about the table asset.
	TableInfo *CleanRoomTableInfo `tfsdk:"table_info"`
	// Time at which this asset was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomAssetInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomAssetInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomCatalog struct {
	// Name of the catalog in the clean room station. Empty for notebooks.
	CatalogName string `tfsdk:"catalog_name"`
	// The details of the shared notebook files.
	NotebookFiles []SharedDataObject `tfsdk:"notebook_files"`
	// The details of the shared tables.
	Tables []SharedDataObject `tfsdk:"tables"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomCatalogUpdate struct {
	// The name of the catalog to update assets.
	CatalogName string `tfsdk:"catalog_name"`
	// The updates to the assets in the catalog.
	Updates *SharedDataObjectUpdate `tfsdk:"updates"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomCatalogUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomCatalogUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomCollaboratorInfo struct {
	// The global Unity Catalog metastore id of the collaborator. Also known as
	// the sharing identifier. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	GlobalMetastoreId string `tfsdk:"global_metastore_id"`
	// The organization name of the collaborator. This is configured in the
	// metastore for Delta Sharing and is used to identify the organization to
	// other collaborators.
	OrganizationName string `tfsdk:"organization_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomCollaboratorInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomCollaboratorInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomInfo struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this clean room was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of clean room creator.
	CreatedBy string `tfsdk:"created_by"`
	// Catalog aliases shared by the current collaborator with asset details.
	LocalCatalogs []CleanRoomCatalog `tfsdk:"local_catalogs"`
	// Name of the clean room.
	Name string `tfsdk:"name"`
	// Username of current owner of clean room.
	Owner string `tfsdk:"owner"`
	// Central clean room details.
	RemoteDetailedInfo *CentralCleanRoomInfo `tfsdk:"remote_detailed_info"`
	// Time at which this clean room was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of clean room updater.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomNotebookInfo struct {
	// The base64 representation of the notebook content in HTML.
	NotebookContent string `tfsdk:"notebook_content"`
	// The name of the notebook.
	NotebookName string `tfsdk:"notebook_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomNotebookInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomNotebookInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomTableInfo struct {
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `tfsdk:"columns"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `tfsdk:"full_name"`
	// Name of table, relative to parent schema.
	Name string `tfsdk:"name"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `tfsdk:"schema_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CleanRoomTableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomTableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`

	Mask *ColumnMask `tfsdk:"mask"`
	// Name of Column.
	Name string `tfsdk:"name"`
	// Whether field may be Null (default: true).
	Nullable bool `tfsdk:"nullable"`
	// Partition index for column.
	PartitionIndex int `tfsdk:"partition_index"`
	// Ordinal position of column (starting at position 0).
	Position int `tfsdk:"position"`
	// Format of IntervalType.
	TypeIntervalType string `tfsdk:"type_interval_type"`
	// Full data type specification, JSON-serialized.
	TypeJson string `tfsdk:"type_json"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName ColumnTypeName `tfsdk:"type_name"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision int `tfsdk:"type_precision"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int `tfsdk:"type_scale"`
	// Full data type specification as SQL/catalogString text.
	TypeText string `tfsdk:"type_text"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName string `tfsdk:"function_name"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string `tfsdk:"using_column_names"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ColumnMask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnMask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Comment string `tfsdk:"comment"`
	// Name of the clean room.
	Name string `tfsdk:"name"`
	// Central clean room details.
	RemoteDetailedInfo CentralCleanRoomInfo `tfsdk:"remote_detailed_info"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateCleanRoom) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCleanRoom) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Description about the provider.
	Comment string `tfsdk:"comment"`
	// The name of the Provider.
	Name string `tfsdk:"name"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr string `tfsdk:"recipient_profile_str"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Description about the recipient.
	Comment string `tfsdk:"comment"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `tfsdk:"data_recipient_global_metastore_id"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Name of Recipient.
	Name string `tfsdk:"name"`
	// Username of the recipient owner.
	Owner string `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `tfsdk:"sharing_code"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Name of the share.
	Name string `tfsdk:"name"`
	// Storage root URL for the share.
	StorageRoot string `tfsdk:"storage_root"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// The name of the clean room.
	Name string `tfsdk:"-" url:"-"`
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name string `tfsdk:"-" url:"-"`
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name string `tfsdk:"-" url:"-"`
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `tfsdk:"-" url:"-"`
}

type GetActivationUrlInfoResponse struct {
}

// Get a clean room
type GetCleanRoomRequest struct {
	// Whether to include remote details (central) on the clean room.
	IncludeRemoteDetails bool `tfsdk:"-" url:"include_remote_details,omitempty"`
	// The name of the clean room.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetCleanRoomRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCleanRoomRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name string `tfsdk:"-" url:"-"`
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name string `tfsdk:"-" url:"-"`
}

type GetRecipientSharePermissionsResponse struct {
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment `tfsdk:"permissions_out"`
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData bool `tfsdk:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetShareRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetShareRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses []string `tfsdk:"allowed_ip_addresses"`
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return. If not set, all the clean rooms
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListCleanRoomsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCleanRoomsResponse struct {
	// An array of clean rooms. Remote details (central) are not included.
	CleanRooms []CleanRoomInfo `tfsdk:"clean_rooms"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListCleanRoomsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCleanRoomsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProviderSharesResponse struct {
	// An array of provider shares.
	Shares []ProviderShare `tfsdk:"shares"`
}

// List providers
type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId string `tfsdk:"-" url:"data_provider_global_metastore_id,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersResponse struct {
	// An array of provider information objects.
	Providers []ProviderInfo `tfsdk:"providers"`
}

// List share recipients
type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId string `tfsdk:"-" url:"data_recipient_global_metastore_id,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListRecipientsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRecipientsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRecipientsResponse struct {
	// An array of recipient information objects.
	Recipients []RecipientInfo `tfsdk:"recipients"`
}

// List shares by Provider
type ListSharesRequest struct {
	// Name of the provider in which to list shares.
	Name string `tfsdk:"-" url:"-"`
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
	Name string `tfsdk:"name"`
	// The operator to apply for the value.
	Op PartitionValueOp `tfsdk:"op"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey string `tfsdk:"recipient_property_key"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value string `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PartitionValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

const PrivilegeSingleUserAccess Privilege = `SINGLE_USER_ACCESS`

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
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE`, `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `MANAGE_ALLOWLIST`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `SINGLE_USER_ACCESS`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "CREATE", "CREATE_CATALOG", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "MANAGE_ALLOWLIST", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "SINGLE_USER_ACCESS", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `tfsdk:"principal"`
	// The privileges assigned to the principal.
	Privileges []Privilege `tfsdk:"privileges"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud string `tfsdk:"cloud"`
	// Description about the provider.
	Comment string `tfsdk:"comment"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of Provider creator.
	CreatedBy string `tfsdk:"created_by"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId string `tfsdk:"data_provider_global_metastore_id"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId string `tfsdk:"metastore_id"`
	// The name of the Provider.
	Name string `tfsdk:"name"`
	// Username of Provider owner.
	Owner string `tfsdk:"owner"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile *RecipientProfile `tfsdk:"recipient_profile"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr string `tfsdk:"recipient_profile_str"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region string `tfsdk:"region"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified Share.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ProviderShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated bool `tfsdk:"activated"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `tfsdk:"activation_url"`
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `tfsdk:"authentication_type"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud string `tfsdk:"cloud"`
	// Description about the recipient.
	Comment string `tfsdk:"comment"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of recipient creator.
	CreatedBy string `tfsdk:"created_by"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `tfsdk:"data_recipient_global_metastore_id"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of Recipient.
	Name string `tfsdk:"name"`
	// Username of the recipient owner.
	Owner string `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region string `tfsdk:"region"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `tfsdk:"sharing_code"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo `tfsdk:"tokens"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of recipient updater.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RecipientInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken string `tfsdk:"bearer_token"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `tfsdk:"endpoint"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion int `tfsdk:"share_credentials_version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RecipientProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `tfsdk:"activation_url"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of recipient token creator.
	CreatedBy string `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `tfsdk:"expiration_time"`
	// Unique ID of the recipient token.
	Id string `tfsdk:"id"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of recipient Token updater.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RecipientTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `tfsdk:"-" url:"-"`
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken string `tfsdk:"bearerToken"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `tfsdk:"endpoint"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime string `tfsdk:"expirationTime"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int `tfsdk:"shareCredentialsVersion"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RetrieveTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `tfsdk:"existing_token_expire_in_seconds"`
	// The name of the recipient.
	Name string `tfsdk:"-" url:"-"`
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`
}

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[string]string

type ShareInfo struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of share creator.
	CreatedBy string `tfsdk:"created_by"`
	// Name of the share.
	Name string `tfsdk:"name"`
	// A list of shared data objects within the share.
	Objects []SharedDataObject `tfsdk:"objects"`
	// Username of current owner of share.
	Owner string `tfsdk:"owner"`
	// Storage Location URL (full path) for the share.
	StorageLocation string `tfsdk:"storage_location"`
	// Storage root URL for the share.
	StorageRoot string `tfsdk:"storage_root"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of share updater.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ShareInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get recipient share permissions
type SharePermissionsRequest struct {
	// The name of the Recipient.
	Name string `tfsdk:"-" url:"-"`
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments"`
	// The share name.
	ShareName string `tfsdk:"share_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ShareToPrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareToPrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt int64 `tfsdk:"added_at"`
	// Username of the sharer.
	AddedBy string `tfsdk:"added_by"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled bool `tfsdk:"cdf_enabled"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment string `tfsdk:"comment"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content string `tfsdk:"content"`
	// The type of the data object.
	DataObjectType SharedDataObjectDataObjectType `tfsdk:"data_object_type"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `tfsdk:"history_data_sharing_status"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name string `tfsdk:"name"`
	// Array of partitions for the shared data.
	Partitions []Partition `tfsdk:"partitions"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs string `tfsdk:"shared_as"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion int64 `tfsdk:"start_version"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus `tfsdk:"status"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs string `tfsdk:"string_shared_as"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Comment string `tfsdk:"comment"`
	// The name of the clean room.
	Name string `tfsdk:"-" url:"-"`
	// Username of current owner of clean room.
	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateCleanRoom) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCleanRoom) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePermissionsResponse struct {
}

type UpdateProvider struct {
	// Description about the provider.
	Comment string `tfsdk:"comment"`
	// Name of the provider.
	Name string `tfsdk:"-" url:"-"`
	// New name for the provider.
	NewName string `tfsdk:"new_name"`
	// Username of Provider owner.
	Owner string `tfsdk:"owner"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr string `tfsdk:"recipient_profile_str"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRecipient struct {
	// Description about the recipient.
	Comment string `tfsdk:"comment"`
	// IP Access List
	IpAccessList *IpAccessList `tfsdk:"ip_access_list"`
	// Name of the recipient.
	Name string `tfsdk:"-" url:"-"`
	// New name for the recipient.
	NewName string `tfsdk:"new_name"`
	// Username of the recipient owner.
	Owner string `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs `tfsdk:"properties_kvpairs"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateResponse struct {
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// The name of the share.
	Name string `tfsdk:"-" url:"-"`
	// New name for the share.
	NewName string `tfsdk:"new_name"`
	// Username of current owner of share.
	Owner string `tfsdk:"owner"`
	// Storage root URL for the share.
	StorageRoot string `tfsdk:"storage_root"`
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate `tfsdk:"updates"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes []catalog.PermissionsChange `tfsdk:"changes"`
	// The name of the share.
	Name string `tfsdk:"-" url:"-"`
}
