// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package catalog_tf

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountsCreateMetastore struct {
	MetastoreInfo []CreateMetastore `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastore) {
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastore) {
}

type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment []CreateMetastoreAssignment `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastoreAssignment) {
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastoreAssignment) {
}

type AccountsCreateStorageCredential struct {
	CredentialInfo []CreateStorageCredential `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateStorageCredential) {
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsCreateStorageCredential) {
}

type AccountsMetastoreAssignment struct {
	MetastoreAssignment []MetastoreAssignment `tfsdk:"metastore_assignment" tf:"optional,object"`
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreAssignment) {
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreAssignment) {
}

type AccountsMetastoreInfo struct {
	MetastoreInfo []MetastoreInfo `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreInfo) {
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreInfo) {
}

type AccountsStorageCredentialInfo struct {
	CredentialInfo []StorageCredentialInfo `tfsdk:"credential_info" tf:"optional,object"`
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsStorageCredentialInfo) {
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState AccountsStorageCredentialInfo) {
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`

	MetastoreInfo []UpdateMetastore `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastore) {
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastore) {
}

type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment []UpdateMetastoreAssignment `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastoreAssignment) {
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastoreAssignment) {
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo []UpdateStorageCredential `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateStorageCredential) {
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateStorageCredential) {
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `tfsdk:"artifact_matchers" tf:"optional"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of the user who set the artifact allowlist.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactAllowlistInfo) {
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringRead(existingState ArtifactAllowlistInfo) {
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact types.String `tfsdk:"artifact" tf:""`
	// The pattern matching type of the artifact
	MatchType types.String `tfsdk:"match_type" tf:""`
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactMatcher) {
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringRead(existingState ArtifactMatcher) {
}

type AssignResponse struct {
}

func (newState *AssignResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AssignResponse) {
}

func (newState *AssignResponse) SyncEffectiveFieldsDuringRead(existingState AssignResponse) {
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:"optional"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken types.String `tfsdk:"session_token" tf:"optional"`
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials) {
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringRead(existingState AwsCredentials) {
}

type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn" tf:""`
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleRequest) {
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleRequest) {
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn" tf:""`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn" tf:"optional"`
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleResponse) {
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleResponse) {
}

type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityRequest) {
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityRequest) {
}

type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Databricks internal ID that represents this managed identity.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityResponse) {
}

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityResponse) {
}

type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId types.String `tfsdk:"application_id" tf:""`
	// The client secret generated for the above app ID in AAD.
	ClientSecret types.String `tfsdk:"client_secret" tf:""`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId types.String `tfsdk:"directory_id" tf:""`
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureServicePrincipal) {
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringRead(existingState AzureServicePrincipal) {
}

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken types.String `tfsdk:"sas_token" tf:"optional"`
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureUserDelegationSas) {
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringRead(existingState AzureUserDelegationSas) {
}

// Cancel refresh
type CancelRefreshRequest struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *CancelRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshRequest) {
}

func (newState *CancelRefreshRequest) SyncEffectiveFieldsDuringRead(existingState CancelRefreshRequest) {
}

type CancelRefreshResponse struct {
}

func (newState *CancelRefreshResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshResponse) {
}

func (newState *CancelRefreshResponse) SyncEffectiveFieldsDuringRead(existingState CancelRefreshResponse) {
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The type of the catalog.
	CatalogType types.String `tfsdk:"catalog_type" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of catalog creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	EffectivePredictiveOptimizationFlag []EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// The full name of the catalog. Corresponds with the name field.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of catalog.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options map[string]types.String `tfsdk:"options" tf:"optional"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name" tf:"optional"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo []ProvisioningInfo `tfsdk:"provisioning_info" tf:"optional,object"`
	// Kind of catalog securable.
	SecurableKind types.String `tfsdk:"securable_kind" tf:"optional"`

	SecurableType types.String `tfsdk:"securable_type" tf:"optional"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified catalog.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *CatalogInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CatalogInfo) {
}

func (newState *CatalogInfo) SyncEffectiveFieldsDuringRead(existingState CatalogInfo) {
}

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:""`
	// The account id associated with the API token.
	AccountId types.String `tfsdk:"account_id" tf:""`
	// The secret access token generated for the access key id
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:""`
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudflareApiToken) {
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringRead(existingState CloudflareApiToken) {
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`

	Mask []ColumnMask `tfsdk:"mask" tf:"optional,object"`
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

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
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

func (newState *ColumnMask) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnMask) {
}

func (newState *ColumnMask) SyncEffectiveFieldsDuringRead(existingState ColumnMask) {
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Unique identifier of the Connection.
	ConnectionId types.String `tfsdk:"connection_id" tf:"optional"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type" tf:"optional"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of connection creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The type of credential.
	CredentialType types.String `tfsdk:"credential_type" tf:"optional"`
	// Full name of connection.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of the connection.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options map[string]types.String `tfsdk:"options" tf:"optional"`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo []ProvisioningInfo `tfsdk:"provisioning_info" tf:"optional,object"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Kind of connection securable.
	SecurableKind types.String `tfsdk:"securable_kind" tf:"optional"`

	SecurableType types.String `tfsdk:"securable_type" tf:"optional"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified connection.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// URL of the remote data source, extracted from options.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConnectionInfo) {
}

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringRead(existingState ConnectionInfo) {
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress []PipelineProgress `tfsdk:"initial_pipeline_sync_progress" tf:"optional,object"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContinuousUpdateStatus) {
}

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringRead(existingState ContinuousUpdateStatus) {
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Name of catalog.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Options map[string]types.String `tfsdk:"options" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name" tf:"optional"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCatalog) {
}

func (newState *CreateCatalog) SyncEffectiveFieldsDuringRead(existingState CreateCatalog) {
}

type CreateConnection struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type" tf:""`
	// Name of the connection.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Options map[string]types.String `tfsdk:"options" tf:""`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
}

func (newState *CreateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateConnection) {
}

func (newState *CreateConnection) SyncEffectiveFieldsDuringRead(existingState CreateConnection) {
}

type CreateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:""`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails []EncryptionDetails `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"name" tf:""`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:""`
}

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExternalLocation) {
}

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringRead(existingState CreateExternalLocation) {
}

type CreateFunction struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type" tf:""`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language" tf:"optional"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name" tf:"optional"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type" tf:""`

	InputParams []FunctionParameterInfos `tfsdk:"input_params" tf:"object"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic" tf:""`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call" tf:""`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name" tf:""`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style" tf:""`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties" tf:"optional"`
	// Table function return parameters.
	ReturnParams []FunctionParameterInfos `tfsdk:"return_params" tf:"optional,object"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body" tf:""`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition" tf:""`
	// Function dependencies.
	RoutineDependencies []DependencyList `tfsdk:"routine_dependencies" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type" tf:""`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name" tf:""`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access" tf:""`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
}

func (newState *CreateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunction) {
}

func (newState *CreateFunction) SyncEffectiveFieldsDuringRead(existingState CreateFunction) {
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo []CreateFunction `tfsdk:"function_info" tf:"object"`
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunctionRequest) {
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringRead(existingState CreateFunctionRequest) {
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastore) {
}

func (newState *CreateMetastore) SyncEffectiveFieldsDuringRead(existingState CreateMetastore) {
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:""`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastoreAssignment) {
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState CreateMetastoreAssignment) {
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir" tf:""`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig []MonitorDataClassificationConfig `tfsdk:"data_classification_config" tf:"optional,object"`
	// Configuration for monitoring inference logs.
	InferenceLog []MonitorInferenceLog `tfsdk:"inference_log" tf:"optional,object"`
	// The notification settings for the monitor.
	Notifications []MonitorNotifications `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule []MonitorCronSchedule `tfsdk:"schedule" tf:"optional,object"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard types.Bool `tfsdk:"skip_builtin_dashboard" tf:"optional"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []types.String `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries []MonitorTimeSeries `tfsdk:"time_series" tf:"optional,object"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMonitor) {
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringRead(existingState CreateMonitor) {
}

// Online Table information.
type CreateOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Specification of the online table.
	Spec []OnlineTableSpec `tfsdk:"spec" tf:"optional,object"`
}

func (newState *CreateOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOnlineTableRequest) {
}

func (newState *CreateOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState CreateOnlineTableRequest) {
}

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the registered model
	Name types.String `tfsdk:"name" tf:""`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
}

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegisteredModelRequest) {
}

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState CreateRegisteredModelRequest) {
}

type CreateResponse struct {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSchema) {
}

func (newState *CreateSchema) SyncEffectiveFieldsDuringRead(existingState CreateSchema) {
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole []AwsIamRoleRequest `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity []AzureManagedIdentityRequest `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal []AzureServicePrincipal `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken []CloudflareApiToken `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageCredential) {
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState CreateStorageCredential) {
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint []TableConstraint `tfsdk:"constraint" tf:"object"`
	// The full name of the table referenced by the constraint.
	FullNameArg types.String `tfsdk:"full_name_arg" tf:""`
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTableConstraint) {
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringRead(existingState CreateTableConstraint) {
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the volume
	Name types.String `tfsdk:"name" tf:""`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	VolumeType types.String `tfsdk:"volume_type" tf:""`
}

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVolumeRequestContent) {
}

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState CreateVolumeRequestContent) {
}

// Currently assigned workspaces
type CurrentWorkspaceBindings struct {
	// A list of workspace IDs.
	Workspaces []types.Int64 `tfsdk:"workspaces" tf:"optional"`
}

func (newState *CurrentWorkspaceBindings) SyncEffectiveFieldsDuringCreateOrUpdate(plan CurrentWorkspaceBindings) {
}

func (newState *CurrentWorkspaceBindings) SyncEffectiveFieldsDuringRead(existingState CurrentWorkspaceBindings) {
}

type DatabricksGcpServiceAccountRequest struct {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountRequest) {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountRequest) {
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The email of the service account. This is an output-only field.
	Email types.String `tfsdk:"email" tf:"optional"`
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountResponse) {
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountResponse) {
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreAssignmentRequest) {
}

func (newState *DeleteAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreAssignmentRequest) {
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *DeleteAccountMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreRequest) {
}

func (newState *DeleteAccountMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreRequest) {
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *DeleteAccountStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountStorageCredentialRequest) {
}

func (newState *DeleteAccountStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountStorageCredentialRequest) {
}

// Delete a Registered Model Alias
type DeleteAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasRequest) {
}

func (newState *DeleteAliasRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAliasRequest) {
}

type DeleteAliasResponse struct {
}

func (newState *DeleteAliasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasResponse) {
}

func (newState *DeleteAliasResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAliasResponse) {
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCatalogRequest) {
}

func (newState *DeleteCatalogRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCatalogRequest) {
}

// Delete a connection
type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteConnectionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteConnectionRequest) {
}

func (newState *DeleteConnectionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteConnectionRequest) {
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteExternalLocationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExternalLocationRequest) {
}

func (newState *DeleteExternalLocationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteExternalLocationRequest) {
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFunctionRequest) {
}

func (newState *DeleteFunctionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteFunctionRequest) {
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteMetastoreRequest) {
}

func (newState *DeleteMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState DeleteMetastoreRequest) {
}

// Delete a Model Version
type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionRequest) {
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionRequest) {
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteOnlineTableRequest) {
}

func (newState *DeleteOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState DeleteOnlineTableRequest) {
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *DeleteQualityMonitorRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQualityMonitorRequest) {
}

func (newState *DeleteQualityMonitorRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQualityMonitorRequest) {
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRegisteredModelRequest) {
}

func (newState *DeleteRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRegisteredModelRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	Force types.Bool `tfsdk:"-"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteSchemaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSchemaRequest) {
}

func (newState *DeleteSchemaRequest) SyncEffectiveFieldsDuringRead(existingState DeleteSchemaRequest) {
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"-"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteStorageCredentialRequest) {
}

func (newState *DeleteStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteStorageCredentialRequest) {
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade types.Bool `tfsdk:"-"`
	// The name of the constraint to delete.
	ConstraintName types.String `tfsdk:"-"`
	// Full name of the table referenced by the constraint.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteTableConstraintRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableConstraintRequest) {
}

func (newState *DeleteTableConstraintRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTableConstraintRequest) {
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableRequest) {
}

func (newState *DeleteTableRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTableRequest) {
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteVolumeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVolumeRequest) {
}

func (newState *DeleteVolumeRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVolumeRequest) {
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]types.String `tfsdk:"delta_runtime_properties" tf:""`
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaRuntimePropertiesKvPairs) {
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState DeltaRuntimePropertiesKvPairs) {
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function []FunctionDependency `tfsdk:"function" tf:"optional,object"`
	// A table that is dependent on a SQL object.
	Table []TableDependency `tfsdk:"table" tf:"optional,object"`
}

