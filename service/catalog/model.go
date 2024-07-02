// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountsCreateMetastore struct {
	MetastoreInfo *CreateMetastore `tfsdk:"metastore_info"`
}

type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment *CreateMetastoreAssignment `tfsdk:"metastore_assignment"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type AccountsCreateStorageCredential struct {
	CredentialInfo *CreateStorageCredential `tfsdk:"credential_info"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
}

type AccountsMetastoreAssignment struct {
	MetastoreAssignment *MetastoreAssignment `tfsdk:"metastore_assignment"`
}

type AccountsMetastoreInfo struct {
	MetastoreInfo *MetastoreInfo `tfsdk:"metastore_info"`
}

type AccountsStorageCredentialInfo struct {
	CredentialInfo *StorageCredentialInfo `tfsdk:"credential_info"`
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`

	MetastoreInfo *UpdateMetastore `tfsdk:"metastore_info"`
}

type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment *UpdateMetastoreAssignment `tfsdk:"metastore_assignment"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo *UpdateStorageCredential `tfsdk:"credential_info"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `tfsdk:"-" url:"-"`
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `tfsdk:"artifact_matchers"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `tfsdk:"created_by"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ArtifactAllowlistInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ArtifactAllowlistInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact string `tfsdk:"artifact"`
	// The pattern matching type of the artifact
	MatchType MatchType `tfsdk:"match_type"`
}

// The artifact type
type ArtifactType string

const ArtifactTypeInitScript ArtifactType = `INIT_SCRIPT`

const ArtifactTypeLibraryJar ArtifactType = `LIBRARY_JAR`

const ArtifactTypeLibraryMaven ArtifactType = `LIBRARY_MAVEN`

// String representation for [fmt.Print]
func (f *ArtifactType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ArtifactType) Set(v string) error {
	switch v {
	case `INIT_SCRIPT`, `LIBRARY_JAR`, `LIBRARY_MAVEN`:
		*f = ArtifactType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INIT_SCRIPT", "LIBRARY_JAR", "LIBRARY_MAVEN"`, v)
	}
}

// Type always returns ArtifactType to satisfy [pflag.Value] interface
func (f *ArtifactType) Type() string {
	return "ArtifactType"
}

type AssignResponse struct {
}

type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `tfsdk:"role_arn"`
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string `tfsdk:"external_id"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `tfsdk:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `tfsdk:"unity_catalog_iam_arn"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AwsIamRoleResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRoleResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `tfsdk:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `tfsdk:"managed_identity_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AzureManagedIdentityRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `tfsdk:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `tfsdk:"credential_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `tfsdk:"managed_identity_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AzureManagedIdentityResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId string `tfsdk:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret string `tfsdk:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string `tfsdk:"directory_id"`
}

// Cancel refresh
type CancelRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `tfsdk:"-" url:"-"`
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

type CancelRefreshResponse struct {
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// The type of the catalog.
	CatalogType CatalogType `tfsdk:"catalog_type"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// The name of the connection to an external data source.
	ConnectionName string `tfsdk:"connection_name"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of catalog creator.
	CreatedBy string `tfsdk:"created_by"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `tfsdk:"enable_predictive_optimization"`
	// The full name of the catalog. Corresponds with the name field.
	FullName string `tfsdk:"full_name"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `tfsdk:"isolation_mode"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of catalog.
	Name string `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `tfsdk:"options"`
	// Username of current owner of catalog.
	Owner string `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `tfsdk:"provider_name"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo `tfsdk:"provisioning_info"`
	// Kind of catalog securable.
	SecurableKind CatalogInfoSecurableKind `tfsdk:"securable_kind"`

	SecurableType string `tfsdk:"securable_type"`
	// The name of the share under the share provider.
	ShareName string `tfsdk:"share_name"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation string `tfsdk:"storage_location"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `tfsdk:"storage_root"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified catalog.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CatalogInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Kind of catalog securable.
type CatalogInfoSecurableKind string

const CatalogInfoSecurableKindCatalogDeltasharing CatalogInfoSecurableKind = `CATALOG_DELTASHARING`

const CatalogInfoSecurableKindCatalogForeignBigquery CatalogInfoSecurableKind = `CATALOG_FOREIGN_BIGQUERY`

const CatalogInfoSecurableKindCatalogForeignDatabricks CatalogInfoSecurableKind = `CATALOG_FOREIGN_DATABRICKS`

const CatalogInfoSecurableKindCatalogForeignMysql CatalogInfoSecurableKind = `CATALOG_FOREIGN_MYSQL`

const CatalogInfoSecurableKindCatalogForeignPostgresql CatalogInfoSecurableKind = `CATALOG_FOREIGN_POSTGRESQL`

const CatalogInfoSecurableKindCatalogForeignRedshift CatalogInfoSecurableKind = `CATALOG_FOREIGN_REDSHIFT`

const CatalogInfoSecurableKindCatalogForeignSnowflake CatalogInfoSecurableKind = `CATALOG_FOREIGN_SNOWFLAKE`

const CatalogInfoSecurableKindCatalogForeignSqldw CatalogInfoSecurableKind = `CATALOG_FOREIGN_SQLDW`

const CatalogInfoSecurableKindCatalogForeignSqlserver CatalogInfoSecurableKind = `CATALOG_FOREIGN_SQLSERVER`

const CatalogInfoSecurableKindCatalogInternal CatalogInfoSecurableKind = `CATALOG_INTERNAL`

const CatalogInfoSecurableKindCatalogOnline CatalogInfoSecurableKind = `CATALOG_ONLINE`

const CatalogInfoSecurableKindCatalogOnlineIndex CatalogInfoSecurableKind = `CATALOG_ONLINE_INDEX`

const CatalogInfoSecurableKindCatalogStandard CatalogInfoSecurableKind = `CATALOG_STANDARD`

const CatalogInfoSecurableKindCatalogSystem CatalogInfoSecurableKind = `CATALOG_SYSTEM`

const CatalogInfoSecurableKindCatalogSystemDeltasharing CatalogInfoSecurableKind = `CATALOG_SYSTEM_DELTASHARING`

// String representation for [fmt.Print]
func (f *CatalogInfoSecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogInfoSecurableKind) Set(v string) error {
	switch v {
	case `CATALOG_DELTASHARING`, `CATALOG_FOREIGN_BIGQUERY`, `CATALOG_FOREIGN_DATABRICKS`, `CATALOG_FOREIGN_MYSQL`, `CATALOG_FOREIGN_POSTGRESQL`, `CATALOG_FOREIGN_REDSHIFT`, `CATALOG_FOREIGN_SNOWFLAKE`, `CATALOG_FOREIGN_SQLDW`, `CATALOG_FOREIGN_SQLSERVER`, `CATALOG_INTERNAL`, `CATALOG_ONLINE`, `CATALOG_ONLINE_INDEX`, `CATALOG_STANDARD`, `CATALOG_SYSTEM`, `CATALOG_SYSTEM_DELTASHARING`:
		*f = CatalogInfoSecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG_DELTASHARING", "CATALOG_FOREIGN_BIGQUERY", "CATALOG_FOREIGN_DATABRICKS", "CATALOG_FOREIGN_MYSQL", "CATALOG_FOREIGN_POSTGRESQL", "CATALOG_FOREIGN_REDSHIFT", "CATALOG_FOREIGN_SNOWFLAKE", "CATALOG_FOREIGN_SQLDW", "CATALOG_FOREIGN_SQLSERVER", "CATALOG_INTERNAL", "CATALOG_ONLINE", "CATALOG_ONLINE_INDEX", "CATALOG_STANDARD", "CATALOG_SYSTEM", "CATALOG_SYSTEM_DELTASHARING"`, v)
	}
}

// Type always returns CatalogInfoSecurableKind to satisfy [pflag.Value] interface
func (f *CatalogInfoSecurableKind) Type() string {
	return "CatalogInfoSecurableKind"
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (f *CatalogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `MANAGED_CATALOG`, `SYSTEM_CATALOG`:
		*f = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "MANAGED_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (f *CatalogType) Type() string {
	return "CatalogType"
}

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	AccessKeyId string `tfsdk:"access_key_id"`
	// The account id associated with the API token.
	AccountId string `tfsdk:"account_id"`
	// The secret access token generated for the access key id
	SecretAccessKey string `tfsdk:"secret_access_key"`
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

type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Unique identifier of the Connection.
	ConnectionId string `tfsdk:"connection_id"`
	// The type of connection.
	ConnectionType ConnectionType `tfsdk:"connection_type"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of connection creator.
	CreatedBy string `tfsdk:"created_by"`
	// The type of credential.
	CredentialType CredentialType `tfsdk:"credential_type"`
	// Full name of connection.
	FullName string `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of the connection.
	Name string `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `tfsdk:"options"`
	// Username of current owner of the connection.
	Owner string `tfsdk:"owner"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `tfsdk:"properties"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo `tfsdk:"provisioning_info"`
	// If the connection is read only.
	ReadOnly bool `tfsdk:"read_only"`
	// Kind of connection securable.
	SecurableKind ConnectionInfoSecurableKind `tfsdk:"securable_kind"`

	SecurableType string `tfsdk:"securable_type"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified connection.
	UpdatedBy string `tfsdk:"updated_by"`
	// URL of the remote data source, extracted from options.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ConnectionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Kind of connection securable.
type ConnectionInfoSecurableKind string

const ConnectionInfoSecurableKindConnectionBigquery ConnectionInfoSecurableKind = `CONNECTION_BIGQUERY`

const ConnectionInfoSecurableKindConnectionDatabricks ConnectionInfoSecurableKind = `CONNECTION_DATABRICKS`

const ConnectionInfoSecurableKindConnectionMysql ConnectionInfoSecurableKind = `CONNECTION_MYSQL`

const ConnectionInfoSecurableKindConnectionOnlineCatalog ConnectionInfoSecurableKind = `CONNECTION_ONLINE_CATALOG`

const ConnectionInfoSecurableKindConnectionPostgresql ConnectionInfoSecurableKind = `CONNECTION_POSTGRESQL`

const ConnectionInfoSecurableKindConnectionRedshift ConnectionInfoSecurableKind = `CONNECTION_REDSHIFT`

const ConnectionInfoSecurableKindConnectionSnowflake ConnectionInfoSecurableKind = `CONNECTION_SNOWFLAKE`

const ConnectionInfoSecurableKindConnectionSqldw ConnectionInfoSecurableKind = `CONNECTION_SQLDW`

const ConnectionInfoSecurableKindConnectionSqlserver ConnectionInfoSecurableKind = `CONNECTION_SQLSERVER`

// String representation for [fmt.Print]
func (f *ConnectionInfoSecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionInfoSecurableKind) Set(v string) error {
	switch v {
	case `CONNECTION_BIGQUERY`, `CONNECTION_DATABRICKS`, `CONNECTION_MYSQL`, `CONNECTION_ONLINE_CATALOG`, `CONNECTION_POSTGRESQL`, `CONNECTION_REDSHIFT`, `CONNECTION_SNOWFLAKE`, `CONNECTION_SQLDW`, `CONNECTION_SQLSERVER`:
		*f = ConnectionInfoSecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONNECTION_BIGQUERY", "CONNECTION_DATABRICKS", "CONNECTION_MYSQL", "CONNECTION_ONLINE_CATALOG", "CONNECTION_POSTGRESQL", "CONNECTION_REDSHIFT", "CONNECTION_SNOWFLAKE", "CONNECTION_SQLDW", "CONNECTION_SQLSERVER"`, v)
	}
}

// Type always returns ConnectionInfoSecurableKind to satisfy [pflag.Value] interface
func (f *ConnectionInfoSecurableKind) Type() string {
	return "ConnectionInfoSecurableKind"
}

// The type of connection.
type ConnectionType string

const ConnectionTypeBigquery ConnectionType = `BIGQUERY`

const ConnectionTypeDatabricks ConnectionType = `DATABRICKS`

const ConnectionTypeMysql ConnectionType = `MYSQL`

const ConnectionTypePostgresql ConnectionType = `POSTGRESQL`

const ConnectionTypeRedshift ConnectionType = `REDSHIFT`

const ConnectionTypeSnowflake ConnectionType = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionType = `SQLDW`

const ConnectionTypeSqlserver ConnectionType = `SQLSERVER`

// String representation for [fmt.Print]
func (f *ConnectionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `DATABRICKS`, `MYSQL`, `POSTGRESQL`, `REDSHIFT`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`:
		*f = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "DATABRICKS", "MYSQL", "POSTGRESQL", "REDSHIFT", "SNOWFLAKE", "SQLDW", "SQLSERVER"`, v)
	}
}

// Type always returns ConnectionType to satisfy [pflag.Value] interface
func (f *ConnectionType) Type() string {
	return "ConnectionType"
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress *PipelineProgress `tfsdk:"initial_pipeline_sync_progress"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `tfsdk:"timestamp"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// The name of the connection to an external data source.
	ConnectionName string `tfsdk:"connection_name"`
	// Name of catalog.
	Name string `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `tfsdk:"options"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `tfsdk:"provider_name"`
	// The name of the share under the share provider.
	ShareName string `tfsdk:"share_name"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `tfsdk:"storage_root"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateConnection struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// The type of connection.
	ConnectionType ConnectionType `tfsdk:"connection_type"`
	// Name of the connection.
	Name string `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `tfsdk:"options"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `tfsdk:"properties"`
	// If the connection is read only.
	ReadOnly bool `tfsdk:"read_only"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `tfsdk:"access_point"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Name of the storage credential used with this location.
	CredentialName string `tfsdk:"credential_name"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `tfsdk:"encryption_details"`
	// Name of the external location.
	Name string `tfsdk:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `tfsdk:"read_only"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `tfsdk:"skip_validation"`
	// Path URL of the external location.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFunction struct {
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Scalar function return data type.
	DataType ColumnTypeName `tfsdk:"data_type"`
	// External function language.
	ExternalLanguage string `tfsdk:"external_language"`
	// External function name.
	ExternalName string `tfsdk:"external_name"`
	// Pretty printed function data type.
	FullDataType string `tfsdk:"full_data_type"`

	InputParams FunctionParameterInfos `tfsdk:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic bool `tfsdk:"is_deterministic"`
	// Function null call.
	IsNullCall bool `tfsdk:"is_null_call"`
	// Name of function, relative to parent schema.
	Name string `tfsdk:"name"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle CreateFunctionParameterStyle `tfsdk:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `tfsdk:"properties"`
	// Table function return parameters.
	ReturnParams FunctionParameterInfos `tfsdk:"return_params"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody CreateFunctionRoutineBody `tfsdk:"routine_body"`
	// Function body.
	RoutineDefinition string `tfsdk:"routine_definition"`
	// Function dependencies.
	RoutineDependencies DependencyList `tfsdk:"routine_dependencies"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `tfsdk:"schema_name"`
	// Function security type.
	SecurityType CreateFunctionSecurityType `tfsdk:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `tfsdk:"specific_name"`
	// Function SQL data access.
	SqlDataAccess CreateFunctionSqlDataAccess `tfsdk:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `tfsdk:"sql_path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Function parameter style. **S** is the value for SQL.
type CreateFunctionParameterStyle string

const CreateFunctionParameterStyleS CreateFunctionParameterStyle = `S`

// String representation for [fmt.Print]
func (f *CreateFunctionParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = CreateFunctionParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (f *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo CreateFunction `tfsdk:"function_info"`
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type CreateFunctionRoutineBody string

const CreateFunctionRoutineBodyExternal CreateFunctionRoutineBody = `EXTERNAL`

const CreateFunctionRoutineBodySql CreateFunctionRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *CreateFunctionRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = CreateFunctionRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (f *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

// Function security type.
type CreateFunctionSecurityType string

const CreateFunctionSecurityTypeDefiner CreateFunctionSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *CreateFunctionSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = CreateFunctionSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (f *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

// Function SQL data access.
type CreateFunctionSqlDataAccess string

const CreateFunctionSqlDataAccessContainsSql CreateFunctionSqlDataAccess = `CONTAINS_SQL`

const CreateFunctionSqlDataAccessNoSql CreateFunctionSqlDataAccess = `NO_SQL`

const CreateFunctionSqlDataAccessReadsSqlData CreateFunctionSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *CreateFunctionSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = CreateFunctionSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (f *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	Name string `tfsdk:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`). If
	// this field is omitted, the region of the workspace receiving the request
	// will be used.
	Region string `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot string `tfsdk:"storage_root"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `tfsdk:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `tfsdk:"data_classification_config"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `tfsdk:"inference_log"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName string `tfsdk:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `tfsdk:"schedule"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `tfsdk:"skip_builtin_dashboard"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `tfsdk:"snapshot"`
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `tfsdk:"time_series"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Online Table information.
type CreateOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `tfsdk:"name"`
	// Specification of the online table.
	Spec *OnlineTableSpec `tfsdk:"spec"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateOnlineTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOnlineTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `tfsdk:"catalog_name"`
	// The comment attached to the registered model
	Comment string `tfsdk:"comment"`
	// The name of the registered model
	Name string `tfsdk:"name"`
	// The name of the schema where the registered model resides
	SchemaName string `tfsdk:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `tfsdk:"storage_location"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateResponse struct {
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Name of schema, relative to parent catalog.
	Name string `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `tfsdk:"storage_root"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment string `tfsdk:"comment"`
	// The <Databricks> managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account"`
	// The credential name. The name must be unique within the metastore.
	Name string `tfsdk:"name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `tfsdk:"read_only"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool `tfsdk:"skip_validation"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint TableConstraint `tfsdk:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg string `tfsdk:"full_name_arg"`
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `tfsdk:"catalog_name"`
	// The comment attached to the volume
	Comment string `tfsdk:"comment"`
	// The name of the volume
	Name string `tfsdk:"name"`
	// The name of the schema where the volume is
	SchemaName string `tfsdk:"schema_name"`
	// The storage location on the cloud
	StorageLocation string `tfsdk:"storage_location"`

	VolumeType VolumeType `tfsdk:"volume_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of credential.
type CredentialType string

const CredentialTypeUsernamePassword CredentialType = `USERNAME_PASSWORD`

// String representation for [fmt.Print]
func (f *CredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialType) Set(v string) error {
	switch v {
	case `USERNAME_PASSWORD`:
		*f = CredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "USERNAME_PASSWORD"`, v)
	}
}