func (newState *Dependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dependency) {
}

func (newState *Dependency) SyncEffectiveFieldsDuringRead(existingState Dependency) {
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies []Dependency `tfsdk:"dependencies" tf:"optional"`
}

func (newState *DependencyList) SyncEffectiveFieldsDuringCreateOrUpdate(plan DependencyList) {
}

func (newState *DependencyList) SyncEffectiveFieldsDuringRead(existingState DependencyList) {
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *DisableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableRequest) {
}

func (newState *DisableRequest) SyncEffectiveFieldsDuringRead(existingState DisableRequest) {
}

type DisableResponse struct {
}

func (newState *DisableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableResponse) {
}

func (newState *DisableResponse) SyncEffectiveFieldsDuringRead(existingState DisableResponse) {
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []EffectivePrivilegeAssignment `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePermissionsList) {
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringRead(existingState EffectivePermissionsList) {
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName types.String `tfsdk:"inherited_from_name" tf:"optional"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType types.String `tfsdk:"inherited_from_type" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePredictiveOptimizationFlag) {
}

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringRead(existingState EffectivePredictiveOptimizationFlag) {
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName types.String `tfsdk:"inherited_from_name" tf:"optional"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType types.String `tfsdk:"inherited_from_type" tf:"optional"`
	// The privilege assigned to the principal.
	Privilege types.String `tfsdk:"privilege" tf:"optional"`
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilege) {
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilege) {
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege `tfsdk:"privileges" tf:"optional"`
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilegeAssignment) {
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilegeAssignment) {
}