// Type always returns CredentialType to satisfy [pflag.Value] interface
func (f *CredentialType) Type() string {
	return "CredentialType"
}

// Currently assigned workspaces
type CurrentWorkspaceBindings struct {
	// A list of workspace IDs.
	Workspaces []int64 `tfsdk:"workspaces"`
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

// String representation for [fmt.Print]
func (f *DataSourceFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `CSV`, `DELTA`, `DELTASHARING`, `JSON`, `ORC`, `PARQUET`, `TEXT`, `UNITY_CATALOG`:
		*f = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "CSV", "DELTA", "DELTASHARING", "JSON", "ORC", "PARQUET", "TEXT", "UNITY_CATALOG"`, v)
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (f *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

type DatabricksGcpServiceAccountRequest struct {
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId string `tfsdk:"credential_id"`
	// The email of the service account. This is an output-only field.
	Email string `tfsdk:"email"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DatabricksGcpServiceAccountResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a Registered Model Alias
type DeleteAliasRequest struct {
	// The name of the alias
	Alias string `tfsdk:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"-" url:"-"`
}

type DeleteAliasResponse struct {
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// The name of the catalog.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a connection
type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name string `tfsdk:"-" url:"-"`
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// Name of the external location.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// Unique ID of the metastore.
	Id string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a Model Version
type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `tfsdk:"-" url:"-"`
	// The integer version number of the model version
	Version int `tfsdk:"-" url:"-"`
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `tfsdk:"-" url:"-"`
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Full name of the schema.
	FullName string `tfsdk:"-" url:"-"`
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `tfsdk:"-" url:"force,omitempty"`
	// Name of the storage credential.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `tfsdk:"-" url:"cascade"`
	// The name of the constraint to delete.
	ConstraintName string `tfsdk:"-" url:"constraint_name"`
	// Full name of the table referenced by the constraint.
	FullName string `tfsdk:"-" url:"-"`
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName string `tfsdk:"-" url:"-"`
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name string `tfsdk:"-" url:"-"`
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]string `tfsdk:"delta_runtime_properties"`
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function *FunctionDependency `tfsdk:"function"`
	// A table that is dependent on a SQL object.
	Table *TableDependency `tfsdk:"table"`
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies []Dependency `tfsdk:"dependencies"`
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `tfsdk:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `tfsdk:"-" url:"-"`
}

type DisableResponse struct {
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []EffectivePrivilegeAssignment `tfsdk:"privilege_assignments"`
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string `tfsdk:"inherited_from_name"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `tfsdk:"inherited_from_type"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value EnablePredictiveOptimization `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EffectivePredictiveOptimizationFlag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePredictiveOptimizationFlag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of the object from which the flag was inherited. If there was no
// inheritance, this field is left blank.
type EffectivePredictiveOptimizationFlagInheritedFromType string

const EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog EffectivePredictiveOptimizationFlagInheritedFromType = `CATALOG`

const EffectivePredictiveOptimizationFlagInheritedFromTypeSchema EffectivePredictiveOptimizationFlagInheritedFromType = `SCHEMA`

// String representation for [fmt.Print]
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Set(v string) error {
	switch v {
	case `CATALOG`, `SCHEMA`:
		*f = EffectivePredictiveOptimizationFlagInheritedFromType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "SCHEMA"`, v)
	}
}

// Type always returns EffectivePredictiveOptimizationFlagInheritedFromType to satisfy [pflag.Value] interface
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Type() string {
	return "EffectivePredictiveOptimizationFlagInheritedFromType"
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName string `tfsdk:"inherited_from_name"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType SecurableType `tfsdk:"inherited_from_type"`
	// The privilege assigned to the principal.
	Privilege Privilege `tfsdk:"privilege"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EffectivePrivilege) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilege) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `tfsdk:"principal"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege `tfsdk:"privileges"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EffectivePrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Whether predictive optimization should be enabled for this object and objects
// under it.
type EnablePredictiveOptimization string

const EnablePredictiveOptimizationDisable EnablePredictiveOptimization = `DISABLE`

const EnablePredictiveOptimizationEnable EnablePredictiveOptimization = `ENABLE`

const EnablePredictiveOptimizationInherit EnablePredictiveOptimization = `INHERIT`

// String representation for [fmt.Print]
func (f *EnablePredictiveOptimization) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EnablePredictiveOptimization) Set(v string) error {
	switch v {
	case `DISABLE`, `ENABLE`, `INHERIT`:
		*f = EnablePredictiveOptimization(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLE", "ENABLE", "INHERIT"`, v)
	}
}

// Type always returns EnablePredictiveOptimization to satisfy [pflag.Value] interface
func (f *EnablePredictiveOptimization) Type() string {
	return "EnablePredictiveOptimization"
}

// Enable a system schema
type EnableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `tfsdk:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `tfsdk:"-" url:"-"`
}

type EnableResponse struct {
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails *SseEncryptionDetails `tfsdk:"sse_encryption_details"`
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	FullName string `tfsdk:"-" url:"-"`
}

type ExternalLocationInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `tfsdk:"access_point"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of external location creator.
	CreatedBy string `tfsdk:"created_by"`
	// Unique ID of the location's storage credential.
	CredentialId string `tfsdk:"credential_id"`
	// Name of the storage credential used with this location.
	CredentialName string `tfsdk:"credential_name"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `tfsdk:"encryption_details"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of the external location.
	Name string `tfsdk:"name"`
	// The owner of the external location.
	Owner string `tfsdk:"owner"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `tfsdk:"read_only"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified the external location.
	UpdatedBy string `tfsdk:"updated_by"`
	// Path URL of the external location.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExternalLocationInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLocationInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp string `tfsdk:"timestamp"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `tfsdk:"child_columns"`
	// The name of the constraint.
	Name string `tfsdk:"name"`
	// Column names for this constraint.
	ParentColumns []string `tfsdk:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable string `tfsdk:"parent_table"`
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string `tfsdk:"function_full_name"`
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of function creator.
	CreatedBy string `tfsdk:"created_by"`
	// Scalar function return data type.
	DataType ColumnTypeName `tfsdk:"data_type"`
	// External function language.
	ExternalLanguage string `tfsdk:"external_language"`
	// External function name.
	ExternalName string `tfsdk:"external_name"`
	// Pretty printed function data type.
	FullDataType string `tfsdk:"full_data_type"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName string `tfsdk:"full_name"`
	// Id of Function, relative to parent schema.
	FunctionId string `tfsdk:"function_id"`

	InputParams *FunctionParameterInfos `tfsdk:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic bool `tfsdk:"is_deterministic"`
	// Function null call.
	IsNullCall bool `tfsdk:"is_null_call"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of function, relative to parent schema.
	Name string `tfsdk:"name"`
	// Username of current owner of function.
	Owner string `tfsdk:"owner"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle FunctionInfoParameterStyle `tfsdk:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `tfsdk:"properties"`
	// Table function return parameters.
	ReturnParams *FunctionParameterInfos `tfsdk:"return_params"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody FunctionInfoRoutineBody `tfsdk:"routine_body"`
	// Function body.
	RoutineDefinition string `tfsdk:"routine_definition"`
	// Function dependencies.
	RoutineDependencies *DependencyList `tfsdk:"routine_dependencies"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `tfsdk:"schema_name"`
	// Function security type.
	SecurityType FunctionInfoSecurityType `tfsdk:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `tfsdk:"specific_name"`
	// Function SQL data access.
	SqlDataAccess FunctionInfoSqlDataAccess `tfsdk:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `tfsdk:"sql_path"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified function.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FunctionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Function parameter style. **S** is the value for SQL.
type FunctionInfoParameterStyle string

const FunctionInfoParameterStyleS FunctionInfoParameterStyle = `S`

// String representation for [fmt.Print]
func (f *FunctionInfoParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = FunctionInfoParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (f *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type FunctionInfoRoutineBody string

const FunctionInfoRoutineBodyExternal FunctionInfoRoutineBody = `EXTERNAL`

const FunctionInfoRoutineBodySql FunctionInfoRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *FunctionInfoRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = FunctionInfoRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (f *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

// Function security type.
type FunctionInfoSecurityType string

const FunctionInfoSecurityTypeDefiner FunctionInfoSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *FunctionInfoSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = FunctionInfoSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (f *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

// Function SQL data access.
type FunctionInfoSqlDataAccess string

const FunctionInfoSqlDataAccessContainsSql FunctionInfoSqlDataAccess = `CONTAINS_SQL`

const FunctionInfoSqlDataAccessNoSql FunctionInfoSqlDataAccess = `NO_SQL`

const FunctionInfoSqlDataAccessReadsSqlData FunctionInfoSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *FunctionInfoSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = FunctionInfoSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (f *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Name of parameter.
	Name string `tfsdk:"name"`
	// Default value of the parameter.
	ParameterDefault string `tfsdk:"parameter_default"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `tfsdk:"parameter_mode"`
	// The type of function parameter.
	ParameterType FunctionParameterType `tfsdk:"parameter_type"`
	// Ordinal position of column (starting at position 0).
	Position int `tfsdk:"position"`
	// Format of IntervalType.
	TypeIntervalType string `tfsdk:"type_interval_type"`
	// Full data type spec, JSON-serialized.
	TypeJson string `tfsdk:"type_json"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName ColumnTypeName `tfsdk:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int `tfsdk:"type_precision"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int `tfsdk:"type_scale"`
	// Full data type spec, SQL/catalogString text.
	TypeText string `tfsdk:"type_text"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	Parameters []FunctionParameterInfo `tfsdk:"parameters"`
}

// The mode of the function parameter.
type FunctionParameterMode string

const FunctionParameterModeIn FunctionParameterMode = `IN`

// String representation for [fmt.Print]
func (f *FunctionParameterMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`:
		*f = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

// The type of function parameter.
type FunctionParameterType string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (f *FunctionParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*f = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `tfsdk:"-" url:"-"`
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `tfsdk:"-" url:"-"`
}

// Get securable workspace bindings
type GetBindingsRequest struct {
	// The name of the securable.
	SecurableName string `tfsdk:"-" url:"-"`
	// The type of the securable.
	SecurableType string `tfsdk:"-" url:"-"`
}

// Get Model Version By Alias
type GetByAliasRequest struct {
	// The name of the alias
	Alias string `tfsdk:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"-" url:"-"`
}

// Get a catalog
type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// The name of the catalog.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	Name string `tfsdk:"-" url:"-"`
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string `tfsdk:"-" url:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `tfsdk:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetEffectiveRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEffectiveRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an external location
type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Name of the external location.
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a function
type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	FullName string `tfsdk:"-" url:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `tfsdk:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetGrantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetGrantRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id string `tfsdk:"-" url:"-"`
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `tfsdk:"cloud"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of metastore creator.
	CreatedBy string `tfsdk:"created_by"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `tfsdk:"default_data_access_config_id"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope `tfsdk:"delta_sharing_scope"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `tfsdk:"global_metastore_id"`
	// Unique identifier of metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// The user-specified name of the metastore.
	Name string `tfsdk:"name"`
	// The owner of the metastore.
	Owner string `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `tfsdk:"privilege_model_version"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot string `tfsdk:"storage_root"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `tfsdk:"storage_root_credential_id"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `tfsdk:"storage_root_credential_name"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified the metastore.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetMetastoreSummaryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetastoreSummaryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type GetMetastoreSummaryResponseDeltaSharingScope string

const GetMetastoreSummaryResponseDeltaSharingScopeInternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL`

const GetMetastoreSummaryResponseDeltaSharingScopeInternalAndExternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *GetMetastoreSummaryResponseDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = GetMetastoreSummaryResponseDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns GetMetastoreSummaryResponseDeltaSharingScope to satisfy [pflag.Value] interface
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Type() string {
	return "GetMetastoreSummaryResponseDeltaSharingScope"
}

// Get a Model Version
type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `tfsdk:"-" url:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// The integer version number of the model version
	Version int `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an Online Table
type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `tfsdk:"-" url:"-"`
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `tfsdk:"-" url:"-"`
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"-" url:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	FullName string `tfsdk:"-" url:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name string `tfsdk:"-" url:"-"`
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName string `tfsdk:"-" url:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `tfsdk:"-" url:"include_delta_metadata,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get catalog workspace bindings
type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name string `tfsdk:"-" url:"-"`
}

// Whether the current securable is accessible from all workspaces or a specific
// set of workspaces.
type IsolationMode string

const IsolationModeIsolated IsolationMode = `ISOLATED`

const IsolationModeOpen IsolationMode = `OPEN`

// String representation for [fmt.Print]
func (f *IsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IsolationMode) Set(v string) error {
	switch v {
	case `ISOLATED`, `OPEN`:
		*f = IsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATED", "OPEN"`, v)
	}
}

// Type always returns IsolationMode to satisfy [pflag.Value] interface
func (f *IsolationMode) Type() string {
	return "IsolationMode"
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds []int64 `tfsdk:"workspace_ids"`
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `tfsdk:"-" url:"-"`
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials []StorageCredentialInfo `tfsdk:"storage_credentials"`
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `tfsdk:"catalogs"`
}

// List connections
type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
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