// Enable a system schema
type EnableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableRequest) {
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringRead(existingState EnableRequest) {
}

type EnableResponse struct {
}

func (newState *EnableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableResponse) {
}

func (newState *EnableResponse) SyncEffectiveFieldsDuringRead(existingState EnableResponse) {
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails []SseEncryptionDetails `tfsdk:"sse_encryption_details" tf:"optional,object"`
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan EncryptionDetails) {
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringRead(existingState EncryptionDetails) {
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

func (newState *ExistsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExistsRequest) {
}

func (newState *ExistsRequest) SyncEffectiveFieldsDuringRead(existingState ExistsRequest) {
}

type ExternalLocationInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of external location creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique ID of the location's storage credential.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails []EncryptionDetails `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the external location.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLocationInfo) {
}

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringRead(existingState ExternalLocationInfo) {
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *FailedStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan FailedStatus) {
}

func (newState *FailedStatus) SyncEffectiveFieldsDuringRead(existingState FailedStatus) {
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []types.String `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
	// Column names for this constraint.
	ParentColumns []types.String `tfsdk:"parent_columns" tf:""`
	// The full name of the parent constraint.
	ParentTable types.String `tfsdk:"parent_table" tf:""`
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForeignKeyConstraint) {
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringRead(existingState ForeignKeyConstraint) {
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName types.String `tfsdk:"function_full_name" tf:""`
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionDependency) {
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringRead(existingState FunctionDependency) {
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of function creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type" tf:"optional"`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language" tf:"optional"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name" tf:"optional"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type" tf:"optional"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Id of Function, relative to parent schema.
	FunctionId types.String `tfsdk:"function_id" tf:"optional"`

	InputParams []FunctionParameterInfos `tfsdk:"input_params" tf:"optional,object"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic" tf:"optional"`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style" tf:"optional"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties" tf:"optional"`
	// Table function return parameters.
	ReturnParams []FunctionParameterInfos `tfsdk:"return_params" tf:"optional,object"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body" tf:"optional"`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition" tf:"optional"`
	// Function dependencies.
	RoutineDependencies []DependencyList `tfsdk:"routine_dependencies" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type" tf:"optional"`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name" tf:"optional"`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access" tf:"optional"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified function.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *FunctionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionInfo) {
}

func (newState *FunctionInfo) SyncEffectiveFieldsDuringRead(existingState FunctionInfo) {
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of parameter.
	Name types.String `tfsdk:"name" tf:""`
	// Default value of the parameter.
	ParameterDefault types.String `tfsdk:"parameter_default" tf:"optional"`
	// The mode of the function parameter.
	ParameterMode types.String `tfsdk:"parameter_mode" tf:"optional"`
	// The type of function parameter.
	ParameterType types.String `tfsdk:"parameter_type" tf:"optional"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:""`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// Full data type spec, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json" tf:"optional"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName types.String `tfsdk:"type_name" tf:""`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type spec, SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:""`
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfo) {
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfo) {
}

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	Parameters []FunctionParameterInfo `tfsdk:"parameters" tf:"optional"`
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfos) {
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfos) {
}

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {
	OauthToken types.String `tfsdk:"oauth_token" tf:"optional"`
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpOauthToken) {
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringRead(existingState GcpOauthToken) {
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// UUID of the table to read or write.
	TableId types.String `tfsdk:"table_id" tf:"optional"`
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialRequest) {
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialRequest) {
}

type GenerateTemporaryTableCredentialResponse struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials []AwsCredentials `tfsdk:"aws_temp_credentials" tf:"optional,object"`
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas []AzureUserDelegationSas `tfsdk:"azure_user_delegation_sas" tf:"optional,object"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken []GcpOauthToken `tfsdk:"gcp_oauth_token" tf:"optional,object"`
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials []R2Credentials `tfsdk:"r2_temp_credentials" tf:"optional,object"`
	// The URL of the storage path accessible by the temporary credential.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialResponse) {
}

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialResponse) {
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreAssignmentRequest) {
}

func (newState *GetAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreAssignmentRequest) {
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *GetAccountMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreRequest) {
}

func (newState *GetAccountMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreRequest) {
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *GetAccountStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountStorageCredentialRequest) {
}

func (newState *GetAccountStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountStorageCredentialRequest) {
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

func (newState *GetArtifactAllowlistRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetArtifactAllowlistRequest) {
}

func (newState *GetArtifactAllowlistRequest) SyncEffectiveFieldsDuringRead(existingState GetArtifactAllowlistRequest) {
}

// Get securable workspace bindings
type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetBindingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBindingsRequest) {
}

func (newState *GetBindingsRequest) SyncEffectiveFieldsDuringRead(existingState GetBindingsRequest) {
}

// Get Model Version By Alias
type GetByAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases types.Bool `tfsdk:"-"`
}

func (newState *GetByAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetByAliasRequest) {
}

func (newState *GetByAliasRequest) SyncEffectiveFieldsDuringRead(existingState GetByAliasRequest) {
}

// Get a catalog
type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *GetCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCatalogRequest) {
}

func (newState *GetCatalogRequest) SyncEffectiveFieldsDuringRead(existingState GetCatalogRequest) {
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
}

func (newState *GetConnectionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetConnectionRequest) {
}

func (newState *GetConnectionRequest) SyncEffectiveFieldsDuringRead(existingState GetConnectionRequest) {
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetEffectiveRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEffectiveRequest) {
}

func (newState *GetEffectiveRequest) SyncEffectiveFieldsDuringRead(existingState GetEffectiveRequest) {
}

// Get an external location
type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

func (newState *GetExternalLocationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExternalLocationRequest) {
}

func (newState *GetExternalLocationRequest) SyncEffectiveFieldsDuringRead(existingState GetExternalLocationRequest) {
}

// Get a function
type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

func (newState *GetFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFunctionRequest) {
}

func (newState *GetFunctionRequest) SyncEffectiveFieldsDuringRead(existingState GetFunctionRequest) {
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetGrantRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGrantRequest) {
}

func (newState *GetGrantRequest) SyncEffectiveFieldsDuringRead(existingState GetGrantRequest) {
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

func (newState *GetMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreRequest) {
}

func (newState *GetMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState GetMetastoreRequest) {
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id" tf:"optional"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled" tf:"optional"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name" tf:"optional"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreSummaryResponse) {
}

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetastoreSummaryResponse) {
}

// Get a Model Version
type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases types.Bool `tfsdk:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionRequest) {
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionRequest) {
}

// Get an Online Table
type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

func (newState *GetOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOnlineTableRequest) {
}

func (newState *GetOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState GetOnlineTableRequest) {
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *GetQualityMonitorRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQualityMonitorRequest) {
}

func (newState *GetQualityMonitorRequest) SyncEffectiveFieldsDuringRead(existingState GetQualityMonitorRequest) {
}

// Get information for a single resource quota.
type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName types.String `tfsdk:"-"`
	// Securable type of the quota parent.
	ParentSecurableType types.String `tfsdk:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName types.String `tfsdk:"-"`
}

func (newState *GetQuotaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaRequest) {
}

func (newState *GetQuotaRequest) SyncEffectiveFieldsDuringRead(existingState GetQuotaRequest) {
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	QuotaInfo []QuotaInfo `tfsdk:"quota_info" tf:"optional,object"`
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaResponse) {
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringRead(existingState GetQuotaResponse) {
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *GetRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRefreshRequest) {
}

func (newState *GetRefreshRequest) SyncEffectiveFieldsDuringRead(existingState GetRefreshRequest) {
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases types.Bool `tfsdk:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

func (newState *GetRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelRequest) {
}

func (newState *GetRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelRequest) {
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

func (newState *GetSchemaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSchemaRequest) {
}

func (newState *GetSchemaRequest) SyncEffectiveFieldsDuringRead(existingState GetSchemaRequest) {
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

func (newState *GetStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStorageCredentialRequest) {
}

func (newState *GetStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetStorageCredentialRequest) {
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata types.Bool `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
}

func (newState *GetTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTableRequest) {
}

func (newState *GetTableRequest) SyncEffectiveFieldsDuringRead(existingState GetTableRequest) {
}

// Get catalog workspace bindings
type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *GetWorkspaceBindingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceBindingRequest) {
}

func (newState *GetWorkspaceBindingRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceBindingRequest) {
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountMetastoreAssignmentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsRequest) {
}

func (newState *ListAccountMetastoreAssignmentsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsRequest) {
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds []types.Int64 `tfsdk:"workspace_ids" tf:"optional"`
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsResponse) {
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsResponse) {
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountStorageCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsRequest) {
}

func (newState *ListAccountStorageCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsRequest) {
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials []StorageCredentialInfo `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsResponse) {
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsResponse) {
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCatalogsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsRequest) {
}

func (newState *ListCatalogsRequest) SyncEffectiveFieldsDuringRead(existingState ListCatalogsRequest) {
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `tfsdk:"catalogs" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsResponse) {
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringRead(existingState ListCatalogsResponse) {
}

// List connections
type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListConnectionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsRequest) {
}

func (newState *ListConnectionsRequest) SyncEffectiveFieldsDuringRead(existingState ListConnectionsRequest) {
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections []ConnectionInfo `tfsdk:"connections" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsResponse) {
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringRead(existingState ListConnectionsResponse) {
}

// List external locations
type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListExternalLocationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsRequest) {
}

func (newState *ListExternalLocationsRequest) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsRequest) {
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations []ExternalLocationInfo `tfsdk:"external_locations" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsResponse) {
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsResponse) {
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// Parent schema of functions.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *ListFunctionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsRequest) {
}

func (newState *ListFunctionsRequest) SyncEffectiveFieldsDuringRead(existingState ListFunctionsRequest) {
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Functions []FunctionInfo `tfsdk:"functions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsResponse) {
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringRead(existingState ListFunctionsResponse) {
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `tfsdk:"metastores" tf:"optional"`
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListMetastoresResponse) {
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringRead(existingState ListMetastoresResponse) {
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName types.String `tfsdk:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListModelVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsRequest) {
}

func (newState *ListModelVersionsRequest) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsRequest) {
}

type ListModelVersionsResponse struct {
	ModelVersions []ModelVersionInfo `tfsdk:"model_versions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsResponse) {
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsResponse) {
}

// List all resource quotas under a metastore.
type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token for the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQuotasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasRequest) {
}

func (newState *ListQuotasRequest) SyncEffectiveFieldsDuringRead(existingState ListQuotasRequest) {
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of returned QuotaInfos.
	Quotas []QuotaInfo `tfsdk:"quotas" tf:"optional"`
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasResponse) {
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringRead(existingState ListQuotasResponse) {
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *ListRefreshesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRefreshesRequest) {
}

func (newState *ListRefreshesRequest) SyncEffectiveFieldsDuringRead(existingState ListRefreshesRequest) {
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
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
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *ListRegisteredModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsRequest) {
}

func (newState *ListRegisteredModelsRequest) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsRequest) {
}

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels []RegisteredModelInfo `tfsdk:"registered_models" tf:"optional"`
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsResponse) {
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsResponse) {
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSchemasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasRequest) {
}

func (newState *ListSchemasRequest) SyncEffectiveFieldsDuringRead(existingState ListSchemasRequest) {
}

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of schema information objects.
	Schemas []SchemaInfo `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasResponse) {
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSchemasResponse) {
}

// List credentials
type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListStorageCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsRequest) {
}

func (newState *ListStorageCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsRequest) {
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	StorageCredentials []StorageCredentialInfo `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsResponse) {
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsResponse) {
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern types.String `tfsdk:"-"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern types.String `tfsdk:"-"`
}

func (newState *ListSummariesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSummariesRequest) {
}

func (newState *ListSummariesRequest) SyncEffectiveFieldsDuringRead(existingState ListSummariesRequest) {
}

// List system schemas
type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults types.Int64 `tfsdk:"-"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSystemSchemasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasRequest) {
}

func (newState *ListSystemSchemasRequest) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasRequest) {
}

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of system schema information objects.
	Schemas []SystemSchemaInfo `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasResponse) {
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasResponse) {
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of table summaries.
	Tables []TableSummary `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTableSummariesResponse) {
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringRead(existingState ListTableSummariesResponse) {
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata types.Bool `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns types.Bool `tfsdk:"-"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties types.Bool `tfsdk:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// Parent schema of tables.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *ListTablesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesRequest) {
}

func (newState *ListTablesRequest) SyncEffectiveFieldsDuringRead(existingState ListTablesRequest) {
}

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of table information objects.
	Tables []TableInfo `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesResponse) {
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringRead(existingState ListTablesResponse) {
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName types.String `tfsdk:"-"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
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
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// The identifier of the schema
	SchemaName types.String `tfsdk:"-"`
}

func (newState *ListVolumesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesRequest) {
}

func (newState *ListVolumesRequest) SyncEffectiveFieldsDuringRead(existingState ListVolumesRequest) {
}

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Volumes []VolumeInfo `tfsdk:"volumes" tf:"optional"`
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesResponseContent) {
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringRead(existingState ListVolumesResponseContent) {
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// The unique ID of the Databricks workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:""`
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreAssignment) {
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState MetastoreAssignment) {
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id" tf:"optional"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled" tf:"optional"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name" tf:"optional"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreInfo) {
}

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringRead(existingState MetastoreInfo) {
}

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	Aliases []RegisteredModelAlias `tfsdk:"aliases" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog containing the model version
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the model version
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The unique identifier of the model version
	Id types.String `tfsdk:"id" tf:"optional"`
	// The unique identifier of the metastore containing the model version
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies []DependencyList `tfsdk:"model_version_dependencies" tf:"optional,object"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId types.Int64 `tfsdk:"run_workspace_id" tf:"optional"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source types.String `tfsdk:"source" tf:"optional"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the model version last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionInfo) {
}

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringRead(existingState ModelVersionInfo) {
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression" tf:""`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId types.String `tfsdk:"timezone_id" tf:""`
}

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorCronSchedule) {
}

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringRead(existingState MonitorCronSchedule) {
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDataClassificationConfig) {
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringRead(existingState MonitorDataClassificationConfig) {
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses []types.String `tfsdk:"email_addresses" tf:"optional"`
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDestination) {
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringRead(existingState MonitorDestination) {
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []types.String `tfsdk:"granularities" tf:""`
	// Optional column that contains the ground truth for the prediction.
	LabelCol types.String `tfsdk:"label_col" tf:"optional"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol types.String `tfsdk:"model_id_col" tf:""`
	// Column that contains the output/prediction from the model.
	PredictionCol types.String `tfsdk:"prediction_col" tf:""`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol types.String `tfsdk:"prediction_proba_col" tf:"optional"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType types.String `tfsdk:"problem_type" tf:""`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col" tf:""`
}

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInferenceLog) {
}

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringRead(existingState MonitorInferenceLog) {
}

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir" tf:"optional"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics" tf:"optional"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig []MonitorDataClassificationConfig `tfsdk:"data_classification_config" tf:"optional,object"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName types.String `tfsdk:"drift_metrics_table_name" tf:""`
	// Configuration for monitoring inference logs.
	InferenceLog []MonitorInferenceLog `tfsdk:"inference_log" tf:"optional,object"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg types.String `tfsdk:"latest_monitor_failure_msg" tf:"optional"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion types.String `tfsdk:"monitor_version" tf:""`
	// The notification settings for the monitor.
	Notifications []MonitorNotifications `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:"optional"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName types.String `tfsdk:"profile_metrics_table_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule []MonitorCronSchedule `tfsdk:"schedule" tf:"optional,object"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []types.String `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// The status of the monitor.
	Status types.String `tfsdk:"status" tf:""`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName types.String `tfsdk:"table_name" tf:""`
	// Configuration for monitoring time series tables.
	TimeSeries []MonitorTimeSeries `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInfo) {
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringRead(existingState MonitorInfo) {
}

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition types.String `tfsdk:"definition" tf:""`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []types.String `tfsdk:"input_columns" tf:""`
	// Name of the metric in the output tables.
	Name types.String `tfsdk:"name" tf:""`
	// The output type of the custom metric.
	OutputDataType types.String `tfsdk:"output_data_type" tf:""`
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
	Type types.String `tfsdk:"type" tf:""`
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorMetric) {
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringRead(existingState MonitorMetric) {
}

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	OnFailure []MonitorDestination `tfsdk:"on_failure" tf:"optional,object"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected []MonitorDestination `tfsdk:"on_new_classification_tag_detected" tf:"optional,object"`
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorNotifications) {
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringRead(existingState MonitorNotifications) {
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms" tf:"optional"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message" tf:"optional"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id" tf:""`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms" tf:""`
	// The current state of the refresh.
	State types.String `tfsdk:"state" tf:""`
	// The method by which the refresh was triggered.
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
}

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshInfo) {
}

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshInfo) {
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	Refreshes []MonitorRefreshInfo `tfsdk:"refreshes" tf:"optional"`
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshListResponse) {
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshListResponse) {
}