func (s *ListConnectionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections []ConnectionInfo `tfsdk:"connections"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListConnectionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List external locations
type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExternalLocationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations []ExternalLocationInfo `tfsdk:"external_locations"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListExternalLocationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `tfsdk:"-" url:"catalog_name"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// Parent schema of functions.
	SchemaName string `tfsdk:"-" url:"schema_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFunctionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Functions []FunctionInfo `tfsdk:"functions"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListFunctionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `tfsdk:"metastores"`
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string `tfsdk:"-" url:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelVersionsResponse struct {
	ModelVersions []ModelVersionInfo `tfsdk:"model_versions"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string `tfsdk:"-" url:"catalog_name,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Max number of registered models to return.
	//
	// If both catalog and schema are specified: - when max_results is not
	// specified, the page length is set to a server configured value (10000, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (10000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (10000, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	//
	// If neither schema nor catalog is specified: - when max_results is not
	// specified, the page length is set to a server configured value (100, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (1000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (100, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string `tfsdk:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListRegisteredModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken string `tfsdk:"next_page_token"`

	RegisteredModels []RegisteredModelInfo `tfsdk:"registered_models"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListRegisteredModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName string `tfsdk:"-" url:"catalog_name"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`
	// An array of schema information objects.
	Schemas []SchemaInfo `tfsdk:"schemas"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List credentials
type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`

	StorageCredentials []StorageCredentialInfo `tfsdk:"storage_credentials"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `tfsdk:"-" url:"catalog_name"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `tfsdk:"-" url:"schema_name_pattern,omitempty"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `tfsdk:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListSummariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSummariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List system schemas
type ListSystemSchemasRequest struct {
	// The ID for the metastore in which the system schema resides.
	MetastoreId string `tfsdk:"-" url:"-"`
}

type ListSystemSchemasResponse struct {
	// An array of system schema information objects.
	Schemas []SystemSchemaInfo `tfsdk:"schemas"`
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`
	// List of table summaries.
	Tables []TableSummary `tfsdk:"tables"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListTableSummariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTableSummariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `tfsdk:"-" url:"catalog_name"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `tfsdk:"-" url:"include_delta_metadata,omitempty"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool `tfsdk:"-" url:"omit_columns,omitempty"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool `tfsdk:"-" url:"omit_properties,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// Parent schema of tables.
	SchemaName string `tfsdk:"-" url:"schema_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListTablesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `tfsdk:"next_page_token"`
	// An array of table information objects.
	Tables []TableInfo `tfsdk:"tables"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName string `tfsdk:"-" url:"catalog_name"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// Maximum number of volumes to return (page length).
	//
	// If not set, the page length is set to a server configured value (10000,
	// as of 1/29/2024). - when set to a value greater than 0, the page length
	// is the minimum of this value and a server configured value (10000, as of
	// 1/29/2024); - when set to 0, the page length is set to a server
	// configured value (10000, as of 1/29/2024) (recommended); - when set to a
	// value less than 0, an invalid parameter error is returned;
	//
	// Note: this parameter controls only the maximum number of volumes to
	// return. The actual number of volumes returned in a page may be smaller
	// than this value, including 0, even if there are more pages.
	MaxResults int `tfsdk:"-" url:"max_results,omitempty"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`
	// The identifier of the schema
	SchemaName string `tfsdk:"-" url:"schema_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListVolumesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken string `tfsdk:"next_page_token"`

	Volumes []VolumeInfo `tfsdk:"volumes"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListVolumesResponseContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesResponseContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The artifact pattern matching type
type MatchType string

const MatchTypePrefixMatch MatchType = `PREFIX_MATCH`

// String representation for [fmt.Print]
func (f *MatchType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MatchType) Set(v string) error {
	switch v {
	case `PREFIX_MATCH`:
		*f = MatchType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREFIX_MATCH"`, v)
	}
}

// Type always returns MatchType to satisfy [pflag.Value] interface
func (f *MatchType) Type() string {
	return "MatchType"
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId int64 `tfsdk:"workspace_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `tfsdk:"cloud"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of metastore creator.
	CreatedBy string `tfsdk:"created_by"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `tfsdk:"default_data_access_config_id"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope MetastoreInfoDeltaSharingScope `tfsdk:"delta_sharing_scope"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `tfsdk:"global_metastore_id"`
	// Unique identifier of metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// The user-specified name of the metastore.
	Name string `tfsdk:"name"`
	// The owner of the metastore.
	Owner string `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `tfsdk:"privilege_model_version"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot string `tfsdk:"storage_root"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `tfsdk:"storage_root_credential_id"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `tfsdk:"storage_root_credential_name"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified the metastore.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MetastoreInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type MetastoreInfoDeltaSharingScope string

const MetastoreInfoDeltaSharingScopeInternal MetastoreInfoDeltaSharingScope = `INTERNAL`

const MetastoreInfoDeltaSharingScopeInternalAndExternal MetastoreInfoDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *MetastoreInfoDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MetastoreInfoDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = MetastoreInfoDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns MetastoreInfoDeltaSharingScope to satisfy [pflag.Value] interface
func (f *MetastoreInfoDeltaSharingScope) Type() string {
	return "MetastoreInfoDeltaSharingScope"
}

type ModelVersionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// The name of the catalog containing the model version
	CatalogName string `tfsdk:"catalog_name"`
	// The comment attached to the model version
	Comment string `tfsdk:"comment"`

	CreatedAt int64 `tfsdk:"created_at"`
	// The identifier of the user who created the model version
	CreatedBy string `tfsdk:"created_by"`
	// The unique identifier of the model version
	Id string `tfsdk:"id"`
	// The unique identifier of the metastore containing the model version
	MetastoreId string `tfsdk:"metastore_id"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `tfsdk:"model_name"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *DependencyList `tfsdk:"model_version_dependencies"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string `tfsdk:"run_id"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int `tfsdk:"run_workspace_id"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `tfsdk:"schema_name"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string `tfsdk:"source"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus `tfsdk:"status"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `tfsdk:"storage_location"`

	UpdatedAt int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the model version last time
	UpdatedBy string `tfsdk:"updated_by"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version int `tfsdk:"version"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Current status of the model version. Newly created model versions start in
// PENDING_REGISTRATION status, then move to READY status once the model version
// files are uploaded and the model version is finalized. Only model versions in
// READY status can be loaded for inference or served.
type ModelVersionInfoStatus string

const ModelVersionInfoStatusFailedRegistration ModelVersionInfoStatus = `FAILED_REGISTRATION`

const ModelVersionInfoStatusPendingRegistration ModelVersionInfoStatus = `PENDING_REGISTRATION`

const ModelVersionInfoStatusReady ModelVersionInfoStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionInfoStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns ModelVersionInfoStatus to satisfy [pflag.Value] interface
func (f *ModelVersionInfoStatus) Type() string {
	return "ModelVersionInfoStatus"
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus MonitorCronSchedulePauseStatus `tfsdk:"pause_status"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `tfsdk:"quartz_cron_expression"`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId string `tfsdk:"timezone_id"`
}

// Read only field that indicates whether a schedule is paused or not.
type MonitorCronSchedulePauseStatus string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatus = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *MonitorCronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorCronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = MonitorCronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	Enabled bool `tfsdk:"enabled"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MonitorDataClassificationConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorDataClassificationConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses []string `tfsdk:"email_addresses"`
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `tfsdk:"granularities"`
	// Optional column that contains the ground truth for the prediction.
	LabelCol string `tfsdk:"label_col"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol string `tfsdk:"model_id_col"`
	// Column that contains the output/prediction from the model.
	PredictionCol string `tfsdk:"prediction_col"`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol string `tfsdk:"prediction_proba_col"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType MonitorInferenceLogProblemType `tfsdk:"problem_type"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `tfsdk:"timestamp_col"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MonitorInferenceLog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInferenceLog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Problem type the model aims to solve. Determines the type of model-quality
// metrics that will be computed.
type MonitorInferenceLogProblemType string

const MonitorInferenceLogProblemTypeProblemTypeClassification MonitorInferenceLogProblemType = `PROBLEM_TYPE_CLASSIFICATION`

const MonitorInferenceLogProblemTypeProblemTypeRegression MonitorInferenceLogProblemType = `PROBLEM_TYPE_REGRESSION`

// String representation for [fmt.Print]
func (f *MonitorInferenceLogProblemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInferenceLogProblemType) Set(v string) error {
	switch v {
	case `PROBLEM_TYPE_CLASSIFICATION`, `PROBLEM_TYPE_REGRESSION`:
		*f = MonitorInferenceLogProblemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PROBLEM_TYPE_CLASSIFICATION", "PROBLEM_TYPE_REGRESSION"`, v)
	}
}

// Type always returns MonitorInferenceLogProblemType to satisfy [pflag.Value] interface
func (f *MonitorInferenceLogProblemType) Type() string {
	return "MonitorInferenceLogProblemType"
}

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `tfsdk:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `tfsdk:"dashboard_id"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `tfsdk:"data_classification_config"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName string `tfsdk:"drift_metrics_table_name"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `tfsdk:"inference_log"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg string `tfsdk:"latest_monitor_failure_msg"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion string `tfsdk:"monitor_version"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName string `tfsdk:"output_schema_name"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName string `tfsdk:"profile_metrics_table_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `tfsdk:"schedule"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `tfsdk:"snapshot"`
	// The status of the monitor.
	Status MonitorInfoStatus `tfsdk:"status"`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName string `tfsdk:"table_name"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `tfsdk:"time_series"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MonitorInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the monitor.
type MonitorInfoStatus string

const MonitorInfoStatusMonitorStatusActive MonitorInfoStatus = `MONITOR_STATUS_ACTIVE`

const MonitorInfoStatusMonitorStatusDeletePending MonitorInfoStatus = `MONITOR_STATUS_DELETE_PENDING`

const MonitorInfoStatusMonitorStatusError MonitorInfoStatus = `MONITOR_STATUS_ERROR`

const MonitorInfoStatusMonitorStatusFailed MonitorInfoStatus = `MONITOR_STATUS_FAILED`

const MonitorInfoStatusMonitorStatusPending MonitorInfoStatus = `MONITOR_STATUS_PENDING`

// String representation for [fmt.Print]
func (f *MonitorInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInfoStatus) Set(v string) error {
	switch v {
	case `MONITOR_STATUS_ACTIVE`, `MONITOR_STATUS_DELETE_PENDING`, `MONITOR_STATUS_ERROR`, `MONITOR_STATUS_FAILED`, `MONITOR_STATUS_PENDING`:
		*f = MonitorInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_STATUS_ACTIVE", "MONITOR_STATUS_DELETE_PENDING", "MONITOR_STATUS_ERROR", "MONITOR_STATUS_FAILED", "MONITOR_STATUS_PENDING"`, v)
	}
}

// Type always returns MonitorInfoStatus to satisfy [pflag.Value] interface
func (f *MonitorInfoStatus) Type() string {
	return "MonitorInfoStatus"
}

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `tfsdk:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `tfsdk:"input_columns"`
	// Name of the metric in the output tables.
	Name string `tfsdk:"name"`
	// The output type of the custom metric.
	OutputDataType string `tfsdk:"output_data_type"`
	// Can only be one of ``"CUSTOM_METRIC_TYPE_AGGREGATE"``,
	// ``"CUSTOM_METRIC_TYPE_DERIVED"``, or ``"CUSTOM_METRIC_TYPE_DRIFT"``. The
	// ``"CUSTOM_METRIC_TYPE_AGGREGATE"`` and ``"CUSTOM_METRIC_TYPE_DERIVED"``
	// metrics are computed on a single table, whereas the
	// ``"CUSTOM_METRIC_TYPE_DRIFT"`` compare metrics across baseline and input
	// table, or across the two consecutive time windows. -
	// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
	// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed
	// aggregate metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously
	// computed aggregate or derived metrics
	Type MonitorMetricType `tfsdk:"type"`
}

// Can only be one of “"CUSTOM_METRIC_TYPE_AGGREGATE"“,
// “"CUSTOM_METRIC_TYPE_DERIVED"“, or “"CUSTOM_METRIC_TYPE_DRIFT"“. The
// “"CUSTOM_METRIC_TYPE_AGGREGATE"“ and “"CUSTOM_METRIC_TYPE_DERIVED"“
// metrics are computed on a single table, whereas the
// “"CUSTOM_METRIC_TYPE_DRIFT"“ compare metrics across baseline and input
// table, or across the two consecutive time windows. -
// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed aggregate
// metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously computed aggregate
// or derived metrics
type MonitorMetricType string

const MonitorMetricTypeCustomMetricTypeAggregate MonitorMetricType = `CUSTOM_METRIC_TYPE_AGGREGATE`

const MonitorMetricTypeCustomMetricTypeDerived MonitorMetricType = `CUSTOM_METRIC_TYPE_DERIVED`

const MonitorMetricTypeCustomMetricTypeDrift MonitorMetricType = `CUSTOM_METRIC_TYPE_DRIFT`

// String representation for [fmt.Print]
func (f *MonitorMetricType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorMetricType) Set(v string) error {
	switch v {
	case `CUSTOM_METRIC_TYPE_AGGREGATE`, `CUSTOM_METRIC_TYPE_DERIVED`, `CUSTOM_METRIC_TYPE_DRIFT`:
		*f = MonitorMetricType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOM_METRIC_TYPE_AGGREGATE", "CUSTOM_METRIC_TYPE_DERIVED", "CUSTOM_METRIC_TYPE_DRIFT"`, v)
	}
}

// Type always returns MonitorMetricType to satisfy [pflag.Value] interface
func (f *MonitorMetricType) Type() string {
	return "MonitorMetricType"
}

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	OnFailure *MonitorDestination `tfsdk:"on_failure"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected *MonitorDestination `tfsdk:"on_new_classification_tag_detected"`
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message string `tfsdk:"message"`
	// Unique id of the refresh operation.
	RefreshId int64 `tfsdk:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs int64 `tfsdk:"start_time_ms"`
	// The current state of the refresh.
	State MonitorRefreshInfoState `tfsdk:"state"`
	// The method by which the refresh was triggered.
	Trigger MonitorRefreshInfoTrigger `tfsdk:"trigger"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MonitorRefreshInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorRefreshInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current state of the refresh.
type MonitorRefreshInfoState string

const MonitorRefreshInfoStateCanceled MonitorRefreshInfoState = `CANCELED`

const MonitorRefreshInfoStateFailed MonitorRefreshInfoState = `FAILED`

const MonitorRefreshInfoStatePending MonitorRefreshInfoState = `PENDING`

const MonitorRefreshInfoStateRunning MonitorRefreshInfoState = `RUNNING`

const MonitorRefreshInfoStateSuccess MonitorRefreshInfoState = `SUCCESS`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = MonitorRefreshInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

// The method by which the refresh was triggered.
type MonitorRefreshInfoTrigger string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTrigger = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTrigger = `SCHEDULE`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoTrigger) Set(v string) error {
	switch v {
	case `MANUAL`, `SCHEDULE`:
		*f = MonitorRefreshInfoTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANUAL", "SCHEDULE"`, v)
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	Refreshes []MonitorRefreshInfo `tfsdk:"refreshes"`
}

type MonitorSnapshot struct {
}

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `tfsdk:"granularities"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `tfsdk:"timestamp_col"`
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name string `tfsdk:"name"`
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `tfsdk:"name"`
	// Specification of the online table.
	Spec *OnlineTableSpec `tfsdk:"spec"`
	// Online Table status
	Status *OnlineTableStatus `tfsdk:"status"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *OnlineTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specification of an online table.
type OnlineTableSpec struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy bool `tfsdk:"perform_full_copy"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId string `tfsdk:"pipeline_id"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `tfsdk:"primary_key_columns"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy `tfsdk:"run_continuously"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy `tfsdk:"run_triggered"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `tfsdk:"source_table_full_name"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `tfsdk:"timeseries_key"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *OnlineTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

// The state of an online table.
type OnlineTableState string

const OnlineTableStateOffline OnlineTableState = `OFFLINE`

const OnlineTableStateOfflineFailed OnlineTableState = `OFFLINE_FAILED`

const OnlineTableStateOnline OnlineTableState = `ONLINE`

const OnlineTableStateOnlineContinuousUpdate OnlineTableState = `ONLINE_CONTINUOUS_UPDATE`

const OnlineTableStateOnlineNoPendingUpdate OnlineTableState = `ONLINE_NO_PENDING_UPDATE`

const OnlineTableStateOnlinePipelineFailed OnlineTableState = `ONLINE_PIPELINE_FAILED`

const OnlineTableStateOnlineTableStateUnspecified OnlineTableState = `ONLINE_TABLE_STATE_UNSPECIFIED`

const OnlineTableStateOnlineTriggeredUpdate OnlineTableState = `ONLINE_TRIGGERED_UPDATE`

const OnlineTableStateOnlineUpdatingPipelineResources OnlineTableState = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const OnlineTableStateProvisioning OnlineTableState = `PROVISIONING`

const OnlineTableStateProvisioningInitialSnapshot OnlineTableState = `PROVISIONING_INITIAL_SNAPSHOT`

const OnlineTableStateProvisioningPipelineResources OnlineTableState = `PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *OnlineTableState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineTableState) Set(v string) error {
	switch v {
	case `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TABLE_STATE_UNSPECIFIED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = OnlineTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TABLE_STATE_UNSPECIFIED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Type always returns OnlineTableState to satisfy [pflag.Value] interface
func (f *OnlineTableState) Type() string {
	return "OnlineTableState"
}

// Status of an online table.
type OnlineTableStatus struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus *ContinuousUpdateStatus `tfsdk:"continuous_update_status"`
	// The state of the online table.
	DetailedState OnlineTableState `tfsdk:"detailed_state"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus *FailedStatus `tfsdk:"failed_status"`
	// A text description of the current state of the online table.
	Message string `tfsdk:"message"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus *ProvisioningStatus `tfsdk:"provisioning_status"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus *TriggeredUpdateStatus `tfsdk:"triggered_update_status"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *OnlineTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []Privilege `tfsdk:"add"`
	// The principal whose privileges we are changing.
	Principal string `tfsdk:"principal"`
	// The set of privileges to remove.
	Remove []Privilege `tfsdk:"remove"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments"`
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `tfsdk:"estimated_completion_time_seconds"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `tfsdk:"latest_version_currently_processing"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64 `tfsdk:"sync_progress_completion"`
	// The number of rows that have been synced in this update.
	SyncedRowCount int64 `tfsdk:"synced_row_count"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64 `tfsdk:"total_row_count"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `tfsdk:"child_columns"`
	// The name of the constraint.
	Name string `tfsdk:"name"`
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

// An object containing map of key-value properties attached to the connection.
type PropertiesKvPairs map[string]string

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	State ProvisioningInfoState `tfsdk:"state"`
}

type ProvisioningInfoState string

const ProvisioningInfoStateActive ProvisioningInfoState = `ACTIVE`

const ProvisioningInfoStateDeleting ProvisioningInfoState = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoState = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoState = `PROVISIONING`

const ProvisioningInfoStateStateUnspecified ProvisioningInfoState = `STATE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ProvisioningInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningInfoState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETING`, `FAILED`, `PROVISIONING`, `STATE_UNSPECIFIED`:
		*f = ProvisioningInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETING", "FAILED", "PROVISIONING", "STATE_UNSPECIFIED"`, v)
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress *PipelineProgress `tfsdk:"initial_pipeline_sync_progress"`
}