type MonitorSnapshot struct {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorSnapshot) {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringRead(existingState MonitorSnapshot) {
}

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []types.String `tfsdk:"granularities" tf:""`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col" tf:""`
}

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorTimeSeries) {
}

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringRead(existingState MonitorTimeSeries) {
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan NamedTableConstraint) {
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringRead(existingState NamedTableConstraint) {
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Specification of the online table.
	Spec []OnlineTableSpec `tfsdk:"spec" tf:"optional,object"`
	// Online Table data synchronization status
	Status []OnlineTableStatus `tfsdk:"status" tf:"optional,object"`
	// Data serving REST API URL for this table
	TableServingUrl          types.String `tfsdk:"table_serving_url" tf:"optional"`
	EffectiveTableServingUrl types.String `tfsdk:"effective_table_serving_url" tf:"computed,optional"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state" tf:"optional"`
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTable) {
	newState.EffectiveTableServingUrl = newState.TableServingUrl
	newState.TableServingUrl = plan.TableServingUrl
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringRead(existingState OnlineTable) {
	newState.EffectiveTableServingUrl = existingState.EffectiveTableServingUrl
	if existingState.EffectiveTableServingUrl.ValueString() == newState.TableServingUrl.ValueString() {
		newState.TableServingUrl = existingState.TableServingUrl
	}
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
	PerformFullCopy types.Bool `tfsdk:"perform_full_copy" tf:"optional"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId          types.String `tfsdk:"pipeline_id" tf:"optional"`
	EffectivePipelineId types.String `tfsdk:"effective_pipeline_id" tf:"computed,optional"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []types.String `tfsdk:"primary_key_columns" tf:"optional"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously []OnlineTableSpecContinuousSchedulingPolicy `tfsdk:"run_continuously" tf:"optional,object"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered []OnlineTableSpecTriggeredSchedulingPolicy `tfsdk:"run_triggered" tf:"optional,object"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name" tf:"optional"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key" tf:"optional"`
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpec) {
	newState.EffectivePipelineId = newState.PipelineId
	newState.PipelineId = plan.PipelineId
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpec) {
	newState.EffectivePipelineId = existingState.EffectivePipelineId
	if existingState.EffectivePipelineId.ValueString() == newState.PipelineId.ValueString() {
		newState.PipelineId = existingState.PipelineId
	}
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecContinuousSchedulingPolicy) {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecContinuousSchedulingPolicy) {
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecTriggeredSchedulingPolicy) {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecTriggeredSchedulingPolicy) {
}

// Status of an online table.
type OnlineTableStatus struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus []ContinuousUpdateStatus `tfsdk:"continuous_update_status" tf:"optional,object"`
	// The state of the online table.
	DetailedState types.String `tfsdk:"detailed_state" tf:"optional"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus []FailedStatus `tfsdk:"failed_status" tf:"optional,object"`
	// A text description of the current state of the online table.
	Message types.String `tfsdk:"message" tf:"optional"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus []ProvisioningStatus `tfsdk:"provisioning_status" tf:"optional,object"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus []TriggeredUpdateStatus `tfsdk:"triggered_update_status" tf:"optional,object"`
}

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableStatus) {
}

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringRead(existingState OnlineTableStatus) {
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []types.String `tfsdk:"add" tf:"optional"`
	// The principal whose privileges we are changing.
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The set of privileges to remove.
	Remove []types.String `tfsdk:"remove" tf:"optional"`
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsChange) {
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringRead(existingState PermissionsChange) {
}

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsList) {
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringRead(existingState PermissionsList) {
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds types.Float64 `tfsdk:"estimated_completion_time_seconds" tf:"optional"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing types.Int64 `tfsdk:"latest_version_currently_processing" tf:"optional"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion types.Float64 `tfsdk:"sync_progress_completion" tf:"optional"`
	// The number of rows that have been synced in this update.
	SyncedRowCount types.Int64 `tfsdk:"synced_row_count" tf:"optional"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount types.Int64 `tfsdk:"total_row_count" tf:"optional"`
}

func (newState *PipelineProgress) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineProgress) {
}

func (newState *PipelineProgress) SyncEffectiveFieldsDuringRead(existingState PipelineProgress) {
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []types.String `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrimaryKeyConstraint) {
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringRead(existingState PrimaryKeyConstraint) {
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges []types.String `tfsdk:"privileges" tf:"optional"`
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment) {
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment) {
}

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningInfo) {
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringRead(existingState ProvisioningInfo) {
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress []PipelineProgress `tfsdk:"initial_pipeline_sync_progress" tf:"optional,object"`
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningStatus) {
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringRead(existingState ProvisioningStatus) {
}

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt types.Int64 `tfsdk:"last_refreshed_at" tf:"optional"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName types.String `tfsdk:"parent_full_name" tf:"optional"`
	// The quota parent securable type.
	ParentSecurableType types.String `tfsdk:"parent_securable_type" tf:"optional"`
	// The current usage of the resource quota.
	QuotaCount types.Int64 `tfsdk:"quota_count" tf:"optional"`
	// The current limit of the resource quota.
	QuotaLimit types.Int64 `tfsdk:"quota_limit" tf:"optional"`
	// The name of the quota.
	QuotaName types.String `tfsdk:"quota_name" tf:"optional"`
}

func (newState *QuotaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuotaInfo) {
}