// Get a Volume
type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `tfsdk:"-" url:"include_browse,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ReadVolumeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadVolumeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName string `tfsdk:"alias_name"`
	// Integer version number of the model version to which this alias points.
	VersionNum int `tfsdk:"version_num"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `tfsdk:"aliases"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `tfsdk:"catalog_name"`
	// The comment attached to the registered model
	Comment string `tfsdk:"comment"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `tfsdk:"created_at"`
	// The identifier of the user who created the registered model
	CreatedBy string `tfsdk:"created_by"`
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"full_name"`
	// The unique identifier of the metastore
	MetastoreId string `tfsdk:"metastore_id"`
	// The name of the registered model
	Name string `tfsdk:"name"`
	// The identifier of the user who owns the registered model
	Owner string `tfsdk:"owner"`
	// The name of the schema where the registered model resides
	SchemaName string `tfsdk:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `tfsdk:"storage_location"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RegisteredModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
}

type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// The type of the parent catalog.
	CatalogType string `tfsdk:"catalog_type"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of schema creator.
	CreatedBy string `tfsdk:"created_by"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `tfsdk:"enable_predictive_optimization"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of schema, relative to parent catalog.
	Name string `tfsdk:"name"`
	// Username of current owner of schema.
	Owner string `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`
	// The unique identifier of the schema.
	SchemaId string `tfsdk:"schema_id"`
	// Storage location for managed tables within schema.
	StorageLocation string `tfsdk:"storage_location"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `tfsdk:"storage_root"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified schema.
	UpdatedBy string `tfsdk:"updated_by"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A map of key-value properties attached to the securable.
type SecurableOptionsMap map[string]string

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[string]string

// The type of Unity Catalog securable
type SecurableType string

const SecurableTypeCatalog SecurableType = `catalog`

const SecurableTypeConnection SecurableType = `connection`

const SecurableTypeExternalLocation SecurableType = `external_location`

const SecurableTypeFunction SecurableType = `function`

const SecurableTypeMetastore SecurableType = `metastore`

const SecurableTypePipeline SecurableType = `pipeline`

const SecurableTypeProvider SecurableType = `provider`

const SecurableTypeRecipient SecurableType = `recipient`

const SecurableTypeSchema SecurableType = `schema`

const SecurableTypeShare SecurableType = `share`

const SecurableTypeStorageCredential SecurableType = `storage_credential`

const SecurableTypeTable SecurableType = `table`

const SecurableTypeVolume SecurableType = `volume`

// String representation for [fmt.Print]
func (f *SecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableType) Set(v string) error {
	switch v {
	case `catalog`, `connection`, `external_location`, `function`, `metastore`, `pipeline`, `provider`, `recipient`, `schema`, `share`, `storage_credential`, `table`, `volume`:
		*f = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "connection", "external_location", "function", "metastore", "pipeline", "provider", "recipient", "schema", "share", "storage_credential", "table", "volume"`, v)
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (f *SecurableType) Type() string {
	return "SecurableType"
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `tfsdk:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `tfsdk:"-" url:"-"`
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias string `tfsdk:"alias" url:"-"`
	// Full name of the registered model
	FullName string `tfsdk:"full_name" url:"-"`
	// The version number of the model version to which the alias points
	VersionNum int `tfsdk:"version_num"`
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm SseEncryptionDetailsAlgorithm `tfsdk:"algorithm"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn string `tfsdk:"aws_kms_key_arn"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SseEncryptionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SseEncryptionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of key encryption to use (affects headers from s3 client).
type SseEncryptionDetailsAlgorithm string

const SseEncryptionDetailsAlgorithmAwsSseKms SseEncryptionDetailsAlgorithm = `AWS_SSE_KMS`

const SseEncryptionDetailsAlgorithmAwsSseS3 SseEncryptionDetailsAlgorithm = `AWS_SSE_S3`

// String representation for [fmt.Print]
func (f *SseEncryptionDetailsAlgorithm) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SseEncryptionDetailsAlgorithm) Set(v string) error {
	switch v {
	case `AWS_SSE_KMS`, `AWS_SSE_S3`:
		*f = SseEncryptionDetailsAlgorithm(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_SSE_KMS", "AWS_SSE_S3"`, v)
	}
}

// Type always returns SseEncryptionDetailsAlgorithm to satisfy [pflag.Value] interface
func (f *SseEncryptionDetailsAlgorithm) Type() string {
	return "SseEncryptionDetailsAlgorithm"
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleResponse `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment string `tfsdk:"comment"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of credential creator.
	CreatedBy string `tfsdk:"created_by"`
	// The <Databricks> managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse `tfsdk:"databricks_gcp_service_account"`
	// The unique identifier of the credential.
	Id string `tfsdk:"id"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// The credential name. The name must be unique within the metastore.
	Name string `tfsdk:"name"`
	// Username of current owner of credential.
	Owner string `tfsdk:"owner"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `tfsdk:"read_only"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified the credential.
	UpdatedBy string `tfsdk:"updated_by"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage bool `tfsdk:"used_for_managed_storage"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *StorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageCredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SystemSchemaInfo struct {
	// Name of the system schema.
	Schema string `tfsdk:"schema"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State SystemSchemaInfoState `tfsdk:"state"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SystemSchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SystemSchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current state of enablement for the system schema. An empty string means
// the system schema is available and ready for opt-in.
type SystemSchemaInfoState string

const SystemSchemaInfoStateAvailable SystemSchemaInfoState = `AVAILABLE`

const SystemSchemaInfoStateDisableInitialized SystemSchemaInfoState = `DISABLE_INITIALIZED`

const SystemSchemaInfoStateEnableCompleted SystemSchemaInfoState = `ENABLE_COMPLETED`

const SystemSchemaInfoStateEnableInitialized SystemSchemaInfoState = `ENABLE_INITIALIZED`

const SystemSchemaInfoStateUnavailable SystemSchemaInfoState = `UNAVAILABLE`

// String representation for [fmt.Print]
func (f *SystemSchemaInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SystemSchemaInfoState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DISABLE_INITIALIZED`, `ENABLE_COMPLETED`, `ENABLE_INITIALIZED`, `UNAVAILABLE`:
		*f = SystemSchemaInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DISABLE_INITIALIZED", "ENABLE_COMPLETED", "ENABLE_INITIALIZED", "UNAVAILABLE"`, v)
	}
}