func (newState *QuotaInfo) SyncEffectiveFieldsDuringRead(existingState QuotaInfo) {
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:"optional"`
	// The secret access key associated with the access key.
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:"optional"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken types.String `tfsdk:"session_token" tf:"optional"`
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan R2Credentials) {
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringRead(existingState R2Credentials) {
}

// Get a Volume
type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

func (newState *ReadVolumeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReadVolumeRequest) {
}

func (newState *ReadVolumeRequest) SyncEffectiveFieldsDuringRead(existingState ReadVolumeRequest) {
}

type RegenerateDashboardRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardRequest) {
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardRequest) {
}

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder types.String `tfsdk:"parent_folder" tf:"optional"`
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardResponse) {
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardResponse) {
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName types.String `tfsdk:"alias_name" tf:"optional"`
	// Integer version number of the model version to which this alias points.
	VersionNum types.Int64 `tfsdk:"version_num" tf:"optional"`
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAlias) {
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAlias) {
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `tfsdk:"aliases" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the registered model
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the registered model
	Name types.String `tfsdk:"name" tf:"optional"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelInfo) {
}

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringRead(existingState RegisteredModelInfo) {
}

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *RunRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunRefreshRequest) {
}

func (newState *RunRefreshRequest) SyncEffectiveFieldsDuringRead(existingState RunRefreshRequest) {
}

type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The type of the parent catalog.
	CatalogType types.String `tfsdk:"catalog_type" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of schema creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	EffectivePredictiveOptimizationFlag []EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
	// The unique identifier of the schema.
	SchemaId types.String `tfsdk:"schema_id" tf:"optional"`
	// Storage location for managed tables within schema.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified schema.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *SchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaInfo) {
}

func (newState *SchemaInfo) SyncEffectiveFieldsDuringRead(existingState SchemaInfo) {
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `tfsdk:"artifact_matchers" tf:""`
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetArtifactAllowlist) {
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringRead(existingState SetArtifactAllowlist) {
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"alias" tf:""`
	// Full name of the registered model
	FullName types.String `tfsdk:"full_name" tf:""`
	// The version number of the model version to which the alias points
	VersionNum types.Int64 `tfsdk:"version_num" tf:""`
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRegisteredModelAliasRequest) {
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringRead(existingState SetRegisteredModelAliasRequest) {
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm types.String `tfsdk:"algorithm" tf:"optional"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn types.String `tfsdk:"aws_kms_key_arn" tf:"optional"`
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan SseEncryptionDetails) {
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringRead(existingState SseEncryptionDetails) {
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole []AwsIamRoleResponse `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity []AzureManagedIdentityResponse `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal []AzureServicePrincipal `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken []CloudflareApiToken `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountResponse `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The unique identifier of the credential.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the credential.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage types.Bool `tfsdk:"used_for_managed_storage" tf:"optional"`
}

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageCredentialInfo) {
}

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState StorageCredentialInfo) {
}

type SystemSchemaInfo struct {
	// Name of the system schema.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SystemSchemaInfo) {
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringRead(existingState SystemSchemaInfo) {
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint []ForeignKeyConstraint `tfsdk:"foreign_key_constraint" tf:"optional,object"`

	NamedTableConstraint []NamedTableConstraint `tfsdk:"named_table_constraint" tf:"optional,object"`

	PrimaryKeyConstraint []PrimaryKeyConstraint `tfsdk:"primary_key_constraint" tf:"optional,object"`
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableConstraint) {
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringRead(existingState TableConstraint) {
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName types.String `tfsdk:"table_full_name" tf:""`
}

func (newState *TableDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableDependency) {
}

func (newState *TableDependency) SyncEffectiveFieldsDuringRead(existingState TableDependency) {
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists types.Bool `tfsdk:"table_exists" tf:"optional"`
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableExistsResponse) {
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringRead(existingState TableExistsResponse) {
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `tfsdk:"columns" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of table creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId types.String `tfsdk:"data_access_configuration_id" tf:"optional"`
	// Data source format
	DataSourceFormat types.String `tfsdk:"data_source_format" tf:"optional"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt types.Int64 `tfsdk:"deleted_at" tf:"optional"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs []DeltaRuntimePropertiesKvPairs `tfsdk:"delta_runtime_properties_kvpairs" tf:"optional,object"`

	EffectivePredictiveOptimizationFlag []EffectivePredictiveOptimizationFlag `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails []EncryptionDetails `tfsdk:"encryption_details" tf:"optional,object"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of table, relative to parent schema.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of table.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`

	RowFilter []TableRowFilter `tfsdk:"row_filter" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName types.String `tfsdk:"storage_credential_name" tf:"optional"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []TableConstraint `tfsdk:"table_constraints" tf:"optional"`
	// The unique identifier of the table.
	TableId types.String `tfsdk:"table_id" tf:"optional"`

	TableType types.String `tfsdk:"table_type" tf:"optional"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the table.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition types.String `tfsdk:"view_definition" tf:"optional"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies []DependencyList `tfsdk:"view_dependencies" tf:"optional,object"`
}

func (newState *TableInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableInfo) {
}

func (newState *TableInfo) SyncEffectiveFieldsDuringRead(existingState TableInfo) {
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:""`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []types.String `tfsdk:"input_column_names" tf:""`
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableRowFilter) {
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringRead(existingState TableRowFilter) {
}

type TableSummary struct {
	// The full name of the table.
	FullName types.String `tfsdk:"full_name" tf:"optional"`

	TableType types.String `tfsdk:"table_type" tf:"optional"`
}

func (newState *TableSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSummary) {
}

func (newState *TableSummary) SyncEffectiveFieldsDuringRead(existingState TableSummary) {
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress []PipelineProgress `tfsdk:"triggered_update_progress" tf:"optional,object"`
}

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggeredUpdateStatus) {
}

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringRead(existingState TriggeredUpdateStatus) {
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId types.String `tfsdk:"-"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UnassignRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignRequest) {
}

func (newState *UnassignRequest) SyncEffectiveFieldsDuringRead(existingState UnassignRequest) {
}

type UnassignResponse struct {
}

func (newState *UnassignResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignResponse) {
}

func (newState *UnassignResponse) SyncEffectiveFieldsDuringRead(existingState UnassignResponse) {
}

type UpdateAssignmentResponse struct {
}

func (newState *UpdateAssignmentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAssignmentResponse) {
}

func (newState *UpdateAssignmentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateAssignmentResponse) {
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// New name for the catalog.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
}

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCatalog) {
}

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringRead(existingState UpdateCatalog) {
}

type UpdateConnection struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
	// New name for the connection.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options map[string]types.String `tfsdk:"options" tf:""`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateConnection) {
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringRead(existingState UpdateConnection) {
}

type UpdateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails []EncryptionDetails `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force types.Bool `tfsdk:"force" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
	// New name for the external location.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExternalLocation) {
}

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringRead(existingState UpdateExternalLocation) {
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateFunction) {
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringRead(existingState UpdateFunction) {
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
	// New name for the metastore.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
}

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastore) {
}

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringRead(existingState UpdateMetastore) {
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastoreAssignment) {
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState UpdateMetastoreAssignment) {
}

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest) {
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest) {
}

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric `tfsdk:"custom_metrics" tf:"optional"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig []MonitorDataClassificationConfig `tfsdk:"data_classification_config" tf:"optional,object"`
	// Configuration for monitoring inference logs.
	InferenceLog []MonitorInferenceLog `tfsdk:"inference_log" tf:"optional,object"`
	// The notification settings for the monitor.
	Notifications []MonitorNotifications `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule []MonitorCronSchedule `tfsdk:"schedule" tf:"optional,object"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []types.String `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries []MonitorTimeSeries `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMonitor) {
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringRead(existingState UpdateMonitor) {
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `tfsdk:"changes" tf:"optional"`
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissions) {
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringRead(existingState UpdatePermissions) {
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// New name for the registered model.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegisteredModelRequest) {
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRegisteredModelRequest) {
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// New name for the schema.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties map[string]types.String `tfsdk:"properties" tf:"optional"`
}

func (newState *UpdateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSchema) {
}

func (newState *UpdateSchema) SyncEffectiveFieldsDuringRead(existingState UpdateSchema) {
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole []AwsIamRoleRequest `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity []AzureManagedIdentityResponse `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal []AzureServicePrincipal `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken []CloudflareApiToken `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"force" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
	// New name for the storage credential.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStorageCredential) {
}

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState UpdateStorageCredential) {
}

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`

	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateTableRequest) {
}

func (newState *UpdateTableRequest) SyncEffectiveFieldsDuringRead(existingState UpdateTableRequest) {
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
	// New name for the volume.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVolumeRequestContent) {
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState UpdateVolumeRequestContent) {
}

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	AssignWorkspaces []types.Int64 `tfsdk:"assign_workspaces" tf:"optional"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces []types.Int64 `tfsdk:"unassign_workspaces" tf:"optional"`
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindings) {
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindings) {
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	Add []WorkspaceBinding `tfsdk:"add" tf:"optional"`
	// List of workspace bindings
	Remove []WorkspaceBinding `tfsdk:"remove" tf:"optional"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindingsParameters) {
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindingsParameters) {
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole []AwsIamRoleRequest `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity []AzureManagedIdentityRequest `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal []AzureServicePrincipal `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken []CloudflareApiToken `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The name of an existing external location to validate.
	ExternalLocationName types.String `tfsdk:"external_location_name" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// The name of the storage credential to validate.
	StorageCredentialName types.String `tfsdk:"storage_credential_name" tf:"optional"`
	// The external location url to validate.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredential) {
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredential) {
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir types.Bool `tfsdk:"isDir" tf:"optional"`
	// The results of the validation check.
	Results []ValidationResult `tfsdk:"results" tf:"optional"`
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredentialResponse) {
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredentialResponse) {
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The operation tested.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result" tf:"optional"`
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidationResult) {
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringRead(existingState ValidationResult) {
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the volume
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails []EncryptionDetails `tfsdk:"encryption_details" tf:"optional,object"`
	// The three-level (fully qualified) name of the volume
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the volume
	Name types.String `tfsdk:"name" tf:"optional"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the volume last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// The unique identifier of the volume
	VolumeId types.String `tfsdk:"volume_id" tf:"optional"`

	VolumeType types.String `tfsdk:"volume_type" tf:"optional"`
}

func (newState *VolumeInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumeInfo) {
}

func (newState *VolumeInfo) SyncEffectiveFieldsDuringRead(existingState VolumeInfo) {
}

type WorkspaceBinding struct {
	BindingType types.String `tfsdk:"binding_type" tf:"optional"`

	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBinding) {
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringRead(existingState WorkspaceBinding) {
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings []WorkspaceBinding `tfsdk:"bindings" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *WorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBindingsResponse) {
}

func (newState *WorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState WorkspaceBindingsResponse) {
}