// Type always returns SystemSchemaInfoState to satisfy [pflag.Value] interface
func (f *SystemSchemaInfoState) Type() string {
	return "SystemSchemaInfoState"
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint *ForeignKeyConstraint `tfsdk:"foreign_key_constraint"`

	NamedTableConstraint *NamedTableConstraint `tfsdk:"named_table_constraint"`

	PrimaryKeyConstraint *PrimaryKeyConstraint `tfsdk:"primary_key_constraint"`
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName string `tfsdk:"table_full_name"`
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists bool `tfsdk:"table_exists"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TableExistsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableExistsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `tfsdk:"access_point"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `tfsdk:"columns"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64 `tfsdk:"created_at"`
	// Username of table creator.
	CreatedBy string `tfsdk:"created_by"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string `tfsdk:"data_access_configuration_id"`
	// Data source format
	DataSourceFormat DataSourceFormat `tfsdk:"data_source_format"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64 `tfsdk:"deleted_at"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs `tfsdk:"delta_runtime_properties_kvpairs"`

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `tfsdk:"enable_predictive_optimization"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `tfsdk:"encryption_details"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// Name of table, relative to parent schema.
	Name string `tfsdk:"name"`
	// Username of current owner of table.
	Owner string `tfsdk:"owner"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId string `tfsdk:"pipeline_id"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`

	RowFilter *TableRowFilter `tfsdk:"row_filter"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `tfsdk:"schema_name"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `tfsdk:"sql_path"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string `tfsdk:"storage_credential_name"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation string `tfsdk:"storage_location"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []TableConstraint `tfsdk:"table_constraints"`
	// The unique identifier of the table.
	TableId string `tfsdk:"table_id"`

	TableType TableType `tfsdk:"table_type"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64 `tfsdk:"updated_at"`
	// Username of user who last modified the table.
	UpdatedBy string `tfsdk:"updated_by"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string `tfsdk:"view_definition"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies *DependencyList `tfsdk:"view_dependencies"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName string `tfsdk:"function_name"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string `tfsdk:"input_column_names"`
}

type TableSummary struct {
	// The full name of the table.
	FullName string `tfsdk:"full_name"`

	TableType TableType `tfsdk:"table_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TableSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeManaged TableType = `MANAGED`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (f *TableType) Type() string {
	return "TableType"
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *PipelineProgress `tfsdk:"triggered_update_progress"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *TriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `tfsdk:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type UnassignResponse struct {
}

type UpdateAssignmentResponse struct {
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `tfsdk:"enable_predictive_optimization"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `tfsdk:"isolation_mode"`
	// The name of the catalog.
	Name string `tfsdk:"-" url:"-"`
	// New name for the catalog.
	NewName string `tfsdk:"new_name"`
	// Username of current owner of catalog.
	Owner string `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateConnection struct {
	// Name of the connection.
	Name string `tfsdk:"-" url:"-"`
	// New name for the connection.
	NewName string `tfsdk:"new_name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `tfsdk:"options"`
	// Username of current owner of the connection.
	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `tfsdk:"access_point"`
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Name of the storage credential used with this location.
	CredentialName string `tfsdk:"credential_name"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `tfsdk:"encryption_details"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `tfsdk:"force"`
	// Name of the external location.
	Name string `tfsdk:"-" url:"-"`
	// New name for the external location.
	NewName string `tfsdk:"new_name"`
	// The owner of the external location.
	Owner string `tfsdk:"owner"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `tfsdk:"read_only"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `tfsdk:"skip_validation"`
	// Path URL of the external location.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `tfsdk:"-" url:"-"`
	// Username of current owner of function.
	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope UpdateMetastoreDeltaSharingScope `tfsdk:"delta_sharing_scope"`
	// Unique ID of the metastore.
	Id string `tfsdk:"-" url:"-"`
	// New name for the metastore.
	NewName string `tfsdk:"new_name"`
	// The owner of the metastore.
	Owner string `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `tfsdk:"privilege_model_version"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `tfsdk:"storage_root_credential_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog for the metastore.
	DefaultCatalogName string `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `tfsdk:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The scope of Delta Sharing enabled for the metastore.
type UpdateMetastoreDeltaSharingScope string

const UpdateMetastoreDeltaSharingScopeInternal UpdateMetastoreDeltaSharingScope = `INTERNAL`

const UpdateMetastoreDeltaSharingScopeInternalAndExternal UpdateMetastoreDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *UpdateMetastoreDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateMetastoreDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = UpdateMetastoreDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns UpdateMetastoreDeltaSharingScope to satisfy [pflag.Value] interface
func (f *UpdateMetastoreDeltaSharingScope) Type() string {
	return "UpdateMetastoreDeltaSharingScope"
}

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	Comment string `tfsdk:"comment"`
	// The three-level (fully qualified) name of the model version
	FullName string `tfsdk:"-" url:"-"`
	// The integer version number of the model version
	Version int `tfsdk:"-" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `tfsdk:"dashboard_id"`
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig `tfsdk:"data_classification_config"`
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog `tfsdk:"inference_log"`
	// The notification settings for the monitor.
	Notifications *MonitorNotifications `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName string `tfsdk:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule `tfsdk:"schedule"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `tfsdk:"snapshot"`
	// Full name of the table.
	TableName string `tfsdk:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `tfsdk:"time_series"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `tfsdk:"changes"`
	// Full name of securable.
	FullName string `tfsdk:"-" url:"-"`
	// Type of securable.
	SecurableType SecurableType `tfsdk:"-" url:"-"`
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment string `tfsdk:"comment"`
	// The three-level (fully qualified) name of the registered model
	FullName string `tfsdk:"-" url:"-"`
	// New name for the registered model.
	NewName string `tfsdk:"new_name"`
	// The identifier of the user who owns the registered model
	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateResponse struct {
}

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment string `tfsdk:"comment"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `tfsdk:"enable_predictive_optimization"`
	// Full name of the schema.
	FullName string `tfsdk:"-" url:"-"`
	// New name for the schema.
	NewName string `tfsdk:"new_name"`
	// Username of current owner of schema.
	Owner string `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `tfsdk:"properties"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment string `tfsdk:"comment"`
	// The <Databricks> managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool `tfsdk:"force"`
	// Name of the storage credential.
	Name string `tfsdk:"-" url:"-"`
	// New name for the storage credential.
	NewName string `tfsdk:"new_name"`
	// Username of current owner of credential.
	Owner string `tfsdk:"owner"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `tfsdk:"read_only"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool `tfsdk:"skip_validation"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	FullName string `tfsdk:"-" url:"-"`

	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment string `tfsdk:"comment"`
	// The three-level (fully qualified) name of the volume
	Name string `tfsdk:"-" url:"-"`
	// New name for the volume.
	NewName string `tfsdk:"new_name"`
	// The identifier of the user who owns the volume
	Owner string `tfsdk:"owner"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	AssignWorkspaces []int64 `tfsdk:"assign_workspaces"`
	// The name of the catalog.
	Name string `tfsdk:"-" url:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces []int64 `tfsdk:"unassign_workspaces"`
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	Add []WorkspaceBinding `tfsdk:"add"`
	// List of workspace bindings
	Remove []WorkspaceBinding `tfsdk:"remove"`
	// The name of the securable.
	SecurableName string `tfsdk:"-" url:"-"`
	// The type of the securable.
	SecurableType string `tfsdk:"-" url:"-"`
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `tfsdk:"cloudflare_api_token"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account"`
	// The name of an existing external location to validate.
	ExternalLocationName string `tfsdk:"external_location_name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `tfsdk:"read_only"`
	// The name of the storage credential to validate.
	StorageCredentialName string `tfsdk:"storage_credential_name"`
	// The external location url to validate.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ValidateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir bool `tfsdk:"isDir"`
	// The results of the validation check.
	Results []ValidationResult `tfsdk:"results"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ValidateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `tfsdk:"message"`
	// The operation tested.
	Operation ValidationResultOperation `tfsdk:"operation"`
	// The results of the tested operation.
	Result ValidationResultResult `tfsdk:"result"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The operation tested.
type ValidationResultOperation string

const ValidationResultOperationDelete ValidationResultOperation = `DELETE`

const ValidationResultOperationList ValidationResultOperation = `LIST`

const ValidationResultOperationPathExists ValidationResultOperation = `PATH_EXISTS`

const ValidationResultOperationRead ValidationResultOperation = `READ`

const ValidationResultOperationWrite ValidationResultOperation = `WRITE`

// String representation for [fmt.Print]
func (f *ValidationResultOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultOperation) Set(v string) error {
	switch v {
	case `DELETE`, `LIST`, `PATH_EXISTS`, `READ`, `WRITE`:
		*f = ValidationResultOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "LIST", "PATH_EXISTS", "READ", "WRITE"`, v)
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (f *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

// The results of the tested operation.
type ValidationResultResult string

const ValidationResultResultFail ValidationResultResult = `FAIL`

const ValidationResultResultPass ValidationResultResult = `PASS`

const ValidationResultResultSkip ValidationResultResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidationResultResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidationResultResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (f *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `tfsdk:"access_point"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the volume are
	CatalogName string `tfsdk:"catalog_name"`
	// The comment attached to the volume
	Comment string `tfsdk:"comment"`

	CreatedAt int64 `tfsdk:"created_at"`
	// The identifier of the user who created the volume
	CreatedBy string `tfsdk:"created_by"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `tfsdk:"encryption_details"`
	// The three-level (fully qualified) name of the volume
	FullName string `tfsdk:"full_name"`
	// The unique identifier of the metastore
	MetastoreId string `tfsdk:"metastore_id"`
	// The name of the volume
	Name string `tfsdk:"name"`
	// The identifier of the user who owns the volume
	Owner string `tfsdk:"owner"`
	// The name of the schema where the volume is
	SchemaName string `tfsdk:"schema_name"`
	// The storage location on the cloud
	StorageLocation string `tfsdk:"storage_location"`

	UpdatedAt int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the volume last time
	UpdatedBy string `tfsdk:"updated_by"`
	// The unique identifier of the volume
	VolumeId string `tfsdk:"volume_id"`

	VolumeType VolumeType `tfsdk:"volume_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *VolumeInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VolumeInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VolumeType string

const VolumeTypeExternal VolumeType = `EXTERNAL`

const VolumeTypeManaged VolumeType = `MANAGED`

// String representation for [fmt.Print]
func (f *VolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VolumeType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`:
		*f = VolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED"`, v)
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (f *VolumeType) Type() string {
	return "VolumeType"
}

type WorkspaceBinding struct {
	BindingType WorkspaceBindingBindingType `tfsdk:"binding_type"`

	WorkspaceId int64 `tfsdk:"workspace_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceBinding) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceBindingBindingType string

const WorkspaceBindingBindingTypeBindingTypeReadOnly WorkspaceBindingBindingType = `BINDING_TYPE_READ_ONLY`

const WorkspaceBindingBindingTypeBindingTypeReadWrite WorkspaceBindingBindingType = `BINDING_TYPE_READ_WRITE`

// String representation for [fmt.Print]
func (f *WorkspaceBindingBindingType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceBindingBindingType) Set(v string) error {
	switch v {
	case `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`:
		*f = WorkspaceBindingBindingType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BINDING_TYPE_READ_ONLY", "BINDING_TYPE_READ_WRITE"`, v)
	}
}

// Type always returns WorkspaceBindingBindingType to satisfy [pflag.Value] interface
func (f *WorkspaceBindingBindingType) Type() string {
	return "WorkspaceBindingBindingType"
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings []WorkspaceBinding `tfsdk:"bindings"`
}
